// web/tests/e2e/character-rating.spec.ts
// The rating inside the character page's Forever Score Section (design/specs/2026-11-06-character.md
// C5; round 5, review 2026-11-09): "Rating 65 of 100", one compact fact under the score and the
// tier, no fight count; with no latest rated fight, or no rating at all (an anonymized character's
// 404), the row is left out rather than drawn empty. The page reads the rating only beside a
// score, so the character carries the fixture's OWNER_SCORE.
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { CHARACTER, CHARACTER_PATH, CHARACTER_URL, OWNER_SCORE } from './support/character-fixture';

const envelope = (data: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'r' }),
});

async function stub(page: Page, rating: { status: number; data: unknown }): Promise<void> {
  await page.route(CHARACTER_URL, (route) => route.fulfill(envelope({ ...CHARACTER.data, ...OWNER_SCORE })));
  await page.route(`${CHARACTER_URL}/rating`, (route) => route.fulfill(envelope(rating.data, rating.status)));
}

test('the rating is a fact, "Rating 65 of 100", and the panel counts the rated fights', async ({ page }) => {
  await stub(page, {
    status: 200,
    data: {
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
    },
  });
  await page.goto(CHARACTER_PATH);

  const rating = page.getByTestId('character-rating');
  await expect(page.getByTestId('character-score').locator('dt', { hasText: 'Rating' })).toBeVisible();
  await expect(rating).toHaveText('65 of 100');
  await expect(page.getByTestId('character-score')).not.toContainText('rated fights');
});

test('rated fights but no latest rating: no rating row', async ({ page }) => {
  await stub(page, {
    status: 200,
    data: {
      player_key: 'us/hardcore/elyra-duskvale',
      sample_size: 2,
      trend: [
        { fought_at: '2026-12-01T00:00:00Z', overall: 40, report_id: 'r1', fight_index: 1 },
        { fought_at: '2026-12-02T00:00:00Z', overall: 44, report_id: 'r1', fight_index: 2 },
      ],
      best_component: '',
      worst_component: '',
      latest: null,
    },
  });
  await page.goto(CHARACTER_PATH);

  await expect(page.getByTestId('character-score-figure')).toBeVisible();
  await expect(page.getByTestId('character-rating')).toHaveCount(0);
});

test('an anonymized character’s rating is left out, not a placeholder', async ({ page }) => {
  await stub(page, { status: 404, data: null });
  await page.goto(CHARACTER_PATH);

  await expect(page.getByTestId('character-score-figure')).toBeVisible();
  await expect(page.getByTestId('character-rating')).toHaveCount(0);
  await expect(page.getByTestId('character-score')).not.toContainText('rated fight');
});
