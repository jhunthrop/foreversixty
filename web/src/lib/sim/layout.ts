// web/src/lib/sim/layout.ts
// The /sim landing area's loading floor (spec 2026-09-23 §3): while the session read is in
// flight and no character is loaded yet, the landing area shows a Skeleton reserving roughly
// the combined SimRunBlock + LandingState.svelte rendered height for four rows, so the
// reveal changes only opacity (lib/account/layout.ts's own reason, same pattern).
//
// 2026-09-26 layout pass re-measure: the Run block (Finding 1) and the restriction caption
// under "Your characters" (Finding 3) both sit ahead of/inside this same area now, so the
// pre-pass 340px figure undercounts it. A real Chromium build+preview, four characters
// (`[data-testid="sim-run-block"]`'s top to `[data-testid="sim-landing"]`'s bottom), measured
// 790.5px at 390x844 (phone) and 549px at desktop widths -- this constant carries no
// responsive variant (unlike the mount's own min-h in sim.astro, which does), so a single
// number here always over- or under-shoots one breakpoint, the same trade-off the pre-pass
// 340px figure already made. Split the difference rather than reserve the phone worst case
// at every width: 700px, comfortably past either figure's midpoint plus the usual CI-font
// buffer, without reserving so much on desktop that the skeleton-to-content reveal shrinks
// the page. This skeleton's own window (a returning signed-in visitor whose session read is
// slow) is narrow and rare; re-measure with `npm run lhci` if it ever shows up there.
export const SIM_LANDING_SKELETON_MIN_H = 'min-h-[700px]';
