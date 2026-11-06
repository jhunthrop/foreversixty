// web/tests/e2e/bis.spec.ts
// The Best in slot page (web layout spec 2026-11-05 §5): the index links to a spec page; one
// island owns faction and band, switched from the control bar with no reload; the paperdoll
// draws every slot (a pick or an honest empty reason); a slot's runners-up are hover targets
// in the shared tooltip; What to stack carries the SimC-convention weights; the controls and
// the answer's links fit a 390 px phone. Picks come from the published nightly data, so the
// assertions read shapes and known copy, never a pinned item.
import { test, expect, type Page, type Locator } from './support/test';
import { unrankedBisRoute } from './support/bis-file';
import { openBis, waitForBis, chooseControl } from './support/bis-page';

const TWO_HANDER_EMPTY_COPY = /is a two-hander; the off hand is taken\./;
const ORDINARY_EMPTY_COPY =
  /raises your damage\.|Nothing here (?:either until \d+|helps at this level either)\.|Nothing sourced at this level helps your DPS|Relic effects aren't simulated yet|No sourced item at this level yet/;
const ANY_EMPTY_COPY = new RegExp(`${TWO_HANDER_EMPTY_COPY.source}|${ORDINARY_EMPTY_COPY.source}`);
const DOLL_SLOTS = 17;

/** Set once per page load; a full navigation wipes it, so reading it back proves "no reload". */
async function markNoReload(page: Page): Promise<void> {
  await page.evaluate(() => {
    (window as unknown as { __noReloadMarker?: boolean }).__noReloadMarker = true;
  });
}

async function expectNoReloadSoFar(page: Page): Promise<void> {
  const marker = await page.evaluate(() => (window as unknown as { __noReloadMarker?: boolean }).__noReloadMarker);
  expect(marker).toBe(true);
}

/** The island is hydrated once the root drops the pending mark and the controls answer clicks. */

function cells(page: Page): Locator {
  return page.locator('[data-testid="bis-doll"] [data-doll-slot]');
}

test('BiS: index links to a spec, faction and band controls switch the doll with no reload', async ({ page }) => {
  await page.goto('/bis');
  await expect(page.getByRole('heading', { name: 'Best in slot', level: 1 })).toBeVisible();
  const marksmanshipLink = page.locator('a[href^="/bis/hunter/marksmanship"]').first();
  await expect(marksmanshipLink).toBeVisible();
  await marksmanshipLink.click();
  await expect(page).toHaveURL(/\/bis\/hunter\/marksmanship/);
  await waitForBis(page);

  const root = page.getByTestId('bis-page');
  await expect(root).toHaveAttribute('data-bis-faction', 'alliance');
  const firstBand = page.locator('[data-testid^="bis-band-alliance-"]');
  await expect(firstBand).toBeVisible();
  const firstLevel = (await firstBand.getAttribute('data-testid'))?.replace('bis-band-alliance-', '');
  expect(Number(firstLevel)).toBeGreaterThan(0);

  // By slot, every slot is drawn: a pick that is its own hover target, or a known empty reason.
  await chooseControl(page, 'view', 'slot');
  await expect(cells(page)).toHaveCount(DOLL_SLOTS);
  await expect(page.locator('[data-testid="bis-doll"] [data-testid^="item-hover-"]').first()).toBeVisible();
  for (const cell of await cells(page).all()) {
    if ((await cell.getAttribute('data-state')) === 'empty') await expect(cell).toContainText(ANY_EMPTY_COPY);
  }

  await markNoReload(page);
  await chooseControl(page, 'faction', 'horde');
  await expect(root).toHaveAttribute('data-bis-faction', 'horde');
  await expect(page).toHaveURL(/[?&]faction=horde/);
  await expectNoReloadSoFar(page);

  await chooseControl(page, 'faction', 'alliance');
  await chooseControl(page, 'band', '30');
  const band30 = page.getByTestId('bis-band-alliance-30');
  await expect(band30).toBeVisible();
  await expect(page).toHaveURL(/[?&]band=30/);
  await expectNoReloadSoFar(page);
  // Band 30's picks include a quest reward, read in full on the doll's source line.
  await expect(band30.getByTestId('bis-doll')).toContainText('Quest:');
});

test('BiS: a spec with no ranked list yet shows the empty state, not a 404', async ({ page }) => {
  const route = unrankedBisRoute();
  test.skip(route === undefined, 'every written spec has a ranked list');
  const response = await page.goto(route!);
  expect(response?.status()).toBe(200);
  await expect(page.getByTestId('bis-empty-state')).toBeVisible();
  await expect(page.getByTestId('bis-empty-state')).toContainText('No leveling BiS list yet');
});

test('BiS: a pick’s runners-up open with the slot, every one its own hover target in the shared tooltip', async ({
  page,
}) => {
  await openBis(page, '/bis/hunter/marksmanship');
  const doll = page.getByTestId('bis-doll');
  // The first slot whose pick has runners-up: open it and read them in the slot's detail.
  let detail: Locator | undefined;
  for (const cell of await cells(page).all()) {
    if ((await cell.getAttribute('data-state')) !== 'item') continue;
    await cell.click();
    const runners = page.locator('[data-testid="bis-detail-runners"]:visible');
    if ((await runners.count()) > 0 && (await runners.locator('li').count()) >= 2) {
      detail = runners;
      break;
    }
  }
  test.skip(detail === undefined, 'no pick with runners-up on this band with the current data');
  const entries = detail!.locator('li');
  await expect(entries.first()).toContainText('Pick');
  await expect(entries.nth(1).locator('.bis-detail-gap')).toHaveText(/^(?:−\d+\.\d DPS|same DPS)$/);
  await expect(entries.nth(1).locator('.bis-detail-name')).not.toBeEmpty();
  await expect(entries.nth(1).locator('.bis-detail-src')).not.toBeEmpty();
  // "+0.0" never shows: a tie reads "Same as next best".
  await expect(doll).not.toContainText('+0.0 DPS');

  const tooltip = page.getByTestId('item-tooltip');
  const pickName = (await page.locator('[data-testid="bis-doll"] [aria-selected="true"] [data-testid^="item-hover-"]').innerText()).trim();
  await page.locator('[data-testid="bis-doll"] [aria-selected="true"] [data-testid^="item-hover-"]').hover();
  await expect(tooltip).toBeVisible();
  await expect(tooltip).toContainText(pickName);
  await page.mouse.move(0, 0);
  await expect(tooltip).toBeHidden();

  const altName = (await entries.nth(1).locator('.bis-detail-name').innerText()).trim();
  await entries.nth(1).locator('.bis-detail-name').hover();
  await expect(tooltip).toBeVisible();
  await expect(tooltip).toContainText(altName);
});

test('BiS: What to stack shows scale factors normalized to the top stat, DPS per point, the haste caption and the addon link', async ({
  page,
}) => {
  await openBis(page, '/bis/hunter/marksmanship?band=60');
  const weights = page.getByTestId('bis-weights-alliance-60');
  await expect(weights).toBeVisible();
  await expect(weights.getByTestId('bis-weights-sentence')).toContainText(/^Stack [a-z ,]+\.$/);
  await weights.getByText('Show the weights and their error').click();
  await expect(weights).toContainText(/Normalized to [A-Za-z ]+ = 1\.00/);
  await expect(weights).toContainText('1.00');
  await expect(weights).toContainText(/\d\.\d\d\d DPS/);
  // Haste has no rating in this client, so it is never a table row: the caption carries it.
  await expect(weights).toContainText(/Haste: \d+\.\d\d per 1%/);
  // The weights feed our own addon; no other addon is ever named (owner rule, 2026-09-30).
  await expect(weights).toContainText('Use these weights in the addon');
  await expect(page.getByTestId('bis-page')).not.toContainText(/Pawn/);
});

test('BiS: an insignificant stat reads "Not significant", never a number', async ({ page }) => {
  await openBis(page, '/bis/hunter/marksmanship');
  const bands = await page.locator('[data-testid^="control-band-"]').evaluateAll((nodes) =>
    nodes.map((node) => node.getAttribute('data-value') ?? ''),
  );
  for (const band of bands) {
    await chooseControl(page, 'band', band);
    const faint = page.getByTestId(`bis-weights-alliance-${band}`).locator('.faint', { hasText: 'Not significant' });
    if ((await faint.count()) > 0) {
      await expect(faint.first()).not.toContainText(/\d\.\d\d/);
      return;
    }
  }
  test.skip(true, 'no insignificant stat on any Alliance band with the current data');
});

test('BiS: an ordinary empty slot (not a two-hander gap) reads the plain no-source copy', async ({ page }) => {
  await page.goto('/bis/druid/balance?view=slot');
  const empty = page.locator('[data-testid="bis-doll"] [data-state="empty"]');
  const texts = await empty.allInnerTexts();
  test.skip(texts.length === 0, 'no empty slot on this band with the current data');
  for (const text of texts) expect(text).toMatch(ANY_EMPTY_COPY);
  const ordinary = texts.filter((text) => !TWO_HANDER_EMPTY_COPY.test(text));
  test.skip(ordinary.length === 0, 'every empty slot on this band is the two-hander gap with the current data');
  for (const text of ordinary) {
    expect(text).toMatch(ORDINARY_EMPTY_COPY);
    expect(text).not.toMatch(TWO_HANDER_EMPTY_COPY);
  }
});

test.describe('the controls at 390', () => {
  test.use({ viewport: { width: 390, height: 844 } });

  test('show all five bands, 44 px tall, with the selected one checked', async ({ page }) => {
    await page.goto('/bis/mage/fire?band=60');
    await waitForBis(page);
    await page.getByTestId('bis-controls').locator('[data-control-bar-toggle]').click();
    for (const band of [20, 30, 40, 50, 60]) {
      const option = page.getByTestId(`control-band-${band}`);
      await expect(option).toBeInViewport({ ratio: 1 });
      const box = await option.boundingBox();
      expect(box?.height ?? 0).toBeGreaterThanOrEqual(44);
      expect((box?.x ?? 0) + (box?.width ?? 0)).toBeLessThanOrEqual(390);
    }
    await expect(page.getByTestId('control-band-60')).toHaveAttribute('aria-checked', 'true');
  });

  test('keep every answer link inside the page', async ({ page }) => {
    await page.goto('/bis/warrior/fury');
    const answer = page.getByTestId('bis-this-set');
    await expect(answer).toBeVisible();
    const links = await answer.locator('a').all();
    expect(links.length).toBeGreaterThan(0);
    for (const link of links) {
      const box = await link.boundingBox();
      expect((box?.x ?? 0) + (box?.width ?? 0)).toBeLessThanOrEqual(390);
    }
  });
});
