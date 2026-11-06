// Command genraidbosses writes api/internal/guilds/raid_bosses_table.go from a build's
// data/builds/<build>/raid-bosses.json (the first raid tier's bosses and loot). Run it through
// `go generate ./internal/guilds/` from api/ after `python -m pipeline raid-bosses`;
// tools/activate-build.sh runs it for each new client build.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jhunthrop/foreversixty/api/internal/raidbosses"
)

func main() {
	build := flag.String("build", "", "the client build whose raid-bosses.json to read")
	data := flag.String("data", "../../../data/builds", "the data/builds directory")
	out := flag.String("out", "raid_bosses_table.go", "the Go file to write")
	flag.Parse()
	if err := run(*build, *data, *out); err != nil {
		fmt.Fprintln(os.Stderr, "genraidbosses:", err)
		os.Exit(1)
	}
}

func run(build, data, out string) error {
	if build == "" {
		return fmt.Errorf("-build is required")
	}
	body, err := os.ReadFile(filepath.Join(data, build, "raid-bosses.json"))
	if err != nil {
		return err
	}
	doc, err := raidbosses.Parse(body)
	if err != nil {
		return err
	}
	src, err := raidbosses.GoSource(doc, raidbosses.SourcePath(build))
	if err != nil {
		return err
	}
	return os.WriteFile(out, src, 0o644)
}
