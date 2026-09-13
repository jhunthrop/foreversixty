import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import FeedRow from './FeedRow.astro';

describe('FeedRow', () => {
  it('renders date, pill, title link, and note', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(FeedRow, {
      props: {
        date: new Date('2026-09-12T00:00:00Z'),
        kind: 'blizzard',
        title: 'Announced',
        href: '/changelog#a',
        note: 'Four zones.',
      },
    });
    expect(html).toContain('Sept 12');
    expect(html).toContain('pill-blizzard');
    expect(html).toContain('href="/changelog#a"');
    expect(html).toContain('Four zones.');
  });
});
