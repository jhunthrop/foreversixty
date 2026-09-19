// web/tests/e2e/sim-saved.spec.ts
// Task 17: /sim/<sim_id> (read-only, from the prerendered fixture and from a stubbed shell
// for an id nothing prerendered) and the save flow on /sim itself.
//
// Read as JSON rather than imported as an ES module: Playwright's own Node runtime needs an
// import attribute this repo's other e2e specs do not carry for a plain `.json` import, the
// same reason sim-sources.spec.ts and sim-specs.spec.ts read their own fixtures this way.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test } from '@playwright/test';

const ROOT = path.join(import.meta.dirname, '..', '..');
const activeBuild = JSON.parse(readFileSync(path.join(ROOT, 'src', 'data', 'active-build.json'), 'utf8')) as {
  build: string;
};
const fixtureResult = JSON.parse(
  readFileSync(path.join(ROOT, 'src', 'fixtures', 'sim', 'result.json'), 'utf8'),
) as {
  dps: { mean: number };
  engine_version: string;
  iterations_run: number;
  duration_ms: number;
  request: { character: { gear: unknown[] } };
};

const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=11726`;

function envelope(data: unknown, status = 200) {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: true, data, error: null, request_id: 'req-test' }),
  };
}

function failure(message: string, status: number) {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: false, data: null, error: { message }, request_id: 'req-test' }),
  };
}

// The minimal shell `sim.astro` (and the Worker's live re-serve of the same static asset
// for any other id) emits: a mount div carrying no `data-sim-id`, so the island falls back
// to reading the id off `window.location.pathname` -- `sim-island.ts`'s own `simIdFor`.
// `astro preview` has no Worker in front of it (that rewrite is Cloudflare-only), so a
// non-prerendered `/sim/<id>` 404s unless the navigation itself is stubbed, the same reason
// shared-build.spec.ts stubs `/b/k7x2qm4a` rather than navigating to it directly.
const SHELL_HTML = `<!doctype html>
<html lang="en"><head><meta charset="utf-8"><title>Saved sim</title>
<link rel="stylesheet" href="/sim-island.css"></head>
<body><main id="main">
<div id="sim" data-sim-mount></div>
<script type="module" src="/sim-island.js"></script>
</main></body></html>`;

test.describe('a saved sim from the prerendered fixture', () => {
  test('renders read-only, with the results and no gear grid claim it cannot make', async ({ page }) => {
    await page.goto('/sim/simfixtureab');

    // Read from the fixture rather than a literal (fixture values drift as the engine
    // golden is refreshed; the plan's own "1,131" is stale against the checked-in file).
    const figure = Math.round(fixtureResult.dps.mean).toLocaleString('en-US');
    await expect(page.getByTestId('sim-dps')).toHaveText(figure);
    await expect(page.getByTestId('sim-saved-encounter')).toHaveText('raid-buffed, 3:00, single target');

    await expect(page.getByTestId('sim-results')).toBeVisible();
    await expect(page.getByTestId('sim-tab-damage')).toHaveAttribute('aria-selected', 'true');

    // The fixture's stored request carries real gear entries, so `gearKnown` is true and
    // the strip renders the grid rather than `simCopy.savedNoGear` -- the inverse of a
    // request with no gear at all, exercised below against a stubbed id.
    await expect(page.getByTestId('sim-slot-head')).toBeVisible();
    await expect(page.getByTestId('sim-no-gear')).toHaveCount(0);
  });

  test('Run this yourself opens /sim with the stored request’s source', async ({ page }) => {
    await page.goto('/sim/simfixtureab');
    await page.getByTestId('sim-run-yourself').click();
    await expect(page).toHaveURL('/sim?source=addon');
  });
});

test('a saved sim whose stored request has no gear shows the line, not the grid', async ({ page }) => {
  const id = 'simnogearabc';
  const noGearResult = {
    ...fixtureResult,
    sim_id: id,
    request: { ...fixtureResult.request, character: { ...fixtureResult.request.character, gear: [] } },
  };
  await page.route(`**/sim/${id}`, (route) =>
    route.fulfill({ status: 200, contentType: 'text/html', body: SHELL_HTML }),
  );
  await page.route(`**/v1/sims/${id}`, (route) => route.fulfill(envelope(noGearResult)));

  await page.goto(`/sim/${id}`);

  await expect(page.getByTestId('sim-no-gear')).toBeVisible();
  await expect(page.getByTestId('sim-slot-head')).toHaveCount(0);
});

test('a stale engine version shows the pill and the sentence, and is never re-run automatically', async ({
  page,
}) => {
  const id = 'simstaleabcd';
  const staleResult = { ...fixtureResult, sim_id: id, engine_version: '6a1c2d9' };
  let calls = 0;
  await page.route(`**/sim/${id}`, (route) =>
    route.fulfill({ status: 200, contentType: 'text/html', body: SHELL_HTML }),
  );
  await page.route(`**/v1/sims/${id}`, (route) => {
    calls += 1;
    return route.fulfill(envelope(staleResult));
  });

  await page.goto(`/sim/${id}`);

  await expect(page.getByTestId('sim-stale-pill')).toHaveText('Engine 6a1c2d9');
  await expect(page.getByTestId('sim-stale')).toBeVisible();

  // No auto re-run: the remedy is the "Run this yourself" link, never a second fetch fired
  // on its own.
  await page.waitForTimeout(500);
  expect(calls).toBe(1);
});

test.describe('saving a browser run from /sim', () => {
  async function runFury(page: import('@playwright/test').Page): Promise<void> {
    await page.goto('/sim');
    await page.getByTestId('sim-addon-input').fill(FURY);
    await page.getByTestId('sim-addon-load').click();
    await expect(page.getByTestId('sim-character')).toBeVisible();
    await page.getByTestId('sim-run-button').click();
    await expect(page.getByTestId('sim-run-button')).toHaveText('Run again', { timeout: 5000 });
  }

  test('opens pre-filled, saves, and shows the link without navigating', async ({ page }) => {
    await runFury(page);

    await expect(page.getByTestId('sim-save-open')).toBeEnabled();
    await page.getByTestId('sim-save-open').click();
    await expect(page.getByTestId('sim-save-title')).toHaveValue('Raid-buffed, 3:00, single target');

    await page.route('**/v1/sims', (route) => route.fulfill(envelope({ sim_id: 'simnew234567' }, 201)));
    await page.getByTestId('sim-save-confirm').click();

    await expect(page.getByTestId('sim-save-link')).toBeVisible();
    await expect(page.getByTestId('sim-save-link')).toHaveValue(/\/sim\/simnew234567$/);
    await expect(page).toHaveURL(/\/sim(\?.*)?$/);
    await expect(page.getByTestId('sim-view')).toBeVisible();
  });

  test('a failed save shows saveFailed and keeps the typed title', async ({ page }) => {
    await runFury(page);

    await page.getByTestId('sim-save-open').click();
    await page.getByTestId('sim-save-title').fill('My opener sim');

    await page.route('**/v1/sims', (route) => route.fulfill(failure('boom', 500)));
    await page.getByTestId('sim-save-confirm').click();

    await expect(page.getByTestId('sim-save-error')).toBeVisible();
    await expect(page.getByTestId('sim-save-title')).toHaveValue('My opener sim');
  });
});
