// web/src/lib/logs/report-summary.ts
// The second line of a report row (design/specs/2026-10-11-logs.md section 2.5): the zone, how
// many bosses died and how long the night ran, "Blackrock Depths · 4 bosses · 2 h 10 m". A part
// the API did not send is left out rather than guessed; the list endpoint does not carry the
// length yet, so it prints when it does.
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

const plural = (count: number, one: string, many: string): string =>
  `${count} ${count === 1 ? one : many}`;

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
    plural(killCount, 'boss', 'bosses'),
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
