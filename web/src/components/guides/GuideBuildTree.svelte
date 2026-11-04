<!-- web/src/components/guides/GuideBuildTree.svelte -->
<!-- The read-only talent tree a spec guide embeds under "Talents and builds" (spec 5):
     decodes the guide's own `build:` FS1 code, fetches that build's talent data, reconstructs
     a legal point order with the planner's own orderFromRanks (lib/guides/tree-store.ts), and
     renders one column per tree through GuideTreeColumns -- never a second tree renderer. States model
     (2026-09-22 spec §1): a skeleton sized to the ready three-column grid, a retry on
     failure, one .reveal on success; nothing moves between them. -->
<script lang="ts">
  import { guidesCopy } from '../../lib/guides/copy';
  import { guideTreeStore, loadGuideTreeSource } from '../../lib/guides/tree-store';
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import LoadError from '../ui/LoadError.svelte';
  import Skeleton from '../ui/Skeleton.svelte';
  import GuideTreeColumns from './GuideTreeColumns.svelte';

  let { code }: { code: string } = $props();

  type Status = 'loading' | 'ready' | 'failed';
  let status = $state<Status>('loading');
  let store = $state<PlannerStore | null>(null);
  let attempt = $state(0);

  async function load(): Promise<void> {
    status = 'loading';
    try {
      store = guideTreeStore(await loadGuideTreeSource(code));
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
     (HomeGuildCard.svelte's own pattern, for the same reason). -->
<span aria-hidden="true"></span>
{#if status === 'loading'}
  <Skeleton lines={3} rowHeight="h-12" minHeight="min-h-[360px]" testid="guide-tree-skeleton" />
{:else if status === 'failed'}
  <LoadError message={guidesCopy.treeLoadFailed} onRetry={retry} testid="guide-tree-error" />
{:else if status === 'ready' && store !== null}
  <GuideTreeColumns {store} />
{/if}
