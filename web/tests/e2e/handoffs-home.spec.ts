// web/tests/e2e/handoffs-home.spec.ts
import { expect, test } from '@playwright/test';

test('the five doors carry Planner, Simulator, Logs, Rankings and Guides, in spec order', async ({
  page,
}) => {
  await page.goto('/');
  const order = ['planner', 'simulator', 'logs', 'rankings', 'guides'];
  const cardHandles = await page.getByTestId('home-five-doors').locator('a').all();
  const testids = await Promise.all(cardHandles.map((card) => card.getAttribute('data-testid')));
  expect(testids).toEqual(order.map((slug) => `home-five-doors-${slug}`));
  await expect(page.getByTestId('home-five-doors-simulator')).toHaveAttribute('href', '/sim');
  await expect(page.getByTestId('home-five-doors-simulator')).toContainText('Open the simulator');
  // Review round 1 fix item 2: Guides is a fifth door here, not just a header nav link.
  await expect(page.getByTestId('home-five-doors-guides')).toHaveAttribute('href', '/guides');
  await expect(page.getByTestId('home-five-doors-guides')).toContainText('Open guides');
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

test('a signed-in visitor with an addon-linked build sees the collapsed sentence, naming the real sync time', async ({
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
              // Home rebuild spec §3.B.6: this line and the hero's own sync line
              // (home-panel.spec.ts) must never disagree about the addon's last sync,
              // both reading the same `build.captured_at`.
              build: { source: 'addon', captured_at: new Date(Date.now() - 10 * 60_000).toISOString() },
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
    'Signed in and addon linked · synced 10 minutes ago → Set up the companion',
  );
});
