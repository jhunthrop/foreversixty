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
// Per-percent stats and sign follow the rule the BiS weights panel shares
// (`lib/bis/weight-reading.ts`): haste and expertise read per 1% outside the normalization,
// a weight inside its error reads "Not significant", and one below zero beyond its error
// reads "No gain", never a negative number; haste no item of the band carries reads "Not on
// any item". The numbered pills rank only the rows that read as a value (`rankedStatLabels`),
// so the pills and the table can never disagree.
import { statLabelForSpec } from '../sim/weights-display';
import type { BisStatWeight } from '../bis/types';
import {
  bandReadingOf,
  isRanked,
  perPercentOf,
  readingOf,
  type BandWeightReading,
  type PerPercentStat,
} from '../bis/weight-reading';
import { guidesCopy } from './copy';

export { readingOf, type PerPercentStat };

/** How a row reads: its number, "Not significant", "No gain" or "Not on any item". */
export type RailStatReading = BandWeightReading;

/** What the band says about haste beyond its weight row. */
export interface HasteFacts {
  /** The band's own published per-1% haste figure (`haste_scale_factor`). */
  scaleFactor: number | null;
  /** Whether any item the band considered carries haste (`haste_on_items`). */
  onItems: boolean;
}

const UNKNOWN_HASTE: HasteFacts = { scaleFactor: null, onItems: true };

export interface RailStatRow {
  label: string;
  /** Set when `reading` is `'value'`, and on a `'not-on-items'` haste row for the note under
   *  the table: the row itself is never dropped (tenet 4). */
  value: number | undefined;
  reading: RailStatReading;
  /** Set on a per-1% stat (`PerPercentStat`), which sits outside the normalization. */
  perPercent?: PerPercentStat;
  /** The band's own rate (DPS, HPS or tank score) per point of the stat, per 1% on a
   *  per-percent row; set only when the file publishes it and `reading` is `'value'`. */
  perPoint?: number;
  /** Rating points per 1% on a rating stat (crit, hit), so a point reads as a rating point. */
  ratingFactor?: number;
}

/**
 * Every guide's own `statPriority` frontmatter says "Critical strike"; `statLabelForSpec`'s
 * engine vocabulary says "Crit" (`sim/copy.ts`'s pinned `statLabel.crit`) -- the one place
 * these two, otherwise identical, vocabularies diverge.
 */
const GUIDE_LABEL_ALIASES: Readonly<Record<string, string>> = { 'Critical strike': 'Crit' };

/** Sort tier: 0 = per-point stat with a value (by scale factor, descending), 1 = per-1% stat
 *  with a value (same), 2 = "No gain", 3 = "Not on any item", 4 = "Not significant" or
 *  unmatched. */
function sortTier(weight: BisStatWeight | undefined, hasteOnItems: boolean): 0 | 1 | 2 | 3 | 4 {
  const reading = bandReadingOf(weight, hasteOnItems);
  if (reading === 'not-significant') return 4;
  if (reading === 'not-on-items') return 3;
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

function rowFor({ label, weight }: MatchedRow, top: number | undefined, haste: HasteFacts): RailStatRow {
  const perPercent = weight === undefined ? undefined : perPercentOf(weight.stat);
  const base = perPercent === undefined ? { label } : { label, perPercent };
  const reading = bandReadingOf(weight, haste.onItems);
  if (reading === 'not-on-items' && weight !== undefined) {
    // Still carries its per-1% figure, for the note under the table that says what it is worth;
    // a weight inside its error has no figure worth printing.
    const value = readingOf(weight) === 'value' ? valueOf(weight, perPercent, top, haste.scaleFactor) : undefined;
    return { ...base, reading, value };
  }
  if (reading !== 'value' || weight === undefined) return { ...base, reading, value: undefined };
  const value = valueOf(weight, perPercent, top, haste.scaleFactor);
  if (value === undefined) return { ...base, reading: 'not-significant', value: undefined };
  const rated =
    weight.unit === 'rating' && weight.rating_factor !== undefined
      ? { ratingFactor: weight.rating_factor }
      : {};
  return weight.dps_per_point === undefined
    ? { ...base, ...rated, reading, value }
    : { ...base, ...rated, reading, value, perPoint: weight.dps_per_point };
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
  haste: HasteFacts = UNKNOWN_HASTE,
): RailStatRow[] {
  const byLabel = new Map(weights.map((w) => [statLabelForSpec(w.stat, spec), w]));
  // Expertise with no weight is a stat the band's character cannot gain (`weightsForGainedStats`
  // dropped it): no row, so the pills and the table never name it.
  const expertiseLabel = statLabelForSpec('expertise', spec);
  const unordered: MatchedRow[] = statPriority
    .map((label) => ({ label, weight: byLabel.get(GUIDE_LABEL_ALIASES[label] ?? label) }))
    .filter(
      (row) =>
        !(row.weight === undefined && (GUIDE_LABEL_ALIASES[row.label] ?? row.label) === expertiseLabel),
    );
  const perPointScales = unordered
    .filter((row) => sortTier(row.weight, haste.onItems) === 0)
    .map((row) => scaleOf(row.weight!));
  const top = perPointScales.length > 0 ? Math.max(...perPointScales) : undefined;

  // `Array#sort` is stable, so ties (and the two tiers with no number) keep the guide's
  // original relative order.
  const ordered = [...unordered].sort((a, b) => {
    const tierA = sortTier(a.weight, haste.onItems);
    const tierB = sortTier(b.weight, haste.onItems);
    if (tierA !== tierB) return tierA - tierB;
    if (tierA > 1) return 0;
    return scaleOf(b.weight!) - scaleOf(a.weight!);
  });
  return ordered.map((row) => rowFor(row, top, haste));
}

/** The numbered pills' stats: the rows that read as a value, in the table's order. A spec
 *  with no weights yet falls back to the guide's own list. */
export function rankedStatLabels(rows: readonly RailStatRow[], fallback: readonly string[]): string[] {
  if (rows.length === 0) return [...fallback];
  return rows.filter((row) => isRanked(row.reading)).map((row) => row.label);
}

/** The weight cell's text: "1.00", "2.5 per 1% haste", "Not significant", "No gain" or
 *  "Not on any item". */
export function railStatValueText(row: RailStatRow): string {
  if (row.reading === 'no-gain') return guidesCopy.noGain;
  if (row.reading === 'not-on-items') return guidesCopy.notOnItems;
  if (row.value === undefined) return guidesCopy.notSignificant;
  return row.perPercent === undefined
    ? row.value.toFixed(2)
    : guidesCopy.perPercentValue(row.value.toFixed(1), row.perPercent);
}
