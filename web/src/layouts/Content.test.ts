// The Svelte renderer has to be handed to the container explicitly, exactly as
// _planner.test.ts and _logs.test.ts do: Content.astro wraps Base.astro, which now mounts
// SessionNav client:load on every page, and Astro's integrations are not loaded in a unit
// test, so without this the container throws NoMatchingRenderer rather than rendering an
// empty shell.
import { createPageContainer } from '../test-support/page-container';
import { describe, expect, it } from 'vitest';
import Content from './Content.astro';

describe('Content layout', () => {
  it('shows confidence, updated stamp, and a sources list', async () => {
    const c = await createPageContainer();
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
    expect(html).toContain('Updated 12 Sept');
    expect(html).toContain('Reported by one outlet');
    expect(html).toContain('href="https://guided.news/x"');
    expect(html).toContain('pill-community');
  });

  it('puts the confidence sentence under the title, before Sources, never in a footer line', async () => {
    const c = await createPageContainer();
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

  it('draws no status and no separator for a page with no confidence (the legal pages)', async () => {
    const c = await createPageContainer();
    const html = await c.renderToString(Content, {
      props: {
        title: 'Privacy policy',
        path: '/privacy',
        updated: new Date('2026-10-28T00:00:00Z'),
        sources: [{ label: 'This site', url: 'https://foreversixty.gg/privacy', kind: 'site' }],
        showSources: false,
      },
      slots: { default: '<p>body</p>' },
    });
    expect(html).toContain('Updated 28 Oct');
    expect(html).not.toContain('Confirmed by Blizzard');
    expect(html).not.toContain('Reported by one outlet');
    expect(html).not.toContain('Our inference');
    expect(html).not.toContain('>Sources<');
    expect(html).not.toMatch(/Updated 28 Oct[^<]*<\/[^>]+>\s*<span>·<\/span>/);
  });
});
