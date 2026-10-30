// The homepage's dates line (polish spec 2026-10-28 §5.6, owner 2026-10-29), which replaced the
// horizontal dates strip (spec 2026-09-24 §2.2). Every status is decided by the clock, never by
// hand, so the line stays true without an edit: a date before `now` is past, a span (a start and an `endIso`) is live
// from its start through the end of its last day, the first date still to come is the next
// one, and everything after that is future. The line shows only what is still to come, in date
// order; a date drops off when its day starts (the changelog keeps the history).
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
  /** What the dates line says while this entry is running ("Name reservation open through Nov 3"). */
  running?: string;
  /** The clock time shown beside the date on the home dates line ("15:00 PST"). */
  time?: string;
  /** The exact moment (ISO instant) the hero's countdown runs to; the day's start when absent. */
  at?: string;
}

export type TimelineStatus = 'past' | 'live' | 'next' | 'future';

export interface TimelineRow extends DateEntry {
  status: TimelineStatus;
}

/** A date is live from its first day through the end of its last (a single day ends the day it starts). */
function statusOf(entry: DateEntry, now: Date): 'past' | 'live' | 'upcoming' {
  const start = new Date(entry.iso);
  const lastDay = new Date(entry.endIso ?? entry.iso);
  if (isPast(new Date(lastDay.getTime() + DAY_MS), now)) return 'past';
  return isPast(start, now) ? 'live' : 'upcoming';
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

/** One event of the dates line: its text, and whether it is the next one (the gold one). */
export interface DatesLineItem {
  text: string;
  next: boolean;
}

/**
 * Still to show: not yet started, or running (a span, or a day that has begun) until it ends. An
 * entry with an exact `at` moment (the launch hour) goes at that moment, not at the day's end.
 */
function isShown(row: TimelineRow, now: Date): boolean {
  if (row.status === 'past') return false;
  if (row.status === 'live' && row.at !== undefined) return !isPast(new Date(row.at), now);
  return true;
}

/** An entry's words on the line: while it runs its `running` phrase when it has one, else "{value} {key}" and its `time`. */
function lineText(row: TimelineRow): string {
  if (row.status === 'live' && row.running !== undefined) return row.running;
  return [`${row.value} ${row.key}`, row.time].filter((part) => part !== undefined).join(' · ');
}

/**
 * The events the dates line shows, in the order the data file lists them: what is running now
 * (a span like the name reservation, "Name reservation open through Nov 3") until it ends, then
 * what is still to come ("Launch Nov 4 · 15:00 PST"). The first is the one the line sets in gold.
 * Empty once every date has passed, which is when the line goes.
 */
export function datesLineItems(rows: readonly TimelineRow[], now: Date = new Date()): DatesLineItem[] {
  return rows
    .filter((row) => isShown(row, now))
    .map((row, index) => ({ text: lineText(row), next: index === 0 }));
}
