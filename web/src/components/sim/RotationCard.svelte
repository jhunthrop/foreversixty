<!-- web/src/components/sim/RotationCard.svelte -->
<!-- Design 5.1. The settings bar's fidelity note says the same thing before a run; this
     says it beside the result, because a saved sim and a shared link have no settings bar
     and the number is exactly as good as the rotation that produced it.

     A validated spec renders the card with no note at all rather than a green "all is
     well" line: the absence of a caution is the message.

     Task 3 (healer review MAJOR): `spec` can be a healer or tank spec here -- a saved sim
     or a share link (`store.result.request.spec`) is not gated by RunControl's own
     `isSimulatedSpec` disable, so this checks it itself rather than trusting that a result
     on screen only ever named a dps spec. An unsimulated spec gets no rotation claim and no
     "what it does" anchor (there is nothing at `/sim/specs#<spec>` for it to point at,
     task-2-brief.md's own split moved it into the unsimulated group with no id anchor) --
     only the honest line. Task 6 replaces the dps arm's plain link with an in-page drawer;
     that branch is untouched here on purpose so Task 6 has one place to change. -->
<script lang="ts">
  import { simCopy } from '../../lib/sim/copy';
  import { isSimulatedSpec, specDisplayName } from '../../lib/sim/spec-label';
  import { needsFidelityNote, specPillClass, specStateLabel, specStateNote } from '../../lib/sim/spec-state';
  import type { SpecFidelity } from '../../lib/sim/types';

  let { spec, fidelity }: { spec: string; fidelity: SpecFidelity | null } = $props();

  const name = $derived(specDisplayName(spec));
  const simulated = $derived(isSimulatedSpec(spec));
  const showNote = $derived(simulated && needsFidelityNote(fidelity));
</script>

<section
  class="border-line bg-raised rounded-panel mx-[18px] flex flex-col gap-2 border p-4 md:mx-0"
  data-testid="sim-rotation-card"
>
  <h2 class="section-title text-[15px]">{simCopy.rotationCard}</h2>
  {#if simulated}
    <p class="text-[13px]">
      {simCopy.rotationCardBody(name)}
      <a href={`/sim/specs#${spec}`} class="ml-1" data-testid="sim-rotation-card-link">
        {simCopy.rotationLink}
      </a>
    </p>
  {:else}
    <p class="text-[13px]" data-testid="sim-rotation-not-simulated">
      {simCopy.rotationNotSimulated(name)}
    </p>
  {/if}
  {#if showNote && fidelity !== null}
    <p class="text-muted text-[13px]" data-testid="sim-rotation-card-note">
      <a href="/sim/specs" class={specPillClass(fidelity.state)}>{specStateLabel(fidelity.state)}</a>
      {specStateNote(fidelity.state)}
    </p>
  {/if}
</section>
