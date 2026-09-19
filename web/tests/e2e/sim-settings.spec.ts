import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test } from '@playwright/test';

// Same fixture as sim-sources.spec.ts: an addon export needs no API stub, so this reads the
// site's own active build id the same way, rather than hardcoding a build the fixture
// server does not publish under.
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

  // Defaults: defaultSettings() (settings.ts) is 3:00, single target, execute phase on,
  // raid-buffed -- and the rotation names the spec the loaded character carries, Fury.
  await expect(page.getByTestId('sim-duration')).toHaveValue('180');
  await expect(page.getByTestId('sim-targets')).toHaveValue('1');
  await expect(page.getByTestId('sim-execute')).toBeChecked();
  await expect(page.getByTestId('sim-preset')).toHaveValue('raid-buffed');
  await expect(page.getByTestId('sim-rotation')).toHaveText('Default for Fury');
  await expect(page.getByTestId('sim-rotation-link')).toHaveAttribute('href', '/sim/specs#warrior-fury');
  await expect(page.getByTestId('sim-settings-footnote')).toHaveText(
    'Fight length varies by 20% between iterations, the way real pulls do.',
  );

  // Fight length: the select is labelled "1:00" .. "8:00", not the raw seconds.
  const duration = page.getByTestId('sim-duration');
  await expect(duration.locator('option', { hasText: '1:00' })).toHaveCount(1);
  await expect(duration.locator('option', { hasText: '8:00' })).toHaveCount(1);
  await duration.selectOption('300');
  await expect(duration).toHaveValue('300');

  // Targets: 1 .. MAX_TARGETS (10).
  const targets = page.getByTestId('sim-targets');
  await expect(targets.locator('option')).toHaveCount(10);
  await targets.selectOption('4');
  await expect(targets).toHaveValue('4');

  // Execute phase: unchecking it zeroes execute_ratio (withExecutePhase); the control
  // itself is the only observable here since the ratio is not rendered anywhere.
  const execute = page.getByTestId('sim-execute');
  await execute.uncheck();
  await expect(execute).not.toBeChecked();
  await execute.check();
  await expect(execute).toBeChecked();

  // Buffs: Raid-buffed, Solo, Custom, and Custom's own title names what choosing it does.
  const preset = page.getByTestId('sim-preset');
  await expect(preset.locator('option')).toHaveCount(3);
  await preset.selectOption('solo');
  await expect(preset).toHaveValue('solo');
  await preset.selectOption('custom');
  await expect(preset).toHaveValue('custom');
  await expect(preset.locator('option[value="custom"]')).toHaveAttribute(
    'title',
    'Keeps the buffs already applied. Choosing each one individually arrives with Top Gear.',
  );

  // The changes hold after the source switcher is reopened and the strip is brought back --
  // the settings bar is keyed to the character, not to the switcher being closed.
  await page.getByTestId('sim-change-source').click();
  await expect(page.getByTestId('sim-sources')).toBeVisible();
  await expect(duration).toHaveValue('300');
  await expect(targets).toHaveValue('4');
});

test('the settings bar is hidden until a character is loaded', async ({ page }) => {
  await page.goto('/sim');
  await expect(page.getByTestId('sim-settings')).toBeHidden();
  await expect(page.getByTestId('sim-empty')).toBeVisible();
});
