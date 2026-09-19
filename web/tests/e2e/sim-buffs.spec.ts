import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test } from '@playwright/test';

const activeBuild = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
) as { build: string };

const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=11726`;

test('Custom opens the whole vocabulary, grouped, and every tick reaches the request', async ({ page }) => {
  // The panel's checkbox changes settings.buffs, and settings.buffs only reaches the engine
  // through the ToWorker message run() posts to the pool (worker.ts's `{ kind: 'run', ...,
  // request }`, a JSON string of the SimRequest). There is no `sim-request-buffs` test id
  // yet -- Task 15 adds the request drawer -- so this intercepts that postMessage instead
  // of reading a control's own state back at itself: a real assertion that fails if toggling
  // thorns did not actually change what the run sends, not a tautology that always passes.
  await page.addInitScript(() => {
    (window as unknown as { __simMessages: unknown[] }).__simMessages = [];
    const proto = window.Worker.prototype;
    const original = proto.postMessage;
    proto.postMessage = function postMessage(this: Worker, message: unknown, transfer?: unknown) {
      (window as unknown as { __simMessages: unknown[] }).__simMessages.push(message);
      return (original as (message: unknown, transfer?: unknown) => void).call(this, message, transfer);
    };
  });

  await page.goto('/sim');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();

  // The panel does not exist until Custom is chosen: a preset is a preset.
  await expect(page.getByTestId('sim-buff-panel')).toBeHidden();
  await page.getByTestId('sim-preset').selectOption('custom');
  await expect(page.getByTestId('sim-buff-panel')).toBeVisible();

  // Every group the design names has a section, and each one is closed on arrival.
  for (const group of [
    'raid-buffs',
    'party-buffs',
    'player-buffs',
    'world-buffs',
    'debuffs',
    'flask',
    'battle-elixir',
    'guardian-elixir',
    'food',
    'weapon-imbue',
    'potion',
    'explosive',
  ]) {
    await expect(page.getByTestId(`sim-buff-group-${group}`)).toBeVisible();
  }
  await expect(page.getByTestId('sim-buff-thorns')).toBeHidden();

  const raid = page.getByTestId('sim-buff-group-raid-buffs');
  await raid.locator('summary').click();
  await expect(page.getByTestId('sim-buff-thorns')).toBeVisible();
  await page.getByTestId('sim-buff-thorns').check();
  await expect(page.getByTestId('sim-buff-thorns')).toBeChecked();

  // A world buff is in its own section, not among the blessings.
  const world = page.getByTestId('sim-buff-group-world-buffs');
  await world.locator('summary').click();
  await expect(page.getByTestId('sim-buff-songflower_serenade')).toBeVisible();

  // The choice reaches the engine: run at the fastest precision and read back every request
  // string the run actually posted to a worker.
  await page.getByTestId('sim-precision').selectOption('fast');
  await page.getByTestId('sim-run-button').click();
  await expect(page.getByTestId('sim-details-card')).toBeVisible({ timeout: 30_000 });

  const sawThorns = await page.evaluate(() => {
    const messages = (window as unknown as { __simMessages: unknown[] }).__simMessages ?? [];
    return messages.some(
      (entry) =>
        typeof entry === 'object' &&
        entry !== null &&
        'request' in entry &&
        typeof (entry as { request: unknown }).request === 'string' &&
        (entry as { request: string }).request.includes('thorns'),
    );
  });
  expect(sawThorns).toBe(true);
});
