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
//
// A step's own engineering commentary (Go/config file paths, commit-round narrative,
// internal spell/aura ids) belongs in that step's `engineeringNotes`, never in `notes` --
// `notes` is rendered verbatim in the live rotation drawer (rotationNotesFor -> steps,
// rotations.ts), so anything written there is player copy. `assertNoEngineeringLeak` is
// the fail-fast guard sync-rotations.mjs runs against every note it publishes, so a leak
// breaks `npm run sync:rotations` (chained into pretest/precheck/predev/prebuild) instead
// of reaching the live page; the same check is exercised directly against every curated
// file in rotation-apl.test.ts.

/** @typedef {{ notes?: unknown }} AplStep */
/** @typedef {{ rotation?: { priorityList?: AplStep[] } }} CuratedApl */

/**
 * A file path this codebase would write (a slash-separated path ending in a source or
 * data extension), e.g. "sim/core/debuffs.go" or "ui/warlock/apls/rotation.apl.json".
 */
const FILE_PATH_RE = /\b(?:[\w-]+\/)+[\w-]+(?:\.[\w-]+)*\.(?:go|mjs|ts|tsx|json)\b/;

/** Phrases that only ever show up in engineering commit narrative, never player copy. */
const LEAK_PHRASES = [/fix round/i, /smoke run/i];

/**
 * True if `text` contains a 7-40 char token that reads as a hex/commit hash -- i.e. one
 * that mixes decimal digits with an a-f letter. A bare decimal id (a spell or item id) is
 * never flagged: it has no letters, so it can't be mistaken for a hash.
 * @param {string} text
 */
function hasCommitLikeToken(text) {
  const tokens = text.match(/\b[0-9a-fA-F]{7,40}\b/g) ?? [];
  return tokens.some((token) => /[a-fA-F]/.test(token));
}

/**
 * Why `note` is not safe to publish as player-facing rotation copy, or null if it is.
 * @param {string} note
 * @returns {string | null}
 */
export function engineeringLeakIn(note) {
  if (FILE_PATH_RE.test(note)) return `contains a file path: ${FILE_PATH_RE.exec(note)?.[0]}`;
  const phraseHit = LEAK_PHRASES.find((re) => re.test(note));
  if (phraseHit) return `contains the engineering phrase ${phraseHit}`;
  if (hasCommitLikeToken(note)) return 'contains a hex/commit-looking token';
  return null;
}

/**
 * Throws with the offending spec/step/reason if `note` leaks engineering narrative.
 * Called by sync-rotations.mjs on every note it is about to publish, so a leak fails the
 * sync step (and therefore pretest/precheck/predev/prebuild) instead of reaching the live
 * rotation drawer.
 * @param {string} note
 * @param {{ spec: string, stepIndex: number }} where
 */
export function assertNoEngineeringLeak(note, { spec, stepIndex }) {
  const reason = engineeringLeakIn(note);
  if (reason) {
    throw new Error(
      `rotation-apl: ${spec} step ${stepIndex}'s player-facing notes ${reason} -- move the ` +
        `engineering narrative to that step's engineeringNotes field instead.`,
    );
  }
}

/**
 * The ordered, non-empty step notes from one curated APL object. Never the file's own
 * top-level `notes` -- that is measurement commentary, not a rotation description. Never a
 * step's own `engineeringNotes` either, by the same rule. Throws via
 * `assertNoEngineeringLeak` if a note leaks engineering narrative; pass `spec` so the error
 * names the offending file.
 * @param {CuratedApl} apl
 * @param {string} [spec]
 * @returns {string[]}
 */
export function rotationNotesOf(apl, spec = '<unknown spec>') {
  const steps = apl?.rotation?.priorityList ?? [];
  const notes = [];
  steps.forEach((step, stepIndex) => {
    const note = typeof step?.notes === 'string' ? step.notes.trim() : '';
    if (note === '') return;
    assertNoEngineeringLeak(note, { spec, stepIndex });
    notes.push(note);
  });
  return notes;
}
