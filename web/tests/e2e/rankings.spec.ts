// web/tests/e2e/rankings.spec.ts
import { expect, test } from '@playwright/test';
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

  await page.goto('/rankings/warden-kelthas');

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
});

test('every filter lands in the URL and in the request', async ({ page }) => {
  const asked: string[] = [];
  await page.route('**/v1/rankings?**', (route) => {
    asked.push(route.request().url());
    return route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(ROWS) });
  });

  await page.goto('/rankings/warden-kelthas');
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

  await page.goto('/rankings/warden-kelthas');
  await page.getByTestId('board-guild').click();

  await expect(page.getByTestId('guild-rows')).toContainText('The Last Watch');
  expect(guildUrl).toContain('kind=progress');
});

test('a failed rankings call says so instead of showing an empty board', async ({ page }) => {
  await page.route('**/v1/rankings?**', (route) => route.abort());
  await page.goto('/rankings/warden-kelthas');
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

  await page.goto('/rankings/warden-kelthas');
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
