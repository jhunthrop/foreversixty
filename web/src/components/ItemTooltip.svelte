<!-- web/src/components/ItemTooltip.svelte -->
<!-- The panel: icon, quality-coloured name (site body face, never the site's display face --
     tooltip-polish brief item 2), item level on its own muted line, slot/type, armor, weapon
     line, green stats, "Equip:"-style effect text, set, required level, then the source block
     behind a hairline -- the client's own tooltip layout and line order, tenet 2's "the real
     item" made visible. Pure presentation over an already-built ItemTooltipModel
     (lib/items/tooltip.ts); it fetches nothing and owns no open/close state -- ItemHover.svelte
     lazy-loads this component and mounts it only while open, so its weight never rides on a
     page's base bundle.

     The name is never truncated (tooltip-polish brief item 1): it wraps instead, inside a
     panel that clamps between 280px and 360px wide (full width minus 32px on a phone). The
     panel's own vertical and horizontal placement -- flip above the anchor, clamp and scroll
     internally when neither side fits, shift left off the viewport's right edge -- is
     `lib/items/tooltip-position.ts`'s job; this component only reads the CSS variables that
     positions it. -->
<script lang="ts">
  import type { ItemTooltipModel } from '../lib/items/tooltip';
  import { bisCopy } from '../lib/bis/copy';
  import { rarityClassFor } from '../lib/planner/items';
  import { dataUrl } from '../lib/planner/load';

  let { model, build, id }: { model: ItemTooltipModel; build: string; id: string } = $props();

  // The item pipeline's own placeholder for an id it has no real icon for is the client's
  // red-bordered "?" texture; the letter mark below is the tooltip's fallback for it, the
  // same way the BiS rows never draw it (tenet: never a red "?").
  let iconBroken = $state(model.icon === 'inv_misc_questionmark');
</script>

<div
  {id}
  role="tooltip"
  data-testid="item-tooltip"
  class="border-line bg-raised rounded-panel absolute z-30 flex w-[clamp(280px,calc(100vw_-_32px),360px)] flex-col gap-2 border p-3 text-[13px] shadow-[0_12px_30px_rgba(0,0,0,.45)]"
  style:left="var(--item-tooltip-shift, 0px)"
  style:top="var(--item-tooltip-top, 100%)"
  style:bottom="var(--item-tooltip-bottom, auto)"
  style:margin-top="var(--item-tooltip-mt, 8px)"
  style:margin-bottom="var(--item-tooltip-mb, 0px)"
  style:max-height="var(--item-tooltip-max-h, none)"
  style:overflow-y="var(--item-tooltip-overflow-y, visible)"
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
    <div class="flex min-w-0 flex-1 flex-col gap-0.5">
      <span class={`text-[15px] font-semibold ${rarityClassFor(model.quality)}`}>
        {model.name}
      </span>
      <span class="tabular text-muted font-mono text-[11px]">Item Level {model.itemLevel}</span>
      {#if model.unique}
        <span class="text-muted text-[11px]">Unique</span>
      {/if}
      {#if model.clientUnconfirmed}
        <span class="text-muted text-[11px]">{bisCopy.clientUnconfirmedTitle}</span>
      {/if}
    </div>
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
        <!-- The client paints a negative stat red. -->
        <span class:text-death={line.startsWith('-')}>{line}</span>
      {/each}
    </div>
  {/if}

  {#if model.effectText}
    <p class="text-rarity-uncommon text-[12px] leading-snug">{model.effectText}</p>
  {/if}

  {#if model.setName}
    <span class="text-rarity-uncommon text-[12px] font-semibold">{model.setName}</span>
  {/if}

  {#if model.requiredLevel > 0}
    <span class="text-muted text-[11px]">Requires Level {model.requiredLevel}</span>
  {/if}

  {#if model.sourceLines.length > 0}
    <div class="border-line-soft text-muted flex flex-col gap-0.5 border-t pt-2 text-[11px]">
      {#each model.sourceLines as line (line)}
        <span>{line}</span>
      {/each}
    </div>
  {/if}
</div>
