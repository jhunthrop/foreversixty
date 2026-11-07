import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { text } from '../../../test-support/raid-fixtures';
import GuildPlanChooser from './GuildPlanChooser.svelte';

const OPEN = new Date('2026-11-15T12:00:00Z');
const CLOSED = new Date('2027-02-01T12:00:00Z');

function html(
  props: Partial<{ count: number; trialUsed: boolean; lapsed: boolean; now: Date }> = {},
): string {
  return render(GuildPlanChooser, {
    props: {
      guildId: 7,
      verifiedCount: props.count ?? 12,
      trialUsed: props.trialUsed ?? false,
      lapsed: props.lapsed ?? false,
      now: props.now ?? OPEN,
    },
  }).body;
}

const checked = (body: string, tier: string): boolean =>
  new RegExp(`<input[^>]*value="${tier}"[^>]*checked`).test(body) ||
  new RegExp(`<input[^>]*checked[^>]*value="${tier}"`).test(body);

describe('GuildPlanChooser', () => {
  it('prices each plan Row with a Figure and shows the founder list price as plain text', () => {
    const body = html({ now: OPEN });
    const price = /data-testid="guild-billing-option-guild-price">([\s\S]*?)<\/label>/.exec(body)?.[1] ?? '';
    expect(text(price)).toContain('$120');
    expect(text(price)).toContain('a year');
    expect(price).toContain('tabular-nums');
    expect(body).not.toContain('<s ');
    expect(body).toMatch(/guild-billing-option-guild-list">\$150 after 6 Jan 2027</);
    expect(body).toContain('max-w-[720px]');
  });

  it('draws the interval as the Segmented control and the plans as Rows, the picked one selected', () => {
    const body = html({ count: 12 });
    expect(body).toMatch(
      /role="radio"[^>]*aria-checked="true"[^>]*data-testid="guild-billing-interval-yearly"/,
    );
    expect(body).toMatch(/ui-row[^"]*bg-selected[^>]*data-testid="guild-billing-option-guild"/);
    expect(body).not.toMatch(/ui-row[^"]*bg-selected[^>]*data-testid="guild-billing-option-raid_guild"/);
  });

  it('preselects Guild at 12 members and Raid guild at 28', () => {
    expect(checked(html({ count: 12 }), 'guild')).toBe(true);
    expect(checked(html({ count: 28 }), 'raid_guild')).toBe(true);
    expect(checked(html({ count: 28 }), 'guild')).toBe(false);
  });

  it('offers the trial, its line and the founder line while the window is open', () => {
    const body = html();
    expect(text(body)).toContain('Start the 7-day trial');
    expect(body).toContain('guild-billing-trial-line');
    expect(body).toContain('guild-billing-founder-line');
    expect(body).toContain('href="/premium/checkout?plan=guild&amp;interval=yearly&amp;guild_id=7"');
  });

  it('with the trial used, the button reads Subscribe with the price and there is no trial line', () => {
    const body = html({ trialUsed: true });
    expect(text(body)).toContain('Subscribe · $120 a year');
    expect(body).not.toContain('guild-billing-trial-line');
    expect(text(body)).not.toContain('Start the 7-day trial');
  });

  it('after the founder window it shows list prices and no founder line', () => {
    const body = html({ trialUsed: true, now: CLOSED });
    expect(text(body)).toContain('Subscribe · $150 a year');
    expect(body).not.toContain('guild-billing-founder-line');
  });

  it('a lapsed guild reads Subscribe again with no trial line', () => {
    const body = html({ lapsed: true });
    expect(text(body)).toContain('Subscribe again');
    expect(body).not.toContain('guild-billing-trial-line');
  });
});

describe('GuildPlanChooser compare link (spec 2026-10-29-premium-entry section 7)', () => {
  it('one action link under the cards, above the subscribe button, to the guild region', () => {
    const body = html();
    const link = /<a[^>]*data-testid="guild-billing-compare"[^>]*>[\s\S]*?<\/a>/.exec(body)?.[0] ?? '';
    expect(link).toContain('href="/premium#guild-leaders"');
    expect(text(link)).toBe('Compare Guild and Raid guild');
    expect(link).toContain('data-glyph="chevron"');
    const at = (id: string): number => body.indexOf(`data-testid="${id}"`);
    expect(at('guild-billing-compare')).toBeGreaterThan(at('guild-billing-option-raid_guild'));
    expect(at('guild-billing-compare')).toBeLessThan(at('guild-billing-subscribe'));
  });
});
