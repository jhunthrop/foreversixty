<!-- web/src/components/sim/SavedWeights.svelte -->
<!-- A saved stat-weights result at /sim/<id>: read-only. The caution travels with the
     numbers -- a shared link is exactly where someone meets these weights without having
     read the page that made them (design 7) -- so the same warning that opens /sim/weights
     opens this page too, in the same words.

     The table is the one the live tool draws (WeightsScaleTable.svelte over weights.ts's
     `toolScaleView`): the published weights' SimulationCraft form, no Pawn string (owner,
     2026-09-30, round 11), the page's one Plate (sim spec 2026-11-06 §4.5 grammar). -->
<script lang="ts">
  import type { WeightsResult } from '../../lib/sim/bulk-types';
  import { bulkCopy } from '../../lib/sim/copy';
  import { toolScaleView } from '../../lib/sim/weights';
  import Notice from '../ui/Notice.svelte';
  import Plate from '../ui/Plate.svelte';
  import WeightsScaleTable from './WeightsScaleTable.svelte';

  let { result }: { result: WeightsResult } = $props();

  const view = $derived(
    toolScaleView(result.weights, result.request.spec, result.request.weights?.reference ?? ''),
  );
</script>

<div class="flex flex-col gap-3" data-testid="sim-weights">
  <Notice
    text={bulkCopy.weightsWarning}
    action={{ label: bulkCopy.weightsWarningLink, href: '/sim/gear' }}
    testid="sim-weights-warning"
  />
  <!-- Task 8, sub-item 1: the same caveat as the live page (StatWeights.svelte), in the
       same words, right before the numbers it caveats -- see that file's own comment. -->
  <p class="text-muted text-[13px] leading-5" data-testid="sim-weights-error-caveat">
    {bulkCopy.weightsErrorCaveat(result.weights.map((row) => row.stat))}
  </p>
  <Plate flush={false} label={bulkCopy.weightsTitle}>
    <WeightsScaleTable {view} />
  </Plate>
</div>
