<!-- web/src/components/sim/SavedSimLoadFailed.svelte -->
<!-- A saved sim the server did not answer for (spec 2026-11-04-sim-results 4.4): one card, what
     happened, and a retry that re-fires the same request. Nothing else renders. -->
<script lang="ts">
  import { pageErrorCopy as copy } from '../../lib/sim/result-copy';
  import { button } from '../../lib/ui/button';
  import Card from '../ui/Card.svelte';

  let { onRetry }: { onRetry: () => void } = $props();
</script>

<div class="pt-[22px] md:pt-8">
  <Card glyph="alert" label={copy.label} title={copy.title} class="mx-[18px] md:mx-0" testid="sim-saved-error">
    <p class="text-text text-[15px]">{copy.body}</p>
    <button
      type="button"
      class={`${button('secondary')} w-fit`}
      onclick={onRetry}
      data-testid="sim-saved-error-retry"
    >
      {copy.retry}
    </button>
  </Card>
</div>
