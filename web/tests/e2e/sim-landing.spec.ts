// web/tests/e2e/sim-landing.spec.ts
// Task 18: the signed-in landing state (design 4.6). A member who opens /sim sees their
// characters and one action each, and no form until they ask for one -- the source switcher
// (tests/e2e/sim-sources.spec.ts) is what every signed-out visitor still gets.
//
// 2026-09-24 landing pass (owner-approved UX review) adds Wendel: a character the site
// holds no build for at all (`build` omitted, unlike Thrallgar and Roland below), whose
// sim-input read 404s -- Finding 1's "Paste export" link and Finding 2's build-missing
// alert both need one.
import { expect, test } from '@playwright/test';
import { landingCopy } from '../../src/lib/sim/landing-copy';
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

// Three characters, each with a build (2026-09-24 landing pass, Finding 1: no build, no Sim
// button) except Wendel: Thrallgar's stored source carries a race (the success path -- the
// API's own answer today is an addon export or a logged fight, and this stub is the former,
// per sources.ts's own `sourcePill`); Roland's does not, which is the second, louder
// limitation the brief calls out -- `fromStoredCharacter` refuses rather than guessing one;
// Wendel carries no `build` at all, so his row gets no Sim button and no pill -- a "Paste
// export" link instead, and his own sim-input read 404s below.
const BUILD = { source: 'addon' as const, captured_at: '2026-09-20T00:00:00Z' };
const ME = {
  user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false, premium: false },
  characters: [
    {
      key: 'us/normal/thrallgar',
      region: 'us',
      ruleset: 'normal',
      name: 'Thrallgar',
      class: 'Warrior',
      build: BUILD,
    },
    { key: 'us/normal/roland', region: 'us', ruleset: 'normal', name: 'Roland', class: 'Mage', build: BUILD },
    { key: 'us/normal/wendel', region: 'us', ruleset: 'normal', name: 'Wendel', class: 'Priest' },
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
  // Wendel carries no build at all: the API's own answer for a character it holds nothing
  // for (Finding 2).
  await page.route('**/v1/characters/us/normal/wendel/sim-input', (route) =>
    route.fulfill(failure('no build', 404)),
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
    // Finding 5 (2026-09-24 landing pass): the landing state no longer carries its own copy
    // of the scope sentence -- ScopeNote.astro's own two sentences, above the island, are
    // the only copy of it now, and no engine-version hash precedes any of this (Finding 6).
    await expect(page.getByTestId('sim-landing-scope-note')).toHaveCount(0);
    await expect(page.getByTestId('sim-scope-note')).toHaveText(simCopy.scopeNote);
    await expect(page.getByTestId('sim-engine-version')).toHaveCount(0);

    // Fix round 1, MEDIUM-1: the row is a real link (`url.ts`'s own `simSearch`), not only
    // the button beside it -- copyable, middle-clickable, opens in a new tab.
    await expect(page.getByTestId('sim-character-link-us/normal/thrallgar')).toHaveAttribute(
      'href',
      '/sim?source=armory&ref=us%2Fnormal%2Fthrallgar',
    );
  });

  // Finding 1: a character with no build gets no Sim button -- a "Paste export" link and no
  // pill instead, since the action already says it.
  test('a character with no build gets a Paste export link, no Sim button, and no pill', async ({ page }) => {
    await page.goto('/sim');

    await expect(page.getByTestId('sim-pick-us/normal/wendel')).toHaveCount(0);
    const pasteLink = page.getByTestId('sim-paste-us/normal/wendel');
    await expect(pasteLink).toHaveText(landingCopy.pasteExport);
    await expect(pasteLink).toHaveAttribute('href', landingCopy.pasteExportHref);
    await expect(page.getByTestId('sim-character-build-us/normal/wendel')).toHaveCount(0);
  });

  // Finding 2: the sim-input 404 gets its own alert, with the character's name and both
  // links, in place of (not alongside) the old plain-text race hint -- which never follows
  // this failure (Ruling: `simCopy.landingNoRace` follows only the race refusal).
  test('picking a character with no build shows the build-missing alert, and the list stays up', async ({
    page,
  }) => {
    await page.goto('/sim');

    // Wendel has no Sim button (Finding 1), but the row's own name link still picks him in
    // place on a plain left click (`follow()`, LandingState.svelte) -- the same
    // fromStoredCharacter/fetchSimInput path a Sim press would take for a character that had
    // one, and the only way this specific failure (a 404, not the race refusal) is reached.
    await page.getByTestId('sim-character-link-us/normal/wendel').click();

    const alert = page.getByTestId('sim-landing-build-missing');
    await expect(alert).toBeVisible();
    await expect(alert).toContainText('No build yet for Wendel.');
    await expect(alert.getByRole('link', { name: landingCopy.buildMissingPasteLink })).toHaveAttribute(
      'href',
      landingCopy.pasteExportHref,
    );
    await expect(alert.getByRole('link', { name: landingCopy.buildMissingAccountLink })).toHaveAttribute(
      'href',
      landingCopy.buildMissingAccountHref,
    );
    // The old plain-text race hint never follows a 404 (Ruling, Finding 2).
    await expect(page.getByTestId('sim-landing-message')).toHaveCount(0);
    await expect(page.getByTestId('sim-landing')).toBeVisible();
    await expect(page.getByTestId('sim-character')).toHaveCount(0);
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

// Task 8: the spine bar mounts above the sim page's own current-character chip, on every
// /sim* page. The chip-slot pre-paint rule (global.css) hides both slots before hydration
// unless `fs_csrf` is present -- the same signed-in-visitor cookie logs-recent-reports.spec.ts's
// own spine-bar test sets -- so this test sets it too, rather than the bar's visibility
// assertion failing regardless of the /v1/me mock.
test('the spine bar reserves its slot before the sim chip slot, with no layout shift once both hydrate', async ({
  page,
}) => {
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME_NO_CHARACTERS)));
  await page.goto('/sim');
  const bar = page.getByTestId('current-character-bar');
  const chip = page.getByTestId('sim-chip-slot');
  await expect(bar).toBeVisible();
  const barBox = await bar.boundingBox();
  const chipBox = await chip.boundingBox();
  expect(barBox).not.toBeNull();
  expect(chipBox).not.toBeNull();
  expect((barBox as { y: number }).y).toBeLessThan((chipBox as { y: number }).y);
});
