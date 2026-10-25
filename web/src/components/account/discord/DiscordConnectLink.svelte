<!-- web/src/components/account/discord/DiscordConnectLink.svelte -->
<!-- A navigation to the API's connect route (a link, so middle-click and long-press work), drawn
     either as the gold button or as the underlined text link. From the click until the page
     unloads it reads `Opening Discord…`; the back button restores the page from the bfcache and
     `pageshow` with `persisted` puts the label back (spec 4.2, 4.3). -->
<script lang="ts">
  import { CONNECT_PATH } from '../../../lib/account/notifications';
  import { notificationsCopy } from '../../../lib/account/notifications-copy';

  let {
    label,
    class: className,
    testid,
    busy = false,
  }: { label: string; class: string; testid: string; busy?: boolean } = $props();

  let opening = $state(false);

  /** A test or remove call is in flight: the click goes nowhere, as a disabled button's would. */
  function onClick(event: MouseEvent): void {
    if (busy) event.preventDefault();
    else opening = true;
  }

  function onPageShow(event: PageTransitionEvent): void {
    if (event.persisted) opening = false;
  }
</script>

<svelte:window onpageshow={onPageShow} />

<a
  href={CONNECT_PATH}
  class="{className} {opening || busy ? 'pointer-events-none' : ''}"
  aria-busy={opening}
  aria-disabled={opening || busy ? 'true' : undefined}
  onclick={onClick}
  data-testid={testid}
>
  {opening ? notificationsCopy.connecting : label}
</a>
