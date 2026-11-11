<!-- web/src/components/guild/report/ReportConsumables.svelte -->
<!-- Consumables by pull (web layout spec 2026-11-05 §7.2 "Raid report", §7.3 Grid): raider by boss
     pull, each cell flask, food and weapon enchant as read from the pull's character sheet against
     the consumables catalogue: up, proven missing (red), or the log cannot say. A pull with no
     sheet (or not read yet) is unknown, never missed: its cells read `?` and one muted line says
     why. A member's report carries only their own rows. -->
<script lang="ts">
  import type { RaidPullConsumables } from '../../../lib/guild/attendance-contract';
  import { consumablesModel } from '../../../lib/guild/report-night-view';
  import { reportNightCopy as copy } from '../../../lib/guild/report-night-copy';
  import Crest from '../../ui/Crest.svelte';
  import Grid from '../../ui/Grid.svelte';
  import Section from '../../ui/Section.svelte';

  let { byPull }: { byPull: readonly RaidPullConsumables[] } = $props();

  const model = $derived(consumablesModel(byPull));
</script>

<!-- The grid sits on its panel; with no sheet to read, the one line is prose on the ground. -->
<Section
  title={copy.consumables.head}
  anchor="raid-consumables"
  panel={model.rows.length > 0}
  testid="raid-consumables"
>
  {#if byPull.length === 0}
    <p class="text-muted text-[14px]" data-testid="raid-consumables-none">{copy.consumables.none}</p>
  {:else}
    {#if model.rows.length > 0}
      <div class="overflow-x-auto">
        <Grid
          caption={copy.consumables.caption}
          labelHeader={copy.consumables.raider}
          columns={model.columns}
          rows={model.rows}
          testid="raid-consumables-grid"
        >
          {#snippet lead(row)}<Crest cls={row.class} size={20} />{/snippet}
        </Grid>
      </div>
      <p class="text-muted mt-2 text-[13px]" data-testid="raid-consumables-legend">
        {copy.consumables.legend}
      </p>
    {/if}
    {#each model.unknown as line (line)}
      <p class="text-muted mt-1 text-[13px]" data-testid="raid-consumables-unknown">{line}</p>
    {/each}
  {/if}
</Section>
