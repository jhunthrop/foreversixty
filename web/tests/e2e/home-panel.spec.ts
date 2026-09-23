import { expect, test } from '@playwright/test';

const fulfil = (body: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify(body),
});

test('the home page offers Battle.net sign-in when signed out', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(fulfil({ ok: false, data: null, error: null, request_id: 'r' }, 401)),
  );
  await page.goto('/');
  await expect(page.getByTestId('home-signed-out')).toBeVisible();
  await expect(page.getByRole('link', { name: 'Sign in with Battle.net' })).toHaveAttribute(
    'href',
    /\/v1\/auth\/battlenet\/start\?next=%2Faccount%3Fsigned_in%3D1$/,
  );
  await expect(page.getByTestId('home-account-panel')).toHaveCount(0);
});

test('a signed-in visitor with zero characters still sees the Battle.net sign-in CTA, not a blank hero', async ({
  page,
}) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      fulfil({
        ok: true,
        data: {
          user: { id: 1, battletag: 'Fixture#1', email: null, role: 'user', anonymize: false },
          characters: [],
          guilds: [],
        },
        error: null,
        request_id: 'r',
      }),
    ),
  );
  await page.goto('/');
  await expect(page.getByTestId('home-account-panel')).toHaveCount(0);
  const signedOut = page.getByTestId('home-signed-out');
  await expect(signedOut).toBeVisible();
  await expect(signedOut).not.toHaveAttribute('inert');
  await expect(page.getByRole('link', { name: 'Sign in with Battle.net' })).toBeVisible();
});

test('the home page shows the current character strip when signed in', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      fulfil({
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
    ),
  );
  await page.addInitScript(() => {
    window.localStorage.setItem(
      'fs.currentCharacter',
      JSON.stringify({
        source: 'armory',
        ref: 'us/normal/kiloz',
        label: 'Kiloz · Warrior',
        classSlug: 'warrior',
        savedAt: new Date().toISOString(),
      }),
    );
  });
  await page.goto('/');
  await expect(page.getByTestId('home-account-panel')).toBeVisible();
  // The header is live on the home page too: a signed-in visitor sees their tag, not a
  // "Sign in" link beside their own character strip (Base.astro's `session`).
  await expect(page.getByTestId('session-nav')).toContainText('Fixture#1');
  await expect(page.getByTestId('session-nav-static')).toHaveCount(0);
  await expect(
    page.getByTestId('home-account-panel').getByRole('link', { name: 'Open in simulator' }),
  ).toHaveAttribute('href', '/sim');
  await expect(
    page.getByTestId('home-account-panel').getByRole('link', { name: 'Your characters' }),
  ).toHaveAttribute('href', '/account');
  // The signed-out "Sign in with Battle.net" link is only ever visually covered by the
  // grid-overlay CLS trick, never removed from the DOM -- without `inert`, a signed-in
  // keyboard/screen-reader user could still tab to, or hear, the duplicate link behind it.
  await expect(page.locator('#home-signed-out')).toHaveAttribute('inert', '');
  await expect(page.locator('#home-signed-out')).toHaveAttribute('aria-hidden', 'true');
});

test('the signed-in hero shows a rating figure once one exists, never before', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      fulfil({
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
    ),
  );
  await page.route('**/v1/characters/us/normal/kiloz/rating', (route) =>
    route.fulfill(
      fulfil({
        ok: true,
        data: {
          player_key: 'us/normal/kiloz',
          sample_size: 4,
          trend: [],
          best_component: 'damage',
          worst_component: 'utility',
          latest: { player_key: 'us/normal/kiloz', overall: 1.08 },
        },
        error: null,
        request_id: 'r',
      }),
    ),
  );
  await page.addInitScript(() => {
    window.localStorage.setItem(
      'fs.currentCharacter',
      JSON.stringify({
        source: 'armory',
        ref: 'us/normal/kiloz',
        label: 'Kiloz · Warrior',
        classSlug: 'warrior',
        savedAt: new Date().toISOString(),
      }),
    );
  });
  await page.goto('/');
  await expect(page.getByTestId('home-hero-rating')).toHaveText('Performance rating 1.08');
});

test('the hub lists the other characters as chips, and a chip makes that character current', async ({
  page,
}) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      fulfil({
        ok: true,
        data: {
          user: { id: 1, battletag: 'Fixture#1', email: null, role: 'user', anonymize: false },
          characters: [
            {
              key: 'us/normal/kiloz',
              region: 'us',
              ruleset: 'normal',
              name: 'Kiloz',
              class: 'warrior',
              level: 60,
            },
            {
              key: 'us/normal/dottzz',
              region: 'us',
              ruleset: 'normal',
              name: 'Dottzz',
              class: 'priest',
              level: 12,
            },
          ],
          guilds: [],
        },
        error: null,
        request_id: 'r',
      }),
    ),
  );
  await page.route('**/v1/characters/**', (route) =>
    route.fulfill(fulfil({ ok: false, data: null, error: { message: 'none' }, request_id: 'r' }, 404)),
  );
  await page.goto('/');
  const panel = page.getByTestId('home-account-panel');
  await expect(panel).toBeVisible();
  // Kiloz is the main character (highest level); Dottzz is the one chip.
  const chip = panel.getByTestId('home-character-chip');
  await expect(chip).toHaveCount(1);
  await expect(chip).toContainText('Dottzz');
  await chip.click();
  await expect(panel).toContainText('Dottzz');
  await expect(panel.getByTestId('home-character-chip')).toContainText('Kiloz');
  const stored = await page.evaluate(() => window.localStorage.getItem('fs.currentCharacter') ?? '');
  expect(stored).toContain('us/normal/dottzz');
});

test('a returning signed-in visitor sees the hub from the session snapshot before /v1/me answers', async ({
  page,
  context,
}) => {
  const body = {
    ok: true,
    data: {
      user: { id: 1, battletag: 'Fixture#1', email: null, role: 'user', anonymize: false },
      characters: [
        {
          key: 'us/normal/kiloz',
          region: 'us',
          ruleset: 'normal',
          name: 'Kiloz',
          class: 'warrior',
          level: 60,
        },
      ],
      guilds: [],
    },
    error: null,
    request_id: 'r',
  };
  // The readable half of the session cookie pair is what makes the snapshot trusted.
  await context.addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await page.route('**/v1/characters/**', (route) =>
    route.fulfill({ status: 404, json: { ok: false, data: null, error: { message: 'none' } } }),
  );
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(body)));
  await page.goto('/');
  await expect(page.getByTestId('home-account-panel')).toBeVisible();
  const stored = await page.evaluate(() => window.localStorage.getItem('fs.me') ?? '');
  expect(stored).toContain('Kiloz');

  // Second load: /v1/me is held for five seconds, yet the hub renders at once from the
  // snapshot, and the signed-out block never shows (Base.astro's session hint).
  await page.unroute('**/v1/me');
  await page.route('**/v1/me', async (route) => {
    await new Promise((resolve) => setTimeout(resolve, 5000));
    await route.fulfill(fulfil(body));
  });
  await page.goto('/');
  await expect(page.getByTestId('home-account-panel')).toBeVisible({ timeout: 3000 });
  await expect(page.getByTestId('home-signed-out')).toBeHidden();
});
