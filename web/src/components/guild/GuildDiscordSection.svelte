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
     webhook's. With the bot off every branch below draws exactly as before. The two returns and
     their held heights live in lib/guild/discord-returns.svelte.ts. -->
<script lang="ts">
  import { onMount } from 'svelte';
  import { createFocusQueue } from '../../lib/account/discord-return.svelte';
  import { notificationsCopy } from '../../lib/account/notifications-copy';
  import { fetchGuildHome, type GuildSettingsData, type GuildSettingsDiscord } from '../../lib/guild/api';
  import { guildDiscordConnectUrl, removeGuildDiscord, testGuildDiscord } from '../../lib/guild/discord-api';
  import { guildBotInstallUrl, removeGuildBot, testGuildBot } from '../../lib/guild/discord-bot-api';
  import { guildBotCopy } from '../../lib/guild/discord-bot-copy';
  import { BOT_FOCUS, clearedBotFailures } from '../../lib/guild/discord-bot-state';
  import { GUILD_BOT_RESERVE, guildBotView } from '../../lib/guild/discord-bot-view';
  import { GuildDiscordReturns } from '../../lib/guild/discord-returns.svelte';
  import {
    DISCORD_FOCUS,
    clearedFailures,
    connectStuck,
    guildDiscordCopy,
    removeRefused,
    removed,
    removedChannel,
    testRefused,
    testSent,
    type DiscordStatus,
  } from '../../lib/guild/discord-settings';
  import { guildDiscordState } from '../../lib/guild/discord-state';
  import { button } from '../../lib/ui/button';
  import DiscordConnectLink from '../account/discord/DiscordConnectLink.svelte';
  import DiscordStatusLine from '../account/discord/DiscordStatusLine.svelte';
  import GuildDiscordAdd from './GuildDiscordAdd.svelte';
  import GuildDiscordBot from './GuildDiscordBot.svelte';
  import GuildDiscordChannel from './GuildDiscordChannel.svelte';
  import GuildPlanGate from './GuildPlanGate.svelte';

  /** How long `Removed.` stays (account spec 4.3); a `stays` line holds until the next press. */
  const STATUS_MS = 4000;
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
  let changing = $state(false);
  let focusSelect = $state(false);
  let pickerSaving = $state(false);
  let section = $state<HTMLElement | null>(null);
  let statusTimer: ReturnType<typeof setTimeout> | undefined;

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
  const returns = new GuildDiscordReturns({
    guildId: () => guildId,
    discord: () => discord,
    onSettings: (fresh) => onSettings(fresh),
    show: (next) => show(next),
    focus: (selector) => focus.request(selector),
    focusSelect: () => (focusSelect = true),
    scrollToSection: () => section?.scrollIntoView({ block: 'start' }),
  });
  const focus = createFocusQueue(() => [
    state,
    view.state,
    confirming,
    busy,
    returns.settling,
    returns.botSettling,
    readOnly,
  ]);
  /** The bot's states hold their reserved heights live, as the chunk placeholder does, so a save
   *  from the can't-be-used notice to state 2 moves the page by the cap change only (A.2, item 17). */
  const botReserve = $derived(!botOn || view.state === 'webhook-only' ? '' : GUILD_BOT_RESERVE[view.state]);
  const held = $derived(returns.held(botReserve));

  $effect(() => () => {
    clearTimeout(statusTimer);
    returns.dispose();
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

  onMount(() => {
    returns.take();
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

  /** Focus returns to the opener: `Choose a channel` on the can't-be-used notice, else `Change channel`.
   *  Read before `changing` drops, which turns the view back from the picker. */
  function cancelChange(): void {
    const opener = view.bot?.channel_unusable != null ? BOT_FOCUS.choose : BOT_FOCUS.change;
    changing = false;
    focusSelect = false;
    focus.request(opener);
  }

  function saved(next: GuildSettingsDiscord): void {
    onDiscord(next);
    changing = false;
    focusSelect = false;
    pickerSaving = false;
    show({ ok: true, message: guildBotCopy.installed.saved, stays: true });
    focus.request(BOT_FOCUS.testButton, { visible: true });
  }

  function onKeydown(event: KeyboardEvent): void {
    if (event.key !== 'Escape' || busy !== null || pickerSaving) return;
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
  <h2 class="section-title">{guildDiscordCopy.heading}</h2>
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
  {#if returns.settling || returns.botSettling}
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
        onSaving={(saving) => (pickerSaving = saving)}
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
        class={button('primary')}
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
