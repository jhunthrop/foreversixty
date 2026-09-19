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
const THRALLGAR_INPUT = {
  spec: 'warrior-fury',
  gear: { head: 12640, main_hand: 11726 },
  talents: [2001, 2001, 2001, 2001, 2001, 2002, 2002],
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
  talents: [],
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
    await expect(page.getByTestId('sim-landing-note')).toHaveText(simCopy.landingSourceNote);
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
});
