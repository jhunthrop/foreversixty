// web/src/lib/rating/types.ts
// The shapes docs/superpowers/specs/2026-09-21-performance-rating-design.md section 5.2
// defines. The API lane has not landed; this plan builds against the spec's own response
// examples (§5.2) per the simulator-parity precedent for building against a signature before
// its provider lands. See docs/superpowers/plans/2026-09-21-rating-web.md's Rulings 1-4 for
// where this module fills a gap the spec leaves open.

export type RatingBasis = 'percentile' | 'absolute' | 'mixed' | '';

export type RatingComponentName =
  'output' | 'survival' | 'mechanics' | 'utility' | 'preparation' | 'activity';

/** The spec §1.4 weight table's own order — every card's six components render in this
 *  order regardless of what order the API's array happens to list them in. */
export const RATING_COMPONENT_ORDER: readonly RatingComponentName[] = [
  'output',
  'survival',
  'mechanics',
  'utility',
  'preparation',
  'activity',
];

/**
 * One piece of evidence behind a component's score. `kind` drives which tab
 * `moments.ts`'s `momentHref` links into; an unrecognised kind renders as plain text
 * rather than a broken link (Ruling 4).
 */
export interface RatingMoment {
  kind: string;
  at_ms?: number;
  spell_id?: number;
  spell_name?: string;
  avoidable?: boolean;
  /** Only meaningful for `kind: "death"` — DeathsTab's own `death-<guid>-<at_ms>` anchor,
   *  supplied by the API. Every other kind's link is built locally (Ruling 4). */
  anchor?: string;
}

export interface RatingComponent {
  name: RatingComponentName;
  /** Null exactly when `excluded` is true. */
  score: number | null;
  /** The renormalised weight actually applied, 0-100. */
  weight: number;
  basis: RatingBasis;
  percentile: number | null;
  bracket_n: number;
  excluded: boolean;
  /** Set iff `excluded`; one of the fixed strings in spec §7.4. */
  reason: string;
  moments: RatingMoment[];
}

export interface RatingCardPlayer {
  player_key: string;
  player_name: string;
  class: string;
  spec: string;
  role: string;
  overall: number;
  overall_uncapped: number;
  overall_capped: boolean;
  basis: RatingBasis;
  /** Always six entries, one per RATING_COMPONENT_ORDER member, per spec §4.1's
   *  Card.Components [6]Component — a component the engine could not score is present
   *  with excluded: true, never omitted. */
  components: RatingComponent[];
}

export interface ReportRatings {
  fight_index: number;
  kill: boolean;
  kill_time_band: string;
  model_version: string;
  players: RatingCardPlayer[];
}

export interface CharacterRatingTrendPoint {
  fought_at: string;
  overall: number;
  report_id: string;
  fight_index: number;
}

export interface CharacterRating {
  player_key: string;
  sample_size: number;
  trend: CharacterRatingTrendPoint[];
  best_component: string;
  worst_component: string;
  /** Null for a character with zero rated fights (Ruling 1) or when `latest` genuinely has
   *  nothing to report; the spec's own example always populates it, but a brand-new
   *  character cannot. */
  latest: RatingCardPlayer | null;
}
