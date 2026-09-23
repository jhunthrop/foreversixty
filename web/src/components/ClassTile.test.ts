import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import ClassTile from './ClassTile.astro';

describe('ClassTile', () => {
  it('defaults to the 84px guide tile, linking to the class guide', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(ClassTile, { props: { name: 'Warrior', slug: 'warrior' } });
    expect(html).toContain('h-[84px]');
    expect(html).toContain('href="/guides/warrior"');
  });

  it('renders a 36px row tile linking to the planner when compact', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(ClassTile, {
      props: { name: 'Warrior', slug: 'warrior', compact: true },
    });
    expect(html).toContain('h-9');
    expect(html).not.toContain('h-[84px]');
    expect(html).toContain('href="/planner?class=warrior"');
  });
});
