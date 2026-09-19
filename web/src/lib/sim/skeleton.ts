// web/src/lib/sim/skeleton.ts
// A saved sim (/sim/<id>) while it loads: the shape of SavedSim.svelte, in placeholder
// blocks, the same pattern report/skeleton.ts uses and for the same reason -- one string,
// used twice (the page shell before the island boots, via set:html, and SimView.svelte
// between mounting and the saved result's first byte, via {@html}), so the two moments
// look identical and nothing shifts between them. Static markup with no data in it.
const block = (classes: string): string => `<span class="skeleton-block ${classes}"></span>`;

/** One gear-grid slot: an icon-sized square over a short label bar, CharacterStrip's own shape. */
const gearSlot = (): string =>
  `<li class="flex flex-col items-center gap-1">${block('h-11 w-11 rounded-control')}${block('h-2 w-10')}</li>`;

/** A results-table row: rank, name, bar, two figures -- report/skeleton.ts's own tableRow. */
const tableRow = (bar: string): string =>
  `<li class="grid min-h-11 grid-cols-[28px_minmax(120px,1.4fr)_minmax(0,3fr)_92px_64px] items-center gap-x-3 border-b border-line-soft px-2 py-2">${block('h-3 w-4')}${block('h-3 w-24')}${block(`h-[6px] ${bar}`)}${block('ml-auto h-3 w-14')}${block('ml-auto h-3 w-10')}</li>`;

/**
 * The skeleton, as SavedSim.svelte is laid out: title, the result header card (figure,
 * encounter, run line, engine), the gear grid, the results tab strip and a table.
 * Announced once to a screen reader; the blocks themselves are decoration and hidden from
 * it, the same split report/skeleton.ts uses.
 */
export const SIM_SAVED_SKELETON_HTML = [
  '<div class="sim-saved-skeleton flex flex-col gap-4" data-testid="sim-saved-skeleton" aria-busy="true">',
  '<p class="sr-only" role="status">Loading the saved sim.</p>',
  '<div aria-hidden="true" class="flex flex-col gap-4">',
  `<div class="px-[18px] md:px-0">${block('h-5 w-48')}</div>`,
  '<div class="bg-raised border-line rounded-panel mx-[18px] flex flex-wrap items-center gap-x-6 gap-y-2 border p-4 md:mx-0">' +
    `${block('h-8 w-24')}${block('h-3 w-12')}${block('h-3 w-32')}${block('h-3 w-24')}${block('h-3 w-16')}` +
    '</div>',
  `<ul class="grid grid-cols-2 gap-2 px-[18px] md:grid-cols-4 md:px-0">${Array.from({ length: 8 }, gearSlot).join('')}</ul>`,
  `<div class="flex flex-wrap gap-2 px-[18px] md:px-0">${['w-20', 'w-24', 'w-20', 'w-16'].map((w) => block(`h-9 ${w} rounded-control`)).join('')}</div>`,
  `<ul class="flex flex-col">${['w-full', 'w-4/5', 'w-2/3', 'w-1/2', 'w-1/3'].map(tableRow).join('')}</ul>`,
  '</div></div>',
].join('');
