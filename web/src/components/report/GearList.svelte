<!-- web/src/components/report/GearList.svelte -->
<!-- What a player wore at pull (web layout spec 2026-11-05 §4.2 `report`; the b1 report mock's
     Gear at the pull): the shared paperdoll frame titled with the player, the spec and the pull's
     rate, the selected slot's client tooltip in its centre, and beside it the worn list, each row
     the item in its quality colour over "Head · Item level 66 · Enchanted" (the log's own item
     level and whether it was enchanted). Under the frame's minimum the frame's list form carries
     those lines itself. The class's items file (the planner's) names the ids, fetched once; an id
     the build does not know is an unknown item with its slot word, never dropped. No gain. -->
<script lang="ts">
  import BisTooltipHost from '../BisTooltipHost.svelte';
  import GearDoll from '../character/GearDoll.svelte';
  import ItemRow from '../ui/ItemRow.svelte';
  import { gearCopy } from '../../lib/gear/copy';
  import { dollCellsFromSlots } from '../../lib/gear/doll-cells';
  import { averageItemLevel } from '../../lib/gear/paperdoll';
  import { FRAME_SLOTS } from '../../lib/ui/paperdoll-frame';
  import { slotsFromWorn } from '../../lib/gear/paperdoll-slots';
  import { loadItems } from '../../lib/planner/load';
  import type { Item } from '../../lib/planner/types';
  import { classSlugOf, wornAtPull } from '../../lib/report/planner-link';
  import type { GearItem } from '../../lib/report/types';

  let {
    gear,
    className,
    dataBuild,
    name = '',
    spec,
    figure,
  }: {
    gear: GearItem[];
    className: string | undefined;
    dataBuild: string;
    /** The combatant: the frame's title. */
    name?: string;
    /** After the title, small: "Fury Warrior". */
    spec?: string;
    /** The pull's rate at the bar's right ("182.4", "DPS"). */
    figure?: { value: string; unit: string };
  } = $props();

  let items = $state<Map<number, Item> | null>(null);
  let failed = $state(false);
  const classSlug = $derived(classSlugOf(className) ?? '');

  $effect(() => {
    const slug = classSlugOf(className);
    if (slug === null) {
      failed = true;
      return;
    }
    void loadItems(dataBuild, slug)
      .then((file) => {
        items = new Map(file.items.map((item) => [item.id, item]));
      })
      .catch(() => {
        failed = true;
      });
  });

  const pieces = $derived(wornAtPull(gear));
  const slots = $derived(
    slotsFromWorn(Object.fromEntries(Object.entries(pieces).map(([slot, piece]) => [slot, piece.ID])), {
      lookup: (id) => items?.get(id),
      failed,
      itemLevel: (slot) => pieces[slot]?.ItemLevel,
      detail: (slot) => {
        const piece = pieces[slot];
        if (piece === undefined) return undefined;
        const enchanted = (piece.Enchants ?? []).some((id) => id > 0);
        return gearCopy.paperdoll.reportDetail(piece.ItemLevel, enchanted);
      },
    }),
  );
  const loading = $derived(items === null && !failed);
  /** The side list: every worn piece in the frame's order, "Head · Item level 66 · Enchanted". */
  const worn = $derived.by(() => {
    if (loading) return [];
    const bySlot = new Map(dollCellsFromSlots(slots, 'report', classSlug).map((cell) => [cell.slot, cell]));
    return FRAME_SLOTS.flatMap((slot) => {
      const cell = bySlot.get(slot);
      return cell?.item ? [cell] : [];
    });
  });
  /** The mean item level of the worn set at pull, once every item is known (the paperdoll's rule). */
  const itemLevel = $derived(items === null || failed ? null : averageItemLevel(slots));
  const resolveItem = async (id: number) => {
    const { fetchItemTooltipModel } = await import('../../lib/items/lookup');
    return fetchItemTooltipModel(dataBuild, classSlug, id);
  };
</script>

{#if itemLevel !== null}
  <p class="text-muted mb-3 text-[13px] leading-5" data-testid="gear-list-ilvl">
    {gearCopy.paperdoll.averageItemLevel(itemLevel)}
  </p>
{/if}
<div class="gear-at-pull-box" data-testid="gear-list">
  <div class="gear-at-pull">
    <div class="gear-at-pull-doll">
      <GearDoll
        {slots}
        mode="report"
        build={dataBuild}
        {classSlug}
        {loading}
        {failed}
        title={name === '' ? gearCopy.paperdoll.atPullLabel : name}
        meta={spec}
        {figure}
        label={name === ''
          ? gearCopy.paperdoll.atPullLabel
          : `${name}: ${gearCopy.paperdoll.atPullLabel.toLowerCase()}`}
        testid="paperdoll"
      />
    </div>
    {#if worn.length > 0}
      <div class="gear-at-pull-list" data-testid="gear-list-rows">
        {#each worn as cell (cell.slot)}
          {#if cell.item}<ItemRow
              item={cell.item}
              build={dataBuild}
              source={cell.source ? `${cell.label} · ${cell.source}` : cell.label}
              unknown={cell.unknown}
              iconSize={28}
              bare
              testid={`gear-list-row-${cell.slot}`}
            />{/if}
        {/each}
      </div>
    {/if}
  </div>
</div>
<BisTooltipHost build={dataBuild} modelsElementId="report-gear-tooltip-models" resolve={resolveItem} />

<style>
  .gear-at-pull-box {
    container-type: inline-size;
  }
  .gear-at-pull {
    display: grid;
    gap: 20px;
  }
  .gear-at-pull-doll {
    width: 100%;
    max-width: 520px;
  }
  .gear-at-pull-list {
    display: flex;
    flex-direction: column;
    min-width: 0;
  }
  .gear-at-pull-list > :global(*) {
    border-bottom: 1px solid var(--color-line-soft);
  }
  /* The frame beside its list once both fit; the frame's own list form says it all on a phone. */
  @container (min-width: 760px) {
    .gear-at-pull {
      grid-template-columns: minmax(0, 520px) minmax(240px, 1fr);
      align-items: start;
    }
  }
  @container (max-width: 422.98px) {
    .gear-at-pull-list {
      display: none;
    }
  }
</style>
