// web/src/lib/bis/weight-reading.ts
// How one published stat weight reads, the one rule the BiS page's weights panel and the spec
// guide's stat table share (2026-11-02 launch defects):
//
// - Per-percent stats. Haste and expertise have no rating in this client. The engine's
//   expertise point takes one percent off the boss's dodge and parry
//   (the engine's attack table), so its weight is per 1%, not per point; read
//   against the per-point stats it topped every tank's table ("Expertise 1547.3 per point").
//   Both sit outside the per-point normalization and read "per 1%".
// - Sign. A weight inside its own error (or flagged insignificant) reads "Not significant";
//   a weight below zero beyond its error reads "No gain". Neither prints a negative number.
import type { BisStatWeight } from './types';

/** The stats that read per 1% rather than per point. */
export type PerPercentStat = 'haste' | 'expertise';

/** How a weight reads: its number, "Not significant", or "No gain". */
export type WeightReading = 'value' | 'not-significant' | 'no-gain';

const PER_PERCENT_STATS: Readonly<Record<string, PerPercentStat>> = {
  melee_haste: 'haste',
  spell_haste: 'haste',
  expertise: 'expertise',
};

/** The per-1% kind of `stat`, or undefined for a per-point stat. */
export function perPercentOf(stat: string): PerPercentStat | undefined {
  return PER_PERCENT_STATS[stat];
}

/** How `weight` reads: inside its error (or flagged insignificant) it has no measurable
 *  value; below zero beyond its error it adds nothing. */
export function readingOf(weight: BisStatWeight | undefined): WeightReading {
  if (weight === undefined || (weight.insignificant ?? false)) return 'not-significant';
  if (Math.abs(weight.weight) <= (weight.error ?? 0)) return 'not-significant';
  return weight.weight < 0 ? 'no-gain' : 'value';
}
