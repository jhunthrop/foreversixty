// web/src/lib/sim/encounter.ts
// How an EncounterSpec reads. Two functions, in the base task rather than beside the
// settings bar, because two groups of this plan need them at the same time: the Worker's
// unfurl copy (the routes group) and the settings bar (the character-model group). Pure
// formatting over the shape types.ts defines, so this is where they belong regardless.
import type { EncounterSpec } from './types';

/** Seconds as a clock: 180 is "3:00", not "180s" and not "3 minutes". */
export function durationLabel(seconds: number): string {
  const minutes = Math.floor(seconds / 60);
  return `${minutes}:${String(seconds % 60).padStart(2, '0')}`;
}

/**
 * The settings clause a saved sim's title and a shared sim's unfurl both carry:
 * "raid-buffed, 3:00, single target".
 */
export function encounterLabel(encounter: EncounterSpec, buffed: boolean): string {
  const targets = encounter.targets === 1 ? 'single target' : `${encounter.targets} targets`;
  return `${buffed ? 'raid-buffed' : 'solo'}, ${durationLabel(encounter.duration_sec)}, ${targets}`;
}
