<!-- web/src/components/guild/RaidLootItem.svelte -->
<!-- One awarded item (attendance spec 3.3, 3.4): the shared ItemHover, with the icon at the caller's
     size and the quality-coloured name. The icon comes with the award (the API's loot table, as the
     Loot tab's); an item with none, or one whose file fails, shows the client's question mark. The
     tooltip opens for every item: an item outside every class file gets the award's own facts. -->
<script lang="ts">
  import type { RaidLootAward } from '../../lib/guild/attendance-contract';
  import { awardTooltipModel, lootIconOf, QUESTION_MARK_ICON } from '../../lib/guild/loot-icon';
  import { fetchItemTooltipModel } from '../../lib/items/lookup';
  import type { ItemTooltipModel } from '../../lib/items/tooltip';
  import { dataUrl } from '../../lib/planner/load';
  import { rarityClassFor } from '../../lib/planner/items';
  import { classSlugFromName } from '../../lib/report/tree-sizes';
  import activeBuild from '../../data/active-build.json';
  import ItemHover from '../ItemHover.svelte';

  let {
    award,
    iconSize,
    nameSize,
  }: {
    award: RaidLootAward;
    /** 28 in the Report, 20 in Your night. */
    iconSize: 28 | 20;
    /** The name's Tailwind text size: `text-[14px]` or `text-[13px]`. */
    nameSize: string;
  } = $props();

  const build = activeBuild.build;
  const classSlug = $derived(classSlugFromName(award.class));
  /** 0: the award's icon; 1: the question mark; 2: neither loaded, the name's two letters. */
  let fallbackStep = $state(0);
  const icon = $derived(fallbackStep === 0 ? lootIconOf(award) : QUESTION_MARK_ICON);
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

<ItemHover itemId={award.item_id} {classSlug} {build} model={ownModel} class="gap-2">
  <span
    class="rounded-control border-line bg-card-top text-muted inline-flex shrink-0 items-center justify-center overflow-hidden border text-[10px] font-bold"
    style={`width:${iconSize}px;height:${iconSize}px`}
    data-testid="raid-loot-icon"
  >
    {#if fallbackStep < 2}
      <img
        src={dataUrl(build, `icons/${icon}.webp`)}
        alt=""
        width={iconSize}
        height={iconSize}
        class="h-full w-full object-cover"
        onerror={() => (fallbackStep += 1)}
      />
    {:else}
      <span aria-hidden="true">{award.item_name.slice(0, 2)}</span>
    {/if}
  </span>
  <span class="{rarityClassFor(award.quality)} {nameSize} truncate">{award.item_name}</span>
</ItemHover>
