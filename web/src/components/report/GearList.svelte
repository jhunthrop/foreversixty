<!-- web/src/components/report/GearList.svelte -->
<!-- What a player wore at pull, as the shared paperdoll in `report` mode in one Plate (web layout
     spec 2026-11-05 §4.2). The combatant row carries item ids in the log's slot order; the class's
     items file (the planner's) names them, fetched once when the row is opened. Each slot's line 2
     is the log's own item level and whether it was enchanted ("Item level 66 · Enchanted"); an id
     the build does not know is drawn as an unknown item with its slot word, never dropped. No
     gain. The doll's own tooltip host opens each item's real tooltip. -->
<script lang="ts">
  import BisTooltipHost from '../BisTooltipHost.svelte';
  import GearDoll from '../character/GearDoll.svelte';
  import Plate from '../ui/Plate.svelte';
  import { gearCopy } from '../../lib/gear/copy';
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
  }: {
    gear: GearItem[];
    className: string | undefined;
    dataBuild: string;
    /** The combatant, for the doll's accessible name. */
    name?: string;
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
  const resolveItem = async (id: number) => {
    const { fetchItemTooltipModel } = await import('../../lib/items/lookup');
    return fetchItemTooltipModel(dataBuild, classSlug, id);
  };
</script>

<Plate label={gearCopy.paperdoll.atPullLabel} testid="gear-list">
  <GearDoll
    {slots}
    mode="report"
    build={dataBuild}
    {classSlug}
    loading={items === null && !failed}
    {failed}
    label={name === ''
      ? gearCopy.paperdoll.atPullLabel
      : `${name}: ${gearCopy.paperdoll.atPullLabel.toLowerCase()}`}
    testid="paperdoll"
  />
</Plate>
<BisTooltipHost build={dataBuild} modelsElementId="report-gear-tooltip-models" resolve={resolveItem} />
