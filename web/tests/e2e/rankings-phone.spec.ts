// web/tests/e2e/rankings-phone.spec.ts
// The phone audit for /rankings/<slug> -- report-phone.spec.ts's counterpart for this page.
// That sweep never visits this route (a different shell, a different island), and a ranked
// table is exactly the kind of wide content that breaks a phone layout, so it gets its own
// sweep in the same idiom: nothing scrolls sideways, and every control -- including a row's
// own links -- clears 44px, with the bar's stack open so its fields are measured too.
import { apiUnreachable, expect, test } from './support/test';
import { barToggle, openFilters } from './support/rankings-filters';
import { gotoHydrated } from './support/hydrated';

// Reads this spec does not answer go unanswered, as the API out of reach (support/test.ts).
test.beforeEach(({ page }) => apiUnreachable(page, '**/v1/rankings/encounters*'));

test.use({ viewport: { width: 360, height: 800 } });
test.skip(() => test.info().project.name !== 'mobile', 'phone layout only');

/** The phone width this file sets, and the number both sweeps measure against. */
const PHONE_WIDTH = 360;
/** The smallest touch target a phone control may have. */
const TARGET_PX = 44;

const ROWS = {
  ok: true,
  data: {
    rows: [
      {
        rank: 1,
        player: {
          key: 'us/hardcore/elyra-duskvale',
          name: 'Elyra Duskvale',
          class: 'Priest',
          spec: 'Shadow',
        },
        guild: { name: 'The Last Watch', ruleset: 'hardcore', region: 'us' },
        value: 1840,
        size: 40,
        fought_at: '2026-12-09T22:10:00Z',
        duration_ms: 240000,
        talent_split: '0/31/20',
        build_id: 'k7x2qm4a',
        trinkets: [],
        buff_count: 11,
        report_id: 'fixture2abcd',
        fight_index: 3,
        state: 'ok',
      },
      {
        rank: 2,
        player: { key: 'eu/normal/other-raider', name: 'Other Raider', class: 'Mage', spec: 'Fire' },
        value: 1600,
        size: 40,
        fought_at: '2026-12-09T21:00:00Z',
        duration_ms: 250000,
        talent_split: '0/20/31',
        trinkets: [],
        buff_count: 8,
        report_id: 'otherreport1',
        fight_index: 2,
        state: 'removed',
      },
    ],
    total: 2,
    page: 1,
    per_page: 100,
    updated_at: '2026-12-09T22:15:00Z',
  },
  error: null,
  request_id: 'r',
};

const GUILD_ROWS = {
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
};

/** Character rows and guild rows: the two boards this page renders. */
const STATES = ['', '?board=guild'];

async function stubApi(page: import('@playwright/test').Page): Promise<void> {
  await page.route('**/v1/rankings?**', (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(ROWS) }),
  );
  await page.route('**/v1/rankings/guilds?**', (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(GUILD_ROWS) }),
  );
}

test('nothing scrolls sideways on either board', async ({ page }) => {
  await stubApi(page);
  for (const query of STATES) {
    await page.goto(`/rankings/warden-kelthas${query}`);
    await expect(page.getByTestId('rankings-filters')).toBeVisible();
    await expect(page.getByTestId('ranking-1').or(page.getByTestId('guild-rows'))).toBeVisible();
    const scrollWidth = await page.evaluate(() => document.documentElement.scrollWidth);
    expect(scrollWidth, query).toBeLessThanOrEqual(PHONE_WIDTH);
    // With the bar's stack open, too: it holds every select and the Filters fields inline.
    await openFilters(page);
    const openWidth = await page.evaluate(() => document.documentElement.scrollWidth);
    expect(openWidth, `${query} stack open`).toBeLessThanOrEqual(PHONE_WIDTH);
  }
});

interface Sweep {
  /** Each control under 44px, named and measured. */
  short: string[];
  /** Of those, the ones a known product gap explains (KNOWN_GAP below). */
  measured: number;
}

/** Opens the board for `query` with the bar's stack open and measures every visible control. */
async function sweep(page: import('@playwright/test').Page, query: string): Promise<Sweep> {
  // The snapshot draws the board's rows in the server render, so a visible row no longer means the
  // island is live: wait for it to hydrate before the bar's toggle is pressed.
  await gotoHydrated(page, `/rankings/warden-kelthas${query}`, 'rankings-filters');
  await expect(page.getByTestId('ranking-1').or(page.getByTestId('guild-rows'))).toBeVisible();
  // The bar collapses to a summary line on a phone: open its stack, so its controls are measured too.
  await openFilters(page);
  await expect(barToggle(page)).toHaveAttribute('aria-expanded', 'true');
  const result: Sweep = { short: [], measured: 0 };
  const controls = await page.locator('button:visible, select:visible, a[href]:visible, input:visible').all();
  for (const control of controls) {
    const box = await control.boundingBox();
    if (box === null) continue;
    const facts = await control.evaluate((element) => ({
      // The skip link is 1px until it takes focus, the one control that is that size on purpose.
      skip: getComputedStyle(element).clipPath !== 'none',
      // `astro dev`'s own toolbar lives in a shadow root; it is not the page and never ships.
      devTool: element.getRootNode() !== document,
      // `#rankings` is this page's own scope (Rankings.svelte's root element), not the shared
      // header and footer, which layout.spec.ts already measures.
      inPage: element.closest('#rankings') !== null,
      named: `<${element.tagName.toLowerCase()}${
        element.getAttribute('data-testid') === null
          ? ''
          : ` data-testid=${element.getAttribute('data-testid')}`
      }> ${(element.textContent ?? '').trim().slice(0, 40)}`,
    }));
    if (facts.skip || facts.devTool) continue;
    result.measured += 1;
    // The larger dimension everywhere, and height as well inside the page -- a 200x20 row link
    // clears the max() line and is still unusable with a thumb, the same reasoning
    // report-phone.spec.ts's own sweep pins for `#report`.
    const clears = Math.max(box.height, box.width) >= TARGET_PX && (!facts.inPage || box.height >= TARGET_PX);
    if (clears) continue;
    const where = `${query} ${facts.named} measured ${Math.round(box.width)}x${Math.round(box.height)}`;
    result.short.push(where);
  }
  return result;
}

test('every filter control, row link and pagination control clears 44px', async ({ page }) => {
  await stubApi(page);
  for (const query of STATES) {
    const { short, measured } = await sweep(page, query);
    expect(short, query).toEqual([]);
    // Every state renders the board switch's two radios plus at least the metric or kind
    // select, so the floor here is three -- a state whose controls all fell through the
    // exclusions above would leave this loop with nothing asserted and still green, which
    // is worse than no audit.
    expect(measured, `${query} measured no controls at all`).toBeGreaterThanOrEqual(3);
  }
});
