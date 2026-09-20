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
     "what it does" trigger (there is nothing at `/sim/specs#<spec>` for it to point at,
     task-2-brief.md's own split moved it into the unsimulated group with no id anchor) --
     only the honest line.

     Task 6 (newcomer BLOCKER, tank MAJOR): "what it does" used to be a plain link to
     `/sim/specs#<spec>`, which explains parse fidelity, not the rotation, and threw away
     the loaded character and the finished run on the way (Back did not restore either).
     It is a Disclosure (Ruling 3's one show/hide primitive) now: the trigger opens an
     in-page drawer with the curated rotation's own step notes (rotations.ts,
     scripts/sync-rotations.mjs), and the page never navigates. The fidelity detail is still
     one click away inside the drawer, `target="_blank"`, so nothing is lost either way. -->
<script lang="ts">
  import { simCopy } from '../../lib/sim/copy';
  import { isSimulatedSpec, specDisplayName } from '../../lib/sim/spec-label';
  import { needsFidelityNote, specPillClass, specStateLabel, specStateNote } from '../../lib/sim/spec-state';
  import { rotationDrawerContent } from '../../lib/sim/rotations';
  import type { SpecFidelity } from '../../lib/sim/types';
  import Disclosure from './Disclosure.svelte';

  let { spec, fidelity }: { spec: string; fidelity: SpecFidelity | null } = $props();

  const name = $derived(specDisplayName(spec));
  const simulated = $derived(isSimulatedSpec(spec));
  const showNote = $derived(simulated && needsFidelityNote(fidelity));
  const drawer = $derived(rotationDrawerContent(spec));
</script>

<section
  class="border-line bg-raised rounded-panel mx-[18px] flex flex-col gap-2 border p-4 md:mx-0"
  data-testid="sim-rotation-card"
>
  <h2 class="section-title text-[15px]">{simCopy.rotationCard}</h2>
  {#if simulated}
    <p class="text-[13px]">{simCopy.rotationCardBody(name)}</p>
    <Disclosure
      label={simCopy.rotationLink}
      id={`sim-rotation-drawer-${spec}`}
      triggerClass="label text-nav text-[13px] underline decoration-dotted underline-offset-2"
      panelClass="border-line-soft rounded-panel flex flex-col gap-2 border p-3"
      triggerTestId="sim-rotation-card-link"
      panelTestId="sim-rotation-drawer-panel"
    >
      <p class="text-[13px]">{drawer.intro}</p>
      {#if drawer.steps.length > 0}
        <ol class="flex list-decimal flex-col gap-1 pl-5 text-[13px]" data-testid="sim-rotation-drawer-steps">
          {#each drawer.steps as step, index (index)}
            <li>{step}</li>
          {/each}
        </ol>
      {:else}
        <p class="text-muted text-[13px]">{simCopy.rotationDrawerEmpty}</p>
      {/if}
      <a
        href={`/sim/specs#${spec}`}
        target="_blank"
        rel="noopener"
        class="text-nav text-[12px] underline decoration-dotted underline-offset-2"
        data-testid="sim-rotation-drawer-fidelity-link"
      >
        {simCopy.rotationDrawerFidelityLink}
      </a>
    </Disclosure>
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
