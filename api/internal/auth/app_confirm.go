package auth

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"image/color"
	"net/http"
	"strings"
	"sync"

	"github.com/jhunthrop/foreversixty/api/internal/brand"
	"github.com/jhunthrop/foreversixty/api/internal/copygen"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// appConfirmCopy is the confirm page's words, shared with its test.
var appConfirmCopy = struct {
	Title, Heading, Button, NotYou, Wordmark string
}{
	Title:    copygen.ConfirmAppTitle,
	Heading:  copygen.ConfirmAppHeading,
	Button:   copygen.ConfirmAppButton,
	NotYou:   copygen.ConfirmAppNotYou,
	Wordmark: copygen.AppName,
}

// appConfirmPage is self-contained: no script, no form, nothing fetched. The site's own faces, the
// header's flat seal and the painted band and ground travel inside the page as data: URLs, the
// only kind its CSP admits. The one link is the click the flow needs. Round 5 within that
// (design/mocks/directions/app_confirm.py): the painted band fading into the painted earth, the
// wordmark as the site's header draws it (seal, then the name in Marcellus), the title, then one
// framed panel with the teal primary and the not-you line (the title and the button say the
// rest; ai-slop review 2026-11-10).
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
<span class="mark"><img src="{{.Seal}}" width="40" height="40" alt="">{{.Copy.Wordmark}}</span>
<h1>{{.Copy.Heading}}</h1>
</div></div>
<main><div class="col panel">
<a class="go" href="{{.Continue}}">{{.Copy.Button}}</a>
<p class="note">{{.Copy.NotYou}}</p>
</div></main>
</body>
</html>
`))

// The wordmark's size and gap, the header bar's own (web/src/styles/nav.css .site-wordmark-type:
// 25 px from 1280, 22 px below it; the seal at 40 px with a 10 px gap). The column is the form
// layout's 560 px.
const (
	wordmarkSize      = 25
	wordmarkSizePhone = 22
	wordmarkGap       = 10
	columnWidth       = 560
	phoneBreakpoint   = 768
)

// The faces the page inlines: the family names the token stacks open with.
var appConfirmFaces = []struct {
	family string
	weight int
	face   brand.Face
}{
	{"Marcellus", 400, brand.Marcellus400},
	{"Open Sans", 400, brand.OpenSans400},
	{"Open Sans", 600, brand.OpenSans600},
}

// fontFaces is one @font-face per inlined face, each a data: URL of the site's own subset TTF.
func fontFaces() (string, error) {
	var b strings.Builder
	for _, f := range appConfirmFaces {
		raw, err := brand.FontFile(f.face)
		if err != nil {
			return "", err
		}
		fmt.Fprintf(&b, "@font-face{font-family:'%s';font-weight:%d;font-display:block;src:url(data:font/ttf;base64,%s) format('truetype')}\n",
			f.family, f.weight, base64.StdEncoding.EncodeToString(raw))
	}
	return b.String(), nil
}

// sealURL is the header's flat seal as a data: URL.
func sealURL() (template.URL, error) {
	raw, err := brand.NavSealSVG()
	if err != nil {
		return "", err
	}
	return template.URL("data:image/svg+xml;base64," + base64.StdEncoding.EncodeToString(raw)), nil
}

// webpURL is WebP art as a data: URL.
func webpURL(art func() ([]byte, error)) (string, error) {
	raw, err := art()
	if err != nil {
		return "", err
	}
	return "data:image/webp;base64," + base64.StdEncoding.EncodeToString(raw), nil
}

// bandFade is how far the band's paint runs under the panel and fades out, the site's
// --band-fade on a phone-first page (web/src/components/ui/band.css).
const bandFade = 96

// typeCSS is a type role as a CSS font shorthand at one width, its tracking after it.
func typeCSS(r brand.TypeRole, size, line float64) string {
	return fmt.Sprintf("font:%d %gpx/%gpx %s;letter-spacing:%gem", r.Weight, size, line, r.Family, r.TrackingEm)
}

// appConfirmStyle is the page's one stylesheet. Its colours, radii, focus stroke, faces and type
// roles are the design tokens (api/internal/brand, generated from design/tokens/tokens.json), so
// a palette or type change reaches it.
func appConfirmStyle() (template.CSS, error) {
	faces, err := fontFaces()
	if err != nil {
		return "", err
	}
	ground, err := webpURL(brand.GroundTileWebP)
	if err != nil {
		return "", err
	}
	band, err := webpURL(brand.BandWebP)
	if err != nil {
		return "", err
	}
	title, body, small := brand.TypeTitle, brand.TypeBody, brand.TypeSmall
	return template.CSS(faces + fmt.Sprintf(`:root{color-scheme:dark}*{box-sizing:border-box}
body{margin:0;min-height:100vh;isolation:isolate;background:%[1]s url(%[28]s) 0 0/%[29]dpx repeat;color:%[2]s;%[3]s;-webkit-font-smoothing:antialiased}
.band{position:relative;padding:0 18px %[32]dpx}
.band::before{content:"";position:absolute;inset:0;z-index:-1;pointer-events:none;
background:linear-gradient(%[31]s,%[31]s),url(%[30]s) center top/cover no-repeat,%[4]s;
mask-image:linear-gradient(to bottom,#000 calc(100%% - %[32]dpx),rgba(0,0,0,.5) calc(100%% - %[32]dpx / 2),transparent)}
.col{max-width:%[19]dpx;margin:0 auto}
.mark{display:flex;align-items:center;gap:%[20]dpx;padding:16px 0 0;font:400 %[21]dpx/1 %[5]s;color:%[6]s;letter-spacing:%[22]gem;white-space:nowrap}
.mark img{display:block;flex:none}
h1{margin:0;padding:22px 0 0;%[23]s;color:%[6]s;text-shadow:0 1px 2px rgba(0,0,0,.45);text-wrap:balance}
main{margin-top:-76px;padding:0 18px 40px}
.panel{position:relative;background:%[7]s;border:1px solid #000;box-shadow:0 0 0 1px %[8]s;border-radius:%[9]dpx;padding:22px 24px}
.panel::before{content:"";position:absolute;inset:6px;border:1px solid %[10]s;border-radius:%[11]dpx;pointer-events:none}
.panel p{margin:0}
.go{position:relative;display:flex;align-items:center;justify-content:center;height:44px;padding:0 20px;
border:1px solid %[12]s;border-radius:%[13]dpx;background:%[14]s;color:%[15]s;font-weight:600;text-decoration:none}
.go:hover{background:%[12]s}.go:active{transform:translateY(1px);filter:brightness(.92)}
.go:focus-visible{outline:%[16]dpx solid %[17]s;outline-offset:2px}
.panel .note{margin-top:16px;%[24]s;color:%[18]s}
@media(max-width:%[25]dpx){.mark{font-size:%[26]dpx}h1{%[27]s}}`,
		cssColor(brand.ColorForeverGround), cssColor(brand.ColorBaseText), typeCSS(body, body.Size, body.LineHeight),
		cssColor(brand.ColorForeverBand), brand.FontTitle, cssColor(brand.ColorForeverCream),
		cssColor(brand.ColorBaseRaised), cssColor(brand.ColorForeverPanelEdge), brand.RadiusPlate,
		cssColor(brand.ColorForeverFrameLine), brand.RadiusPlate-4,
		cssColor(brand.ColorForeverTealEdge), brand.RadiusField, cssColor(brand.ColorForeverTeal), cssColor(brand.ColorBaseStrong),
		brand.StrokeFocus, cssColor(brand.ColorBaseGold), cssColor(brand.ColorForeverMutedCard),
		columnWidth, wordmarkGap, wordmarkSize, title.TrackingEm,
		typeCSS(title, title.Size, title.LineHeight), typeCSS(small, small.Size, small.LineHeight),
		phoneBreakpoint-1, wordmarkSizePhone, typeCSS(title, title.PhoneSize, title.PhoneLineHeight),
		ground, brand.GroundTilePx, band, cssColor(brand.ColorForeverVeilBand), bandFade)), nil
}

// cssColor writes a token as CSS: #rrggbb when opaque, rgba() otherwise.
func cssColor(c color.NRGBA) string {
	if c.A == 0xFF {
		return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
	}
	return fmt.Sprintf("rgba(%d,%d,%d,%.3g)", c.R, c.G, c.B, float64(c.A)/0xFF)
}

// appConfirmCSP lets the page style itself and draw its inlined faces, seal and art, and nothing else:
// no script, no fetch, no form, and no frame around it, so the click cannot be lifted by an overlay.
const appConfirmCSP = "default-src 'none'; style-src 'unsafe-inline'; font-src data:; img-src data:; base-uri 'none'; form-action 'none'; frame-ancestors 'none'"

// appConfirmData is the page's template input.
type appConfirmData struct {
	Copy     any
	Continue string
	Style    template.CSS
	Seal     template.URL
}

// appConfirmAssets are the page's stylesheet and seal, built once: they depend only on the
// embedded tokens, faces and art.
var appConfirmAssets = sync.OnceValues(func() (appConfirmData, error) {
	style, err := appConfirmStyle()
	if err != nil {
		return appConfirmData{}, err
	}
	seal, err := sealURL()
	if err != nil {
		return appConfirmData{}, err
	}
	return appConfirmData{Style: style, Seal: seal}, nil
})

func (s *Service) writeAppConfirm(w http.ResponseWriter, r *http.Request, continuePath string) {
	data, err := appConfirmAssets()
	if err != nil {
		s.logger().Error("auth", "id", httpx.RequestIDFrom(r.Context()), "op", "app_confirm", "err", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	h := w.Header()
	h.Set("Content-Type", "text/html; charset=utf-8")
	h.Set("Cache-Control", "no-store")
	h.Set("Referrer-Policy", "no-referrer")
	h.Set("X-Frame-Options", "DENY")
	h.Set("Content-Security-Policy", appConfirmCSP)
	w.WriteHeader(http.StatusOK)
	data.Copy, data.Continue = appConfirmCopy, continuePath
	if err := appConfirmPage.Execute(w, data); err != nil {
		s.logger().Error("auth", "id", httpx.RequestIDFrom(r.Context()), "op", "app_confirm", "err", err)
	}
}
