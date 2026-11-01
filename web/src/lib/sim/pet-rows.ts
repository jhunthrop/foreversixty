// web/src/lib/sim/pet-rows.ts
// A warlock registers all four demons as pet units, so a sim result carries a damage row
// for each, summoned or not. sim/adapter gives every pet row its real time summoned in the
// fight (the engine's pet_active_seconds_avg); this drops the rows of pets that were never
// out at all, so no tab lists a demon that took no part in the fight.
import type { Actor, Summary } from '../report/types';

/** A pet row with no damage and no time summoned in the fight. A player row never is. */
function isAbsentPet(actor: Actor, players: ReadonlySet<string>): boolean {
  return !players.has(actor.guid) && actor.total === 0 && actor.effective === 0 && actor.active_ms === 0;
}

/**
 * The summary without the pets that were never out in the fight, from every list that
 * names a unit. A zero-damage pet that was out keeps its row and its own uptime. Returns
 * the summary itself when nothing is dropped.
 */
export function withoutAbsentPets(summary: Summary): Summary {
  const players = new Set(summary.roster.map((row) => row.guid));
  const absent = new Set(
    summary.damage_done.filter((actor) => isAbsentPet(actor, players)).map((actor) => actor.guid),
  );
  if (absent.size === 0) return summary;
  return {
    ...summary,
    damage_done: summary.damage_done.filter((actor) => !absent.has(actor.guid)),
    casts: summary.casts.filter((row) => !absent.has(row.guid)),
    auras: summary.auras.filter((track) => !absent.has(track.target_guid)),
    resources: summary.resources.filter((track) => !absent.has(track.guid)),
  };
}
