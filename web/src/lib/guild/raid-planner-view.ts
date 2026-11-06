// web/src/lib/guild/raid-planner-view.ts
// The raid planner's pure models (web layout spec 2026-11-05 §7.2 "Raid planner", §3.4): the roster
// grouped by Status, Role or Class (the `group` query param is the truth), each sign-up row's words
// and state Tag, the bench an officer swaps from, the raiders with no sign-up, the comp target's
// edit draft and when the Discord ping may go out. Role keeps the shipped Raid-Helper groups
// (Tanks, Healers, Melee, Ranged, then Late, Bench, Out; raid comp spec 3.3).
import type { GuildRaidGroup, GuildRaidView, GuildSignupState } from './api';
import { CLASS_ORDER, RAID_GROUPS, classRank, classWord, specClassLine } from './raid-comp';
import { SEP, kindCountsRoles, kindOf, raidCopy } from './raid-copy';
import { roleForSpec } from './raid-roles';
import { rosterGroups, rosterOrder, type RosterRow } from './raid-roster';
import type { RaidChipStatus } from './planned-row';
import type { GuildCompRole } from './overview-api';
import { pingAllowed } from './overview-role-view';

export type PlannerGrouping = 'status' | 'role' | 'class';

/** The query parameter the control bar reads and writes. */
export const GROUP_PARAM = 'group';
const GROUPINGS: readonly PlannerGrouping[] = ['status', 'role', 'class'];
const GROUPING_WORDS: Readonly<Record<PlannerGrouping, string>> = {
  status: 'Status',
  role: 'Role',
  class: 'Class',
};

/** The groupings a night offers: a guild event counts no roles, so it has no Role grouping. */
export function groupingsOf(raid: Pick<GuildRaidView, 'zone_kind'>): PlannerGrouping[] {
  return kindCountsRoles(kindOf(raid)) ? [...GROUPINGS] : GROUPINGS.filter((g) => g !== 'role');
}

/** A raid opens on Role (the shipped Raid-Helper read); an event on Status. */
export function defaultGrouping(raid: Pick<GuildRaidView, 'zone_kind'>): PlannerGrouping {
  return kindCountsRoles(kindOf(raid)) ? 'role' : 'status';
}

/** The grouping the URL asks for, else the night's default. */
export function groupingFrom(search: string, raid: Pick<GuildRaidView, 'zone_kind'>): PlannerGrouping {
  const asked = new URLSearchParams(search).get(GROUP_PARAM);
  const offered = groupingsOf(raid);
  return offered.find((g) => g === asked) ?? defaultGrouping(raid);
}

/** The query string after a choice: the default drops the param, so a plain link stays plain. */
export function groupingSearch(
  search: string,
  grouping: PlannerGrouping,
  raid: Pick<GuildRaidView, 'zone_kind'>,
): string {
  const params = new URLSearchParams(search);
  if (grouping === defaultGrouping(raid)) params.delete(GROUP_PARAM);
  else params.set(GROUP_PARAM, grouping);
  const query = params.toString();
  return query === '' ? '' : `?${query}`;
}

export function groupingOptions(raid: Pick<GuildRaidView, 'zone_kind'>): { value: string; label: string }[] {
  return groupingsOf(raid).map((g) => ({ value: g, label: GROUPING_WORDS[g] }));
}

/** The phone bar's line: `Group by role`. */
export function groupingSummary(grouping: PlannerGrouping): string {
  return `Group by ${GROUPING_WORDS[grouping].toLowerCase()}`;
}

export interface PlannerGroup {
  /** `tank`, `late`, `warrior`: the group's test id suffix. */
  id: string;
  label: string;
  rows: RosterRow[];
  /** Said when the group has nobody (the four role groups only). */
  empty?: string;
}

const STATE_ORDER: readonly GuildSignupState[] = ['in', 'late', 'bench', 'out'];
const STATE_LABELS: Readonly<Record<GuildSignupState, string>> = {
  in: raidCopy.segments.in,
  late: raidCopy.subheads.late,
  bench: raidCopy.subheads.bench,
  out: raidCopy.subheads.out,
};
const GROUP_LABELS: Readonly<Record<GuildRaidGroup, string>> = {
  tank: raidCopy.subheads.tanks,
  healer: raidCopy.subheads.healers,
  melee: raidCopy.subheads.melee,
  ranged: raidCopy.subheads.ranged,
};

function stateRank(state: GuildSignupState | null): number {
  return state === null ? STATE_ORDER.length : STATE_ORDER.indexOf(state);
}

function signedRows(raid: GuildRaidView): RosterRow[] {
  const groups = rosterGroups(raid);
  return [...groups.in, ...groups.late, ...groups.bench, ...groups.out];
}

function byStatus(raid: GuildRaidView): PlannerGroup[] {
  const rows = signedRows(raid);
  return STATE_ORDER.map((state) => ({
    id: state,
    label: STATE_LABELS[state],
    rows: rows.filter((row) => row.state === state).sort(rosterOrder),
  })).filter((group) => group.rows.length > 0);
}

function byRole(raid: GuildRaidView): PlannerGroup[] {
  const groups = rosterGroups(raid);
  const roles = RAID_GROUPS.map((group) => ({
    id: group,
    label: GROUP_LABELS[group],
    rows: groups.groups[group],
    empty: raidCopy.emptyGroup[group],
  }));
  const rest = (['late', 'bench', 'out'] as const)
    .map((state) => ({ id: state, label: STATE_LABELS[state], rows: groups[state] }))
    .filter((group) => group.rows.length > 0);
  return [...roles, ...rest];
}

/** `Warriors` from `warrior`: every class name takes a plain s. */
export function classPlural(classSlug: string): string {
  return `${classWord(classSlug)}s`;
}

function byClass(raid: GuildRaidView): PlannerGroup[] {
  const rows = signedRows(raid);
  const classes = [...new Set(rows.map((row) => row.class.toLowerCase()))].sort(
    (a, b) => classRank(a) - classRank(b) || a.localeCompare(b),
  );
  return classes.map((slug) => ({
    id: CLASS_ORDER.includes(slug) ? slug : `class-${slug}`,
    label: classPlural(slug),
    rows: rows
      .filter((row) => row.class.toLowerCase() === slug)
      .sort((a, b) => stateRank(a.state) - stateRank(b.state) || a.name.localeCompare(b.name)),
  }));
}

/** The Plate's groups in the chosen grouping. */
export function plannerGroups(raid: GuildRaidView, grouping: PlannerGrouping): PlannerGroup[] {
  if (grouping === 'class') return byClass(raid);
  if (grouping === 'role' && kindCountsRoles(kindOf(raid))) return byRole(raid);
  return byStatus(raid);
}

/** The row's sub line: `Holy Paladin · Healer` (a raid or a run), the class alone at an event. */
export function rowSub(row: Pick<RosterRow, 'class' | 'spec_key' | 'role'>, roles: boolean): string {
  if (!roles) return classWord(row.class);
  return `${specClassLine(row.class, row.spec_key)}${SEP}${raidCopy.roleLabels[row.role]}`;
}

export type StateTone = 'success' | 'warn' | 'muted' | 'danger';

const STATE_TONE: Readonly<Record<GuildSignupState, StateTone>> = {
  in: 'success',
  late: 'warn',
  bench: 'muted',
  out: 'danger',
};

export function stateTag(state: GuildSignupState): { text: string; tone: StateTone } {
  return { text: raidCopy.segments[state], tone: STATE_TONE[state] };
}

const STATUS_TAGS: Readonly<Record<RaidChipStatus, { text: string; tone: StateTone | 'gold' }>> = {
  planned: { text: 'Planned', tone: 'gold' },
  locked: { text: 'Locked', tone: 'warn' },
  ended: { text: 'Ended', tone: 'muted' },
  cancelled: { text: 'Cancelled', tone: 'danger' },
  done: { text: 'Done', tone: 'success' },
};

/** The head's state Tag: the raid's status, or the derived `ended`. */
export function plannerStatusTag(status: RaidChipStatus): { text: string; tone: StateTone | 'gold' } {
  return STATUS_TAGS[status];
}

/** The Section's meta and the header's counts: `19 In · 2 Late · 2 Bench`. */
export function stateCountsLine(counts: GuildRaidView['counts']): string {
  return [`${counts.in} In`, `${counts.late} Late`, `${counts.bench} Bench`].join(SEP);
}

/** The officers' Bench and swaps: the benched sign-ups in the roster's order. */
export function benchRows(raid: GuildRaidView): RosterRow[] {
  return rosterGroups(raid).bench;
}

export interface NotSignedRow {
  character_key: string;
  name: string;
  class: string;
  sub: string;
  /** The row an officer's state choice writes (the account's main; its spec's role). */
  roster: RosterRow;
}

/** Officers: one row per account with no sign-up, its main; who has no Discord linked says so. */
export function notSignedRows(raid: GuildRaidView): NotSignedRow[] {
  const roles = kindCountsRoles(kindOf(raid));
  const known = new Map(raid.not_signed.map((m) => [m.character_key, m]));
  return (raid.not_signed_raiders ?? [])
    .map((raider): NotSignedRow => {
      const member = known.get(raider.character_key);
      const role = roleForSpec(member?.spec ?? '') ?? 'dps';
      const specKey = member?.spec_key ?? null;
      const who = rowSub({ class: raider.class, spec_key: specKey, role }, roles);
      return {
        character_key: raider.character_key,
        name: raider.name,
        class: raider.class,
        sub: raider.discord_linked ? who : `${who}${SEP}No Discord linked`,
        roster: {
          character_key: raider.character_key,
          name: raider.name,
          class: raider.class,
          spec: member?.spec ?? '',
          role,
          group: member?.group ?? (role === 'dps' ? 'ranged' : role),
          spec_key: specKey,
          state: null,
          note: '',
          own: false,
          also: null,
        },
      };
    })
    .sort((a, b) => classRank(a.class) - classRank(b.class) || a.name.localeCompare(b.name));
}

/** Why the ping button is off, in words; `null` while it may go out. */
export function pingBlockedLine(
  raid: Pick<GuildRaidView, 'discord_post' | 'ping'>,
  now: Date,
  clock: (at: Date) => string,
): string | null {
  if (raid.discord_post === null) return 'Post the raid to Discord first, then ping under it.';
  if (pingAllowed(raid, now)) return null;
  const next = raid.ping?.next_allowed_at;
  return next == null ? null : `Pinged already. The next ping can go out at ${clock(new Date(next))}.`;
}

export const COMP_ROLES: readonly GuildCompRole[] = ['tanks', 'healers', 'melee', 'ranged'];

export type CompDraft = Readonly<Record<GuildCompRole, string>>;

export function compDraft(target: Readonly<Record<GuildCompRole, number>>): CompDraft {
  return {
    tanks: String(target.tanks),
    healers: String(target.healers),
    melee: String(target.melee),
    ranged: String(target.ranged),
  };
}

export type CompParse =
  | { ok: true; target: Readonly<Record<GuildCompRole, number>> }
  | { ok: false; fields: Readonly<Record<string, string>> };

const WHOLE = /^\d{1,3}$/;

/** The draft as seats; a field that is not a whole number 0 or more is refused before the API. */
export function parseCompDraft(draft: CompDraft): CompParse {
  const fields: Record<string, string> = {};
  const target: Record<string, number> = {};
  for (const role of COMP_ROLES) {
    const raw = draft[role].trim();
    if (WHOLE.test(raw)) target[role] = Number(raw);
    else fields[role] = 'Use a whole number, 0 or more.';
  }
  if (Object.keys(fields).length > 0) return { ok: false, fields };
  return { ok: true, target: target as Record<GuildCompRole, number> };
}
