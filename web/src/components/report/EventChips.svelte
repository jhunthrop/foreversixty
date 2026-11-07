<!-- web/src/components/report/EventChips.svelte -->
<!-- The Events tab's Show chips (logs spec 2026-11-07 section 6.2): one checkbox per kind of line,
     their state in the url's `eventsOff` so a copied link keeps them. Drawn in the rail
     (EventsRail) and, until the frame wires the rail, above the events table. -->
<script lang="ts">
  import { EVENT_CHIPS, chipOn, toggleChip } from '../../lib/report/events';

  let {
    off = [],
    onPatch = () => {},
  }: {
    /** The kinds switched off, from the url. */
    off?: readonly string[];
    onPatch?: (patch: { eventsOff?: string[] }) => void;
  } = $props();
</script>

<div
  class="flex flex-wrap gap-x-4 gap-y-1 text-[14px]"
  role="group"
  aria-label="Show"
  data-testid="events-chips"
>
  {#each EVENT_CHIPS as chip (chip.id)}
    <label class="text-text flex min-h-11 cursor-pointer items-center gap-2 md:min-h-8">
      <input
        type="checkbox"
        class="accent-gold h-4 w-4"
        checked={chipOn(chip, off)}
        onchange={() => onPatch({ eventsOff: toggleChip(chip, off) })}
        data-testid={`events-chip-${chip.id}`}
      />
      {chip.label}
    </label>
  {/each}
</div>
