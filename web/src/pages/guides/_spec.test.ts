// web/src/pages/guides/_spec.test.ts
// Underscore-prefixed for the same reason as every other _*.test.ts under src/pages/: Astro
// skips it, vitest still collects it.
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { getContainerRenderer } from '@astrojs/svelte/container-renderer';
import { loadRenderers } from 'astro:container';
import { getCollection } from 'astro:content';
import { beforeAll, describe, expect, it } from 'vitest';
import SpecPage from './[class]/[spec].astro';

let container: AstroContainer;

beforeAll(async () => {
  const renderers = await loadRenderers([getContainerRenderer()]);
  container = await AstroContainer.create({ renderers });
});

describe('/guides/warrior/fury', () => {
  it('embeds the tree, the build-action buttons, the stat and race pills, and the footer caveat', async () => {
    const guides = await getCollection('guides');
    const entry = guides.find((g) => g.id === 'warrior/fury');
    if (!entry) throw new Error('warrior/fury guide not found');
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
