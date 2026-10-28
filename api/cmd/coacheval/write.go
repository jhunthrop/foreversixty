package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const (
	redactedMark   = "[redacted]"
	resultsFile    = "results.jsonl"
	reportFile     = "report.md"
	privateFile    = "fights.private.json"
	outDirMode     = 0o750
	outFileMode    = 0o600
	percent        = 100
	tokenThousands = 1000
)

// Redactor replaces every identifier it was given in any string it is asked to clean.
type Redactor struct{ secrets []string }

// NewRedactor collects the identifiers of the targets: report ids, player names and keys,
// every roster name and guid. Empty and one-character values are ignored.
func NewRedactor(targets []Target) Redactor {
	seen := map[string]bool{}
	var out []string
	add := func(s string) {
		if len(s) > 1 && !seen[s] {
			seen[s] = true
			out = append(out, s)
		}
	}
	for _, t := range targets {
		add(t.Fight.ReportID)
		for _, r := range t.Fight.Summary.Roster {
			add(r.Name)
			add(r.GUID)
		}
		for name, key := range t.Fight.PlayerKeys {
			add(name)
			add(key)
		}
	}
	sort.Slice(out, func(i, j int) bool { return len(out[i]) > len(out[j]) })
	return Redactor{secrets: out}
}

// Clean returns s with every identifier replaced.
func (r Redactor) Clean(s string) string {
	for _, id := range r.secrets {
		s = strings.ReplaceAll(s, id, redactedMark)
	}
	return s
}

// redact cleans every free-text field of a result.
func (r Redactor) redact(res Result) Result {
	res.Note = r.Clean(res.Note)
	res.FirstProblems, res.LastProblems = r.problems(res.FirstProblems), r.problems(res.LastProblems)
	outs := make([]string, len(res.Outputs))
	for i, o := range res.Outputs {
		outs[i] = r.Clean(o)
	}
	res.Outputs = outs
	if res.Pack != nil {
		var pack any
		if err := json.Unmarshal([]byte(r.Clean(mustJSON(res.Pack))), &pack); err == nil {
			res.Pack = pack
		}
	}
	return res
}

func (r Redactor) problems(ps []Problem) []Problem {
	out := make([]Problem, len(ps))
	for i, p := range ps {
		out[i] = Problem{Kind: p.Kind, Value: r.Clean(p.Value)}
	}
	return out
}

func mustJSON(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		return "null"
	}
	return string(b)
}

// WriteOutputs writes the JSONL, the markdown report and the private fight map into dir.
func WriteOutputs(dir string, targets []Target, results []Result, meta RunMeta) error {
	if err := os.MkdirAll(dir, outDirMode); err != nil {
		return fmt.Errorf("out dir: %w", err)
	}
	red := NewRedactor(targets)
	var jl bytes.Buffer
	for _, r := range results {
		line, err := json.Marshal(red.redact(r))
		if err != nil {
			return fmt.Errorf("encode result: %w", err)
		}
		jl.Write(append(line, '\n'))
	}
	private, err := json.MarshalIndent(privateFights(targets), "", "  ")
	if err != nil {
		return fmt.Errorf("encode fights: %w", err)
	}
	files := map[string][]byte{
		resultsFile: jl.Bytes(),
		reportFile:  []byte(red.Clean(RenderReport(results, meta))),
		privateFile: private,
	}
	for name, data := range files {
		if err := os.WriteFile(filepath.Join(dir, name), data, outFileMode); err != nil {
			return fmt.Errorf("write %s: %w", name, err)
		}
	}
	return nil
}

type privateFight struct {
	Label      string `json:"label"`
	ReportID   string `json:"report_id"`
	FightIndex int    `json:"fight_index"`
	PlayerKey  string `json:"player_key,omitempty"`
}

func privateFights(targets []Target) []privateFight {
	out := make([]privateFight, len(targets))
	for i, t := range targets {
		out[i] = privateFight{Label: t.Label, ReportID: t.Fight.ReportID, FightIndex: t.Fight.FightIndex,
			PlayerKey: t.Fight.PlayerKeys[t.Row.Name]}
	}
	return out
}
