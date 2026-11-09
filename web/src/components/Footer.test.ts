import { existsSync } from 'node:fs';
import { resolve } from 'node:path';
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import {
  FOOTER_COLUMNS,
  FOOTER_LEGAL_LINKS,
  FOOTER_LINKS,
  FOOTER_TRADEMARK,
  footerOperatorLine,
} from '../lib/footer-copy';
import Footer from './Footer.astro';

const SRC = resolve(import.meta.dirname, '..');
const ALL_LINKS = [...FOOTER_LINKS, ...FOOTER_LEGAL_LINKS];

async function renderFooter(): Promise<string> {
  const container = await AstroContainer.create();
  return container.renderToString(Footer);
}

/** A site path resolves when a page file, a directory index or a content page serves it. */
function routeExists(href: string): boolean {
  const path = href.split('#')[0].replace(/^\//, '');
  return [`pages/${path}.astro`, `pages/${path}/index.astro`, `content/pages/${path}.md`].some((candidate) =>
    existsSync(resolve(SRC, candidate)),
  );
}

describe('Footer', () => {
  it('renders the trademark sentence verbatim, with no tagline', async () => {
    const html = await renderFooter();
    expect(html).toContain(FOOTER_TRADEMARK);
    expect(html).not.toContain('Play your class better');
    expect(FOOTER_TRADEMARK).toBe(
      'Not affiliated with or endorsed by Blizzard Entertainment. World of Warcraft and Warcraft are trademarks of Blizzard Entertainment, Inc.',
    );
  });

  it('names Commish LLC as the operator, with the copyright year', async () => {
    const html = await renderFooter();
    expect(html).toContain(
      `Forever Sixty is a registered trade name of Commish LLC. © ${new Date().getFullYear()} Commish LLC.`,
    );
    expect(footerOperatorLine(2026)).toBe(
      'Forever Sixty is a registered trade name of Commish LLC. © 2026 Commish LLC.',
    );
  });

  it('is the compact footer: Get set up and Premium beside the seal, no sitemap columns', async () => {
    const html = await renderFooter();
    expect(FOOTER_LINKS.map((link) => link.label)).toEqual(['Get set up', 'Premium']);
    for (const column of FOOTER_COLUMNS) expect(html).not.toContain(`>${column.heading}</p>`);
  });

  it('renders every site and legal link with its href', async () => {
    const html = await renderFooter();
    for (const { label, href } of ALL_LINKS) {
      expect(html, label).toMatch(new RegExp(`<a[^>]*href="${href}"[^>]*>\\s*${label}\\s*<`));
    }
  });

  it('points every link at a route that exists', () => {
    for (const { label, href } of ALL_LINKS) {
      expect(routeExists(href), `${label} -> ${href}`).toBe(true);
    }
  });

  it('carries the Discord invite and Report a bug, which goes to the Discord too, as links', async () => {
    const html = await renderFooter();
    expect(html).toMatch(/<a[^>]*href="https:\/\/discord\.gg\/[^"]+"[^>]*>\s*Join the Discord\s*</);
    expect(html).toMatch(/<a[^>]*href="https:\/\/discord\.gg\/[^"]+"[^>]*>\s*Report a bug\s*</);
    expect(html).not.toMatch(/github\.com/i);
  });
});
