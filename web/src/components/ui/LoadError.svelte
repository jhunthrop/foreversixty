<!-- web/src/components/ui/LoadError.svelte -->
<!-- The one error state (component audit 2026-11-10 cluster 11): the message and a Try again
     button (the secondary Button, small) that re-fires the same request in place, never "reload
     the page". `line` (the default) sits in the flow where the data would; `panel` gives a region
     that failed whole a strong title over the message, with the button under it. No `onRetry`:
     retrying cannot help, so there is no button. -->
<script lang="ts">
  import { uiCopy } from '../../lib/ui/copy';
  import Button from './Button.svelte';

  let {
    message,
    title,
    onRetry = undefined,
    variant = 'line',
    retrying = false,
    testid = 'load-error',
    retryAriaLabel = undefined,
  }: {
    message: string;
    /** `panel` only: the strong line over the message ("Your guilds did not load."). */
    title?: string;
    onRetry?: () => void;
    variant?: 'line' | 'panel';
    /** The retry is in flight: the button turns busy. */
    retrying?: boolean;
    testid?: string;
    /** The retry's accessible name where a second LoadError can be on screen at once (the planner's
     *  header and tree), so the two buttons differ; elsewhere the visible "Try again" names it. */
    retryAriaLabel?: string;
  } = $props();
</script>

{#snippet retry()}
  {#if onRetry !== undefined}
    <Button
      size="sm"
      onclick={onRetry}
      label={retryAriaLabel}
      busy={retrying}
      busyLabel={uiCopy.loading}
      testid="{testid}-retry">{uiCopy.retry}</Button
    >
  {/if}
{/snippet}

{#if variant === 'panel'}
  <div class="flex flex-col items-start gap-2 py-1" role="alert" data-testid={testid}>
    {#if title !== undefined}<p class="text-strong font-semibold">{title}</p>{/if}
    <p class="text-muted text-[14px] leading-5">{message}</p>
    <div class="mt-1">{@render retry()}</div>
  </div>
{:else}
  <div class="flex min-h-11 flex-wrap items-center gap-3 text-[14px]" role="alert" data-testid={testid}>
    <span class="text-muted">{message}</span>
    {@render retry()}
  </div>
{/if}
