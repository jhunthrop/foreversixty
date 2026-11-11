<script lang="ts">
  import { plannerSimCopy as simCopy } from '../../lib/sim/planner-copy';
  import { plannerCopy } from '../../lib/planner/copy';
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import { LINK } from '../../lib/ui/primitives';
  import Figure from '../ui/Figure.svelte';

  let {
    store,
    simHref,
  }: {
    store: PlannerStore;
    /** `Sim this build ›`; undefined hides it (its line is kept). */
    simHref?: string;
  } = $props();

  const spent = $derived(Math.min(store.spent, store.pointsAvailable));
</script>

<!-- web/src/components/planner/PlannerAnswer.svelte -->
<!-- The planner head's answer (directions mock b2-planner): the points spent of the points there
     are, then `Sim this build ›`, which holds its line while nothing is spent. The planner owns the
     talents and the simulator owns the numbers (one job per page), so no figure of the sim's sits
     here: the link carries the build there. Points spent past a character's level add one line. -->

<div class="pl-ans" data-testid="planner-answer">
  <Figure
    value={String(spent)}
    unit={plannerCopy.pointsOfUnit(store.pointsAvailable)}
    testid="planner-points"
  />
  {#if store.plannedAhead > 0}
    <span class="text-small text-muted" data-testid="planner-spent-line"
      >{plannerCopy.plannedAhead(store.plannedAhead)}</span
    >
  {/if}
  <a
    class="{LINK} pl-sim"
    class:pl-hidden={simHref === undefined}
    href={simHref ?? '/sim'}
    aria-hidden={simHref === undefined ? 'true' : undefined}
    tabindex={simHref === undefined ? -1 : undefined}
    data-testid="planner-sim-link">{simCopy.simThisBuild} ›</a
  >
</div>

<style>
  .pl-ans {
    display: flex;
    flex-direction: column;
    gap: 6px;
    min-width: 0;
  }
  .pl-hidden {
    visibility: hidden;
  }
  @media (min-width: 900px) {
    .pl-ans {
      align-items: flex-end;
      text-align: right;
    }
  }
</style>
