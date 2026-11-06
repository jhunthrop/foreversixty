// web/src/lib/scroll-fade.ts
// A horizontally scrolling strip's edge-fade affordance (polish spec 5.3): the wrapper above the
// scroller carries data-fade-start / data-fade-end while tabs are clipped on that side, and the
// wrapper's own CSS draws the 24px fade. Measured against the first and last child, not the scroll
// width, so trailing padding never keeps the end fade on (the same rule SimTabs.astro follows).

const EPSILON_PX = 1;

export function updateScrollFade(scroller: HTMLElement): void {
  const wrapper = scroller.parentElement;
  const first = scroller.firstElementChild;
  const last = scroller.lastElementChild;
  if (wrapper === null || first === null || last === null) return;
  const edge = scroller.getBoundingClientRect();
  wrapper.toggleAttribute('data-fade-start', first.getBoundingClientRect().left < edge.left - EPSILON_PX);
  wrapper.toggleAttribute('data-fade-end', last.getBoundingClientRect().right > edge.right + EPSILON_PX);
}

/**
 * Scrolls `scroller` along its own axis just far enough to show `tab`, and nothing else: unlike
 * scrollIntoView, the page never moves, so a strip below the fold does not pull the page down to
 * it on load.
 */
export function revealInStrip(scroller: HTMLElement, tab: HTMLElement): void {
  const edge = scroller.getBoundingClientRect();
  const box = tab.getBoundingClientRect();
  if (box.left < edge.left) scroller.scrollLeft -= edge.left - box.left;
  else if (box.right > edge.right) scroller.scrollLeft += box.right - edge.right;
}

/** Svelte action for the scroller: keeps the wrapper's fade attributes current. */
export function scrollFade(scroller: HTMLElement): { destroy: () => void } {
  const update = (): void => updateScrollFade(scroller);
  scroller.addEventListener('scroll', update, { passive: true });
  // Fonts and icons landing after mount resize the tabs without a scroll or window resize.
  const observer = typeof ResizeObserver === 'undefined' ? null : new ResizeObserver(update);
  observer?.observe(scroller);
  for (const child of scroller.children) observer?.observe(child);
  update();
  return {
    destroy: () => {
      scroller.removeEventListener('scroll', update);
      observer?.disconnect();
    },
  };
}
