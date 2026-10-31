// web/tests/e2e/character-score.spec.ts
// Task L8.2: the character page's Forever Score. A character the site has never logged --
// the API's 404, or a 200 whose score is null -- shows the addon's no-logs line where the
// number would be, never a zero; a scored one leads the panel with the badge. The API is
// stubbed with page.route, the same fixture pattern character-rating.spec.ts uses, and the
// panel's own execution-rating read is answered 404 so only the score block is in play.
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import labels from '../../../docs/contracts/fixtures/score-labels.json' with { type: 'json' };
import { meAddonFixture } from '../../src/fixtures/me-addon';
import { CHARACTER, CHARACTER_PATH, CHARACTER_URL, routeOwner } from './support/character-fixture';
import { headerWired } from './support/header';

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

test('a character the API has never heard of is not drawn as a character: not found, not an error and not a zero', async ({
  page,
}) => {
  await stub(page, 'missing');
  await page.goto(CHARACTER_PATH);

  await expect(page.getByTestId('character-not-found-line')).toHaveText(
    'No character called Elyra Duskvale on Hardcore US yet.',
  );
  await expect(page.getByTestId('forever-score')).toHaveCount(0);
  await expect(page.getByTestId('character-error')).toHaveCount(0);
});

test('the owner of a character with no logs yet sees "no logs yet", not "not found"', async ({ page }) => {
  await routeOwner(page);
  await stub(page, 'missing');
  await page.goto(CHARACTER_PATH);

  await expect(page.getByTestId('forever-score-empty')).toContainText('No logs yet');
  await expect(page.getByTestId('character-in-game-score')).toHaveText(labels.noLogs);
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
  await expect(page.getByTestId('forever-score-empty')).toContainText('No logs yet');
  await expect(page.getByTestId('character-in-game-score')).toHaveText(labels.noLogs);
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

// Fix round 1: the reviewers' capture, as the page now draws it.
const CAPTURE = {
  character: { name: 'Elyra Duskvale', region: 'us', ruleset: 'hardcore', class: 'Priest' },
  best: [],
  history: [],
  builds_seen: [],
  score: 1830,
  tier: 'epic',
  as_of: '2026-11-04',
  components: { gear: 412, experience: 620, parses: 540, attendance: 180, discipline: 78 },
  experience: { dungeons: 14, raid_kills: 0, alts: 2 },
};

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

const ok = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
});

test('the capture: captions on both numbers, empty sections and an empty strip are absent', async ({
  page,
}) => {
  await page.route(CHARACTER_URL, (route) => route.fulfill(ok(CAPTURE)));
  await page.route(`${CHARACTER_URL}/rating`, (route) => route.fulfill(ok(RATING)));
  await page.goto(CHARACTER_PATH);

  // A visitor's view (fix round 2 item 1): no "Your" on someone else's profile.
  await expect(page.getByTestId('forever-score-caption')).toHaveText('Whole record, as of 4 Nov.');
  await expect(page.getByTestId('character-rating-caption')).toHaveText('Latest logged fights.');
  await expect(page.getByTestId('character-rating-overall')).toHaveText('65');
  // Every rating segment is empty, so the strip is not drawn.
  await expect(page.getByTestId('character-rating-segments')).toHaveCount(0);
  await expect(page.getByRole('heading', { name: 'Best per encounter' })).toHaveCount(0);
  await expect(page.getByRole('heading', { name: 'Every ranked fight' })).toHaveCount(0);

  await expect(page.getByTestId('forever-score-most-to-gain')).toHaveText(
    'Most to gain: Gear, 88 below its top.',
  );
  await expect(page.getByTestId('forever-score-experience-detail')).toHaveText(
    '14 dungeons · 0 raid bosses · 2 alts',
  );
  await expect(page.getByTestId('tier-strip-epic')).toHaveAttribute('aria-current', 'true');
  // Fix round 2 item 2: the marked bar is the one Most to gain names, with its "+88 pts" (fix
  // round 4 item 2: the cue carries its unit).
  await expect(page.getByTestId('forever-score-gear')).toHaveAttribute('data-needs-work', 'true');
  await expect(page.getByTestId('forever-score-gain-cue')).toHaveText('+88 pts');
  // Fix round 2 item 6: the rating card.
  await expect(page.getByTestId('rating-card-weakest')).toHaveText('Weakest part of a fight: Activity.');
  // L8.5 ruling 1 and its fix round 1: the line carries the count the old 'Last 6 fights' label
  // did, says the 65 is the latest fight's, and makes no claim about the score's Parses.
  await expect(page.getByTestId('rating-card-verdict')).toHaveText(
    'Latest fight (6 Dec) rated 65 of 100 · 6 fights rated',
  );
  // Fix round 2 item 1: a visitor is not told to paste an export for someone else's character.
  await expect(page.getByTestId('character-needs-addon')).toHaveCount(0);
});

test('the fight tables: column headers, "9 Dec 2026" dates, and no Best table that repeats the list', async ({
  page,
}) => {
  await page.route(CHARACTER_URL, (route) => route.fulfill(ok(CHARACTER.data)));
  await page.route(`${CHARACTER_URL}/rating`, (route) => route.fulfill(NOT_FOUND));
  await page.goto(CHARACTER_PATH);

  // The shared fixture's best rows are its history rows, so "Best per encounter" is left out.
  await expect(page.getByTestId('character-history')).toBeVisible();
  await expect(page.getByTestId('character-best')).toHaveCount(0);
  const header = page.getByTestId('character-history').getByTestId('character-table-header');
  await expect(header).toContainText('Boss');
  await expect(header).toContainText('Parse %');
  await expect(header).toContainText('HPS');
  await expect(page.getByTestId('character-history')).toContainText('9 Dec 2026');
  await expect(page.getByTestId('character-builds')).toContainText('first seen 9 Dec 2026');
});

test('"/" opens the header lookup on a page without its own box, and it routes to the profile', async ({
  page,
}) => {
  await page.route(CHARACTER_URL, (route) => route.fulfill(ok(CHARACTER.data)));
  await page.route(`${CHARACTER_URL}/rating`, (route) => route.fulfill(NOT_FOUND));
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

test('the header names the spec of the newest fight: "Discipline Priest"', async ({ page }) => {
  await page.route(CHARACTER_URL, (route) => route.fulfill(ok(CHARACTER.data)));
  await page.route(`${CHARACTER_URL}/rating`, (route) => route.fulfill(NOT_FOUND));
  await page.goto(CHARACTER_PATH);

  await expect(page.getByTestId('character')).toContainText('Discipline Priest');
});

test('In game: the tooltip mock for any visitor, the planner line only for the owner', async ({ page }) => {
  await page.route(CHARACTER_URL, (route) => route.fulfill(ok(CAPTURE)));
  await page.route(`${CHARACTER_URL}/rating`, (route) => route.fulfill(NOT_FOUND));
  await page.goto(CHARACTER_PATH);

  const tooltip = page.getByTestId('character-in-game-tooltip');
  await expect(tooltip).toContainText('Elyra Duskvale');
  await expect(tooltip).toContainText('Forever Score 1830 · Epic · as of 4 Nov');
  await expect(page.getByTestId('character-in-game')).toContainText(
    'Shown on this character’s tooltip when you have the Forever Sixty addon.',
  );
  await expect(page.getByTestId('character-in-game-planner')).toHaveCount(0);
});

test('the owner’s no-logs page keeps the neutral crest and the same card, with what to do next', async ({
  page,
}) => {
  await routeOwner(page);
  await stub(page, 'missing');
  await page.goto(CHARACTER_PATH);

  await expect(page.getByTestId('character-avatar-fallback')).toBeVisible();
  await expect(page.getByTestId('character-avatar-silhouette')).toBeVisible();
  const card = page.getByTestId('forever-score-card');
  await expect(card).toContainText('No logs yet');
  await expect(card).not.toContainText(labels.noLogs);
  await expect(page.getByTestId('character-in-game-tooltip')).toContainText(labels.noLogs);
});

/** The me fixture's addon-synced character, made the owner of this page's character. */
function ownerMe(): typeof meAddonFixture {
  const own = meAddonFixture.characters[0];
  return {
    ...meAddonFixture,
    characters: [
      {
        ...own,
        key: 'us/hardcore/elyra-duskvale',
        region: 'us',
        ruleset: 'hardcore',
        name: 'Elyra Duskvale',
        class: 'Priest',
      },
    ],
  };
}

test('the owner’s header carries the sync line, never "Last seen in game" or "No export yet" (L8.4 fix round 2 item 5)', async ({
  page,
}) => {
  await page.route('**/v1/me', (route) => route.fulfill(ok(ownerMe())));
  await page.route(CHARACTER_URL, (route) => route.fulfill(ok(CAPTURE)));
  await page.route(`${CHARACTER_URL}/rating`, (route) => route.fulfill(NOT_FOUND));
  await page.goto(CHARACTER_PATH);

  await expect(page.getByTestId('character-sync')).toHaveText(/^Addon · \d+ minutes ago$/);
  await expect(page.getByTestId('character')).not.toContainText('Last seen in game');
  await expect(page.getByTestId('character-needs-addon')).toHaveCount(0);
  await expect(page.getByTestId('forever-score-caption')).toHaveText('Your whole record, as of 4 Nov.');
});

test('an owner whose sync time cannot be read gets "Addon", never a dangling dot or "No export yet"', async ({
  page,
}) => {
  const me = ownerMe();
  const own = me.characters[0];
  if (own.build === undefined) throw new Error('the me fixture character has no build');
  me.characters[0] = { ...own, build: { ...own.build, captured_at: 'not a time' } };
  await page.route('**/v1/me', (route) => route.fulfill(ok(me)));
  await page.route(CHARACTER_URL, (route) => route.fulfill(ok(CAPTURE)));
  await page.route(`${CHARACTER_URL}/rating`, (route) => route.fulfill(NOT_FOUND));
  await page.goto(CHARACTER_PATH);

  await expect(page.getByTestId('character-sync')).toHaveText(/^Addon(\s·\s.*)?$/);
  await expect(page.getByTestId('character-sync')).not.toHaveText(/^Addon · $/);
  await expect(page.getByTestId('character')).not.toContainText('Last seen in game');
  await expect(page.getByTestId('character-needs-addon')).toHaveCount(0);
});
