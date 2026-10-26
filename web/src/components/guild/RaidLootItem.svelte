<!-- web/src/components/guild/RaidLootItem.svelte -->
<!-- One awarded item (attendance spec 3.3, 3.4): the shared ItemHover, with the icon at the caller's
     size and the quality-coloured name. The name and quality come from the award, so the row holds
     its height while the icon (the awardee class's item file) loads; the tooltip loads on open. -->
<script lang="ts">
  import type { RaidLootAward } from '../../lib/guild/attendance-contract';
  import { loadItems, dataUrl } from '../../lib/planner/load';
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
  let icon = $state<string | null>(null);

  $effect(() => {
    let live = true;
    loadItems(build, classSlug)
      .then((file) => {
        if (live) icon = file.items.find((item) => item.id === award.item_id)?.icon ?? null;
      })
      .catch(() => {
        // No icon: the box stays and the tooltip still opens from the item id.
      });
    return () => {
      live = false;
    };
  });
</script>

<ItemHover itemId={award.item_id} {classSlug} {build} class="gap-2">
  <span
    class="rounded-control border-line inline-block shrink-0 overflow-hidden border"
    style={`width:${iconSize}px;height:${iconSize}px`}
  >
    {#if icon !== null}
      <img
        src={dataUrl(build, `icons/${icon}.webp`)}
        alt=""
        width={iconSize}
        height={iconSize}
        class="h-full w-full object-cover"
      />
    {/if}
  </span>
  <span class="{rarityClassFor(award.quality)} {nameSize} truncate">{award.item_name}</span>
</ItemHover>
