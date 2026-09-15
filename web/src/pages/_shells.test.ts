// web/src/pages/_shells.test.ts
// The Worker rewrites six head values by selector. If a selector stops matching the
// layout's markup, every unfurl silently freezes at the shell's placeholder text, which no
// other test would notice.
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import Reports from './reports.astro';

describe('the report shell', () => {
  it('carries every data-og hook src/worker.ts rewrites', async () => {
    const container = await AstroContainer.create();
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
    const container = await AstroContainer.create();
    const html = await container.renderToString(Reports);
    expect(html).toContain('data-report-mount');
    expect(html).toContain('id="report"');
    expect(html).toContain('href="/logs"');
  });
});
