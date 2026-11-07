import { expect, test } from './support/test';
import {
  LAST_TALENT,
  NEARLY_FINISHED_BUILD,
  finishBuild,
  scoreFigure,
  showLastTalent,
} from './support/planner';

// Task 20: a talent or slot change runs a 500-iteration sim in the browser pool and updates
// a DPS estimate with its error in under a second (design 4.6). The fake engine (the default
// here: PUBLIC_SIM_ENGINE is unset, so engine.ts loads src/fixtures/sim/engine-fake.ts) makes
// the figure deterministic-ish -- not the exact number, but that one always arrives. Spec
// 2026-11-08 4.1: the figure is the head's answer (`planner-dps`), its caption under it
// (`planner-dps-note`), run on the Level build section's gear once that section has answered.

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
  // One point short: no figure, and the DPS slot reads a plain explainer (never a point count).
  await expect(scoreFigure(page)).toHaveText('—');
  await expect(page.getByTestId('planner-dps-note')).toHaveText('Appears once your points are spent.');
  await page.waitForTimeout(600);
  expect(simAssetRequests).toEqual([]);

  await finishBuild(page);
  await expect(scoreFigure(page)).not.toHaveText('—', { timeout: 3000 });
  await expect(page.getByTestId('planner-dps-note')).toHaveText(/^± \d/);

  // Taking a point back out stops the sim. Spec 2026-11-08 4.1 retired the dimmed stale figure
  // for an unfinished build: the head reads `—` with its waiting line until every point is spent.
  await page.getByTestId(LAST_TALENT).click({ button: 'right' });
  await expect(page.getByTestId('planner-dps-note')).toHaveText('Appears once your points are spent.');
  await expect(scoreFigure(page)).toHaveText('—');

  // Spending it again settles into a fresh estimate of its own.
  await page.getByTestId(LAST_TALENT).click();
  await expect(page.getByTestId('planner-dps-note')).toHaveText(/^± \d/, { timeout: 3000 });

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

// The error line only exists once a run is ready. It used to be removed from the page while
// the next run was pending, so every talent click made the head one line shorter and then one
// line taller again, and the trees under it jumped up and back.
test('a talent click never moves the trees: the head keeps its height through a run', async ({ page }) => {
  await page.goto(NEARLY_FINISHED_BUILD);
  await showLastTalent(page);
  const cell = page.getByTestId(LAST_TALENT);
  // Measured against the document, not the viewport: a click may scroll the page, and that
  // is not the trees moving.
  const top = (): Promise<number> => cell.evaluate((el) => el.getBoundingClientRect().top + window.scrollY);
  // The head's title (Marcellus) settles a few pixels once the font swaps in -- unrelated to
  // anything this test is about (a talent click moving nothing). The settle happens once,
  // well under a second, and never again; waiting it out here is what isolates the one
  // shift this test actually exists to catch.
  await page.evaluate(() => document.fonts.ready);
  await page.waitForTimeout(500);
  const before = await top();

  // Finishing the build brings the figure (or, on a phone, the button that asks for it).
  await finishBuild(page);
  await expect(page.getByTestId('planner-dps-note')).toHaveText(/^± \d/, { timeout: 3000 });
  expect(await top()).toBe(before);

  // Out and back in: the band gives way to "1 point to go" and returns, and the run that
  // replaces it is pending in between. This is the click that used to jump.
  await page.getByTestId(LAST_TALENT).click({ button: 'right' });
  expect(await top()).toBe(before);
  await page.getByTestId(LAST_TALENT).click();
  expect(await top()).toBe(before);
  await expect(page.getByTestId('planner-dps-note')).toHaveText(/^± \d/, { timeout: 3000 });
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
  await expect(page.getByTestId('planner-dps-note')).toHaveText('Runs on this device');
  await page.getByTestId('planner-dps-show').click();
  await expect(scoreFigure(page)).not.toHaveText('—', { timeout: 3000 });

  // Asked once: the next edit sims without asking again.
  await page.getByTestId(LAST_TALENT).click({ button: 'right' });
  await page.getByTestId(LAST_TALENT).click();
  await expect(page.getByTestId('planner-dps-show')).toHaveCount(0);
  await expect(page.getByTestId('planner-dps-note')).toHaveText(/^± \d/, { timeout: 3000 });
});

// Spec 2026-11-08 4.1: "Sim this build" is the head answer's link, under the two figures
// (the build's code or the character's), above everything in the rail -- Share included.
test('"Sim this build" sits in the head answer, under the figures and above the Share panel', async ({
  page,
}) => {
  await page.goto(NEARLY_FINISHED_BUILD);
  const simLink = page.getByTestId('planner-sim-link');
  const shareSection = page.getByTestId('share-panel');
  await expect(simLink).toBeVisible();
  await expect(shareSection).toBeVisible();
  await expect(page.getByTestId('planner-answer').getByTestId('planner-sim-link')).toBeVisible();
  const top = (testid: string): Promise<number> =>
    page.getByTestId(testid).evaluate((el) => el.getBoundingClientRect().top);
  const figureBottom = await page
    .getByTestId('planner-dps')
    .evaluate((el) => el.getBoundingClientRect().bottom);
  const simTop = await top('planner-sim-link');
  expect(simTop).toBeGreaterThanOrEqual(figureBottom);
  expect(simTop).toBeLessThan(await top('share-panel'));
});

// The head's answer is two fixed figure blocks (spec 2026-11-08 5.2): it must not change height
// as the figure fills in -- the same CLS concern the old row-sharing test guarded.
test('the head answer keeps its own height as the DPS figure fills in', async ({ page }) => {
  await page.goto(NEARLY_FINISHED_BUILD);
  await showLastTalent(page);
  const answer = page.getByTestId('planner-answer');
  const height = (): Promise<number> => answer.evaluate((el) => el.getBoundingClientRect().height);
  const before = await height();
  await finishBuild(page);
  await expect(page.getByTestId('planner-dps-note')).toHaveText(/^± \d/, { timeout: 3000 });
  expect(await height()).toBe(before);
});
