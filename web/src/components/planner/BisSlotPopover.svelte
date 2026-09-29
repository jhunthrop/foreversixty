<!-- web/src/components/planner/BisSlotPopover.svelte -->
<!-- "Best in slot at <band>" for one gear slot: fetched on demand (lib/bis/hover.ts) for the
     planner character's spec, faction and level band, rather than bundled -- design step 2
     of this lane's brief. Rendered by GearPanel.svelte inside the same `relative` wrapper as
     the slot button it describes, positioned the way TalentCell.svelte's own tooltip is.

     Every state that is not "here is the pick" (still loading, no file for this spec, or a
     file with nothing known for this slot) shares one message paragraph and one
     `data-testid`, rather than a template per state: this ships in the planner island,
     which has its own tight gzipped budget (scripts/check-island-size.mjs), and Svelte
     compiles a template per branch. -->
<script lang="ts">
  import {
    bandForLevel,
    bisPageHref,
    fetchBisFile,
    slotHoverDiff,
    type SlotHoverDiff,
  } from '../../lib/bis/hover';
  import type { BisFile, Faction } from '../../lib/bis/types';
  import ItemHover from '../ItemHover.svelte';
  import { rarityClassFor } from '../../lib/planner/items';
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import { SLOT_LABELS, type Slot } from '../../lib/planner/types';
  import { specDisplayName } from '../../lib/sim/spec-label';

  let { store, slot, spec, id }: { store: PlannerStore; slot: Slot; spec: string; id: string } = $props();

  const faction = $derived<Faction>(store.raceRow?.faction === 'horde' ? 'horde' : 'alliance');
  const band = $derived(bandForLevel(store.level));

  // undefined: still loading (or not started). null: fetched, and there is no file for this
  // spec yet -- the honest empty state, never a guess.
  let file = $state<BisFile | null | undefined>(undefined);
  let loadedFor = $state('');

  $effect(() => {
    const build = store.treeVersion;
    if (spec === '' || build === '') {
      file = null;
      loadedFor = '';
      return;
    }
    const key = `${build}::${spec}`;
    if (loadedFor === key) return;
    loadedFor = key;
    file = undefined;
    void fetchBisFile(build, spec).then((result) => {
      if (loadedFor === key) file = result;
    });
  });

  interface View {
    message: string | null;
    diff?: SlotHoverDiff;
  }

  const view = $derived<View>(
    file === undefined
      ? { message: 'Loading…' }
      : file === null
        ? { message: `No leveling list for ${specDisplayName(spec)} yet.` }
        : (() => {
            const diff = slotHoverDiff(file, band, faction, slot);
            return diff.pick === undefined
              ? { message: `No known source for ${SLOT_LABELS[slot]} at level ${band} yet.` }
              : { message: null, diff };
          })(),
  );
</script>

<div
  {id}
  role="tooltip"
  data-testid={`bis-hover-${slot}`}
  class="border-line bg-raised rounded-panel absolute top-full z-20 mt-1 w-56 border p-3 text-[13px]"
>
  <span class="label text-muted">Best in slot at level {band}</span>

  {#if view.message !== null}
    <p class="text-muted" data-testid="bis-hover-empty">{view.message}</p>
  {:else if view.diff}
    <div class="flex flex-col gap-1">
      <ItemHover itemId={view.diff.pick.item_id} classSlug={store.classSlug} build={store.treeVersion}>
        <span
          class={`font-semibold ${rarityClassFor(store.itemIndex.get(view.diff.pick.item_id)?.quality ?? 1)}`}
        >
          {view.diff.pick.item_name}
        </span>
      </ItemHover>
      <span class="text-muted text-[11px] uppercase">{view.diff.pick.source}</span>
      {#if view.diff.isNewAtBand}
        <span class="text-gold text-[11px] font-semibold" data-testid="bis-hover-new">New at {band}</span>
      {/if}
      {#if view.diff.previous?.pick && view.diff.previous.pick.item_id !== view.diff.pick.item_id}
        <span class="text-muted text-[11px]" data-testid="bis-hover-was"
          >Was: {view.diff.previous.pick.item_name}</span
        >
      {/if}
    </div>
    <a class="underline" href={bisPageHref(spec, faction, band)}>See the full level {band} list</a>
  {/if}
</div>
