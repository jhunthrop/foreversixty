<!-- web/src/components/ui/ColumnHeads.svelte -->
<!-- A ranked list's one head row (tiers spec 2026-11-06 §4.4): 12 px caps labels on the RankList's
     grid, one line each, numeric heads at the end. Hidden below 1024, where each figure carries its
     unit. Decorative for assistive tech: every cell in a row names its own column. -->
<script lang="ts">
  import type { RankHead } from '../../lib/ui/rank-list';
  import './rank-list.css';

  let { cells, testid }: { cells: readonly RankHead[]; testid?: string } = $props();
</script>

<div class="ui-rank-heads ui-rank-cols" aria-hidden="true" data-testid={testid}>
  {#each cells as cell, index (index)}
    <span
      class="text-label tracking-label text-muted font-bold uppercase"
      data-align={cell.align === 'end' ? 'end' : undefined}
      title={cell.label}>{cell.label}</span
    >
  {/each}
</div>
