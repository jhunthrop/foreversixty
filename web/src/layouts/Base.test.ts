import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import Base from './Base.astro';

describe('Base layout', () => {
  it('sets title, canonical, and og image from props', async () => {
    const container = await AstroContainer.create();
    const html = await container.renderToString(Base, {
      props: { title: 'Dungeons', description: 'All nine.', path: '/dungeons' },
      slots: { default: '<main>body</main>' },
    });
    expect(html).toContain('<title>Dungeons · Forever Sixty</title>');
    expect(html).toContain('href="https://foreversixty.gg/dungeons"');
    expect(html).toContain('content="https://foreversixty.gg/og/dungeons.png"');
    expect(html).toContain('<main>body</main>');
  });

  it('does not double the site name on the homepage', async () => {
    const container = await AstroContainer.create();
    const html = await container.renderToString(Base, {
      props: { title: 'Forever Sixty', description: 'x', path: '/' },
    });
    expect(html).toContain('<title>Forever Sixty</title>');
  });
});
