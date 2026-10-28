package coach

import "testing"

func TestClockAndSpan(t *testing.T) {
	cases := []struct{ name, got, want string }{
		{"Clock", Clock(185_000), "3:05"},
		{"Clock at the pull", Clock(0), "0:00"},
		{"Clock drops milliseconds", Clock(59_999), "0:59"},
		{"Clock before the pull", Clock(-500), "0:00"},
		{"Span", Span(72_000, 85_000), "1:12–1:25"},
		{"Thousands", Thousands(2140), "2,140"},
		{"Thousands small", Thousands(999), "999"},
		{"Thousands millions", Thousands(1234567), "1,234,567"},
		{"Thousands negative", Thousands(-2140), "-2,140"},
		{"Pct", Pct(44.4), "44%"},
		{"Pct rounds", Pct(44.5), "45%"},
		{"PerMinute", PerMinute(9.06), "9.1"},
		{"PerMinute whole", PerMinute(8), "8.0"},
		{"ClassPlural", ClassPlural("Warrior"), "warriors"},
		{"ClassPlural lower", ClassPlural("priest"), "priests"},
	}
	for _, c := range cases {
		if c.got != c.want {
			t.Errorf("%s = %q, want %q", c.name, c.got, c.want)
		}
	}
}
