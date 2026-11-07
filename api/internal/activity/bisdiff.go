package activity

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/sim/specs"
)

// pickKey names one pick in bis_pick_state.
type pickKey struct {
	band                  int
	faction, preset, slot string
}

// bandKey is one row of the feed: a spec's band under one preset, both factions together.
type bandKey struct {
	band   int
	preset string
}

type pickChange struct {
	slot, name string
	itemID     int
}

// DiffBis compares the deployed BiS files of build with the last seen picks and writes one
// bis_change row per spec, band and preset that changed (ruling P5), then stores the new
// picks. A spec seen for the first time is seeded and writes no row. It answers how many
// rows it wrote.
func (s *Store) DiffBis(ctx context.Context, dataDir, build string) (int, error) {
	files, err := loadBisFiles(dataDir, build)
	if err != nil {
		return 0, err
	}
	items := &itemTable{dataDir: dataDir, build: build}
	written := 0
	for _, f := range files {
		n, err := s.diffOne(ctx, f, items)
		if err != nil {
			return written, err
		}
		written += n
	}
	return written, nil
}

func (s *Store) diffOne(ctx context.Context, f bisFile, items *itemTable) (int, error) {
	seen, err := s.picksOf(ctx, f.Spec)
	if err != nil {
		return 0, err
	}
	changes := map[bandKey][]pickChange{}
	var order []bandKey
	for _, b := range sortedBands(f.Bands) {
		k := bandKey{b.Band, b.preset()}
		for _, sl := range b.Slots {
			old, had := seen[pickKey{b.Band, b.Faction, b.preset(), sl.Slot}]
			if len(seen) == 0 || sl.ItemID <= 0 || (had && old == sl.ItemID) || hasChange(changes[k], sl.Slot) {
				continue
			}
			if len(changes[k]) == 0 {
				order = append(order, k)
			}
			changes[k] = append(changes[k], pickChange{slot: sl.Slot, name: sl.ItemName, itemID: sl.ItemID})
		}
	}
	written := 0
	for _, k := range order {
		if _, err := s.Append(ctx, bisEvent(f, k, changes[k], items)); err != nil {
			return written, err
		}
		written++
	}
	return written, s.storePicks(ctx, f)
}

func hasChange(cs []pickChange, slot string) bool {
	for _, c := range cs {
		if c.slot == slot {
			return true
		}
	}
	return false
}

func (s *Store) picksOf(ctx context.Context, spec string) (map[pickKey]int, error) {
	rows, err := s.Pool.Query(ctx, `select band, faction, preset, slot, item_id from bis_pick_state where spec = $1`, spec)
	if err != nil {
		return nil, fmt.Errorf("activity: read picks %s: %w", spec, err)
	}
	defer rows.Close()
	out := map[pickKey]int{}
	for rows.Next() {
		var k pickKey
		var id int
		if err := rows.Scan(&k.band, &k.faction, &k.preset, &k.slot, &id); err != nil {
			return nil, fmt.Errorf("activity: scan picks %s: %w", spec, err)
		}
		out[k] = id
	}
	return out, rows.Err()
}

func (s *Store) storePicks(ctx context.Context, f bisFile) error {
	batch := &pgx.Batch{}
	for _, b := range f.Bands {
		for _, sl := range b.Slots {
			if sl.ItemID <= 0 {
				continue
			}
			batch.Queue(`insert into bis_pick_state (spec, band, faction, preset, slot, item_id) values ($1, $2, $3, $4, $5, $6)
				on conflict (spec, band, faction, preset, slot) do update set item_id = excluded.item_id, updated_at = now()
				where bis_pick_state.item_id <> excluded.item_id`, f.Spec, b.Band, b.Faction, b.preset(), sl.Slot, sl.ItemID)
		}
	}
	if err := s.Pool.SendBatch(ctx, batch).Close(); err != nil {
		return fmt.Errorf("activity: store picks %s: %w", f.Spec, err)
	}
	return nil
}

func bisEvent(f bisFile, k bandKey, cs []pickChange, items *itemTable) Event {
	spec := specs.ByKey[f.Spec]
	words := BandWords(k.band)
	if k.preset != barePreset {
		words += " " + k.preset
	}
	first := cs[0]
	it, _ := items.lookup(spec.ClassSlug, first.itemID)
	segments := []Segment{Strong("Best in slot refreshed"), Text(fmt.Sprintf(" for %s %s: %d new picks", SpecName(f.Spec), words, len(cs)))}
	if len(cs) == 1 {
		segments = []Segment{Strong("Best in slot changed"), Text(fmt.Sprintf(" for %s %s: ", SpecName(f.Spec), words)),
			Item(first.name, first.itemID, it.Quality), Text(" is the new " + slotWords(first.slot))}
	}
	var quality *int
	if it.Icon != "" {
		q := it.Quality
		quality = &q
	}
	return Event{
		Kind: KindBisChange, OccurredAt: f.GeneratedAt, Spec: f.Spec, Band: k.band, TagBase: TagGuides,
		Icon: it.Icon, Quality: quality, Href: "/bis/" + spec.ClassSlug + "/" + spec.SpecSlug,
		Segments:  segments,
		DedupeKey: strings.Join([]string{"bis", f.Spec, strconv.Itoa(k.band), k.preset, f.GeneratedAt.UTC().Format("20060102T150405")}, ":"),
	}
}
