<!-- web/src/components/guild/GuildRosterGate.svelte -->
<!-- The `#roster` anchor for a viewer the roster is closed to: an open Section with one sentence
     naming whom it is for (lib/guild/roster-gate.ts), no rows, a section gap below the tab's content. -->
<script lang="ts">
  import { rosterGateCopy, type RosterGateState } from '../../lib/guild/roster-gate';
  import Section from '../ui/Section.svelte';

  let { state }: { state: RosterGateState } = $props();
</script>

<div class="pt-(--section-gap)">
  <Section title={rosterGateCopy.heading} anchor="roster" testid="guild-roster-gate">
    <p class="text-muted text-[14px]" data-state={state}>{rosterGateCopy[state]}</p>
  </Section>
</div>
