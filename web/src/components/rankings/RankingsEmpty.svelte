<!-- web/src/components/rankings/RankingsEmpty.svelte -->
<!-- A board, or the picker, with nothing ranked yet (spec 2026-09-25 section 6, owner 2026-11-09:
     "every empty state designed and honest"): what is missing in one line, what fills it under it,
     then the one action (Upload a log) and, on a narrowed board, the wider board. Anything the
     caller passes as children (the bosses ranked elsewhere) follows under a hairline. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { rankingsEmptyCopy } from '../../lib/rankings/copy';
  import { LINK } from '../../lib/ui/primitives';
  import Button from '../ui/Button.svelte';
  import Ruled from '../ui/Ruled.svelte';

  let {
    line,
    fill,
    widen,
    testid = 'rankings-empty',
    children,
  }: {
    line: string;
    fill: string;
    /** The wider board, when this one is narrowed to a class or spec. */
    widen?: { label: string; href: string };
    testid?: string;
    children?: Snippet;
  } = $props();
</script>

<div data-testid={testid}>
  <div class="flex flex-col gap-1 px-5 pt-1 pb-5 max-md:px-[14px]">
    <p class="text-strong font-semibold">{line}</p>
    <p class="text-muted text-[14px] leading-5">{fill}</p>
    <div class="mt-3 flex flex-wrap items-center gap-x-5 gap-y-2">
      <Button tone="primary" size="sm" href={rankingsEmptyCopy.href}>{rankingsEmptyCopy.action}</Button>
      {#if widen}<a class={LINK} href={widen.href} data-testid="rankings-empty-widen">{widen.label} ›</a>{/if}
    </div>
  </div>
  {#if children}<Ruled><div>{@render children()}</div></Ruled>{/if}
</div>
