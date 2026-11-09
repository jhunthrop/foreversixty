// web/src/lib/guild/overview-role-view.ts
// The role-aware Overview's pure models (web layout spec 2026-11-05 section 7.1): the member's
// private checklist rows and its one-line verdict (D12: Ready / Almost / Getting started, never a
// floor), the officer's Action needed rows, comp against the target, and the attendance grid.
// Rows whose source is not tracked yet are dropped here, never drawn with a stand-in (7.4).
import { bisHref } from '../character-profile';
import { counted } from '../plural';
import type { CellTone, MarkTone } from '../ui/cells';
import type { GridRowModel } from '../ui/cells';
import type { GuildRaidCounts, GuildRaidCoverRow, GuildRaidView } from './api';
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

/** A quest attunement (Onyxia's Lair): the zone's chain, and the item it hands out when it has one
 *  ("Finish the attunement chain for Onyxia's Lair and keep the Drakefire Amulet in your bags."). */
function questAttunementRow(item: ChecklistItem, zone: string): ChecklistRow {
  const reward = item.key_item?.name;
  const title = `Attuned to ${zone}`;
  if (item.state === 'ok') {
    const sub =
      reward === undefined ? 'Attunement chain done.' : `Attunement chain done. Keep the ${reward} on you.`;
    return { key: item.key, tone: 'ok', title, sub, fix: null };
  }
  const keep = reward === undefined ? '' : ` and keep the ${reward} in your bags`;
  return {
    key: item.key,
    tone: 'miss',
    title,
    sub: `Finish the attunement chain for ${zone}${keep}.`,
    fix: null,
  };
}

function attunementRow(item: ChecklistItem, zone: string): ChecklistRow {
  if (item.attunement_kind === 'quest') return questAttunementRow(item, zone);
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
    return {
      key: item.key,
      tone: 'ok',
      title: 'Flask and food',
      sub: `You carry ${listWords(names)}.`,
      fix: null,
    };
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

/** What the member shares to fix an untracked check, by the API's reason: the addon setup shows how. */
const SHARE_ROWS: Readonly<Record<string, { title: string; sub: string }>> = {
  gear_not_shared: { title: 'Share your gear', sub: 'Your addon export has no gear on file yet.' },
  quests_not_in_export: {
    title: 'Share your quests',
    sub: 'Update the addon and sync to check your attunement.',
  },
  bags_not_shared: {
    title: 'Share your bags',
    sub: 'Allow bag sharing in the addon to check your flask and food.',
  },
};

/** A check we cannot read that the member can fix themselves (share the export); the rest of the
 *  untracked checks (keys, named consumables, resistance, no BiS list) are not drawn. */
function actionableRow(item: ChecklistItem): ChecklistRow[] {
  const fix = { label: 'How', href: SETUP_HREF };
  const share = item.reason === undefined ? undefined : SHARE_ROWS[item.reason];
  if (share !== undefined) return [{ key: item.key, tone: 'warn', ...share, fix }];
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
  key: 'approvals' | 'applications' | 'not_signed' | 'drops';
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
      title: `${counted(approvals.count, 'raider')} waiting for approval`,
      sub: names(approvals.names),
    });
  const applications = actions.applications;
  if (applications !== undefined && applications.count > 0)
    rows.push({
      key: 'applications',
      title: counted(applications.count, 'application'),
      sub: names(applications.names),
    });
  if (notSigned.count > 0 && notSigned.raid_id !== null)
    rows.push({
      key: 'not_signed',
      title: `${notSigned.count} not signed up for ${day ?? 'the next raid'}`,
      sub: names(notSigned.names),
    });
  if (drops !== null && drops.tracked && (drops.count ?? 0) > 0)
    rows.push({ key: 'drops', title: `${counted(drops.count ?? 0, 'drop')} to award`, sub: drops.label });
  return rows;
}

const ROLE_LABEL: Readonly<Record<GuildCompRole, string>> = {
  tanks: 'Tanks',
  healers: 'Healers',
  melee: 'Melee',
  ranged: 'Ranged',
};
const ROLES: readonly GuildCompRole[] = ['tanks', 'healers', 'melee', 'ranged'];

/** Comp against the target (batch 3): per role the seats, who is In and how many short. In is met,
 *  met counting Late (warn) or short; the short cell takes that tone, or a muted dash at none. Late is
 *  named under the role. */
export function compGridRows(counts: GuildRaidCounts, target: GuildCompTarget): GridRowModel[] {
  const inByRole: Readonly<Record<GuildCompRole, number>> = {
    tanks: counts.tanks,
    healers: counts.healers,
    melee: counts.melee ?? 0,
    ranged: counts.ranged ?? 0,
  };
  return ROLES.map((role) => {
    const have = inByRole[role];
    const want = target[role];
    const late = counts.late_groups?.[role] ?? 0;
    const tone: CellTone = have >= want ? 'ok' : have + late >= want ? 'warn' : 'miss';
    const short = Math.max(0, want - have);
    return {
      key: role,
      label: ROLE_LABEL[role],
      sub: late > 0 ? `${late} late` : undefined,
      cells: [
        { tone: 'none', text: String(want), label: `Needs ${want}` },
        { tone, text: String(have), label: `${have} in` },
        short > 0
          ? { tone, text: String(short), label: `${short} short` }
          : { tone: 'none', text: '—', label: 'None short' },
      ],
    };
  });
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

/** The attendance section's title: "Attendance, last raid" for one night, "Attendance, last 4 raids" for four. */
export function attendanceTitle(raids: number): string {
  return raids === 1 ? 'Attendance, last raid' : `Attendance, last ${raids} raids`;
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

/** The cover rows as Grid rows: what the raid needs, then who is signed In or Late to give it. */
export function coverGridRows(
  rows: readonly GuildRaidCoverRow[],
): (GridRowModel & { row: GuildRaidCoverRow })[] {
  return rows.map((row) => {
    const likely = row.likely > 0 ? `${row.likely} talent not confirmed` : '';
    const bench = row.providers_bench > 0 ? `${row.providers_bench} on the bench` : '';
    return {
      key: row.key,
      label: row.name,
      sub: [likely, bench].filter(Boolean).join(' · ') || undefined,
      row,
      cells: [
        { tone: 'none', text: String(row.need), label: `Needs ${row.need}` },
        { tone: coverTone(row), text: String(coverSigned(row)), label: `${coverSigned(row)} signed` },
      ],
    };
  });
}

/** The ping may go out: the raid has a Discord post and the 6 h throttle has passed. */
export function pingAllowed(raid: Pick<GuildRaidView, 'discord_post' | 'ping'>, now: Date): boolean {
  if (raid.discord_post === null) return false;
  const next = raid.ping?.next_allowed_at ?? null;
  return next === null || new Date(next) <= now;
}

/** The ping's answer in words: who was named, and who could not be. */
export function pingLine(result: GuildPingResult): string {
  if (!result.posted) return 'Nobody to ping: none of them has linked Discord.';
  const named = `Pinged ${counted(result.mentioned, 'raider')} on Discord.`;
  if (result.no_discord === 0) return named;
  return `${named} No Discord linked: ${names(result.no_discord_names)}.`;
}
