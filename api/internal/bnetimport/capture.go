// api/internal/bnetimport/capture.go

// Capture support: writing the raw Blizzard bodies bnetapi reads into
// characters' bnet_account/bnet_profile/bnet_equipment columns, and the
// typed fields (race, gender, item level, last login) those bodies carry
// (spec §B, amended by .superpowers/bnet-capture-brief.md).
package bnetimport

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/bnetapi"
)

// maxCaptureBytes is spec §B's cap: "refuse to store a body over 256
// KiB". Blizzard's own responses are ordinarily a few KiB; a body this
// large is either an API change worth investigating or a hostile
// response, not something worth a jsonb column growing unbounded for.
const maxCaptureBytes = 256 * 1024

// capCapture returns raw when it is small enough to store verbatim, or
// nil (after logging op=capture_too_large) when it exceeds
// maxCaptureBytes. A caller writes the result through a
// coalesce($n, column) UPDATE, so a too-large body leaves whatever was
// captured last time in place rather than blanking the column.
func (s *Service) capCapture(field, key string, raw json.RawMessage) json.RawMessage {
	if len(raw) <= maxCaptureBytes {
		return raw
	}
	s.logger().Warn("bnetimport", "op", "capture_too_large", "field", field, "key", key, "bytes", len(raw))
	return nil
}

// rawOrNil turns an empty json.RawMessage into a genuine nil interface
// value, so pgx sends SQL NULL — and the caller's coalesce(...) keeps
// whatever was already stored — instead of an empty-but-non-nil []byte.
func rawOrNil(raw json.RawMessage) any {
	if raw == nil {
		return nil
	}
	return raw
}

// captureProfile stores a character profile fetch's raw body and the
// typed fields it carries — average/equipped item level and last login
// (converted from Blizzard's epoch-millisecond timestamp) — preserving
// whatever was captured before when raw was capped or a field is absent
// from this response (spec §B).
func (s *Service) captureProfile(ctx context.Context, tx pgx.Tx, key string, profile bnetapi.CharacterProfile, raw json.RawMessage) error {
	raw = s.capCapture("bnet_profile", key, raw)
	var lastLogin *time.Time
	if profile.LastLoginTimestamp > 0 {
		t := time.UnixMilli(profile.LastLoginTimestamp)
		lastLogin = &t
	}
	if _, err := tx.Exec(ctx,
		`update characters set
		   bnet_profile = coalesce($2, bnet_profile),
		   average_item_level = coalesce($3, average_item_level),
		   equipped_item_level = coalesce($4, equipped_item_level),
		   last_login_at = coalesce($5, last_login_at),
		   bnet_captured_at = now()
		 where key = $1`,
		key, rawOrNil(raw), profile.AverageItemLevel, profile.EquippedItemLevel, lastLogin); err != nil {
		return fmt.Errorf("bnetimport: capture profile %s: %w", key, err)
	}
	return nil
}

// captureEquipment fetches and stores a character's equipment snapshot
// verbatim (spec §B — not parsed here). A private or missing equipment
// page (403/404) is left exactly as it was, the same rule
// syncCharacterGuild already applies to a private character profile.
func (s *Service) captureEquipment(ctx context.Context, tx pgx.Tx, key, region, realmSlug, name string) error {
	raw, err := s.Client.Equipment(ctx, region, realmSlug, name)
	if err != nil {
		if errors.Is(err, bnetapi.ErrNotFound) || errors.Is(err, bnetapi.ErrForbidden) {
			return nil
		}
		return fmt.Errorf("bnetimport: equipment %s: %w", key, err)
	}
	raw = s.capCapture("bnet_equipment", key, raw)
	if _, err := tx.Exec(ctx,
		`update characters set bnet_equipment = coalesce($2, bnet_equipment), bnet_captured_at = now() where key = $1`,
		key, rawOrNil(raw)); err != nil {
		return fmt.Errorf("bnetimport: capture equipment %s: %w", key, err)
	}
	return nil
}
