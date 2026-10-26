// web/tests/e2e/support/guild-discord.ts
// Stubs for the guild settings' Discord section (design/specs/2026-10-26-guild-raids.md §3.8),
// shared by tests/e2e/guild-settings-discord.spec.ts and the guild-raids acceptance captures. The
// guild is the settings-billing specs' The Last Watch (id 501): its public page, `/v1/me`, the
// settings read (whose `discord` key the stub holds and changes), `/home` for the officer-tools
// flag the standalone settings route reads, and the test and remove routes, which it records.
import type { Page, Route } from '@playwright/test';
import type { GuildSettingsDiscord } from '../../../src/lib/guild/api';
import {
  CLAIM_HOLDER_CAN,
  SETTINGS_GUILD,
  SETTINGS_GUILD_PATH,
  guildBilling,
  meWithGuildPlan,
} from '../../../src/test-support/billing-fixtures';
import { CHANNEL_ID, E2E_API_BASE, GUILD_ID as DISCORD_SERVER_ID } from './discord-notifications';

export { CHANNEL_ID, DISCORD_SERVER_ID, E2E_API_BASE, SETTINGS_GUILD_PATH };

export const DISCORD_SETTINGS_URL = `${SETTINGS_GUILD_PATH}#discord`;
export const GUILD_CONNECT_HREF = `${E2E_API_BASE}/v1/guilds/${SETTINGS_GUILD.id}/notifications/discord/connect`;

export function guildChannel(overrides: Partial<GuildSettingsDiscord> = {}): GuildSettingsDiscord {
  return {
    available: true,
    connected: true,
    channel_id: CHANNEL_ID,
    discord_guild_id: DISCORD_SERVER_ID,
    disabled_reason: null,
    last_error: null,
    failures: 0,
    post_default: true,
    ...overrides,
  };
}

export const NOT_CONNECTED = guildChannel({
  connected: false,
  channel_id: null,
  discord_guild_id: null,
  post_default: false,
});
export const SWITCHED_OFF = guildChannel({
  failures: 5,
  disabled_reason: 'failed five times',
  last_error: 'Unknown Webhook',
});

export interface Refusal {
  status: number;
  code: string;
  message: string;
}

export interface GuildDiscordStubOptions {
  /** The settings read's `discord` key; `null` leaves the key out (a member, connect off). */
  discord: GuildSettingsDiscord | null;
  /** Read once the first settings read has answered (the `connected` return's fresh read). */
  after?: GuildSettingsDiscord;
  /** `home.viewer.officer_tools`. */
  officerTools?: boolean;
  frozen?: boolean;
  /** The test route answers this instead of 204. */
  testRefusal?: Refusal;
}

export interface GuildDiscordStub {
  settingsReads: number;
  tests: number;
  deletes: number;
}

function envelope(data: unknown, status = 200): { status: number; contentType: string; body: string } {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'req-e2e' }),
  };
}

function refused({ status, code, message }: Refusal): { status: number; contentType: string; body: string } {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: false, data: null, error: { code, message }, request_id: 'req-e2e' }),
  };
}

const NO_CONTENT = { status: 204, body: '' };

function settingsBody(discord: GuildSettingsDiscord | null, frozen: boolean): Record<string, unknown> {
  return {
    default_visibility: 'guild',
    officer_max_rank_index: 1,
    claimed_by: { battletag: 'Fixture#1234' },
    claim_pending: null,
    claim: frozen
      ? { state: 'contested', since: '2026-09-01T00:00:00Z', frozen: true }
      : { state: 'claimed', since: '2026-09-01T00:00:00Z', frozen: false },
    invite: { rotated_at: null },
    ...(discord === null ? {} : { discord }),
  };
}

/** Routes everything the standalone settings page reads for The Last Watch, and the two writes. */
export async function stubGuildDiscord(
  page: Page,
  options: GuildDiscordStubOptions,
): Promise<GuildDiscordStub> {
  const seen: GuildDiscordStub = { settingsReads: 0, tests: 0, deletes: 0 };
  let discord = options.discord;
  const { region, ruleset, id } = SETTINGS_GUILD;
  await page.route(`**/v1/guilds/${region}/${ruleset}/the-last-watch`, (route) =>
    route.fulfill(envelope({ guild: SETTINGS_GUILD, progression: [], roster_best: [], reports: [] })),
  );
  await page.route('**/v1/me', (route) =>
    route.fulfill(envelope(meWithGuildPlan(guildBilling(), CLAIM_HOLDER_CAN).data)),
  );
  await page.route(`**/v1/guilds/${id}/home`, (route) =>
    route.fulfill(envelope({ viewer: { officer_tools: options.officerTools ?? true } })),
  );
  await page.route(`**/v1/guilds/${id}/settings`, (route) => {
    seen.settingsReads += 1;
    const body = settingsBody(discord, options.frozen ?? false);
    if (options.after !== undefined) discord = options.after;
    return route.fulfill(envelope(body));
  });
  await page.route(`**/v1/guilds/${id}/notifications/discord/test`, (route: Route) => {
    seen.tests += 1;
    if (options.testRefusal !== undefined) return route.fulfill(refused(options.testRefusal));
    if (discord !== null) discord = { ...discord, failures: 0, disabled_reason: null, last_error: null };
    return route.fulfill(NO_CONTENT);
  });
  await page.route(`**/v1/guilds/${id}/notifications/discord`, (route: Route) => {
    if (route.request().method() !== 'DELETE') return route.fallback();
    seen.deletes += 1;
    if (discord !== null) discord = NOT_CONNECTED;
    return route.fulfill(NO_CONTENT);
  });
  return seen;
}
