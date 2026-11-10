// web/tests/e2e/sim-scope.spec.ts
// task-2-brief.md: the healer and tank persona reviews both found nothing on the site
// states the simulator is DPS-only. The lane's acceptance criterion (constraints.md) is
// specific -- the scope sentence must be visible without scrolling at 390x844 on every
// simulator page -- so this proves that directly, on /sim and every tool page (a saved run is
// always a damage spec, so /sim/<id> carries no note), rather than only
// asserting the copy key is present somewhere in the DOM.
import { expect, test } from './support/test';
import { simCopy } from '../../src/lib/sim/copy';
import { simPageCopy } from '../../src/lib/sim/page-copy';

// 390x844, the exact width and height task-2-brief.md's own acceptance criterion names --
// the same literal viewport tests/e2e/sim-tabs.spec.ts's own phone-width block uses, rather
// than the --project=mobile device (Pixel 7, 412px) most of this suite's phone coverage
// runs under.
const PAGES = ['/sim', '/sim/gear', '/sim/drops', '/sim/talents', '/sim/weights', '/sim/specs'];

test.describe('the DPS-only scope sentence', () => {
  test.use({ viewport: { width: 390, height: 844 } });

  for (const route of PAGES) {
    const isSim = route === '/sim';
    test(
      isSim
        ? `${route} states the DPS-only caveat in its head, visible without scrolling at 390x844`
        : `${route} states the simulator is DPS-only, visible without scrolling at 390x844`,
      async ({ page }) => {
        await page.goto(route);

        // /sim says it in the landing head's one meta line (spec 2026-11-06-sim §4.2, §6:
        // "Damage specs only"), the tool pages in ScopeNote.astro at the top of their shells.
        const note = isSim ? page.getByTestId('sim-head') : page.getByTestId('sim-scope-note');
        if (isSim) await expect(note).toContainText(simPageCopy.landingMeta);
        else await expect(note).toHaveText(simCopy.scopeNote);
        // toBeInViewport() alone would pass for an element scrolled into view; the point is
        // that no scroll is needed, so this also asserts the page has not scrolled.
        await expect(note).toBeInViewport();
        const scrollY = await page.evaluate(() => window.scrollY);
        expect(scrollY, `${route} should not need scrolling to see the scope sentence`).toBe(0);
      },
    );
  }

  test('is the exact sentence task-2-brief.md specifies', () => {
    // h-audit 2026-11-09: cut to the fact; the healer and tank rows on /sim/specs say the rest.
    expect(simCopy.scopeNote).toBe('Damage specs only.');
  });

  // 2026-09-26 layout pass, Finding 3: a signed-out visitor never reads the restriction
  // twice. The head's meta line is the one copy on /sim now.
  test('/sim states the sentence once, not twice, for a signed-out visitor', async ({ page }) => {
    await page.goto('/sim');
    await expect(page.getByText('Damage specs only')).toHaveCount(1);
    await expect(page.getByTestId('sim-sources-scope-note')).toHaveCount(0);
    await expect(page.getByTestId('sim-scope-note')).toHaveCount(0);
  });
});
