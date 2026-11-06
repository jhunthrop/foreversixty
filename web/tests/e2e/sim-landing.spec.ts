// web/tests/e2e/sim-landing.spec.ts
// The signed-in landing (spec 2026-11-06-sim §4.3): a member who opens /sim lands on the
// header's character, loaded -- its name as the title, the meta line, where the gear came
// from, and the run line -- with no character list (the header selector lists them; one way to
// switch). The paste Plate and Other ways in (tests/e2e/sim-sources.spec.ts) are what a
// signed-out visitor, or a member whose character cannot load, still gets.
//
// 2026-09-24 landing pass (owner-approved UX review) adds Wendel: a character the site
// holds no build for at all (`build` omitted, unlike Thrallgar and Roland below), whose
// sim-input read 404s -- Finding 1's "Paste export" link and Finding 2's build-missing
// alert both need one.
import { expect, test } from './support/test';
import { landingCopy } from '../../src/lib/sim/landing-copy';
import { simCopy } from '../../src/lib/sim/copy';
import { simPageCopy } from '../../src/lib/sim/page-copy';
import { gotoSim, simCharacter } from './support/sim-page';

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
/** The same account with another character as its main: the header (and so /sim) picks it. */
const meWithMain = (key: string) => ({ ...ME, main_character_key: key });

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

  // The header's character (here the site's own main-character guess, since this fixture stores
  // no pointer) loads on arrival: no list, no switcher, no engine hash.
  test('opens on the header character, loaded, not a list or the switcher', async ({ page }) => {
    await gotoSim(page);

    await expect(page.getByTestId('page-title')).toHaveText('Thrallgar');
    await expect(simCharacter(page)).toContainText('Fury Warrior');
    // The source line the API actually answered: an addon export, captured just now.
    await expect(page.getByTestId('sim-head-source')).toHaveText(/^Gear from the addon, \d{1,2} [A-Z][a-z]{2}$/);
    await expect(page.getByTestId('sim-settings')).toBeVisible();
    await expect(page.getByTestId('sim-landing')).toHaveCount(0);
    await expect(page.getByTestId('sim-sources')).toHaveCount(0);
    await expect(page.getByTestId('sim-answer-label')).toHaveText(simPageCopy.answer.noSim);
  });

  // The run line's Run sim runs the character the page arrived on.
  test('Run sim runs the header character', async ({ page }) => {
    await gotoSim(page);
    const run = page.getByTestId('sim-run-button');
    await expect(run).toHaveText(simPageCopy.bar.runSim);

    await run.click();

    await expect(page.getByTestId('sim-answer')).toHaveAttribute('data-state', 'result', { timeout: 15_000 });
    await expect(page.getByTestId('page-title')).toHaveText('Thrallgar');
    await expect(run).toHaveText(simPageCopy.bar.runAgain);
  });

  // Finding 2: the sim-input 404 gets its own notice, with the character's name and the way to
  // fix it, and the landing body (the paste Plate) stays up under it.
  test('a header character with no build shows the build-missing notice over the landing', async ({
    page,
  }) => {
    await page.route('**/v1/me', (route) => route.fulfill(envelope(meWithMain('us/normal/wendel'))));
    await gotoSim(page);

    const alert = page.getByTestId('sim-landing-build-missing');
    await expect(alert).toBeVisible();
    await expect(alert).toContainText('No build yet for Wendel.');
    await expect(alert).toContainText(landingCopy.buildMissingPasteLink);
    await expect(alert.getByRole('link', { name: landingCopy.buildMissingAccountLink })).toHaveAttribute(
      'href',
      landingCopy.buildMissingAccountHref,
    );
    // The race hint never follows a 404 (Ruling, Finding 2).
    await expect(page.getByTestId('sim-landing-message')).toHaveCount(0);
    await expect(page.getByTestId('sim-landing')).toBeVisible();
    await expect(page.getByTestId('sim-cold-paste')).toBeVisible();
    await expect(simCharacter(page)).toHaveCount(0);
  });

  // The race refusal names itself; the landing's notice says the character could not be read
  // (spec §5, "Character from the header"). LandingState's appended "paste your export" hint
  // retired with the list.
  test('a header character with no recorded race says why, and the landing stays up', async ({
    page,
  }) => {
    await page.route('**/v1/me', (route) => route.fulfill(envelope(meWithMain('us/normal/roland'))));
    await gotoSim(page);

    await expect(page.getByTestId('sim-landing-message')).toContainText(
      'That export names a race this build does not have: none recorded.',
    );
    await expect(page.getByTestId('sim-landing-build-missing')).toContainText(simCopy.characterFailed);
    // A failed load does not empty the page: the paste Plate is still the way in.
    await expect(page.getByTestId('sim-landing')).toBeVisible();
    await expect(simCharacter(page)).toHaveCount(0);
  });
});

test('a signed-out visitor gets the paste Plate and Other ways in, never a character', async ({ page }) => {
  await page.route('**/v1/me', (route) => route.fulfill(failure('not signed in', 401)));

  await gotoSim(page);

  await expect(page.getByTestId('sim-landing')).toBeVisible();
  await expect(page.getByTestId('sim-sources')).toBeVisible();
  await expect(page.getByTestId('sim-sign-in')).toHaveText(simPageCopy.ways.signIn);
  await expect(simCharacter(page)).toHaveCount(0);
});

test('a signed-in member with no characters gets the landing and the no-characters notice', async ({
  page,
}) => {
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME_NO_CHARACTERS)));

  await gotoSim(page);

  await expect(page.getByTestId('sim-sources')).toBeVisible();
  await expect(page.getByTestId('sim-landing')).toBeVisible();
  await expect(page.getByTestId('sim-no-characters')).toContainText(simPageCopy.noCharacters);
  await expect(page.getByTestId('sim-no-characters').getByRole('link')).toHaveAttribute(
    'href',
    simPageCopy.noCharactersHref,
  );
  await expect(simCharacter(page)).toHaveCount(0);
});

// Spec 2026-09-23 §3: the simulator reads the session through the client cache, so a
// returning signed-in visitor sees their character from the snapshot before /v1/me answers
// (the same shape as home-panel.spec.ts's returning-visitor test).
test('a returning signed-in visitor sees their character from the session snapshot before /v1/me answers', async ({
  page,
  context,
}) => {
  await context.addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await page.route('**/v1/characters/us/normal/thrallgar/sim-input', (route) =>
    route.fulfill(envelope(THRALLGAR_INPUT)),
  );
  await page.route('**/v1/sims?mine=1*', (route) =>
    route.fulfill(envelope({ rows: [], total: 0, page: 1, per_page: 100 })),
  );
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME)));
  await gotoSim(page);
  await expect(page.getByTestId('page-title')).toHaveText('Thrallgar');
  const stored = await page.evaluate(() =>
    Object.keys(window.localStorage)
      .filter((k) => k.startsWith('fs.q.'))
      .map((k) => window.localStorage.getItem(k) ?? '')
      .join('\n'),
  );
  expect(stored).toContain('Thrallgar');

  // Second load: /v1/me is held for five seconds, yet the character renders at once from the
  // snapshot and the landing skeleton never shows.
  await page.unroute('**/v1/me');
  await page.route('**/v1/me', async (route) => {
    await new Promise((resolve) => setTimeout(resolve, 5000));
    await route.fulfill(envelope(ME));
  });
  await page.goto('/sim');
  await expect(page.getByTestId('page-title')).toHaveText('Thrallgar', { timeout: 3000 });
  await expect(page.getByTestId('sim-landing-skeleton')).toHaveCount(0);
});

// 2026-10-31: the "Jump to" spine bar is gone from /sim; the header selector is the one
// current-character control, and nothing on the page repeats it.
test('/sim carries no current-character band: the head opens on the character itself', async ({ page }) => {
  await stubSignedIn(page);
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await gotoSim(page);
  await expect(page.getByTestId('page-title')).toHaveText('Thrallgar');
  await expect(page.getByTestId('current-character-bar')).toHaveCount(0);
  await expect(page.getByTestId('current-character-chip')).toHaveCount(0);
  await expect(page.getByTestId('sim-change-source')).toHaveCount(0);
  const title = await page.getByTestId('page-title').boundingBox();
  const bar = await page.getByTestId('sim-settings').boundingBox();
  expect(title).not.toBeNull();
  expect(bar).not.toBeNull();
  expect((title as { y: number }).y).toBeLessThan((bar as { y: number }).y);
});
