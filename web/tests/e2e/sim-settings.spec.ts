import { expect, test } from './support/test';
import { choose } from './support/select';
import {
  chooseBuffs,
  closeMoreSettings,
  gotoSim,
  openMoreSettings,
  openRequest,
  pasteAndRun,
  SIM_FURY as FURY,
} from './support/sim-page';

// The same character recorded at level 60: the Raid-buffed consumables follow the level
// (sim/request/raid-consumes.json), and a fresh 33's set holds no weapon stone.
const FURY_60 = `${FURY}|level=60`;

// The run line (spec 2026-11-06-sim §4.1) holds Fight, Length, Targets, Buffs and Precision;
// the rest of the request is in More settings (§5.2, ruling S7): the execute phase, the target
// level, the rotation, the consumables, Custom's buffs and the request JSON. The old settings
// card's variation, target armor, target type and dummy fields retired: a fight style ("Target
// dummy" among them) or the request JSON sets those now.
test('the run line reads the defaults and every control changes the settings', async ({ page }) => {
  // The paste runs at once, and the run line is inert while that run is in flight, so this
  // waits it out before touching any of them.
  await pasteAndRun(page);

  const settings = page.getByTestId('sim-settings');
  await expect(settings).toBeVisible();

  // Defaults: defaultSettings() (settings.ts) is Patchwerk, 3:00, single target, and the
  // preset the character's band publishes -- Bare ("None") for this level-33 Fury (owner
  // ruling, round 6).
  const style = page.getByTestId('control-style');
  await expect(style).toHaveValue('patchwerk');
  await expect(page.getByTestId('control-length')).toHaveValue('180');
  await expect(page.getByTestId('control-targets')).toHaveValue('1');
  await expect(page.getByTestId('control-buffs-bare')).toHaveAttribute('aria-checked', 'true');
  await expect(page.getByTestId('control-buffs-bare')).toHaveText('None');

  // Fight style: the contract's nine, plus nothing. Choosing one writes its fields.
  await expect(style.locator('option')).toHaveCount(9);
  await choose(style, 'cleave-3');
  await expect(page.getByTestId('control-targets')).toHaveValue('3');
  await choose(style, 'light-movement');
  await expect(page.getByTestId('control-targets')).toHaveValue('1');

  // Setting targets by hand detaches from the style: the select falls to "Custom".
  await choose(page.getByTestId('control-targets'), '4');
  await expect(style).toHaveValue('');

  // Fight length is labelled as a clock and reaches ten minutes.
  const duration = page.getByTestId('control-length');
  await expect(duration.locator('option', { hasText: '0:20' })).toHaveCount(1);
  await expect(duration.locator('option', { hasText: '10:00' })).toHaveCount(1);
  await choose(duration, '600');
  await expect(duration).toHaveValue('600');

  // The secondary controls live in More settings, closed on arrival.
  await expect(page.getByTestId('sim-more-settings')).toHaveCount(0);
  await openMoreSettings(page);
  const sheet = page.getByTestId('sim-more-settings');
  // The rotation names the spec the loaded character carries, Fury.
  await expect(sheet).toContainText('The Fury rotation');

  await choose(page.getByTestId('sim-target-level'), '60');
  await expect(page.getByTestId('sim-target-level')).toHaveValue('60');

  // Re-attaching a style, then toggling the execute phase, detaches it again: withExecutePhase
  // calls settings.ts's `detached`, the same rule targets used above.
  await closeMoreSettings(page);
  await choose(style, 'patchwerk');
  await expect(style).toHaveValue('patchwerk');
  await openMoreSettings(page);
  await expect(page.getByTestId('sim-execute-on')).toHaveAttribute('aria-checked', 'true');
  await page.getByTestId('sim-execute-off').click();
  await expect(page.getByTestId('sim-execute-off')).toHaveAttribute('aria-checked', 'true');
  await closeMoreSettings(page);
  await expect(style).toHaveValue('');

  // Done keeps the changes on the run line: the length set by hand holds through it all.
  await expect(duration).toHaveValue('600');
  await choose(page.getByTestId('control-targets'), '4');
  await expect(page.getByTestId('control-targets')).toHaveValue('4');
});

test('the run line is absent until a character is loaded', async ({ page }) => {
  await gotoSim(page);
  await expect(page.getByTestId('sim-settings')).toHaveCount(0);
  await expect(page.getByTestId('sim-landing')).toBeVisible();
});

/**
 * Task 4 (dps-minmaxer D16 BLOCKER): the old Raid-buffed preset was five buffs and two
 * consumables, +6% over Solo where vanilla's real answer for a melee is +60% to +120%.
 * This reads the fix from the same three places a reviewer's repro read the bug --
 * what is in it (More settings' consumables), the Custom panel's own counters, and the
 * request the browser actually sends -- rather than only a unit test on settings.ts.
 */
test('Raid-buffed is the full standard set: More settings lists it, Custom’s counters show it, and the request carries it', async ({
  page,
}) => {
  await pasteAndRun(page, FURY_60);
  await expect(page.getByTestId('control-buffs-raid-buffed')).toHaveAttribute('aria-checked', 'true');
  await expect(page.getByTestId('control-buffs-raid-buffed')).toHaveText('Raid');

  // More settings names the preset's consumables, humanised, never bare ids.
  await openMoreSettings(page);
  const consumables = page.getByTestId('sim-consumables');
  await expect(consumables).toBeVisible();
  await expect(consumables).not.toContainText('_');
  await closeMoreSettings(page);

  // Switching to Custom carries every tick across (withPreset(..., 'custom')), so the
  // panel's own counters read the real totals now -- the findings' own quote was "RAID
  // BUFFS 4/30, ... ON THE TARGET 0/24, ... WEAPON OILS AND STONES 0/48".
  await chooseBuffs(page, 'custom');
  await openMoreSettings(page);
  await expect(page.getByTestId('sim-buff-panel')).toBeVisible();
  const groupCount = async (group: string): Promise<number> => {
    // The group's Disclosure summary ends with its "on/size" count (ui/Disclosure's meta).
    const text = await page.getByTestId(`sim-buff-group-${group}`).locator('summary').textContent();
    return Number(/(\d+)\/\d+/.exec(text ?? '')?.[1] ?? 0);
  };
  expect(await groupCount('raid-buffs')).toBeGreaterThan(4);
  expect(await groupCount('debuffs')).toBeGreaterThan(0);
  expect(await groupCount('weapon-imbue')).toBeGreaterThan(0);
  await closeMoreSettings(page);

  // The request the browser would actually send carries the expanded preset -- the request
  // drawer's own rendering of it, not only the settings object a unit test can see.
  await chooseBuffs(page, 'raid-buffed');
  await openRequest(page);
  await expect(page.getByTestId('sim-request-buffs')).toContainText('sunder_armor');
  await expect(page.getByTestId('sim-request-buffs')).toContainText('heart_of_the_lion');
  await expect(page.getByTestId('sim-request-json')).toHaveValue(
    /main_hand_imbue:elemental_sharpening_stone/,
  );
});
