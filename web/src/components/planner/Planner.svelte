<!-- web/src/components/planner/Planner.svelte -->
<!-- The planner island. Mounted two ways: by src/pages/planner.astro with client:load, and
     by src/planner-island.ts on the API-rendered /b/:id page, where `record` is supplied and
     the build starts read-only. Desktop lays the trees out side by side; phone shows one
     tree at a time behind a tab switcher (Task 9). -->
<script lang="ts">
  import { untrack } from 'svelte';
  import { DEFAULT_CLASS_SLUG, ERA_DATA_NOTICE } from '../../lib/planner/config';
  import { DATA_LOAD_FAILED, loadReference, loadSets, loadTalents } from '../../lib/planner/load';
  import { createPlannerStore } from '../../lib/planner/store.svelte';
  import type { BuildRecord } from '../../lib/planner/types';
  import SummaryBar from './SummaryBar.svelte';

  let {
    treeVersion,
    classSlug = DEFAULT_CLASS_SLUG,
    raceSlug,
    record = null,
  }: {
    treeVersion: string;
    classSlug?: string;
    raceSlug?: string;
    record?: BuildRecord | null;
  } = $props();

  // The page is static, so ?class= and ?race= can only be read in the browser. A record
  // (from /b/:id) wins over the query string: that build already names its class and race.
  function fromQuery(name: string): string | undefined {
    if (typeof window === 'undefined') return undefined;
    return new URLSearchParams(window.location.search).get(name) ?? undefined;
  }

  // The store is seeded once, from the props as they arrive. `untrack` says so: without it
  // the compiler warns that these reads only capture the initial value, which is the point --
  // after mount the store owns the class, the race and the order, not the props.
  const store = untrack(() =>
    createPlannerStore({
      treeVersion: record?.tree_version ?? treeVersion,
      classSlug: record ? classSlug : (fromQuery('class') ?? classSlug),
      raceSlug: record ? (raceSlug ?? '') : (fromQuery('race') ?? raceSlug ?? ''),
      order: record?.point_order,
      gear: record?.gear,
      title: record?.title,
      sourceId: record?.id ?? null,
      readOnly: record !== null,
    }),
  );

  let status = $state<'loading' | 'ready' | 'failed'>('loading');
  let attempt = $state(0);

  async function load(): Promise<void> {
    status = 'loading';
    try {
      const reference = await loadReference(store.treeVersion);
      store.setReference(reference);
      store.setTalents(await loadTalents(store.treeVersion, store.classSlug));
      store.setSets(await loadSets(store.treeVersion));
      status = 'ready';
    } catch {
      status = 'failed';
    }
  }

  // Re-runs whenever the class changes (selectClass drops the loaded trees) or Retry bumps
  // `attempt`. Reading both synchronously here is what registers them as dependencies; the
  // writes `load` performs happen after the tracking window, so this cannot loop.
  $effect(() => {
    void store.classSlug;
    void attempt;
    void load();
  });
</script>

<div class="flex flex-col gap-[22px] md:gap-8" data-testid="planner">
  <SummaryBar {store} />

  <p class="text-muted px-[18px] text-[13px] md:px-0">{ERA_DATA_NOTICE}</p>

  {#if status === 'loading'}
    <p class="text-muted px-[18px] text-[14px] md:px-0" role="status">Loading talent data</p>
  {:else if status === 'failed'}
    <div class="border-line bg-raised rounded-panel mx-[18px] flex flex-col gap-3 border p-5 md:mx-0">
      <p class="text-strong text-[15px] font-semibold">{DATA_LOAD_FAILED}</p>
      <!-- Honest for every branch of `load`: the reference files fail into this state too,
           not just talents/<class>.json. -->
      <p class="text-muted text-[13px]">
        Build {store.treeVersion} did not return the files the planner needs.
      </p>
      <button
        type="button"
        class="border-line-warm-strong rounded-control text-text inline-flex h-11 w-fit items-center border px-4 text-[13px] font-bold tracking-[0.06em] uppercase"
        onclick={() => (attempt += 1)}
      >
        Retry
      </button>
    </div>
  {:else if store.talentIndex}
    <div class="grid grid-cols-1 gap-4 px-[18px] md:grid-cols-3 md:px-0" data-testid="tree-columns">
      {#each store.talentIndex.trees as tree, i (tree.id)}
        <section class="border-line bg-raised rounded-panel flex flex-col gap-3 border p-4">
          <header class="flex items-baseline justify-between">
            <h2 class="section-title text-[15px]">{tree.name}</h2>
            <span class="tabular text-gold font-mono text-[15px]">{store.split[i] ?? 0}</span>
          </header>
          <!-- Task 7 renders <TreeGrid {tree} {store} /> here. -->
        </section>
      {/each}
    </div>
  {/if}
</div>
