package coach

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/rating"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/consumables"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

const manifestFile = "manifest.json"

// SpecFidelities is the sims store's per-spec validation rows; *sims.Store satisfies it.
type SpecFidelities interface {
	Specs(ctx context.Context) ([]sims.SpecFidelity, error)
}

// FileTables is the production TableSource. BuildsDir is the image's /data (data/builds in
// tests), CoachDir is the directory of spec rule files. Spell data, trainables and spec
// rules are read once and cached; the returned maps are shared, so callers must not mutate
// them. It is safe for concurrent use.
type FileTables struct {
	BuildsDir, CoachDir string
	Specs               SpecFidelities

	mu      sync.Mutex
	builds  []clientBuild
	classes map[classKey]classData
	rules   map[string]SpecRules
}

var _ TableSource = (*FileTables)(nil)

type clientBuild struct {
	name      string
	fetchedAt time.Time
}

type classKey struct{ build, class string }

type classData struct {
	spells    map[int64]SpellConst
	trainable map[string][]TrainableRank
}

// BuildAt is the client build in force at t: the newest client build (numeric compare, as
// trees.Data.Latest; named data sets such as forever-prebeta never count) whose manifest
// fetched_at is at or before t, else the oldest client build. This approximates the build
// the fight was played on: the manifest records when we fetched a build, not when the
// realm shipped it.
func (f *FileTables) BuildAt(t time.Time) (string, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	builds, err := f.clientBuilds()
	if err != nil {
		return "", err
	}
	pick := builds[0]
	for _, b := range builds {
		if !b.fetchedAt.After(t) {
			pick = b
		}
	}
	return pick.name, nil
}

// clientBuilds lists the client builds oldest first, once. The caller holds f.mu.
func (f *FileTables) clientBuilds() ([]clientBuild, error) {
	if f.builds != nil {
		return f.builds, nil
	}
	entries, err := os.ReadDir(f.BuildsDir)
	if err != nil {
		return nil, fmt.Errorf("coach tables: builds dir: %w", err)
	}
	var out []clientBuild
	for _, e := range entries {
		if !e.IsDir() || !isClientBuildName(e.Name()) {
			continue
		}
		at, err := manifestFetchedAt(filepath.Join(f.BuildsDir, e.Name(), manifestFile))
		if err != nil {
			return nil, err
		}
		out = append(out, clientBuild{name: e.Name(), fetchedAt: at})
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("coach tables: no client builds in %s", f.BuildsDir)
	}
	sort.Slice(out, func(i, j int) bool { return versionLess(out[i].name, out[j].name) })
	f.builds = out
	return out, nil
}

func manifestFetchedAt(path string) (time.Time, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return time.Time{}, fmt.Errorf("coach tables: %s: %w", path, err)
	}
	var m struct {
		FetchedAt time.Time `json:"fetched_at"`
	}
	if err := json.Unmarshal(data, &m); err != nil {
		return time.Time{}, fmt.Errorf("coach tables: %s: %w", path, err)
	}
	return m.FetchedAt, nil
}

// isClientBuildName is true for "1.60.1.69893" and false for a named data set.
func isClientBuildName(name string) bool {
	parts := strings.Split(name, ".")
	if len(parts) < 2 {
		return false
	}
	for _, p := range parts {
		if _, err := strconv.Atoi(p); err != nil {
			return false
		}
	}
	return true
}

// versionLess compares dot-separated numeric versions segment by segment.
func versionLess(a, b string) bool {
	as, bs := strings.Split(a, "."), strings.Split(b, ".")
	for i := 0; i < len(as) && i < len(bs); i++ {
		x, _ := strconv.Atoi(as[i])
		y, _ := strconv.Atoi(bs[i])
		if x != y {
			return x < y
		}
	}
	return len(as) < len(bs)
}

// TablesFor assembles the data tables for one roster row at the build in force at the pull.
// A row with no class or spec, or a class the build ships no spell data for, answers
// ErrNoTables; an unreadable file is a plain error naming it. A spec without a rule file
// gets the zero SpecRules, and an encounter without a mechanics table a nil Mechanics.
func (f *FileTables) TablesFor(ctx context.Context, row summary.RosterRow, encounterID int64, foughtAt time.Time) (Tables, error) {
	if row.Class == "" || row.Spec == "" {
		return Tables{}, fmt.Errorf("%w: row %q has no class or spec", ErrNoTables, row.Name)
	}
	build, err := f.BuildAt(foughtAt)
	if err != nil {
		return Tables{}, err
	}
	slug := rating.DataSpecSlug(row.Class, row.Spec)
	class, err := f.classFor(build, classSlug(row.Class))
	if err != nil {
		return Tables{}, err
	}
	rules, err := f.rulesFor(slug)
	if err != nil {
		return Tables{}, err
	}
	worst, err := f.worstActions(ctx, slug)
	if err != nil {
		return Tables{}, err
	}
	t := Tables{Rules: rules, Spells: class.spells, Trainable: class.trainable, WorstActions: worst}
	t.Rotation, t.RotationOK = sims.RotationFor(build, slug, int(row.Level))
	cat := consumables.Load()
	t.Consumables, _ = cat.For(row.Role)
	t.RaidBuffs = cat.RaidBuffs
	if m, ok := mechanics.Load(encounterID); ok {
		t.Mechanics = &m
	}
	return t, nil
}

func classSlug(class string) string { return strings.ReplaceAll(strings.ToLower(class), " ", "-") }

// classFor loads and caches one class's spell data at a build.
func (f *FileTables) classFor(build, class string) (classData, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	key := classKey{build, class}
	if c, ok := f.classes[key]; ok {
		return c, nil
	}
	spells, err := readBuildFile(f.BuildsDir, build, "spellconst", class, ParseSpellConst)
	if err != nil {
		return classData{}, err
	}
	trainable, err := readBuildFile(f.BuildsDir, build, "trainables", class, ParseTrainables)
	if err != nil {
		return classData{}, err
	}
	if f.classes == nil {
		f.classes = map[classKey]classData{}
	}
	c := classData{spells: spells, trainable: trainable}
	f.classes[key] = c
	return c, nil
}

// readBuildFile parses <builds>/<build>/<dir>/<class>.json; a missing file is ErrNoTables.
func readBuildFile[T any](buildsDir, build, dir, class string, parse func([]byte) (T, error)) (T, error) {
	var zero T
	path := filepath.Join(buildsDir, build, dir, class+".json")
	data, err := os.ReadFile(path)
	if errors.Is(err, fs.ErrNotExist) {
		return zero, fmt.Errorf("%w: %s is absent", ErrNoTables, path)
	}
	if err != nil {
		return zero, fmt.Errorf("coach tables: %s: %w", path, err)
	}
	v, err := parse(data)
	if err != nil {
		return zero, fmt.Errorf("%s: %w", path, err)
	}
	return v, nil
}

// rulesFor loads and caches a spec's rule file; a missing file is the zero value.
func (f *FileTables) rulesFor(slug string) (SpecRules, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if r, ok := f.rules[slug]; ok {
		return r, nil
	}
	r, _, err := LoadSpecRules(f.CoachDir, slug)
	if err != nil {
		return SpecRules{}, err
	}
	if f.rules == nil {
		f.rules = map[string]SpecRules{}
	}
	f.rules[slug] = r
	return r, nil
}

// worstActions is the names on the spec's sims WorstActions list.
func (f *FileTables) worstActions(ctx context.Context, slug string) (map[string]bool, error) {
	out := map[string]bool{}
	if f.Specs == nil {
		return out, nil
	}
	specs, err := f.Specs.Specs(ctx)
	if err != nil {
		return nil, fmt.Errorf("coach tables: sims specs: %w", err)
	}
	for _, s := range specs {
		if s.Spec != slug {
			continue
		}
		for _, w := range s.WorstActions {
			out[w.Name] = true
		}
	}
	return out, nil
}
