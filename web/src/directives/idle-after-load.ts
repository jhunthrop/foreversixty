// web/src/directives/idle-after-load.ts
// A client directive like Astro's own built-in `client:idle`, except the idle wait only
// starts once the window's `load` event has already fired (or immediately if it has by the
// time this runs). Plain `client:idle`'s `requestIdleCallback` can fire within single-digit
// milliseconds on a page that is not yet doing much else -- exactly the race the day-3
// home-lcp lane report traced the homepage's bimodal Lighthouse LCP to: the same island
// sometimes starts hydrating before the hero's own paint (contending for the same throttled
// mobile bandwidth and main thread) and sometimes after, and the CI median flips between the
// two clusters depending on which one wins that race. Waiting for `load` first means the
// browser has already finished fetching everything the initial render needed -- including
// the LCP element's own paint -- before this directive's idle wait even begins.
import type { ClientDirective } from 'astro';

const idleAfterLoad: ClientDirective = (load) => {
  const hydrate = async (): Promise<void> => {
    const start = await load();
    await start();
  };

  const runWhenIdle = (): void => {
    if ('requestIdleCallback' in window) {
      window.requestIdleCallback(() => void hydrate());
    } else {
      setTimeout(() => void hydrate(), 200);
    }
  };

  if (document.readyState === 'complete') {
    runWhenIdle();
  } else {
    window.addEventListener('load', runWhenIdle, { once: true });
  }
};

export default idleAfterLoad;
