// api/internal/dataaddon/write_split.go
package dataaddon

import (
	"bufio"
	"bytes"
	"fmt"
	"sort"
	"strings"
)

// renderSplit is Render's fallback for a combined file over
// maxSingleFileBytes: "Data.lua" becomes a small header/init file (format,
// generated, build, and two empty tables), and one "Data-<region>.lua" per
// region that has any character or guild row merges its own rows into the
// shared global. The TOC must list Data.lua before every Data-<region>.lua
// (addon/ForeverSixtyData/README.md documents the ordering requirement);
// addon-data-release.yml is the one place that ever needs to change the
// TOC's file list to match, since the checked-in TOC only ever lists
// Data.lua (see the plan's Task 12).
func renderSplit(d Data) (map[string][]byte, error) {
	out := map[string][]byte{}

	var header bytes.Buffer
	hw := bufio.NewWriter(&header)
	writeHeader(hw, d)
	fmt.Fprintln(hw, "-- The combined file exceeded 4 MiB, so this build is split by region; see Data-<region>.lua.")
	fmt.Fprintln(hw, "ForeverSixtyData = {")
	fmt.Fprintln(hw, "\tformat = 1,")
	fmt.Fprintf(hw, "\tgenerated = %q,\n", d.Generated.UTC().Format(isoFormat))
	fmt.Fprintf(hw, "\tbuild = %q,\n", d.Build)
	fmt.Fprintln(hw, "\tcharacters = {},")
	fmt.Fprintln(hw, "\tguilds = {},")
	fmt.Fprintln(hw, "}")
	if err := hw.Flush(); err != nil {
		return nil, err
	}
	out["Data.lua"] = header.Bytes()

	for _, region := range regionsOf(d) {
		var buf bytes.Buffer
		bw := bufio.NewWriter(&buf)
		fmt.Fprintf(bw, "-- addon/ForeverSixtyData/Data-%s.lua\n", region)
		fmt.Fprintln(bw, "-- One region's share of the nightly data-addon job's output; see Data.lua.")
		fmt.Fprintln(bw, "for key, row in pairs({")
		for _, key := range sortedKeys(d.Characters) {
			if !strings.HasPrefix(key, region+":") {
				continue
			}
			writeCharacterRow(bw, key, d.Characters[key])
		}
		fmt.Fprintln(bw, "}) do ForeverSixtyData.characters[key] = row end")
		fmt.Fprintln(bw, "for key, row in pairs({")
		for _, key := range sortedKeys(d.Guilds) {
			if !strings.HasPrefix(key, region+":") {
				continue
			}
			writeGuildRow(bw, key, d.Guilds[key])
		}
		fmt.Fprintln(bw, "}) do ForeverSixtyData.guilds[key] = row end")
		if err := bw.Flush(); err != nil {
			return nil, err
		}
		out["Data-"+region+".lua"] = buf.Bytes()
	}
	return out, nil
}

// regionsOf lists every region any character or guild key names, sorted --
// the leading segment before the first ":" of an addon key built by
// characterKey/guildKey.
func regionsOf(d Data) []string {
	set := map[string]bool{}
	for key := range d.Characters {
		set[strings.SplitN(key, ":", 2)[0]] = true
	}
	for key := range d.Guilds {
		set[strings.SplitN(key, ":", 2)[0]] = true
	}
	out := make([]string, 0, len(set))
	for r := range set {
		out = append(out, r)
	}
	sort.Strings(out)
	return out
}
