// web/src/components/guild/GuildDiscordSection.test.ts
// One static render per raids spec §3.8 state (vitest hands components Svelte's server entry, so
// the clicks, the returns and focus are Playwright's: tests/e2e/guild-settings-discord.spec.ts).
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { notificationsCopy } from '../../lib/account/notifications-copy';
import type { GuildSettingsDiscord } from '../../lib/guild/api';
import { guildDiscordConnectUrl } from '../../lib/guild/discord-api';
import { guildDiscordCopy, type DiscordStatus } from '../../lib/guild/discord-settings';
import GuildDiscordChannel from './GuildDiscordChannel.svelte';
import GuildDiscordSection from './GuildDiscordSection.svelte';

const GUILD = 501;
const CONNECTED: GuildSettingsDiscord = {
  available: true,
  connected: true,
  channel_id: '1198765432101234568',
  discord_guild_id: '1087654321098765432',
  disabled_reason: null,
  last_error: null,
  failures: 0,
  post_default: false,
};
const NOT_CONNECTED: GuildSettingsDiscord = {
  ...CONNECTED,
  connected: false,
  channel_id: null,
  discord_guild_id: null,
};
const SWITCHED_OFF: GuildSettingsDiscord = {
  ...CONNECTED,
  failures: 5,
  disabled_reason: 'failed five times',
  last_error: 'Unknown Webhook',
};

interface Options {
  frozen?: boolean;
  officerTools?: boolean;
}

function html(discord: GuildSettingsDiscord, { frozen = false, officerTools = true }: Options = {}): string {
  return render(GuildDiscordSection, {
    props: { guildId: GUILD, discord, frozen, officerTools, onDiscord: () => {}, onSettings: () => {} },
  }).body.replace(/<!--[\s\S]*?-->/g, '');
}

function channel(confirming: boolean, status: DiscordStatus | null = null): string {
  return render(GuildDiscordChannel, {
    props: {
      discord: CONNECTED,
      switchedOff: false,
      connectHref: guildDiscordConnectUrl(GUILD),
      confirming,
      busy: null,
      status,
      gateId: undefined,
      frozen: false,
      onTest: () => {},
      onAskRemove: () => {},
      onRemove: () => {},
      onCancel: () => {},
      onStuck: () => {},
    },
  }).body.replace(/<!--[\s\S]*?-->/g, '');
}

const text = (body: string): string => body.replace(/<[^>]+>/g, ' ').replace(/\s+/g, ' ');
const tag = (body: string, testid: string): string =>
  new RegExp(`<[a-z]+[^>]*data-testid="${testid}"[^>]*>`).exec(body)?.[0] ?? '';
const filled = (body: string): number =>
  (body.match(/<(?:button|a)[^>]*(?:bg-gold|bg-\[linear-gradient)[^>]*>/g) ?? []).length;

describe('GuildDiscordSection', () => {
  it('is the #discord card with the Discord heading in every state', () => {
    for (const discord of [NOT_CONNECTED, CONNECTED, SWITCHED_OFF]) {
      const body = html(discord);
      expect(body).toMatch(/<section id="discord" class="[^"]*scroll-mt-20/);
      expect(body).toMatch(new RegExp(`<h2[^>]*>\\s*${guildDiscordCopy.heading}\\s*</h2>`));
      expect(body).not.toContain('rounded-panel');
    }
  });

  it('not connected: the gold connect link to the guild route, its explanation, the status line', () => {
    const body = html(NOT_CONNECTED);
    const link = tag(body, 'discord-connect');
    expect(link).toContain(`href="${guildDiscordConnectUrl(GUILD)}"`);
    expect(link).toMatch(/href="https?:\/\//);
    expect(link).not.toContain('target=');
    expect(text(body)).toContain(notificationsCopy.connect);
    expect(text(body)).toContain(guildDiscordCopy.explanation);
    expect(body).toContain('data-testid="discord-status"');
    expect(filled(body)).toBe(1);
    expect(body).not.toContain('Paste');
  });

  it('connected: the destination, the channel line with Open in Discord, test, remove, reconnect; no fill', () => {
    const body = html(CONNECTED);
    expect(text(body)).toContain('Posts raids to a Discord channel');
    expect(text(body)).toContain('Channel 1198…4568');
    expect(tag(body, 'discord-open')).toContain(
      'href="https://discord.com/channels/1087654321098765432/1198765432101234568"',
    );
    expect(text(body)).toContain(notificationsCopy.test);
    expect(text(body)).toContain(guildDiscordCopy.remove);
    expect(tag(body, 'discord-replace')).toContain(`href="${guildDiscordConnectUrl(GUILD)}"`);
    expect(text(body)).toContain(notificationsCopy.reconnectNote);
    expect(filled(body)).toBe(0);
    expect(body).not.toContain('data-testid="discord-warning"');
  });

  it('connected without the server id: no Open in Discord link', () => {
    expect(html({ ...CONNECTED, discord_guild_id: null })).not.toContain('discord-open');
  });

  it('connected with 3 failures: the account warning, verbatim', () => {
    const body = html({ ...CONNECTED, failures: 3, last_error: 'Missing Access' });
    const warning = notificationsCopy.warningLines(3, 'Missing Access');
    expect(text(body)).toContain(warning.failures);
    expect(text(body)).toContain(warning.lastError ?? '');
  });

  it('switched off: the notice with the reconnect inside it, Test and switch on as the one fill', () => {
    const body = html(SWITCHED_OFF);
    expect(body).toContain('data-testid="discord-notice"');
    expect(text(body)).toContain(notificationsCopy.switchedOff);
    expect(text(body)).toContain('Last error from Discord: Unknown Webhook.');
    expect(text(body)).toContain(notificationsCopy.switchedOffFix);
    expect(filled(body)).toBe(1);
    expect(tag(body, 'discord-test')).toContain('bg-[linear-gradient');
    expect(text(body)).toContain(notificationsCopy.testAndSwitchOn);
    expect(body.indexOf('discord-replace')).toBeLessThan(body.indexOf('discord-test'));
  });

  it('officer-read-only, not connected: the gate sentence, a connect link with no href', () => {
    const body = html(NOT_CONNECTED, { officerTools: false });
    expect(body).toContain('id="officer-tools-gate-discord"');
    expect(text(body)).toContain('Planning and running raids are part of the guild plan.');
    expect(body).toMatch(/<a[^>]*href="#billing"[^>]*>Subscribe the guild<\/a>/);
    const link = tag(body, 'discord-connect');
    expect(link).not.toContain('href=');
    expect(link).toContain('aria-disabled="true"');
    expect(link).toContain('tabindex="-1"');
    expect(link).toContain('opacity-50');
    expect(link).toContain('aria-describedby="officer-tools-gate-discord"');
  });

  it('officer-read-only, connected: tests disabled and described, Remove live', () => {
    for (const discord of [CONNECTED, SWITCHED_OFF]) {
      const body = html(discord, { officerTools: false });
      const test = tag(body, 'discord-test');
      expect(test).toContain('disabled=""');
      expect(test).toContain('aria-describedby="officer-tools-gate-discord"');
      expect(tag(body, 'discord-remove')).not.toContain('disabled=""');
      expect(tag(body, 'discord-replace')).not.toContain('href=');
    }
  });

  it('frozen: every control off, Remove included, and no gate sentence', () => {
    const body = html(CONNECTED, { frozen: true, officerTools: false });
    expect(body).not.toContain('officer-tools-gate-discord"');
    expect(tag(body, 'discord-test')).toContain('disabled');
    expect(tag(body, 'discord-remove')).toContain('disabled');
    expect(tag(body, 'discord-replace')).toContain('aria-disabled="true"');
    expect(tag(html(NOT_CONNECTED, { frozen: true }), 'discord-connect')).not.toContain('href=');
  });
});

describe('GuildDiscordChannel', () => {
  it('the remove question replaces the actions, with Remove and Cancel, and no reconnect link', () => {
    const body = channel(true);
    expect(text(body)).toContain(guildDiscordCopy.removeQuestion);
    expect(body).toContain('data-testid="discord-remove-confirm"');
    expect(body).toContain('data-testid="discord-remove-cancel"');
    expect(body).not.toContain('data-testid="discord-test"');
    expect(body).not.toContain('data-testid="discord-replace"');
  });

  it('a status line is a focusable target: a failure is an alert, a success a status', () => {
    expect(channel(false, { ok: false, message: guildDiscordCopy.planEnded })).toMatch(
      /id="guild-discord-status" tabindex="-1"[\s\S]*role="alert"[^>]*>\s*Not saved/,
    );
    expect(channel(false, { ok: true, message: notificationsCopy.testSent })).toContain('role="status"');
  });
});
