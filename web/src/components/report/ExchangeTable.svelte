<!-- web/src/components/report/ExchangeTable.svelte -->
<!-- Interrupts and Dispels are the same shape: someone did something to someone else's
     spell. One component, two tabs.

     summary.ExchangeRow carries no timestamp, so nothing in it can be cut to a brushed
     window: scopeSummary (window.ts) passes `interrupts` and `dispels` straight through,
     untouched. These counts are the whole fight's under every window, and this table says
     so outright rather than letting an unchanged count after a brush read as "nothing
     happened in this window" -- a mark and a note a reader could miss beat a raid leader
     concluding the wrong thing from a number that never moved. The `†` mark and the note
     below are unconditional: unlike CastTable's approximate marks, nothing here depends
     on whether the window happens to be whole. -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import { wholeFightAriaLabel, wholeFightMark, wholeFightTitle } from '../../lib/report/format';
  import type { ExchangeRow } from '../../lib/report/types';

  let { rows, emptyText }: { rows: ExchangeRow[]; emptyText: string } = $props();

  const ordered = $derived(
    [...rows].sort((a, b) => b.count - a.count || a.source_name.localeCompare(b.source_name)),
  );
  const mark = wholeFightMark(true);
  const title = wholeFightTitle(true);
</script>

{#if ordered.length === 0}
  <p class="text-muted text-[14px]" data-testid="table-empty">{emptyText}</p>
{:else}
  <div class="flex flex-col" data-testid="exchange-table">
    <div
      class="text-muted label hidden grid-cols-[minmax(120px,1fr)_minmax(120px,1fr)_minmax(120px,1fr)_minmax(120px,1fr)_64px] gap-x-3 px-2 pb-1 md:grid"
    >
      <span>By</span>
      <span>With</span>
      <span>On</span>
      <span>Spell</span>
      <span class="text-right">Count</span>
    </div>
    <ul class="flex flex-col">
      {#each ordered as row (`${row.source_guid}-${row.target_guid}-${row.spell_id}-${row.extra_spell_id}`)}
        <li
          class="border-line-soft grid min-h-11 grid-cols-1 items-center gap-x-3 gap-y-1 border-b px-2 py-2 text-[14px] md:grid-cols-[minmax(120px,1fr)_minmax(120px,1fr)_minmax(120px,1fr)_minmax(120px,1fr)_64px]"
        >
          <span class="truncate font-semibold">{splitUnitName(row.source_name).name}</span>
          <span class="truncate">{row.spell_name}</span>
          <span class="text-muted truncate text-[13px]">{splitUnitName(row.target_name).name}</span>
          <span class="text-muted truncate text-[13px]">{row.extra_spell_name}</span>
          <span
            class="tabular text-right font-mono"
            {title}
            aria-label={wholeFightAriaLabel(true, String(row.count))}
          >
            {mark}{row.count}
          </span>
        </li>
      {/each}
    </ul>
  </div>
  <p class="text-muted text-[12px]" data-testid="exchange-wholefight-note">
    Count is marked {mark}: interrupts and dispels are the whole fight's totals, because the summary keeps no
    timestamp for them and a brushed window cannot cut them down.
  </p>
{/if}
