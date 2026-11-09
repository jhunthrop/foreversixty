import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { choose } from './support/select';
import { openMoreSettings, pasteAndRun } from './support/sim-page';

// The landing's paste Plate loads and runs in one click (SimView.svelte's `runPastedInput`),
// so this helper's postcondition is "character loaded, one run already finished": every caller
// below starts from a settled `sim-run-button` rather than racing that first run.
async function loadFury(page: Page): Promise<void> {
  await pasteAndRun(page);
}

const answer = (page: Page) => page.getByTestId('sim-answer');

test('a run moves running -> done, and a stop keeps the partial figure', async ({ page }) => {
  // The paste's own run (inside loadFury) already proved idle -> running -> done once --
  // sim-cold-paste.spec.ts covers that first run start to finish. This test picks up from
  // its "Run again" state and proves the mechanism survives a second run and a stop.
  await loadFury(page);

  const button = page.getByTestId('sim-run-button');
  await button.click();
  await expect(button).toHaveText('Stop');
  await expect(answer(page)).toHaveAttribute('data-state', 'running');
  await expect(page.getByTestId('sim-error')).toHaveText(/[\d,]+ of 3,000 runs/);

  await expect(button).toHaveText('Run again', { timeout: 5000 });
  await expect(answer(page)).toHaveAttribute('data-state', 'result');
  await expect(page.getByTestId('sim-error')).toHaveText(/^± \d.*\(95%\)$/);
  await expect(page.getByTestId('sim-dps')).toHaveText(/^\s*[\d,]+\.\d\s*DPS\s*$/);

  // Run again at High (10,000 runs) and stop it once a partial figure is up: the number stays,
  // greyed, labelled with how far it got (spec 2026-11-06-sim §5.1), and is not saved.
  await choose(page.getByTestId('control-precision'), 'high');
  await button.click();
  await expect(page.getByTestId('sim-dps')).toHaveText(/\d/);
  await button.click();
  await expect(answer(page)).toHaveAttribute('data-state', 'stopped');
  await expect(page.getByTestId('sim-answer-label')).toHaveText(/^Stopped at [\d,]+ of 10,000 runs$/);
  await expect(page.getByTestId('sim-dps')).toHaveText(/^\s*about [\d,]+\s*DPS\s*$/);
  await expect(page.getByTestId('sim-error')).toContainText('A stopped run is not saved.');
  await expect(button).toHaveText('Run again');
  await expect(page.getByTestId('sim-save-open')).toHaveCount(0);
});

// The run line's groups (and More settings) are inert while a run is in flight (spec §4.4, the
// shipped rule); the primary stays live as Stop.
test('every run-line control is inert while the primary reads Stop, and live again after', async ({
  page,
}) => {
  await loadFury(page);

  const groups = page.locator('#sim-bar-groups');
  const button = page.getByTestId('sim-run-button');
  await expect(groups).not.toHaveAttribute('inert', /.*/);
  await button.click();
  await expect(button).toHaveText('Stop');

  await expect(groups).toHaveAttribute('inert', /.*/);
  await expect(groups).toHaveAttribute('data-busy', '');
  await expect(groups).toHaveAttribute('title', 'Stop the run to change it');
  for (const id of [
    'control-style',
    'control-length',
    'control-targets',
    'control-buffs',
    'control-precision',
  ]) {
    await expect(groups.getByTestId(id).first()).toBeAttached();
  }
  await expect(groups.getByTestId('sim-settings-more')).toBeAttached();

  await button.click();
  // "Run again", not "Run sim": loadFury's own run already finished once.
  await expect(button).toHaveText('Run again');
  await expect(groups).not.toHaveAttribute('inert', /.*/);
  await expect(groups).not.toHaveAttribute('data-busy', /.*/);
});

test('high precision runs 10,000 runs instead of 3,000', async ({ page }) => {
  await loadFury(page);

  await choose(page.getByTestId('control-precision'), 'high');
  await page.getByTestId('sim-run-button').click();
  await expect(page.getByTestId('sim-error')).toHaveText(/of 10,000 runs/, { timeout: 8000 });
});

test('the server lane is not offered to a signed-out visitor', async ({ page }) => {
  await loadFury(page);

  await expect(page.getByTestId('sim-settings')).toBeVisible();
  await openMoreSettings(page);
  await expect(page.getByTestId('sim-server-run')).toHaveCount(0);
});

test('the precision select offers four choices and the report states the run', async ({ page }) => {
  await loadFury(page);

  const precision = page.getByTestId('control-precision');
  await expect(precision).toHaveValue('normal');
  await expect(precision.locator('option')).toHaveCount(4);
  // Precision reads in runs, the one count the page uses (spec §4.1).
  await expect(precision.locator('option')).toHaveText([
    'Fast, 500 runs',
    'Normal, 3,000 runs',
    'High, 10,000 runs',
    'Until ±0.5%',
  ]);

  // Fast keeps the browser suite quick; the report is the same report at every precision.
  await choose(precision, 'fast');
  await page.getByTestId('sim-run-button').click();
  const report = page.getByTestId('sim-live-report');
  await expect(report).toBeVisible({ timeout: 30_000 });
  // This run loads as it nears the viewport (LazySection).
  await report.getByTestId('sim-section-settings').scrollIntoViewIfNeeded();
  await expect(report.getByTestId('sim-settings-runs')).toContainText('500');
  await expect(page.getByTestId('sim-error')).toContainText('(95%)');
  await expect(page.getByTestId('sim-answer-label')).toHaveText('Run in your browser');
  await expect(report.getByTestId('sim-spread-card')).toContainText('500 runs');
});

test('a finished run can be named in the save form', async ({ page }) => {
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
  // Waits out the paste's own run (normal precision) before picking fast and re-running, the
  // same reason `loadFury` above waits: the run line is inert while a run is in flight.
  await loadFury(page);
  await choose(page.getByTestId('control-precision'), 'fast');
  await page.getByTestId('sim-run-button').click();
  await expect(page.getByTestId('sim-live-report')).toBeVisible({ timeout: 30_000 });

  // The title lives in the save form now (spec §4.5: ReportOptions retired), pre-filled with
  // the run's composed name and editable before Save.
  await page.getByTestId('sim-save-open').click();
  const title = page.getByTestId('sim-save-title');
  await expect(title).toHaveValue('Bare, 3:00, single target');
  await title.fill('Pre-raid, no world buffs');
  await expect(title).toHaveValue('Pre-raid, no world buffs');
  await expect(page.getByTestId('sim-save-confirm')).toBeEnabled();
});
