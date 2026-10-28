// web/src/components/claim/claim-next.test.ts
// The Next list's order and single fill, the sign-in that keeps `next`, the plan-less share, the
// setup island's empty state, and the one guild-match rule.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { Me } from '../../lib/account/api';
import { findMyGuild } from '../../lib/guild/my-guild';
import { PRIMARY_BUTTON_FIXED } from '../../lib/planner/styles';
import SignInPrompt from '../SignInPrompt.svelte';
import ClaimMemberHandoff from './ClaimMemberHandoff.svelte';
import ClaimNextSteps from './ClaimNextSteps.svelte';
import SetupBackToClaim from './SetupBackToClaim.svelte';

const PLAN = { plan: 'guild' as const, interval: 'yearly' as const };
const BASE = {
  claimPlan: PLAN,
  guildBase: '/guild/us/pvp/olympus-xxvii',
  settingsBase: '/guild/us/pvp/olympus-xxvii/settings',
};

function testIds(body: string): string[] {
  return [...body.matchAll(/data-testid="(guild-claim-next-[a-z]+)"/g)].map((match) => match[1] ?? '');
}

const primaries = (html: string): number => html.split(PRIMARY_BUTTON_FIXED).length - 1;

describe('Next steps', () => {
  it('leads with the subscribe button, the page one fill', () => {
    const { body } = render(ClaimNextSteps, {
      props: { ...BASE, subscribe: { id: 7, label: 'Start the 7-day trial' } },
    });
    expect(testIds(body)).toEqual([
      'guild-claim-next-subscribe',
      'guild-claim-next-roster',
      'guild-claim-next-invite',
    ]);
    expect(primaries(body)).toBe(1);
    expect(body.indexOf(PRIMARY_BUTTON_FIXED)).toBeLessThan(body.indexOf('guild-claim-next-roster'));
  });

  it('keeps the roster link as the primary when there is nothing to subscribe', () => {
    const { body } = render(ClaimNextSteps, { props: { ...BASE, subscribe: null } });
    expect(testIds(body)).toEqual(['guild-claim-next-roster', 'guild-claim-next-invite']);
    expect(primaries(body)).toBe(1);
    expect(body.indexOf(PRIMARY_BUTTON_FIXED)).toBeLessThan(body.indexOf('guild-claim-next-roster'));
  });
});

describe('sign-in', () => {
  it('sends Battle.net back to the claim page', () => {
    const next = '/guild/us/pvp/olympus-xxvii/claim?plan=guild&interval=yearly';
    const { body } = render(SignInPrompt, { props: { line: 'x', next, testid: 't' } });
    expect(body).toContain(encodeURIComponent(next));
  });
});

describe('member hand-off share', () => {
  it('renders without a plan', () => {
    const { body } = render(ClaimMemberHandoff, { props: { guildName: 'G', url: 'https://x.test/claim' } });
    expect(body).toContain('https://x.test/claim');
  });
});

describe('setup return island', () => {
  it('renders nothing without a valid next', () => {
    expect(render(SetupBackToClaim).body).not.toContain('setup-back-to-claim');
  });
});

describe('findMyGuild', () => {
  const me = {
    guilds: [
      { id: 1, region: 'us', ruleset: 'pvp', name: 'OLYMPUS XXVII' },
      { id: 2, region: 'us', ruleset: 'normal', name: 'OLYMPUS XXVII' },
    ],
  } as unknown as Me;

  it('matches region, ruleset and slugified name', () => {
    expect(findMyGuild(me, { region: 'us', ruleset: 'pvp', slug: 'olympus-xxvii' })?.id).toBe(1);
    expect(findMyGuild(me, { region: 'eu', ruleset: 'pvp', slug: 'olympus-xxvii' })).toBeNull();
    expect(findMyGuild(null, { region: 'us', ruleset: 'pvp', slug: 'olympus-xxvii' })).toBeNull();
  });
});
