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
 * island on the same page that shows it (the header selector). Islands share no store, and
 * the `storage` event does not fire in the tab that made the change.
 */
export const CURRENT_CHARACTER_CHANGED = 'fs:current-character';

/** What a listener learns about a change. `fromPageLoad` is true when the page itself wrote the
 *  pointer because it just loaded that character (a simulator loader recording what it
 *  loaded): the page already shows it, so a page that follows the selector must not load it
 *  again. */
export interface CurrentCharacterChange {
  fromPageLoad: boolean;
  /** True when the session settled its default (the first signed-in answer wrote the account's
   *  main). It moves a page like a choice does, except where that would discard the visitor's
   *  own work (an edited planner, a link's build, a running sim). */
  fromSessionHeal: boolean;
}

/** The `CustomEvent` detail a page-load write carries. */
const PAGE_LOAD_ORIGIN = 'page-load';
const SESSION_HEAL_ORIGIN = 'session-heal';

function hasOrigin(event: Event, origin: string): boolean {
  return event instanceof CustomEvent && (event.detail as { origin?: unknown } | null)?.origin === origin;
}

function dispatchWithOrigin(origin: string): void {
  if (typeof window === 'undefined') return;
  window.dispatchEvent(new CustomEvent(CURRENT_CHARACTER_CHANGED, { detail: { origin } }));
}

/** Tells every listener the pointer changed because this page loaded that character. */
export function announcePageLoadedCharacter(): void {
  dispatchWithOrigin(PAGE_LOAD_ORIGIN);
}

/** Tells every listener the pointer changed because the session settled its default. */
export function announceSessionHeal(): void {
  dispatchWithOrigin(SESSION_HEAL_ORIGIN);
}

/**
 * Calls `listener` whenever the pointer changes: in this tab (`CURRENT_CHARACTER_CHANGED`) or
 * in another one (`storage`, filtered to the pointer's own key). Returns the unsubscribe. The
 * one place the listener wiring lives -- every surface that renders the current character
 * subscribes through this and re-reads with `readCurrent`.
 */
export function onCurrentCharacterChange(listener: (change: CurrentCharacterChange) => void): () => void {
  if (typeof window === 'undefined') return () => undefined;
  const onChanged = (event: Event): void =>
    listener({
      fromPageLoad: hasOrigin(event, PAGE_LOAD_ORIGIN),
      fromSessionHeal: hasOrigin(event, SESSION_HEAL_ORIGIN),
    });
  const onStorage = (event: StorageEvent): void => {
    if (event.key === null || event.key === CURRENT_CHARACTER_STORAGE_KEY)
      listener({ fromPageLoad: false, fromSessionHeal: false });
  };
  window.addEventListener(CURRENT_CHARACTER_CHANGED, onChanged);
  window.addEventListener('storage', onStorage);
  return () => {
    window.removeEventListener(CURRENT_CHARACTER_CHANGED, onChanged);
    window.removeEventListener('storage', onStorage);
  };
}

/** The stored pointer as written (current-character.ts's `parseCurrentCharacter` reads it), or
 *  null when none is stored or storage is out of reach. */
export function readCurrentRaw(): string | null {
  try {
    return globalThis.localStorage?.getItem(CURRENT_CHARACTER_STORAGE_KEY) ?? null;
  } catch {
    return null;
  }
}

/** A choice of character: the stored pointer just before it and just after it. */
export interface CurrentCharacterSwitch {
  readonly previousRaw: string | null;
  readonly nextRaw: string | null;
}

/**
 * `onCurrentCharacterChange` for a page that follows only when it showed the previous character
 * (a guild or character page, character-selector/follow-page.ts): the listener learns the pointer
 * before and after each change. A page-load write and the session settling its default are not
 * choices and call nothing (the page must not leave the address the visitor opened), though each
 * becomes the "before" of the next choice.
 */
export function onCurrentCharacterSwitch(listener: (change: CurrentCharacterSwitch) => void): () => void {
  let previousRaw = readCurrentRaw();
  return onCurrentCharacterChange((change) => {
    const nextRaw = readCurrentRaw();
    const switched = { previousRaw, nextRaw };
    previousRaw = nextRaw;
    if (!change.fromPageLoad && !change.fromSessionHeal) listener(switched);
  });
}
