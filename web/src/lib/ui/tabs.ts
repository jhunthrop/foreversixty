// web/src/lib/ui/tabs.ts
// The one tab strip (site alignment spec 2026-10-29 D3, S7): the report's underline tabs as one
// recipe, shared by ui/TabStrip.svelte (in-page tabs and link tabs) and ui/TabStrip.astro (link
// tabs on static pages). A 44 px tab, 12 px 700 .06em caps, a 2 px gold rule under the current
// one over a 1 px soft base line; the strip scrolls sideways with a 24 px fade on a clipped edge.

export interface TabItem {
  id: string;
  label: string;
  /** A link tab's destination, or an in-page tab's hash (it still selects through onSelect). */
  href?: string;
  disabled?: boolean;
  /** A short muted note after the label (" · read-only"). */
  suffix?: string;
  /** A tabular count after the label, muted (a boss's kills). */
  count?: number;
  /** The tab's accessible name when the visible words say less ("Ragnaros, 4 kills"). */
  ariaLabel?: string;
  testid?: string;
  /** Extra attributes on the tab (a page's own hook, such as data-sim-tab). */
  attrs?: Readonly<Record<string, string>>;
}

export const TAB_STRIP = {
  /** The wrapper that carries data-fade-start / data-fade-end (lib/scroll-fade.ts). */
  wrapper: 'group relative',
  strip:
    'flex flex-nowrap overflow-x-auto overscroll-x-contain [scroll-padding-inline:48px] [scrollbar-width:none] ' +
    'shadow-[inset_0_-1px_0_var(--color-line-soft)] [&::-webkit-scrollbar]:hidden',
  /** Wrap mode (web layout spec 2026-11-05 §7.1): below 768 the tabs wrap to a second line, so none is
   *  cut off; from md the strip scrolls as above. */
  stripWrap:
    'flex flex-wrap md:flex-nowrap md:overflow-x-auto md:overscroll-x-contain [scroll-padding-inline:48px] ' +
    '[scrollbar-width:none] shadow-[inset_0_-1px_0_var(--color-line-soft)] [&::-webkit-scrollbar]:hidden',
  /** Sentence case, 14 px 600, as the mock system's strip (design/mocks/ws_prims.py `.strip-tabs`). */
  tab:
    'inline-flex min-h-11 shrink-0 items-center border-b-2 px-[14px] text-[14px] font-semibold ' +
    'whitespace-nowrap transition-colors duration-[120ms] motion-reduce:transition-none ' +
    'focus-visible:outline-2 focus-visible:-outline-offset-2 focus-visible:outline-gold',
  /** Round 5's one tab look (ui/segmented.css too): the current tab white over the cyan accent rule. */
  active: 'border-selected-mark text-strong',
  idle: 'border-transparent text-nav hover:text-strong',
  disabled: 'border-transparent text-muted cursor-not-allowed',
  suffix: 'text-muted ml-1 text-[10px] tracking-normal normal-case',
  count: 'text-muted ml-2 text-[13px] font-normal tabular-nums',
  fadeStart:
    'pointer-events-none absolute inset-y-0 left-0 w-6 bg-gradient-to-r from-bg to-transparent opacity-0 ' +
    'transition-opacity duration-150 motion-reduce:transition-none group-data-[fade-start]:opacity-100',
  fadeEnd:
    'pointer-events-none absolute inset-y-0 right-0 w-6 bg-gradient-to-l from-bg to-transparent opacity-0 ' +
    'transition-opacity duration-150 motion-reduce:transition-none group-data-[fade-end]:opacity-100',
} as const;

/** A tab's classes for its state. */
export function tabClass(current: boolean, disabled: boolean): string {
  const state = disabled ? TAB_STRIP.disabled : current ? TAB_STRIP.active : TAB_STRIP.idle;
  return `${TAB_STRIP.tab} ${state}`;
}

/** The section-nav tab (saved sim spec 2026-11-04 §3.4): the same underline in sentence case, 48 px. */
export function sectionTabClass(current: boolean): string {
  return (
    'inline-flex h-12 shrink-0 items-center border-b-2 px-3 text-[14px] font-semibold whitespace-nowrap ' +
    'transition-colors duration-[120ms] motion-reduce:transition-none focus-visible:outline-2 ' +
    `focus-visible:-outline-offset-2 focus-visible:outline-gold ${current ? TAB_STRIP.active : TAB_STRIP.idle}`
  );
}

/**
 * The tab a key moves to (WAI-ARIA tabs pattern): ArrowRight/ArrowLeft step and wrap, Home and End
 * jump to the ends; a disabled tab is skipped. Null for any other key, or when no tab is enabled.
 */
export function tabIndexForKey(
  key: string,
  from: number,
  tabs: readonly Pick<TabItem, 'disabled'>[],
): number | null {
  const enabled = tabs.flatMap((tab, index) => (tab.disabled === true ? [] : [index]));
  if (enabled.length === 0) return null;
  if (key === 'Home') return enabled[0];
  if (key === 'End') return enabled[enabled.length - 1];
  if (key !== 'ArrowRight' && key !== 'ArrowLeft') return null;
  const step = key === 'ArrowRight' ? 1 : -1;
  const count = tabs.length;
  for (let offset = 1; offset <= count; offset += 1) {
    const index = (((from + step * offset) % count) + count) % count;
    if (tabs[index]?.disabled !== true) return index;
  }
  return null;
}

/** The id a key selects from the current tab (null: the key is not a tab key, let it through). */
export function tabIdForKey(key: string, tabs: readonly TabItem[], active: string): string | null {
  const from = tabs.findIndex((tab) => tab.id === active);
  const next = tabIndexForKey(key, from < 0 ? 0 : from, tabs);
  return next === null ? null : tabs[next].id;
}

/** The edge fades' test ids: the strip's own id with -fade-start / -fade-end. */
export function tabStripFadeTestids(testid: string | undefined): { start: string; end: string } {
  const base = testid ?? 'tab-strip';
  return { start: `${base}-fade-start`, end: `${base}-fade-end` };
}

/** How much of the previous tab shows beside the current one on a strip that scrolls: enough to
 *  read as "there is more this way", never a stub (spec 2026-11-05 §3.6). */
export const TAB_PEEK_PX = 48;

/** Scrolls a strip sideways to where the current tab's predecessor shows `TAB_PEEK_PX`, never moving
 *  the page. The browser clamps at both ends: the first tab rests at 0, and a strip that fits does
 *  not move. */
export function scrollTabIntoView(strip: HTMLElement, tab: HTMLElement): void {
  strip.scrollLeft = tab.previousElementSibling === null ? 0 : Math.max(0, tab.offsetLeft - TAB_PEEK_PX);
}
