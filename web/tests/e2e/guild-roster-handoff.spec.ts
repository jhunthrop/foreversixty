// web/tests/e2e/guild-roster-handoff.spec.ts
// The Roster tab's Open in simulator / planner slot is decided at first paint from the home's
// `sim_export` (api sims.ExportIsInputSQL): a row whose export the simulator reads holds the slot
// from the start and keeps its height when the links land; a row without one never holds it, so
// no phone row carries a blank line and nothing moves when a lookup finishes.
import { expect, test, apiUnreachable, GUILD_OVERVIEW_READS } from './support/test';
import { GUILD_ID, VIEWER_MEMBER, buildMockHome } from '../../src/fixtures/guild/mock-guild';
import { gotoHydrated } from './support/hydrated';
import { heldRoute } from './support/held-route';
import { GUILD_PAGE, envelope, stubGuild } from './support/row-band-pages';
import { openRosterAccounts } from './support/guild-roster-accounts';

test.beforeEach(({ page }) => apiUnreachable(page, ...GUILD_OVERVIEW_READS));

const PHONE = { width: 390, height: 900 };
const FS1_CODE = 'FS1:1.60.1.70009:hunter:orc:0/0/0:head=12640';

test.use({ viewport: PHONE });
test.skip(() => test.info().project.name !== 'desktop', 'one width, set here');

test('a row with no export the simulator reads holds no handoff slot', async ({ page }) => {
  await stubGuild(page, 'member');
  await gotoHydrated(page, `${GUILD_PAGE}#roster`, 'guild-raid-roster');
  await openRosterAccounts(page);
  await expect(page.getByTestId('guild-roster-handoff')).toHaveCount(0);
});

test('a row whose export the simulator reads holds the slot at first paint, at its final height', async ({
  page,
}) => {
  await stubGuild(page, 'member');
  const home = buildMockHome(VIEWER_MEMBER);
  const withExport = {
    ...home,
    roster: home.roster.map((row) => (row.name === VIEWER_MEMBER ? { ...row, sim_export: true } : row)),
  };
  await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) => route.fulfill(envelope(withExport)));
  const lookup = await heldRoute(page, '**/v1/characters/us/pvp/zulmara/sim-input', (route) =>
    route.fulfill(envelope({ source: 'addon', gear: FS1_CODE, spec: '', talents: '', buffs: [] })),
  );

  await gotoHydrated(page, `${GUILD_PAGE}#roster`, 'guild-raid-roster');

  await openRosterAccounts(page);
  await lookup.started;
  const slots = page.getByTestId('guild-roster-handoff');
  await expect(slots).toHaveCount(1);
  const row = slots.locator('xpath=ancestor::div[contains(@class, "guild-roster-row")]');
  const before = await row.boundingBox();

  lookup.release();
  await expect(page.getByTestId('guild-roster-open-sim')).toBeVisible();
  const after = await row.boundingBox();
  expect(after?.height).toBe(before?.height);
});
