// web/tests/captures/admin-activity/capture.spec.ts
// Review screenshots of /admin/activity, the owner's actives and retention, at 1440, 2000 and 390,
// and the line anyone else sees. Not a test of behaviour (tests/e2e/admin-activity.spec.ts is); it
// writes PNGs to design/reviews/2026-11-13-admin-activity/.
import path from 'node:path';
import { test, type Page } from '@playwright/test';
import { SAMPLE_REPORT } from '../../e2e/support/engagement-report';

const OUT = path.resolve(import.meta.dirname, '../../../../design/reviews/2026-11-13-admin-activity');

const envelope = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
});

const me = (role: string) => ({
  user: { id: 78, battletag: 'Owner#1894', email: null, role, anonymize: false },
  characters: [],
  guilds: [],
});

async function open(page: Page, role: string): Promise<void> {
  await page.route('**/v1/me', (route) => route.fulfill(envelope(me(role))));
  await page.route('**/v1/admin/engagement', (route) => route.fulfill(envelope(SAMPLE_REPORT)));
  await page.goto('/admin/activity');
  await page.waitForSelector(
    role === 'admin' ? '[data-testid="activity-days-grid"]' : '[data-testid="activity-not-owner"]',
  );
  await page.waitForTimeout(600);
}

for (const width of [1440, 2000]) {
  test(`owner at ${width}`, async ({ page }) => {
    await page.setViewportSize({ width, height: 1000 });
    await open(page, 'admin');
    await page.screenshot({ path: path.join(OUT, `admin-activity-${width}.png`), fullPage: true });
  });
}

test('not the owner at 1440', async ({ page }) => {
  await page.setViewportSize({ width: 1440, height: 900 });
  await open(page, 'user');
  await page.screenshot({ path: path.join(OUT, 'admin-activity-not-owner-1440.png'), fullPage: true });
});

test.describe('phone', () => {
  test.use({ viewport: { width: 390, height: 844 }, isMobile: true, hasTouch: true, deviceScaleFactor: 2 });
  test('owner at 390', async ({ page }) => {
    await open(page, 'admin');
    await page.screenshot({ path: path.join(OUT, 'admin-activity-390.png'), fullPage: true });
  });
});
