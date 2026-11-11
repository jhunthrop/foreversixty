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
import { scheduleBoot } from './lib/islands/boot';
import { shellReservations } from './lib/islands/reservations';
import { simIdFrom } from './lib/sim/url';
import type { SimResult } from './lib/sim/types';
import type { BisSetClass } from './lib/sim/bis-set-rail';
import type { BisBenchmarks } from './lib/sim/bis-benchmark';
import type { PrerenderedNote } from './lib/sim/prerender-notes';
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

/** The landing rail's build-time rows (sim.astro), or none on a page that carries no rail. */
export function readBisSets(element: HTMLElement): BisSetClass[] {
  const raw = element.dataset.bisSets;
  if (raw === undefined || raw === '') return [];
  try {
    return JSON.parse(raw) as BisSetClass[];
  } catch (error) {
    console.error('sim island: data-bis-sets is not valid JSON', error);
    return [];
  }
}

/** The published sets a run is read against (sim.astro, sim/[id].astro), or none. */
export function readBisBenchmarks(element: HTMLElement): BisBenchmarks {
  const raw = element.dataset.bisBenchmarks;
  if (raw === undefined || raw === '') return {};
  try {
    return JSON.parse(raw) as BisBenchmarks;
  } catch (error) {
    console.error('sim island: data-bis-benchmarks is not valid JSON', error);
    return {};
  }
}

/** A prerendered saved run's Worth knowing (sim/[id].astro), or none. */
export function readInitialNotes(element: HTMLElement): PrerenderedNote[] | undefined {
  const raw = element.dataset.simNotes;
  if (raw === undefined || raw === '') return undefined;
  try {
    return JSON.parse(raw) as PrerenderedNote[];
  } catch (error) {
    console.error('sim island: data-sim-notes is not valid JSON', error);
    return undefined;
  }
}

type SpecGridComponent = (typeof import('./components/sim/SpecGrid.svelte'))['default'];

/**
 * /sim/specs's grid, fetched before the shell is replaced. Its shell is the grid itself, and
 * SimView would otherwise fetch the chunk only after mounting, painting its short loading
 * skeleton in between: the footer jumped up into view and back (CLS 0.5 under Lighthouse).
 * Null on every other page, or when the chunk will not load (SimView then fetches it itself
 * and shows its error).
 */
export function loadSpecGrid(target: HTMLElement): Promise<SpecGridComponent | null> {
  if (target.dataset.simView !== 'specs') return Promise.resolve(null);
  return import('./components/sim/SpecGrid.svelte').then(
    (module) => module.default,
    (error: unknown) => {
      console.error('sim island: the spec grid chunk did not load', error);
      return null;
    },
  );
}

function boot(): void {
  const target = document.getElementById(MOUNT_ID);
  if (target === null) return;
  void loadSpecGrid(target).then((specGrid) => mountSim(target, specGrid));
}

function mountSim(target: HTMLElement, specGrid: SpecGridComponent | null): void {
  const props = {
    specGrid,
    simId: simIdFor(target, window.location.pathname),
    inlineResult: readInlineResult(target),
    bisSets: readBisSets(target),
    bisBenchmarks: readBisBenchmarks(target),
    initialNotes: readInitialNotes(target),
  };
  // The shell's no-JS paragraph lives inside the mount element; Svelte 5 appends rather
  // than replaces, so it has to go before the mount or it stays under the island.
  target.replaceChildren();
  // The shells' min-h classes (sim.astro, sim/[id].astro) are the pre-hydration CLS reservation
  // (spec 2026-09-25 §3.6, spec 2026-11-04-sim-results 4.2): once SimView mounts it governs its
  // own height, and a reservation left in place floors the page there forever, which on a short
  // saved result is a hole above the footer. Every min-h class goes, whatever its figure, so a
  // re-measure in either page never needs this line to follow it.
  target.classList.remove(...shellReservations(target));
  mount(SimView, { target, props });
}

scheduleBoot(boot);
