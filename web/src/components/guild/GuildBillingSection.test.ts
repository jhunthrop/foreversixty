// web/src/components/guild/GuildBillingSection.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { GuildBillingView } from '../../lib/account/api';
import { notOnPlanLines } from '../../lib/guild/settings-view';
import GuildBillingSection from './GuildBillingSection.svelte';

const SCOPED: GuildBillingView = {
  scope: 'officer_tools',
  status: 'active',
  current_period_end: '2026-12-11T00:00:00Z',
  cancel_at_period_end: false,
  billed_by: '',
  you_are_billing_contact: false,
};
const PLAN: GuildBillingView = {
  ...SCOPED,
  scope: 'plan',
  current_period_end: '2026-12-31T00:00:00Z',
  billed_by: 'hunthrop#1894',
  you_are_billing_contact: true,
};
const EXPLANATION =
  'Granted by Forever Sixty for Award, Undo and Nudge only. It ends on 11 Dec 2026; the guild plan keeps them after that.';
const MASTER_ONLY = 'Only the guild master can release the claim or change billing.';

function html(billing: GuildBillingView | null, mayBill: boolean): string {
  return render(GuildBillingSection, { props: { billing, mayBill, guildId: 7 } }).body.replace(
    /<!--.*?-->/g,
    '',
  );
}

function htmlUnavailable(mayBill: boolean): string {
  return render(GuildBillingSection, {
    props: { billing: PLAN, mayBill, guildId: 7, unavailable: true },
  }).body.replace(/<!--.*?-->/g, '');
}

function slot(body: string, testid: string): string {
  const match = new RegExp(`data-testid="${testid}"[^>]*>([^<]*)<`).exec(body);
  return (match?.[1] ?? '').trim();
}

describe('GuildBillingSection, scoped grant for the claim holder', () => {
  const body = html(SCOPED, true);

  it('draws the plan line, the explanation and the price line', () => {
    expect(slot(body, 'guild-billing-plan')).toBe('Officer tools · until 11 Dec 2026');
    expect(slot(body, 'guild-billing-contact')).toBe(EXPLANATION);
    expect(slot(body, 'guild-billing-price')).toBe(notOnPlanLines().detail);
  });

  it('offers Subscribe the guild with the checkout href and no Manage billing', () => {
    expect(body).toContain('href="/premium/checkout?plan=guild&amp;interval=monthly&amp;guild_id=7"');
    expect(body).toContain('Subscribe the guild');
    expect(body).not.toContain('guild-manage-billing');
    expect(body).not.toContain(MASTER_ONLY);
  });
});

describe('GuildBillingSection, an officer without the claim', () => {
  it('shows the plan line and explanation, then the master-only sentence, no button', () => {
    const body = html(SCOPED, false);
    expect(slot(body, 'guild-billing-plan')).toBe('Officer tools · until 11 Dec 2026');
    expect(slot(body, 'guild-billing-contact')).toBe(EXPLANATION);
    expect(slot(body, 'guild-settings-master-only')).toBe(MASTER_ONLY);
    expect(body).not.toContain('guild-subscribe-link');
    expect(body).not.toContain('guild-manage-billing');
    expect(body).not.toContain('guild-billing-price');
  });

  it('shows the full plan line and billed line for a paid plan', () => {
    const body = html(PLAN, false);
    expect(slot(body, 'guild-billing-plan')).toBe('Guild plan · renews 31 Dec 2026');
    expect(slot(body, 'guild-billing-contact')).toBe('Billed to you.');
    expect(body).not.toContain('guild-manage-billing');
    expect(slot(body, 'guild-settings-master-only')).toBe(MASTER_ONLY);
  });
});

describe('GuildBillingSection, full plan for the claim holder', () => {
  it('keeps Manage billing and no master-only sentence', () => {
    const body = html(PLAN, true);
    expect(body).toContain('guild-manage-billing');
    expect(body).not.toContain(MASTER_ONLY);
  });
});

describe('GuildBillingSection, billing unavailable', () => {
  it('offers retry and the unavailable alert but no Manage billing, even for the claim holder', () => {
    const body = htmlUnavailable(true);
    expect(body).toContain('guild-billing-retry');
    expect(body).toContain('guild-billing-unavailable');
    expect(body).not.toContain('guild-manage-billing');
  });
});
