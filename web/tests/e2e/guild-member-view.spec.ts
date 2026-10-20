// web/tests/e2e/guild-member-view.spec.ts
// A guild member's way into their guild: the header's Roster and Guild settings doors lead
// to it (and to Find a guild / Claim your guild when the session has none), the guild page
// shows the roster to a verified member and one honest sentence to everyone else, and the
// claim entry sits under the header for any signed-in viewer who is not a verified officer.
import { expect, test, type Page } from '@playwright/test';
import {
  GUILD_ID,
  VIEWER_MEMBER,
  VIEWER_OFFICER,
  buildMockGuildPage,
  buildMockHome,
} from '../../src/fixtures/guild/mock-guild';

const GUILD_PATH = '/guild/us/pvp/olympus-xxvii';

function envelope(data: unknown) {
  return {
    status: 200,
    contentType: 'application/json',
    body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
  };
}

type Viewer = 'visitor' | 'stranger' | 'unverified' | 'member' | 'officer';

function meFor(viewer: Exclude<Viewer, 'visitor'>) {
  const name = viewer === 'officer' ? VIEWER_OFFICER : VIEWER_MEMBER;
  const guilds =
    viewer === 'stranger'
      ? []
      : [
          {
            id: GUILD_ID,
            region: 'us',
            ruleset: 'pvp',
            name: 'Olympus XXVII',
            rank: viewer === 'officer' ? 'leader' : 'member',
            consent: 'gear',
            verified: viewer !== 'unverified',
          },
        ];
  return {
    user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
    characters: [
      { key: `us/pvp/${name.toLowerCase()}`, region: 'us', ruleset: 'pvp', name, class: 'warrior' },
    ],
    guilds,
  };
}

async function stub(page: Page, viewer: Viewer): Promise<void> {
  await page.route('**/v1/guilds/us/pvp/olympus-xxvii', (route) =>
    route.fulfill(envelope(buildMockGuildPage())),
  );
  await page.route('**/v1/me', (route) =>
    viewer === 'visitor'
      ? route.fulfill({
          status: 401,
          contentType: 'application/json',
          body: '{"ok":false,"data":null,"error":null,"request_id":"r"}',
        })
      : route.fulfill(envelope(meFor(viewer))),
  );
  await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) =>
    viewer === 'member' || viewer === 'officer'
      ? route.fulfill(envelope(buildMockHome(viewer === 'officer' ? VIEWER_OFFICER : VIEWER_MEMBER)))
      : route.fulfill({
          status: 403,
          contentType: 'application/json',
          body: '{"ok":false,"data":null,"error":{"code":"forbidden","message":"you are not a member of that guild"},"request_id":"r"}',
        }),
  );
  await page.route('**/v1/rankings/guilds*', (route) => route.fulfill(envelope({ rows: [] })));
}

const rosterDoor = (page: Page) => page.getByTestId('site-panel-row-roster');
const settingsDoor = (page: Page) => page.getByTestId('site-panel-foot-guild');

test.describe('the header doors', () => {
  test('lead a member to their guild, even before an officer verifies them', async ({ page }) => {
    await stub(page, 'unverified');
    await page.goto('/');
    await expect(rosterDoor(page)).toHaveAttribute('href', `${GUILD_PATH}#roster`);
    await expect(settingsDoor(page)).toHaveAttribute('href', `${GUILD_PATH}/settings`);
    await expect(rosterDoor(page)).toContainText('Roster');
    await expect(settingsDoor(page)).toHaveText('Guild settings');
  });

  test('offer Find a guild and Claim your guild when signed out', async ({ page }) => {
    await stub(page, 'visitor');
    await page.goto('/');
    await expect(rosterDoor(page)).toHaveAttribute('href', '/guild');
    await expect(rosterDoor(page)).toContainText('Find a guild');
    await expect(settingsDoor(page)).toHaveAttribute('href', '/guild#claim');
    await expect(settingsDoor(page)).toHaveText('Claim your guild');
  });

  test('offer the same to a signed-in account with no guild', async ({ page }) => {
    await stub(page, 'stranger');
    await page.goto('/');
    await expect(rosterDoor(page)).toContainText('Find a guild');
    await expect(settingsDoor(page)).toHaveText('Claim your guild');
  });

  test('Claim your guild lands on the landing page claim sentence', async ({ page }) => {
    await stub(page, 'visitor');
    await page.goto('/guild#claim');
    await expect(page.getByTestId('guild-landing-claim')).toBeInViewport();
  });
});

test.describe('the guild page', () => {
  test('shows a verified member the roster at #roster, and the claim entry', async ({ page }) => {
    await stub(page, 'member');
    await page.goto(`${GUILD_PATH}#roster`);
    await expect(page.getByTestId('guild-roster-tab')).toBeVisible();
    await expect(page.getByTestId('guild-roster-rows').locator('.guild-roster-row').first()).toBeVisible();
    await expect(page.getByTestId('guild-roster-gate')).toHaveCount(0);
    await expect(page.getByTestId('guild-claim-entry')).toHaveText('Run this guild? Claim it.');
    await expect(page.getByTestId('guild-claim-entry-link')).toHaveAttribute('href', `${GUILD_PATH}/claim`);
  });

  test('tells a member still waiting on an officer what opens the roster, with no rows', async ({ page }) => {
    await stub(page, 'unverified');
    await page.goto(`${GUILD_PATH}#roster`);
    const gate = page.getByTestId('guild-roster-gate');
    await expect(gate).toContainText('waiting for an officer to verify');
    await expect(gate).toBeInViewport();
    await expect(page.getByTestId('guild-roster-rows')).toHaveCount(0);
    await expect(page.getByTestId('guild-claim-entry-link')).toBeVisible();
  });

  test('tells a visitor the roster is for members, with no rows and no claim entry', async ({ page }) => {
    await stub(page, 'visitor');
    await page.goto(`${GUILD_PATH}#roster`);
    await expect(page.getByTestId('guild-roster-gate')).toContainText(
      'Roster is for members. Sign in with the account your addon syncs to.',
    );
    await expect(page.getByTestId('guild-roster-rows')).toHaveCount(0);
    await expect(page.getByTestId('guild-claim-entry')).toHaveCount(0);
  });

  test('tells a signed-in account from another guild the roster is not theirs', async ({ page }) => {
    await stub(page, 'stranger');
    await page.goto(GUILD_PATH);
    await expect(page.getByTestId('guild-roster-gate')).toContainText('Roster is for members of this guild');
    await expect(page.getByTestId('guild-claim-entry-link')).toBeVisible();
  });

  test('gives a verified officer Guild settings in place of the claim entry', async ({ page }) => {
    await stub(page, 'officer');
    await page.goto(GUILD_PATH);
    await expect(page.getByTestId('guild-settings-link')).toHaveAttribute('href', `${GUILD_PATH}/settings`);
    await expect(page.getByTestId('guild-claim-entry')).toHaveCount(0);
    await expect(page.getByTestId('guild-roster-gate')).toHaveCount(0);
  });
});
