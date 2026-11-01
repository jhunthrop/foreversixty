// web/tests/e2e/guild-follows-selector.spec.ts
// The owner's 2026-11-01 report: on Employed Millennials' guild page the header selector moved
// to Obnoxious Yell and the page kept showing Employed Millennials. Your guild's page follows
// the selector (to the new character's guild, or to the landing for a guildless one); a guild
// the visitor looked up stays. The header's Guild menu follows the current character strictly,
// and your character's own page moves to the newly chosen character. The guild and character
// addresses are the ones the fixture build prerenders (lib/report/shell-paths.ts).
import type { Page } from '@playwright/test';
import { apiUnreachable, expect, test } from './support/test';
import { buildMockGuildPage } from '../../src/fixtures/guild/mock-guild';
import {
  armoryPointer,
  keyOf,
  signInWith,
  storePointer,
  type FixtureCharacter,
  type StoredPointer,
} from './support/selector';

const LAST_WATCH = { id: 11, name: 'The Last Watch', verified: true };
const IRON_VANGUARD = { id: 12, name: 'Iron Vanguard', verified: true };

const BOW: FixtureCharacter = {
  name: 'Bow',
  class: 'Hunter',
  spec: 'Marksmanship',
  ruleset: 'hardcore',
  guild: LAST_WATCH,
};
const IRONHIDE: FixtureCharacter = {
  name: 'Ironhide',
  class: 'Warrior',
  spec: 'Arms',
  ruleset: 'hardcore',
  guild: IRON_VANGUARD,
};
const OBNOXIOUS: FixtureCharacter = { name: 'Obnoxious Yell', class: 'Warrior', spec: 'Fury', level: 21 };

const LAST_WATCH_PATH = '/guild/us/hardcore/the-last-watch';
const LOOKED_UP_PATH = '/guild/us/pvp/olympus-xxvii';
const ELYRA_PATH = '/character/us/hardcore/elyra-duskvale';

function envelope(data: unknown) {
  return {
    status: 200,
    contentType: 'application/json',
    body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
  };
}

/** Every API read these pages make goes unanswered unless a test stubs it: they follow the address. */
async function quietApi(page: Page): Promise<void> {
  await apiUnreachable(page, '**/v1/guilds/**', '**/v1/characters/**', '**/v1/rankings/**', '**/v1/sims*');
}

async function stubAccount(page: Page): Promise<void> {
  await signInWith(
    page,
    [BOW, IRONHIDE, OBNOXIOUS],
    [
      { ...LAST_WATCH, region: 'us', ruleset: 'hardcore', verified_count: 1 },
      { ...IRON_VANGUARD, region: 'us', ruleset: 'hardcore', verified_count: 1 },
    ],
  );
  await page.route('**/v1/guilds/*/*/*', (route) => route.fulfill(envelope(buildMockGuildPage())));
  await quietApi(page);
}

const selectorRow = (page: Page, character: FixtureCharacter) =>
  page
    .getByTestId('selector-panel')
    .getByTestId(`selector-row-${keyOf(character)}`)
    .getByRole('button')
    .first();

async function choose(page: Page, character: FixtureCharacter): Promise<void> {
  await page.getByTestId('character-selector').click();
  await selectorRow(page, character).click();
  await expect(page.getByTestId('selector-name')).toHaveText(character.name);
}

test.describe('guild and character pages follow the header selector', () => {
  test.use({ viewport: { width: 1440, height: 900 } });
  test.skip(() => test.info().project.name !== 'desktop', 'desktop layout only');
  test.afterEach(async ({ page }) => {
    await page.unrouteAll({ behavior: 'ignoreErrors' });
  });

  test("your guild's page moves to the new character's guild, keeping the tab", async ({ page }) => {
    await stubAccount(page);
    await storePointer(page, armoryPointer(BOW));
    await page.goto(`${LAST_WATCH_PATH}#roster`);
    await expect(page.getByTestId('guild-band-header')).toBeVisible();

    await choose(page, IRONHIDE);

    await expect(page).toHaveURL(/\/guild\/us\/hardcore\/iron-vanguard#roster$/);
  });

  test('a guildless character goes to the find-or-claim landing, and the menu offers its doors', async ({
    page,
  }) => {
    await stubAccount(page);
    await storePointer(page, armoryPointer(BOW));
    await page.goto(LAST_WATCH_PATH);
    await expect(page.getByTestId('guild-band-header')).toBeVisible();

    await choose(page, OBNOXIOUS);

    await expect(page).toHaveURL(/\/guild$/);
    await expect(page.getByTestId('guild-landing')).toBeVisible();
    await expect(page.getByTestId('site-panel-row-roster')).toContainText('Find a guild');
    await expect(page.getByTestId('site-panel-foot-guild')).toHaveText('Claim your guild');
    // The landing does not send a guildless character back to the account's guild.
    await expect(page.locator('[data-site-header][data-session-guild="null"]')).toHaveCount(1);
    await expect(page).toHaveURL(/\/guild$/);
  });

  test('a guild the visitor looked up stays put while the Guild menu follows', async ({ page }) => {
    await stubAccount(page);
    await storePointer(page, armoryPointer(BOW));
    await page.goto(`${LOOKED_UP_PATH}#roster`);
    await expect(page.getByTestId('guild-band-header')).toBeVisible();

    await choose(page, IRONHIDE);

    await expect(page.getByTestId('site-panel-row-roster')).toHaveAttribute(
      'href',
      '/guild/us/hardcore/iron-vanguard#roster',
    );
    await expect(page).toHaveURL(new RegExp(`${LOOKED_UP_PATH}#roster$`));
  });

  test("your character's page moves to the newly chosen character", async ({ page }) => {
    await signInWith(page, [BOW, IRONHIDE]);
    await quietApi(page);
    const elyra: StoredPointer = {
      source: 'armory',
      ref: 'us/hardcore/elyra-duskvale',
      label: 'Elyra Duskvale · Holy Priest',
      classSlug: 'priest',
    };
    await storePointer(page, elyra);
    // This session's default already settled, so the stored character stays current on load.
    await page.addInitScript(() => window.sessionStorage.setItem('fs.pointerHealed', '1'));
    await page.goto(ELYRA_PATH);
    await expect(page.getByTestId('character-selector')).toBeVisible();

    await choose(page, IRONHIDE);

    await expect(page).toHaveURL(new RegExp(`/character/${keyOf(IRONHIDE)}$`));
  });
});
