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
/**
 * `players` is the players; `friendly` is the players and what they own (pets, totems),
 * which the enemies scope leaves out without the friendlies scope taking them in.
 */
export function inSource(
  guid: string,
  source: string,
  players: ReadonlySet<string>,
  friendly: ReadonlySet<string> = players,
): boolean {
  if (source === SOURCE_FRIENDLIES) return players.size === 0 || players.has(guid);
  if (source === SOURCE_ENEMIES) return !friendly.has(guid);
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
export function scopeSource(
  summary: Summary,
  source: string,
  players: ReadonlySet<string>,
  friendly: ReadonlySet<string> = players,
): Summary {
  if (source === SOURCE_FRIENDLIES && players.size === 0) return summary;
  const keep = (guid: string): boolean => inSource(guid, source, players, friendly);
  const actors = (table: Summary['damage_done']): Summary['damage_done'] =>
    table.filter((actor) => keep(actor.guid));
  return {
    ...summary,
    damage_done: actors(summary.damage_done),
    damage_taken: actors(summary.damage_taken),
    healing: actors(summary.healing),
    healing_taken: actors(summary.healing_taken),
    roster: summary.roster.filter((row) => keep(row.guid)),
    combatants: summary.combatants.filter((row) => keep(row.guid)),
    deaths: summary.deaths.filter((death) => keep(death.guid)),
    // An aura belongs to its target, and to the one player who applied it when the
    // scope is a player: Flame Shock on the boss is the shaman's own uptime.
    auras: summary.auras.filter(
      (track) =>
        keep(track.target_guid) ||
        (source !== SOURCE_ENEMIES && players.has(source) && track.appliers.includes(source)),
    ),
    // A pet's casts are its owner's work, the way a pet's damage is its owner's damage:
    // picking the healer shows the statue's casts too. A row written before the engine
    // kept the owner falls back to the caster's own guid, which is what the scope did
    // before, so an older report narrows exactly as it used to.
    casts: summary.casts.filter((row) => keep(row.owner_guid ?? row.guid)),
    resources: summary.resources.filter((track) => keep(track.guid)),
    threat: summary.threat.filter((row) => keep(row.guid)),
    threat_by_target: summary.threat_by_target?.filter((pair) => keep(pair.guid)),
    taunts: summary.taunts?.filter((taunt) => keep(taunt.source_guid)),
    interrupts: summary.interrupts.filter((row) => keep(row.source_guid)),
    dispels: summary.dispels.filter((row) => keep(row.source_guid)),
  };
}
