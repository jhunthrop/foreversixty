// web/src/lib/sim/cooldowns.ts
// Contract 1.7's CooldownSpec, as four choices a player recognises.
//
// The engine takes a list of seconds and an empty list means "on cooldown". "At execute"
// is not a fifth thing the engine knows -- it is the second the execute window opens,
// computed here from the encounter, so changing the fight length moves it with the fight
// rather than leaving a potion at a timestamp that is now past the pull.
import type { CooldownSpec, EncounterSpec } from './types';

export type CooldownMode = 'on-cooldown' | 'on-pull' | 'at-time' | 'at-execute';

export const COOLDOWN_MODES: readonly CooldownMode[] = ['on-cooldown', 'on-pull', 'at-time', 'at-execute'];

export interface CooldownRow {
  id: string;
  mode: CooldownMode;
  /** Meaningful for `at-time`; 0 otherwise. */
  atSec: number;
}

/** Where the execute window opens. With no execute window, the end of the fight. */
export function executeStartSec(encounter: EncounterSpec): number {
  return Math.round(encounter.duration_sec * (1 - encounter.execute_ratio));
}

export function specFor(
  id: string,
  mode: CooldownMode,
  atSec: number,
  encounter: EncounterSpec,
): CooldownSpec {
  switch (mode) {
    case 'on-cooldown':
      return { id, at_sec: [] };
    case 'on-pull':
      return { id, at_sec: [0] };
    case 'at-execute':
      return { id, at_sec: [executeStartSec(encounter)] };
    default:
      return { id, at_sec: [Math.min(encounter.duration_sec, Math.max(0, Math.round(atSec)))] };
  }
}

/**
 * The mode a stored spec was written by. A list of more than one time is a request the
 * panel cannot express, so it reads as `at-time` on its first entry rather than being
 * silently rewritten -- the drawer (Task 15) is where a multi-use schedule is edited.
 */
export function modeOf(spec: CooldownSpec, encounter: EncounterSpec): CooldownMode {
  if (spec.at_sec.length === 0) return 'on-cooldown';
  if (spec.at_sec.length === 1) {
    if (spec.at_sec[0] === 0) return 'on-pull';
    if (spec.at_sec[0] === executeStartSec(encounter)) return 'at-execute';
  }
  return 'at-time';
}

/**
 * A row for every id offered, plus a row for every stored spec whose id is not among them
 * -- so a request pasted into the drawer with a class cooldown in it keeps that cooldown
 * when the panel re-renders, instead of the panel quietly dropping what it cannot offer.
 */
export function rowsFor(
  ids: readonly string[],
  specs: readonly CooldownSpec[],
  encounter: EncounterSpec,
): CooldownRow[] {
  const byId = new Map(specs.map((spec) => [spec.id, spec]));
  const offered: CooldownRow[] = ids.map((id) => {
    const spec = byId.get(id);
    if (spec === undefined) return { id, mode: 'on-cooldown', atSec: 0 };
    return { id, mode: modeOf(spec, encounter), atSec: spec.at_sec[0] ?? 0 };
  });
  const extra: CooldownRow[] = specs
    .filter((spec) => !ids.includes(spec.id))
    .map((spec) => ({ id: spec.id, mode: modeOf(spec, encounter), atSec: spec.at_sec[0] ?? 0 }));
  return [...offered, ...extra];
}

/**
 * One row's answer, written into the spec list. "On cooldown" removes the spec rather than
 * storing an empty one: it is the engine's own default, and a request that spells out
 * every default is a request nobody can read in the drawer.
 */
export function withCooldown(
  specs: readonly CooldownSpec[],
  id: string,
  mode: CooldownMode,
  atSec: number,
  encounter: EncounterSpec,
): CooldownSpec[] {
  const without = specs.filter((spec) => spec.id !== id);
  if (mode === 'on-cooldown') return without;
  return [...without, specFor(id, mode, atSec, encounter)];
}
