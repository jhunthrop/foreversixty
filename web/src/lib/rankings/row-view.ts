// web/src/lib/rankings/row-view.ts
// How a board's rows read, as plain functions: the rank's parse colour and the legend that
// explains it, whether rows carry a crest, how a row's date prints, and where the viewer's own
// row goes (marked in place, or pinned on top when it is off the visible page).
import { formatDate } from '../dates';
import { percentileToken } from '../report/format';
import { rankingsCopy } from './copy';
import type { RankingRow, RankingsPage } from './api';

/** A rank's percentile within its board: first is 100, last is just above 0. */
export function rankPercentile(rank: number, total: number): number {
  return 100 - ((rank - 1) / Math.max(total, 1)) * 100;
}

/** A rank's colour on the Warcraft Logs scale, the percentile floored: the first rank of a board is
 *  its 100th percentile and gold, the next ones pink while they stay at 99 or above. */
export function rankColor(rank: number, total: number): string {
  return percentileToken(Math.max(0, rankPercentile(rank, total)));
}

export interface LegendStep {
  label: string;
  /** The lowest percentile that earns this colour. */
  floor: number;
  color: string;
}

const LEGEND_FLOORS: readonly { label: string; floor: number }[] = [
  { label: '100', floor: 100 },
  { label: '99+', floor: 99 },
  { label: '95+', floor: 95 },
  { label: '75+', floor: 75 },
  { label: '50+', floor: 50 },
  { label: '25+', floor: 25 },
  { label: rankingsCopy.legendUnderLast, floor: 0 },
];

/** The one-line legend, best first, each step in the colour the rank numerals use. */
export const RANK_LEGEND: readonly LegendStep[] = LEGEND_FLOORS.map((step) => ({
  ...step,
  color: percentileToken(step.floor),
}));

/** A board spanning more than one class shows a crest before each name; one class shows none. */
export function showRowCrests(rows: readonly Pick<RankingRow, 'player'>[]): boolean {
  return new Set(rows.map((row) => row.player.class)).size > 1;
}

/** The kill's day, "12 Dec": short enough for the board's date column at every width, and a
 *  board is read against the raid nights, not against today. */
export function rankingDate(iso: string): string {
  return formatDate(new Date(iso));
}

/** Where the viewer's own row goes: in the list where it already is, on top, or nowhere. */
export type OwnRowPlacement = 'in-place' | 'pinned' | 'none';

export function ownRowPlacement(
  rows: readonly Pick<RankingRow, 'player'>[],
  characterRank: RankingsPage['character_rank'],
): OwnRowPlacement {
  if (characterRank === undefined || characterRank === null) return 'none';
  return rows.some((row) => row.player.key === characterRank.row.player.key) ? 'in-place' : 'pinned';
}

/** The one plain line under a name for a kill a moderator has touched; null for a normal one. */
export function moderationLine(state: RankingRow['state']): string | null {
  if (state === 'removed') return rankingsCopy.removed;
  if (state === 'at_risk') return rankingsCopy.atRisk;
  return null;
}

/** The metrics where a 0 means the character did nothing on the board (a healer's 0 DPS). */
const IDLE_METRICS = new Set(['dps', 'hps']);

/**
 * The board without its idle rows: a 0 DPS (or HPS) row beside a percentile reads as broken, so it is
 * left out and the total drops by as many (they sort last, so the ranks above stay as they are).
 */
export function withoutIdleRows<P extends { rows: RankingRow[]; total: number }>(page: P, metric: string): P {
  if (!IDLE_METRICS.has(metric)) return page;
  const rows = page.rows.filter((row) => row.value > 0);
  return rows.length === page.rows.length
    ? page
    : { ...page, rows, total: page.total - (page.rows.length - rows.length) };
}
