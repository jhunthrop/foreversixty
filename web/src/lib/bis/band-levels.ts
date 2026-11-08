// web/src/lib/bis/band-levels.ts
// The BiS level bands' one definition: they run 20..60 in steps of 10 (sim/cmd/leveling-bis's own
// defaultBandsFlag). No imports, so every BiS module (hover.ts, page-state.ts, band-path.ts) can
// read it without an import cycle.
export const MIN_BAND = 20;
export const MAX_BAND = 60;
export const BAND_STEP = 10;
