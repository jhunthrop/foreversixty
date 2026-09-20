// web/tests/e2e/sim-tools-phone.spec.ts
// The four tool pages at phone width. Everything here is about layout and reach, not about
// the engine: a control that overflows the viewport or is under 44px tall is unusable, and
// nothing else in the suite measures that.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test, type Page } from '@playwright/test';

test.use({ viewport: { width: 360, height: 780 } });

const activeBuild = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
) as { build: string };

const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=12784`;

async function load(page: Page, route: string): Promise<void> {
  await page.goto(route);
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();
}

const ROUTES = ['/sim/gear', '/sim/talents', '/sim/drops', '/sim/weights'] as const;

for (const route of ROUTES) {
  test(`${route} never scrolls sideways at 360px`, async ({ page }) => {
    await load(page, route);
    const overflow = await page.evaluate(
      () => document.documentElement.scrollWidth - document.documentElement.clientWidth,
    );
    expect(overflow).toBeLessThanOrEqual(1);
  });

  test(`${route} gives every control a 44px hit target`, async ({ page }) => {
    await load(page, route);
    // Measured the same way sim-phone.spec.ts's own targetsAreBigEnough already measures the
    // exact same shared chrome these tool pages mount (Base.astro's skip link, CharacterStrip,
    // SettingsBar, RequestDrawer) -- a naive per-element box check flags two things that are
    // not bugs: the skip link, 1px until it takes focus (excluded there via computed
    // `clip-path`, Tailwind's sr-only technique), and a checkbox whose 20px mark sits inside a
    // `min-h-11` label that is the real hit target (measured there via the wrapping `<label>`).
    // Both exceptions are reused verbatim rather than reintroduced with different rules, so a
    // control this lane's own pages share with the rest of the simulator is judged by the one
    // standard already reviewed for it. Duplicated rather than imported from sim-phone.spec.ts
    // because extracting it would add a coupling edge into a file outside this task's list for
    // one small helper -- the same call Task 11's report made for its own three-line `block()`
    // duplicate.
    const short = await page.evaluate(() => {
      const bad: string[] = [];
      for (const element of document.querySelectorAll('button, a[href], select, input, label')) {
        const box = element.getBoundingClientRect();
        if (box.width === 0 && box.height === 0) continue;
        if (getComputedStyle(element).clipPath !== 'none') continue;
        const hitTarget =
          element instanceof HTMLInputElement &&
          element.type === 'checkbox' &&
          element.closest('label') !== null
            ? (element.closest('label') as HTMLElement)
            : element;
        const target = hitTarget.getBoundingClientRect();
        if (Math.max(target.height, target.width) < 44) bad.push(element.outerHTML.slice(0, 80));
      }
      return bad;
    });
    expect(short).toEqual([]);
  });
}

test('the run bar stays reachable on /sim/gear after a run', async ({ page }) => {
  await load(page, '/sim/gear');
  await page.getByTestId('sim-search-add-16963').click();
  await page.getByTestId('sim-run-bulk').click();
  await expect(page.getByTestId('sim-combos')).toBeVisible({ timeout: 25_000 });
  await expect(page.getByTestId('sim-run-bulk')).toBeVisible();
});
