<!-- web/src/components/guild/dense/AttendanceStrip.svelte -->
<!-- The last raid nights as cells, oldest first (dense spec 5.3): a tick for present, B for bench, a
     cross for absent, hatched before the raider joined. A glyph in every cell, so the strip reads
     without colour; `legend` draws the key once. -->
<script lang="ts" module>
  import type { AttendanceCell } from '../../../lib/guild/overview-api';

  export const CELL_GLYPH: Readonly<Record<AttendanceCell, string>> = {
    present: '✓',
    bench: 'B',
    absent: '×',
    none: '',
  };
  export const CELL_WORD: Readonly<Record<AttendanceCell, string>> = {
    present: 'Present',
    bench: 'Bench',
    absent: 'Absent',
    none: 'Before joining',
  };
</script>

<script lang="ts">
  let {
    cells = [],
    legend = false,
    testid,
  }: {
    cells?: readonly AttendanceCell[];
    /** Draw the key instead of a raider's nights. */
    legend?: boolean;
    testid?: string;
  } = $props();

  const KEY: readonly AttendanceCell[] = ['present', 'bench', 'absent', 'none'];
</script>

{#if legend}
  <span class="att-legend text-small text-muted" data-testid={testid}>
    {#each KEY as cell (cell)}
      <span class="att-key"
        ><span class="att-cell" data-cell={cell} aria-hidden="true">{CELL_GLYPH[cell]}</span>{CELL_WORD[
          cell
        ]}</span
      >
    {/each}
  </span>
{:else}
  <span
    class="att"
    role="img"
    aria-label={cells.map((c) => CELL_WORD[c]).join(', ')}
    title={cells.map((c) => CELL_WORD[c]).join(', ')}
    data-testid={testid}
  >
    {#each cells as cell, index (index)}<span class="att-cell" data-cell={cell}>{CELL_GLYPH[cell]}</span
      >{/each}
  </span>
{/if}

<style>
  .att,
  .att-legend,
  .att-key {
    display: inline-flex;
    align-items: center;
    vertical-align: middle;
  }
  .att {
    gap: 2px;
  }
  .att-legend {
    flex-wrap: wrap;
    gap: 4px 8px;
  }
  .att-key {
    gap: 4px;
    white-space: nowrap;
  }
  .att-cell {
    display: inline-block;
    width: 14px;
    height: 14px;
    border-radius: 2px;
    font: 700 10px/14px var(--font-body);
    text-align: center;
    color: var(--color-ground);
  }
  .att-cell[data-cell='present'] {
    background: var(--color-kill);
  }
  .att-cell[data-cell='bench'] {
    background: var(--color-gold);
  }
  .att-cell[data-cell='absent'] {
    background: var(--color-wipe);
  }
  .att-cell[data-cell='none'] {
    background: repeating-linear-gradient(45deg, transparent 0 2px, var(--color-muted) 2px 3px);
  }
</style>
