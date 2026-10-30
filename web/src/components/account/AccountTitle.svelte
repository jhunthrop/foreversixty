<!-- web/src/components/account/AccountTitle.svelte -->
<!-- The account page's head, one component so the loading, failed, signed-out and ready states
     draw the same box (the loading reserve pins its height). It is the shared PageHead with the
     account copy; the ready state passes the identity block as the aside. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { accountPageCopy } from '../../lib/account/account-page-copy';
  import PageHead from '../ui/PageHead.svelte';

  let { aside }: { aside?: Snippet } = $props();
</script>

<PageHead eyebrow={accountPageCopy.eyebrow} title={accountPageCopy.title} {aside} />
