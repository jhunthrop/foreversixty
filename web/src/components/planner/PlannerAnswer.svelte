<!-- Its styles travel with its code (css="injected"): it renders only after the island hydrates, and a
     page inlines only the styles of what its server render drew. -->
<svelte:options css="injected" />

<script lang="ts">
  import { confidenceBand, formatMargin } from '../../lib/sim/estimate';
  import { plannerSimCopy as simCopy } from '../../lib/sim/planner-copy';
  import { plannerCopy } from '../../lib/planner/copy';
  import { levelReached } from '../../lib/planner/derive';
  import type { LiveDps } from '../../lib/planner/live-dps.svelte';
  import type { LiveGate } from '../../lib/planner/live-gate';
  import { plannerScoreCopy } from '../../lib/planner/score-unit';
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import type { SlotScoreUnit } from '../../lib/bis/tank-view';
  import { LINK } from '../../lib/ui/primitives';
  import Button from '../ui/Button.svelte';
  import Figure from '../ui/Figure.svelte';

  let {
    store,
    live,
    gate,
    unit,
    bandLabel,
    simHref,
    onshow,
  }: {
    store: PlannerStore;
    live: LiveDps;
    gate: LiveGate;
    unit: SlotScoreUnit;
    /** The Level build section's band (`60`, `30 to 39`) once its gear is the estimate's, else null. */
    bandLabel: string | null;
    /** `Sim this build ›`; undefined hides it (its line is kept). */
    simHref?: string;
    onshow: () => void;
  } = $props();

  const spentLine = $derived.by(() => {
    if (store.characterLevel === null) {
      return plannerCopy.spentOf(
        store.spent,
        store.pointsAvailable,
        store.spent > 0 ? levelReached(store.order) : null,
      );
    }
    const line = plannerCopy.spentAtLevel(
      Math.min(store.spent, store.pointsAvailable),
      store.pointsAvailable,
      store.characterLevel,
    );
    return store.plannedAhead > 0 ? `${line} · ${plannerCopy.plannedAhead(store.plannedAhead)}` : line;
  });

  const ready = $derived(live.state !== 'error' && live.estimate.mean > 0 && gate === 'run');
  const figure = $derived(ready ? plannerScoreCopy.figure(live.estimate.mean, unit) : '—');
  const caption = $derived.by(() => {
    if (gate === 'unfinished') return plannerScoreCopy.waiting;
    if (gate === 'ask') return simCopy.plannerDpsShowNote;
    if (live.state === 'error') return live.message ?? plannerScoreCopy.failed(unit);
    if (live.state !== 'ready' || live.estimate.error <= 0) return '';
    const margin = `± ${formatMargin(confidenceBand(live.estimate))}`;
    return bandLabel === null ? margin : `${margin} · ${plannerCopy.scoreBasis(bandLabel)}`;
  });
  const stale = $derived(live.state === 'pending' || live.state === 'running');
</script>

<!-- web/src/components/planner/PlannerAnswer.svelte -->
<!-- The planner head's answer (spec 2026-11-08 section 4.1, mock web-planner*): Points left in
     gold over the spent line (`21 of 25 spent at level 34`, the level once a point is spent),
     the role's one live score on the Level build's gear (decision P1) over its error and that
     basis, then `Sim this build ›`, which holds its line while nothing is spent. Two fixed
     figure blocks in every state, so an estimate arriving moves nothing. -->

<div class="pl-ans" data-testid="planner-answer">
  <div class="pl-figs">
    <div class="pl-fig">
      <Figure
        value={String(store.pointsLeft)}
        label={plannerCopy.pointsLeft}
        tone="gold"
        size="big"
        testid="planner-points-left"
      />
      <span class="text-small text-muted" data-testid="planner-spent-line">{spentLine}</span>
    </div>
    <div class="pl-fig" class:pl-stale={stale}>
      <Figure value={figure} label={plannerScoreCopy.label(unit)} size="big" testid="planner-dps" />
      {#if gate === 'ask'}
        <Button size="sm" onclick={onshow} testid="planner-dps-show">{plannerScoreCopy.show(unit)}</Button>
      {/if}
      <span class="text-small text-muted" data-testid="planner-dps-note">{caption}</span>
    </div>
  </div>
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
    gap: 10px;
    min-width: 0;
  }
  .pl-figs {
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    gap: 24px;
  }
  .pl-fig {
    display: flex;
    flex-direction: column;
    gap: 2px;
    min-width: 0;
  }
  .pl-stale :global(.ui-figure) {
    opacity: 0.6;
  }
  /* Spec 5.2: the captions are two lines below 900 (one from it), whatever they say. */
  @media (max-width: 899px) {
    .pl-fig > :global(.text-small) {
      min-height: 40px;
    }
  }
  .pl-hidden {
    visibility: hidden;
  }
  @media (min-width: 900px) {
    .pl-ans {
      min-width: 380px;
      min-height: 96px;
      justify-content: flex-end;
      align-items: flex-end;
      text-align: right;
    }
    .pl-figs {
      display: flex;
      gap: 32px;
      justify-content: flex-end;
    }
    .pl-fig {
      max-width: 300px;
      align-items: flex-end;
    }
  }
</style>
