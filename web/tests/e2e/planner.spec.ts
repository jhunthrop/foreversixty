import { expect, test } from '@playwright/test';

test('the planner opens on the default class with an empty build', async ({ page }) => {
  const errors: string[] = [];
  page.on('console', (m) => {
    if (m.type() === 'error') errors.push(m.text());
  });
  await page.goto('/planner');
  await expect(page.getByTestId('planner-level')).toHaveText('9');
  await expect(page.getByTestId('planner-spent')).toHaveText('0/51');
  await expect(page.getByLabel('Class')).toHaveValue('warrior');
  await expect(
    page.getByText(
      'Classic Era trees shown until the beta client exports; Forever revamped talents replace them then.',
    ),
  ).toBeVisible();
  await expect(page.getByRole('heading', { name: 'Arms' })).toBeVisible();
  await expect(page.getByRole('heading', { name: 'Fury' })).toBeVisible();
  expect(errors).toEqual([]);
});

test('?class and ?race preselect the build', async ({ page }) => {
  await page.goto('/planner?class=warrior&race=dwarf');
  await expect(page.getByLabel('Class')).toHaveValue('warrior');
  await expect(page.getByLabel('Race')).toHaveValue('dwarf');
});

test('a failed talent fetch shows the reason and a working retry', async ({ page }) => {
  let attempts = 0;
  await page.route('**/data/*/talents/warrior.json', async (route) => {
    attempts += 1;
    if (attempts === 1) return route.abort('failed');
    return route.continue();
  });
  await page.goto('/planner');
  await expect(page.getByText('Talent data did not load')).toBeVisible();
  await page.getByRole('button', { name: 'Retry' }).click();
  await expect(page.getByRole('heading', { name: 'Arms' })).toBeVisible();
});
