// web/src/lib/sim/lazy-layout.ts
// The Skeleton size SimView.svelte's lazyFallback snippet reserves for CompareView and
// SimResults while their chunk is still loading -- lib/report/layout.ts's own reason,
// mirrored here rather than imported cross-domain (report/skeleton.ts and sim/skeleton.ts
// already keep this same split for the pre-hydration skeleton strings).
export const SIM_LAZY_MIN_H = {
  compare: 'min-h-[360px]',
  results: 'min-h-[420px]',
} as const;
