import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import ToolCard from './ToolCard.astro';

describe('ToolCard', () => {
  it('renders title, description and href with no linkLabel by default', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(ToolCard, {
      props: { title: 'Planner', description: 'Plan your build.', href: '/planner' },
    });
    expect(html).toContain('href="/planner"');
    expect(html).toContain('Planner');
    expect(html).toContain('Plan your build.');
  });

  it('renders the linkLabel as plain text inside the one anchor, never a second <a> (home rebuild spec §3.A.5)', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(ToolCard, {
      props: {
        title: 'Planner',
        description: 'Plan your build.',
        href: '/planner',
        linkLabel: 'Open the planner',
      },
    });
    expect(html).toContain('Open the planner');
    expect((html.match(/<a /g) ?? []).length).toBe(1);
  });

  it('omits the linkLabel line when not given', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(ToolCard, {
      props: { title: 'Get set up', description: 'One page.', href: '/setup' },
    });
    expect(html).not.toContain('Open the planner');
  });

  it('carries a caller-chosen data-testid on its own anchor when given', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(ToolCard, {
      props: { title: 'Planner', description: 'x', href: '/planner', testid: 'home-five-doors-planner' },
    });
    expect(html).toContain('data-testid="home-five-doors-planner"');
  });
});
