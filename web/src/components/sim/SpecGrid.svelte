<!-- web/src/components/sim/SpecGrid.svelte -->
<!-- Three states, one grid: loading (skeleton cards at the grid's own height, so the page
     does not reflow when the answer lands -- the LCP budget is measured on /sim/specs too),
     failed (one alert line and a retry), and answered. The answered state always renders
     `dpsSpecs().length` cards, from mergeSpecRows: a spec the API said nothing about reads
     "Not yet" rather than being missing from the grid. -->
<script lang="ts">
  import { simCopy } from '../../lib/sim/copy';
  import { dpsSpecs } from '../../lib/sim/spec-label';
  import { mergeSpecRows } from '../../lib/sim/spec-state';
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

  const skeletonSlots = Array.from({ length: dpsSpecs().length }, (_, index) => index);
  const grid = 'grid grid-cols-1 gap-3 px-[18px] md:grid-cols-2 md:px-0 lg:grid-cols-3';

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

{#if error !== null}
  <p class="text-muted px-[18px] text-[14px] md:px-0" role="alert" data-testid="specs-error">
    {error}
    <button
      type="button"
      class="text-strong ml-1 inline-flex min-h-11 items-center underline underline-offset-2"
      data-testid="specs-retry"
      onclick={() => onretry()}>{simCopy.tryAgain}</button
    >
  </p>
{:else if rows === null}
  <div class={grid} data-testid="specs-grid">
    {#each skeletonSlots as slot (slot)}
      <div class="bg-card-top border-line rounded-panel min-h-[168px] border" aria-hidden="true"></div>
    {/each}
  </div>
{:else}
  <div class={grid} data-testid="specs-grid">
    {#each mergeSpecRows(rows) as row (row.spec)}
      <SpecCard {row} />
    {/each}
  </div>
{/if}
