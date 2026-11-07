<!-- web/src/components/home/HomeDeletedNotice.svelte -->
<!-- The home page's one-time notice after an account deletion: /account lands on `/?deleted=1`,
     this reads the flag once and drops it from the address bar, so a reload shows nothing. Draws
     nothing on any other visit. -->
<script lang="ts">
  import { onMount } from 'svelte';
  import { accountPageCopy } from '../../lib/account/account-page-copy';
  import { takeDeletedArrival } from '../../lib/account/deleted-arrival';
  import Notice from '../ui/Notice.svelte';

  // Read in the browser only, once: the flag leaves the address bar as it is read.
  let shown = $state(false);
  onMount(() => {
    shown = takeDeletedArrival();
  });
</script>

{#if shown}
  <div role="status">
    <Notice text={accountPageCopy.deletedNotice} tone="neutral" testid="home-deleted-notice" />
  </div>
{/if}
