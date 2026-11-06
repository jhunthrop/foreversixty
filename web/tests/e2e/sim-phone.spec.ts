// web/tests/e2e/sim-phone.spec.ts
// The phone audit for the simulator: report-phone.spec.ts, rankings-phone.spec.ts,
// character-phone.spec.ts and guild-phone.spec.ts's counterpart for every /sim route. Every
// component in this lane was written with its phone layout in mind; this file proves it,
// across every state a player actually reaches -- empty, a loaded character, a finished run
// on every result tab, compare mode, the spec-support grid and a saved sim -- rather than
// trusting that no later change quietly reopened one.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { closeMoreSettings, gotoSim, openMoreSettings, pasteAndRun } from './support/sim-page';
import { assertNoHorizontalScroll } from './support/phone-scroll';
import { fieldBox } from './support/select';

test.describe.configure({ mode: 'serial' });
// No test.use({ viewport }): the phone audit runs under --project=mobile (Pixel 7, 412px),
// which is the width the Global Constraints name, and overriding it here would audit a
// size nothing else in the suite uses.

// H2: every sibling phone spec (report-phone.spec.ts, rankings-phone.spec.ts,
// character-phone.spec.ts, guild-phone.spec.ts) carries this guard; this file relies on
// --project=mobile alone for its 412px width, so without it the desktop project ran it too
// and failed on the shared desktop header's sub-44px nav links -- and, because the file is
// serial, took the five cases after it down with it.
test.skip(() => test.info().project.name !== 'mobile', 'phone layout only');

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

/**
 * Against `page.viewportSize()`, never the live `window.innerWidth` -- Chromium's mobile
 * emulation widens the layout viewport to fit overflowing content, so `innerWidth` and
 * `scrollWidth` grow together and a comparison between them silently reads ~0 overflow
 * even while the page genuinely overflows the device's own width. This is not
 * hypothetical: it is exactly how the per-tab sweep in "after a run..." below missed
 * SimResults' `-mx-[18px]` tab-strip overflow (Fix round 1) until a fixed-width check
 * caught it. See support/phone-scroll.ts's own header note; report-phone.spec.ts hit the
 * same class of false negative first.
 */
async function noHorizontalScroll(page: Page): Promise<void> {
  const width = page.viewportSize()?.width ?? 0;
  await assertNoHorizontalScroll(page, width);
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
 * Every bordered card on the page -- the flat pages' Plate (`.ui-plate`, spec 2026-11-06-sim
 * §4.1) and `.rounded-panel` (CompareView's two tables) -- keeps the page's 18px
 * gutter on both edges, either through its own `mx-[18px]` or a parent's. Scoped to `#sim`,
 * the island's own mount point on all three /sim routes, the same way report-phone.spec.ts
 * scopes its own island-only checks to `#report`: the shared header and footer are every
 * page's, not this one's, and layout.spec.ts already measures those.
 */
async function sectionsKeepGutter(page: Page): Promise<void> {
  const { width } = page.viewportSize() ?? { width: 0 };
  const edges = await page.evaluate(() =>
    [
      ...document.querySelectorAll(
        '#sim .ui-plate, #sim .rounded-panel, #sim .card-surface, #sim .card-surface-quiet',
      ),
    ]
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
 * The shared paperdoll on a phone (web layout spec 2026-11-05 §4.3): one column in the character
 * panel's reading order -- head, then neck under it, the right column's hands further down.
 */
async function gearGridIsOneColumn(page: Page): Promise<void> {
  // One read of all three, so a scroll still settling cannot move one between them.
  const [head, neck, hands] = await page.evaluate(() =>
    ['head', 'neck', 'hands'].map((slot) => {
      const box = document
        .querySelector(`[data-testid="sim-gear-doll-slot-${slot}"]`)
        ?.getBoundingClientRect();
      return box === undefined ? null : { x: box.x, y: box.y };
    }),
  );
  expect(head, 'sim-gear-doll-slot-head missing').not.toBeNull();
  expect(neck, 'sim-gear-doll-slot-neck missing').not.toBeNull();
  expect(hands, 'sim-gear-doll-slot-hands missing').not.toBeNull();

  expect(Math.round(neck!.x), 'neck is not in the head column').toBe(Math.round(head!.x));
  expect(neck!.y, 'neck is not under head').toBeGreaterThan(head!.y);
  expect(hands!.y, 'hands is not under the left column').toBeGreaterThan(neck!.y);
}

// The landing's paste Plate pastes and runs in one click (SimView.svelte's `runPastedInput`),
// so `loadFury` waits out that run.
async function loadFury(page: Page): Promise<void> {
  await pasteAndRun(page, FURY);
}

async function runFury(page: Page): Promise<void> {
  await loadFury(page);
  await expect(page.getByTestId('sim-live-report')).toBeVisible();
}

test('nothing scrolls sideways and every target clears 44px on an empty /sim', async ({ page }) => {
  await gotoSim(page);
  await expect(page.getByTestId('sim-view')).toBeVisible();
  await expect(page.getByTestId('sim-landing')).toBeVisible();

  await noHorizontalScroll(page);
  await targetsAreBigEnough(page);
  await sectionsKeepGutter(page);
});

test('a loaded character reaches the run button and the DPS figure without a scroll after pressing run, in a one-column gear grid, with no sideways scroll', async ({
  page,
}) => {
  // A code link loads the character without running it: the doll is the page's one Plate.
  await gotoSim(page, `/sim?code=${encodeURIComponent(FURY)}`);
  await expect(page.getByTestId('sim-gear-doll-slot-head')).toBeVisible();

  // More settings is closed on arrival: an open sheet would cover the page.
  await expect(page.getByTestId('sim-more-settings')).toHaveCount(0);

  await noHorizontalScroll(page);
  await targetsAreBigEnough(page);
  await sectionsKeepGutter(page);
  await gearGridIsOneColumn(page);

  // The secondary controls are in More settings; the audit measures them open, as a player
  // opens them.
  await openMoreSettings(page);
  await expect(fieldBox(page, 'sim-target-level')).toBeVisible();
  await targetsAreBigEnough(page);
  await closeMoreSettings(page);

  // A player on a phone presses run and sees the number without moving: the answer sits in
  // the head above the run line (spec 2026-11-06-sim §4.4, "the answer is in the first
  // screen"), and the bar is sticky.
  await page.evaluate(() => window.scrollTo(0, 0));
  await page.getByTestId('sim-run-button').click();
  await expect(page.getByTestId('sim-run-button')).toBeInViewport();
  await expect(page.getByTestId('sim-dps')).toBeInViewport();
});

test('after a run, the report clears the width and every target clears 44px', async ({ page }) => {
  await runFury(page);
  await expect(page.getByTestId('sim-damage-card')).toBeVisible();
  // The press that ran it can leave the pointer over a row the report drew under it, and that
  // row's item tooltip is a floating layer, not a card in the column.
  await page.mouse.move(0, 0);
  await page.keyboard.press('Escape');
  await expect(page.getByTestId('item-tooltip')).toHaveCount(0);

  await noHorizontalScroll(page);
  await targetsAreBigEnough(page);
  await sectionsKeepGutter(page);

  // The deeper cards load as they near the viewport; the page still never scrolls sideways.
  await page.getByTestId('sim-section-settings').scrollIntoViewIfNeeded();
  await expect(page.getByTestId('sim-settings-card')).toBeVisible({ timeout: 10_000 });
  await noHorizontalScroll(page);
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
  await expect(page.getByTestId('sim-saved')).toBeVisible();
  await expect(page.getByTestId('sim-damage-card')).toBeVisible();

  await noHorizontalScroll(page);
  await targetsAreBigEnough(page);
  await sectionsKeepGutter(page);
});
