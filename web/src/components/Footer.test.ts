import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import Footer from './Footer.astro';

describe('Footer', () => {
  it('renders the Blizzard disclaimer verbatim', async () => {
    const container = await AstroContainer.create();
    const html = await container.renderToString(Footer);
    expect(html).toContain(
      'Forever Sixty is a fan-run reference. Not affiliated with or endorsed by Blizzard Entertainment. World of Warcraft and Warcraft are trademarks of Blizzard Entertainment, Inc.',
    );
  });

  it('links to About, Sources, Changelog, Contribute, Premium, Get set up', async () => {
    const container = await AstroContainer.create();
    const html = await container.renderToString(Footer);
    for (const label of ['About', 'Sources', 'Changelog', 'Contribute', 'Premium', 'Get set up']) {
      expect(html).toContain(`>${label}<`);
    }
  });

  it('links Get set up to /setup', async () => {
    const container = await AstroContainer.create();
    const html = await container.renderToString(Footer);
    const match = /<a[^>]*href="\/setup"[^>]*>/.exec(html);
    expect(match, 'no href="/setup" anchor found').toBeDefined();
    expect(match?.[0]).toContain('>');
  });
});
