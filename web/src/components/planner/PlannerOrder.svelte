<!-- web/src/components/planner/PlannerOrder.svelte -->
<!-- Point order (spec 2026-11-08 section 4.5, decisions P4 and P7, ruling Q5): the build's spend
     order as rows of ten levels, a socket per level, every point visible at every width. A spent
     point is its talent's icon with the rank it reached; the next point's socket is gold with
     `Next point: level 31`; past a character's level the points are planned ahead, dimmed under
     the level line, then the spec guide's order fills the rest, dimmed. An empty build previews that
     guide's order, dimmed, under one line naming it (h-audit 2026-11-09). The rows are
     lib/planner/point-order.ts's; the drawing is ui/PointOrder. -->
<script lang="ts">
  import { iconUrl } from '../../lib/ui/assets';
  import { plannerCopy } from '../../lib/planner/copy';
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
    if ((store.characterLevel === null && store.spent > 0) || plans !== null) return;
    void import('../../lib/leveling/plans-file')
      .then((m) => m.loadLevelingPlans())
      .then((loaded) => (plans = loaded));
  });
  const ahead = $derived.by((): PointOrderAhead | null => {
    const index = store.talentIndex;
    // A character's build continues in the guide's order; an empty build previews it.
    if (plans === null || index === null || (store.characterLevel === null && store.spent > 0)) return null;
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
          (icon) => iconUrl(store.treeVersion, icon),
          ahead,
        ),
  );
</script>

<Section
  title={store.spent === 0 && ahead !== null
    ? plannerCopy.previewLine(ahead.specName)
    : plannerCopy.pointOrderTitle}
  {panel}
  testid="planner-point-order"
>
  {#if store.spent === 0 && ahead === null}
    <!-- Nothing spent and no guide to preview: one line, never rows of empty sockets (h-audit 2026-11-09). -->
    <p class="text-small text-muted pl-pad">{plannerCopy.noPointsYet}</p>
  {:else if view !== null}
    <!-- The rows meet the heading's hairline (the mock's rhythm). -->
    <div class={panel ? '' : '-mt-3'}>
      <!-- The heads read as levels past the row's first ("+0" to "+9"), never as rank counts (ruling 2026-11-09). -->
      <PointOrder
        rows={view.rows}
        capAfter={view.capAfter}
        capText={view.capText}
        markers={plannerCopy.orderMarkers}
        testid="planner-order"
      />
    </div>
  {/if}
</Section>
