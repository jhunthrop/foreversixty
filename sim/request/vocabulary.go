package request

// The id vocabulary.
//
// The settings bar sends buffs, consumables and professions as strings,
// and an id this module cannot map is an error rather than a silent
// drop, so the web lane has to know exactly which strings exist. The
// buff and consumable lists are not hand-kept - they are read off the
// engine's own protobuf descriptors - so they are generated here rather
// than written down twice: KnownBuffs, KnownConsumables and
// KnownProfessions are the vocabulary, IDS.md is the same vocabulary as
// prose for whoever builds the settings bar, and a test holds the
// committed file to what this code produces.

import (
	"fmt"
	"sort"
	"strings"

	"github.com/jhunthrop/foreversixty/sim/internal/strcase"
	"github.com/wowsims/classic/sim/core/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// buffMessages are the four buff messages, in the precedence a buff id
// is resolved against them: the most specific container that carries a
// field of that name wins, and RaidBuffs is the catch-all.
func buffMessages() []protoreflect.ProtoMessage {
	return []protoreflect.ProtoMessage{
		&proto.Debuffs{}, &proto.IndividualBuffs{}, &proto.PartyBuffs{}, &proto.RaidBuffs{},
	}
}

// KnownBuffs lists every buff id Build accepts, sorted. An id appears
// once however many messages carry a field of that name; it lands in
// the first of them.
func KnownBuffs() []string {
	seen := map[string]bool{}
	for _, id := range buffVocabulary() {
		seen[id.id] = true
	}
	out := make([]string, 0, len(seen))
	for id := range seen {
		out = append(out, id)
	}
	sort.Strings(out)
	return out
}

// KnownProfessions lists every profession slug Build accepts, sorted.
// A character carries at most two of them, and one the engine has no
// enum for is refused at the boundary, so the settings bar needs the
// list as much as it needs the buff one.
func KnownProfessions() []string {
	out := make([]string, 0, len(professions))
	for slug := range professions {
		out = append(out, slug)
	}
	sort.Strings(out)
	return out
}

// KnownConsumables lists every consumable id Build accepts without a
// build's item table, sorted. An id two fields could hold is listed
// only in its qualified forms, so everything listed resolves.
func KnownConsumables() []string {
	out := make([]string, 0, 256)
	for _, entry := range consumeVocabulary() {
		out = append(out, entry.id)
	}
	sort.Strings(out)
	return out
}

// vocabularyEntry is one id and where it lands.
type vocabularyEntry struct {
	id    string
	field string // the proto field it sets
	owner string // the message that field belongs to
}

// professionVocabulary names every profession slug and the engine enum
// value it sets. Unlike the buff and consumable walks this reads the
// site's own map rather than the engine's descriptor: a profession slug
// is an interface the settings bar is built against, so gaining one is
// a decision, not something an engine bump does silently.
// TestTheProfessionMapMatchesTheEngineEnum holds the map to the enum, so
// the decision cannot be missed either.
func professionVocabulary() []vocabularyEntry {
	out := make([]vocabularyEntry, 0, len(professions))
	for slug, p := range professions {
		out = append(out, vocabularyEntry{id: slug, field: p.String(), owner: "Player"})
	}
	return out
}

// buffVocabulary walks the buff messages in precedence order and names
// every field that can be turned on.
func buffVocabulary() []vocabularyEntry {
	var out []vocabularyEntry
	claimed := map[string]bool{}
	for _, msg := range buffMessages() {
		desc := msg.ProtoReflect().Descriptor()
		fields := desc.Fields()
		for i := 0; i < fields.Len(); i++ {
			fd := fields.Get(i)
			switch fd.Kind() {
			case protoreflect.BoolKind, protoreflect.Int32Kind, protoreflect.EnumKind:
			default:
				continue
			}
			id := string(fd.Name())
			if claimed[id] {
				continue
			}
			claimed[id] = true
			out = append(out, vocabularyEntry{id: id, field: id, owner: string(desc.Name())})
		}
	}
	return out
}

// consumeVocabulary names every consumable the engine models, by its
// enum value name, qualified where two fields could hold it.
func consumeVocabulary() []vocabularyEntry {
	desc := consumesDescriptor()
	var out []vocabularyEntry
	fields := desc.Fields()
	for i := 0; i < fields.Len(); i++ {
		fd := fields.Get(i)
		switch fd.Kind() {
		case protoreflect.BoolKind:
			out = append(out, vocabularyEntry{id: string(fd.Name()), field: string(fd.Name()), owner: "Consumes"})
		case protoreflect.EnumKind:
			values := fd.Enum().Values()
			for j := 0; j < values.Len(); j++ {
				v := values.Get(j)
				if v.Number() == 0 {
					continue
				}
				bare := strcase.Snake(string(v.Name()))
				id := bare
				// Both weapon imbue slots hold one enum, so a bare imbue
				// id names two fields and has to say which.
				if len(consumeFields(desc, "", bare)) > 1 {
					id = string(fd.Name()) + ":" + bare
				}
				out = append(out, vocabularyEntry{id: id, field: string(fd.Name()), owner: "Consumes"})
			}
		}
	}
	return out
}

// IDsMarkdown renders the vocabulary as the document the web lane
// reads. sim/internal/genids writes it to IDS.md and a test holds the
// committed file to this output.
func IDsMarkdown() string {
	var b strings.Builder
	b.WriteString(`# The sim request's id vocabulary

Generated by ` + "`go run ./internal/genids`" + ` from the engine's own protobuf
descriptors. Do not edit by hand.

` + "`api.SimRequest.Character`" + `'s ` + "`buffs`" + `, ` + "`consumes`" + ` and ` + "`professions`" + ` are
lists of these ids. An id this module cannot map is an error at the
boundary — a sim that quietly ran without a world buff would report a
DPS number that is wrong and say nothing about why — so the settings
bar must send ids from these lists and nothing else.

## Buffs

A buff id is the engine's own protobuf field name, in lower snake case.
One id lands in exactly one message, the first of
**Debuffs → IndividualBuffs → PartyBuffs → RaidBuffs** that carries a
field of that name, so ` + "`blessing_of_wisdom`" + ` is the blessing on this
player rather than the raid-wide aura of the same name. A boolean field
is turned on, a count is set to one, and a graded field (the engine's
` + "`TristateEffect`" + `) is set to its plain version; the settings bar has no
id for the improved form yet.

| id | lands in |
| --- | --- |
`)
	for _, e := range sortedByID(buffVocabulary()) {
		fmt.Fprintf(&b, "| `%s` | %s |\n", e.id, e.owner)
	}

	b.WriteString(`
## Consumables

A consumable id is one of two things.

**The engine's enum value name**, in lower snake case: ` + "`elixir_of_the_mongoose`" + `.
Where a value repeats its enum's name — ` + "`Food`" + `'s ` + "`FoodGrilledSquid`" + ` — the
name without that prefix (` + "`grilled_squid`" + `) answers to the id as well.

**A client item id** from the build's consumable table,
` + "`data/builds/<build>/simconsumes.json`" + `: ` + "`item:13452`" + `. The table joins an
item onto the engine's value by name, and the caller loads it and
passes it in ` + "`request.Options`" + ` (` + "`request.BuildWith`" + `); ` + "`request.Build`" + ` has
no table, so it rejects item ids. Most of that file's rows are food and
potions no sim models, and those do not resolve — ` + "`(*Consumables).IDs()`" + `
lists the ones that do.

Both weapon imbue slots hold one enum, so an imbue must say which hand:
` + "`main_hand_imbue:shadow_oil`" + `, ` + "`off_hand_imbue:item:3824`" + `. Any id may be
qualified by its field this way; only the ambiguous ones must be.

| id | sets |
| --- | --- |
`)
	for _, e := range sortedByID(consumeVocabulary()) {
		fmt.Fprintf(&b, "| `%s` | Consumes.%s |\n", e.id, e.field)
	}

	b.WriteString(`
## Professions

` + "`api.SimRequest.Character.professions`" + ` is a list of these slugs, at most
two of them, each at most once. The engine models a profession as a
source of self-only recipes and effects — Engineering's grenades and
trinkets, Blacksmithing's extra socket — so a slug this module cannot
map is refused rather than dropped: a sim that quietly ran without the
profession the player counted on would report a wrong number and say
nothing about why. A character with no professions is legal.

| slug | engine enum |
| --- | --- |
`)
	for _, e := range sortedByID(professionVocabulary()) {
		fmt.Fprintf(&b, "| `%s` | Profession.%s |\n", e.id, e.field)
	}
	return b.String()
}

func sortedByID(entries []vocabularyEntry) []vocabularyEntry {
	out := append([]vocabularyEntry(nil), entries...)
	sort.SliceStable(out, func(i, j int) bool { return out[i].id < out[j].id })
	return out
}
