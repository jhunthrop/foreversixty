// web/src/lib/sim/rotations.ts
// The curated rotation's player-facing notes, synced by scripts/sync-rotations.mjs from
// data/curated/apl/<spec>.json into src/data/generated/rotations.json -- gitignored, like
// every other file under data/generated/ (sim-ids.json's own header explains why: derived
// data is not committed). Task 6 (newcomer BLOCKER, tank MAJOR): RotationCard's "what it
// does" used to send the player to /sim/specs#<spec>, which explains parse fidelity, not
// the rotation, and threw away the loaded character and the finished run on the way. This
// is what the in-page drawer reads instead.
//
// The notes for every spec weigh about 12 KB gzipped and are read only inside a closed
// drawer, so they are their own chunk, loaded after first paint rather than carried in the
// sim island's boot bundle (scripts/check-island-size.mjs; round 7, owner ruling 1).
import { specDisplayName } from './spec-label';
import { simCopy } from './copy';

/** Every spec's ordered rotation notes, keyed by spec. */
export type RotationNotes = Readonly<Record<string, readonly string[]>>;

let loading: Promise<RotationNotes> | undefined;

/** Loads the notes once; every later call shares the first load. */
export function loadRotationNotes(): Promise<RotationNotes> {
  loading ??= import('../../data/generated/rotations.json').then((module) => module.default as RotationNotes);
  return loading;
}

/**
 * The spec's ordered rotation notes, in the order the curated priority list casts them. A
 * spec sync-rotations.mjs never wrote an entry for (an unknown spec, or one whose curated
 * file has not landed) answers `[]` -- not fatal, the same as a spec whose curated file
 * carries steps with no player-facing note at all.
 */
export function rotationNotesFor(spec: string, notes: RotationNotes): readonly string[] {
  return notes[spec] ?? [];
}

export interface RotationDrawerContent {
  /** One sentence naming the spec, so the drawer reads on its own. */
  intro: string;
  /** The ordered step notes; empty when the curated file has none yet, null while loading. */
  steps: readonly string[] | null;
}

/**
 * RotationCard's drawer content, and /sim/specs' own per-spec card (task-6-brief.md: "the
 * page a visitor lands on from the strip answers 'what does this rotation do' too"). Pure,
 * so it is unit-testable without mounting either component. `notes` is null until
 * loadRotationNotes() resolves.
 */
export function rotationDrawerContent(spec: string, notes: RotationNotes | null): RotationDrawerContent {
  return {
    intro: simCopy.rotationDrawerIntro(specDisplayName(spec)),
    steps: notes === null ? null : rotationNotesFor(spec, notes),
  };
}
