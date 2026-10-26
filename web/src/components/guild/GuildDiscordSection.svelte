<!-- web/src/components/guild/GuildDiscordSection.svelte -->
<!-- Guild settings · Discord (raids spec §3.8): the guild's one channel for raid posts, a flat
     section under Billing. Not connected: the gold connect link to the guild's route on the API
     origin. Connected and switched off: GuildDiscordChannel. Officer-read-only draws the gate
     sentence, a connect link that goes nowhere and disabled tests, and keeps Remove live; a frozen
     claim turns every control off. A return from Discord (`?discord=`) is read once, dropped from
     the address with `#discord` kept, and the section scrolled to; `connected` holds the connected
     height until the settings read answers. Words and outcomes: lib/guild/discord-settings.ts. -->
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
  let section = $state<HTMLElement | null>(null);
  let statusTimer: ReturnType<typeof setTimeout> | undefined;
  let settleTimer: ReturnType<typeof setTimeout> | undefined;

  const state = $derived(guildDiscordState(discord));
  const tools = $derived(officerTools ?? homeTools ?? true);
  const readOnly = $derived(!frozen && (planEnded || !tools));
  const gateId = $derived(readOnly ? GATE_ID : undefined);
  const connectHref = $derived(guildDiscordConnectUrl(guildId));
  const gateLead = $derived(
    guildDiscordCopy.gate.slice(0, guildDiscordCopy.gate.indexOf(guildDiscordCopy.gateLink)),
  );
  const focus = createFocusQueue(() => [state, confirming, busy, settling, readOnly]);

  $effect(() => () => {
    clearTimeout(statusTimer);
    clearTimeout(settleTimer);
  });

  function show(next: DiscordStatus): void {
    clearTimeout(statusTimer);
    status = next;
    if (next.ok && next.stays !== true) statusTimer = setTimeout(() => (status = null), STATUS_MS);
  }

  function reveal(outcome: GuildReturnOutcome): void {
    clearTimeout(settleTimer);
    settling = false;
    show(outcome.status);
    focus.request(outcome.focus);
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

  onMount(() => {
    takeReturn();
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

  async function test(): Promise<void> {
    const before = discord;
    const error = await run('test', () => testGuildDiscord(guildId));
    if (error === null) {
      if (before.failures > 0 || before.disabled_reason !== null) onDiscord(clearedFailures(before));
      return show(testSent(before.disabled_reason !== null));
    }
    const refused = testRefused(error);
    show(refused);
    if (refused.planEnded !== true) return;
    planEnded = true;
    onPlanRefused();
    focus.request(DISCORD_FOCUS.status);
  }

  async function remove(): Promise<void> {
    const error = await run('remove', () => removeGuildDiscord(guildId));
    if (error !== null) return show(removeRefused());
    confirming = false;
    onDiscord(removedChannel(discord));
    show(removed());
    focus.request(DISCORD_FOCUS.connect);
  }

  function askRemove(): void {
    status = null;
    confirming = true;
    focus.request(DISCORD_FOCUS.removeCancel);
  }

  function cancel(): void {
    confirming = false;
  }

  function onKeydown(event: KeyboardEvent): void {
    if (event.key === 'Escape' && confirming && busy === null) cancel();
  }
</script>

<section
  id="discord"
  bind:this={section}
  class="flex scroll-mt-20 flex-col gap-3 {settling ? GUILD_DISCORD_RESERVE.connected : ''}"
  data-testid="guild-settings-discord"
>
  <h2 class="section-title text-[18px]">{guildDiscordCopy.heading}</h2>
  {#if readOnly}<GuildPlanGate tab="discord" settingsHref="#billing" sentence={gateLead} />{/if}
  {#if settling}
    <!-- the connected height, held until the settings read answers (spec 3.8, Return) -->
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
      />
    </div>
  {/if}
</section>
