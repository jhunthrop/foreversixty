<!-- web/src/components/report/GearList.svelte -->
<!-- What a player wore at pull, named. The combatant row carries item ids in the log's
     slot order; the planner already publishes each class's item file, so the names and
     rarities come from there, fetched once per class when the list is first opened. An
     id the build does not know (a retail log, an item the data lacks) is shown by number
     rather than dropped: the row still says which slot it was and what level it had. -->
<script lang="ts">
  import { loadItems } from '../../lib/planner/load';
  import { rarityClassFor } from '../../lib/planner/items';
  import type { Item } from '../../lib/planner/types';
  import { LOG_GEAR_SLOTS, classSlugOf } from '../../lib/report/planner-link';
  import type { GearItem } from '../../lib/report/types';

  let { gear, className, dataBuild }: { gear: GearItem[]; className: string | undefined; dataBuild: string } =
    $props();

  let items = $state<Map<number, Item> | null>(null);
  let failed = $state(false);

  $effect(() => {
    const slug = classSlugOf(className);
    if (slug === null) {
      failed = true;
      return;
    }
    void loadItems(dataBuild, slug)
      .then((file) => {
        items = new Map(file.items.map((item) => [item.id, item]));
      })
      .catch(() => {
        failed = true;
      });
  });

  const worn = $derived(
    gear
      .map((piece, index) => ({ piece, slot: LOG_GEAR_SLOTS[index] ?? `slot ${index + 1}` }))
      .filter(({ piece }) => piece.ID > 0),
  );
</script>

<ul class="grid grid-cols-1 gap-x-4 text-[13px] md:grid-cols-2" data-testid="gear-list">
  {#each worn as { piece, slot } (`${slot}-${piece.ID}`)}
    {@const item = items?.get(piece.ID)}
    <li class="border-line-soft flex min-h-8 items-center gap-2 border-b py-1">
      <span class="text-muted label w-[72px] shrink-0">{slot}</span>
      {#if item}
        <span class="truncate font-semibold {rarityClassFor(item.quality)}">{item.name}</span>
      {:else}
        <span
          class="text-muted truncate"
          title={failed || items !== null ? 'Not in this build’s item data' : 'Loading'}>Item {piece.ID}</span
        >
      {/if}
      <span class="text-muted tabular ml-auto shrink-0 font-mono text-[12px]">{piece.ItemLevel}</span>
      {#if piece.Enchants && piece.Enchants.some((id) => id > 0)}
        <span class="text-muted shrink-0 text-[11px]" title="Enchanted">ench</span>
      {/if}
    </li>
  {/each}
</ul>
