// api/internal/dataaddon/write_split.go
package dataaddon

import (
	"bufio"
	"bytes"
	"fmt"
	"sort"
)

// renderSplit is Render's fallback for a combined file over maxSingleFileBytes:
// "Data.lua" becomes a small header/init file (format, generated, and two empty tables),
// and one "Data-<region>.lua" per region that has any character or guild record merges its
// own records into the shared global. A record is never cut: each region file carries
// whole records, written exactly as the single file writes them. The TOC must list
// Data.lua before every Data-<region>.lua (addon/ForeverSixtyData/README.md documents the
// ordering requirement); addon-data-release.yml is the one place that ever needs to change
// the TOC's file list to match, since the checked-in TOC only ever lists Data.lua.
func renderSplit(d Data) (map[string][]byte, error) {
	out := map[string][]byte{}

	var header bytes.Buffer
	hw := bufio.NewWriter(&header)
	fmt.Fprintln(hw, "-- The combined file exceeded 4 MiB, so this build is split by region; see Data-<region>.lua.")
	writeOpening(hw, d)
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
			if d.Characters[key].Region == region {
				writeCharacter(bw, key, d.Characters[key])
			}
		}
		fmt.Fprintln(bw, "}) do ForeverSixtyData.characters[key] = row end")
		fmt.Fprintln(bw, "for key, row in pairs({")
		for _, key := range sortedKeys(d.Guilds) {
			if d.Guilds[key].Region == region {
				writeGuild(bw, key, d.Guilds[key])
			}
		}
		fmt.Fprintln(bw, "}) do ForeverSixtyData.guilds[key] = row end")
		if err := bw.Flush(); err != nil {
			return nil, err
		}
		out["Data-"+region+".lua"] = buf.Bytes()
	}
	return out, nil
}

// regionsOf lists every region any character or guild record belongs to, sorted.
func regionsOf(d Data) []string {
	set := map[string]bool{}
	for _, rec := range d.Characters {
		set[rec.Region] = true
	}
	for _, rec := range d.Guilds {
		set[rec.Region] = true
	}
	out := make([]string, 0, len(set))
	for r := range set {
		out = append(out, r)
	}
	sort.Strings(out)
	return out
}
