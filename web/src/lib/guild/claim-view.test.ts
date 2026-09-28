import { describe, expect, it } from 'vitest';
import type { Me } from '../account/api';
import type { ClaimStateView, GuildHome, GuildSettingsData } from './api';
import { claimView } from './claim-view';

const ME = { user: { battletag: 'Ash#1' }, characters: [], guilds: [] } as unknown as Me;
const CLAIMED: ClaimStateView = { state: 'claimed', frozen: false };
const UNCLAIMED: ClaimStateView = { state: 'unclaimed', frozen: false };
const home = (claim: ClaimStateView): GuildHome =>
  ({
    guild: { id: 1, region: 'us', ruleset: 'pvp', name: 'X' },
    claim,
    reports: [],
    roster: [],
  }) as GuildHome;
const settings = (claim: ClaimStateView): GuildSettingsData =>
  ({
    default_visibility: 'public',
    officer_max_rank_index: 1,
    claimed_by: null,
    claim_pending: null,
    claim,
    invite: { rotated_at: null },
  }) as GuildSettingsData;

describe('claimView', () => {
  it('reads the state from the member home when settings is refused (a plain member)', () => {
    // Settings is officer-only, so for a member it is null; the home answers instead.
    const view = claimView({ me: ME, membership: { rank: 'member' }, settings: null, home: home(UNCLAIMED) });
    expect(view).toEqual({ viewer: 'member', rank: 'member', claim: UNCLAIMED });
  });

  it('prefers settings when both answered, since it is the richer read', () => {
    const view = claimView({
      me: ME,
      membership: { rank: 'officer' },
      settings: settings(CLAIMED),
      home: home(UNCLAIMED),
    });
    expect(view.viewer).toBe('eligible');
    expect(view.claim).toBe(settings(CLAIMED).claim);
  });

  it('is eligible only at officer or leader rank', () => {
    expect(claimView({ me: ME, membership: { rank: 'leader' }, settings: null, home: null }).viewer).toBe(
      'eligible',
    );
    expect(claimView({ me: ME, membership: { rank: 'member' }, settings: null, home: null }).viewer).toBe(
      'member',
    );
    expect(claimView({ me: ME, membership: {}, settings: null, home: null }).viewer).toBe('member');
  });

  it('names a signed-in visitor with no character in this guild a stranger, and no session signed-out', () => {
    expect(claimView({ me: ME, membership: null, settings: null, home: null })).toEqual({
      viewer: 'stranger',
      rank: null,
      claim: null,
    });
    expect(claimView({ me: null, membership: null, settings: null, home: null }).viewer).toBe('signed-out');
  });
});
