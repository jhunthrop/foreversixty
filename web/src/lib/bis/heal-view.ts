// web/src/lib/bis/heal-view.ts
// The healer variant of the "This set" panel, as plain data: a healing sim ranks gear under a
// stated incoming-damage profile, so every figure here is built together with that profile's
// name and never on its own. Pure and node-free; all wording lives in `copy.ts`.
import { bisCopy } from './copy';
import type { BisBand, BisHealMetrics, BisHealProfile, RateUnit } from './types';
import { clockSeconds } from '../format/duration';

/** `HPS` for a healer band, `DPS` for every other band. */
export function rateUnitOf(band: BisBand): RateUnit {
  return rateUnitForRole(band.role);
}

/** The same rule from a role string (the spec catalogue's `role`, a band's `role`). */
export function rateUnitForRole(role: string | undefined): RateUnit {
  return role === 'healer' ? 'HPS' : 'DPS';
}

/** The ranker caps `mana_lasts_sec` here; it is the only "never ran out" signal when the
 *  file carries no profile (so no fight length) to compare against. */
const MANA_LASTS_CAP_SEC = 3600;

export type HealMetricKey = 'overheal' | 'mana' | 'hpm';

export interface HealMetricCell {
  key: HealMetricKey;
  label: string;
  /** Short face of the cell, e.g. "14%". */
  value: string;
  /** The same fact as a sentence ("14% overheal", "Mana lasts the whole 5:00 fight"). */
  line: string;
}

export interface HealerSetView {
  /** Effective healing per second, one decimal. */
  figure: string;
  unit: 'HPS';
  /** The profile's own label, or the stated fallback when the file names none. */
  profileLabel: string;
  /** "Effective healing per second under <profile>". */
  caption: string;
  metrics: HealMetricCell[];
}

function isHealMetrics(metrics: BisBand['metrics']): metrics is BisHealMetrics {
  return metrics !== null && metrics !== undefined && 'hps' in metrics;
}

/** `192.7` -> `"3:12"`: rounded down, so a mana bar never lasts longer than the sim saw. */
function manaClock(totalSeconds: number): string {
  return clockSeconds(Math.floor(totalSeconds));
}

function manaCell(manaLastsSec: number, durationSec: number | undefined): HealMetricCell {
  const label = bisCopy.healerManaLabel;
  if (manaLastsSec >= (durationSec ?? MANA_LASTS_CAP_SEC)) {
    return {
      key: 'mana',
      label,
      value: bisCopy.healerManaWholeFightValue,
      line: bisCopy.healerManaWholeFightLine(durationSec === undefined ? undefined : manaClock(durationSec)),
    };
  }
  const clock = manaClock(manaLastsSec);
  return {
    key: 'mana',
    label,
    value: bisCopy.healerManaOutValue(clock),
    line: bisCopy.healerManaOutLine(clock),
  };
}

/** The cells a band's `metrics` can fill, in display order; a missing `metrics` yields none. */
function metricCells(metrics: BisHealMetrics | undefined, durationSec: number | undefined): HealMetricCell[] {
  if (metrics === undefined) return [];
  return [
    {
      key: 'overheal',
      label: bisCopy.healerOverhealLabel,
      value: bisCopy.healerOverhealValue(metrics.overheal_pct),
      line: bisCopy.healerOverhealLine(metrics.overheal_pct),
    },
    manaCell(metrics.mana_lasts_sec, durationSec),
    {
      key: 'hpm',
      label: bisCopy.healerHpmLabel,
      value: bisCopy.healerHpmValue(metrics.hpm),
      line: bisCopy.healerHpmLine(metrics.hpm),
    },
  ];
}

/** The healer "This set" panel's data for `band`, or `undefined` for a band that is not a
 *  healer's. `profile` is the file's `heal_profile`; without it the panel still names the
 *  band's own `profile` id, or says the profile is unstated -- never a bare figure. */
export function healerSetViewFor(
  band: BisBand,
  profile: BisHealProfile | undefined,
): HealerSetView | undefined {
  if (band.role !== 'healer') return undefined;
  const metrics = isHealMetrics(band.metrics) ? band.metrics : undefined;
  const profileLabel = profile?.label ?? band.profile ?? bisCopy.healerProfileUnstated;
  return {
    figure: (metrics?.hps ?? band.set_dps).toFixed(1),
    unit: 'HPS',
    profileLabel,
    caption: bisCopy.healerFigureCaption(profileLabel),
    metrics: metricCells(metrics, profile?.duration_sec),
  };
}
