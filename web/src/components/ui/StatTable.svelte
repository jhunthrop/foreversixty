<!-- web/src/components/ui/StatTable.svelte -->
<!-- A numeric table (sim spec 2026-11-06 §4.7, ruling S8): a lead icon, the name and its sub line,
     an optional Meter column, then right-aligned figure columns, all on one grid (stat-table.css)
     shared by the heads and every StatRow. On the Plate it is the ability table; `open` (no Plate,
     no side padding) it is the Rage sources and the aura uptimes. An ARIA table: the heads are
     columnheaders, a sortable head holds a button and carries aria-sort; with no `nameHead` the
     table draws no head row (an open list whose columns explain themselves). `wrap` lets a reading
     section's lane carry its rows under the prose, the head row repeated (lib/ui/reading-wrap.ts). -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { ariaSortOf, sortMark, type StatColumn, type StatSort } from '../../lib/ui/stat-table';
  import './stat-table.css';

  let {
    columns,
    nameHead,
    barHead = '',
    bar = true,
    open = false,
    sort,
    onsort,
    label,
    wrap = false,
    testid,
    children,
  }: {
    /** The figure columns, in order; the first is the primary figure. */
    columns: readonly StatColumn[];
    /** The name column's head; absent, the table draws no head row. */
    nameHead?: string;
    /** The bar column's head ("Share of damage"). */
    barHead?: string;
    /** Whether the rows carry a Meter column. */
    bar?: boolean;
    open?: boolean;
    sort?: StatSort;
    /** A sortable head was pressed. */
    onsort?: (column: string) => void;
    /** The table's accessible name. */
    label?: string;
    /** In a reading section's lane: the rows may wrap under the prose, the head row with them. */
    wrap?: boolean;
    testid?: string;
    /** The StatRows. */
    children: Snippet;
  } = $props();
</script>

<div
  class="ui-st"
  role="table"
  aria-label={label}
  data-open={open ? 'true' : undefined}
  data-bar={bar ? 'true' : 'false'}
  style:--st-n={columns.length}
  data-wrap-list={wrap ? '' : undefined}
  data-testid={testid}
>
  {#if nameHead !== undefined}
    <div class="ui-st-h ui-st-c" role="row" data-wrap-repeat={wrap ? '' : undefined}>
      <span aria-hidden="true"></span>
      <span class="ui-st-head" role="columnheader">{nameHead}</span>
      {#if bar}
        {#if barHead === ''}<span aria-hidden="true"></span>
        {:else}<span class="ui-st-head" role="columnheader">{barHead}</span>{/if}
      {/if}
      {#each columns as column (column.id)}
        {@const on = sort?.column === column.id}
        <span
          class="ui-st-head"
          role="columnheader"
          data-align="end"
          data-on={on ? 'true' : undefined}
          aria-sort={ariaSortOf(column, sort)}
        >
          {#if column.sortable === true && onsort !== undefined}
            <button
              type="button"
              class="ui-st-sort"
              data-game-tip={column.tip}
              data-testid={`stat-sort-${column.id}`}
              onclick={() => onsort(column.id)}
              >{column.label}{#if on && sort !== undefined}<span class="ui-st-mark" aria-hidden="true"
                  >{sortMark(sort.direction)}</span
                >{/if}</button
            >
          {:else}{column.label}{/if}
        </span>
      {/each}
    </div>
  {/if}
  {@render children()}
</div>
