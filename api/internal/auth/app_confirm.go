package auth

import (
	"html/template"
	"net/http"

	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// appConfirmCopy is the confirm page's words, shared with its test.
var appConfirmCopy = struct {
	Title, Heading, Body, Button, NotYou string
}{
	Title:   "Sign in to the Forever Sixty app",
	Heading: "Sign in to the Forever Sixty app on your phone",
	Body:    "Continue to Battle.net to sign in. The Forever Sixty app on your phone finishes the sign-in.",
	Button:  "Continue to Battle.net",
	NotYou:  "If you didn't start a sign-in from the Forever Sixty app on your phone, close this page.",
}

// appConfirmPage is self-contained: no script, no form, no image, no font, nothing fetched.
// The one link is the click the flow needs.
var appConfirmPage = template.Must(template.New("confirm").Parse(`<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="robots" content="noindex">
<title>{{.Copy.Title}}</title>
<style>
:root { color-scheme: dark; }
body { margin: 0; min-height: 100vh; display: grid; place-items: center; background: #0f0d0a;
  color: #ece4d4; font: 16px/1.5 system-ui, -apple-system, "Segoe UI", sans-serif; }
main { max-width: 420px; margin: 0 16px; padding: 28px 24px; border: 1px solid #3a3226;
  border-radius: 12px; background: #1a1712; }
h1 { margin: 0 0 12px; font-size: 22px; line-height: 1.25; }
p { margin: 0 0 20px; color: #b9ae99; }
a.button { display: block; padding: 14px 16px; border-radius: 8px; background: #d9a441;
  color: #1a1712; font-weight: 700; text-align: center; text-decoration: none; }
p.note { margin: 20px 0 0; font-size: 14px; }
</style>
</head>
<body>
<main>
<h1>{{.Copy.Heading}}</h1>
<p>{{.Copy.Body}}</p>
<a class="button" href="{{.Continue}}">{{.Copy.Button}}</a>
<p class="note">{{.Copy.NotYou}}</p>
</main>
</body>
</html>
`))

// appConfirmCSP lets the page style itself and nothing else: no script, no fetch, no form,
// and no frame around it, so the click cannot be lifted by an overlay.
const appConfirmCSP = "default-src 'none'; style-src 'unsafe-inline'; base-uri 'none'; form-action 'none'; frame-ancestors 'none'"

func (s *Service) writeAppConfirm(w http.ResponseWriter, r *http.Request, continuePath string) {
	h := w.Header()
	h.Set("Content-Type", "text/html; charset=utf-8")
	h.Set("Cache-Control", "no-store")
	h.Set("Referrer-Policy", "no-referrer")
	h.Set("X-Frame-Options", "DENY")
	h.Set("Content-Security-Policy", appConfirmCSP)
	w.WriteHeader(http.StatusOK)
	if err := appConfirmPage.Execute(w, struct {
		Copy     any
		Continue string
	}{appConfirmCopy, continuePath}); err != nil {
		s.logger().Error("auth", "id", httpx.RequestIDFrom(r.Context()), "op", "app_confirm", "err", err)
	}
}
