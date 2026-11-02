package coach

import (
	"context"
	"slices"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// Round 11: the trainables file lists the Ahn'Qiraj book ranks beside the trainer's
// (Revenge rank 6, 25288; Shadow Bolt rank 10). No site rotation casts one, so F5 called
// every level-60 player following the site's rotation "low rank". The coach's ranks drop
// every id spellranks.json flags as a book.
func TestTablesFor_TheTrainableRanksLeaveOutTheBookRanks(t *testing.T) {
	ft := &FileTables{BuildsDir: repoBuilds, CoachDir: t.TempDir()}
	row := summary.RosterRow{Class: "Warrior", Spec: "Protection", Role: "tank", Level: 60}
	tb, err := ft.TablesFor(context.Background(), row, 0, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	revenge := tb.Trainable["Revenge"]
	if len(revenge) == 0 {
		t.Fatal("no Revenge ranks")
	}
	if slices.ContainsFunc(revenge, func(r TrainableRank) bool { return r.ID == 25288 }) {
		t.Errorf("Revenge keeps the book rank 25288: %v", revenge)
	}
	if top := revenge[len(revenge)-1]; top.ID != 11601 || top.Rank != 5 {
		t.Errorf("Revenge's top rank is %+v, want the trainer's rank 5 (11601)", top)
	}
}

func TestRanks_ABookRankIsNeverTheRankF5AsksFor(t *testing.T) {
	in := protSunderRank(t, 60, 4, 22)
	in.Tables.Trainable = withoutBookRanks(in.Tables.Trainable, map[int64]bool{11597: true})
	if o := ranks(in, CountedTime(in)); o.Finding != nil {
		t.Fatalf("rank 4 is the top once rank 5 is a book: %+v", o.Finding)
	}
}

// The rotation the site ships names the rank it casts at the band (a raid-60 Lightning
// Bolt rank 9 beside a trainable rank 10); a player casting that exact rank is following
// the site, so F5 never calls it low.
func TestRanks_TheRankTheBandsRotationNamesIsNeverFlagged(t *testing.T) {
	in := protSunderRank(t, 60, 4, 22)
	in.Tables.Rotation = in.Tables.Rotation.Naming(sunderRank4ID)
	if o := ranks(in, CountedTime(in)); o.Finding != nil {
		t.Fatalf("the band's own rank was flagged: %+v", o.Finding)
	}
	in = protSunderRank(t, 60, 3, 22)
	in.Tables.Rotation = in.Tables.Rotation.Naming(sunderRank4ID)
	if o := ranks(in, CountedTime(in)); o.Finding == nil {
		t.Fatal("a rank below the band's own must still fire")
	}
}
