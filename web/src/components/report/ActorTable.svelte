<!-- web/src/components/report/ActorTable.svelte -->
<script lang="ts">
  import type { Placement } from '../../lib/report/percentile';
  import { formatAmount, formatPerSecond, formatPercent } from '../../lib/report/format';
  import type { Actor } from '../../lib/report/types';
  import type { ExactSplit } from '../../lib/report/exact';
  import ActorRow from './ActorRow.svelte';

  let {
    actors,
    durationMs,
    metricLabel,
    percentiles = new Map<string, Placement>(),
    parseFallback = '',
    pairsLabel = 'Targets',
    approximate = false,
    amountApproximate = false,
    measure = undefined,
  }: {
    actors: Actor[];
    durationMs: number;
    metricLabel: string;
    percentiles?: Map<string, Placement>;
    /** What an empty Parse cell shows: '' on trash, 'wipe', or a dash for not ranked yet. */
    parseFallback?: string;
    pairsLabel?: string;
    approximate?: boolean;
    /** True when the amounts are prorated too: a window met by a target or boss filter. */
    amountApproximate?: boolean;
    measure?: (actor: Actor) => Promise<ExactSplit>;
  } = $props();

  const peak = $derived(actors.reduce((highest, actor) => Math.max(highest, actor.effective), 0));
  const total = $derived(actors.reduce((sum, actor) => sum + actor.effective, 0));
  const totalActive = $derived(actors.reduce((sum, actor) => Math.max(sum, actor.active_ms), 0));

  /** The table as it stands, for a spreadsheet: name, share, amount, per second, active. */
  function csv(): string {
    const lines = [['Rank', 'Name', 'Share', 'Amount', 'Per second', 'Active %']];
    actors.forEach((actor, index) => {
      lines.push([
        String(index + 1),
        actor.name,
        total === 0 ? '0' : ((actor.effective / total) * 100).toFixed(2),
        String(actor.effective),
        durationMs === 0 ? '0' : (actor.effective / (durationMs / 1000)).toFixed(1),
        durationMs === 0 ? '0' : ((actor.active_ms / durationMs) * 100).toFixed(1),
      ]);
    });
    return lines.map((row) => row.map((cell) => `"${cell.replaceAll('"', '""')}"`).join(',')).join('\n');
  }

  let copied = $state('');
  async function copyCsv(): Promise<void> {
    try {
      await navigator.clipboard.writeText(csv());
      copied = 'Copied';
    } catch {
      copied = 'Copy failed';
    }
    setTimeout(() => (copied = ''), 2000);
  }
</script>

{#if actors.length === 0}
  <p class="text-muted text-[14px]" data-testid="table-empty">Nothing in this window.</p>
{:else}
  <div class="flex flex-col" data-testid="actor-table">
    <div
      class="text-muted label hidden grid-cols-[28px_40px_minmax(120px,1.4fr)_52px_minmax(0,3fr)_92px_80px_64px] gap-x-3 px-2 pb-1 md:grid"
    >
      <span title="Rank in this table">#</span>
      <span
        title="Percentile among ranked kills of the same boss by this spec: damage here, healing on the Healing tab. Empty on a wipe, on Damage Taken, or while nothing is ranked yet."
        >Parse</span
      >
      <span>Name</span>
      <span class="text-right" title="Share of this table's total">Share</span>
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
          {amountApproximate}
          {parseFallback}
          {pairsLabel}
          {measure}
          percentile={percentiles.get(actor.guid) ?? null}
          share={total === 0 ? 0 : (actor.effective / total) * 100}
        />
      {/each}
    </ul>
    <div
      class="border-line-soft grid min-h-11 grid-cols-[minmax(0,1fr)_auto] items-center gap-x-3 border-t px-2 py-2 text-[14px] md:grid-cols-[28px_40px_minmax(120px,1.4fr)_52px_minmax(0,3fr)_92px_80px_64px]"
      data-testid="actor-total"
    >
      <span class="text-muted label md:col-span-3">Total</span>
      <span class="tabular hidden text-right font-mono text-[12px] md:inline">100%</span>
      <span class="hidden md:inline"></span>
      <span class="tabular text-right font-mono">{formatAmount(total)}</span>
      <span class="text-muted tabular hidden text-right font-mono text-[13px] md:inline"
        >{formatPerSecond(total, durationMs)}</span
      >
      <span class="text-muted tabular hidden text-right font-mono text-[13px] md:inline"
        >{durationMs === 0 ? '' : formatPercent((totalActive / durationMs) * 100)}</span
      >
    </div>
    <div class="flex justify-end">
      <button
        type="button"
        class="text-nav inline-flex min-h-11 items-center text-[12px] font-bold tracking-[0.06em] uppercase md:min-h-9"
        title="Copy this table as CSV"
        data-testid="copy-csv"
        onclick={() => void copyCsv()}>{copied === '' ? 'Copy CSV' : copied}</button
      >
    </div>
  </div>
{/if}
