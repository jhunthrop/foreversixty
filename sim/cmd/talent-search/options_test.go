package main

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestParseOptionsDefaultsAndOverrides(t *testing.T) {
	o, err := parseOptions([]string{"-spec", "mage-fire", "-repo-root", "/r", "-keep", "A,B", "-level", "40", "-top", "4"})
	if err != nil {
		t.Fatal(err)
	}
	if o.spec != "mage-fire" || o.level != 40 || o.top != 4 || o.keep != "A,B" {
		t.Fatalf("got %+v", o)
	}
	if o.out != filepath.Join("/r", "design", "reviews", "talent-search") {
		t.Errorf("default out = %q", o.out)
	}
	if o.faction != "alliance" || o.seed != 7 || o.limit != 60 || o.refineTop != 3 {
		t.Errorf("defaults changed: %+v", o)
	}
	custom, _ := parseOptions([]string{"-spec", "x", "-out", "/elsewhere"})
	if custom.out != "/elsewhere" {
		t.Errorf("explicit -out ignored: %q", custom.out)
	}
}

func TestParseOptionsRejectsBadInput(t *testing.T) {
	for name, args := range map[string][]string{
		"missing spec": {},
		"level low":    {"-spec", "x", "-level", "9"},
		"level high":   {"-spec", "x", "-level", "61"},
		"unknown flag": {"-spec", "x", "-nope"},
	} {
		if _, err := parseOptions(args); err == nil {
			t.Errorf("%s: accepted", name)
		}
	}
	for _, level := range []string{"10", "60"} {
		if _, err := parseOptions([]string{"-spec", "x", "-level", level}); err != nil {
			t.Errorf("level %s rejected: %v", level, err)
		}
	}
}

func TestRunPropagatesOptionErrors(t *testing.T) {
	if err := run([]string{"-level", "5"}); err == nil || !strings.Contains(err.Error(), "-spec") {
		t.Fatalf("got %v", err)
	}
}

func TestPrepareReadsARealDPSSpecEndToEnd(t *testing.T) {
	o := options{repoRoot: repoRoot, spec: "mage-fire", level: 60, faction: "alliance", preset: "bare", seed: 7}
	in, err := prepare(o)
	if err != nil {
		t.Fatal(err)
	}
	if in.clientBuild == "" || len(in.guide) == 0 || in.guideCode == "" || len(in.modeled) == 0 {
		t.Fatalf("incomplete inputs: %+v", in)
	}
	if in.setup.spec.Spec != "mage-fire" || in.setup.band.Faction != "alliance" || in.setup.band.Band != 60 {
		t.Fatalf("setup = %+v", in.setup)
	}
	if len(in.keep) != 0 {
		t.Fatalf("no -keep given but keep = %v", in.keep)
	}

	var name string
	for id := range in.guide {
		name = in.setup.trees.byID[id].Name
		break
	}
	o.keep = name
	kept, err := prepare(o)
	if err != nil || len(kept.keep) != 1 {
		t.Fatalf("keep %q: %v %v", name, kept.keep, err)
	}
	o.keep = "No Such Talent"
	if _, err := prepare(o); err == nil || !strings.Contains(err.Error(), "-keep") {
		t.Fatalf("typo in -keep must fail, got %v", err)
	}
	o.keep = ""
	o.faction = "neutral"
	if _, err := prepare(o); err == nil {
		t.Fatal("a faction with no band must fail")
	}
}

func TestPrepareFailsOnAnUnknownSpec(t *testing.T) {
	if _, err := prepare(options{repoRoot: repoRoot, spec: "no-such-spec", level: 60}); err == nil {
		t.Fatal("unknown spec accepted")
	}
}
