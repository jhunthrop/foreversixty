<!-- web/src/components/guild/report/ReportLootRow.svelte -->
<!-- One awarded item as an ItemRow (web layout spec 2026-11-05 §3A.2): the icon in its quality
     frame, the quality-coloured name, `Onyxia · to Doomwhisper` under it. The row draws the item's
     icon from the loot table (the client's own art, a quest item's too); ItemIcon draws the empty
     socket for one the build does not ship, never a broken image.
     The page's tooltip host (ReportLoot) opens the real item on hover, focus and long-press. -->
<script lang="ts">
  import activeBuild from '../../../data/active-build.json';
  import type { RaidLootAward } from '../../../lib/guild/attendance-contract';
  import { lootIconOf } from '../../../lib/guild/loot-icon';
  import { reportNightCopy as copy } from '../../../lib/guild/report-night-copy';
  import ItemRow from '../../ui/ItemRow.svelte';

  let { award, boss }: { award: RaidLootAward; boss: string } = $props();

  const build = activeBuild.build;
  const icon = $derived(lootIconOf(award) ?? undefined);
</script>

<div data-testid="raid-loot-award">
  <ItemRow
    item={{ id: award.item_id, name: award.item_name, quality: award.quality, icon }}
    {build}
    iconSize={32}
    source={copy.loot.to(boss, award.name)}
    testid="raid-loot-row"
  />
</div>
