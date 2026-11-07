// web/src/lib/sim/skeleton.ts
// A saved sim (/sim/<id>) while it loads: the shape of RunReport.svelte at its ready heights, in
// placeholder blocks (spec 2026-11-04-sim-results 4.2), the same pattern report/skeleton.ts uses
// and for the same reason -- one string, used twice (the page shell before the island boots, via
// set:html, and SimView.svelte between mounting and the saved result's first byte, via {@html}),
// so the two moments look identical and nothing shifts between them. Static markup, no data.
import { PLATE_SURFACE } from '../ui/primitives';
const block = (classes: string): string => `<span class="skeleton-block ${classes}"></span>`;

/** The page head at its ready geometry: the crest, the title and the meta line at the left, the
 *  answer (label, xl figure, band, two Buttons) in its 189 px slot at the right from 900 px. */
const pageHead = (): string =>
  '<div class="px-[var(--page-inset)] pt-4 pb-[14px] min-[900px]:pt-5 min-[900px]:pb-4">' +
  '<div class="grid items-end gap-4 min-[900px]:grid-cols-[minmax(0,1fr)_auto]">' +
  `<div class="flex items-center gap-4">${block('h-14 w-14 shrink-0 !rounded-full')}<div class="flex flex-col gap-2">${block('h-10 w-56')}${block('h-3.5 w-72 max-w-full')}</div></div>` +
  `<div class="flex min-h-[172px] flex-col justify-end gap-2 min-[900px]:min-h-[189px] min-[900px]:min-w-[360px] min-[900px]:items-end">${block('h-3 w-40')}${block('h-[60px] w-56')}${block('h-3.5 w-40')}<div class="flex gap-4">${block('h-11 w-36')}${block('h-11 w-28')}</div></div>` +
  '</div></div>';

/** A Section's head: the title and its meta over the hairline. */
const sectionHead = (): string =>
  `<div class="flex items-baseline gap-3 pb-[10px]">${block('h-5 w-40')}${block('h-3 w-36')}</div>`;

const abilityRow = (): string =>
  `<div class="flex h-14 items-center gap-3 px-4">${block('h-7 w-7 rounded-[6px]')}${block('h-3 w-40')}<span class="ml-auto"></span>${block('hidden h-2 w-[140px] md:block')}${block('h-3 w-12')}${block('hidden h-3 w-40 md:block')}</div>`;

/**
 * The saved page's first screen at its ready geometry (spec 2026-11-06-sim §5 "Saved page"): the
 * head, the Damage by ability Section with the Plate's heads and 6 rows, and the DPS spread's
 * 168 px in the rail. Announced once to a screen reader; the blocks themselves are hidden from it.
 */
export const SIM_SAVED_SKELETON_HTML = [
  '<div class="sim-saved-skeleton flex flex-col" data-testid="sim-saved-skeleton" aria-busy="true">',
  '<p class="sr-only" role="status">Loading the saved sim.</p>',
  '<div aria-hidden="true" class="flex flex-col">',
  pageHead(),
  '<div class="grid items-start gap-[var(--gap)] px-[var(--page-inset)] pt-5 lg:grid-cols-[minmax(0,1fr)_340px] lg:pt-7 min-[1760px]:grid-cols-[minmax(0,1fr)_var(--rail-w)]">',
  `<div class="flex flex-col">${sectionHead()}<div class="${PLATE_SURFACE} overflow-hidden"><div class="h-10"></div>${Array.from({ length: 6 }, abilityRow).join('')}</div></div>`,
  `<div class="flex flex-col">${sectionHead()}<span class="skeleton-block block h-[168px] w-full"></span></div>`,
  '</div></div></div>',
].join('');
