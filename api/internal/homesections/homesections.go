// Package homesections composes the apps' home: its blocks and their order, computed for the
// viewer (design/specs/2026-11-07-mobile-redesign.md section 2.3, the flows spec 2026-11-08 section 4 and ruling R3), so the
// order changes without an app release. GET /v1/home answers the signed-in list; GET
// /v1/app/config answers the signed-out one. HOME_SECTIONS (a JSON override, like APP_FEATURES)
// switches each rule off, turns blocks off, reorders them and adds a standing notice.
package homesections

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strings"
)

// The block types an app may be sent. An app renders the types it knows, in order, and skips
// any other.
const (
	Head          = "head"
	Upgrades      = "upgrades"
	NextRaid      = "next_raid"
	RaidLive      = "raid_live"
	RaidSummary   = "raid_summary"
	Countdowns    = "countdowns"
	Feed          = "feed"
	GuildSims     = "guild_sims"
	GuildSimsCTA  = "guild_sims_cta"
	SpecWeek      = "spec_week"
	Movers        = "movers"
	Next          = "next"
	Notice        = "notice"
	HeadOut       = "head_out"
	ClassPicker   = "class_picker"
	HowToStartURL = "how_to_start_link"
)

// Styles: the countdowns full or one line; the next raid a card (raid day) or one line; the head
// compact (one line) when a raid block leads.
const (
	StyleFull    = "full"
	StyleLine    = "line"
	StyleCard    = "card"
	StyleCompact = "compact"
)

// The feed filters.
const (
	FilterAll   = "all"
	FilterSpec  = "spec"
	FilterGuild = "guild"
	FilterLogs  = "logs"
)

// The rules HOME_SECTIONS can switch off, by key (design/specs/2026-11-08-app-flows.md section 4):
// the three raid moments (raid day, raid live, the day after), the leveling order under 60, and
// the Join a guild row in place of the ladder.
const (
	RuleRaidMoments = "raid_moments"
	RuleLeveling    = "leveling"
	RuleNoGuildCTA  = "no_guild_cta"
)

const (
	maxLevel        = 60
	upgradesShown   = 3
	feedRows        = 7
	guildSimsShown  = 5
	maxNoticeLength = 280
)

var signedInTypes = []string{Head, Upgrades, NextRaid, RaidLive, RaidSummary, Countdowns, Feed, GuildSims, GuildSimsCTA, SpecWeek, Movers, Next, Notice}

var rules = []string{RuleRaidMoments, RuleLeveling, RuleNoGuildCTA}

//go:embed signed_out.json
var signedOutJSON []byte

// Section is one block. Fields a type does not use are left out of the JSON; an empty filter
// list ([], the signed-out feed with no segments) is kept, a nil one is left out.
type Section struct {
	Type        string   `json:"type"`
	Limit       int      `json:"limit,omitempty"`
	Style       string   `json:"style,omitempty"`
	Filters     []string `json:"filters,omitzero"`
	MoreFilters []string `json:"more_filters,omitzero"`
	// Selected is the feed filter the block opens on (the guild's on raid night); omitted, the first.
	Selected string `json:"selected,omitempty"`
	Text     string `json:"text,omitempty"`
	Href     string `json:"href,omitempty"`
}

// Override is HOME_SECTIONS. The zero value changes nothing.
type Override struct {
	// Rules switches a composition rule off with false; a rule not named stays on.
	Rules map[string]bool `json:"rules"`
	// Off removes these block types.
	Off []string `json:"off"`
	// Order puts the named types first, in this order; the others follow in their own order.
	Order []string `json:"order"`
	// Notice adds a standing notice under the head.
	Notice *NoticeOverride `json:"notice"`
	// SignedOut replaces the bundled signed-out list.
	SignedOut []Section `json:"signed_out"`
}

// NoticeOverride is the owner's standing message.
type NoticeOverride struct {
	Text string `json:"text"`
	Href string `json:"href"`
}

// ParseOverride reads HOME_SECTIONS; empty is no override. Errors name the problem, never the
// value.
func ParseOverride(raw string) (Override, error) {
	var o Override
	if strings.TrimSpace(raw) == "" {
		return o, nil
	}
	dec := json.NewDecoder(strings.NewReader(raw))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&o); err != nil {
		return Override{}, errors.New("HOME_SECTIONS must be a JSON object with rules, off, order, notice or signed_out")
	}
	for key := range o.Rules {
		if !slices.Contains(rules, key) {
			return Override{}, fmt.Errorf("HOME_SECTIONS names an unknown rule %q", key)
		}
	}
	for _, t := range append(slices.Clone(o.Off), o.Order...) {
		if !slices.Contains(signedInTypes, t) {
			return Override{}, fmt.Errorf("HOME_SECTIONS names an unknown block type %q", t)
		}
	}
	if o.Notice != nil && (strings.TrimSpace(o.Notice.Text) == "" || len(o.Notice.Text) > maxNoticeLength) {
		return Override{}, fmt.Errorf("HOME_SECTIONS notice text must be 1 to %d characters", maxNoticeLength)
	}
	if o.Notice != nil && o.Notice.Href != "" && !strings.HasPrefix(o.Notice.Href, "/") {
		return Override{}, errors.New("HOME_SECTIONS notice href must be a site path starting with /")
	}
	for _, s := range o.SignedOut {
		if strings.TrimSpace(s.Type) == "" {
			return Override{}, errors.New("HOME_SECTIONS signed_out sections each need a type")
		}
	}
	return o, nil
}

func (o Override) ruleOn(key string) bool {
	on, set := o.Rules[key]
	return !set || on
}

// SignedOut is the signed-out home: the override's list, else the bundled one
// (apps/shared/home-sections.json).
func SignedOut(o Override) []Section {
	if len(o.SignedOut) > 0 {
		return slices.Clone(o.SignedOut)
	}
	return slices.Clone(bundledSignedOut)
}

var bundledSignedOut = mustSignedOut()

func mustSignedOut() []Section {
	var file struct {
		Sections []Section `json:"sections"`
	}
	if err := json.Unmarshal(signedOutJSON, &file); err != nil || len(file.Sections) == 0 {
		panic("homesections: signed_out.json is not a section list")
	}
	return file.Sections
}

// apply drops the Off types, then puts the Order types first.
func (o Override) apply(in []Section) []Section {
	out := slices.DeleteFunc(slices.Clone(in), func(s Section) bool { return slices.Contains(o.Off, s.Type) })
	if len(o.Order) == 0 {
		return out
	}
	rank := func(s Section) int {
		if i := slices.Index(o.Order, s.Type); i >= 0 {
			return i
		}
		return len(o.Order)
	}
	slices.SortStableFunc(out, func(a, b Section) int { return rank(a) - rank(b) })
	return out
}
