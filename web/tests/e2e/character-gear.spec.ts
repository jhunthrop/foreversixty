// web/tests/e2e/character-gear.spec.ts
// The character page's Gear Section (design/specs/2026-11-06-character.md C1, C4): one Plate with
// the shared Paperdoll in `worn` mode, from the character answer's `gear` (api/openapi.yaml
// `CharacterGear`), stubbed with every slot state (src/fixtures/character-gear.ts), the build's
// items/priest.json reduced to the fixture's real rows, the spec's BiS file stubbed with the band
// the gains are priced from, and the real build's icons.
//   The doll: every slot, quality colours, item level on line 2, the owner's gains in DPS against
//   the pick ("Best in slot", "Not scored", "Not sim-checked"), none for a visitor, the unknown
//   item's slot word, the two-hander over the off hand. The Inspector in place below 1760 and at
//   the side from 1760. The tooltip on hover, focus and tap, for the worn item and the pick. The
//   keyboard. The worst case: nothing clamped or sideways at 360, 390, 1024.
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import {
  CHARACTER,
  CHARACTER_PATH,
  CHARACTER_URL,
  GEAR,
  GEAR_BIS,
  GEAR_ITEMS,
  GEAR_LONGEST,
  OWNER_SCORE,
  routeGearIcons,
  routeOwner,
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

const BIS_HREF = '/bis/priest/discipline/60';
const OWNER_BIS_HREF = '/bis/priest/holy/60';

const slot = (page: Page, key: string) => page.getByTestId(`character-doll-slot-${key}`);

const RGB = {
  rare: 'rgb(61, 148, 240)',
};

interface OpenOptions {
  gear?: unknown;
  /** Sign the character's owner in (the cells carry gains); a visitor otherwise. */
  owner?: boolean;
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
  if (options.owner ?? true) await routeOwner(page);
  await page.route(CHARACTER_URL, (route) =>
    route.fulfill(ok({ ...CHARACTER.data, ...OWNER_SCORE, ...((options.gear ?? GEAR) as object) })),
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
  await page.route('**/data/*/bis/priest-holy.json', (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(GEAR_BIS) }),
  );
  await routeGearIcons(page, options.missingIcons);
  await page.goto(CHARACTER_PATH);
  await expect(page.getByTestId('character-score')).toBeVisible();
  return errors;
}

/** Waits for the doll to be drawn (the items file and the band in, or failed). */
async function drawn(page: Page): Promise<void> {
  await expect(page.getByTestId('character-doll')).not.toHaveAttribute('aria-busy', 'true');
  await expect(slot(page, 'head')).toHaveAttribute('data-state', 'item');
}

/** Waits for the page to settle, so nothing above the section moves a cell from under the pointer. */
async function settled(page: Page): Promise<void> {
  await drawn(page);
  await page.waitForLoadState('networkidle');
}

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

    test('the owner: every slot state, item level on line 2, gains in DPS, no page error', async ({
      page,
    }) => {
      const errors = await open(page);
      await drawn(page);
      const section = page.getByTestId('character-gear-section');
      await expect(section.getByRole('heading', { name: 'Gear' })).toBeVisible();
      await expect(section).toContainText('Level 60 raid');
      await expect(section.getByRole('link', { name: 'Best in slot ›' })).toHaveAttribute(
        'href',
        OWNER_BIS_HREF,
      );
      const head = slot(page, 'head');
      await expect(head.getByTestId('item-hover-10504')).toHaveText('Green Lens');
      await expect(head.getByTestId('item-hover-10504')).toHaveCSS('color', RGB.rare);
      await expect(page.getByTestId('character-doll-line2-head')).toHaveText('Item level 49');
      await expect(head).toContainText('+1.8');
      await expect(head).toContainText('DPS');
      await expect(page.getByTestId('character-doll-basis-head')).toHaveText('vs Halo of Transcendence');
      await expect(slot(page, 'neck')).toContainText('+2.0');
      await expect(slot(page, 'wrist')).toContainText('Not sim-checked');
      await expect(slot(page, 'shoulder')).toContainText('Best in slot');
      await expect(slot(page, 'legs')).toContainText(/Legs · Item \d+/);
      await expect(slot(page, 'legs')).toHaveAttribute('data-state', 'unknown');
      await expect(slot(page, 'legs')).toContainText('Not in our item data yet');
      await expect(slot(page, 'legs')).toContainText('Not scored');
      await expect(slot(page, 'off_hand')).toContainText('Two-hand in main hand');
      await expect(slot(page, 'off_hand').locator('[data-covered="true"]')).toHaveCount(1);
      await expect(slot(page, 'trinket2')).toHaveAttribute('data-state', 'empty');
      await expect(page.getByTestId('character-doll-note')).toHaveText(
        '1 item is not in our item data yet. It is new in the beta, and it gets a name and icon when the client lists it. Item level shows once every item is known.',
      );
      // The head answer is the sum of the priced gains.
      await expect(page.getByTestId('character-answer')).toContainText('+3.8');
      await expect(page.getByTestId('character-answer')).toContainText('Upgrades available');
      await expect(section.getByText(/points/)).toHaveCount(0);
      await noSidewaysScroll(page);
      expect(errors).toEqual([]);
    });

    test('a visitor: the same worn items, and no cell carries a gain', async ({ page }) => {
      const errors = await open(page, { owner: false });
      await drawn(page);
      await expect(slot(page, 'head').getByTestId('item-hover-10504')).toHaveText('Green Lens');
      await expect(page.getByTestId('character-doll-line2-head')).toHaveText('Item level 49');
      await expect(page.locator('[data-testid^="character-doll-basis-"]')).toHaveCount(0);
      await expect(slot(page, 'head')).not.toContainText('+1.8');
      await expect(slot(page, 'wrist')).not.toContainText('Not sim-checked');
      await expect(
        page.getByTestId('character-gear-section').getByRole('link', { name: 'Best in slot ›' }),
      ).toHaveAttribute('href', BIS_HREF);
      expect(errors).toEqual([]);
    });

    test('names the list "Gear by slot"; every cell is at least 44px', async ({ page }) => {
      await open(page);
      await drawn(page);
      await expect(page.getByRole('listbox', { name: 'Gear by slot' })).toBeVisible();
      const heights = await page
        .locator('[data-testid^="character-doll-slot-"]')
        .evaluateAll((els) => els.map((el) => el.getBoundingClientRect().height));
      expect(heights.length).toBe(17);
      for (const height of heights) expect(height).toBeGreaterThanOrEqual(44);
    });

    test('a click on a cell opens the Inspector in place; a second click closes it', async ({ page }) => {
      await open(page);
      await settled(page);
      await slot(page, 'head').click();
      const inspector = page.getByTestId('character-doll-detail-head').getByTestId('character-inspector');
      await expect(inspector).toBeVisible();
      await expect(inspector).toContainText('Best in slot');
      const pick = inspector.getByTestId('character-inspector-pick');
      await expect(pick).toContainText('Halo of Transcendence');
      await expect(pick).toContainText("Onyxia's Lair: Onyxia · 17% drop");
      // The band's runner-up, with its delta against the pick.
      await expect(inspector).toContainText('Green Lens');
      await expect(inspector).toContainText('−1.8');
      await expect(inspector.getByRole('link', { name: 'Open in best in slot ›' })).toHaveAttribute(
        'href',
        OWNER_BIS_HREF,
      );
      await expect(page.getByTestId('character-doll-inspector')).toHaveCount(0);
      await slot(page, 'head').click();
      await expect(page.getByTestId('character-inspector')).toHaveCount(0);
      await noSidewaysScroll(page);
    });

    test('`gear: []`: no Gear section, the rest of the page still draws', async ({ page }) => {
      await open(page, { gear: { gear: [], gear_gap_total: null } });
      await expect(page.getByTestId('character-gear-section')).toHaveCount(0);
      await expect(page.getByTestId('character-doll')).toHaveCount(0);
      await expect(page.getByTestId('character-answer')).toHaveCount(0);
      await expect(page.getByTestId('character-best')).toBeVisible();
    });

    test('never paints a name white while the items file is slow: skeleton rows meanwhile', async ({
      page,
    }) => {
      await open(page, { itemsDelayMs: 1500 });
      await expect(page.getByTestId('character-doll')).toHaveAttribute('aria-busy', 'true');
      await expect(page.getByTestId('item-hover-10504')).toHaveCount(0);
      await drawn(page);
      await expect(slot(page, 'head').getByTestId('item-hover-10504')).toHaveCSS('color', RGB.rare);
    });

    test('a failed items file keeps the names and says so once', async ({ page }) => {
      await open(page, { itemsFailures: 1 });
      await drawn(page);
      await expect(slot(page, 'head')).toContainText('Green Lens');
      await expect(page.getByTestId('character-doll-note')).toHaveText(
        'Item icons and colours did not load. Reload the page to try again.',
      );
    });
  });
}

test.describe('the Inspector at 1800', () => {
  test.use({ viewport: { width: 1800, height: 1000 } });

  test('sits at the side with a slot already shown, and follows the chosen cell', async ({ page }) => {
    await open(page);
    await settled(page);
    const side = page.getByTestId('character-doll-inspector');
    await expect(side.getByTestId('character-inspector')).toBeVisible();
    await slot(page, 'head').click();
    await expect(slot(page, 'head')).toHaveAttribute('aria-selected', 'true');
    await expect(side).toContainText('Green Lens');
    await expect(side.getByTestId('character-inspector-pick')).toContainText('Halo of Transcendence');
    await expect(page.locator('[data-testid^="character-doll-detail-"]')).toHaveCount(0);
    await noSidewaysScroll(page);
  });
});

test.describe('the item tooltip at 1440', () => {
  test.use({ viewport: { width: 1440, height: 900 } });

  test('opens on hover for the worn item and the pick, on keyboard focus; Escape closes it', async ({
    page,
  }) => {
    const errors = await open(page);
    await settled(page);
    const tip = page.getByTestId('item-tooltip');
    await page.getByTestId('item-hover-10504').hover();
    await expect(tip).toBeVisible();
    await expect(tip).toContainText('Green Lens');
    await page.getByTestId('character-doll-basis-head').locator('[data-bis-item="16921"]').hover();
    await expect(tip).toContainText('Halo of Transcendence');
    await page.mouse.move(0, 0);
    await slot(page, 'hands').focus();
    await expect(tip).toBeVisible();
    await page.keyboard.press('Escape');
    await expect(tip).toHaveCount(0);
    expect(errors).toEqual([]);
  });

  test('the keyboard walks the doll in reading order: one tab stop, arrows, Home and End', async ({
    page,
  }) => {
    await open(page);
    await settled(page);
    await slot(page, 'head').focus();
    const focused = () => page.evaluate(() => document.activeElement?.getAttribute('data-testid'));
    await page.keyboard.press('ArrowDown');
    // Neck is empty, so not focusable: down from head is shoulder.
    expect(await focused()).toBe('character-doll-slot-shoulder');
    // Across to the same row of the right column: shoulder is the third row, so legs.
    await page.keyboard.press('ArrowRight');
    expect(await focused()).toBe('character-doll-slot-legs');
    await page.keyboard.press('End');
    expect(await focused()).toBe('character-doll-slot-trinket1');
    await page.keyboard.press('Home');
    expect(await focused()).toBe('character-doll-slot-head');
  });
});

test.describe('the item tooltip on a 390px phone', () => {
  test.use({ viewport: { width: 390, height: 844 }, hasTouch: true, isMobile: true });

  test('a tap on the pick opens it inside the screen; a second tap and Escape close', async ({ page }) => {
    const errors = await open(page);
    await settled(page);
    const pick = page.getByTestId('character-doll-basis-head').locator('[data-bis-item="16921"]');
    await pick.tap();
    const tip = page.getByTestId('item-tooltip');
    await expect(tip).toBeVisible();
    await expect(tip).toContainText('Halo of Transcendence');
    const box = await tip.boundingBox();
    expect(box!.x).toBeGreaterThanOrEqual(15);
    expect(box!.x + box!.width).toBeLessThanOrEqual(390 - 15);
    await pick.tap();
    await expect(tip).toHaveCount(0);
    await page.getByTestId('item-hover-10504').tap();
    await expect(tip).toContainText('Green Lens');
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
        .locator('[data-testid^="item-hover-"]')
        .evaluateAll((els) =>
          els.filter((el) => el.scrollWidth > el.clientWidth + 1).map((el) => el.textContent),
        );
      expect(clipped).toEqual([]);
      await noSidewaysScroll(page);
    });
  });
}
