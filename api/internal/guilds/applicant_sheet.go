// api/internal/guilds/applicant_sheet.go
//
// The applicant sheet's facts on each officer application (spec 2026-11-08-app-flows F3,
// Applicant): the Forever Score, the latest sim, Raid ready (attunement to the guild's next
// raid, enchants, the gap to the spec's 60 raid list, all graded from the export by
// characterPrepOf) and the guilds the character logged raids with. Applying shares the gear
// with the guild's officers; the bags stay private.
package guilds

import (
	"context"
	"fmt"
	"time"
)

// applicantHistoryMax bounds the guilds a sheet lists.
const applicantHistoryMax = 5

// The attunement states a sheet reads.
const (
	attuneOK      = "ok"
	attuneMiss    = "miss"
	attuneUnknown = "unknown"
)

// ApplicantScorer reads a character's Forever Score (rating.Service through an adapter); nil
// leaves the sheet's score null.
type ApplicantScorer interface {
	ScoreOf(ctx context.Context, key string) (*int, error)
}

// ApplicantSheet is what an officer reads before Accept. Every part is null when its source has
// nothing (no export, no sim, no planned raid).
type ApplicantSheet struct {
	Score      *int                 `json:"score"`
	Sim        *ApplicantSim        `json:"sim"`
	Attunement *ApplicantAttunement `json:"attunement"`
	Enchants   *ApplicantEnchants   `json:"enchants"`
	BisGap     *ApplicantBisGap     `json:"bis_gap"`
	History    []ApplicantGuild     `json:"history"`
	ExportedAt *time.Time           `json:"exported_at"`
}

// ApplicantSim is the latest finished sim: its figure and what it counts (dps, hps, tps).
type ApplicantSim struct {
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

// ApplicantAttunement is the attunement to the guild's next raid: ok, miss or unknown (the
// export cannot say), and the key or chain item it names.
type ApplicantAttunement struct {
	ZoneName string   `json:"zone_name"`
	State    string   `json:"state"`
	Item     *KeyItem `json:"item"`
}

// ApplicantEnchants is the enchantable slots with an item, how many carry an enchant, and the
// ones that do not.
type ApplicantEnchants struct {
	Slots     int      `json:"slots"`
	Enchanted int      `json:"enchanted"`
	Missing   []string `json:"missing"`
}

// ApplicantBisGap is the pieces short of the list it names ("Holy Priest, 60 raid").
type ApplicantBisGap struct {
	Pieces int    `json:"pieces"`
	List   string `json:"list"`
}

// ApplicantGuild is a guild the character logged raids with: the raids and the last one.
type ApplicantGuild struct {
	GuildName  string    `json:"guild_name"`
	Raids      int       `json:"raids"`
	LastRaidAt time.Time `json:"last_raid_at"`
}

// simUnits is what a sim's role counts.
var simUnits = map[string]string{RoleDPS: gainUnitDPS, RoleHealer: gainUnitHPS, RoleTank: gainUnitTPS}

// applicantSheet is the sheet for the application's character in guildID, with zone the
// guild's next raid's (nil: none planned).
func (s *Store) applicantSheet(ctx context.Context, guildID int64, a Application, zone *Zone) (ApplicantSheet, error) {
	sh := ApplicantSheet{History: []ApplicantGuild{}}
	var err error
	if sh.Sim, err = s.latestSim(ctx, a.Character.CharacterKey); err != nil {
		return sh, err
	}
	if sh.History, err = s.raidHistory(ctx, a.Character.CharacterKey, guildID); err != nil {
		return sh, err
	}
	p, found, err := s.characterPrepOf(ctx, a.Character.CharacterKey, "gear")
	if err != nil || !found {
		return sh, err
	}
	sh.ExportedAt = p.exportedAt
	if zone != nil {
		sh.Attunement = applicantAttunement(*zone, p.row)
	}
	cr := p.readiness
	if cr.EnchantsChecked {
		slots := equippedEnchantable(p.row.gear)
		missing := cr.MissingEnchantSlots
		if missing == nil {
			missing = []string{}
		}
		sh.Enchants = &ApplicantEnchants{Slots: slots, Enchanted: slots - len(missing), Missing: missing}
	}
	if cr.GearChecked {
		sh.BisGap = &ApplicantBisGap{Pieces: cr.GearGap.Upgrades, List: bisListName(p.row)}
	}
	return sh, nil
}

// applicantAttunement is the zone's attunement reading as the sheet words it; nil for a zone
// that asks for none.
func applicantAttunement(zone Zone, row RosterRow) *ApplicantAttunement {
	r := attunementOf(zone, row)
	if r.State == "" {
		return nil
	}
	state := attuneUnknown
	switch r.State {
	case CheckOK:
		state = attuneOK
	case CheckFail:
		state = attuneMiss
	}
	return &ApplicantAttunement{ZoneName: zone.Name, State: state, Item: zoneAttunementItem(zone)}
}

// equippedEnchantable counts the enchantable slots with an item in them.
func equippedEnchantable(gear map[string]int) int {
	n := 0
	for _, slot := range enchantableSlots {
		if _, ok := gear[slot]; ok {
			n++
		}
	}
	return n
}

// bisListName names the list the gap is measured against: the spec and the 60 raid band.
func bisListName(row RosterRow) string {
	name := row.specStr
	if sp, ok := resolveSpec(&row.className, &row.specStr, ""); ok {
		name = sp.Name
	}
	return fmt.Sprintf("%s, %d raid", name, bisBandLevel)
}

// latestSim is the character's newest finished sim; nil with none.
func (s *Store) latestSim(ctx context.Context, key string) (*ApplicantSim, error) {
	rows, err := s.Pool.Query(ctx, `select sm.dps_mean, cs.role from character_sims cs
		join sims sm on sm.id = cs.sim_id and sm.state = 'done'
		where cs.character_key = $1
		order by coalesce(sm.finished_at, sm.created_at) desc limit 1`, key)
	if err != nil {
		return nil, fmt.Errorf("guilds: latest sim %s: %w", key, err)
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, rows.Err()
	}
	var sim ApplicantSim
	var role string
	if err := rows.Scan(&sim.Value, &role); err != nil {
		return nil, fmt.Errorf("guilds: latest sim %s: %w", key, err)
	}
	sim.Unit = simUnits[role]
	if sim.Unit == "" {
		sim.Unit = gainUnitDPS
	}
	return &sim, nil
}

// raidHistory is the other guilds the character logged raids with in public, unflagged
// reports, newest first.
func (s *Store) raidHistory(ctx context.Context, key string, guildID int64) ([]ApplicantGuild, error) {
	rows, err := s.Pool.Query(ctx, `
		select g.name, count(distinct r.id), max(r.created_at)
		from fight_metrics m join reports r on r.id = m.report_id join guilds g on g.id = r.guild_id
		where m.player_key = $1 and r.guild_id <> $2 and r.visibility = 'public' and r.flagged is null
		group by g.id, g.name order by max(r.created_at) desc limit $3`, key, guildID, applicantHistoryMax)
	if err != nil {
		return nil, fmt.Errorf("guilds: raid history %s: %w", key, err)
	}
	defer rows.Close()
	out := []ApplicantGuild{}
	for rows.Next() {
		var h ApplicantGuild
		if err := rows.Scan(&h.GuildName, &h.Raids, &h.LastRaidAt); err != nil {
			return nil, fmt.Errorf("guilds: raid history %s: %w", key, err)
		}
		out = append(out, h)
	}
	return out, rows.Err()
}
