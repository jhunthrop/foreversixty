// web/src/lib/report/group.ts
// Who "the friendlies" are. An open-world log sees every player in range: the sample's 12:43
// trash pull has 265 players in it and the logger's party is five. The combat log's unit flags
// say who was in the logger's group (MINE, PARTY, RAID; logs/engine/units/units.go), and
// report.json keeps each unit's last-seen flags, so the default scope is the group whenever
// outsiders outnumber it. An instance log, where a few players read as outsiders only because
// they left the group before the log ended, keeps every player.
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
