import { expect, test } from '@playwright/test';
import { LAST_TALENT, NEARLY_FINISHED_BUILD, finishBuild, showLastTalent } from './support/planner';

// Task 20: a talent or slot change runs a 500-iteration sim in the browser pool and updates
// a DPS estimate with its error in under a second (design 4.6). The fake engine (the default
// here: PUBLIC_SIM_ENGINE is unset, so engine.ts loads src/fixtures/sim/engine-fake.ts) makes
// the figure deterministic-ish -- not the exact number, but that one always arrives.

test('a finished build runs a live sim, and an unfinished one asks nothing of the engine', async ({
  page,
}) => {
  // No wasm asset is ever fetched here (the default engine is the fake, not sim.wasm), but
  // the pool itself must not exist before the first edit either -- this proves both: a
  // request would show up here the moment a worker is spawned, long before it could resolve.
  const simAssetRequests: string[] = [];
  page.on('request', (request) => {
    if (request.url().includes('/_sim/')) simAssetRequests.push(request.url());
  });

  await page.goto(NEARLY_FINISHED_BUILD);
  await showLastTalent(page);
  // One point short: no figure, a line saying why, and nothing asked of the pool.
  await expect(page.getByTestId('planner-dps')).toHaveText('—');
  await expect(page.getByTestId('planner-dps-error')).toHaveText('1 point to go');
  await page.waitForTimeout(600);
  expect(simAssetRequests).toEqual([]);

  await finishBuild(page);
  await expect(page.getByTestId('planner-dps')).not.toHaveText('—', { timeout: 3000 });
  await expect(page.getByTestId('planner-dps-error')).toHaveText(/^± \d/);

  // Taking a point back out stops the sim, and the figure from the last run stays on screen
  // (dimmed) rather than blanking -- "a number that vanishes and reappears on every click is
  // worse than a stale number that dims" is the whole design decision this asserts.
  await page.getByTestId(LAST_TALENT).click({ button: 'right' });
  await expect(page.getByTestId('planner-dps-error')).toHaveText('1 point to go');
  await expect(page.getByTestId('planner-dps')).not.toHaveText('—');

  // Spending it again settles into a fresh estimate of its own.
  await page.getByTestId(LAST_TALENT).click();
  await expect(page.getByTestId('planner-dps-error')).toHaveText(/^± \d/, { timeout: 3000 });

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
  await page.goto(NEARLY_FINISHED_BUILD);
  await showLastTalent(page);
  const cell = page.getByTestId(LAST_TALENT);
  // Measured against the document, not the viewport: a click may scroll the page, and that
  // is not the trees moving.
  const top = (): Promise<number> => cell.evaluate((el) => el.getBoundingClientRect().top + window.scrollY);
  const before = await top();

  // Finishing the build brings the figure (or, on a phone, the button that asks for it).
  await finishBuild(page);
  await expect(page.getByTestId('planner-dps-error')).toHaveText(/^± \d/, { timeout: 3000 });
  expect(await top()).toBe(before);

  // Out and back in: the band gives way to "1 point to go" and returns, and the run that
  // replaces it is pending in between. This is the click that used to jump.
  await page.getByTestId(LAST_TALENT).click({ button: 'right' });
  expect(await top()).toBe(before);
  await page.getByTestId(LAST_TALENT).click();
  expect(await top()).toBe(before);
  await expect(page.getByTestId('planner-dps-error')).toHaveText(/^± \d/, { timeout: 3000 });
  expect(await top()).toBe(before);
});

test('a phone is asked before anything runs, and the answer holds for the visit', async ({
  page,
}, testInfo) => {
  test.skip(testInfo.project.name !== 'mobile', 'touch-first devices only');
  await page.goto(NEARLY_FINISHED_BUILD);
  await showLastTalent(page);
  await page.getByTestId(LAST_TALENT).click();

  await expect(page.getByTestId('planner-dps-show')).toBeVisible();
  await expect(page.getByTestId('planner-dps-error')).toHaveText('Runs on this device');
  await page.getByTestId('planner-dps-show').click();
  await expect(page.getByTestId('planner-dps')).not.toHaveText('—', { timeout: 3000 });

  // Asked once: the next edit sims without asking again.
  await page.getByTestId(LAST_TALENT).click({ button: 'right' });
  await page.getByTestId(LAST_TALENT).click();
  await expect(page.getByTestId('planner-dps-show')).toHaveCount(0);
  await expect(page.getByTestId('planner-dps-error')).toHaveText(/^± \d/, { timeout: 3000 });
});

// The DPS column is one line taller than its neighbours (the ± line). Centring the bar's
// columns put its caption and figure above everyone else's.
test('the DPS caption and figure sit on the same lines as the rest of the summary bar', async ({ page }) => {
  await page.goto(NEARLY_FINISHED_BUILD);
  const top = (testid: string): Promise<number> =>
    page.getByTestId(testid).evaluate((el) => Math.round(el.getBoundingClientRect().top));
  const bottom = (testid: string): Promise<number> =>
    page.getByTestId(testid).evaluate((el) => Math.round(el.getBoundingClientRect().bottom));
  // Same row of the bar on both projects: Points and DPS are neighbours.
  expect(await top('planner-dps')).toBe(await top('planner-spent'));
  expect(await bottom('planner-dps')).toBe(await bottom('planner-spent'));
  expect(await top('planner-sim-link')).toBe(await top('planner-dps'));
});
