// api/internal/guilds/raids_view.go
//
// The JSON views of the guild raid calendar (design/specs/2026-10-26-guild-raids.md §9): the
// raid page's RaidView, the Raids tab's planned list and zones, and the guild band's
// NextRaidView. Every view is built in Go from the store rows plus the guild's roster.
package guilds

import (
	"context"
	"sort"
	"strings"
	"time"
)

// RaidCounts are a raid's sign-up totals. Tanks, Healers and DPS count `in` rows only.
type RaidCounts struct {
	In      int `json:"in"`
	Late    int `json:"late"`
	Bench   int `json:"bench"`
	Out     int `json:"out"`
	Tanks   int `json:"tanks"`
	Healers int `json:"healers"`
	DPS     int `json:"dps"`
}

// SignupView is one sign-up as the raid page lists it. Own is true when the character
// belongs to the viewer's account.
type SignupView struct {
	CharacterKey string `json:"character_key"`
	Name         string `json:"name"`
	Class        string `json:"class"`
	Spec         string `json:"spec"`
	Role         string `json:"role"`
	State        string `json:"state"`
	Note         string `json:"note"`
	Own          bool   `json:"own"`
}

// RaidCharacter is a verified roster character with no sign-up yet.
type RaidCharacter struct {
	CharacterKey string `json:"character_key"`
	Name         string `json:"name"`
	Class        string `json:"class"`
	Spec         string `json:"spec"`
}

// ViewerCharacter is one of the viewer's characters in the guild. State is null until it
// signs up; Role is the sign-up's role, else the spec's default. An unverified character is
// listed (Verified false) but cannot sign up.
type ViewerCharacter struct {
	CharacterKey string  `json:"character_key"`
	Name         string  `json:"name"`
	Class        string  `json:"class"`
	Spec         string  `json:"spec"`
	Verified     bool    `json:"verified"`
	State        *string `json:"state"`
	Role         string  `json:"role"`
}

// RaidViewer is who is asking.
type RaidViewer struct {
	Characters []ViewerCharacter `json:"characters"`
}

// RaidDiscord is the guild channel's state on a raid view (Service.raidDiscord): Available
// when the deployment has connect on, Connected when the guild has a channel, PostDefault
// the guild's remembered "Post to Discord" box.
type RaidDiscord struct {
	Available   bool `json:"available"`
	Connected   bool `json:"connected"`
	PostDefault bool `json:"post_default"`
}

// RaidView is GET /v1/guilds/{id}/raids/{raid} and each row of the planned list.
type RaidView struct {
	ID              int64     `json:"id"`
	Zone            string    `json:"zone"`
	ZoneName        string    `json:"zone_name"`
	Title           string    `json:"title"`
	StartsAt        time.Time `json:"starts_at"`
	DurationMinutes int       `json:"duration_minutes"`
	Status          string    `json:"status"`
	Size            *int      `json:"size"`
	GuildTimezone   *string   `json:"guild_timezone"`
	NotesHTML       string    `json:"notes_html"`
	// Notes is the raw markdown the officer typed, for the edit form; members get it too (it
	// is the same text notes_html renders).
	Notes     string          `json:"notes"`
	Counts    RaidCounts      `json:"counts"`
	Signups   []SignupView    `json:"signups"`
	NotSigned []RaidCharacter `json:"not_signed"`
	Viewer    RaidViewer      `json:"viewer"`
	Discord   RaidDiscord     `json:"discord"`
	// PreDoneStatus is what a done raid returns to when its match is cleared; Match, Night and
	// Attendance are null until a log is matched (Attendance: officers every row, a member
	// their own characters').
	PreDoneStatus *string          `json:"pre_done_status"`
	Match         *MatchView       `json:"match"`
	Night         *NightView       `json:"night"`
	Attendance    []AttendanceView `json:"attendance"`
}

// ZoneView is one schedulable zone as the form lists it.
type ZoneView struct {
	Key  string `json:"key"`
	Name string `json:"name"`
	Size *int   `json:"size"`
}

// PlannedSection is the members-only half of GET /v1/guilds/{id}/raids. RaidsPage embeds it
// by pointer, so a nil section leaves all four keys out for everyone else (R1). Discord is the
// guild channel's state, so the create form can offer the box before any raid is planned.
type PlannedSection struct {
	Planned       []RaidView  `json:"planned"`
	Zones         []ZoneView  `json:"zones"`
	GuildTimezone *string     `json:"guild_timezone"`
	Discord       RaidDiscord `json:"discord"`
}

// NextRaidCounts are the band line's totals (the band count is In + Late).
type NextRaidCounts struct {
	In      int `json:"in"`
	Late    int `json:"late"`
	Tanks   int `json:"tanks"`
	Healers int `json:"healers"`
	DPS     int `json:"dps"`
}

// NextRaidOwn is the viewer's own sign-up for the next raid.
type NextRaidOwn struct {
	CharacterKey string `json:"character_key"`
	State        string `json:"state"`
	Role         string `json:"role"`
	// Note travels with the sign-up so a press on the band keeps it without a second read.
	Note string `json:"note"`
}

// NextRaidView is home.next_raid: the guild band's raid line.
type NextRaidView struct {
	ID       int64     `json:"id"`
	Zone     string    `json:"zone"`
	ZoneName string    `json:"zone_name"`
	Title    string    `json:"title"`
	StartsAt time.Time `json:"starts_at"`
	// Status is planned or locked (the band shows the next open raid; a locked one is inert).
	Status        string         `json:"status"`
	Size          *int           `json:"size"`
	GuildTimezone *string        `json:"guild_timezone"`
	Counts        NextRaidCounts `json:"counts"`
	Own           *NextRaidOwn   `json:"own"`
}

// zoneViews lists the curated zones for the form.
func zoneViews() []ZoneView {
	zones := Zones()
	out := make([]ZoneView, len(zones))
	for i, z := range zones {
		out[i] = ZoneView{Key: z.Key, Name: z.Name, Size: z.Size}
	}
	return out
}

// RaidViews builds the views of raids for viewerID in one pass: one roster read, one
// sign-ups read.
func (s *Store) RaidViews(ctx context.Context, guildID, viewerID int64, raids []Raid) ([]RaidView, error) {
	vc, err := s.loadRaidViewContext(ctx, guildID, viewerID)
	if err != nil {
		return nil, err
	}
	ids := make([]int64, len(raids))
	for i, r := range raids {
		ids[i] = r.ID
	}
	byRaid, err := s.SignupsByRaid(ctx, ids)
	if err != nil {
		return nil, err
	}
	out := make([]RaidView, len(raids))
	for i, r := range raids {
		if out[i], err = buildRaidView(r, byRaid[r.ID], vc); err != nil {
			return nil, err
		}
	}
	if err := s.fillMatchViews(ctx, vc, out); err != nil {
		return nil, err
	}
	return out, nil
}

func buildRaidView(r Raid, signups []Signup, vc raidViewContext) (RaidView, error) {
	zone := zoneOf(r.Zone)
	notesHTML, err := RenderNotes(r.Notes)
	if err != nil {
		return RaidView{}, err
	}
	v := RaidView{
		ID: r.ID, Zone: r.Zone, ZoneName: zone.Name, Title: r.Title, StartsAt: r.StartsAt.UTC(),
		DurationMinutes: r.DurationMinutes, Status: r.Status, Size: zone.Size, GuildTimezone: vc.timezone,
		NotesHTML: notesHTML, Notes: r.Notes, Counts: countSignups(signups),
		Signups: make([]SignupView, 0, len(signups)), NotSigned: []RaidCharacter{},
		Viewer: RaidViewer{Characters: []ViewerCharacter{}},
	}
	byKey := make(map[string]Signup, len(signups))
	for _, sg := range signups {
		byKey[sg.CharacterKey] = sg
		v.Signups = append(v.Signups, SignupView{
			CharacterKey: sg.CharacterKey, Name: sg.Name, Class: sg.Class, Spec: sg.Spec,
			Role: sg.Role, State: sg.State, Note: sg.Note, Own: vc.owns(sg),
		})
	}
	sortSignups(v.Signups)
	for _, c := range vc.roster {
		sg, signed := byKey[c.CharacterKey]
		if c.verified && !signed {
			v.NotSigned = append(v.NotSigned, c.RaidCharacter)
		}
		if c.userID == vc.viewerID {
			v.Viewer.Characters = append(v.Viewer.Characters, viewerCharacter(c, sg, signed))
		}
	}
	sortRaidCharacters(v.NotSigned)
	sort.SliceStable(v.Viewer.Characters, func(i, j int) bool {
		return v.Viewer.Characters[i].Name < v.Viewer.Characters[j].Name
	})
	return v, nil
}

func viewerCharacter(c raidRosterChar, sg Signup, signed bool) ViewerCharacter {
	vcx := ViewerCharacter{
		CharacterKey: c.CharacterKey, Name: c.Name, Class: c.Class, Spec: c.Spec,
		Verified: c.verified, Role: c.defaultRole(),
	}
	if signed {
		state := sg.State
		vcx.State, vcx.Role = &state, sg.Role
	}
	return vcx
}

// zoneOf is the curated zone, or a bare one named by its key for a zone no longer curated.
func zoneOf(key string) Zone {
	if z, ok := ZoneByKey(key); ok {
		return z
	}
	return Zone{Key: key, Name: key}
}

func countSignups(signups []Signup) RaidCounts {
	var c RaidCounts
	for _, sg := range signups {
		switch sg.State {
		case SignupIn:
			c.In++
			switch sg.Role {
			case RoleTank:
				c.Tanks++
			case RoleHealer:
				c.Healers++
			default:
				c.DPS++
			}
		case SignupLate:
			c.Late++
		case SignupBench:
			c.Bench++
		case SignupOut:
			c.Out++
		}
	}
	return c
}

// roleOrder and classOrder are the roster's sort keys: tanks, healers, DPS; classes in the
// game's own order.
var (
	roleOrder  = map[string]int{RoleTank: 0, RoleHealer: 1, RoleDPS: 2}
	classOrder = map[string]int{
		"warrior": 0, "paladin": 1, "hunter": 2, "rogue": 3, "priest": 4,
		"shaman": 5, "mage": 6, "warlock": 7, "druid": 8,
	}
)

// rank is a sort key's position, unknown values last.
func rank(order map[string]int, value string) int {
	if i, ok := order[strings.ToLower(value)]; ok {
		return i
	}
	return len(order)
}

func sortSignups(s []SignupView) {
	sort.SliceStable(s, func(i, j int) bool {
		if a, b := rank(roleOrder, s[i].Role), rank(roleOrder, s[j].Role); a != b {
			return a < b
		}
		if a, b := rank(classOrder, s[i].Class), rank(classOrder, s[j].Class); a != b {
			return a < b
		}
		return s[i].Name < s[j].Name
	})
}

func sortRaidCharacters(s []RaidCharacter) {
	sort.SliceStable(s, func(i, j int) bool {
		if a, b := rank(classOrder, s[i].Class), rank(classOrder, s[j].Class); a != b {
			return a < b
		}
		return s[i].Name < s[j].Name
	})
}

// tankSpecs and healerSpecs name specs by their talent tree, as the logs spell them
// (internal/spec); every other spec is DPS. Protection is a tank for warriors and paladins
// alike; Holy, Discipline and Restoration heal for every class that has them.
var (
	tankSpecs   = map[string]bool{"protection": true, "feral combat": true, "feral": true, "guardian": true}
	healerSpecs = map[string]bool{"holy": true, "discipline": true, "restoration": true}
)

// roleForSpec is the default sign-up role of a spec.
func roleForSpec(spec string) string {
	key := strings.ToLower(strings.TrimSpace(spec))
	switch {
	case tankSpecs[key]:
		return RoleTank
	case healerSpecs[key]:
		return RoleHealer
	default:
		return RoleDPS
	}
}

// NextRaidFor is the band line for viewerID: the guild's next raid, nil when none. Own is
// the sign-up of headerKey (the home's viewer character) when it has one, else the viewer's
// first own sign-up, else nil.
func (s *Store) NextRaidFor(ctx context.Context, guildID, viewerID int64, headerKey *string, now time.Time) (*NextRaidView, error) {
	raid, ok, err := s.NextRaid(ctx, guildID, now)
	if err != nil || !ok {
		return nil, err
	}
	vc, err := s.loadRaidOwnership(ctx, guildID, viewerID)
	if err != nil {
		return nil, err
	}
	signups, err := s.RaidSignups(ctx, raid.ID)
	if err != nil {
		return nil, err
	}
	zone := zoneOf(raid.Zone)
	c := countSignups(signups)
	return &NextRaidView{
		ID: raid.ID, Zone: raid.Zone, ZoneName: zone.Name, Title: raid.Title, StartsAt: raid.StartsAt.UTC(),
		Status: raid.Status, Size: zone.Size, GuildTimezone: vc.timezone,
		Counts: NextRaidCounts{In: c.In, Late: c.Late, Tanks: c.Tanks, Healers: c.Healers, DPS: c.DPS},
		Own:    ownSignup(signups, vc, headerKey),
	}, nil
}

func ownSignup(signups []Signup, vc raidViewContext, headerKey *string) *NextRaidOwn {
	var first *Signup
	for i := range signups {
		sg := signups[i]
		if !vc.owns(sg) {
			continue
		}
		if headerKey != nil && sg.CharacterKey == *headerKey {
			return &NextRaidOwn{CharacterKey: sg.CharacterKey, State: sg.State, Role: sg.Role, Note: sg.Note}
		}
		if first == nil {
			first = &signups[i]
		}
	}
	if first == nil {
		return nil
	}
	return &NextRaidOwn{CharacterKey: first.CharacterKey, State: first.State, Role: first.Role, Note: first.Note}
}
