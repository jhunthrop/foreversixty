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
  // Final whole-branch review, C1: this trigger no longer navigates at all (it opens an
  // in-page drawer, tests/e2e/sim-rotation-drawer.spec.ts's own coverage) -- an `href`
  // assertion here used to pin the very navigation that fix removed.

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

/**
 * Task 4 (dps-minmaxer D16 BLOCKER): the old Raid-buffed preset was five buffs and two
 * consumables, +6% over Solo where vanilla's real answer for a melee is +60% to +120%.
 * This reads the fix from the same three places a reviewer's repro read the bug --
 * "what's in it", the Custom panel's own counters, and the request the browser actually
 * sends -- rather than only a unit test on settings.ts's own PRESET_BUFFS constant and
 * presetConsumables function.
 */
test('Raid-buffed is the full standard set: "what’s in it" names it, Custom’s counters show it, and the request carries it', async ({
  page,
}) => {
  await page.goto('/sim');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();
  await expect(page.getByTestId('sim-preset')).toHaveValue('raid-buffed');

  // "What's in it" is one click away (Disclosure.svelte, Ruling 3) and names the preset's
  // own ids -- humanised, never bare, even over this repo's sparse fixture name table
  // (src/fixtures/planner/simbuffs.json has no row for either id below).
  const trigger = page.getByTestId('sim-preset-summary-trigger');
  await expect(trigger).toBeVisible();
  await expect(trigger).toHaveAttribute('aria-expanded', 'false');
  await trigger.click();
  const panel = page.getByTestId('sim-preset-summary-panel');
  await expect(panel).toBeVisible();
  await expect(trigger).toHaveAttribute('aria-expanded', 'true');
  await expect(panel).toContainText(/sunder armor/i);
  await expect(panel).toContainText(/songflower/i);

  // Escape closes it and returns focus to the trigger, wherever the panel's own content
  // put it (Ruling 3's keyboard requirement).
  await page.keyboard.press('Escape');
  await expect(panel).toBeHidden();
  await expect(trigger).toBeFocused();
  await expect(trigger).toHaveAttribute('aria-expanded', 'false');

  // Switching to Custom carries every tick across (withPreset(..., 'custom')), so the
  // panel's own counters read the real totals now -- the findings' own quote was "RAID
  // BUFFS 4/30, ... ON THE TARGET 0/24, ... WEAPON OILS AND STONES 0/48".
  await page.getByTestId('sim-preset').selectOption('custom');
  const groupCount = async (group: string): Promise<number> => {
    const text = await page.getByTestId(`sim-buff-group-${group}`).locator('span.tabular').textContent();
    return Number((text ?? '0/0').split('/')[0]);
  };
  expect(await groupCount('raid-buffs')).toBeGreaterThan(4);
  expect(await groupCount('debuffs')).toBeGreaterThan(0);
  expect(await groupCount('weapon-imbue')).toBeGreaterThan(0);

  // The request the browser would actually send carries the expanded preset -- the REQUEST
  // drawer's own rendering of it, not only the settings object a unit test can see.
  await page.getByTestId('sim-preset').selectOption('raid-buffed');
  await page.getByTestId('sim-request-drawer').locator('summary').click();
  await expect(page.getByTestId('sim-request-buffs')).toContainText('sunder_armor');
  await expect(page.getByTestId('sim-request-buffs')).toContainText('blessing_of_kings');
  await expect(page.getByTestId('sim-request-json')).toHaveValue(
    /main_hand_imbue:elemental_sharpening_stone/,
  );
});
