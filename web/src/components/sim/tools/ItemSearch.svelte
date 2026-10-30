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
    noResultsReason,
    searchItems,
    slotOptions,
    type ItemQuery,
    type SearchContext,
  } from '../../../lib/sim/item-search';
  import { legendaryVariantLabel } from '../../../lib/sim/legendary-variants';
  import { groupSources, type LootFile, sourceLabel } from '../../../lib/sim/loot';
  import { poolQualityCopy } from '../../../lib/sim/pool-quality-copy';
  import { selectCopy } from '../../../lib/ui/copy';
  import Select from '../../ui/Select.svelte';

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
  const emptyReason = $derived(noResultsReason(items, query, ctx));

  /** dps D27: name why an off-hand search excludes a two-hander instead of implying the
   *  item does not exist, when the search can prove one of those two things is true. */
  const noResultsMessage = $derived(
    emptyReason.twoHanded
      ? poolQualityCopy.searchWrongSlotTwoHanded
      : emptyReason.existsElsewhere
        ? poolQualityCopy.searchWrongSlotElsewhere
        : bulkCopy.searchNoResults,
  );
</script>

<section
  class="border-line rounded-panel mx-[18px] flex flex-col gap-3 border p-3 md:mx-0"
  data-testid="sim-item-search"
>
  <h3 class="section-title">{bulkCopy.searchLabel}</h3>

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
    <div class="flex items-center gap-2">
      <label class="quiet-label" for="sim-search-slot-field">{bulkCopy.searchSlot}</label>
      <Select
        id="sim-search-slot-field"
        class="w-[160px]"
        size="compact"
        bind:value={query.slot}
        options={[
          { value: '', label: bulkCopy.searchAnySlot },
          ...slotOptions().map((option) => ({ value: option.slot, label: option.label })),
        ]}
        testid="sim-search-slot"
      />
    </div>
    <div class="flex items-center gap-2">
      <label class="quiet-label" for="sim-search-source-field">{bulkCopy.searchSource}</label>
      <Select
        id="sim-search-source-field"
        class="w-[240px]"
        size="compact"
        bind:value={query.sourceId}
        options={[
          { value: '', label: bulkCopy.searchAnySource },
          ...groups.flatMap((group) =>
            group.sources.map((source) => ({
              value: source.id,
              label: sourceLabel(source),
              group: group.label,
            })),
          ),
        ]}
        searchable
        searchPlaceholder={selectCopy.filterSources}
        testid="sim-search-source"
      />
    </div>
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
    <p class="text-muted text-[13px]">{noResultsMessage}</p>
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
            {item.name}{legendaryVariantLabel(item.id) === '' ? '' : ` (${legendaryVariantLabel(item.id)})`}
          </span>
          <span class="text-muted numeric text-[12px]">{item.item_level}</span>
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
