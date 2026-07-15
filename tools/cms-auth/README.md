# Panel /admin online di Cloudflare (Decap CMS + GitHub OAuth)

Karena Cloudflare Pages tidak punya layanan login bawaan seperti Netlify Identity,
panel `/admin` memakai **GitHub OAuth** yang dijembatani oleh Worker kecil di folder ini.

## Langkah

### 1. Buat GitHub OAuth App
1. Buka <https://github.com/settings/developers> → **OAuth Apps** → **New OAuth App**.
2. Isi:
   - **Application name:** bebas (mis. "Blog CMS")
   - **Homepage URL:** `https://ramadhan.me`
   - **Authorization callback URL:** `https://cms-auth.SUBDOMAIN.workers.dev/callback`
     *(sementara isi tebakan; perbaiki setelah tahu URL Worker di langkah 2)*
3. **Register** → catat **Client ID**, lalu **Generate a new client secret** → catat **Client Secret**.

### 2. Deploy Worker
**Cara mudah (dashboard):**
1. <https://dash.cloudflare.com> → **Workers & Pages** → **Create** → **Create Worker**.
2. Beri nama `cms-auth` → **Deploy** → **Edit code** → tempel isi `worker.js` → **Deploy**.
3. Salin URL Worker (mis. `https://cms-auth.namaakun.workers.dev`).
4. **Settings → Variables and Secrets** → tambah dua secret:
   - `GITHUB_CLIENT_ID` = Client ID dari langkah 1
   - `GITHUB_CLIENT_SECRET` = Client Secret dari langkah 1

**Atau via CLI (wrangler):**
```bash
cd tools/cms-auth
npx wrangler deploy
npx wrangler secret put GITHUB_CLIENT_ID
npx wrangler secret put GITHUB_CLIENT_SECRET
```

### 3. Sinkronkan nilai
- Di **worker.js**, pastikan `ALLOWED_ORIGIN` = `https://ramadhan.me`.
- Di GitHub OAuth App (langkah 1), pastikan **callback URL** memakai URL Worker asli:
  `https://<url-worker-anda>/callback`.
- Di `static/admin/config.yml`, ganti:
  - `repo: USERNAME/blog` → `user/repo` GitHub Anda
  - `base_url: https://cms-auth.SUBDOMAIN.workers.dev` → URL Worker Anda

### 4. Pakai
Commit & push perubahan `config.yml`, tunggu Cloudflare Pages rebuild, lalu buka
`https://ramadhan.me/admin/` → **Login with GitHub**. Setiap simpan di CMS akan
membuat commit ke repo, dan situs otomatis rebuild.

> Uji lokal (opsional): `npx decap-server` lalu buka `/admin/` saat `hugo server` jalan
> (memakai `local_backend: true`).
