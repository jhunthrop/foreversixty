<!-- Its styles travel with its code (css="injected"): it renders only after the island hydrates, and a
     page inlines only the styles of what its server render drew. -->
<svelte:options css="injected" />

<!-- web/src/components/planner/GuideBuilds.svelte -->
<!-- Guide builds (lane P7): the class's published guide builds in the rail, one row per build,
     the spec's icon, its name, the build and its split; the row opens it in the guide's levelling
     order. The build on screen carries a Loaded tag. Nothing draws when the class has no guide. -->
<script lang="ts">
  import { ranksByTalent } from '../../lib/planner/derive';
  import { guideBuildRows, ranksTrees, type GuideBuild } from '../../lib/planner/guide-builds';
  import { plannerCopy } from '../../lib/planner/copy';
  import { dataUrl } from '../../lib/planner/load';
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
          testid={`planner-guide-build-${row.key}`}
        >
          {#snippet lead()}
            {#if row.icon}
              <img
                class="gb-icon"
                src={dataUrl(store.treeVersion, `icons/${row.icon}.webp`)}
                alt=""
                width="28"
                height="28"
                loading="lazy"
                decoding="async"
              />
            {/if}
          {/snippet}
          {#snippet subContent()}
            <span class="gb-sub"
              >{`${row.label} · ${row.split.join('/')}`}{#if row.trees === current}<span class="gb-tag"
                  ><Tag tone="muted">{plannerCopy.guideBuildLoaded}</Tag></span
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
  .gb-tag {
    display: inline-flex;
  }
  .gb-icon {
    display: block;
    width: 28px;
    height: 28px;
    border-radius: 4px;
  }
</style>
