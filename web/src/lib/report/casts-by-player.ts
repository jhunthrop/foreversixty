// web/src/lib/report/casts-by-player.ts
// The Casts tab's By player rail (density pass 2026-11-09): each player's own casts, failed
// casts, spells used and casts per minute over the window. A pet's or totem's casts are left
// out: a Searing Totem's swings are not the shaman pressing a button. Pure, pinned in
// casts-by-player.test.ts.
import type { CastRow } from './types';

export interface PlayerCasts {
  guid: string;
  casts: number;
  failed: number;
  /** Distinct spells cast. */
  spells: number;
  perMinute: number;
}

/** Every player in `players` who cast anything themselves, busiest first. */
export function castsByPlayer(
  rows: readonly CastRow[],
  players: ReadonlySet<string>,
  durationMs: number,
): PlayerCasts[] {
  const minutes = Math.max(durationMs, 1) / 60_000;
  const byPlayer = new Map<string, { casts: number; failed: number; spells: Set<number> }>();
  for (const row of rows) {
    const own = row.owner_guid === undefined || row.owner_guid === row.guid;
    if (!own || !players.has(row.guid) || row.succeeded + row.failed === 0) continue;
    const entry = byPlayer.get(row.guid) ?? { casts: 0, failed: 0, spells: new Set<number>() };
    byPlayer.set(row.guid, {
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
