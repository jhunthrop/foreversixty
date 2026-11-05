// Command genequiprules writes api/internal/guilds/equip_rules_table.go from a build's
// data/builds/<build>/equip-skills.json (the guild Loot tab's who-can-equip-what table). Run it
// through `go generate ./internal/guilds/` from api/ after `python -m pipeline equip-skills`.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jhunthrop/foreversixty/api/internal/equipskills"
)

func main() {
	build := flag.String("build", "", "the client build whose equip-skills.json to read")
	data := flag.String("data", "../../../data/builds", "the data/builds directory")
	out := flag.String("out", "equip_rules_table.go", "the Go file to write")
	flag.Parse()
	if err := run(*build, *data, *out); err != nil {
		fmt.Fprintln(os.Stderr, "genequiprules:", err)
		os.Exit(1)
	}
}

func run(build, data, out string) error {
	if build == "" {
		return fmt.Errorf("-build is required")
	}
	path := filepath.Join(data, build, "equip-skills.json")
	body, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	doc, err := equipskills.Parse(body)
	if err != nil {
		return err
	}
	src, err := equipskills.GoSource(doc, equipskills.SourcePath(build))
	if err != nil {
		return err
	}
	return os.WriteFile(out, src, 0o644)
}
