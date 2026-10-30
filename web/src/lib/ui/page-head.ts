// web/src/lib/ui/page-head.ts
// The page head's recipe (site alignment spec 2026-10-29 §3.1), lifted from /account's head row.
// One table, so PageHead.svelte and its static twin PageHead.astro draw the same box.

export const PAGE_HEAD = {
  /** The head and its `below` row (tabs, a controls row), 16 px apart. */
  root: 'flex flex-col gap-4',
  /** Title block left, the aside right and bottom-aligned from 1024; stacked below it. */
  row: 'flex flex-col gap-4 lg:flex-row lg:items-end lg:justify-between',
  /** The 64 px crest (the `lead`) beside the title block. */
  withLead: 'flex min-w-0 items-center gap-4',
  /** Eyebrow, h1, job line and meta line, 8 px apart. */
  titleBlock: 'flex min-w-0 flex-col gap-2',
  /** A 24 x 1 px gold rule, 12 px gap, then the eyebrow in 11 px 700 .18em caps. */
  eyebrow: 'text-gold flex items-center gap-3 text-[11px] leading-none font-bold tracking-[.18em] uppercase',
  eyebrowRule: 'bg-gold/60 h-px w-6',
  title: 'heading-page',
  /** One job line, 15/22, at most 640 px. */
  job: 'text-text max-w-[640px] text-[15px] leading-[22px]',
  /** The Updated stamp or a descriptor, 13 px muted. */
  meta: 'text-muted text-[13px] leading-5',
  aside: 'min-w-0 lg:shrink-0',
} as const;

/** The h1's classes: the heading face plus a caller's class colour or other title class. */
export function pageHeadTitleClass(titleClass: string | undefined): string {
  return titleClass === undefined || titleClass === '' ? PAGE_HEAD.title : `${PAGE_HEAD.title} ${titleClass}`;
}
