const MONTHS = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'June', 'July', 'Aug', 'Sept', 'Oct', 'Nov', 'Dec'];

/** "Sept 25" in UTC; an invalid date renders as nothing rather than "undefined NaN". */
export function formatDate(d: Date): string {
  if (Number.isNaN(d.getTime())) return '';
  return `${MONTHS[d.getUTCMonth()]} ${d.getUTCDate()}`;
}

export function isPast(d: Date, now: Date = new Date()): boolean {
  return d.getTime() < now.getTime();
}

/** The most recent of a set of dates -- an index page is "updated" when its newest entry is. */
export function latest(dates: Date[]): Date {
  if (dates.length === 0) throw new Error('latest() needs at least one date');
  return dates.reduce((newest, d) => (d.getTime() > newest.getTime() ? d : newest));
}

const MINUTE_MS = 60_000;
const HOUR_MS = 60 * MINUTE_MS;
const DAY_MS = 24 * HOUR_MS;
const RELATIVE_TIME_CEILING_DAYS = 30;

/**
 * "3 hours ago", falling back to `formatDate` past 30 days -- used for "Imported from
 * Battle.net <relative time>" (spec 2026-09-22 §7.2). `elapsed` is floored at zero so a
 * clock skew or a server timestamp a few seconds ahead of the caller's own clock reads as
 * "just now" rather than a negative duration.
 */
export function relativeTime(d: Date, now: Date = new Date()): string {
  const elapsed = Math.max(0, now.getTime() - d.getTime());
  if (elapsed < MINUTE_MS) return 'just now';
  if (elapsed < HOUR_MS) {
    const minutes = Math.floor(elapsed / MINUTE_MS);
    return `${minutes} minute${minutes === 1 ? '' : 's'} ago`;
  }
  if (elapsed < DAY_MS) {
    const hours = Math.floor(elapsed / HOUR_MS);
    return `${hours} hour${hours === 1 ? '' : 's'} ago`;
  }
  const days = Math.floor(elapsed / DAY_MS);
  if (days < RELATIVE_TIME_CEILING_DAYS) return `${days} day${days === 1 ? '' : 's'} ago`;
  return formatDate(d);
}

const RELATIVE_DAY_CEILING_DAYS = 14;
const SHORT_MONTHS = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];

/**
 * The one age rule for dated rows (Logs, the character header): whole days relative inside two
 * weeks ("today", "1 day ago", "3 days ago"), else the date ("Sep 15", UTC). Like `relativeTime`,
 * a timestamp ahead of the caller's clock reads as "today", never a negative duration.
 */
export function relativeDay(d: Date, now: Date = new Date()): string {
  if (Number.isNaN(d.getTime())) return '';
  const days = Math.floor(Math.max(0, now.getTime() - d.getTime()) / DAY_MS);
  if (days === 0) return 'today';
  if (days < RELATIVE_DAY_CEILING_DAYS) return `${days} day${days === 1 ? '' : 's'} ago`;
  return `${SHORT_MONTHS[d.getUTCMonth()]} ${d.getUTCDate()}`;
}
