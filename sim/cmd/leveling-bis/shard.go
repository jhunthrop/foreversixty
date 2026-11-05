package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
)

// A full -all run outgrew bis.yml's 300-minute job on 2026-11-05 (run
// 41984980894: 1,250 re-admitted items and raid enchants on top of a
// 3 h 45 min ranking), and the workflow's timeout cannot change without a
// workflow-scope token. data/curated/bis-shard.json narrows -all to a
// named subset, so one night's ranking can be split across runs that each
// fit. With no file, -all ranks every written spec, as before.

// bisShardFile is the shard's place under the repository root.
var bisShardFile = filepath.Join("data", "curated", "bis-shard.json")

// bisShard is the file's content: the specs this run ranks, and why.
type bisShard struct {
	Specs  []string `json:"specs"`
	Reason string   `json:"reason"`
}

// loadBISShard reads the shard file; ok is false when there is none.
func loadBISShard(repoRoot string) (shard bisShard, ok bool, err error) {
	path := filepath.Join(repoRoot, bisShardFile)
	b, err := os.ReadFile(path)
	if errors.Is(err, fs.ErrNotExist) {
		return bisShard{}, false, nil
	}
	if err != nil {
		return bisShard{}, false, err
	}
	if err := json.Unmarshal(b, &shard); err != nil {
		return bisShard{}, false, fmt.Errorf("decoding %s: %w", path, err)
	}
	if len(shard.Specs) == 0 || shard.Reason == "" {
		return bisShard{}, false, fmt.Errorf("%s: a shard names at least one spec and its reason", path)
	}
	return shard, true, nil
}

// applyBISShard keeps the written specs the shard names, in written
// order. A shard naming a spec that is not written fails: a typo would
// otherwise rank nothing and leave that spec's old file standing.
func applyBISShard(written []string, shard bisShard) ([]string, error) {
	for _, s := range shard.Specs {
		if !slices.Contains(written, s) {
			return nil, fmt.Errorf("%s names %q, which is not a spec with a written rotation", bisShardFile, s)
		}
	}
	out := make([]string, 0, len(shard.Specs))
	for _, s := range written {
		if slices.Contains(shard.Specs, s) {
			out = append(out, s)
		}
	}
	return out, nil
}

// specsToRank is -all's list: the written specs, narrowed by a shard file
// when the repository carries one.
func specsToRank(repoRoot string) ([]string, error) {
	written, err := writtenSpecs(repoRoot)
	if err != nil {
		return nil, err
	}
	shard, ok, err := loadBISShard(repoRoot)
	if err != nil || !ok {
		return written, err
	}
	return applyBISShard(written, shard)
}
