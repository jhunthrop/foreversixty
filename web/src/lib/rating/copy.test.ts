// web/src/lib/rating/copy.test.ts
import { describe, expect, it } from 'vitest';
import { ratingCopy } from './copy';

describe('ratingCopy.characterEmpty', () => {
  it('names what fills the panel in, not just that it is empty (spec 2026-09-25 §3.6)', () => {
    expect(ratingCopy.characterEmpty).toBe(
      'Nothing rated yet. Ratings appear after your first ranked fight.',
    );
  });
});

describe('ratingCopy.excludedReason', () => {
  it('says a log with no character sheet leaves Preparation unscored, in player words', () => {
    expect(ratingCopy.excludedReason('Preparation', 'no_character_sheet')).toBe(
      'Preparation — no character sheet in this log.',
    );
  });

  it('falls back to the bare not-scored line for a reason it has no words for', () => {
    expect(ratingCopy.excludedReason('Utility', 'something_new')).toBe('Utility — not scored.');
  });
});
