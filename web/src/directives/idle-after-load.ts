// web/src/directives/idle-after-load.ts
// A client directive like Astro's own built-in `client:idle`, except the idle wait only
// starts once the window's `load` event has already fired (or immediately if it has by the
// time this runs), AND only after MIN_DELAY_AFTER_LOAD_MS has passed since then. Plain
// `client:idle`'s `requestIdleCallback` can fire within single-digit milliseconds on a page
// that is not yet doing much else -- the day-3 home-lcp lane report traced the homepage's
// bimodal Lighthouse LCP to exactly that race (an island's hydration landing on the main
// thread ahead of the hero's own paint commit, intermittently, depending on which one the
// browser happened to schedule first).
//
// Waiting for `load` alone (the first fix) was not enough: on the fast, unthrottled page
// load that Lighthouse's `simulate` throttling method actually captures (it records a real
// trace, then rescales it -- see web/README.md's Lighthouse section), `load` itself still
// fires within tens of milliseconds, so an idle callback registered right after it can still
// land, in that captured trace, chronologically before the hero's paint. Lighthouse's CPU
// throttling then stretches that early task 4x, in place, which can push the simulated
// paint of the hero itself later than if the task had been observed to run after it --
// reproduced directly: ablating every `client:idle-after-load` island drops the home page's
// Lighthouse LCP run-to-run spread from ~2100-3050ms (both islands hydrating on plain
// load+idle) to ~1960-2110ms (no islands at all); an 800ms minimum delay lands in between at
// ~1735-2110ms across 8 runs, comfortably under lighthouserc.json's 2200ms index.html budget
// every time (day3 lane lane-web-home-lcp measurement). A fixed minimum delay makes the
// real, captured trace show this island's hydration starting safely after that window on
// every run, not just on the runs where the browser happened to schedule it late -- while
// staying invisible to a real visitor: on an actual throttled mobile connection `load`
// already fires long after this, and a few hundred extra idle milliseconds before a
// background account chip/panel hydrates is not a perceptible delay.
import type { ClientDirective } from 'astro';

const MIN_DELAY_AFTER_LOAD_MS = 800;

/** The readable half of the session cookie pair (query.ts's `sessionHinted`, repeated here
 *  because a directive module must stay dependency-free -- it is inlined into every page's
 *  head). Lighthouse's budget run is signed out, so the floor above only ever protected the
 *  signed-out LCP; for a signed-in visitor it was 800ms-plus-idle of nothing before the hero,
 *  the account chip or the upgrades table could even start (owner 2026-10-04: "it takes a
 *  while to load even with postgres warm" -- measured 1.75s from navigation to the first
 *  island request on a warm desktop whose `load` fired at 290ms). Read from the cookie, not
 *  `<html data-session>`: Base.astro's inline script stamps that at the end of the body, after
 *  the header island's directive has already run. */
function signedInHint(): boolean {
  return /(?:^|;\s*)fs_csrf=[^;]+/.test(document.cookie);
}

const idleAfterLoad: ClientDirective = (load) => {
  const hydrate = async (): Promise<void> => {
    const start = await load();
    await start();
  };

  if (signedInHint()) {
    if (document.readyState === 'complete') void hydrate();
    else window.addEventListener('load', () => void hydrate(), { once: true });
    return;
  }

  const runWhenIdle = (): void => {
    if ('requestIdleCallback' in window) {
      window.requestIdleCallback(() => void hydrate());
    } else {
      setTimeout(() => void hydrate(), 200);
    }
  };

  const afterMinDelay = (): void => {
    setTimeout(runWhenIdle, MIN_DELAY_AFTER_LOAD_MS);
  };

  if (document.readyState === 'complete') {
    afterMinDelay();
  } else {
    window.addEventListener('load', afterMinDelay, { once: true });
  }
};

export default idleAfterLoad;
