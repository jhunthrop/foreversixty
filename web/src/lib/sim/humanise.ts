// web/src/lib/sim/humanise.ts
// The one snake_case-to-sentence-case transform this lane uses. Casing only, never
// translation: the engine, the build's own table, or contract 10.8 already chose the
// words -- "rage_gain" reads as "Rage gain", not as anything invented. Shared by
// action-names.ts, buff-names.ts and stats.ts, which each fall back to this whenever
// their own table (spell/item names, the build's simbuffs.json, simCopy.statLabel) does
// not carry an id.
export function humanise(snake: string): string {
  const words = snake.split('_').join(' ');
  return words.charAt(0).toUpperCase() + words.slice(1);
}
