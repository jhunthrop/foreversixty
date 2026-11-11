<!-- web/src/components/guild/loot/LootDrops.svelte -->
<!-- The loot council's one Plate, in its Section (web layout spec 2026-11-05 §7.2): the chosen night's drops as
     ItemRows (icon in its quality frame, quality-coloured name, what it is and who can use it or
     who got it, an Open or Awarded Tag). A row is one button: it opens the drop's candidates (or,
     for a member, their own MS/OS call). On the current night a picker under the Plate opens any
     drop of the loot table that is not on the night yet, grouped by raid then boss; a raid the
     game client does not carry yet reads "Not in the game client yet". -->
<script lang="ts">
  import EmptyState from '../../ui/EmptyState.svelte';
  import type { GuildLootItem, GuildLootRaid } from '../../../lib/guild/api';
  import { pickOptions, type DropRow } from '../../../lib/guild/loot-council-view';
  import ItemRow from '../../ui/ItemRow.svelte';
  import Section from '../../ui/Section.svelte';
  import Select from '../../ui/Select.svelte';
  import Tag from '../../ui/Tag.svelte';

  let {
    heading,
    meta,
    rows,
    openable,
    raids,
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
    /** The first tier's raids, for the picker's raid then boss grouping. */
    raids: readonly GuildLootRaid[];
    selected: number | null;
    build: string;
    emptyText: string;
    onSelect: (row: DropRow) => void;
    onOpen: (itemId: number) => void;
  } = $props();

  const options = $derived(pickOptions(openable, raids));
</script>

<Section title={heading} panel {meta} metaTestid="guild-loot-night-summary" testid="guild-loot-drops">
  <div>
    {#each rows as row (row.itemId)}
      <ItemRow
        item={{
          id: row.itemId,
          name: row.name,
          quality: row.quality,
          icon: row.icon === '' ? undefined : row.icon,
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
      <EmptyState message={emptyText} testid="guild-loot-empty" />
    {/each}
  </div>
  {#if openable.length > 0}
    <div class="mt-3 flex max-w-[360px] flex-col gap-[6px]" data-testid="guild-loot-open-drop">
      <label for="guild-loot-open-drop" class="sr-only">Open a drop from the loot table</label>
      <Select
        id="guild-loot-open-drop"
        size="compact"
        {options}
        value=""
        onchange={(event) => {
          const value = (event.currentTarget as HTMLSelectElement).value;
          if (value !== '') onOpen(Number(value));
        }}
      />
    </div>
  {/if}
</Section>
