<!-- web/src/components/ItemTooltip.svelte -->
<!-- The panel: icon, quality-coloured name (site body face, never the site's display face --
     tooltip-polish brief item 2), item level on its own muted line, slot/type, armor, weapon
     line, white stats, the worn enchant's green line, green "Equip:" sentences and effect text, set, required level, then the source block
     behind a hairline -- the client's own tooltip layout and line order, tenet 2's "the real
     item" made visible. Pure presentation over an already-built ItemTooltipModel
     (lib/items/tooltip.ts); it fetches nothing and owns no open/close state -- ItemHover.svelte
     lazy-loads this component and mounts it only while open, so its weight never rides on a
     page's base bundle.

     The frame is the Overlay primitive's tooltip kind (the in-game tooltip frame); the source block
     is its foot, under the overlay's own hairline, and the icon is ui/ItemIcon in its quality frame.

     The name is never truncated (tooltip-polish brief item 1): it wraps instead, inside a
     panel that clamps between 280px and 360px wide (full width minus 32px on a phone). The
     panel's own vertical and horizontal placement -- flip above the anchor, clamp and scroll
     internally when neither side fits, shift left off the viewport's right edge -- is
     `lib/items/tooltip-position.ts`'s job; this component only reads the CSS variables that
     positions it. `floating={false}` sets it in the flow instead (the paperdoll frame's centre shows the
     selected slot's tooltip in place), full width of its box and named as a group, not a tooltip.
     `tail` ends the body in the client's yellow, as the client ends a tooltip with what an item
     is worth to you (the paperdoll's "+0.3 DPS vs ..."). -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { ItemTooltipModel } from '../lib/items/tooltip';
  import { TOOLTIP_PANEL_PLACEMENT } from '../lib/items/tooltip-position';
  import { bisCopy } from '../lib/bis/copy';
  import { rarityClassFor } from '../lib/planner/items';
  import ItemIcon from './ui/ItemIcon.svelte';
  import Overlay from './ui/Overlay.svelte';

  let {
    model,
    build,
    id,
    enchant,
    floating = true,
    tail,
  }: {
    model: ItemTooltipModel;
    build: string;
    id?: string;
    /** False: in the flow (the paperdoll's centre), not anchored to a trigger. */
    floating?: boolean;
    /** The last lines of the body, in the client's yellow (what the item is worth). */
    tail?: Snippet;
    /** The enchant on the worn item, in the client's own words ("Crusader"): its green line. */
    enchant?: string;
  } = $props();

  // The client prints a white stat ("+18 Strength") in white and the "Equip:" sentences in green,
  // with the enchant's green line between them.
  const EQUIP_PREFIX = 'Equip:';
  const whiteStats = $derived(model.stats.filter((line) => !line.startsWith(EQUIP_PREFIX)));
  const equipStats = $derived(model.stats.filter((line) => line.startsWith(EQUIP_PREFIX)));

  // The item pipeline's own placeholder for an id it has no real icon for is the client's
  // red-bordered "?" texture; ItemIcon's empty socket stands in for it, the same way the BiS rows
  // never draw it (tenet: never a red "?").
  const icon = $derived(model.icon === 'inv_misc_questionmark' || model.icon === '' ? undefined : model.icon);
</script>

<Overlay
  kind="tooltip"
  {id}
  testid="item-tooltip"
  role={floating ? undefined : 'group'}
  label={floating ? undefined : model.name}
  class={floating
    ? 'text-strong absolute z-30 flex w-[clamp(280px,calc(100vw_-_32px),360px)] flex-col text-[13px]'
    : 'text-strong flex w-full flex-col text-[13px]'}
  style={floating ? TOOLTIP_PANEL_PLACEMENT : undefined}
  foot={model.sourceLines.length > 0 ? sources : undefined}
>
  <div class="flex flex-col gap-2 px-[14px] py-3">
    <div class="flex items-start gap-2">
      <ItemIcon {icon} quality={model.quality} {build} size={40} />
      <div class="flex min-w-0 flex-1 flex-col gap-0.5">
        <span class={`text-[15px] font-bold ${rarityClassFor(model.quality)}`}>
          {model.name}
        </span>
        {#if model.itemLevel > 0}
          <span class="text-muted numeric text-[12px]">Item Level {model.itemLevel}</span>
        {/if}
        {#if model.bindLabel}
          <span class="text-[12px]" data-testid="item-tooltip-bind">{model.bindLabel}</span>
        {/if}
        {#if model.unique}
          <span class="text-muted text-[12px]">Unique</span>
        {/if}
        {#if model.clientUnconfirmed}
          <span class="text-muted text-[12px]">{bisCopy.clientUnconfirmedTitle}</span>
        {/if}
      </div>
    </div>

    {#if model.slotLabel !== '' || model.typeLabel}
      <div class="text-muted flex flex-wrap justify-between gap-x-2 text-[12px]">
        <span>{model.slotLabel}</span>
        {#if model.typeLabel}<span>{model.typeLabel}</span>{/if}
      </div>
    {/if}

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

    {#if whiteStats.length > 0}
      <div class="flex flex-col text-[12px]" data-testid="item-tooltip-stats">
        {#each whiteStats as line (line)}
          <!-- The client paints a negative stat red. -->
          <span class:text-death={line.startsWith('-')}>{line}</span>
        {/each}
      </div>
    {/if}

    {#if enchant !== undefined && enchant !== ''}
      <span class="text-rarity-uncommon text-[12px]" data-testid="item-tooltip-enchant">{enchant}</span>
    {/if}

    {#if equipStats.length > 0}
      <div class="text-rarity-uncommon flex flex-col text-[12px]">
        {#each equipStats as line (line)}
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

    {#if model.requiredLevel > 0}
      <span class="text-muted text-[12px]">Requires Level {model.requiredLevel}</span>
    {/if}
    {#if tail}<div class="text-tooltip-yellow text-[12px]">{@render tail()}</div>{/if}
  </div>
</Overlay>

{#snippet sources()}
  <div class="flex flex-col gap-0.5">
    {#each model.sourceLines as line (line)}
      <span>{line}</span>
    {/each}
  </div>
{/snippet}
