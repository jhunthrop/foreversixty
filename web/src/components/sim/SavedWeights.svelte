<!-- web/src/components/sim/SavedWeights.svelte -->
<!-- A saved stat-weights result at /sim/<id>: read-only. The caution travels with the
     numbers -- a shared link is exactly where someone meets these weights without having
     read the page that made them (design 7) -- so the same warning that opens /sim/weights
     opens this page too, in the same words.

     The table is the one the live tool draws (WeightsScaleTable.svelte over weights.ts's
     `toolScaleView`): the published weights' SimulationCraft form, no Pawn string (owner,
     2026-09-30, round 11). -->
<script lang="ts">
  import type { WeightsResult } from '../../lib/sim/bulk-types';
  import { bulkCopy } from '../../lib/sim/copy';
  import { toolScaleView } from '../../lib/sim/weights';
  import WeightsScaleTable from './WeightsScaleTable.svelte';

  let { result }: { result: WeightsResult } = $props();

  const view = $derived(
    toolScaleView(result.weights, result.request.spec, result.request.weights?.reference ?? ''),
  );
</script>

<section
  class="border-line-warm rounded-panel mx-[18px] border p-4 md:mx-0"
  data-testid="sim-weights-warning"
>
  <p class="text-strong text-[14px]">{bulkCopy.weightsWarning}</p>
  <a class="text-nav underline" href="/sim/gear">{bulkCopy.weightsWarningLink}</a>
</section>

<section class="mx-[18px] flex flex-col gap-3 md:mx-0" data-testid="sim-weights">
  <!-- Task 8, sub-item 1: the same caveat as the live page (StatWeights.svelte), in the
       same words, right before the numbers it caveats -- see that file's own comment. -->
  <p class="text-muted text-[12px]" data-testid="sim-weights-error-caveat">
    {bulkCopy.weightsErrorCaveat(result.weights.map((row) => row.stat))}
  </p>
  <WeightsScaleTable {view} />
</section>
