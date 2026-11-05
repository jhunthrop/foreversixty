<!-- web/src/components/sim/SectionNav.svelte -->
<!-- The saved result's sticky section nav (spec 2026-11-04-sim-results 3.4): every card stays open
     on one page, so this only jumps. It is a nav of `#anchor` links on the report's underline, the
     current one found by an IntersectionObserver, with the page's figure at its right end once the
     band's has scrolled out. A click scrolls (instantly under reduced motion) and moves focus to
     the card's h2. Scrolls sideways on a phone with the shared edge fades. -->
<script lang="ts">
  import { scrollFade } from '../../lib/scroll-fade';
  import { bandCopy, navCopy } from '../../lib/sim/result-copy';
  import type { ResultSection } from '../../lib/sim/result-sections';
  import { TAB_STRIP, scrollTabIntoView, sectionTabClass } from '../../lib/ui/tabs';

  let {
    sections,
    figure,
    figureVisible,
  }: {
    sections: readonly ResultSection[];
    /** The page's figure, "904.4". */
    figure: string;
    /** Set once the band's own figure has scrolled out. */
    figureVisible: boolean;
  } = $props();

  let strip = $state<HTMLElement | undefined>(undefined);
  // svelte-ignore state_referenced_locally
  let current = $state<string>(sections[0]?.id ?? '');

  /** A card counts as current when it crosses this share of the viewport from the top. */
  const CURRENT_LINE = 0.3;

  $effect(() => {
    if (typeof IntersectionObserver === 'undefined') return;
    const observer = new IntersectionObserver(
      (entries) => {
        for (const entry of entries) if (entry.isIntersecting) current = entry.target.id;
      },
      { rootMargin: `-${CURRENT_LINE * 100}% 0px -${(1 - CURRENT_LINE) * 100 - 0.1}% 0px` },
    );
    for (const section of sections) {
      const card = document.getElementById(section.id);
      if (card !== null) observer.observe(card);
    }
    return () => observer.disconnect();
  });

  $effect(() => {
    const link = strip?.querySelector<HTMLElement>(`[data-section="${current}"]`);
    if (strip !== undefined && link !== null && link !== undefined) scrollTabIntoView(strip, link);
  });

  /** How long a jump waits for a lazy card's real heading before giving up on moving focus. */
  const HEADING_WAIT_MS = 5000;

  /** Moves focus to the card's h2, waiting for it when the card's chunk is still loading. */
  function focusHeading(card: HTMLElement): void {
    const focus = (): boolean => {
      const heading = card.querySelector<HTMLElement>('h2');
      if (heading === null) return false;
      heading.tabIndex = -1;
      heading.focus({ preventScroll: true });
      return true;
    };
    if (focus()) return;
    const watcher = new MutationObserver(() => {
      if (focus()) watcher.disconnect();
    });
    watcher.observe(card, { childList: true, subtree: true });
    setTimeout(() => watcher.disconnect(), HEADING_WAIT_MS);
  }

  function jump(event: MouseEvent, id: string): void {
    const card = document.getElementById(id);
    if (card === null) return;
    event.preventDefault();
    current = id;
    const smooth = !window.matchMedia('(prefers-reduced-motion: reduce)').matches;
    card.scrollIntoView({ behavior: smooth ? 'smooth' : 'auto', block: 'start' });
    history.replaceState(null, '', `#${id}`);
    focusHeading(card);
  }
</script>

<div
  class="full-bleed border-line-soft sticky top-0 z-[5] border-b bg-[rgba(7,9,13,.94)]"
  data-testid="sim-section-nav-band"
>
  <div class="page-column relative">
    <div class={`${TAB_STRIP.wrapper} lg:pr-[120px]`}>
      <nav
        bind:this={strip}
        use:scrollFade
        class="flex flex-nowrap overflow-x-auto overscroll-x-contain px-[var(--page-inset)] [scrollbar-width:none] [&::-webkit-scrollbar]:hidden"
        aria-label={navCopy.label}
        data-testid="sim-section-nav"
      >
        {#each sections as section (section.id)}
          <a
            href={`#${section.id}`}
            class={sectionTabClass(section.id === current)}
            aria-current={section.id === current ? 'location' : undefined}
            data-section={section.id}
            onclick={(event) => jump(event, section.id)}>{section.label}</a
          >
        {/each}
      </nav>
      <div class={TAB_STRIP.fadeStart} aria-hidden="true"></div>
      <div class={TAB_STRIP.fadeEnd} aria-hidden="true"></div>
    </div>
    <!-- Positioned, so it moves nothing when it fades in; the band's region already states it. -->
    <div
      class={`pointer-events-none absolute inset-y-0 right-0 hidden items-baseline gap-1 self-center pt-[13px] transition-opacity duration-[120ms] motion-reduce:transition-none lg:flex md:right-[var(--page-gutter)] ${figureVisible ? 'opacity-100' : 'opacity-0'}`}
      aria-hidden="true"
      data-testid="sim-section-nav-figure"
    >
      <span class="font-heading text-gold numeric text-[18px] font-extrabold">{figure}</span>
      <span class="text-muted text-[13px]">{bandCopy.dps}</span>
    </div>
  </div>
</div>
