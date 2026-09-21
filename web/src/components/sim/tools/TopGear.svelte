<!-- web/src/components/sim/tools/TopGear.svelte -->
<!-- /sim/gear and /sim/talents are the same island: Top Gear with its gear locked and only
     the loadouts as candidates (design 3.4). `store.tool` is the whole difference, and it
     is read once here rather than branched on in every child.

     This is the composition every later Top Gear task adds a section to:
       - Task 13's ItemSearch and ConsumableCandidates sit here, between the slot grid and
         the combination count. Both only make sense in gear mode (search adds a gear
         candidate; a consumable list only multiplies `gear` mode's product, contract
         10.1 A5), so both live inside the same `gearMode` guard as the slot grid.
       - Task 14's TalentCandidates sits outside the gearMode guard (talent compare is
         loadouts only, contract 10.1's talents mode has no gear candidates) and its
         NamedSets sits inside it, beside ItemSearch and ConsumableCandidates.
       - Task 15's BulkRunBar mounts here (it mounts SettingsBar and RequestDrawer itself,
         so this file needs no import of either), after the candidate sections, owning the
         combination count.
       - Task 16 inserts ComboResults here, after the run bar.
     Each later task's section is its own child component; this file stays a thin
     composition that delegates, never growing past what a single screen's layout needs. -->
<script lang="ts">
  import type { Me } from '../../../lib/account/api';
  import type { Slot } from '../../../lib/planner/types';
  import type { BulkResult } from '../../../lib/sim/bulk-types';
  import type { BulkStore } from '../../../lib/sim/bulk-store.svelte';
  import { SIM_LEVEL } from '../../../lib/sim/character';
  import BulkRunBar from './BulkRunBar.svelte';
  import ComboResults from './ComboResults.svelte';
  import ConsumableCandidates from './ConsumableCandidates.svelte';
  import ItemSearch from './ItemSearch.svelte';
  import NamedSets from './NamedSets.svelte';
  import SlotGrid from './SlotGrid.svelte';
  import TalentCandidates from './TalentCandidates.svelte';

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
      ctx={{ level: SIM_LEVEL, sourcesByItem: store.sourceIndex, known: store.knownItems }}
      treeVersion={store.character?.tree_version ?? ''}
      onadd={(itemId) => store.addSearchItem(itemId)}
    />
    <ConsumableCandidates
      picked={store.consumableIds}
      buffs={store.simBuffs}
      treeVersion={store.character?.tree_version ?? ''}
      spec={store.character?.spec ?? ''}
      ontoggle={(id) => store.toggleConsumable(id)}
    />
    {#if store.character !== null}
      <NamedSets
        character={store.character}
        sets={store.namedSets}
        onadd={(set) => store.addNamedSet(set)}
        onremove={(name) => store.removeNamedSet(name)}
      />
    {/if}
  {/if}

  {#if store.character !== null}
    <TalentCandidates
      character={store.character}
      picked={store.loadouts}
      ontoggle={(loadout, on) => (on ? store.addLoadout(loadout) : store.removeLoadout(loadout.name))}
      onpendingcustom={(pending) => store.setPendingCustomBuild(pending)}
    />
  {/if}

  <BulkRunBar {store} />

  {#if store.result !== null}
    <!-- Fix round 1, finding 1: `runBulkAndSettle` never nulls `result` before a finished
         run replaces it, so `store.result` goes straight from result A to result B and
         this `{#if}` alone never unmounts ComboResults -- its local save/copy/filter state
         (saveOpen, savedUrl, addonCopied, keepSet, ...) would otherwise survive into the
         next run's own display. `{#key store.result}` remounts on every genuinely new
         result (bulk-store-request.ts's `runBulkAndSettle` always assigns a freshly built
         object, never mutates the old one in place), which resets all of it for free. -->
    {#key store.result}
      <ComboResults
        result={store.result as BulkResult}
        items={store.items}
        sets={store.sets}
        treeVersion={store.character?.tree_version ?? ''}
        partial={store.result.aborted === true}
        onsave={(title) => store.save(title)}
      />
    {/key}
  {/if}
</div>
