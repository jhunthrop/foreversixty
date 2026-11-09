// web/src/lib/character/score-parts.ts
// The Forever Score panel's parts (directions mock b2-character): every part's track on one scale
// (as long as its maximum against the largest), so a part's share of the score reads at a glance
// with no "of N"; the fill is the part's share of its own maximum, coloured on the item quality
// ladder players already read (common, uncommon, rare, epic).
import labels from '../../../../docs/contracts/fixtures/score-labels.json';
import { COMPONENT_MAX, type ScoreComponentKey, type ScoreComponents } from '../score/score';

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
