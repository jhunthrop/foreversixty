import { describe, expect, it } from 'vitest';
import { bossCountLabel, distinctBossNames } from './boss-count';
import { accountReportSummary, recentReportSummary } from './logs/report-summary';

const fight = (name: string | null, encounter_id: number | null, kill: boolean) => ({
  name,
  encounter_id,
  kill,
});

describe('distinctBossNames', () => {
  it('counts two Onyxia kills in one night as one boss', () => {
    const night = [
      fight('Onyxian Whelp', null, true),
      fight('Onyxia', 1084, true),
      fight('Onyxian Whelp', null, true),
      fight('Onyxia', 1084, true),
    ];
    expect(distinctBossNames(night)).toEqual(['Onyxia']);
    expect(bossCountLabel(distinctBossNames(night).length)).toBe('1 boss');
  });

  it('skips a wipe, a trash pack and (with the tier list) an encounter the tier does not name', () => {
    const night = [fight('Onyxia', 1084, false), fight('Hakkar', 1085, true), fight('Whelp', null, true)];
    expect(distinctBossNames(night)).toEqual(['Hakkar']);
    expect(distinctBossNames(night, new Set([1084]))).toEqual([]);
  });

  it('keeps the order the bosses fell in', () => {
    const night = [fight('B', 2, true), fight('A', 1, true), fight('B', 2, true)];
    expect(distinctBossNames(night)).toEqual(['B', 'A']);
  });
});

describe('bossCountLabel', () => {
  it('pluralises', () => {
    expect(bossCountLabel(0)).toBe('0 bosses');
    expect(bossCountLabel(1)).toBe('1 boss');
    expect(bossCountLabel(2)).toBe('2 bosses');
  });
});

// The account page and the guild pages say the same number for the same night (round 8, P8-3).
describe('the account page counts distinct bosses', () => {
  it('reads "1 boss · 4 fights" for a night of two whelp packs and two Onyxia pulls, one a kill', () => {
    expect(accountReportSummary({ zone: "Onyxia's Lair", fightCount: 4, killCount: 2, bossCount: 1 })).toBe(
      "Onyxia's Lair · 1 boss · 4 fights",
    );
  });

  it('falls back to the kill count from an API that sends no boss count', () => {
    expect(accountReportSummary({ zone: "Onyxia's Lair", fightCount: 4, killCount: 2 })).toBe(
      "Onyxia's Lair · 2 bosses · 4 fights",
    );
  });

  it('says no bosses at all when none died', () => {
    expect(accountReportSummary({ zone: 'Darnassus', fightCount: 0, killCount: 0, bossCount: 0 })).toBe(
      'Darnassus · 0 fights',
    );
  });

  it('is the public lists count too', () => {
    expect(recentReportSummary({ zone: 'Zone', killCount: 3, bossCount: 1, fightCount: 6 })).toBe(
      'Zone · 1 boss · 6 fights',
    );
  });
});
