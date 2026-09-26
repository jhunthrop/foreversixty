// web/tests/e2e/handoffs-home.spec.ts
import { expect, test } from '@playwright/test';

test('the five next-steps cards carry Planner, Simulator, Logs, Rankings and Guides, in spec order', async ({
  page,
}) => {
  await page.goto('/');
  const grid = page.getByTestId('home-next-steps-signed-out');
  const labels = await grid.locator('[data-testid^="home-next-"] .label').allTextContents();
  expect(labels).toEqual(['Planner', 'Simulator', 'Logs', 'Rankings', 'Guides']);
  await expect(
    page.getByTestId('home-next-simulator-signed-out').getByRole('link', { name: 'Open the simulator' }),
  ).toHaveAttribute('href', '/sim');
  // Review round 1 fix item 2: Guides is a fifth door here, not just a header nav link.
  await expect(
    page.getByTestId('home-next-guides-signed-out').getByRole('link', { name: 'Open guides' }),
  ).toHaveAttribute('href', '/guides');
});

test('the "Get set up" card links to /setup', async ({ page }) => {
  await page.goto('/');
  const row = page.getByTestId('home-companion-row');
  await expect(row.getByRole('link', { name: /Get set up/ })).toHaveAttribute('href', '/setup');
});

test('a signed-in visitor sees the "Get set up" banner collapse to one line, never the three-step pitch', async ({
  page,
}) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({
        ok: true,
        data: {
          user: { id: 1, battletag: 'Fixture#1', email: null, role: 'user', anonymize: false },
          characters: [
            { key: 'us/normal/kiloz', region: 'us', ruleset: 'normal', name: 'Kiloz', class: 'Warrior' },
          ],
          guilds: [],
        },
        error: null,
        request_id: 'r',
      }),
    }),
  );
  await page.goto('/');
  const row = page.getByTestId('home-companion-row');
  await expect(row.getByRole('link', { name: /Get set up/ })).toHaveCount(0);
  const line = page.getByTestId('home-get-set-up');
  await expect(line).toHaveText('Install the addon · Set up the companion →');
  await expect(line).toHaveAttribute('href', '/setup');
});

test('a signed-in visitor with an addon-linked build sees the collapsed one-line sentence', async ({
  page,
}) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({
        ok: true,
        data: {
          user: { id: 1, battletag: 'Fixture#1', email: null, role: 'user', anonymize: false },
          characters: [
            {
              key: 'us/normal/kiloz',
              region: 'us',
              ruleset: 'normal',
              name: 'Kiloz',
              class: 'Warrior',
              build: { source: 'addon', captured_at: '2026-09-20T00:00:00Z' },
            },
          ],
          guilds: [],
        },
        error: null,
        request_id: 'r',
      }),
    }),
  );
  await page.goto('/');
  await expect(page.getByTestId('home-get-set-up')).toHaveText(
    'Signed in and addon linked · Set up the companion →',
  );
});
