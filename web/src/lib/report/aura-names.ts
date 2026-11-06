// web/src/lib/report/aura-names.ts
// Bug B2 (logs spec 2026-11-07 section 11): an aura already up at the pull reaches the summary as
// "Spell #17" when no line of the log named it, though the build's spell table does (Power Word:
// Shield). The Buffs and Debuffs tabs resolve such names through the build's shared spell names
// (`simnames/_shared.json`, the table the simulator already reads), as the cast and damage rows
// read named spells.
import { dataUrl, loadOptional } from '../planner/load';
import type { AuraTrack } from './types';

const UNNAMED = /^Spell #\d+$/;

export type SpellNames = Readonly<Record<string, string>>;

/** Whether a track's name is the engine's placeholder for a spell the log never named. */
export function isUnnamed(name: string): boolean {
  return name === '' || UNNAMED.test(name);
}

/** The tracks with every placeholder name the table knows replaced; the rest unchanged. */
export function nameAuras(tracks: readonly AuraTrack[], names: SpellNames): AuraTrack[] {
  return tracks.map((track) => {
    if (!isUnnamed(track.name)) return track;
    const known = names[String(track.spell_id)];
    return known === undefined || known === '' ? track : { ...track, name: known };
  });
}

let shared: Promise<SpellNames> | null = null;

/** The build's shared spell names, read once; an empty table when the build has none. */
export function loadSharedSpellNames(build: string): Promise<SpellNames> {
  shared ??= loadOptional<{ spell?: Record<string, string> }>(
    dataUrl(build, 'simnames/_shared.json'),
    {},
  ).then((table) => table.spell ?? {});
  return shared;
}
