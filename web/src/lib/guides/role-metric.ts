// web/src/lib/guides/role-metric.ts
// Which number a guide surface shows for a band, by the role it was ranked for (2026-11-02
// launch defect: every role read "DPS"). A damage band's figures are damage; a healer band's
// `*_dps*` keys carry effective healing (types.ts `BisBand.role`); a tank band's per-point
// figures are tank-score points (`score_unit: "tank_score"`, `tankCopy`), and its set is
// judged on damage taken per second, the figure /tiers/tank ranks tanks by. Never relabels
// one role's number as another's.
import { formatOneDecimal, formatWholeNumber, isTankMetrics } from '../bis/tank-view';
import type { BisBand, BisHealMetrics, BisRole } from '../bis/types';
import { guidesCopy } from './copy';

/** A band's role, `'dps'` for a file published before roles (`normaliseBisFile`'s default). */
export function roleOf(band: Pick<BisBand, 'role'>): BisRole {
  return band.role ?? 'dps';
}

/** The class landing card's one figure for a spec's set at level 60. */
export interface SetMetric {
  role: BisRole;
  value: number;
}

function healMetricsOf(band: Pick<BisBand, 'metrics'>): BisHealMetrics | undefined {
  const metrics = band.metrics;
  return metrics !== null && metrics !== undefined && 'hps' in metrics ? metrics : undefined;
}

/** Damage per second for a damage band, healing per second for a healer, damage taken per
 *  second for a tank; `undefined` for a tank band with no metrics (never a stand-in figure). */
export function setMetricOf(band: Pick<BisBand, 'role' | 'set_dps' | 'metrics'>): SetMetric | undefined {
  const role = roleOf(band);
  if (role === 'tank') {
    return isTankMetrics(band.metrics) ? { role, value: band.metrics.dtps } : undefined;
  }
  if (role === 'healer') return { role, value: healMetricsOf(band)?.hps ?? band.set_dps };
  return { role, value: band.set_dps };
}

/** "553.1 DPS", "553.1 HPS", "380 damage taken per second". */
export function setMetricLine(metric: SetMetric): string {
  return metric.role === 'tank'
    ? guidesCopy.setMetricLine(formatWholeNumber(metric.value), metric.role)
    : guidesCopy.setMetricLine(formatOneDecimal(metric.value), metric.role);
}
