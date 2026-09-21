// web/src/pages/_shells.test.ts
// The Worker rewrites six head values by selector. If a selector stops matching the
// layout's markup, every unfurl silently freezes at the shell's placeholder text, which no
// other test would notice.
//
// The Svelte renderer has to be handed to the container explicitly, exactly as
// _planner.test.ts and _logs.test.ts do: reports.astro wraps Base.astro, which now mounts
// SessionNav client:load on every page, and Astro's integrations are not loaded in a unit
// test, so without this the container throws NoMatchingRenderer rather than rendering an
// empty shell.
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { getContainerRenderer } from '@astrojs/svelte/container-renderer';
import { loadRenderers } from 'astro:container';
import { describe, expect, it } from 'vitest';
import Reports from './reports.astro';

describe('the report shell', () => {
  it('carries every data-og hook src/worker.ts rewrites', async () => {
    const renderers = await loadRenderers([getContainerRenderer()]);
    const container = await AstroContainer.create({ renderers });
    const html = await container.renderToString(Reports);
    for (const selector of [
      'data-og="title"',
      'data-og="description"',
      'data-og="canonical"',
      'data-og="og-title"',
      'data-og="og-description"',
      'data-og="og-url"',
      'data-og="og-image"',
    ]) {
      expect(html, selector).toContain(selector);
    }
  });

  it('carries the island’s mount element and an empty state that names where logs come from', async () => {
    const renderers = await loadRenderers([getContainerRenderer()]);
    const container = await AstroContainer.create({ renderers });
    const html = await container.renderToString(Reports);
    expect(html).toContain('data-report-mount');
    expect(html).toContain('id="report"');
    expect(html).toContain('href="/logs"');
  });
});
