// Dashboard lokal untuk mengelola konten blog Hugo.
// Berjalan di localhost, menyunting berkas Markdown di folder content/ secara langsung.
// Hanya memakai pustaka standar Go (tanpa dependensi eksternal).
package main

import (
	"embed"
	"encoding/json"
	"flag"
	"fmt"
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

var contentDir string

func main() {
	addr := flag.String("addr", "127.0.0.1:1414", "alamat server dashboard")
	flag.StringVar(&contentDir, "content", "content", "folder konten Hugo")
	flag.Parse()

	abs, err := filepath.Abs(contentDir)
	if err != nil {
		log.Fatal(err)
	}
	if fi, err := os.Stat(abs); err != nil || !fi.IsDir() {
		log.Fatalf("Folder konten tidak ditemukan: %s\nJalankan dashboard dari root proyek Hugo.", abs)
	}
	contentDir = abs

	mux := http.NewServeMux()
	mux.HandleFunc("/", serveUI)
	mux.HandleFunc("/api/list", handleList)
	mux.HandleFunc("/api/read", handleRead)
	mux.HandleFunc("/api/save", handleSave)
	mux.HandleFunc("/api/create", handleCreate)
	mux.HandleFunc("/api/delete", handleDelete)

	fmt.Printf("\n  ✒  Dashboard Blog berjalan di  http://%s\n", *addr)
	fmt.Printf("     Konten: %s\n\n", contentDir)
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
