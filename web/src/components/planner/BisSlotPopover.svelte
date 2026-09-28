<!-- web/src/components/planner/BisSlotPopover.svelte -->
<!-- "Best in slot at <band>" for one gear slot: fetched on demand (lib/bis/hover.ts) for the
     planner character's spec, faction and level band, rather than bundled -- design step 2
     of this lane's brief. Rendered by GearPanel.svelte inside the same `relative` wrapper as
     the slot button it describes, positioned the way TalentCell.svelte's own tooltip is. -->
<script lang="ts">
  import {
    bandForLevel,
    bisPageHref,
    fetchBisFile,
    slotHoverDiff,
    type SlotHoverDiff,
  } from '../../lib/bis/hover';
  import type { BisFile, Faction } from '../../lib/bis/types';
  import { rarityClassFor } from '../../lib/planner/items';
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import { SLOT_LABELS, type Slot } from '../../lib/planner/types';
  import { specDisplayName } from '../../lib/sim/spec-label';

  let { store, slot, spec, id }: { store: PlannerStore; slot: Slot; spec: string; id: string } = $props();

  const faction = $derived<Faction>(store.raceRow?.faction === 'horde' ? 'horde' : 'alliance');
  const band = $derived(bandForLevel(store.level));
  const specName = $derived(specDisplayName(spec));

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

  const diff = $derived<SlotHoverDiff | undefined>(
    file ? slotHoverDiff(file, band, faction, slot) : undefined,
  );

  function rarityFor(itemId: number): string {
    return rarityClassFor(store.itemIndex.get(itemId)?.quality ?? 1);
  }
</script>

<div
  {id}
  role="tooltip"
  data-testid={`bis-hover-${slot}`}
  class="border-line bg-raised rounded-panel absolute top-full z-30 mt-2 flex w-[240px] flex-col gap-2 border p-3 shadow-[0_12px_30px_rgba(0,0,0,.45)]"
>
  <span class="label text-muted">Best in slot at level {band}</span>

  {#if spec === ''}
    <p class="text-muted text-[13px]">Spend a talent point to see a spec's best-in-slot picks.</p>
  {:else if file === undefined}
    <p class="text-muted text-[13px]">Loading…</p>
  {:else if file === null}
    <p class="text-muted text-[13px]" data-testid="bis-hover-empty">No leveling list for {specName} yet.</p>
  {:else if diff === undefined || diff.pick === undefined}
    <p class="text-muted text-[13px]" data-testid="bis-hover-empty">
      No known source for {SLOT_LABELS[slot]} at level {band} yet.
    </p>
  {:else}
    <div class="flex flex-col gap-1">
      <span class={`text-[13px] font-semibold ${rarityFor(diff.pick.item_id)}`}>{diff.pick.item_name}</span>
      <span
        class="rounded-control border-line inline-flex w-fit items-center border px-2 py-0.5 text-[11px] tracking-[0.04em] uppercase"
      >
        {diff.pick.source}
      </span>
      {#if diff.isNewAtBand}
        <span class="text-gold text-[11px] font-semibold" data-testid="bis-hover-new">New at {band}</span>
      {/if}
      {#if diff.previous?.pick && diff.previous.pick.item_id !== diff.pick.item_id}
        <span class="text-muted text-[12px]" data-testid="bis-hover-was">
          Was: {diff.previous.pick.item_name}
        </span>
      {/if}
    </div>
    <a class="text-[12px] underline" href={bisPageHref(spec, faction, band)}>
      See the full level {band} list
    </a>
  {/if}
</div>
