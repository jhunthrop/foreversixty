<!-- web/src/components/planner/SummaryBar.svelte -->
<!-- design/DESIGN-SYSTEM.md "State panel": label rows on a raised panel, mono numbers,
     gold accent. Pinned to the top of the viewport on phone (design/Mobile.dc.html keeps
     the current state above the fold), static on desktop. -->
<script lang="ts">
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import { MAX_POINTS } from '../../lib/planner/types';

  let { store }: { store: PlannerStore } = $props();

  const controlClass =
    'border-line-warm rounded-control bg-raised text-text h-11 min-w-[8rem] border px-3 text-[14px] font-semibold';
</script>

<div
  class="border-line bg-raised md:rounded-panel sticky top-0 z-20 flex flex-wrap items-center gap-3 border-b px-[18px] py-3 md:static md:gap-5 md:border md:px-5"
>
  <label class="flex flex-col gap-1">
    <span class="label text-muted">Class</span>
    <!-- A function binding rather than a plain `value` attribute: on a <select> the value
         has to be applied after the options exist, and Svelte only does that for bindings. -->
    <select
      class={controlClass}
      bind:value={() => store.classSlug, (slug) => store.selectClass(slug)}
      disabled={store.readOnly}
    >
      {#each store.classes as row (row.slug)}
        <option value={row.slug}>{row.name}</option>
      {/each}
    </select>
  </label>

  <label class="flex flex-col gap-1">
    <span class="label text-muted">Race</span>
    <select
      class={controlClass}
      bind:value={() => store.raceSlug, (slug) => store.selectRace(slug)}
      disabled={store.readOnly}
    >
      {#each store.legalRaces as row (row.slug)}
        <option value={row.slug}>{row.name}</option>
      {/each}
    </select>
  </label>

  <div class="flex flex-col gap-1">
    <span class="label text-muted">Level</span>
    <span class="tabular text-strong font-mono text-[20px] leading-11" data-testid="planner-level">
      {store.level}
    </span>
  </div>

  <div class="flex flex-col gap-1">
    <span class="label text-muted">Left</span>
    <span class="tabular text-gold font-mono text-[20px] leading-11" data-testid="planner-remaining">
      {MAX_POINTS - store.spent}
    </span>
  </div>

  <div class="flex flex-col gap-1">
    <span class="label text-muted">Split</span>
    <span class="tabular text-gold font-mono text-[20px] leading-11" data-testid="planner-split">
      {store.splitLabel}
    </span>
  </div>

  <div class="flex flex-col gap-1">
    <span class="label text-muted">Points</span>
    <span class="tabular text-strong font-mono text-[20px] leading-11" data-testid="planner-spent">
      {store.spent}/{MAX_POINTS}
    </span>
  </div>

  <p
    role="status"
    aria-live="polite"
    class="text-muted order-last min-h-[20px] w-full text-[13px] leading-tight"
    data-testid="planner-refusal"
  >
    {store.refusal ?? ''}
  </p>
</div>
