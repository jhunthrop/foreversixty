// web/src/lib/guild/profile-view.ts
// One member's profile as plain data (web layout spec 2026-11-05 §7.2 Member profile, "What do I
// know about this raider?"): the head's facts and figures, the gear gap in the BiS page's unit
// (totals, as the API gives them), the last 4 nights as a Grid row, loot this tier and the notes.
// Pure, so a test pins the words; the notes are never read for a member (the API leaves them out).
import type { GridCell, GridRowModel } from '../ui/cells';
import { outputFigure, topOutputLabel } from './output-figure';
import { dateLabel } from './raid-time';
import { specClassName } from './spec-name';
import { profileCopy as copy } from './recruit-copy';
import type { MemberNote, MemberProfile, ProfileNight } from './recruit-api';

const SEP = ' · ';

const RANK_WORD: Readonly<Record<string, string>> = {
  leader: 'Guild master',
  officer: 'Officer',
  member: 'Member',
};

/** The head's title, class and fact line: "Frost Mage · Level 60 · Officer". */
export function profileHead(profile: MemberProfile): { name: string; classSlug: string; facts: string[] } {
  const row = profile.member;
  const first = profile.characters[0];
  if (row === null) {
    return { name: first?.name ?? '', classSlug: first?.class ?? '', facts: [] };
  }
  const facts = [
    specClassName(row.spec ?? '', row.class),
    row.level === undefined || row.level === null ? '' : copy.level(row.level),
    RANK_WORD[row.rank] ?? '',
  ].filter((part) => part !== '');
  return { name: row.name, classSlug: row.class ?? '', facts };
}

export interface ProfileFigure {
  key: string;
  label: string;
  value: string;
}

/** Attendance, the best boss figure, item level and items this tier: only what has a source. */
export function profileFigures(profile: MemberProfile): ProfileFigure[] {
  const figures: ProfileFigure[] = [];
  const last4 = profile.attendance?.last_4;
  if (last4 !== undefined && last4.nights > 0) {
    figures.push({
      key: 'attendance',
      label: copy.figures.attendance,
      value: `${last4.present} of ${last4.nights}`,
    });
  }
  const parse = profile.member?.best_parse;
  if (parse !== undefined && parse !== null) {
    figures.push({ key: 'parse', label: topOutputLabel(parse.metric), value: outputFigure(parse.value) });
  }
  const ilvl = profile.member?.item_level;
  if (ilvl !== undefined) figures.push({ key: 'ilvl', label: copy.figures.itemLevel, value: String(ilvl) });
  figures.push({ key: 'items', label: copy.figures.items, value: String(profile.loot.count) });
  return figures;
}

export interface GapLine {
  key: string;
  text: string;
  tone: 'strong' | 'muted';
}

const UNIT_WORD = { dps: 'DPS', hps: 'HPS', tps: 'TPS' } as const;

/** The gear gap's totals in the BiS page's unit; a healer or tank (no unit) gets the count alone. */
export function gearGapLines(profile: MemberProfile): GapLine[] | null {
  const gap = profile.readiness?.gear_gap;
  if (gap === undefined || gap === null) return null;
  const upgrades = gap.upgrades ?? 0;
  if (upgrades === 0) return [{ key: 'none', text: copy.noGap, tone: 'muted' }];
  const lines: GapLine[] = [{ key: 'upgrades', text: copy.upgrades(upgrades), tone: 'strong' }];
  if (gap.unit !== null && gap.gain_dps !== null && gap.gain_dps > 0) {
    lines.push({
      key: 'gain',
      text: copy.gain(gap.gain_dps.toFixed(1), UNIT_WORD[gap.unit]),
      tone: 'strong',
    });
  }
  if (gap.not_sim_checked > 0) {
    lines.push({ key: 'unchecked', text: copy.notSimChecked(gap.not_sim_checked), tone: 'muted' });
  }
  return lines;
}

function nightCell(night: ProfileNight): GridCell {
  if (night.excused) return { tone: 'none', text: 'E', label: copy.cells.excused };
  switch (night.outcome) {
    case 'present':
      return { tone: 'ok', text: '✓', label: copy.cells.present };
    case 'late':
      return { tone: 'warn', text: 'L', label: copy.cells.late };
    case 'absent':
      return { tone: 'miss', text: '×', label: copy.cells.absent };
    case 'no_show':
      return { tone: 'miss', text: '×', label: copy.cells.no_show };
    default:
      return { tone: 'none', text: '—', label: copy.cells.none };
  }
}

export interface AttendanceView {
  /** The last 4 raids oldest first, as a Grid: one column per night, one row. */
  columns: string[];
  rows: GridRowModel[];
  /** "25 of the last 30 raids", or the reports-mode line; '' with nothing logged. */
  line: string;
}

/** Raids mode: the 4 nights one by one and the 30-raid figure; reports mode: the window's line. */
export function attendanceView(
  profile: MemberProfile,
  viewerTz: string,
  locale: string,
): AttendanceView | null {
  const att = profile.attendance;
  if (att === null) return null;
  if (att.source === 'reports') {
    const line =
      att.last_4.nights === 0 ? copy.noNights : copy.reportsMode(att.last_4.present, att.last_4.nights);
    return { columns: [], rows: [], line };
  }
  const nights = [...att.nights].sort((a, b) => a.starts_at.localeCompare(b.starts_at));
  const thirty = att.last_30;
  const line = thirty === null || thirty.nights === 0 ? '' : copy.lastThirty(thirty.present, thirty.nights);
  if (nights.length === 0) return { columns: [], rows: [], line: line === '' ? copy.noNights : line };
  return {
    columns: nights.map((night) => dateLabel(new Date(night.starts_at), viewerTz, locale)),
    rows: [{ key: profile.account_key, label: copy.lastNights(nights.length), cells: nights.map(nightCell) }],
    line,
  };
}

export interface ProfileLootRow {
  key: string;
  id: number;
  name: string;
  quality: number;
  icon: string | undefined;
  source: string;
}

/** Loot this tier as ItemRows, newest first: "Awarded Tue 8 Dec". */
export function lootRows(profile: MemberProfile, viewerTz: string, locale: string): ProfileLootRow[] {
  return [...profile.loot.items]
    .sort((a, b) => b.awarded_at.localeCompare(a.awarded_at))
    .map((item, i) => ({
      key: `${item.item_id}-${item.awarded_at}-${i}`,
      id: item.item_id,
      name: item.name ?? `Item ${item.item_id}`,
      quality: item.quality ?? 1,
      icon: item.icon === null || item.icon === '' ? undefined : item.icon,
      source: copy.awarded(dateLabel(new Date(item.awarded_at), viewerTz, locale)),
    }));
}

/** "Ironbrand · Wed 2 Dec": who wrote it and when; "An officer" for one who left. */
export function noteSub(note: MemberNote, viewerTz: string, locale: string): string {
  const who = note.author?.name ?? copy.anOfficer;
  return `${who}${SEP}${dateLabel(new Date(note.created_at), viewerTz, locale)}`;
}

/** The notes with one added at the top (newest first), or one removed; never mutated. */
export const withNote = (notes: readonly MemberNote[], note: MemberNote): MemberNote[] => [note, ...notes];
export const withoutNote = (notes: readonly MemberNote[], id: number): MemberNote[] =>
  notes.filter((note) => note.id !== id);

export const NOTE_MAX = 1000;
