// web/src/lib/guild/raid-roster.ts
// The Sign-ups roster's groups (raids spec 3.7, raid comp spec 3.3): In rows under Tanks, Healers,
// Melee and Ranged by the API's `group`, then Late, Bench, Out and Not signed up. Within an In
// group: class in the game's order, then name; Late, Bench and Out lead with the role.
import type {
  GuildRaidAlsoSignup,
  GuildRaidGroup,
  GuildRaidRole,
  GuildRaidView,
  GuildSignupState,
} from './api';
import { classThenName, rowGroup, specClassLine } from './raid-comp';
import { raidCopy } from './raid-copy';
import { roleForSpec } from './raid-roles';

const ROLE_ORDER: readonly GuildRaidRole[] = ['tank', 'healer', 'dps'];

export interface RosterRow {
  character_key: string;
  name: string;
  class: string;
  spec: string;
  role: GuildRaidRole;
  /** The API's roster group. */
  group: GuildRaidGroup;
  /** The row's catalogue spec key, `null` when it does not resolve. */
  spec_key: string | null;
  /** `null` for a member who has not signed up. */
  state: GuildSignupState | null;
  note: string;
  /** The viewer's own character (the row's wash). */
  own: boolean;
  /** The muted line naming the account's other answered characters, `null` when it has none. */
  also: string | null;
  /** The picker's Discord name for a character picked on Discord, not yet confirmed by Battle.net. */
  discord_picked_by?: string;
}

/** `Also: Rootbound (Restoration Druid, Bench)`: the account's other answered characters, best
 *  place first as the API lists them, joined by `; `; `null` when there are none (raids spec 3.7). */
export function alsoLine(also: readonly GuildRaidAlsoSignup[]): string | null {
  if (also.length === 0) return null;
  return raidCopy.also(
    also.map((a) =>
      raidCopy.alsoPart(a.name, specClassLine(a.class, a.spec_key), raidCopy.segments[a.state]),
    ),
  );
}

/** On the viewer's answered character that is not its account's roster row: `Officers see
 *  Quillfeather (Late).`, naming the row the officers read; else `null`. */
export function officersSeeLine(raid: Pick<GuildRaidView, 'signups'>, characterKey: string): string | null {
  const row = raid.signups.find((s) => s.own && (s.also ?? []).some((a) => a.character_key === characterKey));
  return row === undefined ? null : raidCopy.stripStatus.officersSee(row.name, raidCopy.segments[row.state]);
}

export interface RosterGroups {
  groups: Readonly<Record<GuildRaidGroup, RosterRow[]>>;
  /** Every In row, class then name: a guild event's In group, which has no roles. */
  in: RosterRow[];
  late: RosterRow[];
  bench: RosterRow[];
  out: RosterRow[];
  notSigned: RosterRow[];
}

function roleRank(role: GuildRaidRole | null): number {
  const index = role === null ? -1 : ROLE_ORDER.indexOf(role);
  return index < 0 ? ROLE_ORDER.length : index;
}

/** Role, then class in the game's order, then name: the order Late, Bench and Out list in. */
export function rosterOrder(
  a: Pick<RosterRow, 'class' | 'name'> & { role: GuildRaidRole | null },
  b: Pick<RosterRow, 'class' | 'name'> & { role: GuildRaidRole | null },
): number {
  return roleRank(a.role) - roleRank(b.role) || classThenName(a, b);
}

function inState(rows: readonly RosterRow[], state: GuildSignupState): RosterRow[] {
  return rows.filter((row) => row.state === state).sort(rosterOrder);
}

export function rosterGroups(raid: GuildRaidView): RosterGroups {
  const signed: RosterRow[] = raid.signups.map((s) => ({
    ...s,
    group: rowGroup(s),
    spec_key: s.spec_key ?? null,
    also: alsoLine(s.also ?? []),
  }));
  const ownKeys = new Set(raid.viewer.characters.map((c) => c.character_key));
  const notSigned = raid.not_signed
    .map((m): RosterRow => {
      const role = roleForSpec(m.spec) ?? 'dps';
      return {
        ...m,
        role,
        group: rowGroup({ ...m, role }),
        spec_key: m.spec_key ?? null,
        state: null,
        note: '',
        own: ownKeys.has(m.character_key),
        also: null,
      };
    })
    .sort(rosterOrder);
  const signedIn = signed.filter((row) => row.state === 'in').sort(classThenName);
  const of = (group: GuildRaidGroup): RosterRow[] => signedIn.filter((row) => row.group === group);
  return {
    groups: { tank: of('tank'), healer: of('healer'), melee: of('melee'), ranged: of('ranged') },
    in: signedIn,
    late: inState(signed, 'late'),
    bench: inState(signed, 'bench'),
    out: inState(signed, 'out'),
    notSigned,
  };
}
