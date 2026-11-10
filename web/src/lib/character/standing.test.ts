// The character page's standing line and activity (character spec 2026-11-06 C16, C18, C19, C21).
import { describe, expect, it } from 'vitest';
import type { NextRaidView } from '../guild/api';
import type { GuildSims, SpecWeek } from '../home/home-answer';
import { activityPath, gradeRow, guildRankRow, ownActivitySegments, raidRow } from './standing';

describe('ownActivitySegments', () => {
  const name = { text: 'Elyra Duskvale', style: 'name' as const, class: 'priest' };

  it('drops the leading own name and capitalises the verb', () => {
    expect(
      ownActivitySegments([name, { text: ' uploaded Molten Core, 10 bosses in 92 min' }], 'Elyra Duskvale'),
    ).toEqual([{ text: 'Uploaded Molten Core, 10 bosses in 92 min' }]);
  });

  it('drops a name mid-sentence without capitalising', () => {
    const climb = [
      { text: 'Holy, 60: ', style: 'strong' as const },
      name,
      { text: ' climbed 6 places to ' },
      { text: '#4', style: 'strong' as const },
    ];
    expect(ownActivitySegments(climb, 'Elyra Duskvale')).toEqual([
      { text: 'Holy, 60: ', style: 'strong' },
      { text: 'climbed 6 places to ' },
      { text: '#4', style: 'strong' },
    ]);
  });

  it("keeps another character's name", () => {
    const other = [
      { text: 'Thalrin', style: 'name' as const, class: 'warrior' },
      { text: ' simmed 900 DPS' },
    ];
    expect(ownActivitySegments(other, 'Elyra Duskvale')).toEqual(other);
  });
});

const row = {
  rank: 3,
  name: 'Elyra',
  class: 'priest',
  key: 'k',
  value: 1,
  delta: null,
  simmed_at: '',
  you: true,
};
const sims: GuildSims = {
  guild_id: 2,
  guild_name: 'Olympus XXVII',
  role: 'dps',
  band: 60,
  preset: 'raid',
  unit: 'DPS',
  of: 14,
  rows: [row],
  you: null,
};

describe('the standing rows', () => {
  it('reads the place in the guild ladder', () => {
    expect(guildRankRow(sims)).toEqual({ label: 'Olympus XXVII · 60 raid', value: '#3 of 14 DPS' });
    expect(guildRankRow({ ...sims, role: 'healer', band: 30, preset: 'bare' })).toEqual({
      label: 'Olympus XXVII · 30 to 39',
      value: '#3 of 14 healers',
    });
    expect(guildRankRow({ ...sims, rows: [{ ...row, you: false }] })).toBeNull();
    expect(guildRankRow(null)).toBeNull();
  });

  it('names the grade, or the place for a list with no letter', () => {
    const week = { grade: 'A', place: 4, of: 23 } as SpecWeek;
    expect(gradeRow(week, 'Holy')).toEqual({ label: 'Holy tier', value: 'A' });
    expect(gradeRow({ ...week, grade: null }, 'Holy')).toEqual({ label: 'Holy tier', value: '#4 of 23' });
    expect(gradeRow(null, 'Holy')).toBeNull();
  });

  it('counts down to the next guild raid until it starts', () => {
    const raid = {
      title: "Onyxia's Lair",
      starts_at: '2026-11-08T20:30:00Z',
      counts: { in: 18 },
    } as NextRaidView;
    expect(raidRow(raid, new Date('2026-11-06T18:00:00Z'))).toEqual({
      label: "Onyxia's Lair",
      value: '2 days, 2 hours, 30 min',
    });
    expect(raidRow(raid, new Date('2026-11-09T00:00:00Z'))).toBeNull();
  });

  it('asks the feed for one character', () => {
    expect(activityPath('us/hardcore/elyra-duskvale')).toBe(
      '/v1/feed?character=us%2Fhardcore%2Felyra-duskvale&limit=5',
    );
  });
});
