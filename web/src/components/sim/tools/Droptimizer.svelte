<!-- web/src/components/sim/tools/Droptimizer.svelte -->
<!-- Placeholder. Task 18 replaces this file with the real source picker and drop results;
     this task only needs the module to exist and resolve so ToolsView.svelte's per-tool
     lazy chunk builds. The prop shape below is pinned to Task 18's own brief
     ({ store: BulkStore; me: Me | null }) so the real component drops in without
     ToolsView changing. -->
<script lang="ts">
  import type { Me } from '../../../lib/account/api';
  import type { BulkStore } from '../../../lib/sim/bulk-store.svelte';
  import { bulkCopy } from '../../../lib/sim/copy';

  let { store, me }: { store: BulkStore; me: Me | null } = $props();
</script>

<p class="text-muted px-[18px] text-[14px] md:px-0" data-testid="sim-tool-placeholder">
  {bulkCopy.dropsTitle}
  {#if me !== null}({store.character?.class_slug ?? ''}){/if}
</p>
