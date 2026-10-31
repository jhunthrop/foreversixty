<!-- web/src/components/report/GearList.svelte -->
<!-- What a player wore at pull, as the shared Paperdoll (design/specs/2026-10-31-paperdoll.md,
     call site 3). The combatant row carries item ids in the log's slot order; the class's items
     file (the planner's) names them, fetched once when the row is opened. Each slot's detail line
     is the log's own item level and whether it was enchanted; an id the build does not know is
     drawn as an unknown item with its slot word, never dropped. The centre's source reads
     "At pull". -->
<script lang="ts">
  import Paperdoll from '../character/Paperdoll.svelte';
  import { gearCopy } from '../../lib/gear/copy';
  import type { PaperdollIdentity } from '../../lib/gear/paperdoll';
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
    spec = '',
  }: {
    gear: GearItem[];
    className: string | undefined;
    dataBuild: string;
    /** The combatant, for the frame's centre. */
    name?: string;
    /** "Protection", when the report names it. */
    spec?: string;
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
  const identity = $derived<PaperdollIdentity>({
    name,
    classSlug,
    race: '',
    spec,
    level: null,
    faction: null,
    source: gearCopy.paperdoll.atPull,
  });
</script>

<Paperdoll
  {identity}
  {slots}
  build={dataBuild}
  {classSlug}
  loading={items === null && !failed}
  {failed}
  testid="gear-list"
/>
