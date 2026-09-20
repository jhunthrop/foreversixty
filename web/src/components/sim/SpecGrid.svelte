<!-- web/src/components/sim/SpecGrid.svelte -->
<!-- Two groups (task-2-brief.md): the simulated (dps) grid, and the unsimulated
     (healer/tank) grid below it, under their own headings so they never interleave.
     The simulated grid has three states, always the same card count (specSkeletonSlots(),
     one per dps spec) so the reserved height never changes across them: loading (skeleton
     cards at the grid's own height, so the page does not reflow when the answer lands --
     the LCP budget is measured on /sim/specs too), failed (the alert takes the first
     card's own slot rather than adding a row above the skeletons -- not a reflow to one
     bare line, which used to collapse the page by the grid's full height the moment a
     fetch failed, and not a taller grid than the loading state either; see the Lighthouse
     findings in the final whole-branch review, round 5), and answered. The answered state
     always renders `dpsSpecs().length` cards, from mergeSpecRows: a spec the API said
     nothing about reads "Not yet" rather than being missing from the grid. The unsimulated
     grid has one state: it renders nonDpsSpecs() directly, with no fidelity pill, engine
     stamp or "Not yet" link, since none of that data will ever exist for a healer or tank.
     sim/specs.astro's static shell renders the identical loading skeleton for both groups
     (spec-skeleton.ts) before this island ever mounts, so first paint already reserves the
     same height too. -->
<script lang="ts">
  import { rowLink } from '../../lib/report/format';
  import { simCopy } from '../../lib/sim/copy';
  import { nonDpsSpecs } from '../../lib/sim/spec-label';
  import { mergeSpecRows } from '../../lib/sim/spec-state';
  import {
    SPEC_CARD_SKELETON_CLASSES,
    SPEC_GRID_CLASSES,
    specSkeletonSlots,
  } from '../../lib/sim/spec-skeleton';
  import type { SpecFidelity } from '../../lib/sim/types';
  import SpecCard from './SpecCard.svelte';
  import UnsimulatedSpecCard from './UnsimulatedSpecCard.svelte';

  let {
    rows,
    error,
    onretry,
  }: {
    rows: SpecFidelity[] | null;
    error: string | null;
    onretry: () => void;
  } = $props();

  const skeletonSlots = specSkeletonSlots();
  const grid = SPEC_GRID_CLASSES;
  // The 7 healer/tank cards carry no fidelity data (they are never in `rows`, task-2-brief.md's
  // second group) so they need no loading state of their own -- derived once from the
  // canonical spec list and rendered immediately, the same way sim/specs.astro's static
  // shell reserves their height with unsimulatedSpecSkeletonSlots() before this island ever
  // mounts.
  const unsimulatedSpecs = nonDpsSpecs();

  // "/sim/specs#warrior-fury" from the settings bar's rotation link: the browser's own
  // fragment scroll fires once, on the shell's first paint, long before this island fetches
  // and renders a single card -- so the scroll has to happen here, once the row it names
  // actually exists in the DOM, rather than being left to a navigation the browser already
  // finished acting on.
  $effect(() => {
    if (rows === null) return;
    const id = window.location.hash.slice(1);
    if (id === '') return;
    document.getElementById(id)?.scrollIntoView();
  });
</script>

<!-- Two groups, never interleaved (task-2-brief.md): the 20 damage specs the simulator
     covers, then the 7 healers and tanks it does not, each under its own heading so the
     page reads as "these 20 work, these 7 do not" rather than one undifferentiated 27-card
     grid. -->
<div class="flex flex-col gap-6">
  <div class="flex flex-col gap-3">
    <h2 class="section-title px-[18px] text-[15px] md:px-0">{simCopy.specsSimulatedHeading}</h2>
    {#if rows === null}
      <div class={grid} data-testid="specs-grid">
        <!-- The error message takes the FIRST skeleton card's own slot rather than adding a
             row above it (round 5): the grid's reserved height must equal what renders in
             both the loaded and the error state, at the same card count sim/specs.astro's
             static shell already reserves for -- adding a row here would make the error state
             taller than everything else that ever occupies this box. -->
        {#each skeletonSlots as slot, index (slot)}
          {#if index === 0 && error !== null}
            <div
              class={`${SPEC_CARD_SKELETON_CLASSES} flex flex-col items-start justify-center gap-2 p-4`}
              role="alert"
              data-testid="specs-error"
            >
              <p class="text-muted text-[14px]">{error}</p>
              <button
                type="button"
                class={`text-strong ${rowLink} underline underline-offset-2`}
                data-testid="specs-retry"
                onclick={() => onretry()}>{simCopy.tryAgain}</button
              >
            </div>
          {:else}
            <div class={SPEC_CARD_SKELETON_CLASSES} aria-hidden="true"></div>
          {/if}
        {/each}
      </div>
    {:else}
      <div class={grid} data-testid="specs-grid">
        {#each mergeSpecRows(rows) as row (row.spec)}
          <SpecCard {row} />
        {/each}
      </div>
    {/if}
  </div>

  <div class="flex flex-col gap-3">
    <h2 class="section-title px-[18px] text-[15px] md:px-0">{simCopy.specsUnsimulatedHeading}</h2>
    <div class={grid} data-testid="specs-grid-unsimulated">
      {#each unsimulatedSpecs as spec (spec.spec)}
        <UnsimulatedSpecCard {spec} />
      {/each}
    </div>
  </div>
</div>
