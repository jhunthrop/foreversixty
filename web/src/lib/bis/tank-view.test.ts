// web/src/lib/bis/tank-view.test.ts
import { describe, expect, it } from 'vitest';
import { tankCopy } from './copy';
import {
  formatChanceOfDeath,
  gapLabelFor,
  headlineDpsOf,
  isTankBand,
  slotScoreUnitFor,
  tankHeadlineFor,
  tankHeadlineForBand,
  tankIndexSummaryOf,
} from './tank-view';
import type { TankMetrics } from './types';

const METRICS: TankMetrics = {
  dtps: 301.24,
  tmi: 1523.4,
  chance_of_death: 0.0123,
  tps: 442.7,
  effective_health: 11235.4,
  dps: 260,
};

describe('formatChanceOfDeath', () => {
  it.each([
    [0, '0.0%'],
    [0.0123, '1.2%'],
    [0.081, '8.1%'],
    [1, '100.0%'],
    [0.0004, '<0.1%'],
    [0.0005, '0.1%'],
  ])('%s -> %s', (chance, expected) => {
    expect(formatChanceOfDeath(chance)).toBe(expected);
  });
});

describe('tankHeadlineFor', () => {
  const headline = tankHeadlineFor(METRICS);

  it('lists the four figures in order: effective health, dtps, chance of death, tps', () => {
    expect(headline.figures.map((figure) => figure.id)).toEqual([
      'effective-health',
      'dtps',
      'chance-of-death',
      'tps',
    ]);
    expect(headline.figures.map((figure) => figure.label)).toEqual([
      tankCopy.effectiveHealthLabel,
      tankCopy.dtpsLabel,
      tankCopy.chanceOfDeathLabel,
      tankCopy.tpsLabel,
    ]);
  });

  it('formats effective health with thousands separators and rates to one decimal', () => {
    expect(headline.figures.map((figure) => figure.value)).toEqual(['11,235', '301.2', '1.2%', '442.7']);
  });

  it('keeps TMI and own damage on the quieter secondary line, with the TMI explanation as title', () => {
    expect(headline.secondaryLine).toBe('TMI 1,523 · own damage 260.0 DPS');
    expect(headline.secondaryTitle).toBe(tankCopy.tmiTitle);
  });
});

describe('the tank/DPS decision', () => {
  const tank = { role: 'tank', metrics: METRICS, set_dps: 260 } as const;
  const dps = { role: 'dps', metrics: null, set_dps: 300 } as const;

  it('reads a band as a tank only when its role says so', () => {
    expect(isTankBand(tank)).toBe(true);
    expect(isTankBand(dps)).toBe(false);
    expect(isTankBand({})).toBe(false);
  });

  it('labels a tank band slot figures as tank score, a DPS band as DPS', () => {
    expect(slotScoreUnitFor(tank)).toBe('tank_score');
    expect(slotScoreUnitFor(dps)).toBe('dps');
  });

  it('gives a tank headline and no DPS headline, and the reverse for a DPS band', () => {
    expect(tankHeadlineForBand(tank)?.figures).toHaveLength(4);
    expect(tankHeadlineForBand(dps)).toBeUndefined();
    expect(headlineDpsOf(tank)).toBeUndefined();
    expect(headlineDpsOf(dps)).toBe(300);
  });

  it('summarises a tank on the index by effective health, a DPS band not at all', () => {
    expect(tankIndexSummaryOf(tank)).toBe('Level 60: 11,235 effective health');
    expect(tankIndexSummaryOf(dps)).toBeUndefined();
  });

  it('words a runner-up gap in the band unit, never DPS on a tank band', () => {
    expect(gapLabelFor(-1.34, 'tank_score')).toBe('−1.3 score');
    expect(gapLabelFor(0.01, 'tank_score')).toBe('same score');
    expect(gapLabelFor(-1.34, 'dps')).toBe('−1.3 DPS');
  });
});
