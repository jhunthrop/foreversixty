import { describe, expect, it } from 'vitest';
import { guidesCopy } from './copy';

describe('guidesCopy', () => {
  it('has no blank strings', () => {
    for (const [key, value] of Object.entries(guidesCopy)) {
      if (typeof value !== 'string') continue;
      expect(value.trim(), `${key} is blank`).not.toBe('');
    }
  });

  it('has no blank-returning copy functions', () => {
    for (const [key, value] of Object.entries(guidesCopy)) {
      if (typeof value !== 'function') continue;
      const sample = (value as (...args: (string | number)[]) => string)(1, 'Example');
      expect(sample.trim(), `${key} returns blank`).not.toBe('');
    }
  });
});
