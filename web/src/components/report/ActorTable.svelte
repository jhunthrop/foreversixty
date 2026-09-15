<!-- web/src/components/report/ActorTable.svelte -->
<script lang="ts">
  import type { Placement } from '../../lib/report/percentile';
  import type { Actor } from '../../lib/report/types';
  import ActorRow from './ActorRow.svelte';

  let {
    actors,
    durationMs,
    metricLabel,
    percentiles = new Map<string, Placement>(),
    parseFallback = '',
    approximate = false,
  }: {
    actors: Actor[];
    durationMs: number;
    metricLabel: string;
    percentiles?: Map<string, Placement>;
    /** What an empty Parse cell shows: '' on trash, 'wipe', or a dash for not ranked yet. */
    parseFallback?: string;
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
      <span title="Rank in this table">#</span>
      <span
        title="Percentile among ranked kills of the same boss by this spec. Empty on a wipe or while nothing is ranked yet."
        >Parse</span
      >
      <span>Name</span>
      <span title="Amount, split by ability. Hover a segment for the ability.">{metricLabel}</span>
      <span class="text-right" title="Total in this window">Amount</span>
      <span class="text-right" title="Amount divided by the window's length">Per sec</span>
      <span class="text-right" title="Share of the window spent casting or attacking">Active</span>
    </div>
    <ul class="flex flex-col">
      {#each actors as actor, index (actor.guid)}
        <ActorRow
          rank={index + 1}
          {actor}
          {peak}
          {durationMs}
          {approximate}
          {parseFallback}
          percentile={percentiles.get(actor.guid) ?? null}
        />
      {/each}
    </ul>
  </div>
{/if}
