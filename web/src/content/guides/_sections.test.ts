// web/src/content/guides/_sections.test.ts
// Underscore-prefixed so Astro's content loader (glob: '**/*.md') never treats it as a
// guide, the same convention src/pages/_classes.test.ts uses for routes.
//
// This reads the guide markdown files straight off disk rather than through
// `getCollection('guides')`: under a plain `vitest run` (no Astro dev/build server behind
// it), `astro:content`'s content layer returns an empty collection for every collection in
// this project, not just this one -- confirmed against `pages` too. Reading the files
// directly and validating frontmatter against the exact same `guideSchema` the real
// collection uses gets the same coverage without depending on that server.
//
// Every spec guide renders the same nine `##` sections in the same order (SPEC_SECTIONS),
// so the table of contents SpecTableOfContents.astro renders is identical to the headings
// the page actually has. This asserts that identity directly against the raw markdown body,
// rather than against a rendered page, so it fails on the exact file that is missing or
// misordered a section.
import { readFileSync, readdirSync } from 'node:fs';
import { dirname, join } from 'node:path';
import { fileURLToPath } from 'node:url';
import { parse as parseYaml } from 'yaml';
import { describe, expect, it } from 'vitest';
import { guideSchema } from '../../content.config';
import { SPEC_SECTIONS, splitSpecSections } from '../../lib/guides/sections';
import { decodeFS1, orderFromRanks } from '../../lib/planner/fs1';
import { indexTalents } from '../../lib/planner/rules';
import { MAX_POINTS, type TalentFile } from '../../lib/planner/types';

const guidesRoot = join(dirname(fileURLToPath(import.meta.url)));

interface LoadedGuide {
  id: string;
  frontmatter: Record<string, unknown>;
  body: string;
}

function loadGuides(): LoadedGuide[] {
  const guides: LoadedGuide[] = [];
  for (const classDir of readdirSync(guidesRoot, { withFileTypes: true })) {
    if (!classDir.isDirectory()) continue;
    for (const file of readdirSync(join(guidesRoot, classDir.name))) {
      if (!file.endsWith('.md')) continue;
      const raw = readFileSync(join(guidesRoot, classDir.name, file), 'utf8');
      const [, frontmatterBlock, ...bodyParts] = raw.split('---');
      guides.push({
        id: `${classDir.name}/${file.slice(0, -'.md'.length)}`,
        frontmatter: parseYaml(frontmatterBlock) as Record<string, unknown>,
        body: bodyParts.join('---'),
      });
    }
  }
  return guides;
}

const guides = loadGuides();
const specGuides = guides.filter((g) => g.frontmatter.spec !== undefined);
const landingPages = guides.filter((g) => g.frontmatter.spec === undefined);

describe('guide content structure', () => {
  it('has at least one spec guide and one class landing page', () => {
    expect(specGuides.length).toBeGreaterThan(0);
    expect(landingPages.length).toBeGreaterThan(0);
  });

  it('gives every class a landing page at <class>/index', () => {
    for (const landing of landingPages) {
      const [classSlug] = landing.id.split('/');
      expect(landing.id).toBe(`${classSlug}/index`);
      expect(landing.frontmatter.classSlug).toBe(classSlug);
    }
  });

  it.each(guides.map((g) => [g.id, g] as const))('%s has schema-valid frontmatter', (id, guide) => {
    const result = guideSchema.safeParse(guide.frontmatter);
    expect(result.success, `${id}: ${result.success ? '' : JSON.stringify(result.error?.issues)}`).toBe(true);
  });

  it.each(specGuides.map((g) => [g.id, g] as const))('%s renders all nine sections in order', (id, guide) => {
    const sections = splitSpecSections(guide.body);
    for (const section of SPEC_SECTIONS) {
      expect(sections.has(section), `missing section "${section}" in ${id}`).toBe(true);
    }
    expect([...sections.keys()], `sections out of order in ${id}`).toEqual(
      SPEC_SECTIONS.filter((section) => sections.has(section)),
    );
  });

  it('gives every spec guide a spec, a role and a classSlug matching its directory', () => {
    for (const guide of specGuides) {
      const [classSlug, spec] = guide.id.split('/');
      expect(guide.frontmatter.classSlug).toBe(classSlug);
      expect(guide.frontmatter.spec).toBe(spec);
      expect(['dps', 'healer', 'tank']).toContain(guide.frontmatter.role);
    }
  });
});

const DATA_BUILD = '1.60.1.69893';
const dataRoot = join(
  dirname(fileURLToPath(import.meta.url)),
  '../../../../data/builds',
  DATA_BUILD,
  'talents',
);
const talentCache = new Map<string, TalentFile>();
function talentsFor(classSlug: string): TalentFile {
  const cached = talentCache.get(classSlug);
  if (cached) return cached;
  const file = JSON.parse(readFileSync(join(dataRoot, `${classSlug}.json`), 'utf8')) as TalentFile;
  talentCache.set(classSlug, file);
  return file;
}

describe('guide build codes', () => {
  it.each(specGuides.map((g) => [g.id, g] as const))(
    '%s has a build: that decodes, names its own class, and is legally reachable within 51 points',
    (id, guide) => {
      const code = guide.frontmatter.build;
      expect(typeof code, `${id}: no build: frontmatter`).toBe('string');
      const decoded = decodeFS1(code as string);
      expect(decoded.ok, `${id}: ${decoded.ok ? '' : decoded.message}`).toBe(true);
      if (!decoded.ok) return;
      expect(decoded.build.classSlug, `${id}: build: is for the wrong class`).toBe(
        guide.frontmatter.classSlug,
      );
      const index = indexTalents(talentsFor(decoded.build.classSlug));
      const { order, dropped } = orderFromRanks(index, decoded.build.treeRanks);
      expect(dropped, `${id}: unreachable talents in build:`).toEqual([]);
      expect(order.length, `${id}: build: spends more than the ${MAX_POINTS}-point cap`).toBeLessThanOrEqual(
        MAX_POINTS,
      );
    },
  );
});
