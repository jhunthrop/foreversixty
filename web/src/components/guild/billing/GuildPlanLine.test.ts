import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { billingState, type BillingViewState } from '../../../lib/guild/billing-state';
import {
  FOUNDER_RAID_GUILD_READ,
  GUILD_TIER_28_READ,
  TRIAL_READ,
  guildBillingRead,
  planView,
} from '../../../test-support/billing-fixtures';
import { text } from '../../../test-support/raid-fixtures';
import type { GuildBillingSettings } from '../../../lib/billing/contract';
import GuildPlanLine from './GuildPlanLine.svelte';

const NOW = new Date('2026-11-15T12:00:00Z');
const NOON = { current_period_end: '2026-12-01T12:00:00Z' };

/** The fixtures sit at midnight UTC, which reads as the day before in a US zone; noon is one day everywhere. */
function atNoon(
  read: GuildBillingSettings,
  overrides: Partial<NonNullable<GuildBillingSettings['billing']>> = {},
): GuildBillingSettings {
  return { ...read, billing: { ...read.billing!, ...NOON, ...overrides } };
}

function html(
  read: GuildBillingSettings,
  payer: boolean,
  showMove = false,
  state?: BillingViewState,
): string {
  return render(GuildPlanLine, {
    props: {
      guildId: 7,
      plan: read.billing!,
      lineState: state ?? billingState(read, NOW),
      payer,
      showMove,
      onMove: () => {},
    },
  }).body;
}

describe('GuildPlanLine', () => {
  it('reads the trial line and its cancel-before date', () => {
    const body = text(html(atNoon(TRIAL_READ, { trial_end: '2026-11-04T12:00:00Z' }), true));
    expect(body).toContain('Guild · trial · first charge 4 Nov 2026: $15');
    expect(body).toContain('Cancel before 4 Nov and nothing is charged. Billed to you.');
  });

  it('reads the active line with the Founder tag', () => {
    const body = html(atNoon(FOUNDER_RAID_GUILD_READ), true);
    expect(text(body)).toContain('Raid guild · renews 1 Dec 2026 · $200 a year');
    expect(body).toContain('guild-billing-founder');
    // The Founder state is a Tag in the plan row's trailing slot (site alignment, settings Plan card).
    expect(body).toMatch(/data-testid="guild-billing-founder" data-tone="gold"[^>]*>[\s\S]*Founder/);
  });

  it('gives a payer Manage billing and Move to Raid guild; a reader neither', () => {
    const payer = html(GUILD_TIER_28_READ, true, true);
    expect(payer).toContain('guild-manage-billing');
    expect(payer).toContain('guild-move-raid-guild');
    const reader = html(guildBillingRead({ billing: planView({ you_are_billing_contact: false }) }), false);
    expect(reader).not.toContain('guild-manage-billing');
    expect(text(reader)).toContain('Billed by Thornwick');
  });

  it('a lapsed plan reads the read-only line and has no Manage billing', () => {
    const read = guildBillingRead({ billing: planView({ status: 'canceled', ...NOON }) });
    const body = html(read, true);
    expect(text(body)).toContain('Guild plan ended 1 Dec 2026.');
    expect(text(body)).toContain('Officer tools are read-only.');
    expect(body).not.toContain('guild-manage-billing');
  });
});
