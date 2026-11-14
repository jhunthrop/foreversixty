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

	"github.com/jhunthrop/foreversixty/api/internal/raidspec"
	"github.com/jhunthrop/foreversixty/api/internal/zoneicons"
)

// RaidCounts are a raid's sign-up totals. Tanks, Healers, DPS, Melee and Ranged count `in` rows
// only (Melee plus Ranged is DPS, split by raidspec.Group). Classes counts In plus Late by class,
// one key for every class among the sign-ups and the verified roster not yet signed, zeros
// included (spec 2026-11-05 §4.1); DPS stays for the readiness page and older clients.
type RaidCounts struct {
	In      int            `json:"in"`
	Late    int            `json:"late"`
	Bench   int            `json:"bench"`
	Out     int            `json:"out"`
	Tanks   int            `json:"tanks"`
	Healers int            `json:"healers"`
	DPS     int            `json:"dps"`
	Melee   int            `json:"melee"`
	Ranged  int            `json:"ranged"`
	Classes map[string]int `json:"classes"`
	// LateGroups counts Late rows by roster group, beside the In-only Tanks..Ranged, so the comp
	// against the target can read "Healers 6/8 (+1 late)" (spec 2026-11-05-web-layout §7.1).
	LateGroups RoleCounts `json:"late_groups"`
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
	// LateMinutes is a Late row's minutes, null when it gave none.
	LateMinutes *int16 `json:"late_minutes"`
	Own         bool   `json:"own"`
	// Group is the roster group the row fights in (tank, healer, melee or ranged), from
	// raidspec.Group; SpecKey is the catalogue key of its spec, null when it does not resolve.
	Group   string  `json:"group"`
	SpecKey *string `json:"spec_key"`
	// Also is the account's other answered characters, best place first (raids spec §3.7): the
	// row is the account's best place, and these show under it without counting anywhere.
	Also []AlsoSignup `json:"also"`
	// DiscordPickedBy is the Discord name of the Discord-only account that picked this character
	// from the guild's directory roster on Discord, until a Battle.net sign-in confirms it:
	// officers see it marked. Empty (absent) for every other character.
	DiscordPickedBy string `json:"discord_picked_by,omitempty"`
}

// AlsoSignup is another answered character of a row's account, as its Also line names it.
type AlsoSignup struct {
	CharacterKey string  `json:"character_key"`
	Name         string  `json:"name"`
	Class        string  `json:"class"`
	Spec         string  `json:"spec"`
	SpecKey      *string `json:"spec_key"`
	State        string  `json:"state"`
}

// RaidCharacter is a verified roster character with no sign-up yet. Group and SpecKey read as
// on SignupView, with the character's default role; DiscordPickedBy as on SignupView.
type RaidCharacter struct {
	DiscordPickedBy string  `json:"discord_picked_by,omitempty"`
	CharacterKey    string  `json:"character_key"`
	Name            string  `json:"name"`
	Class           string  `json:"class"`
	Spec            string  `json:"spec"`
	Group           string  `json:"group"`
	SpecKey         *string `json:"spec_key"`
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
	// RaidSpec is the picked raid spec (a specs.json key, "paladin-holy"); null until picked,
	// when the raid page asks "Which spec tonight?" (spec 2026-10-28 §P.10.7).
	RaidSpec *string `json:"raid_spec"`
	// SpecKey is the catalogue key of the spec the row prints, null when it does not resolve;
	// DefaultSpec is the spec the picker starts on: raid_spec, else the known spec's key, else null.
	SpecKey     *string `json:"spec_key"`
	DefaultSpec *string `json:"default_spec"`
}

// RaidViewer is who is asking.
type RaidViewer struct {
	Characters []ViewerCharacter `json:"characters"`
}

// RaidDiscord is the guild's Discord destination on a raid view (Service.raidDiscord):
// Available when the deployment has connect or the bot on, Connected when the guild's bot is
// active or it has a webhook channel, PostDefault and RemindDefault the guild's remembered
// "Post to Discord" and "Remind the day before" boxes. ChannelName is the bot's channel, null
// for a webhook or when not connected.
type RaidDiscord struct {
	Available     bool    `json:"available"`
	Connected     bool    `json:"connected"`
	PostDefault   bool    `json:"post_default"`
	RemindDefault bool    `json:"remind_default"`
	ChannelName   *string `json:"channel_name"`
}

// RaidDiscordPost is a raid's posting schedule (spec 2026-10-28 §P.8): when it posts or
// posted, whether it reminds, and when the reminder went.
type RaidDiscordPost struct {
	PostedAt   *time.Time `json:"posted_at"`
	PostAt     *time.Time `json:"post_at"`
	RemindedAt *time.Time `json:"reminded_at"`
	Remind     bool       `json:"remind"`
}

// RaidView is GET /v1/guilds/{id}/raids/{raid} and each row of the planned list.
type RaidView struct {
	ID       int64  `json:"id"`
	Zone     string `json:"zone"`
	ZoneName string `json:"zone_name"`
	// ZoneKind is raid, dungeon or event: the page's words follow it (guild events spec §3.5).
	ZoneKind string `json:"zone_kind"`
	// KeyItem is the key the run's door asks for, null when none.
	KeyItem         *KeyItem  `json:"key_item"`
	Title           string    `json:"title"`
	StartsAt        time.Time `json:"starts_at"`
	DurationMinutes int       `json:"duration_minutes"`
	Status          string    `json:"status"`
	Size            *int      `json:"size"`
	GuildTimezone   *string   `json:"guild_timezone"`
	// ZoneIcon is the zone's icon (zoneicons), a name under /data/{build}/icons.
	ZoneIcon string `json:"zone_icon"`
	// AnswerBy is the sign-up deadline; no guild stores one yet, so it is always null.
	AnswerBy  *time.Time `json:"answer_by"`
	NotesHTML string     `json:"notes_html"`
	// Notes is the raw markdown the officer typed, for the edit form; members get it too (it
	// is the same text notes_html renders).
	Notes     string          `json:"notes"`
	Counts    RaidCounts      `json:"counts"`
	Signups   []SignupView    `json:"signups"`
	NotSigned []RaidCharacter `json:"not_signed"`
	Viewer    RaidViewer      `json:"viewer"`
	Discord   RaidDiscord     `json:"discord"`
	// DiscordPost is the raid's posting schedule; null when it has none.
	DiscordPost *RaidDiscordPost `json:"discord_post"`
	// CompTarget is the raid's comp target (raids_comp_target.go), null for a zone with none.
	CompTarget *CompTarget `json:"comp_target"`
	// NotSignedRaiders (one per account that has not answered) and Ping are an officer's only
	// (raids_overview_view.go); both are left out for everyone else.
	NotSignedRaiders *[]NotSignedRaider `json:"not_signed_raiders,omitempty"`
	Ping             *RaidPingState     `json:"ping,omitempty"`
	// PreDoneStatus is what a done raid returns to when its match is cleared; Match, Night and
	// Attendance are null until a log is matched (Attendance: officers every row, a member
	// their own characters').
	PreDoneStatus *string          `json:"pre_done_status"`
	Match         *MatchView       `json:"match"`
	Night         *NightView       `json:"night"`
	Attendance    []AttendanceView `json:"attendance"`
	// Report is the raid-night report, null until a log is matched.
	Report *ReportView `json:"report"`
	// Cover is the buff and debuff cover of the sign-ups (raids_cover.go); null for a dungeon
	// or an event.
	Cover []CoverRow `json:"cover"`
	// Assignments is the raid leader's per-encounter picks (raid_assign_view.go); null for a
	// dungeon or an event.
	Assignments *RaidAssignments `json:"assignments"`
}

// ZoneView is one schedulable zone as the form lists it (guild events spec §3.1).
type ZoneView struct {
	Key      string     `json:"key"`
	Name     string     `json:"name"`
	Kind     string     `json:"kind"`
	Size     *int       `json:"size"`
	Opens    *time.Time `json:"opens"`
	MinLevel *int       `json:"min_level"`
	MaxLevel *int       `json:"max_level"`
	KeyItem  *KeyItem   `json:"key_item"`
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
	Melee   int `json:"melee"`
	Ranged  int `json:"ranged"`
}

// NextRaidOwn is the viewer's own sign-up for the next raid.
type NextRaidOwn struct {
	CharacterKey string `json:"character_key"`
	State        string `json:"state"`
	Role         string `json:"role"`
	// Note travels with the sign-up so a press on the band keeps it without a second read.
	Note string `json:"note"`
	// LateMinutes is a Late answer's minutes, null when it gave none.
	LateMinutes *int16 `json:"late_minutes"`
}

// NextRaidView is home.next_raid: the guild band's raid line.
type NextRaidView struct {
	ID       int64     `json:"id"`
	Zone     string    `json:"zone"`
	ZoneName string    `json:"zone_name"`
	ZoneKind string    `json:"zone_kind"`
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
		out[i] = ZoneView{
			Key: z.Key, Name: z.Name, Kind: z.Kind, Size: z.Size, Opens: z.Opens,
			MinLevel: z.MinLevel, MaxLevel: z.MaxLevel, KeyItem: z.KeyItem,
		}
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
	out, _, err := s.raidViewsIn(ctx, vc, raids)
	return out, err
}

// raidViewsIn is RaidViews over an already loaded context, with the sign-ups it read by raid.
func (s *Store) raidViewsIn(ctx context.Context, vc raidViewContext, raids []Raid) ([]RaidView, map[int64][]Signup, error) {
	ids := make([]int64, len(raids))
	for i, r := range raids {
		ids[i] = r.ID
	}
	byRaid, err := s.SignupsByRaid(ctx, ids)
	if err != nil {
		return nil, nil, err
	}
	out := make([]RaidView, len(raids))
	for i, r := range raids {
		if out[i], err = buildRaidView(r, byRaid[r.ID], vc); err != nil {
			return nil, nil, err
		}
	}
	if err := s.fillMatchViews(ctx, vc, out); err != nil {
		return nil, nil, err
	}
	if err := s.fillAssignments(ctx, vc, out, byRaid); err != nil {
		return nil, nil, err
	}
	return out, byRaid, nil
}

func buildRaidView(r Raid, signups []Signup, vc raidViewContext) (RaidView, error) {
	zone := zoneOf(r.Zone)
	notesHTML, err := RenderNotes(r.Notes)
	if err != nil {
		return RaidView{}, err
	}
	v := RaidView{
		ID: r.ID, Zone: r.Zone, ZoneName: zone.Name, ZoneKind: zone.Kind, KeyItem: zone.KeyItem,
		Title: r.Title, StartsAt: r.StartsAt.UTC(),
		DurationMinutes: r.DurationMinutes, Status: r.Status, Size: zone.Size, GuildTimezone: vc.timezone,
		NotesHTML: notesHTML, Notes: r.Notes, Signups: []SignupView{}, NotSigned: []RaidCharacter{}, ZoneIcon: zoneicons.For(r.Zone),
		Viewer: RaidViewer{Characters: []ViewerCharacter{}},
	}
	byKey := make(map[string]Signup, len(signups))
	answered := make(map[int64]bool, len(signups))
	for _, sg := range signups {
		byKey[sg.CharacterKey] = sg
		answered[sg.SignedBy] = true
	}
	shown, also := onePerAccount(signups)
	v.Counts = countSignups(shown)
	picked := make(map[string]string, len(vc.roster))
	for _, c := range vc.roster {
		picked[c.CharacterKey] = c.DiscordPickedBy
	}
	for _, sg := range shown {
		v.Signups = append(v.Signups, SignupView{
			CharacterKey: sg.CharacterKey, Name: sg.Name, Class: sg.Class, Spec: sg.Spec,
			Role: sg.Role, State: sg.State, Note: sg.Note, LateMinutes: sg.LateMinutes, Own: vc.owns(sg),
			Group: raidspec.Group(sg.Class, sg.Spec, sg.Role), SpecKey: specKey(sg.Class, sg.Spec, sg.Role),
			Also: alsoSignups(also[sg.SignedBy]), DiscordPickedBy: picked[sg.CharacterKey],
		})
	}
	sortSignups(v.Signups)
	v.Cover = raidCover(zone.Kind, v.Signups)
	for _, c := range vc.roster {
		sg, signed := byKey[c.CharacterKey]
		// An account that has answered with any character is not waiting on an answer.
		if c.verified && !answered[c.userID] {
			v.NotSigned = append(v.NotSigned, c.notSigned())
		}
		if c.userID == vc.viewerID {
			v.Viewer.Characters = append(v.Viewer.Characters, viewerCharacter(c, sg, signed))
		}
	}
	sortRaidCharacters(v.NotSigned)
	v.Counts.Classes = classCounts(v.Signups, v.NotSigned)
	sort.SliceStable(v.Viewer.Characters, func(i, j int) bool {
		return v.Viewer.Characters[i].Name < v.Viewer.Characters[j].Name
	})
	return v, nil
}

// placeRank orders an account's places: In over Late over Bench over Out (accountPlaces' order).
var placeRank = map[string]int{SignupIn: 0, SignupLate: 1, SignupBench: 2, SignupOut: 3}

// onePerAccount is each account's best sign-up, in signups' order: a player who switched from
// their main to an alt (R3 benched the main) is one In, not an In and a Bench. Ties keep the
// first row. The account's other rows come back by account, best place first, for the row's
// Also line (raids spec §3.7).
func onePerAccount(signups []Signup) ([]Signup, map[int64][]Signup) {
	best := make(map[int64]int, len(signups))
	for i, sg := range signups {
		j, seen := best[sg.SignedBy]
		if !seen || placeRank[sg.State] < placeRank[signups[j].State] {
			best[sg.SignedBy] = i
		}
	}
	out := make([]Signup, 0, len(best))
	also := map[int64][]Signup{}
	for i, sg := range signups {
		if best[sg.SignedBy] == i {
			out = append(out, sg)
		} else {
			also[sg.SignedBy] = append(also[sg.SignedBy], sg)
		}
	}
	for _, rows := range also {
		sort.SliceStable(rows, func(i, j int) bool {
			if a, b := placeRank[rows[i].State], placeRank[rows[j].State]; a != b {
				return a < b
			}
			return rows[i].Name < rows[j].Name
		})
	}
	return out, also
}

// alsoSignups is an account's other rows as the Also line names them; empty, never null.
func alsoSignups(rows []Signup) []AlsoSignup {
	out := make([]AlsoSignup, len(rows))
	for i, sg := range rows {
		out[i] = AlsoSignup{
			CharacterKey: sg.CharacterKey, Name: sg.Name, Class: sg.Class, Spec: sg.Spec,
			SpecKey: specKey(sg.Class, sg.Spec, sg.Role), State: sg.State,
		}
	}
	return out
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
	vcx.SpecKey = specKey(c.Class, c.Spec, vcx.Role)
	vcx.DefaultSpec = vcx.SpecKey
	if c.raidSpec != "" {
		spec := c.raidSpec
		vcx.RaidSpec, vcx.DefaultSpec = &spec, &spec
	}
	return vcx
}

// notSigned is the roster character as Not signed up lists it, its group read with its default role.
func (c raidRosterChar) notSigned() RaidCharacter {
	out := c.RaidCharacter
	role := c.defaultRole()
	out.Group, out.SpecKey = raidspec.Group(c.Class, c.Spec, role), specKey(c.Class, c.Spec, role)
	return out
}

// specKey is the catalogue key of a class, spec name and role, nil when the catalogue has none.
func specKey(class, spec, role string) *string {
	sp, ok := raidspec.Resolve(class, spec, role)
	if !ok {
		return nil
	}
	return &sp.Key
}

// classCounts is In plus Late by class: one key per class among the sign-ups and the roster not
// yet signed, so a class the guild plays reads 0 rather than going missing (spec 2026-11-05 §3.1).
func classCounts(signups []SignupView, notSigned []RaidCharacter) map[string]int {
	out := map[string]int{}
	for _, c := range notSigned {
		if c.Class != "" {
			out[strings.ToLower(c.Class)] += 0
		}
	}
	for _, sg := range signups {
		if sg.Class == "" {
			continue
		}
		class := strings.ToLower(sg.Class)
		out[class] += 0
		if sg.State == SignupIn || sg.State == SignupLate {
			out[class]++
		}
	}
	return out
}

// zoneOf is the curated zone, or a bare one named by its key for a zone no longer curated.
func zoneOf(key string) Zone {
	if z, ok := ZoneByKey(key); ok {
		return z
	}
	kind := ""
	if strings.HasPrefix(key, raidKeyPrefix) {
		kind = ZoneRaid
	}
	return Zone{Key: key, Name: key, Kind: kind}
}

func countSignups(signups []Signup) RaidCounts {
	var c RaidCounts
	for _, sg := range signups {
		switch sg.State {
		case SignupIn:
			c.In++
			switch raidspec.Group(sg.Class, sg.Spec, sg.Role) {
			case raidspec.GroupTank:
				c.Tanks++
			case raidspec.GroupHealer:
				c.Healers++
			case raidspec.GroupMelee:
				c.DPS++
				c.Melee++
			default:
				c.DPS++
				c.Ranged++
			}
		case SignupLate:
			c.Late++
			c.LateGroups.add(raidspec.Group(sg.Class, sg.Spec, sg.Role))
		case SignupBench:
			c.Bench++
		case SignupOut:
			c.Out++
		}
	}
	return c
}

// roleOrder is the roster's first sort key: tanks, healers, DPS; classes then follow the game's
// own order (raidspec.ClassOrder, the one order the site and the bot share).
var roleOrder = map[string]int{RoleTank: 0, RoleHealer: 1, RoleDPS: 2}

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
		if a, b := raidspec.ClassRank(s[i].Class), raidspec.ClassRank(s[j].Class); a != b {
			return a < b
		}
		return s[i].Name < s[j].Name
	})
}

func sortRaidCharacters(s []RaidCharacter) {
	sort.SliceStable(s, func(i, j int) bool {
		if a, b := raidspec.ClassRank(s[i].Class), raidspec.ClassRank(s[j].Class); a != b {
			return a < b
		}
		return s[i].Name < s[j].Name
	})
}

// tankSpecs and healerSpecs name specs by their talent tree, as the logs spell them
// (internal/spec); every other spec is DPS. Protection is a tank for warriors and paladins
// alike; Holy, Discipline and Restoration heal for every class that has them.
var (
	tankSpecs   = map[string]bool{"protection": true, "feral bear": true, "feral combat": true, "feral": true, "guardian": true}
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

// NextRaidFor is the band line for viewerID: the guild's next raid, dungeon run or guild event
// (guild events spec §3.4), nil when none. Own is
// the sign-up of headerKey (the home's viewer character) when it has one, else the viewer's
// first own sign-up, else nil.
func (s *Store) NextRaidFor(ctx context.Context, guildID, viewerID int64, headerKey *string, now time.Time) (*NextRaidView, error) {
	raid, ok, err := s.NextPlanned(ctx, guildID, now)
	if err != nil || !ok {
		return nil, err
	}
	return s.nextRaidView(ctx, guildID, viewerID, headerKey, raid)
}

// NextRaidFirstFor is /raid's night for viewerID: the next planned raid of kind raid, the one
// Readiness names, else the next dungeon run or guild event; nil when none (bot spec §6.2,
// amended 2026-11-05).
func (s *Store) NextRaidFirstFor(ctx context.Context, guildID, viewerID int64, now time.Time) (*NextRaidView, error) {
	raid, ok, err := s.NextRaid(ctx, guildID, now)
	if err != nil {
		return nil, err
	}
	if !ok {
		return s.NextRaidFor(ctx, guildID, viewerID, nil, now)
	}
	return s.nextRaidView(ctx, guildID, viewerID, nil, raid)
}

// nextRaidView is raid as the band line shows it to viewerID.
func (s *Store) nextRaidView(ctx context.Context, guildID, viewerID int64, headerKey *string, raid Raid) (*NextRaidView, error) {
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
		ID: raid.ID, Zone: raid.Zone, ZoneName: zone.Name, ZoneKind: zone.Kind, Title: raid.Title,
		StartsAt: raid.StartsAt.UTC(), Status: raid.Status, Size: zone.Size, GuildTimezone: vc.timezone,
		Counts: NextRaidCounts{
			In: c.In, Late: c.Late, Tanks: c.Tanks, Healers: c.Healers, DPS: c.DPS, Melee: c.Melee, Ranged: c.Ranged,
		},
		Own: ownSignup(signups, vc, headerKey),
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
			return nextRaidOwn(sg)
		}
		if first == nil {
			first = &signups[i]
		}
	}
	if first == nil {
		return nil
	}
	return nextRaidOwn(*first)
}

func nextRaidOwn(sg Signup) *NextRaidOwn {
	return &NextRaidOwn{CharacterKey: sg.CharacterKey, State: sg.State, Role: sg.Role, Note: sg.Note, LateMinutes: sg.LateMinutes}
}
