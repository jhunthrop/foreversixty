// web/src/lib/sim/skeleton.ts
// A saved sim (/sim/<id>) while it loads: the shape of SavedRun.svelte at its ready heights, in
// placeholder blocks (spec 2026-11-04-sim-results 4.2), the same pattern report/skeleton.ts uses
// and for the same reason -- one string, used twice (the page shell before the island boots, via
// set:html, and SimView.svelte between mounting and the saved result's first byte, via {@html}),
// so the two moments look identical and nothing shifts between them. Static markup, no data.
import { CARD } from '../ui/card';

const block = (classes: string): string => `<span class="skeleton-block ${classes}"></span>`;

/** A card head: the tile, the micro label and the title. */
const cardHead = (): string =>
  `<div class="flex min-h-[68px] items-center gap-3 px-5 py-4">${block('h-9 w-9 rounded-control')}<div class="flex flex-col gap-2">${block('h-2.5 w-16')}${block('h-5 w-44')}</div></div>`;

const abilityRow = (): string =>
  `<li class="border-line-soft flex h-14 items-center gap-3 border-t px-5">${block('h-7 w-7 rounded-[6px]')}${block('h-3 w-40')}<span class="ml-auto"></span>${block('h-3 w-12')}${block('hidden h-2 w-[120px] md:block')}${block('hidden h-3 w-40 md:block')}</li>`;

const fact = (): string =>
  `<div class="flex flex-col gap-2 py-3 md:pr-10">${block('h-3 w-14')}${block('h-3.5 w-16')}</div>`;

/** The band at its ready height on the full fixture: 693 phone, 486 tablet, 262 from 1024. */
const band = (): string =>
  '<div class="full-bleed surface-band h-[693px] overflow-hidden md:h-[486px] lg:h-[262px]"><div class="page-column px-[18px] pt-5 md:pt-7">' +
  '<div class="flex flex-col gap-6 lg:flex-row lg:items-end lg:justify-between">' +
  `<div class="flex items-center gap-5">${block('h-14 w-14 shrink-0 !rounded-full md:h-16 md:w-16')}<div class="flex flex-col gap-2">${block('h-2.5 w-20')}${block('h-9 w-52')}${block('h-3 w-64 max-w-full')}${block('h-3.5 w-60 max-w-full')}</div></div>` +
  `<div class="flex flex-col gap-3 lg:items-end">${block('h-14 w-48')}${block('h-3.5 w-36')}<div class="flex flex-col gap-3 md:flex-row">${block('h-11 w-full md:w-44')}${block('h-11 w-full md:w-32')}</div></div>` +
  '</div>' +
  `<div class="mt-6 grid grid-cols-2 border-t border-[rgba(229,185,85,.14)] md:flex md:flex-wrap">${Array.from({ length: 7 }, fact).join('')}</div>` +
  '</div></div>';

const nav = (): string =>
  `<div class="full-bleed border-line-soft h-[49px] border-b"><div class="page-column flex h-12 items-center gap-6 overflow-hidden px-[18px]">${['w-14', 'w-14', 'w-10', 'w-12', 'w-10', 'w-10', 'w-14', 'w-14'].map((width) => block(`h-3 shrink-0 ${width}`)).join('')}</div></div>`;

const card = (body: string, classes = ''): string =>
  `<div class="${CARD} mx-[18px] overflow-hidden md:mx-0 ${classes}">${cardHead()}${body}</div>`;

/**
 * The page's first screens: the band, the nav, the Damage card (64 head + 40 header row + 8 rows
 * of 56 + 44 total), the spread card beside it and the timeline's head and toolbar. Announced once
 * to a screen reader; the blocks themselves are hidden from it.
 */
export const SIM_SAVED_SKELETON_HTML = [
  '<div class="sim-saved-skeleton flex flex-col" data-testid="sim-saved-skeleton" aria-busy="true">',
  '<p class="sr-only" role="status">Loading the saved sim.</p>',
  '<div aria-hidden="true" class="flex flex-col">',
  band(),
  nav(),
  '<div class="flex flex-col gap-[22px] pt-[22px] md:gap-8 md:pt-8">',
  '<div class="grid grid-cols-1 items-start gap-4 xl:grid-cols-[minmax(0,8fr)_minmax(0,4fr)]">',
  card(
    `<div class="border-line-soft h-10 border-t"></div><ul class="flex flex-col">${Array.from({ length: 8 }, abilityRow).join('')}</ul><div class="border-line-soft h-11 border-t"></div>`,
  ),
  card(
    '<div class="border-line-soft border-t p-5"><span class="skeleton-block block h-[168px] w-full"></span></div>',
    'xl:h-[596px]',
  ),
  '</div>',
  card('<div class="border-line-soft h-14 border-t"></div>'),
  '</div></div></div>',
].join('');
