// api/cmd/api/grant_test.go
package main

import "testing"

func TestParseGrantFlagsRequiresExactlyOneOfUserOrGuild(t *testing.T) {
	_, err := parseGrantFlags([]string{"--plan", "premium", "--note", "x"})
	if err == nil {
		t.Fatal("neither --user nor --guild should refuse")
	}
	_, err = parseGrantFlags([]string{"--user", "42", "--guild", "9", "--plan", "premium", "--note", "x"})
	if err == nil {
		t.Fatal("both --user and --guild should refuse")
	}
}

func TestParseGrantFlagsRequiresPlanAndNote(t *testing.T) {
	if _, err := parseGrantFlags([]string{"--user", "42"}); err == nil {
		t.Fatal("missing --plan/--note should refuse")
	}
	if _, err := parseGrantFlags([]string{"--user", "42", "--plan", "gold", "--note", "x"}); err == nil {
		t.Fatal("an invalid --plan should refuse")
	}
}

func TestParseGrantFlagsHappyPath(t *testing.T) {
	f, err := parseGrantFlags([]string{"--user", "somebattletag#1234", "--plan", "premium", "--until", "2026-12-31", "--note", "beta tester"})
	if err != nil {
		t.Fatal(err)
	}
	if f.user != "somebattletag#1234" || f.plan != "premium" || f.until != "2026-12-31" || f.note != "beta tester" {
		t.Fatalf("flags = %+v", f)
	}
}

func TestParseUntilAcceptsISODateOrEmpty(t *testing.T) {
	if until, err := parseUntil(""); err != nil || until != nil {
		t.Fatalf("empty: %v, %v", until, err)
	}
	until, err := parseUntil("2026-12-31")
	if err != nil || until == nil || until.Year() != 2026 || until.Month() != 12 || until.Day() != 31 {
		t.Fatalf("2026-12-31: %v, %v", until, err)
	}
	if _, err := parseUntil("not-a-date"); err == nil {
		t.Fatal("a garbage date should refuse")
	}
}

func TestParseGrantFlagsFeatures(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr string
	}{
		{"without until", []string{"--guild", "3", "--features", "officer_views", "--note", "x"}, "--until is required with --features"},
		{"unknown feature", []string{"--guild", "3", "--features", "roster_check", "--until", "2026-12-09", "--note", "x"}, "--features must be officer_views"},
		{"with a user", []string{"--user", "3", "--features", "officer_views", "--until", "2026-12-09", "--note", "x"}, "--features requires --guild"},
		{"premium plan", []string{"--guild", "3", "--plan", "premium", "--features", "officer_views", "--until", "2026-12-09", "--note", "x"}, "--features requires --plan guild"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := parseGrantFlags(tc.args)
			if err == nil || err.Error() != tc.wantErr {
				t.Fatalf("err = %v, want %q", err, tc.wantErr)
			}
		})
	}
}

func TestParseGrantFlagsFeaturesDefaultsPlanToGuild(t *testing.T) {
	f, err := parseGrantFlags([]string{"--guild", "3", "--features", "officer_views", "--until", "2026-12-09", "--note", "x"})
	if err != nil {
		t.Fatal(err)
	}
	if f.plan != "guild" || f.guild != "3" || f.features != "officer_views" {
		t.Fatalf("flags = %+v", f)
	}
}
