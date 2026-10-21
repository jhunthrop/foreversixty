// The homepage's horizontal dates row (spec 2026-09-24 §2.2), replacing the "Right now"
// StatePanel. Every status is decided by the clock, never by hand, so the row stays true
// without an edit: a date before `now` is past, a span (a start and an `endIso`) is live
// from its start through the end of its last day, the first date still to come is the next
// one, and everything after that is future. The row is shown next-first (live, next, future)
// with the past dates last and muted, so a visitor meets what is coming before what is done.
// Reuses lib/dates.ts's own `isPast` (strict less-than) rather than a second date comparison.
import { isPast } from '../dates';

const DAY_MS = 24 * 60 * 60 * 1000;

export interface DateEntry {
  key: string;
  value: string;
  note?: string;
  iso: string;
  /** The last day of a span ("through Oct 21"); the entry is live until that day is over. */
  endIso?: string;
  /** Which entry the hero card follows: "beta" (the span the card says is live) or "launch". */
  role?: string;
  /** How the row reads once the span is over: "Sept 17 to Oct 21", "Beta", its note. */
  pastKey?: string;
  pastValue?: string;
  pastNote?: string;
}

export type TimelineStatus = 'past' | 'live' | 'next' | 'future';

export interface TimelineRow extends DateEntry {
  status: TimelineStatus;
}

function statusOf(entry: DateEntry, now: Date): 'past' | 'live' | 'upcoming' {
  const start = new Date(entry.iso);
  if (entry.endIso !== undefined) {
    if (isPast(new Date(new Date(entry.endIso).getTime() + DAY_MS), now)) return 'past';
    return isPast(start, now) ? 'live' : 'upcoming';
  }
  return isPast(start, now) ? 'past' : 'upcoming';
}

/** Every entry with its status, in the order the data file lists them (chronological). */
export function timelineRows(dates: readonly DateEntry[], now: Date = new Date()): TimelineRow[] {
  let nextTaken = false;
  return dates.map((entry) => {
    const state = statusOf(entry, now);
    if (state !== 'upcoming') return { ...entry, status: state };
    if (nextTaken) return { ...entry, status: 'future' };
    nextTaken = true;
    return { ...entry, status: 'next' };
  });
}

const DISPLAY_ORDER: Record<TimelineStatus, number> = { live: 0, next: 1, future: 2, past: 3 };

/** The rows in the order they are shown: live, next, future, then past; stable inside each. */
export function timelineDisplayRows(rows: readonly TimelineRow[]): TimelineRow[] {
  return rows
    .map((row, index) => ({ row, index }))
    .sort((a, b) => DISPLAY_ORDER[a.row.status] - DISPLAY_ORDER[b.row.status] || a.index - b.index)
    .map(({ row }) => row);
}

/** What a row reads: a finished span uses its past-tense wording. */
export function timelineRowText(row: TimelineRow): { key: string; value: string; note?: string } {
  if (row.status === 'past' && row.pastKey !== undefined) {
    return { key: row.pastKey, value: row.pastValue ?? row.value, note: row.pastNote };
  }
  return { key: row.key, value: row.value, note: row.note };
}
