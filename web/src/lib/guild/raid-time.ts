// web/src/lib/guild/raid-time.ts
// Wall-clock and zone arithmetic for guild raids (design/specs/2026-10-26-guild-raids.md 3.5, R4).
// The form takes the guild's wall clock and the API stores UTC instants, so every conversion
// here goes through `Intl.DateTimeFormat` alone: no date library, no hand-kept zone table.

const MINUTE_MS = 60_000;
/** The words (weekdays, months, am/pm) are English everywhere; only the hour cycle follows the viewer. */
export const DEFAULT_LOCALE = 'en-US';

export type HourCycle = 'h12' | 'h23';

type PartName = Intl.DateTimeFormatPartTypes;

function partsOf(formatter: Intl.DateTimeFormat, d: Date): Partial<Record<PartName, string>> {
  const parts: Partial<Record<PartName, string>> = {};
  for (const part of formatter.formatToParts(d)) parts[part.type] = part.value;
  return parts;
}

const zoneFormatters = new Map<string, Intl.DateTimeFormat>();

/** One formatter per zone: every conversion reads through it, so building one each call adds up. */
function zoneFormatter(tz: string): Intl.DateTimeFormat {
  const cached = zoneFormatters.get(tz);
  if (cached !== undefined) return cached;
  const formatter = new Intl.DateTimeFormat(DEFAULT_LOCALE, {
    timeZone: tz,
    hourCycle: 'h23',
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit',
  });
  zoneFormatters.set(tz, formatter);
  return formatter;
}

/** `CST`, `CDT`, `GMT+1`: the zone's short name at that instant; the zone id when it has none. */
export function zoneAbbreviation(tz: string, at: Date): string {
  return (
    new Intl.DateTimeFormat(DEFAULT_LOCALE, { timeZone: tz, timeZoneName: 'short' })
      .formatToParts(at)
      .find((part) => part.type === 'timeZoneName')?.value ?? tz
  );
}

function browserLanguage(): string | undefined {
  return typeof navigator === 'undefined' ? undefined : navigator.language;
}

/** The zone's wall clock at an instant, read back as if it were UTC. */
function wallClockAsUtcMs(d: Date, tz: string): number {
  const p = partsOf(zoneFormatter(tz), d);
  return Date.UTC(
    Number(p.year),
    Number(p.month) - 1,
    Number(p.day),
    Number(p.hour),
    Number(p.minute),
    Number(p.second),
  );
}

/** Minutes the zone is ahead of UTC at that instant (Chicago in winter: -360). */
export function zoneOffsetMinutes(d: Date, tz: string): number {
  const wholeSecondMs = Math.floor(d.getTime() / 1000) * 1000;
  return (wallClockAsUtcMs(d, tz) - wholeSecondMs) / MINUTE_MS;
}

export function offsetsDiffer(d: Date, tzA: string, tzB: string): boolean {
  return zoneOffsetMinutes(d, tzA) !== zoneOffsetMinutes(d, tzB);
}

/** `2027-03-14` and `20:00` on a zone's clock, as the instant they name. Fails fast on a malformed
 *  input or an unknown zone; a time the zone skips (spring forward) lands just after the gap. */
export function wallClockToInstant(date: string, time: string, tz: string): Date {
  const day = /^(\d{4})-(\d{2})-(\d{2})$/.exec(date);
  const clock = /^(\d{2}):(\d{2})$/.exec(time);
  if (day === null || clock === null) throw new Error(`bad wall clock "${date} ${time}"`);
  const asUtc = Date.UTC(
    Number(day[1]),
    Number(day[2]) - 1,
    Number(day[3]),
    Number(clock[1]),
    Number(clock[2]),
  );
  const firstOffset = zoneOffsetMinutes(new Date(asUtc), tz);
  const first = asUtc - firstOffset * MINUTE_MS;
  const secondOffset = zoneOffsetMinutes(new Date(first), tz);
  if (secondOffset === firstOffset) return new Date(first);
  const corrected = asUtc - secondOffset * MINUTE_MS;
  // A corrected instant that does not read back as the input fell in the skipped hour.
  return new Date(wallClockAsUtcMs(new Date(corrected), tz) === asUtc ? corrected : first);
}

/** An instant on a zone's clock, as the form's `date` and `time` field values. */
export function instantToWallClock(d: Date, tz: string): { date: string; time: string } {
  const p = partsOf(zoneFormatter(tz), d);
  return { date: `${p.year}-${p.month}-${p.day}`, time: `${p.hour}:${p.minute}` };
}

/** Whether the locale writes the hour on a 12-hour clock. */
export function usesTwelveHourClock(locale: string): boolean {
  return new Intl.DateTimeFormat(locale, { hour: 'numeric' }).resolvedOptions().hour12 === true;
}

/** The viewer's hour cycle, from their browser language (`de-DE` writes `20:00`). */
export function viewerHourCycle(language: string | undefined = browserLanguage()): HourCycle {
  return language !== undefined && !usesTwelveHourClock(language) ? 'h23' : 'h12';
}

/** The locale every raid clock is written in: English words on the viewer's hour cycle
 *  (`en-US-u-hc-h23` for a German browser), `en-US` on the server. */
export function viewerLocale(language: string | undefined = browserLanguage()): string {
  if (language === undefined) return DEFAULT_LOCALE;
  return viewerHourCycle(language) === 'h23' ? `${DEFAULT_LOCALE}-u-hc-h23` : DEFAULT_LOCALE;
}

/** `8 pm`, `8:30 pm`, `20:00`: minutes only off the hour unless `alwaysMinutes`. */
export function clockText(d: Date, tz: string, locale: string, alwaysMinutes = false): string {
  const twelve = usesTwelveHourClock(locale);
  const p = partsOf(
    new Intl.DateTimeFormat(locale, {
      timeZone: tz,
      hour: 'numeric',
      minute: '2-digit',
      hourCycle: twelve ? 'h12' : 'h23',
    }),
    d,
  );
  if (!twelve) return `${(p.hour ?? '').padStart(2, '0')}:${p.minute}`;
  const minutes = alwaysMinutes || p.minute !== '00' ? `:${p.minute}` : '';
  return `${p.hour}${minutes} ${(p.dayPeriod ?? '').toLowerCase()}`;
}

/** `Thu 10 Dec, 11:42 pm` in the viewer's zone, and ` (10:42 pm guild time)` apart (null) when the
 *  guild's offset is the same, so a caller can set the guild time in its own muted span. */
export function stampParts(
  at: Date,
  viewerTz: string,
  guildTz: string,
  locale: string = DEFAULT_LOCALE,
): { when: string; guildTime: string | null } {
  const when = `${dateLabel(at, viewerTz, locale)}, ${clockText(at, viewerTz, locale, true)}`;
  if (!offsetsDiffer(at, viewerTz, guildTz)) return { when, guildTime: null };
  return { when, guildTime: `(${clockText(at, guildTz, locale, true)} guild time)` };
}

/** `Thu 10 Dec`, on the zone's calendar. */
export function dateLabel(d: Date, tz: string, locale: string = DEFAULT_LOCALE): string {
  const p = partsOf(
    new Intl.DateTimeFormat(locale, { timeZone: tz, weekday: 'short', day: 'numeric', month: 'short' }),
    d,
  );
  return `${p.weekday} ${p.day} ${p.month}`;
}

/** The form's default instant: the last raid's weekday and guild wall time, a whole number of weeks
 *  after it, the first strictly after `now`. Steps the calendar date in `tz`, so daylight saving
 *  keeps the wall time (R4). */
export function nextWeekday(last: Date, now: Date, tz: string): Date {
  const { date, time } = instantToWallClock(last, tz);
  const [year, month, day] = date.split('-').map(Number);
  for (let weeks = 1; ; weeks += 1) {
    const stepped = new Date(Date.UTC(year ?? 0, (month ?? 1) - 1, (day ?? 1) + weeks * 7));
    const candidate = wallClockToInstant(stepped.toISOString().slice(0, 10), time, tz);
    if (candidate.getTime() > now.getTime()) return candidate;
  }
}
