// web/src/components/checkout/CheckoutBand.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { checkoutGuildCopy as copy } from '../../lib/billing/checkout-guild-copy';
import { CHECKOUT_BAND_RESERVE } from '../../lib/billing/checkout-layout';
import type { Interval, PlanKey } from '../../lib/billing/plans';
import CheckoutBand from './CheckoutBand.svelte';

function html(plan: PlanKey | null, interval: Interval = 'yearly', founder = true): string {
  return render(CheckoutBand, { props: { plan, interval, founder } }).body.replace(/<!--.*?-->/g, '');
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

  it('the price is tabular JetBrains Mono', () => {
    expect(html('guild')).toMatch(/class="[^"]*font-mono[^"]*tabular-nums[^"]*"[^>]*>\$120</);
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

  it('before the URL is read: the band ground and its reserve, no h1', () => {
    const body = html(null);
    expect(body).toContain('full-bleed');
    expect(body).toContain(CHECKOUT_BAND_RESERVE);
    expect(body).not.toContain('<h1');
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
