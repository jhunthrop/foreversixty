import { expect, test } from '@playwright/test';
import { collectPageErrors } from './support/console';

test('homepage states the product and stops, not a marketing slogan', async ({ page }) => {
  const errors = collectPageErrors(page);
  // Top guilds (spec §2.4, "Around the site") reads GET /v1/rankings/guilds live; stub it
  // so the ready state is deterministic instead of depending on the real API answering
  // (the way rankings.spec.ts, rankings-phone.spec.ts and rankings-encounter-picker.spec.ts
  // already stub this same endpoint).
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
  // A reference heading, not a slogan: short, and it never ends with "!" or "?" -- a
  // trailing "." is fine (spec 2026-09-23 §2's own headline sentence, reference voice
  // "state the thing and stop").
  expect(heading.length).toBeLessThanOrEqual(90);
  expect(heading).not.toMatch(/[!?]$/);
  expect(heading).toBe('Your character, planned, simmed, logged and ranked.');
  await expect(page.getByTestId('home-timeline')).toBeVisible();
  await expect(page.getByTestId('home-next-planner-signed-out')).toBeVisible();
  await expect(page.getByTestId('home-next-simulator-signed-out')).toBeVisible();
  await expect(page.getByTestId('home-next-logs-signed-out')).toBeVisible();
  await expect(page.getByTestId('home-next-rankings-signed-out')).toBeVisible();
  // Spec §2.4 "Around the site": Recent reports and Top guilds, mounted nested inside
  // the old HomeProductPanel grid before this branch deleted that grid. RecentReports'
  // own wrapping section carries this testid in every load state (loading/failed/empty/
  // ready), so this alone proves the component is mounted at all -- the exact thing an
  // earlier task silently dropped. Top guilds' ready-state <ul> only renders once its
  // client:visible island hydrates and the stubbed fetch above resolves, so it needs a
  // scroll into view first (the same reason logs-recent-reports.spec.ts scrolls to
  // recent-reports before asserting on it).
  await expect(page.getByTestId('recent-reports')).toBeVisible();
  await page.getByTestId('home-around-the-site').scrollIntoViewIfNeeded();
  await expect(page.getByTestId('home-top-guilds')).toBeVisible();
  // The "Your guild" panel is gone (spec 2026-09-28): the hero's own guild card replaced
  // it, and that card is signed-in only, so a signed-out visitor sees no trace of it.
  await expect(page.getByTestId('home-guild-card')).toHaveCount(0);
  await expect(page.getByRole('heading', { name: 'What changed' })).toBeVisible();
  await expect(page.getByText('Not affiliated with or endorsed by Blizzard Entertainment')).toBeVisible();
  expect(errors).toEqual([]);
});

test('the four next-steps cards link to their tools, with the planner and simulator live elements gone from this page', async ({
  page,
}) => {
  await page.goto('/');
  await expect(
    page.getByTestId('home-next-planner-signed-out').getByRole('link', { name: 'Open the planner' }),
  ).toHaveAttribute('href', '/planner');
  await expect(
    page.getByTestId('home-next-simulator-signed-out').getByRole('link', { name: 'Open the simulator' }),
  ).toHaveAttribute('href', '/sim');
  await expect(
    page.getByTestId('home-next-logs-signed-out').getByRole('link', { name: 'Open the logs' }),
  ).toHaveAttribute('href', '/logs');
  await expect(
    page.getByTestId('home-next-rankings-signed-out').getByRole('link', { name: 'Open the rankings' }),
  ).toHaveAttribute('href', '/rankings');
  // The class-tile row and the spec-pill wall left the home page (spec 2026-09-24 §3).
  await expect(page.getByRole('link', { name: 'Warrior', exact: true })).toHaveCount(0);
});

// The reference tiles' and addon/companion row's own ordering and hrefs are
// handoffs-home.spec.ts's job; the assertions above already cover this page's presence.

test('content pages ship no client JavaScript', async ({ page }) => {
  for (const path of ['/about', '/premium']) {
    const scripts: string[] = [];
    page.on('request', (r) => {
      if (r.resourceType() === 'script') scripts.push(r.url());
    });
    await page.goto(path);
    expect(scripts).toEqual([]);
    page.removeAllListeners('request');
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
  await expect(page.getByRole('navigation', { name: 'Primary' })).toBeVisible();
  await expect(page.getByRole('navigation', { name: 'Footer' })).toBeVisible();
});

test('at desktop width, the hero timeline card stays inside its own column instead of the page', async ({
  page,
}, testInfo) => {
  // Fix round (night-site-ux, 2026-09-28): the timeline's grid-area:1/1 overlay wrapper had
  // no min-width:0, so a CSS grid item's default min-width (its own content, not the
  // track) forced the whole lg:col-span-5 column -- and the Nov 4 Launch date, Dec 9 First
  // raids date and "Updated" stamp inside it -- past the viewport and into SkyBand's
  // overflow-hidden, invisible and unreachable by any scroll. Only reproduces at the `lg`
  // breakpoint (1024px+) the two-column hero uses, so this is desktop-only; the mobile
  // project stacks the column full-width and never hit this.
  test.skip(testInfo.project.name !== 'desktop', 'lg two-column hero only');
  await page.goto('/');

  const viewport = page.viewportSize();
  expect(viewport, 'desktop project always sets a viewport').not.toBeNull();
  const viewportWidth = viewport!.width;

  const guildBlock = page.getByTestId('home-guild-block');
  const guildBlockBox = await guildBlock.boundingBox();
  expect(guildBlockBox?.x ?? 0, 'the hero column itself starts on-screen').toBeGreaterThanOrEqual(0);
  expect(
    guildBlockBox!.x + guildBlockBox!.width,
    'the hero column stays inside the viewport',
  ).toBeLessThanOrEqual(viewportWidth + 1);

  // The card that holds the timeline rows must be contained the same way -- not merely the
  // grid cell around it -- since the original bug had the grid cell measuring correctly
  // while its overlaid child still blew out past it.
  const timeline = page.getByTestId('home-timeline');
  const timelineBox = await timeline.boundingBox();
  expect(
    timelineBox!.x + timelineBox!.width,
    'the timeline card stays inside the viewport, not clipped by an ancestor',
  ).toBeLessThanOrEqual(viewportWidth + 1);

  // The last row and the updated stamp are the two elements the original bug hid entirely
  // (they sat past 1280px, inside SkyBand's overflow-hidden, on every viewport narrower
  // than ~1500px) -- both must at least be reachable by scrolling the card horizontally.
  const rows = page.getByTestId('home-timeline-row');
  await expect(rows.last()).toBeAttached();
  await rows.last().scrollIntoViewIfNeeded();
  await expect(rows.last()).toBeInViewport();
  const updated = page.getByTestId('home-timeline-updated');
  await updated.scrollIntoViewIfNeeded();
  await expect(updated).toBeInViewport();
});
