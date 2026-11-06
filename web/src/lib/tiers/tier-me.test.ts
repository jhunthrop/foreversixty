// web/src/lib/tiers/tier-me.test.ts
import { describe, expect, it } from 'vitest';
import { tierMeFrom } from './tier-me';

const me = (character: Record<string, unknown>) => ({
  characters: [
    { key: 'us/normal/t', region: 'us', ruleset: 'normal', name: 'Thandrel', level: 60, ...character },
  ],
});

describe('tierMeFrom', () => {
  it('keys a character with a class and a spec', () => {
    expect(tierMeFrom(me({ class: 'Warrior', spec: 'Fury' }) as never, null)).toEqual({
      kind: 'spec',
      key: 'warrior/fury',
    });
  });

  it('names a character whose spec is not known yet', () => {
    expect(tierMeFrom(me({ class: 'Warrior' }) as never, null)).toEqual({
      kind: 'unknown',
      name: 'Thandrel',
    });
  });

  it('answers nothing signed out or with no character', () => {
    expect(tierMeFrom(null, null)).toEqual({ kind: 'none' });
    expect(tierMeFrom({ characters: [] } as never, null)).toEqual({ kind: 'none' });
  });
});
