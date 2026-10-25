<!-- web/src/components/account/discord/DiscordConnectLink.svelte -->
<!-- A navigation to the API's connect route (a link, so middle-click and long-press work), drawn
     either as the gold button or as the underlined text link. From the click until the page
     unloads it reads `Opening Discord…`; the back button restores the page from the bfcache and
     `pageshow` with `persisted` puts the label back (spec 4.2, 4.3). -->
<script lang="ts">
  import { connectUrl } from '../../../lib/account/notifications';
  import { createOpeningWatchdog } from '../../../lib/account/opening-watchdog';
  import { notificationsCopy } from '../../../lib/account/notifications-copy';

  let {
    label,
    class: className,
    testid,
    busy = false,
    onStuck,
  }: { label: string; class: string; testid: string; busy?: boolean; onStuck: () => void } = $props();

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

<a
  href={connectUrl()}
  class="{className} {opening || busy ? 'pointer-events-none' : ''}"
  aria-busy={opening}
  aria-disabled={opening || busy ? 'true' : undefined}
  onclick={onClick}
  data-testid={testid}
>
  {opening ? notificationsCopy.connecting : label}
</a>
