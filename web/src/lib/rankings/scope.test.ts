// web/src/lib/rankings/scope.test.ts
import { describe, expect, it } from 'vitest';
import { cohortName, scopeHeading } from './scope';

describe('the board scope', () => {
  it('names the cohort and heads the board with its plural', () => {
    expect(cohortName('warrior', 'fury')).toBe('Fury Warrior');
    expect(cohortName('warrior', '')).toBe('Warrior');
    expect(cohortName('', '')).toBe('');
    expect(scopeHeading('warrior', 'Fury')).toBe('Fury Warriors');
    expect(scopeHeading('priest', '')).toBe('Priests');
    expect(scopeHeading('', '')).toBe('Every class and spec');
  });
});
