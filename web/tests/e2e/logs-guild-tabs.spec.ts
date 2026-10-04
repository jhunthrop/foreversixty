// web/tests/e2e/logs-guild-tabs.spec.ts
// Logs landing spec (2026-10-04) §4.C.1: the guild tab strip on "Your reports", backed by
// GET /v1/guilds/{id}/reports (the API lane's new endpoint, merged on main -- same
// MinePage/MyReportPage shape and page-number pagination as GET /v1/reports?mine=1).
import { expect, test } from '@playwright/test';
import { meAddonFixture } from '../../src/fixtures/me-addon';
import { routeSampleReportMeta } from './support/logs-hero-fixture';

function fulfil(body: unknown, status = 200): { status: number; contentType: string; body: string } {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: status < 400, data: body, error: null, request_id: 'e2e' }),
  };
}

test('a guild tab strip shows when the visitor has a guild, and switches the list', async ({ page }) => {
  const me = {
    ...meAddonFixture,
    guilds: [{ id: 5, region: 'us', ruleset: 'normal', name: 'The Last Watch', verified: true }],
  };
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(me)));
  await page.route('**/v1/reports?mine=1**', (route) =>
    route.fulfill(
      fulfil({
        rows: [
          {
            id: 'mine001',
            title: 'Solo farm',
            zone: 'Deadmines',
            status: 'complete',
            visibility: 'private',
            created_at: '2026-10-02T20:00:00Z',
            fight_count: 5,
            kill_count: 2,
          },
        ],
        total: 1,
        page: 1,
        per_page: 100,
      }),
    ),
  );
  await page.route('**/v1/guilds/5/reports**', (route) =>
    route.fulfill(
      fulfil({
        rows: [
          {
            id: 'guildrep01',
            title: 'Raid night, week 4',
            zone: 'Molten Core',
            status: 'complete',
            visibility: 'guild',
            created_at: '2026-10-03T22:00:00Z',
            fight_count: 9,
            kill_count: 5,
          },
        ],
        total: 1,
        page: 1,
        per_page: 100,
      }),
    ),
  );
  await page.route('**/v1/reports/recent**', (route) => route.fulfill(fulfil({ rows: [] })));
  await routeSampleReportMeta(page);

  await page.goto('/logs');
  const tabs = page.getByTestId('my-reports-guild-tabs');
  await expect(tabs).toBeVisible();
  const mine = tabs.getByRole('tab', { name: 'Mine' });
  const guild = tabs.getByTestId('my-reports-guild-tab');
  await expect(mine).toHaveAttribute('aria-selected', 'true');
  await expect(page.getByRole('link', { name: 'Solo farm' })).toBeVisible();

  await guild.click();
  await expect(guild).toHaveAttribute('aria-selected', 'true');
  await expect(mine).toHaveAttribute('aria-selected', 'false');
  await expect(page.getByRole('link', { name: 'Raid night, week 4' })).toBeVisible();
  await expect(page.getByRole('link', { name: 'Solo farm' })).toHaveCount(0);
});

test("a guild tab with no reports yet shows the guild's own empty copy", async ({ page }) => {
  const me = {
    ...meAddonFixture,
    guilds: [{ id: 5, region: 'us', ruleset: 'normal', name: 'The Last Watch', verified: true }],
  };
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(me)));
  await page.route('**/v1/reports?mine=1**', (route) =>
    route.fulfill(fulfil({ rows: [], total: 0, page: 1, per_page: 100 })),
  );
  await page.route('**/v1/guilds/5/reports**', (route) =>
    route.fulfill(fulfil({ rows: [], total: 0, page: 1, per_page: 100 })),
  );
  await page.route('**/v1/reports/recent**', (route) => route.fulfill(fulfil({ rows: [] })));
  await routeSampleReportMeta(page);

  await page.goto('/logs');
  await page.getByTestId('my-reports-guild-tab').click();
  await expect(page.getByTestId('my-reports-empty')).toHaveText('No The Last Watch reports yet.');
});
