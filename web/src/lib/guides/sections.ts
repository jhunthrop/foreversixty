// web/src/lib/guides/sections.ts
// The nine `##` sections every spec guide carries, in the fixed order Icy Veins and
// Wowhead readers already expect. SpecTableOfContents.astro and the section-order test
// both read this list, so a section can never go missing from one without the other
// catching it.
export const SPEC_SECTIONS = [
  'Overview',
  'Talents and builds',
  'Rotation and priority',
  'Stat priority',
  'Gear',
  'Enchants and consumables',
  'Races',
  'Professions',
  'Leveling',
] as const;

export type SpecSection = (typeof SPEC_SECTIONS)[number];

/** Mirrors Astro's built-in github-slugger heading ids for the plain-text headings above:
 *  lowercase, non-alphanumerics become a single hyphen, no leading or trailing hyphen. */
export function slugify(heading: string): string {
  return heading
    .toLowerCase()
    .replace(/[^a-z0-9]+/g, '-')
    .replace(/(^-|-$)/g, '');
}
