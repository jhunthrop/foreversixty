<!-- web/src/components/sim/LazySection.svelte -->
<!-- One lazily mounted card of the saved result (spec 2026-11-04-sim-results 3.7, 4.4, 8): its
     chunk loads as the card nears the viewport, and until then a skeleton holds the card's ready
     height (lib/sim/section-reserve.ts) so nothing below it moves. The reserve goes the moment the
     card renders or fails, so the section is always its card's own height after that. A chunk
     that fails shows the page's one line with Try again; the rest of the page stands. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { LazyLoadState } from '../../lib/report/lazy-component.svelte';
  import { pageErrorCopy } from '../../lib/sim/result-copy';
  import type { SectionReserve } from '../../lib/sim/section-reserve';
  import { CARD } from '../../lib/ui/card';
  import type { GlyphName } from '../../lib/ui/glyphs';
  import CardHead from '../ui/CardHead.svelte';
  import LoadError from '../ui/LoadError.svelte';
  import Skeleton from '../ui/Skeleton.svelte';

  let {
    id,
    head,
    lazy,
    reserve = null,
    ready,
    children,
  }: {
    /** The nav anchor. */
    id: string;
    /** The card's head while its chunk loads. An h3 until the real card's h2 arrives, so a nav
     *  jump's focus lands on the heading that stays. */
    head: { glyph: GlyphName; label: string; title: string; meta?: string | undefined };
    lazy: LazyLoadState;
    /** The card's ready height under and from 768 px, held while it loads; null holds nothing. */
    reserve?: SectionReserve | null;
    /** True once the chunk has resolved. */
    ready: boolean;
    children: Snippet;
  } = $props();

  /** How far ahead of the viewport the chunk starts to load. */
  const APPROACH_MARGIN = '600px';

  let section = $state<HTMLElement | undefined>(undefined);
  /** Only a card still loading holds its reserve: rendered or failed, the section is its card. */
  const held = $derived(!ready && lazy.error === '' && reserve !== null);

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

<section
  bind:this={section}
  {id}
  class={['scroll-mt-16', held && 'reserved']}
  style:--reserve-phone={held ? `${reserve?.phone}px` : undefined}
  style:--reserve-wide={held ? `${reserve?.wide}px` : undefined}
  data-testid={`sim-section-${id}`}
>
  {#if ready}
    <div class="reveal">{@render children()}</div>
  {:else if lazy.error !== ''}
    <div class={`${CARD} mx-[18px] md:mx-0`}>
      <CardHead {...head} headingLevel={3} />
      <div class="border-line-soft border-t p-5">
        <LoadError
          message={pageErrorCopy.partFailed}
          onRetry={() => lazy.load()}
          testid={`sim-section-${id}-error`}
        />
      </div>
    </div>
  {:else}
    <div class={`${CARD} mx-[18px] md:mx-0`} style:min-height="inherit">
      <CardHead {...head} headingLevel={3} />
      <div class="border-line-soft border-t p-5">
        <Skeleton lines={4} testid={`sim-section-${id}-skeleton`} />
      </div>
    </div>
  {/if}
</section>

<style>
  .reserved {
    min-height: var(--reserve-phone);
  }
  @media (min-width: 768px) {
    .reserved {
      min-height: var(--reserve-wide);
    }
  }
</style>
