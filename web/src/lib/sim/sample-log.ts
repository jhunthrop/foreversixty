// web/src/lib/sim/sample-log.ts
// Design 5.1: "one iteration's casts in order with the pre-pull section separated,
// resources at each cast, and the note that it is one iteration and not a guide".
//
// The order is the engine's and is never sorted here: a cast log read out of order is not
// a cast log.
//
// Contract A12: a row carries an action key (`spell:23881`, `item:13503`, `other:melee`)
// and nothing else. Names go through resolveActionName like every other row on this lane,
// so a key never reaches the table and the sample and cast tables can never name one
// action two different ways.
import { resolveActionName, type ActionNames } from './action-names';
import type { SampleCast } from './types';

/**
 * The engine's resources, in the order a player reads them: the caster's pool first, then
 * the melee ones, then combo points. A resource this list does not anticipate is appended
 * rather than dropped -- the engine may well gain one, and a missing column would silently
 * hide it.
 */
export const RESOURCE_ORDER: readonly string[] = ['mana', 'energy', 'rage', 'focus', 'combo_points'];

export interface SampleRow {
  /** A key of its own: two casts can share an instant and a spell id. */
  key: string;
  atMs: number;
  /** True while the fight has not started. */
  prePull: boolean;
  time: string;
  name: string;
  target: string;
  resources: Record<string, number>;
}

export interface SampleLog {
  rows: SampleRow[];
  /** The resource columns this sample actually has, in RESOURCE_ORDER then alphabetical. */
  columns: string[];
  prePullCount: number;
}

/** "-1.5 s" before the pull, "1.9 s" after it. One decimal, because the engine's own ticks are 10ms. */
export function sampleTime(atMs: number): string {
  return `${(atMs / 1000).toFixed(1)} s`;
}

export function sampleLog(sample: readonly SampleCast[] | undefined, names: ActionNames | null): SampleLog {
  if (sample === undefined || sample.length === 0) return { rows: [], columns: [], prePullCount: 0 };

  const seen = new Set<string>();
  for (const cast of sample) {
    for (const key of Object.keys(cast.resources ?? {})) seen.add(key);
  }
  const known = RESOURCE_ORDER.filter((key) => seen.has(key));
  const extra = [...seen].filter((key) => !RESOURCE_ORDER.includes(key)).sort();

  const rows = sample.map((cast, index) => ({
    // The index is in the key because two casts can genuinely share an instant and an
    // action, and Svelte 5 throws on a repeated {#each} key.
    key: `${index}-${cast.at_ms}-${cast.action}`,
    atMs: cast.at_ms,
    prePull: cast.at_ms < 0,
    time: sampleTime(cast.at_ms),
    name: resolveActionName(cast.action, names),
    target: cast.target ?? '',
    resources: { ...(cast.resources ?? {}) },
  }));

  return { rows, columns: [...known, ...extra], prePullCount: rows.filter((row) => row.prePull).length };
}
