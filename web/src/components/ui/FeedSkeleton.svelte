<!-- web/src/components/ui/FeedSkeleton.svelte -->
<!-- A feed (ui/FeedRow) while it loads: `rows` lines the height of a feed row, an icon and a sentence each. -->
<script lang="ts">
  let { rows = 5, testid }: { rows?: number; testid?: string } = $props();
</script>

<div class="flex flex-col" role="status" aria-busy="true" data-testid={testid}>
  {#each Array.from({ length: rows }, (_, index) => index) as index (index)}
    <span class="border-line-soft flex h-[65px] items-center gap-[14px] border-b" aria-hidden="true">
      <span class="skeleton-block block size-10"></span>
      <span class="skeleton-block block h-4 w-3/5"></span>
    </span>
  {/each}
</div>
