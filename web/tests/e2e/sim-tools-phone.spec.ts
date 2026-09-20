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
    // Two exceptions are borrowed from sim-phone.spec.ts's own targetsAreBigEnough, which
    // already audits this exact shared chrome (Base.astro's skip link, CharacterStrip,
    // SettingsBar, RequestDrawer) mounted on /sim: a naive per-element box check flags two
    // things that are not bugs -- the skip link, 1px until it takes focus (excluded here via
    // computed `clip-path`, Tailwind's sr-only technique, same as there), and a checkbox
    // whose 20px mark sits inside a `min-h-11` label that is the real hit target (measured
    // here via the wrapping `<label>`, same as there).
    //
    // Deliberate divergence from sim-phone.spec.ts (controller ruling, Task 21 fix round 1):
    // that file reads `Math.max(height, width) < target`, which passes a control that is
    // wide but short -- a 44x20 row clears it. This file requires BOTH dimensions to reach
    // 44px, on the substituted label where the checkbox substitution applies, never on the
    // hidden input itself. A future reader should not "fix" this file to match that one; the
    // two measure different things on purpose, and this is the stricter of the two.
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
        if (target.height < 44 || target.width < 44) bad.push(element.outerHTML.slice(0, 80));
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
