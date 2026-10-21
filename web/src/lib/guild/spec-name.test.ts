import { describe, expect, it } from 'vitest';
import { specName } from './spec-name';

describe('specName', () => {
  it('splits a camel-cased spec into words', () => {
    expect(specName('BeastMastery')).toBe('Beast Mastery');
  });

  it('reads a slug as a name', () => {
    expect(specName('beast-mastery')).toBe('Beast Mastery');
    expect(specName('holy')).toBe('Holy');
  });

  it('leaves a one-word spec and an empty one alone', () => {
    expect(specName('Holy')).toBe('Holy');
    expect(specName('')).toBe('');
  });
});
