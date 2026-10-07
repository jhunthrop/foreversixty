// web/tests/e2e/bis-preset.spec.ts
// The level-60 "Raid-ready · Bare" control on the BiS page. Hermetic: /bis/warrior/protection
// has no published file, so the page reads src/data/fixtures/bis/warrior-protection.json,
// which carries both presets (bare set DPS 215, raid set DPS 260, invented). Below level 60
// a real file (hunter/marksmanship band 20) proves the control is absent.
import { test, expect, type Page } from '@playwright/test';

const PRESET_PAGE = '/bis/warrior/protection#band-alliance-60';
const BARE_DPS = '215.0';
const RAID_DPS = '260.0';

const raidView = (page: Page) =>
  page.locator('[data-testid="bis-band-alliance-60"] [data-preset-view="raid"]');
const bareView = (page: Page) =>
  page.locator('[data-testid="bis-band-alliance-60"] [data-preset-view="bare"]');
const option = (page: Page, preset: 'raid' | 'bare') => page.getByTestId(`bis-preset-alliance-60-${preset}`);

test('level 60 opens on the raid-ready preset with a caption naming what it includes', async ({ page }) => {
  await page.goto(PRESET_PAGE);
  await expect(option(page, 'raid')).toHaveAttribute('aria-checked', 'true');
  await expect(option(page, 'bare')).toHaveAttribute('aria-checked', 'false');
  await expect(raidView(page).getByTestId('bis-this-set')).toContainText(RAID_DPS);
  await expect(bareView(page)).toBeHidden();
  const caption = page.getByTestId('bis-preset-caption-alliance-60-raid');
  await expect(caption).toContainText('Raid-ready, Phase 1');
  await expect(caption).toContainText('2 buffs, 1 debuff and 2 consumables');
});

test('both presets state what the sim fights, next to the caption', async ({ page }) => {
  await page.goto(PRESET_PAGE);
  const note =
    /Simulated against a level-63 target of no creature type over 180 seconds; slaying bonuses such as attack power against Undead count for nothing here\./;
  await expect(page.getByTestId('bis-preset-caption-alliance-60-raid-target')).toHaveText(note);
  await option(page, 'bare').click();
  await expect(page.getByTestId('bis-preset-caption-alliance-60-bare-target')).toHaveText(note);
});

test('the id lists live in a disclosure, not inline', async ({ page }) => {
  await page.goto(PRESET_PAGE);
  const caption = page.getByTestId('bis-preset-caption-alliance-60-raid');
  await expect(caption.getByText('Fixture Flask')).toBeHidden();
  await caption.getByText('What is included').click();
  await expect(caption.getByText('Fixture Flask', { exact: false })).toBeVisible();
});

test('choosing Bare swaps the numbers and writes ?preset=bare without a reload', async ({ page }) => {
  await page.goto(PRESET_PAGE);
  await page.evaluate(() => {
    (window as unknown as { __marker?: boolean }).__marker = true;
  });
  await option(page, 'bare').click();
  await expect(page).toHaveURL(/[?&]preset=bare/);
  await expect(page).toHaveURL(/#band-alliance-60$/);
  await expect(option(page, 'bare')).toHaveAttribute('aria-checked', 'true');
  await expect(bareView(page).getByTestId('bis-this-set')).toContainText(BARE_DPS);
  await expect(raidView(page)).toBeHidden();
  await expect(page.getByTestId('bis-preset-caption-alliance-60-bare')).toContainText('no raid buffs');
  expect(await page.evaluate(() => (window as unknown as { __marker?: boolean }).__marker)).toBe(true);
  await option(page, 'raid').click();
  await expect(page).toHaveURL(/[?&]preset=raid/);
  await expect(raidView(page).getByTestId('bis-this-set')).toContainText(RAID_DPS);
});

test('a ?preset=bare link opens on the bare numbers, stat weights and list included', async ({ page }) => {
  await page.goto(`/bis/warrior/protection?preset=bare${PRESET_PAGE.slice(PRESET_PAGE.indexOf('#'))}`);
  await expect(option(page, 'bare')).toHaveAttribute('aria-checked', 'true');
  await expect(bareView(page).getByTestId('bis-this-set')).toContainText(BARE_DPS);
  await expect(page.getByTestId('bis-weights-alliance-60-bare')).toBeVisible();
  await expect(page.getByTestId('bis-weights-alliance-60')).toBeHidden();
  await expect(page.getByTestId('bis-list-alliance-60-bare')).toBeVisible();
});

test('the weights carry a hit-to-cap line per preset, white clause only when published', async ({ page }) => {
  await page.goto(PRESET_PAGE);
  const raidLine = page.getByTestId('bis-weights-alliance-60-hit-cap');
  await expect(raidLine).toHaveText('Hit to cap: 6% for specials, 25% for white swings');
  await expect(raidLine).toHaveAttribute('title', /full weight until the cap and nothing past it/);
  await expect(raidLine).toHaveAttribute('title', /10 hit rating and 14 crit rating are 1%/);
  await option(page, 'bare').click();
  await expect(page.getByTestId('bis-weights-alliance-60-bare-hit-cap')).toHaveText(
    'Hit to cap: 6% for specials',
  );
});

test('a band without the key shows no hit-to-cap line and keeps the table in place', async ({ page }) => {
  await page.goto('/bis/warrior/protection#band-horde-60');
  await expect(page.getByTestId('bis-weights-horde-60')).toBeAttached();
  await expect(page.locator('[data-testid^="bis-weights-horde-60"][data-testid$="-hit-cap"]')).toHaveCount(0);
});

test('below level 60 there is no control: bare only', async ({ page }) => {
  await page.goto('/bis/hunter/marksmanship#band-alliance-20');
  const band = page.getByTestId('bis-band-alliance-20');
  await expect(band).toBeVisible();
  await expect(band.locator('[data-preset-option]')).toHaveCount(0);
});

test('capture the preset control for the design review', async ({ page }, testInfo) => {
  const dir = process.env.PRESET_SHOTS_DIR;
  test.skip(dir === undefined || testInfo.project.name !== 'desktop', 'set PRESET_SHOTS_DIR to capture');
  for (const width of [1440, 2000]) {
    await page.setViewportSize({ width, height: 1000 });
    for (const preset of ['raid', 'bare'] as const) {
      await page.goto(`/bis/warrior/protection?preset=${preset}#band-alliance-60`);
      await page.screenshot({ path: `${dir}/bis-preset-${preset}-${width}.png` });
    }
  }
});
