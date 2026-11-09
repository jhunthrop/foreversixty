// The character page's standing line and activity (character spec 2026-11-06 C16, C18, C19, C21).
import { describe, expect, it } from 'vitest';
import type { NextRaidView } from '../guild/api';
import type { GuildSims, SpecWeek } from '../home/home-answer';
import { activityPath, gradeLine, guildRankLine, raidLine } from './standing';

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

describe('the standing line', () => {
  it('reads the place in the guild ladder', () => {
    expect(guildRankLine(sims)).toBe('#3 of 14 DPS in Olympus XXVII, 60 raid');
    expect(guildRankLine({ ...sims, role: 'healer', band: 30, preset: 'bare' })).toBe(
      '#3 of 14 healers in Olympus XXVII, 30 to 39',
    );
    expect(guildRankLine({ ...sims, rows: [{ ...row, you: false }] })).toBeNull();
    expect(guildRankLine(null)).toBeNull();
  });

  it('names the grade, or the place for a list with no letter', () => {
    const week = { grade: 'A', place: 4, of: 23 } as SpecWeek;
    expect(gradeLine(week, 'Holy')).toBe('Holy is graded A this week');
    expect(gradeLine({ ...week, grade: null }, 'Holy')).toBe('Holy is #4 of 23 this week');
    expect(gradeLine(null, 'Holy')).toBeNull();
  });

  it('counts down to the next guild raid until it starts', () => {
    const raid = {
      title: "Onyxia's Lair",
      starts_at: '2026-11-08T20:30:00Z',
      counts: { in: 18 },
    } as NextRaidView;
    expect(raidLine(raid, new Date('2026-11-06T18:00:00Z'))).toBe("Onyxia's Lair in 2 days, 2 hours, 18 in");
    expect(raidLine(raid, new Date('2026-11-09T00:00:00Z'))).toBeNull();
  });

  it('asks the feed for one character', () => {
    expect(activityPath('us/hardcore/elyra-duskvale')).toBe(
      '/v1/feed?character=us%2Fhardcore%2Felyra-duskvale&limit=5',
    );
  });
});
