// web/src/lib/islands/boot.ts
// Runs an island's boot after the shell it replaces has painted, on its own task.
//
// Two reasons, both measured under Lighthouse's mobile profile.
//
// Paint first. Every island page prerenders a static shell of what the island will draw,
// so the page's largest early text exists before any script runs. If the island replaces
// that shell before the browser has painted it -- module scripts run right after parsing,
// often ahead of the first frame -- then the first paint of that text belongs to the
// hydrated copy and Lighthouse charges it to the whole script chain: the simulator's
// intro line swung between an LCP of 2.0 s and 3.3 s on that race alone (2026-09-27, CI
// 0.99 vs 0.92). The swap to the web font counts too, since the re-laid-out text is a new,
// slightly larger paint of the same element, so the boot waits for the fonts the page is
// already loading (capped: a slow connection must not hold the island hostage), then for
// a frame to have been painted.
//
// Own task. A module script's top-level `boot()` runs inside the same task as the bundle's
// evaluation; on the report page that one task is parse + evaluate + a full Svelte mount,
// ~250 ms under the throttled profile, and Total Blocking Time counts every millisecond of
// a task past 50. Yielding once between evaluation and mount splits it.
export type ReadyState = DocumentReadyState;

/** How long the boot waits on `document.fonts.ready` before mounting anyway. */
export const FONT_WAIT_MS = 500;

/** Resolves once the fonts already loading have settled, or after FONT_WAIT_MS, whichever is first. */
function fontsSettled(): Promise<void> {
  const fonts = document.fonts;
  if (fonts === undefined) return Promise.resolve();
  return new Promise((resolve) => {
    const done = (): void => resolve();
    setTimeout(done, FONT_WAIT_MS);
    fonts.ready.then(done, done);
  });
}

/** Runs `run` once a frame has been painted (two frames: the callback runs before its own frame's paint). */
function afterNextPaint(run: () => void): void {
  requestAnimationFrame(() => {
    requestAnimationFrame(run);
  });
}

/** Schedules `boot` for a fresh task once the document has been parsed and its shell has painted. */
export function scheduleBoot(boot: () => void, readyState: ReadyState = document.readyState): void {
  const later = (): void => {
    void fontsSettled().then(() => {
      afterNextPaint(() => {
        setTimeout(boot, 0);
      });
    });
  };
  if (readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', later, { once: true });
  } else {
    later();
  }
}
