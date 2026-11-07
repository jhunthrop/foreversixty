// web/src/lib/account/answer-copy.test.ts
// The signed-in answer's strings live once (guides spec 2026-11-08 G2): the tier list's answer
// and the guides' answer read the same words.
import { describe, expect, it } from 'vitest';
import { tiersCopy } from '../tiers/tier-copy';
import { answerCopy } from './answer-copy';

describe('answerCopy', () => {
  it('is what the tier list’s answer says, key for key', () => {
    for (const [key, value] of Object.entries(answerCopy)) {
      expect(tiersCopy.answer[key as keyof typeof answerCopy], key).toBe(value);
    }
  });

  it('names the character whose spec is not known', () => {
    expect(answerCopy.unknown('Bow')).toBe("Bow's spec is not known yet.");
  });
});
