import { expect, test } from '@playwright/test';
import { collectPageErrors } from './support/console';

test('homepage states the one fixed sentence (tenet 14) and stops, not a marketing slogan', async ({
  page,
}) => {
  const errors = collectPageErrors(page);
  // Top guilds (spec §2.4/§3.A.6, "Around the site") reads GET /v1/rankings/guilds live;
  // stub it so the ready state is deterministic instead of depending on the real API
  // answering (the way rankings.spec.ts, rankings-phone.spec.ts and
  // rankings-encounter-picker.spec.ts already stub this same endpoint).
  await page.route('**/v1/rankings/guilds?**', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({
        ok: true,
        data: {
          rows: [
            {
              rank: 1,
              guild: { name: 'The Last Watch', ruleset: 'hardcore', region: 'us' },
              value: 9,
              fought_at: '2026-12-09T22:10:00Z',
              report_id: 'fixture2abcd',
            },
          ],
        },
        error: null,
        request_id: 'r',
      }),
    }),
  );
  await page.goto('/');
  const h1 = page.locator('h1');
  await expect(h1).toHaveCount(1);
  const heading = (await h1.innerText()).trim();
  // Home rebuild spec §1 / tenet 14: the one fixed sentence, never a slogan -- short, and
  // it never ends with "!" or "?".
  expect(heading.length).toBeLessThanOrEqual(90);
  expect(heading).not.toMatch(/[!?]$/);
  expect(heading).toBe('Play your class better.');
  await expect(page.getByTestId('home-timeline')).toBeVisible();
  // The nine-class picker (§3.A.1) and the larger "Best in slot by class" row (§3.A.4).
  await expect(page.getByTestId('home-class-picker').getByTestId(/^home-class-picker-/)).toHaveCount(9);
  await expect(
    page.getByTestId('home-class-picker-large').getByRole('heading', { name: 'Best in slot by class' }),
  ).toBeVisible();
  // Five doors (§3.A.5): unchanged copy, now with a visible "Open the X" line per card.
  await expect(page.getByTestId('home-five-doors-planner')).toContainText('Open the planner');
  await expect(page.getByTestId('home-five-doors-simulator')).toContainText('Open the simulator');
  await expect(page.getByTestId('home-five-doors-logs')).toContainText('Open the logs');
  await expect(page.getByTestId('home-five-doors-rankings')).toContainText('Open the rankings');
  await expect(page.getByTestId('home-five-doors-guides')).toContainText('Open guides');
  // Spec §3.A.6 "Around the site": Recent reports and Top guilds, unchanged mechanism.
  await expect(page.getByTestId('recent-reports')).toBeVisible();
  await page.getByTestId('home-around-the-site').scrollIntoViewIfNeeded();
  await expect(page.getByTestId('home-top-guilds')).toBeVisible();
  // Signed-in-only regions are absent from the signed-out page.
  await expect(page.getByTestId('home-upgrades')).toBeHidden();
  await expect(page.getByTestId('home-another-class')).toBeHidden();
  await expect(page.getByRole('heading', { name: 'Latest' })).toBeVisible();
  await expect(page.getByText('Not affiliated with or endorsed by Blizzard Entertainment')).toBeVisible();
  expect(errors).toEqual([]);
});

test('every class crest in the hero picker links to its class guide', async ({ page }) => {
  await page.goto('/');
  await expect(page.getByTestId('home-class-picker-warrior')).toHaveAttribute('href', '/guides/warrior');
  await expect(page.getByTestId('home-class-picker-hunter')).toHaveAttribute('href', '/guides/hunter');
  await expect(page.getByRole('link', { name: 'All 28 specs' }).first()).toHaveAttribute('href', '/bis');
});

test("content pages ship only the layout's account menu island", async ({ page }) => {
  for (const path of ['/about', '/premium']) {
    await page.goto(path);
    const islands = page.locator('astro-island');
    await expect(islands).toHaveCount(1);
    await expect(islands.first()).toHaveAttribute('component-url', /AccountMenu/);
  }
});

test('a keyboard user can skip the header straight to the content', async ({ page }) => {
  await page.goto('/');
  const skip = page.getByRole('link', { name: 'Skip to content' });
  const offScreen = await skip.boundingBox();
  expect(offScreen?.height ?? 0).toBeLessThanOrEqual(1);

  await page.keyboard.press('Tab');
  await expect(skip).toBeFocused();
  const onScreen = await skip.boundingBox();
  expect(onScreen?.height ?? 0).toBeGreaterThanOrEqual(44);

  await page.keyboard.press('Enter');
  await expect(page).toHaveURL(/#main$/);
  await expect(page.locator('main')).toBeFocused();
});

test('the header and footer navigations are distinguishable landmarks', async ({ page }) => {
  await page.goto('/');
  const menu = page.getByTestId('menu-button');
  if (await menu.isVisible()) await menu.click();
  await expect(page.getByRole('navigation', { name: 'Primary' })).toBeVisible();
  await expect(page.getByRole('navigation', { name: 'Footer' })).toBeVisible();
});

test('at desktop width, the hero class picker stays inside its own column, never the page', async ({
  page,
}, testInfo) => {
  test.skip(testInfo.project.name !== 'desktop', 'lg two-column hero only');
  await page.goto('/');

  const viewport = page.viewportSize();
  expect(viewport, 'desktop project always sets a viewport').not.toBeNull();
  const viewportWidth = viewport!.width;

  const accountBlock = page.getByTestId('home-account-block');
  const accountBlockBox = await accountBlock.boundingBox();
  expect(accountBlockBox?.x ?? 0, 'the hero column itself starts on-screen').toBeGreaterThanOrEqual(0);
  expect(
    accountBlockBox!.x + accountBlockBox!.width,
    'the hero column stays inside the viewport',
  ).toBeLessThanOrEqual(viewportWidth + 1);

  // The timeline now lives in its own full-width section, outside the sky band's hero grid
  // entirely (home rebuild spec §3.A.3 must-fix), so it can no longer inherit a column's
  // overflow bug the way the old two-column layout could -- this still checks it stays on
  // screen at desktop width.
  const timeline = page.getByTestId('home-timeline');
  const timelineBox = await timeline.boundingBox();
  expect(
    timelineBox!.x + timelineBox!.width,
    'the timeline card stays inside the viewport',
  ).toBeLessThanOrEqual(viewportWidth + 1);
});
