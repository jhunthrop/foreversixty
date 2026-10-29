// web/tests/e2e/premium-entry.spec.ts
// Premium in the flow (design/specs/2026-10-29-premium-entry.md): the header's Premium tag in
// each account state with no layout shift (the pre-paint hide and the after-paint hold), its
// widths, the phone Menu row, the account panel's Plan row, and every contextual link landing on
// its own /premium anchor with the line highlighted.
import { expect, test, type Page } from '@playwright/test';
import { mockEntitlements } from '../../src/test-support/billing-fixtures';
import { headerWired } from './support/header';
import { OBNOXIOUS, armoryPointer, signInWith, signedOut, storePointer } from './support/selector';
import { routeCoach, STATES, ROWS_URL, signInAs } from './support/coach-stubs';

const HINT_KEY = 'fs.premiumSource';
const GRANT_BILLING = {
  plan: 'premium',
  status: 'active',
  current_period_end: null,
  cancel_at_period_end: false,
  founder: false,
  interval: null,
  amount_cents: null,
  source: 'grant',
};
const GRANT_ME = {
  user: { battletag: 'Hunthrop#1894' },
  entitlements: mockEntitlements({
    server_sims: true,
    history: true,
    notifications: true,
    log_coach: true,
    supporter_mark: true,
    billing: GRANT_BILLING as never,
  }),
};
const FREE_ME = { user: { battletag: 'Hunthrop#1894' }, entitlements: mockEntitlements() };
const DEFAULT_GROUND = 'rgba(229, 185, 85, 0.08)';
const TINT = 'rgba(229, 185, 85, 0.08)';
const MENU_ROW_MIN = 52;

const tag = (page: Page) => page.getByTestId('nav-premium');
const navState = (page: Page) => page.evaluate(() => document.documentElement.dataset.premiumNav ?? null);

async function storeHint(page: Page, source: string): Promise<void> {
  await page.addInitScript(([key, value]) => window.localStorage.setItem(key, value), [
    HINT_KEY,
    source,
  ] as const);
}

async function signIn(page: Page, me: Record<string, unknown>): Promise<void> {
  await signInWith(page, [OBNOXIOUS], [], me);
  await storePointer(page, armoryPointer(OBNOXIOUS));
}

declare global {
  interface Window {
    __barShift: number;
    __holdMoved: number | null;
  }
}

/**
 * Sums the layout shifts whose moved nodes sit in the site header (the bar's own CLS). The bar
 * already shifts by about 0.0006 on load without the tag (the doors' web-font swap, measured on a
 * build from before this change), so the tag's states are held to that floor, not to zero.
 */
const BAR_FONT_SWAP_FLOOR = 0.001;

/** Starts the bar's shift sum and the hold-frame probe; call before `page.goto`. */
async function trackBarShift(page: Page): Promise<void> {
  await page.addInitScript(() => {
    window.__barShift = 0;
    window.__holdMoved = null;
    // The hold path: the doors' left edge sampled just before `hold` lands and right after it
    // (a light timer, not a frame loop, so the page's idle-time hydration still runs).
    let lastLeft: number | null = null;
    const sampler = setInterval(() => {
      const setup = document.querySelector('[data-door="setup"]');
      if (setup !== null) lastLeft = setup.getBoundingClientRect().left;
    }, 40);
    setTimeout(() => clearInterval(sampler), 20_000);
    document.addEventListener('DOMContentLoaded', () => {
      new MutationObserver(() => {
        if (document.documentElement.dataset.premiumNav !== 'hold' || window.__holdMoved !== null) return;
        const setup = document.querySelector('[data-door="setup"]');
        if (setup !== null && lastLeft !== null)
          window.__holdMoved = setup.getBoundingClientRect().left - lastLeft;
      }).observe(document.documentElement, { attributes: true, attributeFilter: ['data-premium-nav'] });
    });
    new PerformanceObserver((list) => {
      for (const entry of list.getEntries() as (PerformanceEntry & {
        value: number;
        hadRecentInput: boolean;
        sources?: { node?: Node | null }[];
      })[]) {
        if (entry.hadRecentInput) continue;
        const inBar = (entry.sources ?? []).filter((source) =>
          source.node instanceof Element ? source.node.closest('[data-site-header]') !== null : false,
        );
        if (inBar.length === 0) continue;
        window.__barShift += entry.value;
      }
    }).observe({ type: 'layout-shift', buffered: true });
  });
}

async function readBarShift(page: Page): Promise<number> {
  await page.waitForTimeout(500);
  return page.evaluate(() => window.__barShift);
}

/** Loads `/` and waits for the header wiring and the selector's `/v1/me` answer to settle. */
async function openHome(page: Page): Promise<number> {
  await trackBarShift(page);
  await page.goto('/');
  await headerWired(page);
  await page.waitForLoadState('networkidle');
  return readBarShift(page);
}

test.describe('the bar tag', () => {
  test.use({ viewport: { width: 1440, height: 900 } });
  test.skip(() => test.info().project.name !== 'desktop', 'desktop bar only');

  test('signed out: the quiet gold tag ends the doors, after Get set up, with no shift', async ({ page }) => {
    await signedOut(page);
    const cls = await openHome(page);
    await expect(tag(page)).toBeVisible();
    const last = page.locator('#primary-nav > :last-child');
    await expect(last).toHaveAttribute('data-door', 'premium');
    await expect(page.locator('#primary-nav > :nth-last-child(2)')).toHaveAttribute('data-door', 'setup');
    const pill = tag(page).locator('.pill');
    await expect(pill).toHaveText(/premium/i);
    await expect(pill).toHaveCSS('background-color', DEFAULT_GROUND);
    await expect(tag(page).locator('svg')).toHaveCount(0);
    expect(cls).toBeLessThan(BAR_FONT_SWAP_FLOOR);
  });

  test('free: the tag shows and stays', async ({ page }) => {
    await signIn(page, FREE_ME);
    const cls = await openHome(page);
    await expect(tag(page)).toBeVisible();
    expect(await navState(page)).toBeNull();
    expect(cls).toBeLessThan(BAR_FONT_SWAP_FLOOR);
  });

  test('Premium with the stored hint: hidden before first paint, no shift', async ({ page }) => {
    await storeHint(page, 'grant');
    await signIn(page, GRANT_ME);
    await page.addInitScript(() => {
      document.addEventListener('DOMContentLoaded', () => {
        (window as unknown as { __navAtLoad: string | null }).__navAtLoad =
          document.documentElement.dataset.premiumNav ?? null;
      });
    });
    const cls = await openHome(page);
    expect(await page.evaluate(() => (window as unknown as { __navAtLoad: string | null }).__navAtLoad)).toBe(
      'hide',
    );
    await expect(tag(page)).toBeHidden();
    await expect(tag(page)).toHaveCSS('display', 'none');
    expect(cls).toBeLessThan(BAR_FONT_SWAP_FLOOR);
  });

  test('Premium on a fresh profile: the tag is held after /v1/me, its box kept, no shift', async ({
    page,
  }) => {
    await signIn(page, GRANT_ME);
    const cls = await openHome(page);
    await expect.poll(() => navState(page)).toBe('hold');
    await expect(tag(page)).toHaveCSS('visibility', 'hidden');
    expect(await page.evaluate(() => window.__holdMoved)).toBe(0);
    const box = await tag(page).boundingBox();
    expect(box?.width ?? 0).toBeGreaterThan(40);
    expect(cls).toBeLessThan(BAR_FONT_SWAP_FLOOR);
    // The next page reads the hint the session wrote and hides the tag before paint.
    await page.goto('/guides');
    expect(await navState(page)).toBe('hide');
  });

  for (const source of ['own', 'seat']) {
    test(`the ${source} hint hides it before paint`, async ({ page }) => {
      await storeHint(page, source);
      await signIn(page, GRANT_ME);
      await page.goto('/');
      expect(await navState(page)).toBe('hide');
      await expect(tag(page)).toBeHidden();
    });
  }

  test('marks itself current on /premium', async ({ page }) => {
    await signedOut(page);
    await page.goto('/premium');
    await expect(tag(page)).toHaveAttribute('aria-current', 'page');
    await expect(tag(page).locator('.pill')).toHaveCSS('border-color', 'rgb(229, 185, 85)');
  });
});

test.describe('the bar widths', () => {
  test.skip(() => test.info().project.name !== 'desktop', 'desktop bar only');

  test('the tag stays in the bar at every desktop width, with no overflow', async ({ page }) => {
    await signIn(page, FREE_ME);
    for (const [width, shown] of [
      [1024, true],
      [1100, true],
      [1149, true],
      [1150, true],
      [1280, true],
      [2000, true],
    ] as const) {
      await page.setViewportSize({ width, height: 800 });
      await page.goto('/');
      await headerWired(page);
      if (shown) await expect(tag(page), `${width}`).toBeVisible();
      else await expect(tag(page), `${width}`).toBeHidden();
      const fit = await page.evaluate(() => {
        const nav = document.getElementById('primary-nav')!.getBoundingClientRect();
        const right = document.querySelector('.site-right')!.getBoundingClientRect();
        const header = document.querySelector('.site-header')!;
        return { gap: right.left - nav.right, scroll: header.scrollWidth - header.clientWidth };
      });
      expect(fit.gap, `${width}`).toBeGreaterThan(0);
      expect(fit.scroll, `${width}`).toBe(0);
    }
  });

  test('at 2000 the tag sits inside the 1344 column', async ({ page }) => {
    await signedOut(page);
    await page.setViewportSize({ width: 2000, height: 900 });
    await page.goto('/');
    const box = (await tag(page).boundingBox())!;
    expect(box.x).toBeGreaterThan((2000 - 1344) / 2);
    expect(box.x + box.width).toBeLessThan((2000 + 1344) / 2);
  });
});

test.describe('the phone Menu row', () => {
  test.use({ viewport: { width: 390, height: 844 } });

  test('free: the tag shares the last row with Get set up, right-aligned, 52 px, no chevron', async ({
    page,
  }) => {
    await signIn(page, FREE_ME);
    await page.goto('/');
    await page.getByTestId('menu-button').click();
    await expect(tag(page)).toBeVisible();
    await expect(page.locator('#primary-nav > :last-child')).toHaveAttribute('data-door', 'premium');
    const setup = (await page.locator('#primary-nav [data-door="setup"]').boundingBox())!;
    const row = (await tag(page).boundingBox())!;
    const menu = (await page.getByTestId('primary-nav').boundingBox())!;
    // One row: the same top and height, Get set up at the menu's left edge, the tag at its right.
    expect(Math.abs(row.y - setup.y)).toBeLessThanOrEqual(1);
    expect(row.height).toBeGreaterThanOrEqual(MENU_ROW_MIN);
    expect(setup.height).toBeGreaterThanOrEqual(MENU_ROW_MIN);
    expect(Math.abs(setup.x - menu.x)).toBeLessThanOrEqual(1);
    expect(Math.abs(row.x + row.width - (menu.x + menu.width))).toBeLessThanOrEqual(1);
    expect(setup.x + setup.width).toBeLessThan(row.x);
    await expect(tag(page).locator('svg')).toHaveCount(0);
  });

  test('Premium: no tag; Get set up keeps the whole row', async ({ page }) => {
    await storeHint(page, 'own');
    await signIn(page, GRANT_ME);
    await page.goto('/');
    await page.getByTestId('menu-button').click();
    await expect(tag(page)).toBeHidden();
    const setup = (await page.locator('#primary-nav [data-door="setup"]').boundingBox())!;
    const menu = (await page.getByTestId('primary-nav').boundingBox())!;
    expect(Math.abs(setup.width - menu.width)).toBeLessThanOrEqual(1);
  });
});

test.describe('the account panel Plan row', () => {
  test.use({ viewport: { width: 1440, height: 900 } });
  test.skip(() => test.info().project.name !== 'desktop', 'desktop panel');

  test('free: Plan, Free, See plans as one link that lands on the Premium card', async ({ page }) => {
    await signIn(page, FREE_ME);
    await page.goto('/guides');
    await page.getByTestId('character-selector').click();
    const row = page.getByTestId('selector-row-plan');
    await expect(row).toHaveAccessibleName(/^Plan\s+Free\s+See plans$/);
    await expect(row.getByTestId('selector-plan-see-plans')).toHaveCSS('color', 'rgb(229, 185, 85)');
    await row.click();
    await expect(page).toHaveURL(/\/premium#premium$/);
    await expect(page.locator('#premium')).toBeInViewport();
  });

  test('Premium: no See plans', async ({ page }) => {
    await signIn(page, GRANT_ME);
    await page.goto('/guides');
    await page.getByTestId('character-selector').click();
    await expect(page.getByTestId('selector-row-plan')).toHaveAttribute('href', '/account#plan');
    await expect(page.getByTestId('selector-plan-see-plans')).toHaveCount(0);
  });
});

test.describe('the /premium landings', () => {
  test.use({ viewport: { width: 1440, height: 900 } });
  test.skip(() => test.info().project.name !== 'desktop', 'desktop landings');

  for (const [anchor, lines] of [
    ['coach', 1],
    ['server-sims', 3],
    ['history', 1],
    ['notifications', 1],
  ] as const) {
    test(`#${anchor}: lands in view, ${lines} line(s) tinted, no box changes`, async ({ page }) => {
      await page.goto('/premium');
      const plain = await page.locator(`[data-anchor="${anchor}"]`).first().boundingBox();
      await page.goto(`/premium#${anchor}`);
      await expect(page.locator(`#${anchor}`)).toBeInViewport();
      const tinted = page.locator(`[data-anchor="${anchor}"]`);
      await expect(tinted).toHaveCount(lines);
      for (const line of await tinted.all()) await expect(line).toHaveCSS('background-color', TINT);
      await expect(page.locator('li.feature:not([data-anchor])').first()).not.toHaveCSS(
        'background-color',
        TINT,
      );
      const landed = await tinted.first().boundingBox();
      expect(landed?.height).toBe(plain?.height);
      expect(landed?.width).toBe(plain?.width);
    });
  }

  test('#premium and #guild-leaders land in view', async ({ page }) => {
    for (const anchor of ['premium', 'guild-leaders']) {
      await page.goto(`/premium#${anchor}`);
      await expect(page.locator(`#${anchor}`)).toBeInViewport();
    }
  });
});

test.describe('/premium#coach on a phone', () => {
  test.use({ viewport: { width: 360, height: 640 } });

  test('the card title stays in view above the tinted coach line', async ({ page }) => {
    await page.goto('/premium#coach');
    await expect(page.locator('#coach')).toBeInViewport();
    await expect(page.locator('#premium .plan-name')).toBeInViewport();
    await expect(page.locator('#coach')).toHaveCSS('background-color', TINT);
  });
});

test.describe('the contextual links land on their anchors', () => {
  test.use({ viewport: { width: 1440, height: 900 } });
  test.skip(() => test.info().project.name !== 'desktop', 'desktop');

  test('the report coach teaser lands on #coach', async ({ page }) => {
    await signInAs(page, 'free');
    await routeCoach(page, STATES['teaser-several']);
    await page.goto(ROWS_URL);
    await page.getByTestId('coach-row-open').click();
    await page.getByTestId('coach-upgrade').click();
    await expect(page).toHaveURL(/\/premium#coach$/);
    await expect(page.locator('#coach')).toBeInViewport();
    await expect(page.locator('#coach')).toHaveCSS('background-color', TINT);
  });

  test('the sim server-run link is muted text that lands on #server-sims', async ({ page }) => {
    // A member with one stored build, the sim landing's own fixture shape (sim-landing.spec.ts).
    const envelope = (data: unknown) => ({ ok: true, data, error: null, request_id: 'r' });
    await page.route('**/v1/me', (route) =>
      route.fulfill({
        json: envelope({
          ...FREE_ME,
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
    const link = page.getByTestId('sim-server-run-free');
    await expect(link).toBeVisible({ timeout: 15_000 });
    await expect(link).toHaveText('Run on our servers · Premium plan');
    const run = await page.getByTestId('sim-run-button').evaluate((el) => getComputedStyle(el).fontWeight);
    const quiet = await link.evaluate((el) => ({
      weight: getComputedStyle(el).fontWeight,
      border: getComputedStyle(el).borderTopWidth,
      bg: getComputedStyle(el).backgroundColor,
    }));
    expect(Number(quiet.weight)).toBeLessThan(Number(run));
    expect(quiet.border).toBe('0px');
    expect(quiet.bg).toBe('rgba(0, 0, 0, 0)');
    await link.click();
    await expect(page).toHaveURL(/\/premium#server-sims$/);
    await expect(page.locator('#server-sims')).toBeInViewport();
  });
});
