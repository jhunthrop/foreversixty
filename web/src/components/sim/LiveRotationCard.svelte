<!-- web/src/components/sim/LiveRotationCard.svelte -->
<!-- Design 5.1. The settings bar's fidelity note says the same thing before a run; this
     says it beside the result, because a saved sim and a shared link have no settings bar
     and the number is exactly as good as the rotation that produced it.

     A validated spec renders the card with no note at all rather than a green "all is
     well" line: the absence of a caution is the message.

     Task 3 (healer review MAJOR): `spec` can be a healer or tank spec here -- a saved sim
     or a share link (`store.result.request.spec`) is not gated by RunControl's own
     `isSimulatedSpec` disable, so this checks it itself rather than trusting that a result
     on screen only ever named a dps spec. An unsimulated spec gets no rotation claim and no
     "what it does" trigger (UnsimulatedSpecCard.svelte's own card, not this one, is where a
     healer or tank spec lives on /sim/specs) -- only the honest line.

     Task 6 (newcomer BLOCKER, tank MAJOR); final whole-branch review C1 (RotationDisclosure
     extracted so SettingsBar.svelte's own pre-run copy of this control stopped navigating
     away too): "what it does" used to be a plain link to `/sim/specs#<spec>`, which explains
     parse fidelity, not the rotation, and threw away the loaded character and the finished
     run on the way (Back did not restore either). It is a Disclosure (Ruling 3's one
     show/hide primitive) now: the trigger opens an in-page drawer with the curated
     rotation's own step notes (rotations.ts, scripts/sync-rotations.mjs), and the page
     never navigates. The fidelity detail is still one click away inside the drawer,
     `target="_blank"`, so nothing is lost either way. -->
<script lang="ts">
  import { simCopy } from '../../lib/sim/copy';
  import { isSimulatedSpec, specDisplayName } from '../../lib/sim/spec-label';
  import type { SpecFidelity } from '../../lib/sim/types';
  import FidelityNote from './FidelityNote.svelte';
  import Section from '../ui/Section.svelte';
  import RotationDisclosure from './RotationDisclosure.svelte';

  let { spec, fidelity }: { spec: string; fidelity: SpecFidelity | null } = $props();

  const name = $derived(specDisplayName(spec));
  const simulated = $derived(isSimulatedSpec(spec));
</script>

<Section title={simCopy.rotationCard} panel testid="sim-rotation-card">
  {#if simulated}
    <p class="text-[14px] leading-5">{simCopy.rotationCardBody(name)}</p>
    <RotationDisclosure
      {spec}
      idPrefix="sim-rotation-card"
      triggerClass="text-nav text-[14px] underline decoration-dotted underline-offset-2"
      triggerTestId="sim-rotation-card-link"
    />
    <FidelityNote {fidelity} testid="sim-rotation-card-note" class="pt-2" />
  {:else}
    <p class="text-[14px] leading-5" data-testid="sim-rotation-not-simulated">
      {simCopy.rotationNotSimulated(name)}
    </p>
  {/if}
</Section>
