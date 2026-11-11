<!-- web/src/components/bis/BisItinerary.svelte -->
<!-- The itinerary under the paperdoll (design/specs/2026-11-09-directions.md, Best in slot): one
     panel, the stops in route order (instances by pieces, then crafted, reputation, quest, vendor,
     world drops), each headed by its name and level range, each piece an ItemRow: the item, the
     slot and what the source adds to the heading (wrapping only between its facts), and what it
     is worth: a gain, Worn, or the row's own state; a tie with the next pick shows nothing. -->
<script lang="ts">
  import Plate from '../ui/Plate.svelte';
  import ItemRow from '../ui/ItemRow.svelte';
  import SectionTitle from '../ui/SectionTitle.svelte';
  import { rowFigure, type ItineraryStop } from '../../lib/bis/frame-view';
  import { factSegments } from '../../lib/ui/primitives';

  let {
    stops,
    build,
    label,
  }: {
    stops: readonly ItineraryStop[];
    build: string;
    label: string;
  } = $props();
</script>

<Plate {label} flush={false} testid="bis-itinerary">
  <div class="bis-stops-cols">
    {#each stops as stop (stop.key)}
      <section class="bis-stop" id={stop.anchor} data-testid={`bis-stop-${stop.key}`}>
        <SectionTitle class="bis-stop-title" title={stop.title}>
          {#snippet after()}{#if stop.range}<span class="bis-stop-range">{stop.range}</span>{/if}{/snippet}
        </SectionTitle>
        {#each stop.rows as row (row.cell.slot)}
          <ItemRow
            item={row.cell.item}
            {build}
            gain={rowFigure(row.cell.gain)}
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
  </div>
</Plate>

<style>
  /* One fact of the line, kept whole: it wraps between facts (lib/ui/primitives factSegments). */
  .bis-fact {
    display: inline-block;
    max-width: 100%;
  }
  .bis-stop {
    break-inside: avoid;
    padding-bottom: 18px;
  }
  @media (min-width: 1280px) {
    .bis-stops-cols {
      columns: 2;
      column-gap: 40px;
    }
  }
  .bis-stop :global(.bis-stop-title) {
    display: flex;
    align-items: baseline;
    gap: 10px;
    padding: 4px 0 8px;
    scroll-margin-top: calc(var(--bar-h) + 16px);
  }
  .bis-stop-range {
    font: 600 14px/1 var(--font-body);
    letter-spacing: 0;
    color: var(--color-text);
  }
</style>
