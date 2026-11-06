<!-- web/src/components/guild/report/ReportLootRow.svelte -->
<!-- One awarded item as an ItemRow (web layout spec 2026-11-05 §3A.2): the icon in its quality
     frame, the quality-coloured name, `Onyxia · to Doomwhisper` under it. The build ships an item's
     icon only when a class item file carries it (lib/guild/loot-icon.ts), so the row asks the
     winner's class file first and draws the empty socket for a drop no class file has, never a 404.
     The page's tooltip host (ReportLoot) opens the real item on hover, focus and long-press. -->
<script lang="ts">
  import activeBuild from '../../../data/active-build.json';
  import type { RaidLootAward } from '../../../lib/guild/attendance-contract';
  import { shippedLootIcon, type ClassFileRow } from '../../../lib/guild/loot-icon';
  import { reportNightCopy as copy } from '../../../lib/guild/report-night-copy';
  import { fetchItemTooltipModel } from '../../../lib/items/lookup';
  import { classSlugFromName } from '../../../lib/report/tree-sizes';
  import ItemRow from '../../ui/ItemRow.svelte';

  let { award, boss }: { award: RaidLootAward; boss: string } = $props();

  const build = activeBuild.build;
  let classRow = $state<ClassFileRow>('pending');
  const icon = $derived(shippedLootIcon(award, classRow) ?? undefined);

  $effect(() => {
    const classSlug = classSlugFromName(award.class);
    let live = true;
    const settle = (row: ClassFileRow): void => {
      if (live) classRow = row;
    };
    if (classSlug === '') settle('none');
    else
      fetchItemTooltipModel(build, classSlug, award.item_id)
        .then((model) => settle(model === null ? 'none' : 'found'))
        .catch(() => settle('none'));
    return () => {
      live = false;
    };
  });
</script>

<li data-testid="raid-loot-award">
  <ItemRow
    item={{ id: award.item_id, name: award.item_name, quality: award.quality, icon }}
    {build}
    iconSize={32}
    source={copy.loot.to(boss, award.name)}
    testid="raid-loot-row"
  />
</li>
