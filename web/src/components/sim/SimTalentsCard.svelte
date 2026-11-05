<!-- web/src/components/sim/SimTalentsCard.svelte -->
<!-- The run's talents as the planner's own read-only trees (spec 2026-11-04-sim-results 3.9, 4.7):
     three panels, one tab at a time on a phone, the game tooltip on every cell. A string that does
     not fit today's trees keeps its split and says why nothing is drawn. -->
<script lang="ts">
  import activeBuild from '../../data/active-build.json';
  import { guideTreeStore } from '../../lib/guides/tree-store';
  import { orderFromRanks } from '../../lib/planner/fs1';
  import { loadOptional, dataUrl, loadTalents } from '../../lib/planner/load';
  import { indexTalents, type TalentIndex } from '../../lib/planner/rules';
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import { treeRowColumnsClass } from '../../lib/planner/styles';
  import type { TalentFile } from '../../lib/planner/types';
  import type { ActionNames } from '../../lib/sim/action-names';
  import { plannerHrefForSpec, ranksFromTalentsString } from '../../lib/sim/character';
  import { talentsCardCopy as copy } from '../../lib/sim/input-cards-copy';
  import { talentSplit } from '../../lib/sim/result-band';
  import { setResultTalentTips } from '../../lib/sim/result-tips';
  import { fullestTree, talentsFit } from '../../lib/sim/talents-fit';
  import type { SimResult } from '../../lib/sim/types';
  import TreeTabs from '../planner/TreeTabs.svelte';
  import ActionLink from '../ui/ActionLink.svelte';
  import Card from '../ui/Card.svelte';
  import LoadError from '../ui/LoadError.svelte';
  import Skeleton from '../ui/Skeleton.svelte';

  const props: { result: SimResult; names: ActionNames | null } = $props();
  const character = $derived(props.result.request.character);

  type View =
    | { kind: 'loading' }
    | { kind: 'failed' }
    | { kind: 'unfit' }
    | { kind: 'ready'; store: PlannerStore; index: TalentIndex; treeVersion: string };
  let view = $state.raw<View>({ kind: 'loading' });
  let activeTree = $state(0);

  /** The build the request names when that build is still published, else today's. */
  async function talentFile(): Promise<{ file: TalentFile; treeVersion: string }> {
    const saved = props.result.request.data_build;
    if (saved !== undefined && saved !== activeBuild.build) {
      const file = await loadOptional<TalentFile | null>(dataUrl(saved, `talents/${character.class}.json`), null);
      if (file !== null) return { file, treeVersion: saved };
    }
    return { file: await loadTalents(activeBuild.build, character.class), treeVersion: activeBuild.build };
  }

  function tipsOf(index: TalentIndex): Map<string, { name: string; icon: string; maxRank: number; ranks: string[] }> {
    return new Map(
      [...index.byId.values()].map((talent) => [
        String(talent.id),
        {
          name: talent.name,
          icon: talent.icon,
          maxRank: talent.max_rank,
          ranks: talent.ranks.map((rank) => rank.description),
        },
      ]),
    );
  }

  async function load(): Promise<void> {
    view = { kind: 'loading' };
    try {
      const { file, treeVersion } = await talentFile();
      const index = indexTalents(file);
      if (!talentsFit(character.talents, index)) {
        view = { kind: 'unfit' };
        return;
      }
      setResultTalentTips(tipsOf(index));
      const { order } = orderFromRanks(index, ranksFromTalentsString(character.talents));
      const store = guideTreeStore({
        treeVersion,
        classSlug: character.class,
        raceSlug: character.race,
        talentFile: file,
        order,
      });
      activeTree = fullestTree(character.talents);
      view = { kind: 'ready', store, index, treeVersion };
    } catch (error) {
      console.error('sim talents: the trees did not load', error);
      view = { kind: 'failed' };
    }
  }

  $effect(() => {
    void load();
  });

  const split = $derived(talentSplit(character.talents));
  const plannerHref = $derived(plannerHrefForSpec(character, activeBuild.build));
</script>

<Card
  glyph="tree"
  label={copy.label}
  title={copy.title}
  meta={split}
  body="none"
  class="mx-[18px] md:mx-0"
  testid="sim-talents-card"
>
  <div class="border-line-soft border-t">
    {#if view.kind === 'loading'}
      <div class="p-5"><Skeleton lines={3} rowHeight="h-12" minHeight="min-h-[440px]" testid="sim-talents-skeleton" /></div>
    {:else if view.kind === 'failed'}
      <div class="p-5"><LoadError message={copy.failed} onRetry={() => void load()} testid="sim-talents-error" /></div>
    {:else if view.kind === 'unfit'}
      <p class="text-text px-5 py-4 text-[15px]" data-testid="sim-talents-unfit">{copy.doesNotFit}</p>
    {:else}
      <div class="sim-trees flex flex-col gap-3 py-4 md:p-5" data-testid="sim-talents-trees">
        <TreeTabs
          store={view.store}
          talentIndex={view.index}
          bind:activeTree
          treeColumnsClass={treeRowColumnsClass(view.index.trees.length)}
          readOnly
        />
      </div>
    {/if}
  </div>
  {#if view.kind === 'ready' || view.kind === 'loading'}
    <div class="border-line-soft flex flex-wrap items-center justify-between gap-x-4 border-t px-5 py-1.5">
      <p class="text-muted text-[13px]">
        <span class="md:hidden">{copy.footPhone}</span><span class="hidden md:inline">{copy.footDesktop}</span>
      </p>
      <ActionLink href={plannerHref}>{copy.planner}</ActionLink>
    </div>
  {/if}
</Card>

<style>
  /* Phone tabs read "Fury 32" (spec 3.9): the spec icons leave no room for three names at 360 px. */
  .sim-trees :global([role='tab'] img) {
    display: none;
  }
  /* The saved run's cells (spec 3.9): 40 px from tablet up, a taken talent lit, an untaken one
     greyed with no badge. The planner's and the guides' own cells are untouched. */
  @media (min-width: 768px) {
    .sim-trees :global([role='gridcell']),
    .sim-trees :global([data-game-tip][data-state]) {
      width: 40px;
      height: 40px;
    }
    .sim-trees :global([data-game-tip][data-state] img) {
      width: 36px;
      height: 36px;
    }
  }
  .sim-trees :global([data-game-tip][data-state='maxed']) {
    border: 2px solid var(--color-gold);
  }
  .sim-trees :global([data-game-tip][data-state='filled']) {
    border: 2px solid var(--color-kill);
  }
  .sim-trees :global([data-game-tip][data-state='locked'] img) {
    opacity: 1;
    filter: grayscale(1) brightness(0.45);
  }
  .sim-trees :global([data-game-tip][data-state='locked'] > span) {
    display: none;
  }
</style>
