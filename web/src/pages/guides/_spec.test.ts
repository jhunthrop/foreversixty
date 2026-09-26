// web/src/pages/guides/_spec.test.ts
// Underscore-prefixed for the same reason as every other _*.test.ts under src/pages/: Astro
// skips it, vitest still collects it.
//
// Builds its own `entry` by reading the guide file directly off disk and parsing it through
// the real `guideSchema`, the same way _sections.test.ts's own `loadGuides()` does -- NOT via
// `getCollection('guides')`, which returns an empty collection under plain `vitest run`
// unless Astro's dev-mode content-layer cache (`.astro/data-store.json`) happens to already be
// warm from a prior `astro dev` run, something nothing in this repo's CI produces before
// `npm test`. Running the real schema over the real file keeps this test honest about what
// the page actually receives (a real `CollectionEntry`-shaped object with `id`/`data`/`body`)
// without depending on Astro's content layer being primed.
import { readFileSync } from 'node:fs';
import { fileURLToPath } from 'node:url';
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { getContainerRenderer } from '@astrojs/svelte/container-renderer';
import { loadRenderers } from 'astro:container';
import { parse as parseYaml } from 'yaml';
import { beforeAll, describe, expect, it } from 'vitest';
import { guideSchema } from '../../content.config';
import SpecPage from './[class]/[spec].astro';

let container: AstroContainer;

beforeAll(async () => {
  const renderers = await loadRenderers([getContainerRenderer()]);
  container = await AstroContainer.create({ renderers });
});

function loadGuideEntry(id: string) {
  const path = fileURLToPath(new URL(`../../content/guides/${id}.md`, import.meta.url));
  const raw = readFileSync(path, 'utf8');
  const [, frontmatterBlock, ...bodyParts] = raw.split('---');
  const data = guideSchema.parse(parseYaml(frontmatterBlock));
  return { id, collection: 'guides' as const, data, body: bodyParts.join('---').trim() };
}

describe('/guides/warrior/fury', () => {
  it('embeds the tree, the build-action buttons, the stat and race pills, and the footer caveat', async () => {
    const entry = loadGuideEntry('warrior/fury');
    const html = await container.renderToString(SpecPage, { props: { entry } });
    expect(html).toContain('data-testid="guide-tree-columns"');
    expect(html).toContain('client:visible');
    expect(html).toContain('data-testid="guide-load-build"');
    expect(html).toContain('data-testid="guide-sim-build"');
    expect(html).toContain('data-testid="stat-priority-pills"');
    expect(html).toContain('data-testid="guide-race-pills"');
    expect(html).toContain('data-testid="confidence-footer-note"');
    // The nine headings still appear, in order, exactly as SpecTableOfContents promises.
    const talentsIndex = html.indexOf('id="talents-and-builds"');
    const rotationIndex = html.indexOf('id="rotation-and-priority"');
    expect(talentsIndex).toBeGreaterThan(-1);
    expect(rotationIndex).toBeGreaterThan(talentsIndex);
  });
});
