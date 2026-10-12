// web/src/lib/score/sparkline.ts
// The Forever Score card's sparkline (Task L8.4 ruling 3, fix round 2 items 1 and 2) as a pure
// model: the API's score_history (one point per UTC date, the newest 30, oldest first) in; the
// line's SVG path, the newest point and its tier colour, the two dates as labels, the caption
// ("Up 330 since 10 Nov"), the image's accessible name, and the rows of the table a screen
// reader gets, out. Null with fewer than two points: one point is not a line, and the card then
// draws no box at all; it says "First score recorded 12 Oct." instead (scoreHistoryFirst, fix
// round 3 item 12). Fewer than five points add ", 2 readings" to the caption.
//
// The box is a percentage box (100 x 100): the card stretches it over the whole content column
// at a fixed height (preserveAspectRatio="none", a non-scaling stroke), so x and y are also the
// end dot's `left` and `top` in percent. The oldest point sits on the left edge (the Gear
// label's), the newest on the right edge; the dot may overhang it.
import { formatDayMonthYear } from '../dates';
import { scoreCopy } from './copy';
import { shortDate, tierColor, type ScorePoint } from './score';

export const SPARKLINE_WIDTH = 100;
export const SPARKLINE_HEIGHT = 100;
/** Room above and below for the 8px end dot and the 2px stroke: at the phone's 40px height
 *  12% is 4.8px, more than the dot's 4px radius. Horizontally there is none (edge to edge). */
export const SPARKLINE_PAD = 12;
const PAD = SPARKLINE_PAD;
/** The least the vertical scale spans (fix round 1 item 2): a 10-point move on a 2100-point
 *  score stays a small move, not the line's full height. A narrower range is centred on its
 *  own midpoint. */
const MIN_SPAN = 200;
/** The API sends the newest 30; anything longer is cut to that here too. */
const MAX_POINTS = 30;
/** Below this many points the caption says how many readings the line is drawn through. */
export const SPARSE_POINTS = 5;

export interface ScoreSparkline {
  /** `M x,y L x,y ...`, one vertex per point, evenly spaced, oldest at the left edge. */
  path: string;
  /** The newest point: where its mark goes (percent of the box), its score, and its tier's
   *  quality colour (null for a tier the site does not know, drawn in the line's gold). */
  last: { x: number; y: number; score: number; color: string | null };
  /** "10 Nov" and "9 Dec", the addon's date form; "20 Dec 2026" and "5 Jan 2027" when the
   *  history crosses a year boundary. */
  firstLabel: string;
  lastLabel: string;
  /** "Up 330 since 10 Nov", "Down 40 since 10 Nov" or "No change since 10 Nov"; with fewer
   *  than five points ", 2 readings" after it. */
  caption: string;
  /** "Forever Score history: 1500 on 10 Nov, 1830 on 9 Dec", the SVG's aria-label. */
  ariaLabel: string;
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

/** The readable points, oldest first, the newest 30. */
function readablePoints(history: readonly ScorePoint[] | null | undefined): ScorePoint[] {
  return (history ?? [])
    .filter(readable)
    .sort((a, b) => a.date.localeCompare(b.date))
    .slice(-MAX_POINTS);
}

/** "First score recorded 12 Oct." for a history of exactly one readable point; null otherwise
 *  (none, or a line to draw). */
export function scoreHistoryFirst(history: readonly ScorePoint[] | null | undefined): string | null {
  const points = readablePoints(history);
  if (points.length !== 1) return null;
  const date = shortDate(points[0].date);
  return date === null ? null : scoreCopy.historyFirst(date);
}

export function scoreSparkline(history: readonly ScorePoint[] | null | undefined): ScoreSparkline | null {
  const points = readablePoints(history);
  if (points.length < 2) return null;

  const scores = points.map((point) => point.score);
  const middle = (Math.min(...scores) + Math.max(...scores)) / 2;
  const range = Math.max(MIN_SPAN, Math.max(...scores) - Math.min(...scores));
  const low = middle - range / 2;
  const span = points.length - 1;
  const xy = points.map((point, index) => ({
    x: round((index / span) * SPARKLINE_WIDTH),
    y: round(PAD + (1 - (point.score - low) / range) * (SPARKLINE_HEIGHT - PAD * 2)),
  }));

  const first = points[0];
  const newest = points[points.length - 1];
  const end = xy[xy.length - 1];
  const withYear = first.date.slice(0, 4) !== newest.date.slice(0, 4);
  const label = (date: string): string =>
    withYear ? formatDayMonthYear(new Date(date)) : (shortDate(date) ?? '');
  const firstLabel = label(first.date);
  const lastLabel = label(newest.date);
  return {
    path: xy.map(({ x, y }, index) => `${index === 0 ? 'M' : 'L'}${x},${y}`).join(' '),
    last: { x: end.x, y: end.y, score: newest.score, color: tierColor(newest.tier) },
    firstLabel,
    lastLabel,
    caption:
      scoreCopy.historyChange(newest.score - first.score, firstLabel) +
      (points.length < SPARSE_POINTS ? scoreCopy.historyReadings(points.length) : ''),
    ariaLabel: scoreCopy.historyLabel(first.score, firstLabel, newest.score, lastLabel),
    rows: points.map((point) => ({ date: formatDayMonthYear(new Date(point.date)), score: point.score })),
  };
}
