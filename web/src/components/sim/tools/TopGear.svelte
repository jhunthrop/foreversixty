<!-- web/src/components/sim/tools/TopGear.svelte -->
<!-- /sim/gear and /sim/talents are the same island: Top Gear with its gear locked and only
     the loadouts as candidates (design 3.4). `store.tool` is the whole difference, and it
     is read once here rather than branched on in every child.

     This is the composition every later Top Gear task adds a section to:
       - Task 13's ItemSearch and ConsumableCandidates sit here, between the slot grid and
         the combination count. Both only make sense in gear mode (search adds a gear
         candidate; a consumable list only multiplies `gear` mode's product, contract
         10.1 A5), so both live inside the same `gearMode` guard as the slot grid.
       - Task 14 inserts TalentCandidates and NamedSets here, in the same place.
       - Task 15 inserts BulkRunBar (which itself mounts SettingsPanel and RequestDrawer)
         here, after the candidate sections and before the combination count.
       - Task 16 inserts ComboResults here, after the run bar.
     Each later task's section is its own child component; this file stays a thin
     composition that delegates, never growing past what a single screen's layout needs. -->
<script lang="ts">
  import type { Me } from '../../../lib/account/api';
  import type { Slot } from '../../../lib/planner/types';
  import type { BulkStore } from '../../../lib/sim/bulk-store.svelte';
  import { SIM_LEVEL } from '../../../lib/sim/character';
  import { bulkCopy } from '../../../lib/sim/copy';
  import ConsumableCandidates from './ConsumableCandidates.svelte';
  import ItemSearch from './ItemSearch.svelte';
  import SlotGrid from './SlotGrid.svelte';

  // `me` is part of the pinned prop shape ToolsView.svelte passes to every tool view, but
  // the slot grid itself has no use for it -- Task 15's BulkRunBar is where a signed-out
  // premium prompt would read it. Destructuring only `store` keeps the type annotation
  // (and therefore the contract) intact without an unused local.
  let { store }: { store: BulkStore; me: Me | null } = $props();

  const gearMode = $derived(store.tool === 'gear');
</script>

<div class="flex flex-col gap-[22px] md:gap-8" data-testid="sim-top-gear">
  {#if gearMode}
    <SlotGrid
      rows={store.rows}
      locked={store.locked}
      enchants={store.enchants}
      suffixes={store.suffixes}
      classSlug={store.character?.class_slug ?? ''}
      treeVersion={store.character?.tree_version ?? ''}
      ontoggle={(key) => store.toggleRow(key)}
      oncopy={(key, patch) => store.copyAndModify(key, patch)}
      onlock={(slot: Slot) => store.toggleLock(slot)}
    />
    <ItemSearch
      items={[...store.items.values()]}
      loot={store.loot}
      ctx={{ level: SIM_LEVEL, sourcesByItem: store.sourceIndex }}
      treeVersion={store.character?.tree_version ?? ''}
      onadd={(itemId) => store.addSearchItem(itemId)}
    />
    <ConsumableCandidates
      picked={store.consumableIds}
      buffs={store.simBuffs}
      treeVersion={store.character?.tree_version ?? ''}
      ontoggle={(id) => store.toggleConsumable(id)}
    />
  {/if}

  <p class="text-muted px-[18px] text-[13px] md:px-0" data-testid="sim-combo-count">
    {store.combinations === null ? bulkCopy.combinationsCounting : bulkCopy.combinations(store.combinations)}
  </p>
</div>
