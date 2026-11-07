// web/src/lib/content-pages/sections.test.ts
import { readFileSync, readdirSync } from 'node:fs';
import path from 'node:path';
import { describe, expect, it } from 'vitest';
import { renderGuideMarkdown } from '../guides/markdown';
import { splitDocument, tocFromFor } from './sections';

const PAGES = path.resolve(import.meta.dirname, '../../content/pages');

describe('a content page split at its headings', () => {
  it('keeps the prose before the first heading as the intro and each heading with its id', () => {
    const parts = splitDocument(
      '<p>Hi.</p><h2 id="one">One &amp; done</h2><p>a</p><h2 id="two">Two</h2><p>b</p>',
    );
    expect(parts).toEqual({
      intro: '<p>Hi.</p>',
      sections: [
        { id: 'one', title: 'One & done', html: '<p>a</p>' },
        { id: 'two', title: 'Two', html: '<p>b</p>' },
      ],
    });
  });

  it('fails on a heading with no id (the contents rail could not link it)', () => {
    expect(() => splitDocument('<h2>Loose</h2><p>x</p>')).toThrow(/id/);
  });

  it('shows the contents rail from 1024 for two sections or more, never for one', () => {
    expect(tocFromFor(0)).toBe(0);
    expect(tocFromFor(1)).toBe(0);
    expect(tocFromFor(2)).toBe(1024);
  });

  it('splits every committed page, the text whole', async () => {
    for (const file of readdirSync(PAGES).filter((f) => f.endsWith('.md'))) {
      const body = readFileSync(path.join(PAGES, file), 'utf8').replace(/^---\n[\s\S]*?\n---\n/, '');
      const html = await renderGuideMarkdown(body);
      const { intro, sections } = splitDocument(html);
      const joined = intro + sections.map((s) => s.html).join('');
      expect(joined.replace(/\s/g, '').length, file).toBeGreaterThan(0);
      expect(
        sections.every((s) => s.id !== '' && s.title !== ''),
        file,
      ).toBe(true);
    }
  });
});
