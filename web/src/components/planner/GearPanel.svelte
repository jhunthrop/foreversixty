<!-- web/src/components/planner/GearPanel.svelte -->
<!-- The 17-slot grid, the summed stats, and the active set bonuses. Two columns of slots on
     phone, four from md up; every slot button clears 44px. -->
<script lang="ts">
  import { rarityClassFor } from '../../lib/planner/items';
  import { dataUrl } from '../../lib/planner/load';
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import { SLOTS, SLOT_LABELS, STAT_LABELS, type Slot, type StatKey } from '../../lib/planner/types';
  import ItemPicker from './ItemPicker.svelte';

  let { store }: { store: PlannerStore } = $props();

  let openSlot = $state<Slot | null>(null);

  const totals = $derived(
    (Object.entries(store.statTotals) as [StatKey, number][]).sort(([a], [b]) =>
      STAT_LABELS[a].localeCompare(STAT_LABELS[b]),
    ),
  );
</script>

<section
  class="border-line bg-raised rounded-panel mx-[18px] flex flex-col gap-4 border p-4 md:mx-0"
  data-testid="gear-panel"
>
  <h2 class="section-title text-[15px]">Gear</h2>

  <div class="grid grid-cols-2 gap-2 md:grid-cols-4">
    {#each SLOTS as slot (slot)}
      {@const equippedId = store.gear[slot]}
      {@const item = equippedId === undefined ? undefined : store.itemIndex.get(equippedId)}
      <button
        type="button"
        class="border-line rounded-control bg-card-top flex min-h-11 items-center gap-2 border px-2 py-1 text-left"
        data-testid={`slot-${slot}`}
        aria-label={item ? `${SLOT_LABELS[slot]}: ${item.name}` : `${SLOT_LABELS[slot]}: empty`}
        disabled={store.readOnly}
        onclick={() => (openSlot = openSlot === slot ? null : slot)}
      >
        {#if item}
          <!-- The aria-label above names the slot and the item, so the icon is decorative. -->
          <img
            src={dataUrl(store.treeVersion, `icons/${item.icon}.webp`)}
            alt=""
            width="28"
            height="28"
            loading="lazy"
            decoding="async"
            class="rounded-control border-line h-7 w-7 border object-cover"
          />
        {:else}
          <span class="rounded-control border-line-soft h-7 w-7 border" aria-hidden="true"></span>
        {/if}
        <span class="flex min-w-0 flex-col">
          <span class="label text-muted">{SLOT_LABELS[slot]}</span>
          <span
            class={`truncate text-[13px] font-semibold ${item ? rarityClassFor(item.quality) : 'text-muted'}`}
          >
            {item ? item.name : 'Empty'}
          </span>
        </span>
      </button>
    {/each}
  </div>

  {#if openSlot}
    <ItemPicker {store} slot={openSlot} onclose={() => (openSlot = null)} />
  {/if}

  <div class="grid grid-cols-1 gap-4 md:grid-cols-2">
    <div class="flex flex-col gap-1" data-testid="gear-totals">
      <h3 class="label text-muted">Totals</h3>
      {#if totals.length === 0}
        <p class="text-muted text-[13px]">Nothing equipped.</p>
      {:else}
        <dl class="flex flex-col">
          {#each totals as [key, value] (key)}
            <div class="border-line-soft flex justify-between gap-3 border-b py-1 last:border-b-0">
              <dt class="text-muted text-[13px]">{STAT_LABELS[key]}</dt>
              <dd class="tabular text-strong font-mono text-[13px]">{value}</dd>
            </div>
          {/each}
        </dl>
      {/if}
    </div>

    <div class="flex flex-col gap-1" data-testid="gear-sets">
      <h3 class="label text-muted">Sets</h3>
      {#if store.activeSets.length === 0}
        <p class="text-muted text-[13px]">No set pieces equipped.</p>
      {:else}
        {#each store.activeSets as active (active.set.id)}
          <div class="border-line-soft flex flex-col gap-1 border-b py-2 last:border-b-0">
            <span class="text-strong text-[13px] font-semibold">
              {active.set.name}
              <span class="tabular text-gold font-mono">
                ({active.pieces}/{active.set.item_ids.length})
              </span>
            </span>
            {#each active.active as bonus (bonus.pieces)}
              <span class="text-muted text-[13px]">
                ({bonus.pieces}) {bonus.description}
              </span>
            {/each}
          </div>
        {/each}
      {/if}
    </div>
  </div>
</section>
