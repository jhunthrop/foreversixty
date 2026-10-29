<!-- web/src/components/account/discord/DiscordConnectLink.svelte -->
<!-- A navigation to the API's connect route (a link, so middle-click and long-press work), drawn
     either as the gold button or as the underlined text link. From the click until the page
     unloads it reads `Opening Discord…`; the back button restores the page from the bfcache and
     `pageshow` with `persisted` puts the label back (spec 4.2, 4.3). `href` defaults to the
     account's connect route; the guild passes its own. `gated` (the guild's officer-read-only or
     frozen claim) draws a link that goes nowhere: no href, aria-disabled, out of the tab order,
     described by `describedBy`. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { connectUrl } from '../../../lib/account/notifications';
  import { createOpeningWatchdog } from '../../../lib/account/opening-watchdog';
  import { notificationsCopy } from '../../../lib/account/notifications-copy';

  let {
    label,
    class: className,
    testid,
    busy = false,
    href = connectUrl(),
    gated = false,
    describedBy,
    onStuck,
    content,
  }: {
    label: string;
    class: string;
    testid: string;
    busy?: boolean;
    href?: string;
    gated?: boolean;
    describedBy?: string;
    onStuck: () => void;
    /** Draws the link's inside (the row variant); it receives the label to show. Default: the label text. */
    content?: Snippet<[string]>;
  } = $props();

  let opening = $state(false);
  // No navigation 10 s after the click means no `pageshow` will come: restore the label and say so.
  const watchdog = createOpeningWatchdog(() => {
    opening = false;
    onStuck();
  });

  /** A test or remove call is in flight: the click goes nowhere, as a disabled button's would. */
  function onClick(event: MouseEvent): void {
    if (busy) event.preventDefault();
    else {
      opening = true;
      watchdog.start();
    }
  }

  function onPageShow(event: PageTransitionEvent): void {
    if (!event.persisted) return;
    watchdog.cancel();
    opening = false;
  }

  $effect(() => () => watchdog.cancel());
</script>

<svelte:window onpageshow={onPageShow} />

{#if gated}
  <a
    class="{className} pointer-events-none opacity-50"
    aria-disabled="true"
    aria-describedby={describedBy}
    tabindex="-1"
    data-testid={testid}>{label}</a
  >
{:else}
  <a
    {href}
    class="{className} {opening || busy ? 'pointer-events-none' : ''}"
    aria-busy={opening}
    aria-disabled={opening || busy ? 'true' : undefined}
    onclick={onClick}
    data-testid={testid}
  >
    {#if content !== undefined}
      {@render content(opening ? notificationsCopy.connecting : label)}
    {:else}
      {opening ? notificationsCopy.connecting : label}
    {/if}
  </a>
{/if}
