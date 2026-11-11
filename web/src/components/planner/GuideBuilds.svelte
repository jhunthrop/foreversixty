<!-- Its styles travel with its code (css="injected"): it renders only after the island hydrates, and a
     page inlines only the styles of what its server render drew. -->
<svelte:options css="injected" />

<!-- web/src/components/planner/GuideBuilds.svelte -->
<!-- Guide builds (lane P7): the class's published guide builds in the rail, one row per build,
     the spec's icon, its name, the build and its split; the row opens it in the guide's levelling
     order: build vs guide (h-audit 2026-11-09). The build on screen carries a Loaded tag; every other
     row, once a point is spent, says how many of its points the build on screen lacks, the split's
     change tree by tree, and the first talent it lacks. Nothing draws when the class has no guide. -->
<script lang="ts">
  import { iconUrl } from '../../lib/ui/assets';
  import SpellIcon from '../ui/SpellIcon.svelte';
  import { ranksByTalent } from '../../lib/planner/derive';
  import {
    deltaText,
    guideBuildRows,
    guideDiff,
    ranksTrees,
    type GuideBuild,
  } from '../../lib/planner/guide-builds';
  import { plannerCopy } from '../../lib/planner/copy';
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import Row from '../ui/Row.svelte';
  import Tag from '../ui/Tag.svelte';
  import Section from '../ui/Section.svelte';

  let { builds, store }: { builds: readonly GuideBuild[]; store: PlannerStore } = $props();

  const rows = $derived(guideBuildRows(builds, store.classSlug));
  /** The build on screen in the rows' form, to mark the guide build it is. */
  const current = $derived.by(() => {
    const index = store.talentIndex;
    if (index === null || store.spent === 0) return null;
    const ranks = ranksByTalent(store.order);
    return ranksTrees(index.trees.map((tree) => tree.talents.map((talent) => ranks.get(talent.id) ?? 0)));
  });
  /** Each tree's talent names in the rows' order, for the first talent a guide build takes that this one lacks. */
  const names = $derived(
    store.talentIndex?.trees.map((tree) => tree.talents.map((talent) => talent.name)) ?? [],
  );
  const diffOf = (trees: string) =>
    current === null || trees === current ? null : guideDiff(trees, current, names);
</script>

{#if rows.length > 0}
  <Section title={plannerCopy.guideBuildsTitle} panel testid="planner-guide-builds">
    <div class="gb-rows">
      {#each rows as row (row.key)}
        <Row
          title={row.name}
          sub={`${row.label} · ${row.split.join('/')}`}
          href={row.href}
          density="slim"
          open
          wrapTrail={diffOf(row.trees) !== null}
          testid={`planner-guide-build-${row.key}`}
        >
          {#snippet lead()}
            {#if row.icon}
              <SpellIcon url={iconUrl(store.treeVersion, row.icon)} size={28} />
            {/if}
          {/snippet}
          {#snippet trail()}
            {@const diff = diffOf(row.trees)}
            {#if diff !== null}<span class="gb-off" data-testid={`planner-guide-build-off-${row.key}`}
                ><span class="text-strong">{plannerCopy.guideBuildOff(diff.off)}</span><span
                  class="text-muted">{deltaText(diff.delta)}</span
                ></span
              >{/if}
          {/snippet}
          {#snippet subContent()}
            {@const diff = diffOf(row.trees)}
            <span class="gb-sub"
              >{`${row.label} · ${row.split.join('/')}`}{#if row.trees === current}<span class="gb-tag"
                  ><Tag tone="gold">{plannerCopy.guideBuildLoaded}</Tag></span
                >{:else if diff !== null && diff.missing !== null}<span class="gb-miss"
                  >{plannerCopy.guideBuildMissing(diff.missing)}</span
                >{/if}</span
            >
          {/snippet}
        </Row>
      {/each}
    </div>
  </Section>
{/if}

<style>
  .gb-rows {
    display: flex;
    flex-direction: column;
  }
  .gb-sub {
    display: inline-flex;
    flex-wrap: wrap;
    align-items: center;
    gap: 2px 8px;
  }
  /* The talent a guide row lacks takes its own line under the build and split, at every width. */
  .gb-miss {
    flex-basis: 100%;
  }
  .gb-tag {
    display: inline-flex;
  }
  .gb-off {
    display: flex;
    flex-direction: column;
    align-items: flex-end;
    font-size: 13px;
    line-height: 18px;
    white-space: nowrap;
    font-variant-numeric: tabular-nums;
  }
</style>
