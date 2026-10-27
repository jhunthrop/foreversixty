import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { TRANSFER_RECEIVED, planView } from '../../../test-support/billing-fixtures';
import { text } from '../../../test-support/raid-fixtures';
import GuildTransferOffer from './GuildTransferOffer.svelte';

describe('GuildTransferOffer', () => {
  it('states the offer with the sender, first charge and price, and both buttons', () => {
    const body = render(GuildTransferOffer, {
      props: {
        guildId: 7,
        plan: planView(),
        transfer: { ...TRANSFER_RECEIVED, starts_at: '2026-12-01T12:00:00Z' },
        onChanged: () => Promise.resolve(),
      },
    }).body;
    expect(text(body)).toContain('Thornwick offers you the guild plan.');
    expect(text(body)).toContain('billed from 1 Dec 2026');
    expect(text(body)).toContain('Nothing is charged before then.');
    expect(text(body)).toContain('Take over the plan');
    expect(text(body)).toContain('Not now');
    // The receiver reads what stays with the guild about themselves (persona round 2).
    expect(text(body)).toContain('Seats, raids and records stay with the guild; you manage seats from then.');
    expect(text(body)).not.toContain('Mirelle manages');
  });
});
