import { expect, test } from '@playwright/test';

const SAVED = {
  ok: true,
  data: { id: 'k7x2qm4a', url: 'https://foreversixty.gg/b/k7x2qm4a' },
  error: null,
  request_id: 'req-1',
};

test('sharing a build posts the contract body and shows the link', async ({ page }) => {
  let body: unknown;
  await page.route('**/v1/builds', async (route) => {
    body = route.request().postDataJSON();
    await route.fulfill({ status: 201, contentType: 'application/json', body: JSON.stringify(SAVED) });
  });

  await page.goto('/planner');
  for (let i = 0; i < 3; i += 1) await page.getByTestId('talent-1001').click();
  await page.getByLabel('Title').fill('Arms leveling');
  await page.getByRole('button', { name: 'Share' }).click();

  await expect(page.getByTestId('share-link')).toHaveText('https://foreversixty.gg/b/k7x2qm4a');
  expect(body).toEqual({
    class_id: 1,
    race_id: 1,
    tree_version: '1.15.9.69722',
    point_order: [1001, 1001, 1001],
    gear: {},
    title: 'Arms leveling',
  });
});

test('the copy button reports that it copied', async ({ page, context, browserName }) => {
  test.skip(browserName !== 'chromium', 'clipboard permissions are Chromium-only here');
  await context.grantPermissions(['clipboard-read', 'clipboard-write']);
  await page.route('**/v1/builds', (route) =>
    route.fulfill({ status: 201, contentType: 'application/json', body: JSON.stringify(SAVED) }),
  );
  await page.goto('/planner');
  await page.getByTestId('talent-1001').click();
  await page.getByRole('button', { name: 'Share' }).click();
  await page.getByRole('button', { name: 'Copy link' }).click();
  await expect(page.getByRole('button', { name: 'Copied' })).toBeVisible();
});

test('a rate-limited save keeps the build and explains the wait', async ({ page }) => {
  await page.route('**/v1/builds', (route) =>
    route.fulfill({
      status: 429,
      contentType: 'application/json',
      body: JSON.stringify({ ok: false, data: null, error: { message: 'slow down' }, request_id: 'r' }),
    }),
  );
  await page.goto('/planner');
  await page.getByTestId('talent-1001').click();
  await page.getByRole('button', { name: 'Share' }).click();
  await expect(page.getByRole('alert')).toContainText(
    'Too many saves from this connection; try again in an hour.',
  );
  await expect(page.getByTestId('planner-spent')).toHaveText('1/51');
});

test('a rejected build shows the API field message and retries', async ({ page }) => {
  let attempts = 0;
  await page.route('**/v1/builds', async (route) => {
    attempts += 1;
    if (attempts === 1) {
      return route.fulfill({
        status: 400,
        contentType: 'application/json',
        body: JSON.stringify({
          ok: false,
          data: null,
          error: {
            message: 'build is not valid',
            fields: { 'point_order[0]': 'Talent 1001 is not in this class' },
          },
          request_id: 'r',
        }),
      });
    }
    return route.fulfill({ status: 201, contentType: 'application/json', body: JSON.stringify(SAVED) });
  });
  await page.goto('/planner');
  await page.getByTestId('talent-1001').click();
  await page.getByRole('button', { name: 'Share' }).click();
  await expect(page.getByRole('alert')).toContainText('Talent 1001 is not in this class');
  await page.getByRole('button', { name: 'Retry' }).click();
  await expect(page.getByTestId('share-link')).toBeVisible();
});
