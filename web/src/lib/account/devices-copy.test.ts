// web/src/lib/account/devices-copy.test.ts
import { describe, expect, it } from 'vitest';
import { accountPageCopy } from './account-page-copy';
import { devicesCopy, pairingEntryLabel } from './devices-copy';

describe('pairingEntryLabel', () => {
  it('keeps the plain code fallback, "Use a code instead", when no device is paired', () => {
    expect(pairingEntryLabel(0, accountPageCopy.pairADevice)).toBe('Use a pairing code instead');
  });

  it('becomes the quiet "Pair another computer" once a device is paired', () => {
    expect(pairingEntryLabel(1, accountPageCopy.pairADevice)).toBe('Pair another computer');
    expect(pairingEntryLabel(3, accountPageCopy.pairADevice)).toBe('Pair another computer');
  });
});

describe('devicesCopy', () => {
  it('asks before it revokes, naming the device', () => {
    expect(devicesCopy.revokeQuestion('JH-DESKTOP')).toBe('Revoke JH-DESKTOP? The companion signs out.');
    expect(devicesCopy.revoke).toBe('Revoke');
    expect(devicesCopy.cancel).toBe('Cancel');
  });

  it('puts the browser-less fallback in the line above the code', () => {
    expect(devicesCopy.codeFallbackLine).toBe(
      'Only for a companion that cannot open a browser; the companion signs in with Battle.net by itself.',
    );
  });
});
