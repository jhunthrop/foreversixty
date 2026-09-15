<!-- web/src/components/report/ActorTable.svelte -->
<script lang="ts">
  import type { Actor } from '../../lib/report/types';
  import ActorRow from './ActorRow.svelte';

  let {
    actors,
    durationMs,
    metricLabel,
    percentiles = new Map<string, number>(),
    approximate = false,
  }: {
    actors: Actor[];
    durationMs: number;
    metricLabel: string;
    percentiles?: Map<string, number>;
    approximate?: boolean;
  } = $props();

  const peak = $derived(actors.reduce((highest, actor) => Math.max(highest, actor.effective), 0));
</script>

{#if actors.length === 0}
  <p class="text-muted text-[14px]" data-testid="table-empty">Nothing in this window.</p>
{:else}
  <div class="flex flex-col" data-testid="actor-table">
    <div
      class="text-muted label hidden grid-cols-[28px_40px_minmax(120px,1.4fr)_minmax(0,3fr)_92px_80px_64px] gap-x-3 px-2 pb-1 md:grid"
    >
      <span>#</span>
      <span>Parse</span>
      <span>Name</span>
      <span>{metricLabel}</span>
      <span class="text-right">Amount</span>
      <span class="text-right">Per sec</span>
      <span class="text-right">Active</span>
    </div>
    <ul class="flex flex-col">
      {#each actors as actor, index (actor.guid)}
        <ActorRow
          rank={index + 1}
          {actor}
          {peak}
          {durationMs}
          {approximate}
          percentile={percentiles.get(actor.guid) ?? null}
        />
      {/each}
    </ul>
  </div>
{/if}
