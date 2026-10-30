package discordbot

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

// The item embed's source line is the site BiS page's own (web/src/lib/bis/source-cell.ts
// resolveSourceCell then describeSourceCell, with bis/copy.ts and items/source-labels.ts's
// labels): the pick's source kind resolved against the build's loot.json, so a vendor reads
// "Vendor: Pix Xizzix", a reputation item "Stormpike Guard (exalted)", a quest "Quest:
// Saving the Best for Last · requires level 58". A kind loot.json has nothing for falls back to the page's
// own generic badge word. Change both together.

const lootFileName = "loot.json"

type lootFile struct {
	Sources []lootSource           `json:"sources"`
	Quests  map[string][]lootQuest `json:"quests"`
}

type lootSource struct {
	Kind        string             `json:"kind"`
	Name        string             `json:"name"`
	Items       []int              `json:"items"`
	Trash       []int              `json:"trash"`
	Bosses      []lootBoss         `json:"bosses"`
	ItemChances map[string]float64 `json:"item_chances"`
	Profession  string             `json:"profession"`
	Standing    string             `json:"standing"`
	Rank        *int               `json:"rank"`
	Faction     string             `json:"faction"`
	LevelMin    *int               `json:"level_min"`
	LevelMax    *int               `json:"level_max"`
}

type lootBoss struct {
	Name        string             `json:"name"`
	Items       []int              `json:"items"`
	ItemChances map[string]float64 `json:"item_chances"`
}

type lootQuest struct {
	Name    string `json:"name"`
	Faction string `json:"faction"`
	// MinLevel is the level a player must be to take the quest (classic-db), which the line
	// shows; the quest's own level misled as a requirement (spec 2026-10-29 §A.3).
	MinLevel int `json:"min_level"`
}

// sourceBook is one build's loot.json, read once.
type sourceBook struct {
	mu    sync.Mutex
	build string
	file  *lootFile
}

// lootFor is the build's loot.json; nil when the build has none (every line falls back).
func (b *sourceBook) lootFor(dir, build string) (*lootFile, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.file != nil && b.build == build {
		return b.file, nil
	}
	raw, err := os.ReadFile(filepath.Join(dir, build, lootFileName))
	if errors.Is(err, os.ErrNotExist) {
		b.build, b.file = build, &lootFile{}
		return b.file, nil
	}
	if err != nil {
		return nil, fmt.Errorf("discordbot: read loot for %s: %w", build, err)
	}
	var f lootFile
	if err := json.Unmarshal(raw, &f); err != nil {
		return nil, fmt.Errorf("discordbot: decode loot for %s: %w", build, err)
	}
	b.build, b.file = build, &f
	return b.file, nil
}

// sourceLine is the BiS page's source line for an item of kind, for faction's panel. A loot
// read that fails is logged and the line falls back to the badge word. With no kind (data
// that names none) it is raw, the BiS data's own source.
func (s *Service) sourceLine(itemID int, kind, faction, raw string) string {
	if kind == "" {
		return raw
	}
	loot, err := s.sources.lootFor(s.d.Guilds.BisDataDir(), s.d.Guilds.BisBuild())
	if err != nil {
		s.log.Warn("discordbot", "op", "item_source", "err", err)
		loot = &lootFile{}
	}
	return describeSource(loot, itemID, kind, faction)
}

func itemsOf(src lootSource) []int {
	out := append(append([]int{}, src.Items...), src.Trash...)
	for _, b := range src.Bosses {
		out = append(out, b.Items...)
	}
	return out
}

func names(src lootSource, itemID int) bool {
	for _, id := range itemsOf(src) {
		if id == itemID {
			return true
		}
	}
	return false
}

func findLoot(sources []lootSource, kind string, itemID int) (lootSource, bool) {
	for _, src := range sources {
		if src.Kind == kind && names(src, itemID) {
			return src, true
		}
	}
	return lootSource{}, false
}

func bossOf(src lootSource, itemID int) (lootBoss, bool) {
	for _, b := range src.Bosses {
		for _, id := range b.Items {
			if id == itemID {
				return b, true
			}
		}
	}
	return lootBoss{}, false
}

// chanceOf is the source's (or the item's boss's) drop chance; 0 means unknown.
func chanceOf(src lootSource, itemID int) float64 {
	chances := src.ItemChances
	if b, ok := bossOf(src, itemID); ok && b.ItemChances != nil {
		chances = b.ItemChances
	}
	return max(chances[strconv.Itoa(itemID)], 0)
}

// bestChance is the kind's source naming the item with the highest stated chance; ties and
// sources with no chance keep the first match.
func bestChance(sources []lootSource, kind string, itemID int) (lootSource, bool) {
	var best lootSource
	bestAt, found := -1.0, false
	for _, src := range sources {
		if src.Kind != kind || !names(src, itemID) {
			continue
		}
		at := chanceOf(src, itemID)
		if at == 0 {
			at = -1
		}
		if !found || at > bestAt {
			best, bestAt, found = src, at, true
		}
	}
	return best, found
}

func describeSource(loot *lootFile, itemID int, kind, faction string) string {
	if line, ok := resolvedSource(loot, itemID, kind, faction); ok {
		return line
	}
	return sourceBadge(kind, faction)
}

func resolvedSource(loot *lootFile, itemID int, kind, faction string) (string, bool) {
	switch kind {
	case "quest":
		return questLine(loot.Quests[strconv.Itoa(itemID)], faction)
	case "dungeon", "raid":
		src, ok := bestChance(loot.Sources, kind, itemID)
		if !ok {
			return "", false
		}
		return placeLine(src, itemID), true
	case "crafted":
		src, ok := findLoot(loot.Sources, "crafted", itemID)
		return craftedLine(src), ok
	case "vendor":
		return vendorLine(loot.Sources, itemID)
	case "rep":
		src, ok := findLoot(loot.Sources, "rep", itemID)
		return repLine(src), ok
	case "zone", "world":
		src, ok := findLoot(loot.Sources, kind, itemID)
		if chance := chanceOf(src, itemID); ok && chance > 0 {
			return chanceLine(chance, src.Name), true
		}
		return src.Name, ok
	case "pvp":
		src, ok := findLoot(loot.Sources, "pvp", itemID)
		if !ok || src.Rank == nil || !realFaction(src.Faction) {
			return "", false
		}
		return pvpLine(*src.Rank, src.Faction), true
	case "world_drop":
		src, ok := findLoot(loot.Sources, "world_drop", itemID)
		return worldDropLine(src), ok
	}
	return "", false
}

func realFaction(f string) bool { return f == "alliance" || f == "horde" }

// questLine is the quest of the panel's faction (else the first): "Quest: {name} · requires
// level {min_level}", the name alone when the data has no min_level.
func questLine(options []lootQuest, faction string) (string, bool) {
	if len(options) == 0 {
		return "", false
	}
	q := options[0]
	for _, o := range options {
		if o.Faction == faction {
			q = o
			break
		}
	}
	line := fill(sourceCopy.Quest, "{name}", q.Name)
	if q.MinLevel > 0 {
		line += lineJoin + fill(sourceCopy.QuestLevel, "{min_level}", strconv.Itoa(q.MinLevel))
	}
	return line, true
}

// placeLine is "{chance}% from {boss}" when the drop chance is known, else "{instance} · {boss}".
func placeLine(src lootSource, itemID int) string {
	boss, hasBoss := bossOf(src, itemID)
	bossName := ""
	if hasBoss {
		bossName = boss.Name
		if bossName == "" {
			bossName = fill(sourceCopy.Unnamed, "{zone}", src.Name)
		}
	}
	if chance := chanceOf(src, itemID); chance > 0 {
		from := bossName
		if !hasBoss {
			from = fill(sourceCopy.Trash, "{instance}", src.Name)
		}
		return chanceLine(chance, from)
	}
	if !hasBoss {
		return src.Name
	}
	return src.Name + lineJoin + bossName
}

func chanceLine(chance float64, from string) string {
	figure := strconv.Itoa(int(math.Round(chance)))
	if chance < 1 {
		figure = sourceCopy.UnderOne
	}
	return fill(sourceCopy.Chance, "{chance}", figure, "{from}", from)
}

func craftedLine(src lootSource) string {
	if src.Profession == "" || strings.EqualFold(src.Profession, src.Name) {
		return fill(sourceCopy.Crafted, "{name}", src.Name)
	}
	return fill(sourceCopy.CraftedNamed, "{Profession}", classLabel(src.Profession), "{name}", src.Name)
}

// vendorLine: a vendor that sells a PvP rank's rewards reads as that rank, as the page does.
func vendorLine(sources []lootSource, itemID int) (string, bool) {
	if pvp, ok := findLoot(sources, "pvp", itemID); ok && pvp.Rank != nil && realFaction(pvp.Faction) {
		return pvpLine(*pvp.Rank, pvp.Faction), true
	}
	src, ok := findLoot(sources, "vendor", itemID)
	if !ok {
		return "", false
	}
	if src.Rank != nil && realFaction(src.Faction) {
		return pvpLine(*src.Rank, src.Faction), true
	}
	return fill(sourceCopy.Vendor, "{npc}", src.Name), true
}

func repLine(src lootSource) string {
	if src.Standing == "" {
		return src.Name
	}
	return fill(sourceCopy.Rep, "{faction}", src.Name, "{standing}", src.Standing)
}

// pvpLine is "PvP rank 11 · Knight-Lieutenant · Alliance"; loot.json's rank is the client's
// RequiredPVPRank, the ladder's rank plus 4.
func pvpLine(rank int, faction string) string {
	titles := pvpTitles[faction]
	side := classLabel(faction)
	if i := rank - pvpRankOffset; i >= 0 && i < len(titles) {
		return fill(sourceCopy.PvP, "{rank}", strconv.Itoa(rank), "{title}", titles[i], "{Faction}", side)
	}
	return fill(sourceCopy.PvPUntitled, "{rank}", strconv.Itoa(rank), "{Faction}", side)
}

func worldDropLine(src lootSource) string {
	if src.LevelMin == nil || src.LevelMax == nil {
		return sourceCopy.WorldDrop
	}
	return sourceCopy.WorldDrop + fill(sourceCopy.WorldDropLevels, "{min}", strconv.Itoa(*src.LevelMin), "{max}", strconv.Itoa(*src.LevelMax))
}

// sourceBadge is the page's generic badge word for a kind; a quest names the panel's faction.
func sourceBadge(kind, faction string) string {
	if kind == "quest" {
		if faction == "horde" {
			return sourceCopy.HordeQuest
		}
		return sourceCopy.AllianceQuest
	}
	if word, ok := sourceKindWords[kind]; ok {
		return word
	}
	return kind
}
