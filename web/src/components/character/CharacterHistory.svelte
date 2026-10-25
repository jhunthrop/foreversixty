<!-- web/src/components/character/CharacterHistory.svelte -->
<!-- The History section of the owner's character page (experience spec
     2026-10-24-character-history 3.2 and 4): the score over time and the parse by boss, each
     with the fights behind it. Mounted only for an owner on a plan with history, and loaded
     lazily, so no other viewer pays for it. It holds the box it will fill from the first
     paint (3.11): the last rendered height is stored after every settled paint. -->
<script lang="ts">
  import { untrack } from 'svelte';
  import type { CharacterPath } from '../../lib/characters';
  import { historyCopy } from '../../lib/character/history-copy';
  import {
    createHistoryController,
    type HistoryController,
  } from '../../lib/character/history-controller.svelte';
  import { COLLAPSED_ROWS } from '../../lib/character/history-select';
  import { reserveToStore, scrollBehavior, writeHistoryReserve } from '../../lib/character/history-layout';
  import type { ChartSize } from '../../lib/character/history-scales';
  import EmptyState from '../ui/EmptyState.svelte';
  import LoadError from '../ui/LoadError.svelte';
  import ParseChart from './ParseChart.svelte';
  import ScoreChart from './ScoreChart.svelte';

  /** The anchor the account page's History row links to. */
  const SECTION_ID = 'history';
  /** The skeleton list draws as many rows as a list shows collapsed. */
  const LIST_ROWS = Array.from({ length: COLLAPSED_ROWS }, (_, row) => row);

  let {
    path,
    controller = undefined,
    sizeOverride = null,
  }: {
    path: CharacterPath;
    /** The read of the history; the section makes its own for `path` when none is given. */
    controller?: HistoryController;
    /** A fixed chart size, for tests and captures. */
    sizeOverride?: ChartSize | null;
  } = $props();

  const source = untrack(() => controller ?? createHistoryController(path));
  let section: HTMLElement | undefined = $state();

  $effect(() => {
    if (controller === undefined) untrack(() => source.load());
  });

  // The account page's "Open on your main character" link ends in #history, and the section
  // mounts after /v1/me and the page: the browser's own anchor jump has already missed it.
  let landed = false;
  $effect(() => {
    if (section === undefined || landed) return;
    landed = true;
    if (window.location.hash === `#${SECTION_ID}`) {
      section.scrollIntoView({ behavior: scrollBehavior() });
    }
  });

  // The collapsed, settled box's rendered height, for the next visit's reserve (reserveToStore).
  $effect(() => {
    const element = section;
    if (element === undefined || source.status !== 'ready') return;
    const observer = new ResizeObserver(() => {
      const listOpen = element.querySelector('[aria-expanded="true"]') !== null;
      const height = reserveToStore(source.status, listOpen, element.offsetHeight);
      if (height !== null) writeHistoryReserve(height);
    });
    observer.observe(element);
    return () => observer.disconnect();
  });
</script>

<section
  bind:this={section}
  id={SECTION_ID}
  class="-order-2 flex flex-col gap-2"
  class:history-reserve={source.status === 'loading'}
  aria-busy={source.status === 'loading' ? 'true' : undefined}
  data-testid="character-history-section"
>
  <h2 class="section-title text-[18px]">{historyCopy.heading}</h2>
  <p class="text-muted text-[13px]" data-testid="history-caption">{historyCopy.caption}</p>
  <div class="bg-raised border-line rounded-panel border p-4 md:p-5">
    {#if source.status === 'loading'}
      <div class="flex flex-col gap-7 xl:grid xl:grid-cols-2 xl:gap-x-8" data-testid="history-skeleton">
        <span class="sr-only" role="status">{historyCopy.loadingSr}</span>
        {#each [0, 1] as column (column)}
          <div class="flex flex-col gap-3" aria-hidden="true">
            <!-- The ready column's own rows at its own heights: sub-heading, control row,
                 readout, chart box, the list label and eight list rows. -->
            <span class="skeleton-block h-[15px] w-[140px]"></span>
            <span class="skeleton-block h-11 w-3/4"></span>
            <span class="skeleton-block h-[28px] w-[120px]"></span>
            <span class="skeleton-block aspect-[3/2] w-full sm:aspect-[16/5] xl:aspect-[2/1]"></span>
            <span class="skeleton-block h-4 w-24"></span>
            <span class="skeleton-block h-4 w-1/3"></span>
            <span class="skeleton-block h-9 w-full"></span>
            {#each LIST_ROWS as row (row)}
              <span class="skeleton-block h-[49px] w-full"></span>
            {/each}
            <span class="skeleton-block h-11 w-32"></span>
          </div>
        {/each}
      </div>
    {:else if source.status === 'error' || source.history === null}
      <LoadError message={historyCopy.loadFailed} onRetry={() => source.load()} testid="history-error" />
    {:else if source.history.score.length === 0}
      <EmptyState message={historyCopy.empty} testid="history-empty" />
    {:else}
      <div
        class="flex flex-col gap-7 xl:grid xl:grid-cols-2 xl:gap-x-8 xl:gap-y-0"
        data-testid="history-columns"
      >
        <ScoreChart history={source.history} {sizeOverride} />
        <ParseChart history={source.history} {sizeOverride} />
      </div>
    {/if}
  </div>
</section>
