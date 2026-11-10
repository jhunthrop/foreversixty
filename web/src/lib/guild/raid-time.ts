// web/src/lib/guild/raid-time.ts
// Wall-clock and zone arithmetic for guild raids (design/specs/2026-10-26-guild-raids.md 3.5, R4).
// The form takes the guild's wall clock and the API stores UTC instants, so every conversion
// here goes through `Intl.DateTimeFormat` alone: no date library, no hand-kept zone table. The
// words (dates, clocks) come from lib/dates.ts; this module only does the zone arithmetic.
import { clock, weekdayDayMonth } from '../dates';
import { SITE_LOCALE } from '../format/number';

const MINUTE_MS = 60_000;
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
  const formatter = new Intl.DateTimeFormat(SITE_LOCALE, {
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
    new Intl.DateTimeFormat(SITE_LOCALE, { timeZone: tz, timeZoneName: 'short' })
      .formatToParts(at)
      .find((part) => part.type === 'timeZoneName')?.value ?? tz
  );
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

/** `Thu 10 Dec, 11:42 pm` in the viewer's zone, and ` (10:42 pm guild time)` apart (null) when the
 *  guild's offset is the same, so a caller can set the guild time in its own muted span. */
export function stampParts(
  at: Date,
  viewerTz: string,
  guildTz: string,
  locale: string = SITE_LOCALE,
): { when: string; guildTime: string | null } {
  const when = `${weekdayDayMonth(at, viewerTz)}, ${clock(at, viewerTz, { locale, alwaysMinutes: true })}`;
  if (!offsetsDiffer(at, viewerTz, guildTz)) return { when, guildTime: null };
  return { when, guildTime: `(${clock(at, guildTz, { locale, alwaysMinutes: true })} guild time)` };
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
