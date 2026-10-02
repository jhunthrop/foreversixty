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

test('sign out calls the sessions endpoint and returns to the home page', async ({ page }) => {
  let signOutCalled = false;
  await page.route('**/v1/sessions', async (route) => {
    signOutCalled = signOutCalled || route.request().method() === 'DELETE';
    await route.fulfill({ status: 204, contentType: 'application/json', body: '' });
  });
  // /account also reads the devices list; mocked so the page settles without a real
  // network call, the same fixture shape account-billing.spec.ts uses.
  await page.route('**/v1/devices', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({ ok: true, data: [], error: null, request_id: 'r' }),
    }),
  );

  // Starts away from '/' so the post-sign-out redirect is an observable navigation, not a
  // no-op landing on the page it was already on.
  await page.goto('/account');
  await page.getByTestId('account-menu').locator('summary').click();
  await page.getByTestId('account-menu-signout').click();

  await page.waitForURL('/');
  expect(new URL(page.url()).pathname).toBe('/');
  expect(signOutCalled).toBe(true);
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

// Owner-reported defect, 2026-10-01: "the signed-in status refreshes on every page load" --
// AccountMenu.svelte only replaces the signed-out shell once it hydrates
// (client:idle-after-load) and its own /v1/me read answers, so a returning signed-in
// visitor saw "Sign in" flash on every single page for most of a second. Base.astro's
// pre-paint script now fills the `account-menu-preflight` shell from the persisted /v1/me
// cache entry before first paint. These two tests read the DOM with page.evaluate()
// immediately after goto() resolves (no Playwright auto-retrying `expect`, which could
// itself wait out the island's own hydration delay and hide the bug this is pinning) --
// the real `account-menu` chip is confirmed absent at that instant, proving the assertion
// really observes the pre-hydration window.
test('a returning signed-in visitor sees the real chip before the next page even hydrates', async ({
  page,
  context,
}) => {
  // The readable half of the session cookie pair is what makes a persisted private cache
  // entry trusted (lib/data/query.ts's `sessionHinted`) -- a real sign-in sets it; this
  // test's mocked /v1/me never goes through that flow, so it sets the cookie directly,
  // the same way home-panel.spec.ts's own session-snapshot test does.
  await context.addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  // '**/v1/me' is already routed to ME_BODY by the beforeEach above.
  await page.route('**/v1/devices', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({ ok: true, data: [], error: null, request_id: 'r' }),
    }),
  );

  // First page: a normal sign-in. Wait for the island to actually hydrate and persist
  // /v1/me (lib/data/query.ts) -- this is the fetch the second page's pre-paint script
  // will later read back out of localStorage.
  await page.goto('/account');
  await expect(page.getByTestId('account-menu-portrait')).toBeVisible();
  const persistedBeforeNav = await page.evaluate(() =>
    Object.keys(localStorage).some((k) => k.startsWith('fs.q.')),
  );
  expect(persistedBeforeNav).toBe(true);

  // Second page, a fresh navigation (a fresh document, a fresh hydration delay) --
  // Base.astro's own pre-paint script runs again here, independently of the first page.
  await page.goto('/setup');
  const snapshot = await page.evaluate(() => {
    const shell = document.querySelector('[data-testid="account-menu-preflight"]');
    const fallback = document.querySelector('[data-testid="account-menu-preflight-fallback"]');
    return {
      hydratedMenuPresent: document.querySelector('[data-testid="account-menu"]') !== null,
      ready: shell?.getAttribute('data-ready') ?? null,
      shellDisplay: shell ? getComputedStyle(shell).display : null,
      fallbackDisplay: fallback ? getComputedStyle(fallback).display : null,
      name: document.querySelector('[data-testid="account-menu-preflight-name"]')?.textContent ?? '',
    };
  });
  // The real, interactive chip has not mounted yet -- this really is the pre-hydration
  // window, not a race that happened to resolve before this check ran.
  expect(snapshot.hydratedMenuPresent).toBe(false);
  expect(snapshot.ready).toBe('1');
  expect(snapshot.shellDisplay).toBe('flex');
  expect(snapshot.fallbackDisplay).toBe('none');
  expect(snapshot.name).toBe('Simfury#1234');

  // Hydration then catches up and renders the exact same chip from the exact same cached
  // data -- nothing visibly swaps, and "Sign in" never becomes the active content either.
  await expect(page.getByTestId('account-menu-portrait')).toBeVisible();
  await expect(page.locator('[data-testid="session-nav-static"]')).toHaveCount(0);
});

test('a signed-out visitor never sees the preflight chip shell', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill({
      status: 401,
      contentType: 'application/json',
      body: JSON.stringify({ ok: false, data: null, error: { message: 'signed out' }, request_id: 'r' }),
    }),
  );
  await page.goto('/setup');
  const snapshot = await page.evaluate(() => {
    const shell = document.querySelector('[data-testid="account-menu-preflight"]');
    return {
      ready: shell?.getAttribute('data-ready') ?? null,
      shellDisplay: shell ? getComputedStyle(shell).display : null,
    };
  });
  expect(snapshot.ready).toBeNull();
  expect(snapshot.shellDisplay).toBe('none');
  await expect(page.getByRole('link', { name: 'Sign in' }).first()).toBeVisible();
});
