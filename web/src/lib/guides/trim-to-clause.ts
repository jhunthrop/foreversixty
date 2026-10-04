// web/src/lib/guides/trim-to-clause.ts
// The class landing's spec-card description line (rebuild spec §4.B): the guide's own
// `description` frontmatter, "trimmed in this card to its own natural clause break if over
// ~90 characters" -- never a hard CSS clamp mid-word or mid-sentence (tenet 4 protects a
// sentence from being chopped arbitrarily). Breaks at the last comma/semicolon/period
// before the limit, falling back to the last whitespace, and only ever shortens -- a string
// already at or under the limit is returned unchanged.
const CLAUSE_BREAK = /[,;.]/g;

export function trimToClause(text: string, maxLength = 90): string {
  if (text.length <= maxLength) return text;
  const window = text.slice(0, maxLength);

  let lastBreak = -1;
  for (const match of window.matchAll(CLAUSE_BREAK)) {
    lastBreak = match.index + 1;
  }
  if (lastBreak > 0) return text.slice(0, lastBreak).trim();

  const lastSpace = window.lastIndexOf(' ');
  if (lastSpace > 0) return `${text.slice(0, lastSpace).trim()}…`;

  return `${window.trim()}…`;
}
