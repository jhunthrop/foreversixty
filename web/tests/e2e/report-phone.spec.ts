// web/tests/e2e/report-phone.spec.ts
// The phone audit. Every other report spec asserts one flow; this one asserts the one
// property that holds across all of them -- the page fits a phone and can be operated with
// a finger -- across every tab, view and mode, so a panel added later cannot quietly
// reintroduce a sideways scroll or a 20px button.
//
// 360x800, not the Pixel 7's own 412px: 360 is the narrowest phone design/DESIGN-SYSTEM.md
// covers, and it is the width layout.spec.ts, gear.spec.ts, planner-phone.spec.ts and
// report-brush.spec.ts all measure at. The project skip keeps the run to one device rather
// than one per project, the same pairing report-brush.spec.ts uses.
import { expect, test } from '@playwright/test';

test.use({ viewport: { width: 360, height: 800 } });

// `test.skip(condition, …)` hands the callback the fixtures, not the TestInfo, so the
// project is read off `test.info()` rather than a second parameter.
test.skip(() => test.info().project.name !== 'mobile', 'phone layout only');

const REPORT = '/reports/fixture2abcd?fight=3';

test('nothing scrolls sideways on any tab or view', async ({ page }) => {
  for (const query of [
    'tab=summary',
    'tab=damage-done',
    'tab=deaths',
    'tab=buffs',
    'tab=casts',
    'tab=resources',
    'tab=threat',
    'view=timelines',
    'view=events',
    'view=queries',
    'mode=compare',
    'mode=rankings',
  ]) {
    await page.goto(`${REPORT}&${query}`);
    await expect(page.getByTestId('mode-bar')).toBeVisible();
    const overflows = await page.evaluate(() => document.documentElement.scrollWidth > window.innerWidth);
    expect(overflows, query).toBe(false);
  }
});

test('a damage row is a card with every field labelled', async ({ page }) => {
  await page.goto(`${REPORT}&tab=damage-done`);
  const row = page.getByTestId('actor-Player-4184-000000A1');
  await expect(row.getByTestId('row-name')).toBeVisible();
  await expect(row.getByTestId('row-amount')).toBeVisible();
  await expect(row.getByTestId('row-per-second')).toBeVisible();
  await expect(row.getByTestId('phone-labels')).toBeVisible();

  const box = (await row.boundingBox())!;
  expect(box.width).toBeLessThanOrEqual(360);
});

// The sideways-scroll sweep above visits `mode=compare` with no second fight picked, so
// it never sees the table -- and a table is the one thing on the page that grows past its
// container rather than wrapping inside it. Picked here, and measured against the gutter
// rather than the viewport: eating the 18px gutter does not lengthen the page's own
// scrollWidth, so the sweep would pass with the table flush to the screen edge.
test('the compare table stays inside the page gutter once a second fight is picked', async ({ page }) => {
  await page.goto(`${REPORT}&mode=compare`);
  await page.getByTestId('compare-with').selectOption('1');
  await expect(page.getByTestId('compare-table')).toBeVisible();

  // Measured against the 360 this file sets, not window.innerWidth: a table wider than the
  // phone widens the emulated layout viewport with it, so innerWidth moves to meet the
  // overflow and a comparison against it always passes.
  expect(await page.evaluate(() => document.documentElement.scrollWidth)).toBeLessThanOrEqual(360);

  // Wider than the phone is allowed for a table, but only inside its own scroller and only
  // with the 18px page gutter still standing on both sides.
  const gutter = 18;
  const scroller = page.getByTestId('compare-table').locator('xpath=parent::div');
  const box = (await scroller.boundingBox())!;
  expect(box.x).toBeGreaterThanOrEqual(gutter);
  expect(box.x + box.width).toBeLessThanOrEqual(360 - gutter);

  // Still a real table, with the caption and the column headers a card stack would lose:
  // the header association is the whole reason Compare does not collapse to cards.
  await expect(page.getByTestId('compare-table').locator('th[scope="col"]')).toHaveCount(4);
  await expect(page.getByTestId('compare-table').locator('caption')).toHaveCount(1);
});

test('the chart sits above the table, as the spec asks', async ({ page }) => {
  await page.goto(`${REPORT}&tab=damage-done`);
  const chart = (await page.getByTestId('time-chart').boundingBox())!;
  const table = (await page.getByTestId('actor-table').boundingBox())!;
  expect(chart.y).toBeLessThan(table.y);
});

test('the mode bar stays reachable while a long table scrolls', async ({ page }) => {
  await page.goto(`${REPORT}&tab=damage-done`);
  await page.mouse.wheel(0, 1200);
  await expect(page.getByTestId('mode-bar')).toBeInViewport();
});

test('every interactive control clears 44px', async ({ page }) => {
  for (const query of ['tab=summary', 'tab=damage-done', 'tab=deaths', 'view=queries']) {
    await page.goto(`${REPORT}&${query}`);
    const controls = await page
      .locator('button:visible, select:visible, a[href]:visible, input:visible')
      .all();
    for (const control of controls) {
      // What counts as the target depends on how the control is worked. A range input is
      // dragged, so only its own box is the handle and h-11 goes on the input -- Task 10
      // settled that and report-brush.spec.ts pins it. A checkbox is toggled, and a click
      // anywhere in its wrapping label toggles it, so the label is the target and the box
      // inside it is only the mark. Everything else is measured as itself.
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
      // Inline links inside a sentence are text, not targets; everything else is a target.
      // The skip link is the one control that is 1px on purpose -- it is clipped until it
      // takes focus, at which point Base.astro's `focus:h-11` gives it the same 44px as
      // everything else. `clip-path` is what Tailwind's `sr-only` sets and nothing else
      // in this codebase uses, so it identifies that case without naming the element.
      const skip = await target.evaluate(
        (element) => element.closest('p') !== null || getComputedStyle(element).clipPath !== 'none',
      );
      if (skip) continue;
      // Named as precisely as the markup allows, so a failure says which control is short
      // rather than only that one is.
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
      // Height as well as the larger dimension, for the island this page is here to audit:
      // a 200x20 control clears the line above and is still unusable with a thumb, and
      // height is the dimension layout.spec.ts, planner-phone.spec.ts and
      // report-brush.spec.ts each pin. Scoped to #report because the shared header and
      // footer are every page's, not this one's -- layout.spec.ts measures those.
      const inIsland = await target.evaluate((element) => element.closest('#report') !== null);
      if (inIsland) expect(box.height, where).toBeGreaterThanOrEqual(44);
    }
  }
});
