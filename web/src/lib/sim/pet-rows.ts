// web/src/lib/sim/pet-rows.ts
// A warlock registers all four demons as pet units, so a sim result carries a damage row
// for each, summoned or not. sim/adapter gives every pet row its real time summoned in the
// fight (the engine's pet_active_seconds_avg) and marks it `pet_uptime`; this drops the rows
// of pets that were never out at all, so no tab lists a demon that took no part in the fight.
//
// A result saved before that measurement has no mark, and its adapter stamped every pet
// active for the whole fight. There a pet with no damage is dropped, and a pet that did
// damage keeps its row with Active unknown (`active_unknown`), never a made-up 100%.
import type { Actor, Summary } from '../report/types';

/** A pet row whose Active is the fight's length, not a measurement: an older result's. */
function isUnmeasuredPet(actor: Actor, players: ReadonlySet<string>): boolean {
  return !players.has(actor.guid) && actor.pet_uptime !== true;
}

/** A pet row that took no part in the fight. A player row never is. */
function isAbsentPet(actor: Actor, players: ReadonlySet<string>): boolean {
  if (players.has(actor.guid) || actor.total !== 0 || actor.effective !== 0) return false;
  return actor.active_ms === 0 || isUnmeasuredPet(actor, players);
}

/**
 * The summary without the pets that were never out in the fight, from every list that
 * names a unit, and with an older result's surviving pet rows marked Active unknown. A
 * zero-damage pet measured as out keeps its row and its own uptime. Returns the summary
 * itself when nothing changes.
 */
export function withoutAbsentPets(summary: Summary): Summary {
  const players = new Set(summary.roster.map((row) => row.guid));
  const absent = new Set(
    summary.damage_done.filter((actor) => isAbsentPet(actor, players)).map((actor) => actor.guid),
  );
  const unmeasured = summary.damage_done.some(
    (actor) => !absent.has(actor.guid) && isUnmeasuredPet(actor, players),
  );
  if (absent.size === 0 && !unmeasured) return summary;
  return {
    ...summary,
    damage_done: summary.damage_done
      .filter((actor) => !absent.has(actor.guid))
      .map((actor) => (isUnmeasuredPet(actor, players) ? { ...actor, active_unknown: true } : actor)),
    casts: summary.casts.filter((row) => !absent.has(row.guid)),
    auras: summary.auras.filter((track) => !absent.has(track.target_guid)),
    resources: summary.resources.filter((track) => !absent.has(track.guid)),
  };
}
