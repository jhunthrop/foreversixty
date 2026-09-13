// web/tests/e2e/shared-build.spec.ts
import { expect, test } from '@playwright/test';

// The real /b/:id is server-rendered by the Go API and reaches the browser through the
// Worker proxy (Task 12). Playwright serves the static site only, so the page is stubbed
// with exactly the markup the interface contract specifies -- including the island script
// tag, which resolves against the preview server and loads the bundle this task builds.
const RECORD = {
  id: 'k7x2qm4a',
  class_id: 1,
  race_id: 5,
  tree_version: '1.15.9.69722',
  point_order: [1001, 1001, 1001],
  gear: {},
  title: 'Arms leveling',
  created_at: '2026-09-14T03:12:44Z',
  views: 12,
};

const sharedPage = `<!doctype html>
<html lang="en"><head><meta charset="utf-8"><title>Arms leveling · 3/0 · Forever Sixty</title>
<link rel="stylesheet" href="/planner-island.css"></head>
<body><main id="main">
<div id="planner" data-build='${JSON.stringify(RECORD)}' data-tree-version="${RECORD.tree_version}"></div>
<script type="module" src="/planner-island.js"></script>
</main></body></html>`;

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
