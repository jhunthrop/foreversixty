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
import { assertNoHorizontalScroll } from './support/phone-scroll';

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
  'view=timelines',
  'view=events',
  'view=queries',
  'mode=compare',
  'mode=rankings',
  'mode=mechanics',
];

test('nothing scrolls sideways on any tab or view', async ({ page }) => {
  for (const query of STATES) {
    await page.goto(`${REPORT}&${query}`);
    await expect(page.getByTestId('mode-bar')).toBeVisible();
    // Against PHONE_WIDTH, not window.innerWidth -- see support/phone-scroll.ts's header
    // note: content wider than the phone widens the emulated layout viewport with it, so
    // `scrollWidth > innerWidth` moves its own goal posts and silently passes over exactly
    // the bug this sweep exists to catch.
    await assertNoHorizontalScroll(page, PHONE_WIDTH, query);
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
  expect(box.width).toBeLessThanOrEqual(PHONE_WIDTH);
});

// The sideways-scroll sweep above visits `mode=compare` with no second fight picked, so
// it never sees the table -- and a table is the one thing on the page that grows past its
// container rather than wrapping inside it. Picked here, and measured against the gutter
// rather than the viewport: eating the 18px gutter does not lengthen the page's own
// scrollWidth, so the sweep would pass with the table flush to the screen edge.
test('the compare table stays inside the page gutter once a second fight is picked', async ({ page }) => {
  await page.goto(`${REPORT}&mode=compare`);
  await page.getByTestId('compare-with').selectOption('1');
  // A phone gets one card per player; the table is the desktop's.
  await expect(page.getByTestId('compare-cards')).toBeVisible();

  // Measured against the 360 this file sets, not window.innerWidth: anything wider than the
  // phone widens the emulated layout viewport with it, so innerWidth moves to meet the
  // overflow and a comparison against it always passes.
  expect(await page.evaluate(() => document.documentElement.scrollWidth)).toBeLessThanOrEqual(PHONE_WIDTH);

  // The cards keep the 18px page gutter on both sides, and each names its player beside
  // the difference: a four-column table in this width showed names and no numbers.
  const gutter = 18;
  const box = (await page.getByTestId('compare-cards').boundingBox())!;
  expect(box.x).toBeGreaterThanOrEqual(gutter);
  expect(box.x + box.width).toBeLessThanOrEqual(PHONE_WIDTH - gutter);
  await expect(page.getByTestId('compare-card-delta').first()).toContainText(/[+-]/);
});

test('the chart sits above the table, as the spec asks', async ({ page }) => {
  await page.goto(`${REPORT}&tab=damage-done`);
  const chart = (await page.getByTestId('time-chart').boundingBox())!;
  const table = (await page.getByTestId('actor-table').boundingBox())!;
  expect(chart.y).toBeLessThan(table.y);
});

// The strips once stuck to the top of a phone screen; at 263px they covered every table's
// headings, so they scroll with the page now, and the tab strip wraps rather than hiding
// "Deaths" past the right edge.
test('every tab is visible on a phone without scrolling sideways', async ({ page }) => {
  await page.goto(`${REPORT}&tab=damage-done`);
  const deaths = (await page.getByTestId('tab-deaths').boundingBox())!;
  expect(deaths.x + deaths.width).toBeLessThanOrEqual(390);
  expect(await page.evaluate(() => document.documentElement.scrollWidth)).toBeLessThanOrEqual(390);
});

test('every interactive control clears 44px', async ({ page }) => {
  for (const query of STATES) {
    await page.goto(`${REPORT}&${query}`);
    await expect(page.getByTestId('mode-bar')).toBeVisible();
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
  const row = page.getByTestId('threat-on-target').locator('li').first();
  await expect(row.getByTestId('threat-standing')).toContainText('standing');
  await expect(row.getByTestId('threat-built')).toContainText('built');
});

test('a resource row’s cap figures sit beside the line at 360px', async ({ page }) => {
  await page.goto(`${REPORT}&tab=resources`);
  const figures = page.getByTestId('resource-cap-figures').first();
  await expect(figures).toBeVisible();
  const box = await figures.boundingBox();
  expect(box?.width ?? 0).toBeLessThanOrEqual(PHONE_WIDTH);
});

test('the on-the-chart control is a 44px target on a phone, one tap from the name', async ({ page }) => {
  await page.goto(`${REPORT}&tab=damage-done`);
  await page.getByTestId('actor-Player-4184-000000A1').getByRole('button').first().click();
  const control = page.getByTestId('ability-chart').first();
  await expect(control).toBeVisible();
  const box = await control.boundingBox();
  expect(box?.height ?? 0).toBeGreaterThanOrEqual(44);
  // In the pinned ability column, not seven snapped swipes to the right of it: the
  // control sits inside the first cell, at rest, so a plain tap reaches it. A forced
  // click would pass wherever the button was, which is the whole point of not using one.
  const cell = page.getByTestId('row-abilities').locator('tbody td').first();
  const cellBox = await cell.boundingBox();
  expect(box?.x ?? 0).toBeGreaterThanOrEqual(cellBox?.x ?? 0);
  expect((box?.x ?? 0) + (box?.width ?? 0)).toBeLessThanOrEqual(
    (cellBox?.x ?? 0) + (cellBox?.width ?? 0) + 1,
  );
  await control.click();
});

test('a phase preset is a 44px target on a phone', async ({ page }) => {
  await page.goto(`${REPORT}`);
  const chip = page.getByTestId('window-presets').getByRole('button', { name: /Phase 2/ });
  await expect(chip).toBeVisible();
  const box = await chip.boundingBox();
  expect(box?.height ?? 0).toBeGreaterThanOrEqual(44);
});
