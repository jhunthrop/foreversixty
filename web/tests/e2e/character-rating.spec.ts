// web/tests/e2e/character-rating.spec.ts
// The character page's rating panel: overall + six-segment bar + trend once there are
// enough fights, the too-few-samples copy below that, and the anonymized-character case
// rendering nothing. API stubbed with page.route (the API lane has not landed).
import { expect, test } from '@playwright/test';

const CHARACTER = {
  ok: true,
  data: {
    character: { name: 'Elyra Duskvale', region: 'us', ruleset: 'hardcore', class: 'Priest' },
    best: [],
    history: [],
    builds_seen: [],
  },
  error: null,
  request_id: 'r',
};

const envelope = (data: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'r' }),
});

test('the character page shows the rating panel with a trend once there are enough fights', async ({
  page,
}) => {
  await page.route('**/v1/characters/us/hardcore/elyra-duskvale', (route) =>
    route.fulfill(envelope(CHARACTER.data)),
  );
  await page.route('**/v1/characters/us/hardcore/elyra-duskvale/rating', (route) =>
    route.fulfill(
      envelope({
        player_key: 'us/hardcore/elyra-duskvale',
        sample_size: 6,
        trend: Array.from({ length: 6 }, (_, i) => ({
          fought_at: `2026-12-0${i + 1}T00:00:00Z`,
          overall: 40 + i * 5,
          report_id: 'fixture2abcd',
          fight_index: i + 1,
        })),
        best_component: 'preparation',
        worst_component: 'activity',
        latest: {
          player_key: 'us/hardcore/elyra-duskvale',
          player_name: 'Elyra Duskvale',
          class: 'Priest',
          spec: 'Holy',
          role: 'healer',
          overall: 65,
          overall_uncapped: 65,
          overall_capped: false,
          basis: 'percentile',
          components: [],
        },
      }),
    ),
  );

  await page.goto('/character/us/hardcore/elyra-duskvale');

  await expect(page.getByTestId('character-rating')).toBeVisible();
  await expect(page.getByTestId('character-rating-overall')).toHaveText('65');
  await expect(page.getByTestId('rating-trend')).toBeVisible();
});

test('fewer than five rated fights shows the exact threshold copy, no sparkline', async ({ page }) => {
  await page.route('**/v1/characters/us/hardcore/elyra-duskvale', (route) =>
    route.fulfill(envelope(CHARACTER.data)),
  );
  await page.route('**/v1/characters/us/hardcore/elyra-duskvale/rating', (route) =>
    route.fulfill(
      envelope({
        player_key: 'us/hardcore/elyra-duskvale',
        sample_size: 2,
        trend: [
          { fought_at: '2026-12-01T00:00:00Z', overall: 40, report_id: 'r1', fight_index: 1 },
          { fought_at: '2026-12-02T00:00:00Z', overall: 44, report_id: 'r1', fight_index: 2 },
        ],
        best_component: '',
        worst_component: '',
        latest: null,
      }),
    ),
  );

  await page.goto('/character/us/hardcore/elyra-duskvale');

  await expect(page.getByTestId('character-rating-too-few')).toHaveText(
    'Not enough rated fights yet to show a trend (2 of 5 needed).',
  );
  await expect(page.getByTestId('rating-trend')).toHaveCount(0);
});

test('an anonymized character’s rating panel renders nothing, not a placeholder', async ({ page }) => {
  await page.route('**/v1/characters/us/hardcore/elyra-duskvale', (route) =>
    route.fulfill(envelope(CHARACTER.data)),
  );
  await page.route('**/v1/characters/us/hardcore/elyra-duskvale/rating', (route) =>
    route.fulfill({
      status: 404,
      contentType: 'application/json',
      body: JSON.stringify({ ok: false, data: null, error: 'not_found', request_id: 'r' }),
    }),
  );

  await page.goto('/character/us/hardcore/elyra-duskvale');

  await expect(page.getByTestId('character-rating')).toHaveCount(0);
});
