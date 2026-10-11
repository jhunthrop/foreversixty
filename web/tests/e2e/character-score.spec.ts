// web/tests/e2e/character-score.spec.ts
// Task L8.2: the character page's Forever Score. A character the site has never logged --
// the API's 404, or a 200 whose score is null -- shows the addon's no-logs line where the
// number would be, never a zero; a scored one leads the panel with the badge. The API is
// stubbed with page.route, the same fixture pattern character-rating.spec.ts uses, and the
// panel's own execution-rating read is answered 404 so only the score block is in play.
import { expect, test, type Page } from '@playwright/test';
import labels from '../../../docs/contracts/fixtures/score-labels.json' with { type: 'json' };
import { CHARACTER, CHARACTER_PATH, CHARACTER_URL } from './support/character-fixture';

const NOT_FOUND = {
  status: 404,
  contentType: 'application/json',
  body: JSON.stringify({ ok: false, data: null, error: 'not_found', request_id: 'r' }),
};

async function stub(page: Page, character: { status: number; body: unknown } | 'missing'): Promise<void> {
  await page.route(CHARACTER_URL, (route) =>
    character === 'missing'
      ? route.fulfill(NOT_FOUND)
      : route.fulfill({
          status: character.status,
          contentType: 'application/json',
          body: JSON.stringify(character.body),
        }),
  );
  await page.route(`${CHARACTER_URL}/rating`, (route) => route.fulfill(NOT_FOUND));
}

const withScore = (fields: Record<string, unknown>) => ({
  ...CHARACTER,
  data: { ...CHARACTER.data, ...fields },
});

test('a character the API has never heard of shows "no logs yet", not an error and not a zero', async ({
  page,
}) => {
  await stub(page, 'missing');
  await page.goto(CHARACTER_PATH);

  const none = page.getByTestId('forever-score-none');
  await expect(none).toHaveText(labels.noLogs);
  await expect(page.getByTestId('character-unlogged')).toContainText('Elyra Duskvale');
  await expect(page.getByTestId('character-error')).toHaveCount(0);
  await expect(page.getByTestId('score-badge')).toHaveCount(0);
  await expect(page.getByTestId('forever-score')).not.toContainText(/\b0\b/);
  await expect(page.getByTestId('character-in-game')).toBeVisible();
});

test('a logged character whose score is null reads the same', async ({ page }) => {
  await stub(page, {
    status: 200,
    body: withScore({ score: null, tier: null, as_of: null, components: null, experience: null }),
  });
  await page.goto(CHARACTER_PATH);

  await expect(page.getByTestId('character')).toBeVisible();
  await expect(page.getByTestId('forever-score-none')).toHaveText(labels.noLogs);
  await expect(page.getByTestId('score-badge')).toHaveCount(0);
});

test('a scored character leads with 1830 in the epic colour, its parts and its date', async ({ page }) => {
  await stub(page, {
    status: 200,
    body: withScore({
      score: 1830,
      tier: 'epic',
      as_of: '2026-11-04',
      components: { gear: 412, experience: 640, parses: 510, attendance: 180, discipline: 88 },
      experience: { dungeons: 4, raid_kills: 12, alts: 1 },
    }),
  });
  await page.goto(CHARACTER_PATH);

  const badge = page.getByTestId('forever-score').getByTestId('score-badge');
  await expect(badge).toHaveText('1830');
  await expect(badge).toHaveCSS('color', 'rgb(163, 53, 238)');
  await expect(page.getByTestId('forever-score')).toContainText('as of 4 Nov');
  for (const component of labels.components) {
    await expect(page.getByTestId(`forever-score-${component.key}`)).toContainText(component.label);
  }
  await expect(page.getByTestId('character-in-game').getByTestId('score-badge')).toHaveText('1830');
});
