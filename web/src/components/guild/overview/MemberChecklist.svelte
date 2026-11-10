<!-- web/src/components/guild/overview/MemberChecklist.svelte -->
<!-- Before the raid (web layout spec 2026-11-05 §7.1, owner decision D12): the member's own checks,
     private to them, each failing one with its fix. Only tracked checks are drawn (7.4). -->
<script lang="ts">
  import type { ChecklistRow } from '../../../lib/guild/overview-role-view';
  import Checklist from '../../ui/Checklist.svelte';
  import Section from '../../ui/Section.svelte';

  let { rows }: { rows: readonly ChecklistRow[] } = $props();
</script>

<Section panel title="Before the raid" testid="guild-overview-checklist">
  <Checklist items={rows} testid="guild-checklist" />
</Section>
