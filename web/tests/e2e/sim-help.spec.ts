// web/tests/e2e/sim-help.spec.ts
// Task 7 (newcomer BLOCKER: `[data-tooltip],[role=tooltip],abbr,.tooltip` was 0 on a phone;
// the two explanations /sim had were `title=`, which native touch never fires -- 100% of the
// page's help was unreachable off a mouse, and most controls had none at all). This proves
// HelpNote.svelte's own affordance -- opens, closes again, Escape closes it, every trigger
// clears the 44px floor at 390x844, and no `title=` remains on the page -- plus a sample of
// what the notes actually say, including the fight-style, execute-phase and target-armor
// content the brief calls out by name.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test, type Page } from '@playwright/test';
import { simCopy } from '../../src/lib/sim/copy';

// Same fixture sim-settings.spec.ts and sim-run.spec.ts load: an addon export needs no API
// stub.
const activeBuild = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
) as { build: string };
const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=11726`;

async function loadFury(page: Page): Promise<void> {
  await page.goto('/sim');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();
}

async function openMoreSettings(page: Page): Promise<void> {
  await page.getByTestId('sim-settings-more').locator('summary').click();
}

test('a HelpNote opens on a click, names the control it explains, and closes on a second click or Escape', async ({
  page,
}) => {
  await loadFury(page);

  const trigger = page.getByTestId('sim-style-help-trigger');
  await expect(trigger).toHaveText(simCopy.helpTrigger(simCopy.fightStyle));
  await expect(trigger).toHaveAttribute('aria-expanded', 'false');

  const panel = page.getByTestId('sim-style-help-panel');
  await expect(panel).toBeHidden();

  await trigger.click();
  await expect(trigger).toHaveAttribute('aria-expanded', 'true');
  await expect(panel).toBeVisible();

  await trigger.click();
  await expect(panel).toBeHidden();
  await expect(trigger).toHaveAttribute('aria-expanded', 'false');

  await trigger.click();
  await expect(panel).toBeVisible();
  await page.keyboard.press('Escape');
  await expect(panel).toBeHidden();
  await expect(trigger).toBeFocused();
});

test('Fight style’s help note names all nine styles, including what their names alone do not say', async ({
  page,
}) => {
  await loadFury(page);
  await page.getByTestId('sim-style-help-trigger').click();
  const panel = page.getByTestId('sim-style-help-panel');

  for (const label of Object.values(simCopy.styleLabel)) {
    await expect(panel).toContainText(label);
  }
  // Cleave's targets share Patchwerk's own execute threshold, and the dungeon schedule is
  // fixed at 160 seconds regardless of Fight length -- neither is obvious from the option's
  // name alone (task-7-brief.md's own warning against writing flattering copy over it).
  await expect(panel).toContainText('sharing the same 25% execute threshold as Patchwerk');
  await expect(panel).toContainText('fixed schedule, not a repeating pull');
});

test('Execute phase’s note states the current run’s state, not just the concept', async ({ page }) => {
  await loadFury(page);
  // Patchwerk is the default style; its execute_ratio (0.25) is why the box opens checked.
  await expect(page.getByTestId('sim-execute')).toBeChecked();
  await openMoreSettings(page);

  const panel = page.getByTestId('sim-execute-help-panel');
  await page.getByTestId('sim-execute-help-trigger').click();
  await expect(panel).toContainText(
    'On: this run currently simulates an execute phase below 25% target health, under Patchwerk.',
  );

  // Unticking it by hand detaches the encounter from its style (settings.ts's `detached`),
  // so the style reads Custom -- the note names whichever style actually governs it.
  await page.getByTestId('sim-execute').uncheck();
  await expect(panel).toContainText('Off: Custom has no execute phase, so this run has none.');
});

test('Target armor’s note says whether the placeholder is in effect or overridden', async ({ page }) => {
  await loadFury(page);
  await openMoreSettings(page);

  const panel = page.getByTestId('sim-target-armor-help-panel');
  await page.getByTestId('sim-target-armor-help-trigger').click();
  await expect(panel).toContainText('Empty; the engine uses 3,731, the preset for this level instead.');

  // The input's own onchange (withTargetArmor) fires on blur, not on fill() alone.
  await page.getByTestId('sim-target-armor').fill('5000');
  await page.getByTestId('sim-target-armor').blur();
  await expect(panel).toContainText('Set to 5,000, overriding 3,731, the preset for this level.');
});

test('Buffs keeps its own help note distinct from "what’s in it"', async ({ page }) => {
  await loadFury(page);
  await expect(page.getByTestId('sim-preset-help-trigger')).toBeVisible();
  await expect(page.getByTestId('sim-preset-summary-trigger')).toBeVisible();

  await page.getByTestId('sim-preset-help-trigger').click();
  await expect(page.getByTestId('sim-preset-help-panel')).toContainText(
    'Which buffs and consumables the run applies',
  );
});

test('the finish-notification checkbox explains itself too, once a result is on screen', async ({ page }) => {
  await loadFury(page);
  await page.getByTestId('sim-run-button').click();
  await expect(page.getByTestId('sim-run-button')).toHaveText('Run again', { timeout: 8000 });

  const trigger = page.getByTestId('sim-notify-help-trigger');
  await expect(trigger).toBeVisible();
  await trigger.click();
  await expect(page.getByTestId('sim-notify-help-panel')).toContainText(
    'Asks the browser for permission to show a notification',
  );
});

test.describe('phone (390x844)', () => {
  // 390x844, the exact width and height task-2-brief.md's own acceptance criterion names
  // and sim-scope.spec.ts and sim-tabs.spec.ts already use the same way: an explicit
  // viewport rather than --project=mobile, so the file runs the same under either project.
  test.use({ viewport: { width: 390, height: 844 } });

  test('no title= attribute remains on the page, and every help trigger clears the 44px floor', async ({
    page,
  }) => {
    await loadFury(page);
    await openMoreSettings(page);

    // The newcomer repro this task fixes, verbatim: querying [title] rather than the wider
    // selector the finding used, since HelpNote's own trigger and panel carry neither
    // data-tooltip nor role=tooltip nor .tooltip -- title= is the thing that had to go.
    const titleCount = await page.evaluate(() => document.querySelectorAll('[title]').length);
    expect(titleCount).toBe(0);

    const triggers = page.locator('[data-testid$="-help-trigger"]');
    const count = await triggers.count();
    // Fight style, Fight length, Targets, Buffs, Length varies by, Target level, Target
    // armor, Target type, Execute phase, Target dummy, Precision -- eleven before a run.
    expect(count).toBeGreaterThanOrEqual(11);
    for (let i = 0; i < count; i++) {
      const box = await triggers.nth(i).boundingBox();
      expect(box, `trigger ${i} has a box`).not.toBeNull();
      expect(box!.width, `trigger ${i} width`).toBeGreaterThanOrEqual(44);
      expect(box!.height, `trigger ${i} height`).toBeGreaterThanOrEqual(44);
    }
  });

  test('a tap reveals a help note’s text, a second tap hides it again, and Escape closes it', async ({
    page,
  }) => {
    await loadFury(page);

    const trigger = page.getByTestId('sim-duration-help-trigger');
    const panel = page.getByTestId('sim-duration-help-panel');
    await expect(panel).toBeHidden();

    await trigger.click();
    await expect(panel).toBeVisible();
    await expect(panel).toContainText(simCopy.fightLengthHelp);

    await trigger.click();
    await expect(panel).toBeHidden();

    await trigger.click();
    await expect(panel).toBeVisible();
    await page.keyboard.press('Escape');
    await expect(panel).toBeHidden();
    await expect(trigger).toBeFocused();
  });
});
