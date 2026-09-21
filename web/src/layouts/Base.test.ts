// The Svelte renderer has to be handed to the container explicitly, exactly as
// _planner.test.ts and _logs.test.ts do: Base.astro now mounts SessionNav client:load on
// every page, and Astro's integrations are not loaded in a unit test, so without this the
// container throws NoMatchingRenderer rather than rendering an empty shell.
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { getContainerRenderer } from '@astrojs/svelte/container-renderer';
import { loadRenderers } from 'astro:container';
import { describe, expect, it } from 'vitest';
import Base from './Base.astro';

describe('Base layout', () => {
  it('sets title, canonical, and og image from props', async () => {
    const renderers = await loadRenderers([getContainerRenderer()]);
    const container = await AstroContainer.create({ renderers });
    const html = await container.renderToString(Base, {
      props: { title: 'Dungeons', description: 'All nine.', path: '/dungeons' },
      slots: { default: '<main>body</main>' },
    });
    expect(html).toContain('<title data-og="title">Dungeons · Forever Sixty</title>');
    expect(html).toContain('href="https://foreversixty.gg/dungeons"');
    expect(html).toContain('content="https://foreversixty.gg/og/dungeons.png"');
    expect(html).toContain('<main>body</main>');
  });

  it('does not double the site name on the homepage', async () => {
    const renderers = await loadRenderers([getContainerRenderer()]);
    const container = await AstroContainer.create({ renderers });
    const html = await container.renderToString(Base, {
      props: { title: 'Forever Sixty', description: 'x', path: '/' },
    });
    expect(html).toContain('<title data-og="title">Forever Sixty</title>');
  });
});
