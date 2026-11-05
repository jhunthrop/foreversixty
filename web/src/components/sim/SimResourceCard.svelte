<!-- web/src/components/sim/SimResourceCard.svelte -->
<!-- Where the rage, energy or mana came from and went (spec 2026-11-04-sim-results 3.10): three
     tiles, then the two source lists. The timeline's resource lane is the graph; this card has none. -->
<script lang="ts">
  import type { ActionNames } from '../../lib/sim/action-names';
  import { createCardTables } from '../../lib/sim/card-tables.svelte';
  import { resourceCardCopy as copy } from '../../lib/sim/input-cards-copy';
  import { resourceCard, type ResourceRow } from '../../lib/sim/resource-card';
  import type { SimResult } from '../../lib/sim/types';
  import Card from '../ui/Card.svelte';
  import StatTile from '../ui/StatTile.svelte';
  import ActionMark from './ActionMark.svelte';

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
</script>

{#snippet list(heading: string, total: string, rows: readonly ResourceRow[], testid: string)}
  <div class="min-w-0 px-5 pt-3 pb-2" data-testid={testid}>
    <div class="quiet-label border-line-soft flex justify-between border-b pb-2">
      <span>{heading}</span><span class="numeric">{total}</span>
    </div>
    <ul class="flex flex-col">
      {#each rows as row (row.action)}
        <li
          class="border-line-soft grid min-h-10 grid-cols-[minmax(0,1fr)_56px_112px] items-center gap-3 border-b text-[14px] last:border-b-0 md:grid-cols-[minmax(0,1fr)_96px_136px]"
          data-testid="resource-row"
        >
          <ActionMark action={row.action} {tables}>{row.name}</ActionMark>
          <span class="bg-line-soft block h-1.5 rounded-full">
            <span class="bg-gold block h-full rounded-full" style:width={`${row.barPct}%`}></span>
          </span>
          <span class="flex items-baseline justify-end gap-2 whitespace-nowrap">
            <span class="text-strong numeric font-semibold">{row.amount}</span>
            {#if row.lost !== null}<span class="text-ember numeric text-[13px]">{row.lost}</span>{/if}
          </span>
        </li>
      {/each}
    </ul>
  </div>
{/snippet}

{#if card !== null}
  <Card
    glyph="bolt"
    label={copy.label}
    title={card.title}
    body="none"
    class="mx-[18px] md:mx-0"
    testid="sim-resource-card"
  >
    <div class="border-line-soft grid grid-cols-1 gap-3 border-t p-5 md:grid-cols-3">
      {#each card.tiles as tile (tile.label)}
        <StatTile value={tile.value} label={tile.label} testid="resource-tile" />
      {/each}
    </div>
    {#if card.lists !== null}
      <div class="border-line-soft divide-line-soft grid grid-cols-1 divide-y border-t md:grid-cols-2 md:divide-x md:divide-y-0">
        {@render list(copy.gainedFrom, card.lists.gainedTotal, card.lists.gained, 'resource-gained')}
        {@render list(copy.spentOn, card.lists.spentTotal, card.lists.spent, 'resource-spent')}
      </div>
    {:else}
      <p class="text-muted border-line-soft border-t px-5 py-4 text-[14px]" data-testid="resource-needs-newer">
        {copy.needsNewerRun(card.title.toLowerCase())}
      </p>
    {/if}
  </Card>
{/if}
