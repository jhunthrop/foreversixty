package main

import "testing"

func TestReportNameKeepsTheBareFileAndSuffixesOtherPresets(t *testing.T) {
	if got := reportName(options{spec: "x", preset: "bare"}); got != "x.md" {
		t.Errorf("bare = %q, want x.md", got)
	}
	if got := reportName(options{spec: "x", preset: "raid"}); got != "x-raid.md" {
		t.Errorf("raid = %q, want x-raid.md", got)
	}
}

func TestParseOptionsDefaultsToTheBarePreset(t *testing.T) {
	o, err := parseOptions([]string{"-spec", "x"})
	if err != nil || o.preset != "bare" {
		t.Fatalf("preset = %q, err = %v; want bare", o.preset, err)
	}
}
