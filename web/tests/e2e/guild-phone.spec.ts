// web/tests/e2e/guild-phone.spec.ts
// The phone audit for /guild/<region>/<ruleset>/<name> -- report-phone.spec.ts and
// rankings-phone.spec.ts's counterpart for this page. A guild's progression table is
// exactly the kind of wide, row-heavy content that breaks a phone layout, so this page
// gets its own sweep in the same idiom: nothing scrolls sideways, and every control
// clears 44px in both the max-dimension and the height sense.
import { expect, test } from '@playwright/test';

test.use({ viewport: { width: 360, height: 800 } });
test.skip(() => test.info().project.name !== 'mobile', 'phone layout only');

/** The phone width this file sets, and the number both sweeps measure against. */
const PHONE_WIDTH = 360;

const GUILD = {
  ok: true,
  data: {
    guild: { name: 'The Last Watch', region: 'us', ruleset: 'hardcore' },
    progression: [
      {
        encounter: 'Warden Kelthas',
        encounter_id: 9001,
        difficulty: 8,
        kills: 2,
        pull_count: 14,
        first_kill_at: '2026-12-09T22:10:00Z',
      },
      { encounter: 'Deep Warden', encounter_id: 9002, difficulty: 8, kills: 0, pull_count: 31 },
    ],
    roster_best: [
      {
        player: {
          key: 'us/hardcore/elyra-duskvale',
          name: 'Elyra Duskvale',
          class: 'Priest',
          spec: 'Discipline',
        },
        encounter: 'Warden Kelthas',
        encounter_id: 9001,
        metric: 'hps',
        value: 1840,
        fought_at: '2026-12-09T22:10:00Z',
      },
      {
        player: { key: 'us/hardcore/baelgrim', name: 'Baelgrim', class: 'Warrior', spec: 'Protection' },
        encounter: 'Warden Kelthas',
        encounter_id: 9001,
        metric: 'damage_taken',
        value: 980,
        fought_at: '2026-12-09T22:10:00Z',
      },
    ],
    reports: [
      {
        id: 'fixture2abcd',
        title: 'Sanguine Depths, fixture night',
        zone: 'Sanguine Depths',
        created_at: '2026-09-26T20:09:00Z',
      },
      { id: 'otherreport1', title: '', zone: 'Blackmaw Hold', created_at: '2026-12-08T19:30:00Z' },
    ],
  },
  error: null,
  request_id: 'r',
};

/**
 * This page has no filter or tab state -- one ready state, unlike Rankings' board switch
 * or the report's modes and views. The loop is kept anyway so a state added later slots
 * into the same sweep rather than needing a second copy of it.
 */
const STATES = [''];

async function stubApi(page: import('@playwright/test').Page): Promise<void> {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(GUILD) }),
  );
}

test('nothing scrolls sideways on a guild page', async ({ page }) => {
  await stubApi(page);
  for (const query of STATES) {
    await page.goto(`/guild/us/hardcore/the-last-watch${query}`);
    await expect(page.getByTestId('guild')).toBeVisible();
    const scrollWidth = await page.evaluate(() => document.documentElement.scrollWidth);
    expect(scrollWidth, query).toBeLessThanOrEqual(PHONE_WIDTH);
  }
});

test('every row link on a guild page clears 44px', async ({ page }) => {
  await stubApi(page);
  for (const query of STATES) {
    await page.goto(`/guild/us/hardcore/the-last-watch${query}`);
    await expect(page.getByTestId('guild')).toBeVisible();
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
      // pins for `#report` and rankings-phone.spec.ts pins for `#rankings`. `#guild` is
      // this page's own scope, not the shared header and footer layout.spec.ts measures.
      const inPage = await control.evaluate((element) => element.closest('#guild') !== null);
      if (inPage) expect(box.height, where).toBeGreaterThanOrEqual(44);
      measured += 1;
    }
    // Progression renders two row links, roster bests one link each and reports one link
    // each, so the floor here is six -- a state whose controls all fell through the
    // exclusions above would leave this loop with nothing asserted and still green, which
    // is worse than no audit.
    expect(measured, `${query} measured no controls at all`).toBeGreaterThanOrEqual(6);
  }
});
