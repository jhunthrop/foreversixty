<!-- web/src/components/character/WornList.svelte -->
<!-- The worn list beside a shared paperdoll frame (the report's Gear at the pull, the sim's Gear):
     every worn piece in the frame's order, each row the item in its quality colour over
     "Head · Item level 66" and its enchant in the client's green (lib/gear/doll-cells.ts); an off
     hand a two-hander covers says so, its icon greyed as the frame and the client draw it. No gain.
     Nothing while the items file loads. The caller hides it under the frame's 423 px minimum,
     where the frame's own list form carries these lines. -->
<script lang="ts">
  import ItemIcon from '../ui/ItemIcon.svelte';
  import ItemRow from '../ui/ItemRow.svelte';
  import Ruled from '../ui/Ruled.svelte';
  import { gearCopy } from '../../lib/gear/copy';
  import { dollCellsFromSlots, type WornDollMode } from '../../lib/gear/doll-cells';
  import type { PaperdollSlot } from '../../lib/gear/paperdoll';
  import { FRAME_SLOTS } from '../../lib/gear/slots';
  import { FACT_SEPARATOR } from '../../lib/ui/primitives';

  let {
    slots,
    mode,
    build,
    classSlug = '',
    loading = false,
    testid,
    rowTestid,
  }: {
    slots: readonly PaperdollSlot[];
    mode: WornDollMode;
    /** The data build the icons are read from. */
    build: string;
    classSlug?: string;
    /** The items file is pending: no rows yet. */
    loading?: boolean;
    testid: string;
    /** Each row's testid prefix: "<rowTestid>-<slot>". */
    rowTestid: string;
  } = $props();

  /** The rows: every worn piece in the frame's order, and an off hand a two-hander covers. */
  const worn = $derived.by(() => {
    if (loading) return [];
    const bySlot = new Map(dollCellsFromSlots(slots, mode, classSlug).map((cell) => [cell.slot, cell]));
    return FRAME_SLOTS.flatMap((slot) => {
      const cell = bySlot.get(slot);
      return cell?.item || cell?.coveredBy ? [cell] : [];
    });
  });
</script>

{#if worn.length > 0}
  <Ruled edge="bottom" {testid}>
    {#each worn as cell (cell.slot)}
      {#snippet facts()}{cell.label}{#if cell.source}{FACT_SEPARATOR}{cell.source}{/if}{/snippet}
      {#snippet enchant()}<span class="text-rarity-uncommon">{cell.enchant}</span>{/snippet}
      {#if cell.item}<ItemRow
          item={cell.item}
          {build}
          sourceContent={facts}
          detail={cell.enchant ? enchant : undefined}
          unknown={cell.unknown}
          iconSize={28}
          bare
          testid={`${rowTestid}-${cell.slot}`}
        />{:else if cell.coveredBy}
        <div class="worn-covered" data-testid={`${rowTestid}-${cell.slot}`}>
          <span class="worn-covered-icon"
            ><ItemIcon icon={cell.coveredBy.icon} quality={cell.coveredBy.quality} {build} size={28} /></span
          >
          <span class="min-w-0">
            <span class="text-muted block text-[15px] leading-5">{cell.label}</span>
            <span class="text-muted block text-[13px] leading-[18px]">{gearCopy.paperdoll.twoHanded}</span>
          </span>
        </div>
      {/if}
    {/each}
  </Ruled>
{/if}

<style>
  .worn-covered {
    display: grid;
    grid-template-columns: auto minmax(0, 1fr);
    column-gap: 12px;
    align-items: center;
  }
  .worn-covered-icon {
    filter: grayscale(1) brightness(0.55);
  }
</style>
