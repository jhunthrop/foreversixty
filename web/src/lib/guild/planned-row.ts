// web/src/lib/guild/planned-row.ts
// One planned row of the Raids tab (design/specs/2026-10-26-guild-raids.md 3.1, 3.6) as the words
// and numbers it draws, so PlannedRaidRow.svelte only lays them out. Times read in the viewer's zone;
// the guild-time label shows only when the offsets differ.
import type { GuildRaidRole, GuildRaidStatus, GuildRaidView, GuildSignupState } from './api';
import {
  countdown,
  countLine,
  countRatio,
  kindCountsRoles,
  kindOf,
  yourTimeLine,
  zonedClock,
  zoneLineOf,
} from './raid-copy';
import { raidBreakdownLine } from './raid-comp';
import { dayMonth, longDay, weekdayShort } from '../dates';

const MINUTE_MS = 60_000;
const STATE_ORDER: readonly GuildSignupState[] = ['in', 'late', 'bench', 'out'];

export type RaidChipStatus = GuildRaidStatus | 'ended';

export interface ViewerAnswer {
  state: GuildSignupState | null;
  /** `null` on a guild event, whose chip says `In` with no role (guild round 3). */
  role: GuildRaidRole | null;
}

export interface PlannedRowView {
  weekday: string;
  /** `Thursday 10 December`: the date block is drawn for the eye, so screen readers get this. */
  fullDate: string;
  day: string;
  month: string;
  title: string;
  /** The zone under the title, only when the title differs from it. */
  zoneLine: string | null;
  /** `8:00 pm CST`: the guild's clock and zone (raid comp spec 3.4). */
  time: string;
  /** `9:00 pm EST your time`, only when the viewer's offset differs from the guild's. */
  yourTime: string | null;
  countdown: string;
  /** `23/40`, or `23 signed up` while the size is unknown (R6). */
  count: string;
  breakdown: string;
  /** The bar's fill in percent; `null` (no bar) while the size is unknown. */
  fill: number | null;
  /** The status chip, only when not `planned`. */
  status: RaidChipStatus | null;
  /** The viewer chip; `null` on a cancelled raid, where a sign-up no longer matters. */
  viewer: ViewerAnswer | null;
  cancelled: boolean;
  /** Cancelled, or already started: no gold on its countdown or bar (spec 3.1 row states). */
  quiet: boolean;
}

/** Gold marks only a raid still ahead: a cancelled raid, or one whose start has passed, draws its
 *  countdown and bar fill muted (spec 3.1, 3.2). */
export function raidQuiet(raid: GuildRaidView, now: Date): boolean {
  return raid.status === 'cancelled' || now.getTime() >= Date.parse(raid.starts_at);
}

export function raidEndsAt(raid: GuildRaidView): Date {
  return new Date(Date.parse(raid.starts_at) + raid.duration_minutes * MINUTE_MS);
}

/** Spec 3.6: `ended` is derived, past its end and neither done nor cancelled. */
export function chipStatus(raid: GuildRaidView, now: Date): RaidChipStatus {
  const open = raid.status === 'planned' || raid.status === 'locked';
  return open && now.getTime() >= raidEndsAt(raid).getTime() ? 'ended' : raid.status;
}

/** The viewer's answer for the chip: a signed character (In first), else unsigned; `null` with no
 *  verified character in the guild. */
export function viewerAnswer(raid: GuildRaidView): ViewerAnswer | null {
  const roles = kindCountsRoles(kindOf(raid));
  const verified = raid.viewer.characters.filter((c) => c.verified);
  for (const state of STATE_ORDER) {
    const signed = verified.find((c) => c.state === state);
    if (signed !== undefined) return { state, role: roles ? signed.role : null };
  }
  const first = verified[0];
  return first === undefined ? null : { state: null, role: roles ? first.role : null };
}

export function plannedRowView(
  raid: GuildRaidView,
  now: Date,
  viewerTz: string,
  locale: string,
): PlannedRowView {
  const startsAt = new Date(raid.starts_at);
  const guildTz = raid.guild_timezone ?? viewerTz;
  const n = raid.counts.in + raid.counts.late;
  const status = chipStatus(raid, now);
  // One clock leads everywhere: the date block and the time read in the guild's zone.
  const [day = '', month = ''] = dayMonth(startsAt, guildTz).split(' ');
  return {
    weekday: weekdayShort(startsAt, guildTz),
    fullDate: longDay(startsAt, guildTz),
    day,
    month,
    title: raid.title,
    zoneLine: raid.title === raid.zone_name ? null : zoneLineOf(raid),
    time: zonedClock(startsAt, guildTz, locale),
    yourTime: yourTimeLine(startsAt, guildTz, viewerTz, locale),
    countdown: countdown(startsAt, raidEndsAt(raid), now),
    count: raid.size === null ? countLine(n, null) : countRatio(n, raid.size),
    breakdown: raidBreakdownLine(raid),
    fill: raid.size === null || raid.size <= 0 ? null : Math.min(100, Math.round((n / raid.size) * 100)),
    status: status === 'planned' ? null : status,
    viewer: raid.status === 'cancelled' ? null : viewerAnswer(raid),
    cancelled: raid.status === 'cancelled',
    quiet: raidQuiet(raid, now),
  };
}

/** Rows sort by start. */
export function sortByStart(planned: readonly GuildRaidView[]): GuildRaidView[] {
  return [...planned].sort((a, b) => Date.parse(a.starts_at) - Date.parse(b.starts_at));
}
