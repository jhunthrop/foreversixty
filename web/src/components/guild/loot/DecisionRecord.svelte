<!-- web/src/components/guild/loot/DecisionRecord.svelte -->
<!-- The night's decision record (web layout spec 2026-11-05 §7.2): every award of the night, oldest
     first, as DecisionRows. Officers get Export history in the head, which opens the date range and
     the CSV and JSON downloads under it. -->
<script lang="ts">
  import EmptyState from '../../ui/EmptyState.svelte';
  import type { LootDecision } from '../../../lib/guild/loot-council-api';
  import Section from '../../ui/Section.svelte';
  import DecisionRow from './DecisionRow.svelte';
  import LootExport from './LootExport.svelte';

  let {
    decisions,
    build,
    shipsIcon,
    timeZone,
    exportFor = null,
  }: {
    decisions: readonly LootDecision[];
    build: string;
    shipsIcon: (itemId: number) => boolean;
    timeZone?: string;
    /** Officers: the guild whose history downloads. */
    exportFor?: number | null;
  } = $props();

  let exporting = $state(false);
</script>

<Section title="Decision record" panel testid="guild-loot-decisions">
  {#snippet head()}
    {#if exportFor !== null}
      <button
        type="button"
        class="text-gold hover:text-gold-hover focus-visible:outline-gold rounded-field min-h-11 text-[14px] font-semibold focus-visible:outline-2 focus-visible:outline-offset-2 md:min-h-0"
        aria-expanded={exporting ? 'true' : 'false'}
        onclick={() => (exporting = !exporting)}
        data-testid="guild-loot-export-toggle">Export history</button
      >
    {/if}
  {/snippet}
  {#if exporting && exportFor !== null}
    <div class="pb-3"><LootExport guildId={exportFor} /></div>
  {/if}
  {#each decisions as decision (decision.award_id)}
    <DecisionRow {decision} {build} shipsIcon={shipsIcon(decision.item_id)} {timeZone} />
  {:else}
    <EmptyState message="Nothing awarded on this night yet." testid="guild-loot-decisions-empty" />
  {/each}
</Section>
