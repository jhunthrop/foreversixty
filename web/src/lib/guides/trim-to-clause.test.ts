// web/src/lib/guides/trim-to-clause.test.ts
import { describe, expect, it } from 'vitest';
import { trimToClause } from './trim-to-clause';

describe('trimToClause', () => {
  it('returns short text unchanged', () => {
    expect(trimToClause('A short sentence.')).toBe('A short sentence.');
  });

  it('breaks at the last clause mark before the limit', () => {
    const text =
      'Talents, rotation, stat priority, and race picks for Fury Warrior in Forever, with beta-versus-projection called out.';
    const trimmed = trimToClause(text, 90);
    expect(trimmed).toBe('Talents, rotation, stat priority, and race picks for Fury Warrior in Forever,');
    expect(trimmed.length).toBeLessThanOrEqual(90);
  });

  it('falls back to the last word boundary with an ellipsis when there is no clause mark', () => {
    const text = 'a'.repeat(50) + ' ' + 'b'.repeat(50);
    expect(trimToClause(text, 60)).toBe(`${'a'.repeat(50)}…`);
  });
});
