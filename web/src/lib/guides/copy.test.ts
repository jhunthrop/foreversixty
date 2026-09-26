import { describe, expect, it } from 'vitest';
import { guidesCopy } from './copy';

describe('guidesCopy', () => {
  it('has no blank strings', () => {
    for (const [key, value] of Object.entries(guidesCopy)) {
      expect(value.trim(), `${key} is blank`).not.toBe('');
    }
  });
});
