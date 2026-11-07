<!-- web/src/components/ui/Ordered.svelte -->
<!-- A region's place in Page layout="workbench"'s one-column flow below 1280 (planner spec
     2026-11-08 section 3A): `order` sets its position (the phone order is not the DOM order),
     `span` makes it take both columns from 1024 to 1279. The rules live in Page.svelte. -->
<script lang="ts">
  import type { Snippet } from 'svelte';

  let { order, span = false, children }: { order: number; span?: boolean; children: Snippet } = $props();
</script>

<div data-order style:--o={order} data-span={span ? '' : undefined}>{@render children()}</div>
