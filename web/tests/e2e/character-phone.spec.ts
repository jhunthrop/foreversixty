// web/tests/e2e/character-phone.spec.ts
// The phone audit for /character/<region>/<ruleset>/<name> -- report-phone.spec.ts and
// rankings-phone.spec.ts's counterpart for this page. A character's every ranked fight and
// their best per encounter are exactly the kind of wide, row-heavy content that breaks a
// phone layout, so this page gets its own sweep in the same idiom: nothing scrolls
// sideways, and every control clears 44px in both the max-dimension and the height sense.
import { expect, test } from '@playwright/test';

test.use({ viewport: { width: 360, height: 800 } });
test.skip(() => test.info().project.name !== 'mobile', 'phone layout only');

/** The phone width this file sets, and the number both sweeps measure against. */
const PHONE_WIDTH = 360;

const CHARACTER = {
  ok: true,
  data: {
    character: { name: 'Elyra Duskvale', region: 'us', ruleset: 'hardcore', class: 'Priest' },
    best: [
      {
        encounter: 'Warden Kelthas',
        encounter_id: 9001,
        difficulty: 8,
        metric: 'hps',
        value: 1840,
        percentile: 96.2,
        spec: 'Discipline',
        fought_at: '2026-12-09T22:10:00Z',
        report_id: 'fixture2abcd',
        fight_index: 3,
      },
      {
        encounter: 'Deep Warden',
        encounter_id: 9002,
        difficulty: 8,
        metric: 'hps',
        value: 1420,
        percentile: 71,
        spec: 'Discipline',
        fought_at: '2026-12-08T20:00:00Z',
        report_id: 'otherreport1',
        fight_index: 1,
      },
    ],
    history: [
      {
        encounter: 'Warden Kelthas',
        encounter_id: 9001,
        difficulty: 8,
        metric: 'hps',
        value: 1840,
        percentile: 96.2,
        spec: 'Discipline',
        fought_at: '2026-12-09T22:10:00Z',
        report_id: 'fixture2abcd',
        fight_index: 3,
      },
      {
        encounter: 'Deep Warden',
        encounter_id: 9002,
        difficulty: 8,
        metric: 'hps',
        value: 1420,
        percentile: 71,
        spec: 'Discipline',
        fought_at: '2026-12-08T20:00:00Z',
        report_id: 'otherreport1',
        fight_index: 1,
      },
    ],
    builds_seen: [{ talent_split: '31/20/0', spec: 'Discipline', first_seen: '2026-12-09T22:10:00Z' }],
  },
  error: null,
  request_id: 'r',
};

/**
 * This page has no filter or tab state -- one ready state, unlike Rankings' board switch
 * or the report's modes and views. The loop is kept anyway so a state added later (a
 * failed or missing address, say) slots into the same sweep rather than needing a second
 * copy of it.
 */
const STATES = [''];

async function stubApi(page: import('@playwright/test').Page): Promise<void> {
  await page.route('**/v1/characters/us/hardcore/elyra-duskvale', (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(CHARACTER) }),
  );
}

test('nothing scrolls sideways on a character page', async ({ page }) => {
  await stubApi(page);
  for (const query of STATES) {
    await page.goto(`/character/us/hardcore/elyra-duskvale${query}`);
    await expect(page.getByTestId('character')).toBeVisible();
    const scrollWidth = await page.evaluate(() => document.documentElement.scrollWidth);
    expect(scrollWidth, query).toBeLessThanOrEqual(PHONE_WIDTH);
  }
});

test('every row link on a character page clears 44px', async ({ page }) => {
  await stubApi(page);
  for (const query of STATES) {
    await page.goto(`/character/us/hardcore/elyra-duskvale${query}`);
    await expect(page.getByTestId('character')).toBeVisible();
    let measured = 0;
    const controls = await page
      .locator('button:visible, select:visible, a[href]:visible, input:visible')
      .all();
    for (const control of controls) {
      const box = await control.boundingBox();
      if (box === null) continue;
      // The skip link is 1px until it takes focus, the one control that is that size on
      // purpose; nothing on this page sits inside a `<p>` paragraph the way report-phone's
      // sweep excludes.
      const skip = await control.evaluate((element) => getComputedStyle(element).clipPath !== 'none');
      if (skip) continue;
      const named = await control.evaluate(
        (element) =>
          `<${element.tagName.toLowerCase()}${
            element.getAttribute('data-testid') === null
              ? ''
              : ` data-testid=${element.getAttribute('data-testid')}`
          }> ${(element.textContent ?? '').trim().slice(0, 40)}`,
      );
      const where = `${query} ${named} measured ${Math.round(box.width)}x${Math.round(box.height)}`;
      expect(Math.max(box.height, box.width), where).toBeGreaterThanOrEqual(44);
      // Height as well as the larger dimension: a 200x20 row link clears the max() line
      // above and is still unusable with a thumb, the same reasoning report-phone.spec.ts
      // pins for `#report` and rankings-phone.spec.ts pins for `#rankings`. `#character` is
      // this page's own scope, not the shared header and footer layout.spec.ts measures.
      const inPage = await control.evaluate((element) => element.closest('#character') !== null);
      if (inPage) expect(box.height, where).toBeGreaterThanOrEqual(44);
      measured += 1;
    }
    // Best and history each render two row links, so the floor here is four -- a state
    // whose controls all fell through the exclusions above would leave this loop with
    // nothing asserted and still green, which is worse than no audit.
    expect(measured, `${query} measured no controls at all`).toBeGreaterThanOrEqual(4);
  }
});
