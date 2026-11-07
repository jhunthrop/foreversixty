package activity

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

// Ruling P4's thresholds: a climb of at least climbPlaces ending in the top climbCeiling,
// or any move into the top podium places.
const (
	climbPlaces  = 5
	climbCeiling = 10
	podium       = 3
)

// rankingEvents compares a board with its previous ranks (character key to rank) and
// answers its rows: a new #1, and the climbs worth a row. A board seen for the first time
// (no previous ranks) answers none: there is no movement to tell.
func rankingEvents(id boardID, board []BoardRow, prev map[string]int, now time.Time) []Event {
	if len(prev) == 0 || len(board) == 0 {
		return nil
	}
	var out []Event
	top := board[0]
	if was, ok := prev[top.Key]; !ok || was != 1 {
		out = append(out, rankEvent(id, top, KindRankingTop, 0, now))
	}
	for _, r := range board[1:] {
		was, ok := prev[r.Key]
		if !ok || was <= r.Rank {
			continue
		}
		gain := was - r.Rank
		if (gain >= climbPlaces && r.Rank <= climbCeiling) || r.Rank <= podium {
			out = append(out, rankEvent(id, r, KindRankingClimb, gain, now))
		}
	}
	return out
}

func rankEvent(id boardID, r BoardRow, kind string, gain int, now time.Time) Event {
	board := fmt.Sprintf("%s, %s", SpecName(id.spec), BandWords(id.band))
	segments := []Segment{Name(r.Name, r.Class), Text(" took #1 for " + board)}
	dedupe := fmt.Sprintf("rank-top:%s:%d:%s:%s", id.spec, id.band, r.Key, WeekOf(now).Format("20060102"))
	if kind == KindRankingClimb {
		segments = []Segment{Strong(board + ": "), Name(r.Name, r.Class),
			Text(fmt.Sprintf(" climbed %d %s to ", gain, plural(gain, "place", "places"))), Strong("#" + strconv.Itoa(r.Rank))}
		dedupe = fmt.Sprintf("rank-climb:%s:%d:%s:%d:%s", id.spec, id.band, r.Key, r.Rank, now.UTC().Format("2006010215"))
	}
	return Event{
		Kind: kind, OccurredAt: now, Spec: id.spec, Band: id.band, CharacterKey: r.Key,
		Crest: r.Class, Href: "/character/" + r.Key, Segments: segments, DedupeKey: dedupe,
	}
}

// RefreshBoards recomputes every spec board, writes the ranking rows its movement since
// the last refresh earns, and stores it as this week's board.
func (s *Store) RefreshBoards(ctx context.Context, zoneBands map[string]int, now time.Time) (int, error) {
	week := WeekOf(now)
	prev, err := s.storedRanks(ctx, week)
	if err != nil {
		return 0, err
	}
	if len(prev) == 0 {
		if prev, err = s.storedRanks(ctx, PreviousWeek(now)); err != nil {
			return 0, err
		}
	}
	boards, err := s.computeBoards(ctx, zoneBands)
	if err != nil {
		return 0, err
	}
	written := 0
	for id, b := range boards {
		for _, e := range rankingEvents(id, b, prev[id], now) {
			fresh, err := s.Append(ctx, e)
			if err != nil {
				return written, err
			}
			if fresh {
				written++
			}
		}
	}
	return written, s.replaceWeek(ctx, week, boards)
}
