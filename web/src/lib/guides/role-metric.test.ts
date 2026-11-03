// web/src/lib/guides/role-metric.test.ts
import { describe, expect, it } from 'vitest';
import type { TankMetrics } from '../bis/types';
import { roleOf, setMetricLine, setMetricOf } from './role-metric';

const TANK_METRICS: TankMetrics = {
  dtps: 380.39,
  tmi: 62,
  chance_of_death: 0.07,
  tps: 616.4,
  effective_health: 19635,
  dps: 189.2,
};

describe('setMetricOf', () => {
  it('reads set_dps for a damage band, and for a file published before roles', () => {
    expect(setMetricOf({ role: 'dps', set_dps: 553.12, metrics: undefined })).toEqual({
      role: 'dps',
      value: 553.12,
    });
    expect(setMetricOf({ set_dps: 12, metrics: undefined })).toEqual({ role: 'dps', value: 12 });
  });

  it('reads the healing figure for a healer band', () => {
    const metrics = { hps: 553.1, raw_hps: 700, overheal_pct: 0.2, mana_lasts_sec: 300, hpm: 1.2 };
    expect(setMetricOf({ role: 'healer', set_dps: 553.1, metrics })).toEqual({
      role: 'healer',
      value: 553.1,
    });
    expect(setMetricOf({ role: 'healer', set_dps: 742, metrics: null })).toEqual({
      role: 'healer',
      value: 742,
    });
  });

  it('reads damage taken per second for a tank, never its own damage', () => {
    expect(setMetricOf({ role: 'tank', set_dps: 189.2, metrics: TANK_METRICS })).toEqual({
      role: 'tank',
      value: 380.39,
    });
  });

  it('is undefined for a tank band with no metrics', () => {
    expect(setMetricOf({ role: 'tank', set_dps: 189.2, metrics: null })).toBeUndefined();
  });
});

describe('setMetricLine', () => {
  it('writes each role in its own unit', () => {
    expect(setMetricLine({ role: 'dps', value: 553.12 })).toBe('553.1 DPS');
    expect(setMetricLine({ role: 'healer', value: 742 })).toBe('742.0 HPS');
    expect(setMetricLine({ role: 'tank', value: 380.39 })).toBe('380 damage taken per second');
  });
});

describe('roleOf', () => {
  it('defaults a band with no role to dps', () => {
    expect(roleOf({})).toBe('dps');
    expect(roleOf({ role: 'tank' })).toBe('tank');
  });
});
