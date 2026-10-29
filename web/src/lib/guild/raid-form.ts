// web/src/lib/guild/raid-form.ts
// The raid form's pure logic (design/specs/2026-10-26-guild-raids.md 3.5, R4, R7): defaults from the
// guild's last planned raid, the 6.1 field errors, the repeat helper under the cap, and the payload.
// The form takes the guild's wall clock; `wallClockToInstant` turns it into the stored UTC instant.
import type { FilterOption } from '../rankings/options';
import type { GuildRaidDiscord, GuildRaidDiscordPost, GuildRaidView } from './api';
import { forYouLine, raidCopy } from './raid-copy';
import { instantToWallClock, nextWeekday, wallClockToInstant } from './raid-time';
import type { RaidInput } from './raids-api';

/** The most planned or locked raids a guild holds at once (api `RaidCap`). */
export const RAID_CAP = 12;
/** R7: the repeat box plans the next 3 weeks, four raids in all. */
export const REPEAT_WEEKS = 3;
const REPEAT_TOTAL = REPEAT_WEEKS + 1;
export const TITLE_MAX = 80;
export const NOTES_MAX = 2000;
const NOTES_COUNTER_FROM = 1880;
const DURATION_MIN = 60;
const DURATION_MAX = 240;
const DURATION_STEP = 30;
const DEFAULT_DURATION = 180;
const FIRST_RAID_TIME = '20:00';

export type RaidFormMode = 'create' | 'edit';
export type RaidFormField = 'zone' | 'date' | 'time';
/** A line under a field, or `form` for the line under the actions. */
export type RaidFormErrors = Partial<Record<RaidFormField | 'form', string>>;

export interface RaidFormValues {
  zone: string;
  /** Empty saves as the zone's name (the API fills it). */
  title: string;
  /** `YYYY-MM-DD` on the guild's calendar; empty until picked. */
  date: string;
  /** `HH:MM` on the guild's clock. */
  time: string;
  duration: number;
  timezone: string;
  notes: string;
  repeat: boolean;
  postToDiscord: boolean;
  /** `Remind the day before`; kept when Post to Discord is off. */
  remindDayBefore: boolean;
}

/** `1 h`, `1 h 30 min`. */
export function durationLabel(minutes: number): string {
  const hours = Math.floor(minutes / 60);
  const rest = minutes % 60;
  return rest === 0 ? `${hours} h` : `${hours} h ${rest} min`;
}

export function durationOptions(): FilterOption[] {
  const steps = (DURATION_MAX - DURATION_MIN) / DURATION_STEP + 1;
  return Array.from({ length: steps }, (_, i) => {
    const minutes = DURATION_MIN + i * DURATION_STEP;
    return { value: String(minutes), label: durationLabel(minutes) };
  });
}

function startMs(raid: GuildRaidView): number {
  return Date.parse(raid.starts_at);
}

/** The guild's last planned raid: the latest start that is not cancelled. */
export function lastPlannedRaid(planned: readonly GuildRaidView[]): GuildRaidView | null {
  return planned
    .filter((raid) => raid.status !== 'cancelled')
    .reduce<GuildRaidView | null>(
      (last, raid) => (last === null || startMs(raid) > startMs(last) ? raid : last),
      null,
    );
}

/** The raids the cap counts: planned and locked. */
export function openRaidCount(planned: readonly GuildRaidView[]): number {
  return planned.filter((raid) => raid.status === 'planned' || raid.status === 'locked').length;
}

export interface CreateDefaultsArgs {
  last: GuildRaidView | null;
  now: Date;
  /** The guild's `raid_timezone`, `null` before one is set. */
  guildTimezone: string | null;
  viewerTimezone: string;
  /** The guild channel's state; its remembered Post to Discord answer seeds the box. */
  discord: GuildRaidDiscord | null;
}

export function createDefaults({
  last,
  now,
  guildTimezone,
  viewerTimezone,
  discord,
}: CreateDefaultsArgs): RaidFormValues {
  const timezone = guildTimezone ?? viewerTimezone;
  const postDefault = discord?.post_default ?? false;
  const remindDayBefore = discord?.remind_default ?? true;
  const base = { title: '', notes: '', timezone, repeat: false };
  if (last === null) {
    return {
      ...base,
      zone: '',
      date: '',
      time: FIRST_RAID_TIME,
      duration: DEFAULT_DURATION,
      postToDiscord: postDefault,
      remindDayBefore,
    };
  }
  const { date, time } = instantToWallClock(nextWeekday(new Date(last.starts_at), now, timezone), timezone);
  return {
    ...base,
    zone: last.zone,
    date,
    time,
    duration: last.duration_minutes,
    postToDiscord: postDefault,
    remindDayBefore,
  };
}

/** The edit form reads the raid back on the guild's clock. The view carries the rendered notes
 *  only, so the notes field starts empty until the API answers the raw text. */
export function editDefaults(raid: GuildRaidView, timezone: string): RaidFormValues {
  const { date, time } = instantToWallClock(new Date(raid.starts_at), timezone);
  return {
    zone: raid.zone,
    title: raid.title === raid.zone_name ? '' : raid.title,
    date,
    time,
    duration: raid.duration_minutes,
    timezone,
    notes: raid.notes ?? '',
    repeat: false,
    postToDiscord: raid.discord.post_default,
    remindDayBefore: raid.discord_post?.remind ?? raid.discord.remind_default,
  };
}

export interface ValidateContext {
  now: Date;
  /** `openRaidCount` of the planned list. */
  plannedCount: number;
  mode: RaidFormMode;
  /** The stored start on an edit, which may stay in the past. */
  keepStartsAt?: string;
}

/** The raid's start instant on the guild's clock; `null` until a date is picked. */
export function startOf(values: RaidFormValues): Date | null {
  if (values.date === '' || values.time === '') return null;
  try {
    return wallClockToInstant(values.date, values.time, values.timezone);
  } catch {
    return null;
  }
}

function timeError(values: RaidFormValues, ctx: ValidateContext): string | null {
  const start = startOf(values);
  if (start === null) return values.date === '' && values.time !== '' ? null : raidCopy.form.errors.time;
  if (ctx.keepStartsAt !== undefined && start.getTime() === Date.parse(ctx.keepStartsAt)) return null;
  return start.getTime() > ctx.now.getTime() ? null : raidCopy.form.errors.time;
}

/** The 6.1 form errors, keyed by the field they sit under; `{}` when the form can be sent. */
export function validateRaidForm(values: RaidFormValues, ctx: ValidateContext): RaidFormErrors {
  const errors: RaidFormErrors = {};
  if (values.zone === '') errors.zone = raidCopy.form.errors.raid;
  if (values.date === '') errors.date = raidCopy.form.errors.date;
  const time = timeError(values, ctx);
  if (time !== null) errors.time = time;
  const adding = values.repeat ? REPEAT_TOTAL : 1;
  if (ctx.mode === 'create' && ctx.plannedCount + adding > RAID_CAP) errors.form = raidCopy.form.errors.cap;
  return errors;
}

/** The raids still allowed under the cap. */
export function slotsLeft(plannedCount: number): number {
  return Math.max(0, RAID_CAP - plannedCount);
}

/** Whether all four raids of a repeat fit under the cap, so the API never refuses one. */
export function canRepeat(plannedCount: number): boolean {
  return slotsLeft(plannedCount) >= REPEAT_TOTAL;
}

/** `Only 2 more fit: …` with fewer than four slots left, n the slots left; none at the cap,
 *  where the tab's own cap line says it. */
export function repeatHelper(plannedCount: number): string | null {
  const left = slotsLeft(plannedCount);
  if (left === 0 || canRepeat(plannedCount)) return null;
  return raidCopy.form.onlyFit(left);
}

/** The line under the zone select; an undated form reads as today on the guild's calendar. */
export function forYou(values: RaidFormValues, viewerTimezone: string, now: Date, locale?: string): string {
  const date = values.date === '' ? instantToWallClock(now, values.timezone).date : values.date;
  const time = values.time === '' ? FIRST_RAID_TIME : values.time;
  return forYouLine(date, time, values.timezone, viewerTimezone, locale);
}

/** What the form draws for Discord (spec 3.5 item 8): the box with a connected channel, the
 *  Settings line without one, and nothing when the deployment has connect off or the guild's
 *  channel is not known yet (no planned raid carries it). */
export type DiscordFormPart = 'box' | 'line' | 'none';

export function discordFormPart(discord: GuildRaidDiscord | null): DiscordFormPart {
  if (discord === null || !discord.available) return 'none';
  return discord.connected ? 'box' : 'line';
}

/** A post goes out at once when the raid starts within this many hours (P.4: 7 days). */
const POST_HORIZON_HOURS = 168;
const HOUR_MS = 3_600_000;
const CHANNEL_NAME_MAX = 60;

/** `#raid-signups`, the name cut at 60 characters and a `…`; `null` for a webhook. */
export function channelLabel(name: string | null): string | null {
  if (name === null) return null;
  return name.length > CHANNEL_NAME_MAX ? `#${name.slice(0, CHANNEL_NAME_MAX)}…` : `#${name}`;
}

/** `Thursday 3 December`: the day a week before the start, on the guild's calendar. */
function weekBefore(start: Date, timezone: string): string {
  const day = new Date(start.getTime() - POST_HORIZON_HOURS * HOUR_MS);
  const parts = new Intl.DateTimeFormat('en-US', {
    timeZone: timezone,
    weekday: 'long',
    day: 'numeric',
    month: 'long',
  }).formatToParts(day);
  const at = (type: Intl.DateTimeFormatPartTypes): string => parts.find((p) => p.type === type)?.value ?? '';
  return `${at('weekday')} ${at('day')} ${at('month')}`;
}

/** The lines and rows the Discord block draws (spec P.3). */
export interface DiscordHelper {
  post: string | null;
  remind: string;
  postRow: 'switch' | 'posted' | 'webhookMoved';
  remindRow: 'live' | 'reminded';
}

export interface DiscordHelperArgs {
  values: RaidFormValues;
  discord: GuildRaidDiscord;
  post: GuildRaidDiscordPost | null;
  mode: RaidFormMode;
  startsAt: Date | null;
  now: Date;
  timezone: string;
}

function postRowOf(discord: GuildRaidDiscord, post: GuildRaidDiscordPost | null, mode: RaidFormMode) {
  if (mode !== 'edit' || post?.posted_at == null) return 'switch';
  return discord.channel_name === null ? 'webhookMoved' : 'posted';
}

function postLine(args: DiscordHelperArgs, channel: string | null): string {
  const { values, startsAt, now, timezone } = args;
  if (!values.postToDiscord) return raidCopy.form.postOff;
  const within = startsAt === null || startsAt.getTime() - now.getTime() <= POST_HORIZON_HOURS * HOUR_MS;
  if (values.repeat) return raidCopy.form.postRepeat(within);
  return within
    ? raidCopy.form.postNow(channel)
    : raidCopy.form.postLater(channel, weekBefore(startsAt, timezone));
}

/** Whether the reminder row is switched off for want of a post: only a raid yet to be posted. */
export function remindLocked(values: RaidFormValues, postRow: DiscordHelper['postRow']): boolean {
  return postRow === 'switch' && !values.postToDiscord;
}

/** P.3's choice of helper line and row for the Discord block. */
export function discordHelper(args: DiscordHelperArgs): DiscordHelper {
  const channel = channelLabel(args.discord.channel_name);
  const postRow = postRowOf(args.discord, args.post, args.mode);
  const remindRow = args.post?.reminded_at != null ? 'reminded' : 'live';
  const remindLine = remindLocked(args.values, postRow)
    ? raidCopy.form.remindNeedsPost
    : raidCopy.form.remind;
  return {
    post:
      postRow === 'switch'
        ? postLine(args, channel)
        : postRow === 'posted' && channel !== null
          ? raidCopy.form.posted(channel)
          : null,
    remind: remindRow === 'reminded' ? raidCopy.form.reminded : remindLine,
    postRow,
    remindRow,
  };
}

/** The Settings line's link: the settings page's Discord section (spec 3.8, `#discord`). */
export function discordSettingsHref(settingsHref: string): string {
  return `${settingsHref}#discord`;
}

/** The create or edit body. `discordConnected` gates the Discord flag: a box the form does not
 *  draw never sends a remembered `true`. */
export function raidPayload(
  values: RaidFormValues,
  mode: RaidFormMode,
  discordConnected: boolean,
): RaidInput {
  const body: RaidInput = {
    zone: values.zone,
    title: values.title.trim(),
    starts_at: wallClockToInstant(values.date, values.time, values.timezone).toISOString(),
    duration_minutes: values.duration,
    notes: values.notes,
    timezone: values.timezone,
    post_to_discord: values.postToDiscord && discordConnected,
    remind_day_before: values.remindDayBefore,
  };
  return mode === 'create' ? { ...body, repeat_weeks: values.repeat ? REPEAT_WEEKS : 0 } : body;
}

/** `120 left` from 1,880 characters of notes. */
export function notesCounter(length: number): string | null {
  return length < NOTES_COUNTER_FROM ? null : raidCopy.stripStatus.left(NOTES_MAX - length);
}

const SERVER_FIELDS: Readonly<Record<string, RaidFormField>> = { zone: 'zone', starts_at: 'time' };

/** A 422's field name as the form field whose 6.1 line it shows; `null` for the line under the actions. */
export function serverFieldFor(apiField: string): RaidFormField | null {
  return SERVER_FIELDS[apiField] ?? null;
}

/** A sentence around its inline link: `Connect a Discord channel in ` + link + ` to post raids there.` */
export function splitAtLink(sentence: string, link: string): { before: string; after: string } {
  const at = sentence.indexOf(link);
  if (at < 0) throw new Error(`"${link}" is not in "${sentence}"`);
  return { before: sentence.slice(0, at), after: sentence.slice(at + link.length) };
}

/** What `onSubmit` answers: the caller maps the API's refusal onto one of these. */
export type RaidSubmitResult =
  { ok: true } | { ok: false; code: string; message: string; field?: RaidFormField };
