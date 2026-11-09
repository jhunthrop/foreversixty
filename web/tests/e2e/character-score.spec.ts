// web/tests/e2e/character-score.spec.ts
// The character page's Forever Score Section (design/specs/2026-11-06-character.md C5, C6, C7,
// C10, C11): the score, tier and rating as Figures, the one "Most to gain" line and the parts in
// place; a character with no score reads "No logs yet", never a zero; a 404 is the not-found page.
// Also the best fight per encounter, the builds seen and the head's spec. The API is stubbed with
// page.route from the shared fixture; the rating read is answered 404 unless a test needs it.
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import labels from '../../../docs/contracts/fixtures/score-labels.json' with { type: 'json' };
import type { meAddonFixture } from '../../src/fixtures/me-addon';
import {
  CHARACTER,
  CHARACTER_PATH,
  CHARACTER_URL,
  OWNER_SCORE,
  ownerMeAnswer,
  routeCharacterSearch,
  routeOwner,
} from './support/character-fixture';
import { headerWired } from './support/header';

const NOT_FOUND = {
  status: 404,
  contentType: 'application/json',
  body: JSON.stringify({ ok: false, data: null, error: 'not_found', request_id: 'r' }),
};

const ok = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
});

async function stub(page: Page, data: unknown | 'missing', rating: unknown = null): Promise<void> {
  await page.route(CHARACTER_URL, (route) => route.fulfill(data === 'missing' ? NOT_FOUND : ok(data)));
  await page.route(`${CHARACTER_URL}/rating`, (route) =>
    route.fulfill(rating === null ? NOT_FOUND : ok(rating)),
  );
  if (data === 'missing') await routeCharacterSearch(page);
}

const UNSCORED = { score: null, tier: null, as_of: null, components: null, experience: null };

const RATING = {
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
};

test('a character the API has never heard of is the not-found page, not an error and not a zero', async ({
  page,
}) => {
  await stub(page, 'missing');
  await page.goto(CHARACTER_PATH);

  await expect(page.getByTestId('character-not-found')).toBeVisible();
  await expect(page.getByTestId('character-not-found-line')).toHaveText(
    'No character named Elyra Duskvale on US Hardcore.',
  );
  await expect(page.getByTestId('character-score')).toHaveCount(0);
  await expect(page.getByTestId('character-error')).toHaveCount(0);
});

test('a 404 is the not-found page for the owner too, without the sign-in row', async ({ page }) => {
  await routeOwner(page);
  await stub(page, 'missing');
  await page.goto(CHARACTER_PATH);

  await expect(page.getByTestId('character-not-found-line')).toBeVisible();
  await expect(page.getByTestId('character-add')).not.toContainText('Sign in to see your characters');
  await expect(page.getByTestId('character-error')).toHaveCount(0);
});

test('a character with ranked fights and no score yet says so and draws no Figure', async ({ page }) => {
  await stub(page, { ...CHARACTER.data, ...UNSCORED });
  await page.goto(CHARACTER_PATH);

  await expect(page.getByTestId('character')).toBeVisible();
  await expect(page.getByTestId('character-no-logs')).toHaveText(
    'No Forever Score yet. It is worked out from the rated fights below.',
  );
  await expect(page.getByTestId('character-score-figure')).toHaveCount(0);
  await expect(page.getByTestId('character-score')).not.toContainText(/\b0\b/);
});

test('a score its owner hid is no score block at all: no figure, no "no logs", no label', async ({
  page,
}) => {
  await stub(page, { ...CHARACTER.data, ...UNSCORED, score_history: [], score_hidden: true });
  await page.goto(CHARACTER_PATH);

  await expect(page.getByTestId('character')).toBeVisible();
  await expect(page.getByTestId('character-score')).toHaveCount(0);
  await expect(page.getByTestId('character-no-logs')).toHaveCount(0);
  await expect(page.getByTestId('character-score-figure')).toHaveCount(0);
  await expect(page.getByTestId('character')).not.toContainText(/hidden/i);
});

test('the owner with no logs is told how to get a rating, and fights stay hidden (C11)', async ({ page }) => {
  await routeOwner(page);
  await stub(page, { ...CHARACTER.data, ...UNSCORED, best: [], history: [] });
  await routeCharacterSearch(page);
  await page.goto(CHARACTER_PATH);

  await expect(page.getByTestId('character-no-logs')).toHaveText(
    'No logs yet. Upload a log to get a rating.',
  );
  await expect(page.getByTestId('character-best')).toHaveCount(0);
});

test('a bare character a visitor does not hold is the not-found page', async ({ page }) => {
  await stub(page, { ...CHARACTER.data, ...UNSCORED, best: [], history: [] });
  await routeCharacterSearch(page);
  await page.goto(CHARACTER_PATH);

  await expect(page.getByTestId('character-not-found')).toBeVisible();
  await expect(page.getByTestId('character')).toHaveCount(0);
});

test('a scored character: 1,830, Epic in the epic colour, the rating and how it is scored', async ({
  page,
}) => {
  await stub(page, { ...CHARACTER.data, ...OWNER_SCORE }, RATING);
  await page.goto(CHARACTER_PATH);

  const score = page.getByTestId('character-score');
  await expect(score.getByRole('heading', { name: 'Forever Score' })).toBeVisible();
  await expect(score).toContainText('6 rated fights');
  await expect(page.getByTestId('character-score-figure')).toContainText('1,830');
  const tier = page.getByTestId('character-tier');
  await expect(tier).toContainText('Epic');
  // --color-rarity-epic-text (#b866f5), the epic tone made legible as text.
  await expect(tier.getByText('Epic', { exact: true })).toHaveCSS('color', 'rgb(184, 102, 245)');
  await expect(page.getByTestId('character-rating')).toHaveText('65 of 100');
  await expect(score.getByRole('link', { name: 'How it is scored ›' })).toHaveAttribute('href', '/ratings');
});

test('every part shows its bar on one scale and its points, no "of N"', async ({ page }) => {
  await stub(page, { ...CHARACTER.data, ...OWNER_SCORE });
  await page.goto(CHARACTER_PATH);

  const parts = page.getByTestId('character-parts');
  for (const component of labels.components) await expect(parts).toContainText(component.label);
  await expect(page.getByTestId('character-part-gear')).toContainText('412');
  await expect(parts).not.toContainText(' of ');
});

test('Best per encounter: one Row per boss with its parse and date, then the link to all fights', async ({
  page,
}) => {
  await stub(page, { ...CHARACTER.data, ...OWNER_SCORE });
  await page.goto(CHARACTER_PATH);

  const best = page.getByTestId('character-best');
  await expect(best.getByRole('heading', { name: 'Best per encounter' })).toBeVisible();
  await expect(best).toContainText('Warden Kelthas');
  await expect(best).toContainText('9 Dec · Discipline');
  await expect(best.getByLabel('96 percentile')).toHaveText('96');
  await expect(best).toContainText('Deep Warden');
  await expect(best.getByRole('link', { name: 'All ranked fights (2) ›' })).toHaveAttribute('href', '/logs');
  await expect(page.getByRole('heading', { name: 'Every ranked fight' })).toHaveCount(0);
  await expect(page.getByTestId('character-builds')).toContainText('first seen 9 Dec 2026');
});

test('no best fights: the Section is left out', async ({ page }) => {
  await stub(page, { ...CHARACTER.data, ...OWNER_SCORE, best: [], history: [] }, RATING);
  await page.goto(CHARACTER_PATH);

  await expect(page.getByTestId('character-score')).toBeVisible();
  await expect(page.getByTestId('character-best')).toHaveCount(0);
  await expect(page.getByRole('heading', { name: 'Best per encounter' })).toHaveCount(0);
});

test('"/" opens the header lookup on a page without its own box, and it routes to the profile', async ({
  page,
}) => {
  await stub(page, CHARACTER.data);
  await page.goto('/ratings');
  await headerWired(page);
  await page.locator('body').click({ position: { x: 5, y: 300 } });
  await page.keyboard.press('/');
  const viewport = page.viewportSize();
  const input =
    viewport !== null && viewport.width >= 1024
      ? page.getByTestId('site-search-input')
      : page.getByTestId('site-menu-search-input');
  await expect(input).toBeFocused();
  await input.fill('Elyra Duskvale-Hardcore');
  await input.press('Enter');
  await expect(page).toHaveURL(new RegExp(`${CHARACTER_PATH}$`));
});

test('the head names the spec of the newest fight: "Discipline Priest"', async ({ page }) => {
  await stub(page, CHARACTER.data);
  await page.goto(CHARACTER_PATH);

  await expect(page.getByTestId('character-fact-spec')).toHaveText('Discipline Priest');
});

test('an owner whose sync time cannot be read gets "Addon", never a dangling dot', async ({ page }) => {
  const me = ownerMeAnswer().data as typeof meAddonFixture;
  const own = me.characters[0];
  if (own.build === undefined) throw new Error('the me fixture character has no build');
  const unreadable = { ...me, characters: [{ ...own, build: { ...own.build, captured_at: 'not a time' } }] };
  await page.route('**/v1/me', (route) => route.fulfill(ok(unreadable)));
  await stub(page, { ...CHARACTER.data, ...OWNER_SCORE });
  await page.goto(CHARACTER_PATH);

  const sync = page.getByTestId('character-sync');
  await expect(sync).toHaveText(/^Addon(\s·\s.*)?$/);
  await expect(sync).not.toHaveText(/^Addon · $/);
  await expect(page.getByTestId('character')).not.toContainText('Last seen in game');
});
