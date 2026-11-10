// web/src/lib/report/exchange-rails.ts
// The Interrupts and Dispels tabs' rails (round 5, b1-report-interrupts and b1-report-dispels):
// what the enemies cast, beside the kicks; what landed on the group, beside the dispels.
import type { AuraTrack, CastRow } from './types';

/** How many rows a rail lists: a rail is a glance, the tabs hold the rest. */
export const RAIL_ROWS = 8;

export interface EnemyCast {
  spellId: number;
  name: string;
  casters: string[];
  /** Casts that went off in the window. */
  count: number;
  /** "2.0 s cast" or "instant". */
  castLabel: string;
  castMs: number;
}

/** Enemy spells cast in the window, the ones with a cast bar first, then by how often. */
export function enemyCasts(casts: readonly CastRow[], players: ReadonlySet<string>): EnemyCast[] {
  const bySpell = new Map<number, EnemyCast>();
  for (const cast of casts) {
    const friendly = players.has(cast.guid) || (cast.owner_guid !== undefined && players.has(cast.owner_guid));
    if (friendly || cast.succeeded === 0) continue;
    const found = bySpell.get(cast.spell_id);
    const castMs = Math.max(found?.castMs ?? 0, cast.cast_time_ms);
    bySpell.set(cast.spell_id, {
      spellId: cast.spell_id,
      name: cast.spell_name,
      casters: [...new Set([...(found?.casters ?? []), cast.name])],
      count: (found?.count ?? 0) + cast.succeeded,
      castMs,
      castLabel: castMs > 0 ? `${(castMs / 1000).toFixed(1)} s cast` : 'instant',
    });
  }
  return [...bySpell.values()]
    .sort((a, b) => Number(b.castMs > 0) - Number(a.castMs > 0) || b.count - a.count || a.name.localeCompare(b.name))
    .slice(0, RAIL_ROWS);
}

export interface GroupDebuff {
  spellId: number;
  name: string;
  /** How many of the scope's players it landed on. */
  players: number;
  /** How many times it landed (its segments). */
  landed: number;
}

/** The debuffs that landed on the scope's players, the most often first. */
export function groupDebuffs(auras: readonly AuraTrack[], players: ReadonlySet<string>): GroupDebuff[] {
  const byName = new Map<string, { spellId: number; targets: Set<string>; landed: number }>();
  for (const track of auras) {
    if (track.type !== 'DEBUFF' || !players.has(track.target_guid) || track.segments.length === 0) continue;
    const found = byName.get(track.name) ?? { spellId: track.spell_id, targets: new Set<string>(), landed: 0 };
    byName.set(track.name, {
      spellId: found.spellId,
      targets: new Set([...found.targets, track.target_guid]),
      landed: found.landed + track.segments.length,
    });
  }
  return [...byName.entries()]
    .map(([name, entry]) => ({ spellId: entry.spellId, name, players: entry.targets.size, landed: entry.landed }))
    .sort((a, b) => b.landed - a.landed || a.name.localeCompare(b.name))
    .slice(0, RAIL_ROWS);
}
