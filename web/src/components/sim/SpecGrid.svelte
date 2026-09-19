<!-- web/src/components/sim/SpecGrid.svelte -->
<!-- Three states, one grid container, always the same card count (specSkeletonSlots(),
     one per dps spec) so the reserved height never changes across them: loading (skeleton
     cards at the grid's own height, so the page does not reflow when the answer lands --
     the LCP budget is measured on /sim/specs too), failed (the alert takes the first
     card's own slot rather than adding a row above the skeletons -- not a reflow to one
     bare line, which used to collapse the page by the grid's full height the moment a
     fetch failed, and not a taller grid than the loading state either; see the Lighthouse
     findings in the final whole-branch review, round 5), and answered. The answered state
     always renders `dpsSpecs().length` cards, from mergeSpecRows: a spec the API said
     nothing about reads "Not yet" rather than being missing from the grid. sim/specs.astro's
     static shell renders the identical loading skeleton (spec-skeleton.ts) before this
     island ever mounts, so first paint already reserves the same height too. -->
<script lang="ts">
  import { rowLink } from '../../lib/report/format';
  import { simCopy } from '../../lib/sim/copy';
  import { mergeSpecRows } from '../../lib/sim/spec-state';
  import { SPEC_CARD_SKELETON_CLASSES, SPEC_GRID_CLASSES, specSkeletonSlots } from '../../lib/sim/spec-skeleton';
  import type { SpecFidelity } from '../../lib/sim/types';
  import SpecCard from './SpecCard.svelte';

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
