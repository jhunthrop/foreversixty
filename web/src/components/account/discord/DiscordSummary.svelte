<!-- web/src/components/account/discord/DiscordSummary.svelte -->
<!-- What a saved webhook says about itself: where it posts (never an id; `Open in Discord ›` beside
     it), the 1 to 4 failure warning, and the switched-off or lapsed notice. The name is Discord's own text: the copy
     module cuts it at 40 characters and it is rendered as escaped text only. -->
<script lang="ts">
  import type { DiscordNotifications } from '../../../lib/account/api';
  import { PREMIUM_LINKS } from '../../../lib/billing/premium-links';
  import {
    failureWarning,
    openInDiscordUrl,
    sourceOf,
    type DiscordBlockKey,
  } from '../../../lib/account/notifications';
  import { notificationsCopy } from '../../../lib/account/notifications-copy';
  import { LINK } from '../../../lib/ui/primitives';
  import DiscordNotice from './DiscordNotice.svelte';

  let { state, discord }: { state: DiscordBlockKey; discord: DiscordNotifications } = $props();

  const connected = $derived(sourceOf(discord) === 'connected');
  // A connected webhook's Discord-side name is the application's, so it is never shown (spec 3.2).
  const name = $derived(connected ? '' : (discord.name?.trim() ?? ''));
  // The copy module owns the sentence and its quotes.
  const named = $derived(name === '' ? null : notificationsCopy.destinationNamed(name));
  const destination = $derived(
    named === null ? notificationsCopy.destinationChannel : `${named.lead} ${named.name}`,
  );
  const openUrl = $derived(openInDiscordUrl(discord));
  const warning = $derived(state === 'saved' ? failureWarning(discord) : null);
</script>

<p
  class="flex flex-wrap items-center gap-x-3 gap-y-2 text-[15px] leading-6"
  data-testid="discord-destination"
>
  <span class="[overflow-wrap:anywhere]">{destination}</span>
  {#if openUrl !== null}
    <a class="{LINK} text-[15px]" href={openUrl} target="_blank" rel="noopener" data-testid="discord-open"
      >{notificationsCopy.openInDiscord} ›</a
    >
  {/if}
</p>
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
  </DiscordNotice>
{:else if state === 'lapsed'}
  <DiscordNotice tone="gold" title={notificationsCopy.lapsedTitle} detail={notificationsCopy.lapsedBody}>
    <a class="{LINK} text-[15px]" href={PREMIUM_LINKS.notifications}>{notificationsCopy.seePlans} ›</a>
  </DiscordNotice>
{/if}
