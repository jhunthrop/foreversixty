// web/src/lib/bis/follow-selector-client.ts
// Makes a BiS spec page follow the header character selector. The page only has to listen: it
// never reads the pointer at load (the server-rendered page already shows the default faction
// and band), so the work -- the account API, the character lookup, the upgrades loader, about
// 35 modules -- loads on the first pointer change (follow-selector-run.ts) instead of with the
// page. Loaded eagerly it started fetching just before the page's first paint, and the
// Lighthouse trace then counted all of it toward the BiS LCP, which swung between ~1.95 s and
// ~2.7 s run to run against a 1.9 s budget.
import { CURRENT_CHARACTER_STORAGE_KEY, onCurrentCharacterChange } from '../current-character-events';

export function mountBisFollow(root: HTMLElement): () => void {
  return onCurrentCharacterChange((change) => {
    if (change.fromPageLoad) return;
    void import('./follow-selector-run').then(({ follow }) => follow(root));
  });
}

const BAND_HASH = /^#band-(?:alliance|horde)-\d+$/;
/** A `?band=` the page reads (follow-selector.ts bandFromParam's shape), checked here without
 *  importing the spec tables, so this eagerly loaded script stays a few hundred bytes. */
const BAND_PARAM = /^\d{1,2}$/;

/**
 * The band a BiS page opens on when its URL names none: a shared link's `?band=` (or the legacy
 * `#band-<faction>-<band>` hash) is the island's own to read, so this does nothing then; else the
 * selected character's own band and faction. The character lookup waits for the page to finish
 * loading and go idle, for the same LCP reason as mountBisFollow above.
 */
export function applyBisDefaultBand(root: HTMLElement, win: Window = window): void {
  if (BAND_HASH.test(win.location.hash)) return;
  if (BAND_PARAM.test(new URLSearchParams(win.location.search).get('band') ?? '')) return;
  let stored: string | null = null;
  try {
    stored = win.localStorage.getItem(CURRENT_CHARACTER_STORAGE_KEY);
  } catch {
    return;
  }
  if (stored === null) return;
  const run = (): void => {
    void import('./follow-selector-run').then(({ followDefault }) => followDefault(root));
  };
  const idle = (): void => {
    if (typeof win.requestIdleCallback === 'function') win.requestIdleCallback(run, { timeout: 2000 });
    else setTimeout(run, 200);
  };
  if (win.document.readyState === 'complete') idle();
  else win.addEventListener('load', idle, { once: true });
}
