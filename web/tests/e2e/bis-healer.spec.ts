// web/tests/e2e/bis-healer.spec.ts
// A healer spec's BiS page (priest holy, both presets at 60): the answer headlines HPS with the
// overheal, mana and healing-per-mana cells and no disclosure; nothing at band level
// says DPS; the simulator link is not offered; the weights and runners-up read in HPS.
import type { Page } from '@playwright/test';
import { test, expect } from './support/test';
import { readBisFile } from './support/bis-file';
import { openBis, chooseControl } from './support/bis-page';

const HEALER_PAGE = '/bis/priest/holy?band=60';
const MANA_LINE = /Mana lasts the whole [\d:]+ fight|Out of mana at \d+:\d\d/;
const PROFILE_LABEL = readBisFile('priest-holy').heal_profile?.label ?? '';

const answer = (page: Page) => page.getByTestId('bis-this-set');
const band = (page: Page) => page.getByTestId('bis-band-alliance-60');
const option = (page: Page, preset: 'raid' | 'bare') => page.getByTestId(`control-preset-${preset}`);

const open = (page: Page, url = HEALER_PAGE): Promise<void> => openBis(page, url);

async function expectHealerCells(page: Page): Promise<void> {
  const panel = answer(page);
  await expect(panel.getByTestId('bis-answer-figure')).toContainText(/\d+\.\d/);
  await expect(panel).toContainText('HPS');
  await expect(panel.getByTestId('bis-healer-overheal')).toContainText(/\d+% overheal/);
  await expect(panel.getByTestId('bis-healer-mana')).toContainText(MANA_LINE);
  await expect(panel.getByTestId('bis-healer-hpm')).toContainText(/\d+\.\d/);
}

test('the raid-ready healer answer shows HPS, overheal, mana, healing per mana and the profile', async ({
  page,
}) => {
  await open(page);
  await expectHealerCells(page);
  expect(PROFILE_LABEL).not.toBe('');
  await expect(answer(page).getByTestId('bis-healer-figure')).toHaveAttribute(
    'title',
    `Effective healing per second under ${PROFILE_LABEL}`,
  );
  await expect(answer(page).getByTestId('bis-healer-hpm')).toContainText('Healing per mana');
});

test('the page head has no disclosure and none of the profile notes', async ({ page }) => {
  await open(page);
  await expect(answer(page).locator('details')).toHaveCount(0);
  await expect(answer(page)).not.toContainText('What this profile assumes');
  await expect(answer(page)).not.toContainText('What is included');
});

test('choosing Bare swaps the healer figures and ?preset=bare opens on them', async ({ page }) => {
  await open(page);
  await expect(option(page, 'raid')).toHaveAttribute('aria-checked', 'true');
  const raidFigure = await answer(page).getByTestId('bis-answer-figure').innerText();
  await chooseControl(page, 'preset', 'bare');
  await expect(option(page, 'bare')).toHaveAttribute('aria-checked', 'true');
  await expect(page.getByTestId('bis-page')).toHaveAttribute('data-bis-preset', 'bare');
  await expectHealerCells(page);
  await expect(answer(page).getByTestId('bis-answer-figure')).not.toHaveText(raidFigure);

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
    for (const region of [answer(page), band(page)]) {
      const text = await region.evaluate((el) => el.textContent ?? '');
      expect(text).not.toMatch(/DPS/);
    }
    await expect(answer(page).locator('a', { hasText: 'Open in simulator' })).toHaveCount(0);
    await expect(answer(page).locator('a', { hasText: 'Talents in planner' })).toHaveCount(1);
  }
  await expect(page.getByTestId('bis-weights-alliance-60')).toContainText('HPS');
});

test('every set-bonus note names its set and bonus, and no slot shows more than one', async ({ page }) => {
  await open(page, `${HEALER_PAGE}&view=slot`);
  for (const cell of await page.locator('[data-testid="bis-doll"] [data-state="item"]').all()) {
    await cell.click();
    const notes = page.locator('[data-testid="bis-row-set-bonus"]:visible');
    expect(await notes.count()).toBeLessThanOrEqual(1);
    if ((await notes.count()) === 1) await expect(notes).toHaveText(/^Worn for the .+ \d+-piece bonus$/);
  }
});

test('the healer weights and runners-up read in HPS', async ({ page }) => {
  await open(page, `${HEALER_PAGE}&view=slot`);
  const weights = page.getByTestId('bis-weights-alliance-60');
  await expect(weights).toContainText('HPS per point');
  await expect(weights.getByTestId('bis-weight-row-healing_power')).toContainText(/\d\.\d{3} HPS/);
  await page.locator('[data-testid="bis-doll"] [data-state="item"]').first().click();
  const runners = page.locator('[data-testid="bis-detail-runners"]:visible');
  test.skip((await runners.count()) === 0, 'the first pick has no runners-up with the current data');
  await expect(runners).toContainText('Runners-up · HPS against the pick');
  for (const gap of await runners.locator('.bis-detail-gap').all()) {
    await expect(gap).toHaveText(/^(?:Pick|[+−]\d+\.\d HPS|same HPS)$/);
  }
});

test('the healer band has no horizontal scroll at 390px', async ({ page }) => {
  await page.setViewportSize({ width: 390, height: 844 });
  await open(page);
  await expect(answer(page).getByTestId('bis-answer-figure')).toBeVisible();
  const overflow = await page.evaluate(
    () => document.documentElement.scrollWidth - document.documentElement.clientWidth,
  );
  expect(overflow).toBeLessThanOrEqual(0);
});
