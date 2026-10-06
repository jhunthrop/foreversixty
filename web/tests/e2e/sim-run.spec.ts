import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test, type Page } from '@playwright/test';
import { simCopy } from '../../src/lib/sim/copy';

// Same fixture as sim-settings.spec.ts and sim-sources.spec.ts: an addon export needs no
// API stub, so this reads the site's own active build id rather than hardcoding one.
const activeBuild = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
) as { build: string };

const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=11726`;

// sim-one-paste lane: `loadFury` now goes through ColdPasteHero.svelte, the view's one
// paste surface -- SourceSwitcher's own addon card is gone whenever the hero renders
// (`showAddonCard`). The hero pastes and runs in one click (SimView.svelte's
// `runPastedInput`), so this helper's own postcondition changed from "character loaded,
// never run" to "character loaded, one run already finished" -- waited out here so every
// caller below starts from a settled `sim-run-button` rather than racing the hero's own
// background run.
async function loadFury(page: Page): Promise<void> {
  await page.goto('/sim');
  await page.getByTestId('sim-cold-paste-input').fill(FURY);
  await page.getByTestId('sim-cold-paste-run').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();
  await expect(page.getByTestId('sim-run-button')).toHaveText('Run again', { timeout: 10_000 });
}

test('a run moves running -> done, and a cancel keeps the last figure', async ({ page }) => {
  // The hero's own run (inside loadFury) already proved idle -> running -> done once --
  // sim-cold-paste.spec.ts covers that first run start to finish. This test picks up from
  // its "Run again" state and proves the mechanism survives a second run and a cancel.
  await loadFury(page);

  const button = page.getByTestId('sim-run-button');
  await button.click();
  await expect(button).toHaveText('Stop');
  await expect(page.getByTestId('sim-dps')).not.toHaveText('—');
  await expect(page.getByTestId('sim-progress')).toHaveText(/\d+ of 3,000 iterations/);

  await expect(button).toHaveText('Run again', { timeout: 5000 });
  await expect(page.getByTestId('sim-progress')).toHaveText(/^3,000 iterations/);
  await expect(page.getByTestId('sim-error')).toHaveText(/^± \d/);

  const finishedFigure = await page.getByTestId('sim-dps').textContent();

  // Run again, and stop it before it has a chance to move the figure: the fake engine's
  // first tick lands ~120ms out (engine-fake.ts's own default), which is long past the two
  // clicks below.
  await button.click();
  await button.click();
  await expect(page.getByTestId('sim-message')).toHaveText('Stopped.');
  await expect(page.getByTestId('sim-dps')).toHaveText(finishedFigure ?? '');
});

// The controller addendum (Task 12): the settings bar's own disabled prop is wired to
// store.phase === 'running', which Task 12 could not exercise without a run button.
test('every settings control is disabled while the run control reads Stop, and re-enabled after', async ({
  page,
}) => {
  await loadFury(page);

  const button = page.getByTestId('sim-run-button');
  await button.click();
  await expect(button).toHaveText('Stop');

  await expect(page.getByTestId('sim-duration')).toBeDisabled();
  await expect(page.getByTestId('sim-targets')).toBeDisabled();
  await expect(page.getByTestId('sim-execute')).toBeDisabled();
  await expect(page.getByTestId('sim-preset')).toBeDisabled();

  await button.click();
  // "Run again", not "Run sim": loadFury's own hero run already finished once, so
  // `runAndSettle` (store-request.ts) restores this cancelled run's phase to 'done'
  // rather than 'idle' -- the same "the last completed result wins" rule the sibling
  // test above proves for the figure itself.
  await expect(button).toHaveText('Run again');

  await expect(page.getByTestId('sim-duration')).toBeEnabled();
  await expect(page.getByTestId('sim-targets')).toBeEnabled();
  await expect(page.getByTestId('sim-execute')).toBeEnabled();
  await expect(page.getByTestId('sim-preset')).toBeEnabled();
});

test('high precision runs 10,000 iterations instead of 3,000', async ({ page }) => {
  await loadFury(page);

  await page.getByTestId('sim-precision').selectOption('high');
  await page.getByTestId('sim-run-button').click();
  await expect(page.getByTestId('sim-progress')).toHaveText(/^10,000 iterations/, { timeout: 8000 });
});

test('the server lane is not offered to a signed-out visitor', async ({ page }) => {
  await loadFury(page);

  await expect(page.getByTestId('sim-run')).toBeVisible();
  await expect(page.getByTestId('sim-server-run')).toHaveCount(0);
});

test('the precision select offers four choices and the details card states the run', async ({ page }) => {
  await loadFury(page);

  const precision = page.getByTestId('sim-precision');
  await expect(precision).toHaveValue('normal');
  await expect(precision.locator('option')).toHaveCount(4);
  await expect(page.getByTestId('sim-target-error')).toBeHidden();

  await precision.selectOption('target-error');
  await expect(page.getByTestId('sim-target-error')).toContainText('30,000');

  // Fast keeps the browser suite quick; the card is the same card at every precision.
  await precision.selectOption('fast');
  await page.getByTestId('sim-run-button').click();
  await expect(page.getByTestId('sim-details-card')).toBeVisible({ timeout: 30_000 });
  await expect(page.getByTestId('sim-details-iterations')).toHaveText('500');
  await expect(page.getByTestId('sim-details-margin')).toContainText(simCopy.dps);
  await expect(page.getByTestId('sim-details-margin')).toContainText('%');
  await expect(page.getByTestId('sim-details-lane')).toHaveText(simCopy.detailsLaneBrowser);
  await expect(page.getByTestId('sim-details-engine')).toHaveAttribute('href', '/sim/specs');
  await expect(page.getByTestId('sim-progress')).toContainText('%');
});

test('a finished run can be named, and the saved link opens in a new tab', async ({ page }) => {
  // Persona review 2026-10-06 (retail-raider), §5/§8 item 1: SimSavePanel's Save button
  // now requires sign-in -- a cold-paste run reaching a result stays covered, signed out,
  // by sim-cold-paste.spec.ts; this test opens the save form, so it needs a signed-in
  // `/v1/me`.
  await page.route('**/v1/me', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({
        ok: true,
        data: {
          user: {
            id: 1,
            battletag: 'Fixture#0001',
            email: null,
            role: 'user',
            anonymize: false,
            premium: false,
          },
          characters: [],
          guilds: [],
        },
        error: null,
      }),
    }),
  );
  await page.goto('/sim');
  await page.getByTestId('sim-cold-paste-input').fill(FURY);
  await page.getByTestId('sim-cold-paste-run').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();
  // Waits out the hero's own run (normal precision) before picking fast and re-running,
  // the same reason `loadFury` above waits: the precision select is disabled while a run
  // is in flight.
  await expect(page.getByTestId('sim-run-button')).toHaveText('Run again', { timeout: 10_000 });
  await page.getByTestId('sim-precision').selectOption('fast');
  await page.getByTestId('sim-run-button').click();
  await expect(page.getByTestId('sim-details-card')).toBeVisible({ timeout: 30_000 });

  const title = page.getByTestId('sim-report-title');
  await expect(title).toHaveValue('Raid-buffed, 3:00, single target');
  await title.fill('Pre-raid, no world buffs');
  await page.getByTestId('sim-save-open').click();
  await expect(page.getByTestId('sim-save-title')).toHaveValue('Pre-raid, no world buffs');
});
