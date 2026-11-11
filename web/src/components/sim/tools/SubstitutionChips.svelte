<!-- web/src/components/sim/tools/SubstitutionChips.svelte -->
<!-- What a combination changed, as plain inline item names with their icons (no pill boxes; a
     Tag only for the emptied off-hand, which is a state, not an item).
     The NAME always comes off the substitution -- contract 10.1 A6 fills it for items too,
     from simdb -- so a chip is legible even where the per-class item file has not loaded.
     The item map is consulted for the icon and the rarity colour only, which are the two
     things the result does not carry.

     Engine-lane rule 5: a two-hander replacing a main-plus-off-hand pair emits a SECOND
     substitution, `{kind: "item", slot: "off_hand", item_id: 0, name: "<item removed>"}`.
     `item_id: 0` is never a real item (simdb has no id 0), so `isEmptiedOffHand` catches it
     before the generic branch would look it up (a miss, since nothing is at id 0), fall
     back to `text-text`, and print "<item removed>" verbatim -- a sentence, not a name,
     that a player should never read as if the engine named an item that. It renders as a
     plainly emptied slot instead, through `bulkCopy.offHandEmptied`.

     The predicate is imported from combos.ts, not written again here: rule 5 binds every
     rendered AND exported path, and two copies of it is exactly how the addon string and
     the planner link came to keep the sentinel (final whole-branch review, Important 3).

     Store-free by design: DropResults (Task 18) and SavedCombos (Task 20) reuse this
     component, and neither owns a BulkStore in the shape this one would need. Both inherit
     this fix along with everything else here, since neither has its own chip renderer.

     Task 5 (newcomer MAJOR, review.md:360-363): a chip's source (the boss or vendor it came
     from) used to live only in `title`, invisible on a phone. A `<details>` per chip -- this
     codebase's own house disclosure -- would be absurd at this size, so the source is folded
     straight into the chip's own visible text instead (`substitutionChipLabel`, combos.ts):
     source names here are short (a boss or vendor name), the row already wraps
     (`flex-wrap`), and a reader gets the full answer without a second element to manage. -->
<script lang="ts">
  import ItemName from '../../ui/ItemName.svelte';
  import ItemHover from '../../ItemHover.svelte';
  import { rarityClassFor } from '../../../lib/planner/items';
  import type { Item } from '../../../lib/planner/types';
  import type { Substitution } from '../../../lib/sim/bulk-types';
  import { isEmptiedOffHand, substitutionChipLabel } from '../../../lib/sim/combos';
  import { bulkCopy } from '../../../lib/sim/copy';
  import ItemIcon from '../../ui/ItemIcon.svelte';
  import Tag from '../../ui/Tag.svelte';

  let {
    substitutions,
    items,
    treeVersion,
    classSlug,
  }: {
    substitutions: readonly Substitution[];
    items: ReadonlyMap<number, Item>;
    treeVersion: string;
    /** For ItemHover's fetch key. Optional: a caller with no classSlug in hand (none
     *  today) still gets the chip, just without its own hover tooltip. */
    classSlug?: string;
  } = $props();
</script>

<span class="flex flex-wrap items-center gap-x-4 gap-y-1">
  {#each substitutions as sub, index (`${sub.kind}-${sub.slot ?? ''}-${sub.item_id ?? sub.name ?? index}`)}
    {#if isEmptiedOffHand(sub)}
      <Tag tone="muted">{bulkCopy.offHandEmptied}</Tag>
    {:else}
      {@const item = sub.item_id === undefined ? undefined : items.get(sub.item_id)}
      <span class="inline-flex items-center gap-2 text-[15px] leading-5 font-bold">
        {#if item !== undefined && classSlug !== undefined}
          <ItemHover itemId={item.id} {classSlug} build={treeVersion} class="gap-2">
            <ItemIcon icon={item.icon} quality={item.quality} build={treeVersion} size={28} />
            <ItemName name={substitutionChipLabel(sub)} quality={item.quality} />
          </ItemHover>
        {:else}
          {#if item !== undefined}
            <ItemIcon icon={item.icon} quality={item.quality} build={treeVersion} size={28} />
          {/if}
          <span class={item === undefined ? 'text-strong' : rarityClassFor(item.quality)}>
            {substitutionChipLabel(sub)}
          </span>
        {/if}
        <!-- Task 7: the source name used to be hidden in a title= (never fires on touch).
             Shown inline, muted, only when the substitution actually carries one. -->
        {#if sub.source_name !== undefined && sub.source_name !== ''}
          <span class="text-muted text-[13px] font-normal" data-testid="sim-substitution-source"
            >· {sub.source_name}</span
          >
        {/if}
      </span>
    {/if}
  {/each}
</span>
