// web/src/lib/sim/layout.ts
// The /sim landing area's loading floor (spec 2026-09-23 §3): while the session read is in
// flight and no character is loaded yet, the landing area shows a Skeleton reserving roughly
// LandingState.svelte's rendered height for four rows, so the reveal changes only opacity
// (lib/account/layout.ts's own reason, same pattern). Sized by component-height arithmetic
// (the heading, four bordered rows, the scope note, the "Sim something else" link) and
// checked against `npm run lhci`'s /sim CLS before the branch's final report.
export const SIM_LANDING_SKELETON_MIN_H = 'min-h-[340px]';
