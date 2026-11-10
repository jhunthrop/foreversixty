package notify

import "testing"

// Every footer carries the seal; an embed with no footer, or its own icon, is left as it is.
func TestWithSealFooter(t *testing.T) {
	const seal = "https://foreversixty.gg/seal-400.png"
	if got := WithSealFooter(Embed{Footer: "Forever Sixty"}, seal); got.FooterIcon != seal {
		t.Fatalf("footer icon %q", got.FooterIcon)
	}
	if got := WithSealFooter(Embed{Title: "x"}, seal); got.FooterIcon != "" {
		t.Fatalf("no footer, icon %q", got.FooterIcon)
	}
	if got := WithSealFooter(Embed{Footer: "f", FooterIcon: "own"}, seal); got.FooterIcon != "own" {
		t.Fatalf("own icon replaced: %q", got.FooterIcon)
	}
}
