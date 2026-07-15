// Cloudflare Worker: jembatan OAuth GitHub untuk Decap CMS (panel /admin).
//
// Decap membuka popup ke  {base_url}/auth  → Worker mengarahkan ke GitHub →
// GitHub kembali ke  {base_url}/callback  → Worker menukar code menjadi token →
// mengirim token ke jendela /admin lewat postMessage.
//
// Rahasia diatur sebagai environment variable Worker:
//   GITHUB_CLIENT_ID, GITHUB_CLIENT_SECRET
//
// Ganti nilai ALLOWED_ORIGIN dengan domain situs Anda (untuk keamanan postMessage).

const ALLOWED_ORIGIN = "https://ramadhan.me";

export default {
  async fetch(request, env) {
    const url = new URL(request.url);

    // 1) Mulai login: arahkan ke halaman otorisasi GitHub.
    if (url.pathname === "/auth") {
      const params = new URLSearchParams({
        client_id: env.GITHUB_CLIENT_ID,
        redirect_uri: `${url.origin}/callback`,
        scope: "repo,user",
        state: crypto.randomUUID(),
      });
      return Response.redirect(
        `https://github.com/login/oauth/authorize?${params}`,
        302
      );
    }

    // 2) Callback: tukar 'code' menjadi access token, lalu kirim ke /admin.
    if (url.pathname === "/callback") {
      const code = url.searchParams.get("code");
      if (!code) return new Response("Kode otorisasi tidak ada.", { status: 400 });

      const tokenRes = await fetch(
        "https://github.com/login/oauth/access_token",
        {
          method: "POST",
          headers: { "Content-Type": "application/json", Accept: "application/json" },
          body: JSON.stringify({
            client_id: env.GITHUB_CLIENT_ID,
            client_secret: env.GITHUB_CLIENT_SECRET,
            code,
          }),
        }
      );
      const data = await tokenRes.json();

      const status = data.access_token ? "success" : "error";
      const content = data.access_token
        ? { token: data.access_token, provider: "github" }
        : { error: data.error_description || data.error || "Gagal menukar token" };

      return new Response(renderPage(status, content), {
        headers: { "content-type": "text/html; charset=utf-8" },
      });
    }

    return new Response("Cloudflare CMS OAuth — jalur tidak ditemukan.", { status: 404 });
  },
};

// Halaman yang mengirim hasil ke jendela pembuka (Decap) via postMessage.
function renderPage(status, content) {
  const payload = `authorization:github:${status}:${JSON.stringify(content)}`;
  const target = JSON.stringify(ALLOWED_ORIGIN);
  return `<!doctype html>
<html><head><meta charset="utf-8"><title>Otorisasi…</title></head>
<body>
<p>Menyelesaikan login…</p>
<script>
  (function () {
    var payload = ${JSON.stringify(payload)};
    function send(origin) {
      if (window.opener) window.opener.postMessage(payload, origin);
    }
    function onMsg(e) {
      // Balas ke asal yang mengirim 'authorizing:github'.
      send(e.origin || ${target});
      window.removeEventListener("message", onMsg, false);
      window.close();
    }
    window.addEventListener("message", onMsg, false);
    if (window.opener) window.opener.postMessage("authorizing:github", ${target});
  })();
</script>
</body></html>`;
}
