// web/src/lib/guild/raid-roster.ts
// The Sign-ups roster's groups (design/specs/2026-10-26-guild-raids.md 3.7): In rows under their
// role, then Late, Bench, Out and Not signed up. Within a group: role, then class in the game's
// order, then name.
import type { GuildRaidRole, GuildRaidView, GuildSignupState } from './api';
import { roleForSpec } from './raid-roles';

const ROLE_ORDER: readonly GuildRaidRole[] = ['tank', 'healer', 'dps'];
const CLASS_ORDER: readonly string[] = [
  'warrior',
  'paladin',
  'hunter',
  'rogue',
  'priest',
  'shaman',
  'mage',
  'warlock',
  'druid',
];

export interface RosterRow {
  character_key: string;
  name: string;
  class: string;
  spec: string;
  role: GuildRaidRole;
  /** `null` for a member who has not signed up. */
  state: GuildSignupState | null;
  note: string;
  /** The viewer's own character (the row's wash). */
  own: boolean;
}

export interface RosterGroups {
  roles: Readonly<Record<GuildRaidRole, RosterRow[]>>;
  late: RosterRow[];
  bench: RosterRow[];
  out: RosterRow[];
  notSigned: RosterRow[];
}

function rank(order: readonly string[], value: string): number {
  const index = order.indexOf(value.toLowerCase());
  return index < 0 ? order.length : index;
}

function rosterOrder(a: RosterRow, b: RosterRow): number {
  return (
    rank(ROLE_ORDER, a.role) - rank(ROLE_ORDER, b.role) ||
    rank(CLASS_ORDER, a.class) - rank(CLASS_ORDER, b.class) ||
    a.name.localeCompare(b.name)
  );
}

function inState(rows: readonly RosterRow[], state: GuildSignupState): RosterRow[] {
  return rows.filter((row) => row.state === state).sort(rosterOrder);
}

export function rosterGroups(raid: GuildRaidView): RosterGroups {
  const signed: RosterRow[] = raid.signups.map((s) => ({ ...s }));
  const ownKeys = new Set(raid.viewer.characters.map((c) => c.character_key));
  const notSigned = raid.not_signed
    .map((m): RosterRow => ({
      ...m,
      role: roleForSpec(m.spec) ?? 'dps',
      state: null,
      note: '',
      own: ownKeys.has(m.character_key),
    }))
    .sort(rosterOrder);
  const signedIn = inState(signed, 'in');
  return {
    roles: {
      tank: signedIn.filter((row) => row.role === 'tank'),
      healer: signedIn.filter((row) => row.role === 'healer'),
      dps: signedIn.filter((row) => row.role === 'dps'),
    },
    late: inState(signed, 'late'),
    bench: inState(signed, 'bench'),
    out: inState(signed, 'out'),
    notSigned,
  };
}
