import { describe, expect, it } from 'vitest';
import type { Me } from '../account/me-types';
import { SIGNED_OUT_VIEWER, coachTabListed, ownedKey, viewerFromMe } from './viewer';

function me(overrides: Partial<Me> = {}): Me {
  return {
    user: { id: 1, battletag: null, email: null, role: 'user', anonymize: false },
    characters: [{ key: 'us/normal/bow-jackzon', region: 'us', ruleset: 'normal', name: 'Bow Jackzon' }],
    guilds: [],
    ...overrides,
  } as Me;
}

describe('viewerFromMe', () => {
  it('null is the signed-out viewer', () => {
    expect(viewerFromMe(null)).toBe(SIGNED_OUT_VIEWER);
  });
  it('reads Premium and officer from entitlements', () => {
    const viewer = viewerFromMe(
      me({ entitlements: { log_coach: true, officer_views: true } as Me['entitlements'] }),
    );
    expect(viewer).toMatchObject({ signedIn: true, premium: true, officer: true });
  });
  it('a missing entitlements block is free and not an officer', () => {
    expect(viewerFromMe(me())).toMatchObject({ signedIn: true, premium: false, officer: false });
  });
});

describe('ownedKey', () => {
  const viewer = viewerFromMe(me());
  it('matches the logged unit by display name, ignoring case and the realm segment', () => {
    expect(ownedKey(viewer, 'Bow-Jackzon-ClassicBetaPvP-')).toBe('us/normal/bow-jackzon');
  });
  it('is null for anyone else', () => {
    expect(ownedKey(viewer, 'Simfury-ClassicBetaPvP-')).toBeNull();
  });
});

describe('coachTabListed', () => {
  const roster = [{ name: 'Simfury' }, { name: 'Bow-Jackzon-ClassicBetaPvP-' }];
  it('lists for an owner of a rostered character', () => {
    expect(coachTabListed(viewerFromMe(me()), roster)).toBe(true);
  });
  it('hides from a viewer who owns nobody in the report', () => {
    expect(coachTabListed(viewerFromMe(me()), [{ name: 'Simfury' }])).toBe(false);
  });
  it('lists for an officer', () => {
    const officer = viewerFromMe(me({ entitlements: { officer_views: true } as Me['entitlements'] }));
    expect(coachTabListed(officer, [{ name: 'Simfury' }])).toBe(true);
  });
  it('hides from the signed-out', () => {
    expect(coachTabListed(SIGNED_OUT_VIEWER, roster)).toBe(false);
  });
});
