// Command genraidcover writes api/internal/guilds/raids_cover_catalogue.go from
// data/curated/raid-cover.json (the buff and debuff cover catalogue, owner decision D15), with
// each row's icon stem read off the build's own data. data/curated is not in the API image, so the
// API carries this Go twin; raids_cover_catalogue_test.go pins the two together. Run it through
// `go generate ./internal/guilds/` from api/.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	build := flag.String("build", "", "the client build whose icons to read")
	curated := flag.String("curated", "../../../data/curated/raid-cover.json", "the catalogue")
	data := flag.String("data", "../../../data/builds", "the data/builds directory")
	out := flag.String("out", "raids_cover_catalogue.go", "the Go file to write")
	flag.Parse()
	if err := run(*build, *curated, *data, *out); err != nil {
		fmt.Fprintln(os.Stderr, "genraidcover:", err)
		os.Exit(1)
	}
}

func run(build, curated, data, out string) error {
	if build == "" {
		return fmt.Errorf("-build is required")
	}
	var file catalogueFile
	if err := readJSON(curated, &file); err != nil {
		return err
	}
	icons, err := loadIcons(filepath.Join(data, build))
	if err != nil {
		return err
	}
	src, err := goSource(build, file.Entries, icons)
	if err != nil {
		return err
	}
	return os.WriteFile(out, src, 0o644)
}

func readJSON(path string, into any) error {
	raw, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(raw, into); err != nil {
		return fmt.Errorf("%s: %w", path, err)
	}
	return nil
}
