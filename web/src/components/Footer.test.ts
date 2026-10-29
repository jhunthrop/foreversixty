import { existsSync } from 'node:fs';
import { resolve } from 'node:path';
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import {
  FOOTER_COLUMNS,
  FOOTER_LEGAL_LINKS,
  FOOTER_PROMISE,
  FOOTER_TRADEMARK,
  footerOperatorLine,
} from '../lib/footer-copy';
import Footer from './Footer.astro';

const SRC = resolve(import.meta.dirname, '..');
const ALL_LINKS = [...FOOTER_COLUMNS.flatMap((column) => column.links), ...FOOTER_LEGAL_LINKS];

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
  it('renders the trademark sentence verbatim and the promise as the brand line', async () => {
    const html = await renderFooter();
    expect(html).toContain(FOOTER_TRADEMARK);
    expect(html).toContain(FOOTER_PROMISE);
    expect(FOOTER_TRADEMARK).toBe(
      'Not affiliated with or endorsed by Blizzard Entertainment. World of Warcraft and Warcraft are trademarks of Blizzard Entertainment, Inc.',
    );
  });

  it('names Commish LLC as the operator, with the copyright year', async () => {
    const html = await renderFooter();
    expect(html).toContain(
      `Forever Sixty is operated by Commish LLC. © ${new Date().getFullYear()} Commish LLC.`,
    );
    expect(footerOperatorLine(2026)).toBe('Forever Sixty is operated by Commish LLC. © 2026 Commish LLC.');
  });

  it('has the four labelled columns of the spec in order', () => {
    expect(FOOTER_COLUMNS.map((column) => column.heading)).toEqual([
      'Play',
      'Logs',
      'Guilds',
      'Forever Sixty',
    ]);
  });

  it('renders every column and legal link with its href', async () => {
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

  it('carries the Discord invite and the Contribute issues link as buttons', async () => {
    const html = await renderFooter();
    expect(html).toMatch(/<a[^>]*href="https:\/\/discord\.gg\/[^"]+"[^>]*>[\s\S]*?Join the Discord/);
    expect(html).toMatch(/<a[^>]*href="[^"]+\/issues"[^>]*>[\s\S]*?Contribute/);
  });
});
