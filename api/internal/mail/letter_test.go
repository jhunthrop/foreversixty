package mail

import (
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/brand"
)

var sent = time.Date(2026, time.November, 10, 18, 0, 0, 0, time.UTC)

func sample() Letter {
	return Letter{
		To: "a@b.c", Subject: "Your Premium plan renews on 12 Dec 2026", Preheader: "$50 a year on 12 Dec 2026.",
		Heading: "Your Premium plan renews on 12 Dec 2026", Lead: "To stop it, cancel before then.",
		Facts:  []Fact{{Label: "Plan", Value: "Premium"}, {Label: "Price", Value: "$50 a year"}},
		Note:   "Founder price, kept while your plan stays active.",
		Action: Action{Label: "Manage or cancel", URL: "https://foreversixty.gg/account#plan"},
	}
}

func compose(t *testing.T, l Letter) Message {
	t.Helper()
	m, err := Compose(l, sent)
	if err != nil {
		t.Fatal(err)
	}
	return m
}

func TestComposeKeepsTheEnvelope(t *testing.T) {
	m := compose(t, sample())
	if m.To != "a@b.c" || m.Subject != sample().Subject || m.Text == "" || m.HTML == "" {
		t.Fatalf("message = %+v", m)
	}
}

// The shell is the site's: the hosted raster seal (Gmail drops SVG), the token colours, the
// Marcellus title over Open Sans with client-safe fallbacks, and the site footer's lines.
func TestHTMLIsTheBrandedShell(t *testing.T) {
	h := compose(t, sample()).HTML
	for _, want := range []string{
		`src="https://foreversixty.gg/seal-400.png"`,
		hex(brand.ColorForeverGround), hex(brand.ColorBaseRaised), hex(brand.ColorForeverTeal), `href="mailto:help@foreversixty.gg"`,
		"&#39;Georgia&#39;, serif", "&#39;Arial&#39;, sans-serif", "Forever Sixty</span>",
		"A fan-made player tool. Not affiliated with Blizzard.",
		"© 2026 Commish LLC d/b/a Forever Sixty", "help@foreversixty.gg",
		`role="presentation"`, `max-width:560px`,
		`href="https://foreversixty.gg/account#plan"`, ">Manage or cancel</a>",
		">Plan</td>", ">$50 a year</td>", "$50 a year on 12 Dec 2026.",
	} {
		if !strings.Contains(h, want) {
			t.Errorf("html has no %q", want)
		}
	}
	if strings.Contains(h, "data:image") || strings.Contains(h, "<svg") {
		t.Error("html embeds an image instead of loading the hosted seal")
	}
}

func TestTextCarriesEveryBlockInOrder(t *testing.T) {
	text := compose(t, sample()).Text
	order := []string{"Your Premium plan renews", "To stop it", "Plan: Premium\nPrice: $50 a year", "Founder price",
		"Manage or cancel: https://foreversixty.gg/account#plan", "--\nA fan-made player tool.", "© 2026 Commish LLC",
		"help@foreversixty.gg"}
	at := 0
	for _, want := range order {
		i := strings.Index(text[at:], want)
		if i < 0 {
			t.Fatalf("text has no %q after offset %d:\n%s", want, at, text)
		}
		at += i
	}
}

func TestEmptyBlocksDrawNothing(t *testing.T) {
	l := sample()
	l.Preheader, l.Facts, l.Note = "", nil, ""
	m := compose(t, l)
	if strings.Contains(m.HTML, "display:none") || strings.Contains(m.HTML, "border-top:1px solid") {
		t.Errorf("empty preheader or facts drew markup")
	}
	if strings.Contains(m.Text, "Founder") || strings.Contains(m.Text, "Plan:") {
		t.Errorf("text = %q", m.Text)
	}
}

func TestValuesAreEscaped(t *testing.T) {
	l := sample()
	l.Lead = `<script>alert(1)</script> & "x"`
	l.Action.URL = "javascript:alert(1)"
	h := compose(t, l).HTML
	if strings.Contains(h, "<script>") || strings.Contains(h, `href="javascript:`) {
		t.Errorf("unescaped value in html")
	}
}

func TestFooterYearFollowsTheClock(t *testing.T) {
	m, err := Compose(sample(), time.Date(2031, 1, 1, 0, 0, 0, 0, time.UTC))
	if err != nil || !strings.Contains(m.HTML, "© 2031 ") || !strings.Contains(m.Text, "© 2031 ") {
		t.Fatalf("year not 2031: %v", err)
	}
}
