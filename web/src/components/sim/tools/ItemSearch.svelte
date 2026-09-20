<!-- web/src/components/sim/tools/ItemSearch.svelte -->
<!-- Design 3.1.3: name, minimum item level, slot, source, usable-only. Entirely
     client-side over the items file the grid already has -- no API call, and no second
     index. The design's own caveat stands and the page says it: search can find items this
     character has no way to obtain. -->
<script lang="ts">
  import { rarityClassFor } from '../../../lib/planner/items';
  import { dataUrl } from '../../../lib/planner/load';
  import { SECONDARY_BUTTON } from '../../../lib/planner/styles';
  import type { Item } from '../../../lib/planner/types';
  import { bulkCopy } from '../../../lib/sim/copy';
  import {
    defaultItemQuery,
    matchCount,
    searchItems,
    slotOptions,
    type ItemQuery,
    type SearchContext,
  } from '../../../lib/sim/item-search';
  import { groupSources, type LootFile } from '../../../lib/sim/loot';

  let {
    items,
    loot,
    ctx,
    treeVersion,
    onadd,
  }: {
    items: readonly Item[];
    loot: LootFile;
    ctx: SearchContext;
    treeVersion: string;
    onadd: (itemId: number) => void;
  } = $props();

  let query = $state<ItemQuery>(defaultItemQuery());

  const found = $derived(searchItems(items, query, ctx));
  const matching = $derived(matchCount(items, query, ctx));
  const groups = $derived(groupSources(loot));
</script>

<section
  class="border-line rounded-panel mx-[18px] flex flex-col gap-3 border p-3 md:mx-0"
  data-testid="sim-item-search"
>
  <h3 class="section-title text-[14px]">{bulkCopy.searchLabel}</h3>

  <div class="flex flex-wrap gap-2">
    <input
      type="search"
      bind:value={query.text}
      placeholder={bulkCopy.searchPlaceholder}
      aria-label={bulkCopy.searchLabel}
      class="border-line-warm rounded-control bg-bg text-text placeholder:text-muted h-11 min-w-[180px] flex-1 border px-3 text-[14px]"
    />
    <label class="text-muted flex items-center gap-2 text-[12px]">
      {bulkCopy.searchMinItemLevel}
      <input
        type="number"
        min="0"
        max="120"
        bind:value={query.minItemLevel}
        data-testid="sim-search-ilvl"
        class="border-line-warm rounded-control bg-bg text-text h-11 w-20 border px-2 text-[14px]"
      />
    </label>
    <label class="text-muted flex items-center gap-2 text-[12px]">
      {bulkCopy.searchSlot}
      <select
        bind:value={query.slot}
        data-testid="sim-search-slot"
        class="border-line-warm rounded-control bg-bg text-text h-11 border px-2 text-[14px]"
      >
        <option value="">{bulkCopy.searchAnySlot}</option>
        {#each slotOptions() as option (option.slot)}
          <option value={option.slot}>{option.label}</option>
        {/each}
      </select>
    </label>
    <label class="text-muted flex items-center gap-2 text-[12px]">
      {bulkCopy.searchSource}
      <select
        bind:value={query.sourceId}
        data-testid="sim-search-source"
        class="border-line-warm rounded-control bg-bg text-text h-11 border px-2 text-[14px]"
      >
        <option value="">{bulkCopy.searchAnySource}</option>
        {#each groups as group (group.kind)}
          <optgroup label={group.label}>
            {#each group.sources as source (source.id)}
              <option value={source.id}>{source.name}</option>
            {/each}
          </optgroup>
        {/each}
      </select>
    </label>
    <label class="text-muted flex min-h-11 items-center gap-2 text-[12px]">
      <input
        type="checkbox"
        class="h-5 w-5"
        data-testid="sim-search-usable"
        bind:checked={query.usableOnly}
      />
      {bulkCopy.searchUsableOnly}
    </label>
  </div>

  {#if found.length === 0}
    <p class="text-muted text-[13px]">{bulkCopy.searchNoResults}</p>
  {:else}
    <ul class="max-h-[320px] overflow-y-auto">
      {#each found as item (item.id)}
        <li
          class="border-line-soft flex min-h-11 items-center gap-3 border-b px-1 py-1 last:border-b-0"
          data-testid={`sim-search-result-${item.id}`}
        >
          <img
            src={dataUrl(treeVersion, `icons/${item.icon}.webp`)}
            alt=""
            width="24"
            height="24"
            loading="lazy"
            decoding="async"
            class="rounded-control border-line h-6 w-6 border object-cover"
          />
          <span class={`flex-1 text-[14px] font-semibold ${rarityClassFor(item.quality)}`}>
            {item.name}
          </span>
          <span class="tabular text-muted font-mono text-[12px]">{item.item_level}</span>
          <button
            type="button"
            class="{SECONDARY_BUTTON} border-line-warm text-nav px-3"
            data-testid={`sim-search-add-${item.id}`}
            onclick={() => onadd(item.id)}>{bulkCopy.searchAdd}</button
          >
        </li>
      {/each}
    </ul>
    {#if found.length < matching}
      <p class="text-muted text-[12px]">{bulkCopy.searchTruncated(found.length, matching)}</p>
    {/if}
  {/if}
</section>
