// web/tests/e2e/sim-results.spec.ts
// A finished run on /sim renders the saved page's report body (RunReport.svelte, spec
// 2026-11-06-sim §4.5), fed from the run's own result: the answer in the head carries the
// figure, the run line stays above, then the damage table beside the spread and the lazy
// sections. sim-saved.spec.ts covers the report's sections in depth on /sim/<id>; this file
// covers what is particular to the live page.
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { simPageCopy } from '../../src/lib/sim/page-copy';
import { gotoSim, openSimSection, pasteAndRun, SIM_FURY as FURY, simCharacter } from './support/sim-page';

// The paste Plate's Run loads and runs in one click. The fake engine (engine-fake.ts) answers
// with src/fixtures/sim/result.json's own summary, whatever DPS its seeded normal invents.
async function loadFuryAndRun(page: Page): Promise<void> {
  await pasteAndRun(page);
  await expect(page.getByTestId('sim-live-report')).toBeVisible();
}

test('a finished run shows the report under the run line, with no tab strip', async ({ page }) => {
  await loadFuryAndRun(page);

  const report = page.getByTestId('sim-live-report');
  await expect(page.getByTestId('sim-answer')).toHaveAttribute('data-state', 'result');
  await expect(page.getByTestId('sim-answer-label')).toHaveText('');
  await expect(page.getByTestId('sim-dps')).toHaveText(/^\s*\d[\d,]*\.\d\s*DPS\s*$/);
  await expect(report.getByTestId('sim-damage-card')).toBeVisible();
  await expect(report.getByTestId('sim-spread-card')).toBeVisible();
  // The sticky section nav and the result band retired (ruling S4, spec §4.5).
  await expect(page.getByTestId('sim-section-nav')).toHaveCount(0);
  await expect(page.getByTestId('sim-result-band')).toHaveCount(0);

  // The old eight-tab view is gone, and no brushable time chart came back with the report.
  await expect(page.getByTestId('sim-tab-damage')).toHaveCount(0);
  await expect(page.getByTestId('sim-results')).toHaveCount(0);
  await expect(page.getByTestId('time-chart')).toHaveCount(0);

  // The head (with the answer) comes first, then the run line, then the report. The run line is
  // sticky, so this reads the order from the top of the page.
  await page.evaluate(() => window.scrollTo(0, 0));
  const headBox = await page.getByTestId('sim-answer').boundingBox();
  const barBox = await page.getByTestId('sim-settings').boundingBox();
  const damageBox = await report.getByTestId('sim-damage-card').boundingBox();
  expect(headBox !== null && barBox !== null && damageBox !== null).toBe(true);
  expect(barBox!.y).toBeGreaterThan(headBox!.y);
  expect(damageBox!.y).toBeGreaterThan(barBox!.y);
});

test('the live page has one heading and carries none of the saved page actions', async ({ page }) => {
  await loadFuryAndRun(page);

  await expect(page.locator('h1')).toHaveCount(1);
  await expect(page.locator('h1')).toHaveText('Fury Warrior');
  await expect(page.getByTestId('sim-run-yourself')).toHaveCount(0);
  await expect(page.getByTestId('sim-saved-copy')).toHaveCount(0);
  // The answer is the page's one figure (spec §4.5: the RunControl figure retired).
  await expect(page.getByTestId('sim-dps')).toHaveCount(1);
  await expect(page.getByTestId('sim-error')).toHaveCount(1);
});

test('the damage table reads humanised names, never a raw engine key', async ({ page }) => {
  await loadFuryAndRun(page);

  const table = page.getByTestId('sim-damage-card');
  await expect(table).toBeVisible();
  await expect(table.getByTestId('ability-name').first()).toBeVisible();
  await expect(table).not.toContainText(/\bspell:/);
  await expect(table).not.toContainText(/\bSpell \d+\b/);
});

test('the deeper sections load as they near the viewport, down to This run', async ({ page }) => {
  await loadFuryAndRun(page);

  await page.getByTestId('sim-section-timeline').scrollIntoViewIfNeeded();
  await expect(page.getByTestId('sim-section-timeline')).toBeVisible();
  await openSimSection(page, 'settings');
  await expect(page.getByTestId('sim-settings-card')).toBeVisible({ timeout: 10_000 });
  await expect(page.getByTestId('sim-copy-json')).toBeVisible();
});

test('running again keeps one report, the new run in it', async ({ page }) => {
  await loadFuryAndRun(page);

  await page.getByTestId('sim-run-button').click();
  await expect(page.getByTestId('sim-run-button')).toHaveText('Run again', { timeout: 10_000 });
  await expect(page.getByTestId('sim-live-report')).toHaveCount(1);
  await expect(page.getByTestId('sim-damage-card')).toHaveCount(1);
});

test('run=1 runs the linked build once it loads, with no press', async ({ page }) => {
  await gotoSim(page, `/sim?code=${encodeURIComponent(FURY)}&run=1`);

  await expect(simCharacter(page)).toBeVisible();
  await expect(page.getByTestId('sim-live-report')).toBeVisible({ timeout: 10_000 });
  await expect(page.getByTestId('sim-run-button')).toHaveText('Run again');
});

test('a plain code link still waits for the press', async ({ page }) => {
  await gotoSim(page, `/sim?code=${encodeURIComponent(FURY)}`);

  await expect(simCharacter(page)).toBeVisible();
  await expect(page.getByTestId('sim-run-button')).toHaveText(simPageCopy.bar.runSim);
  await expect(page.getByTestId('sim-answer-label')).toHaveText('');
  await expect(page.getByTestId('sim-live-report')).toHaveCount(0);
});
