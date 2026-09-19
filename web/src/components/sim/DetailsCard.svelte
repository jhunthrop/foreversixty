<!-- web/src/components/sim/DetailsCard.svelte -->
<!-- Design 5.1. Five facts about the run itself, beside the results rather than mixed into
     them: none of them is about the character, and all five are what someone asks when
     they doubt the number. The engine version links to /sim/specs, which is where the
     honesty about each spec lives. -->
<script lang="ts">
  import { simCopy } from '../../lib/sim/copy';
  import { percentLabel, runDetails } from '../../lib/sim/details';
  import type { SimResult } from '../../lib/sim/types';
  import { engineLabel } from '../../lib/sim/version';

  let { result }: { result: SimResult } = $props();

  const details = $derived(runDetails(result));
  const rows = $derived([
    {
      id: 'margin',
      label: simCopy.detailsMargin,
      value: simCopy.detailsMarginValue(
        details.bandDps.toLocaleString('en-US'),
        percentLabel(details.errorPercent),
      ),
    },
    {
      id: 'iterations',
      label: simCopy.detailsIterations,
      value: details.iterations.toLocaleString('en-US'),
    },
    {
      id: 'processing',
      label: simCopy.detailsProcessing,
      value: `${(details.processingMs / 1000).toFixed(1)} s`,
    },
    {
      id: 'lane',
      label: simCopy.detailsLane,
      value: details.lane === 'server' ? simCopy.detailsLaneServer : simCopy.detailsLaneBrowser,
    },
  ]);
</script>

<section
  class="border-line bg-raised rounded-panel mx-[18px] flex flex-col gap-2 border p-4 md:mx-0"
  data-testid="sim-details-card"
>
  <h2 class="section-title text-[15px]">{simCopy.details}</h2>
  <dl class="grid grid-cols-[minmax(0,1fr)_auto] gap-x-4 gap-y-1 text-[13px]">
    {#each rows as row (row.id)}
      <dt class="text-muted">{row.label}</dt>
      <dd class="tabular text-strong text-right font-mono" data-testid={`sim-details-${row.id}`}>
        {row.value}
      </dd>
    {/each}
    <dt class="text-muted">{simCopy.detailsEngine}</dt>
    <dd class="text-right">
      <a class="tabular font-mono text-[13px]" href="/sim/specs" data-testid="sim-details-engine">
        {engineLabel(details.engineVersion)}
      </a>
    </dd>
  </dl>
  {#if details.hitCeiling}
    <p class="text-muted text-[12px]" data-testid="sim-details-ceiling">
      {simCopy.targetErrorCeiling(result.request.iterations.toLocaleString('en-US'))}
    </p>
  {/if}
</section>
