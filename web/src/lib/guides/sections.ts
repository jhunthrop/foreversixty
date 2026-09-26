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

/**
 * Every `## `-prefixed section a guide's raw markdown body carries, sliced from its own
 * heading line through to (not including) the next `## ` heading, keyed by heading text and
 * insertion-ordered to match the body. `[spec].astro` renders each slice through its own
 * markdown pass and plants a component between two of them; `_sections.test.ts` uses this
 * same split so the page and the test can never drift on what counts as "a section".
 * A heading that is not one of SPEC_SECTIONS is not a key here -- its text stays folded into
 * whichever recognised section precedes it, the same tolerance the order test already had.
 */
export function splitSpecSections(body: string): Map<SpecSection, string> {
  const headingMatches = [...body.matchAll(/^##[ \t]+.+$/gm)];
  const sections = new Map<SpecSection, string>();
  for (const [index, match] of headingMatches.entries()) {
    const heading = match[0].replace(/^##[ \t]+/, '').trim();
    if (!(SPEC_SECTIONS as readonly string[]).includes(heading)) continue;
    const start = match.index ?? 0;
    const end = headingMatches[index + 1]?.index ?? body.length;
    sections.set(heading as SpecSection, body.slice(start, end).trimEnd());
  }
  return sections;
}
