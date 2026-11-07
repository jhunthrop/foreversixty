<!-- web/src/components/guild/report/ReportLoot.svelte -->
<!-- Loot awarded, in the report's rail (web layout spec 2026-11-05 §7.2 "Raid report"): every
     award of the night as an ItemRow, boss by boss, and the page's one item tooltip host, which
     reads the winner's class file and falls back to the award's own facts for an item no class
     file carries (a quest item such as Head of Onyxia). -->
<script lang="ts">
  import activeBuild from '../../../data/active-build.json';
  import type { RaidLootGroup } from '../../../lib/guild/attendance-contract';
  import { attendanceCopy } from '../../../lib/guild/attendance-copy';
  import { awardTooltipModel } from '../../../lib/guild/loot-icon';
  import { reportNightCopy as copy } from '../../../lib/guild/report-night-copy';
  import { fetchItemTooltipModel } from '../../../lib/items/lookup';
  import type { ItemTooltipModel } from '../../../lib/items/tooltip';
  import { classSlugFromName } from '../../../lib/report/tree-sizes';
  import BisTooltipHost from '../../BisTooltipHost.svelte';
  import Section from '../../ui/Section.svelte';
  import ReportLootRow from './ReportLootRow.svelte';

  let {
    groups,
    officer,
    lootHref,
  }: { groups: readonly RaidLootGroup[]; officer: boolean; lootHref: string } = $props();

  const build = activeBuild.build;
  const awards = $derived(
    groups.flatMap((group) => group.awards.map((award) => ({ award, boss: group.boss }))),
  );
  const empty = attendanceCopy.loot;

  async function resolve(itemId: number): Promise<ItemTooltipModel | null> {
    const award = awards.find((entry) => entry.award.item_id === itemId)?.award;
    if (award === undefined) return null;
    const slug = classSlugFromName(award.class);
    const model = slug === '' ? null : await fetchItemTooltipModel(build, slug, itemId).catch(() => null);
    return model ?? { ...awardTooltipModel(award), icon: '' };
  }
</script>

<Section title={copy.loot.head} anchor="raid-loot" testid="raid-loot">
  {#if awards.length === 0}
    <p class="text-muted text-[14px]" data-testid="raid-loot-empty">
      {#if officer}{empty.officerEmptyLead}<a class="text-text underline" href={lootHref}
          >{empty.officerEmptyLink}</a
        >{empty.officerEmptyTail}{:else}{empty.memberEmpty}{/if}
    </p>
  {:else}
    <div>
      {#each awards as entry (`${entry.award.item_id}-${entry.award.at}-${entry.award.character_key}`)}
        <ReportLootRow award={entry.award} boss={entry.boss} />
      {/each}
    </div>
    <BisTooltipHost {build} modelsElementId="raid-loot-tooltip-models" {resolve} />
  {/if}
</Section>
