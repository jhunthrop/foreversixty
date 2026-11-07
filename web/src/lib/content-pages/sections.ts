// web/src/lib/content-pages/sections.ts
// A content page (legal and About, pages spec 2026-11-08 sections 4.2 and 4.3) split at its
// `h2` headings: the prose before the first heading (the `intro` region), then one section per
// heading with its id (the markdown's own, so a link to `#cookies` keeps working), title and HTML.
// The contents rail shows from 1024 when there are two sections or more, never for one.

export interface DocSection {
  id: string;
  title: string;
  html: string;
}

export interface DocParts {
  intro: string;
  sections: DocSection[];
}

/** The legal pages: Terms, Privacy and Refunds carry the legal tabs (decision L2). */
export const LEGAL_SLUGS: readonly string[] = ['terms', 'privacy', 'refunds'];

/** The fewest sections that earn a contents rail. */
export const MIN_TOC_SECTIONS = 2;

const H2 = /<h2 id="([^"]+)">([\s\S]*?)<\/h2>/g;

const plain = (html: string): string =>
  html
    .replace(/<[^>]+>/g, '')
    .replace(/&amp;/g, '&')
    .replace(/&#39;|&#x27;/g, "'")
    .replace(/&quot;/g, '"')
    .trim();

/** Splits rendered HTML at its `h2`s; a heading with no id fails (every rendered heading has one). */
export function splitDocument(html: string): DocParts {
  const heads = [...html.matchAll(H2)];
  const loose = (html.match(/<h2[\s>]/g) ?? []).length;
  if (loose !== heads.length) throw new Error('content page: every h2 needs an id for the contents rail');
  const first = heads[0]?.index ?? html.length;
  return {
    intro: html.slice(0, first).trim(),
    sections: heads.map((head, i) => {
      const start = head.index + head[0].length;
      const end = heads[i + 1]?.index ?? html.length;
      return { id: head[1]!, title: plain(head[2]!), html: html.slice(start, end).trim() };
    }),
  };
}

/** The contents rail's breakpoint for a page with this many sections (Page `tocFrom`). */
export function tocFromFor(sectionCount: number): 1024 | 0 {
  return sectionCount >= MIN_TOC_SECTIONS ? 1024 : 0;
}
