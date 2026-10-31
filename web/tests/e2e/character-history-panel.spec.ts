// web/tests/e2e/character-history-panel.spec.ts
// The History section's own behaviour in a real browser (experience spec
// 2026-10-24-character-history 3.7, 3.8, 3.11): selection against the fights list, the keyboard,
// the boss picker's roving focus and the reserve the next visit holds. The lane's acceptance
// e2e is character-history.spec.ts; this file keeps the panel's interaction rules.
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import {
  CHARACTER,
  CHARACTER_PATH,
  CHARACTER_URL,
  ownerMeAnswer,
  routeCharacterHistory,
} from './support/character-fixture';
import { HISTORY_MAX_RESERVE_PX } from '../../src/lib/character/history-layout';

const SCORED = {
  ...CHARACTER,
  data: {
    ...CHARACTER.data,
    score: 1830,
    tier: 'epic',
    as_of: '2026-11-04',
    components: { gear: 412, experience: 640, parses: 510, attendance: 180, discipline: 88 },
    experience: { dungeons: 4, raid_kills: 12, alts: 1 },
  },
};
const RESERVE_KEY = 'fs.history.reserve';

/** Polls until the window has stopped scrolling, so a later reading of scrollY is the settled one. */
async function expectScrollSettled(page: Page): Promise<void> {
  let last = Number.NaN;
  await expect
    .poll(async () => {
      const now = await page.evaluate(() => window.scrollY);
      const still = now === last;
      last = now;
      return still;
    })
    .toBe(true);
}

async function openHistory(page: Page, history: number | undefined = undefined): Promise<void> {
  await page.route(CHARACTER_URL, (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(SCORED) }),
  );
  await page.route(`${CHARACTER_URL}/rating`, (route) => route.fulfill({ status: 404, body: '{}' }));
  const me = ownerMeAnswer() as { data: { entitlements?: Record<string, unknown> } };
  me.data.entitlements = { ...(me.data.entitlements ?? {}), history: true };
  await page.route('**/v1/me', (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(me) }),
  );
  await (history === undefined ? routeCharacterHistory(page) : routeCharacterHistory(page, history));
  await page.goto(CHARACTER_PATH);
}

const reserve = (page: Page): Promise<string | null> =>
  page.evaluate((key) => localStorage.getItem(key), RESERVE_KEY);

test('selecting a hidden row opens the list once and a later hover never moves it', async ({ page }) => {
  await page.setViewportSize({ width: 1440, height: 800 });
  await openHistory(page);
  const score = page.getByTestId('history-score');
  const plot = score.getByTestId('history-plot');
  const toggle = score.getByTestId('history-score-fights-toggle');
  await expect(plot).toBeVisible();
  await expect(toggle).toHaveAttribute('aria-expanded', 'false');
  // The owner's bar above the profile lands after /v1/me; once it is in, nothing above the plot
  // changes height, so scroll anchoring cannot move scrollY under the hover check below.
  await expect(page.getByTestId('current-character-bar').first()).toBeVisible();

  await plot.focus();
  await page.keyboard.press('Home');
  await expect(score.getByTestId('history-readout-value')).toHaveText('740');
  await expect(toggle).toHaveAttribute('aria-expanded', 'true');
  await expect(toggle).toHaveText('Show fewer');
  const marked = score.locator('[data-marked="true"]');
  await expect(marked).toBeAttached();
  await expect(marked).toBeInViewport({ ratio: 0.5, timeout: 15_000 });

  await toggle.click();
  await expect(toggle).toHaveAttribute('aria-expanded', 'false');
  await plot.scrollIntoViewIfNeeded();
  await expectScrollSettled(page);
  const scrolled = await page.evaluate(() => window.scrollY);
  await plot.hover({ position: { x: 300, y: 100 } });
  await plot.hover({ position: { x: 500, y: 100 } });
  await page.mouse.move(0, 0);
  await expect(toggle).toHaveAttribute('aria-expanded', 'false');
  expect(await page.evaluate(() => window.scrollY)).toBe(scrolled);
});

test('the plot answers the arrow keys, Home and End, and a hover previews without selecting', async ({
  page,
}) => {
  await page.setViewportSize({ width: 1440, height: 800 });
  await openHistory(page);
  const score = page.getByTestId('history-score');
  const plot = score.getByTestId('history-plot');
  const value = score.getByTestId('history-readout-value');
  await expect(plot).toBeVisible();
  await plot.focus();
  await page.keyboard.press('ArrowLeft');
  await expect(value).toHaveText('1700');
  await page.keyboard.press('ArrowRight');
  await expect(value).toHaveText('1830');
  await page.keyboard.press('Home');
  await expect(value).toHaveText('740');
  await page.keyboard.press('End');
  await expect(value).toHaveText('1830');
  await plot.hover({ position: { x: 100, y: 100 } });
  await expect(value).not.toHaveText('1830');
  await page.mouse.move(0, 0);
  await expect(value).toHaveText('1830');
});

test('the boss picker is one tab stop with roving focus; Enter selects', async ({ page }) => {
  await page.setViewportSize({ width: 1440, height: 800 });
  await openHistory(page);
  const pills = page.getByTestId('history-pill');
  await expect(pills).toHaveCount(2);
  await expect(pills.nth(0)).toHaveAttribute('tabindex', '0');
  await expect(pills.nth(1)).toHaveAttribute('tabindex', '-1');
  await pills.nth(0).focus();
  await page.keyboard.press('ArrowRight');
  await expect(pills.nth(1)).toBeFocused();
  await page.keyboard.press('ArrowLeft');
  await expect(pills.nth(0)).toBeFocused();
  await page.keyboard.press('ArrowRight');
  await page.keyboard.press('Enter');
  await expect(pills.nth(1)).toHaveAttribute('aria-selected', 'true');
  await expect(pills.nth(1)).toHaveAttribute('tabindex', '0');
  await expect(page.getByTestId('history-parse-one-kill')).toBeVisible();
});

test('the reserve keeps the collapsed box, not an opened list', async ({ page }) => {
  await page.setViewportSize({ width: 1440, height: 800 });
  await openHistory(page);
  const score = page.getByTestId('history-score');
  const toggle = score.getByTestId('history-score-fights-toggle');
  await expect(toggle).toBeVisible();
  await expect.poll(() => reserve(page)).not.toBeNull();
  const collapsed = await reserve(page);
  await toggle.click();
  await expect(toggle).toHaveText('Show fewer');
  await page.waitForTimeout(300);
  // Under the bound the write would be accepted, so an unchanged value proves the skip.
  const expanded = await page.getByTestId('character-history-section').boundingBox();
  expect(expanded?.height ?? Number.POSITIVE_INFINITY).toBeLessThan(HISTORY_MAX_RESERVE_PX);
  expect(await reserve(page)).toBe(collapsed);
});

test('the reserve stores nothing for the error state', async ({ page }) => {
  await page.setViewportSize({ width: 1440, height: 800 });
  await openHistory(page, 500);
  await expect(page.getByTestId('history-error')).toBeVisible();
  await page.waitForTimeout(300);
  expect(await reserve(page)).toBeNull();
});
