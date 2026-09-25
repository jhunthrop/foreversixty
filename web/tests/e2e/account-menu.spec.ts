// web/tests/e2e/account-menu.spec.ts
// The header's account chip and menu (spec 2026-09-25, section 3.3): open, switch character,
// sign out, Escape closes. Section 9 names this as lane A's required evidence for the menu.
import { expect, test } from '@playwright/test';

const ME_BODY = {
  ok: true,
  data: {
    user: { id: 'u1', battletag: 'Simfury#1234' },
    characters: [
      { key: 'us-era-simfury', region: 'us', ruleset: 'era', name: 'Simfury', class: 'Warrior', level: 60 },
      { key: 'us-era-simheal', region: 'us', ruleset: 'era', name: 'Simheal', class: 'Priest', level: 60 },
    ],
    guilds: [],
    entitlements: {},
  },
  error: null,
  request_id: 'r',
};

test.beforeEach(async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(ME_BODY) }),
  );
});

test('opens on click, lists the account items, and closes on Escape', async ({ page }) => {
  await page.goto('/');
  const menu = page.getByTestId('account-menu');
  await menu.locator('summary').click();
  await expect(page.getByRole('link', { name: 'Your account' })).toBeVisible();
  await expect(page.getByRole('link', { name: 'Devices' })).toBeAttached();
  // Scoped to the account menu: the home page's own hero also has a "Plan talents" link and
  // the primary nav's "Planner" link, both of which contain "Plan" as a substring, so an
  // unscoped role query is ambiguous here (Playwright's `name` match is substring by default).
  await expect(menu.getByRole('link', { name: 'Plan' })).toBeAttached();
  await expect(page.getByTestId('account-menu-signout')).toBeVisible();
  await page.keyboard.press('Escape');
  await expect(page.getByRole('link', { name: 'Your account' })).toBeHidden();
});

test('closes on outside click', async ({ page }) => {
  await page.goto('/');
  await page.getByTestId('account-menu').locator('summary').click();
  await expect(page.getByRole('link', { name: 'Your account' })).toBeVisible();
  await page.mouse.click(5, 5);
  await expect(page.getByRole('link', { name: 'Your account' })).toBeHidden();
});

test('switching a character writes the current-character pointer and closes the menu', async ({ page }) => {
  await page.goto('/');
  await page.getByTestId('account-menu').locator('summary').click();
  await page.getByTestId('account-menu-character-us-era-simheal').click();
  await expect(page.getByRole('link', { name: 'Your account' })).toBeHidden();
  const stored = await page.evaluate(() => localStorage.getItem('fs.currentCharacter'));
  expect(stored).not.toBeNull();
  expect(JSON.parse(stored ?? '{}').ref).toBe('us-era-simheal');
});
