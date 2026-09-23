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
  await expect(
    page.getByTestId('home-account-panel').getByRole('link', { name: 'Open in simulator' }),
  ).toHaveAttribute('href', '/sim');
  await expect(
    page.getByTestId('home-account-panel').getByRole('link', { name: 'Your characters' }),
  ).toHaveAttribute('href', '/account');
});
