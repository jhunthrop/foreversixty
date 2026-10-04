// web/src/lib/guild/standing.test.ts
import { describe, expect, it } from 'vitest';
import { EVERY_CHECK_PASSES_LINE, needsBeforeThursdaySentence, standingSentence } from './standing';
import type { GuildHomeStanding } from './api';

function standing(overrides: Partial<GuildHomeStanding> = {}): GuildHomeStanding {
  return {
    spec: 'Marksmanship',
    class: 'hunter',
    same_spec_count: 3,
    rank_by_item_level: 2,
    item_level: 66,
    needs_before_next_raid: [],
    ...overrides,
  };
}

describe('standingSentence', () => {
  it('ranks by item level among same-spec peers', () => {
    expect(standingSentence(standing())).toBe(
      '2 of 3 Marksmanship Hunters this week by item level (66 ilvl)',
    );
  });

  it('reads "the only" when there is exactly one ranked peer', () => {
    expect(standingSentence(standing({ same_spec_count: 1, rank_by_item_level: 1 }))).toBe(
      'the only Marksmanship Hunter in this guild this week',
    );
  });

  it('reads "the only" when there are no ranked peers at all', () => {
    expect(standingSentence(standing({ same_spec_count: 0, rank_by_item_level: 0 }))).toBe(
      'the only Marksmanship Hunter in this guild this week',
    );
  });
});

describe('needsBeforeThursdaySentence', () => {
  it('joins fails with "; " and the exact template', () => {
    expect(needsBeforeThursdaySentence(['Enchant chest', 'Spend 1 talent point'])).toBe(
      'What the guild needs from you before Thursday: Enchant chest; Spend 1 talent point.',
    );
  });

  it('reads the exact all-clear line when every check passes', () => {
    expect(needsBeforeThursdaySentence([])).toBe(EVERY_CHECK_PASSES_LINE);
  });

  it('treats undefined/null the same as empty, never throwing', () => {
    expect(needsBeforeThursdaySentence(undefined)).toBe(EVERY_CHECK_PASSES_LINE);
    expect(needsBeforeThursdaySentence(null)).toBe(EVERY_CHECK_PASSES_LINE);
  });
});
