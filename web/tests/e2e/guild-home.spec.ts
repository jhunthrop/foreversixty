// web/tests/e2e/guild-home.spec.ts
import { expect, test } from '@playwright/test';

function envelope(data: unknown, status = 200) {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'req-test' }),
  };
}

const GUILD_PAGE = {
  guild: { id: 501, name: 'The Last Watch', region: 'us', ruleset: 'hardcore' },
  progression: [],
  roster_best: [],
  reports: [],
};

const ME = {
  user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false, premium: false },
  characters: [],
  guilds: [{ id: 501, region: 'us', ruleset: 'hardcore', name: 'The Last Watch', rank: 'officer' }],
};

const HOME = {
  guild: { id: 501, name: 'The Last Watch', region: 'us', ruleset: 'hardcore' },
  viewer: { rank: 'officer', verified: true, can_manage: true },
  reports: {
    rows: [
      {
        id: 'fixture2abcd',
        title: 'Sanguine Depths',
        zone: 'Sanguine Depths',
        created_at: '2026-09-20T20:00:00Z',
        kill_count: 3,
        wipe_count: 5,
      },
    ],
  },
  roster: [
    {
      character_key: 'us/hardcore/simfury',
      region: 'us',
      ruleset: 'hardcore',
      name: 'Simfury',
      class: 'Warrior',
      spec: 'Fury',
      rank: 'officer',
      verified: true,
      logged_recently: true,
      consent: 'gear',
      user_id: 7,
    },
    {
      character_key: 'us/hardcore/newbie',
      region: 'us',
      ruleset: 'hardcore',
      name: 'Newbie',
      class: 'Priest',
      spec: 'Holy',
      rank: 'member',
      verified: false,
      logged_recently: false,
      consent: 'roster',
      user_id: 42,
    },
  ],
};

async function stub(page: import('@playwright/test').Page) {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) => route.fulfill(envelope(GUILD_PAGE)));
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME)));
  await page.route('**/v1/guilds/501/home', (route) => route.fulfill(envelope(HOME)));
}

test('a signed-in officer sees this week’s reports, who logged, and the roster with approve/remove', async ({
  page,
}) => {
  await stub(page);
  await page.goto('/guild/us/hardcore/the-last-watch');

  await expect(page.getByTestId('guild-home')).toBeVisible();
  await expect(page.getByTestId('guild-home-reports')).toContainText('3 kills · 5 wipes');
  await expect(page.getByTestId('guild-home-roster')).toContainText('Simfury');
  await expect(page.getByTestId('guild-roster-unverified')).toBeVisible();
  await expect(page.getByTestId('guild-roster-approve')).toBeVisible();
  await expect(page.getByTestId('guild-settings-link')).toHaveAttribute(
    'href',
    '/guild/us/hardcore/the-last-watch/settings',
  );
});

test('approving an unverified character calls the approve endpoint and removes the unverified pill', async ({
  page,
}) => {
  await stub(page);
  await page.route('**/v1/guilds/501/characters/us%2Fhardcore%2Fnewbie/approve', (route) =>
    route.fulfill(
      envelope({
        character_key: 'us/hardcore/newbie',
        region: 'us',
        ruleset: 'hardcore',
        name: 'Newbie',
        class: 'Priest',
        spec: 'Holy',
        rank: 'member',
        verified: true,
        logged_recently: false,
        consent: 'roster',
      }),
    ),
  );
  await page.goto('/guild/us/hardcore/the-last-watch');
  await page.getByTestId('guild-roster-approve').click();
  await expect(page.getByTestId('guild-roster-unverified')).toHaveCount(0);
});

test('a roster row at roster-only consent gets no hand-off links; gear consent does', async ({ page }) => {
  await stub(page);
  await page.route('**/v1/characters/us/hardcore/simfury/sim-input', (route) =>
    route.fulfill(
      envelope({
        spec: 'warrior-fury',
        gear: 'FS1:1.60.1.69893:warrior:human:0/0/0:',
        talents: '',
        buffs: [],
        captured_at: new Date().toISOString(),
        source: 'addon',
      }),
    ),
  );
  await page.goto('/guild/us/hardcore/the-last-watch');

  const simfuryRow = page
    .getByTestId('guild-home-roster')
    .getByRole('listitem')
    .filter({ hasText: 'Simfury' });
  await expect(simfuryRow.getByTestId('guild-roster-open-sim')).toBeVisible();

  const newbieRow = page.getByTestId('guild-home-roster').getByRole('listitem').filter({ hasText: 'Newbie' });
  await expect(newbieRow.getByTestId('guild-roster-handoff')).toHaveCount(0);
});

test('no reports this week shows the honest empty state', async ({ page }) => {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) => route.fulfill(envelope(GUILD_PAGE)));
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME)));
  await page.route('**/v1/guilds/501/home', (route) =>
    route.fulfill(envelope({ ...HOME, reports: { rows: [] } })),
  );
  await page.goto('/guild/us/hardcore/the-last-watch');
  await expect(page.getByTestId('guild-home-empty-reports')).toHaveText('No reports this week yet.');
});

test('a signed-out visitor sees only the public page, with no signed-in section and no nag', async ({
  page,
}) => {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) => route.fulfill(envelope(GUILD_PAGE)));
  await page.route('**/v1/me', (route) =>
    route.fulfill({
      status: 401,
      contentType: 'application/json',
      body: '{"ok":false,"data":null,"error":null,"request_id":"r"}',
    }),
  );
  await page.goto('/guild/us/hardcore/the-last-watch');
  await expect(page.getByTestId('guild')).toBeVisible();
  await expect(page.getByTestId('guild-home')).toHaveCount(0);
});

// A guild's real identity is (region, ruleset, name), not (region, ruleset) alone -- many
// guilds share a region and ruleset. Guild B below shares GUILD_PAGE's region and ruleset
// but is a different guild entirely (different id, name, slug).
const GUILD_B_PAGE = {
  guild: { id: 777, name: 'Iron Vanguard', region: 'us', ruleset: 'hardcore' },
  progression: [],
  roster_best: [],
  reports: [],
};

// Signed in, but a verified member of Guild A (id 501) only -- no membership in Guild B.
const ME_GUILD_A_ONLY = {
  ...ME,
  guilds: [{ id: 501, region: 'us', ruleset: 'hardcore', name: 'The Last Watch', rank: 'officer' }],
};

test('a member of one guild visiting a different guild that shares its region and ruleset sees only that guild’s public page, never their own guild’s home panel', async ({
  page,
}) => {
  await page.route('**/v1/guilds/us/hardcore/iron-vanguard', (route) =>
    route.fulfill(envelope(GUILD_B_PAGE)),
  );
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME_GUILD_A_ONLY)));
  // Stubbed so a regression (matching on region+ruleset alone, which resolves to the
  // viewer's OWN guild id 501) is caught by a visible guild-home section instead of
  // silently failing the fetch.
  await page.route('**/v1/guilds/501/home', (route) => route.fulfill(envelope(HOME)));

  await page.goto('/guild/us/hardcore/iron-vanguard');
  await expect(page.getByTestId('guild')).toBeVisible();
  await expect(page.getByRole('heading', { name: 'Iron Vanguard' })).toBeVisible();
  await expect(page.getByTestId('guild-home')).toHaveCount(0);
});

// A solo officer with two verified characters in the guild: roster.length === 2, but both
// rows share one account (user_id), so the empty-roster "just me" message must still show.
const HOME_SOLO_OFFICER = {
  guild: { id: 501, name: 'The Last Watch', region: 'us', ruleset: 'hardcore' },
  viewer: { rank: 'officer', verified: true, can_manage: true },
  reports: { rows: [] },
  roster: [
    {
      character_key: 'us/hardcore/simfury',
      region: 'us',
      ruleset: 'hardcore',
      name: 'Simfury',
      class: 'Warrior',
      spec: 'Fury',
      rank: 'officer',
      verified: true,
      logged_recently: true,
      consent: 'gear',
      user_id: 7,
    },
    {
      character_key: 'us/hardcore/simfuryalt',
      region: 'us',
      ruleset: 'hardcore',
      name: 'Simfuryalt',
      class: 'Priest',
      spec: 'Holy',
      rank: 'officer',
      verified: true,
      logged_recently: false,
      consent: 'gear',
      user_id: 7,
    },
  ],
};

test('a solo officer with two characters in the guild sees the empty-roster message, not the populated roster', async ({
  page,
}) => {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) => route.fulfill(envelope(GUILD_PAGE)));
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME)));
  await page.route('**/v1/guilds/501/home', (route) => route.fulfill(envelope(HOME_SOLO_OFFICER)));

  await page.goto('/guild/us/hardcore/the-last-watch');
  await expect(page.getByTestId('guild-home-empty-roster')).toHaveText(
    "You're the only member the site knows about. Share the invite link to bring the rest of the guild in.",
  );
  await expect(page.getByTestId('guild-home-roster')).toHaveCount(0);
});
