// api/internal/guilds/character_prep.go
//
// One character's readiness read straight from its addon export, outside any guild roster: the
// applicant sheet (applicant_sheet.go) and the ready check (ready_check.go) grade a character the
// Readiness board and the D12 checklist would grade, through the same applyExport and
// computeReadiness, without a second implementation.
package guilds

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/bis"
	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/fs1"
	"github.com/jhunthrop/foreversixty/api/internal/raidspec"
)

// characterPrep is one character's export-read readiness inputs and verdicts.
type characterPrep struct {
	row        RosterRow
	exportedAt *time.Time
	band       bis.Band
	hasBand    bool
	readiness  CharacterReadiness
}

// characterPrepSQL is the character, the spec of its newest public fight, and its export.
var characterPrepSQL = `
select c.key, c.name, lower(coalesce(nullif(c.class, ''), fm.class, '')), coalesce(nullif(fm.spec, ''), ''),
       c.level, coalesce(c.faction, ''), ae.export, ae.captured_at
from characters c
left join lateral (` + newestEvidenceFight("c.key") + `) fm on true
left join addon_exports ae on ae.character_key = c.key
where c.key = $1`

// characterPrepOf reads key's prep as consent shares it (gear or gear_bags); found is false for
// a character the site does not have.
func (s *Store) characterPrepOf(ctx context.Context, key, consent string) (characterPrep, bool, error) {
	var p characterPrep
	var export *string
	err := s.Pool.QueryRow(ctx, characterPrepSQL, key).Scan(&p.row.CharacterKey, &p.row.Name, &p.row.className,
		&p.row.specStr, &p.row.Level, &p.row.faction, &export, &p.exportedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return p, false, nil
	}
	if err != nil {
		return p, false, fmt.Errorf("guilds: character prep %s: %w", key, err)
	}
	p.row.Consent, p.row.Verified = consent, true
	if export != nil {
		s.applyExport(&p.row, *export)
		if p.row.specStr == "" {
			p.row.specStr = exportSpecName(p.row.className, *export)
		}
	}
	p.band, p.hasBand = s.loadBandFor(p.row.className, p.row.specStr, p.row.faction)
	p.readiness = computeReadiness(readinessInputOf(p.row, p.band, p.hasBand, nil))
	return p, true, nil
}

// exportSpecName is the spec the export's talents name, by its catalogue name; "" when they
// name none (character.SpecFromAsk) or the data has no such spec.
func exportSpecName(classSlug, export string) string {
	d, ok := fs1.Decode(export)
	if !ok {
		return ""
	}
	key, source := character.SpecFor(strings.ToLower(classSlug), "", d.Talents.Points)
	if source == character.SpecFromAsk {
		return ""
	}
	if sp, ok := raidspec.ByKey(key); ok {
		return sp.Name
	}
	return ""
}

// nextZone is the zone of the guild's next raid, else its next run; nil with none planned.
func (s *Store) nextZone(ctx context.Context, guildID int64) (*Zone, error) {
	now := time.Now()
	raid, found, err := s.NextRaid(ctx, guildID, now)
	if err == nil && !found {
		raid, found, err = s.NextPlanned(ctx, guildID, now)
	}
	if err != nil || !found {
		return nil, err
	}
	zone := zoneOf(raid.Zone)
	return &zone, nil
}

// zoneAttunementItem is the item a zone's attunement names: its key, else its chain's item.
func zoneAttunementItem(zone Zone) *KeyItem {
	if zone.KeyItem != nil {
		return zone.KeyItem
	}
	if a, ok := attunementFor(zone); ok {
		return a.Item
	}
	return nil
}
