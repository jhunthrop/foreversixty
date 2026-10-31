package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// classicranks.go is the ranker's copy of the pipeline's
// pipeline.loot.pvp_faction.drop_non_rank_rewards: classic-db states every
// classic item's honor rank, zero included, so a classic id it names
// unranked is no rank reward, however it reached a pvp source (its title,
// a rank recovered from an earlier loot.json, or a reputation
// quartermaster's stock backfilled under the one ranked item it sells -
// Kelm Hargunth put Dryad's Wrist Bindings under pvp:rank-12:horde).
// Both read the build's committed raw/classicdb/item_template.json, so the
// rule holds for a loot.json generated before the pipeline learned it.

// classicItemTemplate is the build's committed classic-db item extract.
const classicItemTemplate = "raw/classicdb/item_template.json"

// loadClassicUnranked returns the classic item ids the extract names with
// no honor rank. A build with no extract gets none.
func loadClassicUnranked(buildDir string) (map[int]bool, error) {
	b, err := os.ReadFile(filepath.Join(buildDir, classicItemTemplate))
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", classicItemTemplate, err)
	}
	var extract struct {
		Items []struct {
			ID                int `json:"id"`
			RequiredHonorRank int `json:"required_honor_rank"`
		} `json:"items"`
	}
	if err := json.Unmarshal(b, &extract); err != nil {
		return nil, fmt.Errorf("decoding %s: %w", classicItemTemplate, err)
	}
	unranked := make(map[int]bool, len(extract.Items))
	for _, item := range extract.Items {
		if item.RequiredHonorRank == 0 {
			unranked[item.ID] = true
		}
	}
	return unranked, nil
}

// dropNonRankRewards returns idx without the pvp sources of every unranked
// classic id, and with the rank a vendor row inherited from one
// (vendorInheritsPvpRankGate) taken back off.
func dropNonRankRewards(idx lootIndex, unranked map[int]bool) lootIndex {
	if len(unranked) == 0 {
		return idx
	}
	out := make(lootIndex, len(idx))
	for id, srcs := range idx {
		if !unranked[id] {
			out[id] = srcs
			continue
		}
		kept := make([]itemSource, 0, len(srcs))
		for _, s := range srcs {
			if s.Kind == "pvp" {
				continue
			}
			if s.Kind == "vendor" && s.Rank != 0 {
				s.Rank, s.Title, s.Label = 0, "", s.SourceName
			}
			kept = append(kept, s)
		}
		out[id] = kept
	}
	return out
}
