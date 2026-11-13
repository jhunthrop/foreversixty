// web/tests/e2e/admin-activity.spec.ts
// /admin/activity is the owner's view of who comes back: this week's figures, the weeks, retention
// by signup week and the days, every figure from before the live record shown as a floor. Anyone
// else sees one line, never the figures.
import type { Page } from '@playwright/test';
import { adminActivityCopy as copy } from '../../src/lib/engagement/admin';
import type { EngagementReport } from '../../src/lib/engagement/api';
import { SAMPLE_REPORT as REPORT } from './support/engagement-report';
import { expect, test } from './support/test';

const envelope = (status: number, data: unknown) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify({
    ok: status < 400,
    data: status < 400 ? data : null,
    error: status < 400 ? null : { code: 'x', message: 'refused', fields: null },
    request_id: 'r',
  }),
});

const me = (role: string) => ({
  user: { id: 78, battletag: 'Owner#1894', email: null, role, anonymize: false },
  characters: [],
  guilds: [],
});

async function as(page: Page, role: string | null, report: EngagementReport = REPORT): Promise<void> {
  await page.route('**/v1/me', (route) =>
    route.fulfill(role === null ? envelope(401, null) : envelope(200, me(role))),
  );
  await page.route('**/v1/admin/engagement', (route) => route.fulfill(envelope(200, report)));
}

test('the owner reads this week, the weeks, retention by signup week and the days', async ({ page }) => {
  await as(page, 'admin');
  await page.goto('/admin/activity');
  await expect(page.getByRole('heading', { level: 1, name: copy.pageTitle })).toBeVisible();
  const figures = page.getByTestId('activity-figures');
  await expect(figures).toContainText('30');
  await expect(figures).toContainText(copy.figures.today);
  await expect(figures).toContainText(copy.figures.lastWeek);
  await expect(page.getByTestId('activity-internal')).toHaveText(copy.internal(6));
  await expect(page.getByTestId('activity-floor')).toContainText('Recorded live since 12 Nov');
  await expect(page.getByTestId('activity-weeks-grid-row')).toHaveCount(12);
  await expect(page.getByTestId('activity-days-grid-row')).toHaveCount(28);
  const cohorts = page.getByTestId('activity-cohorts-grid-row');
  await expect(cohorts).toHaveCount(8);
  await expect(cohorts.first()).toContainText('9 Nov');
  await expect(cohorts.last()).toContainText('21 Sep');
  await expect(cohorts.last()).toContainText('10 joined');
  await expect(page.getByTestId('activity-days-grid-row').nth(2)).toContainText('≥ 26');
});

test('anyone else sees one line and never the figures', async ({ page }) => {
  await as(page, 'user');
  await page.goto('/admin/activity');
  await expect(page.getByTestId('activity-not-owner')).toHaveText(copy.notOwner);
  await expect(page.getByTestId('activity-figures')).toHaveCount(0);
});

test('a signed-out visitor is asked to sign in', async ({ page }) => {
  await as(page, null);
  await page.goto('/admin/activity');
  await expect(page.getByTestId('activity-signed-out')).toBeVisible();
});

test('a failed read offers a retry', async ({ page }) => {
  await page.route('**/v1/me', (route) => route.fulfill(envelope(200, me('admin'))));
  await page.route('**/v1/admin/engagement', (route) => route.fulfill(envelope(500, null)));
  await page.goto('/admin/activity');
  await expect(page.getByTestId('activity-load-error')).toBeVisible();
});
