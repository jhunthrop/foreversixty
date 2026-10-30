// web/tests/e2e/rankings.spec.ts
// The Rankings page (design/specs/2026-10-11-rankings.md): signed in, the shared character
// header says where the character stands and the board marks its row in place (or pins it when it
// is off the page); a header switch re-aims the board without a reload; signed out, the shared
// band stands where the header would and the board has its Class and Spec choices; on a phone the
// filters collapse to one line. The older checks (rows, filters in the URL, guild board, empty,
// failed, stale responses) run signed out against a plain two-row board.
import { expect, test, type Page } from '@playwright/test';
import { gotoHydrated } from './support/hydrated';
import { heldRoute } from './support/held-route';
import { openFilters } from './support/rankings-filters';
import { FROST_KEY, OWN_KEY, stubRankingsApi, type BoardOptions } from './support/rankings-fixture';
import { OBNOXIOUS, armoryPointer, keyOf, signInWith, signedOut, storePointer } from './support/selector';
import type { FixtureCharacter } from './support/selector';

const DESKTOP = { width: 1440, height: 900 };
const PHONE = { width: 390, height: 844 };
const FROSTBYTE: FixtureCharacter = {
  name: 'Frostbyte',
  class: 'Mage',
  spec: 'Frost',
  level: 60,
  faction: 'alliance',
  buildDaysAgo: 1,
};

const selector = (page: Page) => page.getByTestId('character-selector');

async function openSignedIn(page: Page, options: BoardOptions = {}) {
  await signInWith(page, [OBNOXIOUS, FROSTBYTE]);
  await storePointer(page, armoryPointer(OBNOXIOUS));
  const api = await stubRankingsApi(page, options);
  await page.goto('/rankings');
  await expect(page.getByTestId('character-header-figure-1')).toBeVisible();
  await expect(page.getByTestId('ranking-1')).toBeVisible();
  return api;
}

test.describe('signed in, desktop', () => {
  test.use({ viewport: DESKTOP });
  test.skip(() => test.info().project.name !== 'desktop', 'desktop layout only');

  test('the band shows the rank, the board opens on the character, and its row is marked in place', async ({
    page,
  }) => {
    const api = await openSignedIn(page);

    await expect(page.locator('h1')).toHaveText('Rankings');
    await expect(page.getByTestId('rankings-job')).toHaveText(
      'Where your character stands among characters like it.',
    );
    await expect(page.getByTestId('character-header-answer-label')).toHaveText('Your rank');
    await expect(page.getByTestId('character-header-figure-1')).toHaveText('14th of 212 Fury Warriors');
    await expect(page.getByTestId('character-header-note')).toHaveText('On Onyxia, Normal ruleset');
    await expect(page.getByTestId('character-header-action')).toHaveText('See your parses →');
    await expect(page.getByTestId('character-header-action')).toHaveAttribute(
      'href',
      '/character/us/normal/obnoxious-yell',
    );

    // The board opened on the encounter the character most recently killed, in its own cohort.
    await expect(page).toHaveURL(/\/rankings\/onyxia$/);
    await expect(page.getByTestId('rankings-scope')).toHaveText('Fury Warriors');
    const board = api.asked.find(
      (url) => url.searchParams.get('character') === OWN_KEY && url.searchParams.get('page') === null,
    );
    expect(board?.searchParams.get('encounter')).toBe('onyxia');
    expect(board?.searchParams.get('class')).toBe('warrior');
    expect(board?.searchParams.get('spec')).toBe('Fury');
    expect(board?.searchParams.get('ruleset')).toBe('normal');
    expect(board?.searchParams.get('metric')).toBe('dps');
    expect(board?.searchParams.get('distinct')).toBe('character');

    // The own row is marked where it is, never pinned; the others are not marked.
    await expect(page.getByTestId('ranking-14')).toHaveAttribute('data-own', 'true');
    await expect(page.getByTestId('ranking-13')).not.toHaveAttribute('data-own', 'true');
    await expect(page.getByTestId('ranking-pinned')).toHaveCount(0);
    await expect(page.getByTestId('ranking-14')).toContainText('Obnoxious Yell');
    await expect(page.getByTestId('ranking-14').getByTestId('ranking-rank')).toHaveText('14');
  });

  test('the page is the spec: one stamp, no pills, no spine bar, a legend, plain Build and gold Report', async ({
    page,
  }) => {
    await openSignedIn(page);

    await expect(page.locator('#rankings .pill')).toHaveCount(0);
    await expect(page.getByTestId('current-character-bar')).toHaveCount(0);
    await expect(page.getByTestId('rankings-updated')).toHaveText('Updated Dec 13');
    await expect(page.getByText(/^Updated /)).toHaveCount(1);
    await expect(page.getByTestId('rankings-tier-link')).toHaveAttribute('href', '/tiers');
    await expect(page.getByTestId('rankings-legend')).toContainText('Rank colour is the percentile:');
    await expect(page.getByTestId('rankings-legend').locator('span')).toHaveText([
      '99+',
      '95+',
      '75+',
      '50+',
      '25+',
      'under 25',
    ]);
    await expect(page.getByTestId('rankings-heading')).not.toContainText(/\d+ ranked/);

    // A single-class board draws no crest in its rows (the header carries it once).
    await expect(page.getByTestId('ranking-rows').locator('img')).toHaveCount(0);

    const row = page.getByTestId('ranking-1');
    const colour = (testid: string) =>
      row.getByTestId(testid).evaluate((node) => getComputedStyle(node).color);
    expect(await colour('ranking-build')).not.toBe(await colour('ranking-report'));
    await expect(row.getByTestId('ranking-build')).toHaveAttribute('href', '/b/b1');
    await expect(row.getByTestId('ranking-report')).toHaveAttribute('href', '/reports/report1?fight=1');

    // No sign-in prose or button of the page's own.
    await expect(page.locator('#rankings').getByText(/sign in/i)).toHaveCount(0);
  });

  test('a header switch re-aims the board at the new character without a reload', async ({ page }) => {
    const api = await openSignedIn(page);
    await page.evaluate(() => {
      (window as unknown as { __kept: boolean }).__kept = true;
    });

    await selector(page).click();
    await page
      .getByTestId('selector-panel')
      .getByTestId(`selector-row-${keyOf(FROSTBYTE)}`)
      .getByRole('button')
      .first()
      .click();

    await expect(page.getByTestId('character-header-name')).toHaveText('Frostbyte');
    await expect(page.getByTestId('rankings-scope')).toHaveText('Frost Mages');
    await expect(page).toHaveURL(/\/rankings\/warden-kelthas$/);
    await expect(page.getByTestId('character-header-figure-1')).toHaveText('3rd of 3 Frost Mages');
    await expect(page.getByTestId('character-header-note')).toHaveText('On Warden Kelthas, Normal ruleset');
    await expect(page.getByTestId('ranking-3')).toHaveAttribute('data-own', 'true');
    // A mage board is one class too: still no crest.
    await expect(page.getByTestId('ranking-rows').locator('img')).toHaveCount(0);
    expect(await page.evaluate(() => (window as unknown as { __kept?: boolean }).__kept)).toBe(true);

    const mage = api.asked.filter((url) => url.searchParams.get('class') === 'mage');
    expect(mage.length).toBeGreaterThan(0);
    expect(mage.every((url) => url.searchParams.get('spec') === 'Frost')).toBe(true);
    expect(mage.some((url) => url.searchParams.get('character') === FROST_KEY)).toBe(true);
  });

  test('a best kill off the visible page is pinned on top at its true rank, once', async ({ page }) => {
    await openSignedIn(page, { ownOffPage: true });

    await expect(page.getByTestId('character-header-figure-1')).toHaveText('204th of 212 Fury Warriors');
    const pinned = page.getByTestId('ranking-pinned');
    await expect(pinned).toContainText('Your best kill');
    await expect(pinned.getByTestId('ranking-pinned-row')).toHaveAttribute('data-own', 'true');
    await expect(pinned.getByTestId('ranking-rank')).toHaveText('204');
    // The pin is the first row of the panel; the board follows from #1, unmarked.
    await expect(page.getByTestId('ranking-1')).toBeVisible();
    await expect(page.getByTestId('ranking-1')).not.toHaveAttribute('data-own', 'true');
    await expect(page.locator('[data-own="true"]')).toHaveCount(1);
    const pinnedBox = await pinned.boundingBox();
    const firstBox = await page.getByTestId('ranking-1').boundingBox();
    expect((pinnedBox?.y ?? Infinity) < (firstBox?.y ?? 0)).toBe(true);
  });

  test('"Another class or spec" opens the pickers, says whose board it is, and goes back', async ({
    page,
  }) => {
    await openSignedIn(page);
    await expect(page.getByTestId('filter-class')).toHaveCount(0);

    await page.getByTestId('rankings-another').click();
    await expect(page.getByTestId('filter-class')).toBeVisible();
    await expect(page.getByTestId('filter-spec')).toBeVisible();

    await page.getByTestId('filter-class').selectOption('mage');
    await expect(page.getByTestId('rankings-scope')).toHaveText('Mages');
    await expect(page.getByTestId('rankings-showing-note')).toContainText(
      'Showing Mages. Your rank is on the Fury Warrior board.',
    );
    // The header's answer stays about the character, and no row on this board is marked.
    await expect(page.getByTestId('character-header-figure-1')).toHaveText('14th of 212 Fury Warriors');
    await expect(page.locator('[data-own="true"]')).toHaveCount(0);
    await expect(page.getByTestId('ranking-pinned')).toHaveCount(0);

    await page.getByTestId('rankings-back').click();
    await expect(page.getByTestId('rankings-scope')).toHaveText('Fury Warriors');
    await expect(page.getByTestId('rankings-showing-note')).toHaveCount(0);
    await expect(page.getByTestId('ranking-14')).toHaveAttribute('data-own', 'true');
  });

  test('an API that does not dedupe yet gives the header no rank to show, and no pin', async ({ page }) => {
    await signInWith(page, [OBNOXIOUS]);
    await storePointer(page, armoryPointer(OBNOXIOUS));
    await stubRankingsApi(page, { legacy: true });
    await page.goto('/rankings');

    await expect(page.getByTestId('character-header-failed')).toContainText('Your rank did not load.');
    await expect(page.getByTestId('ranking-1')).toBeVisible();
    await expect(page.getByTestId('ranking-pinned')).toHaveCount(0);
  });
});

test.describe('signed out, desktop', () => {
  test.use({ viewport: DESKTOP });
  test.skip(() => test.info().project.name !== 'desktop', 'desktop layout only');

  test('the shared band stands where the header would, and the board has its Class and Spec choices', async ({
    page,
  }) => {
    await signedOut(page);
    const api = await stubRankingsApi(page);
    await page.goto('/rankings');
    await expect(page.getByTestId('ranking-1')).toBeVisible();

    await expect(page.getByTestId('character-header-signed-out')).toBeVisible();
    await expect(page.getByTestId('rankings-job')).toHaveText(
      'How characters rank on each encounter, kill by kill.',
    );
    await expect(page.getByTestId('rankings-scope')).toHaveText('Every class and spec');
    // No button of the page's own: the nav selector is the only sign-in control.
    await expect(page.locator('#rankings').getByRole('button', { name: /sign in/i })).toHaveCount(0);
    await expect(page.locator('#rankings').getByRole('link', { name: /sign in/i })).toHaveCount(0);
    await expect(page.getByTestId('rankings-another')).toHaveCount(0);

    // Class and Spec are the first two controls, then the Encounter.
    const firstThree = await page
      .getByTestId('rankings-filter-panel')
      .locator('select')
      .evaluateAll((nodes) => nodes.slice(0, 3).map((node) => node.getAttribute('data-testid')));
    expect(firstThree).toEqual(['filter-class', 'filter-spec', 'filter-encounter']);

    // A board of more than one class shows a ringed crest before each name.
    await expect(page.getByTestId('ranking-1').locator('img')).toHaveCount(1);
    await expect(page.getByTestId('ranking-1').locator('img')).toHaveAttribute('src', /crests\//);
    expect(api.asked.every((url) => url.searchParams.get('character') === null)).toBe(true);
  });
});

test.describe('phone', () => {
  test.use({ viewport: PHONE });
  test.skip(() => test.info().project.name !== 'mobile', 'phone layout only');

  test('the filters collapse to one summary line and a Filters button opens them', async ({ page }) => {
    await openSignedIn(page);

    const summary = page.getByTestId('rankings-summary-line');
    await expect(summary).toContainText('Onyxia · Normal · All regions ·');
    await expect(page.getByTestId('rankings-filter-panel')).toBeHidden();
    const toggle = page.getByTestId('rankings-filters-toggle');
    await expect(toggle).toHaveAttribute('aria-expanded', 'false');
    expect((await toggle.boundingBox())?.height).toBeGreaterThanOrEqual(44);

    await toggle.click();
    await expect(toggle).toHaveAttribute('aria-expanded', 'true');
    await expect(page.getByTestId('rankings-filter-panel')).toBeVisible();
    await expect(page.getByTestId('filter-encounter')).toBeVisible();
    await page.getByTestId('filter-region').selectOption('eu');
    await expect(summary).toContainText('EU');

    // A row is a name over a muted line, the figure and Report; the own row is still marked.
    await expect(page.getByTestId('ranking-14')).toHaveAttribute('data-own', 'true');
    await expect(page.getByTestId('ranking-14').getByTestId('ranking-meta')).toContainText('40-man');
    expect(await page.evaluate(() => document.documentElement.scrollWidth <= window.innerWidth)).toBe(true);
  });
});

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
    deduped: true,
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

  await gotoHydrated(page, '/rankings/warden-kelthas', 'rankings-scope');

  await expect(page.locator('h1')).toHaveText('Rankings');
  await openFilters(page);
  await expect(page.getByTestId('filter-encounter')).toHaveValue('warden-kelthas');
  await expect(page.getByTestId('ranking-1')).toContainText('Elyra Duskvale');
  await expect(page.getByTestId('ranking-1').getByTestId('ranking-build')).toHaveAttribute(
    'href',
    '/b/k7x2qm4a',
  );
  await expect(page.getByTestId('ranking-1').getByTestId('ranking-report')).toHaveAttribute(
    'href',
    '/reports/fixture2abcd?fight=3',
  );
  // The guild is a link in its own column from 1100px; below that it is a muted line or gone
  // (design/specs/2026-10-11-rankings.md section 3, decision 6).
  if (test.info().project.name === 'desktop') {
    await expect(page.getByRole('link', { name: 'The Last Watch' })).toHaveAttribute(
      'href',
      '/guild/us/hardcore/the-last-watch',
    );
  }
  await expect(page.getByTestId('ranking-2').getByTestId('ranking-state')).toHaveText(
    'Removed by a moderator',
  );

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

  await gotoHydrated(page, '/rankings/warden-kelthas', 'rankings-scope');
  await openFilters(page);
  await page.getByTestId('filter-metric').selectOption('hps');
  await page.getByTestId('filter-ruleset').selectOption('hardcore');

  await expect(page).toHaveURL(/\?metric=hps&ruleset=hardcore$/);
  expect(asked[asked.length - 1]).toContain('metric=hps');
  expect(asked[asked.length - 1]).toContain('ruleset=hardcore');

  // "All phases" is a choice, so it survives in the link rather than reverting to the current phase.
  await page.getByTestId('filter-phase').selectOption('');
  await expect(page).toHaveURL(/phase=all/);
  expect(asked[asked.length - 1]).not.toContain('phase=');
});

test('the encounter is a filter: choosing another one moves the board and the path', async ({ page }) => {
  await stubRankingsApi(page);
  await gotoHydrated(page, '/rankings/warden-kelthas', 'rankings-scope');
  await openFilters(page);
  await expect(page.getByTestId('filter-encounter')).toHaveValue('warden-kelthas');

  await page.getByTestId('filter-encounter').selectOption('onyxia');
  await expect(page).toHaveURL(/\/rankings\/onyxia$/);
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

  await gotoHydrated(page, '/rankings/warden-kelthas', 'rankings-scope');
  await page.getByTestId('board-guild').click();

  await expect(page.getByTestId('guild-rows')).toContainText('The Last Watch');
  expect(guildUrl).toContain('kind=progress');
});

// Spec 2026-09-25 section 6: an empty result says what fills it and offers one action,
// through EmptyState -- covering both boards, since both are "no data because nothing has
// been uploaded yet" (2026-09-25-states ledger, ruling 1).
test('the character board with no ranked fights says so and links to Upload a log', async ({ page }) => {
  await page.route('**/v1/rankings?**', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({
        ok: true,
        data: { rows: [], total: 0, page: 1, per_page: 100, updated_at: '2026-12-09T22:15:00Z' },
        error: null,
        request_id: 'r',
      }),
    }),
  );

  await gotoHydrated(page, '/rankings/warden-kelthas', 'rankings-scope');

  const empty = page.getByTestId('rankings-empty');
  await expect(empty).toContainText(
    'No ranked fights yet for this filter. Rankings fill in as reports are uploaded.',
  );
  await expect(empty.getByRole('link', { name: 'Upload a log' })).toHaveAttribute('href', '/logs');
});

test('the guild board with no ranked guilds says so and links to Upload a log', async ({ page }) => {
  await page.route('**/v1/rankings?**', (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(ROWS) }),
  );
  await page.route('**/v1/rankings/guilds?**', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({ ok: true, data: { rows: [] }, error: null, request_id: 'r' }),
    }),
  );

  await gotoHydrated(page, '/rankings/warden-kelthas', 'rankings-scope');
  await page.getByTestId('board-guild').click();

  const empty = page.getByTestId('rankings-empty');
  await expect(empty).toContainText(
    'No ranked fights yet for this filter. Rankings fill in as reports are uploaded.',
  );
  await expect(empty.getByRole('link', { name: 'Upload a log' })).toHaveAttribute('href', '/logs');
});

test('a failed rankings call says so instead of showing an empty board', async ({ page }) => {
  await page.route('**/v1/rankings?**', (route) => route.abort());
  await gotoHydrated(page, '/rankings/warden-kelthas', 'rankings-scope');
  await expect(page.getByTestId('rankings-error')).toBeVisible();
});

// LoadError's Try again has to re-fire the same request in place, never reload the page
// (design 2026-09-22 spec section 1.5). A flag rather than an attempt counter: it flips
// between the two halves of the test, so which request counts as the retry never depends
// on how many the board made on the way in.
test('Try again on a failed board re-fires the same request and paints the rows', async ({ page }) => {
  let failing = true;
  await page.route('**/v1/rankings?**', async (route) => {
    if (failing) {
      await route.abort();
      return;
    }
    await route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(ROWS) });
  });

  await gotoHydrated(page, '/rankings/warden-kelthas', 'rankings-scope');
  await expect(page.getByTestId('rankings-error')).toBeVisible();

  failing = false;
  await page.getByTestId('rankings-error-retry').click();

  await expect(page.getByTestId('ranking-rows').locator('li')).toHaveCount(2);
  await expect(page.getByTestId('rankings-error')).toHaveCount(0);
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

  await gotoHydrated(page, '/rankings/warden-kelthas', 'rankings-scope');
  await slow.started;

  await openFilters(page);
  await page.getByTestId('filter-metric').selectOption('hps');
  await expect(page.getByTestId('ranking-1')).toContainText('Brightwell Solace');

  const staleAnswered = page.waitForResponse((response) => response.url().includes('metric=dps'));
  slow.release();
  await staleAnswered;
  await page.waitForTimeout(250);

  // Still Healing's one row, not Damage's two: the late Damage answer was dropped.
  await expect(page.getByTestId('ranking-1')).toContainText('Brightwell Solace');
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
        execution_score: 0.94,
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
        execution_score: null,
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
        execution_score: 0.94,
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
        execution_score: null,
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
  // L8.4 fix round 2 item 8: the rows carry execution_score, as the API always sends it.
  await expect(page.getByTestId('character-best-execution').first()).toHaveText('94%');
  await expect(page.getByTestId('character')).not.toContainText('NaN');
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

test('a guild page leads with the pull count and kills, and its Progression tab lists roster bests', async ({
  page,
}) => {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(GUILD) }),
  );

  await page.goto('/guild/us/hardcore/the-last-watch');

  await expect(page.locator('h1')).toHaveText('The Last Watch');
  // The public read's progression rows (one boss with kills, 14 + 31 pulls) feed the header.
  await expect(page.getByTestId('guild')).toContainText('1 boss down · 45 pulls');
  await expect(page.getByTestId('guild')).toContainText('Sanguine Depths, fixture night');

  // The per-boss lists moved to the Progression tab, which reads the guild's own progression
  // endpoint (guild-centre.spec.ts covers it); this fixture names no guild id, so the tab
  // shows its empty line above the public read's roster bests.
  await page.getByTestId('guild-tab-progression').click();
  await expect(page.getByTestId('guild-progression-empty')).toBeVisible();
  const best = page.getByTestId('guild-roster-best');
  await expect(best.getByRole('link', { name: 'Elyra Duskvale' })).toHaveAttribute(
    'href',
    '/character/us/hardcore/elyra-duskvale',
  );
  await expect(best).toContainText('Warden Kelthas');
  await expect(best).not.toContainText('NaN');
});
