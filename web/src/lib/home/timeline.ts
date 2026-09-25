// The homepage's horizontal dates row (spec 2026-09-24 §2.2), replacing the "Right now"
// StatePanel: every date before `now` is muted past, the first date at or after `now`
// carries the gold dot, and everything after that is a plain future date. Reuses
// lib/dates.ts's own `isPast` (strict less-than) rather than a second date comparison, so
// "is this the next one" can never disagree with any other page's notion of past/future.
import { isPast } from '../dates';

export interface DateEntry {
  key: string;
  value: string;
  note?: string;
  iso: string;
}

export type TimelineStatus = 'past' | 'next' | 'future';

export interface TimelineRow extends DateEntry {
  status: TimelineStatus;
}

export function timelineRows(dates: readonly DateEntry[], now: Date = new Date()): TimelineRow[] {
  const nextIndex = dates.findIndex((entry) => !isPast(new Date(entry.iso), now));
  return dates.map((entry, index) => ({
    ...entry,
    status: index === nextIndex ? 'next' : index < nextIndex || nextIndex === -1 ? 'past' : 'future',
  }));
}
