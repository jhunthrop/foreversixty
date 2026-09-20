// web/src/lib/sim/tabs.ts
// The one tab strip every simulator page carries (task-1-brief.md: the newcomer and
// raid-leader personas both found /sim/gear, /sim/drops, /sim/talents and /sim/weights
// "from nowhere"). SIM_TABS is the single source both `SimTabs.astro` (the server-rendered
// strip; no page can render a partial copy of it) and this module's own client rewrite
// (the strip keeps the loaded character's query on every href) read.
//
// The array itself lives here rather than in copy.ts's Lane W2 block: its three labels
// that are not literal strings (KIND_TITLES.gear/drops, simCopy.tab*) already live in
// copy.ts, satisfying Global Constraint 6 ("every user-visible string"), and keeping the
// composed array out of copy.ts keeps that file from growing past its own 800-line ceiling
// (Global Constraint 10) for a change that is not itself new copy.
import { KIND_TITLES, simCopy } from './copy';
import type { SourceKind } from './types';
import { defaultSimState, simSearch, withSimState, type SimState } from './url';

/** One entry in the tab strip every simulator page carries. */
export interface SimTabEntry {
  readonly id: 'quick-sim' | 'gear' | 'drops' | 'talents' | 'weights' | 'specs';
  readonly label: string;
  readonly href: string;
}

/**
 * Exactly the six entries task-1-brief.md names, in its order. `SimTabs.astro` renders
 * this directly; this module's `syncTabHrefs` rewrites the `href` on the anchors it
 * produces once a character is loaded.
 */
export const SIM_TABS: readonly SimTabEntry[] = [
  { id: 'quick-sim', label: simCopy.tabQuickSim, href: '/sim' },
  { id: 'gear', label: KIND_TITLES.gear, href: '/sim/gear' },
  { id: 'drops', label: KIND_TITLES.drops, href: '/sim/drops' },
  { id: 'talents', label: simCopy.tabTalents, href: '/sim/talents' },
  { id: 'weights', label: simCopy.tabWeights, href: '/sim/weights' },
  { id: 'specs', label: simCopy.tabSpecs, href: '/sim/specs' },
];

/** The attribute `SimTabs.astro` stamps on each `<a>`, and `syncTabHrefs` below reads to
 *  find them again: the strip is server-rendered above the island's own mount point, so a
 *  client rewrite has to look outside the island's own root to reach it. */
export const SIM_TAB_ATTR = 'data-sim-tab';

/**
 * `basePath` (one of `SIM_TABS`' own `href`s) plus `state`'s query, in the one vocabulary
 * `lib/sim/url.ts` owns -- never built by hand (Global Constraint from task-1-brief.md).
 * `simSearch(defaultSimState())` is `''`, so an unloaded page's tabs come back bare with no
 * branch of this function's own.
 */
export function tabHref(basePath: string, state: SimState): string {
  return `${basePath}${simSearch(state)}`;
}

/** The `source`/`ref` half of a loaded character, the only two `SimState` fields either
 *  island's own URL bootstrap ever carries for a plain "keep what's loaded" link. */
export interface TabSource {
  readonly kind: SourceKind;
  readonly ref: string;
}

/** `source` (a loaded character's own `.source`) as the `SimState` the tab strip's hrefs
 *  should carry: `defaultSimState()` -- so every href stays bare -- when there is none. */
export function tabStateFor(source: TabSource | null): SimState {
  return source === null
    ? defaultSimState()
    : withSimState(defaultSimState(), { source: source.kind, ref: source.ref });
}

/**
 * Rewrites every tab already in the DOM (`SIM_TAB_ATTR` identifies each `<a>`, stamped by
 * `SimTabs.astro`) to carry `state`'s query, so the strip keeps the loaded character across
 * a click to another tool. Called from an `$effect` in both `SimView.svelte` and
 * `ToolsView.svelte` -- the one place either island reacts to its own character/state
 * changes already -- so a link is never stale for longer than a render. A tab missing from
 * the DOM (a page this component is not on) is silently skipped rather than an error: the
 * strip is identical on every page, but `root` may be scoped to less than the whole
 * document in a test.
 */
export function syncTabHrefs(state: SimState, root: ParentNode = document): void {
  for (const tab of SIM_TABS) {
    const anchor = root.querySelector<HTMLAnchorElement>(`[${SIM_TAB_ATTR}="${tab.id}"]`);
    if (anchor !== null) anchor.href = tabHref(tab.href, state);
  }
}
