// web/src/lib/rating/normalise.ts
import type { RatingCardPlayer, RatingComponent, ReportRatings } from './types';

/**
 * The API writes `percentile`, `bracket_n`, `reason` and `moments` with omitempty, so a
 * component that has nothing to say for one of them arrives without the key. The page
 * checks `percentile !== null` and reads `bracket_n` as a number, and `undefined` fails
 * both silently: an absolute component would have rendered "undefined percentile". Fill
 * the gaps once, at the fetch boundary, so the types the page reads are the truth.
 */
function normaliseComponent(part: RatingComponent): RatingComponent {
  return {
    ...part,
    percentile: part.percentile ?? null,
    bracket_n: part.bracket_n ?? 0,
    reason: part.reason ?? '',
    moments: part.moments ?? [],
  };
}

function normalisePlayer(player: RatingCardPlayer): RatingCardPlayer {
  return { ...player, components: (player.components ?? []).map(normaliseComponent) };
}

export function normaliseReportRatings(wire: ReportRatings): ReportRatings {
  return { ...wire, players: (wire.players ?? []).map(normalisePlayer) };
}
