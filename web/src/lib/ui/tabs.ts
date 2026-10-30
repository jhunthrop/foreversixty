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
  testid?: string;
  /** Extra attributes on the tab (a page's own hook, such as data-sim-tab). */
  attrs?: Readonly<Record<string, string>>;
}

export const TAB_STRIP = {
  /** The wrapper that carries data-fade-start / data-fade-end (lib/scroll-fade.ts). */
  wrapper: 'group relative',
  strip:
    'flex flex-nowrap overflow-x-auto overscroll-x-contain [scrollbar-width:none] ' +
    'shadow-[inset_0_-1px_0_var(--color-line-soft)] [&::-webkit-scrollbar]:hidden',
  tab:
    'inline-flex min-h-11 shrink-0 items-center border-b-2 px-3 text-[12px] font-bold tracking-[0.06em] ' +
    'whitespace-nowrap uppercase transition-colors duration-[120ms] motion-reduce:transition-none ' +
    'focus-visible:outline-2 focus-visible:-outline-offset-2 focus-visible:outline-gold',
  active: 'border-gold text-strong',
  idle: 'border-transparent text-nav hover:text-strong',
  disabled: 'border-transparent text-muted cursor-not-allowed',
  suffix: 'text-muted ml-1 text-[10px] tracking-normal normal-case',
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

/** Scrolls a strip sideways so the current tab is in view, never moving the page. */
export function scrollTabIntoView(strip: HTMLElement, tab: HTMLElement): void {
  const left = tab.offsetLeft;
  const right = left + tab.offsetWidth;
  if (left < strip.scrollLeft) strip.scrollLeft = left;
  else if (right > strip.scrollLeft + strip.clientWidth) strip.scrollLeft = right - strip.clientWidth;
}
