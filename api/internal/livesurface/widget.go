package livesurface

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// Widget is GET /v1/me/widget (W3-C12): one small answer every home-screen widget, lock-screen
// widget, Live Activity and ongoing notification draws from, so no two surfaces can disagree.
// GeneratedAt is to the minute, so an unchanged answer keeps its ETag between quick refreshes.
type Widget struct {
	GeneratedAt time.Time              `json:"generated_at"`
	Character   *WidgetCharacter       `json:"character"`
	Guild       *guilds.WidgetGuild    `json:"guild"`
	NextRaid    *guilds.WidgetNextRaid `json:"next_raid"`
	LastRaid    *guilds.WidgetLastRaid `json:"last_raid"`
	Upgrades    *WidgetUpgrades        `json:"upgrades"`
	Live        WidgetLive             `json:"live"`
}

// WidgetCharacter is the character the snapshot is for. SpecName is its primary tree, null
// when not known yet.
type WidgetCharacter struct {
	Key      string  `json:"key"`
	Name     string  `json:"name"`
	Class    *string `json:"class"`
	SpecName *string `json:"spec_name"`
}

// WidgetUpgrades is the Home answer's top rows. State is the Home answer's state (ok,
// no_gear, under_20, no_spec, no_list); Sentence is its words for any state but ok, null
// otherwise. Rows are the first three in the Home answer's order and Count and TotalGain are
// the whole list's, so the medium widget's footer matches Home.
type WidgetUpgrades struct {
	State     string       `json:"state"`
	Sentence  *string      `json:"sentence"`
	Count     int          `json:"count"`
	TotalGain *float64     `json:"total_gain"`
	Unit      string       `json:"unit"`
	Rows      []UpgradeRow `json:"rows"`
}

// UpgradeRow is one upgrade: the slot, the pick and its measured gain (null when the gain is
// not modelled yet; the widget prints `—`).
type UpgradeRow struct {
	Slot      string      `json:"slot"`
	SlotLabel string      `json:"slot_label"`
	Pick      UpgradePick `json:"pick"`
	Gain      *float64    `json:"gain"`
}

// UpgradePick is the item a row recommends.
type UpgradePick struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Quality   int    `json:"quality"`
	Icon      string `json:"icon"`
	ItemLevel int    `json:"item_level"`
}

// WidgetLive is what the live surfaces need beside the raid: the caller's server runs still
// going (0 for an account without server_sims, which has none).
type WidgetLive struct {
	SimRunsRunning int `json:"sim_runs_running"`
}

// widgetRows is how many upgrade rows the snapshot carries (the medium widget's three).
const widgetRows = 3

// widget answers GET /v1/me/widget[?character=<key>].
func (s *Service) widget(w http.ResponseWriter, r *http.Request) {
	userID := auth.ActorFrom(r.Context()).UserID
	out, err := s.WidgetFor(r.Context(), userID, r.URL.Query().Get("character"))
	if err != nil {
		s.fail(w, r, "widget", err, "could not read your widget just now")
		return
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, out)
}

// WidgetFor builds the snapshot for userID. want names the character (the widget's setting);
// an empty or unknown key reads the account's main, else its first character.
func (s *Service) WidgetFor(ctx context.Context, userID int64, want string) (Widget, error) {
	now := s.now()
	out := Widget{GeneratedAt: now.UTC().Truncate(time.Minute)}
	char, err := s.pickCharacter(ctx, userID, want)
	if err != nil || char == nil {
		return out, err
	}
	out.Character = widgetCharacter(*char)
	raid, err := s.Raids.WidgetRaidFor(ctx, userID, char.Key, now)
	if err != nil {
		return Widget{}, err
	}
	out.Guild, out.NextRaid, out.LastRaid = raid.Guild, raid.NextRaid, raid.LastRaid
	if s.Upgrades != nil {
		if out.Upgrades, err = s.Upgrades.WidgetUpgrades(ctx, char.Key); err != nil {
			return Widget{}, err
		}
		out.Upgrades = topRows(out.Upgrades)
	}
	if out.Live.SimRunsRunning, err = s.Sims.RunningServerRuns(ctx, userID); err != nil {
		return Widget{}, err
	}
	return out, nil
}

// pickCharacter is the wanted character when it is the caller's, else the main, else the
// first; nil for an account with no characters.
func (s *Service) pickCharacter(ctx context.Context, userID int64, want string) (*auth.Character, error) {
	chars, err := s.Accounts.Characters(ctx, userID)
	if err != nil || len(chars) == 0 {
		return nil, err
	}
	if c := findCharacter(chars, want); c != nil {
		return c, nil
	}
	u, err := s.Accounts.User(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("livesurface: user: %w", err)
	}
	if u.MainCharacterKey != nil {
		if c := findCharacter(chars, *u.MainCharacterKey); c != nil {
			return c, nil
		}
	}
	return &chars[0], nil
}

func findCharacter(chars []auth.Character, key string) *auth.Character {
	if key == "" {
		return nil
	}
	for i := range chars {
		if chars[i].Key == key {
			return &chars[i]
		}
	}
	return nil
}

func widgetCharacter(c auth.Character) *WidgetCharacter {
	return &WidgetCharacter{Key: c.Key, Name: c.Name, Class: optional(c.Class), SpecName: optional(c.Spec)}
}

// topRows keeps the first widgetRows rows of the Home answer's order; nil stays nil.
func topRows(u *WidgetUpgrades) *WidgetUpgrades {
	if u == nil || len(u.Rows) <= widgetRows {
		return u
	}
	trimmed := *u
	trimmed.Rows = append([]UpgradeRow(nil), u.Rows[:widgetRows]...)
	return &trimmed
}

func optional(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
