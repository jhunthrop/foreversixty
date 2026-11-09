import { describe, expect, it } from 'vitest';
import { profileCopy } from './recruit-copy';

describe('profileCopy attendance lines', () => {
  it('says one raid in the singular', () => {
    expect(profileCopy.lastNights(1)).toBe('Last raid');
    expect(profileCopy.lastThirty(1, 1)).toBe('1 of 1 raid');
  });

  it('counts several raids', () => {
    expect(profileCopy.lastNights(4)).toBe('Last 4 raids');
    expect(profileCopy.lastThirty(25, 30)).toBe('25 of the last 30 raids');
  });
});
