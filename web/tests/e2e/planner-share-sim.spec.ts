import { expect, test } from '@playwright/test';

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

test('a saved build sims itself for the card, without blocking the link', async ({ page }) => {
  await page.route('**/v1/builds', (route) =>
    route.fulfill({ status: 201, contentType: 'application/json', body: JSON.stringify(SAVED) }),
  );

  let releaseSim: () => void = () => {};
  const heldSim = new Promise<void>((resolve) => {
    releaseSim = resolve;
  });
  const simRequests: unknown[] = [];
  await page.route('**/v1/sims', async (route) => {
    simRequests.push(route.request().postDataJSON());
    await heldSim;
    await route.fulfill({
      status: 201,
      contentType: 'application/json',
      body: JSON.stringify({ ok: true, data: { sim_id: 'simnew234567' }, error: null, request_id: 'req-2' }),
    });
  });

  await page.goto('/planner');
  await page.getByTestId('talent-1001').click();
  await expect(page.getByTestId('planner-dps')).not.toHaveText('—', { timeout: 3000 });

  const checkbox = page.getByLabel('Include a simmed DPS on the card');
  await expect(checkbox).toBeChecked();

  await page.getByRole('button', { name: 'Share' }).click();

  // The link is already on screen before the card sim has anywhere near finished.
  await expect(page.getByTestId('share-link')).toHaveText('https://foreversixty.gg/b/k7x2qm4a');
  await expect(page.getByTestId('build-sim-status')).toHaveText('Simming this build for the card…');

  releaseSim();

  await expect(page.getByTestId('build-sim-status')).toHaveText(
    /^The card will show [\d,]+ DPS on engine \S+\.$/,
  );

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

  await page.goto('/planner');
  await page.getByTestId('talent-1001').click();
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

  await page.goto('/planner');
  await page.getByTestId('talent-1001').click();
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
