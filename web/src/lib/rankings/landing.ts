// web/src/lib/rankings/landing.ts
// The /rankings landing (coordinator round, 2026-11-09: "a raider's ranking tool, not a thin
// list"): each ranked encounter with its top three parses inline, and the leader of every spec on
// the raid boss with the most ranked kills. Healers lead by HPS, every other spec by DPS, as a
// raider reads a board. Pure functions over the boards the page reads; no record is made up, so
// a spec with no ranked kill is simply not on the list.
import { SPECS, type SpecRole } from '../sim/specs';
import type { RankFigure } from '../ui/rank-list';
import type { RankedEncounter, RankingRow, RankingsPage } from './api';
import { rankingsCopy } from './copy';
import { percentileLabel, rankRowView, type RankRowView } from './rank-row-view';

/** How many encounters the landing reads a board for, most kills first. */
export const LANDING_ENCOUNTERS = 8;
/** The parses shown under each encounter. */
export const TOP_PARSES = 3;
/** The smallest raid: a boss at this size or over is a raid boss. */
const RAID_SIZE = 10;

export type LandingMetric = 'dps' | 'hps';

/** A landing row: the board row as RankRow draws it, with two figures (the value, the percentile). */
export interface LandingRow extends RankRowView {
  key: string;
}

/** The encounter the spec leaders are read on: the raid boss with the most ranked kills, else the
 *  encounter with the most. Encounters arrive most kills first. */
export function leaderEncounter(encounters: readonly RankedEncounter[]): RankedEncounter | null {
  return encounters.find((encounter) => (encounter.size ?? 0) >= RAID_SIZE) ?? encounters[0] ?? null;
}

/** A spec's role, from its class and spec name; DPS when the pair is unknown. */
export function specRole(className: string, specName: string): SpecRole {
  const classSlug = className.toLowerCase();
  const spec = SPECS.find(
    (entry) => entry.class_slug === classSlug && entry.name.toLowerCase() === specName.toLowerCase(),
  );
  return spec?.role ?? 'dps';
}

const UNIT: Readonly<Record<LandingMetric, string>> = { dps: 'DPS', hps: 'HPS' };

/** A board row with the value (and its unit, the metrics mix on the leaders list) and the percentile. */
export function landingRow(
  row: RankingRow,
  board: Pick<RankingsPage, 'total' | 'rows'>,
  metric: LandingMetric,
): LandingRow {
  const view = rankRowView(row, {
    total: board.total,
    top: board.rows[0]?.value ?? 0,
    unit: UNIT[metric],
    multiClass: true,
  });
  const value = view.figures[0];
  const figures: RankFigure[] = [
    ...(value === undefined ? [] : [{ ...value, unit: UNIT[metric] }]),
    { value: percentileLabel(row.rank, board.total), label: rankingsCopy.percentileLabel },
  ];
  return { ...view, key: row.player.key, figures };
}

/** The first `count` rows of an encounter's board. */
export function topParses(
  board: Pick<RankingsPage, 'total' | 'rows'>,
  count: number = TOP_PARSES,
): LandingRow[] {
  return board.rows.slice(0, count).map((row) => landingRow(row, board, 'dps'));
}

/**
 * The best row of every spec on the boss: healers from the HPS board, everyone else from the DPS
 * board, each at its rank on its own board. Ordered by class, then spec, the way a raid sheet
 * reads. Only the boards' first pages are read, which hold every spec's best on today's boards.
 */
export function specLeaders(
  dps: Pick<RankingsPage, 'total' | 'rows'>,
  hps: Pick<RankingsPage, 'total' | 'rows'> | null,
): LandingRow[] {
  const leaders = new Map<string, LandingRow>();
  const take = (board: Pick<RankingsPage, 'total' | 'rows'>, metric: LandingMetric, healers: boolean) => {
    for (const row of board.rows) {
      const isHealer = specRole(row.player.class, row.player.spec) === 'healer';
      if (isHealer !== healers) continue;
      const key = `${row.player.class.toLowerCase()}/${row.player.spec}`;
      if (!leaders.has(key)) leaders.set(key, landingRow(row, board, metric));
    }
  };
  take(dps, 'dps', false);
  if (hps !== null) take(hps, 'hps', true);
  const order = (row: LandingRow) => row.classSlug;
  return [...leaders.values()].sort((a, b) => order(a).localeCompare(order(b)));
}
