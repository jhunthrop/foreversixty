// web/tests/e2e/sim-phone.spec.ts
// The phone audit for the simulator: report-phone.spec.ts, rankings-phone.spec.ts,
// character-phone.spec.ts and guild-phone.spec.ts's counterpart for every /sim route. Every
// component in this lane was written with its phone layout in mind; this file proves it,
// across every state a player actually reaches -- empty, a loaded character, a finished run
// on every result tab, compare mode, the spec-support grid and a saved sim -- rather than
// trusting that no later change quietly reopened one.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test, type Page } from '@playwright/test';

test.describe.configure({ mode: 'serial' });
// No test.use({ viewport }): the phone audit runs under --project=mobile (Pixel 7, 412px),
// which is the width the Global Constraints name, and overriding it here would audit a
// size nothing else in the suite uses.

const GUTTER = 18;
const TARGET = 44;

const ROOT = path.join(import.meta.dirname, '..', '..');
const activeBuild = JSON.parse(readFileSync(path.join(ROOT, 'src', 'data', 'active-build.json'), 'utf8')) as {
  build: string;
};
// Same addon export sim-run.spec.ts, sim-settings.spec.ts and sim-sources.spec.ts load: it
// needs no API stub, so it is the fastest way onto a page state with a character on it.
const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=11726`;

const meta = JSON.parse(
  readFileSync(path.join(ROOT, 'src', 'fixtures', 'report', 'meta.json'), 'utf8'),
) as Record<string, unknown>;
const specFixture = JSON.parse(
  readFileSync(path.join(ROOT, 'src', 'fixtures', 'sim', 'specs.json'), 'utf8'),
) as unknown[];

// Fight 3 (Warden Kelthas) is the only fixture fight whose COMBATANT_INFO row carries gear
// and talents -- sim-compare.spec.ts's own REF, and the only ref this fixture can open
// compare mode on.
const COMPARE_REF = 'fixture2abcd:3';
// Heroic Strike's client spell id -- the compare fixture's top ability row (Task 16). The
// row is keyed and tested by that integer identity, never by name, which can hold a space
// or a colon that no `slug()` exists to clean.
const HEROIC_STRIKE_SPELL_ID = 25286;

async function noHorizontalScroll(page: Page): Promise<void> {
  const overflow = await page.evaluate(() => document.documentElement.scrollWidth - window.innerWidth);
  expect(overflow, 'page scrolls horizontally').toBeLessThanOrEqual(0);
}

async function targetsAreBigEnough(page: Page): Promise<void> {
  const small = await page.evaluate(
    ({ target }) =>
      [...document.querySelectorAll('button, select, a[href], input, textarea')]
        .filter((element) => {
          const box = element.getBoundingClientRect();
          return box.width > 0 && box.height > 0 && element.closest('p') === null;
        })
        .filter((element) => {
          // Base.astro's skip link is the one control that is 1px on purpose -- it is
          // clipped until it takes focus, at which point its own `focus:h-11` gives it the
          // same 44px as everything else, the same exception report-phone.spec.ts,
          // rankings-phone.spec.ts, character-phone.spec.ts and guild-phone.spec.ts each
          // carve out for it.
          if (getComputedStyle(element).clipPath !== 'none') return false;
          // A checkbox is toggled through the label wrapping it (sim-execute,
          // sim-precision): the label is the target and the 20px box inside it is only the
          // mark, the same reading report-phone.spec.ts's own sweep gives every checkbox.
          const hitTarget =
            element instanceof HTMLInputElement &&
            element.type === 'checkbox' &&
            element.closest('label') !== null
              ? (element.closest('label') as HTMLElement)
              : element;
          const box = hitTarget.getBoundingClientRect();
          return Math.max(box.height, box.width) < target;
        })
        .map((element) => `${element.tagName}:${element.textContent?.trim().slice(0, 30) ?? ''}`),
    { target: TARGET },
  );
  expect(small, 'hit targets under 44px').toEqual([]);
}

/**
 * Every bordered card on the page -- `.rounded-panel` is the one class every such card
 * carries, from CharacterStrip's strip to CompareView's two tables -- keeps the page's 18px
 * gutter on both edges, either through its own `mx-[18px]` or a parent's. Scoped to `#sim`,
 * the island's own mount point on all three /sim routes, the same way report-phone.spec.ts
 * scopes its own island-only checks to `#report`: the shared header and footer are every
 * page's, not this one's, and layout.spec.ts already measures those.
 */
async function sectionsKeepGutter(page: Page): Promise<void> {
  const { width } = page.viewportSize() ?? { width: 0 };
  const edges = await page.evaluate(() =>
    [...document.querySelectorAll('#sim .rounded-panel')]
      .filter((element) => {
        const box = element.getBoundingClientRect();
        return box.width > 0 && box.height > 0;
      })
      .map((element) => {
        const box = element.getBoundingClientRect();
        const named =
          element.getAttribute('data-testid') ??
          element.closest('[data-testid]')?.getAttribute('data-testid') ??
          element.tagName;
        return { left: box.left, right: box.right, named };
      }),
  );
  expect(edges.length, 'no bordered panels found to measure').toBeGreaterThan(0);
  for (const edge of edges) {
    expect(edge.left, `${edge.named} left edge`).toBeGreaterThanOrEqual(GUTTER);
    expect(edge.right, `${edge.named} right edge`).toBeLessThanOrEqual(width - GUTTER);
  }
}

/**
 * SLOTS (planner/types.ts) begins head, neck, shoulder. In CharacterStrip's two-column
 * phone grid (`grid-cols-2`), that means head and neck share row 1 and shoulder starts row
 * 2 -- never head and shoulder sharing a row, which two columns cannot do.
 */
async function gearGridIsTwoColumns(page: Page): Promise<void> {
  const head = await page.getByTestId('sim-slot-head').boundingBox();
  const neck = await page.getByTestId('sim-slot-neck').boundingBox();
  const shoulder = await page.getByTestId('sim-slot-shoulder').boundingBox();
  expect(head, 'sim-slot-head missing').not.toBeNull();
  expect(neck, 'sim-slot-neck missing').not.toBeNull();
  expect(shoulder, 'sim-slot-shoulder missing').not.toBeNull();

  expect(Math.round(head!.y), 'head and neck are not the same row').toBe(Math.round(neck!.y));
  expect(head!.x, 'head is not left of neck').toBeLessThan(neck!.x);
  expect(shoulder!.y, 'shoulder is not a new row').toBeGreaterThan(head!.y);
}

async function loadFury(page: Page): Promise<void> {
  await page.goto('/sim');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();
}

async function runFury(page: Page): Promise<void> {
  await loadFury(page);
  await page.getByTestId('sim-run-button').click();
  await expect(page.getByTestId('sim-run-button')).toHaveText('Run again', { timeout: 10_000 });
  await expect(page.getByTestId('sim-results')).toBeVisible();
}

test('nothing scrolls sideways and every target clears 44px on an empty /sim', async ({ page }) => {
  await page.goto('/sim');
  await expect(page.getByTestId('sim-view')).toBeVisible();
  await expect(page.getByTestId('sim-empty')).toBeVisible();

  await noHorizontalScroll(page);
  await targetsAreBigEnough(page);
  await sectionsKeepGutter(page);
});

test('a loaded character keeps the run button and the DPS figure in view, in a two-column gear grid, with no sideways scroll', async ({
  page,
}) => {
  await loadFury(page);

  await noHorizontalScroll(page);
  await targetsAreBigEnough(page);
  await sectionsKeepGutter(page);
  await gearGridIsTwoColumns(page);

  // A player on a phone has to press run and see the number without moving. CharacterStrip
  // and SettingsBar are both taller than the phone viewport by design (Task 11: "nothing is
  // hidden" on phone), so reaching the button at all takes a scroll -- `.click()`'s own
  // actionability check is that scroll, the same way a player's thumb would reach it. What
  // this proves is the part that is not a given: once there, the figure sits right beside
  // the button rather than a second scroll away.
  await page.getByTestId('sim-run-button').click();
  await expect(page.getByTestId('sim-run-button')).toBeInViewport();
  await expect(page.getByTestId('sim-dps')).toBeInViewport();
});

test('after a run, every result tab clears the width and the tab strip scrolls rather than wraps', async ({
  page,
}) => {
  await runFury(page);

  const tabs = await page.locator('[role="tab"]').all();
  expect(tabs.length, 'no result tabs found').toBeGreaterThan(0);
  for (const tab of tabs) {
    await tab.click();
    await noHorizontalScroll(page);
    await targetsAreBigEnough(page);
  }

  // The strip itself scrolls sideways to reach a tab past the fold -- it does not wrap and
  // it does not widen the page.
  const distribution = page.getByTestId('sim-tab-distribution');
  await distribution.scrollIntoViewIfNeeded();
  await expect(distribution).toBeInViewport();

  const strip = await page.evaluate(() => {
    const tablist = document.querySelector('[role="tablist"]');
    return {
      scrollWidth: tablist?.scrollWidth ?? 0,
      clientWidth: tablist?.clientWidth ?? 0,
      pageScrollWidth: document.documentElement.scrollWidth,
      pageClientWidth: document.documentElement.clientWidth,
    };
  });
  expect(strip.scrollWidth, 'tab strip does not scroll').toBeGreaterThan(strip.clientWidth);
  expect(strip.pageScrollWidth, 'the page scrolled sideways instead of the strip').toBeLessThanOrEqual(
    strip.pageClientWidth,
  );
});

test('compare mode stacks the ability table into one figure per row, with no sideways scroll', async ({
  page,
}) => {
  await page.route('**/v1/reports/fixture2abcd', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({ ok: true, data: meta, error: null, request_id: 'r' }),
    }),
  );
  await page.goto(`/sim?source=fight&ref=${encodeURIComponent(COMPARE_REF)}&mode=compare`);
  await expect(page.getByTestId('compare-headline')).toBeVisible({ timeout: 10_000 });

  await noHorizontalScroll(page);
  await targetsAreBigEnough(page);
  await sectionsKeepGutter(page);

  // Heroic Strike's row: the phone label carrying both figures in one line is visible, and
  // the desktop-only damage cells it stands in for are not -- never a `slug()`ed name,
  // which the row's own test id was written to avoid needing.
  const row = page.getByTestId(`compare-ability-${HEROIC_STRIKE_SPELL_ID}`);
  await expect(row).toBeVisible();
  await expect(row.locator('span[class*="md:hidden"]')).toBeVisible();
  for (const cell of await row.locator('span[class*="md:inline"]').all()) {
    await expect(cell).toBeHidden();
  }
});

test('/sim/specs has no sideways scroll and every card clears 44px', async ({ page }) => {
  await page.route('**/v1/specs**', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({ ok: true, data: { specs: specFixture }, error: null, request_id: 'r' }),
    }),
  );
  await page.goto('/sim/specs');
  await expect(page.getByTestId('specs-grid')).toBeVisible();

  await noHorizontalScroll(page);
  await targetsAreBigEnough(page);
  await sectionsKeepGutter(page);
});

test('a saved sim at /sim/simfixtureab has no sideways scroll and every target clears 44px', async ({
  page,
}) => {
  await page.goto('/sim/simfixtureab');
  await expect(page.getByTestId('sim-saved-header')).toBeVisible();
  await expect(page.getByTestId('sim-results')).toBeVisible();

  await noHorizontalScroll(page);
  await targetsAreBigEnough(page);
  await sectionsKeepGutter(page);
});
