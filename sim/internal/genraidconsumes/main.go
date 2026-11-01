// Command genraidconsumes writes the /sim page's preset tables: sim/request/raid-consumes.json,
// the consumables every spec's raid number carries at each level (the Raid-buffed preset),
// and sim/request/published-presets.json, which preset each band publishes and what the bare
// preset (the class kit) sends at each level.
//
//	go run ./internal/genraidconsumes
//
// TestRaidConsumesJSONIsCommitted and TestPublishedPresetsJSONIsCommitted fail when a
// committed file and this output disagree,
// so a change to data/curated/presets.json or the class kit that forgets the page is a
// failing test rather than a page number that quietly differs from the published one.
package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/jhunthrop/foreversixty/sim/request"
)

func main() {
	out, publishedOut := "request/raid-consumes.json", "request/published-presets.json"
	if len(os.Args) > 2 {
		out, publishedOut = os.Args[1], os.Args[2]
	}
	presets, err := request.LoadPresets("../data/curated/presets.json")
	if err != nil {
		log.Fatal(err)
	}
	build, err := leveling.ReadActiveBuild("..")
	if err != nil {
		log.Fatal(err)
	}
	levels, err := request.LoadItemLevelsFile(filepath.Join("..", "data", "builds", build, "items.json"))
	if err != nil {
		log.Fatal(err)
	}
	b, err := request.RaidConsumesJSON(presets, levels)
	if err != nil {
		log.Fatal(err)
	}
	writeFile(out, b)
	published, err := request.PublishedPresetsJSON()
	if err != nil {
		log.Fatal(err)
	}
	writeFile(publishedOut, published)
}

func writeFile(path string, b []byte) {
	if err := os.WriteFile(path, b, 0o644); err != nil {
		log.Fatal(err)
	}
	log.Printf("wrote %s", path)
}
