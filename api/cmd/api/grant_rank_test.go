// api/cmd/api/grant_rank_test.go
package main

import (
	"testing"
	"time"
)

func TestParseRankFlagsRequiresGuildAndUser(t *testing.T) {
	if _, err := parseRankFlags([]string{"--guild", "3"}); err == nil {
		t.Fatal("missing --user should refuse")
	}
	if _, err := parseRankFlags([]string{"--user", "78"}); err == nil {
		t.Fatal("missing --guild should refuse")
	}
	if _, err := parseRankFlags([]string{"--guild", "3", "--user", "78", "--bogus", "x"}); err == nil {
		t.Fatal("an unknown flag should refuse")
	}
	if _, err := parseRankFlags([]string{"--guild", "3", "--user"}); err == nil {
		t.Fatal("a flag without a value should refuse")
	}
}

func TestGrantRankOverrideValidation(t *testing.T) {
	now := time.Date(2026, 11, 10, 12, 0, 0, 0, time.UTC)
	for _, tc := range []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{"happy", []string{"--rank", "officer", "--until", "2026-11-24", "--note", "testing"}, false},
		{"leader", []string{"--rank", "leader", "--until", "2026-11-24", "--note", "testing"}, false},
		{"bad rank", []string{"--rank", "member", "--until", "2026-11-24", "--note", "testing"}, true},
		{"no until", []string{"--rank", "officer", "--note", "testing"}, true},
		{"past until", []string{"--rank", "officer", "--until", "2026-11-10", "--note", "testing"}, true},
		{"bad until", []string{"--rank", "officer", "--until", "soon", "--note", "testing"}, true},
		{"no note", []string{"--rank", "officer", "--until", "2026-11-24"}, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f, err := parseRankFlags(append([]string{"--guild", "3", "--user", "78"}, tc.args...))
			if err != nil {
				t.Fatal(err)
			}
			o, err := f.grantRankOverride(3, 78, "op", now)
			if (err != nil) != tc.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tc.wantErr)
			}
			want := time.Date(2026, 11, 24, 0, 0, 0, 0, time.UTC)
			if err == nil && (o.GuildID != 3 || o.UserID != 78 || o.CreatedBy != "op" || !o.ExpiresAt.Equal(want)) {
				t.Fatalf("override = %+v", o)
			}
		})
	}
}
