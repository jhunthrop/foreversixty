// web/src/lib/sim/styles.ts
// The fight styles, contract 1.6. A style is a page preset: it expands to encounter fields
// and leaves its own id behind as `encounter.style`, which is a label the engine ignores.
// The table is duplicated as src/fixtures/sim/styles.json because the sim module lane owns
// a Go copy of it; styles.test.ts asserts this file and that one agree, so the fixture --
// not either implementation -- is the artefact the two lanes share.
//
// Nothing here holds a word: labels and notes are simCopy.styleLabel and simCopy.styleNote,
// keyed by these ids, so a copy change is one diff in copy.ts.
import type { EncounterSpec, Movement, TargetCount } from './types';

export type FightStyleId =
  | 'patchwerk'
  | 'execute'
  | 'light-movement'
  | 'heavy-movement'
  | 'cleave-2'
  | 'cleave-3'
  | 'cleave-5'
  | 'dungeon'
  | 'dummy';

export interface FightStyle {
  id: FightStyleId;
  targets: number;
  execute_ratio: number;
  movement: Movement | null;
  targets_over_time: TargetCount[] | null;
  dummy: boolean;
}

/** The style a fresh settings state opens on. */
export const DEFAULT_STYLE_ID: FightStyleId = 'patchwerk';

const PLAIN_EXECUTE = 0.25;

/** A 5 s window out of melee, at the interval the style names. */
function away(intervalSec: number): Movement {
  return { interval_sec: intervalSec, duration_sec: 5, kind: 'away' };
}

function cleave(id: FightStyleId, targets: number): FightStyle {
  return {
    id,
    targets,
    execute_ratio: PLAIN_EXECUTE,
    movement: null,
    targets_over_time: null,
    dummy: false,
  };
}

export const FIGHT_STYLES: readonly FightStyle[] = [
  {
    id: 'patchwerk',
    targets: 1,
    execute_ratio: PLAIN_EXECUTE,
    movement: null,
    targets_over_time: null,
    dummy: false,
  },
  {
    id: 'execute',
    targets: 1,
    execute_ratio: 0.35,
    movement: null,
    targets_over_time: null,
    dummy: false,
  },
  {
    id: 'light-movement',
    targets: 1,
    execute_ratio: PLAIN_EXECUTE,
    movement: away(45),
    targets_over_time: null,
    dummy: false,
  },
  {
    id: 'heavy-movement',
    targets: 1,
    execute_ratio: PLAIN_EXECUTE,
    movement: away(20),
    targets_over_time: null,
    dummy: false,
  },
  cleave('cleave-2', 2),
  cleave('cleave-3', 3),
  cleave('cleave-5', 5),
  {
    id: 'dungeon',
    targets: 1,
    execute_ratio: 0,
    movement: null,
    targets_over_time: [
      { at_sec: 0, count: 1 },
      { at_sec: 40, count: 3 },
      { at_sec: 80, count: 5 },
      { at_sec: 130, count: 3 },
      { at_sec: 160, count: 1 },
    ],
    dummy: false,
  },
  {
    id: 'dummy',
    targets: 1,
    execute_ratio: 0,
    movement: null,
    targets_over_time: null,
    dummy: true,
  },
];

/** The style with this id, or null. An unknown id is never silently treated as Patchwerk. */
export function fightStyle(id: string): FightStyle | null {
  return FIGHT_STYLES.find((style) => style.id === id) ?? null;
}

/**
 * What TARGETS should say about this encounter (tank MAJOR, review.md:325-327): a
 * timeline style's `targets` field is only the ramp's opening count (the dungeon pull
 * settles on 1, not the 5 it actually reaches), so a control reading that field alone
 * understates the run. This file holds no words -- the caller turns these numbers into a
 * sentence via copy.ts.
 *
 * `attached` (fix round 1, reviewer Important): `settings.ts`'s `detached()` never clears
 * `targets_over_time` when a style-owned field like the dummy checkbox is toggled by hand
 * -- clearing it would silently drop the ramp from the run the moment a player ticks a
 * checkbox, which is exactly the kind of dishonesty this whole fix round exists to remove.
 * So a ramp can outlive its style: `timeline` stays keyed on the ramp alone (the control
 * must never show a number the run will not use), and `attached` is a second, independent
 * flag the caller uses only to pick which sentence is still true, not to change whether
 * the control is read-only.
 */
export function targetsSummary(
  encounter: EncounterSpec,
): { timeline: true; first: number; max: number; attached: boolean } | { timeline: false; count: number } {
  const timeline = encounter.targets_over_time;
  if (timeline !== undefined && timeline.length > 0) {
    return {
      timeline: true,
      first: timeline[0].count,
      max: Math.max(...timeline.map((step) => step.count)),
      attached: (encounter.style ?? '') !== '',
    };
  }
  return { timeline: false, count: encounter.targets };
}

/**
 * The style's fields written over an encounter. Every field a style owns is written on
 * every call, including the nulls and the false: switching from Heavy movement to
 * Patchwerk has to clear the movement block, and a partial write would leave the previous
 * style's fight running under the new style's name.
 *
 * Fight length and duration variation are the player's, not the style's, so they survive.
 */
export function applyFightStyle(encounter: EncounterSpec, id: FightStyleId): EncounterSpec {
  const style = fightStyle(id);
  if (style === null) return encounter;
  return {
    ...encounter,
    style: style.id,
    targets: style.targets,
    execute_ratio: style.execute_ratio,
    // `EncounterSpec.movement`/`targets_over_time` are `omitempty` on the Go side: the
    // wire shape the contract specifies is the key absent, not `null`. `undefined` is the
    // only value that reproduces that -- `JSON.stringify` drops an `undefined` field but
    // still emits `"movement":null` for an explicit null.
    movement: style.movement === null ? undefined : { ...style.movement },
    targets_over_time:
      style.targets_over_time === null ? undefined : style.targets_over_time.map((step) => ({ ...step })),
    dummy: style.dummy,
  };
}
