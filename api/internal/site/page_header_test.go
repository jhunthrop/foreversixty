package site

import (
	"os"
	"regexp"
	"strings"
	"testing"
)

// pageHeaderSource is the web's page head (ui/PageHeader.svelte), whose scoped rules the message
// page restates because it renders no Svelte.
const pageHeaderSource = "../../../web/src/components/ui/PageHeader.svelte"

// cssRule is the declarations of the first rule for selector in css, normalised: no whitespace
// around ':' and ';', one per entry.
func cssRule(t *testing.T, css, selector string) []string {
	t.Helper()
	m := regexp.MustCompile(regexp.QuoteMeta(selector) + `\s*\{([^}]*)\}`).FindStringSubmatch(css)
	if m == nil {
		t.Fatalf("no %s rule", selector)
	}
	var out []string
	for _, decl := range strings.Split(m[1], ";") {
		name, value, ok := strings.Cut(decl, ":")
		if !ok {
			continue
		}
		out = append(out, strings.TrimSpace(name)+":"+strings.Join(strings.Fields(value), " "))
	}
	return out
}

// TestMessagePageRestatesThePageHeader holds the message page's .b-head and .b-title to the web's
// .ui-page-header and .ui-title, declaration for declaration, so a change to the site's page head
// fails here until the API page follows (component audit 2026-11-10, cluster 20).
func TestMessagePageRestatesThePageHeader(t *testing.T) {
	src, err := os.ReadFile(pageHeaderSource)
	if err != nil {
		t.Fatal(err)
	}
	var page strings.Builder
	if err := messagePage.Execute(&page, messageData{}); err != nil {
		t.Fatal(err)
	}
	for web, api := range map[string]string{".ui-page-header": ".b-head", ".ui-title": ".b-title"} {
		ours := strings.Join(cssRule(t, page.String(), api), ";")
		for _, decl := range cssRule(t, string(src), web) {
			if !strings.Contains(";"+ours+";", ";"+decl+";") {
				t.Errorf("%s lacks %s's %q (web/src/components/ui/PageHeader.svelte)", api, web, decl)
			}
		}
	}
}
