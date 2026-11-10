<!-- web/src/components/sim/tools/ItemSearch.svelte -->
<!-- Design 3.1.3: name, minimum item level, slot, source, usable-only. Entirely
     client-side over the items file the grid already has -- no API call, and no second
     index. The design's own caveat stands and the page says it: search can find items this
     character has no way to obtain. -->
<script lang="ts">
  import { rarityClassFor } from '../../../lib/planner/items';
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
  import { FIELD_CONTROL, FIELD_FOCUS, FIELD_HEIGHT, FIELD_LABEL, fieldEdge } from '../../../lib/ui/field';
  import Button from '../../ui/Button.svelte';
  import Field from '../../ui/Field.svelte';
  import ItemIcon from '../../ui/ItemIcon.svelte';
  import Row from '../../ui/Row.svelte';
  import Select from '../../ui/Select.svelte';
  import Switch from '../../ui/Switch.svelte';
  import ToolSection from './ToolSection.svelte';

  let {
    items,
    loot,
    ctx,
    treeVersion,
    classSlug = '',
    onadd,
    plate = false,
  }: {
    items: readonly Item[];
    loot: LootFile;
    ctx: SearchContext;
    treeVersion: string;
    /** The character's class: a relic class's ranged slot reads "Relic". */
    classSlug?: string;
    onadd: (itemId: number) => void;
    /** The block sits on the page's one Plate. */
    plate?: boolean;
  } = $props();

  let query = $state<ItemQuery>(defaultItemQuery());

  const found = $derived(searchItems(items, query, ctx));
  /** Rows shown per page of hits: the list grows by a page, never a box that cuts a row in half. */
  const PAGE = 8;
  let shownCount = $state(PAGE);
  // A new search starts back at one page.
  $effect(() => {
    void [query.text, query.minItemLevel, query.slot, query.sourceId, query.usableOnly];
    shownCount = PAGE;
  });
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

<ToolSection title={bulkCopy.searchLabel} testid="sim-item-search" {plate}>
  <div class="flex flex-wrap items-end gap-3">
    <Field
      id="sim-search-text"
      type="search"
      label={bulkCopy.searchLabel}
      labelHidden
      placeholder={bulkCopy.searchPlaceholder}
      grow
      bind:value={query.text}
    />
    <label class="flex flex-col gap-[6px]">
      <span class={FIELD_LABEL}>{bulkCopy.searchMinItemLevel}</span>
      <input
        type="number"
        min="0"
        max="120"
        bind:value={query.minItemLevel}
        data-testid="sim-search-ilvl"
        class={[FIELD_CONTROL, FIELD_HEIGHT, FIELD_FOCUS, fieldEdge(false), 'w-24 tabular-nums']}
      />
    </label>
    <div class="flex flex-col gap-[6px]">
      <label class={FIELD_LABEL} for="sim-search-slot-field">{bulkCopy.searchSlot}</label>
      <Select
        id="sim-search-slot-field"
        class="w-[160px]"
        size="dense"
        bind:value={query.slot}
        options={[
          { value: '', label: bulkCopy.searchAnySlot },
          ...slotOptions(classSlug).map((option) => ({ value: option.slot, label: option.label })),
        ]}
        testid="sim-search-slot"
      />
    </div>
    <div class="flex flex-col gap-[6px]">
      <label class={FIELD_LABEL} for="sim-search-source-field">{bulkCopy.searchSource}</label>
      <Select
        id="sim-search-source-field"
        class="w-[240px]"
        size="dense"
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
    <span class="flex min-h-11 items-center gap-3 text-[14px]">
      <span id="sim-search-usable-label" class="text-text">{bulkCopy.searchUsableOnly}</span>
      <Switch
        checked={query.usableOnly}
        labelledby="sim-search-usable-label"
        testid="sim-search-usable"
        onchange={(next) => (query.usableOnly = next)}
      />
    </span>
  </div>

  {#if found.length === 0}
    <p class="text-muted pt-3 text-[13px]">{noResultsMessage}</p>
  {:else}
    <div class="mt-2">
      {#each found.slice(0, shownCount) as item (item.id)}
        {@const variant = legendaryVariantLabel(item.id)}
        <Row title={item.name} open density="slim" testid={`sim-search-result-${item.id}`}>
          {#snippet lead()}
            <ItemIcon icon={item.icon} quality={item.quality} build={treeVersion} size={28} />
          {/snippet}
          {#snippet titleContent()}
            <span class={rarityClassFor(item.quality)}
              >{item.name}{variant === '' ? '' : ` (${variant})`}</span
            >
          {/snippet}
          {#snippet trail()}
            <span class="text-muted text-[13px] tabular-nums">{item.item_level}</span>
            <Button tone="ghost" size="sm" testid={`sim-search-add-${item.id}`} onclick={() => onadd(item.id)}
              >{bulkCopy.searchAdd}</Button
            >
          {/snippet}
        </Row>
      {/each}
    </div>
    <div class="flex flex-wrap items-center gap-x-3 pt-2">
      {#if shownCount < found.length}
        <Button tone="ghost" size="sm" testid="sim-search-more" onclick={() => (shownCount += PAGE)}
          >{bulkCopy.searchMore(Math.min(PAGE, found.length - shownCount))}</Button
        >
      {/if}
      {#if found.length < matching}
        <p class="text-muted m-0 text-[12px]">{bulkCopy.searchTruncated(found.length, matching)}</p>
      {/if}
    </div>
  {/if}
</ToolSection>
