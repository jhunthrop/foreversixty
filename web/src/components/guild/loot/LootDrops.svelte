<!-- web/src/components/guild/loot/LootDrops.svelte -->
<!-- The loot council's one Plate (web layout spec 2026-11-05 §7.2): the chosen night's drops as
     ItemRows (icon in its quality frame, quality-coloured name, what it is and who can use it or
     who got it, an Open or Awarded Tag). A row is one button: it opens the drop's candidates (or,
     for a member, their own MS/OS call). On the current night a picker under the Plate opens any
     drop of the loot table that is not on the night yet. -->
<script lang="ts">
  import type { GuildLootItem } from '../../../lib/guild/api';
  import type { DropRow } from '../../../lib/guild/loot-council-view';
  import ItemRow from '../../ui/ItemRow.svelte';
  import Plate from '../../ui/Plate.svelte';
  import Select from '../../ui/Select.svelte';
  import Tag from '../../ui/Tag.svelte';

  let {
    heading,
    meta,
    rows,
    openable,
    selected,
    build,
    emptyText,
    onSelect,
    onOpen,
  }: {
    heading: string;
    meta: string;
    rows: readonly DropRow[];
    /** The loot table's drops the current night can still open; empty on a past night. */
    openable: readonly GuildLootItem[];
    selected: number | null;
    build: string;
    emptyText: string;
    onSelect: (row: DropRow) => void;
    onOpen: (itemId: number) => void;
  } = $props();

  const pickOptions = $derived([
    { value: '', label: 'Pick a drop' },
    ...openable.map((item) => ({ value: String(item.item_id), label: item.name })),
  ]);
</script>

<section class="flex flex-col gap-3" aria-labelledby="guild-loot-drops-head" data-testid="guild-loot-drops">
  <div class="flex flex-wrap items-baseline gap-x-3 gap-y-1">
    <h2 id="guild-loot-drops-head" class="text-strong text-[18px] leading-6 font-bold">{heading}</h2>
    <span class="text-small text-muted" data-testid="guild-loot-night-summary">{meta}</span>
  </div>
  <Plate>
    {#each rows as row (row.itemId)}
      <ItemRow
        item={{
          id: row.itemId,
          name: row.name,
          quality: row.quality,
          icon: row.kind === 'gear' ? row.icon : undefined,
        }}
        {build}
        source={row.sub}
        selected={selected === row.itemId}
        onselect={() => onSelect(row)}
        testid={`guild-loot-drop-${row.itemId}`}
      >
        {#snippet trail()}
          {#if row.state === 'awarded'}<Tag tone="success">Awarded</Tag>{:else}<Tag>Open</Tag>{/if}
        {/snippet}
      </ItemRow>
    {:else}
      <p class="text-muted px-4 py-5 text-[14px]" data-testid="guild-loot-empty">{emptyText}</p>
    {/each}
  </Plate>
  {#if openable.length > 0}
    <div class="flex max-w-[360px] flex-col gap-[6px]" data-testid="guild-loot-open-drop">
      <label for="guild-loot-open-drop" class="text-label tracking-label text-muted font-bold uppercase"
        >Open a drop from the loot table</label
      >
      <Select
        id="guild-loot-open-drop"
        size="compact"
        options={pickOptions}
        value=""
        onchange={(event) => {
          const value = (event.currentTarget as HTMLSelectElement).value;
          if (value !== '') onOpen(Number(value));
        }}
      />
    </div>
  {/if}
</section>
