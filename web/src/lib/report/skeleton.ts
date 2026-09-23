// web/src/lib/report/skeleton.ts
// The report page while it loads: the shape of the page it is about to be, drawn in
// placeholder blocks, rather than one line of text above an empty screen. One string,
// used twice -- by the page shell before the island boots (set:html) and by the island
// between mounting and the first byte of data ({@html}) -- so the two moments look the
// same and the layout does not jump between them. Static markup with no data in it.

/** One placeholder block: a bar of the given Tailwind width/height classes. */
const block = (classes: string): string => `<span class="skeleton-block ${classes}"></span>`;

/** A fight-list row: outcome dot, name bar, duration bar. */
const fightRow = (width: string): string =>
  `<li class="flex min-h-11 items-center gap-3 px-2">${block('h-3 w-3 rounded-full')}${block(`h-3 ${width}`)}${block('ml-auto h-3 w-10')}</li>`;

/** A table row: rank, name, bar, two figures. */
const tableRow = (bar: string): string =>
  `<li class="grid min-h-11 grid-cols-[28px_minmax(120px,1.4fr)_minmax(0,3fr)_92px_64px] items-center gap-x-3 border-b border-line-soft px-2 py-2">${block('h-3 w-4')}${block('h-3 w-24')}${block(`h-[6px] ${bar}`)}${block('ml-auto h-3 w-14')}${block('ml-auto h-3 w-10')}</li>`;

/**
 * The skeleton, as the loaded page is laid out: title and subtitle, then the fight list
 * beside the mode bar, the tab strip and a table. Announced once to a screen reader;
 * the blocks themselves are decoration and hidden from it.
 */
export const REPORT_SHELL_TITLE_TESTID = 'report-shell-title';

/** The five characters HTML needs escaped in text content and attributes. */
function escapeHtml(text: string): string {
  return text.replace(
    /[&<>"']/g,
    (c) => ({ '&': '&amp;', '<': '&lt;', '>': '&gt;', '"': '&quot;', "'": '&#39;' })[c] ?? c,
  );
}

/**
 * The skeleton with the report's title already in its heading, when the shell knows it:
 * the prerendered fixture page passes its own, and the Worker fills the same heading
 * through its rewriter (src/worker.ts). The title is the page's largest paint; drawn by
 * the shell it paints on first render instead of after the island hydrates, which is
 * what kept the report page's Lighthouse score at the budget's edge. The island's own
 * loading state uses the bare form (REPORT_SKELETON_HTML) and swaps in the same heading
 * style, so nothing moves.
 */
export function reportSkeletonHtml(title?: string): string {
  const heading =
    title === undefined
      ? `<h1 class="section-title min-h-[26px] text-[18px]" data-testid="${REPORT_SHELL_TITLE_TESTID}">${block('h-5 w-56')}</h1>`
      : `<h1 class="section-title min-h-[26px] text-[18px]" data-testid="${REPORT_SHELL_TITLE_TESTID}">${escapeHtml(title)}</h1>`;
  return [
    '<div class="report-skeleton flex flex-col gap-4 px-[18px] md:px-0" data-testid="report-skeleton" aria-busy="true">',
    '<p class="sr-only" role="status">Loading the report.</p>',
    `<div class="flex flex-col gap-2">${heading}${block('h-3 w-80 max-w-full')}</div>`,
    '<div aria-hidden="true" class="flex flex-col gap-4">',
    '<div class="grid grid-cols-1 gap-[22px] md:grid-cols-[300px_minmax(0,1fr)] md:gap-8">',
    `<ul class="flex flex-col">${['w-28', 'w-36', 'w-24', 'w-32', 'w-28', 'w-36', 'w-20', 'w-32'].map(fightRow).join('')}</ul>`,
    '<div class="flex flex-col gap-4">',
    `<div class="flex flex-wrap gap-2">${['w-20', 'w-24', 'w-20', 'w-16'].map((w) => block(`h-9 ${w} rounded-control`)).join('')}</div>`,
    `<div class="flex flex-wrap gap-3">${['w-16', 'w-24', 'w-24', 'w-16', 'w-14', 'w-16', 'w-16'].map((w) => block(`h-3 ${w}`)).join('')}</div>`,
    `${block('h-24 w-full')}`,
    `<ul class="flex flex-col">${['w-full', 'w-4/5', 'w-2/3', 'w-1/2', 'w-1/3'].map(tableRow).join('')}</ul>`,
    '</div></div></div></div>',
  ].join('');
}

export const REPORT_SKELETON_HTML = reportSkeletonHtml();
