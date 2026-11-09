<!-- web/src/components/planner/PlannerOrder.svelte -->
<!-- Point order (spec 2026-11-08 section 4.5, decisions P4 and P7, ruling Q5): the build's spend
     order as rows of ten levels, a socket per level, every point visible at every width. A spent
     point is its talent's icon with the rank it reached; the next point's socket is gold with
     `Next point: level 31`; past a character's level the points are planned ahead, dimmed under
     the level line, then the spec guide's order fills the rest, dimmed. The rows are
     lib/planner/point-order.ts's; the drawing is ui/PointOrder. -->
<script lang="ts">
  import { plannerCopy } from '../../lib/planner/copy';
  import { dataUrl } from '../../lib/planner/load';
  import { specKeyFor } from '../../lib/addon/score';
  import type { LevelingPlans } from '../../lib/leveling/leveling-plan';
  import { aheadOf, pointOrderRows, type PointOrderAhead } from '../../lib/planner/point-order';
  import { SPECS } from '../../lib/sim/specs';
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import PointOrder from '../ui/PointOrder.svelte';
  import Section from '../ui/Section.svelte';

  let {
    store,
    characterName = null,
    panel = false,
  }: {
    store: PlannerStore;
    characterName?: string | null;
    /** The heading inside its panel (the standalone planner's rail and column). */
    panel?: boolean;
  } = $props();

  // For a character, the sockets past its build carry its spec guide's order (/data/leveling-plans.json's
  // ladder), read on demand: a bare build never asks for it.
  let plans = $state<LevelingPlans | null>(null);
  $effect(() => {
    if (store.characterLevel === null || plans !== null) return;
    void import('../../lib/leveling/plans-file')
      .then((m) => m.loadLevelingPlans())
      .then((loaded) => (plans = loaded));
  });
  const ahead = $derived.by((): PointOrderAhead | null => {
    const index = store.talentIndex;
    if (plans === null || index === null || store.characterLevel === null || store.spent === 0) return null;
    const spec = specKeyFor(store.classSlug, store.split);
    const ladder = plans[spec]?.ladder;
    const name = SPECS.find((row) => row.spec === spec)?.name;
    if (ladder === undefined || name === undefined) return null;
    const points = aheadOf(index, store.order, ladder);
    return points.length === 0 ? null : { specName: name, points };
  });

  const view = $derived(
    store.talentIndex === null
      ? null
      : pointOrderRows(
          store.talentIndex,
          store.order,
          store.characterLevel,
          characterName,
          (icon) => dataUrl(store.treeVersion, `icons/${icon}.webp`),
          ahead,
        ),
  );
</script>

<Section title={plannerCopy.pointOrderTitle} {panel} testid="planner-point-order">
  {#if store.spent === 0}
    <p class="text-small text-muted pl-pad">{plannerCopy.noPointsYet}</p>
  {/if}
  {#if view !== null}
    <!-- The rows meet the heading's hairline once a point is spent (the mock's rhythm); the empty
         line keeps the Section's own gap above it. -->
    <div class={store.spent > 0 && !panel ? '-mt-3' : ''}>
      <PointOrder rows={view.rows} capAfter={view.capAfter} capText={view.capText} testid="planner-order" />
    </div>
  {/if}
</Section>
