// web/src/lib/report/group.ts
// Who "the friendlies" are. An open-world log sees every player in range: the sample's 12:43
// trash pull has 265 players in it and the logger's party is five. The combat log's unit flags
// say who was in the logger's group (MINE, PARTY, RAID; logs/engine/units/units.go), and
// report.json keeps each unit's last-seen flags, so the default scope is the group whenever
// outsiders outnumber it. An instance log, where a few players read as outsiders only because
// they left the group before the log ended, keeps every player.
//
// From engine 0.6.0 each fight carries its own `group`: who was flagged MINE, PARTY or RAID during
// that pull (logs/engine/fight/counts.go). fightGroupScope reads it and needs none of the
// guesswork above; groupScope is the fallback for a fight indexed before then.
import { killCounts } from './kill-counts';
import type { FightEntry } from './types';
import { SOURCE_EVERYONE } from './url';

/** The affiliation bits of a unit's flags. */
const FLAG_MINE = 0x1;
const FLAG_PARTY = 0x2;
const FLAG_RAID = 0x4;
const GROUP_FLAGS = FLAG_MINE | FLAG_PARTY | FLAG_RAID;
/** A Classic party: a group of more than this is a raid. */
export const PARTY_MAX = 5;

/** What the default scope's players are called: a raid, a group, or every player the log saw. */
export type ScopeNoun = 'raid' | 'group' | 'everyone';

export interface GroupScope {
  /** The players in the logger's group. */
  group: ReadonlySet<string>;
  /** Outsiders outnumber the group: an open-world log. */
  crowded: boolean;
  /** The default scope is the group: a crowded log, on a pull the group was in. */
  scoped: boolean;
  /** What the default scope's players are called: everyone on a crowded pull without the group. */
  noun: ScopeNoun;
}

interface UnitFlags {
  guid: string;
  kind: string;
  flags: number;
}

/**
 * The group and whether the default scope narrows to it. `groupList` stands in for the flags
 * when the units have not arrived (the sample's meta carries it). `fightPlayers`, the pull's own
 * players: a crowded pull none of the group was in keeps every player, named as everyone.
 */
export function groupScope(
  units: readonly UnitFlags[],
  groupList?: readonly string[],
  fightPlayers?: readonly string[],
): GroupScope {
  const players = units.filter((unit) => unit.kind === 'player');
  const listed = groupList === undefined ? null : new Set(groupList);
  const inGroup = (unit: UnitFlags): boolean =>
    listed === null ? (unit.flags & GROUP_FLAGS) !== 0 : listed.has(unit.guid);
  const group = new Set(players.filter(inGroup).map((unit) => unit.guid));
  const crowded = group.size > 0 && players.length - group.size > group.size;
  const groupInFight = fightPlayers === undefined || fightPlayers.some((guid) => group.has(guid));
  const scoped = crowded && groupInFight;
  const flaggedRaid = players.some((unit) => inGroup(unit) && (unit.flags & FLAG_RAID) !== 0);
  const size = scoped ? group.size : players.length;
  const noun: ScopeNoun =
    crowded && !scoped ? 'everyone' : flaggedRaid || size > PARTY_MAX ? 'raid' : 'group';
  return { group, crowded, scoped, noun };
}

/** The noun the chart and the tables use under `source`: Everyone in range is everyone. */
export function scopeNoun(scope: GroupScope, source: string): ScopeNoun {
  return source === SOURCE_EVERYONE ? 'everyone' : scope.noun;
}

type FightGroupFacts = Pick<FightEntry, 'players' | 'group' | 'mob_deaths' | 'group_kills'>;

/**
 * The scope for one pull, or the night's pulls, from their own groups. Null when there are no
 * pulls or any of them predates engine 0.6.0: the caller falls back to groupScope. `raidFlagged`
 * is the players the report's units flag RAID, so a raid of five or fewer is still a raid.
 */
export function fightGroupScope(
  fights: readonly FightGroupFacts[],
  raidFlagged: ReadonlySet<string> = new Set(),
): GroupScope | null {
  if (fights.length === 0 || fights.some((fight) => killCounts(fight) === null)) return null;
  const group = new Set(fights.flatMap((fight) => fight.group ?? []));
  const players = new Set(fights.flatMap((fight) => fight.players));
  const outsiders = [...players].some((guid) => !group.has(guid));
  const isRaid = group.size > PARTY_MAX || [...group].some((guid) => raidFlagged.has(guid));
  if (group.size === 0) return { group, crowded: players.size > 0, scoped: false, noun: 'everyone' };
  return { group, crowded: outsiders, scoped: outsiders, noun: isRaid ? 'raid' : 'group' };
}

/** The players a report's units flag RAID, for fightGroupScope's noun. */
export function raidFlaggedPlayers(units: readonly UnitFlags[]): Set<string> {
  return new Set(
    units.filter((unit) => unit.kind === 'player' && (unit.flags & FLAG_RAID) !== 0).map((unit) => unit.guid),
  );
}

type DeathFacts = Pick<FightEntry, 'deaths' | 'group_deaths' | 'group' | 'mob_deaths' | 'group_kills'>;

export interface DeathMark {
  count: number;
  title: string;
}

const deathWord = (count: number): string => `${count} ${count === 1 ? 'death' : 'deaths'}`;

/**
 * A fight row's death mark, or a folded row's over its pulls: the group's deaths, like the row's
 * "your group killed" line, with everyone in range's in the title. A pull the group was not in
 * marks everyone in range's. A pull from before engine 0.6.0 has only the in-range count, worded
 * "among everyone in range" when the report's default scope is the group (`deathsInRange`).
 */
export function fightDeathMark(fights: readonly DeathFacts[], deathsInRange = false): DeathMark | null {
  const inRange = fights.reduce((sum, fight) => sum + fight.deaths, 0);
  const counted = fights.every((fight) => killCounts(fight) !== null && fight.group_deaths !== undefined);
  const withGroup = counted && fights.some((fight) => (fight.group ?? []).length > 0);
  if (withGroup) {
    const ours = fights.reduce((sum, fight) => sum + (fight.group_deaths ?? 0), 0);
    return ours === 0
      ? null
      : { count: ours, title: `${deathWord(ours)} in your group · ${inRange} among everyone in range` };
  }
  if (inRange === 0) return null;
  const amongEveryone = counted || deathsInRange;
  return { count: inRange, title: `${deathWord(inRange)}${amongEveryone ? ' among everyone in range' : ''}` };
}
