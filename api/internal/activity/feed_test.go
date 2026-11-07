package activity

import (
	"context"
	"testing"
	"time"
)

func read(t *testing.T, s *Store, q Query) Page {
	t.Helper()
	if q.Now.IsZero() {
		q.Now = testNow
	}
	if q.GuildIDs == nil {
		q.GuildIDs = []int64{}
	}
	p, err := s.Read(context.Background(), q)
	if err != nil {
		t.Fatal(err)
	}
	return p
}

func texts(p Page) []string {
	out := []string{}
	for _, r := range p.Rows {
		out = append(out, r.Segments[0].Text)
	}
	return out
}

// A guild's rows reach its own verified members only; nobody signed out ever reads one.
func TestFeedKeepsGuildRowsInTheGuild(t *testing.T) {
	s := testStore(t)
	g := guild(t, s, "Olympus XXVII")
	other := guild(t, s, "Other")
	mustAppend(t, s, ev(KindLogUpload, "public", testNow.Add(-time.Minute)))
	own := ev(KindGuildRaid, "own guild", testNow.Add(-2*time.Minute))
	own.GuildID = g
	mustAppend(t, s, own)
	theirs := ev(KindGuildRaid, "their guild", testNow.Add(-3*time.Minute))
	theirs.GuildID = other
	mustAppend(t, s, theirs)

	if got := texts(read(t, s, Query{Filter: FilterAll})); len(got) != 1 || got[0] != "public" {
		t.Fatalf("signed out: %v", got)
	}
	got := read(t, s, Query{Filter: FilterAll, GuildIDs: []int64{g}})
	if len(got.Rows) != 2 || got.Rows[0].Tag != TagYourGuild || got.Rows[1].Tag != TagLogs {
		t.Fatalf("member: %+v", got.Rows)
	}
	if g := texts(read(t, s, Query{Filter: FilterGuild, GuildIDs: []int64{g}})); len(g) != 1 || g[0] != "own guild" {
		t.Fatalf("guild filter: %v", g)
	}
}

// Your spec is the spec and the band; a matching row is tagged so and sorts first in its hour.
func TestFeedSpecFilterAndOwnRowsFirst(t *testing.T) {
	s := testStore(t)
	mustAppend(t, s, ev(KindLogUpload, "newest, not mine", testNow.Add(-time.Minute)))
	mine := ev(KindBisChange, "mine", testNow.Add(-10*time.Minute))
	mine.Spec, mine.Band = "warrior-fury", 20
	mustAppend(t, s, mine)
	other := ev(KindBisChange, "fury at 30", testNow.Add(-20*time.Minute))
	other.Spec, other.Band = "warrior-fury", 30
	mustAppend(t, s, other)
	mustAppend(t, s, ev(KindLogUpload, "last hour", testNow.Add(-90*time.Minute)))

	all := read(t, s, Query{Filter: FilterAll, Spec: "warrior-fury", Band: 20})
	if got := texts(all); got[0] != "mine" || got[1] != "newest, not mine" || got[3] != "last hour" {
		t.Fatalf("order: %v", got)
	}
	if all.Rows[0].Tag != TagYourSpec || all.Rows[2].Tag != TagBestInSlot {
		t.Fatalf("tags: %s %s", all.Rows[0].Tag, all.Rows[2].Tag)
	}
	if got := texts(read(t, s, Query{Filter: FilterSpec, Spec: "warrior-fury", Band: 20})); len(got) != 1 || got[0] != "mine" {
		t.Fatalf("spec filter: %v", got)
	}
	if got := texts(read(t, s, Query{Filter: FilterLogs})); len(got) != 2 {
		t.Fatalf("logs filter: %v", got)
	}
}

// Pages follow on from their cursor with no row twice; the feed reads 7 days back; a dedupe
// key writes once.
func TestFeedPagesWindowAndDedupe(t *testing.T) {
	s := testStore(t)
	for i := 0; i < 5; i++ {
		mustAppend(t, s, ev(KindLogUpload, string(rune('a'+i)), testNow.Add(-time.Duration(i)*time.Hour)))
	}
	mustAppend(t, s, ev(KindLogUpload, "old", testNow.Add(-8*24*time.Hour)))
	if fresh, _ := s.Append(context.Background(), ev(KindLogUpload, "a", testNow)); fresh {
		t.Fatal("a dedupe key wrote twice")
	}
	first := read(t, s, Query{Filter: FilterAll, Limit: 3})
	if len(first.Rows) != 3 || first.NextCursor == nil {
		t.Fatalf("first page: %v %v", texts(first), first.NextCursor)
	}
	second := read(t, s, Query{Filter: FilterAll, Limit: 3, Cursor: *first.NextCursor})
	if got := texts(second); len(got) != 2 || got[0] != "d" || second.NextCursor != nil {
		t.Fatalf("second page: %v", got)
	}
	if _, err := s.Read(context.Background(), Query{Filter: FilterAll, Cursor: "nope", Now: testNow, GuildIDs: []int64{}}); err != ErrBadCursor {
		t.Fatalf("bad cursor: %v", err)
	}
	if n, err := s.Purge(context.Background(), testNow); err != nil || n != 0 {
		t.Fatalf("purge took a row inside 30 days: %d %v", n, err)
	}
	if n, err := s.Purge(context.Background(), testNow.Add(23*24*time.Hour)); err != nil || n != 1 {
		t.Fatalf("purge of the 31-day-old row: %d %v", n, err)
	}
}
