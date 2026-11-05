<!-- web/src/components/sim/FidelityNote.svelte -->
<!-- The spec's fidelity, said beside a result (design 5.1, tenet 8): a spec the sim has not
     validated is labelled, with its state pill and the spec list's own sentence. A validated spec
     (or a fidelity row not yet loaded) renders nothing: the absence of a caution is the message.
     The live rotation card and the saved Rotation card share it. -->
<script lang="ts">
  import { needsFidelityNote, specPillClass, specStateLabel, specStateNote } from '../../lib/sim/spec-state';
  import type { SpecFidelity } from '../../lib/sim/types';

  let {
    fidelity,
    class: extraClass = '',
    testid,
  }: { fidelity: SpecFidelity | null; class?: string; testid: string } = $props();
</script>

{#if fidelity !== null && needsFidelityNote(fidelity)}
  <p class="text-muted text-[13px] {extraClass}" data-testid={testid}>
    <a href="/sim/specs" class={specPillClass(fidelity.state)}>{specStateLabel(fidelity.state)}</a>
    {specStateNote(fidelity.state)}
  </p>
{/if}
