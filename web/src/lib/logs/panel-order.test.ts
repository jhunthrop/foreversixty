import { describe, expect, it } from 'vitest';
import { waysOrder } from './panel-order';

describe('waysOrder', () => {
  it('puts the companion above Upload when a device is paired', () => {
    expect(waysOrder({ signedIn: true, paired: true })).toEqual(['companion', 'upload']);
  });

  it('puts Upload first when signed in with no device', () => {
    expect(waysOrder({ signedIn: true, paired: false })).toEqual(['upload', 'companion']);
  });

  it('puts Upload first when signed out', () => {
    expect(waysOrder({ signedIn: false, paired: false })).toEqual(['upload', 'companion']);
  });

  it('ignores a stale paired flag once signed out', () => {
    expect(waysOrder({ signedIn: false, paired: true })).toEqual(['upload', 'companion']);
  });
});
