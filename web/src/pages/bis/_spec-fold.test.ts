// web/src/pages/bis/_spec-fold.test.ts
// Underscore-prefixed so Astro skips it; vitest still collects it.
//
// The SEO fold (2026-11-15): the mid-level BiS bands (lib/bis/band-path.ts isFoldedBand) have no
// page of their own. Their picks are in the spec page's server HTML (components/bis/BisLeveling.astro),
// no band page is generated for them (so the sitemap, which lists built pages, never names them),
// and level 60 keeps its own page.
import type { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { beforeAll, describe, expect, it } from 'vitest';
import SpecPage from './[class]/[spec].astro';
import { getStaticPaths as bandPaths } from './[class]/[spec]/[band].astro';
import { readSpecCatalog } from '../../lib/bis/load';
import { loadFoldedPanels } from '../../lib/bis/leveling-fold';
import { listedBisRows } from '../../lib/guides/guide-bis';
import { createPageContainer } from '../../test-support/page-container';
import activeBuild from '../../data/active-build.json';

let container: AstroContainer;

beforeAll(async () => {
  container = await createPageContainer();
});

const SPECS = [
  ['warrior', 'fury'],
  ['mage', 'frost'],
] as const;

function specEntry(classSlug: string, specSlug: string) {
  const entry = readSpecCatalog().find((s) => s.class_slug === classSlug && s.spec_slug === specSlug);
  if (entry === undefined) throw new Error(`no catalogue entry for ${classSlug}/${specSlug}`);
  return entry;
}

describe.each(SPECS)('/bis/%s/%s', (classSlug, specSlug) => {
  const entry = specEntry(classSlug, specSlug);
  const panels = loadFoldedPanels(activeBuild.build, classSlug, entry.spec);

  it('folds the 30-39, 40-49 and 50-59 bands, both factions', () => {
    expect([...new Set(panels.map((panel) => panel.band))]).toEqual([30, 40, 50]);
    expect(new Set(panels.map((panel) => panel.faction))).toEqual(new Set(['alliance', 'horde']));
  });

  it('server-renders every folded pick in the Leveling section', async () => {
    const html = await container.renderToString(SpecPage, { props: { spec: entry } });
    const start = html.indexOf('data-testid="bis-leveling"');
    expect(start).toBeGreaterThan(-1);
    const section = html.slice(start);
    expect(section).toContain('Leveling 30 to 59');
    for (const panel of panels) {
      expect(section).toMatch(new RegExp(`data-band="${panel.band}"[^>]*data-faction="${panel.faction}"`));
      for (const row of listedBisRows(panel)) {
        if (row.itemId !== undefined && !row.empty)
          expect(section).toContain(`data-bis-item="${row.itemId}"`);
      }
    }
    expect(html).toMatch(
      new RegExp(`rel="canonical" href="https://foreversixty\\.gg/bis/${classSlug}/${specSlug}"`),
    );
  });
});

/** getStaticPaths builds every spec's page data: about 35 s on the test box. */
const ALL_SPECS_TIMEOUT_MS = 120_000;

describe('band pages', () => {
  it(
    'generate level 60 alone: the folded bands 301 to the spec page instead',
    async () => {
      const slugs = new Set((await bandPaths()).map((path) => path.params.band));
      expect(slugs).toEqual(new Set(['60']));
    },
    ALL_SPECS_TIMEOUT_MS,
  );
});
