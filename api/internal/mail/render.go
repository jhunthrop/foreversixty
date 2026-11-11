package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"image/color"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/brand"
	"github.com/jhunthrop/foreversixty/api/internal/copygen"
)

// The shell's measures. Email clients honour few stylesheets, so the layout is tables with
// inline styles: one column, the site's 560 px reading width, the panel padding stepping down on
// a phone through the one media query clients that support it read.
const (
	columnWidth = 560
	sealSize    = 48
)

// shellStyle is every colour and face the shell paints with, from the design tokens
// (brand/tokens_gen.go, generated from design/tokens/tokens.json): the site's earth ground, the
// raised panel and its text, the teal button, and the footer's muted words. Font
// stacks end in faces every client has (Georgia, Arial), so a client without the site's fonts
// still reads a serif title over a sans body.
type shellStyle struct {
	Ground, Panel, PanelEdge, Rule, Title, Text, Strong, Muted      template.CSS
	Button, ButtonEdge, ButtonText, FooterText, TitleFont, BodyFont template.CSS
	GroundAttr, PanelAttr, ButtonAttr                               string
	ColumnWidth, SealSize                                           int
}

func css(c color.NRGBA) template.CSS { return template.CSS(hex(c)) }

// hex is an opaque token colour as #rrggbb; email clients mishandle alpha.
func hex(c color.NRGBA) string { return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B) }

var style = shellStyle{
	Ground: css(brand.ColorForeverGround), Panel: css(brand.ColorBaseRaised),
	PanelEdge: css(brand.ColorForeverRuleWarm), Rule: css(brand.ColorBaseLine),
	Title: css(brand.ColorForeverCream2), Text: css(brand.ColorBaseText), Strong: css(brand.ColorBaseStrong),
	Muted:  css(brand.ColorBaseMuted),
	Button: css(brand.ColorForeverTeal), ButtonEdge: css(brand.ColorForeverTealEdge),
	ButtonText: css(brand.ColorForeverCream), FooterText: css(brand.ColorForeverMutedGround),
	TitleFont: template.CSS(brand.FontTitle), BodyFont: template.CSS(brand.FontBody),
	GroundAttr: hex(brand.ColorForeverGround), PanelAttr: hex(brand.ColorBaseRaised),
	ButtonAttr: hex(brand.ColorForeverTeal), ColumnWidth: columnWidth, SealSize: sealSize,
}

// shellView is what the template draws.
type shellView struct {
	Letter
	S                                    shellStyle
	SiteURL, SealURL, Wordmark, Fallback string
	Footer                               footer
}

// fontsHref loads the site's two faces where the client allows a stylesheet (Apple Mail, iOS);
// elsewhere the stacks' fallbacks draw.
const fontsHref = "https://fonts.googleapis.com/css2?family=Marcellus&family=Open+Sans:wght@400;600&display=swap"

var shell = template.Must(template.New("mail").Funcs(template.FuncMap{
	"fonts": func() template.URL { return fontsHref },
}).Parse(shellHTML))

func renderHTML(l Letter, now time.Time) (string, error) {
	view := shellView{
		Letter: l, S: style, SiteURL: brand.SiteOrigin, SealURL: brand.SiteOrigin + brand.SealPath,
		Wordmark: copygen.AppName, Fallback: copygen.MailLinkFallback, Footer: footerOf(now),
	}
	var b bytes.Buffer
	if err := shell.Execute(&b, view); err != nil {
		return "", fmt.Errorf("mail: render %q: %w", l.Subject, err)
	}
	return b.String(), nil
}

// shellHTML is the one letter shell. Every value is escaped by html/template; the styles are
// typed CSS from the tokens above.
const shellHTML = `<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="color-scheme" content="dark">
<meta name="supported-color-schemes" content="dark">
<title>{{.Subject}}</title>
<link rel="stylesheet" href="{{fonts}}">
<style>
@media (max-width: 600px) { .fs-pad { padding: 24px 20px !important; } .fs-outer { padding: 20px 12px !important; } }
</style>
</head>
<body style="margin:0;padding:0;background-color:{{.S.Ground}};">
{{- if .Preheader}}
<div style="display:none;max-height:0;max-width:0;overflow:hidden;opacity:0;font-size:1px;line-height:1px;">{{.Preheader}}</div>
{{- end}}
<table role="presentation" width="100%" cellpadding="0" cellspacing="0" border="0" bgcolor="{{.S.GroundAttr}}" style="background-color:{{.S.Ground}};">
<tr><td class="fs-outer" align="center" style="padding:32px 16px;">
<table role="presentation" width="{{.S.ColumnWidth}}" cellpadding="0" cellspacing="0" border="0" style="width:100%;max-width:{{.S.ColumnWidth}}px;">
<tr><td style="padding:0 4px 20px;">
<a href="{{.SiteURL}}" style="text-decoration:none;color:{{.S.Title}};">
<img src="{{.SealURL}}" width="{{.S.SealSize}}" height="{{.S.SealSize}}" alt="" style="display:inline-block;vertical-align:middle;border:0;outline:none;">
<span style="display:inline-block;vertical-align:middle;padding-left:12px;font-family:{{.S.TitleFont}};font-size:24px;line-height:28px;color:{{.S.Title}};">{{.Wordmark}}</span>
</a>
</td></tr>
<tr><td class="fs-pad" bgcolor="{{.S.PanelAttr}}" style="background-color:{{.S.Panel}};border:1px solid {{.S.PanelEdge}};border-radius:8px;padding:32px;">
<h1 style="margin:0 0 12px;font-family:{{.S.TitleFont}};font-weight:400;font-size:26px;line-height:32px;color:{{.S.Title}};">{{.Heading}}</h1>
<p style="margin:0 0 24px;font-family:{{.S.BodyFont}};font-size:16px;line-height:25px;color:{{.S.Text}};">{{.Lead}}</p>
{{- if .Facts}}
<table role="presentation" width="100%" cellpadding="0" cellspacing="0" border="0" style="margin:0 0 24px;border-top:1px solid {{.S.Rule}};">
{{- range .Facts}}
<tr>
<td style="padding:10px 12px 10px 0;border-bottom:1px solid {{$.S.Rule}};font-family:{{$.S.BodyFont}};font-size:14px;line-height:20px;color:{{$.S.Muted}};width:38%;vertical-align:top;">{{.Label}}</td>
<td style="padding:10px 0;border-bottom:1px solid {{$.S.Rule}};font-family:{{$.S.BodyFont}};font-size:15px;line-height:20px;font-weight:600;color:{{$.S.Strong}};vertical-align:top;">{{.Value}}</td>
</tr>
{{- end}}
</table>
{{- end}}
{{- if .Note}}
<p style="margin:0 0 24px;font-family:{{.S.BodyFont}};font-size:15px;line-height:23px;color:{{.S.Text}};">{{.Note}}</p>
{{- end}}
<table role="presentation" cellpadding="0" cellspacing="0" border="0" style="margin:0 0 16px;">
<tr><td bgcolor="{{.S.ButtonAttr}}" style="background-color:{{.S.Button}};border:1px solid {{.S.ButtonEdge}};border-radius:6px;">
<a href="{{.Action.URL}}" style="display:inline-block;padding:12px 24px;font-family:{{.S.BodyFont}};font-size:15px;line-height:20px;font-weight:600;color:{{.S.ButtonText}};text-decoration:none;">{{.Action.Label}}</a>
</td></tr>
</table>
<p style="margin:0;font-family:{{.S.BodyFont}};font-size:12px;line-height:18px;color:{{.S.Muted}};">{{.Fallback}}<br><a href="{{.Action.URL}}" style="color:{{.S.Muted}};text-decoration:underline;word-break:break-all;">{{.Action.URL}}</a></p>
</td></tr>
<tr><td style="padding:24px 4px 0;font-family:{{.S.BodyFont}};font-size:13px;line-height:20px;color:{{.S.FooterText}};">
{{.Footer.FanLine}}<br>{{.Footer.Operator}}<br>{{.Footer.ContactLead}}<a href="mailto:{{.Footer.Email}}" style="color:{{.S.FooterText}};text-decoration:underline;">{{.Footer.Email}}</a>{{.Footer.ContactTail}}
</td></tr>
</table>
</td></tr>
</table>
</body>
</html>
`
