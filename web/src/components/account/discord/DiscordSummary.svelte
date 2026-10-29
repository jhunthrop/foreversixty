<!-- web/src/components/account/discord/DiscordSummary.svelte -->
<!-- What a saved webhook says about itself: a destination row (hash tile; where it posts as the
     label, never an id; `Open in Discord` beside it), the 1 to 4 failure
     warning, and the switched-off or lapsed notice. The name is Discord's own text: the copy
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
  import { button } from '../../../lib/ui/button';
  import { ACTION_LINK } from '../../../lib/ui/link';
  import Glyph from '../../ui/Glyph.svelte';
  import SettingsRow from '../../ui/SettingsRow.svelte';
  import DiscordNotice from './DiscordNotice.svelte';
  import DiscordReplace from './DiscordReplace.svelte';

  let {
    state,
    discord,
    replaceHidden = false,
    replaceDisabled = false,
    onReplace,
    onStuck,
  }: {
    state: DiscordBlockKey;
    discord: DiscordNotifications;
    /** The remove question is open: the replace row is not offered until it is answered. */
    replaceHidden?: boolean;
    replaceDisabled?: boolean;
    onReplace: () => void;
    onStuck: () => void;
  } = $props();

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

<div data-testid="discord-destination">
  <SettingsRow label={destination} glyph="hash" controlBelow={true}>
    {#snippet control()}
      {#if openUrl !== null}
        <a
          class={button('secondary', 'sm', true)}
          href={openUrl}
          target="_blank"
          rel="noopener"
          data-testid="discord-open">{notificationsCopy.openInDiscord}<Glyph name="external" /></a
        >
      {/if}
    {/snippet}
  </SettingsRow>
</div>
{#if warning !== null || state === 'disabled' || state === 'lapsed'}
  <div class="flex flex-col gap-3 px-5 pt-4">
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
        <a class="{ACTION_LINK} -my-2" href={PREMIUM_LINKS.notifications}
          >{notificationsCopy.seePlans}<Glyph name="chevron" /></a
        >
      </DiscordNotice>
    {/if}
  </div>
{/if}
{#if state === 'disabled' && !replaceHidden}
  <div class="mt-2">
    <DiscordReplace
      variant="row"
      source={sourceOf(discord)}
      disabled={replaceDisabled}
      {onReplace}
      {onStuck}
    />
  </div>
{/if}
