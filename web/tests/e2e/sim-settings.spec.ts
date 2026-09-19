import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test } from '@playwright/test';

const activeBuild = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
) as { build: string };

const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=11726`;

test('the settings bar reads the defaults and every control changes the settings', async ({ page }) => {
  await page.goto('/sim');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();

  const settings = page.getByTestId('sim-settings');
  await expect(settings).toBeVisible();

  // Defaults: defaultSettings() (settings.ts) is Patchwerk, 3:00, single target,
  // raid-buffed -- and the rotation names the spec the loaded character carries, Fury.
  await expect(page.getByTestId('sim-style')).toHaveValue('patchwerk');
  await expect(page.getByTestId('sim-duration')).toHaveValue('180');
  await expect(page.getByTestId('sim-targets')).toHaveValue('1');
  await expect(page.getByTestId('sim-preset')).toHaveValue('raid-buffed');
  await expect(page.getByTestId('sim-rotation')).toHaveText('Default for Fury');
  await expect(page.getByTestId('sim-rotation-link')).toHaveAttribute('href', '/sim/specs#warrior-fury');

  // Fight style: the contract's nine, plus nothing. Choosing one writes its fields.
  const style = page.getByTestId('sim-style');
  await expect(style.locator('option')).toHaveCount(9);
  await style.selectOption('cleave-3');
  await expect(page.getByTestId('sim-targets')).toHaveValue('3');
  await expect(page.getByTestId('sim-style-note')).toBeHidden();

  // The two movement styles carry the honesty note; nothing else does.
  await style.selectOption('light-movement');
  await expect(page.getByTestId('sim-targets')).toHaveValue('1');
  await expect(page.getByTestId('sim-style-note')).toBeVisible();

  // Setting targets by hand detaches from the style: the select falls to "Custom".
  await page.getByTestId('sim-targets').selectOption('4');
  await expect(style).toHaveValue('');

  // Fight length is labelled as a clock and reaches ten minutes.
  const duration = page.getByTestId('sim-duration');
  await expect(duration.locator('option', { hasText: '0:20' })).toHaveCount(1);
  await expect(duration.locator('option', { hasText: '10:00' })).toHaveCount(1);
  await duration.selectOption('600');
  await expect(duration).toHaveValue('600');

  // The secondary controls live in the disclosure and are closed on arrival.
  const more = page.getByTestId('sim-settings-more');
  await expect(more).toBeVisible();
  await expect(page.getByTestId('sim-variation')).toBeHidden();
  await more.getByRole('group').or(more).locator('summary').click();

  await expect(page.getByTestId('sim-variation')).toBeVisible();
  await page.getByTestId('sim-variation').selectOption('0');
  await expect(page.getByTestId('sim-variation')).toHaveValue('0');

  await page.getByTestId('sim-target-level').selectOption('60');
  await expect(page.getByTestId('sim-target-level')).toHaveValue('60');

  await page.getByTestId('sim-target-armor').fill('3731');
  await expect(page.getByTestId('sim-target-armor')).toHaveValue('3731');

  await page.getByTestId('sim-target-type').selectOption('undead');
  await expect(page.getByTestId('sim-target-type')).toHaveValue('undead');

  // Re-attaching a style, then toggling execute phase or the dummy, detaches it again:
  // withExecutePhase and withDummy both call settings.ts's `detached`, the same rule
  // targets used above -- SettingsSheet.svelte's own header comment claims this for both.
  await style.selectOption('patchwerk');
  await expect(style).toHaveValue('patchwerk');

  const execute = page.getByTestId('sim-execute');
  await execute.uncheck();
  await expect(execute).not.toBeChecked();
  await expect(style).toHaveValue('');
  await execute.check();

  await style.selectOption('patchwerk');
  await expect(style).toHaveValue('patchwerk');

  const dummy = page.getByTestId('sim-dummy');
  await dummy.check();
  await expect(dummy).toBeChecked();
  await expect(style).toHaveValue('');

  // Re-attaching Patchwerk above reset targets to its own count (1) each time; restore the
  // 4 set by hand earlier so the persistence check below still sees it.
  await page.getByTestId('sim-targets').selectOption('4');

  // The changes hold after the source switcher is reopened and the strip is brought back.
  await page.getByTestId('sim-change-source').click();
  await expect(page.getByTestId('sim-sources')).toBeVisible();
  await expect(duration).toHaveValue('600');
  await expect(page.getByTestId('sim-targets')).toHaveValue('4');
});

test('the settings bar is hidden until a character is loaded', async ({ page }) => {
  await page.goto('/sim');
  await expect(page.getByTestId('sim-settings')).toBeHidden();
  await expect(page.getByTestId('sim-empty')).toBeVisible();
});
