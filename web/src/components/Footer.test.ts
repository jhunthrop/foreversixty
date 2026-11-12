import { existsSync } from 'node:fs';
import { resolve } from 'node:path';
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import { FOOTER_COLUMNS, FOOTER_FAN_LINE, FOOTER_LINKS, footerOperatorLine } from '../lib/footer-copy';
import { LEGAL_OPERATOR, LEGAL_TRADEMARK } from '../lib/legal-notice';
import Footer from './Footer.astro';

const SRC = resolve(import.meta.dirname, '..');

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
  it('is the mock’s one row: the eight links and the fan-made line, no tagline', async () => {
    const html = await renderFooter();
    expect(FOOTER_LINKS.map((link) => link.label)).toEqual([
      'Get set up',
      'Addon',
      'Discord',
      'Premium',
      'Creators',
      'Report a bug',
      'Privacy',
      'Terms',
    ]);
    expect(html).toContain(FOOTER_FAN_LINE);
    expect(FOOTER_FAN_LINE).toBe('A fan-made player tool. Not affiliated with Blizzard.');
    expect(footerOperatorLine(2026)).toBe('© 2026 Commish LLC d/b/a Forever Sixty');
    expect(html).not.toContain('Play your class better');
    for (const column of FOOTER_COLUMNS) expect(html).not.toContain(`>${column.heading}</p>`);
  });

  it('leaves the full trademark sentence and the operator line to /about and /terms', async () => {
    const html = await renderFooter();
    expect(html).not.toContain(LEGAL_TRADEMARK);
    expect(html).not.toContain(LEGAL_OPERATOR);
  });

  it('renders every link with its href', async () => {
    const html = await renderFooter();
    for (const { label, href } of FOOTER_LINKS) {
      expect(html, label).toMatch(new RegExp(`<a[^>]*href="${href}"[^>]*>\\s*${label}\\s*<`));
    }
  });

  it('points every site link at a route that exists, and Discord and Report a bug at the invite', () => {
    for (const { label, href } of FOOTER_LINKS) {
      if (href.startsWith('https://')) expect(href, label).toMatch(/^https:\/\/discord\.gg\//);
      else expect(routeExists(href), `${label} -> ${href}`).toBe(true);
    }
  });
});
