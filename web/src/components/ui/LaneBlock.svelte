<!-- web/src/components/ui/LaneBlock.svelte -->
<!-- One labelled block in a reading section's lane (guides spec 2026-11-08 §3A): a 12 px caps
     label, an optional Small meta at its right, then its rows or table. No box. -->
<script lang="ts">
  import type { Snippet } from 'svelte';

  let {
    label,
    meta,
    testid,
    children,
  }: {
    label: string;
    meta?: string;
    testid?: string;
    children: Snippet;
  } = $props();
</script>

<div class="ui-lane-block" data-testid={testid}>
  <div class="ui-lane-head">
    <h3 class="text-label text-muted m-0 font-bold">{label}</h3>
    {#if meta}<span class="text-small text-muted ml-auto text-right">{meta}</span>{/if}
  </div>
  {@render children()}
</div>

<style>
  .ui-lane-block {
    display: flex;
    flex-direction: column;
    gap: 8px;
    min-width: 0;
  }
  .ui-lane-head {
    display: flex;
    align-items: baseline;
    gap: 12px;
  }
</style>
