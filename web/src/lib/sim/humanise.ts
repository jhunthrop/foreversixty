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

/**
 * A colon-qualified key with no display name for it to carry -- an id the build's own
 * table never got (or has not loaded yet), or a `drop:` origin whose SourceName never
 * rode along. dps-minmaxer review round 2, D48: a saved run rendered "spell:20662" as an
 * ability name, and the newcomer review's own repro rendered "dungeon:blackrock-spire:
 * 175245" as a boss name -- both are a raw wire key reaching the screen verbatim, which
 * this function exists to end. Every segment humanises in place except a trailing numeric
 * id, which stays a plain number rather than folding into a word: "spell:20662" ->
 * "Spell 20662", the same "Item <id>" shape combos.ts's own substitutionLabel already
 * uses for an item substitution with no name. A key with no colon at all is just
 * humanise() -- there is nothing else to split.
 */
export function humaniseKey(key: string): string {
  return key
    .split(':')
    .map((segment) => (/^\d+$/.test(segment) ? segment : humanise(segment.split('-').join('_'))))
    .join(' ');
}
