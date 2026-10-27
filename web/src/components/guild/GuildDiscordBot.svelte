<!-- web/src/components/guild/GuildDiscordBot.svelte -->
<!-- The bot's three states in guild settings (bot spec §3.1): the picker (delegated), installed
     (where raids post, the failure warning, the linked members, the actions or the remove
     question) and removed in Discord (a notice and the add link). The section owns the status,
     the busy flag and every call; this draws and reports presses. Officer-read-only: test and
     change are off, `Remove` stays live; a frozen claim turns every control off. -->
<script lang="ts">
  import { failureWarning, openInDiscordUrl } from '../../lib/account/notifications';
  import type { GuildBot, GuildSettingsDiscord } from '../../lib/guild/api';
  import { guildBotCopy } from '../../lib/guild/discord-bot-copy';
  import type { GuildBotView } from '../../lib/guild/discord-bot-view';
  import type { DiscordStatus } from '../../lib/guild/discord-settings';
  import { PRIMARY_BUTTON_FIXED, SECONDARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import DiscordConnectLink from '../account/discord/DiscordConnectLink.svelte';
  import DiscordNotice from '../account/discord/DiscordNotice.svelte';
  import DiscordStatusLine from '../account/discord/DiscordStatusLine.svelte';
  import GuildBotRemoveQuestion from './GuildBotRemoveQuestion.svelte';
  import GuildDiscordChannelPick from './GuildDiscordChannelPick.svelte';

  const OUTLINE = `${SECONDARY_BUTTON_FIXED} border-line-warm text-text px-4`;
  const DIM = 'cursor-not-allowed opacity-50';

  let {
    guildId,
    discord,
    bot,
    botState,
    installHref,
    changing,
    focusSelect,
    onSaving,
    readOnly,
    frozen,
    gateId,
    confirming,
    busy,
    status,
    onStatus,
    onSaved,
    onTest,
    onChange,
    onCancelChange,
    onAskRemove,
    onRemove,
    onCancelRemove,
    onStuck,
  }: {
    guildId: number;
    discord: GuildSettingsDiscord;
    bot: GuildBot;
    botState: Extract<GuildBotView, { bot: GuildBot }>['state'];
    installHref: string;
    changing: boolean;
    focusSelect: boolean;
    /** The picker's save in flight: Escape must not unmount it and lose a refusal. */
    onSaving: (saving: boolean) => void;
    readOnly: boolean;
    frozen: boolean;
    gateId: string | undefined;
    confirming: boolean;
    busy: 'test' | 'remove' | null;
    status: DiscordStatus | null;
    onStatus: (status: DiscordStatus | null, focusLine?: boolean) => void;
    onSaved: (discord: GuildSettingsDiscord) => void;
    onTest: () => void;
    onChange: () => void;
    onCancelChange: () => void;
    onAskRemove: () => void;
    onRemove: () => void;
    onCancelRemove: () => void;
    onStuck: () => void;
  } = $props();

  const gated = $derived(readOnly || frozen);
  const channelName = $derived(bot.channel_name ?? '');
  const openUrl = $derived(openInDiscordUrl({ guild_id: bot.server_id, channel_id: bot.channel_id ?? '' }));
  const warning = $derived(failureWarning(bot));
  const off = $derived(busy !== null || frozen);
  const dimmed = $derived(gated ? DIM : '');
</script>

{#if botState === 'picker'}
  <GuildDiscordChannelPick
    {guildId}
    server={bot.server_name}
    currentChannelId={bot.channel_id}
    {changing}
    webhook={discord.connected}
    {gated}
    {frozen}
    {status}
    {confirming}
    removing={busy === 'remove'}
    {focusSelect}
    unusable={bot.channel_unusable}
    {onSaving}
    {onStatus}
    {onSaved}
    onCancel={onCancelChange}
    {onAskRemove}
    {onRemove}
    {onCancelRemove}
  />
{:else if botState === 'unusable' && bot.channel_unusable !== null}
  {@const unusable = bot.channel_unusable}
  {@const title = guildBotCopy.channelUnusable.title(unusable.reason, bot.channel_name, bot.server_name)}
  <DiscordNotice
    tone="gold"
    title={title.text}
    strong={title.strong}
    detail={guildBotCopy.channelUnusable.detail(unusable.reason, unusable.missing, discord.connected)}
  />
  <div class="flex flex-col">
    <div class="flex flex-wrap items-center gap-3">
      <button
        class="{PRIMARY_BUTTON_FIXED} px-4 {dimmed}"
        disabled={off || gated}
        aria-describedby={gateId}
        onclick={onChange}
        data-testid="bot-choose-channel"
      >
        {guildBotCopy.channelUnusable.choose}
      </button>
      {#if unusable.reason === 'permissions' && openUrl !== null}
        <a
          class="text-text hover:text-gold-hover -my-3 py-3 text-[13px] underline"
          href={openUrl}
          target="_blank"
          rel="noopener"
          data-testid="bot-open">{guildBotCopy.channelUnusable.open}</a
        >
      {/if}
    </div>
    <div id="guild-discord-status" tabindex="-1" class="outline-none"><DiscordStatusLine {status} /></div>
  </div>
{:else if botState === 'removed'}
  <DiscordNotice
    tone="gold"
    title={guildBotCopy.removedInDiscord.title(bot.server_name)}
    detail={guildBotCopy.removedInDiscord.detail +
      (discord.connected
        ? guildBotCopy.removedInDiscord.withWebhook
        : guildBotCopy.removedInDiscord.withoutWebhook)}
  />
  <div class="flex flex-col">
    <DiscordConnectLink
      label={guildBotCopy.add}
      class="{PRIMARY_BUTTON_FIXED} w-fit px-5 whitespace-nowrap"
      testid="bot-add"
      href={installHref}
      {gated}
      describedBy={gateId}
      {onStuck}
    />
    <div id="guild-discord-status" tabindex="-1" class="outline-none"><DiscordStatusLine {status} /></div>
  </div>
{:else}
  <div class="flex flex-col gap-0.5" data-testid="bot-destination">
    <p class="text-[14px]">
      <span class="text-muted">{guildBotCopy.installed.lead}</span>
      <span class="text-strong font-semibold">{guildBotCopy.installed.channel(channelName)}</span>
      <span class="text-muted">{guildBotCopy.installed.in}</span>
      <span class="text-strong font-semibold">{bot.server_name}</span>
    </p>
    <p class="text-muted text-[13px]">
      {guildBotCopy.installed.withButtons}{#if openUrl !== null}&#32;·&#32;<a
          class="text-text hover:text-gold-hover -my-3 py-3 underline"
          href={openUrl}
          target="_blank"
          rel="noopener"
          data-testid="bot-open">{guildBotCopy.installed.open}</a
        >{/if}
    </p>
  </div>
  {#if warning !== null}
    <div class="flex flex-col gap-0.5 text-[13px]" data-testid="bot-warning">
      <p class="text-text">{warning.failures}</p>
      {#if warning.lastError !== null}<p class="text-muted">{warning.lastError}</p>{/if}
    </div>
  {/if}
  <p class="text-muted text-[13px]" data-testid="bot-linked-members">
    {guildBotCopy.installed.linkedMembers(bot.linked_members, bot.members)}
  </p>
  <div class="flex flex-col">
    {#if readOnly}
      <p class="text-muted mb-3 text-[13px]" data-testid="bot-read-only">{guildBotCopy.installed.readOnly}</p>
    {/if}
    {#if confirming}
      <GuildBotRemoveQuestion
        server={bot.server_name}
        webhook={discord.connected}
        busy={busy === 'remove'}
        disabled={frozen}
        {onRemove}
        onCancel={onCancelRemove}
      />
    {:else}
      <div class="flex flex-wrap gap-3">
        <button
          class="{OUTLINE} {dimmed}"
          disabled={off || readOnly}
          aria-describedby={gateId}
          aria-busy={busy === 'test'}
          onclick={onTest}
          data-testid="bot-test"
        >
          {busy === 'test' ? guildBotCopy.installed.testing : guildBotCopy.installed.test}
        </button>
        <button
          class="{OUTLINE} {dimmed}"
          disabled={off || readOnly}
          aria-describedby={gateId}
          onclick={onChange}
          data-testid="bot-change"
        >
          {guildBotCopy.installed.change}
        </button>
        <button
          class="{OUTLINE} {frozen ? DIM : ''}"
          disabled={off}
          onclick={onAskRemove}
          data-testid="bot-remove"
        >
          {guildBotCopy.installed.remove}
        </button>
      </div>
    {/if}
    <div id="guild-discord-status" tabindex="-1" class="outline-none"><DiscordStatusLine {status} /></div>
    {#if discord.connected}
      <p class="text-muted text-[13px]" data-testid="bot-fallback-note">
        {guildBotCopy.installed.fallbackNote}
      </p>
    {/if}
  </div>
{/if}
