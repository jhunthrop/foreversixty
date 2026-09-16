<!-- web/src/components/report/ActorTable.svelte -->
<script lang="ts">
  import type { Placement } from '../../lib/report/percentile';
  import { formatAmount, formatPerSecond, formatDuration } from '../../lib/report/format';
  import type { Actor } from '../../lib/report/types';
  import CopyCsv from './CopyCsv.svelte';
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
    healing = false,
    splitUnavailable = false,
    splitFilter = '',
    absent = [],
    windowIsWhole = false,
    deadAt = new Map<string, number>(),
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
    /** The Healing tab: the abilities table's last column is what each heal did not need. */
    healing?: boolean;
    /** Over the night under a target or boss filter the per-ability split cannot be measured; rows say so. */
    splitUnavailable?: boolean;
    /** The filter the night cannot split by, in words, for the row's note. */
    splitFilter?: string;
    /** Players with no row in this window, and when they died if they were dead. */
    absent?: { name: string; deadSince: number | null }[];
    windowIsWhole?: boolean;
    /** Players dead at the window's end, by GUID, with when they died. */
    deadAt?: ReadonlyMap<string, number>;
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
    const prorated = approximate && actors.some((actor) => actor.mitigated === undefined);
    // Alphabetical, not by count: the same six numbers must read in the same order under
    // every filter, or a reader compares the wrong pairs.
    const byType = [...avoided.entries()]
      .sort((a, b) => a[0].localeCompare(b[0]))
      .map(([type, count]) => `${prorated ? '~' : ''}${count} ${type.toLowerCase()}`)
      .join(', ');
    return { absorbed, blocked, hits, byType, prorated };
  });

  /** The table as it stands, for a spreadsheet: name, share, amount, per second, active. */
  function csvLines(): string[][] {
    // The healing table's second number is its overheal: the export carries it too.
    const healing = actors.some((actor) => actor.overheal !== undefined);
    const lines = [
      [
        'Rank',
        'Name',
        'Share',
        'Amount',
        ...(healing ? ['Overheal', 'Overheal %'] : []),
        'Per second',
        'Active %',
      ],
    ];
    actors.forEach((actor, index) => {
      // Blank where the screen is blank: a shield healer's absorbs have no overheal figure,
      // and "0" would read as a heal with nothing wasted.
      const overheal = actor.overheal ?? 0;
      const gross = actor.effective + overheal;
      const overhealCells =
        actor.overheal === undefined
          ? ['', '']
          : [String(overheal), gross === 0 ? '0' : ((overheal / gross) * 100).toFixed(1)];
      lines.push([
        String(index + 1),
        actor.name,
        total === 0 ? '0' : ((actor.effective / total) * 100).toFixed(2),
        String(actor.effective),
        ...(healing ? overhealCells : []),
        (actor.time_ms ?? durationMs) === 0
          ? '0'
          : (actor.effective / ((actor.time_ms ?? durationMs) / 1000)).toFixed(1),
        (actor.time_ms ?? durationMs) === 0
          ? '0'
          : Math.min((actor.active_ms / (actor.time_ms ?? durationMs)) * 100, 100).toFixed(1),
      ]);
    });
    return lines;
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
        title="Percentile among ranked kills of the same boss by this spec: damage here, healing on the Healing tab. Reads wipe on a wipe, – while nothing of this spec is ranked on this boss yet, … while the rankings are being asked, ? when they could not be reached; empty on Damage Taken and over the whole night."
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
          {splitFilter}
          {healing}
          deadSince={deadAt.get(actor.guid) ?? null}
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
      <span
        class="tabular text-right font-mono"
        title={amountApproximate ? 'Prorated with its rows, the same way' : undefined}
        >{amountApproximate ? '~' : ''}{formatAmount(total)}</span
      >
      <span
        class="text-muted tabular hidden text-right font-mono text-[13px] md:inline"
        title="Over the whole window. A row divides by its player's own time in it, so a player who died reads higher per second than the total"
        >{formatPerSecond(total, durationMs)}</span
      >
      <!-- No figure: the summary keeps each player's active time, not the raid's, and the
           largest of them said nothing true about the table. -->
      <span class="hidden md:inline"></span>
    </div>
    {#if absent.length > 0}
      <p class="text-muted border-line-soft border-t px-2 py-2 text-[12px]" data-testid="actor-absent">
        No {mitigation ? 'target' : 'source'} row of their own {windowIsWhole
          ? 'under this filter'
          : 'in this window'}: {absent
          .map((entry) =>
            entry.deadSince === null
              ? entry.name
              : `${entry.name} (dead since ${formatDuration(entry.deadSince)})`,
          )
          .join(', ')}.
      </p>
    {/if}
    {#if mitigation && mitigated.prorated}
      <p
        class="text-muted border-line-soft border-t px-2 py-2 text-[12px]"
        data-testid="actor-mitigated"
        title="An absorb, a block and an avoided hit have no damage share to prorate by, so a filter cannot split them from the summary; a pull measures them from that fight's events"
      >
        Mitigated: not split under this filter over this window. Open a pull to measure what was absorbed,
        blocked and avoided under it.
      </p>
    {:else if mitigation && (mitigated.absorbed > 0 || mitigated.blocked > 0 || mitigated.hits > 0)}
      <p
        class="text-muted border-line-soft border-t px-2 py-2 text-[12px]"
        data-testid="actor-mitigated"
        title="What did not land, measured from the fight’s events: absorbed by shields, blocked, and hits avoided outright"
      >
        Mitigated: <span class="tabular font-mono">{formatAmount(mitigated.absorbed)}</span>
        absorbed ·
        <span class="tabular font-mono">{formatAmount(mitigated.blocked)}</span>
        blocked ·
        <span class="tabular font-mono">{mitigated.hits}</span> hits avoided{mitigated.hits > 0
          ? ` (${mitigated.byType})`
          : ''}
      </p>
    {/if}
    <CopyCsv lines={csvLines} />
  </div>
{/if}
