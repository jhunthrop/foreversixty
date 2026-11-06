<!-- web/src/components/sim/SimResourceSection.svelte -->
<!-- Where the rage, energy or mana came from and went (sim spec 2026-11-06 §4.5; spec
     2026-11-04-sim-results 3.10): an open Section titled by the power, three small Figures, then
     two open StatTables side by side, Gained from and Spent on, each row its icon, name, Meter and
     amount, the overflow under the name. The timeline's resource lane is the graph. -->
<script lang="ts">
  import type { ActionNames } from '../../lib/sim/action-names';
  import { createCardTables } from '../../lib/sim/card-tables.svelte';
  import { resourceCardCopy as copy } from '../../lib/sim/input-cards-copy';
  import { resourceCard, type ResourceRow } from '../../lib/sim/resource-card';
  import type { SimResult } from '../../lib/sim/types';
  import Figure from '../ui/Figure.svelte';
  import Section from '../ui/Section.svelte';
  import StatTable from '../ui/StatTable.svelte';
  import ActionStatRow from './ActionStatRow.svelte';

  let { result, names }: { result: SimResult; names: ActionNames | null } = $props();

  const tables = createCardTables(
    () => result,
    () => names,
    () => {
      const sources = result.resource_sources ?? [];
      return sources.flatMap((source) => [...source.gained, ...source.spent].map((row) => row.action));
    },
  );
  const card = $derived(resourceCard(result, tables.names));
  const AMOUNT = [{ id: 'amount', label: '' }];
</script>

{#snippet list(heading: string, total: string, rows: readonly ResourceRow[], testid: string)}
  <div class="min-w-0" data-testid={testid}>
    <h3 class="text-label tracking-label text-muted pt-4 pb-1 font-bold uppercase">{heading} · {total}</h3>
    <StatTable columns={AMOUNT} open label={heading}>
      {#each rows as row (row.action)}
        <ActionStatRow
          action={row.action}
          label={row.name}
          {tables}
          figures={[{ value: row.amount }]}
          fraction={row.barPct / 100}
          note={row.lost ?? undefined}
          noteTone="wipe"
          testid="resource-row"
        />
      {/each}
    </StatTable>
  </div>
{/snippet}

{#if card !== null}
  <Section title={card.title} meta={copy.meta} testid="sim-resource-card">
    <div class="grid grid-cols-3 gap-x-6 gap-y-5">
      {#each card.tiles as tile (tile.label)}
        <Figure label={tile.label} value={tile.value} size="sm" testid="resource-tile" />
      {/each}
    </div>
    {#if card.lists !== null}
      <div class="grid grid-cols-1 gap-x-[var(--gap)] pt-2 lg:grid-cols-2">
        {@render list(copy.gainedFrom, card.lists.gainedTotal, card.lists.gained, 'resource-gained')}
        {@render list(copy.spentOn, card.lists.spentTotal, card.lists.spent, 'resource-spent')}
      </div>
    {:else}
      <p class="text-muted pt-3 text-[14px]" data-testid="resource-needs-newer">
        {copy.needsNewerRun(card.title.toLowerCase())}
      </p>
    {/if}
  </Section>
{/if}
