// web/src/lib/guides/rail-stats.ts
// The spec guide's Stat priority rail card (rebuild spec §4.C.3, round-1 mock review
// finding 4): one line per stat the guide's own `statPriority` names, each carrying its
// real per-point scale number -- `StatWeightsPanel`'s own convention (scale factor
// normalized so the top SIGNIFICANT per-point stat at this band reads exactly 1.00),
// computed fresh over just these stats rather than reused from `panel-view.ts`'s own table.
//
// Ordering (2026-10-05 audit finding): the card's row order comes from the sim, not from
// the guide's own `statPriority` array -- the per-point stats with a value sorted by scale
// factor descending, then the per-1% stats, then "No gain", then "Not significant". A
// guide's written order can disagree with the sim; the card always follows the sim.
//
// Per-percent stats (2026-11-02 launch defect): haste and expertise have no rating in this
// client. The engine's expertise point takes one percent off the boss's dodge and parry
// (`sim/cmd/leveling-bis/expertisecap.go`), so its weight is per 1%, not per point; read
// against the per-point stats it topped every tank's table ("Expertise 1547.3 per point").
// Both sit outside the normalization and print "per 1%".
//
// Sign (same defect): a weight inside its own error reads "Not significant", and a weight
// below zero beyond its error reads "No gain", never a negative number.
import { statLabelForSpec } from '../sim/weights-display';
import type { BisStatWeight } from '../bis/types';
import { guidesCopy } from './copy';

/** The stats the rail prints per 1% rather than per point. */
export type PerPercentStat = 'haste' | 'expertise';

/** How a row reads: its number, "Not significant", or "No gain". */
export type RailStatReading = 'value' | 'not-significant' | 'no-gain';

export interface RailStatRow {
  label: string;
  /** Set only when `reading` is `'value'`: the row itself is never dropped (tenet 4). */
  value: number | undefined;
  reading: RailStatReading;
  /** Set on a per-1% stat (`PerPercentStat`), which sits outside the normalization. */
  perPercent?: PerPercentStat;
  /** The band's own rate (DPS, HPS or tank score) per point of the stat, per 1% on a
   *  per-percent row; set only when the file publishes it and `reading` is `'value'`. */
  perPoint?: number;
}

const PER_PERCENT_STATS: Readonly<Record<string, PerPercentStat>> = {
  melee_haste: 'haste',
  spell_haste: 'haste',
  expertise: 'expertise',
};

function perPercentOf(stat: string): PerPercentStat | undefined {
  return PER_PERCENT_STATS[stat];
}

/**
 * Every guide's own `statPriority` frontmatter says "Critical strike"; `statLabelForSpec`'s
 * engine vocabulary says "Crit" (`sim/copy.ts`'s pinned `statLabel.crit`) -- the one place
 * these two, otherwise identical, vocabularies diverge.
 */
const GUIDE_LABEL_ALIASES: Readonly<Record<string, string>> = { 'Critical strike': 'Crit' };

/** How `weight` reads: inside its error (or flagged insignificant) it has no measurable
 *  value; below zero beyond its error it adds nothing. */
export function readingOf(weight: BisStatWeight | undefined): RailStatReading {
  if (weight === undefined || (weight.insignificant ?? false)) return 'not-significant';
  if (Math.abs(weight.weight) <= (weight.error ?? 0)) return 'not-significant';
  return weight.weight < 0 ? 'no-gain' : 'value';
}

/** Sort tier: 0 = per-point stat with a value (by scale factor, descending), 1 = per-1% stat
 *  with a value (same), 2 = "No gain", 3 = "Not significant" or unmatched. */
function sortTier(weight: BisStatWeight | undefined): 0 | 1 | 2 | 3 {
  const reading = readingOf(weight);
  if (reading === 'not-significant') return 3;
  if (reading === 'no-gain') return 2;
  return perPercentOf(weight!.stat) === undefined ? 0 : 1;
}

function scaleOf(weight: BisStatWeight): number {
  return weight.scale_factor ?? weight.weight;
}

interface MatchedRow {
  label: string;
  weight: BisStatWeight | undefined;
}

function valueOf(
  weight: BisStatWeight,
  perPercent: PerPercentStat | undefined,
  top: number | undefined,
  hasteScaleFactor: number | null,
): number | undefined {
  // Haste reads the band's own published per-1% figure (the BiS page's caption) when it has
  // one; every other row is its scale factor over the top per-point stat shown here.
  if (perPercent === 'haste' && hasteScaleFactor !== null) return hasteScaleFactor;
  return top === undefined ? undefined : scaleOf(weight) / top;
}

function rowFor(
  { label, weight }: MatchedRow,
  top: number | undefined,
  hasteScaleFactor: number | null,
): RailStatRow {
  const perPercent = weight === undefined ? undefined : perPercentOf(weight.stat);
  const base = perPercent === undefined ? { label } : { label, perPercent };
  const reading = readingOf(weight);
  if (reading !== 'value' || weight === undefined) return { ...base, reading, value: undefined };
  const value = valueOf(weight, perPercent, top, hasteScaleFactor);
  if (value === undefined) return { ...base, reading: 'not-significant', value: undefined };
  return weight.dps_per_point === undefined
    ? { ...base, reading, value }
    : { ...base, reading, value, perPoint: weight.dps_per_point };
}

/**
 * `statPriority` (the guide's own frontmatter order) matched against `weights` by its
 * class-aware display label (`statLabelForSpec`), the same vocabulary `StatPriorityPills`
 * renders. A label that matches no weight row reads "Not significant" rather than throw.
 */
export function railStatRows(
  statPriority: readonly string[],
  weights: readonly BisStatWeight[],
  spec: string,
  hasteScaleFactor: number | null = null,
): RailStatRow[] {
  const byLabel = new Map(weights.map((w) => [statLabelForSpec(w.stat, spec), w]));
  const unordered: MatchedRow[] = statPriority.map((label) => ({
    label,
    weight: byLabel.get(GUIDE_LABEL_ALIASES[label] ?? label),
  }));
  const perPointScales = unordered
    .filter((row) => sortTier(row.weight) === 0)
    .map((row) => scaleOf(row.weight!));
  const top = perPointScales.length > 0 ? Math.max(...perPointScales) : undefined;

  // `Array#sort` is stable, so ties (and the two tiers with no number) keep the guide's
  // original relative order.
  const ordered = [...unordered].sort((a, b) => {
    const tierA = sortTier(a.weight);
    const tierB = sortTier(b.weight);
    if (tierA !== tierB) return tierA - tierB;
    if (tierA > 1) return 0;
    return scaleOf(b.weight!) - scaleOf(a.weight!);
  });
  return ordered.map((row) => rowFor(row, top, hasteScaleFactor));
}

/** The weight cell's text: "1.00", "2.5 per 1% haste", "Not significant" or "No gain". */
export function railStatValueText(row: RailStatRow): string {
  if (row.reading === 'no-gain') return guidesCopy.noGain;
  if (row.value === undefined) return guidesCopy.notSignificant;
  return row.perPercent === undefined
    ? row.value.toFixed(2)
    : guidesCopy.perPercentValue(row.value.toFixed(1), row.perPercent);
}
