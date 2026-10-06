// Package enginetalents writes a talent build, named by stable talent
// id, as the positional talent string the COMPILED engine reads.
//
// The engine's core.FillTalentsProto reads one digit per proto field,
// in field-number order, sliced per tree by the class package's
// TalentTreeSizes. That order is generated from whichever client build
// the engine's proto was last regenerated from, which is not always
// the site's active build: at engine 90f9325b0 the paladin message
// still carries Improved Holy Strike and Crusade (both gone in
// 1.60.1.70009) and the shaman message has Elemental Fury and
// Elemental Alacrity in each other's slot. A string written in the
// active build's own (tier, column) order is therefore misread by the
// engine for those classes. This package matches each active-build
// talent to its engine field by stable talent id - the generator writes
// "// node <id>, tier ..." above every field of proto/<class>.proto -
// and writes the digit where the engine will read it. Names are not
// stable: druid's Mangle and Primal Fury nodes are Primal Bite and
// Blood Frenzy in 1.60.1.70009.
package enginetalents

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/druid"
	"github.com/wowsims/classic/sim/hunter"
	"github.com/wowsims/classic/sim/mage"
	"github.com/wowsims/classic/sim/paladin"
	"github.com/wowsims/classic/sim/priest"
	"github.com/wowsims/classic/sim/rogue"
	"github.com/wowsims/classic/sim/shaman"
	"github.com/wowsims/classic/sim/warlock"
	"github.com/wowsims/classic/sim/warrior"
	gproto "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Field is one engine talent field: its proto name, the tree it sits
// in, its index within that tree's digits and whether it is a bool
// (one-rank) field.
type Field struct {
	Name  string
	Tree  int
	Index int
	Bool  bool
}

// GoName is the field's generated Go struct name (improved_holy_strike
// becomes ImprovedHolyStrike), which is how engine code reads it.
func (f Field) GoName() string {
	var b strings.Builder
	for _, part := range strings.Split(f.Name, "_") {
		if part == "" {
			continue
		}
		b.WriteString(strings.ToUpper(part[:1]))
		b.WriteString(part[1:])
	}
	return b.String()
}

// Layout is one class's talent string layout as the compiled engine
// reads it.
type Layout struct {
	TreeSizes [3]int
	byID      map[int]Field
}

// Module is the engine's module path.
const Module = "github.com/wowsims/classic"

// SourceDir is the engine module's source directory as the build
// resolved it (go.mod's replace in development, the module cache in
// CI), from moduleDir, any directory inside the sim module.
func SourceDir(moduleDir string) (string, error) {
	cmd := exec.Command("go", "list", "-m", "-f", "{{.Dir}}", Module)
	cmd.Dir = moduleDir
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("enginetalents: locating the engine source (go list -m %s): %w", Module, err)
	}
	dir := strings.TrimSpace(string(out))
	if dir == "" {
		return "", fmt.Errorf("enginetalents: go list -m %s gave no directory", Module)
	}
	return dir, nil
}

// protoNodeRE is talentgen's comment and the field under it:
//
//	// node 105328, tier 0 col 0, 2 rank(s), spell 1310902
//	int32 improved_holy_strike = 1;
var protoNodeRE = regexp.MustCompile(`// node (\d+), tier[^\n]*\n\s*(?:bool|int32) (\w+) = (\d+);`)

// protoNodeIDs reads proto/<class>.proto's node ids by field name.
func protoNodeIDs(engineDir, class string) (map[string]int, error) {
	path := filepath.Join(engineDir, "proto", class+".proto")
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("enginetalents: %w", err)
	}
	out := map[string]int{}
	for _, m := range protoNodeRE.FindAllStringSubmatch(string(b), -1) {
		id, err := strconv.Atoi(m[1])
		if err != nil {
			return nil, fmt.Errorf("enginetalents: %s: node id %q: %w", path, m[1], err)
		}
		out[m[2]] = id
	}
	return out, nil
}

type classLayout struct {
	msg   gproto.Message
	sizes [3]int
}

func classLayouts() map[string]classLayout {
	return map[string]classLayout{
		"druid":   {&proto.DruidTalents{}, druid.TalentTreeSizes},
		"hunter":  {&proto.HunterTalents{}, hunter.TalentTreeSizes},
		"mage":    {&proto.MageTalents{}, mage.TalentTreeSizes},
		"paladin": {&proto.PaladinTalents{}, paladin.TalentTreeSizes},
		"priest":  {&proto.PriestTalents{}, priest.TalentTreeSizes},
		"rogue":   {&proto.RogueTalents{}, rogue.TalentTreeSizes},
		"shaman":  {&proto.ShamanTalents{}, shaman.TalentTreeSizes},
		"warlock": {&proto.WarlockTalents{}, warlock.TalentTreeSizes},
		"warrior": {&proto.WarriorTalents{}, warrior.TalentTreeSizes},
	}
}

// ForClass reads the compiled engine's layout for a class slug, with
// each field's talent id from the engine's proto source in engineDir
// (SourceDir). Every compiled field must carry an id there, or the
// source and the compiled engine disagree and nothing is trusted.
func ForClass(engineDir, class string) (Layout, error) {
	cl, ok := classLayouts()[class]
	if !ok {
		return Layout{}, fmt.Errorf("enginetalents: unknown class %q", class)
	}
	ids, err := protoNodeIDs(engineDir, class)
	if err != nil {
		return Layout{}, err
	}
	fields := cl.msg.ProtoReflect().Descriptor().Fields()
	total := cl.sizes[0] + cl.sizes[1] + cl.sizes[2]
	if fields.Len() != total {
		return Layout{}, fmt.Errorf("enginetalents: %s message has %d fields but TalentTreeSizes %v sum to %d", class, fields.Len(), cl.sizes, total)
	}
	l := Layout{TreeSizes: cl.sizes, byID: make(map[int]Field, total)}
	tree, offset := 0, 0
	for n := 1; n <= total; n++ {
		fd := fields.ByNumber(protoreflect.FieldNumber(n))
		if fd == nil {
			return Layout{}, fmt.Errorf("enginetalents: %s message has no field number %d", class, n)
		}
		for n-1 >= offset+cl.sizes[tree] {
			offset += cl.sizes[tree]
			tree++
		}
		id, ok := ids[string(fd.Name())]
		if !ok {
			return Layout{}, fmt.Errorf("enginetalents: compiled %s field %s has no node id in the engine's proto source", class, fd.Name())
		}
		l.byID[id] = Field{
			Name:  string(fd.Name()),
			Tree:  tree,
			Index: n - 1 - offset,
			Bool:  fd.Kind() == protoreflect.BoolKind,
		}
	}
	return l, nil
}

// FieldFor is the engine field an active-build talent writes to.
func (l Layout) FieldFor(treePosition int, node leveling.TalentNode) (Field, error) {
	f, ok := l.byID[node.ID]
	if !ok {
		return Field{}, fmt.Errorf("enginetalents: talent %d %q has no engine field", node.ID, node.Name)
	}
	if f.Tree != treePosition {
		return Field{}, fmt.Errorf("enginetalents: talent %d %q is in tree %d but the engine reads %q in tree %d", node.ID, node.Name, treePosition, f.Name, f.Tree)
	}
	return f, nil
}

// Encode writes ranks (talent id -> points) as the engine's positional
// string. Every talent of trees must map onto an engine field, so a
// talent the engine cannot see fails loudly instead of silently
// simming as zero; a rank above a bool field's 1 is refused for the
// same reason (the engine reads "2" in a bool slot as false).
func (l Layout) Encode(trees []leveling.TalentTree, ranks map[int]int) (string, error) {
	digits := make([][]byte, 3)
	for i := range digits {
		digits[i] = []byte(strings.Repeat("0", l.TreeSizes[i]))
	}
	seen := make(map[int]bool, len(ranks))
	for ti, tree := range trees {
		for _, node := range tree.Talents {
			f, err := l.FieldFor(ti, node)
			if err != nil {
				return "", err
			}
			r := ranks[node.ID]
			seen[node.ID] = true
			if r == 0 {
				continue
			}
			if r < 0 || r > 9 || (f.Bool && r > 1) {
				return "", fmt.Errorf("enginetalents: %d points in %q (%s) cannot be written", r, node.Name, f.Name)
			}
			digits[f.Tree][f.Index] = byte('0' + r)
		}
	}
	for id, r := range ranks {
		if r != 0 && !seen[id] {
			return "", fmt.Errorf("enginetalents: talent %d is not in the given trees", id)
		}
	}
	return string(digits[0]) + "-" + string(digits[1]) + "-" + string(digits[2]), nil
}

// Reposition rewrites a positional talent string s - written in
// trees' own (tier, column) order, the format
// leveling.LadderTalentString writes and the published band `talents`
// field carries over the API - as the string the compiled engine
// reads for the same points, by id. It is the one call every site
// that already holds such a string (the ranker's per-band truncated
// build) needs before the string reaches the engine: decode by
// position once (leveling.TalentRanksFromString), then Encode by id,
// which is the only read FieldFor/byID can be trusted to get right
// when the active build's own tree shape has moved since the engine's
// proto was last regenerated (this package's own doc).
func (l Layout) Reposition(trees []leveling.TalentTree, s string) (string, error) {
	ranks, err := leveling.TalentRanksFromString(trees, s)
	if err != nil {
		return "", fmt.Errorf("enginetalents: %w", err)
	}
	return l.Encode(trees, ranks)
}
