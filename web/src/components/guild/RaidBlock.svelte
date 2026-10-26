<!-- web/src/components/guild/RaidBlock.svelte -->
<!-- One block of the Report (attendance spec 3.4): a 15 px section-title sub-head with its mono
     muted count, a 12 px gap, then its rows. -->
<script lang="ts">
  import type { Snippet } from 'svelte';

  let {
    heading,
    count = null,
    testid,
    children,
  }: {
    heading: string;
    /** The mono muted count after the sub-head; none when null. */
    count?: number | null;
    testid: string;
    children: Snippet;
  } = $props();
</script>

<div class="flex min-w-0 flex-col gap-3" data-testid={testid}>
  <h3 class="section-title flex items-baseline gap-2 text-[15px]">
    {heading}{#if count !== null}<span class="text-muted font-mono text-[13px] tabular-nums">{count}</span
      >{/if}
  </h3>
  {@render children()}
</div>
