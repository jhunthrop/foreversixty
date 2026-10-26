// web/src/lib/guild/filter-styles.ts
// The guild tabs' filter chip: the Roster's filters and the Readiness next-raid toggle share it.
// Tailwind utilities only (no scoped <style>): these tabs mount after an async fetch, so a scoped
// block's CSS is dropped from the SSR-critical-CSS inlining pass.
export const FILTER_BTN =
  'border-line h-11 rounded-control border px-3 text-[12px] font-semibold text-text lg:h-8';
export const FILTER_BTN_ON = 'border-gold text-gold bg-[rgba(229,185,85,.08)]';
