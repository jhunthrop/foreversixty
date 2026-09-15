<!-- web/src/components/report/ActorTable.svelte -->
<script lang="ts">
  import type { Placement } from '../../lib/report/percentile';
  import { formatAmount, formatPerSecond, formatDuration } from '../../lib/report/format';
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
    mitigation = false,
    splitUnavailable = false,
    absent = [],
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
    /** Sum what did not land (absorbed, blocked, avoided) under the total: the Damage Taken table's headline. */
    mitigation?: boolean;
    /** Over the night under a target or boss filter the per-ability split cannot be measured; rows say so. */
    splitUnavailable?: boolean;
    /** Players with no row in this window, and when they died if they were dead. */
    absent?: { name: string; deadSince: number | null }[];
    measure?: (actor: Actor) => Promise<ExactSplit>;
  } = $props();

  const peak = $derived(actors.reduce((highest, actor) => Math.max(highest, actor.effective), 0));
  const total = $derived(actors.reduce((sum, actor) => sum + actor.effective, 0));
  /** Absorbed and blocked amounts and the avoided hits by type, over every row's abilities. */
  const mitigated = $derived.by(() => {
    let absorbed = 0;
    let blocked = 0;
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const avoided = new Map<string, number>();
    for (const actor of actors) {
      // A measured row carries its own figures; otherwise the per-ability ones, prorated like them.
      if (actor.mitigated !== undefined) {
        absorbed += actor.mitigated.absorbed;
        blocked += actor.mitigated.blocked;
        for (const [type, count] of Object.entries(actor.mitigated.misses)) {
          avoided.set(type, (avoided.get(type) ?? 0) + count);
        }
        continue;
      }
      for (const ability of actor.abilities) {
        absorbed += ability.absorbed ?? 0;
        blocked += ability.blocked ?? 0;
        for (const [type, count] of Object.entries(ability.misses ?? {})) {
          avoided.set(type, (avoided.get(type) ?? 0) + count);
        }
      }
    }
    const hits = [...avoided.values()].reduce((sum, count) => sum + count, 0);
    const byType = [...avoided.entries()]
      .sort((a, b) => b[1] - a[1])
      .map(([type, count]) => `${count} ${type.toLowerCase()}`)
      .join(', ');
    const prorated = approximate && actors.some((actor) => actor.mitigated === undefined);
    return { absorbed, blocked, hits, byType, prorated };
  });

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
          {splitUnavailable}
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
      <!-- No figure: the summary keeps each player's active time, not the raid's, and the
           largest of them said nothing true about the table. -->
      <span class="hidden md:inline"></span>
    </div>
    {#if absent.length > 0}
      <p class="text-muted border-line-soft border-t px-2 py-2 text-[12px]" data-testid="actor-absent">
        No row in this window: {absent
          .map((entry) =>
            entry.deadSince === null
              ? entry.name
              : `${entry.name} (dead since ${formatDuration(entry.deadSince)})`,
          )
          .join(', ')}.
      </p>
    {/if}
    {#if mitigation && (mitigated.absorbed > 0 || mitigated.blocked > 0 || mitigated.hits > 0)}
      <p
        class="text-muted border-line-soft border-t px-2 py-2 text-[12px]"
        data-testid="actor-mitigated"
        title={mitigated.prorated
          ? 'What did not land, prorated from the whole fight like the split above; a pull measures it exactly'
          : 'What did not land, measured from the fight’s events: absorbed by shields, blocked, and hits avoided outright'}
      >
        Mitigated: <span class="tabular font-mono"
          >{mitigated.prorated ? '~' : ''}{formatAmount(mitigated.absorbed)}</span
        >
        absorbed ·
        <span class="tabular font-mono">{formatAmount(mitigated.blocked)}</span> blocked ·
        <span class="tabular font-mono">{mitigated.hits}</span> hits avoided{mitigated.hits > 0
          ? ` (${mitigated.byType})`
          : ''}
      </p>
    {/if}
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
