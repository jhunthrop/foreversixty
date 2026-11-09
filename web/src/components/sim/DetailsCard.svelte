<!-- web/src/components/sim/DetailsCard.svelte -->
<!-- Design 5.1; sim spec 2026-11-06 §4.5 (This run as Facts). Five facts about the run itself, beside the results rather than mixed into
     them: none of them is about the character, and all five are what someone asks when
     they doubt the number. The engine version links to /sim/specs, which is where the
     honesty about each spec lives. -->
<script lang="ts">
  import { simCopy } from '../../lib/sim/copy';
  import { percentLabel, runDetails } from '../../lib/sim/details';
  import type { Snippet } from 'svelte';
  import type { SimResult } from '../../lib/sim/types';
  import { engineLabel } from '../../lib/sim/version';
  import type { Fact } from '../../lib/ui/facts';
  import Facts from '../ui/Facts.svelte';
  import Section from '../ui/Section.svelte';

  let {
    result,
    children,
  }: {
    result: SimResult;
    /** The section's small actions under the facts (Copy for the game). */
    children?: Snippet;
  } = $props();

  const details = $derived(runDetails(result));
  const facts = $derived<Fact[]>([
    {
      id: 'margin',
      testid: 'sim-details-margin',
      label: simCopy.detailsMargin,
      value: simCopy.detailsMarginValue(details.bandDps, percentLabel(details.errorPercent)),
    },
    {
      id: 'iterations',
      testid: 'sim-details-iterations',
      label: simCopy.detailsIterations,
      value: details.iterations.toLocaleString('en-US'),
    },
    {
      id: 'processing',
      testid: 'sim-details-processing',
      label: simCopy.detailsProcessing,
      value: `${(details.processingMs / 1000).toFixed(1)} s`,
    },
    {
      id: 'lane',
      testid: 'sim-details-lane',
      label: simCopy.detailsLane,
      value: details.lane === 'server' ? simCopy.detailsLaneServer : simCopy.detailsLaneBrowser,
    },
    { id: 'engine', label: simCopy.detailsEngine, content: engineLink },
  ]);
</script>

{#snippet engineLink()}
  <a
    class="text-muted tabular font-mono text-[13px] font-normal"
    href="/sim/specs"
    data-testid="sim-details-engine"
  >
    {engineLabel(details.engineVersion)}
  </a>
{/snippet}

<Section title={simCopy.details} panel testid="sim-details-card">
  <Facts items={facts} />
  {#if details.hitCeiling}
    <p class="text-muted pt-3 text-[13px] leading-5" data-testid="sim-details-ceiling">
      {simCopy.targetErrorCeiling(result.request.iterations.toLocaleString('en-US'))}
    </p>
  {/if}
  {#if children}<div class="flex flex-wrap gap-2 pt-5">{@render children()}</div>{/if}
</Section>
