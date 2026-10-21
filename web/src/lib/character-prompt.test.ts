import { describe, expect, it } from 'vitest';
import { characterPrompt } from './character-prompt';

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
      quietBar: false,
    });
  });

  it('has at most one quiet secondary: sync for the owner, sync or claim for a visitor, none with gear', () => {
    expect(characterPrompt({ hasScore: false, hasGear: false, owned: true }).secondary).toBe('sync');
    expect(characterPrompt({ hasScore: false, hasGear: false, owned: false }).secondary).toBe(
      'sync-or-claim',
    );
    expect(characterPrompt({ hasScore: false, hasGear: true, owned: false }).secondary).toBeNull();
  });

  it('drops the sign-in bar sentence while there is no score', () => {
    expect(characterPrompt({ hasScore: false, hasGear: true, owned: false }).quietBar).toBe(true);
    expect(characterPrompt({ hasScore: true, hasGear: false, owned: false }).quietBar).toBe(false);
  });
});
