// web/tests/e2e/character-gear.spec.ts
// Task L8.6 and its fix round 1: the profile's Gear section, from the character answer's
// `gear` and `gear_gap_total` (api/openapi.yaml `CharacterGear`), stubbed with every row state
// (src/fixtures/character-gear.ts), the build's items/priest.json reduced to the fixture's real
// rows, and the real build's icons.
//   L8.6 ruling 1 / fix 1.   Under the card; the names marked as /bis marks them, opening the
//                            item tooltip on hover, focus and tap with no page error.
//   fix 2.  No white name: the rows wait for the items file; a failed file gives the body
//           colour and the next open retries.
//   fix 3.  The totals line under the title, and the title's variants.
//   fix 4.  Phones (390, 320): label and gap on line 1, worn → pick on line 2, the source on 3.
//   fix 5.  Desktop: the source is a column, the gaps one right edge, a row tint on hover.
//   fix 11. The landing focuses the heading, with gear and without.
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import {
  CHARACTER,
  CHARACTER_PATH,
  CHARACTER_URL,
  GEAR,
  GEAR_ITEMS,
  GEAR_LONG_NAME,
  routeGearIcons,
} from './support/character-fixture';

const NOT_FOUND = {
  status: 404,
  contentType: 'application/json',
  body: JSON.stringify({ ok: false, data: null, error: 'not_found', request_id: 'r' }),
};

const ok = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
});

/** The score whose "Most to gain" is Gear, 88 below its top. */
const SCORE = {
  score: 1830,
  tier: 'epic',
  as_of: '2026-11-04',
  components: { gear: 412, experience: 620, parses: 540, attendance: 180, discipline: 78 },
  experience: { dungeons: 14, raid_kills: 0, alts: 2 },
};

const RGB = {
  white: 'rgb(255, 255, 255)',
  body: 'rgb(233, 228, 216)',
  rare: 'rgb(61, 148, 240)',
  epic: 'rgb(184, 102, 245)',
  muted: 'rgb(154, 148, 132)',
};

interface OpenOptions {
  gear?: unknown;
  /** Hold the items file this long before answering (a failure too). */
  itemsDelayMs?: number;
  /** Answer the items file 500 this many times first. */
  itemsFailures?: number;
  /** Icon file names to answer 404. */
  missingIcons?: string[];
}

async function open(page: Page, options: OpenOptions = {}): Promise<string[]> {
  const errors: string[] = [];
  page.on('pageerror', (error) => errors.push(error.message));
  await page.route(CHARACTER_URL, (route) =>
    route.fulfill(ok({ ...CHARACTER.data, ...SCORE, ...((options.gear ?? GEAR) as object) })),
  );
  await page.route(`${CHARACTER_URL}/rating`, (route) => route.fulfill(NOT_FOUND));
  let failures = options.itemsFailures ?? 0;
  await page.route('**/data/*/items/priest.json', async (route) => {
    if (options.itemsDelayMs !== undefined) await new Promise((r) => setTimeout(r, options.itemsDelayMs));
    if (failures > 0) {
      failures -= 1;
      return route.fulfill({ status: 500, body: '' });
    }
    return route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(GEAR_ITEMS) });
  });
  await routeGearIcons(page, options.missingIcons);
  await page.goto(CHARACTER_PATH);
  await expect(page.getByTestId('forever-score-card')).toBeVisible();
  return errors;
}

/** Waits for the rows to be drawn (the items file in, or failed). */
async function rowsDrawn(page: Page): Promise<void> {
  await expect(page.getByTestId('character-gear-list')).not.toHaveAttribute('aria-busy', 'true');
}

/** Waits for the page to settle (the other islands' reads answered), so nothing above the
 *  section moves a row from under the pointer. */
async function settled(page: Page): Promise<void> {
  await rowsDrawn(page);
  await page.waitForLoadState('networkidle');
}

const row = (page: Page, slot: string) => page.getByTestId(`character-gear-row-${slot}`);

/** Every row's cells, read in one pass so a settling layout cannot split them. */
async function rowBoxes(page: Page) {
  return page.evaluate(() =>
    [...document.querySelectorAll('[data-testid^="character-gear-row-"]')].map((li) => {
      const box = (testid: string) => {
        const el = li.querySelector(`[data-testid="${testid}"]`);
        if (el === null) return null;
        const r = el.getBoundingClientRect();
        return { top: r.top, bottom: r.bottom, left: r.left, right: r.right, height: r.height };
      };
      const r = li.getBoundingClientRect();
      return {
        slot: li.getAttribute('data-testid')?.replace('character-gear-row-', '') ?? '',
        top: r.top,
        height: r.height,
        label: box('character-gear-slot'),
        items: box('character-gear-items'),
        meta: box('character-gear-meta'),
        gap: box('character-gear-gap'),
      };
    }),
  );
}

/** Every name and empty-slot word, and every note, with what would keep it from being seen: its
 *  own visibility, and the opacity of it and each ancestor (fix 6 item 1). */
async function namesSeen(page: Page) {
  return page.getByTestId('character-gear-list').evaluate((list) =>
    [...list.querySelectorAll<HTMLElement>('li')].map((li) => {
      const cells = [
        ...li.querySelectorAll<HTMLElement>(
          '[data-testid="character-gear-worn"], [data-testid="character-gear-pick"], [data-testid="character-gear-empty"], [data-testid="character-gear-meta"]',
        ),
      ];
      return {
        slot: li.dataset.testid?.replace('character-gear-row-', '') ?? '',
        cells: cells.map((cell) => {
          let opacity = 1;
          for (let el: HTMLElement | null = cell; el !== null && el !== list; el = el.parentElement)
            opacity *= Number(getComputedStyle(el).opacity);
          const box = cell.getBoundingClientRect();
          return {
            text: cell.textContent?.trim() ?? '',
            visibility: getComputedStyle(cell).visibility,
            opacity,
            drawn: box.width > 0 && box.height > 0,
          };
        }),
      };
    }),
  );
}

async function noSidewaysScroll(page: Page): Promise<void> {
  const scroll = await page.evaluate(() => ({
    width: document.documentElement.scrollWidth,
    client: document.documentElement.clientWidth,
  }));
  expect(scroll.width).toBeLessThanOrEqual(scroll.client);
}

for (const width of [1440, 390]) {
  test.describe(`the Gear section at ${width} (L8.6 fix round 1)`, () => {
    test.use({ viewport: { width, height: width === 1440 ? 900 : 844 } });

    test('sits under the card; the totals line 4px under the title and 16px above row 1 (fix 3)', async ({
      page,
    }) => {
      await open(page);
      await rowsDrawn(page);
      await expect(page.getByTestId('character-gear-title')).toHaveText(
        'Gear 6 of 15 slots best in slot · +70 points across 7 slots · 1 not scored',
      );
      const totals = page.getByTestId('character-gear-totals');
      await expect(totals).toHaveText(
        'The card’s 88 counts the level range’s best set. These 70 count the picks you can wear.',
      );
      await expect(totals).toHaveCSS('font-size', '13px');
      await expect(totals).toHaveCSS('color', RGB.muted);
      const gaps = await page.evaluate(() => {
        const box = (testid: string) =>
          document.querySelector(`[data-testid="${testid}"]`)!.getBoundingClientRect();
        const card = box('forever-score');
        const section = box('character-gear');
        const title = box('character-gear-title');
        const totals = box('character-gear-totals');
        const first = box('character-gear-row-head');
        return {
          card: section.top - card.bottom,
          title: totals.top - title.bottom,
          row: first.top - totals.bottom,
        };
      });
      expect(gaps.card).toBeGreaterThan(0);
      expect(gaps.card).toBeLessThanOrEqual(33);
      expect(Math.abs(gaps.title - 4)).toBeLessThan(1);
      expect(Math.abs(gaps.row - 16)).toBeLessThan(1);
      await expect(page.getByTestId('character-gear-footer')).toHaveText('Full Discipline BiS list at 60 →');
      await expect(page.getByTestId('character-gear-bis')).toHaveAttribute(
        'href',
        '/bis/priest/discipline?faction=alliance#band-alliance-60',
      );
    });

    test('draws every row state, names in their quality colours with their icons (fixes 1, 7, 8, 9, 10)', async ({
      page,
    }) => {
      const errors = await open(page);
      await rowsDrawn(page);
      const head = row(page, 'head');
      await expect(head.getByTestId('character-gear-slot')).toHaveText('Head');
      await expect(head.getByTestId('character-gear-worn')).toHaveCSS('color', RGB.rare);
      await expect(head.getByTestId('character-gear-pick')).toHaveCSS('color', RGB.epic);
      // The dot between the source and the chance is drawn (::before), not text.
      await expect(head.getByTestId('character-gear-meta')).toContainText("Onyxia's Lair: Onyxia");
      await expect(head.getByTestId('character-gear-chance')).toHaveText('17% drop');
      await expect(head.getByTestId('character-gear-gap')).toHaveText('+18.4 points');
      await expect(head.getByTestId('character-gear-gap')).toHaveCSS('font-weight', '700');
      await expect(row(page, 'wrist').getByTestId('character-gear-gap')).toHaveCSS('font-weight', '400');
      // Icons: a 20px box, the pick's border its quality, the image loaded in it.
      const box = head.locator('[data-bis-item="16921"] [data-testid="character-gear-icon"]');
      await expect(box).toHaveCSS('width', '20px');
      await expect(box).toHaveCSS('height', '20px');
      await expect(box).toHaveCSS('border-top-color', 'rgb(163, 53, 238)');
      // A lazy image: wait for it to decode rather than read it once.
      await expect
        .poll(() => box.locator('img').evaluate((img: HTMLImageElement) => img.naturalWidth))
        .toBeGreaterThan(0);
      await expect(row(page, 'back').getByTestId('character-gear-meta')).toContainText('<1% drop');
      await expect(row(page, 'chest').getByTestId('character-gear-meta')).toHaveText('Tailoring');
      await expect(row(page, 'shoulder').getByTestId('character-gear-meta')).toContainText(
        'Already your best at this level',
      );
      await expect(row(page, 'shoulder').getByTestId('character-gear-gap')).toHaveText('');
      await expect(row(page, 'finger1').getByTestId('character-gear-slot')).toHaveText('Ring 1');
      await expect(row(page, 'trinket2').getByTestId('character-gear-slot')).toHaveText('Trinket 2');
      await expect(row(page, 'trinket2').getByTestId('character-gear-meta')).toHaveText(
        'No pick for your level range yet',
      );
      await expect(row(page, 'ranged').getByTestId('character-gear-meta')).toHaveText(
        'No pick for your level range yet',
      );
      const legs = row(page, 'legs');
      await expect(legs.getByTestId('character-gear-worn')).toHaveText('Not in our item data yet');
      await expect(legs.getByTestId('character-gear-worn')).toHaveCSS('color', RGB.muted);
      await expect(legs.getByTestId('character-gear-gap')).toHaveText('Not scored');
      await expect(legs.getByTestId('character-gear-gap')).toHaveCSS('color', RGB.muted);
      const off = row(page, 'off_hand');
      await expect(off.getByTestId('character-gear-meta')).toContainText('Blackrock Spire: Gyth');
      await expect(off.getByTestId('character-gear-chance')).toHaveText('9% drop');
      await expect(off.getByTestId('character-gear-note')).toHaveText('comes with the main-hand swap');
      await expect(off.getByTestId('character-gear-note')).toHaveCSS('visibility', 'visible');
      await expect(off.getByTestId('character-gear-gap')).toHaveText('');
      expect(errors).toEqual([]);
    });

    test('lands the Gear link on the section and focuses its heading (fix 11)', async ({ page }) => {
      await open(page);
      await rowsDrawn(page);
      const link = page.getByTestId('forever-score-most-to-gain').getByRole('link', { name: 'Gear' });
      await expect(link).toHaveAttribute('href', '#gear');
      await settled(page);
      // A real pointer click: the mouse stays where the link was while the page scrolls.
      await link.click();
      await expect(page).toHaveURL(/#gear$/);
      await expect(page.getByTestId('character-gear-title')).toBeFocused();
      // Whatever row now sits under the stationary pointer opens no tooltip (fix 2 ruling 2).
      await page.waitForTimeout(600);
      await expect(page.getByTestId('item-tooltip')).toHaveCount(0);
      const top = await page
        .getByTestId('character-gear')
        .evaluate((element) => element.getBoundingClientRect().top);
      expect(top).toBeGreaterThanOrEqual(95);
      expect(top).toBeLessThan(110);
    });

    test('`gear: []`: the heading over one line; the link lands there and focuses the heading (fix 11)', async ({
      page,
    }) => {
      await open(page, { gear: { gear: [], gear_gap_total: null } });
      await expect(page.getByTestId('character-gear')).toHaveCount(0);
      const heading = page.getByTestId('character-gear-title');
      const line = page.getByTestId('character-gear-none');
      await expect(heading).toHaveText('Gear');
      await expect(line).toHaveText(
        'No gear on file yet. Sync the addon, or sign in to claim this character.',
      );
      await page.getByTestId('forever-score-most-to-gain').getByRole('link', { name: 'Gear' }).click();
      await expect(page).toHaveURL(/#gear$/);
      await expect(heading).toBeFocused();
      await expect(heading).toBeInViewport();
      await expect(line).toBeInViewport();
    });

    test('never paints a name white while the items file is slow (fix 2)', async ({ page }) => {
      await page.addInitScript(() => {
        (window as unknown as { __white: string[] }).__white = [];
        const sample = () => {
          for (const el of document.querySelectorAll<HTMLElement>(
            '[data-testid="character-gear-worn"], [data-testid="character-gear-pick"]',
          )) {
            const style = getComputedStyle(el);
            const shown = el.closest('.invisible') === null && style.visibility === 'visible';
            if (shown && style.color === 'rgb(255, 255, 255)')
              (window as unknown as { __white: string[] }).__white.push(el.textContent ?? '');
          }
          requestAnimationFrame(sample);
        };
        requestAnimationFrame(sample);
      });
      await open(page, { itemsDelayMs: 1500 });
      await expect(page.getByTestId('character-gear-title')).toBeVisible();
      await expect(page.getByTestId('character-gear-list')).toHaveAttribute('aria-busy', 'true');
      await rowsDrawn(page);
      await expect(row(page, 'head').getByTestId('character-gear-pick')).toHaveCSS('color', RGB.epic);
      await page.waitForTimeout(200);
      expect(await page.evaluate(() => (window as unknown as { __white: string[] }).__white)).toEqual([]);
    });

    test('an icon that fails keeps its bordered box, never the broken-image glyph (fix 2 ruling 1)', async ({
      page,
    }) => {
      await open(page, { missingIcons: ['inv_helmet_24.webp'] });
      await rowsDrawn(page);
      const box = row(page, 'head').locator('[data-bis-item="16921"] [data-testid="character-gear-icon"]');
      await expect(box.locator('img')).toHaveCount(0);
      await expect(box).toBeVisible();
      await expect(box).toHaveCSS('width', '20px');
      // The skeleton's treatment: #14171f, the muted line border, never a white checkbox (fix 4 A).
      await expect(box).toHaveCSS('background-color', 'rgb(20, 23, 31)');
      await expect(box).toHaveCSS('border-top-color', 'rgb(38, 46, 64)');
      // The others load.
      const worn = row(page, 'head').locator('[data-bis-item="10504"] img');
      expect(await worn.evaluate((img: HTMLImageElement) => img.naturalWidth)).toBeGreaterThan(0);
    });

    test('names the list "Gear by slot" and each row by its slot; 24px name targets (fix 2 ruling 9)', async ({
      page,
    }) => {
      await open(page);
      await rowsDrawn(page);
      const list = page.getByRole('list', { name: 'Gear by slot' });
      await expect(list).toBeVisible();
      await expect(list.getByRole('listitem', { name: 'Trinket 2', exact: true })).toHaveCount(1);
      await expect(list.getByRole('listitem', { name: 'Head', exact: true })).toHaveCount(1);
      const name = row(page, 'head').locator('[data-bis-item="16921"]');
      const height = await name.evaluate((element) => element.getBoundingClientRect().height);
      expect(height).toBeGreaterThanOrEqual(24);
      await expect(page.getByTestId('character-gear-bis')).toHaveCSS('padding-top', '12px');
      await expect(page.getByTestId('character-gear-bis')).toHaveCSS('padding-bottom', '12px');
    });

    test('a failed items file draws the names in the body colour; the next open retries (fix 2)', async ({
      page,
    }) => {
      await open(page, { itemsFailures: 1 });
      await rowsDrawn(page);
      await expect(row(page, 'head').getByTestId('character-gear-pick')).toHaveCSS('color', RGB.body);
      // The swap cannot be known: the empty off hand shows no figure (fix 2 ruling 7), its note's
      // line kept blank (fix 4 item 4).
      await expect(row(page, 'off_hand').getByTestId('character-gear-gap')).toHaveText('');
      await expect(row(page, 'off_hand').getByTestId('character-gear-note')).toHaveCSS(
        'visibility',
        'hidden',
      );
      await expect(row(page, 'off_hand').getByTestId('character-gear-note-bar')).toHaveCount(0);
      await expect(page.getByTestId('character-gear-failed')).toHaveText('Item details didn’t load.');
      await expect(page.getByTestId('character-gear-failed')).toHaveCSS('color', RGB.muted);
      // In the footer, under the BiS link (fix 6 item 2).
      await expect(
        page.getByTestId('character-gear-footer').getByTestId('character-gear-failed'),
      ).toHaveCount(1);
      const [link, failed] = await Promise.all(
        ['character-gear-bis', 'character-gear-failed'].map((id) => page.getByTestId(id).boundingBox()),
      );
      expect(failed!.y).toBeGreaterThanOrEqual(link!.y + link!.height - 1);
      // Every row's names and notes are seen, none held: the Ranged row's worn item with no pick
      // too (fix 6 item 1).
      await expect(page.getByTestId('character-gear-skeleton')).toHaveCount(0);
      const seen = await namesSeen(page);
      expect(seen).toHaveLength(17);
      for (const { slot, cells } of seen) {
        expect(cells.length, slot).toBeGreaterThan(0);
        for (const cell of cells)
          expect(cell, `${slot} ${cell.text}`).toMatchObject({
            visibility: 'visible',
            opacity: 1,
            drawn: true,
          });
      }
      const ranged = seen.find((r) => r.slot === 'ranged')!.cells.map((cell) => cell.text);
      expect(ranged).toEqual(["Skul's Ghastly Touch", 'No pick for your level range yet']);
      await expect(row(page, 'ranged').getByTestId('character-gear-worn')).toHaveCSS('color', RGB.body);
      // Every icon box the muted line, no white border.
      const borders = await page
        .getByTestId('character-gear-icon')
        .evaluateAll((boxes) => [...new Set(boxes.map((box) => getComputedStyle(box).borderTopColor))]);
      expect(borders).toEqual(['rgb(38, 46, 64)']);
      await page.reload();
      await rowsDrawn(page);
      await expect(row(page, 'head').getByTestId('character-gear-pick')).toHaveCSS('color', RGB.epic);
      await expect(page.getByTestId('character-gear-failed')).toHaveCount(0);
    });
  });
}

test.describe('the item tooltip at 1440 (fix 1)', () => {
  test.use({ viewport: { width: 1440, height: 900 } });

  test('opens on hover and on keyboard focus, with no page error', async ({ page }, info) => {
    const errors = await open(page);
    await settled(page);
    const pick = row(page, 'head').locator('[data-bis-item="16921"]');
    const tooltip = page.getByTestId('item-tooltip');
    // Hover where there is a pointer (the desktop project); a touch device has no hover.
    if (info.project.use.hasTouch !== true) {
      await pick.hover();
      await expect(tooltip).toBeVisible();
      await expect(tooltip).toContainText('Halo of Transcendence');
      await expect(pick).toHaveCSS('text-decoration-line', 'underline');
      await page.mouse.move(0, 0);
      await expect(tooltip).toHaveCount(0);
    }
    await pick.focus();
    await expect(tooltip).toContainText('Halo of Transcendence');
    await expect(tooltip).toBeVisible();
    await expect(pick).toHaveCSS('outline-style', 'solid');
    await expect(pick).toHaveCSS('outline-width', '2px');
    await expect(pick).toHaveCSS('outline-color', 'rgb(229, 185, 85)');
    expect(errors).toEqual([]);
  });

  test('desktop: the source is a column, the gaps one right edge, a row tint on hover (fix 5)', async ({
    page,
  }, info) => {
    await open(page);
    await settled(page);
    const rows = await rowBoxes(page);
    const metas = rows.flatMap((r) => (r.meta === null ? [] : [r]));
    for (const r of metas) {
      expect(Math.abs(r.meta!.left - metas[0].meta!.left), r.slot).toBeLessThan(1);
      expect(Math.abs(r.meta!.top - r.items!.top), r.slot).toBeLessThan(4);
    }
    const rights = rows.flatMap((r) => (r.gap === null || r.gap.right === r.gap.left ? [] : [r.gap.right]));
    for (const right of rights) expect(Math.abs(right - rights[0])).toBeLessThan(1);
    for (const r of rows) expect(r.height, r.slot).toBeGreaterThanOrEqual(44);
    await expect(row(page, 'head').getByTestId('character-gear-meta')).toHaveCSS('font-size', '13px');
    // A touch device has no hover; the tint is checked where there is a pointer.
    if (info.project.use.hasTouch !== true) {
      await row(page, 'head').getByTestId('character-gear-slot').hover();
      await expect(row(page, 'head')).toHaveCSS('background-color', 'rgba(255, 255, 255, 0.04)');
    }
    await noSidewaysScroll(page);
  });
});

test.describe('the item tooltip on a 390px phone (fix 1)', () => {
  test.use({ viewport: { width: 390, height: 844 } });

  test('opens on a tap, with no page error', async ({ page }, info) => {
    test.skip(info.project.use.hasTouch !== true, 'a tap needs a touch device (the mobile project)');
    const errors = await open(page);
    await settled(page);
    const pick = row(page, 'head').locator('[data-bis-item="16921"]');
    const tooltip = page.getByTestId('item-tooltip');
    await pick.tap();
    await expect(tooltip).toBeVisible();
    await expect(tooltip).toContainText('Halo of Transcendence');
    expect(errors).toEqual([]);
  });

  test('a second tap on the name closes it; so do a tap elsewhere and Escape (fix 2 ruling 2)', async ({
    page,
  }, info) => {
    test.skip(info.project.use.hasTouch !== true, 'a tap needs a touch device (the mobile project)');
    await open(page);
    await settled(page);
    const pick = row(page, 'head').locator('[data-bis-item="16921"]');
    const tooltip = page.getByTestId('item-tooltip');
    await pick.tap();
    await expect(tooltip).toBeVisible();
    await page.waitForTimeout(800);
    await pick.tap();
    await expect(tooltip).toHaveCount(0);
    await pick.tap();
    await expect(tooltip).toBeVisible();
    await page.getByTestId('character-gear-totals').tap();
    await expect(tooltip).toHaveCount(0);
    await pick.tap();
    await expect(tooltip).toBeVisible();
    await page.keyboard.press('Escape');
    await expect(tooltip).toHaveCount(0);
  });
});

for (const width of [390, 320]) {
  test.describe(`phone rows at ${width} (fix 4)`, () => {
    test.use({ viewport: { width, height: width === 390 ? 844 : 568 } });

    test('line 1 the label and the gap on one top, line 2 worn → pick, line 3 the source', async ({
      page,
    }) => {
      await open(page);
      await rowsDrawn(page);
      const rows = await rowBoxes(page);
      expect(rows).toHaveLength(17);
      for (const r of rows) {
        expect(r.height, r.slot).toBeGreaterThanOrEqual(44);
        expect(Math.abs(r.gap!.top - r.label!.top), r.slot).toBeLessThan(1);
        expect(r.items!.top, r.slot).toBeGreaterThanOrEqual(r.label!.bottom - 1);
      }
      const byslot = new Map(rows.map((r) => [r.slot, r]));
      // Three lines where there is a source: it sits under the items.
      for (const slot of ['head', 'neck', 'back', 'chest', 'legs', 'off_hand'])
        expect(byslot.get(slot)!.meta!.top, slot).toBeGreaterThanOrEqual(byslot.get(slot)!.items!.bottom - 1);
      // Best-for-band and no-pick rows: the note is a line of its own under the items, at their
      // left edge, at every phone width (fix 5 item 1).
      for (const slot of ['shoulder', 'trinket2', 'ranged']) {
        const r = byslot.get(slot)!;
        expect(r.meta!.top, slot).toBeGreaterThanOrEqual(r.items!.bottom - 1);
        expect(Math.abs(r.meta!.left - r.items!.left), slot).toBeLessThan(1);
      }
      // The gaps share one right edge.
      const rights = rows.flatMap((r) => (r.gap!.right === r.gap!.left ? [] : [r.gap!.right]));
      for (const right of rights) expect(Math.abs(right - rights[0])).toBeLessThan(1);
      await noSidewaysScroll(page);
    });
  });
}

for (const width of [1440, 1280, 1024]) {
  test.describe(`row rhythm at ${width} (fix 2 ruling 3)`, () => {
    test.use({ viewport: { width, height: 900 } });

    test('every row is 44px, the names column wide enough for the longest pair', async ({ page }) => {
      await open(page);
      await settled(page);
      const rows = await rowBoxes(page);
      expect(rows).toHaveLength(17);
      for (const r of rows) expect(r.height, r.slot).toBeCloseTo(44, 0);
      if (width >= 1280) {
        const column = await row(page, 'main_hand').evaluate(
          (li) => getComputedStyle(li).gridTemplateColumns.split(' ').map(parseFloat)[1],
        );
        expect(column).toBeGreaterThanOrEqual(440);
      }
    });
  });
}

/** A drawn separator, read by no screen reader: the dot glued by no-break spaces, alt text ""
 *  (fix 5 item 2). */
const DRAWN_DOT = '"\u00a0·\u00a0" / ""';

/** The rows with no source line, whose note is the meta line's whole text. */
const NOTE_ROWS = ['shoulder', 'hands', 'feet', 'finger1', 'finger2', 'trinket1', 'trinket2', 'ranged'];

// Fix round 4 item B, fix round 5 item 1: below 1024px a row's note is a line of its own with no
// dot at every width, so no dot can lead or trail a line; each dot left is drawn (a pseudo-element)
// between two texts that share its line. Checked on the fixture and with a long worn name on
// Shoulder, whose note could never share its line.
for (const width of [320, 360, 375, 390, 414]) {
  test.describe(`separators at ${width} (fix 2 ruling 4, fix 5 item 1)`, () => {
    test.use({ viewport: { width, height: width === 320 ? 568 : 844 } });

    for (const [fixture, gear] of [
      ['the fixture', GEAR],
      ['a long worn name', GEAR_LONG_NAME],
    ] as const) {
      test(`${fixture}: no text begins or ends with "·"; no drawn dot starts or ends a line`, async ({
        page,
      }) => {
        await open(page, { gear });
        await rowsDrawn(page);
        const found = await page.getByTestId('character-gear').evaluate((section) => {
          const texts: Text[] = [];
          const edged: string[] = [];
          const separators: Text[] = [];
          const walker = document.createTreeWalker(section, NodeFilter.SHOW_TEXT);
          for (let node = walker.nextNode(); node !== null; node = walker.nextNode()) {
            const text = (node as Text).data.trim();
            if (text === '') continue;
            texts.push(node as Text);
            // The one real separator between a source and its drop chance ("Skum · 50% drop", W.9)
            // is a text node of its own, aria-hidden, glued to both words by no-break spaces.
            const separator =
              text === '·' && (node as Text).parentElement?.getAttribute('aria-hidden') === 'true';
            if (separator) separators.push(node as Text);
            else if (text.startsWith('·') || text.endsWith('·')) edged.push(text);
          }
          /** The box of a text's first or last glyph. */
          const glyph = (node: Text, last: boolean): DOMRect => {
            const at = last
              ? node.data.trimEnd().length - 1
              : node.data.length - node.data.trimStart().length;
            const range = document.createRange();
            range.setStart(node, at);
            range.setEnd(node, at + 1);
            return range.getBoundingClientRect();
          };
          const oneLine = (a: DOMRect, b: DOMRect): boolean => a.top < b.bottom - 1 && b.top < a.bottom - 1;
          // Every dot drawn as a pseudo-element: the glyphs either side of it share its line.
          const drawn: { where: string; content: string; oneLine: boolean }[] = [];
          for (const element of section.querySelectorAll<HTMLElement>('*')) {
            if (getComputedStyle(element).visibility === 'hidden') continue;
            for (const pseudo of ['::before', '::after'] as const) {
              const content = getComputedStyle(element, pseudo).content;
              if (!content.includes('·')) continue;
              const inside = texts.filter((t) => element.contains(t));
              const outside = (flag: number) =>
                texts.filter(
                  (t) => !element.contains(t) && (element.compareDocumentPosition(t) & flag) !== 0,
                );
              const [left, right] =
                pseudo === '::before'
                  ? [outside(Node.DOCUMENT_POSITION_PRECEDING).at(-1), inside[0]]
                  : [inside.at(-1), outside(Node.DOCUMENT_POSITION_FOLLOWING)[0]];
              drawn.push({
                where: `${element.closest('li')?.dataset.testid ?? 'title'} ${element.dataset.testid}${pseudo}`,
                content,
                oneLine:
                  left !== undefined &&
                  right !== undefined &&
                  oneLine(glyph(left, true), glyph(right, false)),
              });
            }
          }
          // Each note-only row: the note's first line against the name's last, and its dot.
          const notes = [...section.querySelectorAll<HTMLElement>('li')].flatMap((li) => {
            const meta = li.querySelector<HTMLElement>('[data-testid="character-gear-meta"]');
            const items = li.querySelector<HTMLElement>('[data-testid="character-gear-items"]');
            if (meta === null || items === null) return [];
            if (
              !['Already your best at this level', 'No pick for your level range yet'].includes(
                meta.textContent?.trim() ?? '',
              )
            )
              return [];
            const name = [...items.getClientRects()].at(-1)!;
            const note = meta.getClientRects()[0];
            return [
              {
                slot: li.dataset.testid?.replace('character-gear-row-', '') ?? '',
                shareLine: oneLine(name, note),
                below: note.top >= name.bottom - 1,
                leftEdge: Math.abs(note.left - items.getBoundingClientRect().left) < 1,
                before: getComputedStyle(meta, '::before').content,
              },
            ];
          });
          // Each real separator: the glyphs either side of it share its line (it never starts or ends one).
          const separated = separators.map((node) => {
            const at = texts.indexOf(node);
            return (
              oneLine(glyph(texts[at - 1], true), glyph(node, false)) &&
              oneLine(glyph(node, true), glyph(texts[at + 1], false))
            );
          });
          return { edged, drawn, notes, separated };
        });
        expect(found.edged).toEqual([]);
        expect(found.separated.length).toBeGreaterThan(0);
        expect(found.separated.every(Boolean)).toBe(true);
        // The chances and the swap note: each dot glued to the texts either side, on their line.
        expect(found.drawn.length).toBeGreaterThan(0);
        for (const dot of found.drawn) {
          expect(dot.content, dot.where).toBe(DRAWN_DOT);
          expect(dot.oneLine, dot.where).toBe(true);
        }
        // A note's dot is drawn only when the note shares the name's line, and below 1024px it
        // never does: the note is its own line, at the items' left edge, with no dot.
        expect(found.notes.map((note) => note.slot)).toEqual(NOTE_ROWS);
        for (const note of found.notes) {
          expect(note, note.slot).toMatchObject({
            shareLine: false,
            below: true,
            leftEdge: true,
            before: 'none',
          });
        }
        if (gear === GEAR_LONG_NAME)
          await expect(row(page, 'shoulder').getByTestId('character-gear-worn')).toHaveText(
            'Mantle of the Unforgiving Majesty',
          );
        // At 480px and under the title's summary is its own block, with no dot after "Gear".
        const heading = page.getByTestId('character-gear-heading');
        expect(await heading.evaluate((element) => getComputedStyle(element, '::after').content)).toBe(
          'none',
        );
        await expect(page.getByTestId('character-gear-summary')).toHaveCSS('display', 'block');
      });
    }
  });
}

// Fix round 5 item 2: above 480px the title's dot after "Gear" is drawn too; no drawn dot is read
// by a screen reader, so the heading's name is its text alone.
test.describe('drawn dots at 1440 (fix 5 item 2)', () => {
  test.use({ viewport: { width: 1440, height: 900 } });

  test('each is drawn with alt text "", so the heading reads without its dot', async ({ page }) => {
    await open(page);
    await rowsDrawn(page);
    const pseudo = (slot: string | null, testid: string, which: '::before' | '::after') =>
      (slot === null ? page : row(page, slot))
        .getByTestId(testid)
        .evaluate((element, p) => getComputedStyle(element, p).content, which);
    expect(await pseudo(null, 'character-gear-heading', '::after')).toBe('"\u00a0·" / ""');
    // The chance is separated from its source by a real " · " in the text now (W.9), not a drawn dot.
    expect(await pseudo('head', 'character-gear-chance', '::before')).toBe('none');
    await expect(row(page, 'head').getByTestId('character-gear-meta')).toContainText('· 17% drop');
    expect(await pseudo('off_hand', 'character-gear-note', '::before')).toBe(DRAWN_DOT);
    await expect(page.getByTestId('character-gear-title')).toHaveAccessibleName(
      'Gear 6 of 15 slots best in slot · +70 points across 7 slots · 1 not scored',
    );
  });
});

// Fix round 2 ruling 8, fix round 3 item 2: mid-load the labels, gaps and sources draw, each name
// is held as its own text, unpainted, on a bar, so every row wraps exactly as it will once the
// names come; nothing below moves, at every width.
for (const [width, height] of [
  [1440, 900],
  [1024, 900],
  [390, 844],
  [320, 568],
] as const) {
  test.describe(`mid-load at ${width} (fix 2 ruling 8, fix 3 item 2)`, () => {
    test.use({ viewport: { width, height } });

    test('17 rows with labels, gaps and sources, the names held as bars; no row and not the footer moves', async ({
      page,
    }) => {
      await open(page, { itemsDelayMs: 2500 });
      const list = page.getByTestId('character-gear-list');
      await expect(list).toHaveAttribute('aria-busy', 'true');
      await expect(list.locator('[data-testid^="character-gear-row-"]')).toHaveCount(17);
      await expect(row(page, 'head').getByTestId('character-gear-slot')).toBeVisible();
      await expect(row(page, 'head').getByTestId('character-gear-gap')).toHaveText('+18.4 points');
      await expect(row(page, 'head').getByTestId('character-gear-chance')).toHaveText('17% drop');
      await expect(page.getByTestId('character-gear-pick')).toHaveCount(0);
      await expect(list.locator('[data-bis-item]')).toHaveCount(0);
      // The held name: a 14px bar, the wrapper's background, its name inside visibility-hidden.
      const bar = row(page, 'head').getByTestId('character-gear-skeleton').last();
      await expect(bar).toHaveCSS('height', '14px');
      await expect(bar).toHaveCSS('background-color', 'rgb(26, 29, 36)');
      await expect(bar.locator('span')).toHaveCSS('visibility', 'hidden');
      // The off hand's note line is reserved as a bar of the note's width (fix 4 item 4).
      const noteBar = row(page, 'off_hand').getByTestId('character-gear-note-bar');
      await expect(noteBar).toHaveCSS('background-color', 'rgb(26, 29, 36)');
      await expect(noteBar.getByTestId('character-gear-note')).toHaveCSS('visibility', 'hidden');
      const footer = page.getByTestId('character-gear-bis');
      await expect(footer).toHaveCSS('visibility', 'hidden');
      const measure = () =>
        page.getByTestId('character-gear').evaluate((section) => {
          const top = section.getBoundingClientRect().top;
          const rows = [...section.querySelectorAll('[data-testid^="character-gear-row-"]')].map((li) => {
            const r = li.getBoundingClientRect();
            return `${li.getAttribute('data-testid')?.slice(19)}:${(r.top - top).toFixed(1)}/${r.height.toFixed(1)}`;
          });
          const link = section.querySelector('[data-testid="character-gear-bis"]')!.getBoundingClientRect();
          return { rows, footer: link.top - top };
        });
      const before = await measure();
      await rowsDrawn(page);
      await expect(footer).toHaveCSS('visibility', 'visible');
      const after = await measure();
      expect(after.rows).toEqual(before.rows);
      expect(Math.abs(after.footer - before.footer)).toBeLessThan(1);
    });

    test('a failed file moves no row and not the footer link; its line comes under the link (fix 6 item 2)', async ({
      page,
    }) => {
      await open(page, { itemsDelayMs: 2500, itemsFailures: 1 });
      const list = page.getByTestId('character-gear-list');
      await expect(list).toHaveAttribute('aria-busy', 'true');
      await expect(list.locator('[data-testid^="character-gear-row-"]')).toHaveCount(17);
      const measure = () =>
        page.getByTestId('character-gear').evaluate((section) => {
          const top = section.getBoundingClientRect().top;
          const rows = [...section.querySelectorAll('[data-testid^="character-gear-row-"]')].map((li) => {
            const r = li.getBoundingClientRect();
            return `${li.getAttribute('data-testid')?.slice(19)}:${(r.top - top).toFixed(1)}/${r.height.toFixed(1)}`;
          });
          const link = section.querySelector('[data-testid="character-gear-bis"]')!.getBoundingClientRect();
          const list = section.querySelector('[data-testid="character-gear-list"]')!.getBoundingClientRect();
          return { rows, list: list.top - top, footer: link.top - top };
        });
      const before = await measure();
      await rowsDrawn(page);
      await expect(page.getByTestId('character-gear-failed')).toBeVisible();
      const after = await measure();
      expect(after.rows).toEqual(before.rows);
      expect(Math.abs(after.list - before.list)).toBeLessThan(1);
      expect(Math.abs(after.footer - before.footer)).toBeLessThan(1);
    });
  });
}

// Fix round 6 item 3: "main-hand swap" is kept on one line, so neither "main-" nor "swap" is left
// at a line's end or start, at every phone width.
for (const [width, height] of [
  [320, 568],
  [360, 740],
  [375, 812],
  [390, 844],
  [414, 896],
] as const) {
  test.describe(`the swap note at ${width} (fix 6 item 3)`, () => {
    test.use({ viewport: { width, height } });

    test('"main-hand swap" is one line box', async ({ page }) => {
      await open(page);
      await rowsDrawn(page);
      const note = row(page, 'off_hand').getByTestId('character-gear-note');
      await expect(note).toHaveText('comes with the main-hand swap');
      // The words' own line boxes, read with a Range over the text, whatever the markup.
      const kept = await note.evaluate((element) => {
        const range = document.createRange();
        const walker = document.createTreeWalker(element, NodeFilter.SHOW_TEXT);
        const nodes: Text[] = [];
        for (let n = walker.nextNode(); n !== null; n = walker.nextNode()) nodes.push(n as Text);
        const all = nodes.map((n) => n.data).join('');
        const start = all.indexOf('main-hand swap');
        const end = start + 'main-hand swap'.length;
        let seen = 0;
        for (const n of nodes) {
          if (start >= seen && start <= seen + n.data.length) range.setStart(n, start - seen);
          if (end >= seen && end <= seen + n.data.length) range.setEnd(n, end - seen);
          seen += n.data.length;
        }
        const tops = new Set(
          [...range.getClientRects()].filter((r) => r.width > 0).map((r) => Math.round(r.top)),
        );
        return { text: range.toString(), lines: tops.size };
      });
      expect(kept).toEqual({ text: 'main-hand swap', lines: 1 });
      await noSidewaysScroll(page);
    });
  });
}

// Fix round 3 item 1: a tooltip opens only from a real mouse move. A page that jumps under a
// parked pointer (an anchor landing, scrollBy) fires a pointerenter at the pointer's old place,
// sometimes before its scroll event: that never opens one.
test.describe('the parked pointer at 1440 (fix 3 item 1)', () => {
  test.use({ viewport: { width: 1440, height: 900 } });

  test('a landing and a scrollBy under a parked pointer open nothing; a 2px move does', async ({ page }) => {
    const errors = await open(page);
    await settled(page);
    const tooltip = page.getByTestId('item-tooltip');
    // Where Halo of Transcendence will sit once the section lands 96px under the top.
    const spot = await page.evaluate(() => {
      const section = document.querySelector('[data-testid="character-gear"]')!.getBoundingClientRect();
      const pick = document
        .querySelector('[data-testid="character-gear-row-head"] [data-bis-item="16921"]')!
        .getBoundingClientRect();
      return { x: pick.left + pick.width / 2, y: 96 + (pick.top + pick.height / 2 - section.top) };
    });
    await page.mouse.move(spot.x, spot.y);
    // Land by keyboard, so the pointer stays parked.
    const link = page.getByTestId('forever-score-most-to-gain').getByRole('link', { name: 'Gear' });
    await link.focus();
    await page.keyboard.press('Enter');
    await expect(page).toHaveURL(/#gear$/);
    const under = await page.evaluate(
      ({ x, y }) =>
        document.elementFromPoint(x, y)?.closest('[data-bis-item]')?.getAttribute('data-bis-item'),
      spot,
    );
    expect(under).toBe('16921');
    await page.waitForTimeout(600);
    await expect(tooltip).toHaveCount(0);
    // A scrollBy away and back, the pointer still parked: nothing opens.
    await page.evaluate(() => window.scrollBy(0, -300));
    await page.waitForTimeout(300);
    await page.evaluate(() => window.scrollBy(0, 300));
    await page.waitForTimeout(600);
    await expect(tooltip).toHaveCount(0);
    // A wheel scroll under the parked pointer: nothing opens either.
    await page.mouse.wheel(0, 44);
    await page.waitForTimeout(600);
    await expect(tooltip).toHaveCount(0);
    await page.mouse.wheel(0, -44);
    await page.waitForTimeout(600);
    await expect(tooltip).toHaveCount(0);
    // A real 2px move opens what the pointer rests on.
    await page.mouse.move(spot.x + 2, spot.y);
    await expect(tooltip).toBeVisible();
    await expect(tooltip).toContainText('Halo of Transcendence');
    expect(errors).toEqual([]);
  });
});

// Fix round 4 item C: hover styling lives in @media (hover: hover) and (pointer: fine), and the
// row's tint follows keyboard focus only, so a landing or a tap on a touch phone leaves nothing lit.
test.describe('no sticky hover on a touch phone at 390 (fix 4 item C)', () => {
  test.use({ viewport: { width: 390, height: 844 } });

  test('a landing by tap and a tap on a name leave no row tinted and no name underlined', async ({
    page,
  }, info) => {
    test.skip(info.project.use.hasTouch !== true, 'a tap needs a touch device (the mobile project)');
    await open(page);
    await settled(page);
    const lit = () =>
      page.getByTestId('character-gear').evaluate((section) => ({
        rows: [...section.querySelectorAll('[data-testid^="character-gear-row-"]')]
          .filter((li) => getComputedStyle(li).backgroundColor !== 'rgba(0, 0, 0, 0)')
          .map((li) => li.getAttribute('data-testid')),
        names: [...section.querySelectorAll('[data-bis-item]')]
          .filter((name) => getComputedStyle(name).textDecorationLine.includes('underline'))
          .map((name) => name.getAttribute('data-bis-item')),
      }));
    await page.getByTestId('forever-score-most-to-gain').getByRole('link', { name: 'Gear' }).tap();
    await expect(page).toHaveURL(/#gear$/);
    await page.waitForTimeout(400);
    expect(await lit()).toEqual({ rows: [], names: [] });
    await row(page, 'head').locator('[data-bis-item="16921"]').tap();
    await expect(page.getByTestId('item-tooltip')).toBeVisible();
    expect(await lit()).toEqual({ rows: [], names: [] });
  });
});

// Fix round 4 item 6: a tooltip dismissed by Escape (or a click) stays shut while the pointer
// rests on that name, whatever it moves; leaving the name and coming back opens it again.
test.describe('dismissal at 1440 (fix 4 item 6)', () => {
  test.use({ viewport: { width: 1440, height: 900 } });

  test('Escape, then a 1px move: still closed; leave and re-enter: open', async ({ page }) => {
    await open(page);
    await settled(page);
    const pick = row(page, 'head').locator('[data-bis-item="16921"]');
    const tooltip = page.getByTestId('item-tooltip');
    await pick.hover();
    await expect(tooltip).toBeVisible();
    const box = (await pick.boundingBox())!;
    await page.keyboard.press('Escape');
    await expect(tooltip).toHaveCount(0);
    await page.mouse.move(box.x + box.width / 2 + 1, box.y + box.height / 2);
    await page.waitForTimeout(300);
    await expect(tooltip).toHaveCount(0);
    await page.mouse.move(box.x + box.width / 2 + 3, box.y + box.height / 2 + 1);
    await page.waitForTimeout(300);
    await expect(tooltip).toHaveCount(0);
    // Leave the name, then come back.
    await page.mouse.move(box.x + box.width / 2, box.y - 60);
    await page.mouse.move(box.x + box.width / 2, box.y + box.height / 2);
    await expect(tooltip).toBeVisible();
    await expect(tooltip).toContainText('Halo of Transcendence');
  });
});

// Fix round 5 item 3: Escape after keyboard focus, the mouse elsewhere, holds nothing against the
// mouse: a hover on that name opens it. A dismissal while the mouse rests on the name (Escape, or a
// click that closes it) still holds until the pointer leaves.
test.describe('dismissal and the mouse at 1440 (fix 5 item 3)', () => {
  test.use({ viewport: { width: 1440, height: 900 } });

  test('Tab to a name, Escape, then hover it: open; Escape again on the hover holds', async ({ page }) => {
    await open(page);
    await settled(page);
    const worn = row(page, 'head').locator('[data-bis-item="10504"]');
    const pick = row(page, 'head').locator('[data-bis-item="16921"]');
    const tooltip = page.getByTestId('item-tooltip');
    await worn.focus();
    await page.keyboard.press('Tab');
    await expect(pick).toBeFocused();
    await expect(tooltip).toContainText('Halo of Transcendence');
    await page.keyboard.press('Escape');
    await expect(tooltip).toHaveCount(0);
    await pick.hover();
    await expect(tooltip).toBeVisible();
    await expect(tooltip).toContainText('Halo of Transcendence');
    // Now the mouse rests on it: Escape holds against the 1px and 3px moves.
    const box = (await pick.boundingBox())!;
    await page.keyboard.press('Escape');
    await expect(tooltip).toHaveCount(0);
    await page.mouse.move(box.x + box.width / 2 + 1, box.y + box.height / 2);
    await page.waitForTimeout(300);
    await expect(tooltip).toHaveCount(0);
    await page.mouse.move(box.x + box.width / 2 + 3, box.y + box.height / 2 + 1);
    await page.waitForTimeout(300);
    await expect(tooltip).toHaveCount(0);
  });

  test('a click that closes a name holds it shut until the pointer leaves', async ({ page }) => {
    await open(page);
    await settled(page);
    const pick = row(page, 'head').locator('[data-bis-item="16921"]');
    const tooltip = page.getByTestId('item-tooltip');
    await pick.hover();
    await expect(tooltip).toBeVisible();
    const box = (await pick.boundingBox())!;
    const [x, y] = [box.x + box.width / 2, box.y + box.height / 2];
    await page.mouse.click(x, y);
    await expect(tooltip).toHaveCount(0);
    await page.mouse.move(x + 1, y);
    await page.waitForTimeout(300);
    await expect(tooltip).toHaveCount(0);
    await page.mouse.move(x + 3, y + 1);
    await page.waitForTimeout(300);
    await expect(tooltip).toHaveCount(0);
    await page.mouse.move(x, box.y - 60);
    await page.mouse.move(x, y);
    await expect(tooltip).toBeVisible();
    await expect(tooltip).toContainText('Halo of Transcendence');
  });
});
