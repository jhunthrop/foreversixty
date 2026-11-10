<!-- web/src/components/sim/SimTalentsSection.svelte -->
<!-- The character's talents as three flat read-only trees (sim spec 2026-11-06 §4.5 item 4, the
     mock's `Trees`; spec 2026-11-04-sim-results 4.7): an open Section with the split as its meta
     and Open in planner as its action, each tree its name and points over the planner's own grid
     (36 px cells, gold ring full, green ring partial, greyscale untaken, rank badge), below 768
     only each tree's name and points, so the phone reads the split (h-audit 2026-11-09); the game
     tooltip on every cell. It reads the request only, so the landing shows
     it before a run. A string that does not fit today's trees keeps its split and says why nothing
     is drawn. -->
<script lang="ts">
  import activeBuild from '../../data/active-build.json';
  import { guideTreeStore } from '../../lib/guides/tree-store';
  import { orderFromRanks } from '../../lib/planner/fs1';
  import { loadOptional, dataUrl, loadTalents } from '../../lib/planner/load';
  import { indexTalents, type TalentIndex } from '../../lib/planner/rules';
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import type { TalentFile } from '../../lib/planner/types';
  import { plannerHrefForSpec, ranksFromTalentsString } from '../../lib/sim/character';
  import { talentsCardCopy as copy } from '../../lib/sim/input-cards-copy';
  import { setResultTalentTips } from '../../lib/sim/result-tips';
  import { talentsFit } from '../../lib/sim/talents-fit';
  import type { SimRequest } from '../../lib/sim/types';
  import TreeGrid from '../planner/TreeGrid.svelte';
  import LoadError from '../ui/LoadError.svelte';
  import Section from '../ui/Section.svelte';
  import Skeleton from '../ui/Skeleton.svelte';

  let { request }: { request: SimRequest } = $props();
  const character = $derived(request.character);

  type View =
    | { kind: 'loading' }
    | { kind: 'failed' }
    | { kind: 'unfit' }
    | { kind: 'ready'; store: PlannerStore; index: TalentIndex; treeVersion: string };
  let view = $state.raw<View>({ kind: 'loading' });

  /** The build the request names when that build is still published, else today's. */
  async function talentFile(): Promise<{ file: TalentFile; treeVersion: string }> {
    const saved = request.data_build;
    if (saved !== undefined && saved !== activeBuild.build) {
      const file = await loadOptional<TalentFile | null>(
        dataUrl(saved, `talents/${character.class}.json`),
        null,
      );
      if (file !== null) return { file, treeVersion: saved };
    }
    return { file: await loadTalents(activeBuild.build, character.class), treeVersion: activeBuild.build };
  }

  function tipsOf(
    index: TalentIndex,
  ): Map<string, { name: string; icon: string; maxRank: number; ranks: string[] }> {
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
      view = { kind: 'ready', store, index, treeVersion };
    } catch (error) {
      console.error('sim talents: the trees did not load', error);
      view = { kind: 'failed' };
    }
  }

  $effect(() => {
    void load();
  });

  const plannerHref = $derived(plannerHrefForSpec(character, activeBuild.build));
</script>

<Section
  panel
  title={copy.title}
  action={{ label: copy.planner, href: plannerHref }}
  testid="sim-talents-card"
>
  <div>
    {#if view.kind === 'loading'}
      <div>
        <Skeleton lines={3} rowHeight="h-12" minHeight="min-h-[328px]" testid="sim-talents-skeleton" />
      </div>
    {:else if view.kind === 'failed'}
      <div>
        <LoadError message={copy.failed} onRetry={() => void load()} testid="sim-talents-error" />
      </div>
    {:else if view.kind === 'unfit'}
      <p class="text-text text-[15px]" data-testid="sim-talents-unfit">{copy.doesNotFit}</p>
    {:else}
      <div class="sim-trees grid grid-cols-3 gap-[var(--gap)]" data-testid="sim-talents-trees">
        {#each view.index.trees as tree, i (tree.id)}
          <div class="min-w-0" data-testid={`tree-panel-${tree.id}`}>
            <h3 class="text-strong flex items-baseline gap-2 pb-2 text-[15px] leading-5 font-bold">
              {tree.name}
              <span class="text-gold numeric" data-testid={`tree-points-${tree.id}`}
                >{view.store.split[i] ?? 0}</span
              >
            </h3>
            <!-- Below 768 a tree is its name and points (the split); the grid is the planner's to show. -->
            <div class="max-md:hidden"><TreeGrid store={view.store} {tree} readOnly plain /></div>
          </div>
        {/each}
      </div>
    {/if}
  </div>
</Section>

<style>
  /* The mock's `Trees` cells (spec §4.5 item 4): 36 px at every width, the ring outside the icon
     (gold when full, green when partial), an untaken talent greyed with no badge. The planner's
     and the guides' own cells are untouched. */
  .sim-trees :global([role='gridcell']),
  .sim-trees :global([data-game-tip][data-state]) {
    width: 36px;
    height: 36px;
    border-width: 0;
  }
  .sim-trees :global([data-game-tip][data-state] img) {
    width: 36px;
    height: 36px;
  }
  .sim-trees :global([data-game-tip][data-state='maxed'] img) {
    box-shadow: 0 0 0 2px var(--color-gold);
  }
  .sim-trees :global([data-game-tip][data-state='filled'] img) {
    box-shadow: 0 0 0 2px var(--color-kill);
  }
  .sim-trees :global([data-game-tip][data-state='locked'] img) {
    opacity: 1;
    filter: grayscale(1) brightness(0.45);
  }
  .sim-trees :global([data-game-tip][data-state='locked'] > span) {
    display: none;
  }
</style>
