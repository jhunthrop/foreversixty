// web/src/lib/score/sparkline.ts
// The profile's sparklines as one pure model (Task L8.4 ruling 3, fix round 2 items 1 and 2;
// L8.5 ruling 1): points in, the line's SVG path, the first and newest values, the newest
// point's mark and colour, the two dates as labels, the caption ("Up 330 since 10 Nov"), the
// image's accessible name, and the rows of the table a screen reader gets, out. Null with
// fewer than two points: one point is not a line.
//
// Two readers share it:
//   - scoreSparkline: the Forever Score card's score_history (one point per UTC date, the
//     newest 30, oldest first). With one point the card draws no box and says "First score
//     recorded 12 Oct." instead (scoreHistoryFirst, fix round 3 item 12); fewer than five
//     points add ", 2 readings" to the caption.
//   - ratingSparkline: the Performance Rating card's trend (one point per rated fight, any
//     order, two on one day included), each overall rounded to whole points (L8.5 ruling 1).
//
// The box is a percentage box (100 x 100): the card stretches it over its column at a fixed
// height (preserveAspectRatio="none", a non-scaling stroke), so x and y are also the end dot's
// `left` and `top` in percent. The oldest point sits on the left edge, the newest on the right
// edge; the dot may overhang it.
import { formatDayMonthYear } from '../dates';
import { ratingCopy } from '../rating/copy';
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
const SCORE_MIN_SPAN = 200;
/** The rating's, in the same proportion of its 100 (L8.5 ruling 1). */
const RATING_MIN_SPAN = 10;
/** The API sends the newest 30; anything longer is cut to that here too. */
const MAX_POINTS = 30;
/** Below this many points the caption says how many readings the line is drawn through. */
export const SPARSE_POINTS = 5;

/** The model's own point: when (YYYY-MM-DD, or an ISO time in UTC), the value, and the mark's
 *  colour for the newest (null draws it in the line's gold). */
interface SparkPoint {
  at: string;
  value: number;
  color: string | null;
}

export interface Sparkline {
  /** `M x,y L x,y ...`, one vertex per point, evenly spaced, oldest at the left edge. */
  path: string;
  /** The oldest point's value and date label. */
  first: { value: number; label: string };
  /** The newest point: where its mark goes (percent of the box), its value, and its colour
   *  (null for none, drawn in the line's gold). */
  last: { x: number; y: number; value: number; color: string | null };
  /** "10 Nov" and "9 Dec", the addon's date form; "20 Dec 2026" and "5 Jan 2027" when the
   *  line crosses a year boundary. */
  firstLabel: string;
  lastLabel: string;
  /** "Up 330 since 10 Nov", "Down 40 since 10 Nov" or "No change since 10 Nov"; with fewer
   *  than five points ", 2 readings" after it. */
  caption: string;
  /** The SVG's aria-label. */
  ariaLabel: string;
  /** The table equivalent, oldest first: "4 Nov 2026" (lib/dates' formatDayMonthYear) and the
   *  value. */
  rows: { date: string; value: number }[];
}

/** The score card's model, in its own words: `score` for the value. */
export interface ScoreSparkline extends Omit<Sparkline, 'last' | 'rows'> {
  last: { x: number; y: number; score: number; color: string | null };
  rows: { date: string; score: number }[];
}

const DATE = /^\d{4}-\d{2}-\d{2}$/;

/** The UTC date of a point, YYYY-MM-DD, or null when it is not a readable time. */
function dayOf(at: string): string | null {
  if (DATE.test(at)) return shortDate(at) === null ? null : at;
  const time = new Date(at);
  return Number.isNaN(time.getTime()) ? null : time.toISOString().slice(0, 10);
}

function round(value: number): number {
  return Math.round(value * 10) / 10;
}

/** The readable points, oldest first, the newest 30. */
function readable(points: readonly SparkPoint[]): SparkPoint[] {
  return points
    .filter((point) => dayOf(point.at) !== null && Number.isFinite(point.value))
    .sort((a, b) => new Date(a.at).getTime() - new Date(b.at).getTime())
    .slice(-MAX_POINTS);
}

function sparkline(
  input: readonly SparkPoint[],
  minSpan: number,
  ariaLabel: (first: number, firstLabel: string, last: number, lastLabel: string) => string,
): Sparkline | null {
  const points = readable(input);
  if (points.length < 2) return null;

  const values = points.map((point) => point.value);
  const middle = (Math.min(...values) + Math.max(...values)) / 2;
  const range = Math.max(minSpan, Math.max(...values) - Math.min(...values));
  const low = middle - range / 2;
  const span = points.length - 1;
  const xy = points.map((point, index) => ({
    x: round((index / span) * SPARKLINE_WIDTH),
    y: round(PAD + (1 - (point.value - low) / range) * (SPARKLINE_HEIGHT - PAD * 2)),
  }));

  const first = points[0];
  const newest = points[points.length - 1];
  const firstDay = dayOf(first.at) ?? '';
  const newestDay = dayOf(newest.at) ?? '';
  const end = xy[xy.length - 1];
  const withYear = firstDay.slice(0, 4) !== newestDay.slice(0, 4);
  const label = (day: string): string =>
    withYear ? formatDayMonthYear(new Date(day)) : (shortDate(day) ?? '');
  const firstLabel = label(firstDay);
  const lastLabel = label(newestDay);
  return {
    path: xy.map(({ x, y }, index) => `${index === 0 ? 'M' : 'L'}${x},${y}`).join(' '),
    first: { value: first.value, label: firstLabel },
    last: { x: end.x, y: end.y, value: newest.value, color: newest.color },
    firstLabel,
    lastLabel,
    caption:
      scoreCopy.historyChange(newest.value - first.value, firstLabel) +
      (points.length < SPARSE_POINTS ? scoreCopy.historyReadings(points.length) : ''),
    ariaLabel: ariaLabel(first.value, firstLabel, newest.value, lastLabel),
    rows: points.map((point) => ({
      date: formatDayMonthYear(new Date(dayOf(point.at) ?? '')),
      value: point.value,
    })),
  };
}

/** The score history's points: dated YYYY-MM-DD (a time is not a score_history date). */
function scorePoints(history: readonly ScorePoint[] | null | undefined): SparkPoint[] {
  return (history ?? [])
    .filter((point) => DATE.test(point.date))
    .map((point) => ({ at: point.date, value: point.score, color: tierColor(point.tier) }));
}

/** "First score recorded 12 Oct." for a history of exactly one readable point; null otherwise
 *  (none, or a line to draw). */
export function scoreHistoryFirst(history: readonly ScorePoint[] | null | undefined): string | null {
  const points = readable(scorePoints(history));
  if (points.length !== 1) return null;
  const date = shortDate(points[0].at);
  return date === null ? null : scoreCopy.historyFirst(date);
}

/** The score history as the shared figure draws it (Sparkline.svelte). */
export function scoreHistoryLine(history: readonly ScorePoint[] | null | undefined): Sparkline | null {
  return sparkline(scorePoints(history), SCORE_MIN_SPAN, scoreCopy.historyLabel);
}

/** The same line in the score's own words: `score` for the value. */
export function scoreSparkline(history: readonly ScorePoint[] | null | undefined): ScoreSparkline | null {
  const line = scoreHistoryLine(history);
  if (line === null) return null;
  const { last, rows, ...rest } = line;
  return {
    ...rest,
    last: { x: last.x, y: last.y, score: last.value, color: last.color },
    rows: rows.map((row) => ({ date: row.date, score: row.value })),
  };
}

/** The Performance Rating card's trend (L8.5 ruling 1): "Up 9 since 4 Dec" from the first
 *  fight to the newest, each overall rounded to whole points as the card shows them. */
export function ratingSparkline(
  trend: readonly { fought_at: string; overall: number }[] | null | undefined,
): Sparkline | null {
  const points = (trend ?? []).map((point) => ({
    at: point.fought_at,
    value: Math.round(point.overall),
    color: null,
  }));
  return sparkline(points, RATING_MIN_SPAN, (first, firstLabel, last, lastLabel) =>
    ratingCopy.trendLabel(readable(points).length, first, firstLabel, last, lastLabel),
  );
}
