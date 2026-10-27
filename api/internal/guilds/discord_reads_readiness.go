// api/internal/guilds/discord_reads_readiness.go
//
// The read behind the Discord bot's /readiness (spec 2026-10-27-discord-bot §3.3.3): the
// readiness board cut to the guild's next raid (the Readiness tab's next-raid filter, Store.
// Readiness), summarised. It is an officer view, gated on every call as the officer routes are:
// rank, the guild plan and the claim freeze (officerGateFor), then the Raid guild tools the
// board belongs to (launch pricing §3.2.8, requireGuildPlan's feature).
package guilds

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
)

const (
	// readinessRowsShown is how many failing rows the answer lists before "+n more".
	readinessRowsShown = 10
	readinessLineJoin  = " · "
)

// ReadinessAnswer is the next raid's board for an officer: Signed is who signed In or Late,
// Failing who has something to fix, Rows the first failing rows in the board's words (most
// checks first; member text unescaped), More the failing rows not listed. Raid is nil when
// no raid is planned.
type ReadinessAnswer struct {
	Raid    *Raid
	Signed  int
	Failing int
	Rows    []string
	More    int
}

// NextRaidReadiness is the next raid's readiness for viewerID. It refuses with ErrNotOfficer,
// ErrPlanRequired, ErrRaidGuildRequired or ErrClaimContested.
func (s *Service) NextRaidReadiness(ctx context.Context, guildID, viewerID int64, now time.Time) (ReadinessAnswer, error) {
	if err := s.readinessGate(ctx, guildID, viewerID); err != nil {
		return ReadinessAnswer{}, err
	}
	raid, found, err := s.Store.NextRaid(ctx, guildID, now)
	if err != nil || !found {
		return ReadinessAnswer{}, err
	}
	view, err := s.Store.Readiness(ctx, guildID, true, &raid)
	if err != nil {
		return ReadinessAnswer{}, fmt.Errorf("guilds: bot readiness: %w", err)
	}
	return readinessAnswerOf(raid, view), nil
}

// readinessGate is the officer routes' gate (officerGateFor), then the Raid guild tools.
func (s *Service) readinessGate(ctx context.Context, guildID, viewerID int64) error {
	if err := s.officerGateFor(ctx, guildID, viewerID); err != nil {
		if errors.Is(err, ErrNotYourCharacter) {
			return ErrNotOfficer
		}
		return err
	}
	ok, _, err := s.Entitlements.CanForGuild(ctx, guildID, viewerID, entitlements.FeatureRaidGuildTools)
	if err != nil {
		return fmt.Errorf("guilds: raid guild tools: %w", err)
	}
	if !ok {
		return ErrRaidGuildRequired
	}
	return nil
}

func readinessAnswerOf(raid Raid, view ReadinessView) ReadinessAnswer {
	out := ReadinessAnswer{Raid: &raid, Signed: len(view.Rows), Rows: []string{}}
	for _, row := range view.Rows {
		if row.Failing == 0 && !row.NoExport {
			continue
		}
		out.Failing++
		if len(out.Rows) < readinessRowsShown {
			out.Rows = append(out.Rows, readinessLine(row, raid.StartsAt))
			continue
		}
		out.More++
	}
	return out
}

// readinessLine is one failing row in the board's own cell words: the name, then each failing
// check (gear gap, enchants, consumables, talent points, a stale sync), or No export yet.
func readinessLine(row ReadinessRow, raidStart time.Time) string {
	parts := []string{row.Name}
	if row.NoExport {
		return strings.Join(append(parts, "No export yet"), readinessLineJoin)
	}
	if g := row.GearGap; g != nil && g.Upgrades > 0 {
		parts = append(parts, gearGapWords(*g))
	}
	if row.Enchants.Checked && len(row.Enchants.MissingSlots) > 0 {
		labels := make([]string, len(row.Enchants.MissingSlots))
		for i, slot := range row.Enchants.MissingSlots {
			labels[i] = enchantableSlotLabels[slot]
		}
		parts = append(parts, "No enchant: "+strings.Join(labels, ", "))
	}
	if row.Consumables.State == "short" {
		parts = append(parts, "Short")
	}
	if n := row.TalentPointsUnspent; n > 0 {
		parts = append(parts, fmt.Sprintf("%d unspent talent point%s", n, plural(n)))
	}
	if row.SyncedBeforeCutoff != nil && !*row.SyncedBeforeCutoff {
		days := int(raidStart.Sub(row.LoggedAt).Hours() / hoursPerDay)
		parts = append(parts, fmt.Sprintf("Synced %d day%s before the raid", days, plural(days)))
	}
	return strings.Join(parts, readinessLineJoin)
}

// gearGapWords is the board's Gear gap cell (web readiness-view.ts gearGapLabel) for a gap.
func gearGapWords(g GearGapView) string {
	upgrades := fmt.Sprintf("%d upgrade%s", g.Upgrades, plural(g.Upgrades))
	gain := int(g.GainDps + 0.5)
	switch {
	case gain <= 0 && g.NotSimChecked > 0:
		return upgrades + readinessLineJoin + "not simmed yet"
	case g.NotSimChecked > 0:
		return fmt.Sprintf("%s · +%d DPS · %d not simmed yet", upgrades, gain, g.NotSimChecked)
	default:
		return fmt.Sprintf("%s · +%d DPS", upgrades, gain)
	}
}
