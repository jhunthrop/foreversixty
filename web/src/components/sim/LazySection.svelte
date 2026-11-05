<!-- web/src/components/sim/LazySection.svelte -->
<!-- One lazily mounted card of the saved result (spec 2026-11-04-sim-results 3.7, 4.4, 8): its
     chunk loads as the card nears the viewport, and until then a skeleton holds the card's ready
     height so nothing below it moves. A chunk that fails shows the page's one line with Try again;
     the rest of the page stands. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { LazyLoadState } from '../../lib/report/lazy-component.svelte';
  import { pageErrorCopy } from '../../lib/sim/result-copy';
  import { CARD } from '../../lib/ui/card';
  import LoadError from '../ui/LoadError.svelte';
  import Skeleton from '../ui/Skeleton.svelte';

  let {
    id,
    lazy,
    reserve = '',
    style = undefined,
    ready,
    children,
  }: {
    /** The nav anchor. */
    id: string;
    lazy: LazyLoadState;
    /** Tailwind min-height classes: the card's ready height at each width. */
    reserve?: string;
    /** The ready height as a style, for a card whose height follows the data. */
    style?: string;
    /** True once the chunk has resolved. */
    ready: boolean;
    children: Snippet;
  } = $props();

  /** How far ahead of the viewport the chunk starts to load. */
  const APPROACH_MARGIN = '600px';

  let section = $state<HTMLElement | undefined>(undefined);

  $effect(() => {
    if (section === undefined) return;
    if (typeof IntersectionObserver === 'undefined') {
      lazy.load();
      return;
    }
    const observer = new IntersectionObserver(
      (entries) => {
        if (entries.some((entry) => entry.isIntersecting)) {
          lazy.load();
          observer.disconnect();
        }
      },
      { rootMargin: APPROACH_MARGIN },
    );
    observer.observe(section);
    return () => observer.disconnect();
  });
</script>

<section bind:this={section} {id} class={`scroll-mt-16 ${reserve}`} {style} data-testid={`sim-section-${id}`}>
  {#if ready}
    <div class="reveal">{@render children()}</div>
  {:else if lazy.error !== ''}
    <div class={`${CARD} mx-[18px] p-5 md:mx-0`}>
      <LoadError message={pageErrorCopy.partFailed} onRetry={() => lazy.load()} testid={`sim-section-${id}-error`} />
    </div>
  {:else}
    <div class={`${CARD} mx-[18px] h-full p-5 md:mx-0`}>
      <Skeleton lines={4} testid={`sim-section-${id}-skeleton`} />
    </div>
  {/if}
</section>
