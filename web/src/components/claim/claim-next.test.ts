// web/src/components/claim/claim-next.test.ts
// The Next list's order and single fill, the sign-in that keeps `next`, the plan-less share, the
// setup island's empty state, and the one guild-match rule.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { Me } from '../../lib/account/api';
import { findMyGuild } from '../../lib/guild/my-guild';
import { BUTTON_PRIMARY } from '../../lib/ui/button';
import SignInPrompt from '../SignInPrompt.svelte';
import ClaimMemberHandoff from './ClaimMemberHandoff.svelte';
import ClaimNextSteps from './ClaimNextSteps.svelte';
import { setupFixture } from '../../test-support/guild-setup-fixture';
import SetupBackToClaim from './SetupBackToClaim.svelte';

const PLAN = { plan: 'guild' as const, interval: 'yearly' as const };
const SETUP = setupFixture();
const BASE = {
  guildId: 7,
  guildName: 'Olympus XXVII',
  claimPlan: PLAN,
  guildBase: '/guild/us/pvp/olympus-xxvii',
  setupHref: '/guild/us/pvp/olympus-xxvii/setup',
  initial: SETUP,
};

function testIds(body: string): string[] {
  return [...body.matchAll(/data-testid="(guild-claim-next-[a-z]+)"/g)].map((match) => match[1] ?? '');
}

const primaries = (html: string): number => html.split(BUTTON_PRIMARY).length - 1;

describe('Set up the guild', () => {
  it('lists the eight steps with their states and leads with Start setup', () => {
    const { body } = render(ClaimNextSteps, { props: { ...BASE, claimPlan: null, subscribe: null } });
    expect(body).toContain('Set up Olympus XXVII');
    expect(body.match(/<li /g)?.length).toBe(8);
    expect(body).toContain('aria-label="Step 5: Done"');
    expect(body).toContain('aria-label="Step 1: Next"');
    expect(testIds(body)).toEqual(['guild-claim-next-setup', 'guild-claim-next-roster']);
    expect(primaries(body)).toBe(1);
    expect(body).toContain(`href="${BASE.setupHref}"`);
  });

  it('keeps the subscribe button as the one fill on a claim made for a plan', () => {
    const { body } = render(ClaimNextSteps, {
      props: { ...BASE, subscribe: { id: 7, label: 'Start the 7-day trial' } },
    });
    expect(testIds(body)).toEqual([
      'guild-claim-next-subscribe',
      'guild-claim-next-setup',
      'guild-claim-next-roster',
    ]);
    expect(primaries(body)).toBe(1);
    expect(body.indexOf(BUTTON_PRIMARY)).toBeLessThan(body.indexOf('guild-claim-next-setup'));
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
