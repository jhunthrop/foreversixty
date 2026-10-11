// web/src/lib/rankings/row-view.ts
// How a board's rows read, as plain functions: the rank's parse colour and the legend that
// explains it, whether rows carry a crest, how a row's date prints, and where the viewer's own
// row goes (marked in place, or pinned on top when it is off the visible page).
import { formatDate, relativeTime } from '../dates';
import { percentileToken } from '../report/format';
import { rankingsCopy } from './copy';
import type { RankingRow, RankingsPage } from './api';

const DAY_MS = 86_400_000;
/** Inside this many days a date reads "3 days ago"; past it, "Nov 25" (the Logs rule). */
export const RELATIVE_DATE_DAYS = 14;

/** A rank's percentile within its board: first is 100, last is just above 0. */
export function rankPercentile(rank: number, total: number): number {
  return 100 - ((rank - 1) / Math.max(total, 1)) * 100;
}

/** The top of the legend; the first rank of a board is its 100th percentile and still reads pink. */
const TOP_LEGEND_PERCENTILE = 99;

export function rankColor(rank: number, total: number): string {
  return percentileToken(Math.min(TOP_LEGEND_PERCENTILE, Math.max(0, rankPercentile(rank, total))));
}

export interface LegendStep {
  label: string;
  /** The lowest percentile that earns this colour. */
  floor: number;
  color: string;
}

const LEGEND_FLOORS: readonly { label: string; floor: number }[] = [
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

/** "3 days ago" inside two weeks, "Nov 25" after. */
export function rankingDate(iso: string, now: Date = new Date()): string {
  const at = new Date(iso);
  if (now.getTime() - at.getTime() < RELATIVE_DATE_DAYS * DAY_MS) return relativeTime(at, now);
  return formatDate(at);
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
