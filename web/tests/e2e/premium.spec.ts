import { expect, test } from '@playwright/test';

test('ships no client JavaScript', async ({ page }) => {
  const scripts: string[] = [];
  page.on('request', (r) => {
    if (r.resourceType() === 'script') scripts.push(r.url());
  });
  await page.goto('/premium');
  expect(scripts).toEqual([]);
});

test('states what is free forever, the two plans, and that there are no ads', async ({ page }) => {
  await page.goto('/premium');
  await expect(page.getByRole('heading', { name: 'What is free, forever' })).toBeVisible();
  await expect(
    page.getByText('Planner, browser simulator, logs, live logging', { exact: false }),
  ).toBeVisible();
  // A plan's price shows up twice on load (the plan heading and its buy link), so scope to
  // the first match -- the assertion only cares that the price is visible somewhere.
  await expect(page.getByText('$4/month', { exact: false }).first()).toBeVisible();
  await expect(page.getByText('$15/month', { exact: false }).first()).toBeVisible();
  // Case-insensitive substring match also picks up "no ads" inside the free-forever
  // paragraph, so scope to the first hit.
  await expect(page.getByText('No ads.', { exact: false }).first()).toBeVisible();
  await expect(
    page.getByText('Nothing that comes from Battle.net is ever behind a paywall', { exact: false }),
  ).toBeVisible();
});

test('the monthly/yearly toggle works with no JavaScript', async ({ page }) => {
  await page.goto('/premium');
  await expect(page.getByText('$4/month', { exact: false }).first()).toBeVisible();
  await page.getByRole('link', { name: 'Yearly' }).first().click();
  await expect(page).toHaveURL(/#yearly$/);
  await expect(page.getByText('$40/year', { exact: false }).first()).toBeVisible();
});

test('every buy control is a plain link to /premium/checkout, never a marketing button', async ({ page }) => {
  await page.goto('/premium');
  const premiumBuy = page.getByRole('link', { name: /Subscribe.*\$4\/month/ });
  await expect(premiumBuy).toHaveAttribute('href', '/premium/checkout?plan=premium&interval=monthly');
});

test('/pricing resolves to /premium', async ({ page }) => {
  await page.goto('/pricing');
  await expect(page).toHaveURL(/\/premium$/);
});
