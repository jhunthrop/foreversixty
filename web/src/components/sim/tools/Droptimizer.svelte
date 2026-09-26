<!-- web/src/components/sim/tools/Droptimizer.svelte -->
<!-- /sim/drops. It is the same store, the same run bar and the same stage loop as Top Gear;
     the source picker replaces the slot grid and the results group by boss instead of
     ranking one flat list.
     "Pin into Top Gear" leaves this page with the item in the query string, because the two
     pages are separate islands and a pin has to survive the navigation. It also carries the
     drop's own origin and source name, so the row Top Gear's ToolsView adds on arrival keeps
     "Ragnaros" as its provenance instead of falling back to a plain search hit. -->
<script lang="ts">
  import { SvelteURLSearchParams } from 'svelte/reactivity';
  import type { Me } from '../../../lib/account/api';
  import type { BulkStore } from '../../../lib/sim/bulk-store.svelte';
  import type { BulkResult } from '../../../lib/sim/bulk-types';
  import { topUpgradeOf } from '../../../lib/sim/combos';
  import { bulkCopy } from '../../../lib/sim/copy';
  import { pickedWithNothingTried } from '../../../lib/sim/drop-picks';
  import { writeLastUpgrade } from '../../../lib/sim/last-upgrade';
  import BulkRunBar from './BulkRunBar.svelte';
  import DropResults from './DropResults.svelte';
  import SourcePicker from './SourcePicker.svelte';

  // `me` is part of the pinned prop shape ToolsView.svelte passes to every tool view (see
  // TopGear.svelte's identical comment); this page has no use for it itself.
  let { store }: { store: BulkStore; me: Me | null } = $props();

  /**
   * Top Gear with this item already ticked, and its drop provenance carried along so the
   * row that lands there is `drop:<source-id>`-sourced rather than a plain search hit
   * (`addRow`'s own provenance merge, candidates.ts, prefers the richer of the two when the
   * item was already a bag/equipped row). `?pin=`/`?pinOrigin=`/`?pinName=` are read by
   * ToolsView's bootstrap the same way `?source=`/`?ref=` are.
   */
  function pin(itemId: number, origin: string, sourceName: string): void {
    const params = new SvelteURLSearchParams(window.location.search);
    params.set('pin', String(itemId));
    if (origin !== '') params.set('pinOrigin', origin);
    else params.delete('pinOrigin');
    if (sourceName !== '') params.set('pinName', sourceName);
    else params.delete('pinName');
    window.location.href = `/sim/gear?${params.toString()}`;
  }

  /**
   * Picks with no trace in the displayed result (task 3b; final whole-branch review,
   * Important 3) -- computed here, where `store.loot` lives, and passed down as plain
   * `{ key, name }` rows so `DropResults` keeps its own no-`loot`-prop rule intact.
   *
   * `store.submittedDropPicks`, not the live `store.pickedBosses`: this must be the pick set
   * that produced `store.result`, or a source ticked after the run would read as "untried"
   * for a result it was never part of.
   */
  const untried = $derived(
    store.result === null
      ? []
      : pickedWithNothingTried(store.submittedDropPicks, store.loot, (store.result as BulkResult).combos),
  );

  // Spec 2026-09-25 §6: the after-sim sentence on plain /sim reads this back later, with no
  // fetch of its own. Only a genuine upgrade (`topUpgradeOf`'s own `> 0` rule, the same one
  // DropResults.svelte's `isUpgrade` uses) is worth recording -- a run with nothing better
  // than what is equipped writes nothing, leaving whatever was last recorded (or nothing)
  // in place rather than overwriting a real upgrade with "no upgrade" from an unrelated
  // later run.
  $effect(() => {
    if (store.result === null) return;
    const upgrade = topUpgradeOf(store.result as BulkResult);
    if (upgrade === null) return;
    writeLastUpgrade({ ...upgrade, savedAt: new Date().toISOString() });
  });
</script>

<div class="flex flex-col gap-[22px] md:gap-8" data-testid="sim-droptimizer">
  {#if store.character !== null}
    <SourcePicker
      loot={store.loot}
      phases={store.phases}
      shownKinds={store.shownKinds}
      showUpcoming={store.showUpcoming}
      picked={store.pickedBosses}
      professions={store.character.professions}
      items={store.items}
      known={store.knownItems}
      now={new Date()}
      ontogglekind={(kind) => store.toggleKind(kind)}
      ontoggleupcoming={(value) => store.setShowUpcoming(value)}
      ontoggle={(sourceId, bossId) => store.toggleSource(sourceId, bossId)}
    />
  {/if}

  <BulkRunBar {store} />

  {#if store.result !== null}
    <DropResults
      result={store.result as BulkResult}
      items={store.items}
      treeVersion={store.character?.tree_version ?? ''}
      {untried}
      onpin={pin}
    />
  {:else if store.pickedBosses.length === 0}
    <p class="text-muted px-[18px] text-[13px] md:px-0">{bulkCopy.dropsNothing}</p>
  {/if}
</div>
