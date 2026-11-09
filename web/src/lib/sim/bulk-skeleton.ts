// web/src/lib/sim/bulk-skeleton.ts
// What each tool page looks like before its island mounts, in placeholder blocks -- the
// same pattern skeleton.ts and report/skeleton.ts use and for the same reason: one string,
// rendered twice (by the Astro shell via set:html, and by ToolsView between mounting and
// the page's own view chunk resolving), so nothing shifts between the two moments.
//
// Static markup with no data in it, so it is safe to {@html} and identical every render.
import { PLATE_SURFACE } from '../ui/primitives';
import type { SimTool } from './bulk-store.svelte';

// A third verbatim copy of this one-liner (already duplicated between report/skeleton.ts and
// sim/skeleton.ts). Left unshared deliberately: extracting it would mean adding an import
// edge from this lane's file into one of those two part-A modules (or a new shared module
// neither currently has), for one line of string-building neither test suite is at risk of
// drifting on -- not worth the cross-lane coupling for this task.
const block = (classes: string): string => `<span class="skeleton-block ${classes}"></span>`;

/** One placeholder row: an icon square, a name bar, two small bars, on the rows' 58 px. */
const slotRow = (): string =>
  '<li class="flex min-h-[var(--row-h)] items-center gap-3 py-2">' +
  `${block('h-8 w-8 rounded-control')}${block('h-3 w-32')}${block('ml-auto h-3 w-10')}${block('h-3 w-8')}` +
  '</li>';

/**
 * The island's first frames at the page's geometry (mock design/mocks/ws_simtools.py): the run
 * bar's band, then the page column with one section panel (round 5: the head inside the panel,
 * ui/Section `panel`) and its rows.
 */
const shell = (label: string, rows: number): string =>
  [
    '<div aria-busy="true">',
    `<p class="sr-only" role="status">${label}</p>`,
    '<div aria-hidden="true">',
    `<div class="flex h-[var(--bar-h)] items-center gap-3 px-[var(--page-inset)]">${block('h-3 w-36')}${block('ml-auto h-9 w-28 rounded-control')}</div>`,
    '<div class="px-[var(--page-inset)] pt-7">',
    `<div class="${PLATE_SURFACE} flex flex-col gap-3 px-5 pt-4 pb-4 max-md:px-[14px]">`,
    `${block('h-5 w-48')}`,
    `<ul class="flex flex-col">${Array.from({ length: rows }, slotRow).join('')}</ul>`,
    '</div></div></div></div>',
  ].join('');

export const TOOL_SKELETONS: Record<SimTool, string> = {
  gear: shell('Loading Top Gear.', 8),
  talents: shell('Loading talent compare.', 3),
  drops: shell('Loading the Droptimizer.', 6),
  weights: shell('Loading stat weights.', 6),
};
