<!-- web/src/components/guild/RaidLootItem.svelte -->
<!-- One awarded item (attendance spec 3.3, 3.4): the shared ItemHover, with the icon at the caller's
     size and the quality-coloured name. The icon comes with the award (the API's loot table, the
     client's own art); an item with none, or one whose file fails, shows the /bis neutral tile with
     the name's letters, never the red question mark. The tooltip opens for every item: an item
     outside every class file gets the award's own facts. -->
<script lang="ts">
  import type { RaidLootAward } from '../../lib/guild/attendance-contract';
  import { awardTooltipModel, lootIconOf, lootInitials } from '../../lib/guild/loot-icon';
  import { fetchItemTooltipModel } from '../../lib/items/lookup';
  import type { ItemTooltipModel } from '../../lib/items/tooltip';
  import { dataUrl } from '../../lib/planner/load';
  import { rarityBorderColorFor, rarityClassFor } from '../../lib/planner/items';
  import { classSlugFromName } from '../../lib/report/tree-sizes';
  import activeBuild from '../../data/active-build.json';
  import ItemHover from '../ItemHover.svelte';

  let {
    award,
    iconSize,
    nameSize,
    show = 'both',
  }: {
    award: RaidLootAward;
    /** 28 in the Report, 20 in Your night, 36 in a loot card's head. */
    iconSize: 36 | 28 | 20;
    /** The name's Tailwind text size: `text-[14px]` or `text-[13px]`. */
    nameSize: string;
    /** A card head draws the icon and the name apart (the lead and the title), each with the tooltip. */
    show?: 'both' | 'icon' | 'name';
  } = $props();

  const build = activeBuild.build;
  const classSlug = $derived(classSlugFromName(award.class));
  /** Set when the icon file fails: the tile shows the name's letters instead. */
  let broken = $state(false);
  const icon = $derived(broken ? null : lootIconOf(award));
  /** Set when the class file has no row for the item (a quest item such as Head of Onyxia). */
  let ownModel = $state<ItemTooltipModel | undefined>(undefined);

  $effect(() => {
    let live = true;
    fetchItemTooltipModel(build, classSlug, award.item_id)
      .then((model) => {
        if (live && model === null) ownModel = awardTooltipModel(award);
      })
      .catch(() => {
        if (live) ownModel = awardTooltipModel(award);
      });
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
  class="gap-2"
  label={show === 'icon' ? award.item_name : undefined}
>
  {#if show !== 'name'}
    <span
      class="rounded-control bg-card-top text-muted inline-flex shrink-0 items-center justify-center overflow-hidden border text-[10px] font-bold"
      style={`width:${iconSize}px;height:${iconSize}px;border-width:2px;border-color:${rarityBorderColorFor(award.quality)}`}
      data-testid="raid-loot-icon"
    >
      {#if icon !== null}
        <img
          src={dataUrl(build, `icons/${icon}.webp`)}
          alt=""
          width={iconSize}
          height={iconSize}
          class="h-full w-full object-cover"
          onerror={() => (broken = true)}
        />
      {:else}
        <span aria-hidden="true" data-testid="raid-loot-initials">{lootInitials(award.item_name)}</span>
      {/if}
    </span>
  {/if}
  {#if show !== 'icon'}
    <span class="{rarityClassFor(award.quality)} {nameSize} truncate">{award.item_name}</span>
  {/if}
</ItemHover>
