import { describe, expect, it } from 'vitest';
import {
  bandCountsLine,
  checksFailingLine,
  encountersDownLine,
  progressionNightDay,
  raidNightDay,
  rosterCardParts,
  showBandCounts,
  updatedLine,
} from './overview-view';

const NOW = new Date(2026, 9, 20, 20, 0); // Tuesday evening, the viewer's own clock

describe('bandCountsLine', () => {
  it('says one boss down, not "1 bosses down"', () => {
    expect(bandCountsLine(1, 4)).toBe('1 boss down · 4 pulls');
    expect(bandCountsLine(2, 1)).toBe('2 bosses down · 1 pull');
    expect(bandCountsLine(0, 0)).toBe('0 bosses down · 0 pulls');
  });
});

describe('showBandCounts', () => {
  it('counts everything for a member and officer', () => {
    expect(showBandCounts('member', 0)).toBe(true);
    expect(showBandCounts('officer', 0)).toBe(true);
  });

  it('gives a visitor counts only when the public tabs list a report', () => {
    expect(showBandCounts('public', 0)).toBe(false);
    expect(showBandCounts('public', 1)).toBe(true);
  });
});

describe('updatedLine and the raid-night days', () => {
  it('reads the day in the viewer zone', () => {
    expect(updatedLine(new Date(2026, 9, 20, 9, 0).toISOString(), NOW)).toBe('Updated today');
    expect(updatedLine(new Date(2026, 9, 19, 9, 0).toISOString(), NOW)).toBe('Updated yesterday');
    expect(updatedLine('not a date', NOW)).toBe('');
  });

  it('words a report instant and a progression date alike', () => {
    expect(raidNightDay(new Date(2026, 9, 20, 19, 0).toISOString(), NOW)).toBe('today');
    expect(progressionNightDay('2026-10-20', NOW)).toBe('today');
    expect(progressionNightDay('2026-10-19', NOW)).toBe('yesterday');
  });
});

describe('the cards', () => {
  it('names the encounters down, singular and plural', () => {
    expect(encountersDownLine(['Onyxia'])).toBe('1 encounter down this tier: Onyxia.');
    expect(encountersDownLine(['Onyxia', 'Molten Core'])).toBe(
      '2 encounters down this tier: Onyxia, Molten Core.',
    );
    expect(encountersDownLine([])).toBe('No encounter down this tier yet.');
  });

  it('counts raiders and checks with the plural helper', () => {
    expect(rosterCardParts({ raiders: 1, waiting_for_approval: 4, below_rating_floor: 0 })).toEqual({
      raiders: '1 raider',
      waiting: '4 waiting for approval',
      belowFloor: '0 below the rating floor',
    });
    expect(checksFailingLine(1)).toBe('1 check failing');
    expect(checksFailingLine(3)).toBe('3 checks failing');
  });
});
