import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { UPGRADE_PREVIEW } from '../../../test-support/billing-fixtures';
import { text } from '../../../test-support/raid-fixtures';
import GuildUpgradeConfirm from './GuildUpgradeConfirm.svelte';

function html(preview: typeof UPGRADE_PREVIEW | null, error = '', committing = false): string {
  return render(GuildUpgradeConfirm, {
    props: { preview, committing, error, onCommit: () => {}, onKeep: () => {} },
  }).body;
}

const commitDisabled = (body: string): boolean =>
  /<button[^>]*disabled[^>]*data-testid="guild-upgrade-commit"/.test(body) ||
  /<button[^>]*data-testid="guild-upgrade-commit"[^>]*disabled/.test(body);

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
