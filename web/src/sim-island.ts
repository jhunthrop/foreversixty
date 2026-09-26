// web/src/sim-island.ts
// Entry point for dist/sim-island.js, the standalone bundle all three /sim routes load.
// It is a separate Vite build rather than an Astro island for the same reason the report
// island is: /sim/<id> is served by the Worker as a static shell for an id that did not
// exist at build time, so the script tag needs one stable, unhashed URL.
//
// The two stylesheet imports are load-bearing, and src/sim-island.test.ts holds them:
// sim-island.css is the only stylesheet these pages link for the island's own markup, so
// it has to carry the tokens, base rules and self-hosted faces the rest of the site uses.
import { mount } from 'svelte';
import SimView from './components/sim/SimView.svelte';
import { simIdFrom } from './lib/sim/url';
import type { SimResult } from './lib/sim/types';
import './styles/fonts.css';
import './styles/global.css';

const MOUNT_ID = 'sim';

/** A prerendered fixture page inlines its result so Lighthouse measures a real render. */
export function readInlineResult(element: HTMLElement): SimResult | null {
  const raw = element.dataset.simResult;
  if (raw === undefined || raw === '') return null;
  try {
    return JSON.parse(raw) as SimResult;
  } catch (error) {
    console.error('sim island: data-sim-result is not valid JSON', error);
    return null;
  }
}

export function simIdFor(element: HTMLElement, pathname: string): string {
  const fromData = element.dataset.simId;
  if (fromData !== undefined && fromData !== '') return fromData;
  return simIdFrom(pathname);
}

function boot(): void {
  const target = document.getElementById(MOUNT_ID);
  if (target === null) return;
  const props = {
    simId: simIdFor(target, window.location.pathname),
    inlineResult: readInlineResult(target),
  };
  // The shell's no-JS paragraph lives inside the mount element; Svelte 5 appends rather
  // than replaces, so it has to go before the mount or it stays under the island.
  target.replaceChildren();
  // Spec 2026-09-25 §3.6: sim.astro's min-h-[…]/md:min-h-[…] is the pre-hydration shell's
  // own CLS reservation (measured against the static hero-card LCP element it renders
  // before hydration) -- it has no reason to keep applying once SimView has mounted and
  // governs its own height, and left in place it floors the page at that height forever.
  // 2026-09-26 layout pass: bumped to the re-measured 1408px/819px figures (sim.astro's own
  // comment has the numbers) -- fixing, in passing, a stale mismatch this call already had
  // (a prior desktop re-measure moved sim.astro's own class from 607 to 620px without this
  // literal following it, so `md:min-h-[607px]` never actually matched anything in the
  // classList and desktop's reservation was never removed).
  target.classList.remove('min-h-[1408px]', 'md:min-h-[819px]');
  mount(SimView, { target, props });
}

if (document.readyState === 'loading') {
  document.addEventListener('DOMContentLoaded', boot, { once: true });
} else {
  boot();
}
