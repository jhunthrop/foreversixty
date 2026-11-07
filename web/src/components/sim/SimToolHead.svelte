<!-- web/src/components/sim/SimToolHead.svelte -->
<!-- The head of every simulator tool page (/sim/gear, /sim/talents, /sim/drops, /sim/weights,
     /sim/specs; mock design/mocks/ws_simtools.py): PageHeader with the tool's title, the DPS-only
     scope line as its one meta line, the tool's one sentence as the lead, and the tool tabs with
     this tool current. The shells render it statically, outside the island, so it paints with the
     page and the island's mount moves nothing above it. The loaded character's query reaches the
     tabs through lib/sim/tabs.ts `syncTabHrefs`. -->
<script lang="ts">
  import { simCopy } from '../../lib/sim/copy';
  import type { SimTabId } from '../../lib/sim/tabs';
  import PageHeader from '../ui/PageHeader.svelte';
  import SimToolTabs from './SimToolTabs.svelte';

  let { tool, title, lead }: { tool: SimTabId; title: string; lead?: string } = $props();
</script>

<PageHeader {title} {lead} testid="sim-page-head">
  {#snippet meta()}<span data-testid="sim-scope-note">{simCopy.scopeNote}</span>{/snippet}
  {#snippet tabs()}<SimToolTabs active={tool} />{/snippet}
</PageHeader>
