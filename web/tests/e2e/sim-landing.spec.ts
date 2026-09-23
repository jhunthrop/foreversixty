// web/tests/e2e/sim-landing.spec.ts
// Task 18: the signed-in landing state (design 4.6). A member who opens /sim sees their
// characters and one button each, and no form until they ask for one -- the source switcher
// (tests/e2e/sim-sources.spec.ts) is what every signed-out visitor still gets.
import { expect, test } from '@playwright/test';
import { simCopy } from '../../src/lib/sim/copy';

function envelope(data: unknown, status = 200) {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'req-test' }),
  };
}

function failure(message: string, status: number) {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: false, data: null, error: { message }, request_id: 'req-test' }),
  };
}

// Two characters: Thrallgar's stored source carries a race (the success path -- the API's
// own answer today is an addon export or a logged fight, and this stub is the former, per
// sources.ts's own `sourcePill`); Roland's does not, which is the second, louder limitation
// the brief calls out -- `fromStoredCharacter` refuses rather than guessing one.
const ME = {
  user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false, premium: false },
  characters: [
    { key: 'us/normal/thrallgar', region: 'us', ruleset: 'normal', name: 'Thrallgar', class: 'Warrior' },
    { key: 'us/normal/roland', region: 'us', ruleset: 'normal', name: 'Roland', class: 'Mage' },
  ],
  guilds: [],
};

const ME_NO_CHARACTERS = { ...ME, characters: [] };

// Computed at test run time so the strip's relative-time pill reads "just now" regardless of
// when the suite runs, the same way sim-sources.spec.ts's own addon-paste test gets there.
// gear and talents match input.go's real shapes (H3, final whole-branch review): gear is
// the source's own opaque JSON, never the planner's slot-to-item map, and talents is
// fight_metrics.talent_split -- points per tree, never a per-talent id array.
const THRALLGAR_INPUT = {
  spec: 'warrior-fury',
  gear: { slots: [12640, 11726] },
  talents: '31/0/20',
  buffs: ['battle_shout'],
  race: 'orc',
  captured_at: new Date().toISOString(),
  source: 'addon',
};

// No `race` field at all: the shape `fromStoredCharacter` reads today, from an API that has
// not started sending one for a stored character (the brief's own "second limitation").
const ROLAND_INPUT = {
  spec: 'mage-fire',
  gear: {},
  talents: '',
  buffs: [],
  captured_at: new Date().toISOString(),
  source: 'fight',
};

async function stubSignedIn(page: import('@playwright/test').Page): Promise<void> {
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME)));
  await page.route('**/v1/characters/us/normal/thrallgar/sim-input', (route) =>
    route.fulfill(envelope(THRALLGAR_INPUT)),
  );
  await page.route('**/v1/characters/us/normal/roland/sim-input', (route) =>
    route.fulfill(envelope(ROLAND_INPUT)),
  );
  // Signed-in fires the history panel's own load (SimView's onMount); stubbed empty so it
  // settles rather than reaching the real API this suite has no server for.
  await page.route('**/v1/sims?mine=1*', (route) =>
    route.fulfill(envelope({ rows: [], total: 0, page: 1, per_page: 100 })),
  );
}

test.describe('a signed-in member with characters', () => {
  test.beforeEach(async ({ page }) => {
    await stubSignedIn(page);
  });

  test('opens on the landing state, not the switcher', async ({ page }) => {
    await page.goto('/sim');

    await expect(page.getByTestId('sim-landing')).toBeVisible();
    await expect(page.getByTestId('sim-character-us/normal/thrallgar')).toBeVisible();
    await expect(page.getByTestId('sim-character-us/normal/roland')).toBeVisible();
    await expect(page.getByTestId('sim-sources')).toHaveCount(0);
    // task-2-brief.md: the landing state is what a signed-in member reads first, so it
    // carries the DPS-only scope sentence too, not only the Astro shell above it.
    await expect(page.getByTestId('sim-landing-scope-note')).toHaveText(simCopy.scopeNote);

    // Fix round 1, MEDIUM-1: the row is a real link (`url.ts`'s own `simSearch`), not only
    // the button beside it -- copyable, middle-clickable, opens in a new tab.
    await expect(page.getByTestId('sim-character-link-us/normal/thrallgar')).toHaveAttribute(
      'href',
      '/sim?source=armory&ref=us%2Fnormal%2Fthrallgar',
    );
  });

  test('pressing Sim loads the strip, with the source pill the API actually answered', async ({ page }) => {
    await page.goto('/sim');

    await page.getByTestId('sim-pick-us/normal/thrallgar').click();

    await expect(page.getByTestId('sim-character')).toBeVisible();
    await expect(page.getByTestId('sim-source-pill')).toHaveText('Addon export, just now');
    await expect(page.getByTestId('sim-landing')).toHaveCount(0);
  });

  test('a character with no recorded race fails with the addon hint, and the list stays up', async ({
    page,
  }) => {
    await page.goto('/sim');

    await page.getByTestId('sim-pick-us/normal/roland').click();

    await expect(page.getByTestId('sim-landing-message')).toContainText(
      'That export names a race this build does not have: none recorded.',
    );
    await expect(page.getByTestId('sim-landing-message')).toContainText(simCopy.landingNoRace);
    // A failed pick does not empty the page: the list is still the whole point of being here.
    await expect(page.getByTestId('sim-landing')).toBeVisible();
    await expect(page.getByTestId('sim-character')).toHaveCount(0);
  });

  test('Sim something else opens the switcher', async ({ page }) => {
    await page.goto('/sim');
    await expect(page.getByTestId('sim-landing')).toBeVisible();

    await page.getByTestId('sim-other-character').click();

    await expect(page.getByTestId('sim-sources')).toBeVisible();
    await expect(page.getByTestId('sim-landing')).toHaveCount(0);
  });

  // Fix round 1, LOW-2: the `onback` round trip had no automated coverage, only a
  // by-hand trace in the review. Both directions the review traced: landing -> switcher ->
  // landing (nothing loaded yet), and strip -> switcher -> strip (a character already up).
  test('Back to your characters returns to the landing state from the switcher', async ({ page }) => {
    await page.goto('/sim');
    await page.getByTestId('sim-other-character').click();
    await expect(page.getByTestId('sim-sources')).toBeVisible();

    await page.getByTestId('sim-back-to-characters').click();

    await expect(page.getByTestId('sim-landing')).toBeVisible();
    await expect(page.getByTestId('sim-sources')).toHaveCount(0);
  });

  test('Back to your characters returns to a loaded strip, not the landing state', async ({ page }) => {
    await page.goto('/sim');
    await page.getByTestId('sim-pick-us/normal/thrallgar').click();
    await expect(page.getByTestId('sim-character')).toBeVisible();

    await page.getByTestId('sim-change-source').click();
    await expect(page.getByTestId('sim-sources')).toBeVisible();

    await page.getByTestId('sim-back-to-characters').click();

    await expect(page.getByTestId('sim-character')).toBeVisible();
    await expect(page.getByTestId('sim-sources')).toHaveCount(0);
  });
});

test('a signed-out visitor gets the switcher, never the landing state', async ({ page }) => {
  await page.route('**/v1/me', (route) => route.fulfill(failure('not signed in', 401)));

  await page.goto('/sim');

  await expect(page.getByTestId('sim-sources')).toBeVisible();
  await expect(page.getByTestId('sim-landing')).toHaveCount(0);
});

test('a signed-in member with no characters gets the switcher and the noCharactersYet line', async ({
  page,
}) => {
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME_NO_CHARACTERS)));

  await page.goto('/sim');

  await expect(page.getByTestId('sim-sources')).toBeVisible();
  await expect(page.getByTestId('sim-landing')).toHaveCount(0);
  await expect(page.getByTestId('sim-no-characters')).toContainText(simCopy.noCharactersYet);
  // Nothing to go back to: the account card carries no "Back to your characters" button
  // here (it used to, and clicking it re-rendered this same state).
  await expect(page.getByTestId('sim-back-to-characters')).toHaveCount(0);
});

// Spec 2026-09-23 §3: the simulator reads the session through the client cache, so a
// returning signed-in visitor sees their characters from the snapshot before /v1/me answers
// (the same shape as home-panel.spec.ts's returning-visitor test).
test('a returning signed-in visitor sees their characters from the session snapshot before /v1/me answers', async ({
  page,
  context,
}) => {
  await context.addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await page.route('**/v1/characters/**', (route) => route.fulfill(failure('none', 404)));
  await page.route('**/v1/sims?mine=1*', (route) =>
    route.fulfill(envelope({ rows: [], total: 0, page: 1, per_page: 100 })),
  );
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME)));
  await page.goto('/sim');
  await expect(page.getByTestId('sim-landing')).toBeVisible();
  const stored = await page.evaluate(() =>
    Object.keys(window.localStorage)
      .filter((k) => k.startsWith('fs.q.'))
      .map((k) => window.localStorage.getItem(k) ?? '')
      .join('\n'),
  );
  expect(stored).toContain('Thrallgar');

  // Second load: /v1/me is held for five seconds, yet the landing state renders at once from
  // the snapshot and the skeleton never shows.
  await page.unroute('**/v1/me');
  await page.route('**/v1/me', async (route) => {
    await new Promise((resolve) => setTimeout(resolve, 5000));
    await route.fulfill(envelope(ME));
  });
  await page.goto('/sim');
  await expect(page.getByTestId('sim-landing')).toBeVisible({ timeout: 3000 });
  await expect(page.getByTestId('sim-landing-skeleton')).toHaveCount(0);
});
