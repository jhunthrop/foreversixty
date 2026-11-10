// web/src/lib/guild/raid-page.ts
// The raid page's pure rules (design/specs/2026-10-26-guild-raids.md 3.2, 4 "Page states", "Poll"):
// when the page polls, which state it is in, the band's when and count lines, and the strip's
// loading reserve kept in localStorage.
import type { GuildRaidView } from './api';
import { countLine, countRatio, countdown, yourTimeLine } from './raid-copy';
import { raidBreakdownLine } from './raid-comp';
import { zoneAbbreviation } from './raid-time';
import { chipStatus, raidEndsAt, type RaidChipStatus } from './planned-row';
import { RAID_DONE_KEY_PREFIX, RAID_ROWS_KEY } from './reserve-keys';
import { clock, longDay } from '../dates';

/** Spec 4 "Poll": the raid page asks again every 10 s. */
export const RAID_POLL_MS = 10_000;
/** Three failed polls in a row pause the poll and show the line. */
export const POLL_FAILURES_TO_PAUSE = 3;
/** The poll stops 6 h after the raid's end (G2's match window). */
const POLL_AFTER_END_MS = 6 * 3_600_000;
const RESERVE_DEFAULT_ROWS = 1;
/** No account holds more characters in one guild than a realm's character slots. */
const RESERVE_MAX_ROWS = 10;

/** The band's `Ended · no log yet` line: the strip's disabled rows read it. */
export const RAID_ENDED_LINE_ID = 'raid-ended-line';

/** Planned or locked, and before end + 6 h. */
export function pollOpen(raid: GuildRaidView, now: Date): boolean {
  const open = raid.status === 'planned' || raid.status === 'locked';
  return open && now.getTime() < raidEndsAt(raid).getTime() + POLL_AFTER_END_MS;
}

export interface RaidPageFlags {
  /** The status chip: the raid's status, or the derived `ended`. */
  status: RaidChipStatus;
  locked: boolean;
  ended: boolean;
  cancelled: boolean;
  done: boolean;
}

export function raidPageFlags(raid: GuildRaidView, now: Date): RaidPageFlags {
  const status = chipStatus(raid, now);
  return {
    status,
    locked: status === 'locked',
    ended: status === 'ended',
    cancelled: status === 'cancelled',
    done: status === 'done',
  };
}

export interface RaidWhen {
  /** `Thursday 10 December`, on the guild's calendar. */
  date: string;
  /** `8:00 pm`, the guild's clock. */
  start: string;
  /** `11:00 pm CST`: the guild's clock, its zone glued on with a no-break space. */
  end: string;
  /** `9:00 pm EST your time`, only when the offsets differ. */
  yourTime: string | null;
  countdown: string;
}

/** Glues the zone to the end time, so a wrap never strands `CST` on its own line. */
const NO_BREAK_SPACE = '\u00a0';

/** The end clock: a raid that runs to midnight reads `midnight`, never `12:00 am` or `00:00`,
 *  which reads as the start of the day it began (guild loop round 3). */
export const MIDNIGHT = 'midnight';
const MIDNIGHT_H23 = '00:00';

function endClock(endsAt: Date, tz: string, locale: string): string {
  return clock(endsAt, tz, { locale: 'en-GB', alwaysMinutes: true }) === MIDNIGHT_H23
    ? MIDNIGHT
    : clock(endsAt, tz, { locale, alwaysMinutes: true });
}

/** The band's when line (raid comp spec 3.1 line 4): the guild's date, times and zone first, then
 *  the viewer's own start when the offsets differ. */
export function raidWhen(raid: GuildRaidView, now: Date, viewerTz: string, locale: string): RaidWhen {
  const startsAt = new Date(raid.starts_at);
  const endsAt = raidEndsAt(raid);
  const guildTz = raid.guild_timezone ?? viewerTz;
  return {
    date: longDay(startsAt, guildTz),
    start: clock(startsAt, guildTz, { locale, alwaysMinutes: true }),
    end: `${endClock(endsAt, guildTz, locale)}${NO_BREAK_SPACE}${zoneAbbreviation(guildTz, startsAt)}`,
    yourTime: yourTimeLine(startsAt, guildTz, viewerTz, locale),
    countdown: countdown(startsAt, endsAt, now),
  };
}

export interface RaidCount {
  /** `23 of 40 signed up`. */
  line: string;
  /** `23/40`, beside the roster heading. */
  ratio: string;
  breakdown: string;
  /** The bar's fill in percent; `null` (no bar) while the size is unknown (R6). */
  fill: number | null;
}

/** Signed up counts In plus Late; Bench and Out are not counted (3.2 line 5). */
export function raidCount(raid: GuildRaidView): RaidCount {
  const n = raid.counts.in + raid.counts.late;
  const { size } = raid;
  return {
    line: countLine(n, size),
    ratio: countRatio(n, size),
    // Nobody signed up: no role breakdown after the line (guild round 2).
    breakdown: n === 0 ? '' : raidBreakdownLine(raid),
    fill: size === null || size <= 0 ? null : Math.min(100, Math.round((n / size) * 100)),
  };
}

/** The strip rows to reserve while the page loads: the count last drawn, else one. */
export function stripReserveRows(storage: Pick<Storage, 'getItem'> | null): number {
  try {
    const raw = storage?.getItem(RAID_ROWS_KEY) ?? null;
    const rows = raw === null ? Number.NaN : Number(raw);
    return Number.isInteger(rows) && rows >= 0 && rows <= RESERVE_MAX_ROWS ? rows : RESERVE_DEFAULT_ROWS;
  } catch {
    return RESERVE_DEFAULT_ROWS;
  }
}

/** Stores the strip's row count after a load. */
export function rememberStripRows(raid: GuildRaidView, storage: Pick<Storage, 'setItem'> | null): void {
  try {
    storage?.setItem(RAID_ROWS_KEY, String(Math.min(raid.viewer.characters.length, RESERVE_MAX_ROWS)));
  } catch {
    // A blocked store only costs the next visit its warm reserve.
  }
}

/** Spec 4 "Page states": the `Your night` rows this raid drew when it was last seen `done`, or
 *  null when it was not; the loading page then reserves the night, not the sign-up strip. */
export function reservedDoneRows(raidId: string, storage: Pick<Storage, 'getItem'> | null): number | null {
  try {
    const raw = storage?.getItem(`${RAID_DONE_KEY_PREFIX}${raidId}`) ?? null;
    const rows = raw === null ? Number.NaN : Number(raw);
    return Number.isInteger(rows) && rows >= 0 && rows <= RESERVE_MAX_ROWS ? rows : null;
  } catch {
    return null;
  }
}

/** Stores a `done` raid's own attendance rows (its `Your night` rows) after a load; any other
 *  status forgets it. */
export function rememberDone(
  raid: GuildRaidView,
  storage: Pick<Storage, 'setItem' | 'removeItem'> | null,
): void {
  const key = `${RAID_DONE_KEY_PREFIX}${raid.id}`;
  try {
    const ownRows = (raid.attendance ?? []).filter((row) => row.own).length;
    if (raid.status === 'done') storage?.setItem(key, String(Math.min(ownRows, RESERVE_MAX_ROWS)));
    else storage?.removeItem(key);
  } catch {
    // A blocked store only costs the next visit its warm reserve.
  }
}
