// web/src/lib/account/account-page-copy.test.ts
import { describe, expect, it } from 'vitest';
import { accountPageCopy } from './account-page-copy';
import { ctaCopy } from '../addon/cta-copy';

describe('accountPageCopy first-run lines', () => {
  it('says the companion signs itself in, and a code is only the fallback', () => {
    expect(accountPageCopy.noDevices).toBe(
      'No devices yet. The companion signs in with Battle.net from its Device page; a pairing code is only for the fallback.',
    );
    // With no companion the button is the download and the code is a text link under it (W.9).
    expect(accountPageCopy.downloadCompanion).toBe('Download the companion');
    expect(accountPageCopy.pairADevice).toBe('Use a pairing code instead');
  });

  it('greets an account with no characters by name instead of naming a character', () => {
    expect(accountPageCopy.signedInBannerEmpty('Hunt#1234', true)).toBe(
      'Signed in as Hunt#1234. Battle.net lists no Forever character on this account yet; make one in game, then refresh from Battle.net.',
    );
    expect(accountPageCopy.signedInBannerEmpty('hunt@example.com', false)).toBe(
      `Signed in as hunt@example.com. Sign in with Battle.net to bring your Forever characters here, or ${ctaCopy.installAddonAndCompanionInline}.`,
    );
  });
});
