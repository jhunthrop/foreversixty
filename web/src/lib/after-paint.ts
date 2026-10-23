// web/src/lib/after-paint.ts
// Runs a page script's work after first paint instead of ahead of it (W.16). A page's
// hoisted `<script>` is a module request chain that Lighthouse's simulation puts ahead of the
// largest contentful paint (the home page paid about 680 ms of simulated LCP for four of
// them), yet each only wires something the markup already works without. The work starts
// when the page has loaded and the browser is idle, after the same 800 ms beat
// src/directives/idle-after-load.ts gives the islands (the captured trace must not start the
// chain ahead of the paint), or at once on the visitor's first pointer, key or focus event,
// so a form or a panel they reach for early is wired by the time they use it. A visitor with
// a session or a stored character starts at `load`: they are the ones the work is for.
const WAKE_EVENTS = ['pointerdown', 'pointermove', 'keydown', 'focusin', 'touchstart'] as const;
const BEAT_AFTER_LOAD_MS = 800;

/** True when the visitor has a session cookie or a stored current character. */
export function hasReturningHint(): boolean {
  if (/(?:^|;\s*)fs_csrf=[^;]+/.test(document.cookie)) return true;
  try {
    return localStorage.getItem('fs.currentCharacter') !== null;
  } catch {
    return false;
  }
}

export function afterPaint(run: () => void): void {
  let done = false;
  const go = (): void => {
    if (done) return;
    done = true;
    for (const name of WAKE_EVENTS) document.removeEventListener(name, go, true);
    run();
  };
  for (const name of WAKE_EVENTS) document.addEventListener(name, go, { capture: true, passive: true });
  const idle = (): void => {
    if ('requestIdleCallback' in window) window.requestIdleCallback(go);
    else setTimeout(go, 200);
  };
  const afterLoad = (): void => {
    if (hasReturningHint()) go();
    else setTimeout(idle, BEAT_AFTER_LOAD_MS);
  };
  if (document.readyState === 'complete') afterLoad();
  else window.addEventListener('load', afterLoad, { once: true });
}
