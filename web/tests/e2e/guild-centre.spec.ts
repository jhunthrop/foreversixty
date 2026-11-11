// web/tests/e2e/guild-centre.spec.ts
// Guild control-centre spec v2 (design/specs/2026-10-04-guild-page.md) acceptance e2e.
// Routes every one of the five control-centre endpoints to
// `web/src/fixtures/guild/mock-guild.ts` -- the same seeded mock the acceptance boards
// (design/mocks/gen_guild.py) render from -- for public, member and officer viewers, and
// asserts the hard rules the build brief names: tabs render and hash-route; the roster
// pins the viewer's own row and sorts; an officer sees Approve all and a Nudge that copies
// real text; a member never sees an officer control; Loot shows the awarded state; the
// Readiness tab sorts worst-first; and nothing scrolls sideways at 390 or 2000px. This
// supersedes guild-home.spec.ts and guild-phone.spec.ts, both of which exercised exactly
// the flat single-page body (the duplicated unverified note, the old flat roster list) this
// rebuild replaces -- see the build report for the full reasoning.
import type { Locator, Page } from '@playwright/test';
import { apiUnreachable, expect, test, GUILD_OVERVIEW_READS } from './support/test';
import {
  GUILD_ID,
  VIEWER_MEMBER,
  VIEWER_OFFICER,
  buildMockGuildPage,
  buildMockHome,
  buildMockLoot,
  buildMockProgression,
  buildMockReadiness,
  buildMockRaids,
} from '../../src/fixtures/guild/mock-guild';
import { guildBillingRead } from '../../src/test-support/billing-fixtures';
import { stubLootCouncil } from './support/officer-tools';
import { openRosterAccounts } from './support/guild-roster-accounts';

test.beforeEach(({ page }) => apiUnreachable(page, ...GUILD_OVERVIEW_READS));

// Reads this spec does not answer go unanswered, as the API out of reach (support/test.ts).
test.beforeEach(({ page }) =>
  apiUnreachable(
    page,
    '**/v1/guilds/*/loot*',
    '**/v1/guilds/*/progression*',
    '**/v1/guilds/*/raids*',
    '**/v1/guilds/*/readiness*',
    '**/v1/guilds/*/settings',
  ),
);

// Not imported from lib/report/shell-paths.ts: that module reaches lib/rankings/api.ts,
// which reads `import.meta.env.PUBLIC_API_BASE_URL` (lib/planner/config.ts) -- fine for
// every in-app import (Vite supplies it), but Playwright's own Node runtime has no
// `import.meta.env` and throws on load (the same class of import sim-tabs.spec.ts's own
// header comment avoids). Both literals must match src/lib/report/shell-paths.ts exactly.
const FIXTURE_GUILD_CENTRE_PATH = 'us/pvp/olympus-xxvii';
const FIXTURE_GUILD_PATH = 'us/hardcore/the-last-watch';

function envelope(data: unknown, status = 200) {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'req-test' }),
  };
}

const [REGION, RULESET, SLUG] = FIXTURE_GUILD_CENTRE_PATH.split('/');
const GUILD_URL = `/guild/${FIXTURE_GUILD_CENTRE_PATH}`;

/** The capability words the API lists on a /v1/me guild entry (api guilds.Capabilities): the
 *  guild master holding the claim has all six, an officer four, a member none. */
function canFor(rank: 'leader' | 'officer' | 'member'): string[] {
  if (rank === 'leader') return ['approve', 'remove', 'settings', 'award', 'billing', 'release'];
  if (rank === 'officer') return ['approve', 'remove', 'settings', 'award'];
  return [];
}

function meFixture(
  viewerName: string | null,
  rank: 'leader' | 'officer' | 'member' = 'member',
  battletag = 'Fixture#1234',
) {
  return {
    user: { id: 7, battletag, email: null, role: 'user', anonymize: false, premium: false },
    characters:
      viewerName === null
        ? []
        : [
            {
              key: `us/pvp/${viewerName.toLowerCase().replace(/[^a-z0-9]+/g, '-')}`,
              region: 'us',
              ruleset: 'pvp',
              name: viewerName,
              class: 'warrior',
            },
          ],
    guilds:
      viewerName === null
        ? []
        : [
            {
              id: GUILD_ID,
              region: REGION,
              ruleset: RULESET,
              name: 'Olympus XXVII',
              rank,
              verified: true,
              can: canFor(rank),
            },
          ],
  };
}

/** The home as the API answers it once `name` has left the waiting list (approved or declined): the
 *  page re-reads the home after a mutation, so a stub has to change its answer too. */
function withoutRow(home: ReturnType<typeof buildMockHome>, name: string) {
  return {
    ...home,
    pending: (home.pending ?? []).filter((r) => r.name !== name),
    roster: home.roster.map((r) => (r.name === name ? { ...r, verified: true } : r)),
  };
}

/** The home once every waiting character is approved. */
function approvedHome(viewer: string) {
  const home = buildMockHome(viewer);
  return {
    ...home,
    pending: [],
    roster: home.roster.map((r) => ({ ...r, verified: true })),
    members: home.members?.map((m) => (m.state === 'waiting' ? { ...m, state: 'approved' as const } : m)),
    counts: home.counts && { ...home.counts, waiting: 0 },
  };
}

/** Routes the public page endpoint plus /v1/me and .../home for the given viewer; the four
 *  depth endpoints (raids/progression/readiness/loot) are routed separately per test so a
 *  test that wants a 404 from one of them can override just that route. */
async function stubCore(
  page: Page,
  viewerName: string | null,
  rank?: 'leader' | 'officer' | 'member',
  battletag?: string,
): Promise<void> {
  await page.route(`**/v1/guilds/${REGION}/${RULESET}/${SLUG}`, (route) =>
    route.fulfill(envelope(buildMockGuildPage())),
  );
  await page.route('**/v1/me', (route) =>
    viewerName === null
      ? route.fulfill({
          status: 401,
          contentType: 'application/json',
          body: '{"ok":false,"data":null,"error":null,"request_id":"r"}',
        })
      : route.fulfill(envelope(meFixture(viewerName, rank, battletag))),
  );
  await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) =>
    route.fulfill(envelope(buildMockHome(viewerName))),
  );
}

async function stubDepthEndpoints(page: Page): Promise<void> {
  await page.route(`**/v1/guilds/${GUILD_ID}/raids*`, (route) => route.fulfill(envelope(buildMockRaids())));
  await page.route(`**/v1/guilds/${GUILD_ID}/progression`, (route) =>
    route.fulfill(envelope(buildMockProgression())),
  );
  await page.route(`**/v1/guilds/${GUILD_ID}/readiness`, (route) =>
    route.fulfill(envelope(buildMockReadiness())),
  );
  await page.route(`**/v1/guilds/${GUILD_ID}/loot*`, (route) => route.fulfill(envelope(buildMockLoot())));
  await stubLootCouncil(page, 'officer');
}

/** Same as `stubCore`, except the public/home endpoints carry the given `faction` --
 *  `stubCore`'s own fixture is always Horde (`GUILD.faction`'s default); these three header-
 *  art tests (spec §12.2) need to override it to Alliance and to null/neutral as well. */
async function stubCoreWithFaction(
  page: Page,
  viewerName: string | null,
  faction: 'alliance' | 'horde' | null,
): Promise<void> {
  await page.route(`**/v1/guilds/${REGION}/${RULESET}/${SLUG}`, (route) =>
    route.fulfill(envelope(buildMockGuildPage(faction))),
  );
  await page.route('**/v1/me', (route) =>
    viewerName === null
      ? route.fulfill({
          status: 401,
          contentType: 'application/json',
          body: '{"ok":false,"data":null,"error":null,"request_id":"r"}',
        })
      : route.fulfill(envelope(meFixture(viewerName))),
  );
  await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) =>
    route.fulfill(envelope(buildMockHome(viewerName, faction))),
  );
}

// Header art round (design/specs/2026-10-04-guild-page.md §12.2, option B, the owner's
// pick): the flat faction logo is the one emblem in the header, ringed beside the h1
// (FactionCrest) and repeated as the band's own cropped watermark over a diagonal faction-
// colour wash. A null/unknown faction renders neither layer -- a neutral band, pixel-
// identical to the page before this round.
test.describe('header art: faction crest, watermark and neutral band (spec §12.2)', () => {
  test('a Horde guild shows the ringed flat Horde logo', async ({ page }) => {
    await stubCoreWithFaction(page, VIEWER_MEMBER, 'horde');
    await page.goto(GUILD_URL);
    const crest = page.getByTestId('guild-faction-crest');
    await expect(crest).toBeVisible();
    await expect(crest).toHaveAttribute('src', /horde-logo-512\.webp$/);
    // The browser re-serialises the style attribute, normalising the authored hex colour to
    // rgb() -- #c0392b is rgb(192, 57, 43).
    await expect(crest).toHaveAttribute('style', /box-shadow: rgb\(192, 57, 43\) 0px 0px 0px 2px/);
  });

  test('an Alliance guild shows the Alliance logo and Alliance bar colour, never Horde’s', async ({
    page,
  }) => {
    await stubCoreWithFaction(page, VIEWER_MEMBER, 'alliance');
    await page.goto(GUILD_URL);
    const crest = page.getByTestId('guild-faction-crest');
    await expect(crest).toHaveAttribute('src', /alliance-logo-512\.webp$/);
    await expect(crest).toHaveAttribute('style', /box-shadow: rgb\(47, 111, 214\) 0px 0px 0px 2px/);
  });

  test('a null-faction guild renders no crest', async ({ page }) => {
    await stubCoreWithFaction(page, VIEWER_MEMBER, null);
    await page.goto(GUILD_URL);
    await expect(page.getByTestId('guild')).toBeVisible();
    await expect(page.getByTestId('guild-faction-crest')).toHaveCount(0);
  });

  // Site alignment D1: no band, faction wash or watermark behind the head, at any width.
  for (const width of [1440, 390]) {
    test(`the head has no band, wash or watermark at ${width}px`, async ({ page }) => {
      await page.setViewportSize({ width, height: 900 });
      await stubCoreWithFaction(page, VIEWER_MEMBER, 'alliance');
      await page.goto(GUILD_URL);
      await expect(page.getByTestId('guild-header')).toBeVisible();
      for (const gone of ['guild-header-band', 'guild-header-vignette', 'guild-header-watermark']) {
        await expect(page.getByTestId(gone)).toHaveCount(0);
      }
    });
  }
});

test.describe('tab strip renders and hash-routes', () => {
  test('officer sees all seven tabs; selecting one updates the hash and shows that panel', async ({
    page,
  }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await stubDepthEndpoints(page);
    await page.goto(GUILD_URL);
    const tabs = page.getByTestId('guild-tabs');
    await expect(tabs).toBeVisible();
    for (const id of ['overview', 'roster', 'raids', 'progression', 'readiness', 'loot', 'settings']) {
      await expect(page.getByTestId(`guild-tab-${id}`)).toBeVisible();
    }
    await page.getByTestId('guild-tab-roster').click();
    await expect(page).toHaveURL(/#roster$/);
    await openRosterAccounts(page);
    await expect(page.getByTestId('guild-roster-tab')).toBeVisible();

    await page.getByTestId('guild-tab-readiness').click();
    await expect(page).toHaveURL(/#readiness$/);
    await expect(page.getByTestId('guild-readiness-tab')).toBeVisible();
  });

  // Site alignment §5.10: the claim and invite strip left the head. Its jobs live in the Settings tab
  // (invite link) and the Roster tab (waiting for approval), so no tab carries it.
  test('the head carries no officer strip on any tab', async ({ page }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await stubDepthEndpoints(page);
    await page.goto(GUILD_URL);
    await expect(page.getByTestId('guild-header')).toBeVisible();
    await expect(page.getByTestId('guild-officer-strip')).toHaveCount(0);
    await page.getByTestId('guild-tab-readiness').click();
    await expect(page.getByTestId('guild-officer-strip')).toHaveCount(0);
  });

  test('a public visitor sees only the three public tabs, never Roster/Readiness/Loot/Settings', async ({
    page,
  }) => {
    await stubCore(page, null);
    await page.goto(GUILD_URL);
    await expect(page.getByTestId('guild-tab-overview')).toBeVisible();
    await expect(page.getByTestId('guild-tab-raids')).toBeVisible();
    await expect(page.getByTestId('guild-tab-progression')).toBeVisible();
    await expect(page.getByTestId('guild-tab-roster')).toHaveCount(0);
    await expect(page.getByTestId('guild-tab-readiness')).toHaveCount(0);
    await expect(page.getByTestId('guild-tab-loot')).toHaveCount(0);
    await expect(page.getByTestId('guild-tab-settings')).toHaveCount(0);
  });

  test('a stale #roster hash on a public visit silently falls back to Overview, never a crash', async ({
    page,
  }) => {
    await stubCore(page, null);
    await page.goto(`${GUILD_URL}#roster`);
    await openRosterAccounts(page);
    await expect(page.getByTestId('guild-overview-tab')).toBeVisible();
    await expect(page.getByTestId('guild-roster-tab')).toHaveCount(0);
  });
});

test.describe('roster: pin, sort, officer controls', () => {
  test('the officer roster pins the viewer’s own row first after sort, and leads with unverified rows', async ({
    page,
  }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await page.goto(`${GUILD_URL}#roster`);
    await openRosterAccounts(page);
    const rows = page.getByTestId('guild-roster-rows').locator('.guild-roster-row');
    await expect(rows.first()).toContainText('(you)');

    await expect(page.getByTestId('guild-roster-waiting')).toContainText('Waiting for approval');
  });

  test('officer sees Approve all, with a distinct accessible name per raider for Approve/Remove', async ({
    page,
  }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await page.goto(`${GUILD_URL}#roster`);
    await openRosterAccounts(page);
    await expect(page.getByTestId('guild-roster-approve-all')).toBeVisible();
    const approveButtons = page.getByTestId('guild-roster-approve');
    const count = await approveButtons.count();
    expect(count).toBeGreaterThan(0);
    const names = await approveButtons.evaluateAll((buttons) =>
      buttons.map((b) => b.getAttribute('aria-label')),
    );
    expect(new Set(names).size).toBe(names.length);
  });

  test('a member never sees Approve, Approve all, or Remove', async ({ page }) => {
    await stubCore(page, VIEWER_MEMBER, 'member');
    await page.goto(`${GUILD_URL}#roster`);
    await openRosterAccounts(page);
    await expect(page.getByTestId('guild-roster-tab')).toBeVisible();
    await expect(page.getByTestId('guild-roster-approve')).toHaveCount(0);
    await expect(page.getByTestId('guild-roster-approve-all')).toHaveCount(0);
    await expect(page.getByTestId('guild-roster-remove')).toHaveCount(0);
  });
});

// Task L8.3: the Roster tab's "In game" column, from the guild's newest addon roster
// snapshot (Task L6.14). The fixture's snapshot came from Zulmara at 21:43 UTC on 22 Dec; the
// page's clock is pinned 3 minutes later so the line reads the same on every run.
test.describe('roster: the In game column from the addon snapshot (L8.3)', () => {
  // The fixture's "seen" times are dated 22 Dec; "today" is the viewer's local date.
  test.use({ timezoneId: 'UTC' });

  async function openRoster(page: Page, withSnapshot: boolean): Promise<void> {
    await page.clock.setFixedTime(new Date('2026-12-22T21:46:00Z'));
    await stubCore(page, VIEWER_MEMBER, 'member');
    await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) =>
      route.fulfill(envelope(buildMockHome(VIEWER_MEMBER, undefined, null, withSnapshot))),
    );
    await page.goto(`${GUILD_URL}#roster`);
    await openRosterAccounts(page);
    await expect(page.getByTestId('guild-roster-rows')).toBeVisible();
  }

  function rowNamed(page: Page, name: string) {
    return page
      .getByTestId('guild-roster-rows')
      .locator('.guild-roster-row')
      .filter({ has: page.getByTestId('guild-roster-name').getByText(name, { exact: false }) });
  }

  test('matched rows read level, spec, gap, build and seen time; an unmatched row reads nothing', async ({
    page,
  }) => {
    await openRoster(page, true);
    // No explaining line above the table (batch 3).
    await expect(page.getByTestId('guild-roster-addon-snapshot')).toHaveCount(0);
    await expect(rowNamed(page, 'Zulmara').getByTestId('guild-roster-ingame')).toHaveText(
      /^Level 60\s+·\s+Gear gap 12 DPS\s+·\s+Gear set: Survival \(raid\)\s+·\s+seen 6 minutes ago\s*$/,
    );
    await expect(rowNamed(page, 'Kraggor').getByTestId('guild-roster-ingame')).toHaveText(
      /^Level 60\s+·\s+Gear gap 4 DPS\s+·\s+Gear set: Protection \(raid\)\s+·\s+seen 8 minutes ago\s*$/,
    );
    await expect(rowNamed(page, 'Windtalon').getByTestId('guild-roster-ingame')).toHaveCount(0);
    await expect(rowNamed(page, 'Windtalon')).not.toContainText('—');
    // The spec line leads with the addon's spec, and names the site's only where it differs
    // (no "in game" word since d42bf7e65).
    await expect(rowNamed(page, 'Zulmara').getByTestId('guild-roster-spec')).toContainText(
      'Survival Hunter · Marksmanship on the site',
    );
    await expect(rowNamed(page, 'Kraggor').getByTestId('guild-roster-spec')).toHaveText(
      'Protection Warrior · Guild master',
    );
  });

  test('the roster has a SORT row of chips and no header row at any width; the chips sort', async ({
    page,
  }) => {
    await openRoster(page, true);
    // The account table (the raid roster above it is a StatTable with its own head).
    await expect(page.getByTestId('guild-roster-tab').locator('[role="columnheader"]')).toHaveCount(0);
    await expect(page.getByTestId('guild-roster-sort-header')).toHaveCount(0);
    const sort = page.getByTestId('guild-roster-sort');
    await expect(sort).toContainText('Sort');
    await expect(sort.getByRole('button')).toHaveText([
      /Rank/,
      /Item level/,
      /Rating/,
      /Attendance/,
      /Last seen/,
    ]);
    await expect(page.getByTestId('guild-roster-sort-ilvl')).toContainText('▾');
    await expect(page.getByTestId('guild-roster-sort-ilvl')).toHaveAttribute('aria-pressed', 'true');
    // Rank puts the leader and officers first (the viewer stays pinned above them).
    await page.getByTestId('guild-roster-sort-rank').click();
    await expect(page.getByTestId('guild-roster-sort-rank')).toContainText('▾');
    await expect(page.getByTestId('guild-roster-sort-rank')).toHaveAttribute('aria-pressed', 'true');
    await expect(page.getByTestId('guild-roster-sort-ilvl')).not.toHaveAttribute('aria-pressed', 'true');
    const names = (
      await page.getByTestId('guild-roster-rows').getByTestId('guild-roster-name').allTextContents()
    ).map((n) => n.replace(' (you)', '').trim());
    expect(names[0]).toBe('Zulmara');
    // The In game column's gap is a fact on the row, not a sort the page offers.
    await expect(page.getByTestId('guild-roster-sort-addon_gap')).toHaveCount(0);
    // Every row's figure columns end at one right edge at lg, under the heads.
    if ((page.viewportSize()?.width ?? 0) >= 1024) {
      const rights = await page
        .getByTestId('guild-roster-rows')
        .getByTestId('guild-roster-columns')
        .evaluateAll((els) => els.map((e) => Math.round(e.getBoundingClientRect().right)));
      expect(new Set(rights).size).toBe(1);
      await expect(page.getByTestId('guild-roster-heads')).toContainText('Item level');
    }
  });

  test('with no snapshot: no addon line, no In game slot, no dashes', async ({ page }) => {
    await openRoster(page, false);
    await expect(page.getByTestId('guild-roster-addon-snapshot')).toHaveCount(0);
    await expect(page.getByTestId('guild-roster-ingame')).toHaveCount(0);
    await expect(page.getByTestId('guild-roster-ingame-slot')).toHaveCount(0);
    await expect(page.getByTestId('guild-roster-rows')).not.toContainText('—');
  });
});

// Lifecycle round 4, the roster model: everyone the addon saw with a word beside each, what
// approving does, and controls that come from the account's capability list.
test.describe('roster: everyone the addon saw, with a word beside each', () => {
  test.use({ timezoneId: 'UTC' });

  test('the count line, a rank tag for officers alone, and the names with no account', async ({ page }) => {
    await page.clock.setFixedTime(new Date('2026-12-22T21:46:00Z'));
    await stubCore(page, VIEWER_MEMBER, 'member');
    await page.goto(`${GUILD_URL}#roster`);
    await openRosterAccounts(page);
    const home = buildMockHome(VIEWER_MEMBER);
    const waiting = home.roster.filter((r) => !r.verified).length;
    await expect(page.getByTestId('guild-roster-counts')).toHaveText(
      `${home.roster.length + 2} on the roster · 1 in game without an account · ${waiting} waiting for approval · 1 not seen in game yet`,
    );
    await expect(page.getByTestId('guild-roster-gear-note')).toHaveCount(0);
    // A plain member carries no rank tag (batch 3): only the guild master and officers are tagged.
    const zulmara = page
      .getByTestId('guild-roster-rows')
      .locator('.guild-roster-row')
      .filter({ has: page.getByTestId('guild-roster-name').getByText('Zulmara') });
    await expect(zulmara.getByTestId('guild-roster-rank')).toHaveCount(0);
    // The names with no account row: a state word each, no controls.
    const entries = page.getByTestId('guild-roster-entries').getByTestId('guild-roster-entry');
    await expect(entries).toHaveCount(2);
    await expect(entries.nth(0)).toContainText('Brightwick');
    await expect(entries.nth(0).getByTestId('guild-roster-entry-state')).toHaveText('No account yet');
    await expect(entries.nth(0).getByTestId('guild-roster-entry-ingame')).toContainText('Level 58');
    await expect(entries.nth(0).getByTestId('guild-roster-entry-ingame')).toContainText('seen 2 hours ago');
    await expect(entries.nth(1)).toContainText('Driftwood');
    await expect(entries.nth(1).getByTestId('guild-roster-entry-state')).toHaveText('Not seen in game yet');
    await expect(entries.locator('button')).toHaveCount(0);
  });

  test('the "No account" chip narrows the roster to the names with no account', async ({ page }) => {
    await stubCore(page, VIEWER_MEMBER, 'member');
    await page.goto(`${GUILD_URL}#roster`);
    await openRosterAccounts(page);
    await page.getByTestId('guild-roster-filter-noaccount').click();
    await expect(page.getByTestId('guild-roster-entries').getByTestId('guild-roster-entry')).toHaveCount(1);
    await expect(page.getByTestId('guild-roster-rows')).toHaveCount(0);
    await page.getByTestId('guild-roster-filter-noaccount').click();
    await expect(page.getByTestId('guild-roster-rows')).toBeVisible();
    // "Verified only" drops the name with no account and keeps the invited member.
    await page.getByTestId('guild-roster-filter-verified').click();
    await expect(page.getByTestId('guild-roster-entries').getByTestId('guild-roster-entry')).toHaveCount(1);
    await expect(page.getByTestId('guild-roster-entries')).toContainText('Driftwood');
  });

  test('the waiting list says what Approve does and shows the rank each character arrived with', async ({
    page,
  }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await page.goto(`${GUILD_URL}#roster`);
    await openRosterAccounts(page);
    await expect(page.getByTestId('guild-roster-approve-explains')).toHaveCount(0);
    const list = page.getByTestId('guild-roster-unverified-list');
    // A member waits with no rank tag (batch 3).
    await expect(list.getByTestId('guild-roster-pending-rank')).toHaveCount(0);
    // The card's head says it; there is no per-row state tag.
    await expect(page.getByTestId('guild-roster-waiting')).toContainText('Waiting for approval');
    await expect(list.getByTestId('guild-roster-state')).toHaveCount(0);
  });

  test('Approve all goes straight through when no officer is waiting', async ({ page }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    let called = false;
    await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) =>
      route.fulfill(envelope(called ? approvedHome(VIEWER_OFFICER) : buildMockHome(VIEWER_OFFICER))),
    );
    await page.route(`**/v1/guilds/${GUILD_ID}/roster/approve-all`, (route) => {
      called = true;
      return route.fulfill(
        envelope({ approved: ['us/pvp/rootgall', 'us/pvp/stormtusk', 'us/pvp/fulmintide'] }),
      );
    });
    await page.goto(`${GUILD_URL}#roster`);
    await openRosterAccounts(page);
    await page.getByTestId('guild-roster-approve-all').click();
    await expect(page.getByTestId('guild-roster-approve-all-confirm')).toHaveCount(0);
    await expect.poll(() => called).toBe(true);
    await expect(page.getByTestId('guild-roster-unverified-list')).toHaveCount(0);
    // The count line and the list agree in the same frame.
    await expect(page.getByTestId('guild-roster-counts')).not.toContainText('waiting for approval');
  });

  test('Approve all asks first when an officer is among the waiting', async ({ page }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    const home = buildMockHome(VIEWER_OFFICER);
    const waiting = home.roster.filter((r) => !r.verified);
    const withOfficer = {
      ...home,
      roster: home.roster.map((r) => (r.name === 'Rootgall' ? { ...r, rank: 'officer' as const } : r)),
      pending: waiting.map((r) => (r.name === 'Rootgall' ? { ...r, rank: 'officer' as const } : r)),
    };
    await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) => route.fulfill(envelope(withOfficer)));
    let called = false;
    await page.route(`**/v1/guilds/${GUILD_ID}/roster/approve-all`, (route) => {
      called = true;
      return route.fulfill(envelope({ approved: waiting.map((r) => r.character_key) }));
    });
    await page.goto(`${GUILD_URL}#roster`);
    await openRosterAccounts(page);
    await page.getByTestId('guild-roster-approve-all').click();
    await expect(page.getByTestId('guild-roster-approve-all-confirm')).toContainText(
      `Approve ${waiting.length}, including 1 officer?`,
    );
    expect(called).toBe(false);
    await page
      .getByTestId('guild-roster-approve-all-confirm')
      .getByRole('button', { name: 'Cancel' })
      .click();
    expect(called).toBe(false);
    await page.getByTestId('guild-roster-approve-all').click();
    await page.getByTestId('guild-roster-approve-all-yes').click();
    await expect.poll(() => called).toBe(true);
  });

  test('Decline asks first, then removes the character from the waiting list', async ({ page }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    let removed = false;
    await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) =>
      route.fulfill(
        envelope(
          removed ? withoutRow(buildMockHome(VIEWER_OFFICER), 'Rootgall') : buildMockHome(VIEWER_OFFICER),
        ),
      ),
    );
    await page.route('**/v1/guilds/*/characters/us/pvp/rootgall', (route) => {
      if (route.request().method() === 'DELETE') {
        removed = true;
        return route.fulfill(envelope({ status: 'removed' }));
      }
      return route.continue();
    });
    await page.goto(`${GUILD_URL}#roster`);
    await openRosterAccounts(page);
    await page.locator('[data-testid="guild-roster-decline"][aria-label="Decline Rootgall"]').click();
    await expect(page.getByTestId('guild-roster-decline-confirm')).toContainText(
      'Remove Rootgall from the waiting list? They can ask again.',
    );
    expect(removed).toBe(false);
    await page.getByTestId('guild-roster-decline-yes').click();
    await expect.poll(() => removed).toBe(true);
    await expect(page.getByTestId('guild-roster-unverified-list')).not.toContainText('Rootgall');
  });

  test('controls come from the capability list: an officer whose list lacks remove and approve sees neither', async ({
    page,
  }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await page.route('**/v1/me', (route) =>
      route.fulfill(
        envelope({
          ...meFixture(VIEWER_OFFICER, 'leader'),
          guilds: [
            {
              id: GUILD_ID,
              region: REGION,
              ruleset: RULESET,
              name: 'Olympus XXVII',
              rank: 'leader',
              verified: true,
              can: [],
            },
          ],
        }),
      ),
    );
    await page.goto(`${GUILD_URL}#roster`);
    await openRosterAccounts(page);
    await expect(page.getByTestId('guild-roster-tab')).toBeVisible();
    await expect(page.getByTestId('guild-roster-approve')).toHaveCount(0);
    await expect(page.getByTestId('guild-roster-approve-all')).toHaveCount(0);
    await expect(page.getByTestId('guild-roster-remove')).toHaveCount(0);
    await expect(page.getByTestId('guild-roster-decline')).toHaveCount(0);
  });

  test('nobody has Remove on their own row: an officer, and the guild master, keep it on everyone else', async ({
    page,
  }) => {
    // The API marks the viewer\'s own officer row may_remove: false (guilds.mayRemoveRow).
    await stubCore(page, 'Sunderfel', 'officer');
    const home = buildMockHome('Sunderfel');
    await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) =>
      route.fulfill(
        envelope({
          ...home,
          roster: home.roster.map((r) => (r.name === 'Sunderfel' ? { ...r, may_remove: false } : r)),
        }),
      ),
    );
    await page.goto(`${GUILD_URL}#roster`);
    await openRosterAccounts(page);
    const own = page
      .getByTestId('guild-roster-rows')
      .locator('.guild-roster-row')
      .filter({ has: page.getByTestId('guild-roster-name').getByText('Sunderfel') });
    await expect(own).toBeVisible();
    await expect(own.getByTestId('guild-roster-remove')).toHaveCount(0);
    // Others' rows still carry Remove for an officer.
    await expect(
      page.locator('[data-testid="guild-roster-remove"][aria-label="Remove Zulmara"]'),
    ).toBeVisible();
  });

  test('the guild master has no Remove on their own row, and a name only the addon saw has no rank word', async ({
    page,
  }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await page.goto(`${GUILD_URL}#roster`);
    await openRosterAccounts(page);
    const own = page
      .getByTestId('guild-roster-rows')
      .locator('.guild-roster-row')
      .filter({ has: page.getByTestId('guild-roster-name').getByText(VIEWER_OFFICER) });
    await expect(own).toBeVisible();
    await expect(own.getByTestId('guild-roster-rank')).toHaveText('Guild master');
    await expect(own.getByTestId('guild-roster-remove')).toHaveCount(0);
    await expect(
      page.locator('[data-testid="guild-roster-remove"][aria-label="Remove Zulmara"]'),
    ).toBeVisible();
    // Brightwick is in game with no account: the state word alone, no "Member" invented for it.
    const noAccount = page.locator('[data-testid="guild-roster-entry"][data-state="no_account"]');
    await expect(noAccount).toContainText('No account yet');
    await expect(noAccount).not.toContainText('Member');
    await expect(noAccount.getByTestId('guild-roster-entry-rank')).toHaveCount(0);
  });

  test('at 390px a no-account row stacks its state and in-game cells under the name', async ({ page }) => {
    await page.setViewportSize({ width: 390, height: 900 });
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await page.goto(`${GUILD_URL}#roster`);
    await openRosterAccounts(page);
    await expect(page.getByTestId('guild-roster-entries')).toBeVisible();
    const entry = page.locator('[data-testid="guild-roster-entry"][data-state="no_account"]');
    const [n, s, g] = await Promise.all([
      entry.getByTestId('guild-roster-entry-name').boundingBox(),
      entry.getByTestId('guild-roster-entry-state').boundingBox(),
      entry.getByTestId('guild-roster-entry-ingame').boundingBox(),
    ]);
    expect(n && s && g).toBeTruthy();
    if (n && s && g) {
      // The state word sits below the name, under it, and the in-game line below the state:
      // stacked, none beside another.
      expect(s.y).toBeGreaterThanOrEqual(n.y + n.height - 1);
      expect(g.y).toBeGreaterThanOrEqual(s.y + s.height - 1);
      expect(Math.abs(s.x - n.x)).toBeLessThan(4);
    }
    const scrolls = await page.evaluate(
      () => document.documentElement.scrollWidth > document.documentElement.clientWidth,
    );
    expect(scrolls).toBe(false);
    const dir = process.env.G7_SHOTS;
    if (dir)
      await page
        .getByTestId('guild-roster-entries')
        .screenshot({ path: `${dir}/g7-roster-entries-${page.viewportSize()?.width}.png` });
  });
});

test.describe('roster at 390px: three blocks under the name, nothing past the edge (round 7)', () => {
  test.use({ timezoneId: 'UTC' });

  test('every row fits, the facts keep one order, Remove is a bordered button that asks with a name', async ({
    page,
  }) => {
    await page.setViewportSize({ width: 390, height: 900 });
    await page.clock.setFixedTime(new Date('2026-12-22T21:46:00Z'));
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await page.goto(`${GUILD_URL}#roster`);
    await openRosterAccounts(page);
    await expect(page.getByTestId('guild-roster-rows')).toBeVisible();
    await expect(page.getByTestId('guild-roster-entries')).toBeVisible();

    // Nothing reaches the edge: the page does not scroll sideways and no row's text runs past it.
    expect(await page.evaluate(() => document.documentElement.scrollWidth)).toBeLessThanOrEqual(390);
    const overflowing = await page.evaluate(() =>
      [...document.querySelectorAll('[data-testid="guild-roster-rows"] *')]
        .filter((el) => el.getBoundingClientRect().right > 390.5)
        .map((el) => el.getAttribute('data-testid') ?? el.tagName),
    );
    expect(overflowing).toEqual([]);

    // Block order under the name: chips, then professions on one line, then the in-game facts.
    const zulmara = page
      .getByTestId('guild-roster-rows')
      .locator('.guild-roster-row')
      .filter({ has: page.getByTestId('guild-roster-name').getByText('Zulmara') });
    const [facts, prof, ingame] = await Promise.all([
      zulmara.getByTestId('guild-roster-facts').boundingBox(),
      zulmara.getByTestId('guild-roster-professions').boundingBox(),
      zulmara.getByTestId('guild-roster-ingame-slot').boundingBox(),
    ]);
    expect(facts && prof && ingame).toBeTruthy();
    if (facts && prof && ingame) {
      expect(prof.y).toBeGreaterThanOrEqual(facts.y);
      expect(ingame.y).toBeGreaterThanOrEqual(prof.y + prof.height - 1);
      expect(prof.height).toBeLessThan(20);
    }
    // Spec words carry the class; units are named; "Gear set:" replaces the old word, and the
    // spec line no longer says "in game" (d42bf7e65).
    await expect(zulmara.getByTestId('guild-roster-spec')).toContainText('Marksmanship on the site');
    await expect(zulmara).toContainText('Rating');
    await expect(zulmara.getByTestId('guild-roster-spec')).toContainText('Survival Hunter');
    await expect(zulmara.getByTestId('guild-roster-spec')).not.toContainText('in game');
    await expect(zulmara).toContainText('Gear set: Survival (raid)');
    await expect(zulmara).not.toContainText('per addon');
    await expect(zulmara).not.toContainText('build:');

    // Remove: the secondary button (a visible border, so it never reads as a link; stage 2 player verdict), 44px high on a phone, sitting on the name row; it asks with a name.
    const remove = page.locator('[data-testid="guild-roster-remove"][aria-label="Remove Windtalon"]');
    const style = await remove.evaluate((node) => {
      const css = getComputedStyle(node);
      return { border: css.borderTopColor, height: node.getBoundingClientRect().height };
    });
    expect(style.border).not.toBe('rgba(0, 0, 0, 0)');
    expect(style.height).toBeGreaterThanOrEqual(44);
    const windtalon = page
      .getByTestId('guild-roster-rows')
      .locator('.guild-roster-row')
      .filter({ has: page.getByTestId('guild-roster-name').getByText('Windtalon') });
    const [nameBox, removeBox] = await Promise.all([
      windtalon.getByTestId('guild-roster-name').boundingBox(),
      remove.boundingBox(),
    ]);
    if (nameBox && removeBox)
      expect(Math.abs(removeBox.y + removeBox.height / 2 - (nameBox.y + 10))).toBeLessThan(30);
    await remove.click();
    await expect(page.getByTestId('guild-roster-remove-confirm')).toContainText(
      'Remove Windtalon from the roster?',
    );
    // One muted line, they never she/he; Remove in the danger token, Cancel focused with the ring.
    await expect(page.getByTestId('guild-roster-remove-rejoin')).toHaveText(
      'They can rejoin with the invite link.',
    );
    await expect(page.getByTestId('guild-roster-remove-cancel')).toBeFocused();
    // The danger button (lib/ui/button.ts BUTTON_DANGER): wipe ink on a wipe hairline, never filled.
    const yes = await page
      .getByTestId('guild-roster-remove-yes')
      .evaluate((n) => [getComputedStyle(n).color, getComputedStyle(n).borderTopWidth]);
    const cancel = await page
      .getByTestId('guild-roster-remove-cancel')
      .evaluate((n) => getComputedStyle(n).color);
    expect(yes[0]).not.toBe(cancel);
    expect(yes[1]).toBe('1px');
  });

  // The line shows only while no rank is known at all (d05e718ba): a roster row's rank is the addon's.
  test('the ranks line sits above the roster only while no rank is known', async ({ page }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) =>
      route.fulfill(envelope({ ...buildMockHome(VIEWER_OFFICER), roster: [], pending: [] })),
    );
    await page.goto(`${GUILD_URL}#roster`);
    await openRosterAccounts(page);
    await expect(page.getByTestId('guild-roster-entries')).toBeVisible();
    await expect(page.getByTestId('guild-roster-ranks-pending')).toHaveText(
      'Ranks show once the addon sends them.',
    );
  });

  test('the ranks line is gone once the addon has sent ranks', async ({ page }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await page.goto(`${GUILD_URL}#roster`);
    await openRosterAccounts(page);
    await expect(page.getByTestId('guild-roster-rows')).toBeVisible();
    await expect(page.getByTestId('guild-roster-ranks-pending')).toHaveCount(0);
  });
});

test.describe('settings: billing and release belong to the guild master', () => {
  // The settings page is prerendered for the fixture guild The Last Watch (id 501) only.
  async function openSettings(page: Page, rank: 'leader' | 'officer'): Promise<void> {
    await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) =>
      route.fulfill(
        envelope({
          guild: { id: 501, name: 'The Last Watch', region: 'us', ruleset: 'hardcore' },
          progression: [],
          roster_best: [],
          reports: [],
        }),
      ),
    );
    await page.route('**/v1/me', (route) =>
      route.fulfill(
        envelope({
          user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
          characters: [],
          guilds: [
            {
              id: 501,
              region: 'us',
              ruleset: 'hardcore',
              name: 'The Last Watch',
              rank,
              verified: true,
              can: canFor(rank),
            },
          ],
        }),
      ),
    );
    // GET /v1/guilds/{id}/billing (launch pricing spec 3.2): the guild master pays, an officer reads.
    await page.route('**/v1/guilds/501/billing', (route) =>
      route.fulfill(
        envelope(guildBillingRead({ can: rank === 'leader' ? ['billing', 'seats', 'transfer'] : [] })),
      ),
    );
    await page.route('**/v1/guilds/501/settings', (route) =>
      route.fulfill(
        envelope({
          default_visibility: 'guild',
          officer_max_rank_index: 2,
          claimed_by: { name: 'Thornhide', class: 'druid', character_key: 'us/hardcore/thornhide' },
          claim_pending: null,
          claim: { state: 'claimed', since: '2026-09-01T00:00:00Z', frozen: false },
          invite: { rotated_at: null, url: null, made_by: null },
        }),
      ),
    );
    await page.goto(`/guild/${FIXTURE_GUILD_PATH}/settings`);
    await expect(page.getByTestId('guild-settings')).toBeVisible();
    // No link made yet (invite.rotated_at null): the row offers to make one, not rotate it.
    await expect(page.getByTestId('guild-invite-make')).toBeVisible();
  }

  test('the guild master sees the billing section', async ({ page }) => {
    await openSettings(page, 'leader');
    await expect(page.getByTestId('guild-settings-billing')).toBeVisible();
    await expect(page.getByTestId('guild-settings-master-only')).toHaveCount(0);
  });

  // Launch pricing spec 3.2: every officer reads the Billing section, with no buttons, and the
  // master-only line (6.2 "Read-only officer").
  test('an officer reads billing with no controls and is told what only the guild master can do', async ({
    page,
  }) => {
    await openSettings(page, 'officer');
    const billing = page.getByTestId('guild-settings-billing');
    await expect(billing).toBeVisible();
    await expect(billing.getByTestId('guild-billing-chooser')).toHaveCount(0);
    await expect(billing.getByTestId('guild-billing-subscribe')).toHaveCount(0);
    await expect(billing.getByTestId('guild-manage-billing')).toHaveCount(0);
    await expect(page.getByTestId('guild-settings-master-only')).toHaveText(
      'The subscriber or the guild leader manages billing and seats.',
    );
    // Everything else on the page is still hers.
    await expect(page.getByTestId('guild-invite-make')).toBeVisible();
  });
});

test.describe('readiness: worst-first sort and officer Nudge', () => {
  test('readiness rows sort worst-first by failing checks', async ({ page }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await stubDepthEndpoints(page);
    await page.goto(`${GUILD_URL}#readiness`);
    const readiness = buildMockReadiness();
    // The board (dense/ReadinessBoard.svelte): one row per raider short of something, worst first.
    const rows = page
      .getByTestId('guild-readiness-board')
      .locator('[data-testid^="guild-readiness-row-"]:visible');
    await expect(rows.first()).toBeVisible();
    const sorted = [...readiness.rows].sort(
      (a, b) =>
        b.failing * 100 + (b.gear_gap?.gain_dps ?? 0) - (a.failing * 100 + (a.gear_gap?.gain_dps ?? 0)),
    );
    const firstRowTestId = await rows.first().getAttribute('data-testid');
    expect(firstRowTestId).toBe(`guild-readiness-row-${sorted[0].character_key}`);
  });

  test('officer Nudge copies real text to the clipboard, with a distinct accessible name per raider', async ({
    page,
    context,
  }) => {
    await context.grantPermissions(['clipboard-read', 'clipboard-write']);
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await stubDepthEndpoints(page);
    await page.goto(`${GUILD_URL}#readiness`);
    const nudgeButtons = page.locator('[data-testid^="guild-readiness-nudge-"]:visible');
    const first = nudgeButtons.first();
    await first.click();
    const copied = await page.evaluate(() => navigator.clipboard.readText());
    expect(copied.length).toBeGreaterThan(0);
    const names = await nudgeButtons.evaluateAll((buttons) =>
      buttons.map((b) => b.getAttribute('aria-label')),
    );
    expect(new Set(names).size).toBe(names.length);
  });

  test('a member sees no Nudge button on the readiness tab', async ({ page }) => {
    await stubCore(page, VIEWER_MEMBER, 'member');
    await stubDepthEndpoints(page);
    await page.goto(`${GUILD_URL}#readiness`);
    await expect(page.getByTestId('guild-readiness-tab')).toBeVisible();
    await expect(page.locator('[data-testid^="guild-readiness-nudge-"]')).toHaveCount(0);
  });
});

// The Loot tab (the loot council) is covered by guild-loot.spec.ts.

test.describe('a 404 from an undeployed endpoint never crashes the page', () => {
  test('the Raids tab shows its own empty state with a next-step line, not a crash', async ({ page }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await page.route(`**/v1/guilds/${GUILD_ID}/raids*`, (route) =>
      route.fulfill({
        status: 404,
        contentType: 'application/json',
        body: '{"ok":false,"data":null,"error":{"message":"not found"},"request_id":"r"}',
      }),
    );
    await page.goto(`${GUILD_URL}#raids`);
    await expect(page.getByTestId('guild-raids-missing')).toBeVisible();
    await expect(page.getByTestId('guild')).toBeVisible();
  });
});

test.describe('the Overview last raid at 390', () => {
  test('every fact is visible and the page does not scroll sideways', async ({ page }) => {
    await page.setViewportSize({ width: 390, height: 900 });
    // The last raid is a member's (batch 3: the officer rail is the cover and attendance).
    await stubCore(page, VIEWER_MEMBER, 'member');
    await stubDepthEndpoints(page);
    await page.goto(GUILD_URL);
    const facts = page.getByTestId('guild-overview-raid-facts');
    await expect(facts).toContainText('boss');
    const pieces = facts.locator(':scope > span');
    const n = await pieces.count();
    expect(n).toBeGreaterThanOrEqual(2);
    for (let i = 0; i < n; i++) {
      const box = await pieces.nth(i).boundingBox();
      expect(box && box.x >= 0 && box.x + box.width <= 390).toBe(true);
    }
    expect(await page.evaluate(() => document.documentElement.scrollWidth)).toBeLessThanOrEqual(390);
  });
});

test.describe('no horizontal overflow', () => {
  for (const width of [390, 2000]) {
    test(`nothing scrolls sideways at ${width}px (officer, every tab)`, async ({ page }) => {
      await page.setViewportSize({ width, height: 900 });
      await stubCore(page, VIEWER_OFFICER, 'leader');
      await stubDepthEndpoints(page);
      for (const hash of ['', '#roster', '#raids', '#progression', '#readiness', '#loot']) {
        await page.goto(`${GUILD_URL}${hash}`);
        await expect(page.getByTestId('guild')).toBeVisible();
        const scrollWidth = await page.evaluate(() => document.documentElement.scrollWidth);
        expect(scrollWidth, hash || '#overview').toBeLessThanOrEqual(width);
      }
    });
  }
});

// Kept from the superseded guild-home.spec.ts: the cross-guild-membership guard (a member
// of one guild visiting a different guild sharing its region/ruleset sees no private
// section of their own guild) is unchanged mechanism and still exercised against the
// older FIXTURE_GUILD_PATH fixture, which this file does not touch.
test('visiting a different guild sharing region/ruleset shows no role-line from the viewer’s own guild', async ({
  page,
}) => {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) =>
    route.fulfill(
      envelope({
        guild: { id: 999, name: 'The Last Watch', region: 'us', ruleset: 'hardcore' },
        progression: [],
        roster_best: [],
        reports: [],
      }),
    ),
  );
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      envelope({
        user: {
          id: 7,
          battletag: 'Fixture#1234',
          email: null,
          role: 'user',
          anonymize: false,
          premium: false,
        },
        characters: [],
        guilds: [
          {
            id: GUILD_ID,
            region: REGION,
            ruleset: RULESET,
            name: 'Olympus XXVII',
            rank: 'leader',
            verified: true,
            can: canFor('leader'),
          },
        ],
      }),
    ),
  );
  await page.goto(`/guild/${FIXTURE_GUILD_PATH}`);
  await expect(page.getByTestId('guild')).toBeVisible();
  await expect(page.getByTestId('guild-role-line')).toHaveCount(0);
});

// Ported from the superseded guild-home.spec.ts: the claim/contest/settings-link role line
// (unchanged v1 mechanism, now living above the tab strip, Guild.svelte) and the roster
// tab's approve/remove mutations.
test.describe('claim, contest and roster mutations (unchanged v1 mechanism)', () => {
  test('the guild master of an unclaimed guild sees the claim row', async ({ page }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) =>
      route.fulfill(
        envelope({ ...buildMockHome(VIEWER_OFFICER), claim: { state: 'unclaimed', frozen: false } }),
      ),
    );
    await page.goto(GUILD_URL);
    await expect(page.getByTestId('guild-claim-row-sentence')).toContainText(
      'Claim it here to approve members',
    );
    await expect(page.getByTestId('guild-claim-link')).toHaveText(`Claim ${buildMockGuildPage().guild.name}`);
    await expect(page.getByTestId('guild-settings-link')).toHaveCount(0);
  });

  test('an officer of a guild with a pending claim sees "Open the claim"', async ({ page }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) =>
      route.fulfill(
        envelope({ ...buildMockHome(VIEWER_OFFICER), claim: { state: 'pending', frozen: false } }),
      ),
    );
    await page.goto(GUILD_URL);
    await expect(page.getByTestId('guild-confirm-claim-link')).toContainText('Open the claim');
  });

  // Fix round 1, item 2: "Contest this claim" moved out of the header band entirely and
  // into the Settings tab, under the claim block -- and it never reaches the account that
  // already holds the claim. buildMockHome's own `claim.yours` is false, so this officer reads
  // as "a verified officer of a different account."
  test('"Contest this claim" never renders in the header, only in Settings for a non-claimant officer', async ({
    page,
  }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    let homeCalls = 0;
    await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) => {
      homeCalls += 1;
      const claim =
        homeCalls === 1
          ? {
              state: 'claimed' as const,
              frozen: false,
              yours: false,
              claimed_by: { name: 'Kraggor', class: 'warrior', character_key: 'us/pvp/kraggor' },
            }
          : { state: 'contested' as const, since: '2026-12-01T00:00:00Z', frozen: true };
      return route.fulfill(envelope({ ...buildMockHome(VIEWER_OFFICER), claim }));
    });
    let contestCalled = false;
    await page.route(`**/v1/guilds/${GUILD_ID}/claim/contest`, (route) => {
      contestCalled = true;
      return route.fulfill(envelope({ status: 'contested' }));
    });
    await page.goto(GUILD_URL);
    await expect(page.getByTestId('guild-home-contest-button')).toHaveCount(0);

    await page.getByTestId('guild-tab-settings').click();
    await page.getByTestId('guild-home-contest-button').click();
    await page.getByTestId('guild-home-contest-confirm-button').click();
    await expect(page.getByTestId('guild-home-frozen')).toBeVisible();
    expect(contestCalled).toBe(true);
  });

  test('the claimant’s own account sees no Contest control anywhere, not even in Settings', async ({
    page,
  }) => {
    // `claim.yours`: this officer's account holds the claim.
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) => {
      const home = buildMockHome(VIEWER_OFFICER);
      return route.fulfill(envelope({ ...home, claim: { ...home.claim, yours: true } }));
    });
    await page.goto(`${GUILD_URL}#settings`);
    await expect(page.getByTestId('guild-settings-claim-contest')).toHaveCount(0);
    await expect(page.getByTestId('guild-home-contest-button')).toHaveCount(0);
  });

  test('a contested claim disables Approve and Remove on the roster tab', async ({ page }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) =>
      route.fulfill(
        envelope({
          ...buildMockHome(VIEWER_OFFICER),
          claim: { state: 'contested', since: '2026-12-01T00:00:00Z', frozen: true },
        }),
      ),
    );
    await page.goto(`${GUILD_URL}#roster`);
    await openRosterAccounts(page);
    await expect(page.getByTestId('guild-roster-approve').first()).toBeDisabled();
    await expect(page.getByTestId('guild-roster-remove').first()).toBeDisabled();
  });

  test('approving an unverified raider calls the approve endpoint and clears their unverified pill', async ({
    page,
  }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    let approveCalled = false;
    await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) =>
      route.fulfill(
        envelope(
          approveCalled
            ? withoutRow(buildMockHome(VIEWER_OFFICER), 'Rootgall')
            : buildMockHome(VIEWER_OFFICER),
        ),
      ),
    );
    await page.route(`**/v1/guilds/${GUILD_ID}/characters/us/pvp/rootgall/approve`, (route) => {
      approveCalled = true;
      return route.fulfill(envelope({ character_key: 'us/pvp/rootgall', status: 'approved' }));
    });
    await page.goto(`${GUILD_URL}#roster`);
    await openRosterAccounts(page);
    const approveButton = page.locator('[data-testid="guild-roster-approve"][aria-label="Approve Rootgall"]');
    await approveButton.click();
    await expect.poll(() => approveCalled).toBe(true);
    await expect(page.getByTestId('guild-roster-unverified-list')).not.toContainText('Rootgall');
  });

  test('a 403 on remove shows the API sentence, not a silent failure', async ({ page }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await page.route('**/v1/guilds/*/characters/us/pvp/sunderfel', (route) =>
      route.fulfill({
        status: 403,
        contentType: 'application/json',
        body: JSON.stringify({
          ok: false,
          data: null,
          error: { message: 'you do not have standing to remove this character' },
          request_id: 'r',
        }),
      }),
    );
    await page.goto(`${GUILD_URL}#roster`);
    await openRosterAccounts(page);
    const removeButton = page.locator('[data-testid="guild-roster-remove"][aria-label="Remove Sunderfel"]');
    await removeButton.click();
    await expect(page.getByTestId('guild-roster-remove-confirm')).toContainText(
      'Remove Sunderfel from the roster?',
    );
    await page.getByTestId('guild-roster-remove-yes').click();
    await expect(page.getByTestId('guild-roster-action-error')).toHaveText(
      'you do not have standing to remove this character',
    );
  });
});

// Guild crest round (docs/contracts/2026-10-05-guild-crest-api.md): the Settings tab's
// "Guild crest" block, officer only -- the web lane's own choice (the contract itself
// mentions a moderator may remove via the API, but this build offers the block to
// `role === 'officer'` alone, matching the build brief).
test.describe('guild crest (Settings tab, officer only)', () => {
  test('an officer sees the block with the faction default', async ({ page }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await page.goto(`${GUILD_URL}#settings`);
    await expect(page.getByTestId('guild-crest-settings')).toBeVisible();
    await expect(page.getByTestId('guild-crest-current')).toHaveAttribute('src', /horde-logo-512\.webp$/);
    await expect(page.getByTestId('guild-crest-default-caption')).toHaveCount(0);
  });

  test('a member never sees the block', async ({ page }) => {
    await stubCore(page, VIEWER_MEMBER);
    await page.goto(`${GUILD_URL}#settings`);
    // The Settings tab itself is officer/moderator only -- a member's stale/bookmarked
    // #settings hash falls back to the first tab they are allowed (Overview).
    await expect(page.getByTestId('guild-crest-settings')).toHaveCount(0);
  });

  test('choosing a small PNG shows the preview and enables Save; Save routes PUT and the header ring switches to the returned crest', async ({
    page,
  }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    let homeCalls = 0;
    await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) => {
      homeCalls += 1;
      const crestUrl = homeCalls > 1 ? 'https://cdn.test/guilds/2024/crest.webp?v=abc' : null;
      return route.fulfill(envelope(buildMockHome(VIEWER_OFFICER, 'horde', crestUrl)));
    });
    let putCalled = false;
    await page.route(`**/v1/guilds/${GUILD_ID}/crest`, (route) => {
      if (route.request().method() !== 'PUT') return route.fallback();
      putCalled = true;
      return route.fulfill(envelope({ crest_url: 'https://cdn.test/guilds/2024/crest.webp?v=abc' }));
    });

    await page.goto(`${GUILD_URL}#settings`);
    const saveButton = page.getByTestId('guild-crest-save');
    await expect(saveButton).toBeDisabled();
    // P.1: a styled "Choose an image" button, not the browser's own file control.
    await expect(page.getByTestId('guild-crest-choose')).toHaveText('Choose an image');
    await expect(page.getByTestId('guild-crest-file-name')).toHaveText('No image chosen');
    await expect(page.getByTestId('guild-crest-file-input')).toHaveCSS('position', 'absolute');

    await page.getByTestId('guild-crest-file-input').setInputFiles({
      name: 'crest.png',
      mimeType: 'image/png',
      buffer: Buffer.from(new Uint8Array(1024)),
    });
    await expect(page.getByTestId('guild-crest-preview')).toBeVisible();
    await expect(page.getByTestId('guild-crest-file-name')).toHaveText('crest.png');
    await p1Shot(page, 'guild-settings-crest-chosen-1280');
    await expect(saveButton).toBeEnabled();

    await saveButton.click();
    // The click resolves before the upload's PUT reaches the route, so poll for it.
    await expect.poll(() => putCalled).toBe(true);
    await expect(page.getByTestId('guild-faction-crest')).toHaveAttribute('src', /crest\.webp\?v=abc$/);
  });

  test('Remove routes DELETE and the ring falls back to the faction logo', async ({ page }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    let homeCalls = 0;
    await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) => {
      homeCalls += 1;
      const crestUrl = homeCalls > 1 ? null : 'https://cdn.test/guilds/2024/crest.webp?v=keep';
      return route.fulfill(envelope(buildMockHome(VIEWER_OFFICER, 'horde', crestUrl)));
    });
    let deleteCalled = false;
    await page.route(`**/v1/guilds/${GUILD_ID}/crest`, (route) => {
      if (route.request().method() !== 'DELETE') return route.fallback();
      deleteCalled = true;
      return route.fulfill({ status: 204 });
    });

    await page.goto(`${GUILD_URL}#settings`);
    await expect(page.getByTestId('guild-faction-crest')).toHaveAttribute('src', /crest\.webp\?v=keep$/);

    await page.getByTestId('guild-crest-remove').click();
    await page.getByTestId('guild-crest-remove-confirm-button').click();
    await expect.poll(() => deleteCalled).toBe(true);
    await expect(page.getByTestId('guild-faction-crest')).toHaveAttribute('src', /horde-logo-512\.webp$/);
  });

  test('an oversize file shows the size message without a request', async ({ page }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    let putCalled = false;
    await page.route(`**/v1/guilds/${GUILD_ID}/crest`, (route) => {
      putCalled = true;
      return route.fulfill(envelope({ crest_url: 'x' }));
    });

    await page.goto(`${GUILD_URL}#settings`);
    await page.getByTestId('guild-crest-file-input').setInputFiles({
      name: 'big.png',
      mimeType: 'image/png',
      buffer: Buffer.alloc(3 * 1024 * 1024),
    });
    await expect(page.getByTestId('guild-crest-file-error')).toHaveText(
      'That file is 3.0 MB; the limit is 2 MB.',
    );
    await expect(page.getByTestId('guild-crest-save')).toBeDisabled();
    expect(putCalled).toBe(false);
  });

  test('the picker has a visible label, Save/Remove have distinct names, and the file error is role=alert', async ({
    page,
  }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) =>
      route.fulfill(
        envelope(buildMockHome(VIEWER_OFFICER, 'horde', 'https://cdn.test/guilds/2024/crest.webp?v=1')),
      ),
    );
    await page.goto(`${GUILD_URL}#settings`);

    await expect(page.getByLabel('Choose a crest image')).toBeVisible();
    await expect(page.getByRole('button', { name: 'Save', exact: true })).toBeVisible();
    await expect(page.getByRole('button', { name: 'Remove', exact: true })).toBeVisible();

    await page.getByTestId('guild-crest-file-input').setInputFiles({
      name: 'bad.gif',
      mimeType: 'image/gif',
      buffer: Buffer.from('x'),
    });
    await expect(page.getByTestId('guild-crest-file-error')).toHaveAttribute('role', 'alert');
  });
});

/** A capture for the review when P1_SHOTS names a folder. */
async function p1Shot(page: Page, name: string): Promise<void> {
  const dir = process.env.P1_SHOTS;
  if (dir === undefined) return;
  await page.screenshot({ path: `${dir}/p1-${name}.png`, fullPage: true });
}

/** The visible words of an element, no-break spaces read as spaces. */
async function words(locator: Locator): Promise<string> {
  return ((await locator.innerText()) ?? '')
    .replace(/\u00a0/g, ' ')
    .replace(/\s+/g, ' ')
    .trim();
}

test.describe('guild facts read as sentences (P.1)', () => {
  for (const width of [1280, 390]) {
    test(`at ${width} the header and Overview facts are body-face sentences with " · " spaced`, async ({
      page,
    }) => {
      await page.setViewportSize({ width, height: 900 });
      await stubCore(page, VIEWER_MEMBER, 'member');
      await stubDepthEndpoints(page);
      await page.goto(GUILD_URL);
      const header = page.getByTestId('guild-header-facts');
      await expect(header).toBeVisible();
      expect(await words(header)).toMatch(/^US PvP( · \d+ raiders?)?$/);
      const raidFacts = page.getByTestId('guild-overview-raid-facts').first();
      expect(await words(raidFacts)).toMatch(/ · /);
      for (const line of [header, raidFacts]) {
        expect(await line.evaluate((el) => getComputedStyle(el).fontFamily)).not.toMatch(/mono/i);
      }
      // Each fact is one piece: no fact's box is split across two lines.
      for (const line of [header, raidFacts]) {
        const split = await line.evaluate(
          (el) =>
            [...el.querySelectorAll('.whitespace-nowrap')].filter(
              (piece) =>
                piece.getBoundingClientRect().height > parseFloat(getComputedStyle(piece).fontSize) * 2.2,
            ).length,
        );
        expect(split).toBe(0);
      }
      expect(await page.evaluate(() => document.documentElement.scrollWidth)).toBeLessThanOrEqual(width);
      await p1Shot(page, `guild-overview-${width}`);
    });

    test(`at ${width} the roster stats read in the body face: a sentence on a phone, columns under heads from lg`, async ({
      page,
    }) => {
      await page.setViewportSize({ width, height: 900 });
      await stubCore(page, VIEWER_OFFICER, 'leader');
      await page.goto(`${GUILD_URL}#roster`);
      await openRosterAccounts(page);
      const stats = page.getByTestId('guild-roster-facts').first();
      await expect(stats).toBeVisible();
      expect(await stats.evaluate((el) => getComputedStyle(el).fontFamily)).not.toMatch(/mono/i);
      if (width >= 1024) {
        await expect(page.getByTestId('guild-roster-heads')).toContainText('Item level');
        expect(await words(stats)).toMatch(/\b\d+\b/);
        expect(await words(stats)).not.toMatch(/Item level/);
      } else {
        expect(await words(stats)).toMatch(/Item level \d+/);
      }
      await p1Shot(page, `guild-roster-${width}`);
    });
  }

  // Raids spec 3.4 put the next-raid line in the band (y < 300 at 1280, spec §1), so the first
  // roster row now starts above the fold rather than ending above it.
  test('at 1280 the claim card is one line, the raid line is in the band and the first roster row starts above the fold', async ({
    page,
  }) => {
    await page.setViewportSize({ width: 1280, height: 800 });
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await page.goto(`${GUILD_URL}#roster`);
    await expect(page.getByTestId('guild-officer-strip')).toHaveCount(0);
    const raidLine = page
      .locator('[data-testid="guild-raid-line"], [data-testid="guild-raid-line-none"]')
      .first();
    await expect(raidLine).toBeVisible();
    const lineBox = await raidLine.boundingBox();
    expect(lineBox).not.toBeNull();
    expect(lineBox!.y).toBeLessThan(300);
    // The raid roster leads the tab (dense spec 5.2): its head starts above the fold.
    const card = page.getByTestId('guild-raid-roster');
    await expect(card).toBeVisible();
    const box = await card.boundingBox();
    expect(box).not.toBeNull();
    expect(box!.y).toBeLessThan(800);
  });

  test('the Overview prints no "no encounter list yet" line of ours', async ({ page }) => {
    await stubCore(page, VIEWER_MEMBER, 'member');
    await stubDepthEndpoints(page);
    await page.goto(GUILD_URL);
    await expect(page.getByTestId('guild-overview-last-raid')).toBeVisible();
    await expect(page.getByTestId('guild-overview-tab')).not.toContainText('no encounter list');
  });
});
