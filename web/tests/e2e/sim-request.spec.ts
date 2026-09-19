import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test } from '@playwright/test';
import { simCopy } from '../../src/lib/sim/copy';

const activeBuild = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
) as { build: string };

const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=11726`;

test('the drawer shows the exact request, validates an edit and runs it', async ({ page }) => {
  await page.goto('/sim');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();

  const drawer = page.getByTestId('sim-request-drawer');
  await drawer.locator('summary').click();
  const editor = page.getByTestId('sim-request-json');
  const json = await editor.inputValue();
  expect(JSON.parse(json)).toMatchObject({ spec: 'warrior-fury', iterations: 3000 });

  // The engine's own Validate, with the field it refuses named beside the message.
  await editor.fill(JSON.stringify({ ...JSON.parse(json), spec: '' }, null, 2));
  await page.getByTestId('sim-request-run').click();
  await expect(page.getByTestId('sim-request-error-spec')).toBeVisible();

  // Syntax errors are caught before the engine is asked anything.
  await editor.fill('{ "spec": ');
  await page.getByTestId('sim-request-run').click();
  await expect(page.getByTestId('sim-request-errors')).toContainText('not JSON');

  // A legal edit runs exactly as written: 500 iterations, whatever the precision select says.
  await editor.fill(JSON.stringify({ ...JSON.parse(json), iterations: 500 }, null, 2));
  await page.getByTestId('sim-request-run').click();
  await expect(page.getByTestId('sim-request-valid')).toBeVisible();
  await expect(page.getByTestId('sim-details-card')).toBeVisible({ timeout: 30_000 });
  await expect(page.getByTestId('sim-details-iterations')).toHaveText('500');
});

test('a pasted request loads the page state', async ({ page }) => {
  await page.goto('/sim');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();

  await page.getByTestId('sim-request-drawer').locator('summary').click();
  const editor = page.getByTestId('sim-request-json');
  const request = JSON.parse(await editor.inputValue());
  request.encounter.duration_sec = 600;
  request.encounter.targets = 5;
  request.character.buffs = ['thorns'];
  request.iterations = 500;
  await editor.fill(JSON.stringify(request, null, 2));
  await page.getByTestId('sim-request-apply').click();

  await expect(page.getByTestId('sim-duration')).toHaveValue('600');
  await expect(page.getByTestId('sim-targets')).toHaveValue('5');
  await expect(page.getByTestId('sim-preset')).toHaveValue('custom');
  await expect(page.getByTestId('sim-precision')).toHaveValue('fast');
  await expect(page.getByTestId('sim-request-buffs')).toHaveText('thorns');
});

test('sharing a request with nowhere to encode to yet shows the error, not a blank field', async ({
  page,
}) => {
  // Task 16 owns the real encoder (url.ts's encodeRequestParam); until it lands,
  // SimView's shareUrlFor honestly answers null, and this is the one path that is
  // actually true today -- a share that returns null must show the error, never a blank
  // or missing field.
  await page.goto('/sim');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();

  await page.getByTestId('sim-request-drawer').locator('summary').click();
  await page.getByTestId('sim-request-share').click();
  await expect(page.getByTestId('sim-request-share-error')).toHaveText(simCopy.requestShareTooLong);
  await expect(page.getByTestId('sim-request-share-link')).toHaveCount(0);
});
