<!-- web/src/components/TierStrip.svelte -->
<!-- The five Forever Score tiers as one strip: each segment a bar in the tier's exact item
     quality colour, the tier's name (rare and epic in their lightened text variants) and its
     range. With `current`, that tier is marked and the others' bars dim; the words never dim,
     so they keep their contrast. Used on the profile's score card and on /ratings (where a
     visually hidden table carries the same facts for assistive tech). Five columns from sm; on a
     phone each tier is a row (swatch, name, range), since five columns at 360px cut the names. -->
<script lang="ts">
  import { scoreCopy } from '../lib/score/copy';
  import { tierStrip } from '../lib/score/score';

  let { current = null }: { current?: string | null } = $props();

  const segments = tierStrip();
</script>

<ol
  class="grid grid-cols-1 gap-1 sm:grid-cols-5"
  aria-label={scoreCopy.tierStripLabel}
  data-testid="tier-strip"
>
  {#each segments as segment (segment.tier)}
    {@const here = segment.tier === current}
    <li
      class={`rounded-control grid min-w-0 grid-cols-[40px_minmax(0,1fr)_auto] items-center gap-x-3 gap-y-1 p-1.5 sm:flex sm:flex-col sm:items-stretch ${here ? 'bg-card-top ring-gold/70 ring-1' : ''}`}
      aria-current={here ? 'true' : undefined}
      data-testid={`tier-strip-${segment.tier}`}
    >
      <span
        class={`block rounded-[2px] ${here ? 'h-2.5' : 'h-1.5'} ${current !== null && !here ? 'opacity-40' : ''}`}
        style={`background-color: ${segment.color}`}
        aria-hidden="true"
      ></span>
      <span class="text-[13px] font-semibold sm:truncate sm:text-[12px]" style={`color: ${segment.textColor}`}
        >{segment.label}</span
      >
      <span class="text-muted tabular font-mono text-[11px] whitespace-nowrap">{segment.range}</span>
      {#if here}<span class="sr-only">{scoreCopy.currentTier}</span>{/if}
    </li>
  {/each}
</ol>
