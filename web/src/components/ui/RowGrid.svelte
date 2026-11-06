<!-- web/src/components/ui/RowGrid.svelte -->
<!-- Open Rows in columns (home spec 2026-11-06 §2): one column below 768, two to 1023, `cols` from
     1024, the column gap the page's --gap. Each Row keeps its own hairline. The home class picker
     first (nine Rows of crests), then any short grid of link rows. No surface of its own. -->
<script lang="ts">
  import type { Snippet } from 'svelte';

  let {
    cols = 3,
    label,
    testid,
    children,
  }: {
    /** Columns from 1024 px. */
    cols?: 2 | 3 | 4;
    /** The list's accessible name, when it is a list of its own. */
    label?: string;
    testid?: string;
    children: Snippet;
  } = $props();
</script>

<div class="ui-row-grid" style:--row-grid-cols={cols} aria-label={label} data-testid={testid}>
  {@render children()}
</div>

<style>
  .ui-row-grid {
    display: grid;
    grid-template-columns: minmax(0, 1fr);
    column-gap: var(--gap);
  }
  /* Every Row keeps its hairline: in columns no Row is the list's last but one. */
  .ui-row-grid > :global(.ui-row) {
    border-bottom-width: 1px;
  }
  @media (min-width: 768px) {
    .ui-row-grid {
      grid-template-columns: repeat(2, minmax(0, 1fr));
    }
  }
  @media (min-width: 1024px) {
    .ui-row-grid {
      grid-template-columns: repeat(var(--row-grid-cols), minmax(0, 1fr));
    }
  }
</style>
