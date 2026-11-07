<!-- web/src/components/sim/SavedSimLoadFailed.svelte -->
<!-- A saved sim the server did not answer for (spec 2026-11-04-sim-results 4.4): the page's one
     Plate (web layout spec 2026-11-05 §3A), what happened, and a retry that re-fires the same
     request. Nothing else renders. -->
<script lang="ts">
  import { pageErrorCopy as copy } from '../../lib/sim/result-copy';
  import { button } from '../../lib/ui/button';
  import Plate from '../ui/Plate.svelte';

  let { onRetry }: { onRetry: () => void } = $props();
</script>

<div class="px-[18px] pt-[22px] md:px-0 md:pt-8">
  <Plate flush={false} label={copy.label} testid="sim-saved-error">
    <div class="flex flex-col items-start gap-3" role="alert">
      <h2 class="text-strong text-[17px] leading-6 font-bold">{copy.title}</h2>
      <p class="text-text text-[15px]">{copy.body}</p>
      <button type="button" class={button('secondary')} onclick={onRetry} data-testid="sim-saved-error-retry">
        {copy.retry}
      </button>
    </div>
  </Plate>
</div>
