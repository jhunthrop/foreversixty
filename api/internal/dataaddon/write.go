// api/internal/dataaddon/write.go
package dataaddon

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"sort"
	"time"
)

// isoFormat is Ratings.lua's own pattern:
// "^(%d%d%d%d)%-(%d%d)%-(%d%d)T(%d%d):(%d%d):(%d%d)Z$".
const isoFormat = "2006-01-02T15:04:05Z"

// dataLuaPath is the checked-in path this content replaces at release
// time -- echoed in the file's own header comment.
const dataLuaPath = "addon/ForeverSixtyData/Data.lua"

// Data is everything one Render call writes.
type Data struct {
	Generated  time.Time
	Build      string
	Characters map[string]characterRow
	Guilds     map[string]guildRow
}

// Render returns the file(s) to publish: one ("Data.lua") when the whole
// file fits in maxSingleFileBytes, or a header plus one per region when it
// does not (renderSplit, write_split.go). Every file returned, together,
// defines the same ForeverSixtyData global the single-file form would --
// Ratings.lua reads one global regardless of how many files built it.
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

func render(w io.Writer, d Data) error {
	bw := bufio.NewWriter(w)
	writeHeader(bw, d)
	fmt.Fprintln(bw, "ForeverSixtyData = {")
	fmt.Fprintln(bw, "\tformat = 1,")
	fmt.Fprintf(bw, "\tgenerated = %q,\n", d.Generated.UTC().Format(isoFormat))
	fmt.Fprintf(bw, "\tbuild = %q,\n", d.Build)
	fmt.Fprintln(bw, "\tcharacters = {")
	for _, key := range sortedKeys(d.Characters) {
		writeCharacterRow(bw, key, d.Characters[key])
	}
	fmt.Fprintln(bw, "\t},")
	fmt.Fprintln(bw, "\tguilds = {")
	for _, key := range sortedKeys(d.Guilds) {
		writeGuildRow(bw, key, d.Guilds[key])
	}
	fmt.Fprintln(bw, "\t},")
	fmt.Fprintln(bw, "}")
	return bw.Flush()
}

func writeHeader(bw *bufio.Writer, d Data) {
	fmt.Fprintf(bw, "-- %s\n", dataLuaPath)
	fmt.Fprintf(bw, "-- Generated %s for data build %s by the nightly data-addon job (api/internal/dataaddon).\n",
		d.Generated.UTC().Format(isoFormat), d.Build)
	fmt.Fprintln(bw, "-- Do not edit by hand: the addon-data-release workflow overwrites this file every night.")
}

func writeCharacterRow(bw *bufio.Writer, key string, row characterRow) {
	fmt.Fprintf(bw, "\t\t[%q] = { rating = %d, mean90 = %d", key, row.Rating, row.Mean90)
	for _, name := range componentNames {
		if v, ok := row.Components[name]; ok {
			fmt.Fprintf(bw, ", %s = %d", name, v)
		}
	}
	fmt.Fprintf(bw, ", fights = %d },\n", row.Fights)
}

func writeGuildRow(bw *bufio.Writer, key string, row guildRow) {
	fmt.Fprintf(bw, "\t\t[%q] = { name = %q", key, row.Name)
	if row.Progress != "" {
		fmt.Fprintf(bw, ", progress = %q", row.Progress)
	}
	fmt.Fprintf(bw, ", nights = %d, roster = %d, members = { ", row.Nights, row.Roster)
	for i, m := range row.Members {
		if i > 0 {
			fmt.Fprint(bw, ", ")
		}
		fmt.Fprintf(bw, "%q", m)
	}
	fmt.Fprintln(bw, " } },")
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
