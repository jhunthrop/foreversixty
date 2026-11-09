// web/src/lib/report/casts-by-player.ts
// The Casts tab's By player rail (density pass 2026-11-09): each player's casts, failed casts,
// spells used and casts per minute over the window, a pet's casts counted to its owner the way
// the Casts table and the damage tables count them. Pure, pinned in casts-by-player.test.ts.
import type { CastRow } from './types';

export interface PlayerCasts {
  guid: string;
  casts: number;
  failed: number;
  /** Distinct spells cast. */
  spells: number;
  perMinute: number;
}

/** Every player in `players` who cast anything, busiest first. */
export function castsByPlayer(
  rows: readonly CastRow[],
  players: ReadonlySet<string>,
  durationMs: number,
): PlayerCasts[] {
  const minutes = Math.max(durationMs, 1) / 60_000;
  const byPlayer = new Map<string, { casts: number; failed: number; spells: Set<number> }>();
  for (const row of rows) {
    const guid = row.owner_guid ?? row.guid;
    if (!players.has(guid) || row.succeeded + row.failed === 0) continue;
    const entry = byPlayer.get(guid) ?? { casts: 0, failed: 0, spells: new Set<number>() };
    byPlayer.set(guid, {
      casts: entry.casts + row.succeeded,
      failed: entry.failed + row.failed,
      spells: row.succeeded > 0 ? new Set([...entry.spells, row.spell_id]) : entry.spells,
    });
  }
  return [...byPlayer.entries()]
    .map(([guid, entry]) => ({
      guid,
      casts: entry.casts,
      failed: entry.failed,
      spells: entry.spells.size,
      perMinute: entry.casts / minutes,
    }))
    .sort((a, b) => b.casts - a.casts || a.guid.localeCompare(b.guid));
}
