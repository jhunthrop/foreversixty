// web/tests/e2e/sim-scope.spec.ts
// task-2-brief.md: the healer and tank persona reviews both found nothing on the site
// states the simulator is DPS-only. The lane's acceptance criterion (constraints.md) is
// specific -- the scope sentence must be visible without scrolling at 390x844 on every
// simulator page -- so this proves that directly, on all seven pages, rather than only
// asserting the copy key is present somewhere in the DOM.
import { expect, test } from '@playwright/test';
import { simCopy } from '../../src/lib/sim/copy';

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
    test(`${route} states the simulator is DPS-only, visible without scrolling at 390x844`, async ({
      page,
    }) => {
      await page.goto(route);

      const note = page.getByTestId('sim-scope-note');
      await expect(note).toHaveText(simCopy.scopeNote);
      // toBeInViewport() alone would pass for an element scrolled into view; the point is
      // that no scroll is needed, so this also asserts the page has not scrolled.
      await expect(note).toBeInViewport();
      const scrollY = await page.evaluate(() => window.scrollY);
      expect(scrollY, `${route} should not need scrolling to see the scope sentence`).toBe(0);
    });
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
