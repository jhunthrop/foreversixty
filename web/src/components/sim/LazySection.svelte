<!-- web/src/components/sim/LazySection.svelte -->
<!-- One lazily mounted section of a sim result (spec 2026-11-04-sim-results 3.7, 8; sim spec
     2026-11-06 §9): its chunk loads as the section nears the viewport, and until then the open
     Section's own head stands over a skeleton that holds the section's ready height
     (lib/sim/section-reserve.ts), so nothing below it moves. The reserve goes the moment the
     section renders or fails. A chunk that fails shows the page's one line with Try again; the
     rest of the page stands. The ready chunk draws its own Section (same title and meta).
     `fold`: below 768 px the section starts as one Disclosure line under its title, so a phone
     reads the answer, the ability table, the spread and the gear before the long secondary
     sections (h-audit 2026-11-09); opened, it loads and draws as on desktop. -->
<script lang="ts">
  import { untrack, type Snippet } from 'svelte';
  import type { LazyLoadState } from '../../lib/report/lazy-component.svelte';
  import { pageErrorCopy } from '../../lib/sim/result-copy';
  import type { SectionReserve } from '../../lib/sim/section-reserve';
  import LoadError from '../ui/LoadError.svelte';
  import Section from '../ui/Section.svelte';
  import Skeleton from '../ui/Skeleton.svelte';
  import Disclosure from '../ui/Disclosure.svelte';

  let {
    id,
    title,
    meta,
    lazy,
    reserve = null,
    ready,
    fold = false,
    children,
  }: {
    /** The section's test hook. */
    id: string;
    /** The ready Section's title and meta, drawn while the chunk loads. */
    title: string;
    meta?: string;
    lazy: LazyLoadState;
    /** The section's ready height under and from 768 px, held while it loads; null holds nothing. */
    reserve?: SectionReserve | null;
    /** True once the chunk has resolved. */
    ready: boolean;
    /** Below 768 px, start folded to its title (a secondary section). */
    fold?: boolean;
    children: Snippet;
  } = $props();

  /** The phone breakpoint the fold applies under. */
  const PHONE_QUERY = '(max-width: 767px)';
  /** Read once at mount: a folded section stays a Disclosure for the page's life. */
  const folded =
    untrack(() => fold) && typeof window !== 'undefined' && window.matchMedia(PHONE_QUERY).matches;
  let opened = $state(false);

  /** How far ahead of the viewport the chunk starts to load. */
  const APPROACH_MARGIN = '600px';

  let holder = $state<HTMLElement | undefined>(undefined);
  /** Only a section still loading holds its reserve: rendered or failed, the holder is its section. */
  const held = $derived(!ready && lazy.error === '' && reserve !== null);

  $effect(() => {
    if (holder === undefined) return;
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
    observer.observe(holder);
    return () => observer.disconnect();
  });
</script>

{#snippet section()}
  <div
    bind:this={holder}
    class={['flex min-w-0 flex-col', held && 'reserved']}
    style:--reserve-phone={held ? `${reserve?.phone}px` : undefined}
    style:--reserve-wide={held ? `${reserve?.wide}px` : undefined}
    data-testid={`sim-section-${id}`}
  >
    {#if ready}
      {@render children()}
    {:else}
      <Section {title} {meta} panel>
        {#if lazy.error !== ''}
          <LoadError
            message={pageErrorCopy.partFailed}
            onRetry={() => lazy.load()}
            testid={`sim-section-${id}-error`}
          />
        {:else}
          <Skeleton lines={4} testid={`sim-section-${id}-skeleton`} />
        {/if}
      </Section>
    {/if}
  </div>
{/snippet}

{#if folded}
  <div class="px-[var(--page-inset)]" data-testid={`sim-section-${id}-folded`}>
    <Disclosure summary={title} bind:open={opened} summaryTestid={`sim-section-${id}-fold`}>
      {@render section()}
    </Disclosure>
  </div>
{:else}
  {@render section()}
{/if}

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
