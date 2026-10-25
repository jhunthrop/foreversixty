<!-- web/src/components/character/HistoryReserve.svelte -->
<!-- The History section's reserved place (experience spec 2026-10-24-character-history 3.11):
     an empty box that global.css draws only under the pre-paint owner hint, at the stored or
     fallback height the section itself uses. `afterSkeleton` is the page's loading skeleton
     standing in for the page above it: the box sits a page gap below it, so the layouts before
     and after the data agree. -->
<script lang="ts">
  let { afterSkeleton = false }: { afterSkeleton?: boolean } = $props();
</script>

<div
  class="history-shell-reserve history-reserve {afterSkeleton ? 'history-after-skeleton' : '-order-2'}"
  aria-hidden="true"
  data-testid="history-reserve"
></div>
