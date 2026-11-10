<!-- web/src/components/report/GearList.svelte -->
<!-- What a player wore at pull (web layout spec 2026-11-05 §4.2 `report`; the b1 report mock's
     Gear at the pull): the shared paperdoll frame titled with the player, the spec and the pull's
     rate, the selected slot's client tooltip in its centre, and beside it the worn list in the
     client's slot order, each row the item in its quality colour over "Head · Item level 66 ·
     Crusader" (the log's own item level, then its enchant in green, named from the build's enchant
     table, else "Enchanted"); a two-hander's off hand says so. Under the frame's minimum the frame's
     list form carries those lines itself. The class's items file (the planner's) names the ids,
     fetched once; an id the build does not know is an unknown item with its slot word, never
     dropped. No gain. The set's item level is the section's meta (the log's own). -->
<script lang="ts">
  import BisTooltipHost from '../BisTooltipHost.svelte';
  import GearDoll from '../character/GearDoll.svelte';
  import ItemIcon from '../ui/ItemIcon.svelte';
  import ItemRow from '../ui/ItemRow.svelte';
  import Ruled from '../ui/Ruled.svelte';
  import { gearCopy } from '../../lib/gear/copy';
  import { dollCellsFromSlots } from '../../lib/gear/doll-cells';
  import { FRAME_SLOTS } from '../../lib/ui/paperdoll-frame';
  import { FACT_SEPARATOR } from '../../lib/ui/primitives';
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
  /** The client's name for each enchant id; empty until the build's table loads (or with none). */
  let enchantNames = $state.raw<ReadonlyMap<number, string>>(new Map());
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
    void Promise.all([import('../../lib/sim/enchants'), import('../../lib/sim/enchant-names')])
      .then(async ([{ loadEnchants }, names]) => names.enchantNames(await loadEnchants(dataBuild)))
      .then(
        (names) => (enchantNames = names),
        (error: unknown) => console.error('report gear: the enchant table did not load', error),
      );
  });

  /** A worn piece's enchant: the client's name, else "Enchanted"; none when the log has no id. */
  function enchantOf(piece: GearItem): string | undefined {
    const id = (piece.Enchants ?? []).find((enchant) => enchant > 0);
    return id === undefined ? undefined : (enchantNames.get(id) ?? gearCopy.paperdoll.enchanted);
  }

  const pieces = $derived(wornAtPull(gear));
  const slots = $derived(
    slotsFromWorn(Object.fromEntries(Object.entries(pieces).map(([slot, piece]) => [slot, piece.ID])), {
      lookup: (id) => items?.get(id),
      failed,
      itemLevel: (slot) => pieces[slot]?.ItemLevel,
      detail: (slot) => {
        const piece = pieces[slot];
        return piece === undefined ? undefined : gearCopy.paperdoll.itemLevelLine(piece.ItemLevel);
      },
    }).map((slot) => {
      const piece = pieces[slot.slot];
      const enchant = piece === undefined ? undefined : enchantOf(piece);
      return enchant === undefined ? slot : { ...slot, enchant };
    }),
  );
  const loading = $derived(items === null && !failed);
  /** The side list: every worn piece in the frame's order, and an off hand a two-hander covers. */
  const worn = $derived.by(() => {
    if (loading) return [];
    const bySlot = new Map(dollCellsFromSlots(slots, 'report', classSlug).map((cell) => [cell.slot, cell]));
    return FRAME_SLOTS.flatMap((slot) => {
      const cell = bySlot.get(slot);
      return cell?.item || cell?.coveredBy ? [cell] : [];
    });
  });
  const resolveItem = async (id: number) => {
    const { fetchItemTooltipModel } = await import('../../lib/items/lookup');
    return fetchItemTooltipModel(dataBuild, classSlug, id);
  };
</script>

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
      <div class="gear-at-pull-list">
        <Ruled edge="bottom" testid="gear-list-rows">
          {#each worn as cell (cell.slot)}
            {#snippet facts()}{cell.label}{#if cell.source}{FACT_SEPARATOR}{cell.source}{/if}{/snippet}
            {#snippet enchant()}<span class="text-rarity-uncommon">{cell.enchant}</span>{/snippet}
            {#if cell.item}<ItemRow
                item={cell.item}
                build={dataBuild}
                sourceContent={facts}
                detail={cell.enchant ? enchant : undefined}
                unknown={cell.unknown}
                iconSize={28}
                bare
                testid={`gear-list-row-${cell.slot}`}
              />{:else if cell.coveredBy}
              <!-- The off hand a two-hander covers: its icon greyed, as the frame and the client draw it. -->
              <div class="gear-at-pull-covered" data-testid={`gear-list-row-${cell.slot}`}>
                <span class="gear-at-pull-covered-icon"
                  ><ItemIcon
                    icon={cell.coveredBy.icon}
                    quality={cell.coveredBy.quality}
                    build={dataBuild}
                    size={28}
                  /></span
                >
                <span class="min-w-0">
                  <span class="text-muted block text-[15px] leading-5">{cell.label}</span>
                  <span class="text-muted block text-[13px] leading-[18px]"
                    >{gearCopy.paperdoll.twoHanded}</span
                  >
                </span>
              </div>
            {/if}
          {/each}
        </Ruled>
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
    max-width: 470px;
  }
  .gear-at-pull-list {
    display: flex;
    flex-direction: column;
    min-width: 0;
  }
  .gear-at-pull-covered {
    display: grid;
    grid-template-columns: auto minmax(0, 1fr);
    column-gap: 12px;
    align-items: center;
  }
  .gear-at-pull-covered-icon {
    filter: grayscale(1) brightness(0.55);
  }
  /* The frame beside its list once both fit; the frame's own list form says it all on a phone. */
  @container (min-width: 760px) {
    .gear-at-pull {
      grid-template-columns: minmax(0, 470px) minmax(280px, 1fr);
      align-items: start;
    }
  }
  @container (max-width: 422.98px) {
    .gear-at-pull-list {
      display: none;
    }
  }
</style>
