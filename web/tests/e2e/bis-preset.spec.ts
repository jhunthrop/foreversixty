// web/tests/e2e/bis-preset.spec.ts
// Level 60's Profile control (Raid-ready · Bare) on the BiS page's control bar, against the
// published tank file carrying both presets, so every expected figure and count is read from
// the file rather than pinned.
import type { Page } from '@playwright/test';
import { test, expect } from './support/test';
import { bisPreset, hitCapText, readBisFile, tankFigureTexts, tankMetrics } from './support/bis-file';
import { openBis as open, chooseControl } from './support/bis-page';

const SPEC = 'warrior-protection';
const PRESET_PAGE = '/bis/warrior/protection?band=60';
const BARE_FIGURE = tankFigureTexts(tankMetrics(SPEC, 'bare')).headline[0]!;
const RAID_FIGURE = tankFigureTexts(tankMetrics(SPEC, 'raid')).headline[0]!;

const countOf = (count: number, noun: string): string => `${count} ${noun}${count === 1 ? '' : 's'}`;
const option = (page: Page, preset: 'raid' | 'bare') => page.getByTestId(`control-preset-${preset}`);
const answer = (page: Page) => page.getByTestId('bis-this-set');
const root = (page: Page) => page.getByTestId('bis-page');
const escapeRegExp = (text: string): string => text.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');

test('level 60 opens on the raid-ready preset with a trust line naming what it includes', async ({
  page,
}) => {
  await open(page, PRESET_PAGE);
  await expect(option(page, 'raid')).toHaveAttribute('aria-checked', 'true');
  await expect(option(page, 'bare')).toHaveAttribute('aria-checked', 'false');
  await expect(root(page)).toHaveAttribute('data-bis-preset', 'raid');
  await expect(answer(page)).toContainText(RAID_FIGURE);
  const trust = page.getByTestId('bis-trust-line');
  await expect(trust).toContainText('Raid-ready, Phase 1');
  const raid = bisPreset(SPEC, 'raid');
  await expect(trust).toContainText(
    `${countOf(raid.buffs.length, 'buff')}, ${countOf(raid.debuffs.length, 'debuff')} and ${countOf(raid.consumes.length, 'consumable')}`,
  );
});

test('both presets state what the sim fights when the run is published, and never invent it', async ({
  page,
}) => {
  const run = readBisFile(SPEC).run;
  await open(page, PRESET_PAGE);
  const trust = page.getByTestId('bis-trust-line');
  for (const preset of ['raid', 'bare'] as const) {
    await chooseControl(page, 'preset', preset);
    if (run === undefined) await expect(trust).not.toContainText(/ s ·|seconds|runs/);
    else await expect(trust).toContainText(`${run.target_type.replace(/_/g, ' ')}, ${run.duration_sec} s`);
  }
});

test('the page head has no disclosure; the raid preset names its effects on the trust line title', async ({
  page,
}) => {
  await open(page, PRESET_PAGE);
  await expect(answer(page).locator('details')).toHaveCount(0);
  await expect(answer(page)).not.toContainText('What is included');
  const firstConsumable = bisPreset(SPEC, 'raid').consumes[0]?.label ?? '';
  expect(firstConsumable).not.toBe('');
  await expect(page.getByTestId('bis-trust-line')).toHaveAttribute(
    'title',
    new RegExp(`Consumables: .*${escapeRegExp(firstConsumable)}`),
  );
});

test('choosing Bare swaps the numbers and writes ?preset=bare without a reload', async ({ page }) => {
  await open(page, PRESET_PAGE);
  await page.evaluate(() => {
    (window as unknown as { __marker?: boolean }).__marker = true;
  });
  await chooseControl(page, 'preset', 'bare');
  await expect(page).toHaveURL(/[?&]preset=bare/);
  await expect(page).toHaveURL(/\/60\?/);
  await expect(option(page, 'bare')).toHaveAttribute('aria-checked', 'true');
  await expect(answer(page)).toContainText(BARE_FIGURE);
  await expect(page.getByTestId('bis-trust-line')).toContainText('no raid buffs');
  expect(await page.evaluate(() => (window as unknown as { __marker?: boolean }).__marker)).toBe(true);
  await chooseControl(page, 'preset', 'raid');
  await expect(page).not.toHaveURL(/[?&]preset=/);
  await expect(answer(page)).toContainText(RAID_FIGURE);
});

test('a ?preset=bare link opens on the bare numbers, stat weights and list included', async ({ page }) => {
  await open(page, `${PRESET_PAGE}&preset=bare`);
  await expect(option(page, 'bare')).toHaveAttribute('aria-checked', 'true');
  await expect(root(page)).toHaveAttribute('data-bis-preset', 'bare');
  await expect(answer(page)).toContainText(BARE_FIGURE);
  await expect(page.getByTestId('bis-weights-alliance-60')).toBeVisible();
  await expect(page.getByTestId('bis-doll')).toBeVisible();
});

test('the legacy ?preset=bare#band link still opens on the bare 60', async ({ page }) => {
  await open(page, '/bis/warrior/protection?preset=bare#band-alliance-60');
  await expect(root(page)).toHaveAttribute('data-bis-band', '60');
  await expect(answer(page)).toContainText(BARE_FIGURE);
});

test('the hit-to-cap line follows the preset, white clause only when published', async ({ page }) => {
  const retribution = (preset: 'raid' | 'bare') =>
    readBisFile('paladin-retribution').bands.find((b) => b.band === 60 && (b.preset ?? 'bare') === preset)!;
  const fury = readBisFile('warrior-fury').bands.find((b) => b.band === 60 && b.preset === 'raid')!;
  expect(retribution('raid').hit_to_cap?.white).toBeUndefined();
  expect(fury.hit_to_cap?.white).toBeDefined();

  await open(page, '/bis/paladin/retribution?band=60');
  const line = page.getByTestId('bis-hit-cap');
  await expect(line).toHaveText(hitCapText(retribution('raid').hit_to_cap!));
  await expect(line).toHaveAttribute('title', /full weight until the cap and nothing past it/);
  await expect(line).toHaveAttribute('title', /\d+ hit rating and \d+ crit rating are 1%/);
  await chooseControl(page, 'preset', 'bare');
  await expect(line).toHaveText(hitCapText(retribution('bare').hit_to_cap!));

  await open(page, '/bis/warrior/fury?band=60');
  await expect(page.getByTestId('bis-hit-cap')).toHaveText(hitCapText(fury.hit_to_cap!));
});

test('a band without the key shows no hit-to-cap line and keeps the weights in place', async ({ page }) => {
  expect(readBisFile('priest-holy').bands.every((band) => !band.hit_to_cap)).toBe(true);
  await open(page, '/bis/priest/holy?faction=horde&band=60');
  await expect(page.getByTestId('bis-weights-horde-60')).toBeAttached();
  await expect(page.getByTestId('bis-hit-cap')).toHaveCount(0);
});

test('below level 60 there is no Profile control: bare only', async ({ page }) => {
  await open(page, '/bis/hunter/marksmanship');
  await expect(page.locator('[data-testid^="bis-band-alliance-"]')).toBeVisible();
  await expect(page.locator('[data-testid^="control-preset-"]')).toHaveCount(0);
});
