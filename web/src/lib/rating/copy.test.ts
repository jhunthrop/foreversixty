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
