// web/src/lib/report/fights.test.ts
// parseReportState accepts any run of digits for ?fight=, because a total parser is what
// makes a pasted link render something rather than throw. It cannot know which fights a
// report actually has -- report.json has not loaded when the url is read -- so the island
// resolves the parsed number against the real list, and this covers the three ways that
// number can miss: 0 (fight_index is 1-based, so no report ever has one), a number past
// the end, and a gap.
import { describe, expect, it } from 'vitest';
import {
  bossPullsLabel,
  combinedFights,
  combinedLabel,
  defaultFightIndex,
  fightRows,
  fightTitle,
  pullNumbers,
  resolveFightIndex,
  trashGroupSelection,
  trashGroups,
} from './fights';
import type { FightEntry } from './types';
import { ALL_FIGHTS, trashScope } from './url';

function fight(index: number): FightEntry {
  return {
    index,
    kind: 'encounter',
    name: `Fight ${index}`,
    kill: true,
    in_progress: false,
    start: '2026-09-26T20:10:00Z',
    end: '2026-09-26T20:11:00Z',
    duration_ms: 60_000,
    players: [],
    deaths: 0,
    mob_deaths: 1,
    group_kills: 1,
  };
}

describe('resolveFightIndex', () => {
  const fights = [fight(1), fight(2), fight(4)];

  it('keeps a fight the report actually has', () => {
    expect(resolveFightIndex(fights, 2, 1)).toBe(2);
    expect(resolveFightIndex(fights, 4, 1)).toBe(4);
  });

  it('reads ?fight=0 as the whole night, which needs a boss pull to mean anything', () => {
    expect(resolveFightIndex(fights, 0, 1)).toBe(ALL_FIGHTS);
    expect(resolveFightIndex([{ ...fight(1), kind: 'trash' }], 0, 1)).toBe(1);
  });

  it('falls back for a fight past the end and for a gap in the numbering', () => {
    expect(resolveFightIndex(fights, 999, 1)).toBe(1);
    expect(resolveFightIndex(fights, 3, 1)).toBe(1);
  });

  it('falls back to the first fight the report has, not to 1', () => {
    expect(resolveFightIndex([fight(7), fight(8)], 3, 7)).toBe(7);
  });

  it('returns the fallback when the report has no fights at all', () => {
    expect(resolveFightIndex([], 3, 1)).toBe(1);
  });
});

describe('pullNumbers', () => {
  it('numbers each boss pull in order and leaves trash unnumbered', () => {
    const fights: FightEntry[] = [
      { ...fight(1), name: 'General Kaal', kill: false },
      { ...fight(2), kind: 'trash', name: 'Trash' },
      { ...fight(3), name: 'General Kaal', kill: true },
      { ...fight(4), name: 'Kryxis the Voracious' },
    ];
    const numbers = pullNumbers(fights);
    expect(numbers.get(1)).toEqual({ pull: 1, of: 2 });
    expect(numbers.get(2)).toBeUndefined();
    expect(numbers.get(3)).toEqual({ pull: 2, of: 2 });
    expect(numbers.get(4)).toEqual({ pull: 1, of: 1 });
  });
});

describe('defaultFightIndex', () => {
  it('opens on the first boss pull, or the first fight when there is no boss', () => {
    const trash = { ...fight(1), kind: 'trash' };
    expect(defaultFightIndex([trash, fight(2), fight(3)])).toBe(2);
    expect(defaultFightIndex([trash, { ...fight(5), kind: 'trash' }])).toBe(1);
    expect(defaultFightIndex([])).toBe(1);
  });

  it('accepts the whole night only when the report has a boss pull', () => {
    expect(resolveFightIndex([fight(1), fight(2)], ALL_FIGHTS, 1)).toBe(ALL_FIGHTS);
    expect(resolveFightIndex([{ ...fight(1), kind: 'trash' }], ALL_FIGHTS, 1)).toBe(1);
  });
});

describe('fightTitle and fightRows', () => {
  const trash = (index: number, zone: string | undefined, ms = 206_000, kills = 10): FightEntry => ({
    ...fight(index),
    kind: 'trash',
    name: 'Trash',
    zone,
    duration_ms: ms,
    mob_deaths: kills,
    group_kills: Math.floor(kills / 2),
  });

  it('titles trash by its length and keeps a boss name', () => {
    expect(fightTitle(trash(1, undefined))).toBe('Trash · 3:26');
    expect(fightTitle(fight(2))).toBe('Fight 2');
  });

  it('folds each zone’s trash into one entry after that zone’s boss pulls, summing kill counts and time', () => {
    const rows = fightRows([
      trash(1, 'Zone A', 60_000, 5),
      trash(2, 'Zone A', 60_000, 7),
      { ...fight(3), zone: 'Zone A' },
      trash(4, 'Zone A'),
    ]);
    expect(rows).toHaveLength(2);
    expect(rows[0]).toMatchObject({ type: 'fight', fight: { index: 3 } });
    expect(rows[1]).toMatchObject({
      type: 'trash-group',
      key: 1,
      zone: 'Zone A',
      kills: { mobDeaths: 22, groupKills: 10 },
      duration_ms: 326_000,
    });
    expect(rows[1].type === 'trash-group' && rows[1].fights.map((entry) => entry.index)).toEqual([1, 2, 4]);
  });

  it('gives a group with a pull from before engine 0.6 no kill count, not a partial one', () => {
    const old: FightEntry = { ...trash(2, 'Zone A'), mob_deaths: undefined, group_kills: undefined };
    const rows = fightRows([trash(1, 'Zone A'), old]);
    expect(rows[0]).toMatchObject({ type: 'trash-group', kills: null });
  });

  it('makes a lone trash pull its zone’s entry too, which opens the pull itself', () => {
    const [group] = trashGroups([fight(1), trash(2, 'Zone A')]);
    expect(group).toMatchObject({ key: 2, zone: 'Zone A' });
    expect(trashGroupSelection(group)).toBe(2);
    expect(trashGroupSelection(trashGroups([trash(2, 'A'), trash(3, 'A')])[0])).toBe(trashScope(2));
  });

  it('gives each zone its own entry, in the order the night reached them', () => {
    const rows = fightRows([trash(1, 'A'), trash(2, 'B'), { ...fight(3), zone: 'A' }, trash(4, 'A')]);
    expect(rows.map((row) => (row.type === 'fight' ? row.fight.index : `trash ${row.zone}`))).toEqual([
      3,
      'trash A',
      'trash B',
    ]);
  });

  it('leaves trash with no zone as its own rows', () => {
    const rows = fightRows([trash(1, undefined), trash(2, undefined)]);
    expect(rows.map((row) => row.type)).toEqual(['fight', 'fight']);
  });
});

describe('combined views', () => {
  const trash = (index: number, zone: string): FightEntry => ({
    ...fight(index),
    kind: 'trash',
    name: 'Trash',
    zone,
  });
  const night = [
    trash(1, 'Zephras Isle'),
    trash(2, 'Zephras Isle'),
    { ...fight(3), zone: 'Zephras Isle' },
    trash(5, 'Zephras Isle'),
  ];

  it('folds every boss pull for the night and one zone’s trash for its scope', () => {
    expect(combinedFights(night, ALL_FIGHTS)?.map((entry) => entry.index)).toEqual([3]);
    expect(combinedFights(night, trashScope(1))?.map((entry) => entry.index)).toEqual([1, 2, 5]);
    expect(combinedFights(night, trashScope(2))).toBeNull();
    expect(combinedFights(night, 3)).toBeNull();
  });

  it('keeps a trash scope the report has and drops one it does not', () => {
    expect(resolveFightIndex(night, trashScope(1), 3)).toBe(trashScope(1));
    expect(resolveFightIndex(night, trashScope(5), 3)).toBe(3);
  });

  it('names the combined view', () => {
    expect(combinedLabel(night, ALL_FIGHTS)).toBe('All boss pulls');
    expect(combinedLabel(night, trashScope(1))).toBe('Trash · Zephras Isle');
  });

  it('counts boss pulls only, and says so on a trash-only night', () => {
    expect(bossPullsLabel(6)).toBe('6 boss pulls');
    expect(bossPullsLabel(1)).toBe('1 boss pull');
    expect(bossPullsLabel(0)).toBe('No boss pulls');
  });
});
