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

  it('links to About, Sources, Changelog, Contribute', async () => {
    const container = await AstroContainer.create();
    const html = await container.renderToString(Footer);
    for (const label of ['About', 'Sources', 'Changelog', 'Contribute']) {
      expect(html).toContain(`>${label}<`);
    }
  });

  it('no longer links to the addon page (promoted to the primary nav)', async () => {
    const container = await AstroContainer.create();
    const html = await container.renderToString(Footer);
    expect(html).not.toContain('href="/addon"');
  });
});
