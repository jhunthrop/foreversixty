package guilds

import (
	"context"
	"testing"
)

// A no-account entry whose snapshot line names no class and a spec two classes play ("Holy")
// takes the class from that character's export, so the roster draws the crest and "Holy
// Paladin" (guild round 3, handed over by the web lane). A line the snapshot already places, or
// a name with no export, is left as it was.
func TestSnapshotLinesTakeTheExportsClass(t *testing.T) {
	pool := testPool(t)
	s := &Store{Pool: pool}
	ctx := context.Background()
	uid := seedUser(t, pool, "removed@example.com")
	if _, err := pool.Exec(ctx, `
		insert into addon_exports (character_key, user_id, region, ruleset, name, export, captured_at, updated_at)
		values ('us/hardcore/ashenveil', $1, 'us', 'hardcore', 'Ashenveil', 'FS1:1.60.1.70790:paladin:human:0/0/0:', now(), now())`,
		uid); err != nil {
		t.Fatal(err)
	}
	g := Guild{Region: "us", Ruleset: "hardcore"}
	lines := []snapshotMember{
		{Name: "Ashenveil-Realm", Spec: "Holy"},
		{Name: "Strayhound", Spec: "Protection"},
		{Name: "Oldmarrow", Spec: "warrior-arms"},
	}
	got, err := s.withExportClasses(ctx, g, lines)
	if err != nil {
		t.Fatal(err)
	}
	if c := snapshotClass(got[0]); c == nil || *c != "paladin" {
		t.Fatalf("Ashenveil's class = %v, want paladin from the export", c)
	}
	if got[1].Class != "" {
		t.Fatalf("Strayhound has no export, class = %q, want none", got[1].Class)
	}
	if c := snapshotClass(got[2]); c == nil || *c != "warrior" || got[2].Class != "" {
		t.Fatalf("Oldmarrow's spec already places a warrior; line = %+v", got[2])
	}
	if lines[0].Class != "" {
		t.Fatal("the input lines were changed in place")
	}
}
