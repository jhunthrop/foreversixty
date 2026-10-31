// web/tests/e2e/guild-discord-bot.spec.ts
// Guild settings · the Forever Sixty bot (design/specs/2026-10-27-discord-bot.md §3.1, §4.1, §9
// items 1 to 7 and 9), API stubbed: the add link and the webhook link, each install return, the
// picker's five list states and its races, the installed state, the remove question, the webhook
// offer, removed in Discord, officer-read-only and frozen, axe, 44 px targets at 360 and the
// §3.1 height caps.
import type { Locator, Page } from '@playwright/test';
import { expect, test } from './support/test';
import { notificationsCopy as copy } from '../../src/lib/account/notifications-copy';
import { guildBotCopy } from '../../src/lib/guild/discord-bot-copy';
import { guildDiscordCopy } from '../../src/lib/guild/discord-copy';
import { GUILD_CONNECT_HREF } from './support/guild-discord';
import {
  BOT,
  BOT_INSTALL_HREF,
  CHANNELS,
  NOT_CONNECTED,
  RAID_CHANNEL_ID,
  SERVER_ID,
  SETTINGS_GUILD_PATH,
  botSettings,
  guildChannel,
  stubGuildBot,
  type GuildBotStubOptions,
} from './support/guild-discord-bot';
import { gotoHydrated } from './support/hydrated';
import { expectNoAxeViolations } from './support/officer-tools';
import { choose, fieldControl } from './support/select';

const SECTION = '#discord';
const PHONE = { width: 360, height: 800 };
const DESKTOP = { width: 1440, height: 1000 };
const MIN_TARGET = 44;
const RETURN = (flag: string): string => `${SETTINGS_GUILD_PATH}?discord_bot=${flag}#discord`;
/** §3.1's caps, phone (360) then desktop (1440); each carries the Card head's 62 px (site alignment). */
const CAPS = {
  nothing: [402, 312],
  picker: [462, 392],
  installed: [512, 392],
  'installed with warning': [582, 442],
  removed: [362, 292],
  'webhook only': [502, 392],
} as const;

const section = (page: Page): Locator => page.locator(SECTION);
const id = (page: Page, testid: string): Locator => section(page).getByTestId(testid);
const WEBHOOK_AND_BOT = botSettings(null, guildChannel({ bot_available: true, bot: null }));

async function open(page: Page, options: GuildBotStubOptions, url = SETTINGS_GUILD_PATH) {
  const seen = await stubGuildBot(page, options);
  await gotoHydrated(page, url, 'guild-settings');
  await expect(section(page).locator('h2')).toHaveText(guildDiscordCopy.heading);
  return seen;
}

test.describe('state 5, nothing connected', () => {
  test('the add link is the install route, the webhook link is the webhook route', async ({ page }) => {
    await open(page, { discord: botSettings(null) });
    await expect(id(page, 'bot-add')).toHaveText(guildBotCopy.add);
    await expect(id(page, 'bot-add')).toHaveAttribute('href', BOT_INSTALL_HREF);
    expect(await id(page, 'bot-add').getAttribute('target')).toBeNull();
    await expect(id(page, 'bot-explanation')).toHaveText(guildBotCopy.addExplanation);
    await expect(id(page, 'bot-webhook')).toHaveText(guildBotCopy.webhookPath);
    await expect(id(page, 'bot-webhook')).toHaveAttribute('href', GUILD_CONNECT_HREF);
    await expect(section(page)).toContainText(guildBotCopy.webhookPathNote);
    await expect(section(page).locator('a[class*="linear-gradient"]')).toHaveCount(1);
  });

  test('with the bot off, G1 draws exactly as shipped', async ({ page }) => {
    await open(page, { discord: { ...NOT_CONNECTED, bot_available: false } });
    await expect(id(page, 'discord-connect')).toHaveAttribute('href', GUILD_CONNECT_HREF);
    await expect(id(page, 'bot-add')).toHaveCount(0);
    await expect(section(page)).toContainText(guildDiscordCopy.explanation);
  });
});

test.describe('the install return', () => {
  test('installed: never state 5, the line, focus on the select, #discord kept with no query', async ({
    page,
  }) => {
    await page.addInitScript(() => {
      const seen = { add: false };
      (window as unknown as { __addSeen: typeof seen }).__addSeen = seen;
      new MutationObserver(() => {
        if (document.querySelector('[data-testid="bot-add"]') !== null) seen.add = true;
      }).observe(document, { childList: true, subtree: true });
    });
    const seen = await open(
      page,
      {
        discord: botSettings(null),
        after: botSettings({ channel_id: null, channel_name: null }),
        channels: CHANNELS,
      },
      RETURN('installed'),
    );
    await expect(section(page).getByRole('status')).toHaveText(guildBotCopy.returns.installed);
    await expect(fieldControl(section(page), 'bot-channel')).toBeFocused();
    await expect(page).toHaveURL(new RegExp(`${SETTINGS_GUILD_PATH}#discord$`));
    expect(
      await page.evaluate(() => (window as unknown as { __addSeen: { add: boolean } }).__addSeen.add),
    ).toBe(false);
    expect(seen.base.settingsReads).toBeGreaterThanOrEqual(2);
  });

  for (const [flag, role] of [
    ['denied', 'status'],
    ['failed', 'alert'],
  ] as const) {
    test(`${flag}: the line as ${role}, focus on the add control, query dropped`, async ({ page }) => {
      await open(page, { discord: botSettings(null) }, RETURN(flag));
      await expect(section(page).getByRole(role)).toHaveText(guildBotCopy.returns[flag]);
      await expect(id(page, 'bot-add')).toBeFocused();
      await expect(page).toHaveURL(new RegExp(`${SETTINGS_GUILD_PATH}#discord$`));
    });
  }

  test('taken names the guild from the read; the expired variant reads without it', async ({ page }) => {
    await open(
      page,
      { discord: botSettings(null, { bot_refusal: { guild_name: 'Mograine Raiders' } }) },
      RETURN('taken'),
    );
    await expect(section(page).getByRole('alert')).toHaveText(guildBotCopy.returns.taken('Mograine Raiders'));
    await expect(id(page, 'bot-add')).toBeFocused();
    const expired = await page.context().newPage();
    await open(expired, { discord: botSettings(null) }, RETURN('taken'));
    await expect(section(expired).getByRole('alert')).toHaveText(guildBotCopy.returns.takenExpired);
  });
});

test.describe('the picker', () => {
  const picker = botSettings({ channel_id: null, channel_name: null });

  test('ready: grouped options, the raid channel preselected, an empty press, a save', async ({ page }) => {
    const seen = await open(page, { discord: picker });
    await expect(id(page, 'bot-picker-server')).toContainText(`${guildBotCopy.picker.added} Olympus Discord`);
    const select = id(page, 'bot-channel');
    await expect(select).toHaveValue(RAID_CHANNEL_ID);
    // The list's order (the native list cannot be screenshotted open): the uncategorised channels
    // first, then one optgroup per Discord category in Discord's order.
    await expect(select.locator('optgroup')).toHaveCount(1);
    await expect(select.locator('optgroup')).toHaveAttribute('label', 'Olympus Raids');
    await expect(select.locator('option').nth(1)).toHaveText('#general');
    await expect(select.locator('optgroup option')).toHaveText(['#raid-signups', '#officers']);
    await choose(select, '');
    await id(page, 'bot-channel-save').click();
    await expect(id(page, 'bot-channel-empty-press')).toHaveText(guildBotCopy.picker.emptyPress);
    await expect(select).toHaveAttribute('aria-invalid', 'true');
    await choose(select, RAID_CHANNEL_ID);
    await id(page, 'bot-channel-save').click();
    await expect(section(page).getByRole('status')).toHaveText(guildBotCopy.installed.saved);
    await expect(id(page, 'bot-test')).toBeFocused();
    expect(await id(page, 'bot-test').evaluate((el) => el.matches(':focus-visible'))).toBe(true);
    await expect(id(page, 'bot-destination')).toContainText('#raid-signups');
    expect(seen.saves).toEqual([RAID_CHANNEL_ID]);
  });

  test('no raid channel: the placeholder shows and nothing is preselected', async ({ page }) => {
    await open(page, { discord: picker, channels: [CHANNELS[0]!, CHANNELS[2]!] });
    await expect(id(page, 'bot-channel')).toHaveValue('');
    await expect(id(page, 'bot-channel')).toContainText(guildBotCopy.picker.placeholder);
  });

  test('loading draws a disabled select that says so', async ({ page }) => {
    await stubGuildBot(page, { discord: picker });
    await page.route('**/discord/bot/channels', async (route) => {
      await new Promise((resolve) => setTimeout(resolve, 600));
      await route.fallback();
    });
    await gotoHydrated(page, SETTINGS_GUILD_PATH, 'guild-settings');
    await expect(id(page, 'bot-channel')).toBeDisabled();
    await expect(id(page, 'bot-channel')).toContainText(guildBotCopy.picker.loading);
    await expect(id(page, 'bot-channel')).toHaveAttribute('aria-busy', 'true');
    await expect(id(page, 'bot-channel-save')).toBeDisabled();
    await expect(id(page, 'bot-channel-save')).toHaveClass(/opacity-50/);
    await expect(id(page, 'bot-channel')).toBeEnabled();
  });

  test('empty and error name themselves and retry', async ({ page }) => {
    await open(page, { discord: picker, channels: [] });
    await expect(id(page, 'bot-channels-empty')).toHaveText(guildBotCopy.picker.empty('Olympus Discord'));
    await expect(id(page, 'bot-channels-reload')).toHaveText(guildBotCopy.picker.reload);
    const failing = await page.context().newPage();
    await open(failing, { discord: picker, channels: 'error' });
    await expect(id(failing, 'bot-channels-error')).toHaveText(guildBotCopy.picker.loadFailed);
    await expect(id(failing, 'bot-channels-retry')).toHaveText(guildBotCopy.picker.retry);
  });

  test('a 422 channel_unusable refetches, clears the select and says so, focused', async ({ page }) => {
    const seen = await open(page, { discord: picker, unusableSaves: 1 });
    await id(page, 'bot-channel-save').click();
    await expect(section(page).getByRole('alert')).toHaveText(guildBotCopy.picker.race);
    await expect(page.locator('#guild-discord-status')).toBeFocused();
    await expect(id(page, 'bot-channel')).toHaveValue('');
    expect(seen.channelReads).toBe(2);
  });

  test('Remove the bot asks first and goes to state 5', async ({ page }) => {
    const seen = await open(page, { discord: picker });
    await id(page, 'bot-remove-bot').click();
    await expect(id(page, 'bot-remove-question')).toHaveText(
      guildBotCopy.remove.question('Olympus Discord', false),
    );
    await expect(id(page, 'bot-remove-cancel')).toBeFocused();
    await id(page, 'bot-remove-confirm').click();
    await expect(id(page, 'bot-add')).toBeFocused();
    await expect(section(page).getByRole('status')).toHaveText(copy.removed);
    expect(seen.botDeletes).toBe(1);
  });
});

test.describe('installed', () => {
  test('names channel and server, links to the channel, counts linked members', async ({ page }) => {
    await open(page, { discord: botSettings({}) });
    await expect(id(page, 'bot-destination')).toContainText(
      `${guildBotCopy.installed.lead} #raid-signups ${guildBotCopy.installed.in} Olympus Discord`,
    );
    await expect(id(page, 'bot-open')).toHaveAttribute(
      'href',
      `https://discord.com/channels/${SERVER_ID}/${RAID_CHANNEL_ID}`,
    );
    await expect(id(page, 'bot-open')).toHaveAttribute('rel', 'noopener');
    await expect(id(page, 'bot-linked-members')).toHaveText(guildBotCopy.installed.linkedMembers(9, 31));
    await expect(
      section(page).locator(
        'button.bg-gold, a.bg-gold, button[class*="linear-gradient"], a[class*="linear-gradient"]',
      ),
    ).toHaveCount(0);
  });

  test('the linked-members line in its three forms', async ({ page }) => {
    for (const linked of [1, 0]) {
      await open(page, { discord: botSettings({ linked_members: linked }) });
      await expect(id(page, 'bot-linked-members')).toHaveText(
        guildBotCopy.installed.linkedMembers(linked, BOT.members),
      );
    }
  });

  test('three failures warn; a test clears them and says to check the channel', async ({ page }) => {
    const seen = await open(page, {
      discord: botSettings({ failures: 3, last_error: 'discord status 403 code 50013' }),
    });
    const warning = id(page, 'bot-warning');
    await expect(warning).toContainText(guildBotCopy.failureWarning.failures(3, 'raid-signups'));
    await expect(warning).toContainText(guildBotCopy.failureWarning.permissions('raid-signups'));
    await expect(warning).not.toContainText('switches off');
    await id(page, 'bot-test').click();
    await expect(section(page).getByRole('status')).toHaveText(copy.testSent);
    await expect(id(page, 'bot-warning')).toHaveCount(0);
    expect(seen.botTests).toBe(1);
  });

  test('Change channel reopens the picker with the current channel and Cancel', async ({ page }) => {
    await open(page, { discord: botSettings({}) });
    await id(page, 'bot-change').click();
    await expect(id(page, 'bot-picker-server')).toContainText(
      `${guildBotCopy.picker.serverLead} Olympus Discord`,
    );
    await expect(id(page, 'bot-channel')).toHaveValue(RAID_CHANNEL_ID);
    await expect(fieldControl(section(page), 'bot-channel')).toBeFocused();
    await id(page, 'bot-channel-cancel').click();
    await expect(id(page, 'bot-change')).toBeFocused();
  });

  test('the remove question: Escape cancels; with a webhook it falls back to it', async ({ page }) => {
    const seen = await open(page, { discord: botSettings({}, guildChannel({ bot_available: true })) });
    await expect(id(page, 'bot-fallback-note')).toHaveText(guildBotCopy.installed.fallbackNote);
    await id(page, 'bot-remove').click();
    await expect(id(page, 'bot-remove-question')).toHaveText(
      guildBotCopy.remove.question('Olympus Discord', true),
    );
    await page.keyboard.press('Escape');
    await expect(id(page, 'bot-test')).toBeVisible();
    await id(page, 'bot-remove').click();
    await id(page, 'bot-remove-confirm').click();
    await expect(id(page, 'bot-offer')).toBeVisible();
    await expect(id(page, 'bot-add')).toBeFocused();
    await expect(section(page).getByRole('status')).toHaveText(copy.removed);
    expect(seen.botDeletes).toBe(1);
  });
});

test.describe('webhook only', () => {
  test('the offer is an outline button and the webhook controls still work', async ({ page }) => {
    const seen = await open(page, { discord: WEBHOOK_AND_BOT });
    await expect(id(page, 'bot-offer')).toContainText(guildBotCopy.offer);
    await expect(id(page, 'bot-add')).toHaveAttribute('href', BOT_INSTALL_HREF);
    await expect(
      section(page).locator(
        'button.bg-gold, a.bg-gold, button[class*="linear-gradient"], a[class*="linear-gradient"]',
      ),
    ).toHaveCount(0);
    await id(page, 'discord-test').click();
    await expect(section(page).getByRole('status')).toHaveText(copy.testSent);
    expect(seen.base.tests).toBe(1);
  });
});

test.describe('removed in Discord', () => {
  test('with and without a webhook', async ({ page }) => {
    await open(page, {
      discord: botSettings({ removed_in_discord: true }, guildChannel({ bot_available: true })),
    });
    await expect(id(page, 'discord-notice')).toContainText(
      guildBotCopy.removedInDiscord.title('Olympus Discord'),
    );
    await expect(id(page, 'discord-notice')).toContainText(guildBotCopy.removedInDiscord.withWebhook.trim());
    await expect(id(page, 'bot-add')).toHaveAttribute('href', BOT_INSTALL_HREF);
    const bare = await page.context().newPage();
    await open(bare, { discord: botSettings({ removed_in_discord: true }) });
    await expect(id(bare, 'discord-notice')).toContainText(
      guildBotCopy.removedInDiscord.withoutWebhook.trim(),
    );
  });
});

test.describe('officer-read-only and frozen', () => {
  test('state 5: add and webhook links are aria-disabled with no href', async ({ page }) => {
    await open(page, { discord: botSettings(null), officerTools: false });
    for (const testid of ['bot-add', 'bot-webhook']) {
      await expect(id(page, testid)).toHaveAttribute('aria-disabled', 'true');
      expect(await id(page, testid).getAttribute('href')).toBeNull();
      await expect(id(page, testid)).toHaveAttribute('aria-describedby', 'officer-tools-gate-discord');
    }
  });

  test('installed: test and change are disabled, the line is drawn, Remove works', async ({ page }) => {
    const seen = await open(page, { discord: botSettings({}), officerTools: false });
    await expect(id(page, 'bot-test')).toBeDisabled();
    await expect(id(page, 'bot-change')).toBeDisabled();
    await expect(id(page, 'bot-read-only')).toHaveText(guildBotCopy.installed.readOnly);
    await id(page, 'bot-remove').click();
    await id(page, 'bot-remove-confirm').click();
    await expect(id(page, 'bot-add')).toHaveAttribute('aria-disabled', 'true');
    expect(seen.botDeletes).toBe(1);
  });

  test('the picker: select and save are disabled, Remove the bot is live', async ({ page }) => {
    await open(page, { discord: botSettings({ channel_id: null, channel_name: null }), officerTools: false });
    await expect(id(page, 'bot-channel')).toBeDisabled();
    await expect(id(page, 'bot-channel-save')).toBeDisabled();
    await expect(id(page, 'bot-remove-bot')).toBeEnabled();
  });

  test('frozen disables Remove too', async ({ page }) => {
    await open(page, { discord: botSettings({}), frozen: true });
    await expect(id(page, 'bot-test')).toBeDisabled();
    await expect(id(page, 'bot-remove')).toBeDisabled();
  });

  test('a 403 guild_plan_required says so, focused, and turns the section read-only', async ({ page }) => {
    await open(page, {
      discord: botSettings({}),
      testRefusal: { status: 403, code: 'guild_plan_required', message: 'plan' },
    });
    await id(page, 'bot-test').click();
    await expect(page.locator('#guild-discord-status')).toBeFocused();
    await expect(section(page).getByRole('alert')).toHaveText(guildBotCopy.planEnded);
    await expect(id(page, 'bot-test')).toBeDisabled();
    await expect(id(page, 'bot-remove')).toBeEnabled();
  });
});

test.describe('axe, 44 px targets at 360, and the §3.1 caps', () => {
  test.use({ reducedMotion: 'reduce' });
  const states: { name: string; options: GuildBotStubOptions; cap: readonly [number, number] }[] = [
    { name: 'nothing', options: { discord: botSettings(null) }, cap: CAPS.nothing },
    {
      name: 'picker',
      options: { discord: botSettings({ channel_id: null, channel_name: null }) },
      cap: CAPS.picker,
    },
    { name: 'installed', options: { discord: botSettings({}) }, cap: CAPS.installed },
    {
      name: 'installed with warning',
      options: { discord: botSettings({ failures: 3, last_error: 'discord status 403 code 50013' }) },
      cap: CAPS['installed with warning'],
    },
    {
      name: 'removed',
      options: { discord: botSettings({ removed_in_discord: true }) },
      cap: CAPS.removed,
    },
    { name: 'webhook only', options: { discord: WEBHOOK_AND_BOT }, cap: CAPS['webhook only'] },
  ];
  for (const { name, options, cap } of states) {
    test(name, async ({ page }) => {
      for (const [index, viewport] of [PHONE, DESKTOP].entries()) {
        await page.setViewportSize(viewport);
        await open(page, options, `${SETTINGS_GUILD_PATH}#discord`);
        await expect(page.locator('[aria-busy="true"]')).toHaveCount(0);
        await expectNoAxeViolations(page, SECTION);
        const height = await section(page).evaluate((node) => node.getBoundingClientRect().height);
        expect(height, `${name} @${viewport.width}`).toBeLessThanOrEqual(cap[index] ?? 0);
        if (viewport === PHONE) {
          const heights = await section(page)
            .locator('button, a[data-testid^="bot-"]:not([data-testid="bot-open"]), select:not([hidden])')
            .evaluateAll((nodes) => nodes.map((node) => node.getBoundingClientRect().height));
          for (const h of heights) expect(h).toBeGreaterThanOrEqual(MIN_TARGET);
        }
      }
    });
  }

  test('the remove question', async ({ page }) => {
    await open(page, { discord: botSettings({}) });
    await id(page, 'bot-remove').click();
    await expectNoAxeViolations(page, SECTION);
  });
});
