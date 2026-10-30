// web/tests/e2e/guild-discord-bot-unusable.spec.ts
// Guild settings · state 2b, the raid channel can't be used (design/specs/2026-10-27-discord-bot.md,
// "Amendment 2026-10-27: channel can't be used", A.5 items 11 to 17), API stubbed: the deleted and
// permissions notices, precedence, fixing by picking and by saving the same channel, the read-only
// and frozen gates, axe, 44 px targets at 360, the height caps and no layout shift.
import { expect, test, type Locator, type Page } from '@playwright/test';
import type { GuildChannelPermission, GuildChannelUnusable } from '../../src/lib/guild/api';
import { guildBotCopy, joinPermissions } from '../../src/lib/guild/discord-bot-copy';
import { guildDiscordCopy } from '../../src/lib/guild/discord-copy';
import {
  CHANNELS,
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
/** A.2's caps, phone then desktop. */
const CAP = [382, 302] as const; // 320 / 240 plus the Card head's 62 px (site alignment)
/** The reserves of 2b and state 2 differ by design (320 / 330 and 240 / 250). */
const RESERVE_DELTA = 10;
const ALL_PERMISSIONS: GuildChannelPermission[] = ['View Channel', 'Send Messages', 'Embed Links'];
const DELETED: GuildChannelUnusable = { reason: 'deleted', missing: [] };
const permissions = (missing: GuildChannelPermission[]): GuildChannelUnusable => ({
  reason: 'permissions',
  missing,
});
const WITH_WEBHOOK = guildChannel({ bot_available: true });

const section = (page: Page): Locator => page.locator(SECTION);
const id = (page: Page, testid: string): Locator => section(page).getByTestId(testid);
const unusable = (channel_unusable: GuildChannelUnusable, extra: Parameters<typeof botSettings>[0] = {}) =>
  botSettings({ channel_unusable, ...extra });

async function open(page: Page, options: GuildBotStubOptions) {
  const seen = await stubGuildBot(page, options);
  await gotoHydrated(page, SETTINGS_GUILD_PATH, 'guild-settings');
  await expect(section(page).locator('h2')).toHaveText(guildDiscordCopy.heading);
  // The Billing section sits above: its read settled, nothing above #discord moves after this.
  await expect(page.getByTestId('guild-billing-planline')).toBeVisible();
  return seen;
}

async function expectNoDestination(page: Page): Promise<void> {
  for (const testid of ['bot-destination', 'bot-test', 'bot-change', 'bot-remove', 'bot-linked-members'])
    await expect(id(page, testid)).toHaveCount(0);
}

test.describe('11 · deleted', () => {
  test('name known, without a webhook', async ({ page }) => {
    await open(page, { discord: unusable(DELETED, {}) });
    await expect(id(page, 'discord-notice')).toContainText(
      guildBotCopy.channelUnusable.title('deleted', 'raid-signups', 'Olympus Discord').text,
    );
    await expect(id(page, 'discord-notice')).toContainText(
      guildBotCopy.channelUnusable.detail('deleted', [], false),
    );
    await expect(id(page, 'bot-choose-channel')).toHaveText(guildBotCopy.channelUnusable.choose);
    await expect(id(page, 'bot-open')).toHaveCount(0);
    await expect(
      section(page).locator(
        'button.bg-gold, a.bg-gold, button[class*="linear-gradient"], a[class*="linear-gradient"]',
      ),
    ).toHaveCount(1);
    await expectNoDestination(page);
  });

  test('name unknown: the webhook sentence with a webhook, the stop sentence without', async ({ page }) => {
    await open(page, {
      discord: botSettings({ channel_unusable: DELETED, channel_name: null }, WITH_WEBHOOK),
    });
    await expect(id(page, 'discord-notice')).toContainText(
      guildBotCopy.channelUnusable.title('deleted', null, 'Olympus Discord').text,
    );
    await expect(id(page, 'discord-notice')).toContainText(
      guildBotCopy.channelUnusable.detail('deleted', [], true),
    );
    await expect(id(page, 'discord-notice')).toContainText(
      'Raids post through your webhook channel until then.',
    );
    await expectNoDestination(page);
    const bare = await page.context().newPage();
    await open(bare, { discord: unusable(DELETED, { channel_name: null }) });
    await expect(id(bare, 'discord-notice')).toContainText('Raids stop posting to Discord until then.');
  });
});

test.describe('12 · permissions', () => {
  const cases: [string, GuildChannelPermission[]][] = [
    ['one', ['Send Messages']],
    ['two', ['Send Messages', 'Embed Links']],
    ['three', ALL_PERMISSIONS],
  ];
  for (const [name, missing] of cases) {
    test(`${name} missing: the list in order and the channel link`, async ({ page }) => {
      await open(page, { discord: unusable(permissions(missing)) });
      const notice = id(page, 'discord-notice');
      await expect(notice).toContainText('The bot cannot post in #raid-signups.');
      await expect(notice.locator('span.font-semibold')).toHaveText('#raid-signups');
      await expect(notice).toContainText(
        `In Discord, give the Forever Sixty role ${joinPermissions(missing)} in that channel, or choose another channel.`,
      );
      const open_ = id(page, 'bot-open');
      await expect(open_).toHaveText(guildBotCopy.channelUnusable.open);
      await expect(open_).toHaveAttribute(
        'href',
        `https://discord.com/channels/${SERVER_ID}/${RAID_CHANNEL_ID}`,
      );
      await expect(open_).toHaveAttribute('target', '_blank');
      await expect(open_).toHaveAttribute('rel', 'noopener');
    });
  }

  test('the joined list reads as specified', async ({ page }) => {
    await open(page, { discord: unusable(permissions(ALL_PERMISSIONS)) });
    await expect(id(page, 'discord-notice')).toContainText(
      'View Channel, Send Messages and Embed Links in that channel',
    );
  });
});

test.describe('13 · precedence', () => {
  test('removed in Discord wins; null draws state 2', async ({ page }) => {
    await open(page, { discord: unusable(DELETED, { removed_in_discord: true }) });
    await expect(id(page, 'discord-notice')).toContainText(
      guildBotCopy.removedInDiscord.title('Olympus Discord'),
    );
    await expect(id(page, 'bot-choose-channel')).toHaveCount(0);
    const working = await page.context().newPage();
    await open(working, { discord: botSettings({}) });
    await expect(id(working, 'bot-destination')).toBeVisible();
    await expect(id(working, 'bot-choose-channel')).toHaveCount(0);
  });

  test('an open picker wins', async ({ page }) => {
    await open(page, { discord: unusable(DELETED) });
    await id(page, 'bot-choose-channel').click();
    await expect(fieldControl(section(page), 'bot-channel')).toBeVisible();
    await expect(id(page, 'discord-notice')).toHaveCount(0);
  });
});

test.describe('14 · fix by picking', () => {
  test('the picker opens focused and unselected; Cancel and Escape return to the notice', async ({
    page,
  }) => {
    await open(page, { discord: unusable(DELETED) });
    await id(page, 'bot-choose-channel').click();
    await expect(fieldControl(section(page), 'bot-channel')).toBeFocused();
    await expect(id(page, 'bot-channel')).toHaveValue('');
    await id(page, 'bot-channel-cancel').click();
    await expect(id(page, 'discord-notice')).toBeVisible();
    await expect(id(page, 'bot-choose-channel')).toBeFocused();
    await id(page, 'bot-choose-channel').click();
    await expect(fieldControl(section(page), 'bot-channel')).toBeFocused();
    await page.keyboard.press('Escape');
    await expect(id(page, 'discord-notice')).toBeVisible();
    await expect(id(page, 'bot-choose-channel')).toBeFocused();
  });

  test('a save shows state 2 with the saved line and focus on the test button', async ({ page }) => {
    const seen = await open(page, { discord: unusable(DELETED) });
    await id(page, 'bot-choose-channel').click();
    await choose(id(page, 'bot-channel'), RAID_CHANNEL_ID);
    await id(page, 'bot-channel-save').click();
    await expect(section(page).getByRole('status')).toHaveText(guildBotCopy.installed.saved);
    await expect(id(page, 'bot-test')).toBeFocused();
    await expect(id(page, 'discord-notice')).toHaveCount(0);
    expect(seen.base.current()?.bot?.channel_unusable).toBeNull();
  });

  test('permissions: the current channel is preselected only when it is in the list', async ({ page }) => {
    await open(page, { discord: unusable(permissions(['Send Messages'])) });
    await id(page, 'bot-choose-channel').click();
    await expect(id(page, 'bot-channel')).toHaveValue(RAID_CHANNEL_ID);
    const gone = await page.context().newPage();
    await open(gone, {
      discord: unusable(permissions(['Send Messages'])),
      channels: [CHANNELS[0]!, CHANNELS[2]!],
    });
    await id(gone, 'bot-choose-channel').click();
    await expect(id(gone, 'bot-channel')).toHaveValue('');
  });
});

test.describe('14b · Escape on the state 2 Change channel picker', () => {
  test('Escape cancels it and focus returns to Change channel', async ({ page }) => {
    await open(page, { discord: botSettings({}) });
    await id(page, 'bot-change').click();
    await expect(fieldControl(section(page), 'bot-channel')).toBeFocused();
    await page.keyboard.press('Escape');
    await expect(id(page, 'bot-destination')).toBeVisible();
    await expect(id(page, 'bot-change')).toBeFocused();
  });
});

test.describe('14c · Escape after a save', () => {
  test('a saved channel does not leave Escape stuck for the next Change channel picker', async ({ page }) => {
    await open(page, { discord: unusable(DELETED) });
    await id(page, 'bot-choose-channel').click();
    await choose(id(page, 'bot-channel'), RAID_CHANNEL_ID);
    await id(page, 'bot-channel-save').click();
    await expect(id(page, 'bot-test')).toBeFocused();
    await id(page, 'bot-change').click();
    await expect(fieldControl(section(page), 'bot-channel')).toBeFocused();
    await page.keyboard.press('Escape');
    await expect(id(page, 'bot-destination')).toBeVisible();
    await expect(id(page, 'bot-change')).toBeFocused();
  });
});

test.describe('15 · fix by saving the same channel', () => {
  test('still unusable: the 422 says so, focused, and the list is fresh; fixed: the 200 clears it', async ({
    page,
  }) => {
    const seen = await open(page, { discord: unusable(permissions(['Embed Links'])), unusableSaves: 1 });
    await id(page, 'bot-choose-channel').click();
    await expect(id(page, 'bot-channel')).toHaveValue(RAID_CHANNEL_ID);
    await id(page, 'bot-channel-save').click();
    await expect(section(page).getByRole('alert')).toHaveText(guildBotCopy.picker.race);
    await expect(page.locator('#guild-discord-status')).toBeFocused();
    expect(seen.channelReads).toBe(2);
    await choose(id(page, 'bot-channel'), RAID_CHANNEL_ID);
    await id(page, 'bot-channel-save').click();
    await expect(section(page).getByRole('status')).toHaveText(guildBotCopy.installed.saved);
    await expect(id(page, 'discord-notice')).toHaveCount(0);
    expect(seen.saves).toEqual([RAID_CHANNEL_ID, RAID_CHANNEL_ID]);
  });
});

test.describe('16 · officer-read-only and frozen', () => {
  for (const [name, gate] of [
    ['officer-read-only', { officerTools: false }],
    ['frozen', { frozen: true }],
  ] as const) {
    test(`${name}: Choose a channel is disabled, Open in Discord works`, async ({ page }) => {
      await open(page, { discord: unusable(permissions(['Send Messages'])), ...gate });
      await expect(id(page, 'bot-choose-channel')).toBeDisabled();
      await expect(id(page, 'bot-open')).toHaveAttribute('href', /^https:\/\/discord\.com\/channels\//);
    });
  }
});

test.describe('17 · axe, 44 px targets, caps and layout', () => {
  test.use({ reducedMotion: 'reduce' });
  const causes: [string, GuildChannelUnusable, GuildBotStubOptions['discord']][] = [
    ['deleted', DELETED, unusable(DELETED, { channel_name: null })],
    ['permissions', permissions(ALL_PERMISSIONS), unusable(permissions(ALL_PERMISSIONS))],
  ];
  for (const [name, , discord] of causes) {
    test(name, async ({ page }) => {
      for (const viewport of [PHONE, DESKTOP]) {
        await page.setViewportSize(viewport);
        await open(page, { discord });
        await expectNoAxeViolations(page, SECTION);
        const height = await section(page).evaluate((node) => node.getBoundingClientRect().height);
        expect(height, `${name} @${viewport.width}`).toBeLessThanOrEqual(CAP[viewport === PHONE ? 0 : 1]);
        if (viewport === PHONE) {
          const heights = await section(page)
            .locator('button, a[data-testid="bot-open"]')
            .evaluateAll((nodes) =>
              nodes.map((node) => {
                return node.getBoundingClientRect().height;
              }),
            );
          for (const h of heights) expect(h).toBeGreaterThanOrEqual(MIN_TARGET);
        }
      }
    });
  }

  test('no layout shift between the notice and state 2', async ({ page }) => {
    await open(page, { discord: unusable(DELETED) });
    const tops = (): Promise<{ section: number; next: number; height: number }> =>
      section(page).evaluate((node) => {
        const box = node.getBoundingClientRect();
        const next = node.nextElementSibling?.getBoundingClientRect().top ?? box.bottom;
        return { section: box.top + window.scrollY, next: next + window.scrollY, height: box.height };
      });
    const before = await tops();
    await id(page, 'bot-choose-channel').click();
    await choose(id(page, 'bot-channel'), RAID_CHANNEL_ID);
    await id(page, 'bot-channel-save').click();
    await expect(id(page, 'bot-test')).toBeFocused();
    const after = await tops();
    expect(after.section).toBe(before.section);
    expect(Math.abs(after.height - before.height)).toBeLessThanOrEqual(RESERVE_DELTA);
    expect(Math.abs(after.next - before.next)).toBeLessThanOrEqual(RESERVE_DELTA);
  });
});
