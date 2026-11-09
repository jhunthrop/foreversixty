// web/tests/e2e/bis-healer.spec.ts
// A healer spec's BiS page (priest holy, both presets at 60): the frame's title bar holds the set's
// HPS, This set beside it the overheal, mana and healing-per-mana cells; nothing at band level
// says DPS; the simulator link is not offered; the weights and the slot lines read in HPS.
import type { Page } from '@playwright/test';
import { test, expect } from './support/test';
import { openBis, chooseControl } from './support/bis-page';

const HEALER_PAGE = '/bis/priest/holy?band=60';
const MANA_VALUE = /^\d+:\d\d$/;

const set = (page: Page) => page.getByTestId('bis-set');
const figure = (page: Page) => page.locator('[data-testid="bis-doll"] .pdf-frame .pdf-figure');
const band = (page: Page) => page.getByTestId('bis-band-alliance-60');
const option = (page: Page, preset: 'raid' | 'bare') => page.getByTestId(`control-preset-${preset}`);

const open = (page: Page, url = HEALER_PAGE): Promise<void> => openBis(page, url);

async function expectHealerCells(page: Page): Promise<void> {
  await expect(figure(page)).toHaveText(/^\d+\.\d HPS$/);
  await expect(set(page).getByTestId('bis-healer-overheal')).toHaveText(/^\d+%$/);
  await expect(set(page).getByTestId('bis-healer-mana')).toHaveText(MANA_VALUE);
  await expect(set(page).getByTestId('bis-healer-hpm')).toHaveText(/\d+\.\d/);
}

test('the raid-ready healer set shows HPS on the frame, overheal, mana and healing per mana beside it', async ({
  page,
  isMobile,
}) => {
  test.skip(isMobile, 'the frame shows from 423 px');
  await open(page);
  await expectHealerCells(page);
  await expect(set(page)).toContainText('Healing per mana');
  await expect(set(page).locator('details')).toHaveCount(0);
});

test('choosing Bare swaps the healer figures and ?preset=bare opens on them', async ({ page, isMobile }) => {
  test.skip(isMobile, 'the frame shows from 423 px');
  await open(page);
  await expect(option(page, 'raid')).toHaveAttribute('aria-checked', 'true');
  const raidFigure = await figure(page).innerText();
  await chooseControl(page, 'preset', 'bare');
  await expect(option(page, 'bare')).toHaveAttribute('aria-checked', 'true');
  await expect(page.getByTestId('bis-page')).toHaveAttribute('data-bis-preset', 'bare');
  await expectHealerCells(page);
  await expect(figure(page)).not.toHaveText(raidFigure);

  await open(page, `${HEALER_PAGE}&preset=bare`);
  await expect(option(page, 'bare')).toHaveAttribute('aria-checked', 'true');
  await expectHealerCells(page);
});

test('no band-level number on a healer band says DPS, and the simulator link is not offered', async ({
  page,
}) => {
  await open(page);
  for (const preset of ['raid', 'bare'] as const) {
    await chooseControl(page, 'preset', preset);
    const text = await band(page).evaluate((el) => el.textContent ?? '');
    expect(text).not.toMatch(/DPS/);
    await expect(page.getByTestId('bis-sim-link')).toHaveCount(0);
  }
  await expect(page.getByTestId('bis-weights-unit')).toContainText('HPS per point');
});

test('the slot lines read in HPS', async ({ page, isMobile }) => {
  test.skip(isMobile, 'the frame shows from 423 px');
  await open(page);
  const gain = page.locator('[data-region="paperdoll-centre"]').getByTestId('bis-slot-gain');
  await expect(gain).toBeVisible();
  await expect(gain).not.toContainText('DPS');
});

test('the healer band has no horizontal scroll at 390px', async ({ page }) => {
  await page.setViewportSize({ width: 390, height: 844 });
  await open(page);
  await expect(page.locator('[data-testid^="paperdoll-list-slot-"]').first()).toBeVisible();
  const overflow = await page.evaluate(
    () => document.documentElement.scrollWidth - document.documentElement.clientWidth,
  );
  expect(overflow).toBeLessThanOrEqual(0);
});
