<!-- web/src/components/guild/dense/CompGaps.svelte -->
<!-- The next raid's seats per raid group, read only (dense spec 5.2, the member's rail): Gap first
     (Late counts as coming, the officer's Wednesday rule), then Need, In and Late. -->
<script lang="ts">
  import type { GuildRaidView } from '../../../lib/guild/api';
  import { compGaps, type CompGapRow } from '../../../lib/guild/dense/week-board';
  import { GROUP_WORD } from '../../../lib/guild/dense/roster-board';
  import Section from '../../ui/Section.svelte';
  import StatRow from '../../ui/StatRow.svelte';
  import StatTable from '../../ui/StatTable.svelte';

  let { raid, title = "Wednesday's comp" }: { raid: GuildRaidView; title?: string } = $props();

  const columns = [
    { id: 'gap', label: 'Gap' },
    { id: 'need', label: 'Need' },
    { id: 'in', label: 'In' },
    { id: 'late', label: 'Late' },
  ];
  const rows = $derived(compGaps(raid));

  function figures(row: CompGapRow) {
    return [
      row.gap > 0
        ? { value: `${row.gap} short`, tone: 'wipe' as const }
        : { value: '✓', tone: 'kill' as const },
      { value: String(row.need), muted: true },
      { value: String(row.in) },
      row.late > 0 ? { value: String(row.late) } : { value: '–', muted: true },
    ];
  }
</script>

{#if rows.length > 0}
  <Section {title} panel flush testid="guild-overview-comp">
    <StatTable {columns} nameHead="Role" bar={false} dense label="Seats against the target">
      {#each rows as row (row.group)}
        <StatRow bar={false} figures={figures(row)}>
          {#snippet name()}{GROUP_WORD[row.group]}{/snippet}
        </StatRow>
      {/each}
    </StatTable>
  </Section>
{/if}
