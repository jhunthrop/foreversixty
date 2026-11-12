// web/src/lib/report/rating-card.ts
// What the report's Rating tab and rail print for a rated player (logs spec 2026-11-07 section
// 6.2, Rating row; copy section 10): the parts in the weight table's order, each part's one-line
// definition without its `{Part}: ` prefix, the basis sentence and column, the score Meter's tone
// (green from 75) and a moment's link text. Browser safe.
import { percentileToken } from './format';
import { ratingCopy } from '../rating/copy';
import { RATING_COMPONENT_ORDER } from '../rating/types';
import type { RatingCardPlayer, RatingComponent, RatingMoment } from '../rating/types';
import { clockDuration } from '../format/duration';

export const RATING_LOADING = 'Loading ratings…';

/** A part scores green from here (section 6.2). */
export const GOOD_SCORE = 75;

/** `{name}’s rating`, the player's first row. */
export function playerRating(name: string): string {
  return `${name}’s rating`;
}

/** The six parts in the weight table's order; a part the API left out reads as excluded. */
export function orderedComponents(player: RatingCardPlayer): RatingComponent[] {
  return RATING_COMPONENT_ORDER.map(
    (name) =>
      player.components.find((part) => part.name === name) ?? {
        name,
        score: null,
        weight: 0,
        basis: '',
        percentile: null,
        bracket_n: 0,
        excluded: true,
        reason: '',
        moments: [],
      },
  );
}

/** The part's definition with the `{Part}: ` prefix the copy module carries dropped. */
export function partDefinition(name: string): string {
  return (ratingCopy.componentDefinition(name) ?? '').replace(/^[^:]+:\s*/, '');
}

/** The Basis column: `Percentile of {n}` or `Absolute`. */
export function basisColumn(part: RatingComponent): string {
  if (part.basis === 'percentile' && part.bracket_n > 0) return `Percentile of ${part.bracket_n}`;
  if (part.basis === 'absolute') return 'Absolute';
  return '–';
}

/** The first row's basis sentence: compared with the spec's logs when any part was, else absolute. */
export function basisSentence(player: RatingCardPlayer): string {
  const ranked = player.components.find(
    (part) => !part.excluded && part.basis === 'percentile' && part.percentile !== null,
  );
  if (ranked !== undefined)
    return ratingCopy.percentileBasis(
      player.spec,
      player.class,
      player.role,
      ranked.percentile ?? 0,
      ranked.bracket_n,
    );
  return player.basis === 'absolute' || player.basis === 'mixed' ? ratingCopy.absoluteBasis : '';
}

export function scoreTone(score: number): 'kill' | 'gold' {
  return score >= GOOD_SCORE ? 'kill' : 'gold';
}

/** An overall's colour: the parse ladder (grey, green, blue, purple, orange, pink, gold), the
 *  scale a raider reads a number out of 100 on at a glance. */
export function ratingColour(score: number): string {
  return percentileToken(score);
}

/** `0:12 · Kick`, or the spell alone when the moment has no time. */
export function momentText(moment: RatingMoment): string {
  const name = moment.spell_name ?? 'Unnamed';
  return moment.at_ms === undefined ? name : `${clockDuration(moment.at_ms)} · ${name}`;
}
