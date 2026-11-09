// web/src/lib/character/score-parts.ts
// The Forever Score panel's parts (directions mock b2-character): every part's track on one scale
// (as long as its maximum against the largest), so a part's share of the score reads at a glance
// with no "of N"; the fill is the part's share of its own maximum, coloured on the item quality
// ladder players already read (common, uncommon, rare, epic).
import labels from '../../../../docs/contracts/fixtures/score-labels.json';
import { COMPONENT_MAX, type ScoreComponentKey, type ScoreComponents, type ScorePoint } from '../score/score';

export interface ScoreTrend {
  /** The score at each point, oldest first, over its lowest (`fromLowest`): the spark's series. */
  values: number[];
  /** The latest score less the first. */
  change: number;
  /** The first point's date (YYYY-MM-DD). */
  since: string;
}

/** The score's trend (`score_history`, one point a day), or null under two points. */
export function scoreTrend(history: readonly ScorePoint[] | undefined): ScoreTrend | null {
  const points = [...(history ?? [])].sort((a, b) => a.date.localeCompare(b.date));
  const first = points[0];
  const last = points.at(-1);
  if (points.length < 2 || first === undefined || last === undefined) return null;
  return {
    values: fromLowest(points.map((point) => point.score)),
    change: last.score - first.score,
    since: first.date,
  };
}

/** A series over its own lowest point, so a spark draws the movement, not the distance from zero. */
export function fromLowest(values: readonly number[]): number[] {
  const low = Math.min(...values);
  return values.map((value) => value - low);
}

export type PartQuality = 'common' | 'uncommon' | 'rare' | 'epic';

/** The ladder's floors, highest first: a part at 90 % of its maximum or more is epic. */
const QUALITY_FLOORS: readonly (readonly [number, PartQuality])[] = [
  [0.9, 'epic'],
  [0.75, 'rare'],
  [0.5, 'uncommon'],
];

export function partQuality(fill: number): PartQuality {
  return QUALITY_FLOORS.find(([floor]) => fill >= floor)?.[1] ?? 'common';
}

export interface ScorePartRow {
  key: ScoreComponentKey;
  label: string;
  points: number;
  /** The track's length, 0 to 1 of the longest. */
  track: number;
  /** The fill along the track, 0 to 1 (a part over its maximum is full). */
  fill: number;
  /** The fill's colour token. */
  color: string;
}

const LABEL = new Map(labels.components.map((part) => [part.key, part.label]));
const KEYS = Object.keys(COMPONENT_MAX) as ScoreComponentKey[];
const LONGEST = Math.max(...Object.values(COMPONENT_MAX));

export function scorePartRows(components: ScoreComponents | null | undefined): ScorePartRow[] {
  return KEYS.map((key) => {
    const points = components?.[key] ?? 0;
    const max = COMPONENT_MAX[key];
    const fill = Math.min(1, points / max);
    return {
      key,
      label: LABEL.get(key) ?? key,
      points,
      track: max / LONGEST,
      fill,
      color: `var(--color-rarity-${partQuality(fill)})`,
    };
  });
}
