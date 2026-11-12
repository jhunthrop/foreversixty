// Your upgrades' words in the band unit (lib/home/upgrades-view.ts, home units lane 2026-11-11).
import { describe, expect, it } from 'vitest';
import type { UpgradesResult } from './upgrades';
import { upgradeGain, upgradesFacts } from './upgrades-view';

describe('upgradeGain', () => {
  it('words the row in the band unit the paperdoll uses: DPS, HPS, a tank score', () => {
    expect(upgradeGain({ gainDps: 1182.64 }, 'tank_score')).toEqual({
      value: '+1,182.6',
      unit: 'score',
      tone: 'kill',
    });
    expect(upgradeGain({ gainDps: 12.34 }, 'hps')).toEqual({ value: '+12.3', unit: 'HPS', tone: 'kill' });
    expect(upgradeGain({ gainDps: 5.06 }, 'dps')).toEqual({ value: '+5.1', unit: 'DPS', tone: 'kill' });
  });

  it('has no figure for a slot with no measured gain', () => {
    expect(upgradeGain({ gainDps: null }, 'dps')).toBeUndefined();
  });
});

describe('upgradesFacts', () => {
  const result = (
    total: number,
    count: number,
    bis: number,
    scoreUnit: UpgradesResult['scoreUnit'] = 'dps',
  ): UpgradesResult =>
    ({
      upgrades: Array.from({ length: count }, () => ({})),
      alreadyBis: Array.from({ length: bis }, () => ({})),
      totalGainDps: total,
      scoreUnit,
      notSimCheckedCount: 0,
    }) as unknown as UpgradesResult;

  it('reads the whole gap and the slots already best in slot', () => {
    expect(upgradesFacts(result(52.08, 12, 5))).toEqual([
      { n: '+52.1', text: 'DPS across all 12' },
      { n: '5', text: 'best in slot' },
    ]);
    expect(upgradesFacts(result(0, 1, 16))).toEqual([{ n: '16', text: 'best in slot' }]);
    expect(upgradesFacts(result(7.8, 5, 0))).toEqual([{ n: '+7.8', text: 'DPS' }]);
    expect(upgradesFacts(result(7.8, 14, 0))).toEqual([{ n: '+7.8', text: 'DPS across all 14' }]);
    expect(upgradesFacts(result(0, 0, 17))).toEqual([]);
    expect(upgradesFacts(undefined)).toEqual([]);
  });

  it('counts the gap in the band unit: a tank score, a healer HPS', () => {
    expect(upgradesFacts(result(3895.7, 5, 1, 'tank_score'))[0]).toEqual({ n: '+3,895.7', text: 'score' });
    expect(upgradesFacts(result(1058.4, 17, 0, 'tank_score'))[0]).toEqual({
      n: '+1,058.4',
      text: 'score across all 17',
    });
    expect(upgradesFacts(result(41.2, 5, 1, 'hps'))[0]).toEqual({ n: '+41.2', text: 'HPS' });
  });
});
