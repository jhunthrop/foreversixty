// web/src/lib/guild/overview-role-view.ts
// The role-aware Overview's pure models (web layout spec 2026-11-05 section 7.1): the member's
// private checklist rows and its one-line verdict (D12: Ready / Almost / Getting started, never a
// floor), the officer's Action needed rows, comp against the target, and the attendance grid.
// Rows whose source is not tracked yet are dropped here, never drawn with a stand-in (7.4).
import { bisHref } from '../character-profile';
import { counted } from '../plural';
import { tallyTone, type CellTone, type MarkTone } from '../ui/cells';
import type { GridRowModel } from '../ui/cells';
import type { GuildRaidCounts, GuildRaidCoverRow } from './api';
import { coverSigned, coverTone } from './cover-view';
import { DEFAULT_LOCALE, clockText, dateLabel, offsetsDiffer, zoneAbbreviation } from './raid-time';
import { enchantSlotLabel as slotWord } from './readiness-view';
import type {
  AttendanceCell,
  ChecklistItem,
  ChecklistVerdict,
  GuildChecklist,
  GuildCompRole,
  GuildCompTarget,
  GuildOfficerActions,
  GuildPingResult,
  GuildRecentAttendance,
} from './overview-api';

export interface ChecklistRow {
  key: ChecklistItem['key'];
  tone: MarkTone;
  title: string;
  sub: string;
  fix: { label: string; href: string } | null;
}

const WORDS = ['no', 'one', 'two', 'three', 'four', 'five', 'six', 'seven', 'eight', 'nine'];

/** "two things" in the verdict line; digits from ten. */
function things(count: number): string {
  const word = WORDS[count] ?? String(count);
  return `${word} ${count === 1 ? 'thing' : 'things'}`;
}

function listWords(items: readonly string[]): string {
  if (items.length <= 1) return items.join('');
  return `${items.slice(0, -1).join(', ')} and ${items[items.length - 1]}`;
}

function levelRow(item: ChecklistItem, zone: string): ChecklistRow {
  const need = item.required_level ?? 60;
  if (item.state === 'ok')
    return {
      key: item.key,
      tone: 'ok',
      title: `Level ${need}`,
      sub: `You are level ${item.level ?? need}.`,
      fix: null,
    };
  const sub = item.level === undefined ? 'Your level is not on file yet.' : `You are level ${item.level}.`;
  return { key: item.key, tone: 'miss', title: `Level ${need} for ${zone}`, sub, fix: null };
}

function attunementRow(item: ChecklistItem, zone: string): ChecklistRow {
  if (item.attunement_kind === 'quest') {
    const chain = item.key_item?.name;
    const title = `Attuned to ${zone}`;
    if (item.state === 'ok')
      return { key: item.key, tone: 'ok', title, sub: 'Your attunement is done.', fix: null };
    const sub =
      chain === undefined ? 'Your attunement quest is not done yet.' : `The ${chain} quest is not done yet.`;
    return { key: item.key, tone: 'miss', title, sub, fix: null };
  }
  const key = item.key_item?.name ?? 'The key';
  if (item.state === 'ok')
    return { key: item.key, tone: 'ok', title: `Key for ${zone}`, sub: `${key} is in your bags.`, fix: null };
  return {
    key: item.key,
    tone: 'miss',
    title: `Key for ${zone}`,
    sub: `${key} is not in your bags.`,
    fix: null,
  };
}

const FLASK_FOOD_MISSING: Readonly<Record<string, string>> = {
  flask_or_elixir: 'No flask or elixir in your bags',
  food: 'No buff food in your bags',
  both: 'No flask, elixir or buff food in your bags',
};

function flaskFoodRow(item: ChecklistItem): ChecklistRow {
  const names = (item.supplies ?? []).map((stack) => stack.name);
  const missing = item.missing ?? [];
  if (item.state === 'ok')
    return { key: item.key, tone: 'ok', title: 'Flask and food', sub: `${listWords(names)}.`, fix: null };
  const title =
    FLASK_FOOD_MISSING[missing.length > 1 ? 'both' : (missing[0] ?? 'both')] ?? FLASK_FOOD_MISSING.both!;
  const sub =
    names.length > 0 ? `You have ${listWords(names)}.` : 'Bring a flask or a raid elixir, and buff food.';
  return { key: item.key, tone: 'warn', title, sub, fix: null };
}

function fireResistanceRow(item: ChecklistItem): ChecklistRow {
  return {
    key: item.key,
    tone: 'ok',
    title: `Fire resistance ${item.resistance ?? 0}`,
    sub: 'From your gear and buffs when you last synced.',
    fix: null,
  };
}

function gearRow(item: ChecklistItem, bis: string | null): ChecklistRow {
  if (item.state === 'ok')
    return {
      key: item.key,
      tone: 'ok',
      title: 'Best in slot',
      sub: 'Every slot is your best in slot.',
      fix: null,
    };
  const upgrades = item.upgrades ?? item.slots?.length ?? 0;
  const slots = item.slots ?? [];
  return {
    key: item.key,
    tone: 'warn',
    title: `${counted(upgrades, 'slot')} ${upgrades === 1 ? 'is' : 'are'} not best in slot`,
    sub: slots.length > 0 ? `${listWords(slots.map(slotWord))}.` : 'The BiS list shows the swaps.',
    fix: bis === null ? null : { label: 'Fix in the BiS list', href: bis },
  };
}

function enchantsRow(item: ChecklistItem): ChecklistRow {
  if (item.state === 'ok')
    return {
      key: item.key,
      tone: 'ok',
      title: 'Enchants',
      sub: 'Every slot that takes one is enchanted.',
      fix: null,
    };
  const missing = item.missing_slots ?? [];
  return {
    key: item.key,
    tone: 'warn',
    title: `${counted(missing.length, 'slot')} without an enchant`,
    sub: `${listWords(missing.map(slotWord))}.`,
    fix: null,
  };
}

/** The checklist as rows, tracked items only, in the API's order. */
export function checklistRows(list: GuildChecklist, race: string | null = null): ChecklistRow[] {
  const zone = list.raid?.zone_name ?? 'the raid';
  const bis = bisHref(list.character.class, list.character.spec || null, list.character.level, race);
  // Gear and enchants both go unread when the export carries no gear: one row asks for it.
  const asked = new Set<string>();
  return list.items.flatMap((item): ChecklistRow[] => {
    if (item.state === 'not_tracked') {
      if (item.reason === undefined || asked.has(item.reason)) return [];
      asked.add(item.reason);
      return actionableRow(item);
    }
    switch (item.key) {
      case 'level':
        return [levelRow(item, zone)];
      case 'attunement':
        return [attunementRow(item, zone)];
      case 'gear_gap':
        return [gearRow(item, bis)];
      case 'enchants':
        return [enchantsRow(item)];
      case 'flask_food':
        return [flaskFoodRow(item)];
      case 'fire_resistance':
        return [fireResistanceRow(item)];
      default:
        return [];
    }
  });
}

const SETUP_HREF = '/setup';

/** A check we cannot read that the member can fix themselves (share the export); the rest of the
 *  untracked checks (keys, named consumables, resistance, no BiS list) are not drawn. */
function actionableRow(item: ChecklistItem): ChecklistRow[] {
  const fix = { label: 'How', href: SETUP_HREF };
  if (item.reason === 'gear_not_shared')
    return [
      {
        key: item.key,
        tone: 'warn',
        title: 'Share your gear',
        sub: 'Your addon export has no gear on file yet.',
        fix,
      },
    ];
  if (item.key === 'level' && item.reason === 'no_level')
    return [
      { key: item.key, tone: 'warn', title: 'Share your level', sub: 'Your level is not on file yet.', fix },
    ];
  return [];
}

const VERDICT_WORD: Readonly<Record<ChecklistVerdict, string>> = {
  ready: 'Ready to raid',
  almost: 'Almost there',
  getting_started: 'Getting started',
};

export function verdictWord(verdict: ChecklistVerdict): string {
  return VERDICT_WORD[verdict];
}

/** The member's Notice: "Almost there: two things before Thursday." (`day` is the raid's weekday). */
export function verdictLine(
  verdict: ChecklistVerdict,
  rows: readonly ChecklistRow[],
  day: string | null,
): string {
  const open = rows.filter((row) => row.tone !== 'ok').length;
  const when = day === null ? 'before the raid' : `before ${day}`;
  if (verdict === 'ready' || open === 0) return day === null ? 'Ready to raid.' : `Ready to raid on ${day}.`;
  return `${VERDICT_WORD[verdict]}: ${things(open)} ${when}.`;
}

export interface ActionRow {
  key: 'approvals' | 'not_signed' | 'drops';
  title: string;
  sub: string;
}

function names(list: readonly string[], max = 3): string {
  if (list.length <= max) return list.join(', ');
  return `${list.slice(0, max).join(', ')} and ${list.length - max} more`;
}

/** Officer Action needed: only what has something waiting, in the order an officer clears it. */
export function actionRows(actions: GuildOfficerActions, day: string | null): ActionRow[] {
  const rows: ActionRow[] = [];
  const { approvals, not_signed: notSigned, drops_to_award: drops } = actions;
  if (approvals.count > 0)
    rows.push({
      key: 'approvals',
      title: `${counted(approvals.count, 'raider')} ${approvals.count === 1 ? 'is' : 'are'} waiting for approval`,
      sub: names(approvals.names),
    });
  if (notSigned.count > 0 && notSigned.raid_id !== null)
    rows.push({
      key: 'not_signed',
      title: `${notSigned.count} ${notSigned.count === 1 ? 'has' : 'have'} not signed up for ${day ?? 'the next raid'}`,
      sub: names(notSigned.names),
    });
  if (drops !== null && drops.tracked && (drops.count ?? 0) > 0)
    rows.push({ key: 'drops', title: `${counted(drops.count ?? 0, 'drop')} to award`, sub: drops.label });
  return rows;
}

export interface TallyRow {
  role: GuildCompRole;
  label: string;
  have: number;
  late: number;
  want: number;
  tone: MarkTone;
}

const ROLE_LABEL: Readonly<Record<GuildCompRole, string>> = {
  tanks: 'Tanks',
  healers: 'Healers',
  melee: 'Melee',
  ranged: 'Ranged',
};
const ROLES: readonly GuildCompRole[] = ['tanks', 'healers', 'melee', 'ranged'];

/** Comp against the target: In per role against the seats; Late is shown beside it, not counted. */
export function compTallyRows(counts: GuildRaidCounts, target: GuildCompTarget): TallyRow[] {
  const inByRole: Readonly<Record<GuildCompRole, number>> = {
    tanks: counts.tanks,
    healers: counts.healers,
    melee: counts.melee ?? 0,
    ranged: counts.ranged ?? 0,
  };
  return ROLES.map((role) => ({
    role,
    label: ROLE_LABEL[role],
    have: inByRole[role],
    late: counts.late_groups?.[role] ?? 0,
    want: target[role],
    tone: tallyTone(inByRole[role], target[role]),
  }));
}

export function compTargetLine(zone: string, target: GuildCompTarget): string {
  const seats = target.tanks + target.healers + target.melee + target.ranged;
  return target.source === 'default'
    ? `${zone} default: ${seats} players`
    : `Set for this raid: ${seats} players`;
}

const CELL: Readonly<Record<AttendanceCell, { tone: CellTone; text: string; label: string }>> = {
  present: { tone: 'ok', text: '✓', label: 'There' },
  absent: { tone: 'miss', text: '×', label: 'Missed' },
  bench: { tone: 'warn', text: 'B', label: 'Bench' },
  none: { tone: 'none', text: '—', label: 'Not in the guild yet' },
};

export interface GridRow {
  key: string;
  label: string;
  sub: string;
  cells: { tone: CellTone; text: string; label: string }[];
}

/** The officer's attendance grid: slipping first, then the bench, each by name. */
export function attendanceGridRows(page: GuildRecentAttendance): GridRow[] {
  return [...page.rows]
    .sort((a, b) => Number(b.slipping) - Number(a.slipping) || a.name.localeCompare(b.name))
    .map((row) => ({
      key: row.account_key,
      label: row.name,
      sub: row.slipping ? 'slipping' : 'bench',
      cells: row.cells.map((cell) => CELL[cell]),
    }));
}

/** The weekday a raid falls on in the guild's zone, for "before Thursday". */
export function raidWeekday(startsAt: string, tz: string, locale: string = DEFAULT_LOCALE): string {
  return new Intl.DateTimeFormat(locale, { weekday: 'long', timeZone: tz }).format(new Date(startsAt));
}

/** `Thu 10 Dec · 8 pm CST`, then `· 9 pm EST your time` when the viewer's zone differs. */
export function raidTimeLine(
  startsAt: string,
  guildTz: string,
  viewerTz: string,
  locale: string = DEFAULT_LOCALE,
): string {
  const at = new Date(startsAt);
  const guild = `${dateLabel(at, guildTz, locale)} · ${clockText(at, guildTz, locale)} ${zoneAbbreviation(guildTz, at)}`;
  if (!offsetsDiffer(at, guildTz, viewerTz)) return guild;
  return `${guild} · ${clockText(at, viewerTz, locale)} ${zoneAbbreviation(viewerTz, at)} your time`;
}

const PER_WORDS: Readonly<Record<GuildRaidCoverRow['per'], string>> = {
  raid: '',
  melee_group: 'per melee group',
  caster_group: 'per caster group',
};

/** The cover rows as Grid rows: what the raid needs, then who is signed In or Late to give it. */
export function coverGridRows(
  rows: readonly GuildRaidCoverRow[],
): (GridRowModel & { row: GuildRaidCoverRow })[] {
  return rows.map((row) => {
    const per = PER_WORDS[row.per];
    const likely = row.likely > 0 ? `${row.likely} on a talent we cannot see` : '';
    const bench = row.providers_bench > 0 ? `${row.providers_bench} on the bench` : '';
    return {
      key: row.key,
      label: row.name,
      sub: [per, likely, bench].filter(Boolean).join(' · ') || undefined,
      row,
      cells: [
        { tone: 'none', text: String(row.need), label: `Needs ${row.need}` },
        { tone: coverTone(row), text: String(coverSigned(row)), label: `${coverSigned(row)} signed` },
      ],
    };
  });
}

/** The ping's answer in words: who was named, and who could not be. */
export function pingLine(result: GuildPingResult): string {
  if (!result.posted) return 'Nobody to ping: none of them has linked Discord.';
  const named = `Pinged ${counted(result.mentioned, 'raider')} on Discord.`;
  if (result.no_discord === 0) return named;
  return `${named} No Discord linked: ${names(result.no_discord_names)}.`;
}

/** The header's next-raid fact on every tab: `Next raid: Onyxia's Lair, Thu 10 Dec` (guild calendar). */
export function nextRaidFact(
  next: { zone_name: string; starts_at: string; guild_timezone: string | null },
  viewerTz: string = Intl.DateTimeFormat().resolvedOptions().timeZone,
): string {
  return `Next raid: ${next.zone_name}, ${dateLabel(new Date(next.starts_at), next.guild_timezone ?? viewerTz)}`;
}
