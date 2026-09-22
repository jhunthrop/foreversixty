// web/tests/e2e/report-rating.spec.ts
// Rating tab coverage: the tab reaches the full card, the Summary dashboard's rating row
// links into it, a scoped one-player link opens directly onto that player, and the
// truthful empty state shows when the API has nothing yet -- API stubbed with page.route
// per docs/superpowers/plans/2026-09-21-rating-web.md (the API lane has not landed).
import { expect, test } from '@playwright/test';

const REPORT = '/reports/fixture2abcd';

const envelope = (data: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'r' }),
});

const PLAYER = {
  player_key: 'us/normal/elyra-duskvale',
  player_name: 'Elyra Duskvale',
  class: 'Priest',
  spec: 'Holy',
  role: 'healer',
  overall: 70,
  overall_uncapped: 70,
  overall_capped: false,
  basis: 'percentile',
  components: [
    {
      name: 'output',
      score: 71,
      weight: 30,
      basis: 'percentile',
      percentile: 71,
      bracket_n: 40,
      excluded: false,
      reason: '',
      moments: [],
    },
    {
      name: 'survival',
      score: 90,
      weight: 10,
      basis: 'percentile',
      percentile: 90,
      bracket_n: 40,
      excluded: false,
      reason: '',
      moments: [],
    },
    {
      name: 'mechanics',
      score: null,
      weight: 0,
      basis: '',
      percentile: null,
      bracket_n: 0,
      excluded: true,
      reason: 'no_mechanics_table',
      moments: [],
    },
    {
      name: 'utility',
      score: 62,
      weight: 20,
      basis: 'percentile',
      percentile: 62,
      bracket_n: 40,
      excluded: false,
      reason: '',
      moments: [],
    },
    {
      name: 'preparation',
      score: 100,
      weight: 10,
      basis: 'absolute',
      percentile: null,
      bracket_n: 0,
      excluded: false,
      reason: '',
      moments: [],
    },
    {
      name: 'activity',
      score: 66,
      weight: 10,
      basis: 'percentile',
      percentile: 66,
      bracket_n: 40,
      excluded: false,
      reason: '',
      moments: [],
    },
  ],
};

test('the Summary dashboard shows a rating row that opens the Rating tab', async ({ page }) => {
  await page.route('**/v1/reports/fixture2abcd/fights/3/ratings', (route) =>
    route.fulfill(
      envelope({
        fight_index: 3,
        kill: true,
        kill_time_band: 'typical',
        model_version: 'v1',
        players: [PLAYER],
      }),
    ),
  );
  await page.goto(`${REPORT}?fight=3`);

  await expect(page.getByTestId('rating-panel')).toBeVisible();
  await expect(page.getByTestId('rating-panel-overall')).toHaveText('70');
  await page.getByRole('button', { name: 'Rating tab' }).click();

  await expect(page).toHaveURL(/tab=rating/);
  await expect(page.getByTestId('rating-tab')).toBeVisible();
  await expect(page.getByTestId('rating-overall')).toHaveText('70');
});

test('an excluded component states why, and a moment link jumps to its tab', async ({ page }) => {
  await page.route('**/v1/reports/fixture2abcd/fights/3/ratings', (route) =>
    route.fulfill(
      envelope({
        fight_index: 3,
        kill: true,
        kill_time_band: 'typical',
        model_version: 'v1',
        players: [PLAYER],
      }),
    ),
  );
  await page.goto(`${REPORT}?fight=3&tab=rating`);

  await expect(page.getByTestId('rating-component-mechanics')).toContainText(
    'has no curated mechanics table yet',
  );
});

test('scoping the url to one player shows only that player’s card', async ({ page }) => {
  await page.route('**/v1/reports/fixture2abcd/fights/3/ratings', (route) =>
    route.fulfill(
      envelope({
        fight_index: 3,
        kill: true,
        kill_time_band: 'typical',
        model_version: 'v1',
        players: [PLAYER],
      }),
    ),
  );
  await page.goto(`${REPORT}?fight=3`);
  const eventsRow = page.locator('[data-testid="source-scope"] option', { hasText: 'Elyra Duskvale' });
  const guid = await eventsRow.getAttribute('value');
  await page.goto(`${REPORT}?fight=3&tab=rating&source=${guid}`);

  await expect(page.getByTestId('rating-tab')).toBeVisible();
  await expect(page.locator('[data-testid^="rating-card-"]')).toHaveCount(1);
});

test('no ratings yet for this fight is a truthful empty state, not a spinner or an error', async ({
  page,
}) => {
  await page.route('**/v1/reports/fixture2abcd/fights/3/ratings', (route) =>
    route.fulfill(
      envelope({ fight_index: 3, kill: true, kill_time_band: 'typical', model_version: 'v1', players: [] }),
    ),
  );
  await page.goto(`${REPORT}?fight=3&tab=rating`);

  await expect(page.getByTestId('rating-tab-empty')).toBeVisible();
});

test('the whole night has no ratings view; it explains why rather than rendering nothing', async ({
  page,
}) => {
  await page.goto(`${REPORT}?fight=0&tab=rating`);
  await expect(page.getByTestId('rating-tab-night')).toBeVisible();
});

test('the Rating tab links to the explanation page', async ({ page }) => {
  await page.route('**/v1/reports/fixture2abcd/fights/3/ratings', (route) =>
    route.fulfill(
      envelope({
        fight_index: 3,
        kill: true,
        kill_time_band: 'typical',
        model_version: 'v1',
        players: [PLAYER],
      }),
    ),
  );
  await page.goto(`${REPORT}?fight=3&tab=rating`);
  await expect(page.getByRole('link', { name: 'How is this calculated?' })).toHaveAttribute(
    'href',
    '/ratings',
  );
});

// The first real production card: a wipe with no mechanics table and no validated spec,
// where only three of six parts were measurable. The engine gives such a card no overall,
// and the page must say so rather than print the renormalised remainder as a judgement.
test('a card with too little measured shows Not rated and the reason, never a number', async ({ page }) => {
  await page.route('**/v1/reports/fixture2abcd/fights/3/ratings', (route) =>
    route.fulfill(
      envelope({
        fight_index: 3,
        kill: false,
        kill_time_band: '',
        model_version: 'v1',
        players: [
          {
            ...PLAYER,
            overall: 0,
            overall_uncapped: 0,
            coverage: 0.3,
            insufficient: true,
            insufficient_reason: 'wipe; no mechanics table for this encounter',
          },
        ],
      }),
    ),
  );
  await page.goto(`${REPORT}?fight=3&tab=rating`);

  await expect(page.getByTestId('rating-tab')).toBeVisible();
  await expect(page.getByTestId('rating-insufficient')).toHaveText('Not rated');
  await expect(page.getByTestId('rating-overall')).toHaveCount(0);
  await expect(page.getByTestId('rating-insufficient-note')).toContainText('no mechanics table');
  // The parts that were measured are still listed.
  await expect(page.getByTestId('rating-components')).toBeVisible();
});
