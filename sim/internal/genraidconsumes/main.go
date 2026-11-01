// Command genraidconsumes writes sim/request/raid-consumes.json, the consumables every
// spec's published level-60 raid number carries, which the /sim page's Raid-buffed preset
// sends by default.
//
//	go run ./internal/genraidconsumes
//
// TestRaidConsumesJSONIsCommitted fails when the committed file and this output disagree,
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
	out := "request/raid-consumes.json"
	if len(os.Args) > 1 {
		out = os.Args[1]
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
	if err := os.WriteFile(out, b, 0o644); err != nil {
		log.Fatal(err)
	}
	log.Printf("wrote %s", out)
}
