// api/internal/guilds/raids_report_digest.go
//
// The per-pull digest (guild_report_pulls, migration 0074): what the report needs from a fight's
// stored summary, read once. Deaths carry the engine's own killing blow (the last hit that
// overkilled, else the last hit) and the last hits before the death; consumables are the log
// coach's at-pull reading (logs/engine/mechanics/consumables), which needs the pull's character
// sheet (COMBATANT_INFO). With no sheet the log cannot say what was up at the pull, so the pull's
// consumables are null: unknown, never missing. No database here.
package guilds

import (
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/consumables"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// deathLastHits is how many hits before a death the report keeps, newest last.
const deathLastHits = 3

// DeathHit is one damage event on a raider before their death.
type DeathHit struct {
	AtMS     int64  `json:"at_ms"`
	SpellID  int64  `json:"spell_id"`
	Ability  string `json:"ability"`
	Source   string `json:"source"`
	Amount   int64  `json:"amount"`
	Overkill int64  `json:"overkill"`
}

// digestDeath is one death as stored: the log's unit name, mapped to a character at read time.
type digestDeath struct {
	Name        string     `json:"name"`
	Class       string     `json:"class"`
	AtMS        int64      `json:"at_ms"`
	KillingBlow *DeathHit  `json:"killing_blow"`
	Last        []DeathHit `json:"last"`
}

// digestConsumable is one raider's pull as stored. A category is true when its buff was up at
// the pull, false when the catalogue's complete list proves it absent, null when the log cannot
// say (an incomplete list proves presence only).
type digestConsumable struct {
	Name          string   `json:"name"`
	Known         bool     `json:"known"`
	Flask         *bool    `json:"flask"`
	Food          *bool    `json:"food"`
	WeaponEnchant *bool    `json:"weapon_enchant"`
	Potions       int64    `json:"potions"`
	Brought       []string `json:"brought"`
}

// pullDigest is one fight's digest. Consumables is nil when no raider of the pull has a
// character sheet.
type pullDigest struct {
	Deaths      []digestDeath
	Consumables []digestConsumable
}

// digestOf reads the report's digest out of one fight's summary.
func digestOf(sum summary.Summary) pullDigest {
	return pullDigest{Deaths: digestDeaths(sum.Deaths), Consumables: digestConsumables(sum)}
}

func digestDeaths(deaths []summary.Death) []digestDeath {
	out := make([]digestDeath, 0, len(deaths))
	for _, d := range deaths {
		dd := digestDeath{Name: d.Name, Class: d.Class, AtMS: d.AtMS, Last: []DeathHit{}}
		if d.KillingBlow != nil {
			kb := hitOf(*d.KillingBlow)
			dd.KillingBlow = &kb
		}
		last := d.Last
		if len(last) > deathLastHits {
			last = last[len(last)-deathLastHits:]
		}
		for _, h := range last {
			dd.Last = append(dd.Last, hitOf(h))
		}
		out = append(out, dd)
	}
	return out
}

func hitOf(r summary.DamageRef) DeathHit {
	return DeathHit{AtMS: r.AtMS, SpellID: r.SpellID, Ability: r.SpellName, Source: r.SourceName,
		Amount: r.Amount, Overkill: r.Overkill}
}

// digestConsumables reads every roster player's pull against their role's catalogue (the
// discipline catalogue when the role is unknown); nil when nobody's sheet is in the log.
func digestConsumables(sum summary.Summary) []digestConsumable {
	cat := consumables.Load()
	out := make([]digestConsumable, 0, len(sum.Roster))
	anyKnown := false
	for _, p := range sum.Roster {
		rc, ok := cat.For(p.Role)
		if !ok {
			rc = cat.AllRoles()
		}
		row := consumableOf(p.Name, rc, consumables.AtPull(sum, p.GUID, rc))
		anyKnown = anyKnown || row.Known
		out = append(out, row)
	}
	if !anyKnown {
		return nil
	}
	return out
}

func consumableOf(name string, rc consumables.RoleCatalogue, p consumables.Pull) digestConsumable {
	row := digestConsumable{Name: name, Known: p.Known, Potions: p.PotionUses, Brought: []string{}}
	if !p.Known {
		return row
	}
	row.Flask = categoryState(rc, p, consumables.CategoryFlask)
	row.Food = categoryState(rc, p, consumables.CategoryFood)
	row.WeaponEnchant = categoryState(rc, p, consumables.CategoryWeaponEnchant)
	for _, cat := range []string{consumables.CategoryFlask, consumables.CategoryFood, consumables.CategoryWeaponEnchant} {
		for _, e := range p.Present[cat] {
			row.Brought = append(row.Brought, e.Name)
		}
	}
	return row
}

func categoryState(rc consumables.RoleCatalogue, p consumables.Pull, category string) *bool {
	switch {
	case p.Has(category):
		v := true
		return &v
	case rc.MissedAtPull(p, category):
		v := false
		return &v
	}
	return nil
}
