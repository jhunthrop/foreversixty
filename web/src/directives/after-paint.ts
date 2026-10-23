// web/src/directives/after-paint.ts
// A client directive for an island whose server-rendered first frame IS the page's largest
// contentful paint (the guild page's lede): hydrate once the browser has painted, not before.
// `client:load` runs the island's scripts as soon as they arrive, which on a fast connection
// can land ahead of the first paint; Lighthouse's `simulate` throttling then rescales that
// early task 4x, in place, and the paint waits behind it (the guild row's bimodal LCP,
// 1.9 s or 2.9 s on the same build; the home page's day-3 report traced the same race).
// Unlike `client:idle-after-load` there is no fixed 800 ms beat: this page's island fetches
// the guild the visitor came for, so it starts 250 ms after the first contentful paint --
// long enough that the captured trace puts the hydration after the paint every run, short
// enough that a real visitor does not notice it. A tab that never paints (opened in the
// background) hydrates after FALLBACK_MS rather than never.
import type { ClientDirective } from 'astro';

const BEAT_AFTER_PAINT_MS = 250;
const FALLBACK_MS = 3000;

const afterPaint: ClientDirective = (load) => {
  let started = false;
  const hydrate = async (): Promise<void> => {
    if (started) return;
    started = true;
    const start = await load();
    await start();
  };
  const soon = (): void => void setTimeout(() => void hydrate(), BEAT_AFTER_PAINT_MS);

  setTimeout(() => void hydrate(), FALLBACK_MS);
  if (typeof PerformanceObserver === 'undefined') {
    soon();
    return;
  }
  try {
    const observer = new PerformanceObserver((list) => {
      if (list.getEntriesByName('first-contentful-paint').length === 0) return;
      observer.disconnect();
      soon();
    });
    observer.observe({ type: 'paint', buffered: true });
  } catch {
    soon();
  }
};

export default afterPaint;
