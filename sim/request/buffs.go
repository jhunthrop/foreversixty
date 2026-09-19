package request

// The settings bar sends buffs and consumables as string ids, and the
// engine holds them as protobuf fields spread over five messages. The
// ids are the engine's own names in lower snake case - a buff is a field
// name (RaidBuffs.battle_shout is "battle_shout"), a consumable is an
// enum value name (AgilityElixir.ElixirOfTheMongoose is
// "elixir_of_the_mongoose") - so the mapping is read off the protobuf
// descriptors rather than copied into a table of several hundred rows
// that would drift from the engine the first time it gained a flask.
//
// Nothing is dropped quietly: an id no message knows is an error at the
// boundary, because a sim that silently ran without a world buff reports
// a DPS number that is wrong and says nothing about why.

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/sim/internal/strcase"
	"github.com/wowsims/classic/sim/core/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	// ErrUnknownBuff is returned for a buff id none of the engine's four
	// buff messages carries a field for.
	ErrUnknownBuff = errors.New("request: unknown buff")
	// ErrUnknownConsume is returned for a consumable id the engine's
	// Consumes message has no enum value for.
	ErrUnknownConsume = errors.New("request: unknown consumable")
	// ErrAmbiguousConsume is returned when a bare consumable id names a
	// value two Consumes fields can both hold - both weapon imbue slots
	// take the same enum - and the caller has to say which.
	ErrAmbiguousConsume = errors.New("request: ambiguous consumable")
)

// improvedSuffix marks the talented version of a graded buff:
// "battle_shout:improved". The plain id is the plain version, which is
// what it has always meant, so every stored request keeps its meaning.
const improvedSuffix = ":improved"

// buffSet is the four buff messages a raid sim request carries. One id
// lands in exactly one of them, so a buff named in two messages - a
// paladin's blessing is both an individual buff and a raid-wide one - is
// not applied twice.
type buffSet struct {
	Individual *proto.IndividualBuffs
	Party      *proto.PartyBuffs
	Raid       *proto.RaidBuffs
	Debuffs    *proto.Debuffs
}

// buffsFor sorts every buff id into the message that carries it. The
// order is the precedence: the most specific container wins, and
// RaidBuffs is the catch-all, so "blessing_of_wisdom" is the blessing on
// this player rather than the raid-wide aura of the same name.
func buffsFor(ids []string) (buffSet, error) {
	set := buffSet{
		Individual: &proto.IndividualBuffs{},
		Party:      &proto.PartyBuffs{},
		Raid:       &proto.RaidBuffs{},
		Debuffs:    &proto.Debuffs{},
	}
	targets := []protoreflect.ProtoMessage{set.Debuffs, set.Individual, set.Party, set.Raid}
	for _, id := range ids {
		if !enableField(targets, id) {
			return buffSet{}, fmt.Errorf("%w: %q", ErrUnknownBuff, id)
		}
	}
	return set, nil
}

// enableField turns on the named field of the first message that has
// one, and reports whether any did. An id carrying improvedSuffix asks
// for the second non-zero value of a graded field, and matches nothing
// else: a boolean buff has no improved form, and answering one with the
// plain buff would apply less than the id promised.
func enableField(targets []protoreflect.ProtoMessage, id string) bool {
	name, improved := strings.CutSuffix(id, improvedSuffix)
	for _, target := range targets {
		msg := target.ProtoReflect()
		fd := msg.Descriptor().Fields().ByName(protoreflect.Name(name))
		if fd == nil {
			continue
		}
		if improved {
			value, ok := gradedValue(fd)
			if !ok {
				return false
			}
			msg.Set(fd, protoreflect.ValueOfEnum(value))
			return true
		}
		switch fd.Kind() {
		case protoreflect.BoolKind:
			msg.Set(fd, protoreflect.ValueOfBool(true))
		case protoreflect.Int32Kind:
			// A count, not a flag: innervates, power infusions, atiesh
			// stacks. One is the smallest thing "on" can mean.
			msg.Set(fd, protoreflect.ValueOfInt32(1))
		case protoreflect.EnumKind:
			msg.Set(fd, protoreflect.ValueOfEnum(onEnumValue(fd)))
		default:
			continue
		}
		return true
	}
	return false
}

// onEnumValue is the first non-zero value of an enum, which is what
// turning a tristate or a graded buff on means.
func onEnumValue(fd protoreflect.FieldDescriptor) protoreflect.EnumNumber {
	values := fd.Enum().Values()
	for i := 0; i < values.Len(); i++ {
		if n := values.Get(i).Number(); n != 0 {
			return n
		}
	}
	return 0
}

// gradedValue is the improved value of a graded field: the SECOND
// non-zero value of its enum. A field with fewer than two is not
// graded and has no improved form.
func gradedValue(fd protoreflect.FieldDescriptor) (protoreflect.EnumNumber, bool) {
	if fd.Kind() != protoreflect.EnumKind {
		return 0, false
	}
	values := fd.Enum().Values()
	var seen int
	for i := 0; i < values.Len(); i++ {
		n := values.Get(i).Number()
		if n == 0 {
			continue
		}
		seen++
		if seen == 2 {
			return n, true
		}
	}
	return 0, false
}

// consumes maps consumable ids onto the engine's Consumes message.
//
// An id is one of two things: an engine enum value name in lower snake
// case ("elixir_of_the_mongoose"), or a client item id from the build's
// simconsumes.json ("item:13452"), which the table joins onto the same
// enum value by name. Either may be qualified by the field it belongs
// in - "main_hand_imbue:shadow_oil", "off_hand_imbue:item:3824" - which
// is how the two weapon imbue slots are told apart.
func consumes(ids []string, table *Consumables) (*proto.Consumes, error) {
	out := &proto.Consumes{}
	msg := out.ProtoReflect()
	for _, id := range ids {
		field, value := splitConsumeID(id)
		key, err := consumeKey(value, table)
		if err != nil {
			return nil, fmt.Errorf("%w: %q", err, id)
		}
		matches := consumeFields(msg.Descriptor(), field, key)
		switch len(matches) {
		case 0:
			return nil, fmt.Errorf("%w: %q", ErrUnknownConsume, id)
		case 1:
			msg.Set(matches[0].field, matches[0].value)
		default:
			return nil, fmt.Errorf("%w: %q is held by %s; name the field, as in %q",
				ErrAmbiguousConsume, id, strings.Join(fieldNames(matches), " and "),
				string(matches[0].field.Name())+":"+value)
		}
	}
	return out, nil
}

// splitConsumeID pulls off a field qualifier, if the id carries one.
// "item" is not a field of Consumes, so a bare "item:13452" is the item
// form rather than a qualified value.
func splitConsumeID(id string) (field, value string) {
	head, tail, ok := strings.Cut(id, ":")
	if !ok || head == itemPrefix {
		return "", id
	}
	return head, tail
}

// consumeKey turns the value half of an id into the lookup key an enum
// value name normalises to. An item id needs the build's table; without
// one, saying so beats resolving it to nothing.
func consumeKey(value string, table *Consumables) (string, error) {
	rest, ok := strings.CutPrefix(value, itemPrefix+":")
	if !ok {
		return value, nil
	}
	itemID, err := strconv.ParseInt(rest, 10, 64)
	if err != nil {
		return "", ErrUnknownConsume
	}
	if table == nil {
		return "", fmt.Errorf("%w: no consumable table is loaded, so an item id cannot be resolved; pass one in Options, from data/builds/<build>/simconsumes.json",
			ErrUnknownConsume)
	}
	key, ok := table.key(itemID)
	if !ok {
		return "", fmt.Errorf("%w: the build's simconsumes.json has no such item", ErrUnknownConsume)
	}
	return key, nil
}

// consumesDescriptor is the Consumes message's descriptor, which the
// vocabulary and the item table both walk.
func consumesDescriptor() protoreflect.MessageDescriptor {
	return (&proto.Consumes{}).ProtoReflect().Descriptor()
}

// consumeMatch is one field of Consumes that can hold a consumable id,
// and the value to set it to.
type consumeMatch struct {
	field protoreflect.FieldDescriptor
	value protoreflect.Value
}

// consumeFields finds every field of Consumes that can express the
// consumable, in field-declaration order so the answer is stable. An
// empty field name matches any field.
func consumeFields(desc protoreflect.MessageDescriptor, field, value string) []consumeMatch {
	var out []consumeMatch
	fields := desc.Fields()
	for i := 0; i < fields.Len(); i++ {
		fd := fields.Get(i)
		if field != "" && string(fd.Name()) != field {
			continue
		}
		switch fd.Kind() {
		case protoreflect.EnumKind:
			values := fd.Enum().Values()
			enum := string(fd.Enum().Name())
			for j := 0; j < values.Len(); j++ {
				v := values.Get(j)
				if v.Number() == 0 || !namesValue(string(v.Name()), enum, value) {
					continue
				}
				out = append(out, consumeMatch{fd, protoreflect.ValueOfEnum(v.Number())})
			}
		case protoreflect.BoolKind:
			// A handful of consumables are their own flag rather than a
			// member of an enum: dragon_breath_chili is one.
			if string(fd.Name()) == value {
				out = append(out, consumeMatch{fd, protoreflect.ValueOfBool(true)})
			}
		}
	}
	return out
}

// namesValue reports whether an id names this enum value. Some values
// repeat their enum's name - Food's GrilledSquid is FoodGrilledSquid -
// and the item they are named after does not, so both the full name and
// the name without that prefix answer to the id.
func namesValue(valueName, enumName, id string) bool {
	if strcase.Snake(valueName) == id {
		return true
	}
	stripped, ok := strings.CutPrefix(valueName, enumName)
	return ok && stripped != "" && strcase.Snake(stripped) == id
}

func fieldNames(matches []consumeMatch) []string {
	names := make([]string, 0, len(matches))
	for _, m := range matches {
		names = append(names, string(m.field.Name()))
	}
	sort.Strings(names)
	return names
}
