// web/src/lib/home-landing-copy.test.ts
import { describe, expect, it } from 'vitest';
import { homeCharacterSearchCopy } from './home-landing-copy';

describe('home-landing-copy', () => {
  it('names the character lookup and its field', () => {
    expect(homeCharacterSearchCopy.heading).toBe('Look up a character');
    expect(homeCharacterSearchCopy.nameLabel.trim().length).toBeGreaterThan(0);
  });
});
