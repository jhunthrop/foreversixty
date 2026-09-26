// web/tests/e2e/sim-scope.spec.ts
// task-2-brief.md: the healer and tank persona reviews both found nothing on the site
// states the simulator is DPS-only. The lane's acceptance criterion (constraints.md) is
// specific -- the scope sentence must be visible without scrolling at 390x844 on every
// simulator page -- so this proves that directly, on all seven pages, rather than only
// asserting the copy key is present somewhere in the DOM.
import { expect, test } from '@playwright/test';
import { simCopy } from '../../src/lib/sim/copy';
import { landingCopy } from '../../src/lib/sim/landing-copy';

// 390x844, the exact width and height task-2-brief.md's own acceptance criterion names --
// the same literal viewport tests/e2e/sim-tabs.spec.ts's own phone-width block uses, rather
// than the --project=mobile device (Pixel 7, 412px) most of this suite's phone coverage
// runs under.
const FIXTURE_SIM_ID = 'simfixtureab';

const PAGES = [
  '/sim',
  '/sim/gear',
  '/sim/drops',
  '/sim/talents',
  '/sim/weights',
  '/sim/specs',
  `/sim/${FIXTURE_SIM_ID}`,
];

test.describe('the DPS-only scope sentence', () => {
  test.use({ viewport: { width: 390, height: 844 } });

  for (const route of PAGES) {
    const isSim = route === '/sim';
    test(
      isSim
        ? `${route} states the merged DPS-only/below-60 caveat`
        : `${route} states the simulator is DPS-only, visible without scrolling at 390x844`,
      async ({ page }) => {
        await page.goto(route);

        const note = page.getByTestId('sim-scope-note');
        const expected = isSim ? landingCopy.scopeCaveat : simCopy.scopeNote;
        await expect(note).toHaveText(expected);
        if (isSim) {
          // Task 8 (spec 2026-09-25 section 6, Ruling 7): /sim's merged caveat moved into
          // SimView's hydrated output, under the character list/source switcher, rather
          // than the static shell's top where ScopeNote.astro used to sit -- deliberately,
          // "caveats move after the thing, never before it." For a signed-out visitor the
          // source switcher's four panels push the note well below the fold, so the
          // "visible without scrolling" guarantee this loop still enforces for the other
          // six pages (whose ScopeNote.astro is unchanged, at the top of their static
          // shells) no longer holds for /sim. That's this task's intended trade-off, not a
          // regression, so /sim is exempted from the viewport/scroll assertions below.
          return;
        }
        // toBeInViewport() alone would pass for an element scrolled into view; the point is
        // that no scroll is needed, so this also asserts the page has not scrolled.
        await expect(note).toBeInViewport();
        const scrollY = await page.evaluate(() => window.scrollY);
        expect(scrollY, `${route} should not need scrolling to see the scope sentence`).toBe(0);
      },
    );
  }

  test('is the exact sentence task-2-brief.md specifies', () => {
    expect(simCopy.scopeNote).toBe(
      'The simulator runs damage specs only. Healing and tanking specs are not simulated yet.',
    );
  });

  // The landing state and the source switcher are what a signed-out visitor reads first
  // (task-2-brief.md), so each carries the same sentence too -- not only the shell.
  test('/sim carries the sentence a second time in the source switcher, for a signed-out visitor', async ({
    page,
  }) => {
    await page.goto('/sim');
    await expect(page.getByTestId('sim-sources-scope-note')).toHaveText(simCopy.scopeNote);
  });
});
