// web/src/components/account/AccountPremiumBlock.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { EntitlementsView, MeCharacter } from '../../lib/account/api';
import type { AccountPlanLine } from '../../lib/account/plan-line';
import { mockEntitlements } from '../../test-support/billing-fixtures';
import AccountPremiumBlock from './AccountPremiumBlock.svelte';

const CHARACTER: MeCharacter = {
  key: 'us/hardcore/elyra-duskvale',
  region: 'us',
  ruleset: 'hardcore',
  name: 'Elyra Duskvale',
  class: 'Priest',
  level: 60,
};

function text(html: string): string {
  return html
    .replace(/<!--.*?-->/g, '')
    .replace(/<[^>]+>/g, ' ')
    .replace(/\s+/g, ' ');
}

function renderBlock(entitlements?: EntitlementsView, characters: MeCharacter[] = []): string {
  return render(AccountPremiumBlock, { props: { entitlements, characters } }).body;
}

describe('AccountPremiumBlock', () => {
  it('has no subheading of its own: the plan line above names the plan', () => {
    const html = render(AccountPremiumBlock, {
      props: { entitlements: mockEntitlements(), notifications: null },
    }).body;
    expect(text(html)).not.toContain('Premium');
  });

  it('shows the live values when the entitlements are on', () => {
    const body = text(
      renderBlock(mockEntitlements({ server_sims: true, supporter_mark: true }), [CHARACTER]),
    );
    expect(body).toContain('Server sims On');
    expect(body).toContain('Supporter mark Shown on your characters');
  });

  it('carries the month count into the server sims row', () => {
    const html = render(AccountPremiumBlock, {
      props: { entitlements: mockEntitlements({ server_sims: true }), serverRunsThisMonth: 4 },
    }).body;
    expect(text(html)).toContain('Server sims On · 4 runs this month');
  });

  it.each([
    [1, 'Server sims On · 1 run this month'],
    [0, 'Server sims On · no runs yet this month'],
    [undefined, 'Server sims On '],
  ])('reads %s runs as %s', (count, expected) => {
    const html = render(AccountPremiumBlock, {
      props: { entitlements: mockEntitlements({ server_sims: true }), serverRunsThisMonth: count },
    }).body;
    expect(text(html)).toContain(expected);
    expect(text(html)).not.toContain(count === undefined ? 'runs' : 'In your browser');
  });

  it('keeps the free words when server sims are off, whatever count is sent', () => {
    const html = render(AccountPremiumBlock, {
      props: { entitlements: mockEntitlements(), serverRunsThisMonth: 4 },
    }).body;
    expect(text(html)).toContain('Server sims In your browser');
  });

  it('shows the free values by default', () => {
    const body = text(renderBlock(mockEntitlements()));
    expect(body).toContain('Server sims In your browser');
    expect(body).toContain('Supporter mark Not shown');
  });

  it('shows the Discord row as muted Premium only by default, never linked', () => {
    const html = renderBlock(mockEntitlements({ server_sims: true, supporter_mark: true }), [CHARACTER]);
    expect(text(html)).toContain('Discord notifications Premium only');
    expect(html).not.toContain('account-premium-notifications-link');
  });

  it('leaves the Discord row out when notifications is null (the deployment has it off)', () => {
    const html = render(AccountPremiumBlock, {
      props: { entitlements: mockEntitlements(), notifications: null },
    }).body;
    expect(html).not.toContain('account-premium-notifications');
    expect(text(html)).not.toContain('Discord notifications');
    expect(html).toContain('account-premium-supporter');
  });

  it('links the Discord row to #discord with the value it is given', () => {
    const html = render(AccountPremiumBlock, {
      props: { entitlements: mockEntitlements(), notifications: { value: 'On · 2 events', linked: true } },
    }).body;
    expect(text(html)).toContain('Discord notifications On · 2 events');
    expect(html).toMatch(/href="#discord"[^>]*account-premium-notifications-link/);
  });

  describe('history row', () => {
    it('links to the main character history when entitled with a main', () => {
      const html = renderBlock(mockEntitlements({ history: true }), [CHARACTER]);
      expect(text(html)).toContain('History charts Open on your main character');
      expect(html).toMatch(
        /href="\/character\/us\/hardcore\/elyra-duskvale#history"[^>]*account-premium-history-link/,
      );
    });

    it('reads plain On when entitled with no main', () => {
      const html = renderBlock(mockEntitlements({ history: true }), []);
      expect(text(html)).toContain('History charts On ');
      expect(html).not.toContain('account-premium-history-link');
    });

    it('reads muted Sparkline only when not entitled', () => {
      const html = renderBlock(mockEntitlements({ history: false }), [CHARACTER]);
      expect(text(html)).toContain('History charts Sparkline only');
      expect(html).toMatch(/text-muted[^>]*data-testid="settings-row-value">Sparkline only</);
      expect(html).not.toContain('account-premium-history-link');
    });
  });

  it('links a live value to where the benefit lives', () => {
    const html = renderBlock(mockEntitlements({ server_sims: true, supporter_mark: true }), [CHARACTER]);
    expect(html).toMatch(/href="\/sim"[^>]*data-testid="account-premium-serverSims-link"/);
    expect(html).toMatch(
      /href="\/character\/us\/hardcore\/elyra-duskvale"[^>]*account-premium-supporter-link/,
    );
  });

  it('keeps a live supporter value plain text when the account has no character', () => {
    const html = renderBlock(mockEntitlements({ supporter_mark: true }), []);
    expect(text(html)).toContain('Shown on your characters');
    expect(html).not.toContain('account-premium-supporter-link');
  });

  it('never links a free value', () => {
    const html = renderBlock(mockEntitlements(), [CHARACTER]);
    expect(html).not.toContain('-link');
  });

  it('renders the free state for an older /v1/me with no entitlements at all', () => {
    const html = renderBlock(undefined);
    expect(html).toContain('data-testid="account-premium-block"');
    expect(text(html)).toContain('Server sims In your browser');
    expect(text(html)).toContain('Supporter mark Not shown');
  });
});

describe('AccountPremiumBlock plan headline', () => {
  const PLAN: AccountPlanLine = {
    head: { name: 'Premium', when: { label: 'Renews', date: '20 Dec 2026' }, extra: '$6 a month' },
    line: 'Premium · renews 20 Dec 2026 · $6 a month',
    line2: null,
    founder: false,
    billingButton: 'manage',
    ended: null,
  };
  const renderPlan = (plan: AccountPlanLine, seePlans = false): string =>
    render(AccountPremiumBlock, { props: { plan, seePlans, notifications: null } }).body;

  it('draws no headline without a plan', () => {
    expect(renderBlock(mockEntitlements())).not.toContain('account-plan-headline');
  });

  it('reads the plan line and a Founder pill with a title and a screen-reader tail', () => {
    const html = renderPlan({ ...PLAN, founder: true });
    expect(text(html)).toContain('Premium · renews 20 Dec 2026 · $6 a month');
    expect(html).toContain('data-testid="account-founder"');
    expect(html.replace(/<!--[^>]*-->/g, '')).toContain('<h2 class="heading-plan">Premium</h2>');
    expect(text(html)).toContain('Renews 20 Dec 2026');
    expect(html).toContain('title="Founder price, kept while your plan stays active."');
    expect(html).toMatch(/class="sr-only"> ?Founder price, kept while your plan stays active\./);
    expect(renderPlan(PLAN)).not.toContain('account-founder');
  });

  it('says cancel anytime under a renewing own plan, and not when it is ending', () => {
    expect(text(renderPlan(PLAN))).toContain('Cancel anytime in Manage billing.');
    const ending = { ...PLAN, head: { ...PLAN.head, when: { label: 'Ends', date: '20 Dec 2026' } } };
    expect(text(renderPlan(ending))).not.toContain('Cancel anytime');
    expect(text(renderPlan({ ...PLAN, billingButton: null }))).not.toContain('Cancel anytime');
  });

  it('the free state links to the plans', () => {
    const html = renderPlan(
      { ...PLAN, line: 'Not on Premium.', head: { name: 'Free', when: null, extra: null } },
      true,
    );
    expect(html).toMatch(/<a class="[^"]*bg-\[linear-gradient[^"]*" href="\/premium#premium"/);
    expect(renderPlan(PLAN)).not.toContain('href="/premium');
  });

  it('a seat holder reads who gave it', () => {
    const html = renderPlan({ ...PLAN, line: 'Premium · a seat from The Last Watch', line2: 'Given by X.' });
    expect(text(html)).toContain('Given by X.');
  });

  it('an ended seat links its resubscribe URL', () => {
    const ended = {
      text: 'Your seat ended 25 Nov. ',
      action: 'Come back at $6 a month',
      href: '/premium/checkout?comeback=9',
    };
    const html = renderPlan({ ...PLAN, ended });
    expect(html).toContain('href="/premium/checkout?comeback=9"');
    expect(text(html)).toContain('Your seat ended 25 Nov. Come back at $6 a month');
  });

  it('an ended seat takes the plan line place: no Not on Premium, no See plans, a gold link', () => {
    const ended = {
      text: 'Your seat ended 18 Nov. ',
      action: 'Come back',
      href: '/premium/checkout?comeback=9',
    };
    const html = renderPlan({ ...PLAN, line: 'Not on Premium.', ended }, true);
    expect(html).not.toContain('account-billing-row');
    expect(html).not.toContain('href="/premium"');
    expect(html).toMatch(/<a class="text-gold[^"]*" href="\/premium\/checkout\?comeback=9"/);
  });
});
