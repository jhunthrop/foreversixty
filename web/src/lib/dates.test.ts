import { describe, expect, it } from 'vitest';
import { formatDate, isPast } from './dates';

describe('formatDate', () => {
  it('uses the site month style', () => {
    expect(formatDate(new Date('2026-09-17T00:00:00Z'))).toBe('Sept 17');
    expect(formatDate(new Date('2026-11-04T00:00:00Z'))).toBe('Nov 4');
    expect(formatDate(new Date('2027-06-01T00:00:00Z'))).toBe('June 1');
  });
});

describe('isPast', () => {
  it('compares against a supplied now', () => {
    const now = new Date('2026-09-12T12:00:00Z');
    expect(isPast(new Date('2026-09-11T00:00:00Z'), now)).toBe(true);
    expect(isPast(new Date('2026-09-17T00:00:00Z'), now)).toBe(false);
  });
});
