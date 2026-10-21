// web/src/lib/guild/nights.ts
// "1 of 1 night", "5 of 8 nights": one rule for the roster's attendance chip and the loot line
// (lifecycle round 8, F8-3). The noun is singular when the denominator is 1; the loot line keeps
// its longer noun, "raid night".
export function nightsLabel(present: number, nights: number, noun = 'night'): string {
  return `${present} of ${nights} ${noun}${nights === 1 ? '' : 's'}`;
}
