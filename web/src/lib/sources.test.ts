import { describe, expect, it } from 'vitest';
import { pillClassFor, pillLabelFor } from './sources';

describe('source pills', () => {
  it('maps each kind to a class and label', () => {
    expect(pillClassFor('blizzard')).toBe('pill-blizzard');
    expect(pillClassFor('datamined')).toBe('pill-datamined');
    expect(pillClassFor('community')).toBe('pill-community');
    expect(pillClassFor('site')).toBe('pill-site');
    expect(pillLabelFor('site')).toBe('This site');
    expect(pillLabelFor('blizzard')).toBe('Blizzard');
  });
});
