// web/tests/e2e/guild-cache-invalidation.spec.ts
// G.4 (lifecycle round 4, walk rows 3, 8c): a player's own guild action shows at once. The
// site keeps private answers in localStorage (`fs.q.*`, gated by the `fs_csrf` cookie hint),
// so these specs plant that cookie and walk across real page loads: the stale answer is
// persisted by the first page, and the page a mutation lands on must still show the new state.
import { expect, test, type Page } from '@playwright/test';
import { guildJoinCopy } from '../../src/lib/guild/copy';
import {
  GUILD_ID,
  VIEWER_MEMBER,
  VIEWER_OFFICER,
  buildMockGuildPage,
  buildMockHome,
} from '../../src/fixtures/guild/mock-guild';

const GUILD_PATH = '/guild/us/pvp/olympus-xxvii';
const INVITE_TOKEN = 'fixtureinvitetoken1';

function envelope(data: unknown, status = 200) {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'r' }),
  };
}

const FORBIDDEN = {
  status: 403,
  contentType: 'application/json',
  body: '{"ok":false,"data":null,"error":{"code":"forbidden","message":"you are not a member of that guild"},"request_id":"r"}',
};

async function signedInCookie(page: Page): Promise<void> {
  await page.context().addCookies([{ name: 'fs_csrf', value: 'test', url: 'http://localhost' }]);
}

function meWith(guilds: unknown[], name: string) {
  return {
    user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
    characters: [
      { key: `us/pvp/${name.toLowerCase()}`, region: 'us', ruleset: 'pvp', name, class: 'warrior' },
    ],
    guilds,
  };
}

test('joining by invite lands on the member view, not the cached stranger view', async ({ page }) => {
  await signedInCookie(page);
  let joined = false;
  let meCalls = 0;
  await page.route('**/v1/me', (route) => {
    meCalls += 1;
    return route.fulfill(
      envelope(
        meWith(
          joined
            ? [
                {
                  id: GUILD_ID,
                  region: 'us',
                  ruleset: 'pvp',
                  name: 'Olympus XXVII',
                  rank: 'member',
                  consent: 'gear',
                  verified: true,
                },
              ]
            : [],
          VIEWER_MEMBER,
        ),
      ),
    );
  });
  await page.route(`**/v1/guilds/us/pvp/olympus-xxvii`, (route) =>
    route.fulfill(envelope(buildMockGuildPage())),
  );
  await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) =>
    joined ? route.fulfill(envelope(buildMockHome(VIEWER_MEMBER))) : route.fulfill(FORBIDDEN),
  );
  // The invite card names the guild before the join (the public token lookup).
  await page.route(`**/v1/guild-invites/${INVITE_TOKEN}`, (route) =>
    route.fulfill(envelope({ guild: { id: GUILD_ID, region: 'us', ruleset: 'pvp', name: 'Olympus XXVII' } })),
  );
  await page.route(`**/v1/guilds/invite/${INVITE_TOKEN}/accept`, (route) => {
    joined = true;
    return route.fulfill(
      envelope({
        guild: { id: GUILD_ID, region: 'us', ruleset: 'pvp', name: 'Olympus XXVII' },
        rank: 'member',
      }),
    );
  });
  await page.route('**/v1/rankings/guilds*', (route) => route.fulfill(envelope({ rows: [] })));

  // 1. The stranger view is read, and /v1/me is remembered in localStorage.
  await page.goto(`${GUILD_PATH}#roster`);
  await expect(page.getByTestId('guild-roster-gate')).toContainText('Roster is for members of this guild');
  expect(
    await page.evaluate(() => Object.keys(localStorage).filter((k) => k.startsWith('fs.q.')).length),
  ).toBeGreaterThan(0);

  // 2. Join through the invite, then open the guild from the joined state.
  await page.goto(`/guild/invite/${INVITE_TOKEN}`);
  await page.getByTestId('guild-join-button').click();
  // The joined state reads the guild's plan from the fresh /v1/me (the entry carries no plan here).
  await expect(page.getByTestId('guild-join-plan')).toHaveText(guildJoinCopy.noPlanYet('Olympus XXVII'));
  await expect(page.getByTestId('guild-join-button')).toHaveCount(0);
  await page.getByTestId('guild-join-open').click();
  await expect(page).toHaveURL(/\/guild\/us\/pvp\/olympus-xxvii/);

  // 3. The member view, from the refetched answer: the Roster tab opens with rows, no gate.
  await page.getByTestId('guild-tab-roster').click();
  await expect(page.getByTestId('guild-roster-rows').locator('.guild-roster-row').first()).toBeVisible();
  await expect(page.getByTestId('guild-roster-gate')).toHaveCount(0);
  expect(meCalls).toBeGreaterThan(1);
});

test('approving a character changes the band count and survives the next visit', async ({ page }) => {
  await signedInCookie(page);
  const approved = new Set<string>();
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      envelope(
        meWith(
          [
            {
              id: GUILD_ID,
              region: 'us',
              ruleset: 'pvp',
              name: 'Olympus XXVII',
              rank: 'leader',
              consent: 'gear',
              verified: true,
              can: ['approve', 'remove', 'settings', 'award', 'billing', 'release'],
            },
          ],
          VIEWER_OFFICER,
        ),
      ),
    ),
  );
  await page.route(`**/v1/guilds/us/pvp/olympus-xxvii`, (route) =>
    route.fulfill(envelope(buildMockGuildPage())),
  );
  // The server's own state: each approval removes the row from the pending list.
  await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) => {
    const home = buildMockHome(VIEWER_OFFICER);
    const pending = (home.pending ?? []).filter((r) => !approved.has(r.character_key));
    const roster = home.roster.map((r) => (approved.has(r.character_key) ? { ...r, verified: true } : r));
    return route.fulfill(
      envelope({
        ...home,
        roster,
        pending,
        summary: { ...home.summary, waiting_for_approval: pending.length },
      }),
    );
  });
  await page.route(`**/v1/guilds/${GUILD_ID}/characters/us/pvp/*/approve`, (route) => {
    const slug = route.request().url().split('/').slice(-2)[0];
    approved.add(`us/pvp/${slug}`);
    return route.fulfill(envelope({ character_key: `us/pvp/${slug}`, status: 'approved' }));
  });
  await page.route('**/v1/rankings/guilds*', (route) => route.fulfill(envelope({ rows: [] })));

  await page.goto(GUILD_PATH);
  // The Overview's Roster card carries the waiting count (the officer strip left the head).
  const waiting = page.getByTestId('guild-overview-roster-facts');
  await expect(waiting).toContainText('waiting for approval');
  const first = Number(((await waiting.textContent()) ?? '').match(/(\d+) waiting for approval/)?.[1]);
  expect(first).toBeGreaterThan(1);

  await page.goto(`${GUILD_PATH}#roster`);
  await page.getByTestId('guild-roster-approve').first().click();
  await expect(page.getByTestId('guild-roster-approve')).toHaveCount(first - 1);

  // The next visit renders the new count, not the persisted pre-approval answer.
  await page.goto(GUILD_PATH);
  await expect(waiting).toContainText(`${first - 1} waiting for approval`);
});
