<!-- web/src/components/report/QueriesRail.svelte -->
<!-- The Queries tab's rail Section (logs spec 2026-11-07 section 6.2): `Columns`, every column a
     query can select from the fight's events.parquet, as literal code. -->
<script lang="ts">
  import { EVENT_COLUMNS } from '../../lib/report/query';
  import Section from '../ui/Section.svelte';

  let { columns = EVENT_COLUMNS }: { columns?: readonly string[] } = $props();
</script>

<Section title="Columns" meta={`${columns.length} in events.parquet`} testid="queries-rail" panel>
  <ul class="text-text flex flex-wrap gap-x-3 gap-y-1 font-mono text-[13px] leading-5" aria-label="Columns">
    {#each columns as column (column)}<li>{column}</li>{/each}
  </ul>
</Section>
