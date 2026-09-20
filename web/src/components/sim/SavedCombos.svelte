<!-- web/src/components/sim/SavedCombos.svelte -->
<!-- A saved Top Gear, talent compare or Droptimizer result at /sim/<id>: read-only, so the
     ranked table and the per-slot summary without the "open in planner" and "copy to
     addon" buttons -- those act on the reader's own character, and a saved page is not
     theirs. "Run this yourself" beside the header (SavedSim.svelte's own button) is the
     way back to a page that is. Task 20.

     The item file is fetched here for icons and rarity colours only -- every name on the
     page comes off the result (contract 10.1 A6), so a saved sim whose class file 404s
     still reads correctly, with slot names and ids in place of icons. -->
<script lang="ts">
  import { loadItems } from '../../lib/planner/load';
  import { SLOT_LABELS, type Item, type Slot } from '../../lib/planner/types';
  import type { BulkResult } from '../../lib/sim/bulk-types';
  import { comboRows, deltaLabel, slotSummary } from '../../lib/sim/combos';
  import { bulkCopy } from '../../lib/sim/copy';
  import { confidenceBand } from '../../lib/sim/estimate';
  import SubstitutionChips from './tools/SubstitutionChips.svelte';

  let { result, treeVersion }: { result: BulkResult; treeVersion: string } = $props();

  let items = $state<Map<number, Item>>(new Map());

  // One load, for this component's one (unchanging) result -- not an `$effect`, the same
  // reason SavedSim.svelte's own item-file load is a plain call rather than one.
  void (async () => {
    try {
      const file = await loadItems(treeVersion, result.request.character.class);
      items = new Map(file.items.map((item) => [item.id, item]));
    } catch {
      // The chips and the slot summary render off the result's own names regardless
      // (contract 10.1 A6); a failed fetch only costs the icon and the rarity colour.
      items = new Map();
    }
  })();

  const rows = $derived(comboRows(result));
  const summary = $derived(slotSummary(result));
  const equippedFigure = $derived(Math.round(result.equipped.mean).toLocaleString('en-US'));
  const equippedBand = $derived(Math.round(confidenceBand(result.equipped)).toLocaleString('en-US'));
</script>

<section class="mx-[18px] flex flex-col gap-3 md:mx-0" data-testid="sim-combos">
  <p class="tabular text-strong font-mono text-[14px]" data-testid="sim-equipped-line">
    {bulkCopy.resultsEquipped}: {equippedFigure} ± {equippedBand}
    <span class="text-muted">{bulkCopy.ranAtStages(result.stages)}</span>
  </p>

  {#if rows.length === 0}
    <p class="text-muted text-[13px]">{bulkCopy.noGain}</p>
  {:else}
    <div role="table" class="flex flex-col">
      <div
        class="text-muted grid grid-cols-[28px_minmax(0,2fr)_84px_96px_56px] items-center gap-x-3 px-2 pb-1 text-[11px] tracking-wide uppercase"
        role="row"
      >
        <span role="columnheader">{bulkCopy.resultsRank}</span>
        <span role="columnheader">{bulkCopy.resultsChange}</span>
        <span role="columnheader" class="ml-auto">{bulkCopy.resultsDps}</span>
        <span role="columnheader" class="ml-auto">{bulkCopy.resultsDelta}</span>
        <span role="columnheader" class="ml-auto">{bulkCopy.resultsPercent}</span>
      </div>
      {#each rows as row, index (row.combo.substitutions
        .map((sub) => `${sub.kind}:${sub.slot ?? ''}:${sub.item_id ?? sub.name ?? ''}`)
        .join('|'))}
        {@const groupEnd = index === rows.length - 1 || rows[index + 1].combo.group !== row.combo.group}
        <div
          class="grid min-h-11 grid-cols-[28px_minmax(0,2fr)_84px_96px_56px] items-center gap-x-3 border-b px-2 py-2 {groupEnd
            ? 'border-line-warm'
            : 'border-line-soft'}"
          role="row"
          aria-label={row.withinError ? bulkCopy.withinError : undefined}
          data-testid="sim-combo-row"
        >
          <span
            class="tabular text-muted font-mono text-[12px]"
            role="cell"
            aria-label={`${bulkCopy.resultsRank} ${row.rank}`}
          >
            {row.rank}
          </span>
          <span role="cell">
            <SubstitutionChips substitutions={row.combo.substitutions} {items} {treeVersion} />
          </span>
          <span
            class="tabular text-strong ml-auto font-mono text-[13px]"
            role="cell"
            aria-label={`${bulkCopy.resultsDps} ${Math.round(row.combo.dps.mean).toLocaleString('en-US')}`}
          >
            {Math.round(row.combo.dps.mean).toLocaleString('en-US')}
          </span>
          <span
            class="tabular text-gold ml-auto font-mono text-[13px]"
            role="cell"
            aria-label={`${bulkCopy.resultsDelta} ${deltaLabel(row.combo.delta)}`}
          >
            {deltaLabel(row.combo.delta)}
          </span>
          <span
            class="tabular text-muted ml-auto font-mono text-[12px]"
            role="cell"
            aria-label={`${bulkCopy.resultsPercent} ${row.percent.toFixed(1)}%`}
          >
            {row.percent.toFixed(1)}%
          </span>
        </div>
      {/each}
    </div>
    {#if rows.length > 1}
      <p class="text-muted text-[12px]">{bulkCopy.withinErrorNote}</p>
    {/if}
  {/if}

  {#if summary.length > 0}
    <section class="border-line rounded-panel border p-3" data-testid="sim-slot-summary">
      <h3 class="section-title text-[14px]">{bulkCopy.slotSummary}</h3>
      <p class="text-muted text-[12px]">{bulkCopy.slotSummaryNote}</p>
      <ul class="flex flex-col">
        {#each summary as row (`${row.slot}-${row.item_id}`)}
          <li class="border-line-soft flex min-h-11 items-center gap-3 border-b px-2 last:border-b-0">
            <span class="text-muted w-24 text-[12px]">{SLOT_LABELS[row.slot as Slot] ?? row.slot}</span>
            <span class="text-text flex-1 text-[13px]">{row.name}</span>
            <span class="tabular text-gold font-mono text-[13px]">
              {row.gain === null ? '—' : `+${Math.round(row.gain).toLocaleString('en-US')}`}
            </span>
          </li>
        {/each}
      </ul>
    </section>
  {/if}
</section>
