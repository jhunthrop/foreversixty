// web/src/lib/guild/discord-returns.svelte.ts
// The guild settings Discord section's two address-bar returns (raids spec §3.8 Return, bot spec
// §3.1): the webhook's `?discord=` and the bot's `?discord_bot=`, each read once and dropped from
// the address with `#discord` kept, the section scrolled to and the outcome shown. A return whose
// settings read is stale (`connected`, `installed`) holds its target height until the fresh read
// answers, or the settle timer runs out. GuildDiscordSection lends the section's status line,
// focus and settings setter (DiscordReturnHost); the API and the address bar come in as deps.
import { tick } from 'svelte';
import { takeReturnFlag } from '../account/discord-actions';
import { discordReturnFlag } from '../account/notifications';
import {
  clearCachedBotRefusal,
  refetchGuildSettings,
  type GuildSettingsData,
  type GuildSettingsDiscord,
} from './api';
import {
  BOT_FOCUS,
  awaitsBotInstallRead,
  botReturnOutcome,
  takeBotReturnFlag,
  type BotReturnOutcome,
} from './discord-bot-state';
import { GUILD_BOT_RESERVE, guildBotView } from './discord-bot-view';
import { guildReturnOutcome, type DiscordStatus, type GuildReturnOutcome } from './discord-settings';
import { GUILD_DISCORD_RESERVE } from './discord-state';

/** The longest a stale return waits for the settings read before drawing what it has. */
export const CONNECT_SETTLE_MS = 4000;

/** What the section lends the returns. */
export interface DiscordReturnHost {
  guildId(): number;
  /** The settings read in hand. */
  discord(): GuildSettingsDiscord;
  /** A fresh settings read, applied to the page. */
  onSettings(fresh: GuildSettingsData): void;
  show(status: DiscordStatus): void;
  focus(selector: string): void;
  /** The picker's select takes focus when it draws. */
  focusSelect(): void;
  scrollToSection(): void;
}

/** The address bar and the API the returns read; tests pass their own. */
export interface DiscordReturnDeps {
  location: () => Location;
  history: () => History;
  refetchSettings: (
    guildId: number,
    onData: (fresh: GuildSettingsData) => void,
    onError?: () => void,
  ) => void;
  clearRefusal: (guildId: number) => Promise<void>;
  setTimer: (run: () => void, ms: number) => ReturnType<typeof setTimeout>;
  clearTimer: (timer: ReturnType<typeof setTimeout> | undefined) => void;
}

export const browserReturnDeps: DiscordReturnDeps = {
  location: () => window.location,
  history: () => window.history,
  refetchSettings: (guildId, onData, onError) => refetchGuildSettings(guildId, onData, undefined, onError),
  clearRefusal: (guildId) => clearCachedBotRefusal(guildId),
  setTimer: (run, ms) => setTimeout(run, ms),
  clearTimer: (timer) => clearTimeout(timer),
};

/** The address bar's query on the client; empty on the server, where no return is read. */
function initialSearch(): string {
  return typeof window === 'undefined' ? '' : window.location.search;
}

export class GuildDiscordReturns {
  /** A `connected` return waits for the settings read: the connected height is held. */
  settling = $state(false);
  /** An `installed` return waits for the settings read: the picker height is held. */
  botSettling = $state(false);

  private botReturnDone = false;
  private settleTimer: ReturnType<typeof setTimeout> | undefined;
  private botSettleTimer: ReturnType<typeof setTimeout> | undefined;
  private readonly host: DiscordReturnHost;
  private readonly deps: DiscordReturnDeps;

  constructor(
    host: DiscordReturnHost,
    deps: DiscordReturnDeps = browserReturnDeps,
    search = initialSearch(),
  ) {
    this.host = host;
    this.deps = deps;
    this.settling = discordReturnFlag(search) === 'connected';
    this.botSettling = awaitsBotInstallRead(search);
  }

  /** The section's held height: a waiting return's target, else the bot state's live reserve. */
  held(botReserve: string): string {
    if (this.settling) return GUILD_DISCORD_RESERVE.connected;
    return this.botSettling ? GUILD_BOT_RESERVE.picker : botReserve;
  }

  /** Reads both returns once; call on mount. */
  take(): void {
    this.takeReturn();
    this.takeBotReturn();
  }

  dispose(): void {
    this.deps.clearTimer(this.settleTimer);
    this.deps.clearTimer(this.botSettleTimer);
  }

  private scroll(): void {
    void tick().then(() => this.host.scrollToSection());
  }

  private reveal(outcome: GuildReturnOutcome): void {
    this.deps.clearTimer(this.settleTimer);
    this.settling = false;
    this.host.show(outcome.status);
    this.host.focus(outcome.focus);
  }

  /** The callback's `?discord=` flag: taken once, the section scrolled to, the outcome shown. */
  private takeReturn(): void {
    const flag = takeReturnFlag(this.deps.location(), this.deps.history());
    if (flag === null) {
      this.settling = false;
      return;
    }
    this.scroll();
    const outcome = guildReturnOutcome(flag);
    if (!outcome.refresh) return this.reveal(outcome);
    this.settleTimer = this.deps.setTimer(() => this.reveal(outcome), CONNECT_SETTLE_MS);
    this.deps.refetchSettings(this.host.guildId(), (fresh) => {
      this.host.onSettings(fresh);
      this.reveal(outcome);
    });
  }

  /** The bot's install return, shown once: the line stays, and focus goes where the outcome says
   *  if that control is drawn for the read in hand (the select in the picker, else Send test). */
  private revealBot(outcome: BotReturnOutcome, read: GuildSettingsDiscord): void {
    if (this.botReturnDone) return;
    this.botReturnDone = true;
    if (read.bot_refusal != null) void this.deps.clearRefusal(this.host.guildId()).catch(() => {});
    this.deps.clearTimer(this.botSettleTimer);
    this.botSettling = false;
    this.host.show({ ok: outcome.role === 'status', message: outcome.message, stays: true });
    if (outcome.focus !== BOT_FOCUS.select) return this.host.focus(outcome.focus);
    const { state: shown } = guildBotView(read, false);
    if (shown === 'picker') this.host.focusSelect();
    else if (shown === 'installed') this.host.focus(BOT_FOCUS.testButton);
  }

  /** The settle timer: reveal on the read in hand only if it holds a bot. A stale read keeps the
   *  picker's held height (no state 5 flash) until the fresh read answers or the refetch fails. */
  private settleBot(outcome: BotReturnOutcome): void {
    const read = this.host.discord();
    if (guildBotView(read, false).bot !== null) this.revealBot(outcome, read);
  }

  /** The callback's `?discord_bot=` flag, read once like the webhook's (bot spec §3.1). */
  private takeBotReturn(): void {
    const flag = takeBotReturnFlag(this.deps.location(), this.deps.history());
    const outcome = flag === null ? null : botReturnOutcome(flag, this.host.discord().bot_refusal ?? null);
    if (outcome === null) {
      this.botSettling = false;
      return;
    }
    this.scroll();
    if (!outcome.holdsPicker) return this.revealBot(outcome, this.host.discord());
    this.botSettleTimer = this.deps.setTimer(() => this.settleBot(outcome), CONNECT_SETTLE_MS);
    this.deps.refetchSettings(
      this.host.guildId(),
      (fresh) => {
        this.host.onSettings(fresh);
        this.revealBot(outcome, fresh.discord ?? this.host.discord());
      },
      () => {
        // The fresh read failed: stop holding and draw the read in hand; `Bot added` only if it
        // holds a bot.
        this.botSettling = false;
        this.settleBot(outcome);
      },
    );
  }
}
