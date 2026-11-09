<!-- web/src/components/bis/BisItinerary.svelte -->
<!-- The itinerary under the paperdoll (design/specs/2026-11-09-directions.md, Best in slot): one
     panel, the stops in route order (instances by pieces, then crafted, reputation, quest, vendor,
     world drops), each headed by its name and level range, each piece an ItemRow: the item, the
     slot and what the source adds to the heading (wrapping only between its facts), and what it
     is worth: a gain, "Same DPS" as the next pick, Worn, or the row's own state. -->
<script lang="ts">
  import Plate from '../ui/Plate.svelte';
  import ItemRow from '../ui/ItemRow.svelte';
  import { rowFigure, type ItineraryStop } from '../../lib/bis/frame-view';
  import { factSegments } from '../../lib/ui/primitives';

  let {
    stops,
    build,
    unit,
    label,
  }: {
    stops: readonly ItineraryStop[];
    build: string;
    /** The band's unit ("DPS"), for a tie's "Same DPS". */
    unit: string;
    label: string;
  } = $props();
</script>

<Plate {label} flush={false} testid="bis-itinerary">
  {#each stops as stop (stop.key)}
    <section class="bis-stop" id={stop.anchor} data-testid={`bis-stop-${stop.key}`}>
      <h2 class="bis-stop-title text-section">
        {stop.title}{#if stop.range}<span class="bis-stop-range">{stop.range}</span>{/if}
      </h2>
      {#each stop.rows as row (row.cell.slot)}
        <ItemRow
          item={row.cell.item}
          {build}
          gain={rowFigure(row.cell.gain, unit)}
          badge={row.cell.badge}
          testid={`bis-row-${row.cell.slot}`}
        >
          {#snippet sourceContent()}{#each factSegments(row.line) as fact, index (index)}{index > 0
                ? ' '
                : ''}<span class="bis-fact">{fact}</span>{/each}{/snippet}
        </ItemRow>
      {/each}
    </section>
  {/each}
</Plate>

<style>
  /* One fact of the line, kept whole: it wraps between facts (lib/ui/primitives factSegments). */
  .bis-fact {
    display: inline-block;
    max-width: 100%;
  }
  .bis-stop + .bis-stop {
    margin-top: 18px;
  }
  .bis-stop-title {
    display: flex;
    align-items: baseline;
    gap: 10px;
    padding: 4px 0 8px;
    font-family: var(--font-title);
    font-weight: 400;
    letter-spacing: var(--tracking-heading);
    color: var(--color-cream);
    scroll-margin-top: calc(var(--bar-h) + 16px);
  }
  .bis-stop-range {
    font: 600 14px/1 var(--font-body);
    letter-spacing: 0;
    color: var(--color-text);
  }
</style>
