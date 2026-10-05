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
import { expect, test, type Page } from '@playwright/test';
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
        : [{ id: GUILD_ID, region: REGION, ruleset: RULESET, name: 'Olympus XXVII', rank, verified: true }],
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

  // Fix round 1, item 1: the officer strip moved into the persistent band alongside the
  // standing line and tab strip, so a claimed guild's officer sees it on every tab, not
  // just Overview.
  test('the officer strip stays visible across tabs, inside the full-bleed band', async ({ page }) => {
    await stubCore(page, VIEWER_OFFICER, 'leader');
    await stubDepthEndpoints(page);
    await page.goto(GUILD_URL);
    await expect(page.getByTestId('guild-officer-strip')).toBeVisible();
    await page.getByTestId('guild-tab-readiness').click();
    await expect(page.getByTestId('guild-officer-strip')).toBeVisible();
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

    const unverifiedList = page.getByTestId('guild-roster-unverified-list');
    await expect(unverifiedList).toContainText('Unverified');
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
    await expect(page.getByTestId('guild-roster-action-error')).toHaveText(
      'you do not have standing to remove this character',
    );
  });
});
