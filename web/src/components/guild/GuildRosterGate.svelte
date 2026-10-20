<!-- web/src/components/guild/GuildRosterGate.svelte -->
<!-- The `#roster` anchor for a viewer the roster is closed to: one sentence naming whom it is
     for (lib/guild/roster-gate.ts), no rows. Tailwind utilities only, like every guild/*
     component (they mount after the guild's own fetch, so scoped CSS would be dropped). -->
<script lang="ts">
  import { rosterGateCopy, type RosterGateState } from '../../lib/guild/roster-gate';

  let { state }: { state: RosterGateState } = $props();
</script>

<section id="roster" class="flex flex-col gap-2" data-testid="guild-roster-gate" data-state={state}>
  <h2 class="label text-gold">{rosterGateCopy.heading}</h2>
  <p class="text-muted text-[14px]">{rosterGateCopy[state]}</p>
</section>
