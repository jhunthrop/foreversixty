import { expect, test } from '@playwright/test';
import { LAST_TALENT, NEARLY_FINISHED_BUILD, finishBuild, showTree } from './support/planner';

import { heldRoute } from './support/held-route';

// Task 21: sharing a build also saves a 3,000-iteration sim whose source is that build's
// own id, so the API's card renderer can join it in without a new field on either shape
// (design 4.7). The save itself never waits on the sim: the link is usable the moment the
// build is saved, and a failed sim only costs the card its DPS line, never the share.

const SAVED = {
  ok: true,
  data: { id: 'k7x2qm4a', url: 'https://foreversixty.gg/b/k7x2qm4a' },
  error: null,
  request_id: 'req-1',
};

const DONE_TEXT = /^The card will show [\d,]+ DPS on engine \S+\.$/;

test('a saved build sims itself for the card, without blocking the link', async ({ page }) => {
  await page.route('**/v1/builds', (route) =>
    route.fulfill({ status: 201, contentType: 'application/json', body: JSON.stringify(SAVED) }),
  );

  const simRequests: unknown[] = [];
  const sim = await heldRoute(page, '**/v1/sims', async (route) => {
    simRequests.push(route.request().postDataJSON());
    await route.fulfill({
      status: 201,
      contentType: 'application/json',
      body: JSON.stringify({ ok: true, data: { sim_id: 'simnew234567' }, error: null, request_id: 'req-2' }),
    });
  });

  await page.goto(NEARLY_FINISHED_BUILD);
  await finishBuild(page);
  await expect(page.getByTestId('planner-dps')).not.toHaveText('—', { timeout: 3000 });

  const checkbox = page.getByLabel('Include a simmed DPS on the card');
  await expect(checkbox).toBeChecked();

  await page.getByRole('button', { name: 'Share' }).click();

  // The link is already on screen before the card sim has anywhere near finished.
  await expect(page.getByTestId('share-link')).toHaveText('https://foreversixty.gg/b/k7x2qm4a');
  await sim.started;
  await expect(page.getByTestId('build-sim-status')).toHaveText(
    'Simming this build for the card, a few seconds…',
  );

  sim.release();

  await expect(page.getByTestId('build-sim-status')).toHaveText(DONE_TEXT);

  expect(simRequests).toHaveLength(1);
  const body = simRequests[0] as { request: { source: { kind: string; ref: string }; iterations: number } };
  expect(body.request.source.kind).toBe('build');
  expect(body.request.source.ref).toBe(SAVED.data.id);
  expect(body.request.iterations).toBe(3000);
});

test('a failed card sim never breaks the share link', async ({ page }) => {
  await page.route('**/v1/builds', (route) =>
    route.fulfill({ status: 201, contentType: 'application/json', body: JSON.stringify(SAVED) }),
  );
  await page.route('**/v1/sims', (route) =>
    route.fulfill({
      status: 500,
      contentType: 'application/json',
      body: JSON.stringify({ ok: false, data: null, error: { message: 'boom' }, request_id: 'req-2' }),
    }),
  );

  await page.goto(NEARLY_FINISHED_BUILD);
  await finishBuild(page);
  await expect(page.getByTestId('planner-dps')).not.toHaveText('—', { timeout: 3000 });

  await page.getByRole('button', { name: 'Share' }).click();

  await expect(page.getByTestId('share-link')).toHaveText('https://foreversixty.gg/b/k7x2qm4a');
  await expect(page.getByTestId('build-sim-status')).toHaveText(
    'The card will not show a DPS figure; the sim did not finish.',
  );
});

test('unchecking the box skips the card sim entirely', async ({ page }) => {
  await page.route('**/v1/builds', (route) =>
    route.fulfill({ status: 201, contentType: 'application/json', body: JSON.stringify(SAVED) }),
  );
  let simCalled = false;
  await page.route('**/v1/sims', async (route) => {
    simCalled = true;
    await route.fulfill({
      status: 201,
      contentType: 'application/json',
      body: JSON.stringify({ ok: true, data: { sim_id: 'simnew234567' }, error: null, request_id: 'req-2' }),
    });
  });

  await page.goto(NEARLY_FINISHED_BUILD);
  await finishBuild(page);
  await expect(page.getByTestId('planner-dps')).not.toHaveText('—', { timeout: 3000 });

  const checkbox = page.getByLabel('Include a simmed DPS on the card');
  await expect(checkbox).toBeChecked();
  await checkbox.uncheck();

  await page.getByRole('button', { name: 'Share' }).click();

  await expect(page.getByTestId('share-link')).toHaveText('https://foreversixty.gg/b/k7x2qm4a');
  // No running/done/skipped line at all -- attachSim returned before touching the pool.
  await expect(page.getByTestId('build-sim-status')).not.toBeVisible();
  expect(simCalled).toBe(false);
});

test("a second save while the first build's sim is still running never overwrites the card", async ({
  page,
}) => {
  // Two different builds, one per Share click, distinguished by id/url. The talent click
  // between them is what makes the second save an edit to a *different* build, not a retry
  // of the same one -- the scenario H1 in the review names: "a talent edit + re-save while
  // the card sim is still running."
  const builds = [
    {
      ok: true,
      data: { id: 'k7x2qm4a', url: 'https://foreversixty.gg/b/k7x2qm4a' },
      error: null,
      request_id: 'r-a',
    },
    {
      ok: true,
      data: { id: 'k9m3wt7c', url: 'https://foreversixty.gg/b/k9m3wt7c' },
      error: null,
      request_id: 'r-b',
    },
  ];
  let buildCall = 0;
  await page.route('**/v1/builds', (route) => {
    const body = builds[Math.min(buildCall, builds.length - 1)];
    buildCall += 1;
    return route.fulfill({ status: 201, contentType: 'application/json', body: JSON.stringify(body) });
  });

  // Only the first build's sim is held -- at the network layer, after its own 3,000-iteration
  // run has already finished, which is enough to prove the race: `cardState` stays 'running'
  // for that first attachSim call the whole time it is held here. `heldRoute` holds every
  // request on a pattern behind one shared gate, which would hold the second build's sim
  // too and defeat the point of this test, so the two responses need distinguishing by the
  // build id each request's own body carries rather than by url alone.
  const firstBuildId = builds[0].data.id;
  const simRefs: string[] = [];
  let releaseFirstSim: () => void = () => {};
  const holdFirstSim = new Promise<void>((resolve) => {
    releaseFirstSim = resolve;
  });
  await page.route('**/v1/sims', async (route) => {
    const posted = route.request().postDataJSON() as { request: { source: { ref: string } } };
    simRefs.push(posted.request.source.ref);
    if (posted.request.source.ref === firstBuildId) await holdFirstSim;
    await route.fulfill({
      status: 201,
      contentType: 'application/json',
      body: JSON.stringify({
        ok: true,
        data: { sim_id: `sim${posted.request.source.ref}` },
        error: null,
        request_id: 'r-sim',
      }),
    });
  });

  await page.goto(NEARLY_FINISHED_BUILD);
  await finishBuild(page);
  await expect(page.getByTestId('planner-dps')).not.toHaveText('—', { timeout: 3000 });

  await page.getByRole('button', { name: 'Share' }).click();
  await expect(page.getByTestId('share-link')).toHaveText(`https://foreversixty.gg/b/${firstBuildId}`);
  await expect(page.getByTestId('build-sim-status')).toHaveText(
    'Simming this build for the card, a few seconds…',
  );

  // A second, different save starts while the first build's sim is still held above.
  // Still a whole build, so it still sims: one point moves from Arms into Flurry.
  await showTree(page, 'Arms');
  await page.getByTestId('talent-1001').click({ button: 'right' });
  await showTree(page, 'Fury');
  await page.getByTestId(LAST_TALENT).click();
  await expect(page.getByTestId('share-link')).not.toBeVisible();
  await page.getByRole('button', { name: 'Share' }).click();

  const secondBuildId = builds[1].data.id;
  await expect(page.getByTestId('share-link')).toHaveText(`https://foreversixty.gg/b/${secondBuildId}`);
  await expect(page.getByTestId('build-sim-status')).toHaveText(DONE_TEXT);

  // The first build's sim finally answers -- its result must not clobber the second build's.
  releaseFirstSim();
  await page.waitForTimeout(300);
  await expect(page.getByTestId('share-link')).toHaveText(`https://foreversixty.gg/b/${secondBuildId}`);
  await expect(page.getByTestId('build-sim-status')).toHaveText(DONE_TEXT);

  expect(simRefs).toEqual([firstBuildId, secondBuildId]);
});
