<!-- web/src/components/sim/UnsimulatedSpecCard.svelte -->
<!-- /sim/specs' healer and tank rows (mock design/mocks/ws_simtools.py `web-sim-specs`), open under
     their own section below the simulated ones: the spec in its class colour and its guide as the
     trail's action, with no state tag, engine stamp or rotation. The section's heading says why; a line under every
     row saying it again was the same sentence eight times. -->
<script lang="ts">
  import { classSpecPath } from '../../lib/sim/gear-card';
  import { classOfSpec, specLabel } from '../../lib/sim/spec-label';
  import { simCopy } from '../../lib/sim/copy';
  import ActionLink from '../ui/ActionLink.svelte';
  import type { Spec } from '../../lib/sim/specs';
  import { classTone } from '../../lib/ui/primitives';
  import Row from '../ui/Row.svelte';
  import { classCrestSrc } from '../../lib/class-crest';

  let { spec }: { spec: Spec } = $props();
</script>

{#snippet crest()}
  <img src={classCrestSrc(classOfSpec(spec.spec))} alt="" width="28" height="28" class="rounded-full" />
{/snippet}

<Row
  id={spec.spec}
  lead={crest}
  density="slim"
  title={specLabel(spec.spec)}
  titleTone={classTone(classOfSpec(spec.spec))}
  open
  testid={`spec-unsimulated-${spec.spec}`}
>
  {#snippet trail()}
    <ActionLink href={`/guides/${classSpecPath(spec.spec)}`} data-testid={`spec-guide-${spec.spec}`}
      >{simCopy.specsGuide}</ActionLink
    >
  {/snippet}
</Row>
