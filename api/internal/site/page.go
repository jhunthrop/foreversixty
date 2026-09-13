package site

import (
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

// BuildJSON is interpolated into a single-quoted attribute; html/template
// escapes the quotes and angle brackets for that context, and the browser
// unescapes them when the island reads the attribute.
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
</head>
<body class="min-h-screen flex flex-col bg-bg">
{{.Header}}
<main id="main">
<div id="planner" data-build='{{.BuildJSON}}' data-tree-version="{{.TreeVersion}}"></div>
</main>
{{.Footer}}
<script type="module" src="{{.IslandURL}}"></script>
</body>
</html>
`))

type messageData struct {
	Title      string
	Heading    string
	Body       string
	StylesURL  string
	FaviconURL string
	PlannerURL string
	Header     template.HTML
	Footer     template.HTML
}

var messagePage = template.Must(template.New("message").Parse(`<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>{{.Title}}</title>
<meta name="robots" content="noindex">
<link rel="icon" href="{{.FaviconURL}}" type="image/svg+xml">
<link rel="stylesheet" href="{{.StylesURL}}">
</head>
<body class="min-h-screen flex flex-col bg-bg">
{{.Header}}
<main id="main" class="px-[18px] md:px-12 py-12 flex flex-col gap-4">
<h1 class="font-display text-[28px] gold-text">{{.Heading}}</h1>
<p class="max-w-[720px] leading-relaxed text-muted">{{.Body}}</p>
<p><a href="{{.PlannerURL}}" class="inline-flex items-center min-h-11 text-nav hover:text-strong">Open the build planner</a></p>
</main>
{{.Footer}}
</body>
</html>
`))

// buildPageHTML renders the shared build page for b.
func (d Deps) buildPageHTML(b builds.Build) (string, error) {
	desc := builds.Describe(d.Data, b)
	record, err := json.Marshal(b)
	if err != nil {
		return "", fmt.Errorf("site: marshal record %s: %w", b.ID, err)
	}
	var out strings.Builder
	err = buildPage.Execute(&out, pageData{
		Title: fmt.Sprintf("%s · %s · Forever Sixty", desc.Title, desc.SplitText),
		Description: fmt.Sprintf("%s build, %s at level %d. Forever Sixty build planner.",
			desc.Heading, desc.SplitText, desc.Level),
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

// messageHTML renders the plain page used for a missing or unloadable build.
func (d Deps) messageHTML(heading, body string) (string, error) {
	var out strings.Builder
	err := messagePage.Execute(&out, messageData{
		Title:      heading + " · Forever Sixty",
		Heading:    heading,
		Body:       body,
		StylesURL:  d.PublicBaseURL + islandStylesPath,
		FaviconURL: d.PublicBaseURL + "/favicon.svg",
		PlannerURL: d.PublicBaseURL + "/planner",
		Header:     Header,
		Footer:     Footer,
	})
	if err != nil {
		return "", fmt.Errorf("site: render message page: %w", err)
	}
	return out.String(), nil
}
