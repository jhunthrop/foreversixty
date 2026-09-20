// web/tests/e2e/sim-tools-phone.spec.ts
// The four tool pages at phone width. Everything here is about layout and reach, not about
// the engine: a control that overflows the viewport or is under 44px tall is unusable, and
// nothing else in the suite measures that.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test, type Page } from '@playwright/test';
import { simCopy } from '../../src/lib/sim/copy';

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
  // Account.svelte's nav widget (client:load) starts on an invisible, same-size placeholder
  // (no href, so it never matches the hit-target query below) and swaps in a real `<a
  // href="/login">` or `<a href="/account">` once its own session check resolves. Waiting for
  // that swap here, rather than racing it, is what makes the hit-target check below
  // deterministic -- without it, "/login" appeared in the failing set on some runs and not
  // others, purely from timing, which is not a real regression to chase.
  await expect(page.getByTestId('session-nav').locator('a')).toBeVisible();
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
    //
    // Fix round 2 removed the last of this spec's KNOWN_SMALL_TARGETS allowlist (Header's
    // "Logs", Footer's "About", Account.svelte's "Sign in", CharacterStrip's "Open in
    // planner" and SettingsBar's rotation link all now clear 44px), so the mechanism itself
    // went with it rather than standing empty -- a plain assertion reads cleaner than an
    // allowlist with nothing left to allow.
    const bad = await page.evaluate(() => {
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
        if (target.height >= 44 && target.width >= 44) continue;
        bad.push(element.outerHTML.slice(0, 80));
      }
      return bad;
    });

    expect(bad, 'a control is under 44px in some dimension').toEqual([]);
  });
}

test('the run bar stays reachable on /sim/gear after a run', async ({ page }) => {
  await load(page, '/sim/gear');
  await page.getByTestId('sim-search-add-16963').click();
  await page.getByTestId('sim-run-bulk').click();
  await expect(page.getByTestId('sim-combos')).toBeVisible({ timeout: 25_000 });
  await expect(page.getByTestId('sim-run-bulk')).toBeVisible();
});

// Task 5 (newcomer MAJOR, review.md:360-363; MINOR, review.md:365-370): the same three tool
// pages' help and run-bar controls, proven at the review's own 390x844 -- a second width
// alongside this file's 360x780 above, not a replacement for it, so a control that only
// just clears 44px at 360 is not quietly let off the hook at the size the review actually
// measured.
test.describe('touch help and hit targets at 390x844 (Task 5)', () => {
  test.use({ viewport: { width: 390, height: 844 } });

  for (const route of ['/sim/gear', '/sim/talents', '/sim/drops'] as const) {
    test(`${route}: "more settings" relies on no hover-only title`, async ({ page }) => {
      await load(page, route);
      await page.getByTestId('sim-settings-more').locator('summary').click();
      // Scoped to the tool island, not `document` page-globally (fix round, Minor 2): a
      // `title=` added anywhere else on the page (header, tab strip, footer) by a sibling
      // lane would fail this spec with a failure that reads as this lane's own bug. The two
      // later cases in this file already scope to `sim-combos`, one level narrower still.
      const titled = await page
        .getByTestId('sim-tools-view')
        .evaluate((root) =>
          [...root.querySelectorAll('[title]')].map(
            (element) => element.getAttribute('data-testid') ?? element.outerHTML.slice(0, 60),
          ),
        );
      expect(titled, 'a control still relies on a hover-only title').toEqual([]);
    });
  }

  test('/sim/gear: the variation and dummy notes are visible text, and their rows clear 44px', async ({
    page,
  }) => {
    await load(page, '/sim/gear');
    await page.getByTestId('sim-settings-more').locator('summary').click();

    // SettingsSheet.svelte:61 and :142 carried `title={simCopy.variationNote}` and
    // `title={simCopy.dummyNote}` -- never reachable on a phone. Both words reach a phone
    // now: Task 5's plain paragraphs became Task 7's HelpNote (Disclosure.svelte), one
    // tappable trigger per control, so the assertion is the same two sentences arriving on
    // screen after a tap rather than a `<p>` that was always there.
    for (const [id, text] of [
      ['sim-variation', simCopy.variationNote],
      ['sim-dummy', simCopy.dummyNote],
    ] as const) {
      const trigger = page.getByTestId(`${id}-help-trigger`);
      await expect(trigger).toBeVisible();
      await trigger.click();
      await expect(page.getByTestId(`${id}-help-panel`)).toHaveText(text);
    }

    // sim-execute/sim-dummy: the drawn mark stays 20x20 on purpose -- native `accent-gold`,
    // never resized, per the brief's own "without changing the visual size of the tick" --
    // but each sits inside a `flex min-h-11` label stretched across the full grid cell, the
    // real hit target a tap actually reaches. Measuring the hidden input itself instead is
    // the exact false positive sim-phone.spec.ts's own targetsAreBigEnough, and this file's
    // own 44px sweep above, already carve out for every checkbox in this suite.
    for (const id of ['sim-execute', 'sim-dummy'] as const) {
      const box = await page.getByTestId(id).evaluate((element) => {
        const rect = (element.closest('label') as HTMLElement).getBoundingClientRect();
        return { width: rect.width, height: rect.height };
      });
      expect(box.width, `${id} label width`).toBeGreaterThanOrEqual(44);
      expect(box.height, `${id} label height`).toBeGreaterThanOrEqual(44);
    }

    // BulkRunBar's precision label: `min-h-11` is explicit now (Task 5) instead of relying
    // on the select child's own height to stretch a plain flex row tall enough by accident.
    const precisionBox = await page.getByTestId('sim-precision').evaluate((element) => {
      const rect = (element.closest('label') as HTMLElement).getBoundingClientRect();
      return { width: rect.width, height: rect.height };
    });
    expect(precisionBox.width, 'precision label width').toBeGreaterThanOrEqual(44);
    expect(precisionBox.height, 'precision label height').toBeGreaterThanOrEqual(44);
  });

  test('/sim/gear: after a run, the results carry no hover-only title', async ({ page }) => {
    await load(page, '/sim/gear');
    await page.getByTestId('sim-search-add-16963').click();
    await page.getByTestId('sim-run-bulk').click();
    await expect(page.getByTestId('sim-combos')).toBeVisible({ timeout: 25_000 });
    const titled = await page
      .getByTestId('sim-combos')
      .evaluate((element) => element.querySelectorAll('[title]').length);
    expect(titled, 'a control inside the results still relies on a hover-only title').toBe(0);
  });

  test('/sim/drops: after a run, a chip names its boss in visible text and the results carry no hover-only title', async ({
    page,
  }) => {
    await load(page, '/sim/drops');
    // Molten Core is gated ahead of today (phases.json), so "show unreleased content" is
    // what makes it pickable at all -- the same setup sim-drops.spec.ts's own "picking a
    // boss..." case uses.
    await page.getByTestId('sim-upcoming').check();
    await page.getByTestId('sim-source-raid:molten-core:11502').check();
    await page.getByTestId('sim-run-bulk').click();
    await expect(page.getByTestId('sim-drops-by-boss')).toBeVisible({ timeout: 25_000 });
    const titled = await page
      .getByTestId('sim-combos')
      .evaluate((element) => element.querySelectorAll('[title]').length);
    expect(titled, 'a control inside the results still relies on a hover-only title').toBe(0);
    // SubstitutionChips.svelte:55 carried the boss name in a `title` alongside the item; it
    // is now part of the chip's own visible text (substitutionChipLabel, combos.ts).
    await expect(page.getByTestId('sim-combo-row').first()).toContainText('Ragnaros');
  });
});
