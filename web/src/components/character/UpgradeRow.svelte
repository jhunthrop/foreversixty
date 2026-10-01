<!-- web/src/components/character/UpgradeRow.svelte -->
<!-- Home rebuild spec §3.B.3's "Your upgrades" table row: a worn-vs-BiS two-column diff,
     built from the same icon/rarity-border/rarity-name/`ItemHover` primitives
     `web/src/components/bis/GearRow.astro` already uses for the `/bis` page's own pick row
     -- not a fork of that component's whole markup (its own shape is a single pick plus a
     runners-up list, not a worn-vs-BiS diff, and it is an Astro component a Svelte island
     cannot render anyway). Every item name opens the shared item tooltip via `ItemHover`,
     the same as everywhere else on the site (tenet 2). -->
<script lang="ts">
  import type { SlotUpgrade } from '../../lib/home/upgrades';
  import { homeUpgradesCopy } from '../../lib/home-panel-copy';
  import { bisCopy } from '../../lib/bis/copy';
  import { slotSourceLine } from '../../lib/home/source-line';
  import { dataUrl } from '../../lib/planner/load';
  import { rarityBorderColorFor, rarityClassFor } from '../../lib/planner/items';
  import { SLOT_DISPLAY_LABELS } from '../../lib/bis/slot-display-labels';
  import type { Slot } from '../../lib/planner/types';
  import type { Item } from '../../lib/planner/types';
  import ItemHover from '../ItemHover.svelte';

  let {
    upgrade,
    wornItem,
    pickItem,
    build,
    classSlug,
  }: {
    upgrade: SlotUpgrade;
    /** `items.get(upgrade.wornItemId)` -- undefined when nothing is worn, or the id is
     *  unknown to this build's item table (`upgrade.wornUnknown`). */
    wornItem: Item | undefined;
    /** `items.get(upgrade.pick.item_id)` -- the pick's own icon/quality; falls back to a
     *  plain, uncoloured pill (same as `GearRow.astro`'s own `hasRealIcon` fallback) when
     *  this build's item table does not carry the id either. */
    pickItem: Item | undefined;
    build: string;
    classSlug: string;
  } = $props();

  const slotLabel = $derived(SLOT_DISPLAY_LABELS[upgrade.slot as Slot] ?? upgrade.slot);
  const pickSourceLine = $derived(slotSourceLine(upgrade.pick));
  const gainLabel = $derived(upgrade.gainDps === null ? '' : `+${upgrade.gainDps.toFixed(1)} DPS`);
</script>

<li class="upgrade-row" data-testid={`home-upgrade-row-${upgrade.slot}`}>
  <span class="upgrade-row-slot">{slotLabel}</span>

  <span class="upgrade-row-cell">
    {#if upgrade.wornItemId === undefined}
      <span class="upgrade-row-icon upgrade-row-icon-empty rounded-control border" aria-hidden="true"></span>
      <span class="upgrade-row-text">
        <span class="text-muted text-[13px]">{homeUpgradesCopy.noItemEquipped}</span>
      </span>
    {:else if upgrade.wornUnknown || wornItem === undefined}
      <span class="upgrade-row-icon upgrade-row-icon-empty rounded-control border" aria-hidden="true"></span>
      <span class="upgrade-row-text">
        <span class="text-muted text-[13px]">{homeUpgradesCopy.unknownItem}</span>
      </span>
    {:else}
      <ItemHover itemId={upgrade.wornItemId} {classSlug} {build} class="upgrade-row-item-hover">
        <img
          src={dataUrl(build, `icons/${wornItem.icon}.webp`)}
          alt=""
          width="36"
          height="36"
          loading="lazy"
          decoding="async"
          class="upgrade-row-icon rounded-control border object-cover"
          style={`border-color:${rarityBorderColorFor(wornItem.quality)}`}
        />
        <span class="upgrade-row-text">
          <span class={rarityClassFor(wornItem.quality)}>{wornItem.name}</span>
          <span class="text-muted text-[12px]">{homeUpgradesCopy.youWearThis}</span>
        </span>
      </ItemHover>
    {/if}
  </span>

  <span class="upgrade-row-arrow" aria-hidden="true">→</span>

  <span class="upgrade-row-cell">
    <ItemHover itemId={upgrade.pick.item_id} {classSlug} {build} class="upgrade-row-item-hover">
      {#if pickItem !== undefined}
        <img
          src={dataUrl(build, `icons/${pickItem.icon}.webp`)}
          alt=""
          width="36"
          height="36"
          loading="lazy"
          decoding="async"
          class="upgrade-row-icon rounded-control border object-cover"
          style={`border-color:${rarityBorderColorFor(pickItem.quality)}`}
        />
      {:else}
        <span class="upgrade-row-icon upgrade-row-icon-empty rounded-control border" aria-hidden="true"
        ></span>
      {/if}
      <span class="upgrade-row-text">
        <span class={rarityClassFor(pickItem?.quality ?? 1)}>{upgrade.pick.item_name}</span>
        <span class="text-muted text-[12px]">{pickSourceLine}</span>
      </span>
    </ItemHover>
  </span>

  {#if upgrade.notSimChecked}
    <!-- Fix round 1 item A.1(d): a weapon slot with no sim-verified alternative never shows a
         scoreItem-diffed number (that's the exact ~13x Ranger Bow overstatement the regression
         was). The BiS page's own vocabulary for "ranked, but no full sim run behind it." -->
    <span
      class="upgrade-row-gain-unchecked"
      data-testid={`home-upgrade-gain-${upgrade.slot}`}
      title={bisCopy.notSimCheckedTitle}
    >
      {bisCopy.notSimCheckedTag}
    </span>
  {:else}
    <span class="upgrade-row-gain font-mono" data-testid={`home-upgrade-gain-${upgrade.slot}`}>
      {gainLabel}
    </span>
  {/if}
</li>

<style>
  .upgrade-row {
    display: grid;
    grid-template-columns: 90px 1fr 20px 1fr 90px;
    column-gap: 12px;
    align-items: center;
    padding: 10px 0;
    border-bottom: 1px solid var(--color-line-soft, #222);
  }
  .upgrade-row:last-child {
    border-bottom: none;
  }
  .upgrade-row-slot {
    font-size: 12px;
    font-weight: 700;
    text-transform: uppercase;
    color: var(--color-muted, #999);
  }
  .upgrade-row-cell {
    display: flex;
    min-width: 0;
    align-items: center;
    gap: 0.6rem;
  }
  /* ItemHover's own default pill classes already give this `inline-flex items-center
     gap-1`; this just widens the icon/name gap to match GearRow's own row rhythm and marks
     it clickable. */
  :global(.upgrade-row-item-hover) {
    gap: 0.6rem !important;
    cursor: pointer;
    min-width: 0;
  }
  .upgrade-row-icon {
    flex-shrink: 0;
    width: 36px;
    height: 36px;
    object-fit: cover;
  }
  .upgrade-row-icon-empty {
    background: transparent;
    border-color: var(--color-line-soft, #222);
    border-style: dashed;
  }
  .upgrade-row-text {
    display: flex;
    min-width: 0;
    flex-direction: column;
    gap: 1px;
    font-size: 13px;
    font-weight: 600;
  }
  .upgrade-row-arrow {
    text-align: center;
    color: var(--color-gold, #c9a35e);
  }
  .upgrade-row-gain {
    text-align: right;
    font-size: 13px;
    font-weight: 700;
    color: var(--color-kill, #7fd48a);
  }
  .upgrade-row-gain-unchecked {
    text-align: right;
    font-size: 11px;
    font-weight: 600;
    text-transform: uppercase;
    letter-spacing: 0.04em;
    color: var(--color-muted, #999);
  }

  @media (max-width: 720px) {
    .upgrade-row {
      grid-template-columns: 1fr;
      row-gap: 6px;
    }
    /* Spec §3.B.3/§5: the stacked phone row keeps the connecting arrow, rotated 90° to point
       down from the worn item to the BiS pick instead of sideways -- never hidden (fix round
       1 item B.6: this used to be `display: none`, dropping the one glyph that tells a
       player which item is "you" and which is "the upgrade" once the two columns stack). */
    .upgrade-row-arrow {
      display: block;
      transform: rotate(90deg);
    }
    .upgrade-row-gain,
    .upgrade-row-gain-unchecked {
      text-align: left;
    }
  }
</style>
