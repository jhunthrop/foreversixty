<!-- web/src/components/ItemTooltip.svelte -->
<!-- The panel: icon, quality-coloured name, item level, slot/type, armor, weapon line,
     green stats, "Equip:"-style effect text, set, source, required level -- the client's
     own tooltip layout, tenet 2's "the real item" made visible. Pure presentation over an
     already-built ItemTooltipModel (lib/items/tooltip.ts); it fetches nothing and owns no
     open/close state -- ItemHover.svelte lazy-loads this component and mounts it only while
     open, so its weight never rides on a page's base bundle. -->
<script lang="ts">
  import type { ItemTooltipModel } from '../lib/items/tooltip';
  import { rarityClassFor } from '../lib/planner/items';
  import { dataUrl } from '../lib/planner/load';

  let { model, build, id }: { model: ItemTooltipModel; build: string; id: string } = $props();

  let iconBroken = $state(false);
</script>

<div
  {id}
  role="tooltip"
  data-testid="item-tooltip"
  class="border-line bg-raised rounded-panel absolute top-full left-0 z-30 mt-2 flex w-72 flex-col gap-2 border p-3 text-[13px] shadow-[0_12px_30px_rgba(0,0,0,.45)]"
  style:left="var(--item-tooltip-shift, 0px)"
>
  <div class="flex items-start gap-2">
    {#if iconBroken}
      <span
        class="rounded-control border-line bg-card-top text-muted flex h-10 w-10 shrink-0 items-center justify-center border text-[11px] font-bold"
        aria-hidden="true"
      >
        {model.name.slice(0, 2)}
      </span>
    {:else}
      <img
        src={dataUrl(build, `icons/${model.icon}.webp`)}
        alt=""
        width="40"
        height="40"
        loading="lazy"
        decoding="async"
        class="rounded-control border-line h-10 w-10 shrink-0 border object-cover"
        onerror={() => (iconBroken = true)}
        onload={(event) => {
          if (event.currentTarget.naturalWidth <= 1) iconBroken = true;
        }}
      />
    {/if}
    <div class="flex min-w-0 flex-col">
      <span class={`font-display truncate text-[14px] font-bold ${rarityClassFor(model.quality)}`}>
        {model.name}
      </span>
      {#if model.unique}
        <span class="text-muted text-[11px]">Unique</span>
      {/if}
    </div>
    <span class="tabular text-muted ml-auto shrink-0 font-mono text-[11px]">iLvl {model.itemLevel}</span>
  </div>

  <div class="text-muted flex flex-wrap justify-between gap-x-2 text-[12px]">
    <span>{model.slotLabel}</span>
    {#if model.typeLabel}<span>{model.typeLabel}</span>{/if}
  </div>

  {#if model.armor !== null}
    <span class="text-[12px]">{model.armor} Armor</span>
  {/if}

  {#if model.weapon}
    <div class="flex flex-col text-[12px]">
      <span>{model.weapon.damageRange}</span>
      <span>{model.weapon.speed}</span>
      <span class="text-muted">{model.weapon.dps}</span>
    </div>
  {/if}

  {#if model.stats.length > 0}
    <div class="text-rarity-uncommon flex flex-col text-[12px]">
      {#each model.stats as line (line)}
        <span>{line}</span>
      {/each}
    </div>
  {/if}

  {#if model.effectText}
    <p class="text-rarity-uncommon text-[12px] leading-snug">{model.effectText}</p>
  {/if}

  {#if model.setName}
    <span class="text-rarity-uncommon text-[12px] font-semibold">{model.setName}</span>
  {/if}

  {#if model.sourceLines.length > 0}
    <div class="border-line-soft text-muted flex flex-col gap-0.5 border-t pt-2 text-[11px]">
      {#each model.sourceLines as line (line)}
        <span>{line}</span>
      {/each}
    </div>
  {/if}

  {#if model.requiredLevel > 0}
    <span class="text-muted text-[11px]">Requires Level {model.requiredLevel}</span>
  {/if}
</div>
