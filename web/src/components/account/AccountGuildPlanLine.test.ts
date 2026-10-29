// web/src/components/account/AccountGuildPlanLine.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { GuildBillingView } from '../../lib/account/api';
import AccountGuildPlanLine from './AccountGuildPlanLine.svelte';

const BASE: GuildBillingView = {
  tier: 'guild',
  you_hold_seat: false,
  scope: 'officer_tools',
  status: 'active',
  current_period_end: '2026-12-12T00:00:00Z',
  cancel_at_period_end: false,
  billed_by: '',
  you_are_billing_contact: false,
};

function renderLine(plan: GuildBillingView): { body: string; text: string } {
  const body = render(AccountGuildPlanLine, {
    props: { plan, region: 'us', ruleset: 'hardcore', name: 'The Last Watch' },
  }).body;
  const text = body
    .replace(/<!--.*?-->/g, '')
    .replace(/<[^>]+>/g, ' ')
    .replace(/\s+/g, ' ')
    .trim();
  return { body, text };
}

describe('AccountGuildPlanLine', () => {
  it('reads a Guild settings link row whose value is the scoped line', () => {
    const out = renderLine(BASE);
    expect(out.body).toContain('data-testid="account-guild-plan"');
    expect(out.text).toBe('Guild settings Officer tools · until 12 Dec 2026');
    expect(out.body).toContain('href="/guild/us/hardcore/the-last-watch/settings"');
    expect(out.body).not.toContain('underline');
  });

  it('reads the paid plan line', () => {
    const out = renderLine({
      ...BASE,
      scope: 'plan',
      current_period_end: '2026-12-31T12:00:00Z',
      billed_by: 'a#1',
    });
    expect(out.text).toBe('Guild settings Guild plan · renews 31 Dec 2026');
  });

  it('names the raid tier and the seat the account holds', () => {
    const out = renderLine({
      ...BASE,
      tier: 'raid_guild',
      scope: 'plan',
      current_period_end: '2026-12-12T12:00:00Z',
      billed_by: 'a#1',
      you_hold_seat: true,
    });
    expect(out.text).toBe('Guild settings Raid guild plan · renews 12 Dec 2026 · you hold a seat');
  });
});
