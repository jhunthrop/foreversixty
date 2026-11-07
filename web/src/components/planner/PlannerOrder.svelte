<!-- web/src/components/planner/PlannerOrder.svelte -->
<!-- Point order (spec 2026-11-08 section 4.5, decisions P4 and P7, ruling Q5): the build's spend
     order as rows of ten levels, a socket per level, every point visible at every width. A spent
     point is its talent's icon with the rank it reached; the next point's socket is gold with
     `Next point: level 31`; past a character's level the points are planned ahead, dimmed under
     the level line. The rows are lib/planner/point-order.ts's; the drawing is ui/PointOrder. -->
<script lang="ts">
  import { plannerCopy } from '../../lib/planner/copy';
  import { dataUrl } from '../../lib/planner/load';
  import { pointOrderRows } from '../../lib/planner/point-order';
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import PointOrder from '../ui/PointOrder.svelte';
  import Section from '../ui/Section.svelte';

  let { store, characterName = null }: { store: PlannerStore; characterName?: string | null } = $props();

  const view = $derived(
    store.talentIndex === null
      ? null
      : pointOrderRows(store.talentIndex, store.order, store.characterLevel, characterName, (icon) =>
          dataUrl(store.treeVersion, `icons/${icon}.webp`),
        ),
  );
</script>

<Section
  title={plannerCopy.pointOrderTitle}
  meta={plannerCopy.pointOrderMeta(store.spent)}
  testid="planner-point-order"
>
  {#if store.spent === 0}
    <p class="text-small text-muted pl-pad">{plannerCopy.noPointsYet}</p>
  {/if}
  {#if view !== null}
    <!-- The rows meet the heading's hairline once a point is spent (the mock's rhythm); the empty
         line keeps the Section's own gap above it. -->
    <div class={store.spent > 0 ? '-mt-3' : ''}>
      <PointOrder rows={view.rows} capAfter={view.capAfter} capText={view.capText} testid="planner-order" />
    </div>
  {/if}
</Section>
