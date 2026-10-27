// @vitest-environment jsdom
// web/src/lib/guild/discord-returns.test.ts
// The guild Discord section's address-bar returns: read once, the flag dropped with #discord kept,
// the held height while a stale read is fetched, and where the outcome's focus goes.
import { afterEach, describe, expect, it, vi } from 'vitest';
import type { GuildBot, GuildSettingsData, GuildSettingsDiscord } from './api';
import { guildBotCopy } from './discord-bot-copy';
import { BOT_FOCUS } from './discord-bot-state';
import { GUILD_BOT_RESERVE } from './discord-bot-view';
import {
  CONNECT_SETTLE_MS,
  GuildDiscordReturns,
  type DiscordReturnDeps,
  type DiscordReturnHost,
} from './discord-returns.svelte';
import { DISCORD_FOCUS, guildDiscordCopy } from './discord-settings';
import { GUILD_DISCORD_RESERVE } from './discord-state';

const GUILD_ID = 7;
const LIVE_RESERVE = 'min-h-[1px]';

const NOTHING: GuildSettingsDiscord = {
  available: true,
  connected: false,
  channel_id: null,
  discord_guild_id: null,
  disabled_reason: null,
  last_error: null,
  failures: 0,
  post_default: false,
  bot_available: true,
  bot: null,
  bot_refusal: null,
};

const BOT: GuildBot = {
  server_id: 'srv1',
  server_name: 'Olympus',
  channel_id: null,
  channel_name: null,
  linked_members: 0,
  members: 4,
  failures: 0,
  last_error: null,
  removed_in_discord: false,
  channel_unusable: null,
};

interface World {
  returns: GuildDiscordReturns;
  host: { [K in keyof DiscordReturnHost]: ReturnType<typeof vi.fn> };
  refetch: ReturnType<typeof vi.fn>;
  clearRefusal: ReturnType<typeof vi.fn>;
  timers: Array<() => void>;
}

function world(address: string, read: GuildSettingsDiscord = NOTHING): World {
  window.history.replaceState({}, '', address);
  const timers: Array<() => void> = [];
  const host = {
    guildId: vi.fn(() => GUILD_ID),
    discord: vi.fn(() => read),
    onSettings: vi.fn(),
    show: vi.fn(),
    focus: vi.fn(),
    focusSelect: vi.fn(),
    scrollToSection: vi.fn(),
  };
  const refetch = vi.fn();
  const clearRefusal = vi.fn(() => Promise.resolve());
  const deps: DiscordReturnDeps = {
    location: () => window.location,
    history: () => window.history,
    refetchSettings: refetch,
    clearRefusal,
    setTimer: (run, ms) => {
      expect(ms).toBe(CONNECT_SETTLE_MS);
      timers.push(run);
      return timers.length as unknown as ReturnType<typeof setTimeout>;
    },
    clearTimer: () => {},
  };
  return {
    returns: new GuildDiscordReturns(host, deps, window.location.search),
    host,
    refetch,
    clearRefusal,
    timers,
  };
}

function settings(discord: GuildSettingsDiscord): GuildSettingsData {
  return { discord } as unknown as GuildSettingsData;
}

function address(): string {
  return window.location.pathname + window.location.search + window.location.hash;
}

afterEach(() => {
  window.history.replaceState({}, '', '/');
});

describe('GuildDiscordReturns, the held height', () => {
  it('holds the connected height on a connected return, the picker on an install, else the live reserve', () => {
    expect(world('/g/settings?discord=connected#discord').returns.held(LIVE_RESERVE)).toBe(
      GUILD_DISCORD_RESERVE.connected,
    );
    expect(world('/g/settings?discord_bot=installed#discord').returns.held(LIVE_RESERVE)).toBe(
      GUILD_BOT_RESERVE.picker,
    );
    expect(world('/g/settings#discord').returns.held(LIVE_RESERVE)).toBe(LIVE_RESERVE);
  });
});

describe('GuildDiscordReturns, the webhook return', () => {
  it('denied: shown at once, focus on connect, the flag gone with the hash and other params kept', () => {
    const w = world('/g/settings?tab=x&discord=denied#discord');
    w.returns.take();
    expect(w.host.show).toHaveBeenCalledWith({
      ok: true,
      message: guildDiscordCopy.returnDenied,
      stays: true,
    });
    expect(w.host.focus).toHaveBeenCalledWith(DISCORD_FOCUS.connect);
    expect(w.refetch).not.toHaveBeenCalled();
    expect(address()).toBe('/g/settings?tab=x#discord');
  });

  it('connected: holds until the fresh read, then applies it and reveals once', () => {
    const w = world('/g/settings?discord=connected#discord');
    w.returns.take();
    expect(w.returns.settling).toBe(true);
    expect(w.host.show).not.toHaveBeenCalled();
    const [guildId, onData] = w.refetch.mock.calls[0] as [number, (fresh: GuildSettingsData) => void];
    expect(guildId).toBe(GUILD_ID);
    const fresh = settings({ ...NOTHING, connected: true });
    onData(fresh);
    expect(w.host.onSettings).toHaveBeenCalledWith(fresh);
    expect(w.returns.settling).toBe(false);
    expect(w.host.focus).toHaveBeenCalledWith(DISCORD_FOCUS.test);
  });

  it('connected: the settle timer reveals on the read in hand when the fetch is slow', () => {
    const w = world('/g/settings?discord=connected#discord');
    w.returns.take();
    w.timers[0]?.();
    expect(w.returns.settling).toBe(false);
    expect(w.host.show).toHaveBeenCalledWith({
      ok: true,
      message: guildDiscordCopy.returnConnected,
      stays: true,
    });
  });

  it('no flag: nothing shown and nothing held', () => {
    const w = world('/g/settings#discord');
    w.returns.take();
    expect(w.host.show).not.toHaveBeenCalled();
    expect(w.host.scrollToSection).not.toHaveBeenCalled();
    expect(w.returns.settling || w.returns.botSettling).toBe(false);
  });
});

describe('GuildDiscordReturns, the bot install return', () => {
  it('installed: the fresh read draws the picker, so the select takes focus; the line stays', () => {
    const w = world('/g/settings?discord_bot=installed#discord');
    w.returns.take();
    expect(w.returns.botSettling).toBe(true);
    const onData = w.refetch.mock.calls[0]?.[1] as (fresh: GuildSettingsData) => void;
    onData(settings({ ...NOTHING, bot: BOT }));
    expect(w.returns.botSettling).toBe(false);
    expect(w.host.show).toHaveBeenCalledWith({
      ok: true,
      message: guildBotCopy.returns.installed,
      stays: true,
    });
    expect(w.host.focusSelect).toHaveBeenCalledTimes(1);
    expect(address()).toBe('/g/settings#discord');
  });

  it('installed: a read that already has a channel sends focus to Send test', () => {
    const w = world('/g/settings?discord_bot=installed#discord');
    w.returns.take();
    const onData = w.refetch.mock.calls[0]?.[1] as (fresh: GuildSettingsData) => void;
    onData(settings({ ...NOTHING, bot: { ...BOT, channel_id: 'c1', channel_name: 'raids' } }));
    expect(w.host.focus).toHaveBeenCalledWith(BOT_FOCUS.testButton);
    expect(w.host.focusSelect).not.toHaveBeenCalled();
  });

  it('installed: a failed fetch stops holding but shows nothing while the read in hand has no bot', () => {
    const w = world('/g/settings?discord_bot=installed#discord');
    w.returns.take();
    const onError = w.refetch.mock.calls[0]?.[2] as () => void;
    onError();
    expect(w.returns.botSettling).toBe(false);
    expect(w.host.show).not.toHaveBeenCalled();
    w.timers[0]?.();
    expect(w.host.show).not.toHaveBeenCalled();
  });

  it('taken: shown once from the read, and the one-time refusal cleared from the cache', () => {
    const read = { ...NOTHING, bot_refusal: { guild_name: 'Mograine Raiders' } };
    const w = world('/g/settings?discord_bot=taken#discord', read);
    w.returns.take();
    expect(w.host.show).toHaveBeenCalledTimes(1);
    expect(w.clearRefusal).toHaveBeenCalledWith(GUILD_ID);
    expect(w.host.focus).toHaveBeenCalledWith(BOT_FOCUS.add);
    expect(w.refetch).not.toHaveBeenCalled();
  });
});
