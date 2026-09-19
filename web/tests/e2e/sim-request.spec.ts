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
  // The drawer tracks the page's own current request until the player edits the textarea
  // (RequestDrawer.svelte's own comment, finding 2 of the final whole-branch review), so
  // an untouched open already carries the style and duration just picked -- no reset
  // needed to pick up settings changed before the drawer was ever opened.
  await page.getByTestId('sim-request-share').click();

  const url = await page.getByTestId('sim-request-share-link').inputValue();
  expect(url).toContain('/sim?req=');

  await page.goto(url);
  await expect(page.getByTestId('sim-character')).toBeVisible();
  await expect(page.getByTestId('sim-targets')).toHaveValue('5');
  await expect(page.getByTestId('sim-duration')).toHaveValue('600');
});

test('the drawer tracks page settings until the player edits, then keeps the edit', async ({ page }) => {
  // Finding 2, final whole-branch review: the drawer used to seed its textarea once, at
  // mount, from whatever request existed then -- so a settings change made after the
  // drawer had already mounted (which is as soon as a character loads, not when the
  // drawer is first opened) went stale in the textarea, and Share or Apply from it carried
  // the old encounter with nothing on screen saying so.
  await page.goto('/sim');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();

  // Change a setting before the drawer is ever opened -- the arrival-request case the
  // fix targets -- then open it and see the change already there.
  await page.getByTestId('sim-duration').selectOption('600');
  await page.getByTestId('sim-request-drawer').locator('summary').click();
  const editor = page.getByTestId('sim-request-json');
  await expect.poll(async () => JSON.parse(await editor.inputValue()).encounter.duration_sec).toBe(600);

  // The player's own edit stops the tracking: a further settings change must not
  // overwrite what they typed.
  const edited = JSON.stringify({ ...JSON.parse(await editor.inputValue()), iterations: 500 }, null, 2);
  await editor.fill(edited);
  await page.getByTestId('sim-duration').selectOption('300');
  await expect(editor).toHaveValue(edited);

  // Reset returns the textarea to tracking the page again.
  await page.getByTestId('sim-request-reset').click();
  await expect.poll(async () => JSON.parse(await editor.inputValue()).encounter.duration_sec).toBe(300);
});

test('Apply keeps a per-slot enchant and suffix', async ({ page }) => {
  await page.goto('/sim');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();

  await page.getByTestId('sim-request-drawer').locator('summary').click();
  const editor = page.getByTestId('sim-request-json');
  const request = JSON.parse(await editor.inputValue());
  request.character.gear = [{ slot: 'head', item_id: 12640, enchant: 2504, suffix: 1820 }];
  await editor.fill(JSON.stringify(request, null, 2));
  await page.getByTestId('sim-request-apply').click();

  // The drawer stopped tracking the page's request the moment the textarea above was
  // edited (finding 2, final whole-branch review), so Reset is what reads back the
  // request Apply actually built rather than the edit still sitting in the textarea.
  await expect(page.getByTestId('sim-character')).toBeVisible();
  await page.getByTestId('sim-request-reset').click();
  const applied = JSON.parse(await editor.inputValue());
  expect(applied.character.gear).toEqual([{ slot: 'head', item_id: 12640, enchant: 2504, suffix: 1820 }]);
});
