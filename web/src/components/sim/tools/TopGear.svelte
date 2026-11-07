<!-- web/src/components/sim/tools/TopGear.svelte -->
<!-- /sim/gear and /sim/talents are the same island (mock design/mocks/ws_simtools.py): the run
     line, then the split page. Top gear: the candidate grid on the page's one Plate in the main
     column, the results above it once a run finishes, and the ways to add candidates (an item
     search, the consumables to try, whole sets, talent builds) as open sections in the rail.
     Talents: gear is locked to what is worn, so the talent builds are the Plate and there is no
     rail. `store.tool` is the whole difference, read once here. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import Page from '../../ui/Page.svelte';
  import type { Me } from '../../../lib/account/api';
  import type { Slot } from '../../../lib/planner/types';
  import type { BulkResult } from '../../../lib/sim/bulk-types';
  import type { BulkStore } from '../../../lib/sim/bulk-store.svelte';
  import { MAX_LEVEL } from '../../../lib/sim/character';
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
  let { store, lead }: { store: BulkStore; me: Me | null; lead?: Snippet } = $props();

  const gearMode = $derived(store.tool === 'gear');
</script>

<Page layout={gearMode ? 'split' : 'ledger'} testid="sim-top-gear">
  {#snippet bar()}<BulkRunBar {store} />{/snippet}
  {#if lead}{@render lead()}{/if}
  {#if store.result !== null}
    <!-- `runBulkAndSettle` replaces `result` with a fresh object on every finished run, so
         `{#key}` remounts ComboResults and resets its save, copy and filter state. -->
    {#key store.result}
      <ComboResults
        result={store.result as BulkResult}
        items={store.items}
        sets={store.sets}
        treeVersion={store.character?.tree_version ?? ''}
        partial={store.result.aborted === true}
        onsave={(title) => store.save(title)}
        talents={store.talentFile}
      />
    {/key}
  {/if}
  {#if gearMode}
    <SlotGrid
      plate
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
  {:else if store.character !== null}
    <TalentCandidates
      plate
      character={store.character}
      picked={store.loadouts}
      ontoggle={(loadout, on) => (on ? store.addLoadout(loadout) : store.removeLoadout(loadout.name))}
      onpendingcustom={(pending) => store.setPendingCustomBuild(pending)}
    />
  {/if}
  {#snippet rail()}
    <ItemSearch
      items={[...store.items.values()]}
      loot={store.loot}
      ctx={{ level: MAX_LEVEL, sourcesByItem: store.sourceIndex, known: store.knownItems }}
      treeVersion={store.character?.tree_version ?? ''}
      classSlug={store.character?.class_slug ?? ''}
      onadd={(itemId) => store.addSearchItem(itemId)}
    />
    <ConsumableCandidates
      picked={store.consumableIds}
      buffs={store.simBuffs}
      treeVersion={store.character?.tree_version ?? ''}
      spec={store.character?.spec ?? ''}
      level={store.character?.level ?? MAX_LEVEL}
      ontoggle={(id) => store.toggleConsumable(id)}
    />
    {#if store.character !== null}
      <NamedSets
        character={store.character}
        sets={store.namedSets}
        onadd={(set) => store.addNamedSet(set)}
        onremove={(name) => store.removeNamedSet(name)}
      />
      <TalentCandidates
        character={store.character}
        picked={store.loadouts}
        ontoggle={(loadout, on) => (on ? store.addLoadout(loadout) : store.removeLoadout(loadout.name))}
        onpendingcustom={(pending) => store.setPendingCustomBuild(pending)}
      />
    {/if}
  {/snippet}
</Page>
