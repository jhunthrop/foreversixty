// web/src/lib/character-selector/stale.ts
// Decision 11 of the nav selector spec: a character is stale when its newest build is older
// than clamp(3 x the player's own median gap between syncs, 3 days, 14 days); with no median
// on record (fewer than three syncs, or an API that does not send it yet) 7 days.
import type { MeCharacter } from '../account/api';
import { STALE_FALLBACK_SEC, STALE_GAP_MULTIPLIER, STALE_MAX_SEC, STALE_MIN_SEC } from './config';

type Build = NonNullable<MeCharacter['build']>;

/** The age, in seconds, past which a build counts as stale. */
export function staleThresholdSec(medianSyncGapSec: number | null | undefined): number {
  if (medianSyncGapSec === null || medianSyncGapSec === undefined || !(medianSyncGapSec > 0)) {
    return STALE_FALLBACK_SEC;
  }
  return Math.min(STALE_MAX_SEC, Math.max(STALE_MIN_SEC, STALE_GAP_MULTIPLIER * medianSyncGapSec));
}

export function isStale(build: Build | undefined, now: Date): boolean {
  if (build === undefined) return false;
  const capturedMs = Date.parse(build.captured_at);
  if (Number.isNaN(capturedMs)) return false;
  const ageSec = (now.getTime() - capturedMs) / 1000;
  return ageSec > staleThresholdSec(build.median_sync_gap_sec);
}

/** A failed Battle.net refresh is a non-null `sync_error`; absent or null is not failed. */
export function isFailed(build: Build | undefined): boolean {
  return build?.sync_error !== undefined && build.sync_error !== null;
}
