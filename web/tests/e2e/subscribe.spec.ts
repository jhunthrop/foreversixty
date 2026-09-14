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

test('an unreachable API offers the Discord fallback instead', async ({ page }) => {
  await page.route('**/v1/subscribe', (route) => route.abort('failed'));
  await page.goto('/');
  await page.getByTestId('subscribe').scrollIntoViewIfNeeded();
  await page.getByLabel('Email updates').fill('player@example.com');
  await page.getByRole('button', { name: 'Subscribe' }).click();
  await expect(page.getByTestId('subscribe-result')).toContainText('The subscribe service is not answering');
  const fallback = page.getByTestId('subscribe-fallback');
  await expect(fallback).toBeVisible();
  await expect(fallback).toHaveAttribute('href', /^https:\/\/discord\.gg\//);
});

// SubscribeBox is client:visible and sits below the fold, so there is a real window between
// the user scrolling to it and its chunk executing. In that window the form has no submit
// handling of its own -- no action, no method, and submit() is what calls preventDefault() --
// so a click would run the browser's native submit and navigate off the page. Aborting the
// island's chunk makes that window permanent and observable; the button must be inert in it.
test('the subscribe button is inert until the island hydrates', async ({ page }) => {
  await page.route('**/SubscribeBox*.js', (route) => route.abort('failed'));
  await page.goto('/');
  await page.getByTestId('subscribe').scrollIntoViewIfNeeded();
  const button = page.getByRole('button', { name: 'Subscribe' });
  await expect(button).toBeDisabled();

  // The property that actually matters: a click must not submit the form and take the page
  // with it. force skips Playwright's own enabled check so the real browser decides.
  const before = page.url();
  await page.getByLabel('Email updates').fill('player@example.com');
  await button.click({ force: true });
  await page.waitForTimeout(500);
  expect(page.url()).toBe(before);
});

test('the subscribe button becomes usable once the island hydrates', async ({ page }) => {
  await page.goto('/');
  await page.getByTestId('subscribe').scrollIntoViewIfNeeded();
  await expect(page.getByRole('button', { name: 'Subscribe' })).toBeEnabled();
});
