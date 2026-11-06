// web/tests/e2e/sim-results.spec.ts
// A finished run on /sim renders the saved page's report (RunReport.svelte, spec
// 2026-11-04-sim-results), fed from the run's own result: the band, the section nav, the damage
// table beside the spread, then the lazy cards. The run controls, the save panel and the
// settings stay where they were. sim-saved.spec.ts covers the report's cards in depth on
// /sim/<id>; this file covers what is particular to the live page.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { simCopy } from '../../src/lib/sim/copy';

// Same fixture as sim-run.spec.ts: an addon export needs no API stub, so this reads the site's
// own active build id rather than hardcoding one.
const activeBuild = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
) as { build: string };

const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=11726`;

// ColdPasteHero's Run loads and runs in one click. The fake engine (engine-fake.ts) answers
// with src/fixtures/sim/result.json's own summary, whatever DPS its seeded normal invents.
async function loadFuryAndRun(page: Page): Promise<void> {
  await page.goto('/sim');
  await page.getByTestId('sim-cold-paste-input').fill(FURY);
  await page.getByTestId('sim-cold-paste-run').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();
  await expect(page.getByTestId('sim-run-button')).toHaveText('Run again', { timeout: 10_000 });
  await expect(page.getByTestId('sim-live-report')).toBeVisible();
}

test('a finished run shows the report under the run control, with no tab strip', async ({ page }) => {
  await loadFuryAndRun(page);

  const report = page.getByTestId('sim-live-report');
  await expect(report.getByTestId('sim-result-band')).toBeVisible();
  await expect(report.getByTestId('sim-result-eyebrow')).toHaveText('RUN IN YOUR BROWSER');
  await expect(report.getByTestId('sim-report-dps')).toHaveText(/^\d[\d,]*\.\d$/);
  await expect(report.getByTestId('sim-fact-runs')).toBeVisible();
  await expect(report.getByTestId('sim-section-nav')).toBeVisible();
  await expect(report.getByTestId('sim-damage-card')).toBeVisible();
  await expect(report.getByTestId('sim-spread-card')).toBeVisible();

  // The old eight-tab view is gone, and no brushable time chart came back with the report.
  await expect(page.getByTestId('sim-tab-damage')).toHaveCount(0);
  await expect(page.getByTestId('sim-results')).toHaveCount(0);
  await expect(page.getByTestId('time-chart')).toHaveCount(0);

  // The run control comes first; the report sits between it and the save and settings panels.
  const runBox = await page.getByTestId('sim-run').boundingBox();
  const bandBox = await report.getByTestId('sim-result-band').boundingBox();
  const titleBox = await page.getByTestId('sim-report-title').boundingBox();
  expect(runBox !== null && bandBox !== null && titleBox !== null).toBe(true);
  expect(bandBox!.y).toBeGreaterThan(runBox!.y);
  expect(titleBox!.y).toBeGreaterThan(bandBox!.y);
});

test('the live band is not the page heading and carries none of the saved page actions', async ({ page }) => {
  await loadFuryAndRun(page);

  await expect(page.locator('h1')).toHaveCount(1);
  await expect(page.locator('h2[data-testid="sim-saved-title"]')).toBeVisible();
  await expect(page.getByTestId('sim-run-yourself')).toHaveCount(0);
  await expect(page.getByTestId('sim-saved-copy')).toHaveCount(0);
  // The run control keeps its own figure; the band's figure has its own test id.
  await expect(page.getByTestId('sim-dps')).toHaveCount(1);
  await expect(page.getByTestId('sim-error')).toHaveCount(1);
});

test('the damage table reads humanised names, never a raw engine key', async ({ page }) => {
  await loadFuryAndRun(page);

  const table = page.getByTestId('ability-table');
  await expect(table).toBeVisible();
  await expect(table).not.toContainText(/\bspell:/);
  await expect(table).not.toContainText(/\bSpell \d+\b/);
});

test('the deeper cards load as they near the viewport, down to the settings card', async ({ page }) => {
  await loadFuryAndRun(page);

  await page.locator('#timeline').scrollIntoViewIfNeeded();
  await expect(page.locator('#timeline')).toBeVisible();
  await page.getByTestId('sim-section-settings').scrollIntoViewIfNeeded();
  await expect(page.getByTestId('sim-settings-card')).toBeVisible({ timeout: 10_000 });
  await expect(page.getByTestId('sim-copy-json')).toBeVisible();
});

test('running again keeps one report, the new run in it', async ({ page }) => {
  await loadFuryAndRun(page);

  await page.getByTestId('sim-run-button').click();
  await expect(page.getByTestId('sim-run-button')).toHaveText('Run again', { timeout: 10_000 });
  await expect(page.getByTestId('sim-live-report')).toHaveCount(1);
  await expect(page.getByTestId('sim-result-band')).toHaveCount(1);
});

test('run=1 runs the linked build once it loads, with no press', async ({ page }) => {
  await page.goto(`/sim?code=${encodeURIComponent(FURY)}&run=1`);

  await expect(page.getByTestId('sim-character')).toBeVisible();
  await expect(page.getByTestId('sim-live-report')).toBeVisible({ timeout: 10_000 });
  await expect(page.getByTestId('sim-run-button')).toHaveText('Run again');
});

test('a plain code link still waits for the press', async ({ page }) => {
  await page.goto(`/sim?code=${encodeURIComponent(FURY)}`);

  await expect(page.getByTestId('sim-character')).toBeVisible();
  await expect(page.getByTestId('sim-run-button')).toHaveText(simCopy.run);
  await expect(page.getByTestId('sim-live-report')).toHaveCount(0);
});
