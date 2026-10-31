import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { gotoHydrated } from './support/hydrated';

function fulfil(body: unknown, status = 200) {
  return { status, contentType: 'application/json', body: JSON.stringify(body) };
}

const ok = (data: unknown) => ({ ok: true, data, error: null, request_id: 'r' });
const failure = (message: string) => ({ ok: false, data: null, error: { message }, request_id: 'r' });

const ME = ok({
  user: { id: 7, battletag: 'Thrall#1234', email: null, role: 'user', anonymize: false },
  characters: [],
  guilds: [],
});

async function signedOut(page: Page) {
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(failure('Sign in to continue'), 401)));
}

async function signedIn(page: Page, lookup: { status: number; body: unknown }) {
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(ME)));
  await page.route('**/v1/devices/connect/lookup**', (route) =>
    route.fulfill(fulfil(lookup.body, lookup.status)),
  );
}

const DEVICE = ok({ name: 'JHUNT-PC', platform: 'Windows', expires_in: 540 });

test('not signed in: the code, the Battle.net hand-off with next, then pending and success', async ({
  page,
}) => {
  await signedOut(page);
  await gotoHydrated(page, '/connect?code=kq47tm', 'connect');

  await expect(page.getByRole('heading', { level: 1 })).toHaveText('Connect your companion');
  await expect(page.getByTestId('connect-code')).toHaveText('Code KQ4-7TM');
  await expect(page.getByText('Forever Sixty will see your BattleTag and character list.')).toBeVisible();
  const href = await page.getByTestId('connect-battlenet').getAttribute('href');
  expect(href).toContain('/v1/auth/battlenet/start?next=');
  expect(decodeURIComponent(href!.split('next=')[1])).toBe('/connect?code=KQ4-7TM');

  // The email form carries the same next.
  let emailBody: unknown = null;
  await page.route('**/v1/auth/email', (route) => {
    emailBody = route.request().postDataJSON();
    return route.fulfill(fulfil(ok({ sent: true }), 202));
  });
  await page.getByTestId('connect-email-toggle').click();
  await page.getByLabel('Email address').fill('thrall@example.com');
  await page.getByRole('button', { name: 'Send link' }).click();
  await expect(page.getByTestId('connect-email-note')).toBeVisible();
  expect(emailBody).toEqual({ email: 'thrall@example.com', next: '/connect?code=KQ4-7TM' });

  // Back from sign-in: pending, with the device panel reserved at 88px.
  await page.unroute('**/v1/me');
  await signedIn(page, { status: 200, body: DEVICE });
  let approveBody: unknown = null;
  let release: () => void = () => {};
  const gate = new Promise<void>((resolve) => (release = resolve));
  await page.route('**/v1/devices/connect/approve', async (route) => {
    approveBody = route.request().postDataJSON();
    await gate;
    await route.fulfill(fulfil(ok({ name: 'JHUNT-PC', platform: 'Windows' })));
  });
  await gotoHydrated(page, '/connect?code=KQ4-7TM', 'connect');
  await expect(page.getByRole('heading', { level: 1 })).toHaveText('Connect JHUNT-PC to Thrall#1234?');
  await expect(page.getByTestId('connect-device')).toHaveText('JHUNT-PC');
  await expect(page.getByTestId('connect-code-check')).toHaveText(
    'Only press Connect if this code matches the one in the companion window.',
  );
  const panel = await page.getByTestId('connect-panel').boundingBox();
  expect(panel?.height).toBe(88);

  const approve = page.getByTestId('connect-approve');
  await approve.click();
  await expect(approve).toBeDisabled();
  await expect(approve).toHaveText('Connecting...');
  release();
  await expect(page.getByRole('heading', { level: 1 })).toHaveText('Connected.');
  await expect(
    page.getByText(
      'Connected. Go back to the companion; it is ready. Now log out once in game, and your characters appear on the site within a minute.',
    ),
  ).toBeVisible();
  await expect(page.getByTestId('connect-account')).toHaveText('Your account');
  await expect(page.getByTestId('connect-account')).toHaveAttribute('href', '/account');
  await expect(page.getByTestId('connect-setup')).toHaveText('Get set up');
  await expect(page.getByTestId('connect-setup')).toHaveAttribute('href', '/setup');
  await expect(page.getByTestId('connect-reports')).toHaveCount(0);
  expect(approveBody).toEqual({ code: 'KQ4-7TM' });
});

test('an unknown or expired code lands on the expired state', async ({ page }) => {
  await signedIn(page, { status: 404, body: failure('not found') });
  await gotoHydrated(page, '/connect?code=KQ4-7TM', 'connect');
  await expect(page.getByRole('heading', { level: 1 })).toHaveText('This link has expired.');
  await expect(page.getByRole('link', { name: 'Back to Forever Sixty' })).toBeVisible();
});

test('already used, unreachable and cancelled', async ({ page }) => {
  await signedIn(page, { status: 200, body: DEVICE });
  await page.route('**/v1/devices/connect/approve', (route) =>
    route.fulfill(fulfil(failure('already decided'), 409)),
  );
  await gotoHydrated(page, '/connect?code=KQ4-7TM', 'connect');
  await page.getByTestId('connect-approve').click();
  await expect(page.getByRole('heading', { level: 1 })).toHaveText('This device is already connected.');

  await page.unroute('**/v1/devices/connect/approve');
  await page.route('**/v1/devices/connect/approve', (route) => route.abort());
  await gotoHydrated(page, '/connect?code=KQ4-7TM', 'connect');
  await page.getByTestId('connect-approve').click();
  await expect(page.getByRole('heading', { level: 1 })).toHaveText('Could not reach the site.');
  await expect(page.getByTestId('connect-retry')).toBeVisible();

  await page.unroute('**/v1/devices/connect/approve');
  let denied: unknown = null;
  await page.route('**/v1/devices/connect/deny', (route) => {
    denied = route.request().postDataJSON();
    return route.fulfill({ status: 204 });
  });
  await gotoHydrated(page, '/connect?code=KQ4-7TM', 'connect');
  await page.getByTestId('connect-deny').click();
  await expect(page.getByRole('heading', { level: 1 })).toHaveText('Not connected.');
  expect(denied).toEqual({ code: 'KQ4-7TM' });
});

test('the card fits above the fold at 360px and the device panel does not move between sign-in and pending', async ({
  page,
}) => {
  await page.setViewportSize({ width: 360, height: 700 });
  await signedOut(page);
  await gotoHydrated(page, '/connect?code=KQ4-7TM', 'connect');
  const before = await page.getByTestId('connect-panel').boundingBox();
  const button = await page.getByTestId('connect-battlenet').boundingBox();
  expect(button!.y + button!.height).toBeLessThanOrEqual(700);
  expect(await page.evaluate(() => document.documentElement.scrollWidth)).toBeLessThanOrEqual(360);

  await page.unroute('**/v1/me');
  await signedIn(page, { status: 200, body: DEVICE });
  await gotoHydrated(page, '/connect?code=KQ4-7TM', 'connect');
  // The card is only as tall as its content now; the 88px panel is the part that must stay put.
  const after = await page.getByTestId('connect-panel').boundingBox();
  expect(after!.height).toBe(before!.height);
  const deny = await page.getByTestId('connect-deny').boundingBox();
  expect(deny!.y + deny!.height).toBeLessThanOrEqual(700);
});

// Review captures: only when CONNECT_SHOTS names a directory.
const SHOTS = process.env.CONNECT_SHOTS;
test.describe('captures', () => {
  test.skip(SHOTS === undefined, 'set CONNECT_SHOTS to a directory to write the review PNGs');
  for (const width of [360, 1280]) {
    test(`site connect at ${width}`, async ({ page }) => {
      await page.setViewportSize({ width, height: width === 360 ? 780 : 800 });
      const shot = (state: string) =>
        page.screenshot({ path: `${SHOTS}/site-connect-${state}-${width}.png` });

      await signedOut(page);
      await gotoHydrated(page, '/connect?code=KQ4-7TM', 'connect');
      await shot('not-signed-in');

      await page.unroute('**/v1/me');
      await signedIn(page, { status: 200, body: DEVICE });
      await page.route('**/v1/devices/connect/approve', (route) =>
        route.fulfill(fulfil(ok({ name: 'JHUNT-PC', platform: 'Windows' }))),
      );
      await gotoHydrated(page, '/connect?code=KQ4-7TM', 'connect');
      await shot('pending');
      await page.getByTestId('connect-approve').click();
      await expect(page.getByTestId('connect-check')).toBeVisible();
      await shot('success');

      await page.unroute('**/v1/devices/connect/lookup**');
      await page.route('**/v1/devices/connect/lookup**', (route) =>
        route.fulfill(fulfil(failure('not found'), 404)),
      );
      await gotoHydrated(page, '/connect?code=KQ4-7TM', 'connect');
      await shot('expired');

      await page.unroute('**/v1/devices/connect/lookup**');
      await signedIn(page, { status: 200, body: DEVICE });
      await page.unroute('**/v1/devices/connect/deny');
      await page.route('**/v1/devices/connect/deny', (route) => route.fulfill({ status: 204 }));
      await gotoHydrated(page, '/connect?code=KQ4-7TM', 'connect');
      await page.getByTestId('connect-deny').click();
      await expect(page.getByTestId('connect-close')).toBeVisible();
      await shot('cancelled');
    });
  }
});
