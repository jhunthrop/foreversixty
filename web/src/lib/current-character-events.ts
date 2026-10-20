// web/src/lib/current-character-events.ts
// How a page hears that the site's "current character" pointer changed, split out of
// current-character.ts so a page that only has to LISTEN (the BiS spec page follows the header
// selector) can attach its listener without bundling the pointer's readers and writers, the
// simulator tab tables and everything they import. Dependency-free on purpose: this file is
// loaded on the first paint's critical path of pages whose real work happens later.
//
// Every read and write is browser-only; importing this during SSR/prerender touches nothing.

/** The localStorage key of the pointer itself (current-character.ts's writers use the same). */
export const CURRENT_CHARACTER_STORAGE_KEY = 'fs.currentCharacter';

/**
 * Dispatched on `window` by a component that has just written a new pointer, for a sibling
 * island on the same page that shows it (CurrentCharacterBar). Islands share no store, and
 * the `storage` event does not fire in the tab that made the change.
 */
export const CURRENT_CHARACTER_CHANGED = 'fs:current-character';

/** What a listener learns about a change. `fromPageLoad` is true when the page itself wrote the
 *  pointer because it just loaded that character (a simulator loader recording what it
 *  loaded): the page already shows it, so a page that follows the selector must not load it
 *  again. */
export interface CurrentCharacterChange {
  fromPageLoad: boolean;
}

/** The `CustomEvent` detail a page-load write carries. */
const PAGE_LOAD_ORIGIN = 'page-load';

function isPageLoadEvent(event: Event): boolean {
  return (
    event instanceof CustomEvent && (event.detail as { origin?: unknown } | null)?.origin === PAGE_LOAD_ORIGIN
  );
}

/** Tells every listener the pointer changed because this page loaded that character. */
export function announcePageLoadedCharacter(): void {
  if (typeof window === 'undefined') return;
  window.dispatchEvent(new CustomEvent(CURRENT_CHARACTER_CHANGED, { detail: { origin: PAGE_LOAD_ORIGIN } }));
}

/**
 * Calls `listener` whenever the pointer changes: in this tab (`CURRENT_CHARACTER_CHANGED`) or
 * in another one (`storage`, filtered to the pointer's own key). Returns the unsubscribe. The
 * one place the listener wiring lives -- every surface that renders the current character
 * subscribes through this and re-reads with `readCurrent`.
 */
export function onCurrentCharacterChange(listener: (change: CurrentCharacterChange) => void): () => void {
  if (typeof window === 'undefined') return () => undefined;
  const onChanged = (event: Event): void => listener({ fromPageLoad: isPageLoadEvent(event) });
  const onStorage = (event: StorageEvent): void => {
    if (event.key === null || event.key === CURRENT_CHARACTER_STORAGE_KEY) listener({ fromPageLoad: false });
  };
  window.addEventListener(CURRENT_CHARACTER_CHANGED, onChanged);
  window.addEventListener('storage', onStorage);
  return () => {
    window.removeEventListener(CURRENT_CHARACTER_CHANGED, onChanged);
    window.removeEventListener('storage', onStorage);
  };
}
