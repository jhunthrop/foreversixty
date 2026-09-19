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

async function loadFury(page: Page): Promise<void> {
  await page.goto('/sim');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();
}

test('the run control moves idle -> running -> done, and a cancel keeps the last figure', async ({
  page,
}) => {
  await loadFury(page);

  const button = page.getByTestId('sim-run-button');
  await expect(button).toHaveText('Run sim');
  await expect(page.getByTestId('sim-dps')).toHaveText('—');
  await expect(page.getByTestId('sim-progress-bar')).toHaveCount(0);

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
  await expect(button).toHaveText('Run sim');

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
  await page.goto('/sim');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();

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
