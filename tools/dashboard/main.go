// Dashboard lokal untuk mengelola konten blog Hugo.
// Berjalan di localhost, menyunting berkas Markdown di folder content/ secara langsung.
// Hanya memakai pustaka standar Go (tanpa dependensi eksternal).
package main

import (
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

//go:embed ui.html
var uiFS embed.FS

var (
	contentDir string
	imagesDir  string // static/images — tujuan unggahan gambar
	configPath string // hugo.toml — sumber identitas situs & teks beranda
)

func main() {
	addr := flag.String("addr", "127.0.0.1:1414", "alamat server dashboard")
	staticDir := flag.String("static", "static", "folder static Hugo (tujuan unggahan)")
	flag.StringVar(&contentDir, "content", "content", "folder konten Hugo")
	flag.StringVar(&configPath, "config", "hugo.toml", "berkas konfigurasi Hugo")
	flag.Parse()

	abs, err := filepath.Abs(contentDir)
	if err != nil {
		log.Fatal(err)
	}
	if fi, err := os.Stat(abs); err != nil || !fi.IsDir() {
		log.Fatalf("Folder konten tidak ditemukan: %s\nJalankan dashboard dari root proyek Hugo.", abs)
	}
	contentDir = abs

	cAbs, err := filepath.Abs(configPath)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := os.Stat(cAbs); err != nil {
		log.Fatalf("Konfigurasi Hugo tidak ditemukan: %s\nJalankan dashboard dari root proyek Hugo.", cAbs)
	}
	configPath = cAbs

	sAbs, err := filepath.Abs(*staticDir)
	if err != nil {
		log.Fatal(err)
	}
	imagesDir = filepath.Join(sAbs, "images")
	if err := os.MkdirAll(imagesDir, 0755); err != nil {
		log.Fatalf("Tidak bisa menyiapkan folder gambar %s: %v", imagesDir, err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", serveUI)
	mux.HandleFunc("/api/list", handleList)
	mux.HandleFunc("/api/read", handleRead)
	mux.HandleFunc("/api/save", handleSave)
	mux.HandleFunc("/api/create", handleCreate)
	mux.HandleFunc("/api/delete", handleDelete)
	mux.HandleFunc("/api/upload", handleUpload)
	mux.HandleFunc("/api/media", handleMedia)
	mux.HandleFunc("/api/site", handleSite)
	// Pratinjau gambar untuk pustaka media di dashboard.
	mux.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir(imagesDir))))

	fmt.Printf("\n  ▪  Dashboard Blog berjalan di  http://%s\n", *addr)
	fmt.Printf("     Konten: %s\n", contentDir)
	fmt.Printf("     Gambar: %s\n\n", imagesDir)
	if err := http.ListenAndServe(*addr, mux); err != nil {
		log.Fatal(err)
	}
}

// ---------- Model ----------

type Item struct {
	Path    string `json:"path"`
	Section string `json:"section"`
	Title   string `json:"title"`
	Date    string `json:"date"`
	Draft   bool   `json:"draft"`
}

type Doc struct {
	Path        string   `json:"path"`
	Section     string   `json:"section"`
	Title       string   `json:"title"`
	Date        string   `json:"date"`
	Draft       bool     `json:"draft"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Categories  []string `json:"categories"`
	Extra       string   `json:"extra"`
	Body        string   `json:"body"`
}

// ---------- Handler ----------

func serveUI(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	b, _ := uiFS.ReadFile("ui.html")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(b)
}

func handleList(w http.ResponseWriter, r *http.Request) {
	var items []Item
	filepath.Walk(contentDir, func(p string, fi os.FileInfo, err error) error {
		if err != nil || fi.IsDir() || !strings.HasSuffix(p, ".md") {
			return nil
		}
		rel, _ := filepath.Rel(filepath.Dir(contentDir), p)
		rel = filepath.ToSlash(rel)
		base := filepath.Base(p)
		if base == "_index.md" {
			return nil // lewati halaman bagian
		}
		raw, _ := os.ReadFile(p)
		fm, _ := splitFrontMatter(string(raw))
		title := fmString(fm, "title")
		if title == "" {
			title = strings.TrimSuffix(base, ".md")
		}
		items = append(items, Item{
			Path:    rel,
			Section: sectionOf(rel),
			Title:   title,
			Date:    fmRaw(fm, "date"),
			Draft:   fmBool(fm, "draft"),
		})
		return nil
	})
	sort.Slice(items, func(i, j int) bool {
		if items[i].Section != items[j].Section {
			return items[i].Section < items[j].Section
		}
		return items[i].Date > items[j].Date
	})
	writeJSON(w, items)
}

func handleRead(w http.ResponseWriter, r *http.Request) {
	p, ok := safePath(r.URL.Query().Get("path"))
	if !ok {
		http.Error(w, "path tidak valid", http.StatusBadRequest)
		return
	}
	raw, err := os.ReadFile(p)
	if err != nil {
		http.Error(w, "tidak ditemukan", http.StatusNotFound)
		return
	}
	fm, body := splitFrontMatter(string(raw))
	rel, _ := filepath.Rel(filepath.Dir(contentDir), p)
	rel = filepath.ToSlash(rel)
	doc := Doc{
		Path:        rel,
		Section:     sectionOf(rel),
		Title:       fmString(fm, "title"),
		Date:        fmRaw(fm, "date"),
		Draft:       fmBool(fm, "draft"),
		Description: fmString(fm, "description"),
		Tags:        fmArray(fm, "tags"),
		Categories:  fmArray(fm, "categories"),
		Extra:       extraLines(fm),
		Body:        strings.TrimLeft(body, "\n"),
	}
	writeJSON(w, doc)
}

func handleSave(w http.ResponseWriter, r *http.Request) {
	var doc Doc
	if err := json.NewDecoder(r.Body).Decode(&doc); err != nil {
		http.Error(w, "json tidak valid", http.StatusBadRequest)
		return
	}
	p, ok := safePath(doc.Path)
	if !ok {
		http.Error(w, "path tidak valid", http.StatusBadRequest)
		return
	}
	if err := os.WriteFile(p, []byte(buildFile(doc)), 0644); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]string{"status": "ok", "path": doc.Path})
}

func handleCreate(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Section string `json:"section"`
		Title   string `json:"title"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "json tidak valid", http.StatusBadRequest)
		return
	}
	req.Section = slugify(req.Section)
	if req.Section == "" || strings.TrimSpace(req.Title) == "" {
		http.Error(w, "bagian dan judul wajib diisi", http.StatusBadRequest)
		return
	}
	slug := slugify(req.Title)
	if slug == "" {
		slug = "tanpa-judul"
	}
	rel := "content/" + req.Section + "/" + slug + ".md"
	p, ok := safePath(rel)
	if !ok {
		http.Error(w, "path tidak valid", http.StatusBadRequest)
		return
	}
	// Hindari menimpa berkas yang sudah ada.
	if _, err := os.Stat(p); err == nil {
		rel = "content/" + req.Section + "/" + slug + "-" + time.Now().Format("150405") + ".md"
		p, _ = safePath(rel)
	}
	if err := os.MkdirAll(filepath.Dir(p), 0755); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	doc := Doc{
		Path:    rel,
		Section: req.Section,
		Title:   req.Title,
		Date:    time.Now().Format("2006-01-02T15:04:05-07:00"),
		Draft:   true,
	}
	if err := os.WriteFile(p, []byte(buildFile(doc)), 0644); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]string{"status": "ok", "path": rel})
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Path string `json:"path"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "json tidak valid", http.StatusBadRequest)
		return
	}
	p, ok := safePath(req.Path)
	if !ok {
		http.Error(w, "path tidak valid", http.StatusBadRequest)
		return
	}
	if err := os.Remove(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]string{"status": "ok"})
}

// ---------- Identitas situs (hugo.toml) ----------

// Kolom yang boleh disunting dari dashboard. Sengaja dibatasi pada teks yang
// tampil ke pembaca — baseURL, menu, dan setelan markup tetap lewat berkas,
// karena salah ketik di sana bisa menggagalkan build.
type SiteField struct {
	Key     string `json:"key"`
	Section string `json:"-"` // "" = tingkat atas, "params" = [params]
	Label   string `json:"label"`
	Kind    string `json:"kind"` // text | area
	Hint    string `json:"hint"`
	Value   string `json:"value"`
}

var siteFields = []SiteField{
	{Key: "title", Section: "", Label: "Judul Situs", Kind: "text",
		Hint: "Tampil di tab peramban dan hasil pencarian."},
	{Key: "author", Section: "params", Label: "Nama Penulis", Kind: "text",
		Hint: "Dipakai di header, byline tulisan, dan monogram."},
	{Key: "role", Section: "params", Label: "Peran", Kind: "text",
		Hint: "Pil kecil di atas judul beranda. Kosongkan untuk menyembunyikan."},
	{Key: "tagline", Section: "params", Label: "Tagline", Kind: "text",
		Hint: "Judul besar di beranda."},
	{Key: "intro", Section: "params", Label: "Intro Beranda", Kind: "area",
		Hint: "Paragraf pengantar di bawah judul beranda."},
	{Key: "description", Section: "params", Label: "Deskripsi Situs", Kind: "area",
		Hint: "Ringkasan untuk mesin pencari & pratinjau tautan."},
	{Key: "email", Section: "params", Label: "Email", Kind: "text", Hint: ""},
	{Key: "github", Section: "params", Label: "GitHub", Kind: "text", Hint: ""},
	{Key: "linkedin", Section: "params", Label: "LinkedIn", Kind: "text",
		Hint: "Kosongkan bila tak dipakai — tautannya otomatis hilang."},
	{Key: "twitter", Section: "params", Label: "Twitter", Kind: "text",
		Hint: "Kosongkan bila tak dipakai."},
}

// tomlSectionOf mengenali baris kepala seksi: [params] atau [[menu.main]].
func tomlSectionOf(t string) (string, bool) {
	if strings.HasPrefix(t, "[[") && strings.HasSuffix(t, "]]") {
		return strings.TrimSpace(t[2 : len(t)-2]), true
	}
	if strings.HasPrefix(t, "[") && strings.HasSuffix(t, "]") {
		return strings.TrimSpace(t[1 : len(t)-1]), true
	}
	return "", false
}

// splitInlineComment memisahkan nilai dari komentar di ujung baris, dengan
// mengabaikan '#' yang berada di dalam tanda kutip.
func splitInlineComment(v string) (val, comment string) {
	inStr := false
	for i := 0; i < len(v); i++ {
		switch v[i] {
		case '"':
			bs := 0
			for j := i - 1; j >= 0 && v[j] == '\\'; j-- {
				bs++
			}
			if bs%2 == 0 {
				inStr = !inStr
			}
		case '#':
			if !inStr {
				return strings.TrimRight(v[:i], " \t"), v[i:]
			}
		}
	}
	return strings.TrimRight(v, " \t"), ""
}

// tomlQuote menyandikan string sebagai basic string TOML satu baris.
// Baris baru sengaja di-escape, bukan ditulis apa adanya, agar berkas tetap sah.
func tomlQuote(s string) string {
	r := strings.NewReplacer(
		`\`, `\\`, `"`, `\"`,
		"\n", `\n`, "\r", `\r`, "\t", `\t`,
	)
	return `"` + r.Replace(s) + `"`
}

func indentOf(line string) string {
	return line[:len(line)-len(strings.TrimLeft(line, " \t"))]
}

func readSiteConfig() ([]SiteField, error) {
	b, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	found := map[string]string{}
	cur := ""
	for _, line := range strings.Split(strings.ReplaceAll(string(b), "\r\n", "\n"), "\n") {
		t := strings.TrimSpace(line)
		if t == "" || strings.HasPrefix(t, "#") {
			continue
		}
		if s, ok := tomlSectionOf(t); ok {
			cur = s
			continue
		}
		eq := strings.Index(t, "=")
		if eq < 0 {
			continue
		}
		key := strings.TrimSpace(t[:eq])
		val, _ := splitInlineComment(strings.TrimSpace(t[eq+1:]))
		for _, f := range siteFields {
			if f.Section == cur && f.Key == key {
				found[key] = unquoteToml(val)
			}
		}
	}
	out := make([]SiteField, 0, len(siteFields))
	for _, f := range siteFields {
		f.Value = found[f.Key]
		out = append(out, f)
	}
	return out, nil
}

// unquoteToml membalik escape basic string TOML.
func unquoteToml(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
		r := strings.NewReplacer(
			`\n`, "\n", `\r`, "\r", `\t`, "\t",
			`\"`, `"`, `\\`, `\`,
		)
		s = r.Replace(s)
	}
	return s
}

func writeSiteConfig(vals map[string]string) error {
	b, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}
	lines := strings.Split(strings.ReplaceAll(string(b), "\r\n", "\n"), "\n")

	byKey := map[string]SiteField{}
	for _, f := range siteFields {
		byKey[f.Key] = f
	}

	written := map[string]bool{}
	// Baris kepala tiap seksi. Kunci baru disisipkan tepat di bawahnya — bukan di
	// ujung seksi — supaya tidak menempel pada komentar milik kunci lain
	// (mis. mendarat di bawah "# Copyright" dan terbaca seolah berkaitan).
	headerLineOf := map[string]int{"": -1} // -1 = tingkat atas, sisipkan di baris pertama
	seen := map[string]bool{"": true}
	cur := ""

	for i, line := range lines {
		t := strings.TrimSpace(line)
		if t == "" || strings.HasPrefix(t, "#") {
			continue
		}
		if s, ok := tomlSectionOf(t); ok {
			cur = s
			if !seen[cur] {
				headerLineOf[cur] = i
				seen[cur] = true
			}
			continue
		}
		eq := strings.Index(t, "=")
		if eq < 0 {
			continue
		}
		key := strings.TrimSpace(t[:eq])
		f, ok := byKey[key]
		if !ok || f.Section != cur {
			continue
		}
		v, ok := vals[key]
		if !ok {
			continue
		}
		_, comment := splitInlineComment(strings.TrimSpace(t[eq+1:]))
		nl := indentOf(line) + key + " = " + tomlQuote(v)
		if comment != "" {
			nl += "  " + comment
		}
		lines[i] = nl
		written[key] = true
	}

	// Kunci yang belum ada di berkas disisipkan di ujung seksinya.
	// Dikerjakan dari indeks terbesar agar penyisipan tidak menggeser sisanya.
	type ins struct {
		at   int
		text string
	}
	var pending []ins
	for _, f := range siteFields {
		v, ok := vals[f.Key]
		if !ok || written[f.Key] {
			continue
		}
		at, ok := headerLineOf[f.Section]
		if !ok {
			return fmt.Errorf("seksi [%s] tidak ada di %s — tambahkan kunci %s secara manual",
				f.Section, filepath.Base(configPath), f.Key)
		}
		indent := "  "
		if f.Section == "" {
			indent = ""
		}
		pending = append(pending, ins{at: at, text: indent + f.Key + " = " + tomlQuote(v)})
	}
	sort.Slice(pending, func(i, j int) bool { return pending[i].at > pending[j].at })
	for _, p := range pending {
		lines = append(lines[:p.at+1], append([]string{p.text}, lines[p.at+1:]...)...)
	}

	out := strings.Join(lines, "\n")
	// Tulis lewat berkas sementara lalu rename: kalau gagal di tengah jalan,
	// hugo.toml lama tetap utuh dan situs tidak ikut rusak.
	tmp := configPath + ".tmp"
	if err := os.WriteFile(tmp, []byte(out), 0644); err != nil {
		return err
	}
	return os.Rename(tmp, configPath)
}

func handleSite(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var req map[string]string
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "json tidak valid", http.StatusBadRequest)
			return
		}
		// Hanya kunci yang dikenal yang diteruskan.
		vals := map[string]string{}
		for _, f := range siteFields {
			if v, ok := req[f.Key]; ok {
				vals[f.Key] = strings.TrimSpace(v)
			}
		}
		if len(vals) == 0 {
			http.Error(w, "tidak ada kolom yang dikenali", http.StatusBadRequest)
			return
		}
		if err := writeSiteConfig(vals); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeJSON(w, map[string]string{"status": "ok"})
		return
	}
	fields, err := readSiteConfig()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, fields)
}

// ---------- Media ----------

// Ekstensi yang boleh disimpan. Peramban sudah mengubah foto raster menjadi
// WebP sebelum dikirim; GIF & SVG dilewatkan apa adanya karena tidak cocok
// dikonversi (animasi hilang / vektor jadi raster).
var allowedExt = map[string]bool{
	".webp": true, ".jpg": true, ".jpeg": true, ".png": true,
	".gif": true, ".svg": true,
}

type MediaItem struct {
	Name string `json:"name"`
	URL  string `json:"url"`  // URL publik di situs, mis. /images/foo.webp
	Src  string `json:"src"`  // URL pratinjau di dashboard
	Size int64  `json:"size"` // byte
	Mod  string `json:"mod"`
}

const maxUpload = 25 << 20 // 25 MB — sudah sangat longgar untuk gambar terkompresi

func handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "metode tidak didukung", http.StatusMethodNotAllowed)
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, maxUpload)
	if err := r.ParseMultipartForm(8 << 20); err != nil {
		http.Error(w, "berkas terlalu besar atau form tidak valid", http.StatusBadRequest)
		return
	}
	file, hdr, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "berkas tidak ditemukan", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Nama diambil dari form bila ada (hasil slugify di peramban), jika tidak
	// pakai nama asli berkas.
	name := strings.TrimSpace(r.FormValue("name"))
	if name == "" {
		name = hdr.Filename
	}
	ext := strings.ToLower(filepath.Ext(name))
	if ext == "" {
		ext = strings.ToLower(filepath.Ext(hdr.Filename))
	}
	if !allowedExt[ext] {
		http.Error(w, "jenis berkas tidak didukung: "+ext, http.StatusBadRequest)
		return
	}
	base := slugify(strings.TrimSuffix(name, filepath.Ext(name)))
	if base == "" {
		base = "gambar"
	}

	// Cari nama yang belum terpakai agar unggahan lama tidak tertimpa.
	target := filepath.Join(imagesDir, base+ext)
	for i := 2; ; i++ {
		if _, err := os.Stat(target); os.IsNotExist(err) {
			break
		}
		target = filepath.Join(imagesDir, fmt.Sprintf("%s-%d%s", base, i, ext))
	}

	out, err := os.Create(target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer out.Close()
	if _, err := io.Copy(out, file); err != nil {
		os.Remove(target)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fname := filepath.Base(target)
	writeJSON(w, map[string]string{
		"status": "ok",
		"name":   fname,
		"url":    "/images/" + fname,
		"src":    "/media/" + fname,
	})
}

func handleMedia(w http.ResponseWriter, r *http.Request) {
	entries, err := os.ReadDir(imagesDir)
	if err != nil {
		writeJSON(w, []MediaItem{})
		return
	}
	items := []MediaItem{}
	for _, e := range entries {
		if e.IsDir() || !allowedExt[strings.ToLower(filepath.Ext(e.Name()))] {
			continue
		}
		fi, err := e.Info()
		if err != nil {
			continue
		}
		items = append(items, MediaItem{
			Name: e.Name(),
			URL:  "/images/" + e.Name(),
			Src:  "/media/" + e.Name(),
			Size: fi.Size(),
			Mod:  fi.ModTime().Format(time.RFC3339),
		})
	}
	// Terbaru dulu.
	sort.Slice(items, func(i, j int) bool { return items[i].Mod > items[j].Mod })
	writeJSON(w, items)
}

// ---------- Front matter (subset TOML datar) ----------

// splitFrontMatter memisahkan blok +++ ... +++ (atau --- ... ---) dari body.
// Mengembalikan peta urut berupa baris front matter mentah dan sisa body.
func splitFrontMatter(raw string) (map[string]string, string) {
	fm := map[string]string{}
	raw = strings.ReplaceAll(raw, "\r\n", "\n")
	var delim string
	if strings.HasPrefix(raw, "+++\n") {
		delim = "+++"
	} else if strings.HasPrefix(raw, "---\n") {
		delim = "---"
	} else {
		return fm, raw
	}
	rest := raw[len(delim)+1:]
	end := strings.Index(rest, "\n"+delim)
	if end < 0 {
		return fm, raw
	}
	block := rest[:end]
	body := rest[end+len(delim)+1:]
	for _, line := range strings.Split(block, "\n") {
		t := strings.TrimSpace(line)
		if t == "" || strings.HasPrefix(t, "#") {
			continue
		}
		eq := strings.Index(t, "=")
		if strings.HasPrefix(t, "  ") { // YAML gaya "key: value"
			eq = -1
		}
		if eq < 0 {
			if c := strings.Index(t, ":"); c > 0 && delim == "---" {
				key := strings.TrimSpace(t[:c])
				fm[key] = strings.TrimSpace(t[c+1:])
			}
			continue
		}
		key := strings.TrimSpace(t[:eq])
		val := strings.TrimSpace(t[eq+1:])
		fm[key] = val
	}
	// Simpan blok mentah untuk merekonstruksi "extra" secara verbatim.
	fm["__block__"] = block
	return fm, body
}

var knownKeys = map[string]bool{
	"title": true, "date": true, "draft": true,
	"description": true, "tags": true, "categories": true,
}

// extraLines mengembalikan baris front matter selain kunci yang sudah dikenal,
// termasuk komentar, dipertahankan apa adanya agar tidak hilang saat disunting.
func extraLines(fm map[string]string) string {
	var out []string
	for _, line := range strings.Split(fm["__block__"], "\n") {
		t := strings.TrimSpace(line)
		if t == "" {
			continue
		}
		if strings.HasPrefix(t, "#") { // komentar: pertahankan
			out = append(out, strings.TrimRight(line, " \t"))
			continue
		}
		if eq := strings.Index(t, "="); eq > 0 {
			key := strings.TrimSpace(t[:eq])
			if knownKeys[key] {
				continue
			}
		}
		out = append(out, strings.TrimRight(line, " \t"))
	}
	return strings.Join(out, "\n")
}

func fmRaw(fm map[string]string, key string) string { return fm[key] }

func fmString(fm map[string]string, key string) string {
	return unquote(fm[key])
}

func fmBool(fm map[string]string, key string) bool {
	return strings.TrimSpace(fm[key]) == "true"
}

func fmArray(fm map[string]string, key string) []string {
	v := strings.TrimSpace(fm[key])
	if !strings.HasPrefix(v, "[") {
		return nil
	}
	v = strings.TrimPrefix(v, "[")
	v = strings.TrimSuffix(v, "]")
	var out []string
	for _, part := range splitTopLevel(v) {
		part = unquote(strings.TrimSpace(part))
		if part != "" {
			out = append(out, part)
		}
	}
	return out
}

// splitTopLevel memecah daftar dengan koma, mengabaikan koma di dalam tanda kutip.
func splitTopLevel(s string) []string {
	var out []string
	var b strings.Builder
	inStr := false
	for _, r := range s {
		switch {
		case r == '"':
			inStr = !inStr
			b.WriteRune(r)
		case r == ',' && !inStr:
			out = append(out, b.String())
			b.Reset()
		default:
			b.WriteRune(r)
		}
	}
	if strings.TrimSpace(b.String()) != "" {
		out = append(out, b.String())
	}
	return out
}

func unquote(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
		s = strings.ReplaceAll(s, `\"`, `"`)
		s = strings.ReplaceAll(s, `\\`, `\`)
	}
	return s
}

func quote(s string) string {
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, `"`, `\"`)
	return `"` + s + `"`
}

// buildFile menyusun ulang berkas Markdown dari Doc.
func buildFile(d Doc) string {
	var b strings.Builder
	b.WriteString("+++\n")
	b.WriteString("title = " + quote(d.Title) + "\n")
	date := strings.TrimSpace(d.Date)
	if date == "" {
		date = time.Now().Format("2006-01-02T15:04:05-07:00")
	}
	// Tanggal ditulis sebagai literal TOML (tanpa kutip) bila berformat tanggal.
	if looksLikeDate(date) {
		b.WriteString("date = " + date + "\n")
	} else {
		b.WriteString("date = " + quote(date) + "\n")
	}
	b.WriteString(fmt.Sprintf("draft = %t\n", d.Draft))
	if strings.TrimSpace(d.Description) != "" {
		b.WriteString("description = " + quote(d.Description) + "\n")
	}
	if len(d.Tags) > 0 {
		b.WriteString("tags = " + tomlArray(d.Tags) + "\n")
	}
	if len(d.Categories) > 0 {
		b.WriteString("categories = " + tomlArray(d.Categories) + "\n")
	}
	if strings.TrimSpace(d.Extra) != "" {
		b.WriteString(strings.TrimRight(d.Extra, "\n") + "\n")
	}
	b.WriteString("+++\n\n")
	b.WriteString(strings.TrimLeft(d.Body, "\n"))
	if !strings.HasSuffix(b.String(), "\n") {
		b.WriteString("\n")
	}
	return b.String()
}

func tomlArray(items []string) string {
	var q []string
	for _, it := range items {
		it = strings.TrimSpace(it)
		if it != "" {
			q = append(q, quote(it))
		}
	}
	return "[" + strings.Join(q, ", ") + "]"
}

var dateRe = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}`)

func looksLikeDate(s string) bool { return dateRe.MatchString(s) }

// ---------- Utilitas ----------

func sectionOf(rel string) string {
	parts := strings.Split(strings.TrimPrefix(rel, "content/"), "/")
	if len(parts) > 1 {
		return parts[0]
	}
	return "halaman"
}

var slugRe = regexp.MustCompile(`[^a-z0-9]+`)

func slugify(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	s = slugRe.ReplaceAllString(s, "-")
	return strings.Trim(s, "-")
}

// safePath memastikan path berada di dalam contentDir dan berakhiran .md.
func safePath(rel string) (string, bool) {
	rel = filepath.ToSlash(strings.TrimSpace(rel))
	if rel == "" || !strings.HasSuffix(rel, ".md") {
		return "", false
	}
	rel = strings.TrimPrefix(rel, "content/")
	clean := filepath.Clean(filepath.Join(contentDir, rel))
	if clean != contentDir && !strings.HasPrefix(clean, contentDir+string(os.PathSeparator)) {
		return "", false
	}
	return clean, true
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(v)
}
