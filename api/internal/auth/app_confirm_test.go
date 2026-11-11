package auth

import (
	"image/color"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/brand"
)

func renderAppConfirm(t *testing.T) string {
	t.Helper()
	rec := httptest.NewRecorder()
	(&Service{}).writeAppConfirm(rec, httptest.NewRequest(http.MethodGet, "/v1/auth/app/start", nil), "/v1/auth/app/continue?x=1")
	return rec.Body.String()
}

// The page is round 5 in the design tokens' own values: the ground, the band, the framed panel and
// the teal primary, none of the retired gold-on-brown palette or a bare system-ui stack.
func TestAppConfirmWearsTheDesignTokens(t *testing.T) {
	body := renderAppConfirm(t)
	for _, token := range []color.NRGBA{
		brand.ColorForeverGround, brand.ColorForeverBand, brand.ColorForeverCream, brand.ColorBaseRaised,
		brand.ColorForeverPanelEdge, brand.ColorForeverFrameLine, brand.ColorForeverTeal, brand.ColorForeverTealEdge,
		brand.ColorForeverMutedCard,
	} {
		if !strings.Contains(body, cssColor(token)) {
			t.Errorf("the page does not use the token %s", cssColor(token))
		}
	}
	for _, retired := range []string{"#d9a441", "#0f0d0a", "#1a1712", "#3a3226", "#b9ae99", "system-ui"} {
		if strings.Contains(body, retired) {
			t.Errorf("the page still carries the retired %q", retired)
		}
	}
	for _, want := range []string{
		`<a class="go" href="/v1/auth/app/continue?x=1">`, `<div class="col panel">`,
		brand.FontTitle, brand.FontBody, `<img src="data:image/svg&#43;xml;base64,`,
		"font-family:'Marcellus';font-weight:400", "font-family:'Open Sans';font-weight:600",
	} {
		if !strings.Contains(body, want) {
			t.Errorf("the page is missing %q", want)
		}
	}
}

func TestCSSColorWritesOpaqueAsHexAndTranslucentAsRGBA(t *testing.T) {
	for _, c := range []struct {
		in   color.NRGBA
		want string
	}{
		{color.NRGBA{R: 0x0F, G: 0x60, B: 0x64, A: 0xFF}, "#0f6064"},
		{color.NRGBA{R: 0x82, G: 0x70, B: 0x5B, A: 0x8C}, "rgba(130,112,91,0.549)"},
	} {
		if got := cssColor(c.in); got != c.want {
			t.Errorf("cssColor(%v) = %q, want %q", c.in, got, c.want)
		}
	}
}
