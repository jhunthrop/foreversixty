// web/tests/e2e/bis-tank.spec.ts
// A tank band on the BiS page: four headline figures in place of the DPS figure, tank
// stats in the weights table, score (never DPS) on the slot figures, and the Raid-ready /
// Bare control switching the figures. Hermetic: /bis/warrior/protection reads
// src/data/fixtures/bis/warrior-protection.json (a tank file, every figure invented).
import { test, expect, type Locator, type Page } from '@playwright/test';

const TANK_PAGE = '/bis/warrior/protection#band-alliance-60';
const RAID_VIEW = '[data-testid="bis-band-alliance-60"] [data-preset-view="raid"]';
const BARE_VIEW = '[data-testid="bis-band-alliance-60"] [data-preset-view="bare"]';
const FIGURE_IDS = ['effective-health', 'dtps', 'chance-of-death', 'tps'] as const;
const WIDE_VIEWPORT = { width: 2000, height: 1000 };

const RAID_FIGURES = ['11,235', '301.2', '<0.1%', '442.7'];
const BARE_FIGURES = ['8,420', '412.6', '8.1%', '318.4'];
const FIGURE_LABELS = ['Effective health', 'Damage taken per second', 'Chance of death', 'Threat per second'];
const TANK_WEIGHT_STATS = [
  'armor',
  'stamina',
  'defense',
  'dodge',
  'parry',
  'block',
  'block_value',
  'expertise',
  'strength',
  'attack_power',
] as const;

const thisSet = (page: Page, view: string): Locator => page.locator(`${view} [data-testid="bis-this-set"]`);
const figure = (panel: Locator, id: string): Locator => panel.getByTestId(`bis-tank-figure-${id}`);

async function expectFigures(panel: Locator, values: readonly string[]): Promise<void> {
  for (const [index, id] of FIGURE_IDS.entries()) {
    await expect(figure(panel, id)).toContainText(values[index] ?? '');
  }
}

test('a tank band headlines the four tank figures, in order, with no DPS figure', async ({ page }) => {
  await page.goto(TANK_PAGE);
  const panel = thisSet(page, RAID_VIEW);
  await expect(panel).toBeVisible();
  await expect(panel.locator('[data-testid^="bis-tank-figure-"]')).toHaveCount(FIGURE_IDS.length);
  const labels = await panel.locator('.tank-figure-label').allTextContents();
  expect(labels).toEqual(FIGURE_LABELS);
  await expectFigures(panel, RAID_FIGURES);
  await expect(panel).not.toContainText('DPS on the training dummy');
  await expect(panel.locator('.this-set-figure')).toHaveCount(0);
});

test('TMI and own damage sit on a quiet secondary line that explains TMI', async ({ page }) => {
  await page.goto(TANK_PAGE);
  const secondary = thisSet(page, RAID_VIEW).getByTestId('bis-tank-secondary');
  await expect(secondary).toHaveText('TMI 1,523 · own damage 260.0 DPS');
  await expect(secondary).toHaveAttribute('title', /risk index/);
});

test('the Raid-ready / Bare control switches the displayed figures', async ({ page }) => {
  await page.goto(TANK_PAGE);
  await expectFigures(thisSet(page, RAID_VIEW), RAID_FIGURES);
  await page.getByTestId('bis-preset-alliance-60-bare').click();
  await expect(page.locator(RAID_VIEW)).toBeHidden();
  await expectFigures(thisSet(page, BARE_VIEW), BARE_FIGURES);
  await page.getByTestId('bis-preset-alliance-60-raid').click();
  await expectFigures(thisSet(page, RAID_VIEW), RAID_FIGURES);
});

test('the weights table renders the tank stats, rating-family rows labelled as ratings', async ({ page }) => {
  await page.goto(TANK_PAGE);
  const weights = page.getByTestId('bis-weights-alliance-60');
  await expect(weights).toBeVisible();
  for (const stat of TANK_WEIGHT_STATS) {
    await expect(weights.getByTestId(`bis-weight-row-${stat}`)).toBeAttached();
  }
  await expect(weights.getByTestId('bis-weight-row-defense')).toContainText('Defense rating');
  await expect(weights.getByTestId('bis-weight-row-defense')).toHaveAttribute(
    'title',
    /2\.4 Defense rating = 1%/,
  );
  await expect(weights.getByTestId('bis-weight-row-hit')).toContainText('Not significant');
});

test('weights and slot figures say score, never DPS, on a tank band', async ({ page }) => {
  await page.goto(TANK_PAGE);
  const weights = page.getByTestId('bis-weights-alliance-60');
  await expect(weights).toContainText('score per point');
  await expect(weights).not.toContainText('DPS');
  const head = page.getByTestId('bis-slot-alliance-60-head');
  await expect(head).toContainText('−1.3 score');
  await expect(head).not.toContainText('DPS');
  await expect(page.getByTestId('bis-list-alliance-60')).not.toContainText('DPS');
  await expect(page.locator(RAID_VIEW).getByText('Runners-up · score vs the pick')).toBeVisible();
});

test('at 2000px the headline stays with equal side margins and nothing scrolls sideways', async ({
  page,
}) => {
  await page.setViewportSize(WIDE_VIEWPORT);
  await page.goto(TANK_PAGE);
  await expect(thisSet(page, RAID_VIEW)).toBeVisible();
  const box = await page.locator(`${RAID_VIEW} .three-panel-row`).boundingBox();
  expect(box).not.toBeNull();
  const left = box?.x ?? 0;
  const right = WIDE_VIEWPORT.width - left - (box?.width ?? 0);
  expect(left).toBeGreaterThanOrEqual(0);
  expect(Math.abs(left - right)).toBeLessThan(2);
  const overflow = await page.evaluate(
    () => document.documentElement.scrollWidth - document.documentElement.clientWidth,
  );
  expect(overflow).toBeLessThanOrEqual(0);
});

test('capture the tank headline for the design review', async ({ page }, testInfo) => {
  const dir = process.env.TANK_SHOTS_DIR;
  test.skip(dir === undefined || testInfo.project.name !== 'desktop', 'set TANK_SHOTS_DIR to capture');
  for (const width of [1440, 2000]) {
    await page.setViewportSize({ width, height: 1000 });
    for (const preset of ['raid', 'bare'] as const) {
      await page.goto(`/bis/warrior/protection?preset=${preset}#band-alliance-60`);
      await page.screenshot({ path: `${dir}/bis-tank-${preset}-${width}.png` });
    }
  }
});
