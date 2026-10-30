// web/tests/e2e/guild-landing.spec.ts
// The bare /guild landing: the header's Guild door lands here. A lookup by name, the top
// guilds, and a redirect to the member's own guild when the session has a verified one.
import { expect, test, type Page } from '@playwright/test';
import { choose, nativeSelect } from './support/select';

function envelope(data: unknown) {
  return {
    status: 200,
    contentType: 'application/json',
    body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
  };
}

async function stubRankings(page: Page): Promise<void> {
  await page.route('**/v1/rankings/guilds*', (route) =>
    route.fulfill(
      envelope({
        rows: [
          {
            rank: 1,
            guild: { name: 'The Last Watch', ruleset: 'hardcore', region: 'us' },
            value: 9,
            fought_at: '2026-12-09T22:10:00Z',
            report_id: 'fixture2abcd',
          },
        ],
      }),
    ),
  );
}

async function stubSignedOut(page: Page): Promise<void> {
  await page.route('**/v1/me', (route) =>
    route.fulfill({
      status: 401,
      contentType: 'application/json',
      body: '{"ok":false,"data":null,"error":null,"request_id":"r"}',
    }),
  );
}

test.describe('guild landing', () => {
  test('shows the lookup and the top guilds list when signed out', async ({ page }) => {
    await stubRankings(page);
    await stubSignedOut(page);
    await page.goto('/guild');
    await expect(page.getByRole('heading', { level: 1, name: 'Guilds' })).toBeVisible();
    await expect(page.getByTestId('guild-lookup')).toBeVisible();
    await expect(page.getByRole('heading', { level: 2, name: 'Top guilds' })).toBeVisible();
    await expect(page.getByTestId('home-top-guilds')).toContainText('The Last Watch');
    await expect(page.getByText('Run a guild? Open its page and choose Claim this guild.')).toBeVisible();
  });

  test('an empty name explains itself and does not navigate', async ({ page }) => {
    await stubRankings(page);
    await stubSignedOut(page);
    await page.goto('/guild');
    await page.getByRole('button', { name: 'Open guild' }).click();
    await expect(page.getByRole('status').filter({ hasText: "Type the guild's name" })).toBeVisible();
    expect(new URL(page.url()).pathname).toBe('/guild');
  });

  test('submitting a name opens that guild page', async ({ page }) => {
    await stubRankings(page);
    await stubSignedOut(page);
    await page.goto('/guild');
    await page.getByLabel('Guild name').fill('  The Last Watch ');
    await choose(nativeSelect(page, 'Region'), 'eu');
    await choose(nativeSelect(page, 'Ruleset'), 'normal');
    await page.getByRole('button', { name: 'Open guild' }).click();
    await expect(page).toHaveURL(/\/guild\/eu\/normal\/the-last-watch$/);
  });

  test('a signed-in member with a verified guild lands on that guild', async ({ page }) => {
    await stubRankings(page);
    await page.route('**/v1/me', (route) =>
      route.fulfill(
        envelope({
          user: {
            id: 7,
            battletag: 'Fixture#1234',
            email: null,
            role: 'user',
            anonymize: false,
            premium: false,
          },
          characters: [],
          guilds: [
            { id: 3, region: 'us', ruleset: 'pvp', name: 'Olympus XXVII', rank: 'member', verified: true },
          ],
        }),
      ),
    );
    await page.goto('/guild');
    await expect(page).toHaveURL(/\/guild\/us\/pvp\/olympus-xxvii$/);
  });

  test('controls keep the 44px floor and nothing scrolls sideways at 390', async ({ page }) => {
    await stubRankings(page);
    await stubSignedOut(page);
    await page.setViewportSize({ width: 390, height: 844 });
    await page.goto('/guild');
    await expect(page.getByTestId('guild-lookup')).toBeVisible();
    const controls = page.getByTestId('guild-lookup').locator('input, select:not([hidden]), button');
    for (const control of await controls.all()) {
      const box = await control.boundingBox();
      expect(box?.height ?? 0).toBeGreaterThanOrEqual(44);
    }
    const overflow = await page.evaluate(
      () => document.documentElement.scrollWidth - document.documentElement.clientWidth,
    );
    expect(overflow).toBeLessThanOrEqual(0);
  });
});
