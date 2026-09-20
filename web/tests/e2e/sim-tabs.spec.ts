// web/tests/e2e/sim-tabs.spec.ts
// task-1-brief.md: the one tab strip every simulator page carries. The newcomer and
// raid-leader persona reviews both found /sim/gear, /sim/drops, /sim/talents and
// /sim/weights "from nowhere" -- this proves the strip is not merely present on /sim but
// genuinely reachable from every page, says where you are, keeps a loaded character across
// a click to another tool, and fits a phone.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test } from '@playwright/test';
import { SIM_TABS } from '../../src/lib/sim/tabs';
import { assertNoHorizontalScroll } from './support/phone-scroll';

// Not imported from lib/report/shell-paths.ts: that module reaches lib/rankings/api.ts,
// which reads `import.meta.env.PUBLIC_API_BASE_URL` (lib/planner/config.ts) -- fine for
// every in-app import (Vite supplies it), but Playwright's own Node runtime has no
// `import.meta.env` and throws on load, the same class of import sim-specs.spec.ts's own
// header comment avoids for `dpsSpecs`. `sim-phone.spec.ts`'s saved-sim case hardcodes the
// same literal for the same reason.
const FIXTURE_SIM_ID = 'simfixtureab';

const ROOT = path.join(import.meta.dirname, '..', '..');
const meta = JSON.parse(
  readFileSync(path.join(ROOT, 'src', 'fixtures', 'report', 'meta.json'), 'utf8'),
) as Record<string, unknown>;

// Fight 3 (Warden Kelthas), the same ref sim-compare.spec.ts and sim-phone.spec.ts's own
// compare-mode case load: the one fixture fight whose COMBATANT_INFO row carries gear and
// talents, so a plain (non-compare) `?source=fight&ref=` bootstrap resolves to a real
// character rather than `simCopy.fightNoCombatant`.
const FIGHT_REF = 'fixture2abcd:3';

async function stubFightMeta(page: import('@playwright/test').Page): Promise<void> {
  await page.route('**/v1/reports/fixture2abcd', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({ ok: true, data: meta, error: null, request_id: 'r' }),
    }),
  );
}

const PAGES = [
  { path: '/sim', current: 'quick-sim' },
  { path: '/sim/gear', current: 'gear' },
  { path: '/sim/drops', current: 'drops' },
  { path: '/sim/talents', current: 'talents' },
  { path: '/sim/weights', current: 'weights' },
  { path: '/sim/specs', current: 'specs' },
  // A saved-sim permalink (task-1-brief.md: "give it the strip too, with Quick Sim
  // current"), not one of the six tools -- Quick Sim is the closest thing it has to a home.
  { path: `/sim/${FIXTURE_SIM_ID}`, current: 'quick-sim' },
] as const;

test.describe('the simulator tab strip', () => {
  for (const { path: route, current } of PAGES) {
    test(`${route} carries all six tabs in order, bare, with "${current}" current`, async ({ page }) => {
      await page.goto(route);
      await expect(page.getByTestId('sim-nav-tabs')).toBeVisible();

      const anchors = page.getByTestId('sim-nav-tabs').locator('a');
      await expect(anchors).toHaveCount(SIM_TABS.length);

      for (const [index, tab] of SIM_TABS.entries()) {
        const anchor = page.getByTestId(`sim-nav-tab-${tab.id}`);
        await expect(anchor).toBeVisible();
        await expect(anchor).toHaveText(tab.label);
        // No character loaded on a bare `goto`: every href is exactly its bare path, the
        // same string SimTabs.astro renders server-side with no rewrite from either island.
        await expect(anchor).toHaveAttribute('href', tab.href);
        await expect(anchors.nth(index)).toHaveAttribute('data-testid', `sim-nav-tab-${tab.id}`);

        if (tab.id === current) await expect(anchor).toHaveAttribute('aria-current', 'page');
        else await expect(anchor).not.toHaveAttribute('aria-current', 'page');
      }
    });
  }

  test('loading a character on /sim carries its query onto the Top Gear tab’s href', async ({ page }) => {
    await stubFightMeta(page);
    await page.goto(`/sim?source=fight&ref=${encodeURIComponent(FIGHT_REF)}`);
    await expect(page.getByTestId('sim-character')).toBeVisible();

    const expectedQuery = `?source=fight&ref=${encodeURIComponent(FIGHT_REF)}`;
    await expect(page.getByTestId('sim-nav-tab-gear')).toHaveAttribute('href', `/sim/gear${expectedQuery}`);
    // Every other tab is rewritten from the same effect, not just the one this test names.
    await expect(page.getByTestId('sim-nav-tab-drops')).toHaveAttribute('href', `/sim/drops${expectedQuery}`);
    await expect(page.getByTestId('sim-nav-tab-quick-sim')).toHaveAttribute('href', `/sim${expectedQuery}`);
  });

  test('the same rewrite runs on the tools island (/sim/gear), not only /sim', async ({ page }) => {
    await stubFightMeta(page);
    await page.goto(`/sim/gear?source=fight&ref=${encodeURIComponent(FIGHT_REF)}`);
    await expect(page.getByTestId('sim-character')).toBeVisible();

    const expectedQuery = `?source=fight&ref=${encodeURIComponent(FIGHT_REF)}`;
    await expect(page.getByTestId('sim-nav-tab-talents')).toHaveAttribute(
      'href',
      `/sim/talents${expectedQuery}`,
    );
  });
});

test.describe('the simulator tab strip at phone width', () => {
  // 390x844, the width task-1-brief.md's own acceptance criterion names.
  test.use({ viewport: { width: 390, height: 844 } });

  test('scrolls sideways within itself rather than the page, and every tab clears 44px', async ({ page }) => {
    await page.goto('/sim');
    await assertNoHorizontalScroll(page, 390);

    for (const tab of SIM_TABS) {
      const box = await page.getByTestId(`sim-nav-tab-${tab.id}`).boundingBox();
      expect(box?.height ?? 0, `${tab.id} height`).toBeGreaterThanOrEqual(44);
    }

    const strip = await page.evaluate(() => {
      const nav = document.querySelector('[data-testid="sim-nav-tabs"]');
      return { scrollWidth: nav?.scrollWidth ?? 0, clientWidth: nav?.clientWidth ?? 0 };
    });
    expect(strip.scrollWidth, 'tab strip does not scroll').toBeGreaterThan(strip.clientWidth);

    // The strip's own overflow never becomes the page's: re-checked after reading the
    // strip's scrollWidth above, the same double-check sim-phone.spec.ts's own result-tab
    // sweep performs for the identical reason (its header comment: Chromium's mobile
    // shrink-to-fit quirk can hide exactly this class of overflow from a single check).
    await assertNoHorizontalScroll(page, 390);
  });
});
