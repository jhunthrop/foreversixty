// web/src/lib/report/exchange-rails.ts
// The Interrupts and Dispels tabs' rails (round 5, b1-report-interrupts and b1-report-dispels):
// what the enemies cast, beside the kicks; what landed on the group, beside the dispels.
import type { AuraTrack, CastRow, Summary } from './types';

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
    const friendly =
      players.has(cast.guid) || (cast.owner_guid !== undefined && players.has(cast.owner_guid));
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
    .sort(
      (a, b) =>
        Number(b.castMs > 0) - Number(a.castMs > 0) || b.count - a.count || a.name.localeCompare(b.name),
    )
    .slice(0, RAIL_ROWS);
}

/**
 * Each spell's school as the summary knows it (the damage taken and the auras carry one, a cast
 * row does not), so an enemy spell with no icon of its own draws its school's base icon.
 */
export function spellSchools(summary: Summary | null): ReadonlyMap<number, number> {
  const schools = new Map<number, number>();
  if (summary === null) return schools;
  for (const actor of [...summary.damage_taken, ...summary.damage_done]) {
    for (const ability of actor.abilities) {
      if (ability.school !== undefined && ability.school > 0) schools.set(ability.spell_id, ability.school);
    }
  }
  for (const track of summary.auras) {
    if (track.school !== undefined && track.school > 0) schools.set(track.spell_id, track.school);
  }
  return schools;
}

export interface GroupDebuff {
  spellId: number;
  name: string;
  /** How many of the scope's players it landed on. */
  players: number;
  /** How many times it landed (its segments). */
  landed: number;
  /** Its time up, summed over every player it was on. */
  uptimeMs: number;
  school?: number;
}

/** The debuffs that landed on the scope's players, the most often first. */
export function groupDebuffs(auras: readonly AuraTrack[], players: ReadonlySet<string>): GroupDebuff[] {
  const byName = new Map<
    string,
    { spellId: number; targets: Set<string>; landed: number; uptimeMs: number; school?: number }
  >();
  for (const track of auras) {
    if (track.type !== 'DEBUFF' || !players.has(track.target_guid) || track.segments.length === 0) continue;
    const found = byName.get(track.name) ?? {
      spellId: track.spell_id,
      targets: new Set<string>(),
      landed: 0,
      uptimeMs: 0,
      school: track.school,
    };
    byName.set(track.name, {
      spellId: found.spellId,
      targets: new Set([...found.targets, track.target_guid]),
      landed: found.landed + track.segments.length,
      uptimeMs: found.uptimeMs + track.uptime_ms,
      school: found.school ?? track.school,
    });
  }
  return [...byName.entries()]
    .map(([name, entry]) => ({
      spellId: entry.spellId,
      name,
      players: entry.targets.size,
      landed: entry.landed,
      uptimeMs: entry.uptimeMs,
      school: entry.school,
    }))
    .sort((a, b) => b.landed - a.landed || a.name.localeCompare(b.name))
    .slice(0, RAIL_ROWS);
}
