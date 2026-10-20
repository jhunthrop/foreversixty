import { expect, test, type Page } from '@playwright/test';

import { meAddonFixture } from '../../src/fixtures/me-addon';

const ok = (data: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
});

const KEY = 'us/hardcore/obnoxious-yell';
const SAVED = { id: 'k7x2qm4a', url: 'https://foreversixty.gg/b/k7x2qm4a' };

function me() {
  const own = meAddonFixture.characters[0];
  return {
    ...meAddonFixture,
    characters: [
      {
        ...own,
        key: KEY,
        region: 'us',
        ruleset: 'hardcore',
        name: 'Obnoxious Yell',
        class: 'Warrior',
        level: 21,
      },
    ],
  };
}

// A leveling list for the 20 to 29 band, so "Load the 20 to 29 build" has a build to load.
const ARMS_LIST = {
  spec: 'warrior-arms',
  build: '1.60.1.70450',
  engine_version: 'test',
  generated_at: '2026-10-01T00:00:00Z',
  bands: (['alliance', 'horde'] as const).map((faction) => ({
    spec: 'warrior-arms',
    band: 20,
    faction,
    race: faction === 'horde' ? 'orc' : 'human',
    talents: '3-0-0',
    talent_points: 3,
    set_dps: 100,
    weights: [],
    slots: [],
  })),
};

async function signIn(page: Page, devices: unknown[] = [{ id: 'd1', name: 'Desk', platform: 'windows' }]) {
  await page.route('**/v1/me', (route) => route.fulfill(ok(me())));
  await page.route('**/v1/devices', (route) => route.fulfill(ok(devices)));
}

test.describe('Send this build to the addon', () => {
  test.skip(({ isMobile }) => isMobile, 'The card is asserted at desktop width');

  test('saves the build, queues it for the character, and reads Sent', async ({ page }) => {
    await signIn(page);
    const saves: unknown[] = [];
    const queued: unknown[] = [];
    await page.route('**/v1/builds', (route) => {
      saves.push(route.request().postDataJSON());
      return route.fulfill(ok(SAVED, 201));
    });
    await page.route('**/v1/addon/inbox', (route) => {
      queued.push(route.request().postDataJSON());
      return route.fulfill(ok({ build_id: SAVED.id }, 201));
    });

    await page.setViewportSize({ width: 1280, height: 900 });
    await page.route('**/data/*/bis/warrior-arms.json', (route) =>
      route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(ARMS_LIST) }),
    );
    await page.goto('/planner');
    const action = page.getByTestId('character-header-action');
    await expect(action).toHaveText('Send this build to the addon');
    await page.getByTestId('planner-header-band').screenshot({
      path: process.env.W2_CAPTURES ? `${process.env.W2_CAPTURES}/send-build-before.png` : undefined,
    });

    // An empty build is never sent: the button is disabled until points are spent.
    await expect(action).toBeDisabled();
    await expect(page.getByTestId('character-header-note')).toHaveText(
      "Spend some points first, or load the band's build.",
    );
    await page.getByTestId('band-compare-load').click();
    await expect(action).toBeEnabled();
    await expect(page.getByTestId('character-header-note')).toHaveCount(0);

    await action.click();

    await expect(action).toHaveText('Sent');
    await expect(action).toBeDisabled();
    await expect(page.getByTestId('character-header-note')).toContainText(
      'Sent to Obnoxious Yell. Log out to the character screen and back in to see it. /reload will not show it.',
    );
    await expect(page.getByTestId('character-header-note-detail')).toHaveText(
      'The companion checks every 10 minutes, and again when you log out.',
    );
    expect(saves).toHaveLength(1);
    expect(queued).toEqual([{ build_id: SAVED.id, character_key: KEY }]);
    await page.getByTestId('planner-header-band').screenshot({
      path: process.env.W2_CAPTURES ? `${process.env.W2_CAPTURES}/send-build-after.png` : undefined,
    });

    // Changing the build puts the button back; a second send reuses nothing stale.
    await page.getByTestId('talent-1002').click();
    await expect(action).toHaveText('Send this build to the addon');
    await expect(action).toBeEnabled();
  });

  test('an account with no companion is pointed at the companion setup', async ({ page }) => {
    await signIn(page, []);
    let saved = false;
    await page.route('**/v1/builds', (route) => {
      saved = true;
      return route.fulfill(ok(SAVED, 201));
    });
    await page.goto('/planner');
    await page.getByTestId('talent-1001').click();
    await page.getByTestId('character-header-action').click();
    await expect(page.getByTestId('character-header-note')).toContainText(
      'Run the companion once so builds can reach your addon.',
    );
    await expect(page.getByTestId('character-header-note-link')).toHaveAttribute('href', '/setup#companion');
    expect(saved).toBe(false);
  });

  test('a 401 from the inbox asks for a sign-in, other failures say try again', async ({ page }) => {
    await signIn(page);
    await page.route('**/v1/builds', (route) => route.fulfill(ok(SAVED, 201)));
    let status = 401;
    await page.route('**/v1/addon/inbox', (route) =>
      route.fulfill({
        status,
        contentType: 'application/json',
        body: JSON.stringify({ ok: false, data: null, error: { message: 'x' }, request_id: 'r' }),
      }),
    );
    await page.goto('/planner');
    await page.getByTestId('talent-1001').click();
    const action = page.getByTestId('character-header-action');
    await action.click();
    await expect(page.getByTestId('character-header-note')).toContainText(
      'Sign in to send builds to your addon.',
    );
    await expect(page.getByTestId('character-header-note-link')).toHaveAttribute('href', '/login');

    status = 500;
    await action.click();
    await expect(page.getByTestId('character-header-note')).toHaveText('Could not send just now. Try again.');
  });

  test('signed out there is no send button, only the way to the addon setup', async ({ page }) => {
    await page.goto('/planner');
    await expect(page.getByTestId('planner-header-band')).toBeVisible();
    await expect(page.getByRole('button', { name: /send this build/i })).toHaveCount(0);
    await expect(page.getByText('Send this build to the addon')).toHaveCount(0);
    await expect(page.getByTestId('character-header-action')).toHaveCount(0);
  });
});
