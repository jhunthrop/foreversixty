<!-- web/src/components/character/ChartSrTable.svelte -->
<!-- The screen-reader form of a History chart, beside the chart (the Sparkline's pattern): a
     captioned table of dates and values, oldest first, each gap a row of its own. -->
<script lang="ts">
  import type { SrRow } from '../../lib/character/history-chart';
  import { historyCopy } from '../../lib/character/history-copy';

  let {
    caption,
    valueHeading,
    rows,
    testid,
  }: { caption: string; valueHeading: string; rows: readonly SrRow[]; testid: string } = $props();
</script>

<table class="sr-only" data-testid={testid}>
  <caption>{caption}</caption>
  <thead>
    <tr><th scope="col">{historyCopy.srColumns.date}</th><th scope="col">{valueHeading}</th></tr>
  </thead>
  <tbody>
    {#each rows as row, index (index)}
      <tr>
        {#if row.value === undefined}
          <td colspan="2">{row.text}</td>
        {:else}
          <td>{row.text}</td>
          <td>{row.value}</td>
        {/if}
      </tr>
    {/each}
  </tbody>
</table>
