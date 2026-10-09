import { describe, expect, it } from 'vitest';
import { isLiveSimmed, plannerScoreCopy, scoreUnitForBuild, scoreUnitForSpec } from './score-unit';

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

  it('promises a live figure only for damage', () => {
    expect(plannerScoreCopy.waiting('dps')).toBe('live when the build reaches 51 points');
    expect(plannerScoreCopy.waiting('hps')).not.toMatch(/reaches 51/);
    expect(plannerScoreCopy.waiting('hps')).toMatch(/hps/);
    expect(isLiveSimmed('dps')).toBe(true);
    expect(isLiveSimmed('hps')).toBe(false);
    expect(isLiveSimmed('tank_score')).toBe(false);
  });

  it('keeps the ask button text for damage and names the unit otherwise', () => {
    expect(plannerScoreCopy.show('dps')).toBe('Show DPS');
    expect(plannerScoreCopy.show('hps')).toBe('Show HPS');
  });
});
