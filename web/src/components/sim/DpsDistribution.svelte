<!-- web/src/components/sim/DpsDistribution.svelte -->
<!-- The run's distribution, as twenty bars and four numbers. Deliberately not a chart
     library and not a canvas: it is a shape, the four rows below it are the data, and a
     dependency for twenty divs would cost more than the whole island's budget allows.
     The bars are aria-hidden; a screen reader reads the numbers. -->
<script lang="ts">
  import { simCopy } from '../../lib/sim/copy';
  import { confidenceBand } from '../../lib/sim/estimate';
  import type { Estimate } from '../../lib/sim/types';

  let { estimate, iterationsRun }: { estimate: Estimate; iterationsRun: number } = $props();

  const BUCKETS = 20;

  /** Normal density at the run's own mean and standard deviation, scaled to the tallest bar. */
  const bars = $derived.by(() => {
    const spread = estimate.max - estimate.min;
    if (spread <= 0 || estimate.stddev <= 0) return [] as { height: number; inBand: boolean }[];
    const band = confidenceBand(estimate);
    const raw = Array.from({ length: BUCKETS }, (_, i) => {
      const at = estimate.min + (spread * (i + 0.5)) / BUCKETS;
      const z = (at - estimate.mean) / estimate.stddev;
      return { density: Math.exp(-0.5 * z * z), inBand: Math.abs(at - estimate.mean) <= band };
    });
    const peak = Math.max(...raw.map((bucket) => bucket.density));
    return raw.map((bucket) => ({
      height: Math.round((bucket.density / peak) * 100),
      inBand: bucket.inBand,
    }));
  });

  const rows = $derived([
    [simCopy.distMean, Math.round(estimate.mean).toLocaleString('en-US')],
    [simCopy.distStdDev, Math.round(estimate.stddev).toLocaleString('en-US')],
    [simCopy.distLowest, Math.round(estimate.min).toLocaleString('en-US')],
    [simCopy.distHighest, Math.round(estimate.max).toLocaleString('en-US')],
    [simCopy.distIterations, iterationsRun.toLocaleString('en-US')],
  ]);
</script>

<div class="flex flex-col gap-4 p-4" data-testid="sim-distribution">
  {#if bars.length > 0}
    <div class="flex h-32 items-end gap-[2px]" aria-hidden="true">
      {#each bars as bar, i (i)}
        <span
          class={`flex-1 rounded-t-[2px] ${bar.inBand ? 'bg-gold' : 'bg-line'}`}
          style={`height:${Math.max(2, bar.height)}%`}
        ></span>
      {/each}
    </div>
  {/if}
  <dl class="flex flex-col">
    {#each rows as [key, value] (key)}
      <div class="border-line-soft flex justify-between gap-3 border-b py-1 last:border-b-0">
        <dt class="text-muted text-[13px]">{key}</dt>
        <dd class="tabular text-strong m-0 font-mono text-[13px]">{value}</dd>
      </div>
    {/each}
  </dl>
</div>
