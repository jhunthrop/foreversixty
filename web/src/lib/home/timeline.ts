// The homepage's horizontal dates row (spec 2026-09-24 §2.2), replacing the "Right now"
// StatePanel. Every status is decided by the clock, never by hand, so the row stays true
// without an edit: a date before `now` is past, a span (a start and an `endIso`) is live
// from its start through the end of its last day, the first date still to come is the next
// one, and everything after that is future. The strip shows only what is live or still to come,
// in date order; a date drops off the day after it ends (the changelog keeps the history).
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

/** The rows the strip shows: everything not yet past, in the order the data file lists them. */
export function timelineDisplayRows(rows: readonly TimelineRow[]): TimelineRow[] {
  return rows.filter((row) => row.status !== 'past');
}

/**
 * Which of the shown rows carries the gold dot: the live one (the Beta tile while the hero card says
 * "Beta is live"), and only when nothing is live, the next one to come. Null for no rows.
 */
export function dotIndex(shown: readonly TimelineRow[]): number | null {
  const live = shown.findIndex((row) => row.status === 'live');
  if (live !== -1) return live;
  const next = shown.findIndex((row) => row.status === 'next');
  return next === -1 ? null : next;
}
