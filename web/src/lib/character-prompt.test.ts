import { describe, expect, it } from 'vitest';
import { characterPrompt, showsNoDataPage } from './character-prompt';

describe('characterPrompt (W.10: one prompt on an empty character page)', () => {
  it('asks for the companion when there are no logs, whatever else is missing', () => {
    expect(characterPrompt({ hasScore: false, hasGear: false, owned: false }).primary).toBe('companion');
    expect(characterPrompt({ hasScore: false, hasGear: true, owned: false }).primary).toBe('companion');
  });

  it('asks for the planner when a score exists but no gear is on file', () => {
    expect(characterPrompt({ hasScore: true, hasGear: false, owned: false }).primary).toBe('planner');
  });

  it('asks nothing of a character with a score and gear', () => {
    expect(characterPrompt({ hasScore: true, hasGear: true, owned: false })).toEqual({
      primary: null,
      secondary: null,
    });
  });

  it('has at most one quiet secondary: sync for the owner, sync or claim for a visitor, none with gear', () => {
    expect(characterPrompt({ hasScore: false, hasGear: false, owned: true }).secondary).toBe('sync');
    expect(characterPrompt({ hasScore: false, hasGear: false, owned: false }).secondary).toBe(
      'sync-or-claim',
    );
    expect(characterPrompt({ hasScore: false, hasGear: true, owned: false }).secondary).toBeNull();
  });
});

describe('showsNoDataPage (P.1, the not-mine page with no data)', () => {
  const EMPTY = { hasScore: false, hasGear: false, hasFights: false };

  it('is the one-sentence page for a visitor and a character the account does not hold', () => {
    expect(showsNoDataPage({ ...EMPTY, owned: false })).toBe(true);
  });

  it('keeps the full set for the signed-in holder', () => {
    expect(showsNoDataPage({ ...EMPTY, owned: true })).toBe(false);
  });

  it('is not used once the character has a score, gear or a fight', () => {
    expect(showsNoDataPage({ ...EMPTY, hasScore: true, owned: false })).toBe(false);
    expect(showsNoDataPage({ ...EMPTY, hasGear: true, owned: false })).toBe(false);
    expect(showsNoDataPage({ ...EMPTY, hasFights: true, owned: false })).toBe(false);
  });
});
