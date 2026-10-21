// web/src/lib/guild/spec-name.ts
// A spec as a player says it. The standing line and the roster rows carry the spec the way the
// export wrote it ("BeastMastery", or a slug "beast-mastery"); the page prints "Beast Mastery".
// Pure string work, with no spec table, so the guild page's first chunk carries no data for it.

/** "BeastMastery" and "beast-mastery" both read "Beast Mastery"; "Holy" stays "Holy"; empty stays empty. */
export function specName(spec: string): string {
  if (spec === '') return spec;
  return spec
    .replace(/([a-z])([A-Z])/g, '$1 $2')
    .split(/[-_\s]+/)
    .filter((word) => word !== '')
    .map((word) => word[0].toUpperCase() + word.slice(1))
    .join(' ');
}

/**
 * A spec with its class, as the character panel says it: "Holy Priest", "Restoration Druid". The
 * class may be a slug ("priest") or a name; with no spec it is the class alone, with no class the
 * spec alone. A spec that already ends in the class ("Holy Priest") is not doubled.
 */
export function specClassName(spec: string, className: string | undefined): string {
  const specWords = specName(spec);
  const classWord = className === undefined ? '' : specName(className);
  if (classWord === '') return specWords;
  if (specWords === '' || specWords.toLowerCase().endsWith(classWord.toLowerCase())) {
    return specWords === '' ? classWord : specWords;
  }
  return `${specWords} ${classWord}`;
}
