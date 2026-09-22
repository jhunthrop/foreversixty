// web/src/lib/report/layout.ts
// The Skeleton size ReportView.svelte's lazyFallback snippet reserves for each lazily
// imported mode/view while its chunk is still loading, so the panel does not sit blank for
// that one round trip (the audit's own finding). One module the snippet's seven call sites
// share, rather than seven inline literals -- report/skeleton.ts's own reason for being one
// module, applied to a new set of numbers.
export const REPORT_LAZY_MIN_H = {
  rating: 'min-h-[320px]',
  timelines: 'min-h-[420px]',
  events: 'min-h-[360px]',
  queries: 'min-h-[320px]',
  compare: 'min-h-[560px]',
  rankings: 'min-h-[420px]',
  mechanics: 'min-h-[480px]',
} as const;
