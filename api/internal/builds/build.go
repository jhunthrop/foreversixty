// Package builds saves, validates, and describes planner builds.
package builds

import (
	"crypto/sha256"
	"encoding/base32"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// MaxPoints is the most talent points a level-60 build can spend.
const MaxPoints = 51

// MaxTitleLen is the longest title a build may carry, in runes.
const MaxTitleLen = 60

// Slots are the seventeen gear slots a build may fill, in the planner's
// display order.
var Slots = []string{
	"head", "neck", "shoulder", "back", "chest", "wrist", "hands", "waist",
	"legs", "feet", "finger1", "finger2", "trinket1", "trinket2",
	"main_hand", "off_hand", "ranged",
}

// slotSet indexes Slots for membership tests.
var slotSet = func() map[string]bool {
	m := make(map[string]bool, len(Slots))
	for _, s := range Slots {
		m[s] = true
	}
	return m
}()

// ItemSlot maps a planner slot to the slot value items carry in the data
// files: the two ring slots and the two trinket slots share one item slot.
func ItemSlot(slot string) string {
	switch slot {
	case "finger1", "finger2":
		return "finger"
	case "trinket1", "trinket2":
		return "trinket"
	default:
		return slot
	}
}

// Input is the POST /v1/builds request body.
type Input struct {
	ClassID     int            `json:"class_id"`
	RaceID      int            `json:"race_id"`
	TreeVersion string         `json:"tree_version"`
	PointOrder  []int          `json:"point_order"`
	Gear        map[string]int `json:"gear"`
	Title       string         `json:"title"`
}

// Build is a stored build record and the JSON the API returns for it.
type Build struct {
	ID          string         `json:"id"`
	ClassID     int            `json:"class_id"`
	RaceID      int            `json:"race_id"`
	TreeVersion string         `json:"tree_version"`
	PointOrder  []int          `json:"point_order"`
	Gear        map[string]int `json:"gear"`
	Title       string         `json:"title,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	Views       int64          `json:"views"`
}

// Normalize fills in the zero values the hash and the store depend on: a nil
// point order and a nil gear map must hash as [] and {}, not null, so that a
// body that omits them yields the same id as one that sends them empty. The
// title is trimmed, since the contract stores a trimmed title.
func (in Input) Normalize() Input {
	if in.PointOrder == nil {
		in.PointOrder = []int{}
	}
	if in.Gear == nil {
		in.Gear = map[string]int{}
	}
	in.Title = strings.TrimSpace(in.Title)
	return in
}

// canonical is the exact shape hashed to produce a build id: field names in
// alphabetical order, title excluded. encoding/json writes struct fields in
// declaration order and map keys sorted, with no whitespace, which is what
// the contract's "sorted keys and no whitespace" means.
type canonical struct {
	ClassID     int            `json:"class_id"`
	Gear        map[string]int `json:"gear"`
	PointOrder  []int          `json:"point_order"`
	RaceID      int            `json:"race_id"`
	TreeVersion string         `json:"tree_version"`
}

// CanonicalJSON returns the bytes ID hashes.
func CanonicalJSON(in Input) ([]byte, error) {
	in = in.Normalize()
	b, err := json.Marshal(canonical{
		ClassID:     in.ClassID,
		Gear:        in.Gear,
		PointOrder:  in.PointOrder,
		RaceID:      in.RaceID,
		TreeVersion: in.TreeVersion,
	})
	if err != nil {
		return nil, fmt.Errorf("builds: canonical json: %w", err)
	}
	return b, nil
}

// ID is the first eight characters of the lowercase, unpadded RFC 4648
// base32 encoding of the SHA-256 of CanonicalJSON. Saving the same build
// twice yields the same id; the title is not part of it, so retitling a
// build is the same build and the first title wins.
func ID(in Input) (string, error) {
	b, err := CanonicalJSON(in)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(b)
	enc := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(sum[:])
	return strings.ToLower(enc)[:IDLen], nil
}

// IDLen is how many characters of the base32 hash an id carries.
const IDLen = 8

// ValidID reports whether s has the shape ID produces: IDLen
// characters of lowercase, unpadded RFC 4648 base32. It says nothing
// about whether that build exists - it is the cheap shape check a
// route does before it goes near the store.
func ValidID(s string) bool {
	if len(s) != IDLen {
		return false
	}
	for _, r := range s {
		if !(r >= 'a' && r <= 'z') && !(r >= '2' && r <= '7') {
			return false
		}
	}
	return true
}

// New turns a validated input into the record to store.
func New(in Input) (Build, error) {
	in = in.Normalize()
	id, err := ID(in)
	if err != nil {
		return Build{}, err
	}
	return Build{
		ID:          id,
		ClassID:     in.ClassID,
		RaceID:      in.RaceID,
		TreeVersion: in.TreeVersion,
		PointOrder:  in.PointOrder,
		Gear:        in.Gear,
		Title:       in.Title,
	}, nil
}
