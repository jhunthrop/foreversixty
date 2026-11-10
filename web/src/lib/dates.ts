// web/src/lib/dates.ts
// The house date and clock (owner rulings 2026-11-10): every date, weekday, month heading, clock
// time and "Updated" stamp the site prints comes from here. Months are three letters ("Sep",
// never "Sept"), the day comes first, and the clock follows the viewer's locale ("8 pm" for a US
// browser, "20:00" for a European one) while the words stay English. A page rendered at build
// time, or on the server, cannot know the viewer and prints the site's own locale. The canonical
// cases live in apps/shared/formats.json, which the server pages, the companion and both apps
// replay too.

import { DAY_MS, HOUR_MS, MINUTE_MS } from './format/duration';
import { SITE_LOCALE } from './format/number';

/** The one month table: three letters, every month. */
export const MONTHS = [
  'Jan',
  'Feb',
  'Mar',
  'Apr',
  'May',
  'Jun',
  'Jul',
  'Aug',
  'Sep',
  'Oct',
  'Nov',
  'Dec',
] as const;
const SHORT_MONTHS = MONTHS;

/** "25 Sep" in UTC, day first (the house date, as every page head's "Updated 6 Nov"); an invalid
 *  date renders as nothing rather than "undefined NaN". */
export function formatDate(d: Date): string {
  if (Number.isNaN(d.getTime())) return '';
  return `${d.getUTCDate()} ${MONTHS[d.getUTCMonth()]}`;
}

/** "Updated 25 Sep": every page head's freshness stamp (UpdatedStamp.astro and its Svelte twin). */
export function updatedLabel(d: Date): string {
  const date = formatDate(d);
  return date === '' ? '' : `Updated ${date}`;
}

const FULL_MONTHS = [
  'January',
  'February',
  'March',
  'April',
  'May',
  'June',
  'July',
  'August',
  'September',
  'October',
  'November',
  'December',
];

/** "November 2026" in UTC: a dated log's month heading (the changelog). */
export function monthHeading(d: Date): string {
  if (Number.isNaN(d.getTime())) return '';
  return `${FULL_MONTHS[d.getUTCMonth()]} ${d.getUTCFullYear()}`;
}

export function isPast(d: Date, now: Date = new Date()): boolean {
  return d.getTime() < now.getTime();
}

/** The most recent of a set of dates -- an index page is "updated" when its newest entry is. */
export function latest(dates: Date[]): Date {
  if (dates.length === 0) throw new Error('latest() needs at least one date');
  return dates.reduce((newest, d) => (d.getTime() > newest.getTime() ? d : newest));
}

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

/**
 * The compact age of a freshness stamp: "just now", "5 min ago", "2 h ago", "3 d ago", then the
 * date past 30 days ("6 Nov"). Like `relativeTime`, a timestamp ahead of the clock is "just now".
 */
export function compactAgo(d: Date, now: Date = new Date()): string {
  if (Number.isNaN(d.getTime())) return '';
  const elapsed = Math.max(0, now.getTime() - d.getTime());
  if (elapsed < MINUTE_MS) return 'just now';
  if (elapsed < HOUR_MS) return `${Math.floor(elapsed / MINUTE_MS)} min ago`;
  if (elapsed < DAY_MS) return `${Math.floor(elapsed / HOUR_MS)} h ago`;
  const days = Math.floor(elapsed / DAY_MS);
  if (days < RELATIVE_TIME_CEILING_DAYS) return `${days} d ago`;
  return formatDate(d);
}

const RELATIVE_DAY_CEILING_DAYS = 14;

/**
 * The one age rule for dated rows (Logs, the character header): whole days relative inside two
 * weeks ("today", "1 day ago", "3 days ago"), else the date ("15 Sep", UTC). Like `relativeTime`,
 * a timestamp ahead of the caller's clock reads as "today", never a negative duration.
 */
export function relativeDay(d: Date, now: Date = new Date()): string {
  if (Number.isNaN(d.getTime())) return '';
  const days = Math.floor(Math.max(0, now.getTime() - d.getTime()) / DAY_MS);
  if (days === 0) return 'today';
  if (days < RELATIVE_DAY_CEILING_DAYS) return `${days} day${days === 1 ? '' : 's'} ago`;
  return formatDate(d);
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

/** "19 Oct, 9:40 pm" ("19 Oct, 21:40" for a 24-hour locale), on the zone's calendar and clock:
 *  when a sync from the game arrived. An unreadable date renders as nothing. */
export function formatSyncStamp(
  d: Date,
  zone: string = viewerZone(),
  locale: string = viewerLocale(),
): string {
  if (Number.isNaN(d.getTime())) return '';
  return `${dayMonth(d, zone)}, ${clock(d, zone, { locale, alwaysMinutes: true })}`;
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

/** Inside this many days a sync reads as words (`3 days ago`); past it, the date. */
const SYNCED_WORDS_DAYS = 7;
/** Past this many days the date carries its year: a late-December sync read in January stays short. */
const SYNCED_YEAR_DAYS = 182;

/**
 * The readiness board's Last synced word, one format for every row (guild loop round 3: `today`
 * beside `27 Oct` beside a weekday): `today`, `yesterday`, `3 days ago` up to 7 days, then
 * `27 Oct` (`27 Oct 2025` once it is over six months old, so the pill stays one short line). Days are the viewer's calendar days. An
 * unreadable date renders as nothing.
 */
export function syncedWord(d: Date, now: Date = new Date()): string {
  if (Number.isNaN(d.getTime())) return '';
  const age =
    dayNumber(now.getFullYear(), now.getMonth(), now.getDate()) -
    dayNumber(d.getFullYear(), d.getMonth(), d.getDate());
  if (age <= 0) return 'today';
  if (age === 1) return 'yesterday';
  if (age <= SYNCED_WORDS_DAYS) return `${age} days ago`;
  const base = `${d.getDate()} ${SHORT_MONTHS[d.getMonth()]}`;
  return age <= SYNCED_YEAR_DAYS ? base : `${base} ${d.getFullYear()}`;
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

/** "31 Dec", always read in UTC, for a stored date-only value or an RFC 3339 instant: the short
 *  sibling of `utcDayMonthYear` for a date that cannot fall in another year. */
export function utcDayMonth(value: string): string {
  return utcDayMonthYear(value).replace(/ \d{4}$/, '');
}

/** The viewer's own IANA zone, the default for the zoned day helpers. */
export function viewerZone(): string {
  return Intl.DateTimeFormat().resolvedOptions().timeZone;
}

function zonedParts(instant: string, zone: string): { day: number; month: string; year: string } | null {
  const date = new Date(instant);
  if (Number.isNaN(date.getTime())) return null;
  const at = zoned(date, zone);
  return { day: at.day, month: MONTHS[at.month], year: String(at.year) };
}

const DATE_ONLY = /^\d{4}-\d{2}-\d{2}$/;

/**
 * "20 Dec 2026" for a stored value: a date-only value is a calendar day and reads as written,
 * an instant reads on the calendar of `zone` (the viewer's by default), so a seat that unlocks
 * at 20:00 UTC on the 20th is the 21st for a viewer in Auckland.
 */
export function zonedDayMonthYear(value: string, zone: string = viewerZone()): string {
  if (DATE_ONLY.test(value)) return utcDayMonthYear(value);
  const parts = zonedParts(value, zone);
  return parts === null ? '' : `${parts.day} ${parts.month} ${parts.year}`;
}

/** "20 Dec": the short sibling of `zonedDayMonthYear`, for a date that cannot fall in another year. */
export function zonedDayMonth(value: string, zone: string = viewerZone()): string {
  return zonedDayMonthYear(value, zone).replace(/ \d{4}$/, '');
}

const SHORT_WEEKDAYS = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'];

/** "Tue 20 Oct" ("Tue 20 Oct 2025" in another year) for a calendar day. */
function nightLabelFor(target: { y: number; m: number; d: number }, now: Date): string {
  const weekday = SHORT_WEEKDAYS[new Date(Date.UTC(target.y, target.m, target.d)).getUTCDay()];
  const base = `${weekday} ${target.d} ${SHORT_MONTHS[target.m]}`;
  return target.y === now.getFullYear() ? base : `${base} ${target.y}`;
}

/**
 * A raid night's date for a moment, in the viewer's own zone: "Tue 20 Oct". It never says "today"
 * or "yesterday", so the same night reads the same on the account's report row, the guild
 * Overview and the Raids tab whenever each is opened (P.1, player review S2-4). An unreadable date
 * renders as nothing.
 */
export function nightDate(d: Date, now: Date = new Date()): string {
  if (Number.isNaN(d.getTime())) return '';
  return nightLabelFor({ y: d.getFullYear(), m: d.getMonth(), d: d.getDate() }, now);
}

/** The same label for a date-only value ("2026-10-20"), read as written. */
export function nightDateOnly(value: string, now: Date = new Date()): string {
  const match = /^(\d{4})-(\d{2})-(\d{2})/.exec(value);
  if (match === null) return '';
  return nightLabelFor({ y: Number(match[1]), m: Number(match[2]) - 1, d: Number(match[3]) }, now);
}

// --- The clock and the zoned calendar ------------------------------------------------------

function browserLanguage(): string | undefined {
  return typeof navigator === 'undefined' ? undefined : navigator.language;
}

const hourCycles = new Map<string, boolean>();

/** Whether the locale writes the hour on a 12-hour clock (`en-US` yes, `en-GB` and `de-DE` no). */
export function usesTwelveHourClock(locale: string): boolean {
  let twelve = hourCycles.get(locale);
  if (twelve === undefined) {
    twelve = new Intl.DateTimeFormat(locale, { hour: 'numeric' }).resolvedOptions().hour12 === true;
    hourCycles.set(locale, twelve);
  }
  return twelve;
}

/** The locale the viewer's clock is written in: their browser's language, the site's own on the
 *  server and at build time (formats.json `site_locale`). Only the hour cycle reads from it. */
export function viewerLocale(language: string | undefined = browserLanguage()): string {
  return language ?? SITE_LOCALE;
}

interface ZonedMoment {
  year: number;
  /** 0-based, as `MONTHS` indexes it. */
  month: number;
  day: number;
  weekday: number;
  hour: number;
  minute: number;
}

const zonedFormats = new Map<string, Intl.DateTimeFormat>();

/** One formatter per zone: the calendar and the 24-hour wall clock of an instant in that zone. */
function zonedFormat(zone: string): Intl.DateTimeFormat {
  let format = zonedFormats.get(zone);
  if (format === undefined) {
    format = new Intl.DateTimeFormat(SITE_LOCALE, {
      timeZone: zone,
      hourCycle: 'h23',
      year: 'numeric',
      month: 'numeric',
      day: 'numeric',
      weekday: 'short',
      hour: 'numeric',
      minute: '2-digit',
    });
    zonedFormats.set(zone, format);
  }
  return format;
}

/** The instant on the zone's calendar and clock. */
function zoned(d: Date, zone: string): ZonedMoment {
  const parts: Partial<Record<Intl.DateTimeFormatPartTypes, string>> = {};
  for (const part of zonedFormat(zone).formatToParts(d)) parts[part.type] = part.value;
  return {
    year: Number(parts.year),
    month: Number(parts.month) - 1,
    day: Number(parts.day),
    weekday: SHORT_WEEKDAYS.indexOf(parts.weekday ?? ''),
    hour: Number(parts.hour) % 24,
    minute: Number(parts.minute),
  };
}

/** The zone's calendar day as a whole-day number, so two days subtract to whole days. */
export function calendarDayIn(d: Date, zone: string): number {
  const at = zoned(d, zone);
  return dayNumber(at.year, at.month, at.day);
}

/** Options for `clock`: the viewer's locale by default, and minutes even on the hour. */
export interface ClockOptions {
  locale?: string;
  /** `8:00 pm` rather than `8 pm`: a stamp, where the minute is part of the record. */
  alwaysMinutes?: boolean;
}

/**
 * The time of day on the zone's clock, in the viewer's hour cycle (owner ruling 2026-11-10):
 * `8 pm`, `8:30 pm` (minutes only off the hour unless `alwaysMinutes`) for a 12-hour locale,
 * `20:00` for a 24-hour one. An unreadable date renders as nothing.
 */
export function clock(d: Date, zone: string = viewerZone(), options: ClockOptions = {}): string {
  if (Number.isNaN(d.getTime())) return '';
  const { hour, minute } = zoned(d, zone);
  const mm = String(minute).padStart(2, '0');
  if (!usesTwelveHourClock(options.locale ?? viewerLocale())) {
    return `${String(hour).padStart(2, '0')}:${mm}`;
  }
  const twelve = hour % 12 === 0 ? 12 : hour % 12;
  const minutes = options.alwaysMinutes === true || minute !== 0 ? `:${mm}` : '';
  return `${twelve}${minutes} ${hour < 12 ? 'am' : 'pm'}`;
}

/** "9 Dec" on the zone's calendar. An unreadable date renders as nothing. */
export function dayMonth(d: Date, zone: string = viewerZone()): string {
  if (Number.isNaN(d.getTime())) return '';
  const at = zoned(d, zone);
  return `${at.day} ${MONTHS[at.month]}`;
}

/** "Wed 9 Dec" on the zone's calendar ("Wed 6 Jan 2027" outside `now`'s year there). */
export function weekdayDayMonth(d: Date, zone: string = viewerZone(), now: Date = new Date()): string {
  if (Number.isNaN(d.getTime())) return '';
  const at = zoned(d, zone);
  const base = `${SHORT_WEEKDAYS[at.weekday]} ${at.day} ${MONTHS[at.month]}`;
  return at.year === zoned(now, zone).year ? base : `${base} ${at.year}`;
}

/** "Wednesday 9 December" on the zone's calendar: a headline day (the first-raids plaque). */
export function longDay(d: Date, zone: string = viewerZone()): string {
  if (Number.isNaN(d.getTime())) return '';
  const at = zoned(d, zone);
  return `${WEEKDAYS[at.weekday]} ${at.day} ${FULL_MONTHS[at.month]}`;
}

/** "Wednesday" on the zone's calendar. */
export function weekdayName(d: Date, zone: string = viewerZone()): string {
  if (Number.isNaN(d.getTime())) return '';
  return WEEKDAYS[zoned(d, zone).weekday];
}

/** "Wed" on the zone's calendar. */
export function weekdayShort(d: Date, zone: string = viewerZone()): string {
  if (Number.isNaN(d.getTime())) return '';
  return SHORT_WEEKDAYS[zoned(d, zone).weekday];
}

export { DAY_MS, HOUR_MS, MINUTE_MS };
