// api/internal/dataaddon/write.go
package dataaddon

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

// dataFormat is contract 5's version, the first field of the file.
const dataFormat = 2

// Data is everything one Render call writes.
type Data struct {
	Generated  time.Time
	Characters map[string]characterRecord // by Name-Realm
	Guilds     map[string]guildRecord     // by guild name
}

// Render returns the file(s) to publish: one ("Data.lua") when the whole
// file fits in maxSingleFileBytes, or a header plus one per region when it
// does not (renderSplit, write_split.go). Every file returned, together,
// defines the same ForeverSixtyData global the single-file form would --
// the addon reads one global regardless of how many files built it.
func Render(d Data) (map[string][]byte, error) {
	return renderWithLimit(d, maxSingleFileBytes)
}

// renderWithLimit is Render with the split threshold spelled out, so a
// test can exercise the split path without generating a real 4 MiB fixture.
func renderWithLimit(d Data, limit int) (map[string][]byte, error) {
	var single bytes.Buffer
	if err := render(&single, d); err != nil {
		return nil, err
	}
	if single.Len() <= limit {
		return map[string][]byte{"Data.lua": single.Bytes()}, nil
	}
	return renderSplit(d)
}

// render writes contract 5's ForeverSixtyData global, exactly the layout of
// docs/contracts/fixtures/data-addon-v2.lua: one field per line, records sorted by key.
func render(w io.Writer, d Data) error {
	bw := bufio.NewWriter(w)
	writeOpening(bw, d)
	fmt.Fprintln(bw, "\tcharacters = {")
	for _, key := range sortedKeys(d.Characters) {
		writeCharacter(bw, key, d.Characters[key])
	}
	fmt.Fprintln(bw, "\t},")
	fmt.Fprintln(bw, "\tguilds = {")
	for _, key := range sortedKeys(d.Guilds) {
		writeGuild(bw, key, d.Guilds[key])
	}
	fmt.Fprintln(bw, "\t},")
	fmt.Fprintln(bw, "}")
	return bw.Flush()
}

// writeOpening writes the global's first three lines: its name, format and the UTC date
// the file was generated.
func writeOpening(bw *bufio.Writer, d Data) {
	fmt.Fprintln(bw, "ForeverSixtyData = {")
	fmt.Fprintf(bw, "\tformat = %d,\n", dataFormat)
	fmt.Fprintf(bw, "\tgenerated = %s,\n", luaString(d.Generated.UTC().Format(time.DateOnly)))
}

// writeCharacter writes one character record. Every number is the record's own, copied
// from rating.ScoreV2; a no-logs record says score = nil and carries only what is known
// about the character besides its page.
func writeCharacter(bw *bufio.Writer, key string, rec characterRecord) {
	fmt.Fprintf(bw, "\t\t[%s] = {\n", luaString(key))
	field := func(name, value string) { fmt.Fprintf(bw, "\t\t\t%s = %s,\n", name, value) }
	known := func() {
		if rec.ItemLevel > 0 {
			field("item_level", strconv.Itoa(rec.ItemLevel))
		}
		if rec.Spec != "" {
			field("spec", luaString(rec.Spec))
		}
	}
	sc := rec.Score
	if sc == nil {
		field("score", "nil")
		known()
		field("profile", luaString(rec.Profile))
		fmt.Fprintln(bw, "\t\t},")
		return
	}
	c := sc.Components
	field("score", strconv.Itoa(sc.Score))
	field("tier", luaString(sc.Tier))
	field("as_of", luaString(sc.AsOf.UTC().Format(time.DateOnly)))
	field("components", fmt.Sprintf("{ gear = %d, experience = %d, parses = %d, attendance = %d, discipline = %d }",
		c.Gear, c.Experience, c.Parses, c.Attendance, c.Discipline))
	field("gear_score", strconv.Itoa(sc.GearScore))
	known()
	field("experience", fmt.Sprintf("{ dungeons = %d, raid_kills = %d, alts = %d }",
		sc.Experience.Dungeons, sc.Experience.RaidKills, sc.Experience.Alts))
	field("best_parses", bestParses(rec.BestParses))
	field("attendance", luaFraction(sc.Attendance))
	field("fights", strconv.Itoa(sc.Fights))
	field("profile", luaString(rec.Profile))
	fmt.Fprintln(bw, "\t\t},")
}

// writeGuild writes one guild record.
func writeGuild(bw *bufio.Writer, key string, rec guildRecord) {
	fmt.Fprintf(bw, "\t\t[%s] = {\n", luaString(key))
	field := func(name, value string) { fmt.Fprintf(bw, "\t\t\t%s = %s,\n", name, value) }
	field("standing", strconv.Itoa(rec.Standing))
	field("nights", strconv.Itoa(rec.Nights))
	field("roster", strconv.Itoa(rec.Roster))
	if rec.Attendance != nil {
		field("attendance", luaFraction(*rec.Attendance))
	}
	members := make([]string, len(rec.Members))
	for i, m := range rec.Members {
		members[i] = luaString(m)
	}
	if len(members) == 0 {
		field("members", "{}")
	} else {
		field("members", "{ "+strings.Join(members, ", ")+" }")
	}
	fmt.Fprintln(bw, "\t\t},")
}

// bestParses is a boss -> percentile table, bosses sorted.
func bestParses(parses map[string]int) string {
	if len(parses) == 0 {
		return "{}"
	}
	entries := make([]string, 0, len(parses))
	for _, boss := range sortedKeys(parses) {
		entries = append(entries, fmt.Sprintf("[%s] = %d", luaString(boss), parses[boss]))
	}
	return "{ " + strings.Join(entries, ", ") + " }"
}

// luaFraction writes a fraction (0.75 is 75 percent) to the hundredth, which is a whole
// percent: the file is shown as percentages, and a 2/3 needs no sixteen digits.
func luaFraction(v float64) string {
	return strconv.FormatFloat(math.Round(v*100)/100, 'f', -1, 64)
}

// luaString quotes s as a Lua string literal any Lua 5.1 reads: a backslash, a double
// quote and every control byte are escaped (a control byte as a three-digit \ddd, so a digit after it is never read into it); every
// other byte, UTF-8 included, is written as is.
func luaString(s string) string {
	var b strings.Builder
	b.WriteByte('"')
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c == '\\' || c == '"':
			b.WriteByte('\\')
			b.WriteByte(c)
		case c == '\n':
			b.WriteString(`\n`)
		case c < 0x20 || c == 0x7f:
			fmt.Fprintf(&b, `\%03d`, c)
		default:
			b.WriteByte(c)
		}
	}
	b.WriteByte('"')
	return b.String()
}

// sortedKeys returns m's keys in ascending byte order, for deterministic
// output regardless of Go's randomized map iteration.
func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
