// web/tests/e2e/guild-week-one.spec.ts
// A leveling guild with zero logs (design/specs/2026-11-05-guild-week-one.md §9): the band, the
// Overview and Progression print one set of numbers, the page never prints NaN, a level 58 row has
// no Nudge and no gear figure, and a healer row never says DPS.
import type { Page } from '@playwright/test';
import { apiUnreachable, expect, test } from './support/test';
import {
  GUILD_ID,
  VIEWER_MEMBER,
  VIEWER_OFFICER,
  buildMockGuildPage,
  buildMockHome,
  buildMockProgression,
  buildMockRaids,
  buildMockReadiness,
} from '../../src/fixtures/guild/mock-guild';
import type {
  GuildHome,
  GuildProgressionPage,
  GuildReadinessPage,
  GuildReadinessRow,
} from '../../src/lib/guild/api';

const GUILD_PATH = '/guild/us/pvp/olympus-xxvii';
/** Set to capture the spec's acceptance screenshots into a directory. */
const SHOTS = process.env.WEEK_ONE_SHOTS;

test.beforeEach(({ page }) => apiUnreachable(page, '**/v1/guilds/*/loot*', '**/v1/guilds/*/raids*'));

function envelope(data: unknown) {
  return {
    status: 200,
    contentType: 'application/json',
    body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
  };
}

function meFor(name: string, officer: boolean) {
  return {
    user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
    characters: [
      {
        key: `us/pvp/${name.toLowerCase()}`,
        region: 'us',
        ruleset: 'pvp',
        name,
        class: 'warrior',
        guild: { id: GUILD_ID, name: 'Olympus XXVII', verified: true },
      },
    ],
    guilds: [
      {
        id: GUILD_ID,
        region: 'us',
        ruleset: 'pvp',
        name: 'Olympus XXVII',
        rank: officer ? 'leader' : 'member',
        consent: 'gear',
        verified: true,
        can: officer ? ['approve', 'remove', 'settings', 'award', 'billing', 'release'] : [],
      },
    ],
  };
}

/** The mock home as a guild with nothing logged, its viewer at `level`. */
function zeroLogHome(viewer: string, level: number): GuildHome {
  const home = buildMockHome(viewer);
  const { summary, members } = home;
  if (summary === undefined || members === undefined)
    throw new Error('the mock home carries a summary and members');
  const levels = [level, 60, 52, 41, 23, 8];
  return {
    ...home,
    standing: undefined,
    summary: {
      ...summary,
      tier: { ...summary.tier, down: 0, pulls_total: 0 },
      named_encounters_down: 0,
      pulls_this_tier: 0,
      waiting_for_approval: 0,
    },
    roster: home.roster.map((row) => ({
      ...row,
      level: row.name === viewer ? level : levels[1 + (row.name.length % 5)],
      professions: row.name === viewer ? ['alchemy', 'cooking'] : row.professions,
    })),
    members: members.map((entry) => ({
      ...entry,
      level: entry.name === viewer ? level : levels[1 + (entry.name.length % 5)],
    })),
    reports: [],
    next_raid: undefined,
  };
}

function zeroLogProgression(): GuildProgressionPage {
  const page = buildMockProgression();
  return {
    tier: { ...page.tier, down: 0, pulls_total: 0 },
    encounters: page.encounters.map((encounter) => {
      const { first_kill_at: _firstKill, best_kill_ms: _bestKill, ...rest } = encounter;
      return {
        ...rest,
        pulls: 0,
        kills: 0,
        pulls_by_night: [],
        deaths_per_pull: 0,
        best_by_role: {},
      };
    }),
    unnamed: page.unnamed.map((zone) => ({ ...zone, pulls: 0, nights: 0 })),
  };
}

function leveling(row: GuildReadinessRow, level: number): GuildReadinessRow {
  return {
    ...row,
    level,
    gear_gap: null,
    enchants: { missing_slots: [], checked: false },
    consumables: { state: 'unknown' },
    talent_points_unspent: 0,
    item_level_delta: null,
    failing: 0,
    nudge_text: undefined,
  };
}

/** The mock board with its last row at 58 and its first a healer's gap. */
function mixedBoard(): GuildReadinessPage {
  const board = buildMockReadiness();
  const [first, ...rest] = board.rows;
  const healer: GuildReadinessRow = {
    ...first,
    gear_gap: { upgrades: 4, gain_dps: 31, not_sim_checked: 2, unit: null },
    nudge_text: `${first.name}: 4 gear upgrades waiting.`,
  };
  const last = rest.at(-1);
  return {
    ...board,
    rows: last === undefined ? [healer] : [healer, ...rest.slice(0, -1), leveling(last, 58)],
  };
}

interface Stub {
  viewer: string;
  officer: boolean;
  home: GuildHome;
  progression: GuildProgressionPage;
  readiness: GuildReadinessPage;
}

async function stub(page: Page, s: Stub): Promise<void> {
  await page.route('**/v1/guilds/us/pvp/olympus-xxvii', (route) =>
    route.fulfill(envelope(buildMockGuildPage())),
  );
  await page.route('**/v1/me', (route) => route.fulfill(envelope(meFor(s.viewer, s.officer))));
  await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) => route.fulfill(envelope(s.home)));
  await page.route('**/v1/guilds/*/progression*', (route) => route.fulfill(envelope(s.progression)));
  await page.route('**/v1/guilds/*/readiness*', (route) => route.fulfill(envelope(s.readiness)));
  await page.route('**/v1/rankings/guilds*', (route) => route.fulfill(envelope({ rows: [] })));
}

async function capture(page: Page, name: string): Promise<void> {
  if (SHOTS === undefined) return;
  const width = page.viewportSize()?.width ?? 0;
  await page.screenshot({ path: `${SHOTS}/${name}-${width}.png`, fullPage: true });
}

test.describe('zero logs', () => {
  test('the band, Overview and Progression all read no pulls, and the card leads', async ({ page }) => {
    await stub(page, {
      viewer: VIEWER_MEMBER,
      officer: false,
      home: zeroLogHome(VIEWER_MEMBER, 52),
      progression: zeroLogProgression(),
      readiness: mixedBoard(),
    });
    await page.goto(GUILD_PATH);

    const card = page.getByTestId('guild-overview-card-week-one');
    await expect(card).toBeVisible();
    await expect(card).toContainText('Before raids open');
    await expect(card).toContainText('AT 60');
    await expect(page.getByTestId('guild-week-one-raids-line')).toContainText('open Wed 9 Dec');
    await expect(page.getByTestId('guild-week-one-sentence')).toHaveText(
      `${VIEWER_MEMBER} is level 52, 8 levels to 60.`,
    );
    await expect(page.getByTestId('guild-week-one-action')).toHaveCount(1);
    await expect(page.getByTestId('guild-week-one-action')).toContainText('Level as');

    const facts = page.getByTestId('guild-header-facts');
    await expect(facts).not.toContainText('bosses');
    await expect(facts).not.toContainText('pull');
    await expect(page.getByTestId('guild-standing-leveling')).toHaveText(
      `${VIEWER_MEMBER} is level 52. Your guild rank by item level starts at 60.`,
    );
    await expect(page.getByTestId('guild-overview-progression-line')).toHaveText(
      'No boss down this tier yet.',
    );
    await capture(page, 'overview-raider-52');

    await page.getByRole('tab', { name: 'Progression' }).click();
    const empty = page.getByTestId('guild-progression-empty');
    await expect(empty).toContainText('Progression starts with your first logged pull.');
    await expect(page.getByTestId('guild-progression-tier')).toHaveCount(0);
    await expect(page.locator('body')).not.toContainText('NaN');
    await capture(page, 'progression-zero');
  });

  test('an officer with nothing planned is pointed at the calendar', async ({ page }) => {
    await stub(page, {
      viewer: VIEWER_OFFICER,
      officer: true,
      home: zeroLogHome(VIEWER_OFFICER, 60),
      progression: zeroLogProgression(),
      readiness: mixedBoard(),
    });
    await page.goto(GUILD_PATH);
    await expect(page.getByTestId('guild-week-one-sentence')).toHaveText(
      'Nothing on the calendar yet. Post a dungeon run or a guild event.',
    );
    await expect(page.getByTestId('guild-week-one-invite')).toContainText('Copy the invite link');
    await capture(page, 'overview-officer-no-event');
    await page.route('**/v1/guilds/*/raids*', (route) =>
      route.fulfill(envelope({ ...buildMockRaids(), planned: [], zones: [] })),
    );
    await page.getByTestId('guild-week-one-action').click();
    await expect(page.getByTestId('guild-raids-tab')).toBeVisible();
    await expect(page.locator('[data-testid^="raid-form-"]').first()).toBeVisible();
  });
});

test.describe('after a kill', () => {
  test('the card is gone and the three surfaces print one set of numbers', async ({ page }) => {
    const progression = buildMockProgression();
    const home = buildMockHome(VIEWER_MEMBER);
    await stub(page, {
      viewer: VIEWER_MEMBER,
      officer: false,
      home,
      progression,
      readiness: buildMockReadiness(),
    });
    await page.goto(GUILD_PATH);

    await expect(page.getByTestId('guild-overview-tab')).toBeVisible();
    await expect(page.getByTestId('guild-overview-card-week-one')).toHaveCount(0);
    const { down, named_encounters: named, pulls_total: pulls } = progression.tier;
    await expect(page.getByTestId('guild-header-facts')).toContainText(`${down} boss down`);
    await expect(page.getByTestId('guild-header-facts')).toContainText(`${pulls} pulls`);
    await expect(page.getByTestId('guild-overview-progression-line')).toHaveText(
      `${down} of ${named} named bosses down: Onyxia.`,
    );
    await capture(page, 'overview-after-kill');

    await page.getByRole('tab', { name: 'Progression' }).click();
    await expect(page.getByTestId('guild-progression-down')).toHaveText(`${down} of ${named}`);
    await expect(page.locator('body')).not.toContainText('NaN');
    await capture(page, 'progression-one-kill');
  });
});

test.describe('readiness', () => {
  test('a level 58 row has no Nudge or gear figure, and a healer row never says DPS', async ({ page }) => {
    const board = mixedBoard();
    const low = board.rows.at(-1);
    const healer = board.rows[0];
    if (low === undefined) throw new Error('fixture board is empty');
    await stub(page, {
      viewer: VIEWER_OFFICER,
      officer: true,
      home: buildMockHome(VIEWER_OFFICER),
      progression: buildMockProgression(),
      readiness: board,
    });
    await page.goto(`${GUILD_PATH}#readiness`);

    const lowRow = page.getByTestId(`guild-readiness-row-${low.character_key}`).locator('visible=true');
    await expect(lowRow).toContainText('Level 58 · 2 levels to 60');
    await expect(lowRow).not.toContainText('upgrade');
    await expect(page.getByTestId(`guild-readiness-nudge-${low.character_key}`)).toHaveCount(0);
    await expect(page.getByTestId('guild-readiness-leveling-group')).toBeVisible();

    const healerRow = page.getByTestId(`guild-readiness-row-${healer.character_key}`).locator('visible=true');
    await expect(healerRow).toContainText('4 upgrades');
    await expect(healerRow).not.toContainText('DPS');
    await capture(page, 'readiness-mixed');
  });
});
