// web/src/lib/sim/rotations.ts
// The curated rotation's player-facing notes, synced by scripts/sync-rotations.mjs from
// data/curated/apl/<spec>.json into src/data/generated/rotations.json -- gitignored, like
// every other file under data/generated/ (sim-ids.json's own header explains why: derived
// data is not committed). Task 6 (newcomer BLOCKER, tank MAJOR): RotationCard's "what it
// does" used to send the player to /sim/specs#<spec>, which explains parse fidelity, not
// the rotation, and threw away the loaded character and the finished run on the way. This
// is what the in-page drawer reads instead.
import generated from '../../data/generated/rotations.json';
import { specDisplayName } from './spec-label';
import { simCopy } from './copy';

const rotations = generated as Record<string, string[]>;

/**
 * The spec's ordered rotation notes, in the order the curated priority list casts them. A
 * spec sync-rotations.mjs never wrote an entry for (an unknown spec, or one whose curated
 * file has not landed) answers `[]` -- not fatal, the same as a spec whose curated file
 * carries steps with no player-facing note at all.
 */
export function rotationNotesFor(spec: string): readonly string[] {
  return rotations[spec] ?? [];
}

export interface RotationDrawerContent {
  /** One sentence naming the spec, so the drawer reads on its own. */
  intro: string;
  /** The ordered step notes; empty when the curated file has none yet. */
  steps: readonly string[];
}

/**
 * RotationCard's drawer content, and /sim/specs' own per-spec card (task-6-brief.md: "the
 * page a visitor lands on from the strip answers 'what does this rotation do' too"). Pure,
 * so it is unit-testable without mounting either component.
 */
export function rotationDrawerContent(spec: string): RotationDrawerContent {
  return {
    intro: simCopy.rotationDrawerIntro(specDisplayName(spec)),
    steps: rotationNotesFor(spec),
  };
}
