// web/src/lib/guides/first-paragraph.ts
// The class landing's own identity paragraph (rebuild spec §4.B): "the content file's own
// first paragraph ... unchanged prose", rendered above the new spec-card row, with the rest
// of the body (the specs list, the race/pairing paragraph, the patch-notes paragraph)
// rendered below it, "unchanged order and content" -- zero content changes (spec §3, §9).
// Every class landing file (`content/guides/<class>/index.md`) opens with a single
// paragraph block before its first blank line (confirmed for all nine), so splitting on
// that first blank line is exact, not a heuristic over markdown structure.

export interface SplitBody {
  first: string;
  rest: string;
}

/** `body` (already the markdown source minus frontmatter) split at its first blank line --
 *  `rest` is `''` when the whole body is one paragraph with nothing after it. */
export function splitFirstParagraph(body: string): SplitBody {
  const trimmed = body.trim();
  const blankLineMatch = /\n[ \t]*\n/.exec(trimmed);
  if (blankLineMatch === null) return { first: trimmed, rest: '' };
  return {
    first: trimmed.slice(0, blankLineMatch.index).trim(),
    rest: trimmed.slice(blankLineMatch.index + blankLineMatch[0].length).trim(),
  };
}
