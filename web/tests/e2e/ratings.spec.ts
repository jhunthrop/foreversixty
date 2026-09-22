// web/tests/e2e/ratings.spec.ts
// /ratings is a static content page: no client JavaScript, no layout shift, and it must
// publish every fact spec section 6.3 requires -- the weight table, the percentile-vs
// -absolute rule, the cap in the coordinator's exact words, and that a rating comes only
// from a public report.
import { expect, test } from '@playwright/test';

test('ships no client JavaScript', async ({ page }) => {
  const scripts: string[] = [];
  page.on('request', (r) => {
    if (r.resourceType() === 'script') scripts.push(r.url());
  });
  await page.goto('/ratings');
  expect(scripts).toEqual([]);
});

test('publishes the weight table for all three roles', async ({ page }) => {
  await page.goto('/ratings');
  await expect(page.getByRole('heading', { name: 'How ratings work' })).toBeVisible();
  const table = page.getByTestId('ratings-weights');
  await expect(table).toBeVisible();
  await expect(table).toContainText('Output');
  await expect(table).toContainText('Survival');
  await expect(table).toContainText('Mechanics');
  await expect(table).toContainText('Utility');
  await expect(table).toContainText('Preparation');
  await expect(table).toContainText('Activity');
  await expect(table).toContainText('DPS');
  await expect(table).toContainText('Healer');
  await expect(table).toContainText('Tank');
});

test('states the cap in the coordinator’s exact words', async ({ page }) => {
  await page.goto('/ratings');
  await expect(
    page.getByText(
      'An avoidable death early in a fight caps the overall score at 40, because nothing else in the fight makes up for it.',
    ),
  ).toBeVisible();
});

test('says ratings are public like Warcraft Logs parses, and only from public reports', async ({ page }) => {
  await page.goto('/ratings');
  await expect(page.getByTestId('ratings-visibility')).toContainText('public');
  await expect(page.getByTestId('ratings-visibility')).toContainText('report');
});

test('explains the percentile-vs-absolute rule, small samples, and role-specific weights', async ({
  page,
}) => {
  await page.goto('/ratings');
  await expect(page.getByTestId('ratings-basis')).toContainText('20');
  await expect(page.getByTestId('ratings-small-samples')).toBeVisible();
  await expect(page.getByTestId('ratings-roles')).toContainText('tanks and healers');
});

test('says a wipe scores differently than a kill', async ({ page }) => {
  await page.goto('/ratings');
  await expect(page.getByTestId('ratings-wipes')).toBeVisible();
});
