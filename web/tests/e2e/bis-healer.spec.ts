// web/tests/e2e/bis-healer.spec.ts
// The healer variant of the /bis "This set" panel. /bis/priest/holy reads the published
// data/builds/<build>/bis/priest-holy.json (the committed fixture only when none exists).
// Assertions are on shape, and the profile caption on the file's own heal_profile label, so
// a re-rank or a retuned profile keeps the spec green.
import { test, expect, type Page } from '@playwright/test';
import { readBisFile } from './support/bis-file';

const HEALER_PAGE = '/bis/priest/holy#band-alliance-60';
// Either mana line: a set that lasts the fight, or the second the healer ran dry.
const MANA_LINE = /Mana lasts the whole [\d:]+ fight|Out of mana at \d+:\d\d/;
// The caption names the profile the file was ranked under, so the label comes from the file.
const PROFILE_LABEL = readBisFile('priest-holy').heal_profile?.label ?? '';

const bandPanel = (page: Page) => page.getByTestId('bis-band-alliance-60');
const raidView = (page: Page) => bandPanel(page).locator('[data-preset-view="raid"]');
const bareView = (page: Page) => bandPanel(page).locator('[data-preset-view="bare"]');
const option = (page: Page, preset: 'raid' | 'bare') => page.getByTestId(`bis-preset-alliance-60-${preset}`);

test('the raid-ready healer panel shows HPS, overheal, mana, healing per mana and the profile', async ({
  page,
}) => {
  await page.goto(HEALER_PAGE);
  const panel = raidView(page).getByTestId('bis-this-set');
  await expect(panel.getByTestId('bis-healer-figure')).toHaveText(/^\d+\.\d$/);
  await expect(panel).toContainText('HPS');
  expect(PROFILE_LABEL).not.toBe('');
  await expect(panel.getByTestId('bis-healer-caption')).toHaveText(
    `Effective healing per second under ${PROFILE_LABEL}`,
  );
  await expect(panel.getByTestId('bis-healer-overheal')).toContainText(/\d+% overheal/);
  await expect(panel.getByTestId('bis-healer-mana')).toContainText(MANA_LINE);
  await expect(panel.getByTestId('bis-healer-hpm')).toContainText(/\d+\.\d/);
  await expect(panel.getByTestId('bis-healer-hpm')).toContainText('Healing per mana');
});

test('the profile is a disclosure that states what the sim assumes', async ({ page }) => {
  await page.goto(HEALER_PAGE);
  const profile = raidView(page).getByTestId('bis-healer-profile');
  await expect(profile.getByText(/health each/)).toBeHidden();
  await profile.getByText('What this profile assumes').click();
  await expect(profile.getByText(/\d[\d,]* health, a [\d,]+ hit every/)).toBeVisible();
  await expect(profile.getByText(/[\d,]+ health each/)).toBeVisible();
  await expect(profile.getByText(/[\d,]+ damage to \d+ members every/)).toBeVisible();
});

test('choosing Bare swaps the healer figures and ?preset=bare opens on them', async ({ page }) => {
  await page.goto(HEALER_PAGE);
  await expect(option(page, 'raid')).toHaveAttribute('aria-checked', 'true');
  await option(page, 'bare').click();
  await expect(option(page, 'bare')).toHaveAttribute('aria-checked', 'true');
  const bare = bareView(page).getByTestId('bis-this-set');
  await expect(bare.getByTestId('bis-healer-figure')).toHaveText(/^\d+\.\d$/);
  await expect(bare.getByTestId('bis-healer-overheal')).toContainText(/\d+% overheal/);
  await expect(bare.getByTestId('bis-healer-mana')).toContainText(MANA_LINE);
  await expect(bare.getByTestId('bis-healer-hpm')).toContainText(/\d+\.\d/);
  await expect(raidView(page)).toBeHidden();

  await page.goto('/bis/priest/holy?preset=bare#band-alliance-60');
  await expect(option(page, 'bare')).toHaveAttribute('aria-checked', 'true');
  await expect(bareView(page).getByTestId('bis-healer-figure')).toHaveText(/^\d+\.\d$/);
});

test('no band-level number on a healer band says DPS, and the simulator link is not offered', async ({
  page,
}) => {
  await page.goto(HEALER_PAGE);
  for (const view of [raidView(page), bareView(page)]) {
    const text = await view.evaluate((el) => (el as HTMLElement).textContent ?? '');
    expect(text).not.toMatch(/DPS/);
    await expect(view.locator('a', { hasText: 'Open in simulator' })).toHaveCount(0);
    await expect(view.locator('a', { hasText: 'Talents in planner' })).toHaveCount(1);
  }
  await expect(raidView(page).getByTestId('bis-weights-alliance-60')).toContainText('HPS');
});

test('every set-bonus note names its set and bonus, and no row shows more than one', async ({ page }) => {
  await page.goto(HEALER_PAGE);
  // Shape, not count: the fixture carries one adopted slot, a published file carries however
  // many the nightly ranker found (possibly none), and the published file wins.
  const notes = await raidView(page).getByTestId('bis-row-set-bonus').all();
  for (const note of notes) {
    await expect(note).toHaveText(/^Worn for the .+ \d+-piece bonus$/);
    await expect(note).toHaveAttribute('title', /^.+ \(\d+\): .+/);
  }
  await expect(raidView(page).locator('li:has([data-testid="bis-row-set-bonus"])')).toHaveCount(notes.length);
});

test('the healer weights and runners-up read in HPS', async ({ page }) => {
  await page.goto(HEALER_PAGE);
  const weights = raidView(page).getByTestId('bis-weights-alliance-60');
  await expect(weights).toContainText('HPS per point');
  await expect(weights.getByTestId('bis-weight-row-healing_power')).toContainText(/\d\.\d{3} HPS/);
  await expect(raidView(page).locator('.the-list-header')).toContainText('Runners-up · HPS vs the pick');
  await expect(raidView(page).getByTestId('bis-row-evidence').first()).toContainText(/\+\d+\.\d HPS/);
});

test('the class index lists a healer spec in HPS, beside no ranking of specs', async ({ page }) => {
  await page.goto('/bis');
  await expect(page.getByTestId('bis-index-dps-priest-holy')).toHaveText(/^Level 60: \d+\.\d HPS$/);
});

test('the healer band has no horizontal scroll at 390px', async ({ page }) => {
  await page.setViewportSize({ width: 390, height: 844 });
  await page.goto(HEALER_PAGE);
  await expect(raidView(page).getByTestId('bis-healer-figure')).toBeVisible();
  const overflow = await page.evaluate(
    () => document.documentElement.scrollWidth - document.documentElement.clientWidth,
  );
  expect(overflow).toBeLessThanOrEqual(0);
});
