// Package homesections composes the apps' home: its blocks and their order, computed for the
// viewer (design/specs/2026-11-07-mobile-redesign.md section 2.3, rulings R0, R2, R3 and R4), so the
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

// The countdowns' two styles.
const (
	StyleFull = "full"
	StyleLine = "line"
)

// The feed filters.
const (
	FilterAll   = "all"
	FilterSpec  = "spec"
	FilterGuild = "guild"
	FilterLogs  = "logs"
)

// The rules HOME_SECTIONS can switch off, by key.
const (
	RuleUnder60Line    = "under_60_line"
	RuleNoGuildCTA     = "no_guild_cta"
	RuleGuildRaidFirst = "guild_raid_first"
)

const (
	maxLevel        = 60
	upgradesShown   = 3
	feedRows        = 7
	guildSimsShown  = 5
	maxNoticeLength = 280
)

var signedInTypes = []string{Head, Upgrades, Countdowns, Feed, GuildSims, GuildSimsCTA, SpecWeek, Movers, Next, Notice}

var rules = []string{RuleUnder60Line, RuleNoGuildCTA, RuleGuildRaidFirst}

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
	Text        string   `json:"text,omitempty"`
	Href        string   `json:"href,omitempty"`
}

// View is what the signed-in composition reads about the viewer's character.
type View struct {
	Level int
	// InGuild: the character is a verified member of a guild.
	InGuild bool
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

// SignedIn is the signed-in home for v under o (rulings R0, R2 and R4): in a guild the next raid and
// your sign-up come second, before the upgrades and the feed; without one the first-raids countdown
// follows the feed and the Join a guild row replaces the ladder; under 60 the countdown is one line.
func SignedIn(v View, o Override) []Section {
	inGuild := v.InGuild || !o.ruleOn(RuleNoGuildCTA)
	filters := []string{FilterAll, FilterSpec}
	guild := Section{Type: GuildSimsCTA}
	if inGuild {
		filters = append(filters, FilterGuild)
		guild = Section{Type: GuildSims, Limit: guildSimsShown}
	}
	countdowns := Section{Type: Countdowns, Style: StyleFull}
	if v.Level < maxLevel && o.ruleOn(RuleUnder60Line) {
		countdowns.Style = StyleLine
	}
	head := Section{Type: Head}
	upgrades := Section{Type: Upgrades, Limit: upgradesShown}
	feed := Section{Type: Feed, Filters: filters, MoreFilters: []string{FilterLogs}, Limit: feedRows}
	rest := []Section{guild, {Type: SpecWeek}, {Type: Movers}, {Type: Next}}

	out := slices.Concat([]Section{head, upgrades, feed, countdowns}, rest)
	if v.InGuild && o.ruleOn(RuleGuildRaidFirst) {
		out = slices.Concat([]Section{head, countdowns, upgrades, feed}, rest)
	}
	if o.Notice != nil {
		out = slices.Insert(out, 1, Section{Type: Notice, Text: o.Notice.Text, Href: o.Notice.Href})
	}
	return o.apply(out)
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
