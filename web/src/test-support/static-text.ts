// web/src/test-support/static-text.ts
// The words a crawler reads in a page's server render: the text inside <main>, without scripts,
// styles, templates, noscript fallbacks or inline SVG. The text-floor guard
// (src/pages/_seo-text.test.ts) counts them.

/** Elements whose contents are never read as page text. */
const HIDDEN_ELEMENTS = /<(script|style|template|noscript|svg)\b[\s\S]*?<\/\1>/g;

/** The words in `html`'s <main> (the whole document when it has none). */
export function staticWords(html: string): string[] {
  const main = /<main\b[\s\S]*<\/main>/.exec(html)?.[0] ?? html;
  return main
    .replace(HIDDEN_ELEMENTS, ' ')
    .replace(/<[^>]+>/g, ' ')
    .replace(/&[a-z0-9#]+;/gi, ' ')
    .split(/\s+/)
    .filter((word) => /[\p{L}\p{N}]/u.test(word));
}
