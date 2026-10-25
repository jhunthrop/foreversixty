// web/src/lib/sim/server-runs.ts
// The view model under "Your server runs": which of seven states a row is in, the plain words
// for an estimate and an age, and which rows the list shows. Pure data in, data out; every
// sentence is in server-runs-copy.ts, which builds on the words here.
import { relativeDay, relativeTime } from '../dates';
import type { SimListRow } from './types';

export type ServerRunState =
  | 'queued'
  | 'running'
  | 'running-overdue'
  | 'running-no-estimate'
  | 'done'
  | 'done-unseen'
  | 'error';

export const MAX_VISIBLE_RUNS = 5;

/** A run is overdue once it has taken more than this many times its estimate. */
export const OVERDUE_FACTOR = 1.5;

const SECONDS_PER_MINUTE = 60;
const MINUTES_PER_HOUR = 60;
/** From 90 minutes the estimate reads in hours. */
const HOURS_FROM_MINUTES = 90;
const RELATIVE_AGE_CEILING_MS = 14 * 24 * 60 * 60 * 1000;

/** "about a minute", "about 3 minutes", "about 2 hours": the estimate rounded to the nearest unit. */
export function estimateWords(seconds: number): string {
  const minutes = Math.round(seconds / SECONDS_PER_MINUTE);
  if (minutes <= 1) return 'about a minute';
  if (minutes < HOURS_FROM_MINUTES) return `about ${minutes} minutes`;
  return `about ${Math.round(minutes / MINUTES_PER_HOUR)} hours`;
}

export function isOverdue(estimateSeconds: number, elapsedSecondsValue: number): boolean {
  return elapsedSecondsValue > OVERDUE_FACTOR * estimateSeconds;
}

/** Whole seconds since the run was created, never negative. */
export function elapsedSeconds(row: SimListRow, now: Date): number {
  const elapsedMs = now.getTime() - new Date(row.created_at).getTime();
  return Math.max(0, Math.floor(elapsedMs / 1000));
}

function runningState(row: SimListRow, now: Date): ServerRunState {
  if (row.estimate_seconds === undefined) return 'running-no-estimate';
  return isOverdue(row.estimate_seconds, elapsedSeconds(row, now)) ? 'running-overdue' : 'running';
}

export function rowState(row: SimListRow, now: Date): ServerRunState {
  switch (row.state) {
    case 'queued':
      return 'queued';
    case 'running':
      return runningState(row, now);
    case 'error':
      return 'error';
    default:
      return row.seen === false ? 'done-unseen' : 'done';
  }
}

/** "14 minutes ago" inside 14 days, else a short date such as "Nov 25". */
export function ageWords(date: Date, now: Date): string {
  const withinRelative = now.getTime() - date.getTime() < RELATIVE_AGE_CEILING_MS;
  return withinRelative ? relativeTime(date, now) : relativeDay(date, now);
}

export interface VisibleRuns {
  rows: SimListRow[];
  hasMore: boolean;
}

/** The newest five by creation time, and whether older rows were left out. */
export function visibleRows(rows: readonly SimListRow[]): VisibleRuns {
  const newestFirst = [...rows].sort((a, b) => b.created_at.localeCompare(a.created_at));
  return { rows: newestFirst.slice(0, MAX_VISIBLE_RUNS), hasMore: newestFirst.length > MAX_VISIBLE_RUNS };
}
