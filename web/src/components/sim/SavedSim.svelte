<!-- web/src/components/sim/SavedSim.svelte -->
<!-- /sim/<sim_id>: a `run` result is the one-page view (RunReport.svelte, spec
     2026-11-04-sim-results); a bulk or weights result keeps its own (SavedKind.svelte). Kind is
     derived from the stored request, never sent and never stored twice (contract 1.1). Opening
     a result marks that server run seen, once (experience spec 2026-10-24 3.6). -->
<script lang="ts">
  import { onMount } from 'svelte';
  import { requestKind } from '../../lib/sim/bulk-types';
  import { markSimSeenOnOpen } from '../../lib/sim/mark-seen-on-open';
  import type { SimResult } from '../../lib/sim/types';
  import SavedKind from './SavedKind.svelte';
  import RunReport from './RunReport.svelte';

  let { result, onrerun, simId = '' }: { result: SimResult; onrerun: () => void; simId?: string } = $props();

  onMount(() => {
    void markSimSeenOnOpen(simId !== '' ? simId : (result.sim_id ?? ''));
  });

  const kind = $derived(requestKind(result.request));
</script>

{#if kind === 'run'}
  <RunReport {result} {onrerun} {simId} />
{:else}
  <div class="flex flex-col gap-[22px] pt-[22px] md:gap-8 md:pt-8">
    <SavedKind {result} {onrerun} />
  </div>
{/if}
