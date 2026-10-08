// web/src/lib/bis/tank-view.ts
// The tank/DPS decision for a BiS band, in one place: which headline the "this set" panel
// shows, which unit a slot's figures are in, and how each tank number is written. Everything
// here is pure -- the Astro components only render what they are handed.
import { bisCopy, tankCopy } from './copy';
import type { BisBand, BisRole, TankMetrics } from './types';

export const DEFAULT_BIS_ROLE: BisRole = 'dps';
export const TANK_ROLE: BisRole = 'tank';

/** The unit a slot's `sim_dps`/`dps_delta` figures are in: real damage, or a tank band's
 *  composite tank-score points. */
export type SlotScoreUnit = 'dps' | 'tank_score';

const NUMBER_LOCALE = 'en-US';
const WHOLE_NUMBER = new Intl.NumberFormat(NUMBER_LOCALE, { maximumFractionDigits: 0 });
const ONE_DECIMAL = new Intl.NumberFormat(NUMBER_LOCALE, {
  minimumFractionDigits: 1,
  maximumFractionDigits: 1,
});
const PERCENT_DECIMALS = 1;
const PERCENT_FACTOR = 100;
/** Under this share a non-zero chance of death is written "<0.1%" rather than "0.0%". */
const CHANCE_OF_DEATH_RESOLUTION = 0.0005;

export type TankFigureId = 'effective-health' | 'dtps' | 'chance-of-death' | 'tps';

export interface TankFigure {
  id: TankFigureId;
  label: string;
  value: string;
  title: string;
}

/** What the "this set" panel shows in place of the DPS figure on a tank band. */
export interface TankHeadline {
  /** Effective health, damage taken per second, chance of death, threat per second. */
  figures: TankFigure[];
  /** TMI and own damage, the quieter line under the four figures. */
  secondaryLine: string;
  secondaryTitle: string;
}

/** A band is a tank band when it says so; `normaliseBisFile` guarantees a tank band has
 *  its metrics. */
export function isTankBand(band: Pick<BisBand, 'role'>): boolean {
  return band.role === TANK_ROLE;
}

export function slotScoreUnitFor(band: Pick<BisBand, 'role'>): SlotScoreUnit {
  return isTankBand(band) ? 'tank_score' : 'dps';
}

/** The word that follows a figure in this unit: "+4.2 DPS", "+4.2 score". */
export function scoreUnitWord(unit: SlotScoreUnit): string {
  return unit === 'tank_score' ? 'score' : 'DPS';
}

export function formatWholeNumber(value: number): string {
  return WHOLE_NUMBER.format(value);
}

export function formatOneDecimal(value: number): string {
  return ONE_DECIMAL.format(value);
}

/** 0..1 as "12.3%"; a non-zero chance under the resolution reads "<0.1%". */
export function formatChanceOfDeath(chance: number): string {
  if (chance > 0 && chance < CHANCE_OF_DEATH_RESOLUTION) return tankCopy.chanceOfDeathBelowResolution;
  return `${(chance * PERCENT_FACTOR).toFixed(PERCENT_DECIMALS)}%`;
}

/** The four headline figures, in the order the panel shows them. */
export function tankHeadlineFor(metrics: TankMetrics): TankHeadline {
  return {
    figures: [
      {
        id: 'effective-health',
        label: tankCopy.effectiveHealthLabel,
        value: formatWholeNumber(metrics.effective_health),
        title: tankCopy.effectiveHealthTitle,
      },
      {
        id: 'dtps',
        label: tankCopy.dtpsLabel,
        value: formatOneDecimal(metrics.dtps),
        title: tankCopy.dtpsTitle,
      },
      {
        id: 'chance-of-death',
        label: tankCopy.chanceOfDeathLabel,
        value: formatChanceOfDeath(metrics.chance_of_death),
        title: tankCopy.chanceOfDeathTitle,
      },
      {
        id: 'tps',
        label: tankCopy.tpsLabel,
        value: formatOneDecimal(metrics.tps),
        title: tankCopy.tpsTitle,
      },
    ],
    secondaryLine: tankCopy.secondaryLine(formatWholeNumber(metrics.tmi), formatOneDecimal(metrics.dps)),
    secondaryTitle: tankCopy.tmiTitle,
  };
}

/** The headline for `band`, or `undefined` for every non-tank band (which keeps the DPS
 *  figure). */
export function tankHeadlineForBand(band: Pick<BisBand, 'role' | 'metrics'>): TankHeadline | undefined {
  return isTankBand(band) && band.metrics ? tankHeadlineFor(band.metrics) : undefined;
}

/** The band's own damage figure when it is a true headline: `set_dps` for a DPS band,
 *  `undefined` for a tank (its `set_dps` is own damage, not what a tank is judged on). */
export function headlineDpsOf(band: Pick<BisBand, 'role' | 'set_dps'>): number | undefined {
  return isTankBand(band) ? undefined : band.set_dps;
}

/** A runner-up's gap-from-the-pick line in the band's unit. */
export function gapLabelFor(delta: number, unit: SlotScoreUnit): string {
  return unit === 'tank_score' ? tankCopy.alternativeGapLabel(delta) : bisCopy.alternativeGapLabel(delta);
}

/** One-line /bis index summary of a tank's level-60 set. */
export function tankIndexSummaryOf(band: Pick<BisBand, 'role' | 'metrics'>): string | undefined {
  return isTankBand(band) && band.metrics
    ? tankCopy.indexSpecSummary(formatWholeNumber(band.metrics.effective_health))
    : undefined;
}
