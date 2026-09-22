// web/tests/e2e/rankings.spec.ts
import { expect, test } from '@playwright/test';
import { gotoHydrated } from './support/hydrated';
import { heldRoute } from './support/held-route';

const ROWS = {
  ok: true,
  data: {
    rows: [
      {
        rank: 1,
        player: {
          key: 'us/hardcore/elyra-duskvale',
          name: 'Elyra Duskvale',
          class: 'Priest',
          spec: 'Shadow',
        },
        guild: { name: 'The Last Watch', ruleset: 'hardcore', region: 'us' },
        value: 1840,
        size: 40,
        fought_at: '2026-12-09T22:10:00Z',
        duration_ms: 240000,
        talent_split: '0/31/20',
        build_id: 'k7x2qm4a',
        trinkets: [],
        buff_count: 11,
        report_id: 'fixture2abcd',
        fight_index: 3,
        state: 'ok',
      },
      {
        rank: 2,
        player: { key: 'eu/normal/other-raider', name: 'Other Raider', class: 'Mage', spec: 'Fire' },
        value: 1600,
        size: 40,
        fought_at: '2026-12-09T21:00:00Z',
        duration_ms: 250000,
        talent_split: '0/20/31',
        trinkets: [],
        buff_count: 8,
        report_id: 'otherreport1',
        fight_index: 2,
        state: 'removed',
      },
    ],
    total: 2,
    page: 1,
    per_page: 100,
    updated_at: '2026-12-09T22:15:00Z',
  },
  error: null,
  request_id: 'r',
};

/** Same shape, a different top row, so a test can tell which answer painted the screen. */
const HPS_ROWS = {
  ...ROWS,
  data: {
    ...ROWS.data,
    rows: [{ ...ROWS.data.rows[0], player: { ...ROWS.data.rows[0].player, name: 'Brightwell Solace' } }],
    total: 1,
  },
};

test('the board lists ranked kills with guilds, splits, reports and moderation states', async ({ page }) => {
  await page.route('**/v1/rankings?**', (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(ROWS) }),
  );

  await gotoHydrated(page, '/rankings/warden-kelthas', 'filter-metric');

  await expect(page.locator('h1')).toHaveText('Warden Kelthas');
  await expect(page.getByTestId('rankings-count')).toContainText('2 ranked kills');
  await expect(page.getByTestId('ranking-1')).toContainText('Elyra Duskvale');
  await expect(page.getByTestId('ranking-1').getByTestId('ranking-build')).toHaveAttribute(
    'href',
    '/b/k7x2qm4a',
  );
  await expect(page.getByTestId('ranking-1').getByTestId('ranking-report')).toHaveAttribute(
    'href',
    '/reports/fixture2abcd?fight=3',
  );
  await expect(page.getByRole('link', { name: 'The Last Watch' })).toHaveAttribute(
    'href',
    '/guild/us/hardcore/the-last-watch',
  );
  await expect(page.getByTestId('ranking-2').getByTestId('ranking-state')).toHaveText('removed');

  // Row 2 has no guild (`player.key: 'eu/normal/other-raider'`), so a fallback that
  // defaulted region/ruleset to `us`/`normal` would link it there instead -- the wrong
  // character page, on a real player, with a link that looks correct. The href must carry
  // the row's own region and ruleset, read out of `player.key`.
  await expect(page.getByTestId('ranking-2').getByTestId('ranking-character')).toHaveAttribute(
    'href',
    '/character/eu/normal/other-raider',
  );
});

test('every filter lands in the URL and in the request', async ({ page }) => {
  const asked: string[] = [];
  await page.route('**/v1/rankings?**', (route) => {
    asked.push(route.request().url());
    return route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(ROWS) });
  });

  await gotoHydrated(page, '/rankings/warden-kelthas', 'filter-metric');
  await page.getByTestId('filter-metric').selectOption('hps');
  await page.getByTestId('filter-ruleset').selectOption('hardcore');

  await expect(page).toHaveURL(/\?metric=hps&ruleset=hardcore$/);
  expect(asked[asked.length - 1]).toContain('metric=hps');
  expect(asked[asked.length - 1]).toContain('ruleset=hardcore');
});

test('the guild board asks the guild endpoint', async ({ page }) => {
  await page.route('**/v1/rankings?**', (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(ROWS) }),
  );
  let guildUrl = '';
  await page.route('**/v1/rankings/guilds?**', (route) => {
    guildUrl = route.request().url();
    return route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({
        ok: true,
        data: {
          rows: [
            {
              rank: 1,
              guild: { name: 'The Last Watch', ruleset: 'hardcore', region: 'us' },
              value: 9,
              fought_at: '2026-12-09T22:10:00Z',
              report_id: 'fixture2abcd',
            },
          ],
        },
        error: null,
        request_id: 'r',
      }),
    });
  });

  await gotoHydrated(page, '/rankings/warden-kelthas', 'filter-metric');
  await page.getByTestId('board-guild').click();

  await expect(page.getByTestId('guild-rows')).toContainText('The Last Watch');
  expect(guildUrl).toContain('kind=progress');
});

test('a failed rankings call says so instead of showing an empty board', async ({ page }) => {
  await page.route('**/v1/rankings?**', (route) => route.abort());
  await gotoHydrated(page, '/rankings/warden-kelthas', 'filter-metric');
  await expect(page.getByTestId('rankings-error')).toBeVisible();
});

// Every filter lives in the URL (src/components/Rankings.svelte's `state`), and changing
// one fires a brand new request without waiting for the one before it to finish -- so the
// default board's own request and the switch to Healing can both be in flight at once and
// settle in either order. Whichever settles last must not speak for a filter combination
// nobody has selected any more: this holds the first (Damage) request open until the
// second (Healing) has already answered and painted the screen, then releases the first
// and checks its answer was dropped rather than overwriting the newer one.
test('a stale filtered response is discarded once a newer filter has already answered', async ({ page }) => {
  const slow = await heldRoute(page, '**/v1/rankings?**metric=dps**', (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(ROWS) }),
  );
  await page.route('**/v1/rankings?**metric=hps**', (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(HPS_ROWS) }),
  );

  await gotoHydrated(page, '/rankings/warden-kelthas', 'filter-metric');
  await slow.started;

  await page.getByTestId('filter-metric').selectOption('hps');
  await expect(page.getByTestId('ranking-1')).toContainText('Brightwell Solace');
  await expect(page.getByTestId('rankings-count')).toContainText('1 ranked kills');

  const staleAnswered = page.waitForResponse((response) => response.url().includes('metric=dps'));
  slow.release();
  await staleAnswered;
  await page.waitForTimeout(250);

  // Still Healing's one row, not Damage's two: the late Damage answer was dropped.
  await expect(page.getByTestId('ranking-1')).toContainText('Brightwell Solace');
  await expect(page.getByTestId('rankings-count')).toContainText('1 ranked kills');
  await expect(page.getByTestId('ranking-rows').locator('li')).toHaveCount(1);
});

const CHARACTER = {
  ok: true,
  data: {
    character: { name: 'Elyra Duskvale', region: 'us', ruleset: 'hardcore', class: 'Priest' },
    best: [
      {
        encounter: 'Warden Kelthas',
        encounter_id: 9001,
        difficulty: 8,
        metric: 'hps',
        value: 1840,
        percentile: 96.2,
        spec: 'Discipline',
        fought_at: '2026-12-09T22:10:00Z',
        report_id: 'fixture2abcd',
        fight_index: 3,
      },
      // Punctuation at the end of the name: the slug it maps to is the one thing two
      // independent derivations of "encounter name to /rankings/<slug>" disagree about.
      {
        encounter: 'Emeriss (Dream)',
        encounter_id: 9003,
        difficulty: 8,
        metric: 'hps',
        value: 1500,
        percentile: 80.1,
        spec: 'Discipline',
        fought_at: '2026-12-09T23:10:00Z',
        report_id: 'fixture2abcd',
        fight_index: 4,
      },
    ],
    history: [
      {
        encounter: 'Warden Kelthas',
        encounter_id: 9001,
        difficulty: 8,
        metric: 'hps',
        value: 1840,
        percentile: 96.2,
        spec: 'Discipline',
        fought_at: '2026-12-09T22:10:00Z',
        report_id: 'fixture2abcd',
        fight_index: 3,
      },
    ],
    builds_seen: [{ talent_split: '31/20/0', spec: 'Discipline', first_seen: '2026-12-09T22:10:00Z' }],
  },
  error: null,
  request_id: 'r',
};

const GUILD = {
  ok: true,
  data: {
    guild: { name: 'The Last Watch', region: 'us', ruleset: 'hardcore' },
    progression: [
      {
        encounter: 'Warden Kelthas',
        encounter_id: 9001,
        difficulty: 8,
        kills: 2,
        pull_count: 14,
        first_kill_at: '2026-12-09T22:10:00Z',
      },
      { encounter: 'Emeriss (Dream)', encounter_id: 9002, difficulty: 8, kills: 0, pull_count: 31 },
    ],
    roster_best: [
      {
        player: {
          key: 'us/hardcore/elyra-duskvale',
          name: 'Elyra Duskvale',
          class: 'Priest',
          spec: 'Discipline',
        },
        encounter: 'Warden Kelthas',
        encounter_id: 9001,
        metric: 'hps',
        value: 1840,
        fought_at: '2026-12-09T22:10:00Z',
      },
    ],
    reports: [
      {
        id: 'fixture2abcd',
        title: 'Sanguine Depths, fixture night',
        zone: 'Sanguine Depths',
        created_at: '2026-09-26T20:09:00Z',
      },
    ],
  },
  error: null,
  request_id: 'r',
};

test('a character page shows bests, history and the builds they were seen in', async ({ page }) => {
  await page.route('**/v1/characters/us/hardcore/elyra-duskvale', (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(CHARACTER) }),
  );

  await page.goto('/character/us/hardcore/elyra-duskvale');

  await expect(page.locator('h1')).toHaveText('Elyra Duskvale');
  await expect(page.getByTestId('character')).toContainText('Hardcore US');
  await expect(page.getByTestId('character-best')).toContainText('Warden Kelthas');
  await expect(page.getByTestId('character-history')).toContainText('Discipline');
  // A split as text, not a planner link: the log does not record the order points were spent in.
  await expect(page.getByTestId('character-builds')).toContainText('31/20/0');
  await expect(page.getByTestId('character-builds').getByRole('link')).toHaveCount(0);

  // rankings/api.ts's encounterSlug is the one derivation of this identifier, trailing
  // hyphen trimmed and all. A second, inlined copy here linked "Emeriss (Dream)" to
  // /rankings/emeriss-dream- , which the Worker accepts and renders as a permanently empty
  // board, while every other page on the site links to /rankings/emeriss-dream.
  await expect(
    page.getByTestId('character-best').getByRole('link', { name: 'Emeriss (Dream)' }),
  ).toHaveAttribute('href', '/rankings/emeriss-dream');
});

test('a guild page leads with progression, pull counts and kill dates', async ({ page }) => {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(GUILD) }),
  );

  await page.goto('/guild/us/hardcore/the-last-watch');

  await expect(page.locator('h1')).toHaveText('The Last Watch');
  await expect(page.getByTestId('guild')).toContainText('1 bosses down · 45 pulls');
  await expect(page.getByTestId('guild-progression')).toContainText('31 pulls');
  await expect(page.getByTestId('guild-kill').nth(1)).toHaveText('not killed');
  await expect(
    page.getByTestId('guild-roster').getByRole('link', { name: 'Elyra Duskvale' }),
  ).toHaveAttribute('href', '/character/us/hardcore/elyra-duskvale');
  await expect(page.getByTestId('guild-reports')).toContainText('Sanguine Depths, fixture night');
  // The same one derivation, for the same reason as on the character page.
  await expect(
    page.getByTestId('guild-progression').getByRole('link', { name: 'Emeriss (Dream)' }),
  ).toHaveAttribute('href', '/rankings/emeriss-dream');
});
