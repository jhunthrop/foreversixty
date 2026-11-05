// web/src/lib/sim/focus-run.ts
// "Open in the sim" leaves a saved result for /sim with the run's setup loaded (spec
// 2026-11-04-sim-results 3.3). The page never runs the sim for the visitor; it focuses the run
// control, so Enter starts it. The ask crosses the navigation in sessionStorage, read once.
const KEY = 'forever-sim-focus-run';

export function askToFocusRun(store: Pick<Storage, 'setItem'> = sessionStorage): void {
  try {
    store.setItem(KEY, '1');
  } catch (error) {
    // A private window can refuse storage: the page still opens, just without the focus.
    console.warn('sim: could not remember to focus the run control', error);
  }
}

/** True once after `askToFocusRun`. */
export function takeFocusRun(store: Pick<Storage, 'getItem' | 'removeItem'> = sessionStorage): boolean {
  try {
    if (store.getItem(KEY) === null) return false;
    store.removeItem(KEY);
    return true;
  } catch {
    return false;
  }
}
