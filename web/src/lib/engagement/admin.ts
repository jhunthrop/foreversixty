// web/src/lib/engagement/admin.ts
// Every word and row on /admin/activity, the owner's view of who comes back. A figure from before
// the live record (the backfill) is a floor and shows as "≥ n".
import { utcDayMonth } from '../dates';
import { EMPTY_FIGURE } from '../format/number';
import { countLabel, shareOf } from '../golinks/funnel';
import type { GridCell, GridRowModel } from '../ui/cells';
import type { CohortWeek, EngagementCohort, EngagementPeriod, EngagementReport } from './api';

/** A headline figure over its label (ui/BigFigures draws it). */
export interface HeadlineFigure {
  value: string;
  label: string;
}

/** The retention table's columns: weeks after signup. */
export const COHORT_WEEKS = 7;

export const adminActivityCopy = {
  pageTitle: 'Activity',
  job: 'Who signs in and who comes back.',
  pageDescription: 'The owner’s daily and weekly active accounts and retention.',
  signedOut: 'Sign in with the owner’s account to see the activity.',
  notOwner: 'This page is for the site’s owner. Your account can’t see the activity.',
  loadFailed: 'The activity didn’t load.',
  nowTitle: 'Now',
  weeksTitle: 'Weeks',
  weeksCaption: 'Active accounts by week',
  daysTitle: 'Days',
  daysCaption: 'Active accounts by day',
  cohortsTitle: 'Retention by signup week',
  cohortsCaption: 'Share of each signup week active in each later week',
  periodColumns: ['Active', 'New', 'Returning', 'Site', 'Addon', 'App'],
  cohortHeader: 'Signed up',
  cohortColumns: Array.from({ length: COHORT_WEEKS }, (_, k) => `+${k + 1}\u00a0wk`),
  weekHeader: 'Week of',
  dayHeader: 'Day',
  joined: (n: number): string => `${countLabel(n)} joined`,
  thisWeek: 'so far',
  /** After a retention share whose week is still running. */
  runningMark: '*',
  cohortsMeta: '* this week, still running; – not reached yet',
  notYet: 'not reached yet',
  weeksMeta: 'Active counts an account once; Site, Addon and App overlap.',
  today: 'today',
  figures: {
    today: 'active today',
    week: 'active this week',
    fresh: 'new this week',
    returning: 'returning this week',
    lastWeek: 'of last week’s signups back so far',
  },
  internal: (n: number): string =>
    `Leaves out ${countLabel(n)} internal ${n === 1 ? 'account' : 'accounts'}: yours and every granted Premium.`,
  floor: (liveSince: string | null): string =>
    liveSince === null
      ? 'Every figure is rebuilt from history (sign-ins, sims, addon syncs) and is a floor: ≥ marks it.'
      : `Recorded live since ${utcDayMonth(liveSince)}. Before that the figures are rebuilt from history (sign-ins, sims, addon syncs) and are a floor: ≥ marks them.`,
  surfaces:
    'Site is a signed-in page, Addon the addon’s syncs through the companion, App the phone app. One account can be on several.',
} as const;

/** The floor mark, held to its figure by a no-break space so a narrow column never splits them. */
const FLOOR = '≥\u00a0';

/** A count, as a floor when partial. */
export function floorLabel(value: number, partial: boolean): string {
  return partial && value > 0 ? `${FLOOR}${countLabel(value)}` : countLabel(value);
}

function periodCells(p: EngagementPeriod): GridCell[] {
  return [p.active, p.new, p.returning, p.site, p.addon, p.app].map((n) => ({
    tone: 'none',
    text: floorLabel(n, p.partial),
  }));
}

/** The weeks Grid: newest first, this week marked as running. */
export function weekRows(weeks: readonly EngagementPeriod[]): GridRowModel[] {
  return weeks.map((w, i) => ({
    key: w.start,
    label: utcDayMonth(w.start),
    sub: i === 0 ? adminActivityCopy.thisWeek : undefined,
    cells: periodCells(w),
  }));
}

/** The days Grid: newest first, today named. */
export function dayRows(days: readonly EngagementPeriod[]): GridRowModel[] {
  return days.map((d, i) => ({
    key: d.start,
    label: utcDayMonth(d.start),
    sub: i === 0 ? adminActivityCopy.today : undefined,
    cells: periodCells(d),
  }));
}

function cohortCell(week: CohortWeek | undefined, joined: number, running: boolean): GridCell {
  if (week === undefined || joined === 0)
    return { tone: 'none', text: EMPTY_FIGURE, label: adminActivityCopy.notYet };
  const share = shareOf(week.active, joined);
  const floored = week.partial && week.active > 0 ? `${FLOOR}${share}` : share;
  const text = running ? `${floored}${adminActivityCopy.runningMark}` : floored;
  const label = `${week.active} of ${joined}, ${floored}${running ? `, ${adminActivityCopy.thisWeek}` : ''}`;
  return { tone: 'none', text, label };
}

/** The retention Grid: newest signup week first, one cell per week after it; a cohort's last
 *  cell is this week, still running, and carries the running mark. */
export function cohortRows(cohorts: readonly EngagementCohort[]): GridRowModel[] {
  return [...cohorts].reverse().map((c) => ({
    key: c.week,
    label: utcDayMonth(c.week),
    sub: adminActivityCopy.joined(c.joined),
    cells: Array.from({ length: COHORT_WEEKS }, (_, k) =>
      cohortCell(c.weeks[k], c.joined, k === c.weeks.length - 1),
    ),
  }));
}

/** Last week's signups who came back this week, the headline retention. */
function lastWeekBack(cohorts: readonly EngagementCohort[]): string {
  const last = cohorts.at(-2);
  if (last === undefined || last.joined === 0 || last.weeks.length === 0) return shareOf(0, 0);
  return shareOf(last.weeks[0].active, last.joined);
}

/** The headline figures: today, this week split new and returning, last week's signups back. */
export function headlineFigures(report: EngagementReport): HeadlineFigure[][] {
  const copy = adminActivityCopy.figures;
  const today = report.days[0];
  const week = report.weeks[0];
  if (today === undefined || week === undefined) return [];
  return [
    [
      { value: floorLabel(today.active, today.partial), label: copy.today },
      { value: floorLabel(week.active, week.partial), label: copy.week },
    ],
    [
      { value: floorLabel(week.new, week.partial), label: copy.fresh },
      { value: floorLabel(week.returning, week.partial), label: copy.returning },
    ],
    [{ value: lastWeekBack(report.cohorts), label: copy.lastWeek }],
  ];
}
