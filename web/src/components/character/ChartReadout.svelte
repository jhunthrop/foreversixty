<!-- web/src/components/character/ChartReadout.svelte -->
<!-- The History section's readout (experience spec 2026-10-24-character-history 3.3 and 3.4):
     the selected point's number, its word, its date line, and the earlier and later step
     buttons. It is the section's primary element and doubles as the tooltip, so nothing
     floats over the chart. -->
<script lang="ts">
  import Icon from '../ui/Icon.svelte';
  import { ICON_CHEVRON_RIGHT } from '../../lib/ui/ui-icons';
  import type { Readout } from '../../lib/character/history-chart';

  let {
    readout,
    liveText,
    earlierLabel,
    laterLabel,
    canEarlier,
    canLater,
    onstep,
    testid,
  }: {
    readout: Readout;
    /** The committed point's readout as one line, for screen readers (a hover preview is not announced). */
    liveText: string;
    earlierLabel: string;
    laterLabel: string;
    canEarlier: boolean;
    canLater: boolean;
    onstep: (move: 'earlier' | 'later') => void;
    testid: string;
  } = $props();

  /** A step is an icon button with no box (the flat system): the row hover ground and gold glyph on
   *  hover, a gold focus ring, a muted glyph at half strength when there is nothing further. */
  const STEP =
    'rounded-control inline-flex size-11 items-center justify-center focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-gold';
  const ENABLED = 'text-text hover:bg-row-hover hover:text-gold-hover active:text-gold-hover';
  const DISABLED = 'text-muted/50 cursor-default';
</script>

{#snippet chevron(direction: 'left' | 'right')}
  <Icon icon={ICON_CHEVRON_RIGHT} size={16} class={direction === 'left' ? 'rotate-180' : ''} />
{/snippet}

<div class="flex min-h-14 items-center justify-between gap-3" data-testid={testid}>
  <div class="flex min-w-0 flex-col">
    <p
      class="flex items-baseline gap-2"
      style={readout.color === null ? undefined : `color: ${readout.color}`}
    >
      <span class="text-[28px] leading-8 font-bold tabular-nums" data-testid="history-readout-value"
        >{readout.value}</span
      >
      {#if readout.word !== ''}
        <span class="text-[15px] font-semibold" data-testid="history-readout-word">{readout.word}</span>
      {/if}
    </p>
    <p class="text-small text-muted tabular-nums" data-testid="history-readout-detail">
      {readout.detail}
    </p>
  </div>
  <div class="flex shrink-0 gap-2">
    <button
      type="button"
      class="{STEP} {canEarlier ? ENABLED : DISABLED}"
      aria-label={earlierLabel}
      aria-disabled={canEarlier ? undefined : 'true'}
      data-testid="history-step-earlier"
      onclick={() => canEarlier && onstep('earlier')}
    >
      {@render chevron('left')}
    </button>
    <button
      type="button"
      class="{STEP} {canLater ? ENABLED : DISABLED}"
      aria-label={laterLabel}
      aria-disabled={canLater ? undefined : 'true'}
      data-testid="history-step-later"
      onclick={() => canLater && onstep('later')}
    >
      {@render chevron('right')}
    </button>
  </div>
  <span class="sr-only" aria-live="polite">{liveText}</span>
</div>
