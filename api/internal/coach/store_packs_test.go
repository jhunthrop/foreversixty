package coach

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPackStore_RefusesAPackWithoutAVersion(t *testing.T) {
	f := newT30(t)
	p := examplePack()
	p.Version = ""
	if err := f.svc.Store.SavePack(context.Background(), t30Report, t30Fight, f.furyKey, p); err == nil {
		t.Fatal("an empty version is refused")
	}
}

func TestPackStore_RoundTrip(t *testing.T) {
	f := newT30(t)
	ctx := context.Background()
	p := examplePack()
	p.Version = "pack-v1.t30"
	if err := f.svc.Store.SavePack(ctx, t30Report, t30Fight, f.furyKey, p); err != nil {
		t.Fatal(err)
	}
	got, ok, err := f.svc.Store.Pack(ctx, t30Report, t30Fight, f.furyKey, p.Version)
	if err != nil || !ok || got.Version != p.Version || len(got.Findings) != len(p.Findings) {
		t.Fatalf("%+v %v %v", got, ok, err)
	}
	if _, ok, _ := f.svc.Store.Pack(ctx, t30Report, t30Fight, f.furyKey, "pack-v1.other"); ok {
		t.Fatal("another version is a miss")
	}
}

func TestPackFor_CachedByVersion(t *testing.T) {
	f := newT30(t)
	ctx := context.Background()
	first, err := f.svc.PackFor(ctx, f.rep(t), t30Fight, f.furyKey)
	if err != nil || first.Status != StatusReady || first.Version == "" || len(first.Findings) == 0 {
		t.Fatalf("%+v %v", first, err)
	}
	reads := f.bucket.reads()
	again, err := f.svc.PackFor(ctx, f.rep(t), t30Fight, f.furyKey)
	if err != nil || again.Version != first.Version || f.bucket.reads() != reads {
		t.Fatalf("the second call reads coach_packs: reads %d -> %d, %v", reads, f.bucket.reads(), err)
	}
	dir := t.TempDir()
	changed := strings.Replace(validRules, `"build": "1.60.1.70630"`, `"build": "1.60.1.70631"`, 1)
	if err := os.WriteFile(filepath.Join(dir, "warrior-fury.json"), []byte(changed), 0o600); err != nil {
		t.Fatal(err)
	}
	moved := f.service(t, dir, false)
	rebuilt, err := moved.PackFor(ctx, f.rep(t), t30Fight, f.furyKey)
	if err != nil || rebuilt.Version == first.Version || f.bucket.reads() != reads+1 {
		t.Fatalf("a changed rule file rebuilds: %s -> %s, reads %d -> %d, %v", first.Version, rebuilt.Version,
			reads, f.bucket.reads(), err)
	}
	if _, cached, _ := moved.Store.Pack(ctx, t30Report, t30Fight, f.furyKey, rebuilt.Version); !cached {
		t.Fatal("the rebuilt pack is cached under its version")
	}
}

func TestPackFor_PreparingIsNotCached(t *testing.T) {
	f := newT30(t)
	svc := f.service(t, coachDirWith(t), true) // a validated DPS spec whose sim reference is not in yet
	ctx := context.Background()
	p, err := svc.PackFor(ctx, f.rep(t), 1, f.furyKey)
	_, cached, _ := svc.Store.Pack(ctx, t30Report, 1, f.furyKey, p.Version)
	if err != nil || p.Status != StatusPreparing || cached {
		t.Fatalf("status %s cached %v err %v", p.Status, cached, err)
	}
	f.mux = http.NewServeMux()
	Mount(f.mux, svc)
	_, _, body := f.get(t, user(f.premium), "")
	if body["status"] != "preparing" || len(f.jobs.Ran()) != 0 {
		t.Fatalf("%v jobs %v", body, f.jobs.Ran())
	}
	_, _, body = f.get(t, user(f.freeOwner), "")
	if body["status"] != "preparing" || body["finding_count"] != float64(0) {
		t.Fatalf("the teaser says preparing too: %v", body)
	}
}

func TestPackFor_DebuffAurasAreFlagged(t *testing.T) {
	at := int64(1000)
	p := Pack{Anchors: []Anchor{{ID: "a1", Kind: AnchorAura, SpellID: 7386, AtMS: &at}, {ID: "a2", Kind: AnchorAura, SpellID: 6673},
		{ID: "a3", Kind: AnchorCast, SpellID: 7386}}}
	rules := SpecRules{Auras: []AuraRule{{SpellID: 7386, Kind: AuraDebuff}, {SpellID: 6673, Kind: AuraSelfBuff}}}
	got := anchorsOf(p, rules)
	if !got["a1"].Debuff || got["a2"].Debuff || got["a3"].Debuff || *got["a1"].AtMS != at {
		t.Fatalf("%+v", got)
	}
}
