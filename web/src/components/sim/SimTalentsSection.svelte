<!-- web/src/components/sim/SimTalentsSection.svelte -->
<!-- The character's talents (sim spec 2026-11-06 §4.5 item 4; spec 2026-11-04-sim-results 4.7): an
     open Section with Open in planner as its action, holding the site's read-only trees (TalentTrees
     view, talent-tree lane 2026-11-10): the planner's cells, arrows, art and tooltip, the tree tabs
     below 768. It reads the request only, so the landing shows it before a run. A string that does
     not fit today's trees keeps its split and says why nothing is drawn. -->
<script lang="ts">
  import activeBuild from '../../data/active-build.json';
  import { viewTrees } from '../../lib/planner/editor-trees';
  import { loadOptional, dataUrl, loadTalents } from '../../lib/planner/load';
  import { indexTalents } from '../../lib/planner/rules';
  import type { TalentFile } from '../../lib/planner/types';
  import { plannerHrefForSpec } from '../../lib/sim/character';
  import { talentsCardCopy as copy } from '../../lib/sim/input-cards-copy';
  import { talentRanks } from '../../lib/sim/talent-loadouts';
  import { talentsFit } from '../../lib/sim/talents-fit';
  import type { SimRequest } from '../../lib/sim/types';
  import type { TalentTreeView } from '../../lib/ui/talent-tree';
  import LoadError from '../ui/LoadError.svelte';
  import Section from '../ui/Section.svelte';
  import Skeleton from '../ui/Skeleton.svelte';
  import TalentTreesWithTips from '../ui/TalentTreesWithTips.svelte';

  let { request }: { request: SimRequest } = $props();
  const character = $derived(request.character);

  type View =
    | { kind: 'loading' }
    | { kind: 'failed' }
    | { kind: 'unfit' }
    | { kind: 'ready'; trees: TalentTreeView[]; treeVersion: string };
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

  async function load(): Promise<void> {
    view = { kind: 'loading' };
    try {
      const { file, treeVersion } = await talentFile();
      const index = indexTalents(file);
      if (!talentsFit(character.talents, index)) {
        view = { kind: 'unfit' };
        return;
      }
      const trees = viewTrees({
        index,
        ranks: talentRanks(index, character.talents),
        classSlug: character.class,
        asset: (path) => dataUrl(treeVersion, path),
      });
      view = { kind: 'ready', trees, treeVersion };
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
      <TalentTreesWithTips trees={view.trees} build={view.treeVersion} testid="sim-talents-trees" />
    {/if}
  </div>
</Section>
