import { describe, expect, it } from 'vitest';
import { reviewedDate } from './dates';

describe('reviewedDate', () => {
  it('reads day and short month', () => {
    expect(reviewedDate('2026-10-28T12:00:00Z')).toBe('28 Oct');
  });
  it('keeps unparseable text', () => {
    expect(reviewedDate('soon')).toBe('soon');
  });
});
