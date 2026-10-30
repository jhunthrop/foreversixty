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
import { expect, test, type Locator, type Page } from '@playwright/test';
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
    const rows = page.getByTestId('guild-roster-rows').locator('.guild-roster-row');
    await expect(rows.first()).toContainText('(you)');

    await expect(page.getByTestId('guild-roster-waiting')).toContainText('Waiting for approval');
  });

  test('officer sees Approve all, with a distinct accessible name per raider for Approve/Remove', async ({
    page,
  }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await page.goto(`${GUILD_URL}#roster`);
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
    await expect(page.getByTestId('guild-roster-rows')).toBeVisible();
  }

  function rowNamed(page: Page, name: string) {
    return page
      .getByTestId('guild-roster-rows')
      .locator('.guild-roster-row')
      .filter({ has: page.getByTestId('guild-roster-name').getByText(name, { exact: false }) });
  }

  test('the line above the table names the snapshot; matched rows read level, spec, gap, build and seen time', async ({
    page,
  }) => {
    await openRoster(page, true);
    await expect(page.getByTestId('guild-roster-addon-snapshot')).toHaveText(
      'Roster from the addon · sent by Zulmara 3 minutes ago',
    );
    await expect(rowNamed(page, 'Zulmara').getByTestId('guild-roster-ingame')).toHaveText(
      /^Level 60\s+·\s+Gear gap 12 DPS\s+·\s+Gear set: Survival \(raid\)\s+·\s+seen 6 minutes ago\s*$/,
    );
    await expect(rowNamed(page, 'Kraggor').getByTestId('guild-roster-ingame')).toHaveText(
      /^Level 60\s+·\s+Gear gap 4 DPS\s+·\s+Gear set: Protection \(raid\)\s+·\s+seen 8 minutes ago\s*$/,
    );
    const none = rowNamed(page, 'Windtalon').getByTestId('guild-roster-ingame-none');
    await expect(none).toHaveText('—');
    await expect(none).toHaveAttribute('title', 'Not seen by the addon in the last snapshot');
    // The spec line says which record it reads: the addon's in game, the site's where it differs.
    await expect(rowNamed(page, 'Zulmara').getByTestId('guild-roster-spec')).toContainText(
      'Survival Hunter in game · Marksmanship on the site',
    );
    await expect(rowNamed(page, 'Kraggor').getByTestId('guild-roster-spec')).toHaveText(
      'Protection Warrior in game',
    );
  });

  test('the roster has a SORT row of chips and no header row at any width; the chips sort', async ({
    page,
  }) => {
    await openRoster(page, true);
    await expect(page.locator('[role="columnheader"]')).toHaveCount(0);
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
    await expect(page.getByTestId('guild-roster-sort-ilvl')).toHaveAttribute('aria-sort', 'descending');
    // Rank puts the leader and officers first (the viewer stays pinned above them).
    await page.getByTestId('guild-roster-sort-rank').click();
    await expect(page.getByTestId('guild-roster-sort-rank')).toContainText('▾');
    await expect(page.getByTestId('guild-roster-sort-rank')).toHaveAttribute('aria-sort', 'descending');
    await expect(page.getByTestId('guild-roster-sort-ilvl')).not.toHaveAttribute('aria-sort', 'descending');
    const names = (
      await page.getByTestId('guild-roster-rows').getByTestId('guild-roster-name').allTextContents()
    ).map((n) => n.replace(' (you)', '').trim());
    expect(names[0]).toBe('Zulmara');
    // The In game column's gap is a fact on the row, not a sort the page offers.
    await expect(page.getByTestId('guild-roster-sort-addon_gap')).toHaveCount(0);
    // Every row's stats start at one left edge at lg.
    if ((page.viewportSize()?.width ?? 0) >= 1024) {
      const lefts = await page
        .getByTestId('guild-roster-rows')
        .getByTestId('guild-roster-facts')
        .evaluateAll((els) =>
          els.map((e) =>
            Math.round(
              e.querySelector('[data-testid="guild-roster-chip"]')?.getBoundingClientRect().left ?? -1,
            ),
          ),
        );
      expect(new Set(lefts).size).toBe(1);
    }
  });

  test('with no snapshot: the null line, and every row "—"', async ({ page }) => {
    await openRoster(page, false);
    await expect(page.getByTestId('guild-roster-addon-snapshot')).toHaveText(
      'No roster from the addon yet. Log in with the addon on your guild master and it sends one.',
    );
    await expect(page.getByTestId('guild-roster-ingame')).toHaveCount(0);
    const rows = await page.getByTestId('guild-roster-rows').locator('.guild-roster-row').count();
    await expect(page.getByTestId('guild-roster-ingame-none')).toHaveCount(rows);
  });
});

// Lifecycle round 4, the roster model: everyone the addon saw with a word beside each, what
// approving does, and controls that come from the account's capability list.
test.describe('roster: everyone the addon saw, with a word beside each', () => {
  test.use({ timezoneId: 'UTC' });

  test('the count line, the rank word on every row and the names with no account', async ({ page }) => {
    await page.clock.setFixedTime(new Date('2026-12-22T21:46:00Z'));
    await stubCore(page, VIEWER_MEMBER, 'member');
    await page.goto(`${GUILD_URL}#roster`);
    const home = buildMockHome(VIEWER_MEMBER);
    const waiting = home.roster.filter((r) => !r.verified).length;
    await expect(page.getByTestId('guild-roster-counts')).toHaveText(
      `${home.roster.length + 2} on the roster · 1 in game without an account · ${waiting} waiting for approval · 1 not seen in game yet`,
    );
    await expect(page.getByTestId('guild-roster-gear-note')).toHaveText(
      'Loot and readiness list members with gear on the site.',
    );
    // Member is a word, not a blank: a plain member's row says so.
    const zulmara = page
      .getByTestId('guild-roster-rows')
      .locator('.guild-roster-row')
      .filter({ has: page.getByTestId('guild-roster-name').getByText('Zulmara') });
    await expect(zulmara.getByTestId('guild-roster-rank')).toHaveText('Member');
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
    await expect(page.getByTestId('guild-roster-approve-explains')).toHaveText(
      'Approving lets this character see the roster, loot and readiness; officers can also remove members and change settings.',
    );
    const list = page.getByTestId('guild-roster-unverified-list');
    await expect(list.getByTestId('guild-roster-pending-rank').first()).toHaveText('Member');
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
    const own = page
      .getByTestId('guild-roster-rows')
      .locator('.guild-roster-row')
      .filter({ has: page.getByTestId('guild-roster-name').getByText(VIEWER_OFFICER) });
    await expect(own).toBeVisible();
    await expect(own.getByTestId('guild-roster-rank')).toHaveText('Leader');
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
    // Spec words carry the class; units are named; "in game" and "Gear set:" replace the old words.
    await expect(zulmara.getByTestId('guild-roster-spec')).toContainText('Marksmanship on the site');
    await expect(zulmara).toContainText('Rating');
    await expect(zulmara).toContainText('Survival Hunter in game');
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
    const yes = await page
      .getByTestId('guild-roster-remove-yes')
      .evaluate((n) => getComputedStyle(n).backgroundColor);
    const cancel = await page
      .getByTestId('guild-roster-remove-cancel')
      .evaluate((n) => getComputedStyle(n).backgroundColor);
    expect(yes).not.toBe(cancel);
    expect(yes).not.toBe('rgba(0, 0, 0, 0)');
  });

  test('the ranks line sits above the roster while a row has no rank', async ({ page }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await page.goto(`${GUILD_URL}#roster`);
    await expect(page.getByTestId('guild-roster-ranks-pending')).toHaveText(
      'Ranks show once the addon sends them.',
    );
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
          claimed_by: { battletag: 'Fixture#1234' },
          claim_pending: null,
          claim: { state: 'claimed', since: '2026-09-01T00:00:00Z', frozen: false },
          invite: { rotated_at: null },
        }),
      ),
    );
    await page.goto(`/guild/${FIXTURE_GUILD_PATH}/settings`);
    await expect(page.getByTestId('guild-settings')).toBeVisible();
    await expect(page.getByTestId('guild-invite-rotate')).toBeVisible();
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
    await expect(page.getByTestId('guild-invite-rotate')).toBeVisible();
  });
});

test.describe('readiness: worst-first sort and officer Nudge', () => {
  test('readiness rows sort worst-first by failing checks', async ({ page }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await stubDepthEndpoints(page);
    await page.goto(`${GUILD_URL}#readiness`);
    const readiness = buildMockReadiness();
    // Each row renders twice (a desktop grid row and a phone stacked card sharing the same
    // testid, GuildReadiness.svelte) -- only one is ever visible at a given viewport, so
    // this must filter to :visible rather than take the first DOM match.
    const rows = page
      .getByTestId('guild-readiness-rows')
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

test.describe('loot: Award asks first and can be undone', () => {
  /** A loot endpoint that keeps the awards the page records, the way the API does. */
  async function stubAwards(page: Page): Promise<{ posts: unknown[]; deletes: string[] }> {
    const record: { posts: unknown[]; deletes: string[] } = { posts: [], deletes: [] };
    const awards = new Map<number, { award_id: number; character_key: string; name: string; at: string }>();
    await page.route(`**/v1/guilds/${GUILD_ID}/loot*`, (route) => {
      const loot = buildMockLoot();
      return route.fulfill(
        envelope({
          ...loot,
          items: loot.items.map((item) => {
            const award = awards.get(item.item_id);
            return {
              ...item,
              awarded_to: award === undefined ? null : { ...award, by_name: VIEWER_OFFICER },
            };
          }),
        }),
      );
    });
    await page.route(`**/v1/guilds/${GUILD_ID}/loot/awards**`, (route) => {
      const request = route.request();
      if (request.method() === 'POST') {
        const body = request.postDataJSON() as { item_id: number; character_key: string };
        record.posts.push(body);
        const candidate = buildMockLoot()
          .items.find((item) => item.item_id === body.item_id)
          ?.candidates.find((c) => c.character_key === body.character_key);
        awards.set(body.item_id, {
          award_id: 41,
          character_key: body.character_key,
          name: candidate?.name ?? '',
          at: new Date().toISOString(),
        });
        return route.fulfill(envelope({ award_id: 41, status: 'awarded' }));
      }
      record.deletes.push(request.url().split('/').pop() ?? '');
      awards.clear();
      return route.fulfill(envelope({ status: 'removed' }));
    });
    return record;
  }

  test('the officer is asked, can Cancel, awards, sees who got it, and can Undo', async ({ page }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await page.route(`**/v1/guilds/${GUILD_ID}/raids*`, (route) => route.fulfill(envelope(buildMockRaids())));
    await page.route(`**/v1/guilds/${GUILD_ID}/progression`, (route) =>
      route.fulfill(envelope(buildMockProgression())),
    );
    const record = await stubAwards(page);
    const item = buildMockLoot().items.find((i) => i.name === 'Stormrage Cover');
    const first = item?.candidates[0];
    expect(first).toBeDefined();
    if (item === undefined || first === undefined) return;
    await page.goto(`${GUILD_URL}#loot`);
    const panel = page.getByTestId(`guild-loot-item-${item.item_id}`);
    await expect(panel).toBeVisible();
    const awardButton = panel.getByTestId(`guild-loot-award-${item.item_id}-${first.character_key}`);
    const confirm = page.getByTestId(`guild-loot-award-confirm-${item.item_id}`);
    const cancel = page.getByTestId(`guild-loot-award-cancel-${item.item_id}`);
    const dir = process.env.G7_SHOTS;

    // Award asks first, naming the item and the candidate, with focus on Cancel; Cancel records nothing.
    await awardButton.click();
    await expect(confirm).toContainText(`Award Stormrage Cover to ${first.name}?`);
    await expect(cancel).toBeFocused();
    await cancel.click();
    await expect(confirm).toHaveCount(0);
    expect(record.posts).toHaveLength(0);

    // Escape backs out too.
    await awardButton.click();
    await page.keyboard.press('Escape');
    await expect(confirm).toHaveCount(0);

    await awardButton.click();
    // While the confirm is open the Award buttons on the rows are disabled, so one live Award exists,
    // and Cancel shows the focus ring.
    for (const row of await panel.locator('[data-testid^="guild-loot-award-' + item.item_id + '-"]').all()) {
      await expect(row).toBeDisabled();
    }
    await expect(page.getByTestId(`guild-loot-award-yes-${item.item_id}`)).toBeEnabled();
    expect(await cancel.evaluate((node) => getComputedStyle(node).outlineStyle)).not.toBe('none');
    await page.keyboard.press('Escape');
    await expect(confirm).toHaveCount(0);
    await awardButton.click();
    if (dir) await panel.screenshot({ path: `${dir}/g7-award-confirm-${page.viewportSize()?.width}.png` });
    await page.getByTestId(`guild-loot-award-yes-${item.item_id}`).click();
    await expect.poll(() => record.posts.length).toBe(1);
    expect(record.posts[0]).toMatchObject({ item_id: item.item_id, character_key: first.character_key });

    // After: who got it, when, from whom, Undo; everyone else reads "Already awarded", never a live button.
    await expect(page.getByTestId(`guild-loot-awarded-${item.item_id}`)).toHaveText(
      `Awarded to ${first.name} · today · by\u00a0${VIEWER_OFFICER}`,
    );
    // One Undo only: the winner row's. The awarded chip carries none.
    await expect(page.getByTestId(`guild-loot-undo-${item.item_id}`)).toHaveCount(0);
    await expect(page.getByTestId(`guild-loot-row-undo-${item.item_id}`)).toBeVisible();
    // The winner's own row reads "Awarded · Undo"; the other candidates keep "Already awarded".
    await expect(panel.getByTestId('guild-loot-winner-slot')).toHaveText(/^Awarded\s+·\s+Undo$/);
    await expect(panel.locator('[data-testid^="guild-loot-award-"]')).toHaveCount(0);
    await expect(panel.getByText('Already awarded')).toHaveCount(item.candidates.length - 1);
    if (dir) await panel.screenshot({ path: `${dir}/g7-award-done-${page.viewportSize()?.width}.png` });

    // Undo removes it by the award's id: the Award buttons come back.
    await page.getByTestId(`guild-loot-row-undo-${item.item_id}`).click();
    await expect.poll(() => record.deletes).toEqual(['41']);
    await expect(page.getByTestId(`guild-loot-awarded-${item.item_id}`)).toHaveCount(0);
    await expect(panel.locator('[data-testid^="guild-loot-award-"]')).toHaveCount(item.candidates.length);
  });

  test('at 390px the awarded line sits under the card title: its box holds its text and clears the title', async ({
    page,
  }) => {
    await page.setViewportSize({ width: 390, height: 900 });
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await page.route(`**/v1/guilds/${GUILD_ID}/raids*`, (route) => route.fulfill(envelope(buildMockRaids())));
    await page.route(`**/v1/guilds/${GUILD_ID}/progression`, (route) =>
      route.fulfill(envelope(buildMockProgression())),
    );
    await stubAwards(page);
    const item = buildMockLoot().items.find((i) => i.name === 'Stormrage Cover');
    const first = item?.candidates[0];
    if (item === undefined || first === undefined) throw new Error('fixture lost Stormrage Cover');
    await page.goto(`${GUILD_URL}#loot`);
    const panel = page.getByTestId(`guild-loot-item-${item.item_id}`);
    await panel.getByTestId(`guild-loot-award-${item.item_id}-${first.character_key}`).click();
    await page.getByTestId(`guild-loot-award-yes-${item.item_id}`).click();
    const chip = page.getByTestId(`guild-loot-awarded-${item.item_id}`);
    await expect(chip).toBeVisible();
    // Sentence case, 14px, and the box grows with its text: nothing overflows it.
    await expect(chip).toHaveText(`Awarded to ${first.name} · today · by\u00a0${VIEWER_OFFICER}`);
    const fit = await chip.evaluate((node) => ({
      overflowY: node.scrollHeight - node.clientHeight,
      overflowX: node.scrollWidth - node.clientWidth,
      size: getComputedStyle(node).fontSize,
      transform: getComputedStyle(node).textTransform,
    }));
    expect(fit.overflowY).toBeLessThanOrEqual(0);
    expect(fit.overflowX).toBeLessThanOrEqual(0);
    expect(fit.size).toBe('14px');
    expect(fit.transform).toBe('none');
    // The title's box and the chip's box do not intersect; the chip is under the title, full width.
    const title = await panel.locator('h2').first().boundingBox();
    const box = await chip.boundingBox();
    const undo = await page.getByTestId(`guild-loot-row-undo-${item.item_id}`).boundingBox();
    expect(title && box && undo).toBeTruthy();
    if (title && box && undo) {
      expect(box.y).toBeGreaterThanOrEqual(title.y + title.height - 1);
      expect(undo.height).toBeGreaterThanOrEqual(44);
      // The one Undo is the winner row's, below the chip.
      expect(undo.y).toBeGreaterThanOrEqual(box.y + box.height);
    }
    expect(await panel.getByText('·', { exact: true }).count()).toBe(0);
    const scrolls = await page.evaluate(() => document.documentElement.scrollWidth);
    expect(scrolls).toBeLessThanOrEqual(390);
  });

  test('a loot candidate reads what it replaces, the spec and the raid nights', async ({ page }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await stubDepthEndpoints(page);
    await page.goto(`${GUILD_URL}#loot`);
    const item = buildMockLoot().items.find((i) => i.name === 'Stormrage Cover');
    const first = item?.candidates[0];
    if (item === undefined || first === undefined) throw new Error('fixture lost Stormrage Cover');
    const line = page.getByTestId(`guild-loot-item-${item.item_id}`).getByTestId('guild-loot-candidate-line');
    await expect(line.first()).toContainText(`Replaces ${first.replaces} · ${first.spec} · `);
    await expect(line.first()).toContainText('raid nights');
    await expect(line.first()).not.toContainText('over their current');
  });

  test('a member sees an awarded item but no Undo and no Award', async ({ page }) => {
    await stubCore(page, VIEWER_MEMBER, 'member');
    await stubDepthEndpoints(page);
    await page.goto(`${GUILD_URL}#loot`);
    await expect(page.locator('[data-testid^="guild-loot-awarded-"]').first()).toBeVisible();
    await expect(page.locator('[data-testid^="guild-loot-undo-"]')).toHaveCount(0);
    await expect(page.locator('[data-testid^="guild-loot-award-"]')).toHaveCount(0);
  });

  test('"Head of Onyxia" is listed once and a paladin helm is offered to no warrior', async ({ page }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await stubDepthEndpoints(page);
    await page.goto(`${GUILD_URL}#loot`);
    await expect(page.getByTestId('guild-loot-items').getByText('Head of Onyxia')).toHaveCount(1);
    const crown = buildMockLoot().items.find((i) => i.name === 'Judgement Crown');
    expect(crown?.candidates).toEqual([]);
    await expect(page.getByTestId(`guild-loot-no-candidate-${crown?.item_id}`)).toBeVisible();
  });
});

test.describe('loot: read-only for a member, awarded state for everyone', () => {
  test('officer sees the Award control and the already-awarded item’s state', async ({ page }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await stubDepthEndpoints(page);
    await page.goto(`${GUILD_URL}#loot`);
    await expect(page.getByTestId('guild-loot-tab')).toBeVisible();
    const loot = buildMockLoot();
    const awardedItem = loot.items.find((item) => item.awarded_to !== null);
    expect(awardedItem).toBeDefined();
    if (awardedItem !== undefined) {
      const itemPanel = page.getByTestId(`guild-loot-item-${awardedItem.item_id}`);
      await expect(itemPanel).toContainText('Awarded');
    }
    const unawarded = loot.items.find((item) => item.awarded_to === null);
    if (unawarded !== undefined) {
      await expect(page.getByTestId(`guild-loot-item-${unawarded.item_id}`).locator('button')).toHaveCount(
        unawarded.candidates.length,
      );
    }
  });

  test('a member sees every ranked candidate but no Award button -- read-only', async ({ page }) => {
    await stubCore(page, VIEWER_MEMBER, 'member');
    await stubDepthEndpoints(page);
    await page.goto(`${GUILD_URL}#loot`);
    await expect(page.getByTestId('guild-loot-tab')).toBeVisible();
    await expect(page.getByTestId('guild-loot-tab').getByText('Read-only')).toBeVisible();
    const loot = buildMockLoot();
    for (const item of loot.items) {
      await expect(page.getByTestId(`guild-loot-item-${item.item_id}`)).toBeVisible();
    }
    await expect(page.locator('[data-testid^="guild-loot-award-"]')).toHaveCount(0);
  });
});

// Live-fix round, defect 5: the Loot tab rendered nothing for the leader on real data --
// GuildLoot.svelte called `candidate.gain_dps.toFixed(0)` unconditionally, which threw for
// every tier-1 (fallback) candidate (`gain_dps: null`), the shape every real candidate on
// a live roster carries today since no BiS file names a raid-tier item yet. The fixture
// now matches the live contract's own 22-item table and tier-1 candidate shape exactly
// (mock-guild.ts), and these two tests reproduce the officer's own repro steps: mount
// Overview (which prefetches .../loot for its own summary card) then switch to Loot.
test.describe('loot tab renders real data after Overview’s own prefetch', () => {
  test('officer: Overview mounts first, then Loot renders every item with its candidates', async ({
    page,
  }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await stubDepthEndpoints(page);
    await page.goto(GUILD_URL);
    await expect(page.getByTestId('guild-overview-tab')).toBeVisible();
    // The Overview's own Loot summary card prefetches GET .../loot before the tab exists.
    await expect(page.getByTestId('guild-overview-card-loot')).toBeVisible();

    await page.getByTestId('guild-tab-loot').click();
    await expect(page.getByTestId('guild-loot-tab')).toBeVisible();

    const loot = buildMockLoot();
    // The API lists a shared name once: 22 drops, 21 rows.
    expect(loot.items).toHaveLength(21);
    for (const item of loot.items) {
      await expect(page.getByTestId(`guild-loot-item-${item.item_id}`)).toBeVisible();
      if (item.slot === '') continue; // a quest/reputation/crafting drop, no candidates
      const panel = page.getByTestId(`guild-loot-item-${item.item_id}`);
      await expect(panel).not.toContainText('NaN');
      for (const candidate of item.candidates) {
        await expect(panel).toContainText(candidate.name);
      }
    }
  });

  test('a 500 from .../loot shows an error line, never the 404 "not live yet" copy', async ({ page }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await page.route(`**/v1/guilds/${GUILD_ID}/raids*`, (route) => route.fulfill(envelope(buildMockRaids())));
    await page.route(`**/v1/guilds/${GUILD_ID}/progression`, (route) =>
      route.fulfill(envelope(buildMockProgression())),
    );
    await page.route(`**/v1/guilds/${GUILD_ID}/readiness`, (route) =>
      route.fulfill(envelope(buildMockReadiness())),
    );
    await page.route(`**/v1/guilds/${GUILD_ID}/loot*`, (route) =>
      route.fulfill({
        status: 500,
        contentType: 'application/json',
        body: '{"ok":false,"data":null,"error":{"message":"the loot ranker fell over"},"request_id":"r"}',
      }),
    );
    await page.goto(`${GUILD_URL}#loot`);
    await expect(page.getByTestId('guild-loot-error')).toBeVisible();
    await expect(page.getByTestId('guild-loot-error')).toContainText('the loot ranker fell over');
    await expect(page.getByTestId('guild-loot-missing')).toHaveCount(0);
    await expect(page.getByText(/isn.t live yet/)).toHaveCount(0);
  });
});

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

test.describe('the Overview raid-night row at 390 (round 8)', () => {
  test('every fact is visible on its own line and the page does not scroll sideways', async ({ page }) => {
    await page.setViewportSize({ width: 390, height: 900 });
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await stubDepthEndpoints(page);
    await page.goto(GUILD_URL);
    const row = page.getByTestId('guild-overview-raid-row').first();
    await expect(row).toBeVisible();
    const facts = row.getByTestId('guild-overview-raid-facts');
    await expect(facts).toContainText('boss');
    const pieces = facts.locator(':scope > span');
    const n = await pieces.count();
    expect(n).toBeGreaterThanOrEqual(2);
    for (let i = 0; i < n; i++) {
      await expect(pieces.nth(i)).toBeVisible();
      const box = await pieces.nth(i).boundingBox();
      expect(box && box.x >= 0 && box.x + box.width <= 390).toBe(true);
    }
    // The facts sit on a line of their own under the title and the day.
    const [title, factsBox] = await Promise.all([
      row.getByTestId('guild-overview-raid-row-label').boundingBox(),
      facts.boundingBox(),
    ]);
    expect(title && factsBox && factsBox.y >= title.y + title.height - 1).toBe(true);
    expect(await page.evaluate(() => document.documentElement.scrollWidth)).toBeLessThanOrEqual(390);
  });

  test('at 1280 the row is a settings row: the title and facts at the left, the day at the right', async ({
    page,
  }) => {
    await page.setViewportSize({ width: 1280, height: 900 });
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await stubDepthEndpoints(page);
    await page.goto(GUILD_URL);
    const row = page.getByTestId('guild-overview-raid-row').first();
    await expect(row).toBeVisible();
    await expect(row.getByTestId('guild-overview-raid-facts')).toContainText(' · ');
    const box = await row.boundingBox();
    expect(box?.height ?? 99).toBeLessThan(80);
    const day = await row.getByTestId('settings-row-value').boundingBox();
    expect(day && box && day.x > box.x + box.width / 2).toBe(true);
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
  test('an officer of an unclaimed guild sees "Claim this guild"', async ({ page }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) =>
      route.fulfill(
        envelope({ ...buildMockHome(VIEWER_OFFICER), claim: { state: 'unclaimed', frozen: false } }),
      ),
    );
    await page.goto(GUILD_URL);
    await expect(page.getByTestId('guild-claim-link')).toHaveText('Claim this guild');
    await expect(page.getByTestId('guild-settings-link')).toHaveCount(0);
  });

  test('an officer of a guild with a pending claim sees "Confirm the claim"', async ({ page }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) =>
      route.fulfill(
        envelope({ ...buildMockHome(VIEWER_OFFICER), claim: { state: 'pending', frozen: false } }),
      ),
    );
    await page.goto(GUILD_URL);
    await expect(page.getByTestId('guild-confirm-claim-link')).toHaveText('Confirm the claim');
  });

  // Fix round 1, item 2: "Contest this claim" moved out of the header band entirely and
  // into the Settings tab, under the claim block -- and it never reaches the account that
  // already holds the claim. This officer's own battletag ('Fixture#1234', meFixture's
  // default) does not match `claim.claimed_by_name` ('Kraggor', buildMockHome's own
  // fixture value), so they read as "a verified officer of a different account."
  test('"Contest this claim" never renders in the header, only in Settings for a non-claimant officer', async ({
    page,
  }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    let homeCalls = 0;
    await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) => {
      homeCalls += 1;
      const claim =
        homeCalls === 1
          ? { state: 'claimed' as const, frozen: false, claimed_by_name: 'Kraggor' }
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
    // This officer's own battletag matches buildMockHome's own `claimed_by_name`
    // ('Kraggor') -- they are the account that holds the claim.
    await stubCore(page, VIEWER_OFFICER, 'leader', 'Kraggor');
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
    const approveButton = page.locator('[data-testid="guild-roster-approve"][aria-label="Approve Rootgall"]');
    await approveButton.click();
    expect(approveCalled).toBe(true);
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
    await expect(page.getByTestId('guild-crest-default-caption')).toHaveText("Default: your faction's logo");
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
    expect(putCalled).toBe(true);
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
    expect(deleteCalled).toBe(true);
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
    test(`at ${width} the header, Overview and roster facts are body-face sentences with " · " spaced`, async ({
      page,
    }) => {
      await page.setViewportSize({ width, height: 900 });
      await stubCore(page, VIEWER_OFFICER, 'leader');
      await stubDepthEndpoints(page);
      await page.goto(GUILD_URL);
      const header = page.getByTestId('guild-header-facts');
      await expect(header).toBeVisible();
      expect(await words(header)).toMatch(/^PvP US · \d+ bosse?s? down · \d+ pulls?( · Updated .+)?$/);
      const rosterFacts = page.getByTestId('guild-overview-roster-facts');
      expect(await words(rosterFacts)).toMatch(
        /^\d+ raiders? with gear on the site · \d+ waiting for approval · \d+ below the rating floor$/,
      );
      const raidFacts = page.getByTestId('guild-overview-raid-facts').first();
      expect(await words(raidFacts)).toMatch(/ · /);
      for (const line of [header, rosterFacts, raidFacts]) {
        expect(await line.evaluate((el) => getComputedStyle(el).fontFamily)).not.toMatch(/mono/i);
      }
      // Each fact is one piece: no fact's box is split across two lines.
      for (const line of [header, rosterFacts, raidFacts]) {
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

    test(`at ${width} the roster stats are one sentence in the body face`, async ({ page }) => {
      await page.setViewportSize({ width, height: 900 });
      await stubCore(page, VIEWER_OFFICER, 'leader');
      await page.goto(`${GUILD_URL}#roster`);
      const chip = page.getByTestId('guild-roster-chip').first();
      await expect(chip).toBeVisible();
      const line = chip;
      expect(await line.evaluate((el) => getComputedStyle(el).fontFamily)).not.toMatch(/mono/i);
      const stats = page.getByTestId('guild-roster-facts').first();
      expect(await words(stats)).toMatch(/Item level \d+/);
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
    // The waiting card leads the tab, so the Roster card's head is what starts above the fold.
    const card = page.getByTestId('guild-roster-card');
    await expect(card).toBeVisible();
    const box = await card.boundingBox();
    expect(box).not.toBeNull();
    expect(box!.y).toBeLessThan(800);
  });

  test('the Overview prints no "no encounter list yet" line of ours', async ({ page }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await stubDepthEndpoints(page);
    await page.goto(GUILD_URL);
    await expect(page.getByTestId('guild-overview-card-progression')).toBeVisible();
    await expect(page.getByTestId('guild-overview-card-progression')).not.toContainText('no encounter list');
  });
});
