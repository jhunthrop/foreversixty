// web/src/lib/report/group.test.ts
import { describe, expect, it } from 'vitest';
import { fightDeathMark, fightGroupScope, groupScope, scopeNoun } from './group';
import { SOURCE_ENEMIES, SOURCE_EVERYONE, SOURCE_FRIENDLIES } from './url';

const MINE = 0x511;
const PARTY = 0x512;
const RAID = 0x514;
const OUTSIDER = 0x518;

const player = (guid: string, flags: number) => ({ guid, kind: 'player', flags });
const crowd = (count: number, flags = OUTSIDER) =>
  Array.from({ length: count }, (_, index) => player(`Player-out-${index}`, flags));

describe('groupScope', () => {
  it('narrows an open-world log to the logger’s own party', () => {
    const units = [
      player('me', MINE),
      ...['a', 'b', 'c', 'd'].map((guid) => player(guid, PARTY)),
      ...crowd(260),
    ];
    const scope = groupScope(units);
    expect(scope.scoped).toBe(true);
    expect([...scope.group].sort()).toEqual(['a', 'b', 'c', 'd', 'me']);
    expect(scope.noun).toBe('group');
  });

  it('names a group of more than five, or one flagged raid, a raid', () => {
    const raid = [player('me', MINE), ...Array.from({ length: 11 }, (_, index) => player(`r${index}`, RAID))];
    expect(groupScope([...raid, ...crowd(100)]).noun).toBe('raid');
    expect(groupScope([player('me', MINE), player('r', RAID), ...crowd(50)]).noun).toBe('raid');
  });

  it('keeps every player in an instance log, where the few outsiders left the group', () => {
    // A raider who left before the log ended reads as an outsider; they were still in the raid.
    const units = [...Array.from({ length: 38 }, (_, index) => player(`r${index}`, RAID)), ...crowd(2)];
    const scope = groupScope(units);
    expect(scope.scoped).toBe(false);
    expect(scope.noun).toBe('raid');
  });

  it('keeps every player when the log flags no group at all', () => {
    expect(groupScope(crowd(40)).scoped).toBe(false);
    expect(groupScope(crowd(40)).noun).toBe('raid');
    expect(groupScope(crowd(3)).noun).toBe('group');
  });

  it('reads a meta’s own group list when the units have not arrived', () => {
    const scope = groupScope(
      [...['me', 'a', 'b'], ...Array.from({ length: 20 }, (_, index) => `x${index}`)].map((guid) => ({
        guid,
        kind: 'player',
        flags: 0,
      })),
      ['me', 'a', 'b'],
    );
    expect(scope.scoped).toBe(true);
    expect(scope.group.size).toBe(3);
  });

  it('keeps every player, named as everyone in range, on a pull the group was not in', () => {
    const units = [player('me', MINE), player('a', PARTY), ...crowd(40)];
    const scope = groupScope(units, undefined, ['Player-out-1', 'Player-out-2']);
    expect(scope.crowded).toBe(true);
    expect(scope.scoped).toBe(false);
    expect(scope.noun).toBe('everyone');
    expect(groupScope(units, undefined, ['a', 'Player-out-1']).scoped).toBe(true);
  });

  it('ignores units that are not players', () => {
    const units = [player('me', MINE), { guid: 'pet', kind: 'pet', flags: MINE }, ...crowd(5)];
    expect([...groupScope(units).group]).toEqual(['me']);
  });
});

describe('fightGroupScope', () => {
  const counted = { mob_deaths: 10, group_kills: 4 };
  const pull = (players: string[], group: string[] | undefined, counts: object = counted) => ({
    players,
    group,
    ...counts,
  });
  const outsiders = (count: number) => Array.from({ length: count }, (_, index) => `x${index}`);

  it('scopes a pull to the group that was in it, whoever left the party before the log ended', () => {
    // The report-wide last-seen flags would call 'c' an outsider; this pull's own flags do not.
    const scope = fightGroupScope([pull(['me', 'a', 'c', ...outsiders(200)], ['me', 'a', 'c'])]);
    expect(scope?.scoped).toBe(true);
    expect([...(scope?.group ?? [])].sort()).toEqual(['a', 'c', 'me']);
    expect(scope?.noun).toBe('group');
  });

  it('scopes an instance pull too, once anyone outside the group is in it', () => {
    const raid = Array.from({ length: 38 }, (_, index) => `r${index}`);
    const scope = fightGroupScope([pull([...raid, 'passer-by'], raid)]);
    expect(scope?.scoped).toBe(true);
    expect(scope?.noun).toBe('raid');
  });

  it('keeps every player, named for the group, when nobody outside it was there', () => {
    const scope = fightGroupScope([pull(['me', 'a'], ['me', 'a'])]);
    expect(scope?.scoped).toBe(false);
    expect(scope?.crowded).toBe(false);
    expect(scope?.noun).toBe('group');
  });

  it('names a pull the group was not in everyone in range', () => {
    const scope = fightGroupScope([pull(outsiders(12), undefined)]);
    expect(scope?.scoped).toBe(false);
    expect(scope?.crowded).toBe(true);
    expect(scope?.noun).toBe('everyone');
  });

  it('calls a small group flagged raid in the report a raid', () => {
    const scope = fightGroupScope([pull(['me', 'a', ...outsiders(9)], ['me', 'a'])], new Set(['a']));
    expect(scope?.noun).toBe('raid');
  });

  it('unions the night’s pulls', () => {
    const scope = fightGroupScope([
      pull(['me', 'a', ...outsiders(5)], ['me', 'a']),
      pull(['me', 'b', ...outsiders(5)], ['me', 'b']),
    ]);
    expect([...(scope?.group ?? [])].sort()).toEqual(['a', 'b', 'me']);
  });

  it('has no answer for a pull from before engine 0.6, or for no pulls at all', () => {
    expect(fightGroupScope([pull(['me', ...outsiders(9)], undefined, {})])).toBeNull();
    expect(fightGroupScope([pull(['me'], ['me']), pull(['me'], undefined, {})])).toBeNull();
    expect(fightGroupScope([])).toBeNull();
  });
});

describe('fightDeathMark', () => {
  it('marks the group’s deaths, and names everyone in range’s in the title', () => {
    expect(fightDeathMark([{ deaths: 3, group_deaths: 1, group: ['me'], ...counted() }])).toEqual({
      count: 1,
      title: '1 death in your group · 3 among everyone in range',
    });
  });

  it('marks nothing when the group lost no one, even if others in range died', () => {
    expect(fightDeathMark([{ deaths: 2, group_deaths: 0, group: ['me'], ...counted() }])).toBeNull();
  });

  it('marks everyone in range’s deaths on a pull the group was not in', () => {
    expect(fightDeathMark([{ deaths: 2, group_deaths: 0, ...counted() }])).toEqual({
      count: 2,
      title: '2 deaths among everyone in range',
    });
  });

  it('keeps an old fight’s in-range count, worded as the report’s scope says', () => {
    expect(fightDeathMark([{ deaths: 2 }], true)).toEqual({
      count: 2,
      title: '2 deaths among everyone in range',
    });
    expect(fightDeathMark([{ deaths: 1 }], false)).toEqual({ count: 1, title: '1 death' });
    expect(fightDeathMark([{ deaths: 0 }])).toBeNull();
  });

  it('sums a folded row’s pulls, and falls back to in-range when any pull is old', () => {
    const a = { deaths: 2, group_deaths: 1, group: ['me'], ...counted() };
    const b = { deaths: 1, group_deaths: 1, group: ['me'], ...counted() };
    expect(fightDeathMark([a, b])).toEqual({
      count: 2,
      title: '2 deaths in your group · 3 among everyone in range',
    });
    expect(fightDeathMark([a, { deaths: 4 }], true)?.count).toBe(6);
  });

  function counted() {
    return { mob_deaths: 4, group_kills: 1 };
  }
});

describe('scopeNoun', () => {
  const scoped = groupScope([player('me', MINE), ...crowd(30)]);
  it('is the group under the default scope, everyone under Everyone in range', () => {
    expect(scopeNoun(scoped, SOURCE_FRIENDLIES)).toBe('group');
    expect(scopeNoun(scoped, SOURCE_EVERYONE)).toBe('everyone');
  });
  it('is the default scope’s noun for any other source', () => {
    expect(scopeNoun(scoped, SOURCE_ENEMIES)).toBe('group');
    expect(scopeNoun(scoped, 'Player-1')).toBe('group');
  });
});
