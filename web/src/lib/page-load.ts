// web/src/lib/page-load.ts
// Two things a static page's browser script asks before it does signed-in work: is a session
// hinted, and has the page settled. Shared so every such page (the tier list, the BiS landing)
// defers its signed-in answer the same way.

/** `<html data-session="1">`: Base.astro's pre-paint hint that a session cookie exists. */
export function hasSessionHint(): boolean {
  return document.documentElement.dataset.session === '1';
}

/** Runs `work` once the page has loaded and the browser is idle, so the signed-in answer never
 *  competes with the first paint. */
export function afterLoad(work: () => void): void {
  const schedule = (): void => {
    if ('requestIdleCallback' in window) window.requestIdleCallback(work);
    else setTimeout(work, 0);
  };
  if (document.readyState === 'complete') schedule();
  else window.addEventListener('load', schedule, { once: true });
}
