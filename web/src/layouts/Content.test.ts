// The Svelte renderer has to be handed to the container explicitly, exactly as
// _planner.test.ts and _logs.test.ts do: Content.astro wraps Base.astro, which now mounts
// SessionNav client:load on every page, and Astro's integrations are not loaded in a unit
// test, so without this the container throws NoMatchingRenderer rather than rendering an
// empty shell.
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { getContainerRenderer } from '@astrojs/svelte/container-renderer';
import { loadRenderers } from 'astro:container';
import { describe, expect, it } from 'vitest';
import Content from './Content.astro';

describe('Content layout', () => {
  it('shows confidence, updated stamp, and a sources list', async () => {
    const renderers = await loadRenderers([getContainerRenderer()]);
    const c = await AstroContainer.create({ renderers });
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

  it('defaults confidencePlacement to header, matching every pages-collection entry', async () => {
    const renderers = await loadRenderers([getContainerRenderer()]);
    const c = await AstroContainer.create({ renderers });
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
    const titleIndex = html.indexOf('Hall of Thanes</h1>');
    const sourcesIndex = html.indexOf('>Sources<');
    expect(titleIndex).toBeGreaterThan(-1);
    expect(sourcesIndex).toBeGreaterThan(titleIndex);
    expect(html.slice(titleIndex, sourcesIndex)).toContain('Reported by one outlet');
    expect(html).not.toContain('data-testid="confidence-footer-note"');
  });

  it('moves the confidence sentence to a quiet line above Sources when confidencePlacement is footer', async () => {
    const renderers = await loadRenderers([getContainerRenderer()]);
    const c = await AstroContainer.create({ renderers });
    const html = await c.renderToString(Content, {
      props: {
        title: 'Hall of Thanes',
        path: '/dungeons/hall-of-thanes',
        updated: new Date('2026-09-12T00:00:00Z'),
        confidence: 'single-source',
        sources: [{ label: 'guided.news roundup', url: 'https://guided.news/x', kind: 'community' }],
        confidencePlacement: 'footer',
      },
      slots: { default: '<p>body</p>' },
    });
    const titleIndex = html.indexOf('Hall of Thanes</h1>');
    const bodyEnd = html.indexOf('</header>', titleIndex);
    const footerNoteIndex = html.indexOf('data-testid="confidence-footer-note"');
    const sourcesIndex = html.indexOf('>Sources<');
    expect(html.slice(titleIndex, bodyEnd)).not.toContain('Reported by one outlet');
    expect(footerNoteIndex).toBeGreaterThan(bodyEnd);
    expect(footerNoteIndex).toBeLessThan(sourcesIndex);
    expect(html).toContain('Reported by one outlet');
  });
});
