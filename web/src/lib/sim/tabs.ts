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
import { defaultSimState, MAX_CODE, simSearch, withSimState, type SimState } from './url';

/** One entry in the tab strip every simulator page carries. */
export interface SimTabEntry {
  readonly id: 'quick-sim' | 'gear' | 'drops' | 'talents' | 'weights' | 'specs';
  readonly label: string;
  readonly href: string;
  /**
   * Whether this tab's own destination can bootstrap a character from `?code=` -- true for
   * `quick-sim` and `specs` (SimView.svelte's store, `store.svelte.ts`'s `fromPlannerCode`
   * branch of its init), false for the four tools-island tabs (ToolsView.svelte /
   * bulk-store.svelte.ts, which bootstrap only from `?source=&ref=` and have never read
   * `?code=` at all). `syncTabHrefs` reads this per tab (fix round 1, Finding A) so a
   * fallback FS1 code -- the only escape hatch for a `ref`-less character -- is never
   * offered to a tab that cannot consume it: that would be exactly the inert-query problem
   * this fallback exists to avoid, just moved rather than fixed. Adding `?code=` support
   * to the tools island is a real feature, not a fix-round-sized change, so it stays out of
   * scope here.
   */
  readonly supportsCode: boolean;
}

/**
 * Exactly the six entries task-1-brief.md names, in its order. `SimTabs.astro` renders
 * this directly; this module's `syncTabHrefs` rewrites the `href` on the anchors it
 * produces once a character is loaded.
 */
export const SIM_TABS: readonly SimTabEntry[] = [
  { id: 'quick-sim', label: simCopy.tabQuickSim, href: '/sim', supportsCode: true },
  { id: 'gear', label: KIND_TITLES.gear, href: '/sim/gear', supportsCode: false },
  { id: 'drops', label: KIND_TITLES.drops, href: '/sim/drops', supportsCode: false },
  { id: 'talents', label: simCopy.tabTalents, href: '/sim/talents', supportsCode: false },
  { id: 'weights', label: simCopy.tabWeights, href: '/sim/weights', supportsCode: false },
  { id: 'specs', label: simCopy.tabSpecs, href: '/sim/specs', supportsCode: true },
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

/**
 * `source` (a loaded character's own `.source`) as the `SimState` the tab strip's hrefs
 * should carry.
 *
 * `source.ref` is empty for an `addon`- or `manual`-kind character (sources.ts,
 * `store.svelte.ts`'s `fromPlannerCode`): neither a pasted export nor a hand-entered FS1
 * code has a server-side ref to round-trip through `?source=&ref=`, and the newcomer
 * persona's addon paste is the product's primary entry path (task-1-brief.md fix round 1,
 * Finding A) -- so `?source=addon` alone would be an inert query that looks like it
 * restores the character and does not. `fallbackCode` is the same escape hatch
 * `store-request.ts`'s own "Run this yourself" link already falls back to for the
 * identical case: a fresh FS1 v2 code derived from the character as currently loaded
 * (`codeForCharacterSpec`), passed in by the caller (never computed here -- deriving one
 * needs an engine-shaped `CharacterSpec`, which only the calling store can build).
 * `fallbackCode` past `MAX_CODE` is treated as unusable, the same as `parseSimState` would
 * silently drop it on the other end: a link this function will not itself honour is not
 * emitted at all, so the strip degrades to bare rather than to another inert query.
 *
 * Falls back to `defaultSimState()` -- every href stays bare -- when there is no character,
 * or no source, or no usable ref and no usable fallback code.
 */
export function tabStateFor(source: TabSource | null, fallbackCode: string | null = null): SimState {
  if (source === null) return defaultSimState();
  if (source.ref !== '') return withSimState(defaultSimState(), { source: source.kind, ref: source.ref });
  if (fallbackCode !== null && fallbackCode !== '' && fallbackCode.length <= MAX_CODE) {
    return withSimState(defaultSimState(), { code: fallbackCode });
  }
  return defaultSimState();
}

/**
 * Rewrites every tab already in the DOM (`SIM_TAB_ATTR` identifies each `<a>`, stamped by
 * `SimTabs.astro`) to carry the loaded character's query, so the strip keeps it across a
 * click to another tool. Called from an `$effect` in both `SimView.svelte` and
 * `ToolsView.svelte` -- the one place either island reacts to its own character/state
 * changes already -- so a link is never stale for longer than a render. A tab missing from
 * the DOM (a page this component is not on) is silently skipped rather than an error: the
 * strip is identical on every page, but `root` may be scoped to less than the whole
 * document in a test.
 *
 * `state` is computed once per tab, not once for the whole strip: `fallbackCode` (see
 * `tabStateFor`) is only ever passed on to a tab whose own `supportsCode` is true, so a tab
 * that cannot bootstrap from `?code=` gets `source`/`ref` or a bare href instead, never a
 * query it cannot itself honour.
 */
export function syncTabHrefs(
  source: TabSource | null,
  fallbackCode: string | null,
  root: ParentNode = document,
): void {
  for (const tab of SIM_TABS) {
    const anchor = root.querySelector<HTMLAnchorElement>(`[${SIM_TAB_ATTR}="${tab.id}"]`);
    if (anchor === null) continue;
    const state = tabStateFor(source, tab.supportsCode ? fallbackCode : null);
    anchor.href = tabHref(tab.href, state);
  }
}
