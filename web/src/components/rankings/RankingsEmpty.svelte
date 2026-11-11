<!-- web/src/components/rankings/RankingsEmpty.svelte -->
<!-- A board, or the picker, with nothing ranked yet (spec 2026-09-25 section 6, owner 2026-11-09:
     "every empty state designed and honest"): the shared designed empty (ui/EmptyState with `fill`)
     with the board's one action, Upload a log, and on a narrowed board the wider board as the second
     link; anything passed as children (the bosses ranked elsewhere) follows under a hairline. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { rankingsEmptyCopy } from '../../lib/rankings/copy';
  import EmptyState from '../ui/EmptyState.svelte';

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

<div class="px-5 pt-1 pb-5 max-md:px-[14px]">
  <EmptyState
    message={line}
    {fill}
    action={{ label: rankingsEmptyCopy.action, href: rankingsEmptyCopy.href }}
    secondary={widen}
    {testid}
    {children}
  />
</div>
