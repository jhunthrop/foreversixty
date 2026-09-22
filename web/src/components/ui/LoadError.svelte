<!-- web/src/components/ui/LoadError.svelte -->
<!-- The error state every island shows when its fetch fails: the message, and a Try again
     button that re-fires the same request in place. Never "reload the page". -->
<script lang="ts">
  import { uiCopy } from '../../lib/ui/copy';

  let {
    message,
    onRetry = undefined,
    testid = 'load-error',
  }: { message: string; onRetry?: () => void; testid?: string } = $props();
</script>

<div class="flex min-h-11 flex-wrap items-center gap-3 text-[14px]" role="alert" data-testid={testid}>
  <span class="text-muted">{message}</span>
  {#if onRetry !== undefined}
    <button
      type="button"
      class="border-line-warm text-text rounded-control inline-flex min-h-11 items-center border px-3 text-[13px] font-semibold md:min-h-0 md:py-1"
      onclick={onRetry}
      data-testid="{testid}-retry"
    >
      {uiCopy.retry}
    </button>
  {/if}
</div>
