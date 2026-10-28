// web/src/components/claim/claim-handoff.test.ts
// The claim page's plan line, member hand-off and subscribe button (spec 2026-10-27 §6.2, §9.3).
// svelte/server only renders pre-effect state, so GuildClaim's post-load states are held by the
// pure rules (claim-plan.ts) and the components each state mounts.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { checkoutGuildCopy } from '../../lib/billing/checkout-guild-copy';
import {
  claimLinkFor,
  claimPlanLine,
  parseClaimPlan,
  setupHrefFor,
  subscribeLabel,
} from '../../lib/guild/claim-plan';
import { guildClaimCopy } from '../../lib/guild/copy';
import { specString } from '../../test-support/spec-table';
import ClaimMemberHandoff from './ClaimMemberHandoff.svelte';
import ClaimPlanLine from './ClaimPlanLine.svelte';
import ClaimSubscribeButton from './ClaimSubscribeButton.svelte';

const SPEC = new URL('../../../../design/specs/2026-10-27-subscribe-and-join.md', import.meta.url);
const HEADING = '### 9.3';
const GUILD = 'OLYMPUS XXVII';
const PATH = { region: 'us' as const, ruleset: 'pvp' as const, slug: 'olympus-xxvii' };
const NOT_FOUNDER = false;

describe('plan line', () => {
  it('reads the spec sentence, with the plan name and yearly price', () => {
    const plan = parseClaimPlan('?plan=guild&interval=yearly');
    if (plan === null) throw new Error('plan should parse');
    const expected = specString(SPEC, HEADING, 'Claim page, with `plan`', 0, {
      Guild: GUILD,
      Plan: 'Guild',
      price: '$150 a year',
    });
    expect(claimPlanLine(GUILD, plan, NOT_FOUNDER)).toBe(expected);
  });

  it('renders the line on a gold left rule', () => {
    const { body } = render(ClaimPlanLine, { props: { line: 'x' } });
    expect(body).toContain('border-gold');
    expect(body).toContain('border-l-2');
  });

  it.each(['?plan=premium', '?plan=nope', '', '?interval=yearly'])('has no plan for %j', (search) => {
    expect(parseClaimPlan(search)).toBeNull();
  });

  it('defaults a missing or tampered interval to yearly', () => {
    expect(parseClaimPlan('?plan=raid_guild')?.interval).toBe('yearly');
    expect(parseClaimPlan('?plan=guild&interval=weekly')?.interval).toBe('yearly');
  });
});

describe('member hand-off', () => {
  it('reads the spec sentence and holds the claim URL with the same plan', () => {
    const plan = parseClaimPlan('?plan=guild&interval=yearly');
    const url = claimLinkFor(PATH, plan, 'https://pvpsixty.gg');
    expect(url).toBe('https://pvpsixty.gg/guild/us/pvp/olympus-xxvii/claim?plan=guild&interval=yearly');
    const { body } = render(ClaimMemberHandoff, { props: { guildName: GUILD, url, tier: 'guild' } });
    const sentence = specString(SPEC, HEADING, 'Claim page, member', 0, { Guild: GUILD });
    expect(guildClaimCopy.memberHandoff(GUILD)).toBe(sentence);
    expect(body).toContain(sentence);
    expect(body).toContain(url.replaceAll('&', '&amp;'));
    expect(body).toContain(checkoutGuildCopy.copyField.labelGuildMaster(GUILD));
  });

  it('carries no plan when the page was opened without one', () => {
    expect(claimLinkFor(PATH, null, 'https://pvpsixty.gg')).toBe(
      'https://pvpsixty.gg/guild/us/pvp/olympus-xxvii/claim',
    );
  });
});

describe('subscribe button', () => {
  const plan = { plan: 'guild' as const, interval: 'yearly' as const };

  it('is the trial while it is unused, else the price', () => {
    expect(subscribeLabel(plan, false, NOT_FOUNDER)).toBe(checkoutGuildCopy.ready.startTrial);
    expect(subscribeLabel(plan, true, NOT_FOUNDER)).toBe('Subscribe · $150 a year');
  });

  it('renders as the page primary with its label', () => {
    const { body } = render(ClaimSubscribeButton, {
      props: { guildId: 1, claimPlan: plan, label: 'Start the 7-day trial' },
    });
    expect(body).toContain('Start the 7-day trial');
    expect(body).toContain('data-testid="guild-claim-next-subscribe"');
  });
});

describe('get set up', () => {
  it('sets next to the claim URL, encoded', () => {
    expect(setupHrefFor('/setup', '/guild/us/pvp/olympus-xxvii/claim?plan=guild&interval=yearly')).toBe(
      '/setup?next=%2Fguild%2Fus%2Fpvp%2Folympus-xxvii%2Fclaim%3Fplan%3Dguild%26interval%3Dyearly',
    );
  });

  it('reads the spec back-link', () => {
    expect(guildClaimCopy.backToClaim(GUILD)).toBe(
      specString(SPEC, HEADING, '`/setup` after a save with `next` to a claim', 0, { Guild: GUILD }),
    );
  });
});
