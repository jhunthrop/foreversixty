import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import { addonCopy } from '../lib/addon/copy';
import Footer from './Footer.astro';

describe('Footer', () => {
  it('renders the Blizzard disclaimer verbatim', async () => {
    const container = await AstroContainer.create();
    const html = await container.renderToString(Footer);
    expect(html).toContain(
      'Forever Sixty is a fan-run reference. Not affiliated with or endorsed by Blizzard Entertainment. World of Warcraft and Warcraft are trademarks of Blizzard Entertainment, Inc.',
    );
  });

  it('links to About, Sources, Changelog, the addon page, Contribute', async () => {
    const container = await AstroContainer.create();
    const html = await container.renderToString(Footer);
    for (const label of ['About', 'Sources', 'Changelog', addonCopy.pageTitle, 'Contribute']) {
      expect(html).toContain(`>${label}<`);
    }
    expect(html).toContain('href="/addon"');
  });
});
