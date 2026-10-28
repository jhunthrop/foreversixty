package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

const summaryExt = ".json"

// dirSource reads a directory of fight summary JSON files, the offline corpus. The file
// stem stands in for the report id. No pool is available offline, so the pack's pool
// comparison is absent; the raid size is the roster's headcount.
type dirSource struct {
	dir string
	now func() time.Time
}

func (d dirSource) Fights(_ context.Context, n int) ([]Fight, error) {
	paths, err := filepath.Glob(filepath.Join(d.dir, "*"+summaryExt))
	if err != nil {
		return nil, fmt.Errorf("corpus %s: %w", d.dir, err)
	}
	sort.Strings(paths)
	var out []Fight
	for _, p := range paths {
		if len(out) == n {
			break
		}
		f, err := readSummaryFile(p, d.now())
		if err != nil {
			return nil, err
		}
		out = append(out, f)
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("corpus %s: no %s files", d.dir, summaryExt)
	}
	return out, nil
}

func readSummaryFile(path string, fought time.Time) (Fight, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Fight{}, fmt.Errorf("corpus: %w", err)
	}
	var s summary.Summary
	if err := json.Unmarshal(data, &s); err != nil {
		return Fight{}, fmt.Errorf("corpus %s: %w", filepath.Base(path), err)
	}
	stem := strings.TrimSuffix(filepath.Base(path), summaryExt)
	return Fight{ReportID: stem, FightIndex: s.FightIndex, Size: len(s.Roster), FoughtAt: fought, Summary: s}, nil
}
