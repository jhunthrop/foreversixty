// web/tests/e2e/planner-dps.spec.ts
// The planner owns the talents and the simulator owns the numbers (one job per page; directions
// mock b2-planner): the head answers with the points spent, and "Sim this build ›" carries the
// build to the simulator. The standalone planner never runs the engine; the sim tools that embed
// the planner keep their live estimate (SummaryBar, sim-talents.spec.ts).
import { expect, test } from './support/test';
import {
  LAST_TALENT,
  NEARLY_FINISHED_BUILD,
  finishBuild,
  pointsSpent,
  showLastTalent,
  spentIs,
} from './support/planner';

test('the head answers with the points, never a sim figure, and asks nothing of the engine', async ({
  page,
}) => {
  // A request would show up here the moment a sim worker is spawned.
  const simAssetRequests: string[] = [];
  page.on('request', (request) => {
    if (request.url().includes('/_sim/')) simAssetRequests.push(request.url());
  });

  await page.goto(NEARLY_FINISHED_BUILD);
  await showLastTalent(page);
  await expect(pointsSpent(page)).toHaveText(spentIs(50));
  await expect(page.getByTestId('planner-answer')).toContainText('of 51 points');

  await finishBuild(page);
  await expect(page.getByTestId('planner-answer')).not.toContainText('DPS');
  await expect(page.getByTestId('planner-dps')).toHaveCount(0);
  await page.waitForTimeout(600);
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
  await expect(page.getByTestId('sim-view')).toBeVisible();
});

test('"Sim this build" sits in the head answer, under the points and above the Share panel', async ({
  page,
}) => {
  await page.goto(NEARLY_FINISHED_BUILD);
  const simLink = page.getByTestId('planner-sim-link');
  await expect(simLink).toBeVisible();
  await expect(page.getByTestId('share-panel')).toBeVisible();
  await expect(page.getByTestId('planner-answer').getByTestId('planner-sim-link')).toBeVisible();
  const top = (testid: string): Promise<number> =>
    page.getByTestId(testid).evaluate((el) => el.getBoundingClientRect().top);
  const pointsBottom = await page
    .getByTestId('planner-points')
    .evaluate((el) => el.getBoundingClientRect().bottom);
  const simTop = await top('planner-sim-link');
  expect(simTop).toBeGreaterThanOrEqual(pointsBottom);
  expect(simTop).toBeLessThan(await top('share-panel'));
});

test('a talent click never moves the trees: the head keeps its height', async ({ page }) => {
  await page.goto(NEARLY_FINISHED_BUILD);
  await showLastTalent(page);
  const cell = page.getByTestId(LAST_TALENT);
  // Measured against the document: a click may scroll the page, and that is not the trees moving.
  const top = (): Promise<number> => cell.evaluate((el) => el.getBoundingClientRect().top + window.scrollY);
  // The title (Marcellus) settles a few pixels once the font swaps in, once.
  await page.evaluate(() => document.fonts.ready);
  await page.waitForTimeout(500);
  const before = await top();

  await finishBuild(page);
  expect(await top()).toBe(before);
  await page.getByTestId(LAST_TALENT).click({ button: 'right' });
  expect(await top()).toBe(before);
  await page.getByTestId(LAST_TALENT).click();
  expect(await top()).toBe(before);
});
