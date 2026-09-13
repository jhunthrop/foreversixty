import { expect, test } from '@playwright/test';

test('the homepage tool card points at the live planner', async ({ page }) => {
  await page.goto('/');
  const card = page.getByRole('link', { name: /Build planner/ });
  await expect(card).toHaveAttribute('href', '/planner');
  await expect(card).toContainText('Live');
  await card.click();
  await expect(page).toHaveURL(/\/planner$/);
});

test('subscribing accepts an address and says what happens next', async ({ page }) => {
  let body: unknown;
  await page.route('**/v1/subscribe', async (route) => {
    body = route.request().postDataJSON();
    await route.fulfill({
      status: 202,
      contentType: 'application/json',
      body: JSON.stringify({ ok: true, data: { status: 'check your email' }, error: null, request_id: 'r' }),
    });
  });
  await page.goto('/');
  await page.getByTestId('subscribe').scrollIntoViewIfNeeded();
  await page.getByLabel('Email updates').fill('player@example.com');
  await page.getByRole('button', { name: 'Subscribe' }).click();
  await expect(page.getByTestId('subscribe-result')).toHaveText('Check your email and confirm the address.');
  expect(body).toEqual({ email: 'player@example.com' });
});

test('an unreachable API offers the mailto fallback instead', async ({ page }) => {
  await page.route('**/v1/subscribe', (route) => route.abort('failed'));
  await page.goto('/');
  await page.getByTestId('subscribe').scrollIntoViewIfNeeded();
  await page.getByLabel('Email updates').fill('player@example.com');
  await page.getByRole('button', { name: 'Subscribe' }).click();
  await expect(page.getByTestId('subscribe-result')).toContainText('The subscribe service is not answering');
  const mailto = page.getByTestId('subscribe-mailto');
  await expect(mailto).toBeVisible();
  await expect(mailto).toHaveAttribute('href', /^mailto:/);
});
