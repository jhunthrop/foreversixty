// api/internal/bnetimport/build.go
package bnetimport

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/bnetapi"
	"github.com/jhunthrop/foreversixty/api/internal/bnetbuild"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// buildAndWriteExport runs bnetbuild.Encode over a character's freshly-captured profile,
// equipment and specializations (spec
// docs/superpowers/specs/2026-09-22-battlenet-first-design.md §2.3) and, on success,
// upserts addon_exports with the newest-captured_at-wins rule (spec §2.4). It is a no-op —
// never an error — whenever any of the three captures came back empty (a private profile, a
// 403/404 equipment or specializations read — a Season of Discovery character has none of
// them) or this process holds no active-build data tables: "all three answered" is the
// gate, and honesty (spec §1.3) means a character this cannot build for keeps whatever
// export it already had.
func (s *Service) buildAndWriteExport(ctx context.Context, tx pgx.Tx, userID int64, key, region, ruleset string,
	profile bnetapi.CharacterProfile, rawEquipment, rawSpecializations json.RawMessage) error {
	if rawEquipment == nil || rawSpecializations == nil || s.Tables.Trees == nil {
		return nil
	}
	classID, ok := classIDFor(s.Tables.Trees, profile.ClassSlug)
	if !ok {
		return nil
	}
	code, report, err := bnetbuild.Encode(bnetbuild.Inputs{
		Build:     s.Tables.Build,
		Profile:   profile,
		Equipment: rawEquipment,
		Talents:   rawSpecializations,
		Talent:    bnetbuild.TalentTable{Build: s.Tables.Trees, ClassID: classID},
		Enchants:  s.Tables.Enchants,
		Suffixes:  s.Tables.Suffixes,
		Races:     s.Tables.Races,
	})
	if err != nil {
		s.logger().Warn("bnetimport", "op", "build_encode", "key", key, "err", err)
		return nil
	}

	capturedAt := time.Now()
	if profile.LastLoginTimestamp > 0 {
		capturedAt = time.UnixMilli(profile.LastLoginTimestamp)
	}
	tag, err := tx.Exec(ctx,
		`insert into addon_exports (character_key, user_id, region, ruleset, name, export, source, captured_at, updated_at)
		 values ($1, $2, $3, $4, $5, $6, 'blizzard', $7, now())
		 on conflict (character_key) do update set
		   user_id = excluded.user_id, region = excluded.region, ruleset = excluded.ruleset, name = excluded.name,
		   export = excluded.export, source = 'blizzard', captured_at = excluded.captured_at, updated_at = now()
		 where addon_exports.user_id = excluded.user_id
		   and (addon_exports.source = 'blizzard' or excluded.captured_at > addon_exports.captured_at)`,
		key, userID, region, ruleset, profile.Name, code, capturedAt)
	if err != nil {
		return fmt.Errorf("bnetimport: write export %s: %w", key, err)
	}
	if tag.RowsAffected() == 0 {
		// The newest-wins guard (spec §2.4) blocked the write — a different account
		// already owns this key's addon_exports row, or an addon export taken since the
		// character's last Blizzard session is still the freshest source. Not an error:
		// this is the guard doing exactly what it's for.
		s.logger().Info("bnetimport", "op", "build_guarded", "key", key)
		return nil
	}
	s.logger().Info("bnetimport", "op", "build_encoded", "key", key,
		"unmatched_talents", len(report.UnmatchedTalents), "clamped", len(report.Clamped),
		"no_suffix", len(report.NoSuffix), "skipped_slots", len(report.SkippedSlots))
	return nil
}

// classIDFor resolves a class slug to its id within the active build's data — the one
// lookup bnetbuild.TalentTable needs that trees.Build indexes by id rather than slug.
func classIDFor(b *trees.Build, slug string) (int, bool) {
	for _, c := range b.Classes() {
		if c.Slug == slug {
			return c.ID, true
		}
	}
	return 0, false
}
