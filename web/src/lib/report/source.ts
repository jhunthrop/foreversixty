// web/src/lib/report/source.ts
// The Source control (url.ts's `state.source`) narrows every table to one unit, to the
// friendlies or to the enemies. The three Actor tables already scope themselves through
// ReportView's tabActors; this is the same rule for everything else the page shows -- the
// roster, auras, casts, deaths, resources, threat, interrupts and dispels -- so picking a
// player from the Source list narrows the whole page to them rather than only the damage
// tabs, which is what a reader who picked their own name expects.
import type { Summary } from './types';
import { SOURCE_ENEMIES, SOURCE_FRIENDLIES } from './url';

/** Whether a unit is inside the chosen source scope. */
export function inSource(guid: string, source: string, players: ReadonlySet<string>): boolean {
  if (source === SOURCE_FRIENDLIES) return players.size === 0 || players.has(guid);
  if (source === SOURCE_ENEMIES) return !players.has(guid);
  return guid === source;
}

/**
 * The summary with everything outside the source scope removed. `players` is the set of
 * player GUIDs from report.json; when it is empty (a report with no unit list) nothing
 * is cut. The roster and combatants are players already, and under `friendlies` they are
 * left alone; the cast, aura, resource and threat tables carry enemies too, and the
 * default scope cuts those down to the players so a Casts tab does not open on sixty
 * rows of trash spells.
 */
export function scopeSource(summary: Summary, source: string, players: ReadonlySet<string>): Summary {
  if (source === SOURCE_FRIENDLIES && players.size === 0) return summary;
  const keep = (guid: string): boolean => inSource(guid, source, players);
  return {
    ...summary,
    roster: summary.roster.filter((row) => keep(row.guid)),
    combatants: summary.combatants.filter((row) => keep(row.guid)),
    deaths: summary.deaths.filter((death) => keep(death.guid)),
    auras: summary.auras.filter((track) => keep(track.target_guid)),
    casts: summary.casts.filter((row) => keep(row.guid)),
    resources: summary.resources.filter((track) => keep(track.guid)),
    threat: summary.threat.filter((row) => keep(row.guid)),
    interrupts: summary.interrupts.filter((row) => keep(row.source_guid)),
    dispels: summary.dispels.filter((row) => keep(row.source_guid)),
  };
}
