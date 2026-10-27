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
        onChanged: () => {},
      },
    }).body;
    expect(text(body)).toContain('Fixture#1234 offers you the guild plan.');
    expect(text(body)).toContain('billed from 1 Dec 2026');
    expect(text(body)).toContain('Nothing is charged before then.');
    expect(text(body)).toContain('Take over the plan');
    expect(text(body)).toContain('Not now');
  });
});
