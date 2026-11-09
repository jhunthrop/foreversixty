import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { gotoHydrated } from './support/hydrated';
import { CONNECT_STATES, openConnectState } from './support/connect-page';

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
  await expect(page.getByTestId('connect-code')).toHaveText('KQ4-7TM');
  await expect(page.getByRole('region', { name: 'The companion' })).toBeVisible();
  // Device and Platform keep their rows until a session can look the code up (ruling C2).
  await expect(page.getByTestId('connect-facts').getByText('Shown once you sign in')).toHaveCount(2);
  // The box signing in never moves: the same three rows at the same height once signed in. Read
  // off the page, not pinned: a Facts row is 41 px on macOS and rounds to 41.33 on Linux.
  const signedOutFacts = await page.getByTestId('connect-facts').boundingBox();
  await expect(page.getByText('Forever Sixty will see your BattleTag and character list.')).toBeVisible();
  const href = await page.getByTestId('connect-battlenet').getAttribute('href');
  expect(href).toContain('/v1/auth/battlenet/start?next=');
  expect(decodeURIComponent(href!.split('next=')[1])).toBe('/connect?code=KQ4-7TM');

  await expect(page.getByTestId('connect-email-toggle')).toHaveCount(0);

  // Back from sign-in: pending, the Facts in the same order and at the same height.
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
  await expect(page.getByTestId('connect-facts').locator('dt')).toHaveText(['Code', 'Device', 'Platform']);
  await expect(page.getByTestId('connect-signout')).toHaveText('Sign out.');
  await expect(page.getByTestId('connect-code-check')).toHaveText(
    'Only press Connect if this code matches the one in the companion window.',
  );
  const facts = await page.getByTestId('connect-facts').boundingBox();
  expect(facts?.height).toBe(signedOutFacts?.height);

  const approve = page.getByTestId('connect-approve');
  await approve.click();
  await expect(approve).toBeDisabled();
  await expect(approve).toHaveText('Connecting...');
  await expect(page.getByTestId('connect-deny')).toBeDisabled();
  release();
  await expect(page.getByRole('heading', { level: 1 })).toHaveText('Connected.');
  await expect(page.getByTestId('connect-check')).toBeVisible();
  await expect(page.getByTestId('connect-connected')).toHaveText('Connected');
  await expect(
    page.getByText(
      'Go back to the companion; it is ready. Now log out once in game, and your characters appear on the site within a minute.',
    ),
  ).toBeVisible();
  await expect(page.getByTestId('connect-account')).toHaveText('Your account');
  await expect(page.getByTestId('connect-account')).toHaveAttribute('href', '/account');
  await expect(page.getByTestId('connect-setup')).toHaveText('Get set up');
  await expect(page.getByTestId('connect-setup')).toHaveAttribute('href', '/setup');
  // Get set up is the one primary and comes first (rulings C3).
  const setupBox = await page.getByTestId('connect-setup').boundingBox();
  const accountBox = await page.getByTestId('connect-account').boundingBox();
  expect(setupBox!.x).toBeLessThan(accountBox!.x);
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

test('the plate fits above the fold at 360 x 700 and the Facts do not move between sign-in and pending', async ({
  page,
}) => {
  await page.setViewportSize({ width: 360, height: 700 });
  await signedOut(page);
  await gotoHydrated(page, '/connect?code=KQ4-7TM', 'connect');
  const before = await page.getByTestId('connect-facts').boundingBox();
  const button = await page.getByTestId('connect-battlenet').boundingBox();
  expect(button!.y + button!.height).toBeLessThanOrEqual(700);
  expect(await page.evaluate(() => document.documentElement.scrollWidth)).toBeLessThanOrEqual(360);

  await page.unroute('**/v1/me');
  await signedIn(page, { status: 200, body: DEVICE });
  await gotoHydrated(page, '/connect?code=KQ4-7TM', 'connect');
  await expect(page.getByTestId('connect-device')).toBeVisible();
  // The same box: signing in fills Device and Platform in their rows (the title above may wrap).
  const after = await page.getByTestId('connect-facts').boundingBox();
  expect(after!.height).toBe(before!.height);
  const deny = await page.getByTestId('connect-deny').boundingBox();
  expect(deny!.y + deny!.height).toBeLessThanOrEqual(700);
});

test('the platform reads as a player says it, and the form column is 560 px from 768', async ({ page }) => {
  await page.setViewportSize({ width: 1440, height: 900 });
  await signedIn(page, {
    status: 200,
    body: ok({ name: 'JHUNT-PC', platform: 'windows/amd64', expires_in: 540 }),
  });
  await gotoHydrated(page, '/connect?code=KQ4-7TM', 'connect');
  await expect(page.getByTestId('connect-facts').locator('dd').nth(2)).toHaveText('Windows');
  const plate = await page.getByRole('region', { name: 'The companion' }).boundingBox();
  expect(plate!.width).toBe(560);
});

// Review captures: only when CONNECT_SHOTS names a directory. Every state at the review widths,
// signed out and pending at 360 x 700 too (setup spec section 13).
const SHOTS = process.env.CONNECT_SHOTS;
test.describe('captures', () => {
  test.skip(SHOTS === undefined, 'set CONNECT_SHOTS to a directory to write the review PNGs');
  test.skip(({ isMobile }) => isMobile, 'the captures set their own viewports');
  const sizes = [
    ...[390, 1440, 1920, 2000].flatMap((width) =>
      CONNECT_STATES.map((state) => ({ state, width, height: 1000 })),
    ),
    { state: 'signed-out' as const, width: 360, height: 700 },
    { state: 'pending' as const, width: 360, height: 700 },
  ];
  for (const { state, width, height } of sizes) {
    test(`connect ${state} at ${width}x${height}`, async ({ page }) => {
      await page.setViewportSize({ width, height });
      await openConnectState(page, state);
      await page.evaluate(() => document.fonts.ready);
      const name = height === 700 ? `${width}x700` : `${width}`;
      await page.screenshot({ path: `${SHOTS}/connect-${state}-${name}.png`, fullPage: height !== 700 });
    });
  }
});
