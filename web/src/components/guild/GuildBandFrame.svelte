<!-- web/src/components/guild/GuildBandFrame.svelte -->
<!-- The raid page's head frame (raids spec 3.2, site alignment 2026-10-29 D1): the head lines and the
     tab strip on the page ground, in one column. There is no band, faction wash or watermark: the
     TabStrip's edge fade runs to the page ground. -->
<script lang="ts">
  import type { Snippet } from 'svelte';

  let {
    minHeight = '',
    testid = 'guild-header-band',
    children,
  }: {
    /** Tailwind min-height classes the head holds while it loads. */
    minHeight?: string;
    testid?: string;
    children: Snippet;
  } = $props();
</script>

<div class="flex flex-col gap-3 {minHeight}" data-testid={testid}>
  {@render children()}
</div>
