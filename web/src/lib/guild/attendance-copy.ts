// web/src/lib/guild/attendance-copy.ts
// Every string of design/specs/2026-10-27-guild-attendance.md 6.1, in one module: constants where
// the words are fixed, functions where the data fills them. No number is typed into prose;
// attendance-copy.test.ts reads the spec table itself, so a drift in either file fails.
import type {
  ConsumableBefore,
  MoveMetric,
  RaidConsumableRow,
  RaidKill,
  RaidMatch,
  RaidNight,
  RaidRef,
  RaidUnnamed,
} from './attendance-contract';
import { nightsLabel } from './nights';
import { SEP } from './raid-copy';
import { DEFAULT_LOCALE, clockText, dateLabel, stampParts } from './raid-time';

export const SECONDS_PER_MINUTE = 60;
const MINUTE_MS = 60_000;
const SECOND_MS = 1000;
const HOUR_MS = 60 * MINUTE_MS;
export const DAY_MS = 24 * HOUR_MS;
/** The match picker lists logs uploaded within this many days. */
export const PICKER_WINDOW_DAYS = 14;
/** The roster's two attendance windows, in logged raids (spec 3.6). */
export const ROSTER_SHORT_WINDOW = 8;
export const ROSTER_LONG_WINDOW = 30;
const MINUTES_PER_HOUR = 60;
/** The auto-matcher's window after a raid ends; past it the page asks for an upload (spec 3.1). */
export const ENDED_MATCH_WINDOW_MS = 6 * HOUR_MS;
/** Discord's copied report is at most this many lines (spec 4 "Copy for Discord"). */
export const DISCORD_REPORT_MAX_LINES = 12;
/** A death or consumables list shows this many rows before `Show all` (spec 7). */
export const REPORT_ROWS_BEFORE_SHOW_ALL = 10;
export const BIGGEST_CHANGES_ROWS = 5;
export const PICKER_MAX_ROWS = 20;
/** How long `Copied` stays up. */
export const COPIED_VISIBLE_MS = 4000;
const MINUS = '−'; // a real minus sign: a fall is `−12`, not `-12`

/** `3 wipes`, `1 wipe`. */
function unit(n: number, singular: string): string {
  return `${n} ${singular}${n === 1 ? '' : 's'}`;
}

export const kills = (n: number): string => unit(n, 'kill');
export const wipes = (n: number): string => unit(n, 'wipe');
export const deaths = (n: number): string => unit(n, 'death');

/** `2 h 41 min`, `2 h`, `41 min`; a log under a minute reads `1 min`. */
export function durationLabel(ms: number): string {
  const total = Math.max(1, Math.round(ms / MINUTE_MS));
  const hours = Math.floor(total / MINUTES_PER_HOUR);
  const minutes = total % MINUTES_PER_HOUR;
  if (hours === 0) return `${minutes} min`;
  return minutes === 0 ? `${hours} h` : `${hours} h ${minutes} min`;
}

/** A kill's duration as the game's clock: `1:42`, `12:05`, `1:02:09`. */
export function killTimeLabel(ms: number): string {
  const totalSeconds = Math.round(ms / SECOND_MS);
  const seconds = totalSeconds % SECONDS_PER_MINUTE;
  const totalMinutes = Math.floor(totalSeconds / SECONDS_PER_MINUTE);
  const pad = (n: number): string => String(n).padStart(2, '0');
  if (totalMinutes < MINUTES_PER_HOUR) return `${totalMinutes}:${pad(seconds)}`;
  return `${Math.floor(totalMinutes / MINUTES_PER_HOUR)}:${pad(totalMinutes % MINUTES_PER_HOUR)}:${pad(seconds)}`;
}

export const attendanceCopy = {
  cameWord: 'came',
  match: {
    loggedAsNight: 'Logged as a raid night.',
    openNight: 'Open the night',
    openRaid: 'Open the raid',
    byTime: 'matched by time',
    byOfficer: 'matched by an officer',
    notThisLog: 'Not this log',
    clearing: 'Clearing…',
  },
  ended: {
    noLog: 'Ended · no log yet',
    matchSoon: ' We match the log when it is uploaded.',
    uploadLead: ' ',
    uploadLink: 'Upload the log',
    uploadTail: ', then match it here.',
    matchALog: 'Match a log',
  },
  clearConfirm: {
    prompt: "Clear the match? This raid's attendance is removed until another log is matched.",
    confirm: 'Clear the match',
    keep: 'Keep it',
  },
  picker: {
    head: 'Match a log',
    radiogroupLabel: 'Logs to match',
    closest: 'Closest',
    match: 'Match this log',
    matching: 'Matching…',
    cancel: 'Cancel',
    loadFailed: 'These logs did not load.',
    tryAgain: 'Try again',
    matchedElsewhere: 'That log was just matched to another raid.',
  },
  yourNight: {
    head: 'Your night',
    excused: 'Excused',
    noLoot: 'No loot this raid.',
  },
  report: {
    head: 'Report',
    copy: 'Copy for Discord',
    copied: 'Copied',
    post: 'Post to Discord',
    posting: 'Posting…',
    kills: 'Kill order',
    deaths: 'Deaths',
    consumables: 'Consumables',
    loot: 'Loot',
    firstPull: 'first pull',
  },
  deaths: { none: 'Nobody died.' },
  consumables: {
    before: 'Before',
    inTheLog: 'In the log',
    state: { stocked: 'Stocked', short: 'Short', unknown: 'Not shared' } satisfies Record<
      ConsumableBefore,
      string
    >,
    noBuffData: 'This log carries no buff data, so only the readiness check shows.',
  },
  loot: {
    officerEmptyLead: 'No loot awarded for this raid. Award it on the ',
    officerEmptyLink: 'Loot tab',
    officerEmptyTail: '.',
    memberEmpty: 'No loot recorded for this raid.',
  },
  attendance: {
    head: 'Attendance',
    groups: {
      noShow: 'Signed up, not in the log',
      late: 'Late',
      joined: 'Joined from the bench or without a sign-up',
      present: 'Present',
      benchOut: 'Bench and out',
    },
    present: 'Present',
    notInLog: 'Not in the log',
    excused: 'Excused',
  },
  roster: {
    noRaids: 'No raids yet',
    footnote: `Came: raids attended of raids signed up for or attended, over this guild's last ${ROSTER_SHORT_WINDOW} and ${ROSTER_LONG_WINDOW} logged raids.`,
  },
  readiness: {
    noExport: 'No export yet',
    nobodySigned: 'Nobody signed up for this raid yet.',
  },
  errors: {
    buttons: 'That did not save; try again.',
    rows: 'That did not save. Try again.',
    post: 'Could not post. Try again.',
    reportLoad: 'This report did not load.',
    officerToolsEnded: "Not saved. This guild's officer tools have ended.",
    copyFailed: 'Could not copy. Select the report and copy it.',
  },
  liveRegion: "The raid's log was matched.",
  changelog:
    "Added attendance and raid-night reports. A raid's log is matched when it is uploaded, and the raid page shows who came, kill times, deaths, loot and the biggest changes since last time.",
} as const;

/** The night line, the figures apart from the `came` ratio: a zero clause is left out. */
export function nightFacts(night: RaidNight): string[] {
  const facts: string[] = [];
  if (night.kills > 0) facts.push(kills(night.kills));
  if (night.wipes > 0) facts.push(wipes(night.wipes));
  if (night.deaths > 0) facts.push(deaths(night.deaths));
  if (night.duration_ms > 0) facts.push(durationLabel(night.duration_ms));
  return facts;
}

/** `22 of 23 came`. */
export function cameRatio(night: Pick<RaidNight, 'came' | 'signed'>): string {
  return `${night.came} of ${night.signed} ${attendanceCopy.cameWord}`;
}

/** `22 of 23 came · 1 kill · 3 wipes · 14 deaths · 2 h 41 min` as one string. */
export function nightLine(night: RaidNight): string {
  return [cameRatio(night), ...nightFacts(night)].join(SEP);
}

/** `Thu 10 Dec, 11:42 pm` in the viewer's zone, with ` (10:42 pm guild time)` when the guild's
 *  offset differs. */
export function stampLabel(
  at: Date,
  viewerTz: string,
  guildTz: string,
  locale: string = DEFAULT_LOCALE,
): string {
  const { when, guildTime } = stampParts(at, viewerTz, guildTz, locale);
  return guildTime === null ? when : `${when} ${guildTime}`;
}

/** Line 2 of the match line: `Log uploaded Thu 10 Dec, 11:42 pm by Kraggor · matched by time`. */
export function uploadedLine(
  match: Pick<RaidMatch, 'uploaded_at' | 'uploaded_by_name' | 'how'>,
  viewerTz: string,
  guildTz: string,
  locale: string = DEFAULT_LOCALE,
): string {
  const when = stampLabel(new Date(match.uploaded_at), viewerTz, guildTz, locale);
  const how = match.how === 'hand' ? attendanceCopy.match.byOfficer : attendanceCopy.match.byTime;
  return `Log uploaded ${when} by ${match.uploaded_by_name}${SEP}${how}`;
}

/** The tail of the ended line: the match promise inside the window, the upload prompt after it. */
export type EndedHint =
  { kind: 'soon'; text: string } | { kind: 'upload'; lead: string; link: string; tail: string };

export function endedHint(endsAt: Date, now: Date): EndedHint {
  const ended = attendanceCopy.ended;
  if (now.getTime() - endsAt.getTime() <= ENDED_MATCH_WINDOW_MS)
    return { kind: 'soon', text: ended.matchSoon };
  return { kind: 'upload', lead: ended.uploadLead, link: ended.uploadLink, tail: ended.uploadTail };
}

export function pickerHelper(zoneName: string, zoneKnown: boolean): string {
  return zoneKnown
    ? `${zoneName} logs this guild uploaded in the last ${PICKER_WINDOW_DAYS} days.`
    : `Logs this guild uploaded in the last ${PICKER_WINDOW_DAYS} days. ${zoneName}'s log name is not known yet, so every zone shows.`;
}

export function pickerEmpty(zoneName: string): string {
  return `No ${zoneName} logs uploaded in the last ${PICKER_WINDOW_DAYS} days. Upload the log, then match it here.`;
}

/** `by Kraggor`: the uploader in the picker, the awarding officer on a loot row. */
export const byName = (name: string): string => `by ${name}`;

/** `9 pulls · 1 kill`. */
export function pickerFigures(pulls: number, killCount: number): string {
  return `${unit(pulls, 'pull')}${SEP}${kills(killCount)}`;
}

/** `Matched to {Title} · Thu 3 Dec. Clear it there first.`, the title apart so the page links it. */
export function matchedElsewhereParts(
  ref: RaidRef,
  tz: string,
  locale: string = DEFAULT_LOCALE,
): { lead: string; title: string; tail: string } {
  return {
    lead: 'Matched to ',
    title: ref.title,
    tail: `${SEP}${dateLabel(new Date(ref.starts_at), tz, locale)}. Clear it there first.`,
  };
}

export const postedLine = (time: string): string => `Posted to Discord at ${time}.`;

export function postedTime(postedAt: string, tz: string, locale: string = DEFAULT_LOCALE): string {
  return clockText(new Date(postedAt), tz, locale, true);
}

export const changesHead = (date: string): string => `Biggest changes since ${date}`;

export const firstRaidLine = (zoneName: string): string =>
  `First ${zoneName} raid logged. Changes show from the next one.`;

export const showAllRaiders = (n: number): string => `Show all ${n} raiders`;

export const deathSplit = (onWipes: number, onKills: number): string =>
  `${onWipes} on wipes${SEP}${onKills} on kills`;

/** `after 3 wipes`, or `first pull` for a kill with no wipe before it. */
export function killAfter(wipesBefore: number): string {
  return wipesBefore === 0 ? attendanceCopy.report.firstPull : `after ${wipes(wipesBefore)}`;
}

/** ` · 2 deaths`; empty when nobody died on the kill. */
export function killDeaths(n: number): string {
  return n === 0 ? '' : `${SEP}${deaths(n)}`;
}

export const notKilled = (wipesBefore: number): string => `Not killed${SEP}${wipes(wipesBefore)}`;

export function unnamedLine(unnamed: RaidUnnamed): string {
  return `${unit(unnamed.pulls, 'pull')} with no named encounter published: ${kills(unnamed.kills)}, ${wipes(unnamed.wipes)}.`;
}

/** The kill-order row's words, by boss. */
export interface KillRowText {
  /** `1:42`, or `null` for a boss never killed. */
  time: string | null;
  /** `after 3 wipes` / `first pull`, or `Not killed · 4 wipes`. */
  note: string;
  /** ` · 2 deaths` or empty. */
  deaths: string;
  /** Only `first pull` reads in `--kill`. */
  firstPull: boolean;
}

export function killRowText(entry: RaidKill): KillRowText {
  if (!entry.killed || entry.kill_ms === null) {
    return { time: null, note: notKilled(entry.wipes_before), deaths: '', firstPull: false };
  }
  return {
    time: killTimeLabel(entry.kill_ms),
    note: killAfter(entry.wipes_before),
    deaths: killDeaths(entry.deaths),
    firstPull: entry.wipes_before === 0,
  };
}

export const metricLabel = (metric: MoveMetric): string => metric.toUpperCase();

/** `DPS 238 → 291`. */
export function moveFigures(metric: MoveMetric, before: number, after: number): string {
  return `${metricLabel(metric)} ${before} → ${after}`;
}

export const onBoss = (boss: string): string => `on ${boss}`;

/** `+53` (rising) or `−12`; a flat move reads `+0`. */
export function moveDelta(before: number, after: number): { text: string; rising: boolean; size: number } {
  const delta = after - before;
  const size = Math.abs(delta);
  return delta >= 0
    ? { text: `+${size}`, rising: delta > 0, size }
    : { text: `${MINUS}${size}`, rising: false, size };
}

/** The consumables chips and muted words for one row (`log_has_buffs` true only). */
export interface UsedChip {
  text: string;
  used: boolean;
}

export function usedChips(row: Pick<RaidConsumableRow, 'flask' | 'food' | 'potions'>): UsedChip[] {
  const chips: UsedChip[] = [];
  if (row.flask !== null) chips.push({ text: row.flask ? 'Flask' : 'No flask', used: row.flask });
  if (row.food !== null) chips.push({ text: row.food ? 'Food' : 'No food', used: row.food });
  if (row.potions !== null) {
    chips.push(
      row.potions > 0 ? { text: `Potions ${row.potions}`, used: true } : { text: 'No potions', used: false },
    );
  }
  return chips;
}

/** `22 of 23 came · 1 late · 1 not in the log · 2 more joined · 1 excused`; zero clauses are left out. */
export function attendanceSummary(night: RaidNight): string {
  const clauses = [cameRatio(night)];
  if (night.late > 0) clauses.push(`${night.late} late`);
  if (night.not_in_log > 0) clauses.push(`${night.not_in_log} not in the log`);
  if (night.joined > 0) clauses.push(`${night.joined} more joined`);
  if (night.excused > 0) clauses.push(`${night.excused} excused`);
  return clauses.join(SEP);
}

/** `Late, first seen 22 min after the first pull`. */
export const lateOutcome = (minutes: number): string =>
  `Late, first seen ${minutes} min after the first pull`;

export const cameAsOutcome = (alt: string): string => `Came as ${alt}`;

/** `9/11`. */
export const pullsRatio = (present: number, total: number): string => `${present}/${total}`;

/** The roster chip: `Came 7 of 8 · 26 of 30`, or `Came 3 of 3` alone while the longer window adds nothing. */
export function rosterChip(
  last8: { present: number; nights: number },
  last30: { present: number; nights: number } | undefined,
): string {
  const head = `Came ${last8.present} of ${last8.nights}`;
  if (last30 === undefined || last30.nights <= last8.nights) return head;
  return `${head}${SEP}${last30.present} of ${last30.nights}`;
}

/** The loot line in raid mode: `7 of 8 raids`. */
export const raidsAttended = (present: number, nights: number): string =>
  nightsLabel(present, nights, 'raid');

/** The character guild line's tail in raid mode: ` · came 7 of 8 raids`. */
export const cameGuildLine = (present: number, nights: number): string =>
  `${SEP}came ${raidsAttended(present, nights)}`;

export const nextRaidToggle = (date: string): string => `Next raid${SEP}${date}`;

/** `Showing the 23 signed In or Late for {Title}, Thu 10 Dec, 8:00 pm (9:00 pm guild time).`, the
 *  title apart so the page links it. */
export function showingParts(
  n: number,
  title: string,
  when: string,
): { lead: string; title: string; tail: string } {
  return { lead: `Showing the ${n} signed In or Late for `, title, tail: `, ${when}.` };
}

export const syncedBefore = (days: number): string =>
  `Synced ${days} ${days === 1 ? 'day' : 'days'} before the raid`;
