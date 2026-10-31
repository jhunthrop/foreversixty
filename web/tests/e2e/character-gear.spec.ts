// web/tests/e2e/character-gear.spec.ts
// The profile's Gear section as the shared paperdoll (design/specs/2026-10-31-paperdoll.md, call
// site 2), from the character answer's `gear` and `gear_gap_total` (api/openapi.yaml
// `CharacterGear`), stubbed with every slot state (src/fixtures/character-gear.ts), the build's
// items/priest.json reduced to the fixture's real rows, and the real build's icons.
//   The frame: every slot, quality colours, gains, "Best in slot", "Not scored", the unknown
//   item's slot word, the two-hander's empty off hand, no item level with an unknown item.
//   The tooltip on hover, focus and tap, toward the centre from the right column (§4.2).
//   The worst case: no name clamped at 360, 390 and 1024; a long word breaks with a hyphen.
//   The landing focuses the heading, with gear and without.
import { expect, test, type Page } from '@playwright/test';
import {
  CHARACTER,
  CHARACTER_PATH,
  CHARACTER_URL,
  GEAR,
  GEAR_ITEMS,
  GEAR_LONGEST,
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

/** Waits for the frame to be drawn (the items file in, or failed). */
async function drawn(page: Page): Promise<void> {
  await expect(page.getByTestId('paperdoll-slot-head')).not.toHaveAttribute('data-state', 'loading');
}

/** Waits for the page to settle, so nothing above the section moves a cell from under the pointer. */
async function settled(page: Page): Promise<void> {
  await drawn(page);
  await page.waitForLoadState('networkidle');
}

const slot = (page: Page, key: string) => page.getByTestId(`paperdoll-slot-${key}`);

async function noSidewaysScroll(page: Page): Promise<void> {
  const scroll = await page.evaluate(() => ({
    width: document.documentElement.scrollWidth,
    client: document.documentElement.clientWidth,
  }));
  expect(scroll.width).toBeLessThanOrEqual(scroll.client);
}

for (const width of [1440, 390]) {
  test.describe(`the Gear paperdoll at ${width}`, () => {
    test.use({ viewport: { width, height: width === 1440 ? 900 : 844 } });

    test('draws every slot state in the frame, with no page error', async ({ page }) => {
      const errors = await open(page);
      await drawn(page);
      await expect(page.getByTestId('character-gear-title')).toHaveText('Gear');
      await expect(page.getByTestId('character-gear-summary')).toHaveText('+70 points across 7 slots');
      await expect(page.getByTestId('character-gear-summary')).toHaveCSS('color', 'rgb(229, 185, 85)');
      const head = slot(page, 'head');
      await expect(head.getByTestId('paperdoll-item-head')).toHaveAttribute('aria-label', 'Head: Green Lens');
      await expect(head.locator('.pd-name')).toHaveCSS('color', RGB.rare);
      await expect(slot(page, 'shoulder')).toContainText('Best in slot');
      await expect(slot(page, 'legs')).toContainText(/Legs · Item \d+/);
      await expect(slot(page, 'legs').getByTestId('paperdoll-unknown')).toHaveCSS('color', RGB.muted);
      await expect(slot(page, 'off_hand')).toHaveAttribute('data-state', 'two-hand');
      await expect(slot(page, 'off_hand')).toContainText('Two-hand in main hand');
      await expect(slot(page, 'trinket2')).toHaveAttribute('data-state', 'empty');
      await expect(page.getByTestId('paperdoll-ilvl')).toHaveCount(0);
      await expect(page.getByTestId('paperdoll-note')).toHaveText(
        '1 item is not in our item data yet. It is new in the beta, and it gets a name and icon when the client lists it. Item level shows once every item is known.',
      );
      await expect(page.getByTestId('character-gear-bis')).toHaveAttribute(
        'href',
        '/bis/priest/discipline?faction=alliance#band-alliance-60',
      );
      await noSidewaysScroll(page);
      expect(errors).toEqual([]);
    });

    test('names the list "Gear by slot"; every target is at least 44px', async ({ page }) => {
      await open(page);
      await drawn(page);
      await expect(page.getByRole('list', { name: 'Gear by slot' })).toBeVisible();
      const targets = page.locator('.pd-known:visible, .pd-up:visible');
      const heights = await targets.evaluateAll((els) => els.map((el) => el.getBoundingClientRect().height));
      expect(heights.length).toBeGreaterThan(0);
      for (const height of heights) expect(height).toBeGreaterThanOrEqual(44);
    });

    test('lands the Gear link on the section and focuses its heading', async ({ page }) => {
      await open(page);
      await settled(page);
      const link = page.getByTestId('forever-score-most-to-gain').getByRole('link', { name: 'Gear' });
      await expect(link).toHaveAttribute('href', '#gear');
      await link.click();
      await expect(page).toHaveURL(/#gear$/);
      await expect(page.getByTestId('character-gear-title')).toBeFocused();
      const top = await page
        .getByTestId('character-gear')
        .evaluate((element) => element.getBoundingClientRect().top);
      expect(top).toBeGreaterThanOrEqual(95);
      expect(top).toBeLessThan(110);
    });

    test('`gear: []`: the heading over one line; the link lands there and focuses the heading', async ({
      page,
    }) => {
      await open(page, { gear: { gear: [], gear_gap_total: null } });
      await expect(page.getByTestId('character-gear')).toHaveCount(0);
      const heading = page.getByTestId('character-gear-title');
      await expect(heading).toHaveText('Gear');
      await page.getByTestId('forever-score-most-to-gain').getByRole('link', { name: 'Gear' }).click();
      await expect(page).toHaveURL(/#gear$/);
      await expect(heading).toBeFocused();
    });

    test('never paints a name white while the items file is slow, and shimmers meanwhile', async ({
      page,
    }) => {
      await open(page, { itemsDelayMs: 1500 });
      await expect(slot(page, 'head')).toHaveAttribute('data-state', 'loading');
      await expect(page.getByTestId('paperdoll-name')).toHaveText('Elyra Duskvale');
      await drawn(page);
      await expect(slot(page, 'head').locator('.pd-name')).toHaveCSS('color', RGB.rare);
    });

    test('a failed items file draws plain names and says so once', async ({ page }) => {
      await open(page, { itemsFailures: 1 });
      await drawn(page);
      await expect(slot(page, 'head').locator('.pd-name')).toHaveCSS('color', RGB.body);
      await expect(page.getByTestId('paperdoll-note')).toHaveText(
        'Item icons and colours did not load. Reload the page to try again.',
      );
    });
  });
}

test.describe('the item tooltip at 1440', () => {
  test.use({ viewport: { width: 1440, height: 900 } });

  test('opens on hover and keyboard focus for an item and a pick; Escape closes it', async ({ page }) => {
    const errors = await open(page);
    await settled(page);
    await page.getByTestId('paperdoll-item-head').hover();
    await expect(page.getByTestId('item-tooltip')).toBeVisible();
    await expect(page.getByTestId('item-tooltip')).toContainText('Green Lens');
    await page.getByTestId('paperdoll-pick-head').hover();
    await expect(page.getByTestId('item-tooltip')).toContainText('Halo of Transcendence');
    await page.mouse.move(0, 0);
    await page.getByTestId('paperdoll-item-hands').focus();
    await expect(page.getByTestId('item-tooltip')).toBeVisible();
    await page.keyboard.press('Escape');
    await expect(page.getByTestId('item-tooltip')).toHaveCount(0);
    expect(errors).toEqual([]);
  });

  test('the right column opens toward the centre and never covers the crest (§4.2)', async ({ page }) => {
    await open(page);
    await settled(page);
    const cell = page.getByTestId('paperdoll-item-hands');
    await cell.hover();
    const tip = page.getByTestId('item-tooltip');
    await expect(tip).toBeVisible();
    const [cellBox, tipBox, crest] = await Promise.all([
      cell.boundingBox(),
      tip.boundingBox(),
      page.locator('.pd-crest').boundingBox(),
    ]);
    expect(Math.abs(cellBox!.x + cellBox!.width - (tipBox!.x + tipBox!.width))).toBeLessThan(2);
    expect(tipBox!.x).toBeGreaterThanOrEqual(crest!.x + crest!.width);
  });

  test('Tab walks the frame in reading order, picks after their cells', async ({ page }) => {
    await open(page);
    await settled(page);
    await page.getByTestId('paperdoll-item-head').focus();
    const order: string[] = [];
    for (let step = 0; step < 5; step++) {
      order.push((await page.evaluate(() => document.activeElement?.getAttribute('data-testid'))) ?? '');
      await page.keyboard.press('Tab');
    }
    // Head and its pick, Hands (worn as the pick: no pick to visit), then the empty Neck's pick
    // and Waist: the frame's pairs in reading order.
    expect(order).toEqual([
      'paperdoll-item-head',
      'paperdoll-pick-head',
      'paperdoll-item-hands',
      'paperdoll-pick-neck',
      'paperdoll-item-waist',
    ]);
  });
});

test.describe('the item tooltip on a 390px phone', () => {
  test.use({ viewport: { width: 390, height: 844 }, hasTouch: true, isMobile: true });

  test('a tap on an upgrade row opens the pick; a tap elsewhere and Escape close it', async ({ page }) => {
    const errors = await open(page);
    await settled(page);
    const row = page.getByTestId('paperdoll-uprow-head');
    await expect(row).toHaveAttribute(
      'aria-label',
      'Head upgrade: Halo of Transcendence, 18.4 points to gain',
    );
    await row.tap();
    const tip = page.getByTestId('item-tooltip');
    await expect(tip).toBeVisible();
    await expect(tip).toContainText('Halo of Transcendence');
    const box = await tip.boundingBox();
    expect(box!.x).toBeGreaterThanOrEqual(15);
    expect(box!.x + box!.width).toBeLessThanOrEqual(390 - 15);
    await row.tap();
    await expect(tip).toHaveCount(0);
    await page.getByTestId('paperdoll-item-head').tap();
    await expect(tip).toBeVisible();
    await page.keyboard.press('Escape');
    await expect(tip).toHaveCount(0);
    expect(errors).toEqual([]);
  });
});

for (const width of [360, 390, 1024]) {
  test.describe(`the worst case at ${width} (§3.3)`, () => {
    test.use({ viewport: { width, height: 900 } });

    test('no name is clamped, and nothing scrolls sideways', async ({ page }) => {
      await open(page, { gear: GEAR_LONGEST });
      await drawn(page);
      const clipped = await page
        .locator('.pd-name')
        .evaluateAll((els) =>
          els.filter((el) => el.scrollHeight > el.clientHeight + 1).map((el) => el.textContent),
        );
      expect(clipped).toEqual([]);
      await noSidewaysScroll(page);
    });
  });
}

test.describe('hyphenation at 360', () => {
  test.use({ viewport: { width: 360, height: 900 } });

  test('a word longer than its line breaks at a soft hyphen, drawn as "-"', async ({ page }) => {
    await open(page, { gear: GEAR_LONGEST });
    await drawn(page);
    const name = slot(page, 'head').locator('.pd-name');
    await expect(name).toContainText('Elec\u00adtro\u00admag\u00adne\u00adtic');
    // The word wraps inside itself: more lines than words would need without a break.
    const lines = await name.evaluate((el) => {
      const range = document.createRange();
      range.selectNodeContents(el);
      return new Set([...range.getClientRects()].map((rect) => Math.round(rect.top))).size;
    });
    expect(lines).toBeGreaterThanOrEqual(3);
  });
});
