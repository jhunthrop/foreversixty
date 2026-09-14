<!-- web/src/components/planner/OrderStrip.svelte -->
<!-- The order the points were spent in, which is the part of a build the site keeps that a
     talent calculator does not. One horizontal row of icons with the level under each; it
     scrolls sideways on phone rather than wrapping, and collapses to its header. -->
<script lang="ts">
  import { levelForIndex } from '../../lib/planner/derive';
  import { dataUrl } from '../../lib/planner/load';
  import { SECONDARY_BUTTON } from '../../lib/planner/styles';
  import type { PlannerStore } from '../../lib/planner/store.svelte';

  let { store }: { store: PlannerStore } = $props();

  let expanded = $state(true);

  const points = $derived(
    store.order.map((id, i) => ({
      key: `${i}-${id}`,
      level: levelForIndex(i),
      talent: store.talentIndex?.byId.get(id) ?? null,
    })),
  );
</script>

<section
  class="border-line bg-raised rounded-panel mx-[18px] flex flex-col gap-3 border p-4 md:mx-0"
  data-testid="order-strip"
>
  <header class="flex items-center justify-between gap-3">
    <h2 class="section-title text-[15px]">Point order</h2>
    <div class="flex items-center gap-3">
      <span class="tabular text-muted font-mono text-[13px]">{store.order.length}</span>
      <button
        type="button"
        class="{SECONDARY_BUTTON} border-line-warm text-text px-3"
        aria-expanded={expanded}
        onclick={() => (expanded = !expanded)}
      >
        {expanded ? 'Hide point order' : 'Show point order'}
      </button>
    </div>
  </header>

  {#if expanded}
    <!-- One fixed-height row, whatever it holds. A 36px icon, the 4px gap and the 14px level
         line under it come to 54px; the rest is room for a classic horizontal scrollbar on the
         platforms that still reserve one. Points arrive one at a time as a build is planned, so
         a row that grew from nothing to its first icon would push the page down on the first
         click; this keeps the height reserved from first paint, empty message and all. -->
    <div class="flex h-18 items-center">
      {#if points.length === 0}
        <p class="text-muted text-[13px]">No points spent yet.</p>
      {:else}
        <ol class="-mx-1 flex h-full w-full flex-nowrap items-center gap-2 overflow-x-auto px-1">
          {#each points as point (point.key)}
            <!-- `relative` so the absolutely-positioned sr-only name below resolves against
                 this cell. Without it the name's containing block sits outside the scrolling
                 list, which leaves it unclipped and widens the whole page once the order is
                 long enough to scroll. -->
            <li class="relative flex w-11 shrink-0 flex-col items-center gap-1">
              {#if point.talent}
                <!-- The name carries the alt text of this cell, so the icon itself is
                     decorative: a screen reader reads "Improved Heroic Strike 10". -->
                <img
                  src={dataUrl(store.treeVersion, `icons/${point.talent.icon}.webp`)}
                  alt=""
                  width="36"
                  height="36"
                  loading="lazy"
                  decoding="async"
                  class="rounded-control border-line h-9 w-9 border object-cover"
                />
                <span class="sr-only">{point.talent.name}</span>
              {/if}
              <span class="tabular text-muted font-mono text-[11px] leading-[14px]">
                {point.level}
              </span>
            </li>
          {/each}
        </ol>
      {/if}
    </div>
  {/if}
</section>
