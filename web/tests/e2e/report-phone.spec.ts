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
import { apiUnreachable, expect, test } from './support/test';
import { assertNoHorizontalScroll } from './support/phone-scroll';
import { openChartIfCollapsed } from './support/report-chart';
import { choose } from './support/select';

// Reads this spec does not answer go unanswered, as the API out of reach (support/test.ts).
test.beforeEach(({ page }) => apiUnreachable(page, '**/v1/rankings?*'));

test.use({ viewport: { width: 360, height: 800 } });

// `test.skip(condition, …)` hands the callback the fixtures, not the TestInfo, so the
// project is read off `test.info()` rather than a second parameter.
test.skip(() => test.info().project.name !== 'mobile', 'phone layout only');

const REPORT = '/reports/fixture2abcd?fight=3';

/** The phone width this file sets, and the number both sweeps measure against. */
const PHONE_WIDTH = 360;

/**
 * Every state the report can be in, swept by both audits below. One list, so a panel added
 * to a new tab cannot be covered by the scroll sweep and missed by the hit-target sweep.
 */
const STATES = [
  'tab=summary',
  'tab=damage-done',
  'tab=deaths',
  'tab=buffs',
  'tab=casts',
  'tab=resources',
  'tab=threat',
  'tab=timeline',
  'tab=events',
  'tab=queries',
  'tab=compare',
  'tab=mechanics',
];

test('nothing scrolls sideways on any tab or view', async ({ page }) => {
  for (const query of STATES) {
    await page.goto(`${REPORT}&${query}`);
    await expect(page.getByTestId('report-bar')).toBeVisible();
    // Against PHONE_WIDTH, not window.innerWidth -- see support/phone-scroll.ts's header
    // note: content wider than the phone widens the emulated layout viewport with it, so
    // `scrollWidth > innerWidth` moves its own goal posts and silently passes over exactly
    // the bug this sweep exists to catch.
    await assertNoHorizontalScroll(page, PHONE_WIDTH, query);
  }
});

test('a damage row keeps its name, its amount and the rate under it', async ({ page }) => {
  await page.goto(`${REPORT}&tab=damage-done`);
  const row = page.getByTestId('actor-Player-4184-000000A1');
  await expect(row.getByTestId('row-name')).toBeVisible();
  // Below 1024 a StatRow keeps its primary figure with the rate under it (logs spec 12).
  await expect(row.getByRole('cell').nth(1)).toBeVisible();
  await expect(row.getByRole('cell').nth(1)).toContainText('DPS');

  const box = (await row.boundingBox())!;
  expect(box.width).toBeLessThanOrEqual(PHONE_WIDTH);
});

// The sideways-scroll sweep above visits `tab=compare` with no second fight picked, so
// it never sees the table -- and a table is the one thing on the page that grows past its
// container rather than wrapping inside it. Picked here, and measured against the gutter
// rather than the viewport: eating the 18px gutter does not lengthen the page's own
// scrollWidth, so the sweep would pass with the table flush to the screen edge.
test('the compare table stays inside the page gutter once a second fight is picked', async ({ page }) => {
  await page.goto(`${REPORT}&tab=compare`);
  await choose(page.getByTestId('compare-with'), '1');
  // One table at every width; a phone folds the other figures under the name.
  await expect(page.getByTestId('compare-table')).toBeVisible();

  // Measured against the 360 this file sets, not window.innerWidth: anything wider than the
  // phone widens the emulated layout viewport with it, so innerWidth moves to meet the
  // overflow and a comparison against it always passes.
  expect(await page.evaluate(() => document.documentElement.scrollWidth)).toBeLessThanOrEqual(PHONE_WIDTH);

  // The cards keep the 18px page gutter on both sides, and each names its player beside
  // the difference: a four-column table in this width showed names and no numbers.
  const gutter = 18;
  const box = (await page.getByTestId('compare-table').boundingBox())!;
  expect(box.x).toBeGreaterThanOrEqual(gutter);
  expect(box.x + box.width).toBeLessThanOrEqual(PHONE_WIDTH - gutter);
  await expect(page.getByTestId('compare-Player-4184-000000A1')).toContainText(/[+\u2212]/);
});

test('the table leads and the chart follows it, as the spec asks (decision R4)', async ({ page }) => {
  await page.goto(`${REPORT}&tab=damage-done`);
  await openChartIfCollapsed(page);
  const chart = (await page.getByTestId('time-chart').boundingBox())!;
  const table = (await page.getByTestId('actor-table').boundingBox())!;
  expect(table.y).toBeLessThan(chart.y);
});

// The strips once stuck to the top of a phone screen; at 263px they covered every table's
// headings, so they scroll with the page now. Spec 2026-09-25 section 6 replaced the
// thirteen-pill strip below lg with a single select, which cannot itself wrap past the
// right edge the way the pill row once could -- so this now guards the select the same way.
test('the tab strip is one line that scrolls inside itself, never the page', async ({ page }) => {
  await page.goto(`${REPORT}&tab=damage-done`);
  const strip = (await page.getByTestId('report-tabs').boundingBox())!;
  expect(strip.height).toBeLessThanOrEqual(48);
  expect(await page.evaluate(() => document.documentElement.scrollWidth)).toBeLessThanOrEqual(PHONE_WIDTH);
});

test('every interactive control clears 44px', async ({ page }) => {
  // One round trip per control, for every control, in every state: twenty seconds on a
  // laptop and past the 45-second default on a two-core CI runner, where it timed out in
  // the middle of measuring rather than on any control that was too small.
  test.slow();
  for (const query of STATES) {
    await page.goto(`${REPORT}&${query}`);
    await expect(page.getByTestId('report-bar')).toBeVisible();
    let measured = 0;
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
      measured += 1;
    }
    // A state whose controls all fell through the exclusions above would leave this loop
    // with nothing asserted and the test still green, which is worse than no audit. Every
    // state renders the mode bar, so the floor is its three mode buttons.
    expect(measured, `${query} measured no controls at all`).toBeGreaterThanOrEqual(3);
  }
});

test('a windowed threat row says which figure is standing and which is built', async ({ page }) => {
  await page.goto(
    `${REPORT}&tab=threat&start=3000&end=14000&target=Creature-0-2085-2284-7855-169754-0000AA0002`,
  );
  const row = page.getByTestId('threat-on-target').locator('[role="row"][data-testid^="threat-"]').first();
  // The primary figure's phone line names it and carries the built figure.
  await expect(row.getByTestId('threat-standing')).toContainText('standing');
  await expect(row.getByTestId('threat-standing')).toContainText('built');
});

test('a resource row’s cap figures sit beside the line at 360px', async ({ page }) => {
  await page.goto(`${REPORT}&tab=resources`);
  const figures = page.getByTestId('resource-cap-figures').first();
  await expect(figures).toBeVisible();
  const box = await figures.boundingBox();
  expect(box?.width ?? 0).toBeLessThanOrEqual(PHONE_WIDTH);
});

test('the chart folds to a 44px line on a phone, and the Window select is a 44px field', async ({ page }) => {
  await page.goto(`${REPORT}`);
  const fold = page.getByTestId('chart-collapsed').locator('summary');
  await expect(fold).toBeVisible();
  expect((await fold.boundingBox())?.height ?? 0).toBeGreaterThanOrEqual(44);
  await page.getByTestId('report-bar').locator('[data-control-bar-toggle]').click();
  const select = page.getByTestId('window-select');
  await expect(select).toContainText('Phase 2');
  const box = await page.locator('.sel:has(> select[data-testid="window-select"])').boundingBox();
  expect(box?.height ?? 0).toBeGreaterThanOrEqual(44);
});
