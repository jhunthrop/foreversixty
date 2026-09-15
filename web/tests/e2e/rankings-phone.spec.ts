// web/tests/e2e/rankings-phone.spec.ts
// The phone audit for /rankings/<slug> -- report-phone.spec.ts's counterpart for this page.
// That sweep never visits this route (a different shell, a different island), and a ranked
// table is exactly the kind of wide content that breaks a phone layout, so it gets its own
// sweep in the same idiom: nothing scrolls sideways, and every control -- including a row's
// own links, which the report's rankings mode does not carry -- clears 44px.
import { expect, test } from '@playwright/test';

test.use({ viewport: { width: 360, height: 800 } });
test.skip(() => test.info().project.name !== 'mobile', 'phone layout only');

/** The phone width this file sets, and the number both sweeps measure against. */
const PHONE_WIDTH = 360;

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
    const scrollWidth = await page.evaluate(() => document.documentElement.scrollWidth);
    expect(scrollWidth, query).toBeLessThanOrEqual(PHONE_WIDTH);
  }
});

test('every filter control, row link and pagination control clears 44px', async ({ page }) => {
  await stubApi(page);
  for (const query of STATES) {
    await page.goto(`/rankings/warden-kelthas${query}`);
    await expect(page.getByTestId('rankings-filters')).toBeVisible();
    let measured = 0;
    const controls = await page
      .locator('button:visible, select:visible, a[href]:visible, input:visible')
      .all();
    for (const control of controls) {
      // A checkbox is toggled through the label wrapping it (`Today only`), so the label
      // is the target and the box inside it is only the mark -- report-phone.spec.ts's
      // own sweep settles this the same way.
      const target = (await control.evaluate(
        (element) =>
          element instanceof HTMLInputElement &&
          element.type === 'checkbox' &&
          element.closest('label') !== null,
      ))
        ? control.locator('xpath=ancestor::label[1]')
        : control;
      const box = await target.boundingBox();
      if (box === null) continue;
      // The skip link is 1px until it takes focus, the one control that is that size on
      // purpose; nothing else on this page is inside a `<p>`, so the exclusion is narrower
      // than report-phone.spec.ts's -- this page has no prose paragraph carrying a link.
      const skip = await target.evaluate((element) => getComputedStyle(element).clipPath !== 'none');
      if (skip) continue;
      const named = await target.evaluate(
        (element) =>
          `<${element.tagName.toLowerCase()}${
            element.getAttribute('data-testid') === null
              ? ''
              : ` data-testid=${element.getAttribute('data-testid')}`
          }> ${(element.textContent ?? '').trim().slice(0, 40)}`,
      );
      const where = `${query} ${named} measured ${Math.round(box.width)}x${Math.round(box.height)}`;
      expect(Math.max(box.height, box.width), where).toBeGreaterThanOrEqual(44);
      // Height as well as the larger dimension, for the page this audit covers -- a
      // 200x20 row link clears the max() line above and is still unusable with a thumb,
      // the same reasoning report-phone.spec.ts's own sweep pins for `#report`.
      // `#rankings` is this page's equivalent scope (Rankings.svelte's own root element),
      // not the shared header and footer, which layout.spec.ts already measures.
      const inPage = await target.evaluate((element) => element.closest('#rankings') !== null);
      if (inPage) expect(box.height, where).toBeGreaterThanOrEqual(44);
      measured += 1;
    }
    // Every state renders the board switch's two tabs plus at least the metric or kind
    // filter, so the floor here is three -- a state whose controls all fell through the
    // exclusions above would leave this loop with nothing asserted and still green, which
    // is worse than no audit.
    expect(measured, `${query} measured no controls at all`).toBeGreaterThanOrEqual(3);
  }
});
