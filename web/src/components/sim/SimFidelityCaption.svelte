<!-- web/src/components/sim/SimFidelityCaption.svelte -->
<!-- The spec's fidelity as a caption in the head's answer (spec 2026-11-06-sim §4.6): a spec in
     progress or not yet measured keeps the shipped note, one line under the error, its muted Tag
     linking the spec list. A validated spec (or a row not yet loaded) renders nothing. /sim and
     /sim/<id> share it. -->
<script lang="ts">
  import { needsFidelityNote, specStateLabel, specStateNote } from '../../lib/sim/spec-state';
  import type { SpecFidelity } from '../../lib/sim/types';
  import Tag from '../ui/Tag.svelte';

  let { fidelity }: { fidelity: SpecFidelity | null } = $props();
</script>

{#if fidelity !== null && needsFidelityNote(fidelity)}
  <span class="text-muted text-[13px] leading-5" data-testid="spec-fidelity-note">
    <a href="/sim/specs"><Tag tone="muted">{specStateLabel(fidelity.state)}</Tag></a>
    {specStateNote(fidelity.state)}
  </span>
{/if}
