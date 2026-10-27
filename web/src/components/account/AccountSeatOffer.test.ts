import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { ME_SEAT_OFFER } from '../../test-support/billing-fixtures';
import AccountSeatOffer from './AccountSeatOffer.svelte';

function text(html: string): string {
  return html
    .replace(/<!--.*?-->/g, '')
    .replace(/<[^>]+>/g, ' ')
    .replace(/\s+/g, ' ')
    .trim();
}

describe('AccountSeatOffer', () => {
  const body = render(AccountSeatOffer, { props: { offer: ME_SEAT_OFFER, zone: 'UTC' } }).body;

  it('names the guild, the sender and what happens to their own subscription', () => {
    const out = text(body);
    expect(out).toContain('The Last Watch offers you a Premium seat.');
    expect(out).toContain('From Thornwick, 10 Nov.');
    expect(out).toContain('your own Premium stops renewing and ends 20 Dec 2026, and the seat starts then.');
    expect(out).toContain('come back at $6 a month');
  });

  it('draws both answers as buttons and focuses nothing', () => {
    expect(body).toContain('data-testid="account-seat-offer-accept"');
    expect(body).toContain('data-testid="account-seat-offer-decline"');
    expect(body).not.toContain('autofocus');
    expect(body).not.toContain('role="alert"');
  });

  it('an open offer carries the gold rule, its buttons full width below 768 with the commit first', () => {
    expect(body).toMatch(/class="flex flex-col gap-2 border-gold border-l-2 pl-3"/);
    expect(body).toContain('flex flex-col items-stretch');
    expect(body.indexOf('account-seat-offer-accept')).toBeLessThan(
      body.indexOf('account-seat-offer-decline'),
    );
  });
});
