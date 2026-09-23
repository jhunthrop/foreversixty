import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import HomeProductPanel from './HomeProductPanel.astro';

describe('HomeProductPanel', () => {
  it('renders the label, sentence, slotted content and the link', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(HomeProductPanel, {
      props: {
        label: 'Planner',
        sentence: 'Every talent tree for 1.60, every race and class, shared by link.',
        href: '/planner',
        linkLabel: 'Open the planner',
      },
      slots: { default: '<span data-testid="live-element">nine tiles</span>' },
    });
    expect(html).toContain('data-testid="home-product-planner"');
    expect(html).toContain('Planner');
    expect(html).toContain('Every talent tree for 1.60, every race and class, shared by link.');
    expect(html).toContain('data-testid="live-element"');
    expect(html).toContain('href="/planner"');
    expect(html).toContain('Open the planner');
  });
});
