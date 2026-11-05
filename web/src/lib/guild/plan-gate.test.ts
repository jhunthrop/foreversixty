import { describe, expect, it } from 'vitest';
import communityLinks from '../../data/links.json';
import { gateAction, gateFactsFrom, type GateFacts } from './plan-gate';

const SETTINGS = '/guild/us/pvp/olympus/settings';
const CLAIM = '/guild/us/pvp/olympus/claim';
const HOLDER = { name: 'Thornhide', class: 'druid', character_key: 'us/pvp/thornhide' };

const facts = (over: Partial<GateFacts> = {}): GateFacts => ({
  claim: { state: 'claimed', yours: true, claimed_by: HOLDER },
  guildMaster: true,
  trialUsed: false,
  lapsed: false,
  guildTier: false,
  granted: false,
  settingsHref: SETTINGS,
  claimHref: CLAIM,
  ...over,
});

describe('gateAction, the spec table', () => {
  it('claim holder, no plan, trial unused: the trial, with Raid guild preselected from a Raid guild feature', () => {
    expect(gateAction('raids', facts())).toEqual({
      kind: 'link',
      label: 'Start the 7-day trial',
      href: `${SETTINGS}#billing`,
    });
    expect(gateAction('discord', facts())).toMatchObject({ href: `${SETTINGS}#billing` });
    expect(gateAction('loot', facts())).toEqual({
      kind: 'link',
      label: 'Start the 7-day trial',
      href: `${SETTINGS}?plan=raid_guild#billing`,
    });
    expect(gateAction('readiness', facts())).toMatchObject({ href: `${SETTINGS}?plan=raid_guild#billing` });
  });

  it('claim holder, trial used or lapsed: subscribe; an unknown trial never promises one', () => {
    for (const over of [{ trialUsed: true }, { lapsed: true }, { trialUsed: null }]) {
      expect(gateAction('raids', facts(over))).toEqual({
        kind: 'link',
        label: 'Subscribe your guild',
        href: `${SETTINGS}#billing`,
      });
    }
    expect(gateAction('loot', facts({ trialUsed: true }))).toMatchObject({
      href: `${SETTINGS}?plan=raid_guild#billing`,
    });
  });

  it('claim holder on Guild, a Raid guild feature: the Move confirm', () => {
    expect(gateAction('loot', facts({ guildTier: true }))).toEqual({
      kind: 'link',
      label: 'Move to Raid guild',
      href: `${SETTINGS}?upgrade=1#billing`,
    });
    expect(gateAction('readiness', facts({ guildTier: true }))).toMatchObject({
      label: 'Move to Raid guild',
    });
  });

  it('claim holder on a granted Guild plan, a Raid guild feature: ask us, never a dead Move link (guild round 3)', () => {
    const granted = facts({ guildTier: true, granted: true });
    for (const feature of ['loot', 'readiness'] as const) {
      expect(gateAction(feature, granted)).toEqual({
        kind: 'contact',
        before: 'Your plan was granted by Forever Sixty. ',
        link: 'Ask us on Discord',
        after: ' to move it to Raid guild.',
        href: communityLinks.discordInvite,
      });
    }
    // A granted plan's Guild features are not gated: Raids reads its usual action.
    expect(gateAction('raids', granted)).toMatchObject({ kind: 'link' });
  });

  it('an officer when someone else holds the claim: ask them, whatever the billing says', () => {
    const claim = { state: 'claimed' as const, yours: false, claimed_by: HOLDER };
    expect(gateAction('loot', facts({ claim, guildMaster: false }))).toEqual({ kind: 'ask', holder: HOLDER });
    expect(gateAction('raids', facts({ claim, guildMaster: false, guildTier: true }))).toEqual({
      kind: 'ask',
      holder: HOLDER,
    });
  });

  it('a holder with no character on the roster is asked for by role', () => {
    const claim = { state: 'contested' as const, yours: false, claimed_by: null };
    expect(gateAction('raids', facts({ claim }))).toEqual({ kind: 'ask', holder: null });
  });

  it('unclaimed or pending: the guild master claims, an officer is told so', () => {
    for (const state of ['unclaimed', 'pending'] as const) {
      const claim = { state, yours: false, claimed_by: null };
      expect(gateAction('loot', facts({ claim }))).toEqual({
        kind: 'link',
        label: 'Claim the guild first',
        href: CLAIM,
      });
      expect(gateAction('loot', facts({ claim, guildMaster: false }))).toEqual({
        kind: 'note',
        text: 'Your guild master claims the guild first.',
      });
    }
  });
});

describe('gateFactsFrom', () => {
  const input = {
    claim: null,
    membership: null,
    officerTools: true,
    raidGuildTools: false,
    settingsHref: SETTINGS,
    claimHref: CLAIM,
  };

  it('reads the Guild plan off the two tools flags', () => {
    expect(gateFactsFrom(input).guildTier).toBe(true);
    expect(gateFactsFrom({ ...input, officerTools: false }).guildTier).toBe(false);
    expect(gateFactsFrom({ ...input, raidGuildTools: true }).guildTier).toBe(false);
  });

  it('reads rank, trial and lapse off the /v1/me entry', () => {
    const got = gateFactsFrom({
      ...input,
      membership: { rank: 'leader', covered: { state: 'lapsed' }, subscribe: { trial_used: true } },
    });
    expect(got).toMatchObject({ guildMaster: true, lapsed: true, trialUsed: true });
    expect(gateFactsFrom(input)).toMatchObject({ guildMaster: false, lapsed: false, trialUsed: null });
  });

  it('reads a grant off an empty billing contact on a live plan, never on a lapsed or paid one', () => {
    const granted = (covered: { state: 'active' | 'lapsed'; billed_by?: string } | null) =>
      gateFactsFrom({ ...input, membership: { rank: 'leader', covered } }).granted;
    expect(granted({ state: 'active', billed_by: '' })).toBe(true);
    expect(granted({ state: 'active', billed_by: 'Thornhide' })).toBe(false);
    expect(granted({ state: 'lapsed', billed_by: '' })).toBe(false);
    expect(granted({ state: 'active' })).toBe(false);
    expect(granted(null)).toBe(false);
  });
});
