// web/tests/captures/premium-entry/capture.spec.ts
// Acceptance screenshots for design/specs/2026-10-29-premium-entry.md §12: the header in each
// state at 1440, 2000 and 390, the phone Menu, the account panel free beside Premium, the sim run
// area free and Premium, and the /premium landings. Not a test of behaviour; the e2e spec
// tests/e2e/premium-entry.spec.ts asserts. PNGs go to design/reviews/2026-10-29-premium-entry/.
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { test, type Page } from '@playwright/test';
import { mockEntitlements } from '../../../src/test-support/billing-fixtures';
import { OBNOXIOUS, armoryPointer, signInWith, signedOut, storePointer } from '../../e2e/support/selector';

const OUT = fileURLToPath(new URL('../../../../design/reviews/2026-10-29-premium-entry', import.meta.url));
const shot = (name: string): string => path.join(OUT, `${name}.png`);

const GRANT_ME = {
  user: { battletag: 'Hunthrop#1894' },
  entitlements: mockEntitlements({
    server_sims: true,
    history: true,
    notifications: true,
    log_coach: true,
    supporter_mark: true,
    billing: {
      plan: 'premium',
      status: 'active',
      current_period_end: null,
      cancel_at_period_end: false,
      founder: false,
      interval: null,
      amount_cents: null,
      source: 'grant',
    } as never,
  }),
};
const FREE_ME = { user: { battletag: 'Hunthrop#1894' }, entitlements: mockEntitlements() };

type State = 'signed-out' | 'free' | 'premium';

async function as(page: Page, state: State): Promise<void> {
  if (state === 'signed-out') {
    await signedOut(page);
    return;
  }
  if (state === 'premium') await page.addInitScript(() => localStorage.setItem('fs.premiumSource', 'grant'));
  await signInWith(page, [OBNOXIOUS], [], state === 'premium' ? GRANT_ME : FREE_ME);
  await storePointer(page, armoryPointer(OBNOXIOUS));
}

async function settle(page: Page): Promise<void> {
  await page.locator('[data-site-header][data-header-wired]').waitFor();
  await page.waitForLoadState('networkidle');
  await page.evaluate(() => document.fonts.ready);
}

for (const width of [1440, 2000, 1024, 390]) {
  for (const state of ['signed-out', 'free', 'premium'] as const) {
    test(`header ${state} ${width}`, async ({ page }) => {
      await page.setViewportSize({ width, height: 900 });
      await as(page, state);
      await page.goto('/');
      await settle(page);
      await page.locator('[data-site-header]').screenshot({ path: shot(`header-${state}-${width}`) });
    });
  }
}

test('header states: hover, focus-visible and current at 1440', async ({ page }) => {
  await page.setViewportSize({ width: 1440, height: 900 });
  await as(page, 'signed-out');
  await page.goto('/premium');
  await settle(page);
  await page.locator('[data-site-header]').screenshot({ path: shot('header-current-1440') });
  await page.goto('/');
  await settle(page);
  await page.getByTestId('nav-premium').hover();
  await page.locator('[data-site-header]').screenshot({ path: shot('header-hover-1440') });
  await page.mouse.move(0, 400);
  await page.getByTestId('nav-premium').focus();
  await page.keyboard.press('Shift+Tab');
  await page.keyboard.press('Tab');
  await page.locator('[data-site-header]').screenshot({ path: shot('header-focus-1440') });
});

for (const state of ['free', 'premium'] as const) {
  test(`phone menu ${state} 390`, async ({ page }) => {
    await page.setViewportSize({ width: 390, height: 844 });
    await as(page, state);
    await page.goto('/');
    await settle(page);
    await page.getByTestId('menu-button').click();
    await page.locator('[data-site-header]').screenshot({ path: shot(`phone-menu-${state}-390`) });
  });

  test(`account panel ${state} 1440`, async ({ page }) => {
    await page.setViewportSize({ width: 1440, height: 900 });
    await as(page, state);
    await page.goto('/guides');
    await settle(page);
    await page.getByTestId('character-selector').click();
    await page.getByTestId('selector-row-plan').waitFor();
    await page.screenshot({ path: shot(`account-panel-${state}-1440`) });
  });
}

test('account sheet free 390', async ({ page }) => {
  await page.setViewportSize({ width: 390, height: 844 });
  await as(page, 'free');
  await page.goto('/guides');
  await settle(page);
  await page.locator('[data-testid="character-selector"]:visible').first().click();
  await page.getByTestId('selector-row-plan').waitFor();
  await page.screenshot({ path: shot('account-sheet-free-390') });
});

for (const state of ['free', 'premium'] as const) {
  test(`sim run area ${state} 1440`, async ({ page }) => {
    await page.setViewportSize({ width: 1440, height: 900 });
    const envelope = (data: unknown) => ({ ok: true, data, error: null, request_id: 'r' });
    await page.route('**/v1/me', (route) =>
      route.fulfill({
        json: envelope({
          ...(state === 'premium' ? GRANT_ME : FREE_ME),
          user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
          characters: [
            {
              key: 'us/normal/thrallgar',
              region: 'us',
              ruleset: 'normal',
              name: 'Thrallgar',
              class: 'Warrior',
              build: { source: 'addon', captured_at: '2026-09-20T00:00:00Z' },
            },
          ],
          guilds: [],
        }),
      }),
    );
    await page.route('**/v1/characters/us/normal/thrallgar/sim-input', (route) =>
      route.fulfill({
        json: envelope({
          spec: 'warrior-fury',
          gear: { slots: [12640, 11726] },
          talents: '31/0/20',
          buffs: ['battle_shout'],
          race: 'orc',
          captured_at: new Date().toISOString(),
          source: 'addon',
        }),
      }),
    );
    await page.route('**/v1/sims?mine=1*', (route) =>
      route.fulfill({ json: envelope({ rows: [], total: 0, page: 1, per_page: 100 }) }),
    );
    await page.goto('/sim');
    await page.locator('[data-testid="sim-run-block-action"]:visible').first().click();
    const run = page.getByTestId('sim-run');
    await run.waitFor();
    await page.getByTestId('sim-run-button').filter({ hasText: /Run/ }).waitFor();
    await run.screenshot({ path: shot(`sim-run-${state}-1440`) });
  });
}

for (const [anchor, width, height] of [
  ['coach', 360, 640],
  ['coach', 1440, 900],
  ['server-sims', 1440, 900],
  ['server-sims', 390, 844],
  ['premium', 1440, 900],
] as const) {
  test(`/premium#${anchor} ${width}`, async ({ page }) => {
    await page.setViewportSize({ width, height });
    await as(page, 'signed-out');
    await page.goto(`/premium#${anchor}`);
    await page.evaluate(() => document.fonts.ready);
    await page.screenshot({ path: shot(`premium-${anchor}-${width}`) });
  });
}
