<!-- web/src/components/planner/OrderStrip.svelte -->
<!-- The order the points were spent in, which is the part of a build the site keeps that a
     talent calculator does not. One horizontal row of icons with the level under each; it
     scrolls sideways on phone rather than wrapping.

     Design loop, planner round (build review round 1, findings 3 and 4): this used to be a
     full-width panel with its own custom expand/collapse button on every breakpoint, sized
     the same as Share/Import/Gear even though its content is at most a short row of icons.
     From md up it is now a plain, always-open panel in the rail beside the trees, and it
     reserves no height at all while it is empty -- a single line says so instead. Below md
     it folds behind a native, closed-by-default <details>, the same collapse ImportBox now
     uses, so a visitor reaches Gear without scrolling past a panel most never open. -->
<script lang="ts">
  import { levelForIndex } from '../../lib/planner/derive';
  import { dataUrl } from '../../lib/planner/load';
  import { CELL_BORDER, cellState } from '../../lib/planner/styles';
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import { plannerCopy } from '../../lib/planner/copy';
  import { CARD_BODY, CARD_HEAD, cardClass } from '../../lib/ui/card';
  import Card from '../ui/Card.svelte';
  import IconTile from '../ui/IconTile.svelte';

  let {
    store,
    phone = false,
    class: className = '',
  }: {
    store: PlannerStore;
    /** Below md, closed by default behind a native <details> rather than always open. */
    phone?: boolean;
    /** Merged onto the root element, so the caller's own responsive grid can place this
     *  panel without either side knowing about the other's layout. */
    class?: string;
  } = $props();

  const points = $derived(
    store.order.map((id, i) => {
      const talent = store.talentIndex?.byId.get(id) ?? null;
      const rank = store.order.slice(0, i + 1).filter((other) => other === id).length;
      return {
        key: `${i}-${id}`,
        level: levelForIndex(i),
        talent,
        // A point in the strip is always spent, so it is filled or maxed; the
        // third argument only decides between available and locked.
        state: talent ? cellState(rank, talent.max_rank, true) : 'locked',
      };
    }),
  );
</script>

{#snippet list()}
  {#if points.length === 0}
    <p class="text-muted text-[13px]">No points spent yet.</p>
  {:else}
    <!-- One fixed-height row, whatever it holds. A 36px icon, the 4px gap and the 14px level
         line under it come to 54px; the rest is room for a classic horizontal scrollbar on
         the platforms that still reserve one. Points arrive one at a time as a build is
         planned, so a row that grew from nothing to its first icon would push the page down
         on the first click -- this keeps that height reserved once there is a first icon to
         reserve it for, but not before (the empty branch above has no height of its own). -->
    <div class="flex h-18 items-center">
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
                class={`rounded-control h-9 w-9 border object-cover ${CELL_BORDER[point.state]}`}
              />
              <span class="sr-only">{point.talent.name}</span>
            {/if}
            <span class="text-muted numeric text-[11px] leading-[14px]">
              {point.level}
            </span>
          </li>
        {/each}
      </ol>
    </div>
  {/if}
{/snippet}

{#if phone}
  <!-- Closed by default on a phone: the card's own surface and compact head recipe, as a
       native <details> (a `Card` is a section, so it cannot fold). -->
  <details class="{cardClass(false)} mx-[18px] {className}" data-testid="order-strip">
    <summary class="{CARD_HEAD.compact} min-h-11 cursor-pointer list-none">
      <IconTile glyph="list" size={28} />
      <span class="section-title flex-1 leading-6">{plannerCopy.pointOrderTitle}</span>
      <span class="text-muted flex items-center gap-3 text-[13px]">
        <span class="numeric">{plannerCopy.pointOrderMeta(store.order.length)}</span>
        <span class="when-closed" aria-hidden="true">Show</span>
        <span class="when-open" aria-hidden="true">Hide</span>
      </span>
    </summary>
    <div class={CARD_BODY.text}>
      {@render list()}
    </div>
  </details>
{:else}
  <Card
    glyph="list"
    compact
    title={plannerCopy.pointOrderTitle}
    meta={plannerCopy.pointOrderMeta(store.order.length)}
    testid="order-strip"
    class="mx-[18px] md:mx-0 {className}"
  >
    {@render list()}
  </Card>
{/if}
