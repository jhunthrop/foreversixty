// web/src/lib/report/rating-card.test.ts
import { describe, expect, it } from 'vitest';
import type { RatingCardPlayer, RatingComponent } from '../rating/types';
import {
  basisColumn,
  basisSentence,
  momentText,
  orderedComponents,
  partDefinition,
  playerRating,
  ratingColour,
  scoreTone,
} from './rating-card';

function part(overrides: Partial<RatingComponent>): RatingComponent {
  return {
    name: 'output',
    score: 80,
    weight: 40,
    basis: 'percentile',
    percentile: 80,
    bracket_n: 40,
    excluded: false,
    reason: '',
    moments: [],
    ...overrides,
  };
}

const PLAYER = {
  player_key: 'k',
  player_name: 'Simfury',
  class: 'Warrior',
  spec: 'Fury',
  role: 'dps',
  overall: 70,
  overall_uncapped: 70,
  overall_capped: false,
  coverage: 1,
  insufficient: false,
  insufficient_reason: '',
  basis: 'percentile',
  components: [part({})],
} satisfies RatingCardPlayer;

describe('rating card', () => {
  it('orders all six parts and fills the missing ones as excluded', () => {
    const parts = orderedComponents(PLAYER);
    expect(parts).toHaveLength(6);
    expect(parts[0].excluded).toBe(false);
    expect(parts.slice(1).every((entry) => entry.excluded)).toBe(true);
  });

  it('drops the part prefix from a definition', () => {
    expect(partDefinition('survival')).toBe('whether you died, and how much avoidable damage you took.');
  });

  it('reads the basis column and sentence', () => {
    expect(basisColumn(part({}))).toBe('Percentile of 40');
    expect(basisColumn(part({ basis: 'absolute' }))).toBe('Absolute');
    expect(basisSentence(PLAYER)).toBe('Compared with other Fury Warriors on this fight (40 logs).');
  });

  it('scores green from 75, names the rating, and times a moment', () => {
    expect(scoreTone(75)).toBe('kill');
    expect(scoreTone(74)).toBe('gold');
    expect(playerRating('Simfury')).toBe('Simfury’s rating');
    expect(momentText({ kind: 'cast', at_ms: 12_000, spell_name: 'Kick' })).toBe('0:12 · Kick');
    expect(momentText({ kind: 'cast' })).toBe('Unnamed');
  });

  it('colours an overall on the parse ladder every raider reads', () => {
    expect(ratingColour(84)).toBe('var(--color-rarity-epic-text)');
    expect(ratingColour(58)).toBe('var(--color-rarity-rare-text)');
    expect(ratingColour(96)).toBe('var(--color-parse-orange)');
  });
});
