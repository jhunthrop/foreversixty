<!-- web/src/components/sim/UnsimulatedSpecCard.svelte -->
<!-- /sim/specs' healer and tank rows (mock design/mocks/ws_simtools.py `web-sim-specs`), open under
     their own section below the simulated ones: the spec in its class colour and one line, with no
     state tag, engine stamp or rotation, since none of that will exist for a healer or tank. -->
<script lang="ts">
  import { simCopy } from '../../lib/sim/copy';
  import { classOfSpec, specLabel } from '../../lib/sim/spec-label';
  import type { Spec } from '../../lib/sim/specs';
  import { classTone } from '../../lib/ui/primitives';
  import Row from '../ui/Row.svelte';

  let { spec }: { spec: Spec } = $props();
</script>

<Row
  id={spec.spec}
  title={specLabel(spec.spec)}
  titleTone={classTone(classOfSpec(spec.spec))}
  sub={simCopy.specsUnsimulatedBody}
  open
  testid={`spec-unsimulated-${spec.spec}`}
/>
