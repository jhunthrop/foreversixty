// web/src/lib/sim/bulk-skeleton.ts
// What each tool page looks like before its island mounts, in placeholder blocks -- the
// same pattern skeleton.ts and report/skeleton.ts use and for the same reason: one string,
// rendered twice (by the Astro shell via set:html, and by ToolsView between mounting and
// the page's own view chunk resolving), so nothing shifts between the two moments.
//
// Static markup with no data in it, so it is safe to {@html} and identical every render.
import type { SimTool } from './bulk-store.svelte';

// A third verbatim copy of this one-liner (already duplicated between report/skeleton.ts and
// sim/skeleton.ts). Left unshared deliberately: extracting it would mean adding an import
// edge from this lane's file into one of those two part-A modules (or a new shared module
// neither currently has), for one line of string-building neither test suite is at risk of
// drifting on -- not worth the cross-lane coupling for this task.
const block = (classes: string): string => `<span class="skeleton-block ${classes}"></span>`;

/** One slot row of the candidate grid: an icon square, a name bar, two small bars. */
const slotRow = (): string =>
  '<li class="flex min-h-11 items-center gap-3 border-b border-line-soft px-2 py-2">' +
  `${block('h-8 w-8 rounded-control')}${block('h-3 w-32')}${block('ml-auto h-3 w-10')}${block('h-3 w-8')}` +
  '</li>';

/** One ranked results row: rank, chips, figure, delta, percent. */
const comboRow = (): string =>
  '<li class="grid min-h-11 grid-cols-[28px_minmax(120px,2fr)_84px_84px_56px] items-center gap-x-3 border-b border-line-soft px-2 py-2">' +
  `${block('h-3 w-4')}${block('h-3 w-40')}${block('ml-auto h-3 w-14')}${block('ml-auto h-3 w-12')}${block('ml-auto h-3 w-8')}` +
  '</li>';

const shell = (label: string, body: string): string =>
  [
    `<div class="flex flex-col gap-4" aria-busy="true">`,
    `<p class="sr-only" role="status">${label}</p>`,
    '<div aria-hidden="true" class="flex flex-col gap-4">',
    body,
    '</div></div>',
  ].join('');

/** The character strip's reserved band, which every tool page opens with. */
const strip = (): string =>
  '<div class="bg-raised border-line rounded-panel mx-[18px] flex flex-wrap items-center gap-x-6 gap-y-2 border p-4 md:mx-0">' +
  `${block('h-6 w-40')}${block('h-3 w-24')}${block('h-3 w-16')}` +
  '</div>';

const runBar = (): string =>
  '<div class="border-line rounded-panel mx-[18px] flex flex-wrap items-center gap-3 border p-4 md:mx-0">' +
  `${block('h-3 w-36')}${block('h-9 w-28 rounded-control')}${block('h-9 w-24 rounded-control')}` +
  '</div>';

const grid = (count: number): string =>
  `<ul class="mx-[18px] flex flex-col md:mx-0">${Array.from({ length: count }, slotRow).join('')}</ul>`;

const table = (count: number): string =>
  `<ul class="mx-[18px] flex flex-col md:mx-0">${Array.from({ length: count }, comboRow).join('')}</ul>`;

export const TOOL_SKELETONS: Record<SimTool, string> = {
  gear: shell('Loading Top Gear.', [strip(), grid(8), runBar()].join('')),
  talents: shell('Loading talent compare.', [strip(), grid(3), runBar()].join('')),
  drops: shell('Loading the Droptimizer.', [strip(), grid(6), runBar()].join('')),
  weights: shell('Loading stat weights.', [strip(), table(6), runBar()].join('')),
};
