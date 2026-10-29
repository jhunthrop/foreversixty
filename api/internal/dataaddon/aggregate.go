// api/internal/dataaddon/aggregate.go
package dataaddon

import (
	"math"

	"github.com/jhunthrop/foreversixty/api/internal/fs1"
	"github.com/jhunthrop/foreversixty/api/internal/rating"
	talentspec "github.com/jhunthrop/foreversixty/api/internal/spec"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// characterRecord is one character's entry in ForeverSixtyData.characters (contract 5).
// Every number in it is copied from rating.ScoreV2, the formula's only home; the writer
// never works out a tier or a component itself.
type characterRecord struct {
	// Region is the character's region: which file the record goes to when the render
	// splits (write_split.go). It is not written.
	Region string
	// Score is the character's Forever Score; nil when rating.ScoreV2 answers ErrNoLogs
	// (no rated fight in a public report), and then the record says score = nil (the addon
	// shows "no logs yet") so the guild's roster still lists it.
	Score *rating.ScoreV2
	// BestParses is the best overall percentile per boss name from the character's rated
	// fights in the window.
	BestParses map[string]int
	// ItemLevel (0 unknown) and Spec ("" unknown) are written only when known: from
	// ScoreV2 for a scored character, and from the item-level reader and the stored
	// export's talents for a no-logs one.
	ItemLevel int
	Spec      string
	Profile   string
}

// percentile is a rated fight's overall as a whole percentile, half away from zero.
func percentile(overall float64) int {
	return int(math.Round(overall))
}

// exportSpec is the build data's spec slug ("warrior-fury") the stored export's talents
// name: the class's tree with strictly the most points, the same rule GET /v1/me names a
// character's spec by. "" when the export does not decode, the build data has no such
// class, no points are spent, or two trees tie: a record would rather omit a spec than
// guess one.
func exportSpec(data *trees.Data, export string) string {
	if data == nil {
		return ""
	}
	decoded, ok := fs1.Decode(export)
	if !ok {
		return ""
	}
	build, ok := data.Build(decoded.DataBuild)
	if !ok {
		if build, ok = data.Active(); !ok {
			return ""
		}
	}
	most, ties := 0, 0
	for _, p := range decoded.Talents.Points {
		switch {
		case p > most:
			most, ties = p, 1
		case p == most && p > 0:
			ties++
		}
	}
	if most == 0 || ties > 1 {
		return ""
	}
	tree := talentspec.New(build).Name(decoded.ClassSlug, decoded.Talents.Points)
	if tree == "" {
		return ""
	}
	return rating.DataSpecSlug(decoded.ClassSlug, tree)
}
