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
import { expect, test, type Page } from '@playwright/test';
import {
  CHARACTER,
  CHARACTER_PATH,
  CHARACTER_URL,
  GEAR,
  GEAR_ITEMS,
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
  /** Hold the items file this long before answering. */
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
    if (failures > 0) {
      failures -= 1;
      return route.fulfill({ status: 500, body: '' });
    }
    if (options.itemsDelayMs !== undefined) await new Promise((r) => setTimeout(r, options.itemsDelayMs));
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
        'Gear +70 pts across 7 slots · 1 not scored',
      );
      const totals = page.getByTestId('character-gear-totals');
      await expect(totals).toHaveText(
        'The card’s 88 counts the band’s best set. These 70 count the picks you can wear.',
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
      await expect(head.getByTestId('character-gear-meta')).toHaveText("Onyxia's Lair: Onyxia · 17% drop");
      await expect(head.getByTestId('character-gear-gap')).toHaveText('+18.4 pts');
      await expect(head.getByTestId('character-gear-gap')).toHaveCSS('font-weight', '700');
      await expect(row(page, 'wrist').getByTestId('character-gear-gap')).toHaveCSS('font-weight', '400');
      // Icons: a 20px box, the pick's border its quality, the image loaded in it.
      const box = head.locator('[data-bis-item="16921"] [data-testid="character-gear-icon"]');
      await expect(box).toHaveCSS('width', '20px');
      await expect(box).toHaveCSS('height', '20px');
      await expect(box).toHaveCSS('border-top-color', 'rgb(163, 53, 238)');
      expect(await box.locator('img').evaluate((img: HTMLImageElement) => img.naturalWidth)).toBeGreaterThan(
        0,
      );
      await expect(row(page, 'back').getByTestId('character-gear-meta')).toContainText('<1% drop');
      await expect(row(page, 'chest').getByTestId('character-gear-meta')).toHaveText('Tailoring');
      await expect(row(page, 'shoulder').getByTestId('character-gear-meta')).toContainText(
        'Best for your band',
      );
      await expect(row(page, 'shoulder').getByTestId('character-gear-gap')).toHaveText('');
      await expect(row(page, 'finger1').getByTestId('character-gear-slot')).toHaveText('Ring 1');
      await expect(row(page, 'trinket2').getByTestId('character-gear-slot')).toHaveText('Trinket 2');
      await expect(row(page, 'trinket2').getByTestId('character-gear-meta')).toHaveText(
        'No pick for your band yet',
      );
      await expect(row(page, 'ranged').getByTestId('character-gear-meta')).toHaveText(
        'No pick for your band yet',
      );
      const legs = row(page, 'legs');
      await expect(legs.getByTestId('character-gear-worn')).toHaveText('Unknown item');
      await expect(legs.getByTestId('character-gear-worn')).toHaveCSS('color', RGB.muted);
      await expect(legs.getByTestId('character-gear-gap')).toHaveText('Not scored');
      await expect(legs.getByTestId('character-gear-gap')).toHaveCSS('color', RGB.muted);
      const off = row(page, 'off_hand');
      await expect(off.getByTestId('character-gear-meta')).toHaveText(
        'Blackrock Spire: Gyth · 9% drop · comes with the main-hand swap',
      );
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
      await expect(line).toHaveText('No gear on file. Upload a log or sync the addon.');
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
      await expect(box).toHaveCSS('background-color', 'rgb(18, 21, 28)');
      await expect(box).toHaveCSS('border-top-color', 'rgb(163, 53, 238)');
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
      // The swap cannot be known: the empty off hand shows no figure (fix 2 ruling 7).
      await expect(row(page, 'off_hand').getByTestId('character-gear-gap')).toHaveText('');
      await page.reload();
      await rowsDrawn(page);
      await expect(row(page, 'head').getByTestId('character-gear-pick')).toHaveCSS('color', RGB.epic);
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
      // Best-for-band and no-pick rows: the note follows the items on their line at 390 (two
      // lines); at 320 a long name can push it under them.
      for (const slot of ['shoulder', 'trinket2', 'ranged']) {
        const r = byslot.get(slot)!;
        if (width === 390) expect(Math.abs(r.meta!.top - r.items!.top), slot).toBeLessThan(6);
        else expect(r.meta!.top, slot).toBeGreaterThanOrEqual(r.items!.top - 1);
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

for (const width of [390, 320]) {
  test.describe(`separators at ${width} (fix 2 ruling 4)`, () => {
    test.use({ viewport: { width, height: width === 390 ? 844 : 568 } });

    test('no text begins with "·": each dot trails the word before it', async ({ page }) => {
      await open(page);
      await rowsDrawn(page);
      const found = await page.getByTestId('character-gear').evaluate((section) => {
        const leading: string[] = [];
        const walker = document.createTreeWalker(section, NodeFilter.SHOW_TEXT);
        for (let node = walker.nextNode(); node !== null; node = walker.nextNode()) {
          const text = node.textContent ?? '';
          if (/^[ \t\r\n]*·/.test(text)) leading.push(text);
        }
        const afters: string[] = [];
        for (const element of section.querySelectorAll('*')) {
          const content = getComputedStyle(element, '::after').content;
          if (content.includes('·')) afters.push(content);
        }
        return { leading, afters };
      });
      expect(found.leading).toEqual([]);
      // Every drawn dot is a no-break space then the dot, so it stays on its word's line.
      for (const content of found.afters) expect(content.charCodeAt(1), content).toBe(160);
      expect(found.afters.length).toBeGreaterThan(0);
      // At 480px and under the title's summary is its own block, with no dot after "Gear".
      const heading = page.getByTestId('character-gear-heading');
      expect(await heading.evaluate((element) => getComputedStyle(element, '::after').content)).toBe('none');
      await expect(page.getByTestId('character-gear-summary')).toHaveCSS('display', 'block');
    });
  });
}

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
      await expect(row(page, 'head').getByTestId('character-gear-gap')).toHaveText('+18.4 pts');
      await expect(row(page, 'head').getByTestId('character-gear-meta')).toHaveText(
        "Onyxia's Lair: Onyxia · 17% drop",
      );
      await expect(page.getByTestId('character-gear-pick')).toHaveCount(0);
      await expect(list.locator('[data-bis-item]')).toHaveCount(0);
      // The held name: its own text, unpainted, on a 14px bar.
      const bar = row(page, 'head').getByTestId('character-gear-skeleton').last();
      await expect(bar).toHaveCSS('height', '14px');
      await expect(bar).toHaveCSS('background-color', 'rgb(26, 29, 36)');
      await expect(bar).toHaveCSS('color', 'rgba(0, 0, 0, 0)');
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
