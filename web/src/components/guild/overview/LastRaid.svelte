<!-- web/src/components/guild/overview/LastRaid.svelte -->
<!-- Last raid (web layout spec 2026-11-05 §7.1, demoted): past nights collapse to one Row linking
     the report, with progression as the section's one line. -->
<script lang="ts">
  import { guildLastRaidCopy } from '../../../lib/guild/copy';
  import EmptyState from '../../ui/EmptyState.svelte';
  import FactLine from '../FactLine.svelte';
  import Row from '../../ui/Row.svelte';
  import Section from '../../ui/Section.svelte';

  let {
    report,
    facts,
    progression,
  }: {
    report: { id: string; title: string; zone: string; day: string } | null;
    facts: readonly string[];
    progression: string;
  } = $props();
</script>

<Section panel title={guildLastRaidCopy.title} meta={progression} testid="guild-overview-last-raid">
  {#if report === null}
    <EmptyState
      message={guildLastRaidCopy.empty}
      action={guildLastRaidCopy.emptyAction}
      testid="last-raid-empty"
    />
  {:else}
    <Row title={report.title === '' ? report.zone : report.title} href={`/reports/${report.id}`} open>
      {#snippet subContent()}<FactLine
          facts={[report.day, ...facts]}
          testid="guild-overview-raid-facts"
        />{/snippet}
      {#snippet trail()}<span class="text-gold font-semibold">Report ›</span>{/snippet}
    </Row>
  {/if}
</Section>
