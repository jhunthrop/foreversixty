<!-- web/src/components/sim/UnsimulatedSpecCard.svelte -->
<!-- One of /sim/specs' healer and tank specs (mock design/mocks/ws_simtools.py `web-sim-specs`): the
     class crest and the spec's name in its class colour, a link to its guide, set inline with the
     others on one wrapping line under their heading (h-audit 2026-11-09: eight rows of "Guide ›"
     said nothing a link line does not). No state tag, engine stamp or rotation. -->
<script lang="ts">
  import { classCrestSrc } from '../../lib/class-crest';
  import { classSpecPath } from '../../lib/sim/gear-card';
  import { classOfSpec, specLabel } from '../../lib/sim/spec-label';
  import type { Spec } from '../../lib/sim/specs';
  import { classColorVar } from '../../lib/report/format';
  import { simCopy } from '../../lib/sim/copy';

  let { spec }: { spec: Spec } = $props();
</script>

<span id={spec.spec} class="inline-flex min-h-11 items-center" data-testid={`spec-unsimulated-${spec.spec}`}>
  <a
    href={`/guides/${classSpecPath(spec.spec)}`}
    class="inline-flex items-center gap-2 text-[15px] font-bold hover:underline"
    style:color={classColorVar(classOfSpec(spec.spec))}
    data-testid={`spec-guide-${spec.spec}`}
    ><img
      src={classCrestSrc(classOfSpec(spec.spec))}
      alt=""
      width="20"
      height="20"
      class="rounded-full"
    />{specLabel(spec.spec)}<span class="text-muted text-[13px] font-normal"
      >{simCopy.specsRole[spec.role]}</span
    ></a
  >
</span>
