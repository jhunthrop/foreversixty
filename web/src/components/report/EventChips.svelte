<!-- web/src/components/report/EventChips.svelte -->
<!-- The Events tab's Show chips (logs spec 2026-11-07 section 6.2, round 5 b1-report-events): one
     toggle chip per kind of line under the panel head, their state in the url's `eventsOff` so a
     copied link keeps them. -->
<script lang="ts">
  import { EVENT_CHIPS, chipOn, toggleChip } from '../../lib/report/events';
  import Chips from '../ui/Chips.svelte';

  let {
    off = [],
    onPatch = () => {},
  }: {
    /** The kinds switched off, from the url. */
    off?: readonly string[];
    onPatch?: (patch: { eventsOff?: string[] }) => void;
  } = $props();

  const chips = $derived(
    EVENT_CHIPS.map((chip) => ({ id: chip.id, label: chip.label, on: chipOn(chip, off) })),
  );
  function toggle(id: string): void {
    const chip = EVENT_CHIPS.find((one) => one.id === id);
    if (chip !== undefined) onPatch({ eventsOff: toggleChip(chip, off) });
  }
</script>

<Chips {chips} label="Show" onToggle={toggle} testid="events-chips" />
