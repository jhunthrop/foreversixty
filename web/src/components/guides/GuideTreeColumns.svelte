<!-- web/src/components/guides/GuideTreeColumns.svelte -->
<!-- A guide build's three talent trees, one column each from md and stacked below, drawn by
     the planner's read-only TreeGrid: shared by the Talents and builds trees and the Leveling
     tree, so both read and tooltip alike. Keyed on the tree id, so a new store (another level)
     updates the cells in place rather than remounting them. -->
<script lang="ts">
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import TreeGrid from '../planner/TreeGrid.svelte';

  let { store }: { store: PlannerStore } = $props();
</script>

{#if store.talentIndex !== null}
  <div class="reveal grid grid-cols-1 gap-4 md:grid-cols-3" data-testid="guide-tree-columns">
    {#each store.talentIndex.trees as tree (tree.id)}
      <div class="flex flex-col gap-2">
        <h3 class="section-title">{tree.name}</h3>
        <TreeGrid {store} {tree} readOnly />
      </div>
    {/each}
  </div>
{/if}
