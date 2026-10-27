import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { UPGRADE_PREVIEW } from '../../../test-support/billing-fixtures';
import { text } from '../../../test-support/raid-fixtures';
import GuildUpgradeConfirm from './GuildUpgradeConfirm.svelte';

type Plan = { interval: 'monthly' | 'yearly' | null; founder: boolean; current_period_end: string | null };

function html(
  preview: typeof UPGRADE_PREVIEW | null,
  error = '',
  committing = false,
  plan: Plan | null = null,
): string {
  return render(GuildUpgradeConfirm, {
    props: { plan, preview, committing, error, onCommit: () => {}, onKeep: () => {} },
  }).body;
}

const FOUNDER_YEARLY: Plan = {
  interval: 'yearly',
  founder: true,
  current_period_end: '2026-12-01T12:00:00Z',
};

const commitDisabled = (body: string): boolean =>
  /<button[^>]*\sdisabled(=|\s|>)[^>]*data-testid="guild-upgrade-commit"/.test(body) ||
  /<button[^>]*data-testid="guild-upgrade-commit"[^>]*\sdisabled(=|\s|>)/.test(body);

describe('GuildUpgradeConfirm', () => {
  it('shows a Skeleton and holds the commit until the preview lands', () => {
    const body = html(null);
    expect(body).toContain('guild-upgrade-skeleton');
    expect(commitDisabled(body)).toBe(true);
    expect(text(body)).toContain('Keep Guild');
  });

  it('states the prorated amount and the next price once loaded', () => {
    const body = html({ ...UPGRADE_PREVIEW, renews_at: '2026-12-01T12:00:00Z' });
    expect(body).not.toContain('guild-upgrade-skeleton');
    expect(commitDisabled(body)).toBe(false);
    expect(text(body)).toContain(
      'Move to Raid guild? Today you pay $6 for the rest of this month. From 1 Dec 2026 it is $25 a month.',
    );
  });

  it('while loading, only the amount is a Skeleton: the period, date and founder price stand', () => {
    const body = html(null, '', false, FOUNDER_YEARLY);
    expect(body).toContain('guild-upgrade-skeleton');
    expect(text(body)).toContain('Move to Raid guild? Today you pay');
    expect(text(body)).toContain('for the rest of this year. From 1 Dec 2026 it is $200 a year.');
  });

  it('the date is the plan line own, so the two never disagree', () => {
    const body = html({ ...UPGRADE_PREVIEW, renews_at: '2026-12-02T12:00:00Z' }, '', false, {
      interval: 'monthly',
      founder: false,
      current_period_end: '2026-12-01T12:00:00Z',
    });
    expect(text(body)).toContain('From 1 Dec 2026 it is $25 a month.');
  });

  it('Keep Guild shows its gold ring on focus moved in code, and the buttons go full width below 768', () => {
    const body = html(UPGRADE_PREVIEW);
    expect(body).toMatch(/class="[^"]*focus:outline-gold[^"]*"[^>]*data-testid="guild-upgrade-keep"/);
    expect(body).toContain('flex flex-col items-stretch');
    expect(body).toMatch(/class="[^"]*w-full[^"]*md:w-fit[^"]*"[^>]*data-testid="guild-upgrade-commit"/);
  });

  it('shows the failure line', () => {
    const body = html(UPGRADE_PREVIEW, 'That did not go through. Nothing was charged. Try again.');
    expect(body).toContain('guild-upgrade-error');
  });

  it('reads Moving… and holds the commit while it is in flight', () => {
    const body = html(UPGRADE_PREVIEW, '', true);
    expect(text(body)).toContain('Moving…');
    expect(commitDisabled(body)).toBe(true);
  });
});
