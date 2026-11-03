// web/src/lib/sim/buff-id.ts
// The suffix an improved raid buff's id carries ("battle_shout:improved"). Its own module, with
// no data import, so buff-names.ts (which the e2e specs load straight into Node) can read it
// without pulling in the generated sim-ids JSON.
export const IMPROVED_SUFFIX = ':improved';
