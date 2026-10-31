// The band's "your server sim finished" line (experience spec 2026-10-24 section 3.4) on the
// signed-in home page, with /v1/me stubbed.
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { mockEntitlements } from '../../src/test-support/billing-fixtures';
import { readLayoutShift, trackLayoutShift } from './support/cls';
import { ownerMeAnswer } from './support/character-fixture';

const MAX_CLS = 0.02;
const GATE_TIMEOUT_MS = 10_000;

function run(id: string, dps: number): Record<string, unknown> {
  return { sim_id: id, spec: 'mage-frost', title: 'Frost night', dps, finished_at: '2026-10-24T11:00:00Z' };
}

async function stubMe(page: Page, finished: unknown[], server = true): Promise<string[]> {
  const seen: string[] = [];
  const me = ownerMeAnswer();
  const data = me.data as Record<string, unknown>;
  await page
    .context()
    .addCookies([
      { name: 'fs_csrf', value: 'token', url: `http://localhost:${process.env.E2E_PORT ?? '4321'}` },
    ]);
  await page.route('**/v1/me', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({
        ...me,
        data: {
          ...data,
          entitlements: mockEntitlements({ server_sims: server }),
          finished_sims: finished,
        },
      }),
    }),
  );
  await page.route('**/v1/sims/*/seen', (route) => {
    seen.push(route.request().url().split('/').at(-2) ?? '');
    return route.fulfill({ status: 204, body: '' });
  });
  return seen;
}

test('one unseen run: the line sits inside the band, links to the result, and dismiss clears it', async ({
  page,
}) => {
  await page.setViewportSize({ width: 1440, height: 1000 });
  const seen = await stubMe(page, [run('sim-1', 1204)]);
  await page.goto('/');
  const band = page.getByTestId('character-header');
  const line = page.getByTestId('finished-sim-line');
  await expect(line).toBeVisible();
  await expect(line).toContainText('Your Frost Mage sim finished: 1,204 DPS');
  await expect(page.getByTestId('finished-sim-link')).toHaveAttribute('href', '/sim/sim-1');
  await expect(page.getByTestId('finished-sim-more')).toHaveCount(0);

  const withLine = (await band.boundingBox())!.height;
  const lineBox = (await line.boundingBox())!;
  expect(lineBox.height).toBeGreaterThanOrEqual(44);
  expect(lineBox.height).toBeLessThan(46);

  await page.getByTestId('finished-sim-dismiss').click();
  await expect(line).toHaveCount(0);
  expect(seen).toEqual(['sim-1']);
  await expect(page.getByTestId('character-header-action')).toBeFocused();
  const without = (await band.boundingBox())!.height;
  expect(withLine - without).toBeCloseTo(lineBox.height, 0);
  expect(without).toBeGreaterThanOrEqual(136);
});

test('three unseen runs: and 2 more, and dismiss marks all three seen', async ({ page }) => {
  await page.setViewportSize({ width: 1440, height: 1000 });
  const seen = await stubMe(page, [run('sim-3', 1300), run('sim-2', 1250), run('sim-1', 1204)]);
  await page.goto('/');
  await expect(page.getByTestId('finished-sim-link')).toHaveAttribute('href', '/sim/sim-3');
  await expect(page.getByTestId('finished-sim-more')).toHaveText('and 2 more');
  await expect(page.getByTestId('finished-sim-more')).toHaveAttribute('href', '/sim#server-runs');
  await page.getByTestId('finished-sim-dismiss').click();
  await expect(page.getByTestId('finished-sim-line')).toHaveCount(0);
  expect([...seen].sort()).toEqual(['sim-1', 'sim-2', 'sim-3']);
});

test('no unseen runs, or no server sims: no line', async ({ page }) => {
  await stubMe(page, []);
  await page.goto('/');
  await expect(page.getByTestId('character-header')).toBeVisible();
  await expect(page.getByTestId('finished-sim-line')).toHaveCount(0);

  await page.unroute('**/v1/me');
  await stubMe(page, [run('sim-1', 1204)], false);
  await page.reload();
  await expect(page.getByTestId('character-header')).toBeVisible();
  await expect(page.getByTestId('finished-sim-line')).toHaveCount(0);
});

test('a failed dismiss swaps the text for four seconds, then the sentence returns', async ({ page }) => {
  await page.setViewportSize({ width: 1440, height: 1000 });
  await stubMe(page, [run('sim-1', 1204)]);
  await page.unroute('**/v1/sims/*/seen');
  await page.route('**/v1/sims/*/seen', (route) =>
    route.fulfill({
      status: 500,
      contentType: 'application/json',
      body: JSON.stringify({ data: null, error: { code: 'x', message: 'no' } }),
    }),
  );
  await page.goto('/');
  await page.getByTestId('finished-sim-dismiss').click();
  const failure = page.getByTestId('finished-sim-failed');
  await expect(failure).toHaveText('Could not dismiss that. Try again.');
  await page.waitForTimeout(2000);
  await expect(failure).toBeVisible();
  await expect(page.getByTestId('finished-sim-link')).toBeVisible({ timeout: 4000 });
  await expect(failure).toHaveCount(0);
});

test('while the seen call is in flight the dismiss is busy and the line stays, then it goes', async ({
  page,
}) => {
  await page.setViewportSize({ width: 1440, height: 1000 });
  await stubMe(page, [run('sim-1', 1204)]);
  await page.unroute('**/v1/sims/*/seen');
  let release: () => void = () => {};
  const gate = new Promise<void>((resolve) => {
    release = resolve;
    // Fallback so a failing assertion never leaves the handler pending until teardown.
    setTimeout(resolve, GATE_TIMEOUT_MS);
  });
  await page.route('**/v1/sims/*/seen', async (route) => {
    await gate;
    await route.fulfill({ status: 204, body: '' });
  });
  try {
    await page.goto('/');
    const dismiss = page.getByTestId('finished-sim-dismiss');
    await dismiss.click();
    await expect(dismiss).toHaveAttribute('aria-busy', 'true');
    await expect(page.getByTestId('finished-sim-line')).toBeVisible();
    release();
    await expect(page.getByTestId('finished-sim-line')).toHaveCount(0);
  } finally {
    release();
  }
});

test('phone: the arrow replaces Open result, the line is at least 56 px, the dismiss is 44 px', async ({
  page,
}) => {
  await page.setViewportSize({ width: 390, height: 900 });
  await stubMe(page, [run('sim-1', 1204)]);
  await page.goto('/');
  await expect(page.getByTestId('finished-sim-open-phone')).toBeVisible();
  await expect(page.getByTestId('finished-sim-open')).toBeHidden();
  expect((await page.getByTestId('finished-sim-line').boundingBox())!.height).toBeGreaterThanOrEqual(56);
  const dismiss = (await page.getByTestId('finished-sim-dismiss').boundingBox())!;
  expect(dismiss.width).toBe(44);
  expect(dismiss.height).toBe(44);
});

for (const viewport of [
  { width: 390, height: 900 },
  { width: 1440, height: 1000 },
]) {
  test(`the hinted line appears without layout shift at ${viewport.width}px`, async ({ page }) => {
    await page.setViewportSize(viewport);
    await stubMe(page, [run('sim-1', 1204)]);
    // The first visit stores the /v1/me snapshot the pre-paint script reads on the second.
    await page.goto('/');
    await expect(page.getByTestId('finished-sim-line')).toBeVisible();

    await trackLayoutShift(page);
    await page.reload();
    await expect(page.locator('html')).toHaveAttribute('data-finished', '1');
    await expect(page.getByTestId('finished-sim-line')).toBeVisible();
    expect(await readLayoutShift(page)).toBeLessThanOrEqual(MAX_CLS);
  });
}
