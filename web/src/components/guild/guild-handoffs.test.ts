// web/src/components/guild/guild-handoffs.test.ts
// Task 21: the guild page, invite and settings hand-off lines of the subscribe-and-join spec.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { specString } from '../../test-support/spec-table';
import type { GuildBillingSettings, GuildCoveredView } from '../../lib/billing/contract';
import { billingState } from '../../lib/guild/billing-state';
import {
  billingHandoffUrl,
  chooserPreselect,
  parseBillingPreselect,
  shouldOpenUpgrade,
  showsBillingHandoff,
} from '../../lib/guild/billing-handoff';
import { guildHomeCopy, guildJoinCopy, guildSettingsCopy } from '../../lib/guild/copy';
import { joinedPlanLine } from '../../lib/guild/join-plan';
import { claimEntryFor, claimEntryNotes } from '../../lib/guild/roster-gate';
import GuildJoinedBlock from './GuildJoinedBlock.svelte';
import BillingHandoff from './BillingHandoff.svelte';
import GuildUnverifiedNote from './GuildUnverifiedNote.svelte';

const SPEC = new URL('../../../../design/specs/2026-10-27-subscribe-and-join.md', import.meta.url);
const row = (label: string, index = 0, fill: Record<string, string> = {}): string =>
  specString(SPEC, '### 9.3', label, index, fill);
const G = 'OLYMPUS XXVII';
const PATH = { region: 'us', ruleset: 'pvp', slug: 'olympus-xxvii' } as const;
const NOW = new Date('2026-10-27T12:00:00Z');

function view(over: Partial<GuildBillingSettings> = {}): GuildBillingSettings {
  return {
    billing: null,
    trial_used: false,
    verified_count: 12,
    seats: [],
    transfer: null,
    transfer_candidates: [],
    can: [],
    ...over,
  };
}

const livePlan = {
  tier: 'guild',
  status: 'active',
  scope: 'plan',
  billed_by: 'Someone',
  cancel_at_period_end: false,
  current_period_end: '2027-01-01T00:00:00Z',
  trial_end: null,
  amount_cents: 12000,
  interval: 'yearly',
} as unknown as GuildBillingSettings['billing'];

describe('copy pinned to spec 9.3', () => {
  it('unverified note without an approver', () => {
    expect(guildHomeCopy.notVerifiedNoApprover(G)).toBe(
      row('Guild page, unverified, no approver', 0, { Guild: G }),
    );
    expect(guildHomeCopy.sendGuildMasterLink).toBe(row('Guild page, unverified, no approver', 1));
  });

  it('claim slot, settings and invite lines', () => {
    expect(claimEntryNotes.getClaimed).toBe(row('Guild page claim slot, member of an unclaimed guild'));
    expect(guildSettingsCopy.sendThemThisLink).toBe(row('Settings, not the claim holder, no plan'));
    expect(guildJoinCopy.onPlan(G, 'Guild')).toBe(row('Invite joined', 0, { Guild: G, Tier: 'Guild' }));
    expect(guildJoinCopy.noPlanYet(G)).toBe(row('Invite joined', 1, { Guild: G }));
  });
});

describe('GuildUnverifiedNote', () => {
  const props = { guildName: G, claimHref: '/guild/us/pvp/olympus-xxvii/claim' };

  it("says today's sentence with an approver or an unknown answer", () => {
    for (const hasApprover of [true, undefined]) {
      const body = render(GuildUnverifiedNote, { props: { ...props, hasApprover } }).body;
      expect(body).toContain(guildHomeCopy.notVerifiedNote);
      expect(body).not.toContain('guild-home-send-claim-link');
    }
  });

  it('says nobody can approve and links the claim page without one', () => {
    const body = render(GuildUnverifiedNote, { props: { ...props, hasApprover: false } }).body;
    expect(body).toContain(guildHomeCopy.notVerifiedNoApprover(G));
    expect(body).toContain('href="/guild/us/pvp/olympus-xxvii/claim"');
    expect(body).toContain(guildHomeCopy.sendGuildMasterLink);
  });
});

describe('the claim slot', () => {
  it('reads Get the guild claimed for a member of an unclaimed guild', () => {
    const entry = claimEntryFor({
      signedIn: true,
      canManage: false,
      membership: { rank: 'member' },
      claimState: 'unclaimed',
      claimedBy: null,
    });
    expect(entry).toEqual({ kind: 'getClaimed' });
  });
});

describe('joinedPlanLine', () => {
  const covered = (over: Partial<GuildCoveredView>): GuildCoveredView => ({
    tier: 'raid_guild',
    state: 'active',
    ends_at: null,
    billed_by: 'x',
    you_pay: false,
    you_hold_seat: false,
    seats_given: 0,
    seats_total: 3,
    ...over,
  });

  it('names the tier, or says no plan yet', () => {
    expect(joinedPlanLine(G, covered({}))).toBe(`${G} is on the Raid guild plan.`);
    expect(joinedPlanLine(G, null)).toBe(`${G} has no plan yet.`);
    expect(joinedPlanLine(G, covered({ state: 'lapsed' }))).toBe(`${G} has no plan yet.`);
  });
});

describe('BillingHandoff', () => {
  it('uses fitsTier: 12 verified is guild, 28 is raid_guild, yearly', () => {
    expect(billingHandoffUrl(PATH, 12)).toBe(
      'https://foreversixty.gg/guild/us/pvp/olympus-xxvii/settings?plan=guild&interval=yearly#billing',
    );
    expect(billingHandoffUrl(PATH, 28)).toContain('plan=raid_guild&interval=yearly');
  });

  it('renders the line and a copy field holding the link', () => {
    const body = render(BillingHandoff, { props: { path: PATH, guildName: G, verifiedCount: 28 } }).body;
    expect(body).toContain('Send them this link.');
    expect(body).toContain('plan=raid_guild');
  });

  it('draws only for a read-only reader of a guild with no live plan', () => {
    const none = view();
    expect(showsBillingHandoff(none, billingState(none, NOW))).toBe(true);
    const payer = view({ can: ['billing'] });
    expect(showsBillingHandoff(payer, billingState(payer, NOW))).toBe(false);
    const planned = view({ billing: livePlan });
    expect(showsBillingHandoff(planned, billingState(planned, NOW))).toBe(false);
  });
});

describe('settings preselect', () => {
  it('reads plan, interval and upgrade; drops anything else', () => {
    expect(parseBillingPreselect('?plan=raid_guild&interval=monthly')).toEqual({
      tier: 'raid_guild',
      interval: 'monthly',
      upgrade: false,
    });
    expect(parseBillingPreselect('?plan=premium&interval=weekly&upgrade=1')).toEqual({
      tier: null,
      interval: null,
      upgrade: true,
    });
  });

  it('opens the upgrade confirm once, for a payer on a Guild plan', () => {
    const payer = view({ can: ['billing'], billing: livePlan });
    const state = billingState(payer, NOW);
    expect(shouldOpenUpgrade(true, payer, state, false)).toBe(true);
    expect(shouldOpenUpgrade(true, payer, state, true)).toBe(false);
    expect(shouldOpenUpgrade(false, payer, state, false)).toBe(false);
    expect(shouldOpenUpgrade(true, view(), billingState(view(), NOW), false)).toBe(false);
  });
});

describe('chooser preselect', () => {
  // The section's effects do not run under svelte/server, so the props it hands the chooser are
  // pinned through this pure helper.
  it('turns ?plan=raid_guild&interval=monthly into the chooser start state', () => {
    expect(chooserPreselect(parseBillingPreselect('?plan=raid_guild&interval=monthly'))).toEqual({
      initialTier: 'raid_guild',
      initialInterval: 'monthly',
    });
    expect(chooserPreselect(parseBillingPreselect(''))).toEqual({
      initialTier: null,
      initialInterval: 'yearly',
    });
  });
});

describe('GuildJoinedBlock', () => {
  it('shows the joined sentence, the plan line and Open the guild, with no join button', () => {
    const body = render(GuildJoinedBlock, {
      props: { planLine: `${G} has no plan yet.`, openHref: '/guild/us/pvp/olympus-xxvii' },
    }).body;
    expect(body).toContain('role="status"');
    expect(body).toContain(guildJoinCopy.joined);
    expect(body).toContain(`${G} has no plan yet.`);
    expect(body).toContain('href="/guild/us/pvp/olympus-xxvii"');
    expect(body).toContain(guildJoinCopy.openGuild);
    expect(body).not.toContain('guild-join-button');
  });

  it('reserves the plan line while the re-read is in flight', () => {
    const body = render(GuildJoinedBlock, { props: { planLine: null, openHref: '/g' } }).body;
    expect(body).toContain('min-h-[21px]');
    expect(body).toContain('data-testid="guild-join-plan"');
  });
});
