import { describe, expect, it } from 'vitest';
import { killCounts, killCountsLabel, sumKillCounts } from './kill-counts';

describe('killCounts', () => {
  it('reads an engine 0.6 fight’s true figures', () => {
    expect(killCounts({ mob_deaths: 177, group_kills: 26 })).toEqual({ mobDeaths: 177, groupKills: 26 });
    expect(killCounts({ mob_deaths: 0, group_kills: 0 })).toEqual({ mobDeaths: 0, groupKills: 0 });
  });

  it('has no figure for a fight indexed before engine 0.6, whose npc_kills double counted', () => {
    expect(killCounts({})).toBeNull();
    expect(killCounts({ npc_kills: 203 } as never)).toBeNull();
    expect(killCounts({ mob_deaths: 4 })).toBeNull();
    expect(killCounts({ mob_deaths: Number.NaN, group_kills: 1 })).toBeNull();
  });
});

describe('sumKillCounts', () => {
  it('sums a run of counted pulls', () => {
    expect(
      sumKillCounts([
        { mob_deaths: 5, group_kills: 2 },
        { mob_deaths: 7, group_kills: 7 },
      ]),
    ).toEqual({ mobDeaths: 12, groupKills: 9 });
  });

  it('has no figure when any pull in the run has none, rather than a partial one', () => {
    expect(sumKillCounts([{ mob_deaths: 5, group_kills: 2 }, {}])).toBeNull();
    expect(sumKillCounts([])).toBeNull();
  });
});

describe('killCountsLabel', () => {
  it('says how many mobs died and how many of them the group killed', () => {
    expect(killCountsLabel({ mobDeaths: 177, groupKills: 26 })).toBe('177 mobs died · your group killed 26');
    expect(killCountsLabel({ mobDeaths: 1, groupKills: 1 })).toBe('1 mob died · your group killed 1');
    expect(killCountsLabel({ mobDeaths: 3, groupKills: 0 })).toBe('3 mobs died · your group killed 0');
    expect(killCountsLabel({ mobDeaths: 0, groupKills: 0 })).toBe('No mobs died');
  });
});
