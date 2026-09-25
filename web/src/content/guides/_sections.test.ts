// web/src/content/guides/_sections.test.ts
// Underscore-prefixed so Astro's content loader (glob: '**/*.md') never treats it as a
// guide, the same convention src/pages/_classes.test.ts uses for routes.
//
// Every spec guide renders the same nine `##` sections in the same order (SPEC_SECTIONS),
// so the table of contents SpecTableOfContents.astro renders is identical to the headings
// the page actually has. This asserts that identity directly against the raw markdown body,
// rather than against a rendered page, so it fails on the exact file that is missing or
// misordered a section.
import { getCollection } from 'astro:content';
import { describe, expect, it } from 'vitest';
import { SPEC_SECTIONS } from '../../lib/guides/sections';

const guides = await getCollection('guides');
const specGuides = guides.filter((g) => g.data.spec !== undefined);
const landingPages = guides.filter((g) => g.data.spec === undefined);

describe('guide content structure', () => {
  it('has at least one spec guide and one class landing page', () => {
    expect(specGuides.length).toBeGreaterThan(0);
    expect(landingPages.length).toBeGreaterThan(0);
  });

  it('gives every class a landing page at <class>/index', () => {
    for (const landing of landingPages) {
      expect(landing.id).toBe(`${landing.data.classSlug}/index`);
    }
  });

  it.each(specGuides.map((g) => [g.id, g] as const))(
    '%s renders all nine sections in order',
    (_id, guide) => {
      const headings = [...guide.body!.matchAll(/^##\s+(.+)$/gm)].map((m) => m[1].trim());
      const positions = SPEC_SECTIONS.map((section) => headings.indexOf(section));
      for (const [index, position] of positions.entries()) {
        expect(position, `missing section "${SPEC_SECTIONS[index]}" in ${guide.id}`).toBeGreaterThanOrEqual(0);
      }
      const sorted = [...positions].sort((a, b) => a - b);
      expect(positions, `sections out of order in ${guide.id}`).toEqual(sorted);
    },
  );

  it('gives every spec guide a spec, a role and a classSlug matching its directory', () => {
    for (const guide of specGuides) {
      const [classSlug, spec] = guide.id.split('/');
      expect(guide.data.classSlug).toBe(classSlug);
      expect(guide.data.spec).toBe(spec);
      expect(['dps', 'healer', 'tank']).toContain(guide.data.role);
    }
  });
});
