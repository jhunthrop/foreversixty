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

// The addon-paste path (sim-sources.spec.ts's own fixture): an `addon`-kind character,
// whose `source.ref` is always '' (sources.ts's `fromAddonExport`) -- fix round 1, Finding
// A's own case.
const activeBuild = JSON.parse(readFileSync(path.join(ROOT, 'src', 'data', 'active-build.json'), 'utf8')) as {
  build: string;
};
const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=11726`;

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

  // Fix round 1, Finding A: an addon paste is "the primary path into the product" (the
  // newcomer persona review) and stamps `source.ref: ''` -- before this fix every tab's
  // href read `?source=addon`, a query that looks like it restores the character and
  // silently does not. Quick Sim (SimView.svelte's own store reads `?code=`) gets a working
  // fallback link; Top Gear (ToolsView.svelte / bulk-store.svelte.ts, which has never read
  // `?code=`) gets a bare href instead -- never a `?code=` it cannot itself honour, which
  // would just move the inert-query problem rather than fix it (`SIM_TABS`' own
  // `supportsCode`).
  test('an addon-pasted character gets a working ?code= href only where the destination can use it', async ({
    page,
  }) => {
    await page.goto('/sim');
    await page.getByTestId('sim-addon-input').fill(FURY);
    await page.getByTestId('sim-addon-load').click();
    await expect(page.getByTestId('sim-character')).toBeVisible();

    // The fallback code needs the store's own talent index, which resolves slightly after
    // `character` itself does (bulk-store.svelte.ts's own `adopt()`, same ordering
    // store.svelte.ts uses) -- `toHaveAttribute` polls for it rather than reading the
    // attribute once, racing that resolve.
    const quickSim = page.getByTestId('sim-nav-tab-quick-sim');
    await expect(quickSim).toHaveAttribute('href', /^\/sim\?code=/);
    const quickSimHref = await quickSim.getAttribute('href');
    const query = new URL(quickSimHref ?? '', page.url()).searchParams;
    expect(query.get('source'), 'no source= for a ref-less character').toBeNull();
    expect(query.get('ref'), 'no ref= for a ref-less character').toBeNull();

    // Top Gear cannot bootstrap from ?code= (ToolsView.svelte has never read it), so it
    // stays bare rather than carrying a query it cannot itself honour.
    await expect(page.getByTestId('sim-nav-tab-gear')).toHaveAttribute('href', '/sim/gear');

    // Proves the Quick Sim link, once followed, genuinely restores the character -- not
    // merely that it carries a code= parameter.
    await page.goto(quickSimHref ?? '/sim');
    await expect(page.getByTestId('sim-character')).toBeVisible();
    await expect(page.getByTestId('sim-character-descriptor')).toContainText('Fury Warrior');
  });

  // The tools island's own `characterCode` fallback (bulk-store.svelte.ts): from /sim/gear,
  // the four tools tabs stay bare (none of them read ?code=) while the two SimView-served
  // tabs on the same strip still get the working fallback link.
  test('from the tools island, only the SimView-served tabs get the ?code= fallback', async ({ page }) => {
    await page.goto('/sim/gear');
    await page.getByTestId('sim-addon-input').fill(FURY);
    await page.getByTestId('sim-addon-load').click();
    await expect(page.getByTestId('sim-character')).toBeVisible();

    const specs = page.getByTestId('sim-nav-tab-specs');
    await expect(specs).toHaveAttribute('href', /^\/sim\/specs\?code=/);
    const specsHref = await specs.getAttribute('href');
    expect(new URL(specsHref ?? '', page.url()).searchParams.get('source')).toBeNull();

    for (const id of ['gear', 'drops', 'talents', 'weights']) {
      await expect(
        page.getByTestId(`sim-nav-tab-${id}`),
        `${id} cannot bootstrap from ?code=`,
      ).toHaveAttribute('href', `/sim/${id}`);
    }
  });
});

test.describe('the simulator tab strip at phone width', () => {
  // 390x844, the width task-1-brief.md's own acceptance criterion names.
  test.use({ viewport: { width: 390, height: 844 } });

  test('scrolls sideways within itself rather than the page, and every tab clears 44x44', async ({
    page,
  }) => {
    await page.goto('/sim');
    await assertNoHorizontalScroll(page, 390);

    for (const tab of SIM_TABS) {
      const box = await page.getByTestId(`sim-nav-tab-${tab.id}`).boundingBox();
      // design/DESIGN-SYSTEM.md's "44px minimum hit targets" is both dimensions, not just
      // height (fix round 1, Finding B) -- a tab this short in either one is not a real
      // 44x44 target even where the other clears it.
      expect(box?.height ?? 0, `${tab.id} height`).toBeGreaterThanOrEqual(44);
      expect(box?.width ?? 0, `${tab.id} width`).toBeGreaterThanOrEqual(44);
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
