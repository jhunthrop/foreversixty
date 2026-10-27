// web/tests/e2e/support/officer-tools.ts
// Stubs for the officer tools gate (design/specs/2026-10-25-officer-tools-gate.md), shared by
// tests/e2e/guild-officer-tools.spec.ts and the acceptance captures in
// tests/captures/officer-tools-gate/. The guild control centre is the seeded mock guild
// (src/fixtures/guild/mock-guild.ts); `officerTools` stubs `home.viewer.officer_tools`.
import { createRequire } from 'node:module';
import { expect, type Locator, type Page } from '@playwright/test';
import {
  GUILD,
  GUILD_ID,
  VIEWER_MEMBER,
  VIEWER_OFFICER,
  buildMockGuildPage,
  buildMockHome,
  buildMockLoot,
  buildMockProgression,
  buildMockRaids,
  buildMockReadiness,
} from '../../../src/fixtures/guild/mock-guild';
import type { GuildBillingView } from '../../../src/lib/account/api';
import {
  CLAIM_HOLDER_CAN,
  guildBilling,
  mockEntitlements,
  scopedGrant,
} from '../../../src/test-support/billing-fixtures';

export const GUILD_SLUG = 'olympus-xxvii';
export const GUILD_URL = `/guild/${GUILD.region}/${GUILD.ruleset}/${GUILD_SLUG}`;
export const GUILD_SETTINGS_PATH = `${GUILD_URL}/settings`;

export const REFUSAL_MESSAGE = 'Award and Undo are part of the guild plan.';
export const REFUSAL_CODE = 'guild_plan_required';

export type Viewer = 'officer' | 'member';

export interface GuildCentreStub {
  viewer: Viewer;
  /** `home.viewer.officer_tools`; left out, the field is absent. */
  officerTools?: boolean;
  /** `home.viewer.raid_guild_tools`; defaults to true, false is a Guild-tier guild. */
  raidGuildTools?: boolean;
  /** No boss killed yet this tier: the Loot tab's empty state. */
  emptyLoot?: boolean;
}

export function envelope(data: unknown, status = 200) {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'req-test' }),
  };
}

/** The API's error envelope, as account/api.ts and guild/api.ts read it. */
export function refusal(code: string, message: string, status: number) {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: false, data: null, error: { code, message }, request_id: 'req-test' }),
  };
}

function viewerName(viewer: Viewer): string {
  return viewer === 'officer' ? VIEWER_OFFICER : VIEWER_MEMBER;
}

/** `/v1/me` for a viewer of the mock guild: an officer carries the officer capabilities. */
function guildCentreMe(viewer: Viewer) {
  const name = viewerName(viewer);
  return {
    user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false, premium: false },
    characters: [
      { key: `us/pvp/${name.toLowerCase()}`, region: 'us', ruleset: 'pvp', name, class: 'warrior' },
    ],
    guilds: [
      {
        id: GUILD_ID,
        region: GUILD.region,
        ruleset: GUILD.ruleset,
        name: GUILD.name,
        rank: viewer,
        verified: true,
        can: viewer === 'officer' ? ['approve', 'remove', 'settings', 'award'] : [],
      },
    ],
  };
}

/** Routes the public page, /v1/me, the home and the four depth endpoints of the mock guild. */
export async function stubGuildCentre(page: Page, stub: GuildCentreStub): Promise<void> {
  const name = viewerName(stub.viewer);
  const home = buildMockHome(name, GUILD.faction, null, true, {
    officerTools: stub.officerTools,
    raidGuildTools: stub.raidGuildTools,
  });
  const loot = stub.emptyLoot === true ? { ...buildMockLoot(), items: [] } : buildMockLoot();
  await page.route(`**/v1/guilds/${GUILD.region}/${GUILD.ruleset}/${GUILD_SLUG}`, (route) =>
    route.fulfill(envelope(buildMockGuildPage())),
  );
  await page.route('**/v1/me', (route) => route.fulfill(envelope(guildCentreMe(stub.viewer))));
  await page.route(`**/v1/guilds/${GUILD_ID}/home`, (route) => route.fulfill(envelope(home)));
  await page.route(`**/v1/guilds/${GUILD_ID}/raids*`, (route) => route.fulfill(envelope(buildMockRaids())));
  await page.route(`**/v1/guilds/${GUILD_ID}/progression`, (route) =>
    route.fulfill(envelope(buildMockProgression())),
  );
  await page.route(`**/v1/guilds/${GUILD_ID}/readiness`, (route) =>
    route.fulfill(envelope(buildMockReadiness())),
  );
  await page.route(`**/v1/guilds/${GUILD_ID}/loot*`, (route) => route.fulfill(envelope(loot)));
}

/** Answers every Award and Undo with the 403 `guild_plan_required`; returns the requests seen.
 *  Registered after stubGuildCentre, so it wins over the `loot*` route for these URLs. */
export async function refuseAwards(page: Page): Promise<string[]> {
  const seen: string[] = [];
  await page.route(`**/v1/guilds/${GUILD_ID}/loot/awards**`, (route) => {
    seen.push(route.request().method());
    return route.fulfill(refusal(REFUSAL_CODE, REFUSAL_MESSAGE, 403));
  });
  return seen;
}

/** The mock loot's first item with a candidate and no award, and its first candidate. */
export function unawardedItem(): { itemId: number; characterKey: string } {
  const item = buildMockLoot().items.find((i) => i.awarded_to === null && i.candidates.length > 0);
  if (item === undefined) throw new Error('the mock loot lost its unawarded items');
  return { itemId: item.item_id, characterKey: item.candidates[0].character_key };
}

/** The mock loot's awarded item (its winner row carries the one Undo). */
export function awardedItemId(): number {
  const item = buildMockLoot().items.find((i) => i.awarded_to !== null);
  if (item === undefined) throw new Error('the mock loot lost its awarded item');
  return item.item_id;
}

export const awardButtons = (page: Page): Locator => page.locator('button[data-testid^="guild-loot-award-"]');
export const undoButtons = (page: Page): Locator =>
  page.locator('button[data-testid^="guild-loot-row-undo-"]');
export const nudgeButtons = (page: Page): Locator =>
  page.locator('button[data-testid^="guild-readiness-nudge-"]');

/** Two guilds on the account page: Olympus XXVII on the scoped grant, The Last Watch on the paid plan. */
/** The account page's third guild: no plan, so its row carries no plan line. */
const NO_PLAN_GUILD_ID = 502;

/** `/v1/me` with three guild rows: a scoped grant, a paid plan, and a guild with no plan. */
export function accountMeWithGuildRows() {
  const row = (
    id: number,
    region: string,
    ruleset: string,
    name: string,
    plan: GuildBillingView | null,
  ): Record<string, unknown> => ({
    id,
    region,
    ruleset,
    name,
    rank: 'officer',
    consent: 'gear',
    verified: true,
    can: CLAIM_HOLDER_CAN,
    plan,
  });
  return {
    ok: true,
    data: {
      user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
      characters: [],
      guilds: [
        row(GUILD_ID, GUILD.region, GUILD.ruleset, GUILD.name, scopedGrant()),
        row(
          501,
          'us',
          'hardcore',
          'The Last Watch',
          guildBilling({ current_period_end: '2026-12-31T00:00:00Z' }),
        ),
        row(NO_PLAN_GUILD_ID, 'us', 'hardcore', 'Ashen Vigil', null),
      ],
      entitlements: mockEntitlements(),
      notifications: { discord: null },
    },
    error: null,
    request_id: 'r',
  };
}

/** Routes /v1/me with the three guild rows and the device list the account page reads. */
export async function stubAccountGuildRows(page: Page): Promise<void> {
  await page.route('**/v1/me', (route) => route.fulfill(envelope(accountMeWithGuildRows().data)));
  await page.route('**/v1/devices', (route) => route.fulfill(envelope([])));
}

const AXE_SOURCE = createRequire(import.meta.url).resolve('axe-core/axe.min.js');

interface AxeResults {
  violations: { id: string; nodes: { target: string[] }[] }[];
}

/** Runs axe on the element matching `selector` and expects no violations. */
export async function expectNoAxeViolations(page: Page, selector: string): Promise<void> {
  await page.evaluate(() => Promise.all(document.getAnimations().map((animation) => animation.finished)));
  await page.addScriptTag({ path: AXE_SOURCE });
  const results = await page.evaluate(async (context): Promise<AxeResults> => {
    const axe = (window as unknown as { axe: { run: (context: string) => Promise<AxeResults> } }).axe;
    return axe.run(context);
  }, selector);
  expect(results.violations.map((v) => ({ id: v.id, targets: v.nodes.map((n) => n.target) }))).toEqual([]);
}
