<!-- web/src/components/guild/GuildDiscordChannel.svelte -->
<!-- A saved guild channel (raids spec §3.8, Connected and Switched off): where it posts, the 1 to 4
     failure warning or the switched-off notice, the actions row (or the remove question), the
     status line, and `Connect a different channel` with its note. The status line is a focus
     target for the plan-ended line, so its wrapper takes tabindex -1. -->
<script lang="ts">
  import { failureWarning, openInDiscordUrl } from '../../lib/account/notifications';
  import { notificationsCopy } from '../../lib/account/notifications-copy';
  import type { GuildSettingsDiscord } from '../../lib/guild/api';
  import { guildDiscordCopy, type DiscordStatus } from '../../lib/guild/discord-settings';
  import DiscordActions from '../account/discord/DiscordActions.svelte';
  import DiscordNotice from '../account/discord/DiscordNotice.svelte';
  import DiscordReplace from '../account/discord/DiscordReplace.svelte';
  import DiscordStatusLine from '../account/discord/DiscordStatusLine.svelte';

  let {
    discord,
    switchedOff,
    connectHref,
    confirming,
    busy,
    status,
    gateId,
    frozen,
    onTest,
    onAskRemove,
    onRemove,
    onCancel,
    onStuck,
  }: {
    discord: GuildSettingsDiscord;
    switchedOff: boolean;
    connectHref: string;
    confirming: boolean;
    busy: 'test' | 'remove' | null;
    status: DiscordStatus | null;
    /** The gate sentence's id in officer-read-only: tests and connect are off, Remove stays. */
    gateId: string | undefined;
    frozen: boolean;
    onTest: () => void;
    onAskRemove: () => void;
    onRemove: () => void;
    onCancel: () => void;
    onStuck: () => void;
  } = $props();

  const channelId = $derived(discord.channel_id ?? '');
  const lineTwo = $derived(notificationsCopy.channelLine(channelId));
  const idAt = $derived(lineTwo.lastIndexOf(' ') + 1);
  const openUrl = $derived(openInDiscordUrl({ guild_id: discord.discord_guild_id, channel_id: channelId }));
  const warning = $derived(switchedOff ? null : failureWarning(discord));
  const gated = $derived(frozen || gateId !== undefined);
</script>

{#snippet replace()}
  <DiscordReplace
    source="connected"
    disabled={busy !== null}
    href={connectHref}
    {gated}
    describedBy={gateId}
    {onStuck}
  />
{/snippet}

<div class="flex flex-col gap-0.5" data-testid="discord-destination">
  <p class="text-[14px]">
    <span class="text-muted">{guildDiscordCopy.destination.lead}</span>
    <span class="text-strong font-semibold">{guildDiscordCopy.destination.name}</span>
  </p>
  <p class="text-muted text-[13px]">
    {lineTwo.slice(0, idAt)}<span class="text-text font-mono whitespace-nowrap tabular-nums"
      >{lineTwo.slice(idAt)}</span
    >{#if openUrl !== null}&#32;·&#32;<a
        class="text-text hover:text-gold-hover -my-3 py-3 underline"
        href={openUrl}
        target="_blank"
        rel="noopener"
        data-testid="discord-open">{notificationsCopy.openInDiscord}</a
      >{/if}
  </p>
</div>
{#if warning !== null}
  <div class="flex flex-col gap-0.5 text-[13px]" data-testid="discord-warning">
    <p class="text-text">{warning.failures}</p>
    {#if warning.lastError !== null}<p class="text-muted">{warning.lastError}</p>{/if}
  </div>
{/if}
{#if switchedOff}
  <DiscordNotice
    tone="wipe"
    title={notificationsCopy.switchedOff}
    detail={discord.last_error === null ? null : notificationsCopy.lastError(discord.last_error)}
  >
    <p class="text-muted text-[13px]">{notificationsCopy.switchedOffFix}</p>
    {#if !confirming}<div class="-mt-1">{@render replace()}</div>{/if}
  </DiscordNotice>
{/if}
<div class="flex flex-col">
  <DiscordActions
    state={switchedOff ? 'disabled' : 'saved'}
    {confirming}
    removeQuestion={guildDiscordCopy.removeQuestion}
    {busy}
    {gateId}
    disabled={frozen}
    {onTest}
    {onAskRemove}
    {onRemove}
    {onCancel}
  />
  <div id="guild-discord-status" tabindex="-1" class="outline-none"><DiscordStatusLine {status} /></div>
  {#if !switchedOff && !confirming}{@render replace()}{/if}
</div>
