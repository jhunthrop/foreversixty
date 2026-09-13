import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import Panel from './Panel.astro';

describe('Panel', () => {
  it('renders the title and a Sample pill when both are given', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(Panel, { props: { title: 'Right now', sample: true } });
    expect(html).toContain('Right now');
    expect(html).toContain('pill-sample');
  });

  it('still renders the Sample pill when there is no title', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(Panel, { props: { sample: true } });
    expect(html).toContain('pill-sample');
  });

  it('still renders the aside slot when there is no title', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(Panel, {
      slots: { aside: '<a href="/x">More</a>' },
    });
    expect(html).toContain('href="/x"');
    expect(html).toContain('More');
  });
});
