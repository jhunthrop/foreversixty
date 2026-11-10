<!-- web/src/components/sim/SimFidelityCaption.svelte -->
<!-- The spec's calibration in the head's answer (spec 2026-11-06-sim §4.6), on the error's line: a
     muted Tag naming it ("No parses yet", "Calibrating"), linking the spec's own row on the spec
     list, where its sentence is said; the same sentence is the Tag's tooltip and a screen reader's
     line, so the head keeps the compact answer contract (ui/PageHeader). The rotation card says it
     in full too (FidelityNote). A validated spec (or a row not yet loaded) renders nothing. /sim and
     /sim/<id> share it. -->
<script lang="ts">
  import { needsFidelityNote, specHeadLabel, specStateNote } from '../../lib/sim/spec-state';
  import type { SpecFidelity } from '../../lib/sim/types';
  import Tag from '../ui/Tag.svelte';

  let { fidelity }: { fidelity: SpecFidelity | null } = $props();
</script>

{#if fidelity !== null && needsFidelityNote(fidelity) && fidelity.state !== 'validated'}
  <span class="sim-fidelity" data-testid="spec-fidelity-note">
    <a href="/sim/specs#{fidelity.spec}" title={specStateNote(fidelity.state)}
      ><Tag tone="muted">{specHeadLabel(fidelity.state)}</Tag><span class="sr-only">
        {specStateNote(fidelity.state)}</span
      ></a
    >
  </span>
{/if}

<style>
  .sim-fidelity {
    display: inline-flex;
  }
  /* A 44 px tap target round the 22 px Tag, without growing the caption's line. */
  .sim-fidelity a {
    display: inline-flex;
    align-items: center;
    min-height: 44px;
    margin-block: -11px;
  }
</style>
