<!-- web/src/components/ui/Skeleton.svelte -->
<!-- The loading state every island shows while its data is in flight: shimmer lines that
     reserve the ready content's height, so the reveal changes nothing but opacity. `lines`
     draws that many placeholder rows at the given `rowHeight`; `minHeight` (a Tailwind
     min-h class, e.g. "min-h-[320px]") reserves more than the rows when the ready view is
     taller. Screen readers hear one "Loading" and the region is aria-busy. -->
<script lang="ts">
  import { uiCopy } from '../../lib/ui/copy';

  let {
    lines = 3,
    rowHeight = 'h-4',
    minHeight = '',
    label = uiCopy.loading,
    testid = 'skeleton',
  }: { lines?: number; rowHeight?: string; minHeight?: string; label?: string; testid?: string } = $props();

  // Widths cycle so a stack of rows reads as text, not bars; fixed so SSR and client agree.
  const WIDTHS = ['w-3/4', 'w-1/2', 'w-5/6', 'w-2/3'] as const;
</script>

<div class="flex flex-col gap-3 {minHeight}" role="status" aria-busy="true" data-testid={testid}>
  <span class="sr-only">{label}</span>
  {#each Array.from({ length: lines }, (_, index) => index) as index (index)}
    <span class="skeleton-block {rowHeight} {WIDTHS[index % WIDTHS.length]}" aria-hidden="true"></span>
  {/each}
</div>
