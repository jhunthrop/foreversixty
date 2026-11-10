<!-- web/src/components/report/QueriesRail.svelte -->
<!-- The Queries tab's rail (round 5, b1-report-queries): `Columns`, the ones a query reaches for
     first as literal code, and `All 66` in the head to list every column of the fight's
     events.parquet. -->
<script lang="ts">
  import { COMMON_EVENT_COLUMNS, EVENT_COLUMNS } from '../../lib/report/query';
  import Section from '../ui/Section.svelte';

  let { columns = EVENT_COLUMNS }: { columns?: readonly string[] } = $props();

  let all = $state(false);
  const shown = $derived(all ? columns : columns.filter((column) => COMMON_EVENT_COLUMNS.includes(column)));
</script>

<Section title="Columns" testid="queries-rail" panel>
  {#snippet head()}<button
      type="button"
      class="text-gold text-[14px] font-semibold hover:underline"
      aria-expanded={all}
      data-testid="queries-rail-all"
      onclick={() => (all = !all)}>{all ? 'Fewer' : `All ${columns.length}`}</button
    >{/snippet}
  <ul class="text-text flex flex-wrap gap-x-3 gap-y-1 font-mono text-[13px] leading-6" aria-label="Columns">
    {#each shown as column (column)}<li>{column}</li>{/each}
  </ul>
</Section>
