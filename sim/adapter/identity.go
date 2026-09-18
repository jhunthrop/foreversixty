package adapter

// Row identity.
//
// Every row of a summary is keyed by an integer id, not by its name:
// logs/engine/summary keys abilities by {spell_id, via}, casts by
// {guid, spell_id} and auras by {target_guid, spell_id}, and the report
// components key their {#each} blocks on the same tuples, where a
// repeated key is a runtime error in Svelte 5. An engine ActionID is not
// an integer, though - it is a spell, an item or an "other" action, each
// with a tag and a rank - so several engine actions can share one spell
// id (the three tags of a white swing) or one raw number (item 5513 and
// spell 5513 are different things).
//
// So an untagged spell keeps its own id, which is what the web resolves
// a name with from the build's spells.json, and everything else gets a
// derived id from the reserved space below. The name always carries the
// client id the web would look up, so nothing is lost by the derivation.
//
// The reserved space cannot collide with a client spell id: the largest
// id the client carries today is just above 1,000,000 (Forever's own new
// objects), every derived id is at least syntheticBase, and Summarize
// checks its own output for duplicate keys before returning it, so a
// future id that broke this arithmetic would fail loudly rather than
// reach the report.

import (
	"fmt"
	"strings"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/wowsims/classic/sim/core/proto"
)

// The derived id is base + variant block + kind block + the raw number:
//
//	id = syntheticBase + variant*variantStride + kind*kindStride + raw
//	variant = tag*rankStride + rank
//
// with raw < kindStride, kind < 4, rank < rankStride. Within those
// bounds the packing is injective, which TestDerivedIDsAreInjective
// pins; outside them Summarize's duplicate check is the backstop.
const (
	// syntheticBase is the first id the adapter allocates for itself.
	// Client spell ids are below it.
	syntheticBase = 2_000_000
	// kindStride is the width of one kind's block: five times the
	// largest id a client carries.
	kindStride = 10_000_000
	// variantStride is the width of one (tag, rank) variant: the whole
	// kind space over again.
	variantStride = 100_000_000
	// rankStride is the width of one rank inside a tag.
	rankStride = 100
)

// The kinds of engine action, in the order their blocks sit in the
// reserved space.
const (
	kindSpell = iota
	kindItem
	kindOther
	kindUnknown
)

// ActionName returns the row id and display name for an engine ActionID.
//
// The id is the action's own spell id when it is a plain untagged spell,
// and a derived id from the reserved space otherwise, so that no two
// engine actions ever produce the same summary row key. The name is the
// id in a readable form - the engine's ActionIDs carry no names - and
// always spells out the client id, the tag and the rank, so the web can
// resolve a real name from the build's own spells.json, which it already
// loads for tooltips. Forever keeps vanilla ids for abilities that
// already existed and uses ids above 1,000,000 only for new objects, so
// no translation table is needed on either side.
func ActionName(id *proto.ActionID) (int64, string) {
	if id == nil {
		return derivedID(kindUnknown, 0, 0, 0), "unknown"
	}
	suffix := variantSuffix(id.Tag, id.Rank)
	switch raw := id.RawId.(type) {
	case *proto.ActionID_SpellId:
		if suffix == "" {
			return int64(raw.SpellId), fmt.Sprintf("spell:%d", raw.SpellId)
		}
		return derivedID(kindSpell, int64(raw.SpellId), id.Tag, id.Rank),
			fmt.Sprintf("spell:%d%s", raw.SpellId, suffix)
	case *proto.ActionID_ItemId:
		// An item id and a spell id are different numbers in the same
		// range, so an item never keeps its raw number as a row id.
		return derivedID(kindItem, int64(raw.ItemId), id.Tag, id.Rank),
			fmt.Sprintf("item:%d%s", raw.ItemId, suffix)
	case *proto.ActionID_OtherId:
		return derivedID(kindOther, int64(raw.OtherId), id.Tag, id.Rank),
			fmt.Sprintf("other:%s%s", otherActionName(raw.OtherId), suffix)
	}
	return derivedID(kindUnknown, 0, id.Tag, id.Rank), "unknown" + suffix
}

// derivedID packs an action into the reserved space. See the block
// comment above for the layout.
func derivedID(kind int, raw int64, tag, rank int32) int64 {
	variant := int64(tag)*rankStride + int64(rank)
	return syntheticBase + variant*variantStride + int64(kind)*kindStride + raw
}

// variantSuffix names a tag and a rank, and is empty for the plain
// action that has neither.
func variantSuffix(tag, rank int32) string {
	var b strings.Builder
	if tag != 0 {
		fmt.Fprintf(&b, "/%d", tag)
	}
	if rank != 0 {
		fmt.Fprintf(&b, "+r%d", rank)
	}
	return b.String()
}

// otherActionName is the engine's own name for a non-spell action, in
// lower snake case and without its enum prefix: OtherActionAttack reads
// "attack". A value the engine has not named reads as its number, so a
// new one is still legible rather than blank.
func otherActionName(id proto.OtherAction) string {
	name, ok := proto.OtherAction_name[int32(id)]
	if !ok {
		return fmt.Sprintf("%d", int32(id))
	}
	return snake(strings.TrimPrefix(name, "OtherAction"))
}

// snake turns an upper-camel protobuf name into lower snake case.
func snake(name string) string {
	var b strings.Builder
	b.Grow(len(name) + 4)
	for i, r := range name {
		if r >= 'A' && r <= 'Z' {
			if i > 0 {
				b.WriteByte('_')
			}
			b.WriteRune(r + ('a' - 'A'))
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}

// rowKey is one summary row's identity, as the logs engine's own maps
// and the report's keyed blocks compute it: an ability by its actor,
// spell and pet (summary/damage.go's abilityKey, ActorRow.svelte's
// abilityKey), a cast by its caster and spell (CastTable.svelte), an
// aura by its target and spell (AuraTable.svelte).
type rowKey struct {
	table string
	scope string
	id    int64
}

func (k rowKey) String() string { return fmt.Sprintf("%s row (%s, %d)", k.table, k.scope, k.id) }

// checkRowIdentity reports the first duplicate row key in a summary. The
// logs engine's maps make a duplicate impossible for a real fight, and
// the report's keyed {#each} blocks raise a runtime error on one, so a
// sim that produced one would break the page it exists to feed. Rather
// than trust the arithmetic above to stay injective forever, Summarize
// checks what it built.
func checkRowIdentity(s summary.Summary) error {
	seen := map[rowKey]bool{}
	add := func(k rowKey) error {
		if seen[k] {
			return fmt.Errorf("%w: two %s", ErrDuplicateRow, k)
		}
		seen[k] = true
		return nil
	}
	for _, a := range s.DamageDone {
		for _, ab := range a.Abilities {
			if err := add(rowKey{"ability", a.GUID + "|" + ab.Via, ab.SpellID}); err != nil {
				return err
			}
		}
	}
	for _, c := range s.Casts {
		if err := add(rowKey{"cast", c.GUID, c.SpellID}); err != nil {
			return err
		}
	}
	for _, t := range s.Auras {
		if err := add(rowKey{"aura", t.TargetGUID, t.SpellID}); err != nil {
			return err
		}
	}
	return nil
}
