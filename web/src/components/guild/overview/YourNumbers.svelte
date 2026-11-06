<!-- web/src/components/guild/overview/YourNumbers.svelte -->
<!-- Your numbers (web layout spec 2026-11-05 §7.1, members): plain figures from the roster row, each
     only when the roster carries it. -->
<script lang="ts">
  import type { GuildRosterRow as RosterRow } from '../../../lib/guild/api';
  import Figure from '../../ui/Figure.svelte';
  import Section from '../../ui/Section.svelte';

  let { row }: { row: RosterRow } = $props();
</script>

<Section title="Your numbers" testid="guild-overview-numbers">
  <div class="flex flex-wrap gap-x-10 gap-y-4">
    {#if row.attendance !== undefined && row.attendance.nights > 0}
      <Figure
        value={`${row.attendance.present} of ${row.attendance.nights}`}
        label="Raids attended"
        size="sm"
      />
    {/if}
    {#if row.item_level != null}<Figure
        value={String(row.item_level)}
        label="Your item level"
        size="sm"
      />{/if}
  </div>
</Section>
