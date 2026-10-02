<!-- web/src/components/ui/LoadError.svelte -->
<!-- The error state every island shows when its fetch fails: the message, and a Try again
     button that re-fires the same request in place. Never "reload the page". -->
<script lang="ts">
  import { uiCopy } from '../../lib/ui/copy';

  let {
    message,
    onRetry = undefined,
    testid = 'load-error',
    retryAriaLabel = undefined,
  }: {
    message: string;
    onRetry?: () => void;
    testid?: string;
    /** Fix round 4 (CI: "strict mode violation", two "Try again" buttons on the same
     *  planner page -- PlannerCharacterCard's own and Planner's own tree-load error share
     *  this component and its default accessible name). Undefined everywhere this
     *  component mounts only once per page (every other caller today), so the visible
     *  "Try again" text stays the accessible name there too -- set only where a second
     *  LoadError can be on screen at once. */
    retryAriaLabel?: string;
  } = $props();
</script>

<div class="flex min-h-11 flex-wrap items-center gap-3 text-[14px]" role="alert" data-testid={testid}>
  <span class="text-muted">{message}</span>
  {#if onRetry !== undefined}
    <button
      type="button"
      class="border-line-warm text-text rounded-control inline-flex min-h-11 items-center border px-3 text-[13px] font-semibold md:min-h-0 md:py-1"
      onclick={onRetry}
      aria-label={retryAriaLabel}
      data-testid="{testid}-retry"
    >
      {uiCopy.retry}
    </button>
  {/if}
</div>
