// web/tests/e2e/logs-companion.spec.ts
// Logs landing spec (2026-10-04) §4.D.1 (the per-device status line) and §4.D.2 (the
// pairing poll's success state).
import { expect, test, type Page } from '@playwright/test';
import { meAddonFixture } from '../../src/fixtures/me-addon';
import { routeSampleReportMeta } from './support/logs-hero-fixture';

function fulfil(body: unknown, status = 200): { status: number; contentType: string; body: string } {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: status < 400, data: body, error: null, request_id: 'e2e' }),
  };
}

async function signIn(page: Page): Promise<void> {
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(meAddonFixture)));
  await page.route('**/v1/reports?mine=1**', (route) =>
    route.fulfill(fulfil({ rows: [], total: 0, page: 1, per_page: 100 })),
  );
  await page.route('**/v1/reports/recent**', (route) => route.fulfill(fulfil({ rows: [] })));
  await routeSampleReportMeta(page);
}

test('a device that has reported in shows a relative last-seen line, never "Connected"', async ({ page }) => {
  await signIn(page);
  await page.route('**/v1/devices', (route) =>
    route.fulfill(
      fulfil([
        {
          id: 'dev1',
          name: 'MacBook Pro',
          platform: 'macOS',
          created_at: '2026-10-01T00:00:00Z',
          last_seen_at: new Date(Date.now() - 4 * 60_000).toISOString(),
        },
      ]),
    ),
  );

  await page.goto('/logs');
  const status = page.getByTestId('companion-status-line');
  await expect(status).toContainText('MacBook Pro');
  await expect(status).toContainText('macOS');
  await expect(status).toContainText('last seen');
  await expect(status).not.toContainText('Connected');
});

test('a device paired but never seen reads honestly, no fabricated status', async ({ page }) => {
  await signIn(page);
  await page.route('**/v1/devices', (route) =>
    route.fulfill(
      fulfil([
        {
          id: 'dev1',
          name: 'Gaming PC',
          platform: 'Windows',
          created_at: '2026-10-01T00:00:00Z',
          last_seen_at: null,
        },
      ]),
    ),
  );

  await page.goto('/logs');
  await expect(page.getByTestId('companion-status-line')).toHaveText(
    'Gaming PC · Windows · paired, not seen yet',
  );
});

test('no status line at all when nothing is paired yet', async ({ page }) => {
  await signIn(page);
  await page.route('**/v1/devices', (route) => route.fulfill(fulfil([])));

  await page.goto('/logs');
  await expect(page.getByTestId('companion-status')).toHaveCount(0);
});

test('pairing: code shown, polls while waiting, then shows the success line once the companion uses it', async ({
  page,
}) => {
  await signIn(page);
  // Elapsed wall-clock time, not a call count: the initial page load and the poll both
  // read this same route (the shared query cache means only one of them is a real
  // network hit at page load, and exactly which one is not worth pinning a test to),
  // so "paired after N seconds" is the deterministic signal instead.
  const start = Date.now();
  const PAIR_AFTER_MS = 9_000;
  await page.route('**/v1/devices', (route) => {
    const paired = Date.now() - start > PAIR_AFTER_MS;
    return route.fulfill(
      fulfil(
        paired
          ? [
              {
                id: 'dev-new',
                name: "Zulmara's PC",
                platform: 'Windows',
                created_at: new Date().toISOString(),
                last_seen_at: null,
              },
            ]
          : [],
      ),
    );
  });
  await page.route('**/v1/devices/pair', (route) =>
    route.fulfill(fulfil({ code: '4821-9930', expires_in: 600 })),
  );

  await page.goto('/logs');
  await page.getByRole('button', { name: 'Show pairing code' }).click();
  await expect(page.getByTestId('pairing-code')).toHaveText('4821-9930');

  // Still polling, code still on screen, no success yet -- well before PAIR_AFTER_MS.
  await page.waitForTimeout(6_000);
  await expect(page.getByTestId('pairing-code')).toBeVisible();
  await expect(page.getByTestId('pairing-success')).toHaveCount(0);

  // The next poll tick after PAIR_AFTER_MS finds the new device.
  await expect(page.getByTestId('pairing-success')).toHaveText(
    "Zulmara's PC paired. It starts uploading as soon as you are logging.",
    { timeout: 10_000 },
  );
  await expect(page.getByTestId('pairing-code')).toHaveCount(0);
});
