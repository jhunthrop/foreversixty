// web/tests/e2e/onboarding-second-screen.spec.ts
// W.9: the second screen of onboarding -- /setup (one gold button per step and a done line each),
// /logs' first look (one row style, the sign-in button, a card sized to its content), a real
// character with no logs, the landing's hero card, and the sample report's table on a phone. Each
// is captured at 1280 and 390 for the review (W9_SHOTS names the folder).
import { mkdirSync } from 'node:fs';
import { join } from 'node:path';
import { expect, test, type Page } from '@playwright/test';
import {
  CHARACTER,
  CHARACTER_PATH,
  CHARACTER_URL,
  GEAR,
  GEAR_ITEMS,
  routeGearIcons,
} from './support/character-fixture';
import { routeSampleReportMeta } from './support/logs-sample-fixture';

const SHOTS = process.env.W9_SHOTS ?? 'test-results/w9';
const SIZES = [
  { name: '1280', width: 1280, height: 900 },
  { name: '390', width: 390, height: 844 },
] as const;

const fulfil = (body: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify({ ok: status < 400, data: body, error: null, request_id: 'r' }),
});

const ago = (ms: number): string => new Date(Date.now() - ms).toISOString();
const HOUR = 3_600_000;

const ME_DONE = {
  user: { id: 1, battletag: 'hunthrop#1894', email: null, role: 'user', anonymize: false },
  characters: [
    {
      key: 'us/pvp/bow-jackzon',
      region: 'us',
      ruleset: 'pvp',
      name: 'Bow Jackzon',
      class: 'hunter',
      level: 24,
      build: { source: 'addon', captured_at: ago(HOUR) },
    },
  ],
  guilds: [],
};
const ME_NEW = { ...ME_DONE, characters: [] };

async function signedIn(page: Page, me: unknown, devices: unknown[]): Promise<void> {
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(me)));
  await page.route('**/v1/devices', (route) => route.fulfill(fulfil(devices)));
  await page.route('**/v1/reports?mine=1**', (route) =>
    route.fulfill(fulfil({ rows: [], total: 0, page: 1, per_page: 100 })),
  );
}

async function signedOut(page: Page): Promise<void> {
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(null, 401)));
  await page.route('**/v1/devices', (route) => route.fulfill(fulfil(null, 401)));
}

async function settle(page: Page): Promise<void> {
  await page.waitForLoadState('networkidle');
  await page.waitForTimeout(500);
}

for (const size of SIZES) {
  test.describe(`at ${size.name}`, () => {
    test.use({ viewport: { width: size.width, height: size.height } });
    test.beforeEach(() => mkdirSync(SHOTS, { recursive: true }));

    test('/setup signed out: three steps, one gold button each, the manual folded', async ({ page }) => {
      await signedOut(page);
      await page.goto('/setup');
      await settle(page);
      await expect(page.getByRole('heading', { name: 'Get set up', level: 1 })).toBeVisible();
      await expect(page.getByRole('link', { name: 'Sign in with Battle.net' })).toBeVisible();
      await expect(page.getByRole('link', { name: 'Install from CurseForge' })).toBeVisible();
      await expect(page.getByRole('link', { name: 'Download for Windows' })).toBeVisible();
      await expect(page.getByText('Windows only; Mac is not supported yet.')).toBeVisible();
      await expect(page.getByTestId('setup-line-signin')).toBeHidden();
      for (const id of ['setup-try-export', 'setup-everything']) {
        await expect(page.getByTestId(id)).not.toHaveAttribute('open', '');
      }
      await expect(page.getByTestId('setup-manual-install')).toHaveAttribute('open', '');
      await expect(page.getByText('Current data build')).toBeVisible();
      await page.screenshot({ path: join(SHOTS, `w9-setup-signed-out-${size.name}.png`), fullPage: true });
    });

    test('/setup signed in, everything done: a line and a gold chip per step', async ({ page }) => {
      await signedIn(page, ME_DONE, [
        {
          id: 'd1',
          name: 'JH-DESKTOP',
          platform: 'windows',
          created_at: ago(48 * HOUR),
          last_seen_at: ago(3 * 60_000),
        },
      ]);
      await page.goto('/setup');
      await settle(page);
      await expect(page.getByTestId('setup-line-signin')).toHaveText('Signed in as hunthrop#1894');
      await expect(page.getByTestId('setup-line-addon')).toHaveText(
        'Bow Jackzon arrived from the addon 1 hour ago',
      );
      await expect(page.getByTestId('setup-line-companion')).toHaveText(
        'Companion signed in on JH-DESKTOP · last seen 3 minutes ago',
      );
      await expect(page.locator('.setup-chip[data-done="true"]')).toHaveCount(3);
      await expect(page.getByRole('link', { name: 'Sign in with Battle.net' })).toBeHidden();
      await page.screenshot({
        path: join(SHOTS, `w9-setup-signed-in-done-${size.name}.png`),
        fullPage: true,
      });
    });

    test('/setup signed in, addon and companion pending', async ({ page }) => {
      await signedIn(page, ME_NEW, []);
      await page.goto('/setup');
      await settle(page);
      await expect(page.getByTestId('setup-line-signin')).toHaveText('Signed in as hunthrop#1894');
      await expect(page.getByTestId('setup-line-addon')).toHaveText(
        'No character yet. Log out once with the addon on and the companion running, or paste an export.',
      );
      await expect(page.getByTestId('setup-line-companion')).toHaveText('No companion signed in yet');
      await expect(page.locator('.setup-chip[data-done="true"]')).toHaveCount(1);
      await page.screenshot({
        path: join(SHOTS, `w9-setup-signed-in-pending-${size.name}.png`),
        fullPage: true,
      });
    });

    test('#paste opens the export disclosure that holds it', async ({ page }) => {
      await signedOut(page);
      await page.goto('/setup#paste');
      await settle(page);
      await expect(page.getByTestId('setup-try-export')).toHaveAttribute('open', '');
      await expect(page.getByTestId('addon-paste-box')).toBeVisible();
    });

    test('/logs: one row style, the sample dated like its page, the sign-in button, cards sized to content', async ({
      page,
    }) => {
      await signedOut(page);
      await page.route('**/v1/reports/recent**', (route) =>
        route.fulfill(
          fulfil({
            rows: [
              {
                id: 'named0000001',
                title: 'Blackrock Depths',
                created_at: new Date().toISOString(),
                fight_count: 14,
                kill_count: 3,
                zone: 'Blackrock Depths',
                logging_character: { key: 'us/pvp/bow-jackzon', name: 'Bow Jackzon' },
              },
            ],
          }),
        ),
      );
      await routeSampleReportMeta(page);
      await page.goto('/logs');
      await settle(page);
      const rows = page.getByTestId('report-row');
      await expect(rows).toHaveCount(2);
      await expect(rows.first()).toContainText('Bow Jackzon');
      await expect(rows.first()).toContainText('Blackrock Depths · 3 bosses · 14 fights · today');
      const titles = rows.locator('.report-title');
      const [a, b] = await Promise.all([
        titles.nth(0).evaluate((el) => getComputedStyle(el).color + getComputedStyle(el).fontSize),
        titles.nth(1).evaluate((el) => getComputedStyle(el).color + getComputedStyle(el).fontSize),
      ]);
      expect(a).toBe(b);
      expect(await titles.nth(0).evaluate((el) => getComputedStyle(el).opacity)).toBe('1');
      await expect(page.getByTestId('reports-sign-in')).toBeVisible();
      await expect(page.getByText('Sign in or paste an export to see your character here')).toHaveCount(0);
      await expect(page.getByTestId('logs-companion-link')).toHaveAttribute('href', '/setup#companion');
      await expect(page.getByTestId('pairing')).toHaveCount(0);
      await page.screenshot({ path: join(SHOTS, `w9-logs-${size.name}.png`), fullPage: true });
    });

    test('a real character with no logs: no typo hint, the gear card above, class and race tidy', async ({
      page,
    }) => {
      await signedOut(page);
      const data = {
        ...CHARACTER.data,
        character: { ...CHARACTER.data.character, class: 'priest' },
        gear: [],
        best: [],
        history: [],
        builds_seen: [],
        score_history: [],
      };
      await page.route(CHARACTER_URL, (route) =>
        route.fulfill({
          status: 200,
          contentType: 'application/json',
          body: JSON.stringify({ ...CHARACTER, data }),
        }),
      );
      await page.goto(CHARACTER_PATH);
      await settle(page);
      // P.1: a visitor to a character with nothing on it gets one sentence and the lookup box.
      await expect(page.getByTestId('character-not-found-line')).toHaveText(
        'Elyra Duskvale has no Forever Sixty data yet.',
      );
      await expect(page.getByText('Check the name and ruleset')).toHaveCount(0);
      await expect(page.getByText('(Player)')).toHaveCount(0);
      await page.screenshot({ path: join(SHOTS, `w9-character-no-logs-${size.name}.png`), fullPage: true });
    });

    test('the landing hero card follows the clock, and no stale This site row is shown', async ({ page }) => {
      await page.clock.install({ time: new Date('2026-10-23T12:00:00Z') });
      await signedOut(page);
      await page.route('**/v1/reports/recent**', (route) => route.fulfill(fulfil({ rows: [] })));
      await page.route('**/v1/rankings/guilds?**', (route) => route.fulfill(fulfil({ rows: [] })));
      await routeSampleReportMeta(page);
      await page.goto('/');
      await page.clock.resume();
      await settle(page);
      const card = page.getByTestId('home-hero-card');
      await expect(card).toContainText('Launch Nov 4 · 15:00 PST');
      await expect(page.getByTestId('home-hero-countdown-days')).toContainText('12');
      await expect(page.getByTestId('home-hero-countdown')).toContainText('Minutes');
      await expect(card).not.toContainText('Beta is live');
      await expect(
        page.getByTestId('home-news-item').filter({ has: page.locator('[href*="simulator-tools"]') }),
      ).toBeHidden();
      await page.screenshot({ path: join(SHOTS, `w9-landing-after-beta-${size.name}.png`), fullPage: false });
    });
  });
}

test.describe('the beta card while the beta runs', () => {
  test.use({ viewport: { width: 1280, height: 900 } });
  test('reads live with its cap on Oct 20', async ({ page }) => {
    await page.clock.install({ time: new Date('2026-10-20T12:00:00Z') });
    await signedOut(page);
    await page.route('**/v1/reports/recent**', (route) => route.fulfill(fulfil({ rows: [] })));
    await page.route('**/v1/rankings/guilds?**', (route) => route.fulfill(fulfil({ rows: [] })));
    await page.goto('/');
    await page.clock.resume();
    await settle(page);
    const card = page.getByTestId('home-hero-card');
    await expect(card).toContainText('Beta is live');
    await expect(card).toContainText('level cap 30');
    await mkdirSync(SHOTS, { recursive: true });
    await page.screenshot({ path: join(SHOTS, 'w9-landing-beta-live-1280.png') });
  });
});

test.describe('the sample report on a phone', () => {
  test.use({ viewport: { width: 390, height: 844 } });
  test('folds the fight list to one row, so the damage rows start on the first screen', async ({ page }) => {
    await signedOut(page);
    await page.goto('/reports/sample');
    await settle(page);
    const toggle = page.getByTestId('fight-selector-toggle');
    await expect(toggle).toBeVisible();
    await expect(page.getByTestId('fight-selector-body')).toBeHidden();
    const table = page.getByTestId('roster-name').first();
    await expect(table).toBeVisible();
    const top = await table.evaluate((el) => el.getBoundingClientRect().top + window.scrollY);
    expect(top).toBeLessThan(844);
    mkdirSync(SHOTS, { recursive: true });
    await page.screenshot({ path: join(SHOTS, 'w9-sample-report-390.png'), fullPage: false });
    await toggle.click();
    await expect(page.getByTestId('fight-selector-body')).toBeVisible();
  });
});

test.describe('a character with gear and no score', () => {
  test.use({ viewport: { width: 1280, height: 900 } });
  test('leads with the gear card, above the empty score card', async ({ page }) => {
    await signedOut(page);
    const data = { ...CHARACTER.data, ...GEAR, best: [], history: [], builds_seen: [], score_history: [] };
    await routeGearIcons(page);
    await page.route('**/data/*/items/priest.json', (route) =>
      route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(GEAR_ITEMS) }),
    );
    await page.route(CHARACTER_URL, (route) =>
      route.fulfill({
        status: 200,
        contentType: 'application/json',
        body: JSON.stringify({ ...CHARACTER, data }),
      }),
    );
    await page.goto(CHARACTER_PATH);
    await settle(page);
    const gear = await page.getByTestId('character-gear').boundingBox();
    const score = await page.getByTestId('forever-score').boundingBox();
    expect(gear).not.toBeNull();
    expect(score).not.toBeNull();
    expect(gear!.y).toBeLessThan(score!.y);
    mkdirSync(SHOTS, { recursive: true });
    await page.screenshot({ path: join(SHOTS, 'w9-character-gear-first-1280.png'), fullPage: true });
  });
});

test.describe('the not-found character page', () => {
  test.use({ viewport: { width: 1280, height: 900 } });
  test('reads at full strength once settled', async ({ page }) => {
    await signedOut(page);
    await page.route(CHARACTER_URL, (route) => route.fulfill(fulfil(null, 404)));
    await page.goto(CHARACTER_PATH);
    await settle(page);
    await expect(page.getByTestId('character-not-found')).toBeVisible();
    expect(await page.getByTestId('character-name').evaluate((el) => getComputedStyle(el).opacity)).toBe('1');
    mkdirSync(SHOTS, { recursive: true });
    await page.screenshot({ path: join(SHOTS, 'w9-character-not-found-1280.png') });
  });
});
