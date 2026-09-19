<!-- web/src/components/sim/RotationCard.svelte -->
<!-- Design 5.1. The settings bar's fidelity note says the same thing before a run; this
     says it beside the result, because a saved sim and a shared link have no settings bar
     and the number is exactly as good as the rotation that produced it.

     A validated spec renders the card with no note at all rather than a green "all is
     well" line: the absence of a caution is the message. -->
<script lang="ts">
  import { simCopy } from '../../lib/sim/copy';
  import { specRow } from '../../lib/sim/spec-label';
  import { specPillClass, specStateLabel, specStateNote } from '../../lib/sim/spec-state';
  import type { SpecFidelity } from '../../lib/sim/types';

  let { spec, fidelity }: { spec: string; fidelity: SpecFidelity | null } = $props();

  const name = $derived(specRow(spec)?.name ?? spec);
  const showNote = $derived(fidelity !== null && fidelity.state !== 'validated');
</script>

<section
  class="border-line bg-raised rounded-panel mx-[18px] flex flex-col gap-2 border p-4 md:mx-0"
  data-testid="sim-rotation-card"
>
  <h2 class="section-title text-[15px]">{simCopy.rotationCard}</h2>
  <p class="text-[13px]">
    {simCopy.rotationCardBody(name)}
    <a href={`/sim/specs#${spec}`} class="ml-1" data-testid="sim-rotation-card-link">
      {simCopy.rotationLink}
    </a>
  </p>
  {#if showNote && fidelity !== null}
    <p class="text-muted text-[13px]" data-testid="sim-rotation-card-note">
      <a href="/sim/specs" class={specPillClass(fidelity.state)}>{specStateLabel(fidelity.state)}</a>
      {specStateNote(fidelity.state)}
    </p>
  {/if}
</section>
