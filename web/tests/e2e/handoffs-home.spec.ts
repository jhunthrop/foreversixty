// web/tests/e2e/handoffs-home.spec.ts
import { expect, test } from './support/test';

test('the landing carries three outcome rows in the player’s words, no feature cards', async ({ page }) => {
  await page.goto('/');
  const rows = page.getByTestId('home-outcomes').locator('a');
  await expect(rows).toHaveText([/Plan your talents/, /Sim your gear/, /Upload tonight’s raid/]);
  await expect(page.getByTestId('home-outcome-planner')).toHaveAttribute('href', '/planner');
  await expect(page.getByTestId('home-outcome-sim')).toHaveAttribute('href', '/sim');
  await expect(page.getByTestId('home-outcome-logs')).toHaveAttribute('href', '/logs');
  await expect(page.getByText('What the site does')).toHaveCount(0);
  await expect(page.getByTestId('home-five-doors')).toHaveCount(0);
});

test('with a current character the planner and sim rows open that character', async ({ page }) => {
  await page.addInitScript(() => {
    localStorage.setItem(
      'fs.currentCharacter',
      JSON.stringify({
        source: 'code',
        ref: 'FS1CODE',
        label: 'Bow Jackzon · Marksmanship Hunter',
        classSlug: 'hunter',
        savedAt: '2026-10-20T10:00:00Z',
      }),
    );
  });
  await page.goto('/');
  await expect(page.getByTestId('home-outcome-planner')).toHaveAttribute('href', '/planner?code=FS1CODE');
  await expect(page.getByTestId('home-outcome-sim')).toHaveAttribute('href', /^\/sim.*FS1CODE/);
  await expect(page.getByTestId('home-outcome-logs')).toHaveAttribute('href', '/logs');
});

test('the "Get set up" card links to /setup', async ({ page }) => {
  await page.goto('/');
  const row = page.getByTestId('home-companion-row');
  await expect(row.getByRole('link', { name: /Get set up/ })).toHaveAttribute('href', '/setup');
  await expect(page.getByTestId('home-get-set-up-card').getByTestId('tool-card-button')).toHaveText(
    'Open setup',
  );
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
  // Review round 3 ("islands" item): this card is below the signed-out fold and now
  // mounts `client:visible`, like every other below-fold island on this page, so it only
  // hydrates once scrolled into view. `HomeGetSetUp` itself renders nothing before it is
  // ready, so its own `<astro-island>` wrapper has zero height pre-hydration --
  // `scrollIntoViewIfNeeded` aligns the (tall) static pitch card's own top edge, which can
  // leave that zero-height sibling just past the bottom of a short viewport. Centering the
  // row instead keeps both safely inside the intersection observer's bounds.
  await row.evaluate((el) => el.scrollIntoView({ block: 'center' }));
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
  await page.getByTestId('home-companion-row').evaluate((el) => el.scrollIntoView({ block: 'center' }));
  await expect(page.getByTestId('home-get-set-up')).toHaveText(
    'Signed in and addon linked · synced 10 minutes ago → Set up the companion',
  );
});
