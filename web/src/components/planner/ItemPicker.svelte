<!-- web/src/components/planner/ItemPicker.svelte -->
<!-- The searchable list for one slot: filtered to the slot and the class, sorted by
     required level then name, coloured by rarity. -->
<script lang="ts">
  import { addonCopy } from '../../lib/addon/copy';
  import { scoreItem, type SpecWeights } from '../../lib/addon/score';
  import { itemsForSlot, rarityClassFor, searchItems } from '../../lib/planner/items';
  import { dataUrl } from '../../lib/planner/load';
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import { SECONDARY_BUTTON } from '../../lib/planner/styles';
  import { SLOT_LABELS, type Slot } from '../../lib/planner/types';

  let {
    store,
    slot,
    weights,
    onclose,
  }: { store: PlannerStore; slot: Slot; weights?: SpecWeights; onclose: () => void } = $props();

  let query = $state('');
  let sortByScore = $state(false);

  const all = $derived(itemsForSlot([...store.itemIndex.values()], slot));
  const searched = $derived(searchItems(all, query));
  const shown = $derived(
    sortByScore && weights !== undefined
      ? [...searched].sort((a, b) => scoreItem(b, weights) - scoreItem(a, weights))
      : searched,
  );
</script>

<div class="border-line bg-raised rounded-panel flex flex-col gap-3 border p-4" data-testid="item-picker">
  <header class="flex items-center justify-between gap-3">
    <h3 class="section-title text-[15px]">{SLOT_LABELS[slot]}</h3>
    <button type="button" class="{SECONDARY_BUTTON} border-line-warm text-text px-3" onclick={onclose}>
      Close
    </button>
  </header>

  <!-- Tailwind v4's preflight sets `border: 0 solid`, so the bare `border` is what makes
       border-line-warm visible at all. -->
  <input
    type="search"
    bind:value={query}
    placeholder="Filter by name"
    aria-label={`Filter ${SLOT_LABELS[slot]} items`}
    class="border-line-warm rounded-control bg-bg text-text placeholder:text-muted h-11 border px-3 text-[14px]"
  />

  {#if weights !== undefined}
    <label class="text-muted flex min-h-11 w-fit items-center gap-2 text-[12px]">
      <input type="checkbox" class="h-5 w-5" data-testid="sort-by-score" bind:checked={sortByScore} />
      {addonCopy.sortByScore}
    </label>
  {/if}

  {#if store.gear[slot] !== undefined}
    <button
      type="button"
      class="{SECONDARY_BUTTON} border-line-warm text-text w-fit px-3"
      onclick={() => {
        store.unequip(slot);
        onclose();
      }}
    >
      Clear slot
    </button>
  {/if}

  {#if shown.length === 0}
    <p class="text-muted text-[13px]">No item in this class list fits {SLOT_LABELS[slot]}.</p>
  {:else}
    <ul class="max-h-[320px] overflow-y-auto">
      {#each shown as item (item.id)}
        <!-- The rule belongs to the row, not the button: the button is its li's only child,
             so `last:` on the button would match every row and draw no rule at all. -->
        <li class="border-line-soft border-b last:border-b-0">
          <button
            type="button"
            class="flex min-h-11 w-full items-center gap-3 px-1 py-2 text-left"
            data-testid={`item-${item.id}`}
            onclick={() => {
              store.equip(slot, item.id);
              onclose();
            }}
          >
            <!-- The name beside it carries the item, so the icon is decorative. -->
            <img
              src={dataUrl(store.treeVersion, `icons/${item.icon}.webp`)}
              alt=""
              width="28"
              height="28"
              loading="lazy"
              decoding="async"
              class="rounded-control border-line h-7 w-7 border object-cover"
            />
            <span class={`flex-1 text-[14px] font-semibold ${rarityClassFor(item.quality)}`}>
              {item.name}
            </span>
            {#if weights !== undefined}
              <span class="tabular text-muted font-mono text-[12px]" data-testid="item-score">
                {scoreItem(item, weights).toFixed(1)}
              </span>
            {/if}
            <span class="tabular text-muted font-mono text-[12px]">{item.required_level}</span>
          </button>
        </li>
      {/each}
    </ul>
  {/if}
</div>
