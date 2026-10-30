// web/src/lib/report/group.test.ts
import { describe, expect, it } from 'vitest';
import { groupScope, scopeNoun } from './group';
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
