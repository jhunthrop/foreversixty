<!-- web/src/components/sim/tools/TopGear.svelte -->
<!-- Placeholder. Task 12 replaces this file with the real slot grid, candidates, locks and
     enchants (and Tasks 13-17 add sections to it); this task only needs the module to exist
     and resolve so ToolsView.svelte's per-tool lazy chunk builds. The prop shape below is
     pinned to Task 12's own brief ({ store: BulkStore; me: Me | null }) so the real
     component drops in without ToolsView changing. -->
<script lang="ts">
  import type { Me } from '../../../lib/account/api';
  import type { BulkStore } from '../../../lib/sim/bulk-store.svelte';
  import { bulkCopy } from '../../../lib/sim/copy';

  let { store, me }: { store: BulkStore; me: Me | null } = $props();
</script>

<p class="text-muted px-[18px] text-[14px] md:px-0" data-testid="sim-tool-placeholder">
  {store.tool === 'talents' ? bulkCopy.talentsTitle : bulkCopy.gearTitle}
  {#if me !== null}({me.user.email}){/if}
</p>
