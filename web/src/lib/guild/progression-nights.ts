// web/src/lib/guild/progression-nights.ts
// One boss's nights on the Progression tab (guild dense spec, h-guild-tabs): every night it was
// pulled, newest first, with its kill time against the night before, its pulls, wipes and deaths,
// and the boss's headline figures (best kill, how much faster than the first, deaths a pull).
// Pure: the page prints only what the API sent (a night without a kill has no kill time).
import type { GuildProgressionEncounter } from './api';

export interface ProgressionNightRow {
  reportId: string;
  /** YYYY-MM-DD, the night's date as the API counts it. */
  date: string;
  killMs: number | null;
  /** Seconds against the previous night with a kill: negative is faster. */
  deltaS: number | null;
  pulls: number;
  killed: boolean;
  wipes: number;
  deaths: number;
  first: boolean;
  best: boolean;
}

export function progressionNights(encounter: GuildProgressionEncounter): ProgressionNightRow[] {
  const nights = encounter.pulls_by_night;
  const kills = nights.map((n) => n.kill_ms).filter((ms): ms is number => ms != null);
  const best = kills.length > 1 ? Math.min(...kills) : null;
  const firstKill = nights.findIndex((n) => n.killed);
  let before: number | null = null;
  const rows = nights.map((n, index) => {
    const killMs = n.kill_ms ?? null;
    const deltaS = killMs === null || before === null ? null : Math.round((killMs - before) / 1000);
    if (killMs !== null) before = killMs;
    return {
      reportId: n.report_id,
      date: n.date,
      killMs,
      deltaS,
      pulls: n.pulls,
      killed: n.killed,
      wipes: Math.max(0, n.pulls - (n.killed ? 1 : 0)),
      deaths: n.deaths ?? 0,
      first: index === firstKill,
      best: best !== null && killMs === best,
    };
  });
  return rows.reverse();
}

/** Seconds the best kill is under the first kill; null with fewer than two timed kills. */
export function fasterSinceFirst(encounter: GuildProgressionEncounter): number | null {
  const kills = encounter.pulls_by_night.map((n) => n.kill_ms).filter((ms): ms is number => ms != null);
  if (kills.length < 2) return null;
  return Math.round((kills[0] - Math.min(...kills)) / 1000);
}

/** The slowest kill of the boss's nights: the Meter's full width. */
export function slowestKillMs(rows: readonly ProgressionNightRow[]): number | null {
  const kills = rows.map((r) => r.killMs).filter((ms): ms is number => ms !== null);
  return kills.length === 0 ? null : Math.max(...kills);
}
