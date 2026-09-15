// web/tests/e2e/shared-build.spec.ts
import { expect, test } from '@playwright/test';

import { ACTIVE_BUILD } from './support/active-build';

// The real /b/:id is server-rendered by the Go API and reaches the browser through the
// Worker proxy (Task 12). Playwright serves the static site only, so the page is stubbed
// with exactly the markup the interface contract specifies -- including the island script
// tag, which resolves against the preview server and loads the bundle this task builds.
const RECORD = {
  id: 'k7x2qm4a',
  class_id: 1,
  race_id: 5,
  tree_version: ACTIVE_BUILD,
  point_order: [1001, 1001, 1001],
  gear: {},
  title: 'Arms leveling',
  created_at: '2026-09-14T03:12:44Z',
  views: 12,
};

const pageFor = (build: string): string => `<!doctype html>
<html lang="en"><head><meta charset="utf-8"><title>Arms leveling · 3/0 · Forever Sixty</title>
<link rel="stylesheet" href="/planner-island.css"></head>
<body><main id="main">
<div id="planner" data-build='${build}' data-tree-version="${RECORD.tree_version}"></div>
<script type="module" src="/planner-island.js"></script>
</main></body></html>`;

const sharedPage = pageFor(JSON.stringify(RECORD));

test('the island bundle and its stylesheet are published at fixed paths', async ({ request }) => {
  const script = await request.get('/planner-island.js');
  expect(script.status()).toBe(200);
  expect(script.headers()['content-type']).toContain('javascript');
  expect((await script.body()).length).toBeGreaterThan(1000);

  const styles = await request.get('/planner-island.css');
  expect(styles.status()).toBe(200);
  expect(styles.headers()['content-type']).toContain('css');
});

test('a shared build opens read-only and Fork makes it editable', async ({ page }) => {
  await page.route('**/b/k7x2qm4a', (route) =>
    route.fulfill({ status: 200, contentType: 'text/html', body: sharedPage }),
  );
  await page.goto('/b/k7x2qm4a');

  await expect(page.getByTestId('talent-1001')).toHaveAttribute('data-rank', '3');
  await expect(page.getByTestId('planner-split')).toHaveText('3/0');
  await expect(page.getByTestId('planner-level')).toHaveText('12');
  await expect(page.getByLabel('Class')).toBeDisabled();

  await page.getByTestId('talent-1002').click();
  await expect(page.getByTestId('planner-refusal')).toHaveText(
    'This build was opened from a share link; fork it to edit',
  );
  await expect(page.getByTestId('talent-1002')).toHaveAttribute('data-rank', '0');

  await page.getByRole('button', { name: 'Fork' }).click();
  await expect(page.getByLabel('Class')).toBeEnabled();
  await page.getByTestId('talent-1002').click();
  await expect(page.getByTestId('talent-1002')).toHaveAttribute('data-rank', '1');
  await expect(page.getByTestId('planner-spent')).toHaveText('4/51');
  await expect(page.getByRole('button', { name: 'Reset' })).toBeVisible();
});

// The API writes data-build itself, so malformed JSON there means the API is broken, not the
// visitor. The island logs it and mounts anyway rather than leaving a blank page: an empty
// planner someone can use beats nothing at all, and it cannot be read-only -- there is no
// build to be read-only about.
test('a build attribute that is not JSON still mounts an editable planner', async ({ page }) => {
  await page.route('**/b/broken', (route) =>
    route.fulfill({ status: 200, contentType: 'text/html', body: pageFor('{"class_id":') }),
  );
  await page.goto('/b/broken');

  await expect(page.getByRole('grid', { name: 'Arms talents' })).toBeVisible();
  await expect(page.getByTestId('talent-1001')).toHaveAttribute('data-rank', '0');
  await expect(page.getByLabel('Class')).toBeEnabled();
  await expect(page.getByRole('button', { name: 'Reset' })).toBeVisible();
  await expect(page.getByRole('button', { name: 'Fork' })).toBeHidden();
});

// boot() resolves the record's class and race ids through loadReference before it mounts.
// When that fetch fails it swallows the error deliberately, because Planner runs the same
// call again and has a state for it -- so the visitor gets the retry rather than a blank div.
test('a reference fetch that fails still mounts, and Planner offers the retry', async ({ page }) => {
  await page.route('**/data/*/classes.json', (route) => route.abort('failed'));
  await page.route('**/b/k7x2qm4a', (route) =>
    route.fulfill({ status: 200, contentType: 'text/html', body: sharedPage }),
  );
  await page.goto('/b/k7x2qm4a');

  await expect(page.getByText('Talent data did not load')).toBeVisible();
  await expect(page.getByRole('button', { name: 'Retry' })).toBeVisible();
});
