// web/src/lib/logs/report-summary.ts
// The second line of a report row (design/specs/2026-10-11-logs.md section 2.5): the zone, how
// many bosses died and how long the night ran, "Blackrock Depths · 4 bosses · 2 h 10 m". A part
// the API did not send is left out rather than guessed; the list endpoint does not carry the
// length yet, so it prints when it does.
import { bossCountLabel } from '../boss-count';
import { reportCredit } from '../report/credit';

const MS_PER_MINUTE = 60_000;
const MINUTES_PER_HOUR = 60;
const SEPARATOR = ' · ';

/** "48 m", "2 h 10 m", "1 h 05 m"; the minute is padded only once an hour precedes it. */
export function formatReportLength(durationMs: number): string {
  const totalMinutes = Math.round(durationMs / MS_PER_MINUTE);
  const hours = Math.floor(totalMinutes / MINUTES_PER_HOUR);
  const minutes = totalMinutes % MINUTES_PER_HOUR;
  if (hours === 0) return `${minutes} m`;
  return `${hours} h ${String(minutes).padStart(2, '0')} m`;
}

const plural = (count: number, one: string, many: string): string => `${count} ${count === 1 ? one : many}`;

function join(parts: readonly (string | undefined)[]): string {
  return parts.filter((part): part is string => part !== undefined && part !== '').join(SEPARATOR);
}

export interface OwnReportFacts {
  zone: string;
  killCount: number;
  durationMs?: number;
}

/** A report you own: zone, bosses, length. */
export function ownReportSummary({ zone, killCount, durationMs }: OwnReportFacts): string {
  return join([
    zone,
    bossCountLabel(killCount),
    durationMs === undefined ? undefined : formatReportLength(durationMs),
  ]);
}

export interface PublicReportFacts {
  zone?: string;
  killCount: number;
  guildName?: string;
}

/** A public report: the zone when known, the kills, the guild when it has one. */
export function publicReportSummary({ zone, killCount, guildName }: PublicReportFacts): string {
  return join([zone, plural(killCount, 'kill', 'kills'), guildName]);
}

export interface RecentReportFacts {
  zone?: string;
  killCount: number;
  /** Distinct bosses killed, when the API sends it; absent, the kill count stands in. */
  bossCount?: number;
  fightCount: number;
  guildName?: string;
}

/** A row of the public list, in the shape "Your reports" uses (the character and the badge are
 *  ReportRow's own): zone, bosses when there are any, fights, then the guild when it has one. */
export function recentReportSummary({
  zone,
  killCount,
  bossCount,
  fightCount,
  guildName,
}: RecentReportFacts): string {
  const bosses = bossCount ?? killCount;
  return join([
    zone === undefined || zone === '' ? undefined : zone,
    bosses > 0 ? bossCountLabel(bosses) : undefined,
    plural(fightCount, 'fight', 'fights'),
    guildName,
  ]);
}

export interface AccountReportFacts {
  zone: string;
  fightCount: number;
  killCount: number;
  /** Distinct bosses killed, when the API sends it; absent, the kill count stands in. */
  bossCount?: number;
}

/**
 * The account page's and /logs' own-report summary: where, then the figures that mean
 * something to the player. "Darnassus · 0 fights", "Blackrock Depths · 3 bosses · 14 fights";
 * bosses print only when there are any, and a report with no zone says so.
 */
export function accountReportSummary({ zone, fightCount, killCount, bossCount }: AccountReportFacts): string {
  const bosses = bossCount ?? killCount;
  return join([
    zone === '' ? 'No zone yet' : zone,
    bosses > 0 ? bossCountLabel(bosses) : undefined,
    plural(fightCount, 'fight', 'fights'),
  ]);
}

export interface PublicRowFacts extends Omit<RecentReportFacts, 'zone'> {
  /** The report's own title, or the zone the API resolved it to. */
  title: string;
  zone?: string;
  /** The credited character as the API sends it; null or absent when none is credited. */
  loggingCharacter?: string | { key?: string; name?: string } | null;
}

/**
 * A row of the public list: the credited character is the title (a player looks for a name) and
 * the zone leads the muted line under it; with nobody credited the title is the report's own, and
 * the zone is not said a second time when it already is the title. The age is the row's to add.
 */
export function publicReportRow(facts: PublicRowFacts): { title: string; summary: string } {
  const credited = creditedName(facts.loggingCharacter);
  const zone = (facts.zone ?? '').trim();
  const title = credited ?? (facts.title === '' ? zone : facts.title);
  const zoneIsTitle = credited === null && title === zone;
  return {
    title,
    summary: recentReportSummary({
      zone: zoneIsTitle ? undefined : zone,
      killCount: facts.killCount,
      fightCount: facts.fightCount,
      guildName: facts.guildName,
    }),
  };
}

/** The statuses a row names: a report still being written, read or sent. A finished one
 *  ("complete", "ready") carries no Tag (logs spec 2026-11-07 MUST 4). */
const UNFINISHED_STATUS: Readonly<Record<string, string>> = {
  live: 'Live',
  processing: 'Processing',
  uploading: 'Uploading',
};

/** The row's Tag: "Live" while the log is still being written, "Processing" while it is read. Nothing once finished. */
export function reportStatusLabel(status: string): string {
  return UNFINISHED_STATUS[status] ?? '';
}

/** Shown on a live row the viewer owns that no character is credited to yet. */
export const LOG_OUT_HINT = 'Log out once in game.';

/** The character a row names, from `logging_character` as the API sends it; null when not credited. */
export function creditedName(
  loggingCharacter: string | { key?: string; name?: string } | null | undefined,
): string | null {
  const credit = reportCredit({ logging_character: loggingCharacter });
  return credit.credited ? credit.name : null;
}

export interface OwnReportRowFacts {
  title: string;
  zone: string;
  /** When the night began (its first pull), else the upload. */
  startedAt: string;
  fightCount: number;
  killCount: number;
  bossCount?: number;
  loggingCharacter?: string | { key?: string; name?: string } | null;
}

/**
 * A report row in "Your reports" (logs spec 2026-11-07 section 5): the title falls back to the
 * zone, then to what the night held ("Trash only", "3 bosses"; production's first report had
 * neither); the line under it (directions b1) is the zone when known, the pulls, the bosses or
 * "Trash only", and who it is credited to when someone is.
 */
export function ownReportRow(facts: OwnReportRowFacts): { title: string; sub: string } {
  const zone = facts.zone.trim();
  const bosses = facts.bossCount ?? facts.killCount;
  const content = bosses > 0 ? bossCountLabel(bosses) : TRASH_ONLY;
  // No title and no zone: the row is named by what the night held; the row's one date sits at its right.
  const title = facts.title.trim() || zone || content;
  return {
    title,
    sub: join([
      zone === '' || zone === title ? undefined : zone,
      plural(facts.fightCount, 'pull', 'pulls'),
      content === title ? undefined : content,
      creditedName(facts.loggingCharacter) ?? undefined,
    ]),
  };
}

const TRASH_ONLY = 'Trash only';
