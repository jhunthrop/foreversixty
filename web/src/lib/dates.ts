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

/**
 * "9 Dec 2026", in UTC: the Forever Score pages' table date (fix round 2 item 6). Day first,
 * the same month names the addon's "as of 4 Nov" uses, so a table and the card read alike.
 * An unreadable date renders as nothing.
 */
export function formatDayMonthYear(d: Date): string {
  if (Number.isNaN(d.getTime())) return '';
  return `${d.getUTCDate()} ${SHORT_MONTHS[d.getUTCMonth()]} ${d.getUTCFullYear()}`;
}

/** "19 Oct, 21:40", in the viewer's own time zone: when a sync from the game arrived. An
 *  unreadable date renders as nothing. */
export function formatSyncStamp(d: Date): string {
  if (Number.isNaN(d.getTime())) return '';
  const pad = (value: number): string => String(value).padStart(2, '0');
  return `${d.getDate()} ${SHORT_MONTHS[d.getMonth()]}, ${pad(d.getHours())}:${pad(d.getMinutes())}`;
}

const WEEKDAYS = ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday'];
const WEEKDAY_WINDOW_DAYS = 6;

/** The calendar day as a day number, so two days subtract to whole days with no clock or DST in it. */
function dayNumber(year: number, month: number, day: number): number {
  return Math.floor(Date.UTC(year, month, day) / DAY_MS);
}

/** "today", "yesterday", a weekday inside 6 days, else "20 Oct" ("20 Oct 2025" in another year). */
function dayWordFor(target: { y: number; m: number; d: number }, now: Date): string {
  const age =
    dayNumber(now.getFullYear(), now.getMonth(), now.getDate()) - dayNumber(target.y, target.m, target.d);
  if (age <= 0) return 'today';
  if (age === 1) return 'yesterday';
  if (age <= WEEKDAY_WINDOW_DAYS)
    return WEEKDAYS[new Date(Date.UTC(target.y, target.m, target.d)).getUTCDay()];
  const base = `${target.d} ${SHORT_MONTHS[target.m]}`;
  return target.y === now.getFullYear() ? base : `${base} ${target.y}`;
}

/**
 * The guild pages' day word for a moment ("Updated today", a raid night): the day is resolved in
 * the viewer's own zone, so 20:00 in Chicago on the 20th is "today" even where UTC has already
 * reached the 21st. An unreadable date renders as nothing.
 */
export function dayWord(d: Date, now: Date = new Date()): string {
  if (Number.isNaN(d.getTime())) return '';
  return dayWordFor({ y: d.getFullYear(), m: d.getMonth(), d: d.getDate() }, now);
}

/**
 * The same word for a date-only value ("2026-10-20"): a calendar date has no zone, so it is read
 * as written (never through `new Date`, which would move it by the viewer's offset) and only then
 * set against the viewer's today. Anything else renders as nothing.
 */
export function dateOnlyWord(value: string, now: Date = new Date()): string {
  const match = /^(\d{4})-(\d{2})-(\d{2})/.exec(value);
  if (match === null) return '';
  return dayWordFor({ y: Number(match[1]), m: Number(match[2]) - 1, d: Number(match[3]) }, now);
}

/** "31 Dec 2026" for a stored date-only value or an RFC 3339 instant, always read in UTC: a
 *  grant "until 2026-12-31" is the 31st for everyone, not the 30th west of Greenwich. */
export function utcDayMonthYear(value: string): string {
  const dateOnly = /^(\d{4})-(\d{2})-(\d{2})$/.exec(value);
  if (dateOnly !== null) {
    return `${Number(dateOnly[3])} ${SHORT_MONTHS[Number(dateOnly[2]) - 1]} ${dateOnly[1]}`;
  }
  return formatDayMonthYear(new Date(value));
}
