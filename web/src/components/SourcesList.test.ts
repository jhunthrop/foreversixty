import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import SourcesList from './SourcesList.astro';

describe('SourcesList', () => {
  it('renders one labelled, pilled link per source', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(SourcesList, {
      props: {
        sources: [
          {
            label: 'BlizzCon announcement',
            url: 'https://worldofwarcraft.blizzard.com/forever',
            kind: 'blizzard',
          },
          { label: 'guided.news roundup', url: 'https://guided.news/x', kind: 'community' },
        ],
      },
    });
    expect(html).toContain('Sources');
    expect(html).toContain('href="https://worldofwarcraft.blizzard.com/forever"');
    expect(html).toContain('BlizzCon announcement');
    expect(html).toContain('pill-blizzard');
    expect(html).toContain('href="https://guided.news/x"');
    expect(html).toContain('pill-community');
  });

  it('deduplicates by url so a shared source is listed once', async () => {
    const c = await AstroContainer.create();
    const source = { label: 'guided.news roundup', url: 'https://guided.news/x', kind: 'community' };
    const html = await c.renderToString(SourcesList, { props: { sources: [source, { ...source }, source] } });
    expect(html.split('href="https://guided.news/x"')).toHaveLength(2);
  });
});
