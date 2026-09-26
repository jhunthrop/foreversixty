<!-- web/src/components/guides/GuideBuildTree.svelte -->
<!-- The read-only talent tree a spec guide embeds under "Talents and builds" (spec 5):
     decodes the guide's own `build:` FS1 code, fetches that build's talent data, reconstructs
     a legal point order with the planner's own orderFromRanks, and renders one column per
     tree through TreeGrid's own read-only mode -- never a second tree renderer. States model
     (2026-09-22 spec §1): a skeleton sized to the ready three-column grid, a retry on
     failure, one .reveal on success; nothing moves between them. -->
<script lang="ts">
  import { guidesCopy } from '../../lib/guides/copy';
  import { decodeFS1, orderFromRanks } from '../../lib/planner/fs1';
  import { loadTalents } from '../../lib/planner/load';
  import { indexTalents } from '../../lib/planner/rules';
  import { createPlannerStore, type PlannerStore } from '../../lib/planner/store.svelte';
  import LoadError from '../ui/LoadError.svelte';
  import Skeleton from '../ui/Skeleton.svelte';
  import TreeGrid from '../planner/TreeGrid.svelte';

  let { code }: { code: string } = $props();

  type Status = 'loading' | 'ready' | 'failed';
  let status = $state<Status>('loading');
  let store = $state<PlannerStore | null>(null);
  let attempt = $state(0);

  async function load(): Promise<void> {
    status = 'loading';
    const decoded = decodeFS1(code);
    if (!decoded.ok) {
      status = 'failed';
      return;
    }
    try {
      const talentFile = await loadTalents(decoded.build.dataBuild, decoded.build.classSlug);
      const index = indexTalents(talentFile);
      const { order } = orderFromRanks(index, decoded.build.treeRanks);
      const next = createPlannerStore({
        treeVersion: decoded.build.dataBuild,
        classSlug: decoded.build.classSlug,
        raceSlug: decoded.build.raceSlug,
        order,
        readOnly: true,
      });
      next.setTalents(talentFile);
      store = next;
      status = 'ready';
    } catch {
      status = 'failed';
    }
  }

  $effect(() => {
    void attempt;
    void load();
  });

  function retry(): void {
    attempt += 1;
  }
</script>

<!-- Always-present anchor so client:visible's observer has a real element from mount
     (HomeGuildLink.svelte's own pattern, for the same reason). -->
<span aria-hidden="true"></span>
{#if status === 'loading'}
  <Skeleton lines={3} rowHeight="h-12" minHeight="min-h-[360px]" testid="guide-tree-skeleton" />
{:else if status === 'failed'}
  <LoadError message={guidesCopy.treeLoadFailed} onRetry={retry} testid="guide-tree-error" />
{:else if status === 'ready' && store !== null && store.talentIndex !== null}
  {@const talentIndex = store.talentIndex}
  <div class="reveal grid grid-cols-1 gap-4 md:grid-cols-3" data-testid="guide-tree-columns">
    {#each talentIndex.trees as tree (tree.id)}
      <div class="flex flex-col gap-2">
        <h3 class="section-title text-[14px]">{tree.name}</h3>
        <TreeGrid {store} {tree} readOnly />
      </div>
    {/each}
  </div>
{/if}
