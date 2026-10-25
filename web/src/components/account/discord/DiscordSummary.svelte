<!-- web/src/components/account/discord/DiscordSummary.svelte -->
<!-- What a saved webhook says about itself: where it posts (the webhook's name, then the shortened
     channel id), the 1 to 4 failure warning, and the switched-off or lapsed notice. The name is
     Discord's own text: the copy module cuts it at 40 characters and it is rendered as escaped
     text only. -->
<script lang="ts">
  import type { DiscordNotifications } from '../../../lib/account/api';
  import {
    failureWarning,
    openInDiscordUrl,
    sourceOf,
    type DiscordBlockKey,
  } from '../../../lib/account/notifications';
  import { notificationsCopy } from '../../../lib/account/notifications-copy';
  import DiscordNotice from './DiscordNotice.svelte';
  import DiscordReplace from './DiscordReplace.svelte';

  let {
    state,
    discord,
    replaceHidden = false,
    replaceDisabled = false,
    onReplace,
  }: {
    state: DiscordBlockKey;
    discord: DiscordNotifications;
    /** The remove question is open: the replace link is not offered until it is answered. */
    replaceHidden?: boolean;
    replaceDisabled?: boolean;
    onReplace: () => void;
  } = $props();

  const connected = $derived(sourceOf(discord) === 'connected');
  // A connected webhook's Discord-side name is the application's, so it is never shown (spec 3.2).
  const name = $derived(connected ? '' : (discord.name?.trim() ?? ''));
  // The copy module owns the sentence and its quotes; the name part is emphasised on its own.
  const named = $derived(name === '' ? null : notificationsCopy.destinationNamed(name));
  const openUrl = $derived(openInDiscordUrl(discord));
  const lineTwo = $derived(notificationsCopy.channelLine(discord.channel_id));
  const idAt = $derived(lineTwo.lastIndexOf(' ') + 1);
  const warning = $derived(state === 'saved' ? failureWarning(discord) : null);
</script>

<div class="flex flex-col gap-0.5" data-testid="discord-destination">
  <p class="text-[14px] [overflow-wrap:anywhere]">
    {#if connected}
      <span class="text-muted">{notificationsCopy.destinationConnected.lead}</span>
      <span class="text-strong font-semibold">{notificationsCopy.destinationConnected.name}</span>
    {:else if named !== null}
      <span class="text-muted">{named.lead}</span>
      <span class="text-strong font-semibold">{named.name}</span>
    {:else}
      <span class="text-muted">{notificationsCopy.destinationUnnamed}</span>
    {/if}
  </p>
  <p class="text-muted text-[13px]">
    {lineTwo.slice(0, idAt)}<span class="text-text font-mono whitespace-nowrap tabular-nums"
      >{lineTwo.slice(idAt)}</span
    >{#if openUrl !== null}{' · '}<a
        class="text-text hover:text-gold-hover -my-3 py-3 underline"
        href={openUrl}
        target="_blank"
        rel="noopener"
        data-testid="discord-open">{notificationsCopy.openInDiscord}</a
      >{/if}
  </p>
</div>
{#if warning !== null}
  <div class="flex flex-col gap-0.5 text-[13px]">
    <p class="text-text">{warning.failures}</p>
    {#if warning.lastError !== null}<p class="text-muted">{warning.lastError}</p>{/if}
  </div>
{/if}
{#if state === 'disabled'}
  <DiscordNotice
    tone="wipe"
    title={notificationsCopy.switchedOff}
    detail={discord.last_error === null ? null : notificationsCopy.lastError(discord.last_error)}
  >
    <p class="text-muted text-[13px]">{notificationsCopy.switchedOffFix}</p>
    {#if !replaceHidden}
      <div class="-mt-1">
        <DiscordReplace {discord} disabled={replaceDisabled} {onReplace} />
      </div>
    {/if}
  </DiscordNotice>
{:else if state === 'lapsed'}
  <DiscordNotice tone="gold" title={notificationsCopy.lapsedTitle} detail={notificationsCopy.lapsedBody}>
    <a class="text-text hover:text-gold-hover w-fit text-[13px] underline" href="/premium"
      >{notificationsCopy.seePlans}</a
    >
  </DiscordNotice>
{/if}
