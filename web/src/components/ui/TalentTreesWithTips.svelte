<script lang="ts">
  import { onMount } from 'svelte';
  import { cellTipModel, type TalentTreeView } from '../../lib/ui/talent-tree';
  import type { GameTipModel } from '../../lib/tooltips/game-tip';
  import BisTooltipHost from '../BisTooltipHost.svelte';
  import TalentTrees from './TalentTrees.svelte';

  let {
    trees,
    build,
    label,
    testid,
  }: {
    /** Cells that carry their rank text (`viewTrees` with its text). */
    trees: readonly TalentTreeView[];
    build: string;
    label?: string;
    testid?: string;
  } = $props();

  let root: HTMLDivElement | undefined = $state();
  // The host listens on the document, so it mounts with the island, never in the server render.
  let mounted = $state(false);
  onMount(() => {
    mounted = true;
  });

  const TALENT_KEY = /^talent:(\d+):\d+$/;
  function resolveTip(key: string): Promise<GameTipModel | null> {
    const id = Number(TALENT_KEY.exec(key)?.[1]);
    const cell = trees.flatMap((tree) => tree.cells).find((candidate) => candidate.id === id);
    return Promise.resolve(cell === undefined ? null : cellTipModel(cell));
  }
</script>

<!-- web/src/components/ui/TalentTreesWithTips.svelte -->
<!-- A read-only build on a page with no talent text of its own (the sim, /sim/talents): the site's
     one tree (TalentTrees, view) plus the site's tooltip host scoped to these trees, serving each
     talent's game tooltip from the cells' own text. Hover, keyboard focus, a tap and Escape are the
     host's, as on every guide tree, spell and item. Kept apart from TalentTrees so the planner,
     whose cells open their own tooltip, never carries a host. -->

<TalentTrees mode="view" {trees} {build} {label} {testid} tipScope bind:root />
{#if mounted && root}
  <BisTooltipHost {build} modelsElementId="talent-tree-tooltip-models" {resolveTip} scope={root} />
{/if}
