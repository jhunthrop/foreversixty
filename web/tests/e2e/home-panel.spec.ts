import { expect, test } from '@playwright/test';

const fulfil = (body: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify(body),
});

test('the home page offers Battle.net sign-in when signed out', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(fulfil({ ok: false, data: null, error: null, request_id: 'r' }, 401)),
  );
  await page.goto('/');
  const signedOut = page.getByTestId('home-signed-out');
  await expect(signedOut).toBeVisible();
  // The pitch heading moved inside the signed-out block in this redesign.
  await expect(signedOut.getByRole('heading', { level: 1 })).toHaveText(
    'Your character, planned, simmed, logged and ranked.',
  );
  await expect(page.getByRole('link', { name: 'Sign in with Battle.net' })).toHaveAttribute(
    'href',
    /\/v1\/auth\/battlenet\/start\?next=%2Faccount%3Fsigned_in%3D1$/,
  );
  await expect(page.getByTestId('home-account-panel')).toHaveCount(0);
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

test('a character with no build shows Get the build, not a Sim button, with the portrait fallback', async ({
  page,
}, testInfo) => {
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
  const panel = page.getByTestId('home-account-panel');
  await expect(panel).toBeVisible();
  // The header is live on the home page too: a signed-in visitor sees their tag, not a
  // "Sign in" link beside their own hero.
  await expect(page.getByTestId('session-nav')).toContainText('Fixture#1');
  await expect(page.getByTestId('session-nav-static')).toHaveCount(0);
  // The character name is the page's only heading -- no testid is passed for it, so it is
  // located by role/name, the way a real user (or a screen reader) would find it.
  await expect(panel.getByRole('heading', { name: 'Kiloz', level: 1 })).toBeVisible();
  await expect(panel.getByRole('link', { name: 'Get the build' })).toHaveAttribute(
    'href',
    '/account#characters',
  );
  await expect(panel.getByRole('link', { name: 'Plan talents' })).toHaveAttribute('href', '/planner');
  await expect(panel.getByRole('link', { name: 'Logs' })).toHaveAttribute('href', '/logs');
  await expect(panel.getByRole('link', { name: 'Your characters' })).toHaveAttribute('href', '/account');
  // The signed-out "Sign in with Battle.net" link is only ever visually covered by the
  // grid-overlay CLS trick, never removed from the DOM -- without `inert`, a signed-in
  // keyboard/screen-reader user could still tab to, or hear, the duplicate link behind it.
  await expect(page.locator('#home-signed-out')).toHaveAttribute('inert', '');
  await expect(page.locator('#home-signed-out')).toHaveAttribute('aria-hidden', 'true');
  // No render_url on this fixture: CharacterIdentity's own inline portrait (the one
  // beside the name, testid "home-hero" passed through to its CharacterPortrait) is the
  // hero's sole visual identity here -- HomeAccountPanel renders no separate portrait of
  // its own when there is no render, so there is exactly one portrait, not two.
  await expect(panel.getByTestId('home-hero-avatar-fallback')).toBeVisible();
  await expect(panel.getByTestId('home-hero-render')).toHaveCount(0);
  // Every hero gets the class's tree art as a backdrop at desktop width, render or not
  // (lib/home/class-art.ts); it is decoration, so it is absent from the phone layout.
  const art = panel.getByTestId('home-hero-art');
  await expect(art).toHaveAttribute('style', /trees\/warriorarms\.webp/);
  if (testInfo.project.name === 'desktop') await expect(art).toBeVisible();
  else await expect(art).toBeHidden();
  // No character has a build: the reason is Blizzard's, and the hero says so once.
  await expect(panel.getByTestId('home-hero-no-bnet-data')).toHaveText(
    'Blizzard serves no data for this realm type yet.',
  );
});

test('a character with a build shows a Sim button to the armory-source href and its render image', async ({
  page,
}, testInfo) => {
  // A real (tiny) image, not just a mocked /v1/me url: the img has no explicit
  // width/height, so an unresolved src collapses its rendered box to 0x0 and a
  // visibility assertion below would pass or fail for the wrong reason.
  await page.route('**/render.jpg', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'image/png',
      // A minimal valid 1x1 transparent PNG.
      body: Buffer.from(
        'iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII=',
        'base64',
      ),
    }),
  );
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
              class: 'Warrior',
              render_url: 'https://example.test/render.jpg',
              build: { source: 'addon', captured_at: '2026-09-20T00:00:00Z' },
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
        ref: 'us/normal/kiloz',
        label: 'Kiloz · Warrior',
        classSlug: 'warrior',
        savedAt: new Date().toISOString(),
      }),
    );
  });
  await page.goto('/');
  const panel = page.getByTestId('home-account-panel');
  await expect(panel.getByRole('link', { name: 'Sim Kiloz' })).toHaveAttribute(
    'href',
    '/sim?source=armory&ref=us%2Fnormal%2Fkiloz',
  );
  await expect(panel.getByRole('link', { name: 'Get the build' })).toHaveCount(0);
  const renderImage = panel.getByTestId('home-hero-render');
  await expect(renderImage).toHaveAttribute('src', 'https://example.test/render.jpg');
  // `hidden lg:block`: visible once the image data loads on the desktop-width project,
  // and genuinely hidden (not just untested) below the `lg` breakpoint on mobile.
  if (testInfo.project.name === 'desktop') {
    await expect(renderImage).toBeVisible();
  } else {
    await expect(renderImage).toBeHidden();
  }
});

test('the hero shows a guild line when the character has one, with a verified mark', async ({ page }) => {
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
              class: 'Warrior',
              guild: { id: 9, name: 'Emerald Dream', verified: true },
            },
          ],
          guilds: [],
        },
        error: null,
        request_id: 'r',
      }),
    ),
  );
  await page.route('**/v1/characters/**', (route) =>
    route.fulfill(fulfil({ ok: false, data: null, error: { message: 'none' }, request_id: 'r' }, 404)),
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
  const panel = page.getByTestId('home-account-panel');
  const guildLine = panel.getByTestId('home-hero-guild-line');
  await expect(guildLine).toContainText('Emerald Dream');
  await expect(guildLine.getByTestId('home-hero-guild-verified')).toHaveText('Verified');
});

test('the signed-in hero shows a rating figure once one exists, never before', async ({ page }) => {
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

test('the hub lists the other characters as chips, and a chip makes that character current', async ({
  page,
}) => {
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
  await page.route('**/v1/characters/**', (route) =>
    route.fulfill(fulfil({ ok: false, data: null, error: { message: 'none' }, request_id: 'r' }, 404)),
  );
  await page.goto('/');
  const panel = page.getByTestId('home-account-panel');
  await expect(panel).toBeVisible();
  // Kiloz is the main character (highest level); Dottzz is the one chip.
  const chip = panel.getByTestId('home-character-chip');
  await expect(chip).toHaveCount(1);
  await expect(chip).toContainText('Dottzz');
  await chip.click();
  await expect(panel).toContainText('Dottzz');
  await expect(panel.getByTestId('home-character-chip')).toContainText('Kiloz');
  const stored = await page.evaluate(() => window.localStorage.getItem('fs.currentCharacter') ?? '');
  expect(stored).toContain('us/normal/dottzz');
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
  // The readable half of the session cookie pair is what makes the snapshot trusted.
  await context.addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await page.route('**/v1/characters/**', (route) =>
    route.fulfill({ status: 404, json: { ok: false, data: null, error: { message: 'none' } } }),
  );
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(body)));
  await page.goto('/');
  await expect(page.getByTestId('home-account-panel')).toBeVisible();
  // The data module persists under its own keys (lib/data/query.ts); the session entry is
  // whichever one holds the account.
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
