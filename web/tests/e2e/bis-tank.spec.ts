// web/tests/e2e/bis-tank.spec.ts
// A tank spec's BiS page (warrior protection, both presets published at 60): the answer
// headlines the four tank figures instead of a DPS figure, the secondary line explains TMI,
// the Profile control swaps them, the weights table carries the tank stats in score units, and
// nothing on the page says DPS. Figures are read from the published file, never pinned.
import type { Locator, Page } from '@playwright/test';
import { test, expect } from './support/test';
import { bisBand, tankFigureTexts, tankMetrics } from './support/bis-file';
import { openBis, chooseControl } from './support/bis-page';

const TANK_PAGE = '/bis/warrior/protection?band=60';
const FIGURE_IDS = ['effective-health', 'dtps', 'chance-of-death', 'tps'] as const;
const WIDE_VIEWPORT = { width: 2000, height: 1000 };

const RAID_TEXTS = tankFigureTexts(tankMetrics('warrior-protection', 'raid'));
const BARE_TEXTS = tankFigureTexts(tankMetrics('warrior-protection', 'bare'));
const FIGURE_LABELS = ['Effective health', 'Damage taken per second', 'Chance of death', 'Threat per second'];
const TANK_WEIGHT_STATS = [
  'armor',
  'stamina',
  'defense',
  'dodge',
  'parry',
  'block',
  'block_value',
  'strength',
  'attack_power',
] as const;

const thisSet = (page: Page): Locator => page.getByTestId('bis-this-set');
const figure = (panel: Locator, id: string): Locator => panel.getByTestId(`bis-tank-figure-${id}`);

const open = (page: Page, url = TANK_PAGE): Promise<void> => openBis(page, url);

async function expectFigures(panel: Locator, values: readonly string[]): Promise<void> {
  for (const [index, id] of FIGURE_IDS.entries()) {
    await expect(figure(panel, id)).toContainText(values[index] ?? '');
  }
}

test('a tank band headlines the four tank figures, in order, with no DPS figure', async ({ page }) => {
  await open(page);
  const panel = thisSet(page);
  await expect(panel).toBeVisible();
  await expect(panel.locator('[data-testid^="bis-tank-figure-"]')).toHaveCount(FIGURE_IDS.length);
  const labels = await panel.locator('.tank-figure-label').allTextContents();
  expect(labels).toEqual(FIGURE_LABELS);
  await expectFigures(panel, RAID_TEXTS.headline);
  await expect(panel).not.toContainText('DPS on the training dummy');
  await expect(panel.getByTestId('bis-set-figure')).toHaveCount(0);
});

test('TMI and own damage sit on a quiet secondary line that explains TMI', async ({ page }) => {
  await open(page);
  const secondary = thisSet(page).getByTestId('bis-tank-secondary');
  await expect(secondary).toHaveText(RAID_TEXTS.secondary);
  await expect(secondary).toHaveAttribute('title', /risk index/);
});

test('the Raid-ready / Bare control switches the displayed figures', async ({ page }) => {
  await open(page);
  await expectFigures(thisSet(page), RAID_TEXTS.headline);
  await chooseControl(page, 'preset', 'bare');
  await expect(page.getByTestId('bis-page')).toHaveAttribute('data-bis-preset', 'bare');
  await expectFigures(thisSet(page), BARE_TEXTS.headline);
  await chooseControl(page, 'preset', 'raid');
  await expectFigures(thisSet(page), RAID_TEXTS.headline);
});

test('the weights table renders the tank stats, rating-family rows labelled as ratings', async ({ page }) => {
  await open(page);
  const weights = page.getByTestId('bis-weights-alliance-60');
  await expect(weights).toBeVisible();
  await weights.getByText('Show the weights and their error').click();
  for (const stat of TANK_WEIGHT_STATS) {
    await expect(weights.getByTestId(`bis-weight-row-${stat}`)).toBeAttached();
  }
  // Expertise is per 1%, not per point: a caption, never a per-point table row.
  await expect(weights.getByTestId('bis-weight-row-expertise')).toHaveCount(0);
  await expect(weights).toContainText(/Expertise: (\d+\.\d\d per 1%|no gain|not significant)/);
  await expect(weights.getByTestId('bis-weight-row-defense')).toContainText('Defense rating');
  const published = bisBand('warrior-protection', 'alliance', 60).weights ?? [];
  const defenseFactor = published.find((weight) => weight.stat === 'defense')?.rating_factor;
  expect(defenseFactor).toBeDefined();
  await expect(weights.getByTestId('bis-weight-row-defense')).toHaveAttribute(
    'title',
    new RegExp(`${defenseFactor} Defense rating = 1%`),
  );
  const hitRow = weights.getByTestId('bis-weight-row-hit');
  const hitInsignificant = published.find((weight) => weight.stat === 'hit')?.insignificant === true;
  if (hitInsignificant) await expect(hitRow).toContainText('Not significant');
  else await expect(hitRow).not.toContainText('Not significant');
});

test('weights and slot figures say score, never DPS, on a tank band', async ({ page }) => {
  await open(page);
  const weights = page.getByTestId('bis-weights-alliance-60');
  await expect(weights).toContainText('score per point');
  await expect(weights).not.toContainText('DPS');
  // A tied slot with no different alternative carries no figure ("No close alternative"), so the
  // unit is asserted over the doll's figures, not one slot's.
  const head = page.getByTestId('bis-doll-slot-head');
  await expect(head).not.toContainText('DPS');
  await expect(page.getByTestId('bis-doll')).toContainText(/score vs next best|[−+]\d[\d,.]* score/);
  await expect(page.getByTestId('bis-doll')).not.toContainText('DPS');
  await head.click();
  const runners = page.locator('[data-testid="bis-detail-runners"]:visible');
  await expect(runners).toContainText('Runners-up · score against the pick');
  await expect(runners).toContainText(/(same|[−+]\d[\d,.]*) score/);
  await expect(runners).not.toContainText('DPS');
});

test('at 2000px the headline stays with equal side margins and nothing scrolls sideways', async ({
  page,
}) => {
  await page.setViewportSize(WIDE_VIEWPORT);
  await open(page);
  await expect(thisSet(page)).toBeVisible();
  const box = await page.getByTestId('bis-plate').boundingBox();
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
