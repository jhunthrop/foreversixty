// web/tests/e2e/guild-settings-discord.spec.ts
// Guild settings · Discord (design/specs/2026-10-26-guild-raids.md §3.8, §9's Settings · Discord
// paragraph), API stubbed: the connect link on the API origin, the three returns from Discord, test,
// the remove question, officer-read-only (Remove stays live, `#billing` scrolls to Billing), the
// 403 that ends the officer tools, the switched-off fill, a non-officer's nothing, the raid form's
// line after a remove, axe in every state, 44 px controls at 360 and the §3.8 height caps.
import type { Locator, Page } from '@playwright/test';
import { apiUnreachable, expect, test } from './support/test';
import { notificationsCopy as copy } from '../../src/lib/account/notifications-copy';
import { guildDiscordCopy } from '../../src/lib/guild/discord-copy';
import { GUILD_ID } from '../../src/fixtures/guild/mock-guild';
import { CLAIM_HOLDER_CAN, guildBilling } from '../../src/test-support/billing-fixtures';
import { GUILD_URL, raidsMe, stubGuildRaids } from './support/guild-raids';
import {
  DISCORD_SETTINGS_URL,
  GUILD_CONNECT_HREF,
  NOT_CONNECTED,
  SETTINGS_GUILD_PATH,
  SWITCHED_OFF,
  guildChannel,
  settingsBody,
  stubGuildDiscord,
  type GuildDiscordStubOptions,
} from './support/guild-discord';
import { gotoHydrated } from './support/hydrated';
import { envelope, expectNoAxeViolations } from './support/officer-tools';
import { fieldBox } from './support/select';

// Reads this spec does not answer go unanswered, as the API out of reach (support/test.ts).
test.beforeEach(({ page }) => apiUnreachable(page, '**/v1/guilds/*/billing'));

const SECTION = '#discord';
const PHONE = { width: 360, height: 800 };
const DESKTOP = { width: 1440, height: 1000 };
const MIN_TARGET = 44;
/** §3.8's caps, phone (360) then desktop (1440); each carries the Card head's 62 px (site alignment). */
const CAPS = {
  'not-connected': [272, 224],
  connected: [362, 296],
  'switched-off': [502, 392],
} as const;

const section = (page: Page): Locator => page.locator(SECTION);

async function open(page: Page, options: GuildDiscordStubOptions, url = SETTINGS_GUILD_PATH) {
  const seen = await stubGuildDiscord(page, options);
  await gotoHydrated(page, url, 'guild-settings');
  await expect(section(page).locator('h2')).toHaveText(guildDiscordCopy.heading);
  return seen;
}

test.describe('not connected', () => {
  test('the connect link is the guild route on the API origin, same tab, under Billing', async ({ page }) => {
    await open(page, { discord: NOT_CONNECTED });
    const connect = section(page).getByTestId('discord-connect');
    await expect(connect).toHaveText(copy.connect);
    await expect(connect).toHaveAttribute('href', GUILD_CONNECT_HREF);
    expect(await connect.getAttribute('target')).toBeNull();
    await expect(section(page)).toContainText(guildDiscordCopy.explanation);
    const order = await page.evaluate(() =>
      ['guild-settings-billing', 'guild-settings-discord', 'guild-visibility-select'].map((id) => {
        const node = document.querySelector(`[data-testid="${id}"]`);
        return node === null ? -1 : [...document.querySelectorAll('*')].indexOf(node);
      }),
    );
    expect(order[0]).toBeGreaterThan(-1);
    expect(order[0]).toBeLessThan(order[1]);
    expect(order[1]).toBeLessThan(order[2]);
  });

  test('a member gets no discord key and no section', async ({ page }) => {
    await stubGuildDiscord(page, { discord: null });
    await gotoHydrated(page, SETTINGS_GUILD_PATH, 'guild-settings');
    await expect(fieldBox(page, 'guild-visibility-select')).toBeVisible();
    await expect(page.locator(SECTION)).toHaveCount(0);
    await expect(page.getByTestId('discord-connect')).toHaveCount(0);
  });
});

test.describe('the return from Discord', () => {
  test('connected: never the not-connected state, the line, focus on Send test message, #discord kept', async ({
    page,
  }) => {
    await page.addInitScript(() => {
      const seen = { connect: false };
      (window as unknown as { __connectSeen: typeof seen }).__connectSeen = seen;
      new MutationObserver(() => {
        if (document.querySelector('[data-testid="discord-connect"]') !== null) seen.connect = true;
      }).observe(document, { childList: true, subtree: true });
    });
    const seen = await open(
      page,
      { discord: NOT_CONNECTED, after: guildChannel() },
      `${SETTINGS_GUILD_PATH}?discord=connected#discord`,
    );
    await expect(section(page).getByRole('status')).toHaveText(guildDiscordCopy.returnConnected);
    await expect(section(page).getByTestId('discord-test')).toBeFocused();
    await expect(page).toHaveURL(new RegExp(`${SETTINGS_GUILD_PATH}#discord$`));
    expect(
      await page.evaluate(
        () => (window as unknown as { __connectSeen: { connect: boolean } }).__connectSeen.connect,
      ),
    ).toBe(false);
    expect(seen.settingsReads).toBeGreaterThanOrEqual(2);
    const top = await section(page).evaluate((node) => node.getBoundingClientRect().top);
    expect(top).toBeLessThan(200);
  });

  test('denied: the status line, focus back on the connect control', async ({ page }) => {
    await open(page, { discord: NOT_CONNECTED }, `${SETTINGS_GUILD_PATH}?discord=denied#discord`);
    await expect(section(page).getByRole('status')).toHaveText(guildDiscordCopy.returnDenied);
    await expect(section(page).getByTestId('discord-connect')).toBeFocused();
    await expect(page).toHaveURL(new RegExp(`${SETTINGS_GUILD_PATH}#discord$`));
    await page.reload();
    await expect(section(page).getByTestId('discord-connect')).toBeVisible();
    await expect(section(page).getByRole('status')).toHaveCount(0);
  });

  test('failed: the alert, focus back on the connect control used (the reconnect link)', async ({ page }) => {
    await open(page, { discord: guildChannel() }, `${SETTINGS_GUILD_PATH}?discord=failed#discord`);
    await expect(section(page).getByRole('alert')).toHaveText(guildDiscordCopy.returnFailed);
    await expect(section(page).getByTestId('discord-replace')).toBeFocused();
    await expect(page).toHaveURL(new RegExp(`${SETTINGS_GUILD_PATH}#discord$`));
  });
});

test.describe('connected', () => {
  test('test calls the route and says to check the channel', async ({ page }) => {
    const seen = await open(page, { discord: guildChannel() });
    await expect(section(page)).toContainText('Posts raids to a Discord channel');
    await section(page).getByTestId('discord-test').click();
    await expect(section(page).getByRole('status')).toHaveText(copy.testSent);
    expect(seen.tests).toBe(1);
  });

  test('the remove question: focus on Cancel, Escape cancels, Remove removes', async ({ page }) => {
    const seen = await open(page, { discord: guildChannel() });
    await section(page).getByTestId('discord-remove').click();
    await expect(section(page)).toContainText(guildDiscordCopy.removeQuestion);
    await expect(section(page).getByTestId('discord-remove-cancel')).toBeFocused();
    await expect(section(page).getByTestId('discord-replace')).toHaveCount(0);
    await page.keyboard.press('Escape');
    await expect(section(page).getByTestId('discord-test')).toBeVisible();
    await section(page).getByTestId('discord-remove').click();
    await section(page).getByTestId('discord-remove-confirm').click();
    await expect(section(page).getByTestId('discord-connect')).toBeFocused();
    await expect(section(page).getByRole('status')).toHaveText(copy.removed);
    expect(seen.deletes).toBe(1);
  });

  test('switched off: Test and switch on is the one fill, and a test switches posts back on', async ({
    page,
  }) => {
    await open(page, { discord: SWITCHED_OFF });
    await expect(
      section(page).locator(
        'button.bg-gold, a.bg-gold, button[class*="linear-gradient"], a[class*="linear-gradient"]',
      ),
    ).toHaveCount(1);
    await expect(section(page).getByTestId('discord-test')).toHaveText(copy.testAndSwitchOn);
    await section(page).getByTestId('discord-test').click();
    await expect(section(page).getByRole('status')).toHaveText(copy.testSentSwitchedOn);
    await expect(section(page).getByTestId('discord-notice')).toHaveCount(0);
  });

  test('a 403 guild_plan_required turns the section read-only and focuses the line', async ({ page }) => {
    await open(page, {
      discord: guildChannel(),
      testRefusal: {
        status: 403,
        code: 'guild_plan_required',
        message: 'Planning and running raids are part of the guild plan.',
      },
    });
    await section(page).getByTestId('discord-test').click();
    await expect(page.locator('#guild-discord-status')).toBeFocused();
    await expect(section(page).getByRole('alert')).toHaveText(guildDiscordCopy.planEnded);
    await expect(section(page).getByTestId('discord-test')).toBeDisabled();
    await expect(page.locator('#officer-tools-gate-discord')).toBeVisible();
    await expect(section(page).getByTestId('discord-remove')).toBeEnabled();
  });
});

test.describe('officer-read-only', () => {
  test('not connected: connect has aria-disabled and no href; Subscribe the guild scrolls to Billing', async ({
    page,
  }) => {
    await page.setViewportSize(PHONE);
    await open(page, { discord: NOT_CONNECTED, officerTools: false });
    const connect = section(page).getByTestId('discord-connect');
    await expect(connect).toHaveAttribute('aria-disabled', 'true');
    expect(await connect.getAttribute('href')).toBeNull();
    await expect(connect).toHaveAttribute('aria-describedby', 'officer-tools-gate-discord');
    await page.locator('#officer-tools-gate-discord a').click();
    await expect(page).toHaveURL(/#billing$/);
    await expect(page.getByTestId('guild-settings-billing')).toBeInViewport();
  });

  test('on the guild page Settings tab, Subscribe the guild scrolls to Billing and the tab stays', async ({
    page,
  }) => {
    await stubGuildRaids(page, { viewer: 'officer', officerTools: false });
    const me = raidsMe('officer');
    const billingMe = {
      ...me,
      guilds: me.guilds.map((g) => ({ ...g, can: [...CLAIM_HOLDER_CAN], plan: guildBilling() })),
    };
    await page.route('**/v1/me', (route) => route.fulfill(envelope(billingMe)));
    await page.route(`**/v1/guilds/${GUILD_ID}/settings`, (route) =>
      route.fulfill(envelope(settingsBody(NOT_CONNECTED, false))),
    );
    await gotoHydrated(page, `${GUILD_URL}#settings`, 'guild-settings-tab');
    await page.locator('#officer-tools-gate-discord a').click();
    await expect(page.getByTestId('guild-settings-tab')).toBeVisible();
    await expect(page.getByTestId('guild-settings-billing')).toBeInViewport();
    await expect(page.getByTestId('guild-tab-settings')).toHaveAttribute('aria-selected', 'true');
  });

  test('connected: test disabled, Remove still removes', async ({ page }) => {
    const seen = await open(page, { discord: guildChannel(), officerTools: false });
    await expect(section(page).getByTestId('discord-test')).toBeDisabled();
    await section(page).getByTestId('discord-remove').click();
    await section(page).getByTestId('discord-remove-confirm').click();
    await expect(section(page).getByTestId('discord-connect')).toHaveAttribute('aria-disabled', 'true');
    expect(seen.deletes).toBe(1);
  });

  test('frozen: every control is off, Remove included', async ({ page }) => {
    await open(page, { discord: guildChannel(), frozen: true });
    await expect(section(page).getByTestId('discord-test')).toBeDisabled();
    await expect(section(page).getByTestId('discord-remove')).toBeDisabled();
    expect(await section(page).getByTestId('discord-replace').getAttribute('href')).toBeNull();
  });
});

test('after a remove, the raid form draws the Settings line in place of the box', async ({ page }) => {
  const stub = await stubGuildRaids(page, { viewer: 'officer', discordConnected: true });
  await gotoHydrated(page, `${GUILD_URL}#raids`, 'guild-raids-tab');
  await page.getByTestId('guild-raids-plan').click();
  await expect(page.getByTestId('raid-form-discord')).toBeVisible();
  stub.discordConnected = false;
  await page.reload();
  await gotoHydrated(page, `${GUILD_URL}#raids`, 'guild-raids-tab');
  await page.getByTestId('guild-raids-plan').click();
  await expect(page.getByTestId('raid-form-discord')).toHaveCount(0);
  await expect(page.getByTestId('raid-form-settings-link')).toHaveAttribute(
    'href',
    `${GUILD_URL}/settings#discord`,
  );
});

test.describe('axe, 44 px targets at 360, and the §3.8 caps', () => {
  test.use({ reducedMotion: 'reduce' });
  /** The §3.8 cap, phone then desktop, for the states it is written for (the connected cap
   *  counts the state at rest; the 1 to 4 failure warning adds its two lines on top). */
  const states: { name: string; options: GuildDiscordStubOptions; cap?: readonly [number, number] }[] = [
    { name: 'not connected', options: { discord: NOT_CONNECTED }, cap: CAPS['not-connected'] },
    { name: 'connected', options: { discord: guildChannel() }, cap: CAPS.connected },
    {
      name: 'connected, 3 failures',
      options: { discord: guildChannel({ failures: 3, last_error: 'Missing Access' }) },
    },
    { name: 'switched off', options: { discord: SWITCHED_OFF }, cap: CAPS['switched-off'] },
    { name: 'read-only', options: { discord: guildChannel(), officerTools: false } },
  ];
  for (const { name, options, cap } of states) {
    test(name, async ({ page }) => {
      for (const [index, viewport] of [PHONE, DESKTOP].entries()) {
        await page.setViewportSize(viewport);
        await open(page, options, DISCORD_SETTINGS_URL);
        await expectNoAxeViolations(page, SECTION);
        if (cap !== undefined) {
          const height = await section(page).evaluate((node) => node.getBoundingClientRect().height);
          expect(height, `${name} @${viewport.width}`).toBeLessThanOrEqual(cap[index]);
        }
        if (viewport === PHONE) {
          const heights = await section(page)
            .locator('button, a[data-testid^="discord-"]:not([data-testid="discord-open"])')
            .evaluateAll((nodes) => nodes.map((node) => node.getBoundingClientRect().height));
          for (const h of heights) expect(h).toBeGreaterThanOrEqual(MIN_TARGET);
        }
      }
    });
  }

  test('the remove question', async ({ page }) => {
    await open(page, { discord: guildChannel() });
    await section(page).getByTestId('discord-remove').click();
    await expectNoAxeViolations(page, SECTION);
  });
});
