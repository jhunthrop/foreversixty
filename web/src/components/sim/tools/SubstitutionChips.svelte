<!-- web/src/components/sim/tools/SubstitutionChips.svelte -->
<!-- What a combination changed, as icon chips.
     The NAME always comes off the substitution -- contract 10.1 A6 fills it for items too,
     from simdb -- so a chip is legible even where the per-class item file has not loaded.
     The item map is consulted for the icon and the rarity colour only, which are the two
     things the result does not carry.

     Store-free by design: DropResults (Task 18) and SavedCombos (Task 20) reuse this
     component, and neither owns a BulkStore in the shape this one would need. -->
<script lang="ts">
  import { rarityClassFor } from '../../../lib/planner/items';
  import { dataUrl } from '../../../lib/planner/load';
  import type { Item } from '../../../lib/planner/types';
  import type { Substitution } from '../../../lib/sim/bulk-types';
  import { substitutionLabel } from '../../../lib/sim/combos';

  let {
    substitutions,
    items,
    treeVersion,
  }: {
    substitutions: readonly Substitution[];
    items: ReadonlyMap<number, Item>;
    treeVersion: string;
  } = $props();
</script>

<span class="flex flex-wrap items-center gap-1">
  {#each substitutions as sub, index (`${sub.kind}-${sub.slot ?? ''}-${sub.item_id ?? sub.name ?? index}`)}
    {@const item = sub.item_id === undefined ? undefined : items.get(sub.item_id)}
    <span
      class="border-line rounded-pill inline-flex items-center gap-1 border px-2 py-[2px] text-[12px]"
      title={sub.source_name !== undefined && sub.source_name !== ''
        ? `${substitutionLabel(sub)} · ${sub.source_name}`
        : substitutionLabel(sub)}
    >
      {#if item !== undefined}
        <img
          src={dataUrl(treeVersion, `icons/${item.icon}.webp`)}
          alt=""
          width="16"
          height="16"
          loading="lazy"
          decoding="async"
          class="rounded-control h-4 w-4 object-cover"
        />
      {/if}
      <span class={item === undefined ? 'text-text' : rarityClassFor(item.quality)}>
        {substitutionLabel(sub)}
      </span>
    </span>
  {/each}
</span>
