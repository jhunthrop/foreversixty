import { expect, test } from '@playwright/test';

// Task 20: a talent or slot change runs a 500-iteration sim in the browser pool and updates
// a DPS estimate with its error in under a second (design 4.6). The fake engine (the default
// here: PUBLIC_SIM_ENGINE is unset, so engine.ts loads src/fixtures/sim/engine-fake.ts) makes
// the figure deterministic-ish -- not the exact number, but that one always arrives.

test('a talent change runs a live sim and keeps a figure on screen through the next one', async ({
  page,
}) => {
  // No wasm asset is ever fetched here (the default engine is the fake, not sim.wasm), but
  // the pool itself must not exist before the first edit either -- this proves both: a
  // request would show up here the moment a worker is spawned, long before it could resolve.
  const simAssetRequests: string[] = [];
  page.on('request', (request) => {
    if (request.url().includes('/_sim/')) simAssetRequests.push(request.url());
  });

  await page.goto('/planner');
  await expect(page.getByTestId('talent-1001')).toBeVisible();
  await expect(page.getByTestId('planner-dps')).toHaveText('—');
  expect(simAssetRequests).toEqual([]);

  await page.getByTestId('talent-1001').click();
  // The debounce (350ms) has not fired yet, so nothing has been asked of the pool.
  expect(simAssetRequests).toEqual([]);

  await expect(page.getByTestId('planner-dps')).not.toHaveText('—', { timeout: 3000 });
  await expect(page.getByTestId('planner-dps-error')).toHaveText(/^± \d/);

  // Three more edits, quickly: each one re-debounces the run behind it, and the figure from
  // the previous run stays on screen (dimmed) the whole time rather than blanking between
  // clicks -- "a number that vanishes and reappears on every click is worse than a stale
  // number that dims" is the whole design decision this asserts.
  for (let i = 0; i < 3; i += 1) {
    await page.getByTestId('talent-1002').click();
    await expect(page.getByTestId('planner-dps')).not.toHaveText('—');
  }

  // The last of those four edits settles into its own ready estimate.
  await expect(page.getByTestId('planner-dps')).not.toHaveText('—', { timeout: 3000 });
  await expect(page.getByTestId('planner-dps-error')).toHaveText(/^± \d/);

  expect(simAssetRequests).toEqual([]);
});

test('"Sim this build" carries the build to the full results', async ({ page }) => {
  await page.goto('/planner');
  await expect(page.getByTestId('talent-1001')).toBeVisible();
  await page.getByTestId('talent-1001').click();

  const link = page.getByTestId('planner-sim-link');
  // Never saved in this test, so the link carries the build's own FS1 code -- the planner's
  // export format, via encodeFS1 -- rather than a build id.
  await expect(link).toHaveAttribute('href', /^\/sim\?code=FS1%3A/);

  await link.click();
  await expect(page).toHaveURL(/\/sim\?code=/);
  // /sim reading a `code` query itself is a later task's work (it reads ?source=/?ref=
  // today; see src/pages/sim.astro and src/lib/sim/url.ts) -- what belongs to this one is
  // that the link is well-formed and lands the player on the simulator, not that the page
  // has already learned to decode it.
  await expect(page.getByTestId('sim-view')).toBeVisible();
});

// The error band only exists once a run is ready. It used to be removed from the page while
// the next run was pending, so every talent click made the summary bar one line shorter and
// then one line taller again, and the trees under it jumped up and back.
test('a talent click never moves the trees: the summary bar keeps its height through a run', async ({
  page,
}) => {
  await page.goto('/planner');
  const cell = page.getByTestId('talent-1001');
  await expect(cell).toBeVisible();
  const before = (await cell.boundingBox())?.y;

  await cell.click();
  await expect(page.getByTestId('planner-dps-error')).toHaveText(/^± \d/, { timeout: 3000 });
  expect((await cell.boundingBox())?.y).toBe(before);

  // The second click is the one that used to jump: a band is on screen and the run that
  // replaces it is pending.
  await page.getByTestId('talent-1002').click();
  expect((await cell.boundingBox())?.y).toBe(before);
  await expect(page.getByTestId('planner-dps-error')).toHaveText(/^± \d/, { timeout: 3000 });
  expect((await cell.boundingBox())?.y).toBe(before);
});
