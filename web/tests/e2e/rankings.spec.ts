// web/tests/e2e/rankings.spec.ts
// The Rankings page (design/specs/2026-11-07-logs-reports.md section 7; the rules are
// 2026-10-11-rankings.md's): signed in, the PageHeader's answer says where the character stands
// and the board marks its row in place (or pins it on top when it is off the page); a header
// switch re-aims the board without a reload; signed out, a board has no answer and every class and
// spec; the ControlBar carries the filters (Region, Phase, Faction and Today only behind Filters),
// and on a phone collapses to one summary line. The older checks (rows, filters in the URL, guild
// board, empty, failed, stale responses) run signed out against a plain two-row board.
import type { Page } from '@playwright/test';
import { apiUnreachable, expect, test } from './support/test';
import { gotoHydrated } from './support/hydrated';
import { heldRoute } from './support/held-route';
import { barToggle, moreFilter, openFilters } from './support/rankings-filters';
import { FROST_KEY, OWN_KEY, stubRankingsApi, type BoardOptions } from './support/rankings-fixture';
import { OBNOXIOUS, armoryPointer, keyOf, signInWith, signedOut, storePointer } from './support/selector';
import type { FixtureCharacter } from './support/selector';
import { choose, fieldBox } from './support/select';

// Reads this spec does not answer go unanswered, as the API out of reach (support/test.ts).
test.beforeEach(({ page }) => apiUnreachable(page, '**/v1/encounters*'));

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
const answer = (page: Page) => page.getByTestId('rankings-answer');
const heading = (page: Page) => page.getByTestId('rankings-heading').locator('h2');
/** The board's own rows, marked as the viewer's (selected: the gold ground and rail). */
const ownRows = (page: Page) => page.locator('#rankings li[aria-current="true"]');

async function openSignedIn(page: Page, options: BoardOptions = {}) {
  await signInWith(page, [OBNOXIOUS, FROSTBYTE]);
  await storePointer(page, armoryPointer(OBNOXIOUS));
  const api = await stubRankingsApi(page, options);
  await page.goto('/rankings');
  await expect(page.getByTestId('rankings-answer-rank')).toBeVisible();
  await expect(page.getByTestId('ranking-1')).toBeVisible();
  return api;
}

test.describe('signed in, desktop', () => {
  test.use({ viewport: DESKTOP });
  test.skip(() => test.info().project.name !== 'desktop', 'desktop layout only');

  test('the header answers with the rank, the board opens on the character, and its row is marked in place', async ({
    page,
  }) => {
    const api = await openSignedIn(page);

    // The encounter is the title; the meta names the board and the ruleset.
    await expect(page.locator('main h1')).toHaveText('Onyxia');
    await expect(page.locator('#rankings header')).toContainText('Rankings · Fury Warriors · Normal ruleset');
    await expect(answer(page)).toContainText('Your rank');
    await expect(answer(page)).toContainText('Obnoxious Yell');
    await expect(page.getByTestId('rankings-answer-rank')).toHaveText(/^\s*14th\s*of 212 Fury Warriors\s*$/);
    await expect(answer(page)).toContainText('On Onyxia, Normal ruleset');
    await expect(page.getByTestId('rankings-answer-parses')).toHaveText('See your parses ›');
    await expect(page.getByTestId('rankings-answer-parses')).toHaveAttribute(
      'href',
      '/character/us/normal/obnoxious-yell',
    );

    // The board opened on the encounter the character most recently killed, in its own cohort.
    await expect(page).toHaveURL(/\/rankings\/onyxia$/);
    await expect(heading(page)).toHaveText('Fury Warriors');
    await expect(page.getByTestId('rankings-heading')).toContainText(
      '212 ranked, each character’s best kill',
    );
    const board = api.asked.find(
      (url) => url.searchParams.get('character') === OWN_KEY && url.searchParams.get('page') === null,
    );
    expect(board?.searchParams.get('encounter')).toBe('onyxia');
    expect(board?.searchParams.get('class')).toBe('warrior');
    expect(board?.searchParams.get('spec')).toBe('Fury');
    expect(board?.searchParams.get('ruleset')).toBe('normal');
    expect(board?.searchParams.get('metric')).toBe('dps');
    expect(board?.searchParams.get('distinct')).toBe('character');

    // The own row is marked where it is, never pinned, with the You tag; the others are not marked.
    const own = page.getByTestId('ranking-14');
    await expect(own).toHaveAttribute('aria-current', 'true');
    await expect(page.getByTestId('ranking-13')).not.toHaveAttribute('aria-current', 'true');
    await expect(ownRows(page)).toHaveCount(1);
    await expect(page.getByTestId('ranking-pinned-row')).toHaveCount(0);
    await expect(own.getByTestId('rank-name')).toHaveText('Obnoxious Yell');
    await expect(own).toContainText('You');
    await expect(own.getByTestId('rank-rank')).toHaveText('14');
  });

  test('the page is the spec: one stamp, no pills, no spine bar, a legend, plain Build and gold Report', async ({
    page,
  }) => {
    await openSignedIn(page);

    await expect(page.locator('#rankings .pill')).toHaveCount(0);
    await expect(page.getByTestId('current-character-bar')).toHaveCount(0);
    // The retired shared band: the board's answer is the PageHeader's own.
    await expect(page.locator('#rankings [data-testid^="character-header"]')).toHaveCount(0);
    await expect(page.locator('#rankings header')).toContainText('Updated 13 Dec');
    await expect(page.getByText(/^Updated /)).toHaveCount(1);
    await expect(page.locator('#rankings a[href="/tiers"]')).toHaveText('Tier list ›');
    await expect(page.locator('#rankings')).toContainText('Comparing specs, not characters?');
    await expect(page.getByTestId('rankings-legend')).toContainText('Rank colour is the percentile:');
    await expect(page.getByTestId('rankings-legend').locator('b')).toHaveText([
      '100',
      '99+',
      '95+',
      '75+',
      '50+',
      '25+',
      'under 25',
    ]);
    // "100" is gold, the colour of a board's first rank.
    await expect(page.getByTestId('rankings-legend').locator('b').first()).toHaveAttribute(
      'style',
      /--color-parse-gold/,
    );

    // A single-class board draws no crest in its rows (the header carries it once).
    await expect(page.getByTestId('ranking-board').locator('img')).toHaveCount(0);

    const row = page.getByTestId('ranking-1');
    const colour = (testid: string) =>
      row.getByTestId(testid).evaluate((node) => getComputedStyle(node).color);
    expect(await colour('ranking-build')).not.toBe(await colour('ranking-report'));
    await expect(row.getByTestId('ranking-build')).toHaveAttribute('href', '/b/b1');
    await expect(row.getByTestId('ranking-report')).toHaveAttribute('href', '/reports/report1?fight=1');

    // No sign-in prose or button of the page's own.
    await expect(page.locator('#rankings').getByText(/sign in/i)).toHaveCount(0);
  });

  test('each rank is in its percentile colour, floored: 1 of 212 gold, 2 and 3 pink, 4 orange', async ({
    page,
  }) => {
    await openSignedIn(page);

    const rankColour = (rank: number) => page.getByTestId(`ranking-${rank}`).getByTestId('rank-rank');
    await expect(rankColour(1)).toHaveAttribute('style', /var\(--color-parse-gold\)/);
    await expect(rankColour(2)).toHaveAttribute('style', /var\(--color-parse-pink\)/);
    await expect(rankColour(3)).toHaveAttribute('style', /var\(--color-parse-pink\)/);
    await expect(rankColour(4)).toHaveAttribute('style', /var\(--color-parse-orange\)/);
    // The colour is applied, not only declared: gold paints differently from pink.
    const painted = (rank: number) => rankColour(rank).evaluate((node) => getComputedStyle(node).color);
    expect(await painted(1)).not.toBe(await painted(2));
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

    await expect(answer(page)).toContainText('Frostbyte');
    await expect(heading(page)).toHaveText('Frost Mages');
    await expect(page).toHaveURL(/\/rankings\/warden-kelthas$/);
    await expect(page.locator('main h1')).toHaveText('Warden Kelthas');
    await expect(page.getByTestId('rankings-answer-rank')).toHaveText(/^\s*3rd\s*of 3 Frost Mages\s*$/);
    await expect(answer(page)).toContainText('On Warden Kelthas, Normal ruleset');
    await expect(page.getByTestId('ranking-3')).toHaveAttribute('aria-current', 'true');
    await expect(ownRows(page)).toHaveCount(1);
    // A mage board is one class too: still no crest.
    await expect(page.getByTestId('ranking-board').locator('img')).toHaveCount(0);
    expect(await page.evaluate(() => (window as unknown as { __kept?: boolean }).__kept)).toBe(true);

    const mage = api.asked.filter((url) => url.searchParams.get('class') === 'mage');
    expect(mage.length).toBeGreaterThan(0);
    expect(mage.every((url) => url.searchParams.get('spec') === 'Frost')).toBe(true);
    expect(mage.some((url) => url.searchParams.get('character') === FROST_KEY)).toBe(true);
  });

  test('a best kill off the visible page is pinned on top at its true rank, once', async ({ page }) => {
    await openSignedIn(page, { ownOffPage: true });

    await expect(page.getByTestId('rankings-answer-rank')).toHaveText(/^\s*204th\s*of 212 Fury Warriors\s*$/);
    const pinned = page.getByTestId('ranking-pinned-row');
    await expect(pinned).toHaveAttribute('aria-current', 'true');
    await expect(pinned).toContainText('You');
    await expect(pinned.getByTestId('rank-rank')).toHaveText('204');
    // The pin is the first row of the board; the board follows from #1, unmarked.
    await expect(page.getByTestId('ranking-board').locator('li').first()).toHaveAttribute(
      'data-testid',
      'ranking-pinned-row',
    );
    await expect(page.getByTestId('ranking-1')).toBeVisible();
    await expect(page.getByTestId('ranking-1')).not.toHaveAttribute('aria-current', 'true');
    await expect(ownRows(page)).toHaveCount(1);
    const pinnedBox = await pinned.boundingBox();
    const firstBox = await page.getByTestId('ranking-1').boundingBox();
    expect((pinnedBox?.y ?? Infinity) < (firstBox?.y ?? 0)).toBe(true);
  });

  test('another class on the bar says whose board it is, and goes back', async ({ page }) => {
    await openSignedIn(page);

    await choose(page.getByTestId('control-class'), 'mage');
    await expect(heading(page)).toHaveText('Mages');
    await expect(page).toHaveURL(/class=mage/);
    await expect(page.getByTestId('rankings-showing-note')).toContainText(
      'Showing mages. Your rank is on the Fury Warrior board.',
    );
    // The header's answer stays about the character, and no row on this board is marked.
    await expect(page.getByTestId('rankings-answer-rank')).toHaveText(/^\s*14th\s*of 212 Fury Warriors\s*$/);
    await expect(ownRows(page)).toHaveCount(0);
    await expect(page.getByTestId('ranking-pinned-row')).toHaveCount(0);

    await page.getByTestId('rankings-back').click();
    await expect(heading(page)).toHaveText('Fury Warriors');
    await expect(page.getByTestId('rankings-showing-note')).toHaveCount(0);
    await expect(page.getByTestId('ranking-14')).toHaveAttribute('aria-current', 'true');
  });

  test('the pager walks the board in place: Previous muted on page one, Next a link', async ({ page }) => {
    const api = await openSignedIn(page);

    const pager = page.getByTestId('rankings-pager');
    await expect(pager).toContainText('Page 1 of 3');
    await expect(page.getByTestId('rankings-prev')).not.toHaveAttribute('href');
    await expect(page.getByTestId('rankings-prev')).toHaveAttribute('aria-disabled', 'true');
    await expect(page.getByTestId('rankings-next')).toHaveAttribute('href', /page=2/);

    await page.getByTestId('rankings-next').click();
    await expect(page).toHaveURL(/page=2/);
    await expect(pager).toContainText('Page 2 of 3');
    await expect(page.getByTestId('rankings-prev')).toHaveAttribute('href', /\/rankings\/onyxia/);
    await expect.poll(() => api.asked.at(-1)?.searchParams.get('page')).toBe('2');
  });

  test('a character with no ranked kill is told how to get on the board', async ({ page }) => {
    await signInWith(page, [OBNOXIOUS]);
    await storePointer(page, armoryPointer(OBNOXIOUS));
    await stubRankingsApi(page, { unranked: true });
    await page.goto('/rankings');

    await expect(page.getByTestId('rankings-answer-unranked')).toHaveText(
      'Not ranked yet on Onyxia. Upload a log of a kill to get on the board.',
    );
    await expect(answer(page).getByRole('link', { name: 'Upload a log ›' })).toHaveAttribute('href', '/logs');
    await expect(page.getByTestId('ranking-1')).toBeVisible();
    await expect(ownRows(page)).toHaveCount(0);
    await expect(page.getByTestId('ranking-pinned-row')).toHaveCount(0);
  });

  test('an API that does not dedupe yet gives the header no rank to show, and no pin', async ({ page }) => {
    await signInWith(page, [OBNOXIOUS]);
    await storePointer(page, armoryPointer(OBNOXIOUS));
    await stubRankingsApi(page, { legacy: true });
    await page.goto('/rankings');

    await expect(answer(page)).toContainText('Your rank did not load.');
    await expect(page.getByTestId('rankings-answer-rank')).toHaveCount(0);
    await expect(page.getByTestId('ranking-1')).toBeVisible();
    await expect(page.getByTestId('ranking-pinned-row')).toHaveCount(0);
  });
});

test.describe('signed out, desktop', () => {
  test.use({ viewport: DESKTOP });
  test.skip(() => test.info().project.name !== 'desktop', 'desktop layout only');

  test('a board has no answer, every class and spec, and the bar in the spec order', async ({ page }) => {
    await signedOut(page);
    const api = await stubRankingsApi(page);
    await gotoHydrated(page, '/rankings/warden-kelthas', 'rankings-heading');
    await expect(page.getByTestId('ranking-1')).toBeVisible();

    await expect(page.getByTestId('rankings-answer')).toHaveCount(0);
    await expect(page.locator('main h1')).toHaveText('Warden Kelthas');
    await expect(heading(page)).toHaveText('Every class and spec');
    await expect(page.getByTestId('rankings-showing-note')).toHaveCount(0);
    // No button of the page's own: the nav selector is the only sign-in control.
    await expect(page.locator('#rankings').getByRole('button', { name: /sign in/i })).toHaveCount(0);
    await expect(page.locator('#rankings').getByRole('link', { name: /sign in/i })).toHaveCount(0);

    // Encounter, Class, Spec and Ruleset are the bar's selects, in that order (spec section 7).
    const selects = await page
      .getByTestId('rankings-filters')
      .locator('.ui-cbar-select select')
      .evaluateAll((nodes) => nodes.map((node) => node.getAttribute('data-testid')));
    expect(selects).toEqual(['control-encounter', 'control-class', 'control-spec', 'control-ruleset']);
    await expect(page.getByTestId('control-metric-dps')).toHaveAttribute('aria-checked', 'true');
    await expect(page.getByTestId('control-board-character')).toHaveAttribute('aria-checked', 'true');

    // A board of more than one class shows a ringed crest before each name.
    await expect(page.getByTestId('ranking-1').locator('img')).toHaveCount(1);
    await expect(page.getByTestId('ranking-1').locator('img')).toHaveAttribute('src', /crests\//);
    expect(api.asked.every((url) => url.searchParams.get('character') === null)).toBe(true);
  });

  test('the Filters popover holds Region, Phase, Faction and Today only, counts them and clears them', async ({
    page,
  }) => {
    await signedOut(page);
    const api = await stubRankingsApi(page);
    await gotoHydrated(page, '/rankings/warden-kelthas', 'rankings-heading');

    const button = page.getByTestId('rankings-filters-button');
    await expect(button).toHaveText('Filters');
    await expect(page.getByTestId('rankings-filters-popover')).toBeHidden();
    await button.click();
    await expect(button).toHaveAttribute('aria-expanded', 'true');
    const popover = page.getByTestId('rankings-filters-popover');
    await expect(popover).toBeVisible();
    await expect(fieldBox(popover, 'rankings-pop-region')).toBeVisible();
    await expect(fieldBox(popover, 'rankings-pop-phase')).toBeVisible();
    await expect(fieldBox(popover, 'rankings-pop-faction')).toBeVisible();

    await choose(page.getByTestId('rankings-pop-region'), 'eu');
    await page.getByTestId('rankings-pop-today').click();
    await expect(page).toHaveURL(/region=eu/);
    await expect(page).toHaveURL(/since=today/);
    await expect(page.getByTestId('rankings-pop-today')).toHaveAttribute('aria-checked', 'true');
    await expect(button).toHaveText('Filters (2)');
    await expect.poll(() => api.asked.at(-1)?.searchParams.get('region')).toBe('eu');
    await expect.poll(() => api.asked.at(-1)?.searchParams.get('since')).not.toBeNull();

    await page.getByTestId('rankings-filters-clear').click();
    await expect(button).toHaveText('Filters');
    await expect(page).not.toHaveURL(/region=|since=/);

    // Escape closes it and gives focus back to the button.
    await page.getByTestId('rankings-filters-clear').focus();
    await page.keyboard.press('Escape');
    await expect(popover).toBeHidden();
    await expect(button).toBeFocused();
  });
});

test.describe('phone', () => {
  test.use({ viewport: PHONE });
  test.skip(() => test.info().project.name !== 'mobile', 'phone layout only');

  test('the bar collapses to one summary line and its toggle opens the stack', async ({ page }) => {
    await openSignedIn(page);

    const toggle = barToggle(page);
    await expect(toggle).toContainText('Onyxia · Fury');
    await expect(toggle).toHaveAttribute('aria-expanded', 'false');
    await expect(fieldBox(page, 'control-encounter')).toBeHidden();
    await expect(page.getByTestId('rankings-filters-button')).toBeHidden();
    expect((await toggle.boundingBox())?.height).toBeGreaterThanOrEqual(44);

    await toggle.click();
    await expect(toggle).toHaveAttribute('aria-expanded', 'true');
    await expect(fieldBox(page, 'control-encounter')).toBeVisible();
    // The Filters fields sit inline in the stack, not behind a popover.
    await expect(fieldBox(page, 'rankings-stack-region')).toBeVisible();
    await choose(page.getByTestId('rankings-stack-region'), 'eu');
    await expect(page).toHaveURL(/region=eu/);

    // A row is a name over a muted line, the figure and Report; the own row is still marked.
    await expect(page.getByTestId('ranking-14')).toHaveAttribute('aria-current', 'true');
    await expect(page.getByTestId('ranking-14')).toContainText('40-man');
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

  await gotoHydrated(page, '/rankings/warden-kelthas', 'rankings-heading');

  await expect(page.locator('main h1')).toHaveText('Warden Kelthas');
  await openFilters(page);
  await expect(page.getByTestId('control-encounter')).toHaveValue('warden-kelthas');
  await expect(page.getByTestId('ranking-1').getByTestId('rank-name')).toHaveText('Elyra Duskvale');
  // The sub line: guild, spec and class, raid size.
  await expect(page.getByTestId('ranking-1')).toContainText('The Last Watch · Shadow Priest · 40-man');
  await expect(page.getByTestId('ranking-1').getByTestId('ranking-value')).toHaveText('1,840');
  await expect(page.getByTestId('ranking-1').getByTestId('ranking-build')).toHaveAttribute(
    'href',
    '/b/k7x2qm4a',
  );
  await expect(page.getByTestId('ranking-1').getByTestId('ranking-report')).toHaveAttribute(
    'href',
    '/reports/fixture2abcd?fight=3',
  );
  await expect(page.getByTestId('ranking-2').getByTestId('rank-note')).toHaveText('Removed by a moderator');
  await expect(page.getByTestId('ranking-1').getByTestId('rank-note')).toHaveCount(0);

  // Row 2 has no guild (`player.key: 'eu/normal/other-raider'`), so a fallback that
  // defaulted region/ruleset to `us`/`normal` would link it there instead -- the wrong
  // character page, on a real player, with a link that looks correct. The href must carry
  // the row's own region and ruleset, read out of `player.key`.
  await expect(page.getByTestId('ranking-2').getByTestId('rank-name')).toHaveAttribute(
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

  await gotoHydrated(page, '/rankings/warden-kelthas', 'rankings-heading');
  await openFilters(page);
  await page.getByTestId('control-metric-hps').click();
  await expect(page.getByTestId('control-metric-hps')).toHaveAttribute('aria-checked', 'true');
  await openFilters(page);
  await choose(page.getByTestId('control-ruleset'), 'hardcore');

  await expect(page).toHaveURL(/\?metric=hps&ruleset=hardcore$/);
  expect(asked[asked.length - 1]).toContain('metric=hps');
  expect(asked[asked.length - 1]).toContain('ruleset=hardcore');

  // "All phases" is a choice, so it survives in the link rather than reverting to the current phase.
  await choose(await moreFilter(page, 'phase'), '');
  await expect(page).toHaveURL(/phase=all/);
  expect(asked[asked.length - 1]).not.toContain('phase=');
});

test('the encounter is a filter: choosing another one moves the board and the path', async ({ page }) => {
  await stubRankingsApi(page);
  await gotoHydrated(page, '/rankings/warden-kelthas', 'rankings-heading');
  await openFilters(page);
  await expect(page.getByTestId('control-encounter')).toHaveValue('warden-kelthas');
  const before = await page.evaluate(() => performance.getEntriesByType('navigation').length);

  await choose(page.getByTestId('control-encounter'), 'onyxia');
  await expect(page).toHaveURL(/\/rankings\/onyxia$/);
  await expect(page.locator('main h1')).toHaveText('Onyxia');
  await expect(page.getByTestId('control-encounter')).toHaveValue('onyxia');
  // In place: the path moved without a page load.
  expect(await page.evaluate(() => performance.getEntriesByType('navigation').length)).toBe(before);
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

  await gotoHydrated(page, '/rankings/warden-kelthas', 'rankings-heading');
  await openFilters(page);
  await page.getByTestId('control-board-guild').click();

  await expect(page.getByTestId('control-board-guild')).toHaveAttribute('aria-checked', 'true');
  await expect(page).toHaveURL(/board=guild/);
  await expect(page.getByTestId('guild-rows')).toContainText('The Last Watch');
  await expect(page.getByTestId('guild-rows').getByRole('link', { name: /The Last Watch/ })).toHaveAttribute(
    'href',
    '/guild/us/hardcore/the-last-watch',
  );
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

  await gotoHydrated(page, '/rankings/warden-kelthas', 'rankings-heading');

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

  await gotoHydrated(page, '/rankings/warden-kelthas', 'rankings-heading');
  await openFilters(page);
  await page.getByTestId('control-board-guild').click();

  const empty = page.getByTestId('rankings-empty');
  await expect(empty).toContainText(
    'No ranked fights yet for this filter. Rankings fill in as reports are uploaded.',
  );
  await expect(empty.getByRole('link', { name: 'Upload a log' })).toHaveAttribute('href', '/logs');
});

test('a failed rankings call says so instead of showing an empty board', async ({ page }) => {
  await page.route('**/v1/rankings?**', (route) => route.abort());
  await gotoHydrated(page, '/rankings/warden-kelthas', 'rankings-heading');
  // The spec's copy ends with a full stop; api.ts's RANKINGS_FAILED does not yet (reported).
  await expect(page.getByTestId('rankings-error')).toContainText('Rankings did not load');
  await expect(page.getByTestId('rankings-error').getByRole('button', { name: 'Try again' })).toBeVisible();
  await expect(page.getByTestId('rankings-empty')).toHaveCount(0);
  await expect(page.getByTestId('ranking-board')).toHaveCount(0);
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

  await gotoHydrated(page, '/rankings/warden-kelthas', 'rankings-heading');
  await expect(page.getByTestId('rankings-error')).toBeVisible();

  failing = false;
  await page.getByTestId('rankings-error').getByRole('button', { name: 'Try again' }).click();

  await expect(page.getByTestId('ranking-board').locator('li')).toHaveCount(2);
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

  await gotoHydrated(page, '/rankings/warden-kelthas', 'rankings-heading');
  await slow.started;

  await openFilters(page);
  await page.getByTestId('control-metric-hps').click();
  await expect(page.getByTestId('ranking-1')).toContainText('Brightwell Solace');

  const staleAnswered = page.waitForResponse((response) => response.url().includes('metric=dps'));
  slow.release();
  await staleAnswered;
  await page.waitForTimeout(250);

  // Still Healing's one row, not Damage's two: the late Damage answer was dropped.
  await expect(page.getByTestId('ranking-1')).toContainText('Brightwell Solace');
  await expect(page.getByTestId('ranking-board').locator('li')).toHaveCount(1);
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

test('a character page shows the best per encounter and the builds they were seen in', async ({ page }) => {
  await page.route('**/v1/characters/us/hardcore/elyra-duskvale', (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(CHARACTER) }),
  );

  await page.goto('/character/us/hardcore/elyra-duskvale');

  await expect(page.getByTestId('character-head').getByTestId('page-title')).toHaveText('Elyra Duskvale');
  await expect(page.getByTestId('character-fact-chip')).toHaveText('Hardcore US');
  const best = page.getByTestId('character-best');
  await expect(best).toContainText('Warden Kelthas');
  await expect(best.getByRole('link', { name: /Warden Kelthas/ })).toHaveAttribute(
    'href',
    '/rankings/warden-kelthas',
  );
  // The parse percentile and the amount on each row, never NaN.
  await expect(best.getByLabel('96 percentile')).toHaveText('96');
  await expect(best).toContainText('1,840');
  await expect(page.getByTestId('character')).not.toContainText('NaN');
  // Every ranked fight lives on /logs, counted from the history.
  await expect(best.getByRole('link', { name: 'All fights (1) ›' })).toHaveAttribute('href', '/logs');
  const builds = page.getByTestId('character-builds');
  await expect(builds).toContainText('31/20/0');
  await expect(builds).toContainText('Discipline · first seen 9 Dec 2026');

  // rankings/api.ts's encounterSlug is the one derivation of this identifier, trailing
  // hyphen trimmed and all. A second, inlined copy here linked "Emeriss (Dream)" to
  // /rankings/emeriss-dream- , which the Worker accepts and renders as a permanently empty
  // board, while every other page on the site links to /rankings/emeriss-dream.
  await expect(best.getByRole('link', { name: /Emeriss \(Dream\)/ })).toHaveAttribute(
    'href',
    '/rankings/emeriss-dream',
  );
});

test('a guild page leads with the pull count and kills, and its Progression tab lists roster bests', async ({
  page,
}) => {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(GUILD) }),
  );

  await page.goto('/guild/us/hardcore/the-last-watch');

  await expect(page.locator('main h1')).toHaveText('The Last Watch');
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

// Stage 2 fix round, on the new frame: the bar's board switch never overlaps the board's heading,
// at any width (below 768 with the bar's stack open, where the switch is shown).
test('the board switch and the board heading never overlap from 320 to 2000', async ({ page }) => {
  await stubRankingsApi(page);
  await gotoHydrated(page, '/rankings/warden-kelthas', 'rankings-heading');
  for (const width of [320, 360, 390, 480, 768, 1024, 1440, 2000]) {
    await page.setViewportSize({ width, height: 900 });
    await openFilters(page);
    const title = await heading(page).boundingBox();
    const switchBox = await page.getByTestId('control-board').boundingBox();
    expect(title, `title at ${width}`).not.toBeNull();
    expect(switchBox, `switch at ${width}`).not.toBeNull();
    const apart =
      (title?.x ?? 0) + (title?.width ?? 0) <= (switchBox?.x ?? 0) ||
      (switchBox?.x ?? 0) + (switchBox?.width ?? 0) <= (title?.x ?? 0) ||
      (title?.y ?? 0) + (title?.height ?? 0) <= (switchBox?.y ?? 0) ||
      (switchBox?.y ?? 0) + (switchBox?.height ?? 0) <= (title?.y ?? 0);
    expect(apart, `overlap at ${width}`).toBe(true);
  }
});
