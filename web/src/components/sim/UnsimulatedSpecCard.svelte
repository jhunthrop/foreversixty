<!-- web/src/components/sim/UnsimulatedSpecCard.svelte -->
<!-- task-2-brief.md: /sim/specs' 7 healer/tank cards, grouped under their own heading below
     the 20 simulated ones. A distinct component rather than SpecCard.svelte with an
     `unsupported`-shaped row: the brief calls for a visibly different card -- no fidelity
     pill, no engine stamp, no "Not yet" badge link -- and building that as a SpecCard
     branch would mean threading three separate "hide this" props through a component whose
     entire job today is showing exactly those things. -->
<script lang="ts">
  import { classColorVar } from '../../lib/report/format';
  import { simCopy } from '../../lib/sim/copy';
  import { classOfSpec, specLabel } from '../../lib/sim/spec-label';
  import type { Spec } from '../../lib/sim/specs';

  let { spec }: { spec: Spec } = $props();

  const colour = $derived(classColorVar(classOfSpec(spec.spec)));
</script>

<article
  id={spec.spec}
  class="border-line bg-card-top rounded-panel flex flex-col gap-3 border p-4"
  data-testid={`spec-unsimulated-${spec.spec}`}
>
  <h3 class="text-[15px] font-semibold" style={`color: ${colour}`}>{specLabel(spec.spec)}</h3>
  <p class="text-muted text-[13px]">{simCopy.specsUnsimulatedBody}</p>
</article>
