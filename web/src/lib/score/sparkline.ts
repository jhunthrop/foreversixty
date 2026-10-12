// web/src/lib/score/sparkline.ts
// The Forever Score card's sparkline (Task L8.4 ruling 3) as a pure model: the API's
// score_history (one point per UTC date, the newest 30, oldest first) in; the line's SVG path,
// the newest point and its tier colour, the first and last dates as labels, and the rows of the
// table a screen reader gets instead, out. Null with fewer than two points: one point is not a
// line, and the card then draws no box at all.
import { formatDayMonthYear } from '../dates';
import { shortDate, tierColor, type ScorePoint } from './score';

export const SPARKLINE_WIDTH = 240;
export const SPARKLINE_HEIGHT = 48;
/** Room for the newest point's mark (radius 3.5) and the stroke, so neither is clipped. */
export const SPARKLINE_PAD = 5;
const PAD = SPARKLINE_PAD;
/** The least the vertical scale spans (fix round 1 item 2): a 10-point move on a 2100-point
 *  score stays a small move, not the line's full height. A narrower range is centred on its
 *  own midpoint. */
const MIN_SPAN = 200;
/** The API sends the newest 30; anything longer is cut to that here too. */
const MAX_POINTS = 30;

export interface ScoreSparkline {
  /** `M x,y L x,y ...`, one vertex per point, evenly spaced, oldest at the left. */
  path: string;
  /** The newest point: where its mark goes, its score, and its tier's quality colour (null for
   *  a tier the site does not know, drawn in the line's gold). */
  last: { x: number; y: number; score: number; color: string | null };
  /** "12 Oct" and "4 Nov", the addon's date form. */
  firstLabel: string;
  lastLabel: string;
  /** The table equivalent, oldest first: "4 Nov 2026" (lib/dates' formatDayMonthYear) and the
   *  score. */
  rows: { date: string; score: number }[];
}

const DATE = /^\d{4}-\d{2}-\d{2}$/;

function readable(point: ScorePoint): boolean {
  return DATE.test(point.date) && shortDate(point.date) !== null && Number.isFinite(point.score);
}

function round(value: number): number {
  return Math.round(value * 10) / 10;
}

export function scoreSparkline(history: readonly ScorePoint[] | null | undefined): ScoreSparkline | null {
  const points = (history ?? [])
    .filter(readable)
    .sort((a, b) => a.date.localeCompare(b.date))
    .slice(-MAX_POINTS);
  if (points.length < 2) return null;

  const scores = points.map((point) => point.score);
  const middle = (Math.min(...scores) + Math.max(...scores)) / 2;
  const range = Math.max(MIN_SPAN, Math.max(...scores) - Math.min(...scores));
  const low = middle - range / 2;
  const span = points.length - 1;
  const xy = points.map((point, index) => ({
    x: round(PAD + (index / span) * (SPARKLINE_WIDTH - PAD * 2)),
    y: round(PAD + (1 - (point.score - low) / range) * (SPARKLINE_HEIGHT - PAD * 2)),
  }));

  const first = points[0];
  const newest = points[points.length - 1];
  const end = xy[xy.length - 1];
  return {
    path: xy.map(({ x, y }, index) => `${index === 0 ? 'M' : 'L'}${x},${y}`).join(' '),
    last: { x: end.x, y: end.y, score: newest.score, color: tierColor(newest.tier) },
    firstLabel: shortDate(first.date) ?? '',
    lastLabel: shortDate(newest.date) ?? '',
    rows: points.map((point) => ({ date: formatDayMonthYear(new Date(point.date)), score: point.score })),
  };
}
