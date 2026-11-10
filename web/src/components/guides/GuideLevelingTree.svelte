<!-- web/src/components/guides/GuideLevelingTree.svelte -->
<!-- The Leveling lane's tree (design/specs/2026-11-04-leveling-tree.md, restyled by the guides
     spec 2026-11-08 §3A/§4.3): the guide's own ladder from level 10 to 60, one point a level in
     the order the sim, the addon and the planner share (`guideLadderSteps`). The level line names
     the point taken there and what is spent, the LevelSlider moves the level, the TalentTrees
     show the build at it (each step's talent string over the build's own layout, so nothing
     loads), and Load in planner follows it. The server renders level 60; a signed-in visitor on
     a character of this spec below 60 opens at its level once the answer settles. -->
<script lang="ts">
  import { onMount } from 'svelte';
  import { guidesCopy } from '../../lib/guides/copy';
  import { onGuideCharacter, MAX_LEVEL } from '../../lib/guides/guide-character';
  import type { LadderStep } from '../../lib/guides/leveling-builds';
  import { treesAtStep } from '../../lib/guides/talent-cells';
  import { LINK } from '../../lib/ui/primitives';
  import LevelSlider from '../ui/LevelSlider.svelte';
  import TalentTrees, { type TalentTree } from '../ui/TalentTrees.svelte';

  let {
    build,
    spec,
    classSlug,
    specSlug,
    trees,
    steps,
  }: {
    build: string;
    /** The spec's full id (`rogue-combat`), for the planner link. */
    spec: string;
    classSlug: string;
    specSlug: string;
    /** The guide's level 60 build as cells: the layout every level draws on. */
    trees: TalentTree[];
    /** One per level, 10 to 60 (`guideLadderSteps`). */
    steps: LadderStep[];
  } = $props();

  const first = $derived(steps[0]?.level ?? 10);
  const last = $derived(steps[steps.length - 1]?.level ?? MAX_LEVEL);
  let level = $state(MAX_LEVEL);
  let sliding = $state(false);

  const step = $derived(steps.find((candidate) => candidate.level === level) ?? steps[steps.length - 1]);
  const shown = $derived(step === undefined ? trees : treesAtStep(trees, step.talents));
  const plannerHref = $derived(
    `/planner?spec=${encodeURIComponent(spec)}&talents=${encodeURIComponent(step?.talents ?? '')}`,
  );

  onMount(() =>
    onGuideCharacter((character) => {
      if (
        character !== null &&
        character.classSlug === classSlug &&
        character.specSlug === specSlug &&
        character.level !== undefined &&
        character.level < MAX_LEVEL
      ) {
        level = Math.min(last, Math.max(first, character.level));
      }
    }),
  );
</script>

<div class="guide-leveling" data-testid="guide-leveling-tree">
  <p
    class="guide-leveling-line"
    aria-live={sliding ? 'off' : 'polite'}
    aria-atomic="true"
    data-testid="guide-leveling-caption"
  >
    <b class="text-strong"
      >{step === undefined ? '' : guidesCopy.levelingPoint(step.level, step.name, step.rank)}</b
    >
    {#if last > level}<span class="text-small text-muted"
        >{guidesCopy.levelingPointsSpent(level - first + 1, last - level)}</span
      >{/if}
  </p>
  <LevelSlider
    {level}
    min={first}
    max={last}
    label={guidesCopy.levelSliderLabel}
    valueText={guidesCopy.levelValueText(level)}
    previousLabel={guidesCopy.previousLevel}
    nextLabel={guidesCopy.nextLevel}
    onchange={(next) => (level = next)}
    onslide={(on) => (sliding = on)}
    testid="guide-leveling-slider"
  />
  <TalentTrees trees={shown} {build} testid="guide-leveling-trees" />
  <p class="m-0">
    <a
      class={LINK}
      href={plannerHref}
      aria-label={guidesCopy.loadLevelBuildAriaLabel(level)}
      data-testid="guide-leveling-load">{guidesCopy.loadInPlanner} ›</a
    >
  </p>
</div>

<style>
  .guide-leveling {
    display: flex;
    flex-direction: column;
    gap: 8px;
  }
  .guide-leveling-line {
    display: flex;
    flex-direction: column;
    margin: 0;
  }
</style>
