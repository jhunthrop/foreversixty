// web/tests/e2e/character-gear.spec.ts
// The character page's Gear panel (round 5, directions mock b2-character; character spec
// 2026-11-06 C1, C4): the page's framed panel holding the shared PaperdollFrame on the owner's
// frame, from the character answer's `gear` (api/openapi.yaml `CharacterGear`), stubbed with every
// slot state (src/fixtures/character-gear.ts), the build's items/priest.json reduced to the
// fixture's real rows, the spec's BiS file stubbed with the band the gains are priced from, and the
// real build's icons.
//   The frame: each worn item in its socket, named for assistive tech; the two-hander greyed over
//   the off hand; the title bar (spec and level). The centre: the Upgrades (the band's picks by
//   their priced gain) until a slot is chosen, then that slot (the pick with its source, the
//   runner-up) and back. A phone: the list form, the Upgrades under it. The tooltip on hover. The
//   worst case: nothing clamped or sideways at 360, 390, 1024.
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

const frameSlot = (page: Page, key: string) => page.getByTestId(`paperdoll-frame-slot-${key}`);
const listSlot = (page: Page, key: string) => page.getByTestId(`paperdoll-list-slot-${key}`);

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
  await expect(page.getByTestId('character-gear')).toHaveAttribute('aria-busy', 'false');
}

async function noSidewaysScroll(page: Page): Promise<void> {
  const scroll = await page.evaluate(() => ({
    width: document.documentElement.scrollWidth,
    client: document.documentElement.clientWidth,
  }));
  expect(scroll.width).toBeLessThanOrEqual(scroll.client);
}

test.describe('the Gear frame at 1440', () => {
  test.use({ viewport: { width: 1440, height: 900 } });

  test('the owner: every worn item in its socket, the title bar, the Upgrades, no page error', async ({
    page,
  }) => {
    const errors = await open(page);
    await drawn(page);
    const section = page.getByTestId('character-gear-section');
    await expect(section.getByRole('heading', { name: 'Gear', exact: true })).toBeVisible();
    await expect(section.getByRole('link', { name: 'Best in slot ›' })).toHaveAttribute(
      'href',
      OWNER_BIS_HREF,
    );
    await expect(section.getByRole('heading', { name: 'Holy Priest' })).toBeVisible();
    await expect(section.getByTestId('character-doll').locator('.pdf-meta')).toHaveText('Level 60');
    await expect(frameSlot(page, 'head')).toHaveAccessibleName('Head: Green Lens');
    await expect(frameSlot(page, 'off_hand')).toHaveAccessibleName(/two-handed$/);
    await expect(frameSlot(page, 'trinket2')).toHaveCount(0);
    // No slot line carries the item level (ai-slop review): the frame shows icons alone.
    await expect(section).not.toContainText('Item level 49');
    const upgrades = page.locator('.pdf-frame').getByTestId('character-upgrades');
    await expect(upgrades).toBeVisible();
    await expect(upgrades.getByRole('heading', { name: 'Upgrades' })).toBeVisible();
    await expect(upgrades).toContainText('Halo of Transcendence');
    await expect(upgrades).toContainText('+1.8 DPS');
    await expect(upgrades).toContainText('+2.0 DPS');
    await expect(page.getByTestId('character-doll-note')).toHaveText(
      '1 item is not in our item data yet. It is new in the beta, and it gets a name and icon when the client lists it. Item level shows once every item is known.',
    );
    // The band's answer is the sum of the priced gains.
    await expect(page.getByTestId('character-answer')).toHaveText('+3.8DPS in upgrades');
    await noSidewaysScroll(page);
    expect(errors).toEqual([]);
  });

  test('a visitor sees the same worn set and the visitor link to best in slot', async ({ page }) => {
    const errors = await open(page, { owner: false });
    await drawn(page);
    await expect(frameSlot(page, 'head')).toHaveAccessibleName('Head: Green Lens');
    await expect(
      page.getByTestId('character-gear-section').getByRole('link', { name: 'Best in slot ›' }),
    ).toHaveAttribute('href', BIS_HREF);
    expect(errors).toEqual([]);
  });

  test('choosing a slot shows it in the centre: the pick, its source, the runner-up; then back', async ({
    page,
  }) => {
    await open(page);
    await drawn(page);
    await page.waitForLoadState('networkidle');
    await frameSlot(page, 'head').click();
    await expect(frameSlot(page, 'head')).toHaveAttribute('aria-pressed', 'true');
    const inspector = page.locator('.pdf-frame').getByTestId('character-inspector');
    await expect(inspector).toBeVisible();
    const pick = inspector.getByTestId('character-inspector-pick');
    await expect(pick).toContainText('Halo of Transcendence');
    await expect(pick).toContainText("Onyxia's Lair: Onyxia · 17% drop");
    await expect(inspector).toContainText('Green Lens');
    await expect(inspector).toContainText('−1.8');
    await expect(inspector.getByRole('link', { name: 'Open in best in slot ›' })).toHaveAttribute(
      'href',
      OWNER_BIS_HREF,
    );
    await inspector.getByTestId('character-upgrades-back').click();
    await expect(page.locator('.pdf-frame').getByTestId('character-upgrades')).toBeVisible();
    await noSidewaysScroll(page);
  });

  test('`gear: []`: no Gear panel, the rest of the page still draws', async ({ page }) => {
    await open(page, { gear: { gear: [], gear_gap_total: null } });
    await expect(page.getByTestId('character-gear-section')).toHaveCount(0);
    await expect(page.getByTestId('character-doll')).toHaveCount(0);
    await expect(page.getByTestId('character-answer')).toHaveCount(0);
    await expect(page.getByTestId('character-best')).toBeVisible();
  });

  test('the frame waits empty while the items file is slow: never a white name', async ({ page }) => {
    await open(page, { itemsDelayMs: 1500 });
    await expect(page.getByTestId('character-gear')).toHaveAttribute('aria-busy', 'true');
    await expect(frameSlot(page, 'head')).toHaveCount(0);
    await drawn(page);
    await expect(frameSlot(page, 'head')).toHaveAccessibleName('Head: Green Lens');
  });

  test('a failed items file keeps the sockets and says so once', async ({ page }) => {
    await open(page, { itemsFailures: 1 });
    await drawn(page);
    await expect(frameSlot(page, 'head')).toHaveAccessibleName('Head: Green Lens');
    await expect(page.getByTestId('character-doll-note')).toHaveText(
      'Item icons and colours did not load. Reload the page to try again.',
    );
  });

  test('a worn item opens its client tooltip on hover', async ({ page }) => {
    await open(page);
    await drawn(page);
    await page.waitForLoadState('networkidle');
    await frameSlot(page, 'head').hover();
    await expect(page.getByTestId('item-tooltip').first()).toContainText('Green Lens');
  });
});

test.describe('the Gear list form at 390', () => {
  test.use({ viewport: { width: 390, height: 844 } });

  test('every slot in two columns with its name, the Upgrades under them, never the frame image', async ({
    page,
  }) => {
    const errors = await open(page);
    await drawn(page);
    await expect(page.locator('.pdf-frame')).toBeHidden();
    await expect(listSlot(page, 'head')).toContainText('Green Lens');
    await expect(listSlot(page, 'head')).not.toContainText('Item level');
    const below = page.getByTestId('character-gear').getByTestId('character-upgrades').last();
    await expect(below).toBeVisible();
    await expect(below).toContainText('Halo of Transcendence');
    const heights = await page
      .locator('[data-testid^="paperdoll-list-slot-"]')
      .evaluateAll((els) => els.map((el) => el.getBoundingClientRect().height));
    for (const height of heights) expect(height).toBeGreaterThanOrEqual(44);
    await noSidewaysScroll(page);
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
        .locator('.pdf-name')
        .evaluateAll((els) =>
          els.filter((el) => el.scrollWidth > el.clientWidth + 1).map((el) => el.textContent),
        );
      expect(clipped).toEqual([]);
      await noSidewaysScroll(page);
    });
  });
}
