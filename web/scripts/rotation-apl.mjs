// web/scripts/rotation-apl.mjs
// data/curated/apl/<spec>.json's rotation.priorityList, reduced to the ordered
// player-facing notes -- what "what it does" (Task 6, newcomer BLOCKER, tank MAJOR) shows
// instead of sending the player to /sim/specs, which explains parse fidelity, not the
// rotation.
//
// The file's own top-level `notes` is developer prose about measurement (e.g. "the engine
// lane's spec tasks measured a 2.0 percent ... DPS loss") and is never read here: only a
// step's own `notes` is player-facing. A step with no `notes` is skipped -- most
// cooldown-only entries (autocastOtherCooldowns and the like) carry none.
//
// A plain .mjs so scripts/sync-rotations.mjs can import the same parser this runs under
// node before vitest starts; src/lib/sim/rotation-apl.test.ts drives it through vitest all
// the same, the way ids-md.mjs/ids-md.test.ts already do for sync-sim-ids.mjs.

/** @typedef {{ notes?: unknown }} AplStep */
/** @typedef {{ rotation?: { priorityList?: AplStep[] } }} CuratedApl */

/**
 * The ordered, non-empty step notes from one curated APL object. Never the file's own
 * top-level `notes` -- that is measurement commentary, not a rotation description.
 * @param {CuratedApl} apl
 * @returns {string[]}
 */
export function rotationNotesOf(apl) {
  const steps = apl?.rotation?.priorityList ?? [];
  return steps
    .map((step) => (typeof step?.notes === 'string' ? step.notes.trim() : ''))
    .filter((note) => note !== '');
}
