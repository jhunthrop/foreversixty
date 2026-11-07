// web/src/layouts/Content.test.ts
// A content page in the flat system (pages spec 2026-11-08 §4.2, §4.3): the PageHeader with the
// one Updated stamp, the lead only when the body opens on a heading, the legal tabs on Terms,
// Privacy and Refunds; the intro region, one ReadingSection per heading, the contents rail and
// its phone Disclosure from two sections. No confidence sentence, no Sources footer (PG4).
//
// The Svelte renderer has to be handed to the container explicitly (test-support/page-container):
// Content.astro wraps Base.astro, which mounts islands on every page.
import { createPageContainer } from '../test-support/page-container';
import { describe, expect, it } from 'vitest';
import Content from './Content.astro';

const UPDATED = new Date('2026-10-28T00:00:00Z');
const LEGAL_HTML =
  '<h2 id="what-is-collected">What is collected</h2><p>a</p><h2 id="cookies">Cookies</h2><p><code>fs_session</code></p>';

async function page(
  slug: string,
  html: string,
  description = 'What is collected and kept.',
): Promise<string> {
  const c = await createPageContainer();
  return c.renderToString(Content, {
    props: { slug, title: 'Privacy policy', description, updated: UPDATED, html },
  });
}

describe('Content layout', () => {
  it('heads a legal page with its title, one Updated stamp, the description as its lead, and the legal tabs', async () => {
    const html = await page('privacy', LEGAL_HTML);
    expect(html.match(/<h1[^>]*>Privacy policy<\/h1>/g)).toHaveLength(1);
    expect(html.match(/Updated 28 Oct/g)).toHaveLength(1);
    expect(html).toContain('What is collected and kept.');
    expect(html).toMatch(/<nav[^>]*aria-label="Legal pages"/);
    expect(html).toMatch(/href="\/privacy"[^>]*aria-current="page"|aria-current="page"[^>]*href="\/privacy"/);
    expect(html).not.toMatch(/>LEGAL<|>ABOUT</);
  });

  it('draws each heading as a reading section with its own id, the rail from 1024 and the phone contents', async () => {
    const html = await page('privacy', LEGAL_HTML);
    expect(html).toMatch(/<section id="what-is-collected"[^>]*data-region="section-what-is-collected"/);
    expect(html).toMatch(/<section id="cookies"[^>]*data-region="section-cookies"/);
    expect(html).toContain('data-toc-from="1024"');
    expect(html).toMatch(/data-toc-link="cookies"/);
    expect(html).toContain('2 sections');
    expect(html).toContain('<code>fs_session</code>');
  });

  it('leads an About page with its own first paragraph (the intro), no rail, no tabs', async () => {
    const html = await page('about', '<p>Who runs this.</p>', 'Who runs this site.');
    expect(html).toMatch(/data-region="intro"/);
    expect(html).toContain('Who runs this.');
    expect(html.slice(html.indexOf('<main'))).not.toContain('Who runs this site.');
    expect(html).toContain('data-toc-from="0"');
    expect(html).not.toContain('data-toc-link="');
    expect(html).not.toContain('aria-label="Legal pages"');
  });

  it('draws no confidence sentence and no Sources footer', async () => {
    const html = await page('privacy', LEGAL_HTML);
    for (const gone of ['Confirmed by Blizzard', 'Reported by one outlet', 'Our inference', '>Sources<']) {
      expect(html).not.toContain(gone);
    }
  });
});
