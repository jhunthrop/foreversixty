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
