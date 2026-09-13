import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import Content from './Content.astro';

describe('Content layout', () => {
  it('shows confidence, updated stamp, and a sources list', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(Content, {
      props: {
        title: 'Hall of Thanes',
        path: '/dungeons/hall-of-thanes',
        updated: new Date('2026-09-12T00:00:00Z'),
        confidence: 'single-source',
        sources: [{ label: 'guided.news roundup', url: 'https://guided.news/x', kind: 'community' }],
      },
      slots: { default: '<p>body</p>' },
    });
    expect(html).toContain('Updated Sept 12');
    expect(html).toContain('Reported by one outlet');
    expect(html).toContain('href="https://guided.news/x"');
    expect(html).toContain('pill-community');
  });
});
