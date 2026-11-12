<!-- web/src/components/ui/LoadOlder.svelte -->
<!-- The one "older" control (component audit 2026-11-10 cluster 16) under a list that reads its
     next page from the API: the small secondary Button, busy (spinner and "Loading…") while the page
     is in flight, and on a failure the muted line with Try again in its place. The caller hides it
     when there is no next page. -->
<script lang="ts">
  import { uiCopy } from '../../lib/ui/copy';
  import Button from './Button.svelte';
  import LoadError from './LoadError.svelte';

  let {
    label = uiCopy.loadOlder,
    loading,
    failed = false,
    failedText = uiCopy.olderFailed,
    onLoad,
    testid,
  }: {
    /** "Older reports", "Older reviews". */
    label?: string;
    loading: boolean;
    failed?: boolean;
    failedText?: string;
    onLoad: () => void;
    testid: string;
  } = $props();
</script>

<div class="pt-2">
  {#if failed && !loading}
    <LoadError message={failedText} onRetry={onLoad} testid="{testid}-error" />
  {:else}
    <Button size="sm" busy={loading} busyLabel={uiCopy.loadingOlder} onclick={onLoad} {testid}>{label}</Button
    >
  {/if}
</div>
