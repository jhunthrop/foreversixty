// web/src/components/checkout/CheckoutBand.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { checkoutGuildCopy as copy } from '../../lib/billing/checkout-guild-copy';
import { CHECKOUT_BAND_RESERVE } from '../../lib/billing/checkout-layout';
import type { Interval, PlanKey } from '../../lib/billing/plans';
import CheckoutBand from './CheckoutBand.svelte';

function html(plan: PlanKey | null, interval: Interval = 'yearly', founder = true, invalid = false): string {
  return render(CheckoutBand, { props: { plan, interval, founder, invalid } }).body.replace(
    /<!--.*?-->/g,
    '',
  );
}

/** The text of the element carrying `testid`, as the browser lays it out inline. */
function textOf(body: string, testid: string): string {
  const open = body.indexOf(`data-testid="${testid}"`);
  const start = body.indexOf('>', open) + 1;
  let depth = 1;
  let index = start;
  while (depth > 0) {
    const next = body.indexOf('<', index);
    depth += body.startsWith('</', next) ? -1 : 1;
    index = body.indexOf('>', next) + 1;
  }
  return text(body.slice(start, body.lastIndexOf('<', index - 1)));
}

function text(body: string): string {
  return body.replace(/<[^>]+>/g, '');
}

describe('CheckoutBand', () => {
  it('Guild yearly founder: the band, the h1 and the whole plan summary', () => {
    const body = html('guild');
    expect(body).toContain('full-bleed');
    expect(body).toContain('page-column');
    expect(body).toMatch(/<h1[^>]*>Subscribe a guild<\/h1>/);
    expect(body).toContain('>Guild<');
    expect(body).toContain('>$120<');
    expect(body).toContain('> a year<');
    expect(body).toContain('>Founder<');
    expect(body).toContain('pill-site');
    expect(text(body)).toContain(copy.sizeLine.guild);
    expect(body).toContain(`>${copy.planSummary.trial}<`);
    expect(copy.planSummary.trial).toBe('First 7 days free');
    expect(body).toMatch(/<a[^>]*href="\/premium"[^>]*>Change plan<\/a>/);
  });

  it('the price and its unit share one inline box and read "$120 a year"', () => {
    const body = html('guild');
    expect(textOf(body, 'checkout-plan-price')).toBe('$120 a year');
    expect(textOf(html('raid_guild', 'monthly', false), 'checkout-plan-price')).toBe('$25 a month');
  });

  it('the price is the body face with tabular numerals (type pass: no code face on UI numbers)', () => {
    expect(html('guild')).toMatch(/class="[^"]*numeric[^"]*"[^>]*>\$120</);
  });

  it('Raid guild renders its own size line and price, not the Guild one', () => {
    const body = text(html('raid_guild'));
    expect(body).toContain(copy.sizeLine.raidGuild);
    expect(body).not.toContain(copy.sizeLine.guild);
    expect(body).toContain('Raid guild');
    expect(body).toContain('$200');
  });

  it('monthly after the founder window: list price, a month, no Founder pill', () => {
    const body = html('guild', 'monthly', false);
    expect(body).toContain('>$15<');
    expect(body).toContain('> a month<');
    expect(body).not.toContain('Founder');
  });

  it('the CHECKOUT label is gold and hidden below 768', () => {
    const label = html('guild').match(/<p class="([^"]*)"[^>]*>(?:(?!<\/p>).)*CHECKOUT/s);
    expect(label?.[1].split(' ')).toEqual(
      expect.arrayContaining(['label', 'text-gold', 'hidden', 'md:block']),
    );
  });

  it('Premium: h1 Premium and no guild plan summary (Task 12 fills the body)', () => {
    const body = text(html('premium'));
    expect(body).toContain('Premium');
    expect(body).not.toContain('Subscribe a guild');
    expect(body).not.toContain(copy.planSummary.changePlan);
  });

  it('an invalid link: exactly one h1 Checkout, no plan summary', () => {
    const body = html(null, 'yearly', true, true);
    expect(body.match(/<h1/g)).toHaveLength(1);
    expect(body).toMatch(new RegExp(`<h1[^>]*>${copy.band.invalidTitle}</h1>`));
    expect(copy.band.invalidTitle).toBe('Checkout');
    expect(body).not.toContain('checkout-plan-summary');
    expect(body).toContain(CHECKOUT_BAND_RESERVE);
  });

  it('before the URL is read: the static frame carries the LCP size line, marked pending', () => {
    const body = html(null);
    expect(body).toContain('full-bleed');
    expect(body).toContain(CHECKOUT_BAND_RESERVE);
    expect(body).toContain('checkout-pending');
    expect(textOf(body, 'checkout-size-line')).toBe(copy.sizeLine.guild);
    expect(body).not.toContain('checkout-plan-price');
  });

  it('a resolved plan is not marked pending', () => {
    for (const plan of ['guild', 'premium'] as const) expect(html(plan)).not.toContain('checkout-pending');
  });

  it('reserves the same height whatever the plan', () => {
    for (const plan of ['guild', 'raid_guild', 'premium', null] as const) {
      expect(html(plan)).toContain(CHECKOUT_BAND_RESERVE);
    }
  });

  it('never claims a 40-person raid fits', () => {
    for (const plan of ['guild', 'raid_guild'] as const) {
      expect(text(html(plan))).not.toMatch(/40-person|40 players/);
    }
  });
});
