// web/src/lib/guild/overview-role-view.ts
// The role-aware Overview's pure models (web layout spec 2026-11-05 section 7.1): the member's
// private checklist rows and its one-line verdict (D12: Ready / Almost / Getting started, never a
// floor), the officer's Action needed rows, comp against the target, and the attendance grid.
// Rows whose source is not tracked yet are dropped here, never drawn with a stand-in (7.4).
import { bisHref } from '../character-profile';
import { counted } from '../plural';
import { tallyTone, type CellTone, type MarkTone } from '../ui/cells';
import type { GuildRaidCounts } from './api';
import { enchantSlotLabel as slotWord } from './readiness-view';
import type {
  AttendanceCell,
  ChecklistItem,
  ChecklistVerdict,
  GuildChecklist,
  GuildCompRole,
  GuildCompTarget,
  GuildOfficerActions,
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
  return list.items.flatMap((item): ChecklistRow[] => {
    if (item.state === 'not_tracked') return [];
    switch (item.key) {
      case 'level':
        return [levelRow(item, zone)];
      case 'attunement':
        return [attunementRow(item, zone)];
      case 'gear_gap':
        return [gearRow(item, bis)];
      case 'enchants':
        return [enchantsRow(item)];
      default:
        return [];
    }
  });
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
