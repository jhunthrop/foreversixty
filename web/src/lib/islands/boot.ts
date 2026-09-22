// web/src/lib/islands/boot.ts
// Runs an island's boot on its own task, after the task that evaluated its module.
//
// A module script runs once the document has been parsed, so `document.readyState` is
// already "interactive" and a `boot()` called from the top level runs inside the same task
// as the bundle's evaluation. On the report page that one task is parse + evaluate + a
// full Svelte mount of the landing view, ~100 ms on a laptop and ~250 ms under Lighthouse's
// throttled CI profile; Total Blocking Time counts every millisecond of a task past 50, so
// one long task costs more than two shorter ones doing the same work. Yielding once between
// evaluation and mount splits it, and costs the first paint a single timer tick.
export type ReadyState = DocumentReadyState;

/** Schedules `boot` for a fresh task once the document has been parsed. */
export function scheduleBoot(boot: () => void, readyState: ReadyState = document.readyState): void {
  const later = (): void => {
    setTimeout(boot, 0);
  };
  if (readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', later, { once: true });
  } else {
    later();
  }
}
