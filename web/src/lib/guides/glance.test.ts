import { describe, expect, it } from 'vitest';
import {
  atLevel60,
  isCheapRank,
  level60Clause,
  markCheapRanks,
  nextTalentPoint,
  racesByFaction,
} from './glance';
import type { RotationItem } from './rotation-lists';
import type { GuideRaceRow } from './race-table';

const item = (name: string, clause: string, extra: Partial<RotationItem> = {}): RotationItem => ({
  name,
  clause,
  ...extra,
});

describe('level60Clause', () => {
  it('drops the rank swaps a leveller makes', () => {
    expect(
      level60Clause(
        'On every swing it is ready, with an older rank on your bar at 26 to 29, at 33, and at 46 and 47, and to 49 on most Horde characters',
      ),
    ).toBe('On every swing it is ready');
    expect(level60Clause('The filler once you have Arcane Blast, which the leveling build takes at 57')).toBe(
      'The filler once you have Arcane Blast',
    );
  });

  it('keeps a clause that names no level', () => {
    const clause = 'To build Improved Scorch to 5 stacks, or to refresh it with 5 seconds or less left';
    expect(level60Clause(clause)).toBe(clause);
  });
});

describe('nextTalentPoint', () => {
  it('is the next level the ladder places a point at', () => {
    const steps = [
      { level: 34, talentId: 1, name: 'Lacerate', rank: 1, talents: '' },
      { level: 35, talentId: null, name: null, rank: 0, talents: '' },
      { level: 36, talentId: 2, name: 'Expose Prey', rank: 1, talents: '' },
    ];
    expect(nextTalentPoint(steps, 34)).toEqual({ level: 36, name: 'Expose Prey' });
    expect(nextTalentPoint(steps, 36)).toBeUndefined();
  });
});

describe('racesByFaction', () => {
  it('groups Alliance then Horde with the recommended races first', () => {
    const row = (name: string, faction: string, recommended: boolean): GuideRaceRow => ({
      slug: name.toLowerCase(),
      name,
      faction,
      recommended,
      racials: [],
    });
    const groups = racesByFaction([
      row('Human', 'alliance', false),
      row('Orc', 'horde', false),
      row('Night Elf', 'alliance', true),
      row('Troll', 'horde', true),
    ]);
    expect(groups.map((g) => [g.faction, g.rows.map((r) => r.name)])).toEqual([
      ['alliance', ['Night Elf', 'Human']],
      ['horde', ['Troll', 'Orc']],
    ]);
  });
});

describe('cheap ranks', () => {
  const prose = 'Renew always up at rank 7 or higher. Downrank: Greater Heal rank 1 is the one to keep.';

  it('marks a lower rank only where the prose chose it', () => {
    expect(isCheapRank(item('Renew', '', { rank: 7 }), 10, prose)).toBe(true);
    expect(isCheapRank(item('Greater Heal', '', { rank: 4 }), 5, prose)).toBe(false);
    expect(isCheapRank(item('Renew', '', { rank: 10 }), 10, prose)).toBe(false);
  });

  it('marks every list', () => {
    const lists = { opener: [], priority: [item('Greater Heal', '', { rank: 1 })], cooldowns: [] };
    expect(markCheapRanks(lists, () => 5, prose).priority[0]?.cheapRank).toBe(true);
  });
});

describe('atLevel60', () => {
  it('drops the levelling rank swaps from every list, keeping each row one line', () => {
    const row = (clause: string) => item('Raptor Strike', clause);
    const lists = atLevel60({
      opener: [row('20 s before the pull')],
      priority: [
        row('On every swing it is ready, with an older rank on your bar at 26 to 29, at 31, and to 49'),
      ],
      cooldowns: [row('5 min cooldown')],
    });
    expect(lists.priority[0]?.clause).toBe('On every swing it is ready');
    expect(lists.opener[0]?.clause).toBe('20 s before the pull');
    expect(lists.cooldowns[0]?.clause).toBe('5 min cooldown');
  });
});
