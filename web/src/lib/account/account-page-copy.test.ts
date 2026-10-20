// web/src/lib/account/account-page-copy.test.ts
import { describe, expect, it } from 'vitest';
import { accountPageCopy } from './account-page-copy';

describe('accountPageCopy first-run lines', () => {
  it('says the companion signs itself in, and a code is only the fallback', () => {
    expect(accountPageCopy.noDevices).toBe(
      'No devices yet. The companion signs in with Battle.net from its Device page; a code is only for the fallback.',
    );
    expect(accountPageCopy.pairADevice).toBe('Get a pairing code');
  });

  it('greets an account with no characters by name instead of naming a character', () => {
    expect(accountPageCopy.signedInBannerEmpty('Hunt#1234')).toBe(
      'Signed in as Hunt#1234. Install the addon and the companion to bring your characters here.',
    );
  });
});
