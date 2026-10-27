<!-- web/src/components/guild/GuildDiscordSection.svelte -->
<!-- Guild settings · Discord (raids spec §3.8): the guild's one channel for raid posts, a flat
     section under Billing. Not connected: the gold connect link to the guild's route on the API
     origin. Connected and switched off: GuildDiscordChannel. Officer-read-only draws the gate
     sentence, a connect link that goes nowhere and disabled tests, and keeps Remove live; a frozen
     claim turns every control off. A return from Discord (`?discord=`) is read once, dropped from
     the address with `#discord` kept, and the section scrolled to; `connected` holds the connected
     height until the settings read answers. Words and outcomes: lib/guild/discord-settings.ts.
     With the bot on for the deployment (discord bot spec §3.1) the section adds the bot's states:
     add (nothing connected, and the offer above a webhook), the picker, installed and removed in
     Discord (GuildDiscordBot); its install return (`?discord_bot=`) is read once like the
     webhook's. With the bot off every branch below draws exactly as before. -->
<script lang="ts">
  import { onMount, tick } from 'svelte';
  import { createFocusQueue } from '../../lib/account/discord-return.svelte';
  import { takeReturnFlag } from '../../lib/account/discord-actions';
  import { discordReturnFlag } from '../../lib/account/notifications';
  import { notificationsCopy } from '../../lib/account/notifications-copy';
  import {
    fetchGuildHome,
    refetchGuildSettings,
    type GuildSettingsData,
    type GuildSettingsDiscord,
  } from '../../lib/guild/api';
  import { guildDiscordConnectUrl, removeGuildDiscord, testGuildDiscord } from '../../lib/guild/discord-api';
  import { guildBotInstallUrl, removeGuildBot, testGuildBot } from '../../lib/guild/discord-bot-api';
  import { guildBotCopy } from '../../lib/guild/discord-bot-copy';
  import {
    BOT_FOCUS,
    botReturnOutcome,
    clearedBotFailures,
    takeBotReturnFlag,
    type BotReturnOutcome,
  } from '../../lib/guild/discord-bot-state';
  import { BOT_RETURN_PARAM, GUILD_BOT_RESERVE, guildBotView } from '../../lib/guild/discord-bot-view';
  import {
    DISCORD_FOCUS,
    clearedFailures,
    connectStuck,
    guildDiscordCopy,
    guildReturnOutcome,
    removeRefused,
    removed,
    removedChannel,
    testRefused,
    testSent,
    type DiscordStatus,
    type GuildReturnOutcome,
  } from '../../lib/guild/discord-settings';
  import { GUILD_DISCORD_RESERVE, guildDiscordState } from '../../lib/guild/discord-state';
  import { PRIMARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import DiscordConnectLink from '../account/discord/DiscordConnectLink.svelte';
  import DiscordStatusLine from '../account/discord/DiscordStatusLine.svelte';
  import GuildDiscordAdd from './GuildDiscordAdd.svelte';
  import GuildDiscordBot from './GuildDiscordBot.svelte';
  import GuildDiscordChannel from './GuildDiscordChannel.svelte';
  import GuildPlanGate from './GuildPlanGate.svelte';

  /** How long `Removed.` stays (account spec 4.3); a `stays` line holds until the next press. */
  const STATUS_MS = 4000;
  /** The longest a `connected` return waits for the settings read before drawing what it has. */
  const CONNECT_SETTLE_MS = 4000;
  const GATE_ID = 'officer-tools-gate-discord';

  let {
    guildId,
    discord,
    frozen,
    officerTools,
    onDiscord,
    onSettings,
    onPlanRefused = () => {},
  }: {
    guildId: number;
    discord: GuildSettingsDiscord;
    frozen: boolean;
    /** From the guild page (`home.viewer.officer_tools`); left out on the standalone settings
     *  route, where the section reads the guild's `/home` itself. */
    officerTools?: boolean;
    /** The channel after a test or a remove, applied to the page's settings in place. */
    onDiscord: (next: GuildSettingsDiscord) => void;
    /** A fresh settings read (the `connected` return's). */
    onSettings: (fresh: GuildSettingsData) => void;
    /** A `guild_plan_required` refusal: the guild page turns its other tabs read-only too. */
    onPlanRefused?: () => void;
  } = $props();

  let homeTools = $state<boolean | null>(null);
  let planEnded = $state(false);
  let confirming = $state(false);
  let busy = $state<'test' | 'remove' | null>(null);
  let status = $state<DiscordStatus | null>(null);
  let settling = $state(
    typeof window !== 'undefined' && discordReturnFlag(window.location.search) === 'connected',
  );
  let changing = $state(false);
  let focusSelect = $state(false);
  let botReturnDone = false;
  let botSettling = $state(
    typeof window !== 'undefined' &&
      new URLSearchParams(window.location.search).get(BOT_RETURN_PARAM) === 'installed',
  );
  let section = $state<HTMLElement | null>(null);
  let statusTimer: ReturnType<typeof setTimeout> | undefined;
  let settleTimer: ReturnType<typeof setTimeout> | undefined;

  const state = $derived(guildDiscordState(discord));
  const botOn = $derived(discord.bot_available === true);
  const view = $derived(guildBotView(discord, changing));
  const installHref = $derived(guildBotInstallUrl(guildId));
  const tools = $derived(officerTools ?? homeTools ?? true);
  const readOnly = $derived(!frozen && (planEnded || !tools));
  const gateId = $derived(readOnly ? GATE_ID : undefined);
  const connectHref = $derived(guildDiscordConnectUrl(guildId));
  const gateLead = $derived(
    guildDiscordCopy.gate.slice(0, guildDiscordCopy.gate.indexOf(guildDiscordCopy.gateLink)),
  );
  const focus = createFocusQueue(() => [
    state,
    view.state,
    confirming,
    busy,
    settling,
    botSettling,
    readOnly,
  ]);
  const held = $derived(
    settling ? GUILD_DISCORD_RESERVE.connected : botSettling ? GUILD_BOT_RESERVE.picker : '',
  );

  $effect(() => () => {
    clearTimeout(statusTimer);
    clearTimeout(settleTimer);
  });

  function show(next: DiscordStatus): void {
    clearTimeout(statusTimer);
    status = next;
    if (next.ok && next.stays !== true) statusTimer = setTimeout(() => (status = null), STATUS_MS);
  }

  /** A status from a child: a plan refusal turns the section read-only; `focusLine` moves focus. */
  function report(next: DiscordStatus | null, focusLine = false): void {
    if (next === null) {
      clearTimeout(statusTimer);
      status = null;
      return;
    }
    show(next);
    if (next.planEnded === true) {
      planEnded = true;
      onPlanRefused();
    }
    if (focusLine) focus.request(DISCORD_FOCUS.status);
  }

  function reveal(outcome: GuildReturnOutcome): void {
    clearTimeout(settleTimer);
    settling = false;
    show(outcome.status);
    focus.request(outcome.focus);
  }

  /** The bot's install return, shown once: the line stays, and focus goes where the outcome says
   *  if that control is drawn for the read in hand (the select in the picker, else Send test). */
  function revealBot(outcome: BotReturnOutcome, read: GuildSettingsDiscord): void {
    if (botReturnDone) return;
    botReturnDone = true;
    clearTimeout(settleTimer);
    botSettling = false;
    show({ ok: outcome.role === 'status', message: outcome.message, stays: true });
    if (outcome.focus !== BOT_FOCUS.select) return focus.request(outcome.focus);
    const { state: shown } = guildBotView(read, false);
    if (shown === 'picker') focusSelect = true;
    else if (shown === 'installed') focus.request(BOT_FOCUS.testButton);
  }

  /** The settle timer: draw the read in hand, and say `Bot added` only if it holds a bot (a
   *  stale read is not one; the fresh read, when it lands, shows the line). */
  function settleBot(outcome: BotReturnOutcome): void {
    botSettling = false;
    if (guildBotView(discord, false).bot !== null) revealBot(outcome, discord);
  }

  /** The callback's `?discord=` flag: taken once, the section scrolled to, the outcome shown. */
  function takeReturn(): void {
    const flag = takeReturnFlag(window.location, window.history);
    if (flag === null) {
      settling = false;
      return;
    }
    void tick().then(() => section?.scrollIntoView({ block: 'start' }));
    const outcome = guildReturnOutcome(flag);
    if (!outcome.refresh) return reveal(outcome);
    settleTimer = setTimeout(() => reveal(outcome), CONNECT_SETTLE_MS);
    refetchGuildSettings(guildId, (fresh) => {
      onSettings(fresh);
      reveal(outcome);
    });
  }

  /** The callback's `?discord_bot=` flag, read once like the webhook's (bot spec §3.1). */
  function takeBotReturn(): void {
    const flag = takeBotReturnFlag(window.location, window.history);
    const outcome = flag === null ? null : botReturnOutcome(flag, discord.bot_refusal ?? null);
    if (outcome === null) {
      botSettling = false;
      return;
    }
    void tick().then(() => section?.scrollIntoView({ block: 'start' }));
    if (!outcome.holdsPicker) return revealBot(outcome, discord);
    settleTimer = setTimeout(() => settleBot(outcome), CONNECT_SETTLE_MS);
    refetchGuildSettings(guildId, (fresh) => {
      onSettings(fresh);
      revealBot(outcome, fresh.discord ?? discord);
    });
  }

  onMount(() => {
    takeReturn();
    takeBotReturn();
    if (officerTools !== undefined) return;
    fetchGuildHome(guildId)
      .then((home) => (homeTools = home.viewer?.officer_tools ?? true))
      .catch(() => (homeTools = null));
  });

  async function run(kind: 'test' | 'remove', action: () => Promise<void>): Promise<unknown> {
    busy = kind;
    status = null;
    focus.request(null);
    try {
      await action();
      return null;
    } catch (error) {
      return error ?? new Error('failed');
    } finally {
      busy = null;
    }
  }

  /** The bot is what a test and a remove reach while it is installed (picker or installed). */
  const botActive = $derived(view.bot !== null && view.state !== 'removed');

  async function test(): Promise<void> {
    const before = discord;
    const viaBot = botActive;
    const error = await run('test', () => (viaBot ? testGuildBot(guildId) : testGuildDiscord(guildId)));
    if (error === null) {
      if (viaBot) onDiscord(clearedBotFailures(before));
      else if (before.failures > 0 || before.disabled_reason !== null) onDiscord(clearedFailures(before));
      return show(testSent(!viaBot && before.disabled_reason !== null));
    }
    report(testRefused(error), true);
  }

  async function remove(): Promise<void> {
    const viaBot = botActive;
    const error = await run('remove', () => (viaBot ? removeGuildBot(guildId) : removeGuildDiscord(guildId)));
    if (error !== null)
      return show(viaBot ? { ok: false, message: guildBotCopy.remove.failed } : removeRefused());
    confirming = false;
    changing = false;
    onDiscord(viaBot ? { ...discord, bot: null } : removedChannel(discord));
    show(removed());
    focus.request(botOn ? BOT_FOCUS.add : DISCORD_FOCUS.connect);
  }

  function askRemove(): void {
    status = null;
    confirming = true;
    focus.request(botActive ? BOT_FOCUS.removeCancel : DISCORD_FOCUS.removeCancel);
  }

  function cancel(): void {
    confirming = false;
  }

  function change(): void {
    status = null;
    focusSelect = true;
    changing = true;
  }

  function cancelChange(): void {
    changing = false;
    focusSelect = false;
    focus.request(view.bot?.channel_unusable != null ? BOT_FOCUS.choose : BOT_FOCUS.change);
  }

  function saved(next: GuildSettingsDiscord): void {
    onDiscord(next);
    changing = false;
    focusSelect = false;
    show({ ok: true, message: guildBotCopy.installed.saved, stays: true });
    focus.request(BOT_FOCUS.testButton);
  }

  function onKeydown(event: KeyboardEvent): void {
    if (event.key !== 'Escape' || busy !== null) return;
    if (confirming) cancel();
    else if (changing && view.state === 'picker') cancelChange();
  }
</script>

<section
  id="discord"
  bind:this={section}
  class="flex scroll-mt-20 flex-col gap-3 {held}"
  data-testid="guild-settings-discord"
>
  <h2 class="section-title text-[18px]">{guildDiscordCopy.heading}</h2>
  {#if readOnly}<GuildPlanGate tab="discord" settingsHref="#billing" sentence={gateLead} />{/if}
  {#snippet addParts(variant: 'nothing' | 'offer')}
    <GuildDiscordAdd
      {variant}
      {installHref}
      webhookHref={connectHref}
      gated={frozen || readOnly}
      {gateId}
      busy={busy !== null}
      {status}
      onStuck={() => show(connectStuck())}
    />
  {/snippet}
  {#snippet offerPart()}{@render addParts('offer')}{/snippet}
  {#if settling || botSettling}
    <!-- the connected or picker height, held until the settings read answers (spec 3.8, Return) -->
  {:else if botOn && view.state === 'nothing'}
    {@render addParts('nothing')}
  {:else if view.bot !== null}
    <div class="flex max-w-[720px] flex-col gap-3" onkeydown={onKeydown} role="presentation">
      <GuildDiscordBot
        {guildId}
        {discord}
        bot={view.bot}
        botState={view.state}
        {installHref}
        {changing}
        {focusSelect}
        {readOnly}
        {frozen}
        {gateId}
        {confirming}
        {busy}
        {status}
        onStatus={report}
        onSaved={saved}
        onTest={() => void test()}
        onChange={change}
        onCancelChange={cancelChange}
        onAskRemove={askRemove}
        onRemove={() => void remove()}
        onCancelRemove={cancel}
        onStuck={() => show(connectStuck())}
      />
    </div>
  {:else if state === 'not-connected'}
    <!-- One line at 1440: the 140 cap counts the explanation as one line, so it is not held to 720. -->
    <div class="flex flex-col items-start">
      <DiscordConnectLink
        label={notificationsCopy.connect}
        class="{PRIMARY_BUTTON_FIXED} px-5 whitespace-nowrap"
        testid="discord-connect"
        href={connectHref}
        busy={busy !== null}
        gated={frozen || readOnly}
        describedBy={gateId}
        onStuck={() => show(connectStuck())}
      />
      <p class="text-muted mt-2 text-[13px]" data-testid="discord-explanation">
        {guildDiscordCopy.explanation}
      </p>
      <div id="guild-discord-status" tabindex="-1" class="w-full outline-none">
        <DiscordStatusLine {status} />
      </div>
    </div>
  {:else}
    <div class="flex max-w-[720px] flex-col gap-3" onkeydown={onKeydown} role="presentation">
      <GuildDiscordChannel
        {discord}
        switchedOff={state === 'switched-off'}
        {connectHref}
        {confirming}
        {busy}
        {status}
        {gateId}
        {frozen}
        onTest={() => void test()}
        onAskRemove={askRemove}
        onRemove={() => void remove()}
        onCancel={cancel}
        onStuck={() => show(connectStuck())}
        offer={botOn ? offerPart : undefined}
      />
    </div>
  {/if}
</section>
