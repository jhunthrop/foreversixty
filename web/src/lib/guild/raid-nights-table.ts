// web/src/lib/guild/raid-nights-table.ts
// The Raids tab's night table (guild dense spec, h-guild-tabs): one row a logged night with the
// time the bosses took to die, that time against the night before, the wipes, deaths and raiders,
// so a raid leader reads the trend down one column. Boss time is the summed kill pulls of the
// tier's bosses (trash packs and wipes left out); with the tier's boss list unknown, every named
// kill counts. Rows arrive newest first (the API's order); a night is compared with the one after it.
import type { GuildRaidFight, GuildRaidRow } from './api';
import { raidTally } from './raid-view';
import { counted } from '../plural';
import { clockDuration } from '../format/duration';

export interface RaidNightFigures {
  row: GuildRaidRow;
  /** The bosses' kill pulls summed; null when no boss died. */
  bossMs: number | null;
  /** Seconds faster (negative) or slower than the night before; null without both times. */
  deltaS: number | null;
  /** The fastest boss time of the listed nights. */
  best: boolean;
  wipes: number;
}

function isBossKill(fight: GuildRaidFight, bossIds: ReadonlySet<number> | null): boolean {
  if (!fight.kill || fight.name === null) return false;
  if (bossIds === null) return fight.encounter_id !== null;
  return fight.encounter_id !== null && bossIds.has(fight.encounter_id);
}

/** The pulls on the tier's bosses (kills and wipes; trash packs left out). */
export function bossPulls(row: GuildRaidRow, bossIds: ReadonlySet<number> | null): number {
  return row.fights.filter(
    (fight) =>
      fight.name !== null &&
      fight.encounter_id !== null &&
      (bossIds === null || bossIds.has(fight.encounter_id)),
  ).length;
}

export function bossTimeMs(row: GuildRaidRow, bossIds: ReadonlySet<number> | null): number | null {
  const kills = row.fights.filter((fight) => isBossKill(fight, bossIds));
  return kills.length === 0 ? null : kills.reduce((sum, fight) => sum + fight.duration_ms, 0);
}

export function raidNightFigures(
  rows: readonly GuildRaidRow[],
  bossIds: ReadonlySet<number> | null,
): RaidNightFigures[] {
  const times = rows.map((row) => bossTimeMs(row, bossIds));
  const known = times.filter((ms): ms is number => ms !== null);
  const fastest = known.length === 0 ? null : Math.min(...known);
  return rows.map((row, index) => {
    const bossMs = times[index];
    const before = times[index + 1] ?? null;
    return {
      row,
      bossMs,
      deltaS: bossMs === null || before === null ? null : Math.round((bossMs - before) / 1000),
      best: known.length > 1 && bossMs !== null && bossMs === fastest,
      wipes: bossIds === null ? row.wipe_count : raidTally(row, bossIds).wipes,
    };
  });
}

/** `-6 s`, `+4 s`, `0 s`; the figure's tone: faster is the kill colour. */
export function deltaWords(deltaS: number): { value: string; tone: 'kill' | 'wipe' | undefined } {
  const sign = deltaS > 0 ? '+' : deltaS < 0 ? '−' : '';
  return {
    value: `${sign}${Math.abs(deltaS)} s`,
    tone: deltaS < 0 ? 'kill' : deltaS > 0 ? 'wipe' : undefined,
  };
}

/** The council's awards per night's report id, from the loot nights read. */
export function awardsByReport(
  nights: readonly { report_id: string | null; decisions: readonly unknown[] }[],
): ReadonlyMap<string, number> {
  return new Map(
    nights.flatMap((night) =>
      night.report_id === null ? [] : [[night.report_id, night.decisions.length] as const],
    ),
  );
}

/** The Raid nights head: `5 nights · 5 kills · best 3:12`. */
export function raidNightsMeta(rows: readonly GuildRaidRow[], bossIds: ReadonlySet<number> | null): string {
  const kills = rows.reduce((sum, row) => sum + row.fights.filter((f) => isBossKill(f, bossIds)).length, 0);
  const times = rows.map((row) => bossTimeMs(row, bossIds)).filter((ms): ms is number => ms !== null);
  const parts = [counted(rows.length, 'night'), counted(kills, 'kill')];
  if (times.length > 0) parts.push(`best ${clockDuration(Math.min(...times))}`);
  return parts.join(' · ');
}

/** The attendance grid's nights the Raids tab reads (the API's most). */
export const ATTENDANCE_GRID_NIGHTS = 12;

/** A raider a night missed: the name with its class, so the line reads it in class colour. */
export interface MissedRaider {
  name: string;
  class: string;
}

/** One guild raid's misses: the raiders absent and those benched. */
export interface NightMisses {
  absent: readonly MissedRaider[];
  bench: readonly MissedRaider[];
}

/** Officers: who missed each guild raid and who sat on the bench, from the attendance grid. */
export function absentByRaid(grid: {
  raids: readonly { id: number }[];
  rows: readonly { name: string; class: string; cells: readonly string[] }[];
}): ReadonlyMap<number, NightMisses> {
  const named = (index: number, cell: string): MissedRaider[] =>
    grid.rows.filter((row) => row.cells[index] === cell).map((row) => ({ name: row.name, class: row.class }));
  return new Map(
    grid.raids.map((raid, index) => [
      raid.id,
      { absent: named(index, 'absent'), bench: named(index, 'bench') },
    ]),
  );
}

/** The slowest boss time of the listed nights: the Meter's full width. */
export function slowestMs(figures: readonly RaidNightFigures[]): number | null {
  const known = figures.map((f) => f.bossMs).filter((ms): ms is number => ms !== null);
  return known.length === 0 ? null : Math.max(...known);
}
