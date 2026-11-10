<!-- web/src/components/guild/RaidLootItem.svelte -->
<!-- One awarded item (attendance spec 3.3, 3.4): the shared ItemHover, with ui/ItemIcon at the
     caller's size and the quality-coloured name. The icon comes with the award (the API's loot
     table, the client's own art, a quest item's too); an item with none, or one whose file the
     build does not ship, shows ItemIcon's dark socket, the same as the Loot tab, never the red
     question mark. With no class to read (a drop nobody here is eligible for) no class file is
     fetched at all. The tooltip opens for every item: an item outside every class file gets the
     award's own facts. -->
<script lang="ts">
  import ItemName from '../ui/ItemName.svelte';
  import type { RaidLootAward } from '../../lib/guild/attendance-contract';
  import { awardTooltipModel } from '../../lib/guild/loot-icon';
  import { fetchItemTooltipModel } from '../../lib/items/lookup';
  import type { ItemTooltipModel } from '../../lib/items/tooltip';
  import { classSlugFromName } from '../../lib/report/tree-sizes';
  import activeBuild from '../../data/active-build.json';
  import ItemHover from '../ItemHover.svelte';
  import ItemIcon from '../ui/ItemIcon.svelte';

  let {
    award,
    iconSize,
    nameSize,
    show = 'both',
  }: {
    award: RaidLootAward;
    /** 28 in the Report, 20 in Your night, 40 in a loot card's head, 24 in a collapsed drop, 16 in a
     *  line of text. */
    iconSize: 40 | 28 | 24 | 20 | 16;
    /** The name's Tailwind text size: `text-[14px]` or `text-[13px]`. */
    nameSize: string;
    /** A card head draws the icon and the name apart (the lead and the title), each with the tooltip. */
    show?: 'both' | 'icon' | 'name';
  } = $props();

  const build = activeBuild.build;
  const classSlug = $derived(classSlugFromName(award.class));
  /** Set when the class file has no row for the item (a quest item such as Head of Onyxia). */
  let ownModel = $state<ItemTooltipModel | undefined>(undefined);

  $effect(() => {
    let live = true;
    const none = (): void => {
      if (!live) return;
      ownModel = { ...awardTooltipModel(award), icon: '' };
    };
    if (classSlug === '') {
      none();
      return;
    }
    fetchItemTooltipModel(build, classSlug, award.item_id)
      .then((model) => {
        if (model === null) none();
      })
      .catch(none);
    return () => {
      live = false;
    };
  });
</script>

<ItemHover
  itemId={award.item_id}
  {classSlug}
  {build}
  model={ownModel}
  class={iconSize === 16 ? 'gap-1' : 'gap-2'}
  label={show === 'icon' ? award.item_name : undefined}
>
  {#if show !== 'name'}
    <ItemIcon icon={award.icon} quality={award.quality} {build} size={iconSize} testid="raid-loot-icon" />
  {/if}
  {#if show !== 'icon'}
    <ItemName name={award.item_name} quality={award.quality} class="{nameSize} truncate" />
  {/if}
</ItemHover>
