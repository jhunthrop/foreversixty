// web/src/lib/guides/guide-header.test.ts
import { describe, expect, it } from 'vitest';
import { pageUpdatedAt } from './guide-header';

describe('pageUpdatedAt', () => {
  const updated = new Date('2026-10-07T00:00:00Z');

  it('takes the BiS file’s day when it is later than the guide’s own', () => {
    expect(pageUpdatedAt(updated, '2026-10-22T23:19:15Z').toISOString()).toBe('2026-10-22T23:19:15.000Z');
  });

  it('keeps the guide’s own date when the file is older, missing or unreadable', () => {
    expect(pageUpdatedAt(updated, '2026-09-01T00:00:00Z')).toBe(updated);
    expect(pageUpdatedAt(updated, undefined)).toBe(updated);
    expect(pageUpdatedAt(updated, 'garbage')).toBe(updated);
  });
});
