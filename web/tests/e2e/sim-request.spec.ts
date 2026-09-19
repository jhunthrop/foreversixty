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
  await expect(page.getByTestId('sim-request-errors')).toContainText(simCopy.requestNotJson);

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

test('sharing a request past the URL budget shows the error, not a blank field', async ({ page }) => {
  // The encoder (url.ts's encodeRequestParam) exists now (Task 16), so the only way
  // shareUrlFor still answers null is a request genuinely past MAX_REQUEST_PARAM -- padded
  // here with an oversized settings clause the engine happily accepts before the drawer
  // even asks for a share link (settings aren't re-validated on share, only on
  // apply/run/share's own verify() -- so this exercises the size gate on its own, not the
  // engine's).
  await page.goto('/sim');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();

  await page.getByTestId('sim-request-drawer').locator('summary').click();
  const editor = page.getByTestId('sim-request-json');
  const request = JSON.parse(await editor.inputValue());
  // Padding source.ref (types.ts: "character_key | "" | build id | ...", a free string,
  // never validated against a known format) pushes well past the size gate without
  // changing what the engine simulates, so this stays a request the engine's own
  // Validate accepts -- the drawer's verify() must pass before onshare ever runs.
  request.source.ref = 'x'.repeat(20_000);
  await editor.fill(JSON.stringify(request, null, 2));
  await page.getByTestId('sim-request-share').click();
  await expect(page.getByTestId('sim-request-share-error')).toHaveText(simCopy.requestShareTooLong);
  await expect(page.getByTestId('sim-request-share-link')).toHaveCount(0);
});

test('a shared request link reproduces the whole page state', async ({ page }) => {
  await page.goto('/sim');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();

  await page.getByTestId('sim-style').selectOption('cleave-5');
  await page.getByTestId('sim-duration').selectOption('600');
  await page.getByTestId('sim-request-drawer').locator('summary').click();
  // The drawer seeds its textarea once, from whatever request existed at mount, and then
  // "owns" it for the player (RequestDrawer.svelte's own comment) -- it does not
  // re-seed on every settings change, since that would throw away an in-progress edit.
  // Reset is the drawer's own way back to the page's current request, so this shares
  // the style and duration just picked rather than the pre-mount defaults.
  await page.getByTestId('sim-request-reset').click();
  await page.getByTestId('sim-request-share').click();

  const url = await page.getByTestId('sim-request-share-link').inputValue();
  expect(url).toContain('/sim?req=');

  await page.goto(url);
  await expect(page.getByTestId('sim-character')).toBeVisible();
  await expect(page.getByTestId('sim-targets')).toHaveValue('5');
  await expect(page.getByTestId('sim-duration')).toHaveValue('600');
});
