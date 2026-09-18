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
	"strings"
	"unicode"

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
// one, and reports whether any did.
func enableField(targets []protoreflect.ProtoMessage, id string) bool {
	for _, target := range targets {
		msg := target.ProtoReflect()
		fd := msg.Descriptor().Fields().ByName(protoreflect.Name(id))
		if fd == nil {
			continue
		}
		switch fd.Kind() {
		case protoreflect.BoolKind:
			msg.Set(fd, protoreflect.ValueOfBool(true))
		case protoreflect.Int32Kind:
			// A count, not a flag: innervates, power infusions, atiesh
			// stacks. One is the smallest thing "on" can mean.
			msg.Set(fd, protoreflect.ValueOfInt32(1))
		case protoreflect.EnumKind:
			// TristateEffect and its kin: value one is the plain version
			// of the buff, value two the talented one. The settings bar
			// has no id for the improved form yet, so "on" is plain.
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

// consumes maps consumable ids onto the engine's Consumes message. An id
// is an enum value name in lower snake case, optionally qualified by the
// field it belongs in - "main_hand_imbue:shadow_oil" - which is how the
// two weapon imbue slots are told apart.
func consumes(ids []string) (*proto.Consumes, error) {
	out := &proto.Consumes{}
	msg := out.ProtoReflect()
	for _, id := range ids {
		field, value, qualified := strings.Cut(id, ":")
		if !qualified {
			field, value = "", id
		}
		matches := consumeFields(msg.Descriptor(), field, value)
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
			for j := 0; j < values.Len(); j++ {
				v := values.Get(j)
				if v.Number() == 0 || snake(string(v.Name())) != value {
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

func fieldNames(matches []consumeMatch) []string {
	names := make([]string, 0, len(matches))
	for _, m := range matches {
		names = append(names, string(m.field.Name()))
	}
	sort.Strings(names)
	return names
}

// snake turns a protobuf enum value name into the id the settings bar
// sends: ElixirOfTheMongoose becomes elixir_of_the_mongoose.
func snake(name string) string {
	var b strings.Builder
	b.Grow(len(name) + 4)
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i > 0 {
				b.WriteByte('_')
			}
			b.WriteRune(unicode.ToLower(r))
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}
