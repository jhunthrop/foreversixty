<!-- web/src/components/guild/RaidBlock.svelte -->
<!-- One block of the Report (attendance spec 3.4): a 15 px section-title sub-head with its mono
     muted count, a 12 px gap, then its rows. -->
<script lang="ts">
  import type { Snippet } from 'svelte';

  let {
    heading,
    tail = null,
    count = null,
    testid,
    children,
  }: {
    heading: string;
    /** A tail after the heading that breaks as one unit (`since Thu 3 Dec`). */
    tail?: string | null;
    /** The mono muted count after the sub-head; none when null or 0. */
    count?: number | null;
    testid: string;
    children: Snippet;
  } = $props();

  /** An expression, so Svelte keeps the space at the block edge before the tail. */
  const SPACE = ' ';
</script>

<div class="flex min-w-0 flex-col gap-3" data-testid={testid}>
  <h3 class="section-title flex items-baseline gap-2">
    <span
      >{heading}{#if tail !== null}{SPACE}<span class="whitespace-nowrap">{tail}</span>{/if}</span
    >{#if count !== null && count > 0}<span class="text-muted numeric text-[13px]">{count}</span>{/if}
  </h3>
  {@render children()}
</div>
