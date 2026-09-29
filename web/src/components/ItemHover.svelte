<!-- web/src/components/ItemHover.svelte -->
<!-- STUB -- lane `web-item-tooltips` owns this file's real implementation
     (`web/src/components/ItemTooltip.svelte` and `web/src/lib/items/*`); this lane
     (`bis-page-ux`) was told to add a minimal local stub with the same prop contract when
     the real one is not on its branch yet, and to note it in its report. The controller
     merges the real file over this one -- do not build on this beyond what `/bis` needs.

     Props per the brief: `{ itemId, classSlug, build, model? }`. This stub never fetches or
     reads off disk itself (`classSlug`/`build` are accepted but otherwise unused here) --
     every /bis caller already has the item resolved at build time (`lib/bis/load.ts`'s
     `itemDetails`) and passes it as `model`, so this stays a dumb renderer: an icon-less
     name pill (falling back to a bare id) coloured by quality, with a CSS-only (no JS, no
     hydration -- this page ships no island above the fold) hover/focus card showing item
     level, required level and the stats the real tooltip will one day show in full. -->
<script lang="ts">
  import { rarityClassFor } from '../lib/planner/items';
  import type { ItemHoverModel } from '../lib/bis/types';

  interface Props {
    itemId: number;
    classSlug: string;
    build: string;
    model?: ItemHoverModel;
  }

  let { itemId, model }: Props = $props();

  const label = $derived(model?.name ?? `Item #${itemId}`);
  const rarityClass = $derived(rarityClassFor(model?.quality ?? 1));
  const statEntries = $derived(Object.entries(model?.stats ?? {}));
</script>

<span class="item-hover" tabindex="0" data-testid={`item-hover-${itemId}`}>
  <span class={`font-semibold ${rarityClass}`}>{label}</span>
  {#if model}
    <span class="item-hover-card border-line bg-raised rounded-panel border p-2 text-[12px]" role="tooltip">
      <span class={`block font-semibold ${rarityClass}`}>{model.name}</span>
      <span class="text-muted block">Item level {model.itemLevel}</span>
      {#if model.requiredLevel > 0}
        <span class="text-muted block">Requires level {model.requiredLevel}</span>
      {/if}
      {#each statEntries as [stat, value] (stat)}
        <span class="block">+{value} {stat}</span>
      {/each}
    </span>
  {/if}
</span>

<style>
  .item-hover {
    position: relative;
    display: inline-block;
    cursor: default;
  }
  .item-hover-card {
    display: none;
    position: absolute;
    left: 0;
    top: 100%;
    z-index: 30;
    margin-top: 0.25rem;
    width: 14rem;
    text-align: left;
  }
  .item-hover:hover .item-hover-card,
  .item-hover:focus-within .item-hover-card,
  .item-hover:focus .item-hover-card {
    display: block;
  }
</style>
