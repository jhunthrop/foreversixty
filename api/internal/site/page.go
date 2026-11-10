package site

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/builds"
)

// islandPath is where the site publishes the planner island bundle; the
// contract fixes it. islandStylesPath is the contract's "emitted as
// planner-island.css and linked by the API page" option, which this page
// takes because the site's header and footer markup below needs the site's
// CSS to look like the site.
const (
	islandPath       = "/planner-island.js"
	islandStylesPath = "/planner-island.css"
)

type pageData struct {
	Title        string
	Description  string
	CanonicalURL string
	CardURL      string
	StylesURL    string
	IslandURL    string
	FaviconURL   string
	Header       template.HTML
	Footer       template.HTML
	BuildJSON    string
	TreeVersion  string
}

// The page column is the web Base layout's one `.page-column` (web/src/layouts/Base.astro,
// global.css), so the shared build sits in the site's column, not edge to edge (premium spec
// 2026-11-08 bug B4). The island mounts the planner read-only into #planner and draws the whole
// page head itself (the record's title, views, the answer); the record's title and views reach it
// in data-build. Until it mounts, .b-reserve holds a screen's height so the footer, the only
// thing under it, never paints in the first screen and then jumps (spec section 13).
//
// BuildJSON is interpolated into a single-quoted attribute; html/template escapes the quotes and
// angle brackets for that context, and the browser unescapes them when the island reads it.
var buildPage = template.Must(template.New("build").Parse(`<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>{{.Title}}</title>
<meta name="description" content="{{.Description}}">
<link rel="canonical" href="{{.CanonicalURL}}">
<link rel="icon" href="{{.FaviconURL}}" type="image/svg+xml">
<link rel="stylesheet" href="{{.StylesURL}}">
<meta property="og:title" content="{{.Title}}">
<meta property="og:description" content="{{.Description}}">
<meta property="og:image" content="{{.CardURL}}">
<meta property="og:url" content="{{.CanonicalURL}}">
<meta name="twitter:card" content="summary_large_image">
<style>.b-reserve{min-height:100vh}</style>
</head>
<body class="min-h-screen flex flex-col bg-bg">
{{.Header}}
<div class="page-column">
<main id="main" tabindex="-1">
<div id="planner" data-build='{{.BuildJSON}}' data-tree-version="{{.TreeVersion}}"><div class="b-reserve" aria-hidden="true"></div></div>
</main>
</div>
{{.Footer}}
<script type="module" src="{{.IslandURL}}"></script>
</body>
</html>
`))

// askSenderLine is the not-found page's line above the actions (premium spec 2026-11-08 ruling
// Q12). The headings and their sentences are handler.go's; the two actions' words are written in
// the template, where design/mocks/ws_premium.py reads the planner's.
const askSenderLine = "Ask whoever sent it for a fresh link."

// The system Button's primary recipe and the flat link (web/src/lib/ui/button.ts
// `button('primary')`, lib/ui/primitives.ts `LINK`): Tailwind utilities the site already uses, so
// planner-island.css carries every one. page_test.go holds each token to those two files.
const (
	primaryButtonClass = "rounded-field button-label inline-flex items-center justify-center gap-2 border px-[18px] whitespace-nowrap " +
		"transition-[filter,background-color,border-color] duration-[120ms] " +
		"focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-gold " +
		"border-transparent text-bg bg-[linear-gradient(180deg,#f0cc6c,#c99a3a)] " +
		"shadow-[inset_0_0_0_1px_var(--color-gold-deep),0_6px_20px_rgba(229,185,85,.22)] " +
		"[@media(hover:hover)_and_(pointer:fine)]:hover:hover:brightness-[1.08] active:translate-y-px active:brightness-[.92] h-11"
	linkClass = "text-gold font-semibold hover:text-gold-hover active:text-gold-hover rounded-field " +
		"focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-gold"
)

type messageData struct {
	Title       string
	Heading     string
	Body        string
	Ask         string
	StylesURL   string
	FaviconURL  string
	PlannerURL  string
	GuidesURL   string
	ButtonClass string
	LinkClass   string
	Header      template.HTML
	Footer      template.HTML
}

// messagePage is the web system's message page (premium spec 2026-11-08 section 8, mocks
// web-b-missing and web-b-error): the PageHeader (the title in Marcellus, the sentence as its meta
// line) over the form layout's 560 px column, then the planner as the primary and the guides'
// builds as a link. The PageHeader and Page rules are the Svelte primitives' (ui/PageHeader,
// ui/Page layout form), restated here because this page renders no Svelte; the parity spec
// (web/tests/e2e/parity-b.spec.ts) holds it to the mock.
var messagePage = template.Must(template.New("message").Parse(`<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>{{.Title}}</title>
<meta name="robots" content="noindex">
<link rel="icon" href="{{.FaviconURL}}" type="image/svg+xml">
<link rel="stylesheet" href="{{.StylesURL}}">
<style>
.b-head{padding:16px var(--page-inset) 14px}
.b-title-row{display:flex;align-items:center;gap:16px;min-width:0}
.b-title-text{min-width:0}
.b-title{margin:0;font-family:var(--font-title);font-weight:400;font-size:var(--text-title-phone);line-height:var(--text-title-phone--line-height);letter-spacing:.01em;color:var(--color-strong);overflow-wrap:anywhere}
.b-meta{margin-top:6px;font-size:var(--text-small);line-height:var(--text-small--line-height);color:var(--color-muted)}
.b-form{padding:20px var(--page-inset) 0}
.b-stack{display:flex;flex-direction:column;gap:12px;max-width:560px}
.b-acts{display:flex;flex-wrap:wrap;align-items:center;gap:10px 16px}
.b-acts p{margin:0}
@media (min-width:768px){.b-title{font-size:var(--text-title);line-height:var(--text-title--line-height)}}
@media (min-width:900px){.b-head{padding-block:20px 16px}}
</style>
</head>
<body class="min-h-screen flex flex-col bg-bg">
{{.Header}}
<div class="page-column">
<main id="main" tabindex="-1">
<header class="b-head" data-region="head">
<div class="b-title-row" data-region="head-title">
<div class="b-title-text">
<h1 class="b-title" data-testid="page-title">{{.Heading}}</h1>
<div class="b-meta" data-region="head-meta">{{.Body}}</div>
</div>
</div>
</header>
<div class="b-form">
<div class="b-stack">
{{if .Ask}}<p class="b-meta" data-testid="b-message-ask">{{.Ask}}</p>
{{end}}<div class="b-acts">
<p><a href="{{.PlannerURL}}" class="{{.ButtonClass}}">Open the build planner</a></p>
<p><a class="{{.LinkClass}}" href="{{.GuidesURL}}">Talent builds in the class guides ›</a></p>
</div>
</div>
</div>
</main>
</div>
{{.Footer}}
</body>
</html>
`))

// buildPageHTML renders the shared build page for b. A sim lookup
// failure is logged and otherwise ignored: the page still renders,
// simply without a DPS line in its description.
func (d Deps) buildPageHTML(ctx context.Context, b builds.Build) (string, error) {
	desc, err := builds.Describe(d.Data, b).WithSimDPS(ctx, b.ID, d.Sims)
	if err != nil {
		d.logger().Error("site", "op", "describe", "build", b.ID, "err", err)
	}
	record, err := json.Marshal(b)
	if err != nil {
		return "", fmt.Errorf("site: marshal record %s: %w", b.ID, err)
	}
	var out strings.Builder
	err = buildPage.Execute(&out, pageData{
		Title:        fmt.Sprintf("%s · %s · Forever Sixty", desc.Title, desc.SplitText),
		Description:  pageDescription(desc),
		CanonicalURL: d.PublicBaseURL + "/b/" + b.ID,
		CardURL:      d.PublicBaseURL + "/b/" + b.ID + "/card.png",
		StylesURL:    d.PublicBaseURL + islandStylesPath,
		IslandURL:    d.PublicBaseURL + islandPath,
		FaviconURL:   d.PublicBaseURL + "/favicon.svg",
		Header:       Header,
		Footer:       Footer,
		BuildJSON:    string(record),
		TreeVersion:  b.TreeVersion,
	})
	if err != nil {
		return "", fmt.Errorf("site: render build page %s: %w", b.ID, err)
	}
	return out.String(), nil
}

// pageDescription is the meta and og description text: the build's
// shape, and its simmed DPS when Description carries a line for it.
func pageDescription(desc builds.Description) string {
	s := fmt.Sprintf("%s build, %s at level %d", desc.Heading, desc.SplitText, desc.Level)
	if desc.DPSLine != "" {
		s += ", simmed to " + desc.DPSLine
	}
	return s + ". Forever Sixty build planner."
}

// messageHTML renders the page used for a missing or unloadable build. ask is the line above
// the actions ("" for none).
func (d Deps) messageHTML(heading, body, ask string) (string, error) {
	var out strings.Builder
	err := messagePage.Execute(&out, messageData{
		Title:       heading + " · Forever Sixty",
		Heading:     heading,
		Body:        body,
		Ask:         ask,
		StylesURL:   d.PublicBaseURL + islandStylesPath,
		FaviconURL:  d.PublicBaseURL + "/favicon.svg",
		PlannerURL:  d.PublicBaseURL + "/planner",
		GuidesURL:   d.PublicBaseURL + "/guides",
		ButtonClass: primaryButtonClass,
		LinkClass:   linkClass,
		Header:      Header,
		Footer:      Footer,
	})
	if err != nil {
		return "", fmt.Errorf("site: render message page: %w", err)
	}
	return out.String(), nil
}
