// web/src/lib/guild/attendance-copy.ts
// Every string of design/specs/2026-10-27-guild-attendance.md 6.1, in one module: constants where
// the words are fixed, functions where the data fills them. No number is typed into prose;
// attendance-copy.test.ts reads the spec table itself, so a drift in either file fails.
import type { ConsumableBefore, MoveMetric, RaidMatch, RaidNight, RaidRef } from './attendance-contract';
import { NO_RAID_NIGHTS, nightsLabel } from './nights';
import { SEP } from './raid-copy';
import { stampParts } from './raid-time';
import { clock, weekdayDayMonth } from '../dates';
import { SITE_LOCALE } from '../format/number';
import { HOUR_MS, spokenDuration } from '../format/duration';
import { uiCopy, withRetry } from '../ui/copy';

export const SECONDS_PER_MINUTE = 60;
export { DAY_MS } from '../format/duration';
/** The match picker lists logs uploaded within this many days. */
export const PICKER_WINDOW_DAYS = 14;
/** The auto-matcher's window after a raid ends; past it the page asks for an upload (spec 3.1). */
export const ENDED_MATCH_WINDOW_MS = 6 * HOUR_MS;
/** Discord's copied report is at most this many lines (spec 4 "Copy for Discord"). */
export const DISCORD_REPORT_MAX_LINES = 12;
export const BIGGEST_CHANGES_ROWS = 5;
export const PICKER_MAX_ROWS = 20;
const MINUS = '−'; // a real minus sign: a fall is `−12`, not `-12`

/** `3 wipes`, `1 wipe`. */
function unit(n: number, singular: string): string {
  return `${n} ${singular}${n === 1 ? '' : 's'}`;
}

export const kills = (n: number): string => unit(n, 'kill');
export const wipes = (n: number): string => unit(n, 'wipe');
export const deaths = (n: number): string => unit(n, 'death');

export const attendanceCopy = {
  cameWord: 'came',
  /** The night line's head when nobody signed up beforehand (the bot's words): never `0 of 0 came`. */
  cameUnsigned: (n: number): string => `${n} came, nobody signed up beforehand`,
  nobodySigned: 'Nobody signed up beforehand',
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
    tryAgain: uiCopy.retry,
    matchedElsewhere: 'That log was just matched to another raid.',
    /** Before a candidate's time: it is the upload time, not the raid's (WP 11, interim). */
    uploaded: 'Uploaded',
  },
  yourNight: {
    head: 'Your night',
    excused: 'Excused',
    noLoot: 'No loot this raid.',
  },
  report: {
    head: 'Report',
    copy: 'Copy for Discord',
    copied: uiCopy.copied,
    post: 'Post to Discord',
    posting: 'Posting…',
    kills: 'Kill order',
    deaths: 'Deaths',
    consumables: 'Consumables',
    loot: 'Loot',
    firstPull: 'first pull',
    changes: 'Biggest changes',
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
      joined: 'Joined',
      present: 'Present',
      benchOut: 'Bench and out',
    },
    present: 'Present',
    notInLog: 'Not in the log',
    excused: 'Excused',
  },
  roster: {
    noNights: NO_RAID_NIGHTS,
  },
  readiness: {
    noExport: 'No export yet',
    nobodySigned: 'Nobody signed up for this raid yet.',
  },
  errors: {
    buttons: 'That did not save; try again.',
    rows: withRetry('That did not save.'),
    post: withRetry('Could not post.'),
    reportLoad: 'This report did not load.',
    officerToolsEnded: "Not saved. This guild's officer tools have ended.",
    copyFailed: 'Could not copy. Select the report and copy it.',
  },
  liveRegion: "The raid's log was matched.",
} as const;

/** Nobody signed up beforehand: the joiners are the whole night, so the head counts them. */
const nobodySignedUp = (night: Pick<RaidNight, 'signed'>): boolean => night.signed === 0;

/** The night line, the figures apart from the `came` head: a zero clause is left out, and the
 *  joiners are left to the head when nobody signed up. */
export function nightFacts(night: RaidNight): string[] {
  const facts: string[] = night.joined > 0 && !nobodySignedUp(night) ? [`${night.joined} joined`] : [];
  if (night.kills > 0) facts.push(kills(night.kills));
  if (night.wipes > 0) facts.push(wipes(night.wipes));
  if (night.deaths > 0) facts.push(deaths(night.deaths));
  if (night.duration_ms > 0) facts.push(spokenDuration(night.duration_ms));
  return facts;
}

/** `22 of 23 came`; with nobody signed up beforehand, the bot's `5 came, nobody signed up
 *  beforehand` (or `Nobody signed up beforehand` with nobody in the log either). */
export function cameRatio(night: Pick<RaidNight, 'came' | 'signed' | 'joined'>): string {
  if (nobodySignedUp(night)) {
    const came = night.came + night.joined;
    return came > 0 ? attendanceCopy.cameUnsigned(came) : attendanceCopy.nobodySigned;
  }
  return `${night.came} of ${night.signed} ${attendanceCopy.cameWord}`;
}

/** `22 of 23 came · 2 joined · 1 kill · 3 wipes · 14 deaths · 2 h 41 min` as one string. */
export function nightLine(night: RaidNight): string {
  return [cameRatio(night), ...nightFacts(night)].join(SEP);
}

/** `Thu 10 Dec, 11:42 pm` in the viewer's zone, with ` (10:42 pm guild time)` when the guild's
 *  offset differs. */
export function stampLabel(
  at: Date,
  viewerTz: string,
  guildTz: string,
  locale: string = SITE_LOCALE,
): string {
  const { when, guildTime } = stampParts(at, viewerTz, guildTz, locale);
  return guildTime === null ? when : `${when} ${guildTime}`;
}

/** Line 2 of the match line: `Log uploaded Thu 10 Dec, 11:42 pm by Kraggor · matched by time`. */
export function uploadedLine(
  match: Pick<RaidMatch, 'uploaded_at' | 'uploaded_by_name' | 'how'>,
  viewerTz: string,
  guildTz: string,
  locale: string = SITE_LOCALE,
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
): { lead: string; title: string; tail: string } {
  return {
    lead: 'Matched to ',
    title: ref.title,
    tail: `${SEP}${weekdayDayMonth(new Date(ref.starts_at), tz)}. Clear it there first.`,
  };
}

export const postedLine = (time: string): string => `Posted to Discord at ${time}.`;

export function postedTime(postedAt: string, tz: string, locale: string = SITE_LOCALE): string {
  return clock(new Date(postedAt), tz, { locale });
}

/** `after 3 wipes`, or `first pull` for a kill with no wipe before it. */
export function killAfter(wipesBefore: number): string {
  return wipesBefore === 0 ? attendanceCopy.report.firstPull : `after ${wipes(wipesBefore)}`;
}

export const metricLabel = (metric: MoveMetric): string => metric.toUpperCase();

export const onBoss = (boss: string): string => `on ${boss}`;

/** `+53` (rising) or `−12`; a flat move reads `+0`. */
export function moveDelta(before: number, after: number): { text: string; rising: boolean; size: number } {
  const delta = after - before;
  const size = Math.abs(delta);
  return delta >= 0
    ? { text: `+${size}`, rising: delta > 0, size }
    : { text: `${MINUS}${size}`, rising: false, size };
}

/** `22 of 23 came · 1 late · 1 not in the log · 2 more joined · 1 excused`; zero clauses are left out. */
export function attendanceSummary(night: RaidNight): string {
  const clauses = [cameRatio(night)];
  if (night.late > 0) clauses.push(`${night.late} late`);
  if (night.not_in_log > 0) clauses.push(`${night.not_in_log} not in the log`);
  if (night.joined > 0 && !nobodySignedUp(night)) clauses.push(`${night.joined} more joined`);
  if (night.excused > 0) clauses.push(`${night.excused} excused`);
  return clauses.join(SEP);
}

/** `Late, first seen 22 min after the first pull`. */
export const lateOutcome = (minutes: number): string =>
  `Late, first seen ${minutes} min after the first pull`;

export const cameAsOutcome = (alt: string): string => `Came as ${alt}`;

/** The roster chip, the bot's own (api guilds attendanceChip): `7 of 8 raid nights · 26 of the
 *  last 30`, or `1 of 1 raid night` alone while the longer window adds nothing; `No raid nights
 *  yet` with no eligible night in either window. */
export function rosterChip(
  last8: { present: number; nights: number },
  last30: { present: number; nights: number } | undefined,
): string {
  if (last8.nights === 0 && last30 !== undefined && last30.nights > 0) {
    // Nothing in the short window but older nights in the long one: the long ratio alone, true.
    return raidsAttended(last30.present, last30.nights);
  }
  const head = raidsAttended(last8.present, last8.nights);
  if (last30 === undefined || last30.nights <= last8.nights) return head;
  return `${head}${SEP}${last30.present} of the last ${last30.nights}`;
}

/** The one attendance figure on the site and in the bot (guild loop round 3): `7 of 8 raid
 *  nights`, `1 of 1 raid night`, `0 of 1 raid night`, and `No raid nights yet` with no eligible
 *  night. */
export const raidsAttended = nightsLabel;

/** `came 7 of 8 raid nights`, the clause the guild line appends after its separator; `No raid
 *  nights yet` with no eligible night (never `came 0 of 0`). */
export const cameGuildClause = (present: number, nights: number): string =>
  nights === 0 ? NO_RAID_NIGHTS : `came ${raidsAttended(present, nights)}`;
export const cameGuildLine = (present: number, nights: number): string =>
  `${SEP}${cameGuildClause(present, nights)}`;

/** `Showing the 23 signed In or Late for {Title}, Thu 10 Dec, 8:00 pm (9:00 pm guild time).`, the
 *  title apart so the page links it. */
export function showingParts(
  n: number,
  title: string,
  when: string,
): { lead: string; title: string; tail: string } {
  return { lead: `Showing the ${n} signed In or Late for `, title, tail: `, ${when}.` };
}

/** `Synced 27 Oct`: a stale export (the readiness board's cutoff is 7 days, guild round 3), its
 *  day in the Last synced column's one format (`syncedWord`). */
export const syncedLine = (when: string): string => `Synced ${when}`;
