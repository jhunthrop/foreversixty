// web/src/lib/tiers/tier-list.ts
// The tier list's ranking, as pure functions of the BiS bands (design/specs/2026-10-09-
// tier-list.md, section 9). Which number sorts a role, what a gap and a bar are, where the
// tier a gap falls in and which rows tie are rules, not nightly figures: every value comes from the
// published band at build time and nothing here is typed in.
import type { BisBand, BisRole, SpecCatalogEntry, TankMetrics } from '../bis/types';
import { isTankMetrics } from '../bis/tank-view';
import { bisBandPath } from '../bis/band-path';
import { MAX_BAND, MIN_BAND } from '../bis/band-levels';
import { HEALER_ERROR_Z, TIE_MARGIN_PERCENT, TIER_BANDS, type TierBand, type TierLetter } from './tier-rules';

/** The roles, in the order of the page's tabs. */
export const TIER_ROLES: readonly BisRole[] = ['dps', 'tank', 'healer'];

const PERCENT = 100;

export interface TierInput {
  entry: SpecCatalogEntry;
  band: BisBand;
  /** "View ranking sim": the band opened and run in the simulator. Absent where /sim cannot
   *  rerun the figure (tank and healer bands: the browser sim runs damage specs only). */
  simHref?: string;
}

export type TierTie = 'above' | 'below' | 'both';

/** A spec a row ties with: enough to name it ("≈ tie with Arms") and to say when the two sit
 *  in different letters. */
export interface TierTieMate {
  classSlug: string;
  specSlug: string;
  /** "Arms" */
  name: string;
  tier: TierLetter | null;
}

export interface TierRow {
  /** 1-based. Healers the error cannot tell apart share the rank of the first of them. */
  rank: number;
  /** False on a healer list published without the measured error and tuning bounds: the
   *  rows are sorted by value, but the page claims neither an order nor a tie. */
  ranked: boolean;
  spec: string;
  classSlug: string;
  specSlug: string;
  name: string;
  role: BisRole;
  /** "Night Elf": the race the band was simmed as. */
  race: string;
  /** The sorted figure: set DPS, effective healing per second, or damage taken per second. */
  metric: number;
  /** Percent behind the top spec (tank: percent more damage taken than the best). 0 on the top row. */
  gapPercent: number;
  /** 0..1 bar length: value over top (tank: best over value). */
  fraction: number;
  /** The neighbour(s) within the tie margin (a healer: within error), or `null`. */
  tie: TierTie | null;
  /** The specs this row ties with, best first: on a DPS list every spec within the tie margin
   *  of this one (no chain carries a mark past the margin); on a healer list the neighbours the
   *  error cannot order it against. Empty when `tie` is `null`. */
  tiedWith: readonly TierTieMate[];
  /** S to D from the gap to the top. `null` on tank and healer lists, which carry no letters. */
  tier: TierLetter | null;
  lowConfidence: boolean;
  /** The planner link's talent string. */
  talents: string;
  /** The ranking sim behind the number (`TierInput.simHref`); absent on tank and healer rows. */
  simHref?: string;
  /** Tank only. `health` is the maximum health, absent on a band ranked before it was published. */
  health?: number;
  threat?: number;
  isBestHealth?: boolean;
  isBestThreat?: boolean;
}

export interface TierGroup {
  letter: TierLetter;
  /** Percent behind the top where the tier starts (inclusive) and where it ends (`null`: no end). */
  from: number;
  to: number | null;
  rows: TierRow[];
}

/** `night-elf` -> `Night Elf`. */
export function raceLabel(race: string): string {
  return race
    .split('-')
    .map((part) => (part.length === 0 ? part : part[0]!.toUpperCase() + part.slice(1)))
    .join(' ');
}

export function roleOf(band: Pick<BisBand, 'role'>): BisRole {
  return band.role ?? 'dps';
}

function tankMetricsOf(band: BisBand): TankMetrics {
  if (!isTankMetrics(band.metrics)) {
    throw new Error(`tier list: tank band ${band.spec} ${band.faction} has no tank metrics`);
  }
  return band.metrics;
}

/** The figure a role is sorted on: damage taken per second for a tank, `set_dps` otherwise
 *  (a healer's equals its effective healing per second). */
export function sortMetricOf(band: BisBand): number {
  return roleOf(band) === 'tank' ? tankMetricsOf(band).dtps : band.set_dps;
}

interface Scored {
  input: TierInput;
  metric: number;
}

function sortScored(scored: Scored[], tank: boolean): Scored[] {
  const direction = tank ? 1 : -1;
  return [...scored].sort(
    (a, b) => direction * (a.metric - b.metric) || a.input.entry.name.localeCompare(b.input.entry.name),
  );
}

function gapAndFraction(metric: number, top: number, tank: boolean): { gap: number; fraction: number } {
  return tank
    ? { gap: (metric / top - 1) * PERCENT, fraction: top / metric }
    : { gap: (1 - metric / top) * PERCENT, fraction: metric / top };
}

const GAP_PRECISION = 1e6;

/** The tier a gap falls in. The gap is rounded past float noise first, so 10% behind is B, not A. */
export function tierOf(gapPercent: number): TierLetter {
  const gap = Math.round(gapPercent * GAP_PRECISION) / GAP_PRECISION;
  return TIER_BANDS.findLast((band) => gap >= band.from)!.letter;
}

function withinTieMargin(metric: number, neighbour: number): boolean {
  return Math.abs(1 - metric / neighbour) * PERCENT <= TIE_MARGIN_PERCENT;
}

/** Both rows of a tied pair carry the mark; a row tied on both sides says so. */
function tieMarks(metrics: readonly number[]): (TierTie | null)[] {
  return metrics.map((metric, index) => {
    const above = index > 0 && withinTieMargin(metric, metrics[index - 1]!);
    const below = index < metrics.length - 1 && withinTieMargin(metrics[index + 1]!, metric);
    if (above && below) return 'both';
    if (above) return 'above';
    return below ? 'below' : null;
  });
}

/** The indexes a DPS row ties with: every other row within the margin of it, never one only
 *  chained to it through a third spec. Sorted best first, so the margin is read off the higher one. */
function marginMates(metrics: readonly number[], index: number): number[] {
  const metric = metrics[index]!;
  return metrics.flatMap((other, j) => {
    if (j === index) return [];
    const [lower, higher] = j < index ? [metric, other] : [other, metric];
    return withinTieMargin(lower, higher) ? [j] : [];
  });
}

/** The neighbours a mark names: above for 'above', below for 'below', both for 'both'. */
function neighbourMates(tie: TierTie | null, index: number): number[] {
  if (tie === null) return [];
  if (tie === 'both') return [index - 1, index + 1];
  return [tie === 'above' ? index - 1 : index + 1];
}

/** One measurement of a healer: effective HPS and its standard error. */
interface HealerPoint {
  hps: number;
  error: number;
}

/** A healer's measurements: the tuned profile and each tuning bound. `null` until the file
 *  publishes the error and the bounds, and without them no order is measured. */
function healerPoints(band: BisBand, metric: number): HealerPoint[] | null {
  const metrics = band.metrics;
  if (metrics === null || metrics === undefined || !('hps' in metrics) || metrics.hps_error === undefined) {
    return null;
  }
  const bounds = metrics.hps_at_bounds ?? [];
  if (bounds.length === 0) return null;
  return [
    { hps: metric, error: metrics.hps_error },
    ...bounds.map((row) => ({ hps: row.hps, error: row.hps_error })),
  ];
}

/** Whether the healer above leads the one below at one measurement, beyond the error. */
function leadsBeyondError(above: HealerPoint, below: HealerPoint): boolean {
  return above.hps - below.hps > HEALER_ERROR_Z * Math.hypot(above.error, below.error);
}

/** Whether two neighbouring healers tie: unless the one above leads beyond the error at the
 *  tuned profile and at every tuning bound, the order is the error's or the tuning's, not the
 *  healers'. A pair missing the error or the bounds is never ordered (owner ruling, round 5). */
function healerTiedWith(above: Scored, below: Scored): boolean {
  const pointsAbove = healerPoints(above.input.band, above.metric);
  const pointsBelow = healerPoints(below.input.band, below.metric);
  if (pointsAbove === null || pointsBelow === null || pointsAbove.length !== pointsBelow.length) return true;
  return !pointsAbove.every((point, index) => leadsBeyondError(point, pointsBelow[index]!));
}

/** Whether every healer band carries the error and both bounds, so the list can order at all. */
export function healerOrderMeasured(inputs: readonly TierInput[]): boolean {
  return inputs
    .filter((input) => roleOf(input.band) === 'healer')
    .every((input) => healerPoints(input.band, sortMetricOf(input.band)) !== null);
}

/** Healer tie marks and ranks. A healer within error of the one above shares its rank, and
 *  the chain carries on, so no two healers are ever ordered by a gap the error cannot tell. */
function healerTiesAndRanks(sorted: readonly Scored[]): { ties: (TierTie | null)[]; ranks: number[] } {
  const tiedAbove = sorted.map((scored, index) => index > 0 && healerTiedWith(sorted[index - 1]!, scored));
  const ranks = sharedRanks(tiedAbove);
  const ties = tiedAbove.map((above, index): TierTie | null => {
    const below = tiedAbove[index + 1] === true;
    if (above && below) return 'both';
    if (above) return 'above';
    return below ? 'below' : null;
  });
  return { ties, ranks };
}

/** Ranks where a row tied with the one above shares that row's rank, down the chain. */
function sharedRanks(tiedAbove: readonly boolean[]): number[] {
  return tiedAbove.reduce<number[]>((acc, tied, index) => [...acc, tied ? acc[index - 1]! : index + 1], []);
}

/** The tie marks and ranks of one sorted role: healers by error, DPS and tanks by the fixed
 *  margin. Healers and tanks within it share a rank (a short list has no letter to carry the
 *  tie); DPS rows keep their own. */
function tiesAndRanks(
  sorted: readonly Scored[],
  role: BisRole,
  ranked: boolean,
): { ties: (TierTie | null)[]; ranks: number[] } {
  if (!ranked) return { ties: sorted.map(() => null), ranks: sorted.map((_scored, index) => index + 1) };
  if (role === 'healer') return healerTiesAndRanks(sorted);
  const ties = tieMarks(sorted.map((scored) => scored.metric));
  const ranks =
    role === 'tank'
      ? sharedRanks(ties.map((tie) => tie === 'above' || tie === 'both'))
      : sorted.map((_scored, index) => index + 1);
  return { ties, ranks };
}

function bestOf(values: readonly number[]): number {
  return Math.max(...values);
}

/** The ranked rows of one role, best first. An empty role gives an empty list. */
export function rankRole(inputs: readonly TierInput[], role: BisRole): TierRow[] {
  const tank = role === 'tank';
  const pool = inputs.filter((input) => roleOf(input.band) === role);
  if (pool.length === 0) return [];
  const sorted = sortScored(
    pool.map((input) => ({ input, metric: sortMetricOf(input.band) })),
    tank,
  );
  const top = sorted[0]!.metric;
  // A healer list published without the measured error claims neither an order nor a tie.
  const ranked = role !== 'healer' || healerOrderMeasured(pool);
  const { ties, ranks } = tiesAndRanks(sorted, role, ranked);
  const sortedMetrics = sorted.map((scored) => scored.metric);
  const tiers = sortedMetrics.map((metric) =>
    roleHasTiers(role) ? tierOf(gapAndFraction(metric, top, tank).gap) : null,
  );
  const mateOf = (j: number): TierTieMate => {
    const { entry } = sorted[j]!.input;
    return {
      classSlug: entry.class_slug,
      specSlug: entry.spec_slug,
      name: entry.name,
      tier: tiers[j] ?? null,
    };
  };
  const matesOf = (index: number): TierTieMate[] => {
    const tie = ties[index] ?? null;
    if (tie === null) return [];
    const indexes = role === 'healer' ? neighbourMates(tie, index) : marginMates(sortedMetrics, index);
    return indexes.map(mateOf);
  };
  const healths = tank ? sorted.map((s) => tankMetricsOf(s.input.band).health) : [];
  // Health is compared only when every band publishes it: a column half filled ranks nothing.
  const bestHealth = healths.every((h) => h !== undefined) ? bestOf(healths as number[]) : undefined;
  const bestThreat = tank ? bestOf(sorted.map((s) => tankMetricsOf(s.input.band).tps)) : 0;
  return sorted.map(({ input, metric }, index): TierRow => {
    const { entry, band } = input;
    const { gap, fraction } = gapAndFraction(metric, top, tank);
    const row: TierRow = {
      rank: ranks[index]!,
      ranked,
      spec: entry.spec,
      classSlug: entry.class_slug,
      specSlug: entry.spec_slug,
      name: entry.name,
      role,
      race: raceLabel(band.race),
      metric,
      gapPercent: index === 0 ? 0 : gap,
      fraction,
      tie: ties[index] ?? null,
      tiedWith: matesOf(index),
      tier: tiers[index] ?? null,
      lowConfidence: band.weights_low_confidence === true,
      talents: band.talents,
      ...(input.simHref === undefined ? {} : { simHref: input.simHref }),
    };
    if (!tank) return row;
    const metrics = tankMetricsOf(band);
    return {
      ...row,
      ...(bestHealth === undefined || metrics.health === undefined
        ? {}
        : { health: metrics.health, isBestHealth: metrics.health === bestHealth }),
      threat: metrics.tps,
      isBestThreat: metrics.tps === bestThreat,
    };
  });
}

function bandEnd(bands: readonly TierBand[], index: number): number | null {
  return bands[index + 1]?.from ?? null;
}

/** A DPS list cut into its tiers, best first. A tier with no spec is not drawn. */
export function groupByTier(rows: readonly TierRow[]): TierGroup[] {
  return TIER_BANDS.flatMap((band, index): TierGroup[] => {
    const inTier = rows.filter((row) => row.tier === band.letter);
    return inTier.length === 0
      ? []
      : [{ letter: band.letter, from: band.from, to: bandEnd(TIER_BANDS, index), rows: inTier }];
  });
}

/** Whether a role's list carries tier letters: the DPS list only (20 rows; the others are short). */
export function roleHasTiers(role: BisRole): boolean {
  return role === 'dps';
}

export interface TierRowHrefs {
  bis: string;
  guide: string;
  planner: string;
  /** `undefined` where the row's number cannot be rerun on /sim. */
  sim: string | undefined;
}

/** Where a row's next steps go: its BiS page, its guide, its level 60 build in the planner and,
 *  on a DPS row, the ranking sim behind its number. */
export function hrefsFor(
  row: Pick<TierRow, 'classSlug' | 'specSlug' | 'spec' | 'talents' | 'simHref'>,
): TierRowHrefs {
  const path = `${row.classSlug}/${row.specSlug}`;
  return {
    bis: bisBandPath(row.classSlug, row.specSlug, MAX_BAND, MIN_BAND),
    guide: `/guides/${path}`,
    planner: `/planner?spec=${row.spec}&talents=${encodeURIComponent(row.talents)}`,
    sim: row.simHref,
  };
}
