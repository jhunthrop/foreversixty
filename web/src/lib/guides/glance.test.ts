import { describe, expect, it } from 'vitest';
import {
  fillerOf,
  glanceRotation,
  glanceStatRows,
  isCheapRank,
  isRaceOnlyRow,
  level60Clause,
  markCheapRanks,
  nextTalentPoint,
  racesByFaction,
  specRowTag,
} from './glance';
import type { RotationItem } from './rotation-lists';
import type { RailStatRow } from './rail-stats';
import type { GuideRaceRow } from './race-table';

const item = (name: string, clause: string, extra: Partial<RotationItem> = {}): RotationItem => ({
  name,
  clause,
  ...extra,
});
const RACES = ['Human', 'Undead', 'Night Elf'];

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

describe('glanceRotation', () => {
  const priority = [
    item('Dark Sacrifice', 'Undead only, once you are under 70% mana'),
    item('Evocation', 'When dry'),
    item('Scorch', 'To build stacks'),
    item('Pyroblast', 'On Heating Up'),
    item('Fire Blast', 'Every time it is off cooldown'),
    item('Fireball', 'The filler once you have Arcane Blast, which the leveling build takes at 57', {
      spellId: 1,
    }),
  ];

  it('never gives a race-only row a top spot, and adds the filler below them', () => {
    const glance = glanceRotation(priority, RACES);
    expect(glance.rows.map((row) => row.name)).toEqual(['Evocation', 'Scorch', 'Pyroblast', 'Fire Blast']);
    expect(glance.filler?.name).toBe('Fireball');
  });

  it('has no filler line when the filler is already shown', () => {
    expect(glanceRotation(priority.slice(3), RACES).filler).toBeUndefined();
  });

  it('finds race-only rows and the filler', () => {
    expect(isRaceOnlyRow(priority[0]!, RACES)).toBe(true);
    expect(isRaceOnlyRow(priority[1]!, RACES)).toBe(false);
    expect(fillerOf(priority)?.name).toBe('Fireball');
  });
});

describe('glanceStatRows', () => {
  it('keeps every stat a player can gear for', () => {
    const rows = [
      { label: 'Hit', reading: 'value', value: 1 },
      { label: 'Attack speed', reading: 'not-on-items', value: 8 },
      { label: 'Spell penetration', reading: 'not-significant' },
    ] as RailStatRow[];
    expect(glanceStatRows(rows).map((row) => row.label)).toEqual(['Hit', 'Spell penetration']);
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

describe('specRowTag', () => {
  it('tags healers, tanks and the melee spec of a ranged class only', () => {
    expect(specRowTag('priest', 'priest-holy', 'healer')).toBe('healer');
    expect(specRowTag('warrior', 'warrior-protection', 'tank')).toBe('tank');
    expect(specRowTag('hunter', 'hunter-survival', 'dps')).toBe('melee');
    expect(specRowTag('hunter', 'hunter-marksmanship', 'dps')).toBeUndefined();
    expect(specRowTag('druid', 'druid-feral', 'dps')).toBeUndefined();
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
