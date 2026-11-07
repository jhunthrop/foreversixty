<!-- web/src/components/ui/Divided.svelte -->
<!-- A block under a hairline inside an open Section (setup spec 2026-11-07 section 4): the
     hand install under The addon, the pairing block under The companion. A 1 px --color-line-soft
     rule, 12 px above the content, the content in a column. No box. -->
<script lang="ts">
  import type { Snippet } from 'svelte';

  let { gap = 6, testid, children }: { gap?: number; testid?: string; children: Snippet } = $props();
</script>

<div class="ui-divided" style:gap="{gap}px" data-testid={testid}>{@render children()}</div>

<style>
  .ui-divided {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    padding-top: 12px;
    border-top: 1px solid var(--color-line-soft);
  }
</style>
