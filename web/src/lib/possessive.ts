// web/src/lib/possessive.ts
// The possessive of a stored name in player-facing copy: `OLYMPUS XXVII's`, but `Employed
// Millennials'` (never `Millennials's`). The name keeps its stored case.

const ENDS_IN_S = /s$/i;

export function possessive(name: string): string {
  return ENDS_IN_S.test(name) ? `${name}'` : `${name}'s`;
}
