// web/src/lib/tiers/tier-list.ts
// The tier list's ranking, as pure functions of the BiS bands (design/specs/2026-10-09-
// tier-list.md, section 9). Which number sorts a role, what a gap and a bar are, where the
// tier a gap falls in and which rows tie are rules, not nightly figures: every value comes from the
// published band at build time and nothing here is typed in.
import type { BisBand, BisRole, SpecCatalogEntry, TankMetrics } from '../bis/types';
import { isTankMetrics } from '../bis/tank-view';
import { TIE_MARGIN_PERCENT, TIER_BANDS, type TierBand, type TierLetter } from './tier-rules';

/** The roles, in the order of the page's tabs. */
export const TIER_ROLES: readonly BisRole[] = ['dps', 'tank', 'healer'];

const PERCENT = 100;

export interface TierInput {
  entry: SpecCatalogEntry;
  band: BisBand;
}

export type TierTie = 'above' | 'below' | 'both';

export interface TierRow {
  rank: number;
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
  /** The neighbour(s) within the tie margin, or `null`. Always `null` on a tank list. */
  tie: TierTie | null;
  /** S to D from the gap to the top. `null` on tank and healer lists, which carry no letters. */
  tier: TierLetter | null;
  lowConfidence: boolean;
  /** The planner link's talent string. */
  talents: string;
  /** Tank only. */
  effectiveHealth?: number;
  threat?: number;
  isBestEffectiveHealth?: boolean;
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
function tieMarks(metrics: readonly number[], enabled: boolean): (TierTie | null)[] {
  return metrics.map((metric, index) => {
    if (!enabled) return null;
    const above = index > 0 && withinTieMargin(metric, metrics[index - 1]!);
    const below = index < metrics.length - 1 && withinTieMargin(metrics[index + 1]!, metric);
    if (above && below) return 'both';
    if (above) return 'above';
    return below ? 'below' : null;
  });
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
  const ties = tieMarks(
    sorted.map((scored) => scored.metric),
    !tank,
  );
  const bestEffectiveHealth = tank
    ? bestOf(sorted.map((s) => tankMetricsOf(s.input.band).effective_health))
    : 0;
  const bestThreat = tank ? bestOf(sorted.map((s) => tankMetricsOf(s.input.band).tps)) : 0;
  return sorted.map(({ input, metric }, index): TierRow => {
    const { entry, band } = input;
    const { gap, fraction } = gapAndFraction(metric, top, tank);
    const row: TierRow = {
      rank: index + 1,
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
      tier: roleHasTiers(role) ? tierOf(gap) : null,
      lowConfidence: band.weights_low_confidence === true,
      talents: band.talents,
    };
    if (!tank) return row;
    const metrics = tankMetricsOf(band);
    return {
      ...row,
      effectiveHealth: metrics.effective_health,
      threat: metrics.tps,
      isBestEffectiveHealth: metrics.effective_health === bestEffectiveHealth,
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
}

/** Where a row's three next steps go: its BiS page, its guide and its level 60 build in the planner. */
export function hrefsFor(row: Pick<TierRow, 'classSlug' | 'specSlug' | 'spec' | 'talents'>): TierRowHrefs {
  const path = `${row.classSlug}/${row.specSlug}`;
  return {
    bis: `/bis/${path}`,
    guide: `/guides/${path}`,
    planner: `/planner?spec=${row.spec}&talents=${encodeURIComponent(row.talents)}`,
  };
}
