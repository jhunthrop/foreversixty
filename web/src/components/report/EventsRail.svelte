<!-- web/src/components/report/EventsRail.svelte -->
<!-- The Events tab's rail Section (logs spec 2026-11-07 section 6.2): `Show`, the checkbox chips
     that pick which kinds of line the table lists. It drives the same url state EventsView reads
     (`eventsOff`), so the frame passes it the page's `state.eventsOff` and `patch`. -->
<script lang="ts">
  import Section from '../ui/Section.svelte';
  import EventChips from './EventChips.svelte';

  let {
    off = [],
    onPatch = () => {},
  }: {
    off?: readonly string[];
    onPatch?: (patch: { eventsOff?: string[] }) => void;
  } = $props();
</script>

<Section title="Show" testid="events-rail" panel>
  <EventChips {off} {onPatch} />
</Section>
