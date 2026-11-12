<!-- web/src/components/sim/SavedSimLoadFailed.svelte -->
<!-- A saved sim the server did not answer for (spec 2026-11-04-sim-results 4.4): the page's one
     Plate (web layout spec 2026-11-05 §3A), what happened, and a retry that re-fires the same
     request. Nothing else renders. -->
<script lang="ts">
  import { pageErrorCopy as copy } from '../../lib/sim/result-copy';
  import LoadError from '../ui/LoadError.svelte';
  import Plate from '../ui/Plate.svelte';

  let { onRetry }: { onRetry: () => void } = $props();
</script>

<div class="px-(--page-inset) pt-[22px] md:pt-8">
  <Plate flush={false} label={copy.label} testid="sim-saved-error-plate">
    <LoadError title={copy.title} message={copy.body} variant="panel" {onRetry} testid="sim-saved-error" />
  </Plate>
</div>
