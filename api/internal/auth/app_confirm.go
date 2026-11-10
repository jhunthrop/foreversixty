package auth

import (
	"fmt"
	"html/template"
	"image/color"
	"net/http"

	"github.com/jhunthrop/foreversixty/api/internal/brand"
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
// The one link is the click the flow needs. Round 5 within that (design/mocks/directions/app_confirm.py):
// the band colour fading into the earth ground (the painted art's own fallbacks), the wordmark as type,
// the title, then one framed panel with the sentence, the teal primary and the not-you line.
var appConfirmPage = template.Must(template.New("confirm").Parse(`<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="robots" content="noindex">
<title>{{.Copy.Title}}</title>
<style>{{.Style}}</style>
</head>
<body>
<div class="band"><div class="col">
<span class="mark">Forever Sixty</span>
<h1>{{.Copy.Heading}}</h1>
</div></div>
<main><div class="col panel">
<p>{{.Copy.Body}}</p>
<a class="go" href="{{.Continue}}">{{.Copy.Button}}</a>
<p class="note">{{.Copy.NotYou}}</p>
</div></main>
</body>
</html>
`))

// The type stacks: the site's faces (Marcellus, Open Sans) where the device has them installed, since
// the CSP fetches none, then the nearest system faces: a flared roman for the title, a humanist sans
// for the text.
const (
	appConfirmTitleFont = "'Marcellus',Optima,Candara,Georgia,serif"
	appConfirmBodyFont  = "'Open Sans','Segoe UI',Roboto,'Helvetica Neue',Arial,sans-serif"
)

// appConfirmStyle is the page's one stylesheet. Its colours, radii and focus stroke are the design
// tokens (api/internal/brand, generated from design/tokens/tokens.json), so a palette change reaches it.
var appConfirmStyle = template.CSS(fmt.Sprintf(`:root{color-scheme:dark}*{box-sizing:border-box}
body{margin:0;min-height:100vh;background:%[1]s;color:%[2]s;font:16px/1.55 %[3]s;-webkit-font-smoothing:antialiased}
.band{background:linear-gradient(%[4]s 0,%[4]s calc(100%% - 96px),%[1]s 100%%);padding:0 18px 96px}
.col{max-width:560px;margin:0 auto}
.mark{display:block;padding:18px 0 0;font:400 22px/1 %[5]s;color:%[6]s;letter-spacing:-.015em}
h1{margin:0;padding:22px 0 0;font:400 21px/25px %[5]s;letter-spacing:-.015em;color:%[6]s;text-shadow:0 1px 2px rgba(0,0,0,.45)}
main{margin-top:-76px;padding:0 18px 40px}
.panel{position:relative;background:%[7]s;border:1px solid #000;box-shadow:0 0 0 1px %[8]s;border-radius:%[9]dpx;padding:22px 24px}
.panel::before{content:"";position:absolute;inset:6px;border:1px solid %[10]s;border-radius:%[11]dpx;pointer-events:none}
.panel p{margin:0}
.go{position:relative;display:flex;align-items:center;justify-content:center;height:44px;margin-top:18px;padding:0 20px;
border:1px solid %[12]s;border-radius:%[13]dpx;background:%[14]s;color:%[15]s;font-weight:600;text-decoration:none}
.go:hover{background:%[12]s}.go:active{transform:translateY(1px);filter:brightness(.92)}
.go:focus-visible{outline:%[16]dpx solid %[17]s;outline-offset:2px}
.panel .note{margin-top:16px;font-size:14px;line-height:20px;color:%[18]s}
@media(min-width:768px){.mark{padding-top:24px}h1{padding-top:28px;font-size:26px;line-height:30px}}`,
	cssColor(brand.ColorForeverGround), cssColor(brand.ColorBaseText), appConfirmBodyFont,
	cssColor(brand.ColorForeverBand), appConfirmTitleFont, cssColor(brand.ColorForeverCream),
	cssColor(brand.ColorBaseRaised), cssColor(brand.ColorForeverPanelEdge), brand.RadiusPlate,
	cssColor(brand.ColorForeverFrameLine), brand.RadiusPlate-4,
	cssColor(brand.ColorForeverTealEdge), brand.RadiusField, cssColor(brand.ColorForeverTeal), cssColor(brand.ColorBaseStrong),
	brand.StrokeFocus, cssColor(brand.ColorBaseGold), cssColor(brand.ColorForeverMutedCard)))

// cssColor writes a token as CSS: #rrggbb when opaque, rgba() otherwise.
func cssColor(c color.NRGBA) string {
	if c.A == 0xFF {
		return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
	}
	return fmt.Sprintf("rgba(%d,%d,%d,%.3g)", c.R, c.G, c.B, float64(c.A)/0xFF)
}

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
		Style    template.CSS
	}{appConfirmCopy, continuePath, appConfirmStyle}); err != nil {
		s.logger().Error("auth", "id", httpx.RequestIDFrom(r.Context()), "op", "app_confirm", "err", err)
	}
}
