// web/tests/e2e/bis-tank.spec.ts
// A tank spec's BiS page (warrior protection, both presets published at 60): This set beside the
// paperdoll carries the four tank figures (the frame's bar holds no DPS figure), the secondary
// line explains TMI, the Profile control swaps them, What to stack reads in score units, and
// nothing on the page says DPS. Figures are read from the published file, never pinned.
import type { Locator, Page } from '@playwright/test';
import { test, expect } from './support/test';
import { tankFigureTexts, tankMetrics } from './support/bis-file';
import { openBis, chooseControl } from './support/bis-page';

const TANK_PAGE = '/bis/warrior/protection?band=60';
const FIGURE_IDS = ['effective-health', 'dtps', 'chance-of-death', 'tps'] as const;
const WIDE_VIEWPORT = { width: 2000, height: 1000 };

const RAID_TEXTS = tankFigureTexts(tankMetrics('warrior-protection', 'raid'));
const BARE_TEXTS = tankFigureTexts(tankMetrics('warrior-protection', 'bare'));
const FIGURE_LABELS = ['Effective health', 'Damage taken per second', 'Chance of death', 'Threat per second'];

const thisSet = (page: Page): Locator => page.getByTestId('bis-set');
const figure = (panel: Locator, id: string): Locator => panel.getByTestId(`bis-tank-figure-${id}`);

const open = (page: Page, url = TANK_PAGE): Promise<void> => openBis(page, url);

async function expectFigures(panel: Locator, values: readonly string[]): Promise<void> {
  for (const [index, id] of FIGURE_IDS.entries()) {
    await expect(figure(panel, id)).toContainText(values[index] ?? '');
  }
}

test('a tank band lists the four tank figures, in order, and the frame shows no DPS figure', async ({
  page,
}) => {
  await open(page);
  const panel = thisSet(page);
  await expect(panel).toBeVisible();
  await expect(panel.locator('[data-testid^="bis-tank-figure-"]')).toHaveCount(FIGURE_IDS.length);
  expect(await panel.locator('dt').allTextContents()).toEqual(FIGURE_LABELS);
  await expectFigures(panel, RAID_TEXTS.headline);
  await expect(page.locator('[data-testid="bis-doll"] .pdf-figure')).toHaveCount(0);
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

test('What to stack reads in score, never DPS, with the expertise line', async ({ page }) => {
  await open(page);
  const weights = page.getByTestId('bis-weights');
  await expect(weights.locator('dt')).not.toHaveCount(0);
  await expect(weights.getByTestId('bis-weights-unit')).toHaveText(/^1\.00 = \d+\.\d{3} score$/);
  await expect(weights).not.toContainText('DPS');
  await expect(weights.getByTestId('bis-expertise-cap')).toContainText(/^Expertise to cap: /);
});

test('nothing at band level says DPS on a tank band but its own damage', async ({ page }) => {
  await open(page);
  // The TMI line names the tank's own damage in DPS; everything else reads in score.
  const text = await page
    .getByTestId('bis-band-alliance-60')
    .evaluate((el) =>
      (el.textContent ?? '').replace(
        el.querySelector('[data-testid="bis-tank-secondary"]')?.textContent ?? '',
        '',
      ),
    );
  expect(text).not.toMatch(/DPS/);
  await expect(page.getByTestId('bis-itinerary')).toContainText(/[−+]\d[\d,.]*\s*score/);
});

test('at 2000px the page stays centred and nothing scrolls sideways', async ({ page }) => {
  await page.setViewportSize(WIDE_VIEWPORT);
  await open(page);
  await expect(thisSet(page)).toBeVisible();
  const box = await page.getByTestId('bis-itinerary').boundingBox();
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
