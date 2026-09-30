import { expect, test } from '@playwright/test';

const fulfil = (body: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify(body),
});

const NO_SIMS_PAGE = { rows: [], total: 0, page: 1, per_page: 20 };
const NO_RATING = null;

/** Stubs every fetch the signed-in hero and its cards make, beyond `/v1/me` itself, so a
 *  test that does not care about sims/ratings/talents still settles instead of hanging on
 *  a real network call (the real API's CORS policy refuses `localhost` outright, so an
 *  un-stubbed request never resolves at all in a local run). */
async function stubHeroExtras(page: import('@playwright/test').Page): Promise<void> {
  await page.route('**/v1/sims**', (route) => route.fulfill(fulfil(NO_SIMS_PAGE)));
  await page.route('**/rating', (route) => route.fulfill(fulfil(NO_RATING)));
  await page.route('**/v1/characters/**', (route) =>
    route.fulfill(fulfil({ ok: false, data: null, error: { message: 'none' }, request_id: 'r' }, 404)),
  );
}

test('the home page offers Battle.net sign-in when signed out', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(fulfil({ ok: false, data: null, error: null, request_id: 'r' }, 401)),
  );
  await page.goto('/');
  const signedOut = page.getByTestId('home-signed-out');
  await expect(signedOut).toBeVisible();
  // Home rebuild spec §1 / tenet 14: the one sentence this page is allowed to say about
  // itself, replacing the old pitch headline.
  await expect(signedOut.getByRole('heading', { level: 1 })).toHaveText('Play your class better.');
  await expect(page.getByRole('link', { name: 'Sign in with Battle.net' })).toHaveAttribute(
    'href',
    /\/v1\/auth\/battlenet\/start\?next=%2Faccount%3Fsigned_in%3D1$/,
  );
  await expect(page.getByTestId('home-account-panel')).toHaveCount(0);
  // The nine-class picker (§3.A.1): one row, in web/src/data/classes.json's own order.
  const picker = page.getByTestId('home-class-picker');
  await expect(picker.getByTestId(/^home-class-picker-/)).toHaveCount(9);
  await expect(picker.getByTestId('home-class-picker-warrior')).toContainText('Warrior');
  // §3.A.4: the signed-out "Best in slot by class" row is present; the signed-in-only
  // regions are not (display:none, per the session hint).
  await expect(page.getByTestId('home-class-picker-large')).toBeVisible();
  await expect(page.getByTestId('home-upgrades')).toBeHidden();
  await expect(page.getByTestId('home-another-class')).toBeHidden();
});

test('a signed-in visitor with zero characters still sees the Battle.net sign-in CTA, not a blank hero', async ({
  page,
}) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      fulfil({
        ok: true,
        data: {
          user: { id: 1, battletag: 'Fixture#1', email: null, role: 'user', anonymize: false },
          characters: [],
          guilds: [],
        },
        error: null,
        request_id: 'r',
      }),
    ),
  );
  await page.goto('/');
  await expect(page.getByTestId('home-account-panel')).toHaveCount(0);
  const signedOut = page.getByTestId('home-signed-out');
  await expect(signedOut).toBeVisible();
  await expect(signedOut).not.toHaveAttribute('inert');
  await expect(page.getByRole('link', { name: 'Sign in with Battle.net' })).toBeVisible();
});

test('a signed-in hero shows the descriptor, the three next-action cards, and hides the signed-out door', async ({
  page,
  context,
}) => {
  await stubHeroExtras(page);
  // The signed-out/signed-in region swap (Best in slot by class vs. Your upgrades/Another
  // class) reads the session cookie's readable half before first paint (Base.astro), not
  // the mocked /v1/me response -- the same pre-paint hint `data-session-hide` already
  // relies on elsewhere on this page.
  await context.addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      fulfil({
        ok: true,
        data: {
          user: { id: 1, battletag: 'Fixture#1', email: null, role: 'user', anonymize: false },
          characters: [
            {
              key: 'us/normal/zulmara',
              region: 'us',
              ruleset: 'normal',
              name: 'Zulmara',
              class: 'hunter',
              race: 'Troll',
              realm: 'Skyborne',
              level: 24,
              faction: 'horde',
              guild: { id: 9, name: 'Sample Guild', verified: true },
              build: { source: 'addon', captured_at: new Date(Date.now() - 4 * 60_000).toISOString() },
            },
          ],
          guilds: [],
        },
        error: null,
        request_id: 'r',
      }),
    ),
  );
  await page.addInitScript(() => {
    window.localStorage.setItem(
      'fs.currentCharacter',
      JSON.stringify({
        source: 'armory',
        ref: 'us/normal/zulmara',
        label: 'Zulmara · Hunter',
        classSlug: 'hunter',
        savedAt: new Date().toISOString(),
      }),
    );
  });
  await page.goto('/');
  const panel = page.getByTestId('home-account-panel');
  await expect(panel).toBeVisible();
  await expect(panel.getByRole('heading', { name: 'Zulmara', level: 1 })).toBeVisible();
  await expect(panel.getByTestId('home-hero-descriptor')).toContainText('Level 24 Troll Hunter');
  await expect(panel.getByTestId('home-hero-descriptor')).toContainText('Horde');
  await expect(panel.getByTestId('home-hero-guild-line')).toContainText('Sample Guild');
  await expect(panel.getByTestId('home-hero-sync')).toContainText('from the addon');
  // Best in slot / Talents: not yet computable against live worn-gear data (§3.B.2) --
  // an honest, settled line, never a fabricated figure.
  await expect(panel.getByTestId('home-hero-card-bis')).toContainText('Not available yet');
  await expect(panel.getByTestId('home-hero-card-talents')).toContainText('Not available yet');
  // Simulator: the one card with a real, live figure -- no saved sim for this fixture, so
  // the existing empty-state copy/action shows.
  await expect(panel.getByTestId('home-hero-card-sim')).toContainText('No sim yet.');
  await expect(page.locator('#home-signed-out')).toHaveAttribute('inert', '');
  await expect(page.locator('#home-signed-out')).toHaveAttribute('aria-hidden', 'true');
  // §3.A.4/§3.B.3/§3.B.5: the signed-out class-picker row is gone, replaced by the
  // signed-in-only "Your upgrades" (not-yet-available) and "Another class" regions.
  await expect(page.getByTestId('home-class-picker-large')).toBeHidden();
  await expect(page.getByTestId('home-upgrades')).toBeVisible();
  await expect(page.getByTestId('home-upgrades-not-yet-available')).toBeVisible();
  await expect(page.getByTestId('home-another-class').getByTestId(/^class-crest-/)).toHaveCount(9);
});

test('no Battle.net data for this realm type yet is shown once, when every character has no build', async ({
  page,
}) => {
  await stubHeroExtras(page);
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      fulfil({
        ok: true,
        data: {
          user: { id: 1, battletag: 'Fixture#1', email: null, role: 'user', anonymize: false },
          characters: [
            { key: 'us/normal/kiloz', region: 'us', ruleset: 'normal', name: 'Kiloz', class: 'warrior' },
          ],
          guilds: [],
        },
        error: null,
        request_id: 'r',
      }),
    ),
  );
  await page.goto('/');
  const panel = page.getByTestId('home-account-panel');
  await expect(panel.getByTestId('home-hero-no-bnet-data')).toHaveText(
    'Blizzard serves no data for this realm type yet.',
  );
  // No build at all: the sync line has nothing honest to say, so it is absent rather than
  // invented.
  await expect(panel.getByTestId('home-hero-sync')).toHaveCount(0);
});

test('the signed-in hero shows a rating figure once one exists, never before', async ({ page }) => {
  await page.route('**/v1/sims**', (route) => route.fulfill(fulfil(NO_SIMS_PAGE)));
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      fulfil({
        ok: true,
        data: {
          user: { id: 1, battletag: 'Fixture#1', email: null, role: 'user', anonymize: false },
          characters: [
            { key: 'us/normal/kiloz', region: 'us', ruleset: 'normal', name: 'Kiloz', class: 'Warrior' },
          ],
          guilds: [],
        },
        error: null,
        request_id: 'r',
      }),
    ),
  );
  await page.route('**/v1/characters/us/normal/kiloz/rating', (route) =>
    route.fulfill(
      fulfil({
        ok: true,
        data: {
          player_key: 'us/normal/kiloz',
          sample_size: 4,
          trend: [],
          best_component: 'damage',
          worst_component: 'utility',
          latest: { player_key: 'us/normal/kiloz', overall: 1.08 },
        },
        error: null,
        request_id: 'r',
      }),
    ),
  );
  await page.addInitScript(() => {
    window.localStorage.setItem(
      'fs.currentCharacter',
      JSON.stringify({
        source: 'armory',
        ref: 'us/normal/kiloz',
        label: 'Kiloz · Warrior',
        classSlug: 'warrior',
        savedAt: new Date().toISOString(),
      }),
    );
  });
  await page.goto('/');
  await expect(page.getByTestId('home-hero-rating')).toHaveText('Performance rating 1.08');
});

test('the Switch character panel lists every character, hero first, with a Current marker', async ({
  page,
}) => {
  await stubHeroExtras(page);
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      fulfil({
        ok: true,
        data: {
          user: { id: 1, battletag: 'Fixture#1', email: null, role: 'user', anonymize: false },
          characters: [
            {
              key: 'us/normal/kiloz',
              region: 'us',
              ruleset: 'normal',
              name: 'Kiloz',
              class: 'warrior',
              level: 60,
            },
            {
              key: 'us/normal/dottzz',
              region: 'us',
              ruleset: 'normal',
              name: 'Dottzz',
              class: 'priest',
              level: 12,
            },
          ],
          guilds: [],
        },
        error: null,
        request_id: 'r',
      }),
    ),
  );
  await page.goto('/');
  const switchPanel = page.getByTestId('home-switch-character-panel');
  await expect(switchPanel).toBeVisible();
  await expect(switchPanel).toContainText('Switch character');
  await expect(switchPanel.getByRole('link', { name: 'Add one' })).toHaveAttribute(
    'href',
    '/account#add-character',
  );
  await expect(switchPanel.getByTestId('current-character-bar-switch-current')).toBeVisible();
  const switchButton = switchPanel.getByTestId('current-character-bar-switch-us/normal/dottzz');
  await switchButton.click();
  await expect(page.getByTestId('home-account-panel').getByRole('heading', { level: 1 })).toHaveText(
    'Dottzz',
  );
});

test('a returning signed-in visitor sees the hub from the session snapshot before /v1/me answers', async ({
  page,
  context,
}) => {
  const body = {
    ok: true,
    data: {
      user: { id: 1, battletag: 'Fixture#1', email: null, role: 'user', anonymize: false },
      characters: [
        {
          key: 'us/normal/kiloz',
          region: 'us',
          ruleset: 'normal',
          name: 'Kiloz',
          class: 'warrior',
          level: 60,
        },
      ],
      guilds: [],
    },
    error: null,
    request_id: 'r',
  };
  await stubHeroExtras(page);
  // The readable half of the session cookie pair is what makes the snapshot trusted.
  await context.addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(body)));
  await page.goto('/');
  await expect(page.getByTestId('home-account-panel')).toBeVisible();
  const stored = await page.evaluate(() =>
    Object.keys(window.localStorage)
      .filter((k) => k.startsWith('fs.q.'))
      .map((k) => window.localStorage.getItem(k) ?? '')
      .join('\n'),
  );
  expect(stored).toContain('Kiloz');

  // Second load: /v1/me is held for five seconds, yet the hub renders at once from the
  // snapshot, and the signed-out block never shows (Base.astro's session hint).
  await page.unroute('**/v1/me');
  await page.route('**/v1/me', async (route) => {
    await new Promise((resolve) => setTimeout(resolve, 5000));
    await route.fulfill(fulfil(body));
  });
  await page.goto('/');
  await expect(page.getByTestId('home-account-panel')).toBeVisible({ timeout: 3000 });
  await expect(page.getByTestId('home-signed-out')).toBeHidden();
});
