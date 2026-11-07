// web/src/lib/html.ts
// Text into HTML built as a string at build time (a rendered section, a table, an attribute):
// the four characters that could end a text node or a quoted attribute.

export function escapeHtml(text: string): string {
  return text.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;').replace(/"/g, '&quot;');
}
