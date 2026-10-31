// web/src/lib/items/drop-chance.ts
// The one way the site prints a drop chance: a whole percent, as the client and Wowhead round
// it, and "<1%" for anything under 1 (so a 0.3% drop never reads "0%"). The tooltip's source
// lines, the BiS source cell and the character page's gear rows all print through it.

/** "20%", "<1%" from a chance already in percent (19.6, 0.3). */
export function dropPercentLabel(percent: number): string {
  return percent < 1 ? '<1%' : `${Math.round(percent)}%`;
}
