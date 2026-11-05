import { readFileSync } from 'node:fs';
import { describe, expect, it } from 'vitest';
import {
  NO_PUBLIC_LOGS,
  bandCountsLine,
  downEncounterNames,
  bossesDownCountLine,
  bossesDownLine,
  progressionNightDay,
  raidNightDay,
  rosterCardFacts,
  showBandCounts,
  showsWeekOne,
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

  it('words a report instant and a progression date alike, as the night’s date and never today', () => {
    expect(raidNightDay(new Date(2026, 9, 20, 19, 0).toISOString(), NOW)).toBe('Tue 20 Oct');
    expect(raidNightDay(new Date(2026, 9, 19, 19, 0).toISOString(), NOW)).toBe('Mon 19 Oct');
    expect(progressionNightDay('2026-10-20', NOW)).toBe('Tue 20 Oct');
    expect(progressionNightDay('2026-10-19', NOW)).toBe('Mon 19 Oct');
  });
});

describe('the cards', () => {
  it('names the bosses down against the named count, never every encounter', () => {
    expect(bossesDownLine(['Onyxia'], 1)).toBe('1 of 1 named bosses down: Onyxia.');
    expect(bossesDownLine([], 1)).toBe('No boss down this tier yet.');
    expect(bossesDownCountLine(0, 1)).toBe('No boss down this tier yet.');
    expect(bossesDownCountLine(1, 1)).toBe('1 of 1 named bosses down.');
  });

  it('shows the Before raids open card to members while nothing is logged in the tier', () => {
    expect(showsWeekOne('member', { pulls_total: 0 })).toBe(true);
    expect(showsWeekOne('officer', undefined)).toBe(true);
    expect(showsWeekOne('member', { pulls_total: 3 })).toBe(false);
    expect(showsWeekOne('public', { pulls_total: 0 })).toBe(false);
  });

  it('names only the encounters killed: one pulled and never downed reads as none down', () => {
    expect(downEncounterNames([{ name: 'Onyxia', kills: 0 }])).toEqual([]);
    expect(bossesDownLine(downEncounterNames([{ name: 'Onyxia', kills: 0 }]), 1)).toBe(
      'No boss down this tier yet.',
    );
    expect(downEncounterNames([{ name: 'Onyxia', kills: 2 }])).toEqual(['Onyxia']);
  });

  it('counts raiders with the plural helper and leaves zero clauses out', () => {
    expect(rosterCardFacts({ raiders: 1, waiting_for_approval: 4, below_rating_floor: 0 })).toEqual([
      { text: '1 raider with gear on the site' },
      { text: '4 waiting for approval', class: 'text-gold' },
    ]);
    expect(rosterCardFacts({ raiders: 5, waiting_for_approval: 0, below_rating_floor: 0 })).toEqual([
      { text: '5 raiders with gear on the site' },
    ]);
    expect(
      rosterCardFacts({ raiders: 5, waiting_for_approval: 0, below_rating_floor: 2 }).map((f) => f.text),
    ).toEqual(['5 raiders with gear on the site', '2 below the rating floor']);
  });
});

describe('NO_PUBLIC_LOGS', () => {
  it("is the guild-page spec's Progression empty line, never weekly (guild loop round 3)", () => {
    const spec = readFileSync(
      new URL('../../../../design/specs/2026-10-04-guild-page.md', import.meta.url),
      'utf8',
    );
    expect(NO_PUBLIC_LOGS).toBe('No public logs yet.');
    expect(spec).toContain(`\`${NO_PUBLIC_LOGS}\``);
    expect(NO_PUBLIC_LOGS).not.toContain('week');
  });
});
