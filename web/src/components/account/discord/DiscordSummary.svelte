<!-- web/src/components/account/discord/DiscordSummary.svelte -->
<!-- What a saved webhook says about itself: where it posts, the 1 to 4 failure warning, and the
     switched-off or lapsed notice. The channel id is not a credential and is shown whole. -->
<script lang="ts">
  import type { DiscordNotifications } from '../../../lib/account/api';
  import { channelLabel, failureWarning, type DiscordBlockKey } from '../../../lib/account/notifications';
  import { notificationsCopy } from '../../../lib/account/notifications-copy';
  import DiscordNotice from './DiscordNotice.svelte';

  let { state, discord }: { state: DiscordBlockKey; discord: DiscordNotifications } = $props();

  const destination = $derived(channelLabel(discord.channel_id));
  const warning = $derived(state === 'saved' ? failureWarning(discord) : null);
</script>

<p class="text-[14px]" data-testid="discord-destination">
  <span class="text-muted">{notificationsCopy.destinationLead}</span>
  {#if destination !== null}
    <span>{notificationsCopy.destinationNoun}</span>
    <span class="text-strong font-mono text-[13px] whitespace-nowrap tabular-nums">{discord.channel_id}</span>
  {/if}
</p>
{#if warning !== null}
  <p class="text-muted text-[13px]">{warning}</p>
{/if}
{#if state === 'disabled'}
  <DiscordNotice
    tone="wipe"
    title={discord.disabled_reason === 'failed five times'
      ? notificationsCopy.switchedOff
      : notificationsCopy.switchedOffGeneric}
    detail={discord.last_error === null ? null : notificationsCopy.lastError(discord.last_error)}
  />
  <p class="text-muted text-[13px]">{notificationsCopy.switchedOffHelper}</p>
{:else if state === 'lapsed'}
  <DiscordNotice tone="gold" title={notificationsCopy.lapsedTitle} detail={notificationsCopy.lapsedBody}>
    <a class="text-text hover:text-gold-hover w-fit text-[13px] underline" href="/premium"
      >{notificationsCopy.seePlans}</a
    >
  </DiscordNotice>
{/if}
