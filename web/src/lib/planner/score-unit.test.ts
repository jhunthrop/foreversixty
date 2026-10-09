import { describe, expect, it } from 'vitest';
import type { SimResult } from '../sim/types';
import {
  liveFigureOf,
  needsRoleMetrics,
  plannerScoreCopy,
  scoreUnitForBuild,
  scoreUnitForSpec,
} from './score-unit';

describe('scoreUnitForSpec', () => {
  it.each([
    ['priest-holy', 'hps'],
    ['druid-restoration', 'hps'],
    ['warrior-protection', 'tank_score'],
    ['mage-fire', 'dps'],
    ['not-a-spec', 'dps'],
  ] as const)('%s reads %s', (spec, unit) => {
    expect(scoreUnitForSpec(spec)).toBe(unit);
  });
});

describe('scoreUnitForBuild', () => {
  it('follows the tree with the most points', () => {
    expect(scoreUnitForBuild('priest', [0, 11, 0])).toBe('hps');
    expect(scoreUnitForBuild('priest', [0, 0, 11])).toBe('dps');
    expect(scoreUnitForBuild('warrior', [0, 0, 11])).toBe('tank_score');
  });
});

describe('plannerScoreCopy', () => {
  it('names the unit per role', () => {
    expect(plannerScoreCopy.label('dps')).toBe('DPS');
    expect(plannerScoreCopy.label('hps')).toBe('HPS');
    expect(plannerScoreCopy.label('tank_score')).toBe('Tank score');
    expect(plannerScoreCopy.ratePhrase('hps')).toBe('healing per second');
    expect(plannerScoreCopy.ratePhrase('tank_score')).toBe('tank score');
    expect(plannerScoreCopy.ratePhrase('dps')).toBe('damage per second');
  });

  it('writes the band figure in the band unit', () => {
    expect(plannerScoreCopy.bandFigure(40.9088, 'hps')).toBe('40.9 HPS');
    expect(plannerScoreCopy.bandFigure(40.9088, 'dps')).toBe('40.9 DPS');
  });

  it('promises the same live figure for every role', () => {
    expect(plannerScoreCopy.waiting()).toBe('live when the build reaches 51 points');
  });

  it("prints the figure in the role's own digits", () => {
    expect(plannerScoreCopy.figure(12345.6, 'dps')).toBe('12,346');
    expect(plannerScoreCopy.figure(381.04, 'hps')).toBe('381.0');
    expect(plannerScoreCopy.figure(27517.4, 'tank_score')).toBe('27,517');
  });

  it('names the failure by unit', () => {
    expect(plannerScoreCopy.failed('dps')).toBe('DPS estimate unavailable for this build.');
    expect(plannerScoreCopy.failed('tank_score')).toBe('Tank score estimate unavailable for this build.');
  });

  it('keeps the ask button text for damage and names the unit otherwise', () => {
    expect(plannerScoreCopy.show('dps')).toBe('Show DPS');
    expect(plannerScoreCopy.show('hps')).toBe('Show HPS');
  });
});

describe('the live figure', () => {
  const estimate = (mean: number) => ({ mean, stddev: 0, error: 1, min: 0, max: 0 });
  const result = {
    dps: estimate(900),
    healing: { effective_hps: estimate(380), hps: estimate(450), mana_lasts_sec: 1, hpm: 1 },
    tank: { score: estimate(27_000) },
  } as unknown as SimResult;

  it("is the role's own estimate", () => {
    expect(liveFigureOf(result, 'dps')?.mean).toBe(900);
    expect(liveFigureOf(result, 'hps')?.mean).toBe(380);
    expect(liveFigureOf(result, 'tank_score')?.mean).toBe(27_000);
  });

  it("is missing, never the damage number, when the role's block is absent", () => {
    const plain = { dps: estimate(900) } as unknown as SimResult;
    expect(liveFigureOf(plain, 'hps')).toBeNull();
    expect(liveFigureOf(plain, 'tank_score')).toBeNull();
  });

  it('asks only a healer and a tank for role metrics', () => {
    expect(needsRoleMetrics('dps')).toBe(false);
    expect(needsRoleMetrics('hps')).toBe(true);
    expect(needsRoleMetrics('tank_score')).toBe(true);
  });
});
